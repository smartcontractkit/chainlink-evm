// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// solhint-disable
// gas-custom-errors,interface-starts-with-i,chainlink-solidity/all-caps-constant-storage-variables,chainlink-solidity/prefix-storage-variables-with-s-underscore

import {AutomationReceiver} from "../../AutomationReceiver.sol";

import {IReceiver} from "../../IReceiver.sol";
import {ReceiverTemplate} from "../../ReceiverTemplate.sol";
import {Ownable} from "@openzeppelin/contracts@5.1.0/access/Ownable.sol";
import {Pausable} from "@openzeppelin/contracts@5.1.0/utils/Pausable.sol";

interface Vm {
  function prank(
    address msgSender
  ) external;

  function expectRevert(
    bytes4 revertData
  ) external;

  function expectRevert(
    bytes calldata revertData
  ) external;

  /// @dev Marks `target` cold, as if it had never been accessed in the current transaction.
  ///      Used to simulate production delivery, where setCallAllowed ran in an earlier
  ///      transaction and the target account is therefore cold at delivery time — unlike a
  ///      same-transaction Foundry test, which would otherwise leave it warm.
  function cool(
    address target
  ) external;
}

/// @dev Minimal Automation-style target. `performUpkeep` records the last performData and
///      can be toggled to revert, to exercise the execution-failure path.
contract MockUpkeep {
  bool public shouldRevert;
  uint256 public performCount;
  bytes public lastPerformData;

  function setShouldRevert(
    bool value
  ) external {
    shouldRevert = value;
  }

  function performUpkeep(
    bytes calldata performData
  ) external {
    if (shouldRevert) {
      revert("upkeep failed");
    }
    performCount++;
    lastPerformData = performData;
  }
}

/// @dev Burns a fixed amount of gas on every call so we can test the insufficient-gas guard.
contract MockGasHog {
  uint256 public callCount;

  function performUpkeep(
    bytes calldata
  ) external {
    callCount++;
  }
}

/// @dev Records the gas remaining at the start of performUpkeep.
contract MockGasRecorder {
  uint256 public gasOnEntry;

  function performUpkeep(
    bytes calldata
  ) external {
    gasOnEntry = gasleft();
  }
}

/// @dev Worst case for GAS_OVERHEAD: consumes every unit of forwarded gas and reverts (OOG).
contract MockGasBurner {
  function performUpkeep(
    bytes calldata
  ) external pure {
    // solhint-disable-next-line no-empty-blocks
    while (true) {}
  }
}

contract AutomationReceiverTest {
  Vm private constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

  address private constant FORWARDER = address(uint160(1));
  address private constant ATTACKER = address(uint160(3));
  bytes4 private constant PERFORM_SELECTOR = bytes4(keccak256("performUpkeep(bytes)"));

  bytes32 private constant WORKFLOW_ID = bytes32(uint256(42));
  address private constant WORKFLOW_OWNER = address(uint160(5));

  uint256 private constant GAS_OVERHEAD = 7000;

  AutomationReceiver private receiver;
  MockUpkeep private target;
  MockGasHog private gasHog;
  MockGasRecorder private gasRecorder;
  MockGasBurner private gasBurner;

  constructor() {
    receiver = new AutomationReceiver(FORWARDER);
    receiver.setExpectedWorkflowId(WORKFLOW_ID);
    receiver.setExpectedAuthor(WORKFLOW_OWNER);
    target = new MockUpkeep();
    gasHog = new MockGasHog();
    gasRecorder = new MockGasRecorder();
    gasBurner = new MockGasBurner();
  }

  // ─── helpers
  // ────────────────────────────────────────────────
  function _performCall(
    bytes memory performData
  ) private pure returns (bytes memory) {
    return abi.encodeWithSignature("performUpkeep(bytes)", performData);
  }

  function _report(address tgt, bytes memory callData) private pure returns (bytes memory) {
    return abi.encode(tgt, uint256(0), callData);
  }

  function _reportAtBlock(address tgt, uint256 blockNumber, bytes memory callData) private pure returns (bytes memory) {
    return abi.encode(tgt, blockNumber, callData);
  }

  function _metadata(bytes32 wfId, address wfOwner) private pure returns (bytes memory) {
    return abi.encodePacked(wfId, bytes10(0), wfOwner);
  }

  function _deliver(
    bytes memory report
  ) private {
    vm.prank(FORWARDER);
    receiver.onReport(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report);
  }

  function _deliverOutcome(address tgt, uint256 gasAmount, bytes memory report) private returns (uint8) {
    vm.cool(tgt);
    vm.prank(FORWARDER);
    try receiver.onReport{gas: gasAmount}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report) {
      return 0;
    } catch (bytes memory data) {
      bytes4 sel;
      if (data.length >= 4) {
        assembly {
          sel := mload(add(data, 32))
        }
      }
      return sel == AutomationReceiver.InsufficientGas.selector ? 1 : 2;
    }
  }

  // ─── inbound auth
  // ────────────────────────────────────────────
  function test_onReport_RevertWhen_InvalidSender() external {
    bytes memory report = _report(address(target), _performCall(hex"01"));

    vm.expectRevert(abi.encodeWithSelector(_invalidSenderSelector(), ATTACKER, FORWARDER));
    vm.prank(ATTACKER);
    receiver.onReport(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report);
  }

  // ─── forwarder-zero guard
  // ────────────────────────────────────
  function test_onReport_RevertWhen_ForwarderIsZero() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.setForwarderAddress(address(0));

    vm.expectRevert(ReceiverTemplate.InvalidForwarderAddress.selector);
    receiver.onReport(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), _report(address(target), _performCall(hex"01")));
  }

  // ─── workflow identity guard
  // ─────────────────────────────────
  function test_onReport_WorkflowIdAloneSuffices() external {
    AutomationReceiver freshReceiver = new AutomationReceiver(FORWARDER);
    freshReceiver.setExpectedWorkflowId(WORKFLOW_ID);
    freshReceiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    vm.prank(FORWARDER);
    freshReceiver.onReport(_metadata(WORKFLOW_ID, address(0)), _report(address(target), _performCall(hex"01")));

    _assertEq(target.performCount(), 1);
  }

  function test_onReport_OwnerAndNameSufficeWithoutWorkflowId() external {
    AutomationReceiver freshReceiver = new AutomationReceiver(FORWARDER);
    freshReceiver.setExpectedAuthor(WORKFLOW_OWNER);
    freshReceiver.setExpectedWorkflowName("my-workflow");
    freshReceiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    bytes10 wfName = freshReceiver.getExpectedWorkflowName();

    vm.prank(FORWARDER);
    freshReceiver.onReport(
      abi.encodePacked(bytes32(0), wfName, WORKFLOW_OWNER), _report(address(target), _performCall(hex"01"))
    );

    _assertEq(target.performCount(), 1);
  }

  function test_onReport_RevertWhen_OwnerSetButNameMissing() external {
    AutomationReceiver freshReceiver = new AutomationReceiver(FORWARDER);
    freshReceiver.setExpectedAuthor(WORKFLOW_OWNER);
    freshReceiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    vm.expectRevert(AutomationReceiver.WorkflowIdentityNotConfigured.selector);
    vm.prank(FORWARDER);
    freshReceiver.onReport(_metadata(bytes32(0), WORKFLOW_OWNER), _report(address(target), _performCall(hex"01")));
  }

  function test_onReport_RevertWhen_NoIdentityConfigured() external {
    AutomationReceiver freshReceiver = new AutomationReceiver(FORWARDER);
    freshReceiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    vm.expectRevert(AutomationReceiver.WorkflowIdentityNotConfigured.selector);
    vm.prank(FORWARDER);
    freshReceiver.onReport(_metadata(bytes32(0), address(0)), _report(address(target), _performCall(hex"01")));
  }

  function test_onReport_CombinedIdentityAccepted() external {
    AutomationReceiver freshReceiver = new AutomationReceiver(FORWARDER);
    freshReceiver.setExpectedWorkflowId(WORKFLOW_ID);
    freshReceiver.setExpectedAuthor(WORKFLOW_OWNER);
    freshReceiver.setExpectedWorkflowName("my-workflow");
    freshReceiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    bytes10 wfName = freshReceiver.getExpectedWorkflowName();
    vm.prank(FORWARDER);
    freshReceiver.onReport(
      abi.encodePacked(WORKFLOW_ID, wfName, WORKFLOW_OWNER), _report(address(target), _performCall(hex"01"))
    );

    _assertEq(target.performCount(), 1);
  }

  // ─── outbound allowlist
  // ─────────────────────────────────────
  function test_onReport_RevertWhen_CallNotAllowed() external {
    bytes memory report = _report(address(target), _performCall(hex"01"));

    vm.expectRevert(
      abi.encodeWithSelector(AutomationReceiver.CallNotAllowed.selector, address(target), PERFORM_SELECTOR)
    );
    _deliver(report);
  }

  function test_onReport_AllowedCallExecutes() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    bytes memory performData = hex"deadbeef";
    _deliver(_report(address(target), _performCall(performData)));

    _assertEq(target.performCount(), 1);
    _assertEq(keccak256(target.lastPerformData()), keccak256(performData));
  }

  function test_onReport_RevertWhen_CallRevoked() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, false);

    vm.expectRevert(
      abi.encodeWithSelector(AutomationReceiver.CallNotAllowed.selector, address(target), PERFORM_SELECTOR)
    );
    _deliver(_report(address(target), _performCall(hex"01")));
  }

  function test_onReport_FailingCallDoesNotRevert() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    target.setShouldRevert(true);

    _deliver(_report(address(target), _performCall(hex"01")));

    _assertEq(target.performCount(), 0);
  }

  // ─── malformed reports
  // ──────────────────────────────────────
  function test_onReport_RevertWhen_ZeroTarget() external {
    vm.expectRevert(AutomationReceiver.InvalidTargetAddress.selector);
    _deliver(_report(address(0), _performCall(hex"01")));
  }

  function test_onReport_RevertWhen_MissingSelector() external {
    vm.expectRevert(AutomationReceiver.MissingSelector.selector);
    _deliver(_report(address(target), hex"010203"));
  }

  // ─── allowlist administration
  // ───────────────────────────────
  function test_setCallAllowed_RevertWhen_NotOwner() external {
    vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, ATTACKER));
    vm.prank(ATTACKER);
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
  }

  function test_setCallAllowed_RevertWhen_ZeroTarget() external {
    vm.expectRevert(AutomationReceiver.InvalidTargetAddress.selector);
    receiver.setCallAllowed(address(0), PERFORM_SELECTOR, true);
  }

  function test_setCallAllowed_RevertWhen_CodelessTarget() external {
    address eoa = address(uint160(99));
    vm.expectRevert(abi.encodeWithSelector(AutomationReceiver.TargetHasNoCode.selector, eoa));
    receiver.setCallAllowed(eoa, PERFORM_SELECTOR, true);
  }

  function test_setCallAllowed_RevocationSkipsCodeCheck() external {
    address eoa = address(uint160(99));
    receiver.setCallAllowed(eoa, PERFORM_SELECTOR, false);
    _assertFalse(receiver.isCallAllowed(eoa, PERFORM_SELECTOR));
  }

  function test_isCallAllowed_ReflectsState() external {
    _assertFalse(receiver.isCallAllowed(address(target), PERFORM_SELECTOR));
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    _assertTrue(receiver.isCallAllowed(address(target), PERFORM_SELECTOR));
  }

  // ─── emergency pause
  // ────────────────────────────────────────
  function test_pause_RevertWhen_NotOwner() external {
    vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, ATTACKER));
    vm.prank(ATTACKER);
    receiver.pause(true);
  }

  function test_unpause_RevertWhen_NotOwner() external {
    receiver.pause(true);
    vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, ATTACKER));
    vm.prank(ATTACKER);
    receiver.unpause();
  }

  function test_onReport_RevertWhen_PausedRetryable() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.pause(true);

    vm.expectRevert(Pausable.EnforcedPause.selector);
    _deliver(_report(address(target), _performCall(hex"01")));

    _assertEq(target.performCount(), 0);
  }

  function test_onReport_PausedNonRetryableConsumesReport() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.pause(false);

    _deliver(_report(address(target), _performCall(hex"01")));

    _assertEq(target.performCount(), 0);
  }

  function test_unpause_ResumesDelivery() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    receiver.pause(true);
    vm.expectRevert(Pausable.EnforcedPause.selector);
    _deliver(_report(address(target), _performCall(hex"01")));

    receiver.unpause();
    _deliver(_report(address(target), _performCall(hex"01")));

    _assertEq(target.performCount(), 1);
  }

  function test_paused_ReflectsState() external {
    _assertFalse(receiver.paused());
    receiver.pause(true);
    _assertTrue(receiver.paused());
    receiver.unpause();
    _assertFalse(receiver.paused());
  }

  function test_retryableWhilePaused_ReflectsMode() external {
    receiver.pause(true);
    _assertTrue(receiver.retryableWhilePaused());
    receiver.unpause();

    receiver.pause(false);
    _assertFalse(receiver.retryableWhilePaused());
  }

  function test_pause_RevertWhen_AlreadyPaused() external {
    receiver.pause(true);
    vm.expectRevert(Pausable.EnforcedPause.selector);
    receiver.pause(true);
  }

  function test_unpause_RevertWhen_NotPaused() external {
    vm.expectRevert(Pausable.ExpectedPause.selector);
    receiver.unpause();
  }

  // ─── consumer gas limit administration ─────────────────────
  function test_setConsumerGasLimit_RevertWhen_NotOwner() external {
    vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, ATTACKER));
    vm.prank(ATTACKER);
    receiver.setConsumerGasLimit(address(gasHog), PERFORM_SELECTOR, 500_000);
  }

  function test_setConsumerGasLimit_RevertWhen_ZeroTarget() external {
    vm.expectRevert(AutomationReceiver.InvalidTargetAddress.selector);
    receiver.setConsumerGasLimit(address(0), PERFORM_SELECTOR, 100_000);
  }

  function test_setConsumerGasLimit_SetAndGet() external {
    _assertEq(receiver.getConsumerGasLimit(address(target), PERFORM_SELECTOR), 0);

    receiver.setConsumerGasLimit(address(target), PERFORM_SELECTOR, 300_000);
    _assertEq(receiver.getConsumerGasLimit(address(target), PERFORM_SELECTOR), 300_000);

    receiver.setConsumerGasLimit(address(target), PERFORM_SELECTOR, 0);
    _assertEq(receiver.getConsumerGasLimit(address(target), PERFORM_SELECTOR), 0);
  }

  function test_setConsumerGasLimit_IsPerPair() external {
    receiver.setConsumerGasLimit(address(gasHog), PERFORM_SELECTOR, 200_000);

    _assertEq(receiver.getConsumerGasLimit(address(gasHog), PERFORM_SELECTOR), 200_000);
    _assertEq(receiver.getConsumerGasLimit(address(target), PERFORM_SELECTOR), 0);
    _assertEq(receiver.getConsumerGasLimit(address(gasHog), bytes4(keccak256("otherFn()"))), 0);
  }

  // ─── gas guard
  // ──────────────────────────────────────────────
  function test_onReport_RevertWhen_InsufficientGas() external {
    receiver.setCallAllowed(address(gasHog), PERFORM_SELECTOR, true);
    uint256 limit = 200_000;
    receiver.setConsumerGasLimit(address(gasHog), PERFORM_SELECTOR, limit);

    bytes memory report = _report(address(gasHog), _performCall(hex""));
    bool reverted;
    vm.prank(FORWARDER);
    try receiver.onReport{gas: limit + limit / 63 + GAS_OVERHEAD - 1}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report) {
      reverted = false;
    } catch (bytes memory data) {
      bytes4 sel;
      assembly {
        sel := mload(add(data, 32))
      }
      if (sel != AutomationReceiver.InsufficientGas.selector) revert("wrong revert selector");
      reverted = true;
    }
    _assertTrue(reverted);
  }

  function test_onReport_SufficientGasWithLimitSucceeds() external {
    receiver.setCallAllowed(address(gasHog), PERFORM_SELECTOR, true);
    uint256 limit = 50_000;
    receiver.setConsumerGasLimit(address(gasHog), PERFORM_SELECTOR, limit);

    bytes memory report = _report(address(gasHog), _performCall(hex""));
    vm.prank(FORWARDER);
    receiver.onReport{gas: limit + limit / 63 + GAS_OVERHEAD + 50_000}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report);

    _assertEq(gasHog.callCount(), 1);
  }

  function test_onReport_GasLimitZeroPreservesUnboundedBehavior() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    target.setShouldRevert(true);

    _deliver(_report(address(target), _performCall(hex"01")));
    _assertEq(target.performCount(), 0);
  }

  // ─── EIP-150 (63/64 rule)
  // ────────────────────────────────────
  function test_onReport_RequiredIncludesEIP150Buffer() external {
    receiver.setCallAllowed(address(gasHog), PERFORM_SELECTOR, true);
    uint256 limit = 200_000;
    receiver.setConsumerGasLimit(address(gasHog), PERFORM_SELECTOR, limit);

    bytes memory report = _report(address(gasHog), _performCall(hex""));
    bool reverted;
    uint256 emittedRequired;
    vm.prank(FORWARDER);
    try receiver.onReport{gas: limit + limit / 63 + GAS_OVERHEAD - 1}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report) {
      reverted = false;
    } catch (bytes memory errData) {
      bytes4 sel;
      assembly {
        sel := mload(add(errData, 32))
      }
      if (sel != AutomationReceiver.InsufficientGas.selector) revert("wrong revert selector");
      assembly {
        emittedRequired := mload(add(errData, 68))
      }
      reverted = true;
    }
    _assertTrue(reverted);
    _assertEq(emittedRequired, limit + limit / 63 + GAS_OVERHEAD);
  }

  function test_onReport_EIP150TermEnsuresFullGasForwardedAtHighLimit() external {
    receiver.setCallAllowed(address(gasRecorder), PERFORM_SELECTOR, true);
    uint256 limit = 1_000_000;
    receiver.setConsumerGasLimit(address(gasRecorder), PERFORM_SELECTOR, limit);

    bytes memory report = _report(address(gasRecorder), _performCall(hex""));
    vm.prank(FORWARDER);
    receiver.onReport{gas: limit + limit / 63 + GAS_OVERHEAD + 60_000}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report);

    uint256 gasReceived = gasRecorder.gasOnEntry();
    _assertTrue(gasReceived >= limit - 500);
    _assertTrue(gasReceived <= limit);
  }

  // ─── GAS_OVERHEAD accuracy
  // ──────────────────────────────────
  function test_onReport_GasOverheadCorrectlyCoversNoOpConsumer() external {
    receiver.setCallAllowed(address(gasRecorder), PERFORM_SELECTOR, true);
    uint256 limit = 30_000;
    receiver.setConsumerGasLimit(address(gasRecorder), PERFORM_SELECTOR, limit);

    bytes memory report = _report(address(gasRecorder), _performCall(hex""));
    vm.cool(address(gasRecorder));
    vm.prank(FORWARDER);
    receiver.onReport{gas: limit + limit / 63 + GAS_OVERHEAD + 60_000}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report);

    uint256 gasReceived = gasRecorder.gasOnEntry();
    _assertTrue(gasReceived >= limit - 500);
    _assertTrue(gasReceived <= limit);
  }

  function test_onReport_GasOverheadCoversWorstCaseGasBurningConsumer() external {
    receiver.setCallAllowed(address(gasBurner), PERFORM_SELECTOR, true);
    uint256 limit = 1000;
    receiver.setConsumerGasLimit(address(gasBurner), PERFORM_SELECTOR, limit);
    bytes memory report = _report(address(gasBurner), _performCall(hex""));

    uint256 lo = 0;
    uint256 hi = 500_000;
    _assertEq(_deliverOutcome(address(gasBurner), hi, report), 0);

    while (hi - lo > 1) {
      uint256 mid = (lo + hi) / 2;
      if (_deliverOutcome(address(gasBurner), mid, report) == 0) {
        hi = mid;
      } else {
        lo = mid;
      }
    }

    _assertEq(_deliverOutcome(address(gasBurner), lo, report), 1);
  }

  // ─── block-number monotonicity
  // ───────────────────────────────
  function test_onReport_BlockNumberCheckDisabledAllowsOutOfOrder() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);

    _deliver(_reportAtBlock(address(target), 100, _performCall(hex"01")));
    _deliver(_reportAtBlock(address(target), 50, _performCall(hex"01")));

    _assertEq(target.performCount(), 2);
  }

  function test_onReport_BlockNumberCheckWithExplicitFloor() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.setBlockNumberCheck(address(target), PERFORM_SELECTOR, true, 100);

    _deliver(_reportAtBlock(address(target), 99, _performCall(hex"01")));
    _assertEq(target.performCount(), 0);

    _deliver(_reportAtBlock(address(target), 100, _performCall(hex"01")));
    _assertEq(target.performCount(), 1);

    _deliver(_reportAtBlock(address(target), 101, _performCall(hex"01")));
    _assertEq(target.performCount(), 2);

    _deliver(_reportAtBlock(address(target), 100, _performCall(hex"01")));
    _assertEq(target.performCount(), 2);
  }

  function test_onReport_BlockNumberCheckEqualBlockAccepted() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.setBlockNumberCheck(address(target), PERFORM_SELECTOR, true, 100);

    _deliver(_reportAtBlock(address(target), 100, _performCall(hex"01")));
    _deliver(_reportAtBlock(address(target), 100, _performCall(hex"01")));

    _assertEq(target.performCount(), 2);
  }

  function test_onReport_BlockNumberCheckSnapshotUsesCurrentBlock() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.setBlockNumberCheck(address(target), PERFORM_SELECTOR, true, 0);

    (bool enabled, uint256 floor) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertTrue(enabled);
    _assertEq(floor, block.number);

    if (block.number > 0) {
      _deliver(_reportAtBlock(address(target), block.number - 1, _performCall(hex"01")));
      _assertEq(target.performCount(), 0);
    }

    _deliver(_reportAtBlock(address(target), block.number, _performCall(hex"01")));
    _assertEq(target.performCount(), 1);
  }

  function test_getBlockNumberCheck_ReflectsState() external {
    (bool enabled, uint256 last) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertFalse(enabled);
    _assertEq(last, 0);

    receiver.setBlockNumberCheck(address(target), PERFORM_SELECTOR, true, 500);
    (enabled, last) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertTrue(enabled);
    _assertEq(last, 500);

    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    _deliver(_reportAtBlock(address(target), 600, _performCall(hex"01")));
    (, last) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertEq(last, 600);

    receiver.setBlockNumberCheck(address(target), PERFORM_SELECTOR, false, 0);
    (enabled, last) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertFalse(enabled);
    _assertEq(last, 0);
  }

  function test_setBlockNumberCheck_IsPerPair() external {
    receiver.setBlockNumberCheck(address(gasHog), PERFORM_SELECTOR, true, 300);

    (bool enabledHog, uint256 lastHog) = receiver.getBlockNumberCheck(address(gasHog), PERFORM_SELECTOR);
    _assertTrue(enabledHog);
    _assertEq(lastHog, 300);

    (bool enabledTarget, uint256 lastTarget) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertFalse(enabledTarget);
    _assertEq(lastTarget, 0);
  }

  function test_setBlockNumberCheck_RevertWhen_NotOwner() external {
    vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, ATTACKER));
    vm.prank(ATTACKER);
    receiver.setBlockNumberCheck(address(target), PERFORM_SELECTOR, true, 0);
  }

  function test_setBlockNumberCheck_RevertWhen_ZeroTarget() external {
    vm.expectRevert(AutomationReceiver.InvalidTargetAddress.selector);
    receiver.setBlockNumberCheck(address(0), PERFORM_SELECTOR, true, 0);
  }

  function test_onReport_InsufficientGasDoesNotAdvanceBlockNumber() external {
    receiver.setCallAllowed(address(target), PERFORM_SELECTOR, true);
    receiver.setBlockNumberCheck(address(target), PERFORM_SELECTOR, true, 100);
    uint256 limit = 200_000;
    receiver.setConsumerGasLimit(address(target), PERFORM_SELECTOR, limit);

    bytes memory report = _reportAtBlock(address(target), 100, _performCall(hex"01"));

    bool reverted;
    vm.prank(FORWARDER);
    try receiver.onReport{gas: limit + limit / 63 + GAS_OVERHEAD - 1}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report) {
      reverted = false;
    } catch (bytes memory data) {
      bytes4 sel;
      assembly {
        sel := mload(add(data, 32))
      }
      if (sel != AutomationReceiver.InsufficientGas.selector) revert("wrong revert selector");
      reverted = true;
    }
    _assertTrue(reverted);

    (, uint256 lastBlock) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertEq(lastBlock, 100);
    _assertEq(target.performCount(), 0);

    vm.prank(FORWARDER);
    receiver.onReport{gas: limit + limit / 63 + GAS_OVERHEAD + 100_000}(_metadata(WORKFLOW_ID, WORKFLOW_OWNER), report);
    _assertEq(target.performCount(), 1);
    (, lastBlock) = receiver.getBlockNumberCheck(address(target), PERFORM_SELECTOR);
    _assertEq(lastBlock, 100);
  }

  // ─── supportsInterface
  // ───────────────────────────────────────
  function test_supportsInterface_ERC165() external view {
    _assertTrue(receiver.supportsInterface(0x01ffc9a7));
  }

  function test_supportsInterface_IReceiver() external view {
    _assertTrue(receiver.supportsInterface(type(IReceiver).interfaceId));
  }

  function test_supportsInterface_ReturnsFalseForUnknown() external view {
    _assertFalse(receiver.supportsInterface(0xdeadbeef));
  }

  // ─── getters
  // ────────────────────────────────────────────────
  function test_getters_ReflectState() external view {
    _assertEq(receiver.getForwarderAddress(), FORWARDER);
    _assertEq(receiver.getExpectedAuthor(), WORKFLOW_OWNER);
    _assertEq(receiver.getExpectedWorkflowId(), WORKFLOW_ID);
  }

  // ─── tiny assertion helpers (no forge-std dependency) ───────
  function _invalidSenderSelector() private pure returns (bytes4) {
    return bytes4(keccak256("InvalidSender(address,address)"));
  }

  function _assertEq(uint256 actual, uint256 expected) private pure {
    if (actual != expected) revert("uint mismatch");
  }

  function _assertEq(bytes32 actual, bytes32 expected) private pure {
    if (actual != expected) revert("bytes32 mismatch");
  }

  function _assertEq(address actual, address expected) private pure {
    if (actual != expected) revert("address mismatch");
  }

  function _assertTrue(
    bool value
  ) private pure {
    if (!value) revert("expected true");
  }

  function _assertFalse(
    bool value
  ) private pure {
    if (value) revert("expected false");
  }
}
