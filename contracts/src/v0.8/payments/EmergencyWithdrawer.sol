// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {PausableWithAccessControl} from "./PausableWithAccessControl.sol";
import {Common} from "./libraries/Common.sol";
import {Errors} from "./libraries/Errors.sol";

import {IERC20} from "@openzeppelin/contracts@5.0.2/interfaces/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

/// @notice Base contract that adds ERC20 emergencyWithdraw functionality.
abstract contract EmergencyWithdrawer is PausableWithAccessControl {
  using SafeERC20 for IERC20;

  /// @notice This event is emitted when an asset is withdrawn from the contract by the admin during
  /// an emergency
  /// @param to The address of the admin
  /// @param asset The address of the asset that was withdrawn
  /// @param amount The amount of assets that was withdrawn
  event AssetEmergencyWithdrawn(address indexed to, address indexed asset, uint256 amount);

  /// @notice This error is thrown when a native token transfer fails
  /// @param to The address of the recipient
  /// @param amount The amount of native token transferred - address(0) is used for native token
  /// @param data The bubbled up revert data
  error FailedNativeTokenTransfer(address to, uint256 amount, bytes data);

  constructor(uint48 adminRoleTransferDelay, address admin) PausableWithAccessControl(adminRoleTransferDelay, admin) {}

  /// @notice Withdraws assets from the contract to the specified address
  /// @dev precondition - The contract must be paused
  /// @dev precondition - The caller must have the DEFAULT_ADMIN_ROLE
  /// @dev precondition - The assetAmounts list must not be empty
  /// @param to The address to transfer the assets to
  /// @param assetAmounts The list of assets and amounts to transfer
  function emergencyWithdraw(
    address to,
    Common.AssetAmount[] calldata assetAmounts
  ) external whenPaused onlyRole(DEFAULT_ADMIN_ROLE) {
    if (assetAmounts.length == 0) {
      revert Errors.EmptyList();
    }

    for (uint256 i; i < assetAmounts.length; ++i) {
      address asset = assetAmounts[i].asset;
      uint256 amount = assetAmounts[i].amount;
      _transferAsset(to, asset, amount);
      emit AssetEmergencyWithdrawn(to, asset, amount);
    }
  }

  /// @notice Withdraws native token from the contract to the specified address
  /// @dev precondition The contract must be paused
  /// @dev precondition The caller must have the DEFAULT_ADMIN_ROLE
  /// @param to The address to transfer the native tokens to
  /// @param amount The amount of native token to transfer
  function emergencyWithdrawNative(address payable to, uint256 amount) external whenPaused onlyRole(DEFAULT_ADMIN_ROLE) {
    _transferNative(to, amount);
    emit AssetEmergencyWithdrawn(to, address(0), amount);
  }

  /// @dev Helper function to withdraw native tokens and perform sanity checks
  /// @dev precondition The recipient must not be the zero address
  /// @dev precondition The amount must be greater than zero
  /// @param to The address to transfer the native tokens to
  /// @param amount The amount of native tokens to transfer
  function _transferNative(address payable to, uint256 amount) internal {
    if (to == address(0)) {
      revert Errors.InvalidZeroAddress();
    }
    if (amount == 0) {
      revert Errors.InvalidZeroAmount();
    }

    (bool success, bytes memory data) = to.call{value: amount}("");

    if (!success) {
      revert FailedNativeTokenTransfer(to, amount, data);
    }
  }

  /// @dev Helper function to transfer a list of assets
  /// @dev precondition The transferred assets must not be the zero address
  /// @dev precondition The amounts must be greater than zero
  /// @param to The address to transfer the asset to
  /// @param asset The asset to transfer
  /// @param amount The amount of asset to transfer
  function _transferAsset(address to, address asset, uint256 amount) internal {
    if (to == address(0) || asset == address(0)) {
      revert Errors.InvalidZeroAddress();
    }
    if (amount == 0) {
      revert Errors.InvalidZeroAmount();
    }

    IERC20(asset).safeTransfer(to, amount);
  }
}
