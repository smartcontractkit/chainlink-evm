// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_batchPauseWorkflows is WorkflowRegistrySetup {
  function test_batchPauseWorkflows_WhenCallerIsNOTALinkedOwner() external {
    // it reverts with OwnershipLinkDoesNotExist
    bytes32[] memory workflowIds = new bytes32[](2);
    bytes32 wfId2 = keccak256("workflow-id2");
    workflowIds[0] = s_workflowId;
    workflowIds[1] = wfId2;

    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_owner));
    s_registry.batchPauseWorkflows(workflowIds);
  }

  modifier whenCallerISALinkedOwner() {
    _linkOwner(s_user);
    _;
  }

  function test_batchPauseWorkflows_WhenWorkflowIdsLengthIs0() external whenCallerISALinkedOwner {
    // it reverts with EmptyUpdateBatch
    bytes32[] memory workflowIds = new bytes32[](0);
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.EmptyUpdateBatch.selector));
    s_registry.batchPauseWorkflows(workflowIds);
  }

  function test_batchPauseWorkflows_WhenWorkflowIdsContainsAnUnknownID() external whenCallerISALinkedOwner {
    // it reverts with WorkflowDoesNotExist
    bytes32[] memory workflowIds = new bytes32[](1);
    workflowIds[0] = s_workflowId;

    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowDoesNotExist.selector, s_workflowId));
    s_registry.batchPauseWorkflows(workflowIds);
  }

  function test_batchPauseWorkflows_WhenWorkflowIdsContainsAnIDNotOwnedByCaller() external whenCallerISALinkedOwner {
    // it reverts with CallerIsNotWorkflowOwner
    _setDONLimit();
    vm.prank(s_user);
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

    bytes32[] memory workflowIds = new bytes32[](1);
    workflowIds[0] = s_workflowId;

    _linkOwner(s_stranger);
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.CallerIsNotWorkflowOwner.selector, s_stranger));
    s_registry.batchPauseWorkflows(workflowIds);
  }

  function test_batchPauseWorkflows_WhenEveryListedWorkflowIsAlreadyPAUSED() external whenCallerISALinkedOwner {
    // it emits no WorkflowPaused events and leaves state unchanged
    _setDONLimit();
    vm.startPrank(s_user);
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

    bytes32[] memory workflowIds = new bytes32[](1);
    workflowIds[0] = s_workflowId;

    vm.recordLogs();
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
    vm.stopPrank();
    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("WorkflowPaused(bytes32,address,string,string)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("WorkflowPaused was emitted when it should not have been");
        fail();
      }
    }
  }

  function test_batchPauseWorkflows_WhenEveryListedWorkflowIsACTIVE() external whenCallerISALinkedOwner {
    // it pauses each workflow and emits a WorkflowPaused event for each
    _setDONLimit();
    bytes32[] memory workflowIds = new bytes32[](2);
    string memory wfName2 = "workflow-2";
    bytes32 wfId2 = keccak256("workflow-2");
    workflowIds[0] = s_workflowId;
    workflowIds[1] = wfId2;

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
      wfName2,
      s_tag,
      wfId2,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.WorkflowPaused(s_workflowId, s_user, s_donFamily, s_workflowName);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.WorkflowPaused(wfId2, s_user, s_donFamily, wfName2);
    s_registry.batchPauseWorkflows(workflowIds);
    vm.stopPrank();
  }

  function test_batchPauseWorkflows_WhenTheListMixesPAUSEDAndACTIVEWorkflows() external whenCallerISALinkedOwner {
    // it pauses only the ACTIVE ones and emits events just for them
    _setDONLimit();
    bytes32[] memory workflowIds = new bytes32[](2);
    string memory wfName2 = "workflow-2";
    bytes32 wfId2 = keccak256("workflow-2");
    workflowIds[0] = s_workflowId;
    workflowIds[1] = wfId2;

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
      wfName2,
      s_tag,
      wfId2,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    vm.recordLogs();
    s_registry.batchPauseWorkflows(workflowIds);
    vm.stopPrank();

    // now inspect the logs
    Vm.Log[] memory logs = vm.getRecordedLogs();
    bytes32 sig = keccak256("WorkflowPaused(bytes32,address,string,string)");

    bool sawA = false;
    bool sawB = false;

    for (uint256 i = 0; i < logs.length; i++) {
      if (logs[i].topics[0] != sig) continue;

      // topics[1] is indexed workflowId
      bytes32 wid = logs[i].topics[1];
      if (wid == s_workflowId) sawA = true;
      if (wid == wfId2) sawB = true;
    }

    assertTrue(sawA, "expected WorkflowPaused(s_workflowId) to be emitted");
    assertFalse(sawB, "did not expect WorkflowPaused(wfId2) to be emitted");
  }
}
