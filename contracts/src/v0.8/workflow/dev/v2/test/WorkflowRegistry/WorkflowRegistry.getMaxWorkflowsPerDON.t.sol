// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getMaxWorkflowsPerDON is WorkflowRegistrySetup {
  function test_getMaxWorkflowsPerDON_WhenDonLabelHasNeverBeenConfigured() external view {
    // It should return 0
    assertEq(s_registry.getMaxWorkflowsPerDON(s_donLabel), 0);
  }

  function test_getMaxWorkflowsPerDON_WhenDonLabelWasConfiguredToLimit() external {
    // It should return limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donLabel, 100, true);
    assertEq(s_registry.getMaxWorkflowsPerDON(s_donLabel), 100);
  }
}
