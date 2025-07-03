// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_pauseWorkflow is WorkflowRegistrySetup {
  function test_pauseWorkflow_WhenCallerIsNotLinkedAsAnOwner() external {
    // It reverts with OwnershipLinkDoesNotExist
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_owner));
    s_registry.pauseWorkflow(s_workflowId);
  }

  modifier whenCallerIsLinked() {
    _;
  }

  function test_pauseWorkflow_WhenNoWorkflowExistsForTheGivenWorkflowId() external whenCallerIsLinked {
    // It reverts with WorkflowDoesNotExist
    _linkOwner(s_owner);
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowDoesNotExist.selector, s_workflowId));
    s_registry.pauseWorkflow(s_workflowId);
  }

  function test_pauseWorkflow_WhenTheWorkflowExistsButOwnerIsNotCaller() external whenCallerIsLinked {
    // It reverts with CallerIsNotWorkflowOwner
    // set DON limit first
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);

    address user2 = makeAddr("user2");
    _linkOwner(s_user);
    _linkOwner(user2);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);
    vm.prank(user2);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.CallerIsNotWorkflowOwner.selector, user2));
    s_registry.pauseWorkflow(s_workflowId);
  }

  function test_pauseWorkflow_WhenTheWorkflowExistsOwnerMatchesButStatusIsPAUSED() external whenCallerIsLinked {
    // It returns immediately (no state change, no event)
    // set DON limit first
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);
    _linkOwner(s_user);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.PAUSED, false, s_user);
    vm.prank(s_user);

    // should not emit any logs since the workflow is already paused
    vm.recordLogs();
    s_registry.pauseWorkflow(s_workflowId);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("WorkflowPaused(bytes32,address,string,string)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("WorkflowPaused was emitted when it should not have been");
        fail();
      }
    }
  }

  function test_pauseWorkflow_WhenTheWorkflowExistsOwnerMatchesAndStatusIsACTIVE() external whenCallerIsLinked {
    // It calls pauses the workflow and emits WorkflowPaused
    // set DON limit first
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);
    _linkOwner(s_user);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);
    vm.prank(s_user);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.WorkflowPaused(s_workflowId, s_user, s_donFamily, s_workflowName);
    s_registry.pauseWorkflow(s_workflowId);
  }
}
