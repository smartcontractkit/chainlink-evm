// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

/// @notice Library for payments contract roles IDs to use with the OpenZeppelin AccessControl contracts.
library Roles {
  /// @notice This is the ID for the pauser role, which is given to the addresses that can pause and
  /// the contract.
  /// @dev Hash: 0x65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a
  bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
  /// @notice This is the ID for the unpauser role, which is given to the addresses that can unpause
  /// the contract.
  /// @dev Hash: 0x427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a
  bytes32 public constant UNPAUSER_ROLE = keccak256("UNPAUSER_ROLE");
  /// @notice This is the ID for the payment validator role, which is given to addresses that are able to
  /// sign payment requests off-chain
  /// @dev Hash: 0xa04d4bae570effd1b0024bf8c7251040ba88950a58517fa84a34199781f19fb0
  bytes32 public constant PAYMENT_VALIDATOR_ROLE = keccak256("PAYMENT_VALIDATOR_ROLE");
}
