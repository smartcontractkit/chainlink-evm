// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getUserDONOverrides is WorkflowRegistrySetup {
  function test_getUserDONOverrides_WhenThereAreNoUserOverridesForASpecificDON() external view {
    // it should return an empty list
    WorkflowRegistry.UserOverrideView[] memory list = s_registry.getUserDONOverrides(s_donFamily, 0, 100);
    assertEq(list.length, 0);
  }

  function test_getUserDONOverrides_WhenThereAreUserOverridesForASpecificDON() external {
    // it should return the user overrides for that DON

    // set the don configs
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, true);
    s_registry.setDONLimit("fast-pool", 200, true);
    s_registry.setDONLimit("slow-pool", 150, true);
    s_registry.setUserDONOverride(s_user, s_donFamily, 10, true);
    s_registry.setUserDONOverride(s_user, "fast-pool", 20, true);
    s_registry.setUserDONOverride(s_user, "slow-pool", 30, true);
    vm.stopPrank();

    WorkflowRegistry.UserOverrideView[] memory list = s_registry.getUserDONOverrides(s_donFamily, 0, 100);
    assertEq(list.length, 1);
    assertEq(list[0].user, s_user);
    assertEq(list[0].limit, 10);

    list = s_registry.getUserDONOverrides("fast-pool", 0, 100);
    assertEq(list.length, 1);
    assertEq(list[0].user, s_user);
    assertEq(list[0].limit, 20);

    list = s_registry.getUserDONOverrides("slow-pool", 0, 100);
    assertEq(list.length, 1);
    assertEq(list[0].user, s_user);
    assertEq(list[0].limit, 30);
  }
}
