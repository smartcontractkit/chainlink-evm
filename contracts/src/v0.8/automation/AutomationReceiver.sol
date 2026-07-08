// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./ReceiverTemplate.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

/**
 * @title AutomationReceiver
 * @notice Generic bridge that executes Automation-style upkeeps delivered by a CRE workflow.
 *
 * @dev Two independent authorization layers protect this contract:
 *
 *      1. INBOUND — answers "who may deliver a report?":
 *         a) {ReceiverTemplate} enforces the CRE Forwarder address check and optional
 *            workflowId / workflowName / workflowOwner identity checks.
 *         b) {_processReport} additionally requires that the forwarder is non-zero (closing
 *            the gap left by `ReceiverTemplate.setForwarderAddress` which permits address(0))
 *            and that at least one complete workflow identity option is configured before any
 *            report is accepted: either (1) workflowId is set, or (2) both workflowOwner and
 *            workflowName are set. Neither option alone is sufficient for option 2.
 *
 *      2. OUTBOUND (this contract) — answers "what may a report make this contract do?":
 *         a closed-by-default allowlist of (target, function-selector) pairs. The inbound checks
 *         do NOT constrain the `(target, data)` a report carries, so without this layer any
 *         authorized report could call any contract/function this receiver is trusted by. The
 *         owner must explicitly allow each (target, selector) before it can be executed.
 *
 *      Migration rule of thumb: inbound authorizes the workflow; outbound authorizes the action.
 */
contract AutomationReceiver is ReceiverTemplate, Pausable {
    // Non-target gas costs reserved by the guard (gasleft() < required).
    // GAS_OVERHEAD covers only the post-check path: everything from the gasleft()
    // comparison through function return. The s_consumerGasLimit SLOAD immediately
    // preceding the check is NOT included — it is already reflected in the lower
    // gasleft() value at the check point and must not be double-counted.
    //
    // [Post-check — must complete with remaining gas at the guard point:]
    //   Pre-call ops (GAS, ADD, LT, JUMPI, stack)          ~50
    //   CALL to target (cold account access — in production  ~2,600
    //     setCallAllowed runs in an earlier transaction, so
    //     the target account has not yet been touched in the
    //     delivery transaction and the CALL pays the full
    //     EIP-2929 cold-access surcharge, not the ~100 warm rate)
    //   Post-call (success flag, JUMPI, LOG3)              ~2,200
    //     ├─ LOG3 base + 3 topics (375 + 3×375)          1,500
    //     ├─ LOG3 data  (64 B ABI-encoded empty bytes)     512
    //     └─ misc (returnData mem, JUMPI, stack)           188
    //   Misc stack / memory                                ~50
    //                                          Subtotal:  ~4,900
    //   Conservative headroom (future opcode repricing,     ~2,100
    //     estimation error)
    //                                          Total:      ~7,000
    //
    // EIP-150 (63/64 rule): a CALL can forward at most 63/64 of available gas.
    // Without the `consumerGasLimit / 63` term, the guard under-provisions whenever
    // consumerGasLimit > 63 × GAS_OVERHEAD (~441,000). _processReport therefore adds
    // consumerGasLimit / 63 to `required` dynamically so that:
    //   63/64 × available  ≥  consumerGasLimit
    //
    // Pre-guard overhead (paid BEFORE the check point, so NOT part of `required`):
    // Up to four cold SLOADs in ReceiverTemplate.onReport when identity checks are
    // active (s_forwarderAddress, s_expectedWorkflowId, s_expectedAuthor,
    // s_expectedWorkflowName: 4 × 2,100 = 8,400 gas), warm re-reads of those slots
    // in _processReport (~400 gas), abi.decode of the report (~300 gas), the cold
    // nested-mapping read of s_callAllowed (~4,200 gas on first access), a cold
    // s_consumerGasLimit SLOAD (~2,100 gas), and paused() (~2,100 gas cold) add up to
    // ~17,500 gas on the first delivery to a pair. After slots warm, pre-guard drops to
    // ~5,500 gas. When the block-number monotonicity check is enabled for a pair, budget
    // a further ~10,000 gas for its two cold SLOADs and the advancing SSTORE.
    //
    // This pre-guard gas is intentionally NOT added to `required` and needs no per-feature constant:
    // it is spent before `gasleft()` is read, so the guard already sees it (it observes the reduced
    // remaining gas) and cannot manufacture it. Adding it to `required` would double-count and cause
    // spurious InsufficientGas reverts. The only caller-side action is to size writeGasLimit to cover
    // this pre-guard overhead on top of consumerGasLimit + GAS_OVERHEAD (see README Gas Limit
    // section). Under-budgeting fails cleanly (InsufficientGas if the guard is reached, otherwise
    // OOG) — both are recorded by the Forwarder as retryable, so this is a tuning, not a safety,
    // concern.
    uint256 private constant GAS_OVERHEAD = 7000;

    /// @notice Closed-by-default allowlist of callable (target, selector) pairs.
    mapping(address target => mapping(bytes4 selector => bool allowed)) private s_callAllowed;

    /// @notice Per-(target, selector) minimum gas the consumer needs to execute. 0 = no limit.
    mapping(address target => mapping(bytes4 selector => uint256 gasLimit)) private s_consumerGasLimit;

    /// @notice Opt-in per-(target, selector) block-number monotonicity switch. Closed by default.
    mapping(address target => mapping(bytes4 selector => bool enabled)) private s_blockNumberCheckEnabled;

    /// @notice Highest report block number accepted so far for a (target, selector) pair. Doubles as
    ///         the configurable initial floor set by {setBlockNumberCheck}.
    mapping(address target => mapping(bytes4 selector => uint256 blockNumber)) private s_lastReportBlock;
    /// @notice While paused, controls how incoming reports are handled. The owner picks the mode
    ///         at {pause} time: true = reverts (reports stay unconsumed and retryable, resuming
    ///         after {unpause}); false = the report is consumed (dropped, not retried). Only
    ///         meaningful while {paused} is true.
    bool private s_retryableWhilePaused;

    /// @notice Emitted when a target call succeeds.
    event CallExecuted(address indexed target, bytes4 indexed selector, bytes returnData);
    /// @notice Emitted when an allowed target call reverts. The report is still consumed.
    event CallFailed(address indexed target, bytes4 indexed selector, bytes reason);
    /// @notice Emitted when the owner updates the outbound allowlist.
    event CallAllowedSet(address indexed target, bytes4 indexed selector, bool allowed);
    /// @notice Emitted when the owner updates the consumer gas limit for a (target, selector) pair.
    event ConsumerGasLimitSet(address indexed target, bytes4 indexed selector, uint256 previousLimit, uint256 newLimit);
    /// @notice Emitted when the owner enables/disables the block-number monotonicity check for a pair.
    ///         When enabled, `initialBlockNumber` is the floor the next report must meet or exceed.
    event BlockNumberCheckSet(address indexed target, bytes4 indexed selector, bool enabled, uint256 initialBlockNumber);
    /// @notice Emitted when a report is skipped because its block number is older than the last
    ///         accepted one for the pair. The report is consumed (not reverted), mirroring
    ///         Chainlink Automation's fire-and-forget semantics.
    event StaleReportSkipped(address indexed target, bytes4 indexed selector, uint256 reportBlockNumber, uint256 lastReportBlock);
    /// @notice Emitted when a report is dropped (consumed, not retried) because the receiver is
    ///         paused in non-retryable mode (see {pause}).
    event ReportSkippedWhilePaused();

    /// @notice Thrown when the decoded target is the zero address.
    error InvalidTargetAddress();
    /// @notice Thrown when the target address has no deployed code (EOA, mistyped address, or never-deployed contract).
    error TargetHasNoCode(address target);
    /// @notice Thrown when the report carries fewer than 4 bytes of calldata (no selector).
    error MissingSelector();
    /// @notice Thrown when (target, selector) is not on the outbound allowlist.
    error CallNotAllowed(address target, bytes4 selector);
    /// @notice Thrown when there is not enough gas to safely forward consumerGasLimit to the target.
    ///         Causes the forwarder to record the transmission as failed so it can be retried.
    error InsufficientGas(uint256 available, uint256 required);
    /// @notice Thrown when onReport is called without a complete workflow identity configuration.
    ///         The receiver requires at least one of the two valid options to be satisfied
    ///         (satisfying both is also fine):
    ///         (1) workflowId is set, or (2) both workflowOwner and workflowName are set.
    ///         Without at least one complete option the receiver cannot be bound to a specific
    ///         workflow and would accept reports from any DON-signed payload.
    error WorkflowIdentityNotConfigured();

    constructor(address _forwarder) ReceiverTemplate(_forwarder) {}

    /// @notice Emergency stop: halt all report processing until {unpause} is called. Owner-only.
    /// @dev Global equivalent of Chainlink Automation's registry-wide pause. DON pause/delete only
    ///      stops NEW reports; already-signed reports remain permissionlessly deliverable, and this
    ///      on-chain switch also rejects those in-flight reports. The owner chooses, per pause, how
    ///      reports that arrive while paused are handled via `retryable`:
    ///      - `retryable = true`: `_processReport` reverts with `EnforcedPause`. The revert bubbles
    ///        up through `onReport`, so the CRE Forwarder records the transmission as failed and it
    ///        stays retryable — pending reports resume delivery once {unpause} is called.
    ///      - `retryable = false`: `_processReport` emits `ReportSkippedWhilePaused` and returns,
    ///        consuming the report. Such reports are dropped and will NOT be redelivered after
    ///        {unpause}. Use this when a backlog of retried reports would be undesirable.
    /// @param retryable Whether reports delivered while paused should remain retryable (true) or be
    ///        dropped (false). The chosen mode applies until the next {pause} or {unpause}.
    function pause(bool retryable) external onlyOwner {
        s_retryableWhilePaused = retryable;
        _pause();
    }

    /// @notice Resume report processing after a {pause}. Owner-only.
    function unpause() external onlyOwner {
        _unpause();
    }

    /// @notice Returns the pause mode selected at the last {pause}: true if reports delivered while
    ///         paused revert (retryable), false if they are dropped. Only meaningful while paused.
    function retryableWhilePaused() external view returns (bool) {
        return s_retryableWhilePaused;
    }

    /// @notice Allow or disallow the receiver to call `selector` on `target`.
    /// @dev Closed by default. Register every (target, selector) the migrated upkeep needs,
    ///      e.g. `performUpkeep(bytes)` for custom-logic/log upkeeps, or your specific
    ///      time-based function. Owner-only.
    ///      When `allowed` is true, validates that `target` has deployed code at the time of
    ///      registration; passing an EOA, a mistyped address, or a never-deployed address reverts
    ///      with `TargetHasNoCode`. The check is skipped when `allowed` is false so that an
    ///      allowlist entry can always be revoked, even if the target has since self-destructed
    ///      and its code has become empty.
    /// @param target The contract the receiver is permitted to call.
    /// @param selector The 4-byte function selector permitted on `target`.
    /// @param allowed True to permit, false to revoke.
    function setCallAllowed(address target, bytes4 selector, bool allowed) external onlyOwner {
        if (target == address(0)) {
            revert InvalidTargetAddress();
        }
        if (allowed && target.code.length == 0) {
            revert TargetHasNoCode(target);
        }
        s_callAllowed[target][selector] = allowed;
        emit CallAllowedSet(target, selector, allowed);
    }

    /// @notice Returns whether the receiver may call `selector` on `target`.
    function isCallAllowed(address target, bytes4 selector) external view returns (bool) {
        return s_callAllowed[target][selector];
    }

    /// @notice Set the minimum gas required to execute `selector` on `target`.
    /// @dev When non-zero, `_processReport` will revert with `InsufficientGas` before calling
    ///      the target if available gas is below `gasLimit + gasLimit / 63 + GAS_OVERHEAD`.
    ///      The `gasLimit / 63` term compensates for the EIP-150 (63/64) rule: a CALL can
    ///      forward at most 63/64 of the available gas, so without this buffer a high gas limit
    ///      would cause the target to receive less than `gasLimit`. This causes the CRE
    ///      Forwarder to record the transmission as failed (retryable) rather than permanently
    ///      consuming the report. Set this to the `performGasLimit` tuned in Automation for the
    ///      specific function being migrated. Each (target, selector) pair has its own limit.
    ///      Zero (the default) disables the guard for that pair and preserves fire-and-forget.
    ///      Note: the on-chain formula only covers costs from the guard check onward. The
    ///      workflow's writeGasLimit must also budget for pre-guard overhead (~17,500 gas
    ///      on the first delivery to a pair, ~5,500 gas thereafter) on top of this limit.
    /// @param target  The contract the limit applies to. Must not be the zero address.
    /// @param selector The 4-byte function selector the limit applies to.
    /// @param gasLimit Minimum gas required by the consumer. 0 = no guard.
    function setConsumerGasLimit(address target, bytes4 selector, uint256 gasLimit) external onlyOwner {
        if (target == address(0)) revert InvalidTargetAddress();
        uint256 previous = s_consumerGasLimit[target][selector];
        s_consumerGasLimit[target][selector] = gasLimit;
        emit ConsumerGasLimitSet(target, selector, previous, gasLimit);
    }

    /// @notice Returns the configured consumer gas limit for a (target, selector) pair (0 = no guard).
    function getConsumerGasLimit(address target, bytes4 selector) external view returns (uint256) {
        return s_consumerGasLimit[target][selector];
    }

    /// @notice Enable or disable the block-number monotonicity (staleness) check for a (target, selector).
    /// @dev Closed by default. When enabled, `_processReport` rejects any report whose encoded block
    ///      number is strictly less than the last accepted one for the pair (CLA-parity: a report at
    ///      the same block is still accepted). Stale reports are consumed, not reverted, so the
    ///      forwarder does not pointlessly retry an always-stale delivery.
    ///
    ///      Under Chainlink Automation the registry rejected any report whose trigger block preceded
    ///      the last performed block. The CRE delivery path does not preserve this ordering, so this
    ///      opt-in check restores it for conditional-trigger / log upkeeps that relied on it. The
    ///      workflow always encodes a block number in the report payload, so this can be turned on
    ///      later without any workflow change.
    ///
    ///      When enabling, the initial floor is configurable:
    ///        - `initialBlockNumber == 0`: snapshot the current `block.number` as the floor. Note this
    ///          does NOT reject reports minted between deployment/config and the moment this check is
    ///          enabled if they carry a block number above the snapshot; prefer an explicit floor when
    ///          that matters.
    ///        - `initialBlockNumber != 0`: use the provided value as the floor.
    ///      The very first accepted report must have a block number >= this floor.
    ///
    ///      Disabling (enabled = false) resets the floor to 0, so out-of-order reports execute
    ///      again going forward. This is NOT retroactive: a report already rejected as stale while
    ///      the check was enabled was consumed at that moment (see `StaleReportSkipped`) and is gone
    ///      — the CRE Forwarder does not hold it for redelivery. Turning the check off cannot revive
    ///      it; only reports the workflow re-submits after disabling will be considered.
    ///      Each (target, selector) pair is configured independently. Owner-only.
    /// @param target The contract the check applies to. Must not be the zero address.
    /// @param selector The 4-byte function selector the check applies to.
    /// @param enabled True to enable the check, false to disable it (and clear the stored floor).
    /// @param initialBlockNumber Floor for the first report when enabling; 0 snapshots `block.number`.
    function setBlockNumberCheck(address target, bytes4 selector, bool enabled, uint256 initialBlockNumber)
        external
        onlyOwner
    {
        if (target == address(0)) revert InvalidTargetAddress();
        s_blockNumberCheckEnabled[target][selector] = enabled;
        uint256 floor = enabled ? (initialBlockNumber == 0 ? block.number : initialBlockNumber) : 0;
        s_lastReportBlock[target][selector] = floor;
        emit BlockNumberCheckSet(target, selector, enabled, floor);
    }

    /// @notice Returns the block-number monotonicity configuration for a (target, selector) pair.
    /// @return enabled Whether the check is active.
    /// @return lastReportBlock The highest report block number accepted so far (or the floor set when
    ///         the check was enabled, if no report has been accepted since).
    function getBlockNumberCheck(address target, bytes4 selector)
        external
        view
        returns (bool enabled, uint256 lastReportBlock)
    {
        return (s_blockNumberCheckEnabled[target][selector], s_lastReportBlock[target][selector]);
    }

    /// @notice Decodes and executes the call on the target contract.
    /// @param report ABI-encoded (address target, uint256 blockNumber, bytes data), where
    ///        `blockNumber` is the block the report was produced for (used only by the opt-in
    ///        monotonicity check) and `data` is a full function call (4-byte selector followed by its
    ///        arguments) executed as `target.call(data)`.
    /// @dev Two pre-conditions are enforced before any decoding:
    ///      1. The forwarder address must not be zero. `ReceiverTemplate.setForwarderAddress`
    ///         does not block address(0), so this guard closes that gap: if the owner ever
    ///         sets the forwarder to zero (disabling the caller check in onReport), every
    ///         subsequent report delivery is rejected here instead.
    ///      2. A complete workflow identity option must be configured. Two options are accepted:
    ///         (a) workflowId is set — binds the receiver to one specific workflow instance; or
    ///         (b) both workflowOwner and workflowName are set — binds the receiver to a named
    ///         workflow from a specific owner. Either piece of option (b) alone is insufficient:
    ///         owner alone allows any workflow from that owner; name alone is globally ambiguous.
    ///         Requiring a complete option closes the cross-receiver replay vector from audit M-02.
    ///      Authorization failures (zero target, missing selector, not-allowlisted) revert loudly —
    ///      they indicate misconfiguration or a malformed report. Execution failures (an allowed call that
    ///      reverts) are swallowed: `CallFailed` is emitted and the report is consumed, matching
    ///      Chainlink Automation's fire-and-forget semantics where the next trigger re-evaluates
    ///      eligibility.
    ///      Staleness: when the block-number monotonicity check is enabled for the pair (see
    ///      {setBlockNumberCheck}), a report whose block number is older than the last accepted one
    ///      is skipped — `StaleReportSkipped` is emitted and the report is consumed without executing
    ///      the target call or reverting.
    ///      Gas-guard: when `s_consumerGasLimit[target][selector]` is non-zero, the function
    ///      reverts with `InsufficientGas` before the target call if available gas is below
    ///      `gasLimit + gasLimit / 63 + GAS_OVERHEAD`. The `gasLimit / 63` term accounts for
    ///      the EIP-150 (63/64) rule: a CALL forwards at most 63/64 of available gas, so
    ///      without this buffer a high gas limit (above ~441,000) would cause the target to
    ///      receive less than configured. This ensures a low-gas delivery is recorded as failed
    ///      by the forwarder and can be retried, preventing griefing attacks. Each
    ///      (target, selector) pair has its own configurable limit.
    ///      Pause guard: when the contract is paused (see {pause}), this function short-circuits
    ///      before any decoding or execution, and its behavior depends on the mode the owner
    ///      selected at pause time. In retryable mode it reverts with `EnforcedPause` (the revert
    ///      propagates through `onReport`, so the forwarder records the transmission as failed and
    ///      it resumes after {unpause}). In non-retryable mode it emits `ReportSkippedWhilePaused`
    ///      and returns, consuming (dropping) the report so it is not redelivered.
    function _processReport(bytes calldata report) internal override {
        if (paused()) {
            if (s_retryableWhilePaused) {
                revert EnforcedPause();
            }
            emit ReportSkippedWhilePaused();
            return;
        }
        // Read the inherited permission fields directly (warm SLOADs). ReceiverTemplate.onReport
        // already loaded all four slots before dispatching here, so these reads cost ~100 gas
        // each — far cheaper than a self-STATICCALL to the external getters.
        if (s_forwarderAddress == address(0)) {
            revert InvalidForwarderAddress();
        }
        if (s_expectedWorkflowId == bytes32(0) &&
            (s_expectedAuthor == address(0) || s_expectedWorkflowName == bytes10(0))) {
            revert WorkflowIdentityNotConfigured();
        }

        (address target, uint256 reportBlockNumber, bytes memory data) =
            abi.decode(report, (address, uint256, bytes));

        if (target == address(0)) {
            revert InvalidTargetAddress();
        }
        if (data.length < 4) {
            revert MissingSelector();
        }

        // Read the leading 4-byte selector from the in-memory `data`. Safe: length >= 4 is
        // checked above, and a 32-byte mload assigned to bytes4 keeps the high (first) 4 bytes.
        bytes4 selector;
        assembly {
            selector := mload(add(data, 0x20))
        }
        if (!s_callAllowed[target][selector]) {
            revert CallNotAllowed(target, selector);
        }

        // Opt-in staleness protection: reject reports older than the last accepted one for this pair.
        // CLA-parity comparison (>=): a report at the same block is still accepted. The stored block
        // is updated before the gas guard, but an InsufficientGas revert rolls it back so a retry at
        // the same block stays valid.
        if (s_blockNumberCheckEnabled[target][selector]) {
            uint256 lastReportBlock = s_lastReportBlock[target][selector];
            if (reportBlockNumber < lastReportBlock) {
                emit StaleReportSkipped(target, selector, reportBlockNumber, lastReportBlock);
                return;
            }
            s_lastReportBlock[target][selector] = reportBlockNumber;
        }

        uint256 consumerGasLimit = s_consumerGasLimit[target][selector];
        bool success;
        bytes memory returnData;
        if (consumerGasLimit > 0) {
            // consumerGasLimit / 63 compensates for EIP-150: a CALL forwards at most
            // 63/64 of available gas. Without this term, limits above ~441,000
            // (63 × GAS_OVERHEAD, ~441,000) would cause the target to receive less than requested.
            uint256 required = consumerGasLimit + consumerGasLimit / 63 + GAS_OVERHEAD;
            if (gasleft() < required) {
                revert InsufficientGas(gasleft(), required);
            }
            (success, returnData) = target.call{gas: consumerGasLimit}(data);
        } else {
            (success, returnData) = target.call(data);
        }

        if (success) {
            emit CallExecuted(target, selector, returnData);
        } else {
            emit CallFailed(target, selector, returnData);
        }
    }
}