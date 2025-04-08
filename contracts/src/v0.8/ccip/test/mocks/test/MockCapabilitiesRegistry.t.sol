// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {MockCapabilitiesRegistry} from "../MockCapabilitiesRegistry.sol";
import {Test} from "forge-std/Test.sol";

contract MockCapabilitiesRegistryTest is Test {
  MockCapabilitiesRegistry private s_mockCapabilitiesRegistry;
  uint32 internal constant INITIAL_DON_ID = 100;
  address private s_deployer = address(0x1);

  function setUp() public {
    s_mockCapabilitiesRegistry = new MockCapabilitiesRegistry(INITIAL_DON_ID);
  }

  function test_GetNextDONId() public {
    vm.prank(s_deployer);

    uint32 nextDonId = s_mockCapabilitiesRegistry.getNextDONId();
    assertEq(nextDonId, INITIAL_DON_ID, "The next DON ID should match the initial value.");
  }
}
