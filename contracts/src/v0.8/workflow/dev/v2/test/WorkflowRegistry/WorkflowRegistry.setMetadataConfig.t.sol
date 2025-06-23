// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_setMetadataConfig is WorkflowRegistrySetup {
  function test_WhenTheCallerIsNOTTheContractOwner() external {
    WorkflowRegistry.MetadataConfig memory newCfg = WorkflowRegistry.MetadataConfig({
      maxWorkflowNameLength: 10,
      maxWorkflowTagLength: 8,
      maxUrlLength: 150,
      maxAttributesLength: 256
    });

    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setMetadataConfig(newCfg);
  }

  //whenTheCallerISTheContractOwner
  function test_WhenConfigFieldsAreNon_zero() external {
    WorkflowRegistry.MetadataConfig memory newCfg = WorkflowRegistry.MetadataConfig({
      maxWorkflowNameLength: 12,
      maxWorkflowTagLength: 6,
      maxUrlLength: 180,
      maxAttributesLength: 512
    });

    vm.prank(s_owner);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.MetadataConfigUpdated(12, 6, 180, 512);

    s_registry.setMetadataConfig(newCfg);

    WorkflowRegistry.MetadataConfig memory stored = s_registry.getMetadataConfig();
    assertEq(stored.maxWorkflowNameLength, 12);
    assertEq(stored.maxWorkflowTagLength, 6);
    assertEq(stored.maxUrlLength, 180);
    assertEq(stored.maxAttributesLength, 512);
  }

  // whenTheCallerISTheContractOwner
  function test_WhenAllConfigFieldsAreZero() external {
    // Zeroing every field
    WorkflowRegistry.MetadataConfig memory newCfg = WorkflowRegistry.MetadataConfig({
      maxWorkflowNameLength: 0,
      maxWorkflowTagLength: 0,
      maxUrlLength: 0,
      maxAttributesLength: 0
    });

    vm.prank(s_owner);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.MetadataConfigUpdated(0, 0, 0, 0);

    s_registry.setMetadataConfig(newCfg);

    WorkflowRegistry.MetadataConfig memory stored = s_registry.getMetadataConfig();
    assertEq(stored.maxWorkflowNameLength, 0);
    assertEq(stored.maxWorkflowTagLength, 0);
    assertEq(stored.maxUrlLength, 0);
    assertEq(stored.maxAttributesLength, 0);
  }
}
