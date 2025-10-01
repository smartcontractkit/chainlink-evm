// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_adminPauseAllByDON is WorkflowRegistrySetup {
  function test_adminPauseAllByDON_WhenCallerIsNOTTheContractOwner() external {
    // it reverts with Ownable2StepMsgSender caller is not the owner
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.adminPauseAllByDON(s_donFamily, 0); // 0 means no limit
  }

  // whenCallerIsTheContractOwner
  function test_adminPauseAllByDON_WhenThereAreActiveWorkflows() external {
    // it pauses all of the workflows
    bytes32 wfId2 = keccak256("workflow-id2");
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, 10);
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
    s_registry.adminPauseAllByDON(s_donFamily, 0); // 0 means no limit

    // confirm the workflows are now paused
    wf1 = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf1.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    wf2 = s_registry.getWorkflowById(wfId2);
    assertEq(uint8(wf2.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
  }

  // whenCallerIsTheContractOwner
  function test_adminPauseAllByDON_WhenThereAreActiveWorkflowsButLimitIsSet() external {
    // it pauses only up to the limit number of workflows
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, 10);
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
      keccak256("workflow-id2"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    s_registry.upsertWorkflow(
      "workflow-3",
      s_tag,
      keccak256("workflow-id3"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    s_registry.upsertWorkflow(
      "workflow-4",
      s_tag,
      keccak256("workflow-id4"),
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
    WorkflowRegistry.WorkflowMetadataView memory wf2 = s_registry.getWorkflowById(keccak256("workflow-id2"));
    assertEq(uint8(wf2.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));
    WorkflowRegistry.WorkflowMetadataView memory wf3 = s_registry.getWorkflowById(keccak256("workflow-id3"));
    assertEq(uint8(wf3.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));
    WorkflowRegistry.WorkflowMetadataView memory wf4 = s_registry.getWorkflowById(keccak256("workflow-id4"));
    assertEq(uint8(wf4.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));

    vm.prank(s_owner);
    s_registry.adminPauseAllByDON(s_donFamily, 2); // pause first 2 workflows

    // confirm the two workflows are now paused
    wf1 = s_registry.getWorkflowById(keccak256("workflow-id4"));
    assertEq(uint8(wf1.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    wf2 = s_registry.getWorkflowById(keccak256("workflow-id3"));
    assertEq(uint8(wf2.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    // but the other two are still active
    wf3 = s_registry.getWorkflowById(keccak256("workflow-id1"));
    assertEq(uint8(wf3.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));
    wf4 = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf4.status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE));

    vm.prank(s_owner);
    s_registry.adminPauseAllByDON(s_donFamily, 2); // pause the remaining 2 workflows

    // confirm that all workflows are now paused
    wf1 = s_registry.getWorkflowById(s_workflowId);
    assertEq(uint8(wf1.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    wf2 = s_registry.getWorkflowById(keccak256("workflow-id2"));
    assertEq(uint8(wf2.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    wf3 = s_registry.getWorkflowById(keccak256("workflow-id3"));
    assertEq(uint8(wf3.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
    wf4 = s_registry.getWorkflowById(keccak256("workflow-id4"));
    assertEq(uint8(wf4.status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED));
  }
}
