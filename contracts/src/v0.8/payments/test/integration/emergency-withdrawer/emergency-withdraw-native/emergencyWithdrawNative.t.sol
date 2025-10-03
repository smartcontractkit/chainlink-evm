// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {EmergencyWithdrawer} from "../../../../EmergencyWithdrawer.sol";
import {PausableWithAccessControl} from "../../../../PausableWithAccessControl.sol";
import {Errors} from "../../../../libraries/Errors.sol";
import {BaseIntegrationTest} from "../../BaseIntegrationTest.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.0.2/access/AccessControl.sol";
import {Pausable} from "@openzeppelin/contracts@5.0.2/utils/Pausable.sol";

contract EmergencyWithdrawer_EmergencyWithdrawNativeIntegrationTest is BaseIntegrationTest {
  function setUp() public {
    for (uint256 i; i < s_commonContracts[CommonContracts.EMERGENCY_WITHDRAWER].length; ++i) {
      address commonContract = s_commonContracts[CommonContracts.EMERGENCY_WITHDRAWER][i];
      deal(commonContract, 1 ether);
      _changePrank(i_pauser);
      EmergencyWithdrawer(commonContract).emergencyPause();
      _changePrank(i_owner);
    }
  }

  function test_emergencyWithdrawNative() public performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER) {
    deal(i_owner, 0);

    vm.expectEmit(s_contractUnderTest);
    emit EmergencyWithdrawer.AssetEmergencyWithdrawn(i_owner, address(0), 1 ether);

    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdrawNative(payable(i_owner), 1 ether);

    assertEq(s_contractUnderTest.balance, 0);
    assertEq(i_owner.balance, 1 ether);
  }

  function test_emrgencyWithrawNative_RevertWhen_NotPaused()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    _changePrank(i_unpauser);
    PausableWithAccessControl(s_contractUnderTest).emergencyUnpause();

    vm.expectRevert(Pausable.ExpectedPause.selector);
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdrawNative(payable(i_owner), 1 ether);
  }

  function test_emergencyWithdrawNative_RevertWhen_CallerDoesNotHaveADMIN_ROLE()
    public
    whenCallerIsNotAdmin
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    vm.expectRevert(
      abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, i_nonOwner, DEFAULT_ADMIN_ROLE)
    );
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdrawNative(payable(i_owner), 1 ether);
  }

  function test_emergencyWithdrawNative_RevertWhen_ToEqAddressZero()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    vm.expectRevert(Errors.InvalidZeroAddress.selector);
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdrawNative(payable(address(0)), 1 ether);
  }

  function test_emergencyWithdrawNative_RevertWhen_AmountIsZero()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    vm.expectRevert(Errors.InvalidZeroAmount.selector);
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdrawNative(payable(i_owner), 0);
  }

  function test_emergencyWithdrawNative_RevertWhen_FailedNativeTokenTransfer()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    vm.expectRevert(
      abi.encodeWithSelector(
        EmergencyWithdrawer.FailedNativeTokenTransfer.selector,
        address(this),
        1 ether,
        abi.encodeWithSelector(bytes4(keccak256(bytes("Error(string)"))), "Receive not allowed")
      )
    );
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdrawNative(payable(address(this)), 1 ether);
  }

  receive() external payable {
    // solhint-disable-next-line gas-custom-errors
    revert("Receive not allowed");
  }
}
