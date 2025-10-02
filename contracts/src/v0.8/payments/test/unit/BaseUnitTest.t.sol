// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {PaymentTokenOnRamp} from "../../PaymentTokenOnRamp.sol";
import {Roles} from "../../libraries/Roles.sol";
import {BaseTest} from "../BaseTest.t.sol";

abstract contract BaseUnitTest is BaseTest {
  address internal immutable i_token1 = makeAddr("token1");
  address internal immutable i_token2 = makeAddr("token2");
  address internal immutable i_feeAggregator = makeAddr("feeAggregator");

  PaymentTokenOnRamp internal s_paymentTokenOnRamp;
  address internal s_authority;
  uint256 internal s_authorityPk;
  address[] internal s_paymentRequestSigners;

  constructor() {
    // Increment block.timestamp to avoid underflows
    skip(1 weeks);

    (s_authority, s_authorityPk) = makeAddrAndKey("authority");

    s_paymentRequestSigners.push(s_authority);

    s_paymentTokenOnRamp = new PaymentTokenOnRamp(
      PaymentTokenOnRamp.ConstructorParams({
        admin: i_owner,
        adminRoleTransferDelay: DEFAULT_ADMIN_TRANSFER_DELAY,
        feeAggregator: i_feeAggregator,
        chainSelector: CHAIN_SELECTOR,
        paymentRequestSigners: s_paymentRequestSigners
      })
    );

    _changePrank(i_owner);
    s_paymentTokenOnRamp.grantRole(Roles.PAYMENT_VALIDATOR_ROLE, s_authority);
    s_paymentTokenOnRamp.grantRole(Roles.PAUSER_ROLE, i_pauser);
    s_paymentTokenOnRamp.grantRole(Roles.UNPAUSER_ROLE, i_unpauser);

    // Add contracts to the list of contracts that are PausableWithAccessControl
    s_commonContracts[CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL].push(address(s_paymentTokenOnRamp));

    vm.label(address(s_paymentTokenOnRamp), "PaymentTokenOnRamp");
    vm.label(i_owner, "Owner");
    vm.label(i_unpauser, "Unpauser");
    vm.label(i_nonOwner, "Non-Owner");
    vm.label(s_authority, "Authority");
  }

  /// @notice Empty test function to ignore file in coverage report
  function test_baseUnitTest() public {}
}
