// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_approveOperation is WorkflowRegistrySetup {
  function test_approveOperation_WhenTheUserIsNotLinked() external {
    // it should revert with OwnershipLinkDoesNotExist
    bytes32 operationDigest = keccak256("operation-digest");
    uint256 expiryTimestamp = block.timestamp + 1 hours;

    address vaultNode = address(0x89652);
    vm.prank(vaultNode);
    assertFalse(s_registry.isOperationApproved(s_user, operationDigest), "Operation should not be approved");

    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_user));
    vm.prank(s_user);
    s_registry.approveOperation(operationDigest, expiryTimestamp);
  }

  function test_approveOperation_WhenTheUserIsLinked() external {
    // it should accept the operation digest
    bytes32 operationDigest = keccak256("operation-digest");
    uint256 expiryTimestamp = block.timestamp + 1 hours;

    // link the owner first to ensure the operation can be approved
    _linkOwner(s_user);
    address vaultNode = address(0x89652);
    vm.prank(vaultNode);
    assertFalse(s_registry.isOperationApproved(s_user, operationDigest), "Operation should not be approved");

    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OperationApproved(s_user, operationDigest, expiryTimestamp);
    vm.prank(s_user);
    s_registry.approveOperation(operationDigest, expiryTimestamp);

    vm.prank(vaultNode);
    assertTrue(s_registry.isOperationApproved(s_user, operationDigest), "Operation should be approved");

    bytes32 newOperationDigest = keccak256("new-operation-digest");
    uint256 newExpiryTimestamp = block.timestamp + 1 hours; // same timestamp as the previous operation
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OperationApproved(s_user, newOperationDigest, newExpiryTimestamp);
    vm.prank(s_user);
    s_registry.approveOperation(newOperationDigest, newExpiryTimestamp);

    vm.prank(vaultNode);
    assertTrue(s_registry.isOperationApproved(s_user, newOperationDigest), "New operation should be approved");
    assertTrue(s_registry.isOperationApproved(s_user, operationDigest), "Old operation should still be approved");

    vm.warp(block.timestamp + 1 hours); // Advances the block timestamp by 1 hour only for the next call
    vm.prank(vaultNode);
    assertFalse(s_registry.isOperationApproved(s_user, newOperationDigest), "New operation should expire");
    assertFalse(s_registry.isOperationApproved(s_user, operationDigest), "Old operation should expire");

    newExpiryTimestamp = block.timestamp + 2 hours; // same digest, but one hour ahead of block time
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OperationApproved(s_user, newOperationDigest, newExpiryTimestamp);
    vm.prank(s_user);
    s_registry.approveOperation(newOperationDigest, newExpiryTimestamp);

    vm.prank(vaultNode);
    assertFalse(s_registry.isOperationApproved(s_user, operationDigest), "Old operation should be expired");
    assertTrue(s_registry.isOperationApproved(s_user, newOperationDigest), "New operation should be approved");
  }
}
