// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_adminBatchPauseWorkflows is WorkflowRegistrySetup {
  function test_adminBatchPauseWorkflows_WhenCallerIsNOTTheContractOwner() external {
    // it reverts with OnlyOwner
    bytes32[] memory workflowIds = new bytes32[](2);
    workflowIds[0] = s_workflowId;
    workflowIds[1] = keccak256("workflow-id2");
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.adminBatchPauseWorkflows(workflowIds);
  }

  // whenCallerIsTheContractOwner
  function test_adminBatchPauseWorkflows_WhenWorkflowIdsLengthIs0() external {
    // it reverts EmptyUpdateBatch
    bytes32[] memory workflowIds = new bytes32[](0);
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.EmptyUpdateBatch.selector));
    s_registry.adminBatchPauseWorkflows(workflowIds);
  }

  // whenCallerIsTheContractOwner
  function test_adminBatchPauseWorkflows_WhenWorkflowIdsIsNotZero() external {
    // it pauses each workflow in workflowIds
    bytes32[] memory workflowIds = new bytes32[](2);
    bytes32 wfId2 = keccak256("workflow-id2");
    workflowIds[0] = s_workflowId;
    workflowIds[1] = wfId2;

    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);
    _linkOwner(s_user);

    // add some workflows for a different user
    vm.startPrank(s_user);
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    s_registry.upsertWorkflow(
      "workflow-2",
      s_tag,
      wfId2,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    vm.stopPrank();

    vm.prank(s_owner);
    s_registry.adminBatchPauseWorkflows(workflowIds);

    WorkflowRegistry.WorkflowMetadataView memory wf1 = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf1.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    WorkflowRegistry.WorkflowMetadataView memory wf2 = s_registry.getWorkflowById(wfId2);
    assertEq(uint8(wf2.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
  }
}
