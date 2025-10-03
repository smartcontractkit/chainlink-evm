// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

interface IFeeWithdrawer {
  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @param feeTokens The fee tokens to withdraw.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external;
}
