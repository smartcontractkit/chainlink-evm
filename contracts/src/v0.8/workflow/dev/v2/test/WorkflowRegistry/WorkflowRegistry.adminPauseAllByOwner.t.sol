// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_adminPauseAllByOwner is WorkflowRegistrySetup {
  function test_adminPauseAllByOwner_WhenCallerIsNOTTheContractOwner() external {
    // it reverts with Ownable2StepMsgSender caller is not the owner
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.adminPauseAllByOwner(s_user);
  }

  // whenCallerIsTheContractOwner
  function test_adminPauseAllByOwner_WhenThereAreActiveWorkflows() external {
    // it pauses all of the workflows
    bytes32 wfId2 = keccak256("workflow-id2");
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);
    _linkOwner(s_user);

    // add some workflows
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

    // check the workflows are active
    WorkflowRegistry.WorkflowMetadataView memory wf1 = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf1.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));
    WorkflowRegistry.WorkflowMetadataView memory wf2 = s_registry.getWorkflowById(wfId2);
    assertEq(uint8(wf2.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));

    vm.prank(s_owner);
    s_registry.adminPauseAllByOwner(s_user);

    // confirm the workflows are now paused
    wf1 = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf1.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    wf2 = s_registry.getWorkflowById(wfId2);
    assertEq(uint8(wf2.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
  }
}
