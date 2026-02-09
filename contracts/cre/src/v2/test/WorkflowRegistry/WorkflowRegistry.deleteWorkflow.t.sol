// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistrydeleteWorkflow is WorkflowRegistrySetup {
  function test_WhenTheOwnerIsNotLinked() external {
    // It should revert with OwnershipLinkDoesNotExist
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_owner));
    s_registry.deleteWorkflow(s_workflowId);
  }

  modifier whenTheOwnerIsLinked() {
    _linkOwner(s_owner);
    _;
  }

  function test_WhenTheWorkflowDoesNotExist() external whenTheOwnerIsLinked {
    // It should revert with WorkflowDoesNotExist
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowDoesNotExist.selector, s_workflowId));
    s_registry.deleteWorkflow(s_workflowId);
  }

  modifier whenTheWorkflowExists() {
    _;
  }

  function test_WhenCallerIsNotTheOwner() external whenTheOwnerIsLinked whenTheWorkflowExists {
    // It should revert with CallerIsNotWorkflowOwner
    _linkOwner(s_user);
    vm.prank(s_owner);
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
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.CallerIsNotWorkflowOwner.selector, s_user));
    s_registry.deleteWorkflow(s_workflowId);
  }

  function test_WhenCallerIsTheOwner() external whenTheOwnerIsLinked whenTheWorkflowExists {
    // It should delete the workflow and emit WorkflowDeleted
    vm.startPrank(s_owner);
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
    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_owner, 0, 100);
    assertEq(wrs.length, 1, "There should be 1 workflow for the s_owner");

    s_registry.deleteWorkflow(s_workflowId);
    vm.stopPrank();

    wrs = s_registry.getWorkflowListByOwner(s_owner, 0, 100);
    assertEq(wrs.length, 0, "There should be 0 workflows for the s_owner");
  }

  // deleting a PAUSED workflow should not leave any stale entries
  function test_deleteWorkflow_pausedWorkflowShouldNotLeaveAnyStaleEntries()
    external
    whenTheOwnerIsLinked
    whenTheWorkflowExists
  {
    string memory donFamily = "test-don";
    vm.prank(s_owner);
    s_registry.setDONLimit(donFamily, 10, 5);

    string memory wfName = "paused-workflow";
    string memory tag = "v1";
    bytes32 wfId = keccak256("paused-wf-id");

    vm.prank(s_owner);
    s_registry.upsertWorkflow(
      wfName,
      tag,
      wfId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    WorkflowRegistry.WorkflowMetadataView[] memory beforeList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(beforeList.length, 1, "Should have 1 workflow before delete");
    assertEq(beforeList[0].workflowId, wfId, "Workflow ID mismatch");
    assertEq(beforeList[0].owner, s_owner, "Owner mismatch");

    vm.prank(s_owner);
    s_registry.deleteWorkflow(wfId);

    WorkflowRegistry.WorkflowMetadataView[] memory afterList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(afterList.length, 0, "DON list should be empty after deleting the only workflow");

    // Recreate deleted workflow to verify DON index is clean
    vm.prank(s_owner);
    s_registry.upsertWorkflow(
      wfName,
      tag,
      wfId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    WorkflowRegistry.WorkflowMetadataView[] memory wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 1, "Should have exactly 1 new workflow");
    assertEq(wfList[0].workflowId, wfId, "New workflow ID should match");
    assertEq(wfList[0].workflowName, wfName, "New workflow name should match");
    assertEq(wfList[0].owner, s_owner, "New workflow owner should match");
    // Verify no ghost entries (all returned workflows have valid metadata)
    assertTrue(wfList[0].owner != address(0), "New workflow should have valid owner");

    vm.prank(s_owner);
    s_registry.pauseWorkflow(wfId);
    wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 1, "Should have exactly 1 workflow");
    assertTrue(uint8(wfList[0].status) == uint8(WorkflowRegistry.WorkflowStatus.PAUSED), "Status should be PAUSED");
    assertEq(wfList[0].donFamily, donFamily, "DON family should match the updated value");
    assertEq(wfList[0].workflowId, wfId, "Workflow ID should match the updated value");
    assertEq(wfList[0].owner, s_owner, "Owner should match the updated value");
    assertEq(wfList[0].workflowName, wfName, "Workflow name should match the updated value");

    vm.prank(s_owner);
    s_registry.activateWorkflow(wfId, donFamily);
    wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 1, "Should have exactly 1 workflow");
    assertTrue(uint8(wfList[0].status) == uint8(WorkflowRegistry.WorkflowStatus.ACTIVE), "Status should be ACTIVE");
    assertEq(wfList[0].donFamily, donFamily, "DON family should match the updated value");
    assertEq(wfList[0].workflowId, wfId, "Workflow ID should match the updated value");
    assertEq(wfList[0].owner, s_owner, "Owner should match the updated value");
    assertEq(wfList[0].workflowName, wfName, "Workflow name should match the updated value");

    vm.prank(s_owner);
    s_registry.pauseWorkflow(wfId);

    vm.prank(s_owner);
    string memory changeDonFamily = "different-test-don";
    s_registry.setDONLimit(changeDonFamily, 10, 5);

    vm.prank(s_owner);
    s_registry.activateWorkflow(wfId, changeDonFamily);
    wfList = s_registry.getWorkflowListByDON(changeDonFamily, 0, 10);
    assertEq(wfList.length, 1, "Should have exactly 1 workflow");
    assertTrue(uint8(wfList[0].status) == uint8(WorkflowRegistry.WorkflowStatus.ACTIVE), "Status should be ACTIVE");
    assertEq(wfList[0].donFamily, changeDonFamily, "DON family should match the updated value");
    assertEq(wfList[0].workflowId, wfId, "Workflow ID should match the updated value");
    assertEq(wfList[0].owner, s_owner, "Owner should match the updated value");
    assertEq(wfList[0].workflowName, wfName, "Workflow name should match the updated value");

    wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 0, "Original DON family list should be empty");

    vm.prank(s_owner);
    s_registry.deleteWorkflow(wfId);

    wfList = s_registry.getWorkflowListByDON(changeDonFamily, 0, 10);
    assertEq(wfList.length, 0, "DON list should be empty after deletion");

    wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 0, "Owner list should be empty after deletion");
  }

  // deleting ACTIVE workflow turned into PAUSED should not leave any stale entries
  function test_deleteWorkflow_activeThenPausedWorkflowShouldNotLeaveAnyStaleEntries()
    external
    whenTheOwnerIsLinked
    whenTheWorkflowExists
  {
    string memory donFamily = "test-don-2";
    vm.prank(s_owner);
    s_registry.setDONLimit(donFamily, 10, 5);

    string memory wfName = "active-then-paused";
    string memory tag = "v1";
    bytes32 wfId = keccak256("active-wf-id");

    vm.prank(s_owner);
    s_registry.upsertWorkflow(
      wfName,
      tag,
      wfId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    vm.prank(s_owner);
    s_registry.pauseWorkflow(wfId);

    WorkflowRegistry.WorkflowMetadataView[] memory beforeList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(beforeList.length, 1, "Should have 1 workflow before delete");
    assertTrue(uint8(beforeList[0].status) == uint8(WorkflowRegistry.WorkflowStatus.PAUSED), "Status should be PAUSED");

    vm.prank(s_owner);
    s_registry.deleteWorkflow(wfId);

    WorkflowRegistry.WorkflowMetadataView[] memory afterList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(afterList.length, 0, "DON list should be empty after deletion");

    // Recreate deleted workflow to verify DON index is clean
    vm.prank(s_owner);
    s_registry.upsertWorkflow(
      wfName,
      tag,
      wfId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    WorkflowRegistry.WorkflowMetadataView[] memory wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 1, "Should have exactly 1 new workflow");
    assertEq(wfList[0].workflowId, wfId, "New workflow ID should match");
    assertEq(wfList[0].workflowName, wfName, "New workflow name should match");
    assertEq(wfList[0].owner, s_owner, "New workflow owner should match");
    // Verify no ghost entries (all returned workflows have valid metadata)
    assertTrue(wfList[0].owner != address(0), "New workflow should have valid owner");

    vm.prank(s_owner);
    s_registry.activateWorkflow(wfId, donFamily);
    wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 1, "Should have exactly 1 workflow");
    assertTrue(uint8(wfList[0].status) == uint8(WorkflowRegistry.WorkflowStatus.ACTIVE), "Status should be ACTIVE");
    assertEq(wfList[0].donFamily, donFamily, "DON family should match the updated value");
    assertEq(wfList[0].workflowId, wfId, "Workflow ID should match the updated value");
    assertEq(wfList[0].owner, s_owner, "Owner should match the updated value");
    assertEq(wfList[0].workflowName, wfName, "Workflow name should match the updated value");

    vm.prank(s_owner);
    string memory changeDonFamily = "different-test-don";
    s_registry.setDONLimit(changeDonFamily, 10, 5);

    vm.prank(s_owner);
    s_registry.updateWorkflowDONFamily(wfId, changeDonFamily);
    wfList = s_registry.getWorkflowListByDON(changeDonFamily, 0, 10);
    assertEq(wfList.length, 1, "Should have exactly 1 workflow");
    assertTrue(uint8(wfList[0].status) == uint8(WorkflowRegistry.WorkflowStatus.ACTIVE), "Status should be ACTIVE");
    assertEq(wfList[0].donFamily, changeDonFamily, "DON family should match the updated value");
    assertEq(wfList[0].workflowId, wfId, "Workflow ID should match the updated value");
    assertEq(wfList[0].owner, s_owner, "Owner should match the updated value");
    assertEq(wfList[0].workflowName, wfName, "Workflow name should match the updated value");

    wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 0, "Original DON family list should be empty");

    vm.prank(s_owner);
    s_registry.deleteWorkflow(wfId);

    wfList = s_registry.getWorkflowListByDON(changeDonFamily, 0, 10);
    assertEq(wfList.length, 0, "DON list should be empty after deletion");

    wfList = s_registry.getWorkflowListByDON(donFamily, 0, 10);
    assertEq(wfList.length, 0, "Owner list should be empty after deletion");
  }
}
