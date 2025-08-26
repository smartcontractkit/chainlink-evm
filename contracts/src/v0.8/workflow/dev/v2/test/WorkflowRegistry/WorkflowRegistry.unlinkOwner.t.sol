// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_unlinkOwner is WorkflowRegistrySetup {
  function setUp() public override {
    super.setUp();
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 1000, true); // 1000 workflows on the test DON
  }

  modifier whenTheCallerIsTheOwner() {
    s_user = s_owner;
    _;
  }

  modifier whenTheOwnerIsLinked() {
    _linkOwner(s_owner);
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerWhenRequestTimestampHasExpired()
    external
    whenTheCallerIsTheOwner
    whenTheOwnerIsLinked
  {
    // it should revert with expiration error
    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(s_user);
    (, bytes memory sig) = _getUnlinkProofSignature(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.UnlinkOwnerRequestExpired.selector, s_owner, block.timestamp, s_validityTimestamp
      )
    );
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
  }

  modifier whenTheRequestTimestampHasNotExpired() {
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerGivenTheOwnerIsNotLinked()
    external
    whenTheCallerIsTheOwner
    whenTheRequestTimestampHasNotExpired
  {
    // it should revert with not linked error
    vm.prank(s_user);
    (, bytes memory sig) = _getUnlinkProofSignature(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_owner));
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should not be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerWhenTheProofIsValid()
    external
    whenTheCallerIsTheOwner
    whenTheRequestTimestampHasNotExpired
    whenTheOwnerIsLinked
    whenTheProofMatchesTheStoredProof
  {
    // it should unlink the s_owner
    (bytes32 ownerProof, bytes memory sig) = _getUnlinkProofSignature(s_owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, ownerProof, false);
    vm.prank(s_user);
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerWhenTheProofIsNotValid()
    external
    whenTheCallerIsTheOwner
    whenTheRequestTimestampHasNotExpired
    whenTheOwnerIsLinked
    whenTheProofMatchesTheStoredProof
  {
    // it should revert with signature error
    uint256 differentValidityTimestamp = uint256(block.timestamp + 2 hours);
    (bytes32 storedProof,) = _getUnlinkProofSignature(s_user);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(s_registry), s_owner, differentValidityTimestamp, storedProof
      )
    );
    bytes memory invalidSig = abi.encodePacked(r, s, v);

    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidOwnershipLink.selector, s_owner, s_validityTimestamp, storedProof, invalidSig
      )
    );
    // calling with validity timestamp that does not match the one from the signature
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, invalidSig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerGivenTheProofDoesNotMatchTheStoredProof()
    external
    whenTheCallerIsTheOwner
    whenTheRequestTimestampHasNotExpired
    whenTheOwnerIsLinked
  {
    // it should revert with s_proof does not match error
    bytes32 differentProof = keccak256("different-s_proof");
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(s_registry), s_owner, s_validityTimestamp, differentProof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    (bytes32 storedProof,) = _getUnlinkProofSignature(s_user);
    vm.prank(s_user); // s_user = s_owner
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidOwnershipLink.selector, s_owner, s_validityTimestamp, storedProof, sig
      )
    );
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
  }

  modifier whenCallerIsDifferentFromTheOwnerAddress() {
    s_user = s_stranger; // s_user is not the s_owner
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerWhenRequestTimestampHasExpired()
    external
    whenCallerIsDifferentFromTheOwnerAddress
  {
    // it should revert with expiration error
    // first link the owner
    _linkOwner(s_owner);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.UnlinkOwnerRequestExpired.selector, s_owner, block.timestamp, s_validityTimestamp
      )
    );
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerGivenTheOwnerIsNotLinked()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
  {
    // it should revert with not linked error
    address unlinkedOwner = makeAddr("unlinked-owner");
    (, bytes memory sig) = _getUnlinkProofSignature(unlinkedOwner);
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, unlinkedOwner));
    s_registry.unlinkOwner(unlinkedOwner, s_validityTimestamp, sig);
    assertFalse(s_registry.isOwnerLinked(unlinkedOwner), "Owner should not be linked");
  }

  modifier whenTheProofMatchesTheStoredProof() {
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerWhenTheProofIsValid()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    whenTheOwnerIsLinked
    whenTheProofMatchesTheStoredProof
  {
    // it should unlink the s_owner
    (bytes32 proof, bytes memory sig) = _getUnlinkProofSignature(s_owner);
    vm.prank(s_user);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, proof, false);
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerWhenTheProofIsNotValid()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    whenTheOwnerIsLinked
    whenTheProofMatchesTheStoredProof
  {
    // it should revert with signature error
    (bytes32 validProof,) = _getUnlinkProofSignature(s_owner);
    uint256 differentValidityTimestamp = uint256(block.timestamp + 2 hours);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(s_registry), s_owner, differentValidityTimestamp, validProof
      )
    );
    bytes memory invalidSig = abi.encodePacked(r, s, v);

    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidOwnershipLink.selector, s_owner, s_validityTimestamp, validProof, invalidSig
      )
    );
    // calling with validity timestamp that does not match the one from the signature
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, invalidSig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerGivenTheProofDoesNotMatchTheStoredProof()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    whenTheOwnerIsLinked
  {
    // it should revert with s_proof does not match error
    bytes32 differentProof = keccak256("different-s_proof");
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(s_registry), s_owner, s_validityTimestamp, differentProof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);
    (bytes32 storedProof,) = _getUnlinkProofSignature(s_owner);

    vm.prank(s_user); // s_user = not s_owner
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidOwnershipLink.selector, s_owner, s_validityTimestamp, storedProof, sig
      )
    );
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
  }

  modifier givenThatOwnerHasNoActiveWorkflows() {
    // create 5 random paused workflows
    _linkOwner(s_owner);
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.PAUSED, false, s_owner);
    _;
  }

  modifier givenThatOwnerHasActiveWorkflows() {
    // create 5 random active workflows
    _linkOwner(s_owner);
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_owner);
    _;
  }

  function test_unlinkOwner_AutomaticallyDeletesWorkflowsWhenOwnerIsLinked()
    external
    whenTheCallerIsTheOwner
    givenThatOwnerHasActiveWorkflows
  {
    // it should automatically delete workflows and unlink the owner
    (bytes32 proof, bytes memory sig) = _getUnlinkProofSignature(s_owner);

    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the s_owner");

    vm.prank(s_user); // s_user = s_owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, proof, false);
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);

    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should be unlinked");

    // Verify all workflows were deleted
    WorkflowRegistry.WorkflowMetadataView[] memory remainingWrs = s_registry.getWorkflowListByOwner(s_owner, 0, 100);
    assertEq(remainingWrs.length, 0, "All workflows should be deleted");
  }

  modifier whenTheCallerIsNotTheOwner() {
    s_user = s_stranger; // s_user is not the s_owner
    _;
  }

  modifier givenThatCallerHasNoActiveWorkflows() {
    // create 5 random paused workflows
    _linkOwner(s_owner);
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.PAUSED, false, s_owner);
    _;
  }

  modifier givenThatCallerHasActiveWorkflows() {
    // create 5 random active workflows
    _linkOwner(s_owner);
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_owner);
    _;
  }

  function test_unlinkOwner_AutomaticallyDeletesWorkflowsWhenCallerIsNotOwner()
    external
    whenTheCallerIsNotTheOwner
    givenThatCallerHasActiveWorkflows
  {
    // it should automatically delete workflows and unlink the owner even when caller is not the owner
    (bytes32 proof, bytes memory sig) = _getUnlinkProofSignature(s_owner);

    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the s_owner");

    vm.prank(s_user); // s_user = not s_owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, proof, false);
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, sig);

    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should be unlinked");

    // Verify all workflows were deleted
    WorkflowRegistry.WorkflowMetadataView[] memory remainingWrs = s_registry.getWorkflowListByOwner(s_owner, 0, 100);
    assertEq(remainingWrs.length, 0, "All workflows should be deleted");
  }
}
