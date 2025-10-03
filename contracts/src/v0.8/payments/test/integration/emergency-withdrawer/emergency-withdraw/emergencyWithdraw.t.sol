// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {EmergencyWithdrawer} from "../../../../EmergencyWithdrawer.sol";
import {PausableWithAccessControl} from "../../../../PausableWithAccessControl.sol";
import {Common} from "../../../../libraries/Common.sol";
import {Errors} from "../../../../libraries/Errors.sol";
import {BaseIntegrationTest} from "../../BaseIntegrationTest.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.0.2/access/AccessControl.sol";
import {Pausable} from "@openzeppelin/contracts@5.0.2/utils/Pausable.sol";

contract EmergencyWithdrawer_EmergencyWithdrawIntegrationTest is BaseIntegrationTest {
  Common.AssetAmount[] private s_assetAmounts;

  function setUp() public {
    for (uint256 i; i < s_commonContracts[CommonContracts.EMERGENCY_WITHDRAWER].length; ++i) {
      address commonContract = s_commonContracts[CommonContracts.EMERGENCY_WITHDRAWER][i];
      deal(address(s_mockWETH), commonContract, 1 ether);
      deal(address(s_mockUSDC), commonContract, 1000e6);
      _changePrank(i_pauser);
      EmergencyWithdrawer(commonContract).emergencyPause();
      _changePrank(i_owner);
    }

    s_assetAmounts.push(Common.AssetAmount({asset: address(s_mockWETH), amount: 1 ether}));
    s_assetAmounts.push(Common.AssetAmount({asset: address(s_mockUSDC), amount: 1000e6}));
  }

  function test_emergencyWithdraw_RevertWhen_ContractIsNotPaused()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    _changePrank(i_unpauser);
    PausableWithAccessControl(s_contractUnderTest).emergencyUnpause();
    vm.expectRevert(Pausable.ExpectedPause.selector);
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdraw(i_owner, s_assetAmounts);
  }

  function test_emergencyWithdraw_RevertWhen_CallerDoesNotHaveDEFAULT_ADMIN_ROLE()
    public
    whenCallerIsNotAdmin
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    vm.expectRevert(
      abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, i_nonOwner, DEFAULT_ADMIN_ROLE)
    );
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdraw(i_owner, s_assetAmounts);
  }

  function test_emergencyWithdraw_RevertWhen_EmptyAssetList()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    vm.expectRevert(Errors.EmptyList.selector);
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdraw(i_owner, new Common.AssetAmount[](0));
  }

  function test_emergencyWithdraw_RevertWhen_AssetIsAddressZero()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    s_assetAmounts[0].asset = address(0);
    vm.expectRevert(Errors.InvalidZeroAddress.selector);
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdraw(i_owner, s_assetAmounts);
  }

  function test_emergencyWithdraw_RevertWhen_AmountIsZero()
    public
    performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER)
  {
    s_assetAmounts[0].amount = 0;
    vm.expectRevert(Errors.InvalidZeroAmount.selector);
    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdraw(i_owner, s_assetAmounts);
  }

  function test_emergencyWithdraw_SingleAsset() public performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER) {
    Common.AssetAmount[] memory assetAmounts = new Common.AssetAmount[](1);
    assetAmounts[0] = s_assetAmounts[0];
    deal(s_assetAmounts[0].asset, i_owner, 0);

    vm.expectEmit(s_contractUnderTest);
    emit EmergencyWithdrawer.AssetEmergencyWithdrawn(i_owner, assetAmounts[0].asset, assetAmounts[0].amount);

    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdraw(i_owner, assetAmounts);

    assertEq(s_mockWETH.balanceOf(i_owner), assetAmounts[0].amount);
    assertEq(s_mockWETH.balanceOf(s_contractUnderTest), 0);
  }

  function test_emergencyWithdraw_MultipleAssets() public performForAllContracts(CommonContracts.EMERGENCY_WITHDRAWER) {
    deal(s_assetAmounts[0].asset, i_owner, 0);
    deal(s_assetAmounts[1].asset, i_owner, 0);

    vm.expectEmit(s_contractUnderTest);
    emit EmergencyWithdrawer.AssetEmergencyWithdrawn(i_owner, s_assetAmounts[0].asset, s_assetAmounts[0].amount);
    vm.expectEmit(s_contractUnderTest);
    emit EmergencyWithdrawer.AssetEmergencyWithdrawn(i_owner, s_assetAmounts[1].asset, s_assetAmounts[1].amount);

    EmergencyWithdrawer(s_contractUnderTest).emergencyWithdraw(i_owner, s_assetAmounts);

    assertEq(s_mockWETH.balanceOf(i_owner), s_assetAmounts[0].amount);
    assertEq(s_mockUSDC.balanceOf(i_owner), s_assetAmounts[1].amount);
    assertEq(s_mockWETH.balanceOf(s_contractUnderTest), 0);
    assertEq(s_mockUSDC.balanceOf(s_contractUnderTest), 0);
  }
}
