// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

/// @notice Interface for pausable & unpausable contracts
interface IPausable {
  /// @notice This function pauses the contract
  /// @dev Sets the pause flag to true
  function emergencyPause() external;

  /// @notice This function unpauses the contract
  /// @dev Sets the pause flag to false
  function emergencyUnpause() external;
}
