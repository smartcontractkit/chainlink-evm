// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getWorkflow is WorkflowRegistrySetup {
  function test_getWorkflow_WhenTheWorkflowExistsWithTheOwnerAndNameAndTag() external {
    // it returns the correct metadata
    _setDONLimit();
    _linkOwner(s_owner);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, true, s_owner);

    // try to fetch by workfowId first
    WorkflowRegistry.WorkflowMetadataView memory metadata = s_registry.getWorkflowById(s_workflowId);
    assertEq(metadata.owner, s_owner, "Expected owner to match");

    // try to fetch by owner, workflowName and tag next
    metadata = s_registry.getWorkflow(s_owner, s_workflowName, s_tag);
    assertEq(metadata.workflowId, s_workflowId, "Expected workflowId to match");
  }

  function test_getWorkflow_WhenTheWorkflowDoesNotExist() external view {
    // it returns nothing with no errors

    // try to fetch by workfowId first
    s_registry.getWorkflowById(s_workflowId);

    // try to fetch by owner, workflowName and tag next
    s_registry.getWorkflow(s_owner, s_workflowName, s_tag);
  }
}
