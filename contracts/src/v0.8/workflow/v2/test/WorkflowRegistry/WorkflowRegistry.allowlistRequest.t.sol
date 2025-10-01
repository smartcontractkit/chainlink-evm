// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_allowlistRequest is WorkflowRegistrySetup {
  // NOTE: lock and control current timestamp this way due to issues when via-ir is enabled:
  // https://github.com/foundry-rs/foundry/issues/1373
  uint256 public currentTimestamp = block.timestamp;

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
    // It should fail because the previous request is still valid
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

    // Try to update the same request with new expiry (this should fail because previous is still valid)
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.PreviousAllowlistedRequestStillValid.selector, s_user, requestDigest, initialExpiry
      )
    );
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest, updatedExpiry);

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
    bool stopSearch;
    assertEq(s_registry.totalAllowlistedRequests(), 3, "Should have exactly 3 unique requests");

    // Verify contents of all requests match expectations, requests are returned in the reverse order they were added
    (requests, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(2, 0);
    assertEq(requests.length, 3, "Should return exactly 3 requests");

    assertEq(requests[0].expiryTimestamp, expiry3, "Third request expiry should match");
    assertEq(requests[0].owner, s_user, "Owner should match");
    assertEq(requests[0].requestDigest, requestDigest3, "Third request digest should match");

    assertEq(requests[1].expiryTimestamp, expiry2, "Initial second request expiry should match");
    assertEq(requests[1].owner, s_user, "Owner should match");
    assertEq(requests[1].requestDigest, requestDigest2, "Second request digest should match");

    assertEq(requests[2].expiryTimestamp, initialExpiry, "Initial first request expiry should match");
    assertEq(requests[2].owner, s_user, "Owner should match");
    assertEq(requests[2].requestDigest, requestDigest, "First request digest should match");

    // Try to update the second request and verify that this will fail
    uint32 newExpiry2 = uint32(block.timestamp + 5 hours);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.PreviousAllowlistedRequestStillValid.selector, s_user, requestDigest2, expiry2
      )
    );
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest2, newExpiry2);

    // Fast forward the block time beyond the initial expiry of the second request, this will allowlist a new one
    vm.warp(currentTimestamp + 3 hours); // Advances the block timestamp by 3 hours only for the next call
    emit WorkflowRegistry.RequestAllowlisted(s_user, requestDigest2, newExpiry2);
    vm.prank(s_user);
    s_registry.allowlistRequest(requestDigest2, newExpiry2);

    uint256 totalRequests = s_registry.totalAllowlistedRequests();

    // Verify that we have 2 requests, because the first one is expired now
    vm.warp(currentTimestamp + 3 hours); // Advances the block timestamp by 3 hours only for the next call
    (requests, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(totalRequests - 1, 0);
    assertEq(requests.length, 2, "Should return exactly 2 requests");

    // Verify contents of all requests match expectations, requests are returned in the reverse order they were added
    assertEq(requests[0].expiryTimestamp, newExpiry2, "Update second request expiry should match");
    assertEq(requests[0].owner, s_user, "Owner should match");
    assertEq(requests[0].requestDigest, requestDigest2, "Second request digest should match");

    assertEq(requests[1].expiryTimestamp, expiry3, "Third request expiry should match");
    assertEq(requests[1].owner, s_user, "Owner should match");
    assertEq(requests[1].requestDigest, requestDigest3, "Third request digest should match");
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
