// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {PaymentTokenOnRamp} from "../../../../PaymentTokenOnRamp.sol";
import {Errors} from "../../../../libraries/Errors.sol";
import {BaseUnitTest} from "../../BaseUnitTest.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.0.2/access/IAccessControl.sol";

contract PaymentTokenOnRamp_setFeeAggregatorUnitTest is BaseUnitTest {
  address private immutable i_newFeeAggregatorReciever = makeAddr("newFeeAggregatorReciever");

  function test_setFeeAggregator_RevertWhen_CallerDoesNotHaveDEFAULT_ADMIN_ROLE() public whenCallerIsNotAdmin {
    vm.expectRevert(
      abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, i_nonOwner, DEFAULT_ADMIN_ROLE)
    );
    s_paymentTokenOnRamp.setFeeAggregator(i_newFeeAggregatorReciever);
  }

  function test_setFeeAggregator_RevertWhen_FeeAggregatorAddressZero() public {
    vm.expectRevert(Errors.InvalidZeroAddress.selector);
    s_paymentTokenOnRamp.setFeeAggregator(address(0));
  }

  function test_setFeeAggregator_RevertWhen_FeeAggregatorAddressNotUpdated() public {
    vm.expectRevert(Errors.ValueNotUpdated.selector);
    s_paymentTokenOnRamp.setFeeAggregator(i_feeAggregator);
  }

  function test_setFeeAggregator_UpdatesFeeAggregator() external {
    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.FeeAggregatorSet(i_newFeeAggregatorReciever);
    s_paymentTokenOnRamp.setFeeAggregator(i_newFeeAggregatorReciever);
    assertEq(address(s_paymentTokenOnRamp.getFeeAggregator()), i_newFeeAggregatorReciever);
  }
}
