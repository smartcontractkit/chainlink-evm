// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {PausableWithAccessControl} from "../../../../PausableWithAccessControl.sol";
import {BaseUnitTest} from "../../BaseUnitTest.t.sol";

import {IAccessControlDefaultAdminRules} from
  "@openzeppelin/contracts@5.0.2/access/extensions/IAccessControlDefaultAdminRules.sol";

contract GrantRoleUnitTests is BaseUnitTest {
  function setUp() public {
    _changePrank(i_owner);
  }

  function test_grantRole_RevertWhen_GrantingDefaultAdminRole()
    public
    performForAllContracts(CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL)
  {
    vm.expectRevert(IAccessControlDefaultAdminRules.AccessControlEnforcedDefaultAdminRules.selector);
    PausableWithAccessControl(s_contractUnderTest).grantRole(DEFAULT_ADMIN_ROLE, address(this));
  }

  function test_grantRole_ShouldEnumerateRoleAfterGranting()
    public
    performForAllContracts(CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL)
  {
    assertEq(PausableWithAccessControl(s_contractUnderTest).getRoleMemberCount(TEST_ROLE), 0);
    assertEq(PausableWithAccessControl(s_contractUnderTest).getRoleMembers(TEST_ROLE), new address[](0));

    PausableWithAccessControl(s_contractUnderTest).grantRole(TEST_ROLE, i_unpauser);

    assertEq(PausableWithAccessControl(s_contractUnderTest).getRoleMemberCount(TEST_ROLE), 1);
    assertEq(PausableWithAccessControl(s_contractUnderTest).getRoleMember(TEST_ROLE, 0), i_unpauser);

    address[] memory members = new address[](1);
    members[0] = i_unpauser;

    assertEq(PausableWithAccessControl(s_contractUnderTest).getRoleMembers(TEST_ROLE), members);
  }
}
