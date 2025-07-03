// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_totalEvents is WorkflowRegistrySetup {
  function test_totalEvents() external {
    // it should return the total capacity events count
    _linkOwner(s_owner);
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, true); // should add one event
    s_registry.setUserDONOverride(s_user, s_donFamily, 2, true); // should not add an event
    vm.stopPrank();
    vm.prank(s_stranger);
    uint256 total = s_registry.totalEvents();
    assertEq(1, total);
  }
}
