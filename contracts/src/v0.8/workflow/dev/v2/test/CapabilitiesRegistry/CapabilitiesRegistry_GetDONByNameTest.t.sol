// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";

contract CapabilitiesRegistry_GetDONByNameTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);

    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    vm.stopPrank();
    vm.startPrank(ADMIN);
    s_CapabilitiesRegistry.addDONs(s_paramsForTwoDONs);
  }

  function test_RevertWhen_DONDoesNotExist() public {
    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.DONWithNameDoesNotExist.selector, "non-existent-don-name")
    );
    s_CapabilitiesRegistry.getDONByName("non-existent-don-name");
  }

  function test_CorrectlyFetchesDONByName() public view {
    CapabilitiesRegistry.DONInfo memory don = s_CapabilitiesRegistry.getDONByName(TEST_DON_NAME_ONE);
    assertEq(don.id, DON_ID, "DON ID mismatch");
    assertEq(don.configCount, 1, "Config count mismatch");
    assertEq(don.isPublic, true, "Is public mismatch");
    assertEq(don.acceptsWorkflows, true, "Accepts workflows mismatch");
    assertEq(don.f, 1, "F mismatch");
    assertEq(don.name, TEST_DON_NAME_ONE, "Name mismatch");
    assertEq(don.config, bytes(""), "Config mismatch");
    assertEq(don.donFamilies.length, 1, "Don families length mismatch");
    assertEq(don.donFamilies[0], TEST_DON_FAMILY_ONE, "Don family mismatch");
    assertEq(don.capabilityConfigurations.length, s_capabilityConfigs.length, "Capability configs length mismatch");
    assertEq(don.capabilityConfigurations[0].capabilityId, s_basicCapabilityId, "Capability ID mismatch");

    don = s_CapabilitiesRegistry.getDONByName(TEST_DON_NAME_TWO);
    assertEq(don.id, DON_ID_TWO, "DON ID mismatch");
    assertEq(don.configCount, 1, "Config count mismatch");
    assertEq(don.isPublic, false, "Is public mismatch");
    assertEq(don.acceptsWorkflows, false, "Accepts workflows mismatch");
    assertEq(don.f, 1, "F mismatch");
    assertEq(don.capabilityConfigurations.length, s_capabilityConfigs.length, "Capability configs length mismatch");
    assertEq(don.capabilityConfigurations[0].capabilityId, s_basicCapabilityId, "Capability ID mismatch");
    assertEq(don.name, TEST_DON_NAME_TWO, "Name mismatch");
    assertEq(don.config, TEST_DON_CONFIG, "Config mismatch");
    assertEq(don.donFamilies.length, 1, "Don families length mismatch");
    assertEq(don.donFamilies[0], TEST_DON_FAMILY_ONE, "Don family mismatch");
  }
}
