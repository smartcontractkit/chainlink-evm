// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_GetDONByNameTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);

    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    vm.stopPrank();
    vm.startPrank(ADMIN);
  }

  function test_CorrectlyFetchesDONFamilies() public {
    s_CapabilitiesRegistry.addDONs(s_paramsForTwoDONs);

    string[] memory families = s_CapabilitiesRegistry.getDONFamilies();
    assertEq(families.length, 1, "Families length mismatch");
    assertEq(families[0], TEST_DON_FAMILY_ONE, "Expected only default family");

    uint256[] memory donIds = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(donIds.length, 2, "Expected 2 DONs in default family");
    assertEq(donIds[0], DON_ID, "First DON doesn't belong to the DON family as expected");
    assertEq(donIds[1], DON_ID_TWO, "Second DON doesn't belong to the DON family as expected");
  }
}
