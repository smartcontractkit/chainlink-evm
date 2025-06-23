// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getDONRegistry is WorkflowRegistrySetup {
  function test_getDONRegistry_WhenTheRegistryHasnNotBeenSetYet() external view {
    // it should return address 0, 0

    (address donRegValue, uint64 chainSelValue) = s_registry.getDONRegistry();
    assertEq(chainSelValue, 0);
    assertEq(donRegValue, address(0));
  }

  function test_getDONRegistry_WhenTheRegistryHasBeenSet() external {
    // it should return the DON registry values

    // set the don registry
    vm.prank(s_owner);
    address donRegAddr = makeAddr("don-registry-address");
    uint64 chainSel = 123456;
    s_registry.setDONRegistry(donRegAddr, chainSel);

    (address donRegValue, uint64 chainSelValue) = s_registry.getDONRegistry();
    assertEq(chainSelValue, chainSel);
    assertEq(donRegValue, donRegAddr);
  }
}
