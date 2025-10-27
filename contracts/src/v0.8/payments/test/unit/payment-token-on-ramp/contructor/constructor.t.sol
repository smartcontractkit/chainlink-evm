// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {IFeeWithdrawer} from "../../../../interfaces/IFeeWithdrawer.sol";

import {PaymentTokenOnRamp} from "../../../../PaymentTokenOnRamp.sol";
import {Errors} from "../../../../libraries/Errors.sol";
import {Roles} from "../../../../libraries/Roles.sol";
import {BaseUnitTest} from "../../BaseUnitTest.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.0.2/access/IAccessControl.sol";

contract PaymentTokenOnRamp_constructor is BaseUnitTest {
  function test_constructor_RevertWhen_ChainSelectorEqZero() external {
    vm.expectRevert(PaymentTokenOnRamp.InvalidChainSelector.selector);

    new PaymentTokenOnRamp(
      PaymentTokenOnRamp.ConstructorParams({
        admin: i_owner,
        adminRoleTransferDelay: DEFAULT_ADMIN_TRANSFER_DELAY,
        feeAggregator: i_feeAggregator,
        chainSelector: 0,
        paymentRequestSigners: s_paymentRequestSigners
      })
    );
  }

  function test_constructor_RevertWhen_PaymentRequestSignerEqAddressZero() external {
    address[] memory paymentRequestSigners = new address[](1);
    paymentRequestSigners[0] = address(0);

    vm.expectRevert(Errors.InvalidZeroAddress.selector);

    new PaymentTokenOnRamp(
      PaymentTokenOnRamp.ConstructorParams({
        admin: i_owner,
        adminRoleTransferDelay: DEFAULT_ADMIN_TRANSFER_DELAY,
        feeAggregator: i_feeAggregator,
        chainSelector: CHAIN_SELECTOR,
        paymentRequestSigners: paymentRequestSigners
      })
    );
  }

  function test_constructor_RevertWhen_FeeAggregatorEqAddressZero() external {
    vm.expectRevert(Errors.InvalidZeroAddress.selector);

    new PaymentTokenOnRamp(
      PaymentTokenOnRamp.ConstructorParams({
        admin: i_owner,
        adminRoleTransferDelay: DEFAULT_ADMIN_TRANSFER_DELAY,
        feeAggregator: address(0),
        chainSelector: CHAIN_SELECTOR,
        paymentRequestSigners: s_paymentRequestSigners
      })
    );
  }

  function test_constructor() external {
    vm.expectEmit();
    emit PaymentTokenOnRamp.ChainSelectorSet(CHAIN_SELECTOR);
    vm.expectEmit();
    emit IAccessControl.RoleGranted(Roles.PAYMENT_VALIDATOR_ROLE, s_authority, i_owner);
    vm.expectEmit();
    emit PaymentTokenOnRamp.FeeAggregatorSet(i_feeAggregator);

    PaymentTokenOnRamp paymentTokenOnRamp = new PaymentTokenOnRamp(
      PaymentTokenOnRamp.ConstructorParams({
        admin: i_owner,
        adminRoleTransferDelay: DEFAULT_ADMIN_TRANSFER_DELAY,
        feeAggregator: i_feeAggregator,
        chainSelector: CHAIN_SELECTOR,
        paymentRequestSigners: s_paymentRequestSigners
      })
    );

    assertTrue(paymentTokenOnRamp.hasRole(Roles.PAYMENT_VALIDATOR_ROLE, s_authority));
    assertEq(paymentTokenOnRamp.getFeeAggregator(), i_feeAggregator);
    assertEq(paymentTokenOnRamp.typeAndVersion(), "PaymentTokenOnRamp 1.0.0");
  }

  function test_supportsInterface() external view {
    assertTrue(s_paymentTokenOnRamp.supportsInterface(type(IFeeWithdrawer).interfaceId));
  }
}
