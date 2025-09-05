// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getMaxWorkflowsPerUserDON is WorkflowRegistrySetup {
  // whenNoUserOverrideExistsForUserDONLabel
  function test_getMaxWorkflowsPerUserDON_WhenGlobalDONLimitIsUnset() external view {
    // It should return 0
    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 0);
  }

  // whenNoUserOverrideExistsForUserDONLabel
  function test_getMaxWorkflowsPerUserDON_WhenGlobalDONLimitIsSetToL() external {
    // It should return L

    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, true);
    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 100);
  }

  function test_getMaxWorkflowsPerUserDON_WhenAUserOverrideExistsAndIsEnabled() external {
    // It should return the value

    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, true);
    s_registry.setUserDONOverride(s_user, s_donFamily, 5, true);
    vm.stopPrank();
    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 5);
  }
}
