// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getAllowlistedRequests is WorkflowRegistrySetup {
  // NOTE: lock and control current timestamp this way due to issues when via-ir is enabled:
  // https://github.com/foundry-rs/foundry/issues/1373
  uint256 public currentTimestamp = block.timestamp;

  function test_getAllowlistedRequests_WhenNoRequestsAreAllowlisted() external view {
    // it should return an empty array
    uint256 total = s_registry.totalAllowlistedRequests();
    WorkflowRegistry.OwnerAllowlistedRequest[] memory requests = s_registry.getAllowlistedRequests(0, 100);
    assertEq(total, 0, "Total number of allowlisted requests should be 0");
    assertEq(requests.length, 0, "Zero requests should be returned");
  }

  modifier whenSomeRequestsAreAllowlisted() {
    _;
  }

  function test_getAllowlistedRequests_WhenNoneOfTheRequestsHaveExpired() external whenSomeRequestsAreAllowlisted {
    // it should return all requests
    _linkTestOwners();
    _allowlistValidTestRequests();

    uint256 total = s_registry.totalAllowlistedRequests();
    WorkflowRegistry.OwnerAllowlistedRequest[] memory requests = s_registry.getAllowlistedRequests(0, 100);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 6, "All 6 requests should be returned");
    assertEq(keccak256("request-digest-1-owner-1"), requests[0].requestDigest, "First request digest should match");
    assertEq(keccak256("request-digest-2-owner-1"), requests[1].requestDigest, "Second request digest should match");
    assertEq(keccak256("request-digest-1-owner-2"), requests[2].requestDigest, "Third request digest should match");
    assertEq(keccak256("request-digest-1-owner-3"), requests[3].requestDigest, "Fourth request digest should match");
    assertEq(keccak256("request-digest-2-owner-3"), requests[4].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-3-owner-3"), requests[5].requestDigest, "Sixth request digest should match");

    // try out pagination - page size 2
    requests = s_registry.getAllowlistedRequests(0, 2);
    assertEq(requests.length, 2, "2 requests should be returned");
    assertEq(keccak256("request-digest-1-owner-1"), requests[0].requestDigest, "First request digest should match");
    assertEq(keccak256("request-digest-2-owner-1"), requests[1].requestDigest, "Second request digest should match");

    requests = s_registry.getAllowlistedRequests(2, 2);
    assertEq(requests.length, 2, "2 requests should be returned");
    assertEq(keccak256("request-digest-1-owner-2"), requests[0].requestDigest, "Third request digest should match");
    assertEq(keccak256("request-digest-1-owner-3"), requests[1].requestDigest, "Fourth request digest should match");

    requests = s_registry.getAllowlistedRequests(4, 2);
    assertEq(requests.length, 2, "2 requests should be returned");
    assertEq(keccak256("request-digest-2-owner-3"), requests[0].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-3-owner-3"), requests[1].requestDigest, "Sixth request digest should match");

    // try out pagination - page size 4
    requests = s_registry.getAllowlistedRequests(0, 4);
    assertEq(requests.length, 4, "4 requests should be returned");
    assertEq(keccak256("request-digest-1-owner-1"), requests[0].requestDigest, "First request digest should match");
    assertEq(keccak256("request-digest-2-owner-1"), requests[1].requestDigest, "Second request digest should match");
    assertEq(keccak256("request-digest-1-owner-2"), requests[2].requestDigest, "Third request digest should match");
    assertEq(keccak256("request-digest-1-owner-3"), requests[3].requestDigest, "Fourth request digest should match");

    requests = s_registry.getAllowlistedRequests(4, 4);
    assertEq(requests.length, 2, "2 requests should be returned");
    assertEq(keccak256("request-digest-2-owner-3"), requests[0].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-3-owner-3"), requests[1].requestDigest, "Sixth request digest should match");

    // try out pagination - out of bounds
    requests = s_registry.getAllowlistedRequests(8, 4);
    assertEq(requests.length, 0, "No requests should be returned");
  }

  function test_getAllowlistedRequests_WhenSomeOfTheRequestsHaveExpired() external whenSomeRequestsAreAllowlisted {
    // it should return only the non-expired requests
    _linkTestOwners();
    _allowlistValidTestRequests();

    vm.warp(currentTimestamp + 1 hours);
    uint256 total = s_registry.totalAllowlistedRequests();
    // this will time out request-digest-1-owner-1, request-digest-2-owner-1 and request-digest-1-owner-3
    vm.warp(currentTimestamp + 1 hours);
    WorkflowRegistry.OwnerAllowlistedRequest[] memory requests = s_registry.getAllowlistedRequests(0, 100);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 3, "3 requests should be returned");
    assertEq(keccak256("request-digest-1-owner-2"), requests[0].requestDigest, "Third request digest should match");
    assertEq(keccak256("request-digest-2-owner-3"), requests[1].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-3-owner-3"), requests[2].requestDigest, "Sixth request digest should match");

    vm.warp(currentTimestamp + 2 hours);
    total = s_registry.totalAllowlistedRequests();
    // this will time out all requests aside from request-digest-3-owner-3
    vm.warp(currentTimestamp + 2 hours);
    requests = s_registry.getAllowlistedRequests(0, 100);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 1, "1 request should be returned");
    assertEq(keccak256("request-digest-3-owner-3"), requests[0].requestDigest, "Sixth request digest should match");

    vm.warp(currentTimestamp + 3 hours);
    total = s_registry.totalAllowlistedRequests();
    // this will time out all requests
    vm.warp(currentTimestamp + 3 hours);
    requests = s_registry.getAllowlistedRequests(0, 100);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 0, "No requests should be returned");
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
    uint32 expiryTimestamp = uint32(currentTimestamp + 1 hours);
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-2-owner-1");
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner2 - 1 request digest
    address owner2 = address(0x2);
    requestDigest = keccak256("request-digest-1-owner-2");
    expiryTimestamp = uint32(currentTimestamp + 2 hours);
    vm.prank(owner2);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner3 - 3 request digests
    address owner3 = address(0x3);
    requestDigest = keccak256("request-digest-1-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 1 hours);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-2-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 2 hours);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-3-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 3 hours);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
  }
}
