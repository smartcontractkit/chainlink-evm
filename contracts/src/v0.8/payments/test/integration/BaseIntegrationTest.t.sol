// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {PaymentTokenOnRamp} from "../../PaymentTokenOnRamp.sol";
import {Roles} from "../../libraries/Roles.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {MockERC20} from "../mocks/MockERC20.sol";
import {MockLinkToken} from "../mocks/MockLinkToken.sol";
import {MockWrappedNative} from "../mocks/MockWrappedNative.sol";

// @notice Base contract for integration tests. Tests the interactions between multiple contracts in a simulated
// environment.
abstract contract BaseIntegrationTest is BaseTest {
  PaymentTokenOnRamp internal s_paymentTokenOnRamp;

  MockWrappedNative internal s_mockWETH;
  MockLinkToken internal s_mockLINK;
  MockERC20 internal s_mockUSDC;

  address internal s_authority;
  uint256 internal s_authorityPk;
  address[] internal s_paymentRequestSigners;
  address internal immutable i_feeAggregator = makeAddr("feeAggregator");

  constructor() {
    // Increment block.timestamp to avoid underflows
    skip(1 weeks);

    s_mockWETH = new MockWrappedNative();
    s_mockLINK = new MockLinkToken();
    s_mockUSDC = new MockERC20("USDC", "USDC", 6);

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

    s_paymentTokenOnRamp.grantRole(Roles.PAUSER_ROLE, i_pauser);
    s_paymentTokenOnRamp.grantRole(Roles.UNPAUSER_ROLE, i_unpauser);

    // Add contracts to the list of contracts that are EmergencyWithdrawer
    s_commonContracts[CommonContracts.EMERGENCY_WITHDRAWER].push(address(s_paymentTokenOnRamp));

    vm.label(i_owner, "Owner");
    vm.label(i_unpauser, "Unpauser");
    vm.label(address(s_mockLINK), "Mock LINK");
    vm.label(address(s_mockWETH), "Mock WETH");
    vm.label(address(s_mockUSDC), "Mock USDC");
  }

  /// @notice Empty test function to ignore file in coverage report
  function test_baseUnitTest() public {}
}
