// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistrygetWorkflowMetadataByOwnerAndNameAndTag is WorkflowRegistrySetup {
  function test_getWorkflowMetadataByOwnerAndNameAndTag_WhenAWorkflowExistsForTheSpecifiedOwnerAndNameAndTag() external {
    // It should return the expected WorkflowMetadataView
    _linkOwner(s_owner);
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.PAUSED, false, s_owner);
    WorkflowRegistry.WorkflowMetadataView memory wf =
      s_registry.getWorkflowMetadataByOwnerAndNameAndTag(s_owner, s_workflowName, s_tag);
    assertEq(wf.workflowName, s_workflowName);
    assertEq(wf.workflowId, s_workflowId);
    assertEq(wf.owner, s_owner);
    assertEq(wf.tag, s_tag);
  }

  function test_getWorkflowMetadataByOwnerAndNameAndTag_WhenNoWorkflowExistsForTheSpecifiedOwnerAndNameAndTag()
    external
  {
    // It should revert with WorkflowDoesNotExist
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowDoesNotExist.selector));
    s_registry.getWorkflowMetadataByOwnerAndNameAndTag(s_owner, s_workflowName, s_tag);
  }
}
