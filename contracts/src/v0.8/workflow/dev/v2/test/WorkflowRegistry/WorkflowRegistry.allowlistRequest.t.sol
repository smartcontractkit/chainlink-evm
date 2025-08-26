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

  function test_allowlistRequest_WhenTheUserIsLinked() external {
    //it should allowlist the request digest
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
