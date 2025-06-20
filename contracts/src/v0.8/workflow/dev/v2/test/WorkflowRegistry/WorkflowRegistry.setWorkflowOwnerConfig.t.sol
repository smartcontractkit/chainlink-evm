// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_setWorkflowOwnerConfig is WorkflowRegistrySetup {
  function test_setWorkflowOwnerConfig_WhenTheCallerIsNOTTheContractOwner() external {
    // it should revert with OnlyCallableByOwner
    bytes memory blob = abi.encodePacked(uint256(1), uint8(2));
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setWorkflowOwnerConfig(s_user, blob);
  }

  modifier whenTheCallerISTheContractOwner() {
    _;
  }

  function test_setWorkflowOwnerConfig_WhenSettingANonEmptyConfigForTheFirstTime()
    external
    whenTheCallerISTheContractOwner
  {
    // it should store the blob and emit WorkflowOwnerConfigUpdated
    bytes memory blob = abi.encodePacked("hello", uint16(0x1234));

    vm.prank(s_owner);
    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.WorkflowOwnerConfigUpdated(s_user, blob);
    s_registry.setWorkflowOwnerConfig(s_user, blob);

    bytes memory out = s_registry.getWorkflowOwnerConfig(s_user);
    assertEq(keccak256(out), keccak256(blob));
  }

  function test_setWorkflowOwnerConfig_WhenUpdatingToADifferentBlob() external whenTheCallerISTheContractOwner {
    // it should overwrite the blob and emit WorkflowOwnerConfigUpdated
    bytes memory blob1 = abi.encodePacked("foo");
    bytes memory blob2 = abi.encodePacked("bar", uint8(99));

    // first set
    vm.prank(s_owner);
    s_registry.setWorkflowOwnerConfig(s_user, blob1);
    assertEq(keccak256(s_registry.getWorkflowOwnerConfig(s_user)), keccak256(blob1));

    // overwrite
    vm.prank(s_owner);
    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.WorkflowOwnerConfigUpdated(s_user, blob2);
    s_registry.setWorkflowOwnerConfig(s_user, blob2);

    bytes memory out2 = s_registry.getWorkflowOwnerConfig(s_user);
    assertEq(keccak256(out2), keccak256(blob2));
  }

  function test_setWorkflowOwnerConfig_WhenSettingAnEmptyBlob() external whenTheCallerISTheContractOwner {
    // it should clear the stored bytes and emit WorkflowOwnerConfigUpdated
    bytes memory blob1 = abi.encodePacked("keep");
    bytes memory emptyBlob = "";

    // first set
    vm.prank(s_owner);
    s_registry.setWorkflowOwnerConfig(s_user, blob1);
    assertEq(keccak256(s_registry.getWorkflowOwnerConfig(s_user)), keccak256(blob1));

    // clear
    vm.prank(s_owner);
    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.WorkflowOwnerConfigUpdated(s_user, emptyBlob);
    s_registry.setWorkflowOwnerConfig(s_user, emptyBlob);

    bytes memory out3 = s_registry.getWorkflowOwnerConfig(s_user);
    assertEq(out3.length, 0);
  }
}
