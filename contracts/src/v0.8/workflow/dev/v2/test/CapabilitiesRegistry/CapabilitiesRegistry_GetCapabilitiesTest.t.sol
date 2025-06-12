// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_GetCapabilitiesTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;
    s_CapabilitiesRegistry.addCapabilities(capabilities);
  }

  function test_ReturnsCapabilities() public {
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_basicCapabilityId;
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);

    CapabilitiesRegistry.CapabilityInfo[] memory capabilities = s_CapabilitiesRegistry.getCapabilities();

    assertEq(capabilities.length, 2);

    assertEq(capabilities[0].capabilityId, s_basicCapabilityId);
    assertEq(capabilities[0].metadata, TEST_CAPABILITY_METADATA);
    assertEq(capabilities[0].configurationContract, address(0));
    assertEq(capabilities[0].isDeprecated, true);

    assertEq(capabilities[1].capabilityId, s_capabilityWithConfigurationContractId);
    assertEq(capabilities[1].metadata, bytes(""));
    assertEq(capabilities[1].configurationContract, address(s_capabilityConfigurationContract));
    assertEq(capabilities[1].capabilityId, s_capabilityWithConfigurationContractId);
    assertEq(capabilities[1].isDeprecated, false);
  }
}
