// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getActiveAllowlistedRequestsReverse is WorkflowRegistrySetup {
  // NOTE: lock and control current timestamp this way due to issues when via-ir is enabled:
  // https://github.com/foundry-rs/foundry/issues/1373
  uint256 public currentTimestamp = block.timestamp;

  function test_getActiveAllowlistedRequestsReverse_WhenNoRequestsAreAllowlisted() external view {
    // it should return an empty array
    uint256 total = s_registry.totalAllowlistedRequests();
    (WorkflowRegistry.OwnerAllowlistedRequest[] memory requests, uint256 nextIndex, bool stopSearch) =
      s_registry.getActiveAllowlistedRequestsReverse(0, 100, 0);
    assertEq(total, 0, "Total number of allowlisted requests should be 0");
    assertEq(requests.length, 0, "Zero requests should be returned");
    assertEq(nextIndex, 0, "Next index should be 0");
    assertEq(stopSearch, true, "Stop search should be true");
  }

  modifier whenSomeRequestsAreAllowlisted() {
    _;
  }

  function test_getActiveAllowlistedRequestsReverse_WhenNoneOfTheRequestsHaveExpired()
    external
    whenSomeRequestsAreAllowlisted
  {
    // it should return all requests
    _linkTestOwners();
    _allowlistValidTestRequests();

    uint256 total = s_registry.totalAllowlistedRequests();
    (WorkflowRegistry.OwnerAllowlistedRequest[] memory requests, uint256 nextIndex, bool stopSearch) =
      s_registry.getActiveAllowlistedRequestsReverse(0, 100, 0);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 6, "All 6 requests should be returned");
    assertEq(stopSearch, true, "Stop search should be true because we scanned all requests");
    assertEq(nextIndex, 6, "Next index should be 6");
    assertEq(
      keccak256("request-digest-3-owner-3"), requests[0].requestDigest, "All - Sixth request digest should match"
    );
    assertEq(
      keccak256("request-digest-2-owner-3"), requests[1].requestDigest, "All - Fifth request digest should match"
    );
    assertEq(
      keccak256("request-digest-1-owner-3"), requests[2].requestDigest, "All - Fourth request digest should match"
    );
    assertEq(
      keccak256("request-digest-1-owner-2"), requests[3].requestDigest, "All - Third request digest should match"
    );
    assertEq(
      keccak256("request-digest-2-owner-1"), requests[4].requestDigest, "All - Second request digest should match"
    );
    assertEq(
      keccak256("request-digest-1-owner-1"), requests[5].requestDigest, "All - First request digest should match"
    );

    // try out pagination - page size 2
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 2, 0);
    assertEq(requests.length, 2, "Page 1 - 2 requests should be returned");
    assertEq(nextIndex, 2, "Page 1 - Next index should be 2");
    assertEq(stopSearch, false, "Page 1 - Stop search should be false");
    assertEq(
      keccak256("request-digest-3-owner-3"), requests[0].requestDigest, "Page 1 - Sixth request digest should match"
    );
    assertEq(
      keccak256("request-digest-2-owner-3"), requests[1].requestDigest, "Page 1 - Fifth request digest should match"
    );

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(2, 2, 0);
    assertEq(requests.length, 2, "Page 2 - 2 requests should be returned");
    assertEq(nextIndex, 4, "Page 2 - Next index should be 4");
    assertEq(stopSearch, false, "Page 2 - Stop search should be false");
    assertEq(
      keccak256("request-digest-1-owner-3"), requests[0].requestDigest, "Page 2 - Fourth request digest should match"
    );
    assertEq(
      keccak256("request-digest-1-owner-2"), requests[1].requestDigest, "Page 2 - Third request digest should match"
    );

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(4, 2, 0);
    assertEq(requests.length, 2, "Page 3 - 2 requests should be returned");
    assertEq(nextIndex, 6, "Page 3 - Next index should be 6");
    assertEq(
      stopSearch, true, "Page 3 - Stop search should be true because we scanned all requests and returned the last 2"
    );
    assertEq(
      keccak256("request-digest-2-owner-1"), requests[0].requestDigest, "Page 3 - Second request digest should match"
    );
    assertEq(
      keccak256("request-digest-1-owner-1"), requests[1].requestDigest, "Page 3 - First request digest should match"
    );

    // try out pagination - page size 4
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 4, 0);
    assertEq(requests.length, 4, "4 requests should be returned");
    assertEq(nextIndex, 4, "Next index should be 4");
    assertEq(stopSearch, false, "Stop search should be false");
    assertEq(keccak256("request-digest-3-owner-3"), requests[0].requestDigest, "Sixth request digest should match");
    assertEq(keccak256("request-digest-2-owner-3"), requests[1].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-1-owner-3"), requests[2].requestDigest, "Fourth request digest should match");
    assertEq(keccak256("request-digest-1-owner-2"), requests[3].requestDigest, "Third request digest should match");

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(4, 4, 0);
    assertEq(requests.length, 2, "2 requests should be returned");
    assertEq(nextIndex, 6, "Next index should be 6");
    assertEq(stopSearch, true, "Stop search should be true because we scanned all requests and returned the last two");
    assertEq(keccak256("request-digest-2-owner-1"), requests[0].requestDigest, "Second request digest should match");
    assertEq(keccak256("request-digest-1-owner-1"), requests[1].requestDigest, "First request digest should match");

    // try out pagination - out of bounds
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(8, 4, 0);
    assertEq(requests.length, 0, "No requests should be returned");
    assertEq(nextIndex, 0, "Next index should be 0");
    assertEq(stopSearch, true, "Stop search should be true");

    // allowlist a few more requests to try out bounding the results to only the new requests by using the
    // lastTotalCount parameter
    address owner1 = address(0x1);
    bytes32 requestDigest = keccak256("request-digest-new-1");
    uint32 expiryTimestamp = uint32(currentTimestamp + 10 hours);
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-new-2");
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-new-3");
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-new-4");
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    assertEq(s_registry.totalAllowlistedRequests(), 10, "Should have exactly 10 unique requests");
    // last known total count was 6, so we will try to skip all the seen requests
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 100, 6);
    assertEq(requests.length, 4, "4 new requests should be returned");
    // index 5 is aligned with the requested total count of 6
    assertEq(nextIndex, 5, "Next index should be 5");
    assertEq(stopSearch, true, "Stop search should be true, because we scanned all NEW requests");
    assertEq(keccak256("request-digest-new-4"), requests[0].requestDigest, "Newest request digest should match");
    assertEq(keccak256("request-digest-new-3"), requests[1].requestDigest, "2nd newest request digest should match");
    assertEq(keccak256("request-digest-new-2"), requests[2].requestDigest, "3rd newest request digest should match");
    assertEq(keccak256("request-digest-new-1"), requests[3].requestDigest, "4th newest request digest should match");

    // try out the same but with pagination (page size 2)
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 2, 6);
    assertEq(requests.length, 2, "Page 1 - 2 new requests should be returned");
    assertEq(nextIndex, 2, "Page 1 - Next index should be 2");
    assertEq(stopSearch, false, "Page 1 - Stop search should be false");
    assertEq(
      keccak256("request-digest-new-4"), requests[0].requestDigest, "Page 1 - Newest request digest should match"
    );
    assertEq(
      keccak256("request-digest-new-3"), requests[1].requestDigest, "Page 1 - 2nd newest request digest should match"
    );

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(2, 2, 6);
    assertEq(requests.length, 2, "Page 2 - 2 new requests should be returned");
    assertEq(nextIndex, 4, "Page 2 - Next index should be 4");
    assertEq(stopSearch, false, "Page 2 - Stop search should be false");
    assertEq(
      keccak256("request-digest-new-2"), requests[0].requestDigest, "Page 2 - 3rd newest request digest should match"
    );
    assertEq(
      keccak256("request-digest-new-1"), requests[1].requestDigest, "Page 2 - 4th newest request digest should match"
    );

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(4, 2, 6);
    assertEq(requests.length, 0, "Page 3 - No new requests should be returned");
    assertEq(nextIndex, 5, "Page 3 - Next index should be 5");
    assertEq(stopSearch, true, "Page 3 - Stop search should be true, because we have scanned all NEW requests");

    // try out the same but with pagination (page size 3)
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 3, 6);
    assertEq(requests.length, 3, "Page 1 - 3 new requests should be returned");
    assertEq(nextIndex, 3, "Page 1 - Next index should be 3");
    assertEq(stopSearch, false, "Page 1 - Stop search should be false");
    assertEq(
      keccak256("request-digest-new-4"), requests[0].requestDigest, "Page 1 - Newest request digest should match"
    );
    assertEq(
      keccak256("request-digest-new-3"), requests[1].requestDigest, "Page 1 - 2nd newest request digest should match"
    );
    assertEq(
      keccak256("request-digest-new-2"), requests[2].requestDigest, "Page 1 - 3rd newest request digest should match"
    );

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(3, 3, 6);
    assertEq(requests.length, 1, "Page 2 - 1 new request should be returned");
    assertEq(nextIndex, 5, "Page 2 - Next index should be 5");
    assertEq(stopSearch, true, "Page 2 - Stop search should be true, because we have scanned all NEW requests");
    assertEq(
      keccak256("request-digest-new-1"), requests[0].requestDigest, "Page 2 - 4th newest request digest should match"
    );
  }

  function test_getActiveAllowlistedRequestsReverse_WhenSomeOfTheRequestsHaveExpired()
    external
    whenSomeRequestsAreAllowlisted
  {
    // it should return only the non-expired requests
    _linkTestOwners();
    _allowlistValidTestRequests();

    vm.warp(currentTimestamp + 1 hours);
    uint256 total = s_registry.totalAllowlistedRequests();
    // this will time out request-digest-1-owner-1, request-digest-2-owner-1 and request-digest-1-owner-3
    vm.warp(currentTimestamp + 1 hours);
    (WorkflowRegistry.OwnerAllowlistedRequest[] memory requests, uint256 nextIndex, bool stopSearch) =
      s_registry.getActiveAllowlistedRequestsReverse(0, 100, 0);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 3, "3 requests should be returned");
    assertEq(nextIndex, 6, "Next index should be 6");
    assertEq(stopSearch, true, "Stop search should be true because we scanned all requests");
    assertEq(keccak256("request-digest-3-owner-3"), requests[0].requestDigest, "Sixth request digest should match");
    assertEq(keccak256("request-digest-2-owner-3"), requests[1].requestDigest, "Fifth request digest should match");
    assertEq(keccak256("request-digest-1-owner-2"), requests[2].requestDigest, "Third request digest should match");

    vm.warp(currentTimestamp + 2 hours);
    total = s_registry.totalAllowlistedRequests();
    // this will time out all requests aside from request-digest-3-owner-3
    vm.warp(currentTimestamp + 2 hours);
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 100, 0);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 1, "1 request should be returned");
    assertEq(keccak256("request-digest-3-owner-3"), requests[0].requestDigest, "Sixth request digest should match");

    vm.warp(currentTimestamp + 3 hours);
    total = s_registry.totalAllowlistedRequests();
    // this will time out all requests
    vm.warp(currentTimestamp + 3 hours);
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 100, 0);
    assertEq(total, 6, "Total number of allowlisted requests should be 6");
    assertEq(requests.length, 0, "No requests should be returned");
  }

  function test_getActiveAllowlistedRequestsReverse_WhenStopCriteriaWasReached()
    external
    whenSomeRequestsAreAllowlisted
  {
    // it should return only the non-expired requests with stopSearch true

    // reduce max allowed expiry to 5 minutes
    vm.prank(s_owner);
    s_registry.setConfig(64, 32, 200, 1024, 300);

    _linkTestOwners();
    _allowlistMoreValidTestRequests();

    uint256 total = s_registry.totalAllowlistedRequests();
    assertEq(total, 12, "Total number of allowlisted requests should be 12");
    (WorkflowRegistry.OwnerAllowlistedRequest[] memory requests, uint256 nextIndex, bool stopSearch) =
      s_registry.getActiveAllowlistedRequestsReverse(0, 100, 0);
    assertEq(requests.length, 6, "6 requests should be returned");
    // We retrieved all 6 active requests. Because we start at index 0, the 7th active request will be at index 6
    assertEq(nextIndex, 7, "Next index should be 7, this item is considered too old");
    assertEq(stopSearch, true, "Stop search should be true");

    // this will time out all request with expiry less than currentTimestamp + 1 minutes
    vm.warp(currentTimestamp + 1 minutes);
    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(0, 2, 0);
    assertEq(requests.length, 2, "2 requests should be returned");
    assertEq(nextIndex, 3, "Next index should be 3");
    assertEq(stopSearch, false, "Stop search should be false");
    assertEq(keccak256("request-digest-6-owner-3"), requests[0].requestDigest, "12th request digest should match");
    assertEq(keccak256("request-digest-4-owner-3"), requests[1].requestDigest, "10th request digest should match");

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(3, 2, 0);
    assertEq(requests.length, 2, "2 requests should be returned");
    assertEq(nextIndex, 6, "Next index should be 6");
    assertEq(stopSearch, false, "Stop search should be false");
    assertEq(keccak256("request-digest-2-owner-2"), requests[0].requestDigest, "9th request digest should match");
    assertEq(keccak256("request-digest-3-owner-1"), requests[1].requestDigest, "7th request digest should match");

    (requests, nextIndex, stopSearch) = s_registry.getActiveAllowlistedRequestsReverse(6, 2, 0);
    assertEq(requests.length, 0, "0 requests should be returned");
    assertEq(nextIndex, 7, "Next index should be 7, this item is considered too old");
    assertEq(stopSearch, true, "Stop search should be true");
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

  function _allowlistMoreValidTestRequests() internal {
    // owner1 - 2 request digests
    address owner1 = address(0x1);
    bytes32 requestDigest = keccak256("request-digest-1-owner-1");
    uint32 expiryTimestamp = uint32(currentTimestamp + 1 minutes);
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-2-owner-1");
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner2 - 1 request digest
    address owner2 = address(0x2);
    requestDigest = keccak256("request-digest-1-owner-2");
    expiryTimestamp = uint32(currentTimestamp + 2 minutes);
    vm.prank(owner2);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner3 - 3 request digests
    address owner3 = address(0x3);
    requestDigest = keccak256("request-digest-1-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 1 minutes);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-2-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 2 minutes);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-3-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 3 minutes);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // double the maximum expiration time, all the requests above expired
    currentTimestamp += 10 minutes;
    vm.warp(currentTimestamp);

    // owner1 - 2 more request digests
    requestDigest = keccak256("request-digest-3-owner-1");
    expiryTimestamp = uint32(currentTimestamp + 3 minutes);
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-4-owner-1");
    expiryTimestamp = uint32(currentTimestamp + 1 minutes);
    vm.prank(owner1);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner2 - 1 more request digest
    requestDigest = keccak256("request-digest-2-owner-2");
    expiryTimestamp = uint32(currentTimestamp + 4 minutes);
    vm.prank(owner2);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);

    // owner3 - 3 more request digests
    requestDigest = keccak256("request-digest-4-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 3 minutes);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-5-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 1 minutes);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
    requestDigest = keccak256("request-digest-6-owner-3");
    expiryTimestamp = uint32(currentTimestamp + 3 minutes);
    vm.prank(owner3);
    s_registry.allowlistRequest(requestDigest, expiryTimestamp);
  }
}
