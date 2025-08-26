// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_setConfig is WorkflowRegistrySetup {
  function test_setConfig_WhenTheCallerIsNOTTheContractOwner() external {
    // it should revert
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setConfig(10, 8, 150, 256, 3600);
  }

  //whenTheCallerISTheContractOwner
  function test_setConfig_WhenConfigFieldsAreNon_zero() external {
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.ConfigUpdated(12, 6, 180, 512, 3600);
    s_registry.setConfig(12, 6, 180, 512, 3600);
    assertEq(s_registry.getConfig().maxNameLen, 12);
    assertEq(s_registry.getConfig().maxTagLen, 6);
    assertEq(s_registry.getConfig().maxUrlLen, 180);
    assertEq(s_registry.getConfig().maxAttrLen, 512);
    assertEq(s_registry.getConfig().maxExpiryLen, 3600);
  }

  // whenTheCallerISTheContractOwner
  function test_setConfig_WhenSomeConfigFieldsAreZero() external {
    // it should emit ConfigUpdated and store the new config values
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.ConfigUpdated(12, 6, 0, 0, 0);
    s_registry.setConfig(12, 6, 0, 0, 0);
    assertEq(s_registry.getConfig().maxNameLen, 12);
    assertEq(s_registry.getConfig().maxTagLen, 6);
    assertEq(s_registry.getConfig().maxUrlLen, 0);
    assertEq(s_registry.getConfig().maxAttrLen, 0);
    assertEq(s_registry.getConfig().maxExpiryLen, 0);
  }

  // whenTheCallerISTheContractOwner
  function test_setConfig_WhenAllConfigFieldsAreZero() external {
    // it should emit ConfigUpdated and restore default immutable values
    vm.startPrank(s_owner);
    // set value to something else first
    s_registry.setConfig(12, 6, 180, 512, 3600);
    assertEq(s_registry.getConfig().maxNameLen, 12);
    assertEq(s_registry.getConfig().maxTagLen, 6);
    assertEq(s_registry.getConfig().maxUrlLen, 180);
    assertEq(s_registry.getConfig().maxAttrLen, 512);
    assertEq(s_registry.getConfig().maxExpiryLen, 3600);

    // set value to all zero now
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.ConfigUpdated(0, 0, 0, 0, 0);
    s_registry.setConfig(0, 0, 0, 0, 0);
    vm.stopPrank();
    assertEq(s_registry.getConfig().maxNameLen, 0);
    assertEq(s_registry.getConfig().maxTagLen, 0);
    assertEq(s_registry.getConfig().maxUrlLen, 0);
    assertEq(s_registry.getConfig().maxAttrLen, 0);
    assertEq(s_registry.getConfig().maxExpiryLen, 0);
  }
}
