// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

library Common {
  /// @notice Parameters to transfer assets
  struct AssetAmount {
    address asset; // asset address.
    uint256 amount; // Amount of assets.
  }
}
