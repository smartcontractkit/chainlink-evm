// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";

contract CapabilitiesRegistry_GetDONsTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);

    changePrank(NODE_OPERATOR_ONE_ADMIN);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    changePrank(ADMIN);

    s_CapabilitiesRegistry.addDONs(s_paramsForTwoDONs);
  }

  function test_CorrectlyFetchesDONs() public view {
    CapabilitiesRegistry.DONInfo[] memory dons = s_CapabilitiesRegistry.getDONs();
    assertEq(dons.length, 2);
    assertEq(dons[0].id, DON_ID);
    assertEq(dons[0].configCount, 1);
    assertEq(dons[0].isPublic, true);
    assertEq(dons[0].acceptsWorkflows, true);
    assertEq(dons[0].f, 1);
    assertEq(dons[0].capabilityConfigurations.length, s_capabilityConfigs.length);
    assertEq(dons[0].capabilityConfigurations[0].capabilityId, s_basicCapabilityId);

    assertEq(dons[1].id, DON_ID_TWO);
    assertEq(dons[1].configCount, 1);
    assertEq(dons[1].isPublic, false);
    assertEq(dons[1].capabilityConfigurations.length, s_capabilityConfigs.length);
    assertEq(dons[1].capabilityConfigurations[0].capabilityId, s_basicCapabilityId);
  }

  function test_DoesNotIncludeRemovedDONs() public {
    uint32[] memory removedDONIDs = new uint32[](1);
    removedDONIDs[0] = DON_ID;
    s_CapabilitiesRegistry.removeDONs(removedDONIDs);

    CapabilitiesRegistry.DONInfo[] memory dons = s_CapabilitiesRegistry.getDONs();
    assertEq(dons.length, 1);
    assertEq(dons[0].id, DON_ID_TWO);
    assertEq(dons[0].configCount, 1);
    assertEq(dons[0].isPublic, false);
    assertEq(dons[0].acceptsWorkflows, false);
    assertEq(dons[0].f, 1);
    assertEq(dons[0].capabilityConfigurations.length, s_capabilityConfigs.length);
    assertEq(dons[0].capabilityConfigurations[0].capabilityId, s_basicCapabilityId);
  }
}
