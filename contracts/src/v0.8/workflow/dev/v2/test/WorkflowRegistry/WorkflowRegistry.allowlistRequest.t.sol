// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_allowlistRequest is WorkflowRegistrySetup {
  function test_allowlistRequest_WhenTheUserIsNotLinked() external {
    // it should revert with OwnershipLinkDoesNotExist
    bytes32 requestDigest = keccak256("request-digest");
    uint32 expiryTimestamp = uint32(block.timestamp + 1 hours);

    address vaultNode = address(0x89652);
    vm.prank(vaultNode);
    assertFalse(s_registry.isRequestAllowlisted(s_user, requestDigest), "Request should not be allowlisted");

    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_user));
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // old timestamp should revert
    expiryTimestamp = uint32(block.timestamp - 1 hours);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidExpiryTimestamp.selector,
        requestDigest,
        expiryTimestamp,
        s_registry.getConfig().maxExpiryLen
      )
    );
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // timestamp equal to current block timestamp should revert
    expiryTimestamp = uint32(block.timestamp);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidExpiryTimestamp.selector,
        requestDigest,
        expiryTimestamp,
        s_registry.getConfig().maxExpiryLen
      )
    );
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
  }

  // When the user is linked
  function test_allowlistRequest_WhenTheUserAlreadyHasARequest() external {
    // It should update the existing request in-place without growing the array
    bytes32 requestDigest = keccak256("duplicate-test-request");
    uint32 initialExpiry = uint32(block.timestamp + 1 hours);
    uint32 updatedExpiry = uint32(block.timestamp + 2 hours);

    // Link the owner first
    _linkOwner(s_user);

    // Initial allowlist
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.RequestAllowlisted(s_user, requestDigest, initialExpiry);
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest, initialExpiry);

    // Verify the request is allowlisted with initial expiry
    assertTrue(s_registry.isRequestAllowlisted(s_user, requestDigest), "Request should be allowlisted");
    assertEq(s_registry.totalAllowlistedRequests(), 1, "Should have exactly 1 request in storage");

    // Get the request details to verify initial expiry
    WorkflowRegistry.OwnerAllowlistedRequest[] memory requests = s_registry.getAllowlistedRequests(0, 10);
    assertEq(requests.length, 1, "Should return exactly 1 request");
    assertEq(requests[0].expiryTimestamp, initialExpiry, "Initial expiry should match");
    assertEq(requests[0].owner, s_user, "Owner should match");
    assertEq(requests[0].requestDigest, requestDigest, "Request digest should match");

    // Update the same request with new expiry (this should update in-place, not add new entry)
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.RequestAllowlisted(s_user, requestDigest, updatedExpiry);
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest, updatedExpiry);

    // Verify the request is still allowlisted but with updated expiry
    assertTrue(s_registry.isRequestAllowlisted(s_user, requestDigest), "Request should still be allowlisted");
    assertEq(s_registry.totalAllowlistedRequests(), 1, "Should still have exactly 1 request in storage (no duplicates)");

    // Get the updated request details
    requests = s_registry.getAllowlistedRequests(0, 10);
    assertEq(requests.length, 1, "Should still return exactly 1 request");
    assertEq(requests[0].expiryTimestamp, updatedExpiry, "Expiry should be updated");
    assertEq(requests[0].owner, s_user, "Owner should still match");
    assertEq(requests[0].requestDigest, requestDigest, "Request digest should still match");

    // Add multiple different requests to verify they are stored separately
    bytes32 requestDigest2 = keccak256("different-request-1");
    bytes32 requestDigest3 = keccak256("different-request-2");
    uint32 expiry2 = uint32(block.timestamp + 3 hours);
    uint32 expiry3 = uint32(block.timestamp + 4 hours);

    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest2, expiry2);
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest3, expiry3);

    // Verify all 3 unique requests are stored
    assertEq(s_registry.totalAllowlistedRequests(), 3, "Should have exactly 3 unique requests");
    requests = s_registry.getAllowlistedRequests(0, 10);
    assertEq(requests.length, 3, "Should return exactly 3 requests");

    // Update the second request and verify no duplicates
    uint32 newExpiry2 = uint32(block.timestamp + 5 hours);
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest2, newExpiry2);

    // Should still have exactly 3 unique requests
    assertEq(s_registry.totalAllowlistedRequests(), 3, "Should still have exactly 3 unique requests");
    requests = s_registry.getAllowlistedRequests(0, 10);
    assertEq(requests.length, 3, "Should still return exactly 3 requests");

    // Find and verify the updated request
    bool foundUpdatedRequest = false;
    for (uint256 i = 0; i < requests.length; i++) {
      if (requests[i].requestDigest == requestDigest2) {
        assertEq(requests[i].expiryTimestamp, newExpiry2, "Second request expiry should be updated");
        foundUpdatedRequest = true;
        break;
      }
    }
    assertTrue(foundUpdatedRequest, "Should find the updated second request");
  }

  // When the user is linked
  function test_allowlistRequest_WhenTheUserHasNoExistingRequest() external {
    // It should allowlist the request digest with a new one
    bytes32 requestDigest = keccak256("request-digest");
    uint32 expiryTimestamp = uint32(block.timestamp + 1 hours);

    // link the owner first to ensure the request can be allowlisted
    _linkOwner(s_user);
    address vaultNode = address(0x89652);
    vm.prank(vaultNode);
    assertFalse(s_registry.isRequestAllowlisted(s_user, requestDigest), "Request should not be allowlisted");

    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.RequestAllowlisted(s_user, requestDigest, expiryTimestamp);
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    vm.prank(vaultNode);
    assertTrue(s_registry.isRequestAllowlisted(s_user, requestDigest), "Request should be allowlisted");

    bytes32 newRequestDigest = keccak256("new-request-digest");
    uint32 newExpiryTimestamp = uint32(block.timestamp + 1 hours); // same timestamp as the previous request
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.RequestAllowlisted(s_user, newRequestDigest, newExpiryTimestamp);
    vm.prank(s_user);
    s_registry.allowlistRequest(newRequestDigest, newExpiryTimestamp);

    vm.prank(vaultNode);
    assertTrue(s_registry.isRequestAllowlisted(s_user, newRequestDigest), "New request should be allowlisted");
    assertTrue(s_registry.isRequestAllowlisted(s_user, requestDigest), "Old request should still be allowlisted");

    vm.warp(block.timestamp + 1 hours); // Advances the block timestamp by 1 hour only for the next call
    vm.prank(vaultNode);
    assertFalse(s_registry.isRequestAllowlisted(s_user, newRequestDigest), "New request should expire");
    assertFalse(s_registry.isRequestAllowlisted(s_user, requestDigest), "Old request should expire");

    newExpiryTimestamp = uint32(block.timestamp + 2 hours); // same digest, but one hour ahead of block time
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.RequestAllowlisted(s_user, newRequestDigest, newExpiryTimestamp);
    vm.prank(s_user);
    s_registry.allowlistRequest(newRequestDigest, newExpiryTimestamp);

    vm.prank(vaultNode);
    assertFalse(s_registry.isRequestAllowlisted(s_user, requestDigest), "Old request should be expired");
    assertTrue(s_registry.isRequestAllowlisted(s_user, newRequestDigest), "New request should be allowlisted");

    // revert if expiration timestamp is much greater than maxAllowedExpiry
    newRequestDigest = keccak256("new-request-digest-2");
    newExpiryTimestamp = uint32(block.timestamp + 8 days); // much more than maxAllowedExpiry
    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidExpiryTimestamp.selector,
        newRequestDigest,
        newExpiryTimestamp,
        s_registry.getConfig().maxExpiryLen
      )
    );
    s_registry.allowlistRequest(newRequestDigest, newExpiryTimestamp);

    // don't revert if expiration time is equal to maxAllowedExpiry
    newRequestDigest = keccak256("new-request-digest-2");
    uint32 maxExpiry = s_registry.getConfig().maxExpiryLen;
    newExpiryTimestamp = uint32(block.timestamp + maxExpiry);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.RequestAllowlisted(s_user, newRequestDigest, newExpiryTimestamp);
    vm.prank(s_user);
    s_registry.allowlistRequest(newRequestDigest, newExpiryTimestamp);

    // don't revert if maxAllowedExpiry is set to unlimited
    WorkflowRegistry.Config memory config = s_registry.getConfig();
    vm.prank(s_owner);
    // set only the maxAllowedExpiry to unlimited
    s_registry.setConfig(config.maxNameLen, config.maxTagLen, config.maxUrlLen, config.maxAttrLen, 0);
    newRequestDigest = keccak256("new-request-digest-3");
    newExpiryTimestamp = uint32(block.timestamp + 8 days); // much more than default maxAllowedExpiry
    vm.prank(s_user);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.RequestAllowlisted(s_user, newRequestDigest, newExpiryTimestamp);
    s_registry.allowlistRequest(newRequestDigest, newExpiryTimestamp);
  }
}
