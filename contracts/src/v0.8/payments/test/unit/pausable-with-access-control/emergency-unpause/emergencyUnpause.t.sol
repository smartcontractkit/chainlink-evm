// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {PausableWithAccessControl} from "../../../../PausableWithAccessControl.sol";
import {Roles} from "../../../../libraries/Roles.sol";
import {BaseUnitTest} from "../../BaseUnitTest.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.0.2/access/IAccessControl.sol";
import {Pausable} from "@openzeppelin/contracts@5.0.2/utils/Pausable.sol";

contract EmergencyUnpauseUnitTest is BaseUnitTest {
  function setUp() public {
    _changePrank(i_unpauser);
  }

  function test_emergencyUnpause_RevertWhen_CallerDoesNotHaveUNPAUSER_ROLE()
    public
    performForAllContracts(CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL)
  {
    _changePrank(i_nonOwner);
    vm.expectRevert(
      abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, i_nonOwner, Roles.UNPAUSER_ROLE)
    );
    PausableWithAccessControl(s_contractUnderTest).emergencyUnpause();
  }

  function test_emergencyUnpause() public performForAllContracts(CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL) {
    _changePrank(i_pauser);
    PausableWithAccessControl(s_contractUnderTest).emergencyPause();
    _changePrank(i_unpauser);
    vm.expectEmit(address(s_contractUnderTest));
    emit Pausable.Unpaused(i_unpauser);
    PausableWithAccessControl(s_contractUnderTest).emergencyUnpause();

    assertFalse(PausableWithAccessControl(s_contractUnderTest).paused());
  }
}
