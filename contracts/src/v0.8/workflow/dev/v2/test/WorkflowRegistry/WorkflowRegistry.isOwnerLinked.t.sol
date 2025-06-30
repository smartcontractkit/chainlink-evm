// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistryisOwnerLinked is WorkflowRegistrySetup {
  function test_WhenTheSpecifiedOwnerIsLinked() external {
    // It should return true
    _linkOwner(s_user);
    assertTrue(s_registry.isOwnerLinked(s_user));
  }

  function test_WhenTheSpecifiedOwnerHasNotLinked() external {
    // It should return false
    assertFalse(s_registry.isOwnerLinked(s_user));
  }
}
