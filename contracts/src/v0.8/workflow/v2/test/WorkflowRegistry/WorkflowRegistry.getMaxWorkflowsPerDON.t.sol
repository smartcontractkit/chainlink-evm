// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getMaxWorkflowsPerDON is WorkflowRegistrySetup {
  function test_getMaxWorkflowsPerDON_WhenDonLabelHasNeverBeenConfigured() external view {
    // It should return 0
    (uint32 donLimit, uint32 defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, 0);
    assertEq(defaultUserLimit, 0);
  }

  function test_getMaxWorkflowsPerDON_WhenDonLabelWasConfiguredToLimit() external {
    // It should return limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, 10);
    (uint32 donLimit, uint32 defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, 100);
    assertEq(defaultUserLimit, 10);
  }
}
