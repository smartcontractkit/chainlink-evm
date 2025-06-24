// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_DeprecateCapabilitiesTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    s_CapabilitiesRegistry.addCapabilities(capabilities);
  }

  function test_RevertWhen_CalledByNonAdmin() public {
    changePrank(STRANGER);
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_basicCapabilityId;

    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);
  }

  function test_RevertWhen_CapabilityDoesNotExist() public {
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_nonExistentCapabilityId;

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.CapabilityDoesNotExist.selector, s_nonExistentCapabilityId)
    );
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);
  }

  function test_RevertWhen_CapabilityIsDeprecated() public {
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_basicCapabilityId;

    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.CapabilityIsDeprecated.selector, s_basicCapabilityId));
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);
  }

  function test_DeprecatesCapability() public {
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_basicCapabilityId;

    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);
    assertEq(s_CapabilitiesRegistry.isCapabilityDeprecated(s_basicCapabilityId), true);
  }

  function test_EmitsEvent() public {
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_basicCapabilityId;

    vm.expectEmit(address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.CapabilityDeprecated(s_basicCapabilityId);
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);
  }
}
