// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_activateWorkflow is WorkflowRegistrySetup {
  function test_activateWorkflow_WhenCallerIsNotLinkedAsAnOwner() external {
    // It reverts with OwnershipLinkDoesNotExist
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_owner));
    s_registry.activateWorkflow(s_workflowId, s_donFamily);
  }

  modifier whenCallerIsLinked() {
    _;
  }

  function test_activateWorkflow_WhenNoWorkflowExistsForTheGivenWorkflowId() external whenCallerIsLinked {
    // It reverts with WorkflowDoesNotExist
    _linkOwner(s_owner);
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowDoesNotExist.selector, s_workflowId));
    s_registry.activateWorkflow(s_workflowId, s_donFamily);
  }

  function test_activateWorkflow_WhenTheWorkflowExistsButOwnerDoesNotEqualCaller() external whenCallerIsLinked {
    // It reverts with CallerIsNotWorkflowOwner
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);

    address user2 = makeAddr("user2");
    _linkOwner(s_user);
    _linkOwner(user2);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);
    vm.prank(user2);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.CallerIsNotWorkflowOwner.selector, user2));
    s_registry.activateWorkflow(s_workflowId, s_donFamily);
  }

  function test_activateWorkflow_WhenTheWorkflowExistsOwnerMatchesButStatusIsACTIVE() external whenCallerIsLinked {
    // It returns immediately (no state change, no event)
    // set DON limit first
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, true);
    _linkOwner(s_user);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);
    vm.prank(s_user);

    // should not emit any logs since the workflow is already paused
    vm.recordLogs();
    s_registry.activateWorkflow(s_workflowId, s_donFamily);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("WorkflowActivated(bytes32,address,string,string)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("WorkflowActivated was emitted when it should not have been");
        fail();
      }
    }
  }

  modifier whenTheWorkflowExistsOwnerMatchesAndStatusIsPAUSED() {
    _;
  }

  function test_activateWorkflow_WhenThereAreAlreadyTooManyWorkflowsInTheDON()
    external
    whenCallerIsLinked
    whenTheWorkflowExistsOwnerMatchesAndStatusIsPAUSED
  {
    // It reverts with MaxWorkflowsPerUserDONExceeded
    bytes32 wfId2 = keccak256("workflow-id2");
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 1, true);
    _linkOwner(s_user);

    // add 2 worflows 1 active and 1 paused in a don with a limit of 1
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

    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerUserDONExceeded.selector, s_user, s_donFamily)
    );
    s_registry.activateWorkflow(s_workflowId, s_donFamily);
    vm.stopPrank();
  }

  function test_activateWorkflow_WhenNoDONLimitIsSetGloballyForTheDonFamily()
    external
    whenCallerIsLinked
    whenTheWorkflowExistsOwnerMatchesAndStatusIsPAUSED
  {
    // It reverts with DonLimitNotSet
    _linkOwner(s_user);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.PAUSED, false, s_user);
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.DonLimitNotSet.selector, s_donFamily));
    s_registry.activateWorkflow(s_workflowId, s_donFamily);
  }

  function test_activateWorkflow_WhenThereIsEnoughSpaceForTheWorkflowInTheDON()
    external
    whenCallerIsLinked
    whenTheWorkflowExistsOwnerMatchesAndStatusIsPAUSED
  {
    // It activates the workflow and emits WorkflowActivated
    bytes32 wfId2 = keccak256("workflow-id2");
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 2, true);
    _linkOwner(s_user);

    // add 2 worflows 1 active and 1 paused in a don with a limit of 1
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

    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.WorkflowActivated(s_workflowId, s_user, s_donFamily, s_workflowName);
    s_registry.activateWorkflow(s_workflowId, s_donFamily);
    vm.stopPrank();
  }
}
