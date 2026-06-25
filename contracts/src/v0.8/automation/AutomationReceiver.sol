// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ReceiverTemplate} from "./ReceiverTemplate.sol";

/**
 * @title AutomationReceiver
 * @notice Generic bridge that executes Automation-style upkeeps delivered by a CRE workflow.
 *
 * @dev Two independent authorization layers protect this contract:
 *
 *      1. INBOUND (inherited from {ReceiverTemplate}) — answers "who may deliver a report?":
 *         the CRE Forwarder address plus the optional workflowId / workflowName / workflowOwner
 *         identity checks.
 *
 *      2. OUTBOUND (this contract) — answers "what may a report make this contract do?":
 *         a closed-by-default allowlist of (target, function-selector) pairs. The owner must
 *         explicitly allow each (target, selector) before it can be executed.
 *
 *      Migration rule of thumb: inbound authorizes the workflow; outbound authorizes the action.
 */
contract AutomationReceiver is ReceiverTemplate {
  mapping(address target => mapping(bytes4 selector => bool allowed)) private s_callAllowed;

  event CallExecuted(address indexed target, bytes4 indexed selector, bytes returnData);
  event CallFailed(address indexed target, bytes4 indexed selector, bytes reason);
  event CallAllowedSet(address indexed target, bytes4 indexed selector, bool allowed);

  error InvalidTargetAddress();
  error MissingSelector();
  error CallNotAllowed(address target, bytes4 selector);

  constructor(address _forwarder) ReceiverTemplate(_forwarder) {}

  /// @notice Allow or disallow the receiver to call `selector` on `target`. Owner-only.
  function setCallAllowed(address target, bytes4 selector, bool allowed) external onlyOwner {
    if (target == address(0)) revert InvalidTargetAddress();
    s_callAllowed[target][selector] = allowed;
    emit CallAllowedSet(target, selector, allowed);
  }

  /// @notice Returns whether the receiver may call `selector` on `target`.
  function isCallAllowed(address target, bytes4 selector) external view returns (bool) {
    return s_callAllowed[target][selector];
  }

  /// @notice Decodes and executes the call encoded in the CRE report.
  /// @param report ABI-encoded (address target, bytes data).
  function _processReport(bytes calldata report) internal override {
    (address target, bytes memory data) = abi.decode(report, (address, bytes));

    if (target == address(0)) revert InvalidTargetAddress();
    if (data.length < 4) revert MissingSelector();

    bytes4 selector;
    assembly {
      selector := mload(add(data, 0x20))
    }
    if (!s_callAllowed[target][selector]) revert CallNotAllowed(target, selector);

    // solhint-disable-next-line avoid-low-level-calls
    (bool success, bytes memory returnData) = target.call(data);
    if (success) {
      emit CallExecuted(target, selector, returnData);
    } else {
      emit CallFailed(target, selector, returnData);
    }
  }
}
