// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_totalLinkedOwners is WorkflowRegistrySetup {
  function test_totalLinkedOwners() external {
    // it should return the total capacity events count
    _linkOwner(s_owner);
    _linkOwner(s_user);
    _linkOwner(s_stranger);
    uint256 total = s_registry.totalLinkedOwners();
    assertEq(3, total);
  }
}
