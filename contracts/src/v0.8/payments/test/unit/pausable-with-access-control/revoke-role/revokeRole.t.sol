// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {PausableWithAccessControl} from "../../../../PausableWithAccessControl.sol";
import {BaseUnitTest} from "../../BaseUnitTest.t.sol";

import {IAccessControlDefaultAdminRules} from
  "@openzeppelin/contracts@5.0.2/access/extensions/IAccessControlDefaultAdminRules.sol";

contract RevokeRoleUnitTests is BaseUnitTest {
  function setUp() public {
    _changePrank(i_owner);
  }

  function test_revokeRole_RevertWhen_RevokingDefaultAdminRole()
    public
    performForAllContracts(CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL)
  {
    vm.expectRevert(IAccessControlDefaultAdminRules.AccessControlEnforcedDefaultAdminRules.selector);
    PausableWithAccessControl(s_contractUnderTest).revokeRole(DEFAULT_ADMIN_ROLE, address(this));
  }

  function test_revokeRole_ShouldEnumerateRoleAfterRevoking()
    public
    performForAllContracts(CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL)
  {
    PausableWithAccessControl pausableContractUnderTest = PausableWithAccessControl(s_contractUnderTest);
    pausableContractUnderTest.grantRole(TEST_ROLE, i_unpauser);

    assertEq(pausableContractUnderTest.getRoleMemberCount(TEST_ROLE), 1);
    assertEq(pausableContractUnderTest.getRoleMember(TEST_ROLE, 0), i_unpauser);

    address[] memory members = new address[](1);
    members[0] = i_unpauser;

    assertEq(pausableContractUnderTest.getRoleMembers(TEST_ROLE), members);

    pausableContractUnderTest.revokeRole(TEST_ROLE, i_unpauser);

    assertEq(pausableContractUnderTest.getRoleMemberCount(TEST_ROLE), 0);
    assertEq(pausableContractUnderTest.getRoleMembers(TEST_ROLE), new address[](0));
  }
}
