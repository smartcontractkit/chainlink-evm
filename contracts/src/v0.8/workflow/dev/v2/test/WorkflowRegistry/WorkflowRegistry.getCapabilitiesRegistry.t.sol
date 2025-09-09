// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getCapabilitiesRegistry is WorkflowRegistrySetup {
  function test_getCapabilitiesRegistry_WhenTheRegistryHasnNotBeenSetYet() external view {
    // it should return address 0, 0

    (address capRegValue, uint64 chainSelValue) = s_registry.getCapabilitiesRegistry();
    assertEq(chainSelValue, 0);
    assertEq(capRegValue, address(0));
  }

  function test_getCapabilitiesRegistry_WhenTheRegistryHasBeenSet() external {
    // it should return the Capabilities Registry values

    // set the Capabilities Registry
    vm.prank(s_owner);
    address capRegAddr = makeAddr("cap-registry-address");
    uint64 chainSel = 123_456;
    s_registry.setCapabilitiesRegistry(capRegAddr, chainSel);

    (address capRegValue, uint64 chainSelValue) = s_registry.getCapabilitiesRegistry();
    assertEq(chainSelValue, chainSel);
    assertEq(capRegValue, capRegAddr);
  }
}
