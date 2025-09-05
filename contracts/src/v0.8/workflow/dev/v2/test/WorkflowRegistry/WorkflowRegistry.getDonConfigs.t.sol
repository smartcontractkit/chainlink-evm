// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getDonConfigs is WorkflowRegistrySetup {
  function test_getDonConfigs_WhenThereAreNoDONConfigs() external view {
    // it should return an empty list
    WorkflowRegistry.DonConfigView[] memory list = s_registry.getDonConfigs(0, 100);
    assertEq(list.length, 0);
  }

  function test_getDonConfigs_WhenThereAreDONConfigsSet() external {
    // it should return the DON configs

    // set the don configs
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, true);
    s_registry.setDONLimit("fast-pool", 200, true);
    s_registry.setDONLimit("slow-pool", 150, true);
    vm.stopPrank();
    WorkflowRegistry.DonConfigView[] memory list = s_registry.getDonConfigs(0, 100);
    assertEq(list.length, 3);
    assertEq(list[0].family, s_donFamily);
    assertEq(list[0].limit, 100);
    assertEq(list[1].family, "fast-pool");
    assertEq(list[1].limit, 200);
    assertEq(list[2].family, "slow-pool");
    assertEq(list[2].limit, 150);
  }
}
