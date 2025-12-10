// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IFeeManager} from "./interfaces/IFeeManager.sol";
import {IVerifierFeeManager} from "./interfaces/IVerifierFeeManager.sol";
import {Common} from "../libraries/Common.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

/**
 * @title NoOpFeeManager
 * @notice A no-op implementation of IFeeManager that does not collect fees.
 * @dev All functions return successfully without performing any fee collection or state changes.
 *      Any ETH sent to payable functions is refunded to the subscriber.
 */
contract NoOpFeeManager is IFeeManager {
  /// @notice Error thrown when ETH refund fails
  error RefundFailed();

  /// @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) external pure override returns (bool) {
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
    // Return zero fee, zero reward, zero discount
    return (fee, reward, appliedDiscount);
  }

  /// @inheritdoc IFeeManager
  function setNativeSurcharge(uint64) external override {
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
  function payLinkDeficit(bytes32) external override {
    // No-op
  }

  /// @inheritdoc IFeeManager
  function updateSubscriberGlobalDiscount(address, address, uint64) external override {
    // No-op
  }

  /**
   * @notice Refunds any ETH sent to the contract
   * @param recipient The address to refund ETH to
   */
  function _refund(address recipient) internal {
    if (msg.value > 0) {
      (bool success, ) = payable(recipient).call{value: msg.value}("");
      if (!success) revert RefundFailed();
    }
  }
}



