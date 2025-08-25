// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_batchActivateWorkflows is WorkflowRegistrySetup {
  function test_batchActivateWorkflows_WhenCallerIsNOTALinkedOwner() external {
    // it reverts with OwnershipLinkDoesNotExist
    bytes32[] memory workflowIds = new bytes32[](2);
    bytes32 wfId2 = keccak256("workflow-id2");
    workflowIds[0] = s_workflowId;
    workflowIds[1] = wfId2;

    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_user));
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
  }

  modifier whenCallerISALinkedOwner() {
    _linkOwner(s_user);
    _;
  }

  function test_batchActivateWorkflows_WhenWorkflowIdsLengthIs0() external whenCallerISALinkedOwner {
    // it reverts with EmptyUpdateBatch
    bytes32[] memory workflowIds = new bytes32[](0);
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.EmptyUpdateBatch.selector));
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
  }

  function test_batchActivateWorkflows_WhenEveryWorkflowIdIsUnknown() external whenCallerISALinkedOwner {
    // it reverts with WorkflowDoesNotExist
    bytes32[] memory workflowIds = new bytes32[](1);
    workflowIds[0] = s_workflowId;

    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowDoesNotExist.selector, s_workflowId));
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
  }

  function test_batchActivateWorkflows_WhenAtLeastOneWorkflowIdIsNotOwnedByCaller() external whenCallerISALinkedOwner {
    // it reverts with CallerIsNotWorkflowOwner
    // add a workflow
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
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
  }

  function test_batchActivateWorkflows_WhenDONFamilyHasNoGlobalLimitSet() external whenCallerISALinkedOwner {
    // it reverts with DonLimitNotSet
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
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.DonLimitNotSet.selector, s_donFamily));
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
  }

  function test_batchActivateWorkflows_WhenActivationWouldExceedCallersPer_DONCap() external whenCallerISALinkedOwner {
    // it reverts with MaxWorkflowsPerUserDONExceeded
    bytes32[] memory workflowIds = new bytes32[](2);
    bytes32 wfId2 = keccak256("workflow-id2");
    workflowIds[0] = s_workflowId;
    workflowIds[1] = wfId2;

    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 1, true);

    // add some workflows
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
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerUserDONExceeded.selector, s_user, s_donFamily)
    );
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
    vm.stopPrank();
  }

  function test_batchActivateWorkflows_WhenAllListedWorkflowsAreAlreadyACTIVE() external whenCallerISALinkedOwner {
    // it emits no WorkflowActivated events and leaves state unchanged
    _setDONLimit();
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

    bytes32[] memory workflowIds = new bytes32[](1);
    workflowIds[0] = s_workflowId;

    vm.recordLogs();
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
    vm.stopPrank();
    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("WorkflowActivated(bytes32,address,string,string)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("WorkflowActivated was emitted when it should not have been");
        fail();
      }
    }
  }

  function test_batchActivateWorkflows_WhenListMixesACTIVEAndPAUSEDWorkflowsWhereTheOnesToActivateAreWithinCap()
    external
    whenCallerISALinkedOwner
  {
    // it activates each PAUSED workflow and emits a WorkflowActivated event for each
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
      WorkflowRegistry.WorkflowStatus.PAUSED,
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

    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.WorkflowActivated(s_workflowId, s_user, s_donFamily, s_workflowName);
    vm.expectEmit(true, true, true, true);
    emit WorkflowRegistry.WorkflowActivated(wfId2, s_user, s_donFamily, wfName2);
    s_registry.batchActivateWorkflows(workflowIds, s_donFamily);
    vm.stopPrank();
  }
}
