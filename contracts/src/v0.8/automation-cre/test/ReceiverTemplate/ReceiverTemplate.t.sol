// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// solhint-disable
// gas-custom-errors,interface-starts-with-i,chainlink-solidity/all-caps-constant-storage-variables,chainlink-solidity/prefix-storage-variables-with-s-underscore

import {ReceiverTemplate} from "../../ReceiverTemplate.sol";

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
}

contract MockReceiver is ReceiverTemplate {
  uint256 public reportCount;
  bytes32 public lastReportHash;

  constructor(
    address forwarder
  ) ReceiverTemplate(forwarder) {}

  function _processReport(
    bytes calldata report
  ) internal override {
    reportCount++;
    lastReportHash = keccak256(report);
  }
}

contract ReceiverTemplateTest {
  Vm private constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

  address private constant FORWARDER = address(uint160(1));
  address private constant AUTHOR = address(uint160(2));
  address private constant ATTACKER = address(uint160(3));
  bytes32 private constant WORKFLOW_ID = bytes32(uint256(0x1234));

  function test_constructor_RevertWhen_ZeroForwarder() external {
    vm.expectRevert(ReceiverTemplate.InvalidForwarderAddress.selector);
    new MockReceiver(address(0));
  }

  function test_onReport_RevertWhen_InvalidSender() external {
    MockReceiver receiver = new MockReceiver(FORWARDER);

    vm.expectRevert(abi.encodeWithSelector(ReceiverTemplate.InvalidSender.selector, ATTACKER, FORWARDER));
    vm.prank(ATTACKER);
    receiver.onReport("", "report");
  }

  /// @dev setForwarderAddress(address(0)) is permitted at the ReceiverTemplate level —
  ///      it emits a SecurityWarning but does not revert. Concrete contracts such as
  ///      AutomationReceiver close this gap inside _processReport.
  function test_setForwarderAddress_ZeroSucceedsAtTemplateLevel() external {
    MockReceiver receiver = new MockReceiver(FORWARDER);
    receiver.setForwarderAddress(address(0));
  }

  /// @dev setExpectedWorkflowName succeeds even without an author set — the dependency is
  ///      enforced lazily inside onReport (WorkflowNameRequiresAuthorValidation).
  function test_setExpectedWorkflowName_CanBeSetWithoutAuthor() external {
    MockReceiver receiver = new MockReceiver(FORWARDER);
    receiver.setExpectedWorkflowName("game-resolution");
  }

  /// @dev Clearing the author while a workflow name is set does not revert on the setter itself.
  function test_setExpectedAuthor_ClearingAuthorWhileNameSetDoesNotRevert() external {
    MockReceiver receiver = new MockReceiver(FORWARDER);
    receiver.setExpectedAuthor(AUTHOR);
    receiver.setExpectedWorkflowName("game-resolution");
    receiver.setExpectedAuthor(address(0));
  }

  function test_onReport_AcceptsWhenMetadataMatches() external {
    MockReceiver receiver = new MockReceiver(FORWARDER);
    receiver.setExpectedAuthor(AUTHOR);
    receiver.setExpectedWorkflowId(WORKFLOW_ID);
    receiver.setExpectedWorkflowName("game-resolution");

    bytes memory metadata = abi.encodePacked(WORKFLOW_ID, _workflowName("game-resolution"), AUTHOR);
    bytes memory report = abi.encode("verified report");

    vm.prank(FORWARDER);
    receiver.onReport(metadata, report);

    _assertEq(receiver.reportCount(), 1);
    _assertEq(receiver.lastReportHash(), keccak256(report));
  }

  function test_onReport_RevertWhen_WrongWorkflowOwner() external {
    MockReceiver receiver = new MockReceiver(FORWARDER);
    receiver.setExpectedAuthor(AUTHOR);

    bytes memory metadata = abi.encodePacked(WORKFLOW_ID, _workflowName("game-resolution"), ATTACKER);

    vm.expectRevert(abi.encodeWithSelector(ReceiverTemplate.InvalidAuthor.selector, ATTACKER, AUTHOR));
    vm.prank(FORWARDER);
    receiver.onReport(metadata, "report");
  }

  function _workflowName(
    string memory name
  ) private pure returns (bytes10) {
    bytes32 hash = sha256(bytes(name));
    bytes16 hexChars = "0123456789abcdef";
    bytes memory hexString = new bytes(64);

    for (uint256 i = 0; i < 32; i++) {
      hexString[i * 2] = hexChars[uint8(hash[i] >> 4)];
      hexString[i * 2 + 1] = hexChars[uint8(hash[i] & 0x0f)];
    }

    bytes memory first10 = new bytes(10);
    for (uint256 i = 0; i < 10; i++) {
      first10[i] = hexString[i];
    }

    return bytes10(first10);
  }

  function _assertEq(uint256 actual, uint256 expected) private pure {
    if (actual != expected) revert("uint mismatch");
  }

  function _assertEq(bytes32 actual, bytes32 expected) private pure {
    if (actual != expected) revert("bytes32 mismatch");
  }
}
