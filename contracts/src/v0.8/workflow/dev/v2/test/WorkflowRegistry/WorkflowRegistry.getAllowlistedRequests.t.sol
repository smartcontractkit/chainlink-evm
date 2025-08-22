// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getAllowlistedRequests is WorkflowRegistrySetup {
  function test_getAllowlistedRequests_WhenNoRequestsAreAllowlisted() external {
    // it should return an empty array
  }

  modifier whenSomeRequestsAreAllowlisted() {
    _;
  }

  function test_getAllowlistedRequests_WhenNoneOfTheRequestsHaveExpired() external whenSomeRequestsAreAllowlisted {
    // it should return all requests
    _linkTestOwners();
    _allowlistValidTestRequests();

    (WorkflowRegistry.OwnerAllowlistedRequest[] memory requests, uint256 total) =
      s_registry.getAllowlistedRequests(0, 100);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 6, "All 6 requests should be returned");
    assertEq(keccak256("request-digest-1-owner-1"), requests[0].requestDigest, "First request digest should match");
    assertEq(keccak256("request-digest-2-owner-1"), requests[1].requestDigest, "Second request digest should match");
    assertEq(keccak256("request-digest-1-owner-2"), requests[2].requestDigest, "Third request digest should match");
    assertEq(keccak256("request-digest-1-owner-3"), requests[3].requestDigest, "Fourth request digest should match");
    assertEq(keccak256("request-digest-2-owner-3"), requests[4].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-3-owner-3"), requests[5].requestDigest, "Sixth request digest should match");
  }

  function test_getAllowlistedRequests_WhenSomeOfTheRequestsHaveExpired() external whenSomeRequestsAreAllowlisted {
    // it should return only the non-expired requests
    _linkTestOwners();
    _allowlistValidTestRequests();

    // Advances the block timestamp by 1 hour only for the next call
    // this will time out request-digest-1-owner-1, request-digest-2-owner-1 and request-digest-1-owner-3
    vm.warp(block.timestamp + 1 hours);
    (WorkflowRegistry.OwnerAllowlistedRequest[] memory requests, uint256 total) =
      s_registry.getAllowlistedRequests(0, 100);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 3, "3 requests should be returned");
    assertEq(keccak256("request-digest-1-owner-2"), requests[0].requestDigest, "Third request digest should match");
    assertEq(keccak256("request-digest-2-owner-3"), requests[1].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-3-owner-3"), requests[2].requestDigest, "Sixth request digest should match");
  }

  function _linkTestOwners() internal {
    _linkOwner(address(0x1)); // owner1
    _linkOwner(address(0x2)); // owner2
    _linkOwner(address(0x3)); // owner3
  }

  // total of 6 valid request digests
  function _allowlistValidTestRequests() internal {
    // owner1 - 2 request digests
    address owner1 = address(0x1);
    bytes32 requestDigest = keccak256("request-digest-1-owner-1");
    uint32 expiryTimestamp = uint32(block.timestamp + 1 hours);
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-2-owner-1");
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner2 - 1 request digest
    address owner2 = address(0x2);
    requestDigest = keccak256("request-digest-1-owner-2");
    expiryTimestamp = uint32(block.timestamp + 2 hours);
    vm.prank(owner2);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner3 - 3 request digests
    address owner3 = address(0x3);
    requestDigest = keccak256("request-digest-1-owner-3");
    expiryTimestamp = uint32(block.timestamp + 1 hours);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-2-owner-3");
    expiryTimestamp = uint32(block.timestamp + 2 hours);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-3-owner-3");
    expiryTimestamp = uint32(block.timestamp + 3 hours);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
  }
}
