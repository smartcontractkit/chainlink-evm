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
}
