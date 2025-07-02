// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_setMetadataConfig is WorkflowRegistrySetup {
  function test_setMetadataConfig_WhenTheCallerIsNOTTheContractOwner() external {
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    // passes four args instead of a struct
    s_registry.setMetadataConfig(10, 8, 150, 256);
  }

  //whenTheCallerISTheContractOwner
  function test_setMetadataConfig_WhenConfigFieldsAreNon_zero() external {
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.MetadataConfigUpdated(12, 6, 180, 512);
    s_registry.setMetadataConfig(12, 6, 180, 512);
    assertEq(s_registry.maxNameLen(), 12);
    assertEq(s_registry.maxTagLen(), 6);
    assertEq(s_registry.maxUrlLen(), 180);
    assertEq(s_registry.maxAttrLen(), 512);
  }

  // whenTheCallerISTheContractOwner
  function test_setMetadataConfig_WhenSomeConfigFieldsAreZero() external {
    // it should emit MetadataConfigUpdated and store the new config values
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.MetadataConfigUpdated(12, 6, 0, 0);
    s_registry.setMetadataConfig(12, 6, 0, 0);
    assertEq(s_registry.maxNameLen(), 12);
    assertEq(s_registry.maxTagLen(), 6);
    assertEq(s_registry.maxUrlLen(), 0);
    assertEq(s_registry.maxAttrLen(), 0);
  }

  // whenTheCallerISTheContractOwner
  function test_setMetadataConfig_WhenAllConfigFieldsAreZero() external {
    // it should emit MetadataConfigUpdated and restore default immutable values
    vm.startPrank(s_owner);
    // set value to something else first
    s_registry.setMetadataConfig(12, 6, 180, 512);
    assertEq(s_registry.maxNameLen(), 12);
    assertEq(s_registry.maxTagLen(), 6);
    assertEq(s_registry.maxUrlLen(), 180);
    assertEq(s_registry.maxAttrLen(), 512);

    // set value to all zero now
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.MetadataConfigUpdated(0, 0, 0, 0);
    s_registry.setMetadataConfig(0, 0, 0, 0);
    vm.stopPrank();
    // since we cleared the override, we now get the immutable defaults
    assertEq(s_registry.maxNameLen(), 64);
    assertEq(s_registry.maxTagLen(), 32);
    assertEq(s_registry.maxUrlLen(), 200);
    assertEq(s_registry.maxAttrLen(), 1024);
  }
}
