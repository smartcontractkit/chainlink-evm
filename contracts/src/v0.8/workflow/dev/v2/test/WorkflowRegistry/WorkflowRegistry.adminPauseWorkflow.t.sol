// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_adminPauseWorkflow is WorkflowRegistrySetup {
  function test_adminPauseWorkflow_WhenCallerIsNOTTheContractOwner() external {
    // when caller is NOT the contract owner
    // it reverts with Ownable
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.adminPauseWorkflow(s_workflowId);
  }

  // whenCallerIsTheContractOwner
  function test_adminPauseWorkflow_WhenWorkflowStatusIsPAUSED() external {
    // setup: configure a DON limit and register a workflow as PAUSED
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);
    (bytes32 ownerProof, bytes memory sig) = _getLinkProofSignature(s_owner);
    s_registry.linkOwner(s_validityTimestamp, ownerProof, sig);

    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    // precondition check
    WorkflowRegistry.WorkflowMetadataView memory before = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(before.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));

    // when workflow is already PAUSED
    // it returns immediately, no change
    s_registry.adminPauseWorkflow(s_workflowId);
    WorkflowRegistry.WorkflowMetadataView memory wf = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    vm.stopPrank();
  }

  // whenCallerIsTheContractOwner
  function test_adminPauseWorkflow_WhenWorkflowStatusIsACTIVE() external {
    // setup: configure a DON limit and register a workflow as ACTIVE
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);
    (bytes32 ownerProof, bytes memory sig) = _getLinkProofSignature(s_owner);
    s_registry.linkOwner(s_validityTimestamp, ownerProof, sig);

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

    // precondition check
    WorkflowRegistry.WorkflowMetadataView memory before = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(before.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));

    // when workflow is ACTIVE
    // it pauses the workflow
    s_registry.adminPauseWorkflow(s_workflowId);
    WorkflowRegistry.WorkflowMetadataView memory wf = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    vm.stopPrank();
  }
}
