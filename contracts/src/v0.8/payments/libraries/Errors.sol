// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

/// @notice Library for common custom errors used across multiple contracts
library Errors {
  /// @notice This error is thrown whenever a zero-address is supplied when
  /// a non-zero address is required
  error InvalidZeroAddress();
  /// @notice This error is thrown when trying to pass in an empty list as an argument
  error EmptyList();
  /// @notice This error is thrown when passing in a zero amount as a function parameter
  error InvalidZeroAmount();
  /// @notice This error is thrown when a value is not updated e.g. when trying to configure a state variable the same
  /// value as the one already configured
  error ValueNotUpdated();
}
