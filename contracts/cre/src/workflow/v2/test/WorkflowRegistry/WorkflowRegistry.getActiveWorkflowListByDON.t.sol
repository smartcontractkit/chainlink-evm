// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getActiveWorkflowListByDON is WorkflowRegistrySetup {
  address private s_owner1 = makeAddr("owner1");
  address private s_owner2 = makeAddr("owner2");
  string private s_donFamily1 = "DON-Family-1";
  string private s_donFamily2 = "DON-Family-2";

  function test_getActiveWorkflowListByDON_WhenTheDONFamilyHasNoWorkflowsRegistered() external view {
    // it should return an empty array
    WorkflowRegistry.WorkflowMetadataView[] memory workflows =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows.length, 0, "Expected no workflows");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 1);
    assertEq(workflows.length, 0, "Expected no workflows");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 5, 10);
    assertEq(workflows.length, 0, "Expected no workflows");
  }

  modifier whenTheDONFamilyHasWorkflowsRegistered() {
    // Set up DON limits for both families
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);
    s_registry.setDONLimit(s_donFamily2, 50, 5);
    vm.stopPrank();

    // Link owners
    _linkOwner(s_owner1);
    _linkOwner(s_owner2);

    // Create 3 ACTIVE workflows for DON Family 1
    vm.startPrank(s_owner1);
    _createActiveWorkflowForDON("Active-Workflow-1", "tag1", keccak256("active1"), s_donFamily1);
    _createActiveWorkflowForDON("Active-Workflow-2", "tag2", keccak256("active2"), s_donFamily1);
    vm.stopPrank();

    vm.startPrank(s_owner2);
    _createActiveWorkflowForDON("Active-Workflow-3", "tag3", keccak256("active3"), s_donFamily1);
    vm.stopPrank();

    // Create 2 PAUSED workflows for DON Family 1 (should NOT appear in active list)
    vm.startPrank(s_owner1);
    _createPausedWorkflowForDON("Paused-Workflow-1", "paused1", keccak256("paused1"), s_donFamily1);
    _createPausedWorkflowForDON("Paused-Workflow-2", "paused2", keccak256("paused2"), s_donFamily1);
    vm.stopPrank();

    // Create 2 ACTIVE workflows for DON Family 2 (different DON)
    vm.startPrank(s_owner1);
    _createActiveWorkflowForDON("Other-Active-1", "other1", keccak256("other1"), s_donFamily2);
    _createActiveWorkflowForDON("Other-Active-2", "other2", keccak256("other2"), s_donFamily2);
    vm.stopPrank();
    _;
  }

  function test_getActiveWorkflowListByDON_WhenStartIsGreaterThanOrEqualToTotalActiveWorkflows()
    external
    whenTheDONFamilyHasWorkflowsRegistered
  {
    // it should return an empty array
    WorkflowRegistry.WorkflowMetadataView[] memory workflows =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 3, 10);
    assertEq(workflows.length, 0, "Expected no workflows when start equals total active");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 4, 5);
    assertEq(workflows.length, 0, "Expected no workflows when start is greater than total active");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 10, 1);
    assertEq(workflows.length, 0, "Expected no workflows when start is much greater than total active");
  }

  modifier whenStartIsLessThanTotalActiveWorkflows() {
    _;
  }

  function test_getActiveWorkflowListByDON_WhenLimitIsZero()
    external
    whenTheDONFamilyHasWorkflowsRegistered
    whenStartIsLessThanTotalActiveWorkflows
  {
    // it should return an empty array
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 0);
    assertEq(workflows.length, 0, "Expected no workflows when limit is 0");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 1, 0);
    assertEq(workflows.length, 0, "Expected no workflows when limit is 0");
  }

  function test_getActiveWorkflowListByDON_WhenLimitIsLessThanTotalMinusStart()
    external
    whenTheDONFamilyHasWorkflowsRegistered
    whenStartIsLessThanTotalActiveWorkflows
  {
    // it should return exactly limit workflows starting from start index
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 2);
    assertEq(workflows.length, 2, "Expected exactly 2 workflows");
    assertEq(workflows[0].workflowName, "Active-Workflow-1", "Expected first active workflow");
    assertEq(workflows[1].workflowName, "Active-Workflow-2", "Expected second active workflow");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 1, 1);
    assertEq(workflows.length, 1, "Expected exactly 1 workflow starting from index 1");
    assertEq(workflows[0].workflowName, "Active-Workflow-2", "Expected second active workflow");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 2, 1);
    assertEq(workflows.length, 1, "Expected exactly 1 workflow");
    assertEq(workflows[0].workflowName, "Active-Workflow-3", "Expected third active workflow");
  }

  function test_getActiveWorkflowListByDON_WhenLimitIsGreaterThanOrEqualToTotalMinusStart()
    external
    whenTheDONFamilyHasWorkflowsRegistered
    whenStartIsLessThanTotalActiveWorkflows
  {
    // it should return all active workflows from start index to the end
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 3);
    assertEq(workflows.length, 3, "Expected all 3 active workflows");
    assertEq(workflows[0].workflowName, "Active-Workflow-1", "Expected first active workflow");
    assertEq(workflows[1].workflowName, "Active-Workflow-2", "Expected second active workflow");
    assertEq(workflows[2].workflowName, "Active-Workflow-3", "Expected third active workflow");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows.length, 3, "Expected all 3 active workflows when limit exceeds total");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 1, 10);
    assertEq(workflows.length, 2, "Expected last 2 active workflows");
    assertEq(workflows[0].workflowName, "Active-Workflow-2", "Expected second active workflow");
    assertEq(workflows[1].workflowName, "Active-Workflow-3", "Expected third active workflow");

    workflows = s_registry.getActiveWorkflowListByDON(s_donFamily1, 2, 5);
    assertEq(workflows.length, 1, "Expected last active workflow");
    assertEq(workflows[0].workflowName, "Active-Workflow-3", "Expected third active workflow");
  }

  function test_getActiveWorkflowListByDON_ShouldOnlyReturnActiveWorkflowsFromSpecifiedDON()
    external
    whenTheDONFamilyHasWorkflowsRegistered
  {
    // Verify DON Family 1 active workflows (should be 3, not 5 total)
    WorkflowRegistry.WorkflowMetadataView[] memory workflows1 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows1.length, 3, "Expected 3 ACTIVE workflows for DON Family 1");

    // Verify all returned workflows are ACTIVE
    for (uint256 i = 0; i < workflows1.length; i++) {
      assertEq(
        uint256(workflows1[i].status), uint256(WorkflowRegistry.WorkflowStatus.ACTIVE), "Expected workflow to be ACTIVE"
      );
    }

    // Verify DON Family 2 active workflows
    WorkflowRegistry.WorkflowMetadataView[] memory workflows2 =
      s_registry.getActiveWorkflowListByDON(s_donFamily2, 0, 10);
    assertEq(workflows2.length, 2, "Expected 2 ACTIVE workflows for DON Family 2");
    assertEq(workflows2[0].workflowName, "Other-Active-1", "Expected first other active workflow");
    assertEq(workflows2[1].workflowName, "Other-Active-2", "Expected second other active workflow");

    // Verify workflow names don't overlap
    for (uint256 i = 0; i < workflows1.length; i++) {
      for (uint256 j = 0; j < workflows2.length; j++) {
        assertTrue(
          keccak256(bytes(workflows1[i].workflowName)) != keccak256(bytes(workflows2[j].workflowName)),
          "Active workflows from different DONs should not have the same name"
        );
      }
    }
  }

  function test_getActiveWorkflowListByDON_ShouldOnlyReturnActiveWorkflows()
    external
    whenTheDONFamilyHasWorkflowsRegistered
  {
    // All workflows returned should be ACTIVE only
    WorkflowRegistry.WorkflowMetadataView[] memory workflows =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);

    assertEq(workflows.length, 3, "Expected only 3 ACTIVE workflows (PAUSED ones excluded)");

    for (uint256 i = 0; i < workflows.length; i++) {
      assertEq(
        uint256(workflows[i].status), uint256(WorkflowRegistry.WorkflowStatus.ACTIVE), "Expected workflow to be ACTIVE"
      );
    }
  }

  function test_getActiveWorkflowListByDON_ShouldNotIncludePausedWorkflows() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);
    // Create ACTIVE workflow
    bytes32 activeId = keccak256("active");
    s_registry.upsertWorkflow(
      "Active-Workflow",
      "active",
      activeId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily1,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // Create PAUSED workflow
    _createPausedWorkflowForDON("Paused-Workflow", "paused", keccak256("paused"), s_donFamily1);
    vm.stopPrank();

    WorkflowRegistry.WorkflowMetadataView[] memory workflows =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows.length, 1, "Expected only 1 ACTIVE workflow");
    assertEq(workflows[0].workflowName, "Active-Workflow", "Expected the ACTIVE workflow");
    assertEq(
      uint256(workflows[0].status), uint256(WorkflowRegistry.WorkflowStatus.ACTIVE), "Expected workflow to be ACTIVE"
    );
  }

  function test_getActiveWorkflowListByDON_ShouldExcludeWorkflowThatWasPaused() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);
    // Create ACTIVE workflow
    bytes32 activeId = keccak256("active");
    s_registry.upsertWorkflow(
      "Active-Workflow",
      "active",
      activeId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily1,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // Verify it appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsBefore =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsBefore.length, 1, "Expected 1 ACTIVE workflow before pause");

    // Pause the workflow
    s_registry.pauseWorkflow(activeId);
    vm.stopPrank();

    // Verify it no longer appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsAfter =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsAfter.length, 0, "Expected no ACTIVE workflows after pause");
  }

  function test_getActiveWorkflowListByDON_ShouldIncludeWorkflowThatWasReactivated() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);
    // Create PAUSED workflow
    bytes32 wfId = keccak256("workflow");
    _createPausedWorkflowForDON("Test-Workflow", "test", wfId, s_donFamily1);

    // Verify it doesn't appear in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsBefore =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsBefore.length, 0, "Expected no ACTIVE workflows when paused");

    // Activate the workflow
    s_registry.activateWorkflow(wfId, s_donFamily1);

    // Verify it now appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsAfter =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsAfter.length, 1, "Expected 1 ACTIVE workflow after activation");
    assertEq(workflowsAfter[0].workflowName, "Test-Workflow", "Expected the reactivated workflow");
    assertEq(
      uint256(workflowsAfter[0].status),
      uint256(WorkflowRegistry.WorkflowStatus.ACTIVE),
      "Expected workflow to be ACTIVE"
    );
    vm.stopPrank();
  }

  function test_getActiveWorkflowListByDON_ShouldNotIncludeDeletedWorkflow() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);
    // Create ACTIVE workflow
    bytes32 activeId = keccak256("active");
    s_registry.upsertWorkflow(
      "Active-Workflow",
      "active",
      activeId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily1,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // Verify it appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsBefore =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsBefore.length, 1, "Expected 1 ACTIVE workflow before deletion");
    assertEq(workflowsBefore[0].workflowName, "Active-Workflow", "Expected the ACTIVE workflow");

    // Delete the workflow
    s_registry.deleteWorkflow(activeId);
    vm.stopPrank();

    // Verify it no longer appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsAfter =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsAfter.length, 0, "Expected no ACTIVE workflows after deletion");
  }

  function test_getActiveWorkflowListByDON_ShouldIncludeRecreatedWorkflow() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);
    // Create and delete ACTIVE workflow
    bytes32 wfId = keccak256("workflow");
    s_registry.upsertWorkflow(
      "Test-Workflow",
      "v1",
      wfId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily1,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    s_registry.deleteWorkflow(wfId);

    // Verify list is empty
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsAfterDelete =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsAfterDelete.length, 0, "Expected no workflows after deletion");

    // Recreate with same ID but different tag
    s_registry.upsertWorkflow(
      "Test-Workflow-V2",
      "v2",
      wfId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily1,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
    vm.stopPrank();

    // Verify recreated workflow appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsAfterRecreate =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsAfterRecreate.length, 1, "Expected 1 ACTIVE workflow after recreation");
    assertEq(workflowsAfterRecreate[0].workflowName, "Test-Workflow-V2", "Expected the recreated workflow");
    assertEq(workflowsAfterRecreate[0].tag, "v2", "Expected the new tag");
    assertEq(
      uint256(workflowsAfterRecreate[0].status),
      uint256(WorkflowRegistry.WorkflowStatus.ACTIVE),
      "Expected workflow to be ACTIVE"
    );
  }

  function test_getActiveWorkflowListByDON_ShouldHandleMultipleDeletesAndRecreations() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);

    // Create 3 ACTIVE workflows
    bytes32 wfId1 = keccak256("wf1");
    bytes32 wfId2 = keccak256("wf2");
    bytes32 wfId3 = keccak256("wf3");

    _createActiveWorkflowForDON("Workflow-1", "v1", wfId1, s_donFamily1);
    _createActiveWorkflowForDON("Workflow-2", "v1", wfId2, s_donFamily1);
    _createActiveWorkflowForDON("Workflow-3", "v1", wfId3, s_donFamily1);

    // Verify all 3 appear
    WorkflowRegistry.WorkflowMetadataView[] memory workflows1 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows1.length, 3, "Expected 3 ACTIVE workflows");

    // Delete workflow 2
    s_registry.deleteWorkflow(wfId2);

    // Verify only 2 remain
    WorkflowRegistry.WorkflowMetadataView[] memory workflows2 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows2.length, 2, "Expected 2 ACTIVE workflows after first deletion");

    // Recreate workflow 2
    s_registry.upsertWorkflow(
      "Workflow-2",
      "v1",
      wfId2,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily1,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // Verify all 3 are back
    WorkflowRegistry.WorkflowMetadataView[] memory workflows3 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows3.length, 3, "Expected 3 ACTIVE workflows after recreation");

    // Verify names (order may vary)
    bool found1 = false;
    bool found2Recreated = false;
    bool found3 = false;

    for (uint256 i = 0; i < workflows3.length; i++) {
      if (keccak256(bytes(workflows3[i].workflowName)) == keccak256(bytes("Workflow-1"))) found1 = true;
      if (keccak256(bytes(workflows3[i].workflowName)) == keccak256(bytes("Workflow-2"))) {
        found2Recreated = true;
      }
      if (keccak256(bytes(workflows3[i].workflowName)) == keccak256(bytes("Workflow-3"))) found3 = true;
    }

    assertTrue(found1, "Expected Workflow-1 to be present");
    assertTrue(found2Recreated, "Expected recreated Workflow-2 to be present");
    assertTrue(found3, "Expected Workflow-3 to be present");

    // Delete all and recreate in different order
    s_registry.deleteWorkflow(wfId1);
    s_registry.deleteWorkflow(wfId2);
    s_registry.deleteWorkflow(wfId3);

    WorkflowRegistry.WorkflowMetadataView[] memory workflows4 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows4.length, 0, "Expected no workflows after deleting all");

    // Recreate in reverse order
    _createActiveWorkflowForDON("Workflow-3", "v1", wfId3, s_donFamily1);
    _createActiveWorkflowForDON("Workflow-2", "v1", wfId2, s_donFamily1);
    _createActiveWorkflowForDON("Workflow-1", "v1", wfId1, s_donFamily1);

    WorkflowRegistry.WorkflowMetadataView[] memory workflows5 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows5.length, 3, "Expected 3 ACTIVE workflows after final recreation");

    vm.stopPrank();
  }

  function test_getActiveWorkflowListByDON_DeletedPausedWorkflowShouldNotLeaveGhostEntry() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);

    // Create PAUSED workflow
    bytes32 pausedId = keccak256("paused");
    _createPausedWorkflowForDON("Paused-Workflow", "v1", pausedId, s_donFamily1);

    // Verify it doesn't appear in active list (it's paused)
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsBefore =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsBefore.length, 0, "Expected no ACTIVE workflows (workflow is paused)");

    // Delete the paused workflow
    s_registry.deleteWorkflow(pausedId);

    // Create a new ACTIVE workflow
    bytes32 newActiveId = keccak256("new-active");
    _createActiveWorkflowForDON("New-Active-Workflow", "v1", newActiveId, s_donFamily1);

    vm.stopPrank();

    // Verify only the new active workflow appears (no ghost entries from deleted paused workflow)
    WorkflowRegistry.WorkflowMetadataView[] memory workflowsAfter =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflowsAfter.length, 1, "Expected exactly 1 ACTIVE workflow");
    assertEq(workflowsAfter[0].workflowName, "New-Active-Workflow", "Expected the new active workflow");
    assertEq(workflowsAfter[0].workflowId, newActiveId, "Expected correct workflow ID");
    // Ensure no ghost entries (all returned workflows have valid metadata)
    assertTrue(workflowsAfter[0].owner != address(0), "Expected valid owner (no ghost entry)");
  }

  function test_getActiveWorkflowListByDON_RecreateAsPausedAfterDeletingActive() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 100, 10);

    // Link owner
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);

    // Create ACTIVE workflow
    bytes32 wfId = keccak256("workflow");
    _createActiveWorkflowForDON("Test-Workflow", "v1", wfId, s_donFamily1);

    // Verify it appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflows1 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows1.length, 1, "Expected 1 ACTIVE workflow");

    // Delete it
    s_registry.deleteWorkflow(wfId);

    // Verify list is empty
    WorkflowRegistry.WorkflowMetadataView[] memory workflows2 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows2.length, 0, "Expected no ACTIVE workflows after deletion");

    // Recreate as PAUSED with same ID
    _createPausedWorkflowForDON("Test-Workflow-Paused", "v2", wfId, s_donFamily1);

    vm.stopPrank();

    // Verify it does NOT appear in active list (it's paused)
    WorkflowRegistry.WorkflowMetadataView[] memory workflows3 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows3.length, 0, "Expected no ACTIVE workflows (recreated as paused)");

    // Activate it
    vm.prank(s_owner1);
    s_registry.activateWorkflow(wfId, s_donFamily1);

    // Verify it now appears in active list
    WorkflowRegistry.WorkflowMetadataView[] memory workflows4 =
      s_registry.getActiveWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows4.length, 1, "Expected 1 ACTIVE workflow after activation");
    assertEq(workflows4[0].workflowName, "Test-Workflow-Paused", "Expected the reactivated workflow");
  }

  // Helper functions
  function _createActiveWorkflowForDON(
    string memory workflowName,
    string memory tag,
    bytes32 workflowId,
    string memory donFamily
  ) internal {
    s_registry.upsertWorkflow(
      workflowName,
      tag,
      workflowId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  function _createPausedWorkflowForDON(
    string memory workflowName,
    string memory tag,
    bytes32 workflowId,
    string memory donFamily
  ) internal {
    s_registry.upsertWorkflow(
      workflowName,
      tag,
      workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }
}
