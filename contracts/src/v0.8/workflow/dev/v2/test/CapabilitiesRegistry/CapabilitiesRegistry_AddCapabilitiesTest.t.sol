// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {ICapabilityConfiguration} from "../../interfaces/ICapabilityConfiguration.sol";

import {BaseTest} from "./BaseTest.t.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

contract CapabilitiesRegistry_AddCapabilitiesTest is BaseTest {
  function test_RevertWhen_CalledByNonAdmin() public {
    changePrank(STRANGER);

    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](1);
    capabilities[0] = s_basicCapability;

    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    s_CapabilitiesRegistry.addCapabilities(capabilities);
  }

  function test_RevertWhen_CapabilityExists() public {
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](1);
    capabilities[0] = s_basicCapability;

    // Successfully add the capability the first time
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    // Try to add the same capability again
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.CapabilityAlreadyExists.selector, s_basicCapabilityId));
    s_CapabilitiesRegistry.addCapabilities(capabilities);
  }

  function test_RevertWhen_ConfigurationContractNotDeployed() public {
    address nonExistentContract = address(1);
    s_capabilityWithConfigurationContract.configurationContract = nonExistentContract;

    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](1);
    capabilities[0] = s_capabilityWithConfigurationContract;

    vm.expectRevert(
      abi.encodeWithSelector(
        CapabilitiesRegistry.InvalidCapabilityConfigurationContractInterface.selector, nonExistentContract
      )
    );
    s_CapabilitiesRegistry.addCapabilities(capabilities);
  }

  function test_RevertWhen_ConfigurationContractDoesNotMatchInterface() public {
    address contractWithoutERC165 = address(9999);
    vm.mockCall(
      contractWithoutERC165,
      abi.encodeWithSelector(
        IERC165.supportsInterface.selector,
        ICapabilityConfiguration.getCapabilityConfiguration.selector
          ^ ICapabilityConfiguration.beforeCapabilityConfigSet.selector
      ),
      abi.encode(false)
    );
    s_capabilityWithConfigurationContract.configurationContract = contractWithoutERC165;
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](1);
    capabilities[0] = s_capabilityWithConfigurationContract;

    vm.expectRevert(
      abi.encodeWithSelector(
        CapabilitiesRegistry.InvalidCapabilityConfigurationContractInterface.selector, contractWithoutERC165
      )
    );
    s_CapabilitiesRegistry.addCapabilities(capabilities);
  }

  function test_AddCapability_NoConfigurationContract() public {
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](1);
    capabilities[0] = s_basicCapability;

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.CapabilityConfigured(s_basicCapability.capabilityId);
    s_CapabilitiesRegistry.addCapabilities(capabilities);
    CapabilitiesRegistry.CapabilityInfo memory storedCapability =
      s_CapabilitiesRegistry.getCapability(s_basicCapability.capabilityId);

    assertEq(storedCapability.capabilityId, s_basicCapability.capabilityId);
    assertEq(storedCapability.metadata, s_basicCapability.metadata);
    assertEq(storedCapability.configurationContract, s_basicCapability.configurationContract);
  }

  function test_AddCapability_WithConfiguration() public {
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](1);
    capabilities[0] = s_capabilityWithConfigurationContract;

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.CapabilityConfigured(s_capabilityWithConfigurationContract.capabilityId);
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    CapabilitiesRegistry.CapabilityInfo memory storedCapability =
      s_CapabilitiesRegistry.getCapability(s_capabilityWithConfigurationContract.capabilityId);

    assertEq(storedCapability.capabilityId, s_capabilityWithConfigurationContract.capabilityId);
    assertEq(storedCapability.metadata, s_capabilityWithConfigurationContract.metadata);
    assertEq(storedCapability.configurationContract, s_capabilityWithConfigurationContract.configurationContract);
  }
}
