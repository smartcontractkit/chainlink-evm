// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_isAllowedSigner is WorkflowRegistrySetup {
  function test_isAllowedSigner_WhenTheSignerAddressHasNeverBeenConfigured() external view {
    // It should return false
    assertEq(s_registry.isAllowedSigner(s_user), false);
  }

  function test_isAllowedSigner_WhenTheSignerIsConfigured() external view {
    // It should return true
    assertEq(s_registry.isAllowedSigner(s_allowedSigner), true);
  }
}
