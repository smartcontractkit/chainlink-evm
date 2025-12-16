// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {Common} from "../libraries/Common.sol";
import {IFeeManager} from "./interfaces/IFeeManager.sol";
import {IVerifierFeeManager} from "./interfaces/IVerifierFeeManager.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

/**
 * @title NoOpFeeManager
 * @notice A no-op implementation of IFeeManager that does not collect fees.
 * @dev All functions return successfully without performing any fee collection or state changes.
 *      Any ETH sent to payable functions is refunded to the subscriber.
 */
contract NoOpFeeManager is IFeeManager, ITypeAndVersion {
  /// @notice Error thrown when ETH refund fails
  error RefundFailed();

  /// @notice The scalar representing 100% discount (1e18 = 100%)
  uint64 private constant PERCENTAGE_SCALAR = 1e18;

  /// @inheritdoc ITypeAndVersion
  function typeAndVersion() external pure override returns (string memory) {
    return "NoOpFeeManager 0.5.0";
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) external pure override returns (bool) {
    return interfaceId == this.processFee.selector || interfaceId == this.processFeeBulk.selector;
  }

  /// @inheritdoc IVerifierFeeManager
  function processFee(bytes calldata, bytes calldata, address subscriber) external payable override {
    // Refund any ETH sent
    _refund(subscriber);
  }

  /// @inheritdoc IVerifierFeeManager
  function processFeeBulk(bytes[] calldata, bytes calldata, address subscriber) external payable override {
    // Refund any ETH sent
    _refund(subscriber);
  }

  /// @inheritdoc IVerifierFeeManager
  function setFeeRecipients(bytes32, Common.AddressAndWeight[] calldata) external override {
    // No-op
  }

  /// @inheritdoc IFeeManager
  function getFeeAndReward(
    address,
    bytes memory,
    address
  ) external pure override returns (Common.Asset memory fee, Common.Asset memory reward, uint256 appliedDiscount) {
    // Return zero fee, zero reward, 100% discount (1e18) to indicate no fees are charged
    return (fee, reward, PERCENTAGE_SCALAR);
  }

  /// @inheritdoc IFeeManager
  function setNativeSurcharge(
    uint64
  ) external override {
    // No-op
  }

  /// @inheritdoc IFeeManager
  function updateSubscriberDiscount(address, bytes32, address, uint64) external override {
    // No-op
  }

  /// @inheritdoc IFeeManager
  function withdraw(address, address, uint192) external override {
    // No-op
  }

  /// @inheritdoc IFeeManager
  function linkAvailableForPayment() external pure override returns (uint256) {
    return 0;
  }

  /// @inheritdoc IFeeManager
  function payLinkDeficit(
    bytes32
  ) external override {
    // No-op
  }

  /// @inheritdoc IFeeManager
  function updateSubscriberGlobalDiscount(address, address, uint64) external override {
    // No-op
  }

  /**
   * @notice Returns 100% discount for any subscriber/feedId/token combination
   * @dev Replicates public mapping getter signature from FeeManager for backwards compatibility
   */
  // solhint-disable-next-line func-name-mixedcase
  function s_subscriberDiscounts(address, bytes32, address) external pure returns (uint256) {
    return PERCENTAGE_SCALAR;
  }

  /**
   * @notice Returns 100% discount for any subscriber/token combination
   * @dev Replicates public mapping getter signature from FeeManager for backwards compatibility
   */
  // solhint-disable-next-line func-name-mixedcase
  function s_globalDiscounts(address, address) external pure returns (uint256) {
    return PERCENTAGE_SCALAR;
  }

  /**
   * @notice Refunds any ETH sent to the contract
   * @param recipient The address to refund ETH to
   */
  function _refund(
    address recipient
  ) internal {
    if (msg.value > 0) {
      (bool success,) = payable(recipient).call{value: msg.value}("");
      if (!success) revert RefundFailed();
    }
  }
}
