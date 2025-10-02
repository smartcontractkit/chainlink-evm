// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {PaymentTokenOnRamp} from "../../../../PaymentTokenOnRamp.sol";
import {Errors} from "../../../../libraries/Errors.sol";
import {BaseIntegrationTest} from "../../BaseIntegrationTest.t.sol";

import {Pausable} from "@openzeppelin/contracts@5.0.2/utils/Pausable.sol";

contract PaymentTokenOnRamp_withdrawFeeTokensTest is BaseIntegrationTest {
  address[] private s_feeTokens;

  function setUp() external {
    s_feeTokens.push(address(s_mockLINK));
    s_feeTokens.push(address(s_mockUSDC));

    deal(address(s_mockLINK), address(s_paymentTokenOnRamp), 1 ether);
    deal(address(s_mockUSDC), address(s_paymentTokenOnRamp), 10e6);
  }

  function test_RevertWhen_ContractIsPaused() external givenContractIsPaused(address(s_paymentTokenOnRamp)) {
    vm.expectRevert(Pausable.EnforcedPause.selector);

    s_paymentTokenOnRamp.withdrawFeeTokens(new address[](0));
  }

  function test_withdrawFeeTokens_RevertWhen_FeeTokensListIsEmpty() external {
    vm.expectRevert(Errors.EmptyList.selector);

    s_paymentTokenOnRamp.withdrawFeeTokens(new address[](0));
  }

  function test_withdrawFeeTokens_ZeroAmount() external {
    s_feeTokens.pop();
    s_feeTokens[0] = address(s_mockWETH);

    vm.recordLogs();

    s_paymentTokenOnRamp.withdrawFeeTokens(s_feeTokens);

    assertEq(vm.getRecordedLogs().length, 0);
    assertEq(s_mockWETH.balanceOf(address(s_paymentTokenOnRamp)), 0);
    assertEq(s_mockWETH.balanceOf(i_feeAggregator), 0);
  }

  function test_withdrawFeeTokens_SingleToken() external {
    s_feeTokens.pop();

    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.FeeTokenWithdrawn(i_feeAggregator, address(s_mockLINK), 1 ether);

    s_paymentTokenOnRamp.withdrawFeeTokens(s_feeTokens);

    assertEq(s_mockLINK.balanceOf(address(s_paymentTokenOnRamp)), 0);
    assertEq(s_mockLINK.balanceOf(i_feeAggregator), 1 ether);
  }

  function test_withdrawFeeTokens_MultipleTokens() external {
    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.FeeTokenWithdrawn(i_feeAggregator, address(s_mockLINK), 1 ether);
    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.FeeTokenWithdrawn(i_feeAggregator, address(s_mockUSDC), 10e6);

    s_paymentTokenOnRamp.withdrawFeeTokens(s_feeTokens);

    assertEq(s_mockLINK.balanceOf(address(s_paymentTokenOnRamp)), 0);
    assertEq(s_mockLINK.balanceOf(i_feeAggregator), 1 ether);
    assertEq(s_mockUSDC.balanceOf(address(s_paymentTokenOnRamp)), 0);
    assertEq(s_mockUSDC.balanceOf(i_feeAggregator), 10e6);
  }
}
