// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getWorkflowListByDON is WorkflowRegistrySetup {
  address private s_owner1 = makeAddr("owner1");
  address private s_owner2 = makeAddr("owner2");
  string private s_donFamily1 = "DON-Family-1";
  string private s_donFamily2 = "DON-Family-2";

  function test_getWorkflowListByDON_WhenTheDONFamilyHasNoWorkflowsRegistered() external view {
    // it should return an empty array
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows.length, 0, "Expected no workflows");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 1);
    assertEq(workflows.length, 0, "Expected no workflows");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 5, 10);
    assertEq(workflows.length, 0, "Expected no workflows");
  }

  modifier whenTheDONFamilyHasWorkflowsRegistered() {
    // Set up DON limits for both families
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 10, true);
    s_registry.setDONLimit(s_donFamily2, 5, true);
    vm.stopPrank();

    // Link owners and create workflows for DON Family 1
    _linkOwner(s_owner1);
    _linkOwner(s_owner2);

    // Create 5 workflows for DON Family 1
    vm.startPrank(s_owner1);
    _createWorkflowForDON("Workflow-1", "tag1", keccak256("id1"), s_donFamily1);
    _createWorkflowForDON("Workflow-2", "tag2", keccak256("id2"), s_donFamily1);
    _createWorkflowForDON("Workflow-3", "tag3", keccak256("id3"), s_donFamily1);
    vm.stopPrank();

    vm.startPrank(s_owner2);
    _createWorkflowForDON("Workflow-4", "tag4", keccak256("id4"), s_donFamily1);
    _createWorkflowForDON("Workflow-5", "tag5", keccak256("id5"), s_donFamily1);
    vm.stopPrank();

    // Create 2 workflows for DON Family 2 (different DON)
    vm.startPrank(s_owner1);
    _createWorkflowForDON("Other-Workflow-1", "other1", keccak256("other1"), s_donFamily2);
    _createWorkflowForDON("Other-Workflow-2", "other2", keccak256("other2"), s_donFamily2);
    vm.stopPrank();
    _;
  }

  function test_getWorkflowListByDON_WhenStartIsGreaterThanOrEqualToTotalWorkflows()
    external
    whenTheDONFamilyHasWorkflowsRegistered
  {
    // it should return an empty array
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getWorkflowListByDON(s_donFamily1, 5, 10);
    assertEq(workflows.length, 0, "Expected no workflows when start equals total");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 6, 5);
    assertEq(workflows.length, 0, "Expected no workflows when start is greater than total");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 10, 1);
    assertEq(workflows.length, 0, "Expected no workflows when start is much greater than total");
  }

  modifier whenStartIsLessThanTotalWorkflows() {
    _;
  }

  function test_getWorkflowListByDON_WhenLimitIsZero()
    external
    whenTheDONFamilyHasWorkflowsRegistered
    whenStartIsLessThanTotalWorkflows
  {
    // it should return an empty array
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 0);
    assertEq(workflows.length, 0, "Expected no workflows when limit is 0");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 2, 0);
    assertEq(workflows.length, 0, "Expected no workflows when limit is 0");
  }

  function test_getWorkflowListByDON_WhenLimitIsLessThanTotalMinusStart()
    external
    whenTheDONFamilyHasWorkflowsRegistered
    whenStartIsLessThanTotalWorkflows
  {
    // it should return exactly limit workflows starting from start index
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 2);
    assertEq(workflows.length, 2, "Expected exactly 2 workflows");
    assertEq(workflows[0].workflowName, "Workflow-1", "Expected first workflow");
    assertEq(workflows[1].workflowName, "Workflow-2", "Expected second workflow");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 1, 2);
    assertEq(workflows.length, 2, "Expected exactly 2 workflows starting from index 1");
    assertEq(workflows[0].workflowName, "Workflow-2", "Expected second workflow at index 0");
    assertEq(workflows[1].workflowName, "Workflow-3", "Expected third workflow at index 1");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 3, 1);
    assertEq(workflows.length, 1, "Expected exactly 1 workflow");
    assertEq(workflows[0].workflowName, "Workflow-4", "Expected fourth workflow");
  }

  function test_getWorkflowListByDON_WhenLimitIsGreaterThanOrEqualToTotalMinusStart()
    external
    whenTheDONFamilyHasWorkflowsRegistered
    whenStartIsLessThanTotalWorkflows
  {
    // it should return all workflows from start index to the end
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 5);
    assertEq(workflows.length, 5, "Expected all 5 workflows");
    assertEq(workflows[0].workflowName, "Workflow-1", "Expected first workflow");
    assertEq(workflows[1].workflowName, "Workflow-2", "Expected second workflow");
    assertEq(workflows[2].workflowName, "Workflow-3", "Expected third workflow");
    assertEq(workflows[3].workflowName, "Workflow-4", "Expected fourth workflow");
    assertEq(workflows[4].workflowName, "Workflow-5", "Expected fifth workflow");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows.length, 5, "Expected all 5 workflows when limit exceeds total");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 2, 10);
    assertEq(workflows.length, 3, "Expected last 3 workflows");
    assertEq(workflows[0].workflowName, "Workflow-3", "Expected third workflow");
    assertEq(workflows[1].workflowName, "Workflow-4", "Expected fourth workflow");
    assertEq(workflows[2].workflowName, "Workflow-5", "Expected fifth workflow");

    workflows = s_registry.getWorkflowListByDON(s_donFamily1, 4, 5);
    assertEq(workflows.length, 1, "Expected last workflow");
    assertEq(workflows[0].workflowName, "Workflow-5", "Expected fifth workflow");
  }

  function test_getWorkflowListByDON_ShouldOnlyReturnWorkflowsFromSpecifiedDON()
    external
    whenTheDONFamilyHasWorkflowsRegistered
  {
    // Verify DON Family 1 workflows
    WorkflowRegistry.WorkflowMetadataView[] memory workflows1 = s_registry.getWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows1.length, 5, "Expected 5 workflows for DON Family 1");

    // Verify DON Family 2 workflows
    WorkflowRegistry.WorkflowMetadataView[] memory workflows2 = s_registry.getWorkflowListByDON(s_donFamily2, 0, 10);
    assertEq(workflows2.length, 2, "Expected 2 workflows for DON Family 2");
    assertEq(workflows2[0].workflowName, "Other-Workflow-1", "Expected first other workflow");
    assertEq(workflows2[1].workflowName, "Other-Workflow-2", "Expected second other workflow");

    // Verify workflow names don't overlap
    for (uint256 i = 0; i < workflows1.length; i++) {
      for (uint256 j = 0; j < workflows2.length; j++) {
        assertTrue(
          keccak256(bytes(workflows1[i].workflowName)) != keccak256(bytes(workflows2[j].workflowName)),
          "Workflows from different DONs should not have the same name"
        );
      }
    }
  }

  function test_getWorkflowListByDON_ShouldReturnWorkflowsWithCorrectStatus()
    external
    whenTheDONFamilyHasWorkflowsRegistered
  {
    // All workflows should be PAUSED by default from our setup
    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 10);

    for (uint256 i = 0; i < workflows.length; i++) {
      assertEq(
        uint256(workflows[i].status), uint256(WorkflowRegistry.WorkflowStatus.PAUSED), "Expected workflow to be PAUSED"
      );
    }
  }

  function test_getWorkflowListByDON_ShouldIncludeBothActiveAndPausedWorkflows() external {
    // Set up DON limit
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily1, 10, true);

    // Link owner and create workflows
    _linkOwner(s_owner1);

    vm.startPrank(s_owner1);
    // Create PAUSED workflow
    _createWorkflowForDON("Paused-Workflow", "paused", keccak256("paused"), s_donFamily1);

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
    vm.stopPrank();

    WorkflowRegistry.WorkflowMetadataView[] memory workflows = s_registry.getWorkflowListByDON(s_donFamily1, 0, 10);
    assertEq(workflows.length, 2, "Expected 2 workflows (ACTIVE and PAUSED)");

    // Verify we have both statuses
    bool hasPaused = false;
    bool hasActive = false;
    for (uint256 i = 0; i < workflows.length; i++) {
      if (workflows[i].status == WorkflowRegistry.WorkflowStatus.PAUSED) {
        hasPaused = true;
      }
      if (workflows[i].status == WorkflowRegistry.WorkflowStatus.ACTIVE) {
        hasActive = true;
      }
    }
    assertTrue(hasPaused, "Expected to find PAUSED workflow");
    assertTrue(hasActive, "Expected to find ACTIVE workflow");
  }

  // Helper function to create a workflow for a specific DON
  function _createWorkflowForDON(
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
