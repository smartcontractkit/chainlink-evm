// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";

import {Test} from "forge-std/Test.sol";

contract WorkflowRegistry_unlinkOwner is Test {
  WorkflowRegistry public wr;
  address public owner = address(0xabcd);
  uint256 public allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
  address public allowedSigner;
  uint256 public validityTimestamp = uint256(block.timestamp + 1 hours);
  bytes32 public proof = keccak256("test-proof");
  address public notOwner = address(0x1234);
  address public caller;
  string public donLabel = "my-don";

  function setUp() public {
    // hardcode the signer's private key into test environment (so that vm.sign can be used)
    allowedSigner = vm.addr(allowedSignerPrivateKey);
    assertEq(allowedSigner, address(0x86f2cE81640Fd86e68CF3EB25c2801D6E1C62bd0));

    vm.startPrank(owner);
    wr = new WorkflowRegistry();
    wr.setDONLimit(donLabel, 1000, true); // 1000 workflows on the test DON
    address[] memory signers = new address[](1);
    signers[0] = allowedSigner;
    wr.updateAllowedSigners(signers, true);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);
    wr.linkOwner(validityTimestamp, proof, sig);
    vm.stopPrank();
  }

  modifier whenCallerIsEqualToTheOwnerAddress() {
    caller = owner; // caller is the owner
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerWhenRequestTimestampHasExpired()
    external
    whenCallerIsEqualToTheOwnerAddress
  {
    // it should revert with expiration error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(caller); // caller = owner
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.UnlinkOwnerRequestExpired.selector, owner, block.timestamp, validityTimestamp
      )
    );
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  modifier whenTheRequestTimestampHasNotExpired() {
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerGivenTheOwnerIsNotLinked()
    external
    whenCallerIsEqualToTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
  {
    // it should revert with not linked error
    address unlinkedOwner = address(0x5678);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), unlinkedOwner, validityTimestamp, proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = owner
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, unlinkedOwner));
    wr.unlinkOwner(unlinkedOwner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(unlinkedOwner), "Owner should not be linked");
  }

  modifier givenTheOwnerIsLinked() {
    _;
  }

  modifier givenTheProofMatchesTheStoredProof() {
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerWhenTheProofIsValid()
    external
    whenCallerIsEqualToTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    givenTheOwnerIsLinked
    givenTheProofMatchesTheStoredProof
  {
    // it should unlink the owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerWhenTheProofIsNotValid()
    external
    whenCallerIsEqualToTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    givenTheOwnerIsLinked
    givenTheProofMatchesTheStoredProof
  {
    // it should revert with signature error
    uint256 differentValidityTimestamp = uint256(block.timestamp + 2 hours);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, differentValidityTimestamp, proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = owner
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    // calling with validity timestamp that does not match the one from the signature
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsOwnerGivenTheProofDoesNotMatchTheStoredProof()
    external
    whenCallerIsEqualToTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    givenTheOwnerIsLinked
  {
    // it should revert with proof does not match error
    bytes32 differentProof = keccak256("different-proof");
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, differentProof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = owner
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  modifier whenCallerIsDifferentFromTheOwnerAddress() {
    caller = notOwner; // caller is not the owner
    _;
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerWhenRequestTimestampHasExpired()
    external
    whenCallerIsDifferentFromTheOwnerAddress
  {
    // it should revert with expiration error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(caller); // caller = not owner
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.UnlinkOwnerRequestExpired.selector, owner, block.timestamp, validityTimestamp
      )
    );
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerGivenTheOwnerIsNotLinked()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
  {
    // it should revert with not linked error
    address unlinkedOwner = address(0x5678);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), unlinkedOwner, validityTimestamp, proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = not owner
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, unlinkedOwner));
    wr.unlinkOwner(unlinkedOwner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(unlinkedOwner), "Owner should not be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerWhenTheProofIsValid()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    givenTheOwnerIsLinked
    givenTheProofMatchesTheStoredProof
  {
    // it should unlink the owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = not owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerWhenTheProofIsNotValid()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    givenTheOwnerIsLinked
    givenTheProofMatchesTheStoredProof
  {
    // it should revert with signature error
    uint256 differentValidityTimestamp = uint256(block.timestamp + 2 hours);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, differentValidityTimestamp, proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = not owner
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    // calling with validity timestamp that does not match the one from the signature
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  function test_unlinkOwner_UnlinkOwnerWhenCallerIsNotOwnerGivenTheProofDoesNotMatchTheStoredProof()
    external
    whenCallerIsDifferentFromTheOwnerAddress
    whenTheRequestTimestampHasNotExpired
    givenTheOwnerIsLinked
  {
    // it should revert with proof does not match error
    bytes32 differentProof = keccak256("different-proof");
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, differentProof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(caller); // caller = not owner
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  modifier whenTheCallerIsTheOwner() {
    caller = owner; // caller is the owner
    _;
  }

  modifier givenThatOwnerHasNoActiveWorkflows() {
    // create 5 random paused workflows
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.PAUSED, false);
    _;
  }

  function test_unlinkOwner_preUnlinkActionsWhenNONEIsSelectedAsTheUnlinkAction()
    external
    whenTheCallerIsTheOwner
    givenThatOwnerHasNoActiveWorkflows
  {
    // it should unlink the owner without any additional actions
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_preUnlinkActionsWhenREMOVEIsSelectedAsTheUnlinkAction()
    external
    whenTheCallerIsTheOwner
    givenThatOwnerHasNoActiveWorkflows
  {
    // it should unlink the owner without any additional actions
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.REMOVE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_preUnlinkActionsWhenPAUSEIsSelectedAsTheUnlinkAction()
    external
    whenTheCallerIsTheOwner
    givenThatOwnerHasNoActiveWorkflows
  {
    // it should unlink the owner without any additional actions
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.PAUSE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  modifier givenThatOwnerHasActiveWorkflows() {
    // create 5 random active workflows
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.ACTIVE, false);
    _;
  }

  function test_unlinkOwner_preUnlinkActionsWhenNONEIsTheUnlinkAction()
    external
    whenTheCallerIsTheOwner
    givenThatOwnerHasActiveWorkflows
  {
    // it should revert with active workflows error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = owner
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.CannotUnlinkWithActiveWorkflows.selector));
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  function test_unlinkOwner_preUnlinkActionsWhenREMOVETheUnlinkAction()
    external
    whenTheCallerIsTheOwner
    givenThatOwnerHasActiveWorkflows
  {
    // it should remove the workflows and unlink the owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.REMOVE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");

    wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 0, "There should be 0 workflows for the owner");
  }

  function test_unlinkOwner_preUnlinkActionsWhenPAUSEIsTheUnlinkAction()
    external
    whenTheCallerIsTheOwner
    givenThatOwnerHasActiveWorkflows
  {
    // it should pause the workflows and unlink the owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.PAUSE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");

    wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 0, "There should be still 5 workflows for the owner");

    for (uint256 i = 0; i < wrs.length; ++i) {
      assertEq(uint8(wrs[i].status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED), "Workflow should be paused");
    }
  }

  modifier whenTheCallerIsNotTheOwner() {
    caller = notOwner; // caller is not the owner
    _;
  }

  modifier givenThatCallerHasNoActiveWorkflows() {
    // create 5 random paused workflows
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.PAUSED, false);
    _;
  }

  function test_unlinkOwner_preUnlinkActionsWhenNONEIsChosenAsTheUnlinkAction()
    external
    whenTheCallerIsNotTheOwner
    givenThatCallerHasNoActiveWorkflows
  {
    // it should unlink the owner without any additional actions
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = not owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_preUnlinkActionsWhenREMOVEIsChosenAsTheUnlinkAction()
    external
    whenTheCallerIsNotTheOwner
    givenThatCallerHasNoActiveWorkflows
  {
    // it should unlink the owner without any additional actions
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = not owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.REMOVE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  function test_unlinkOwner_preUnlinkActionsWhenPAUSEIsChosenAsTheUnlinkAction()
    external
    whenTheCallerIsNotTheOwner
    givenThatCallerHasNoActiveWorkflows
  {
    // it should unlink the owner without any additional actions
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = not owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.PAUSE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  modifier givenThatCallerHasActiveWorkflows() {
    // create 5 random active workflows
    _upsertTestWorklows(WorkflowRegistry.WorkflowStatus.ACTIVE, false);
    _;
  }

  function test_unlinkOwner_preUnlinkActionsWhenNONEIsEqualToTheUnlinkAction()
    external
    whenTheCallerIsNotTheOwner
    givenThatCallerHasActiveWorkflows
  {
    // it should revert with active workflows error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = not owner
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.CannotUnlinkWithActiveWorkflows.selector));
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  function test_unlinkOwner_preUnlinkActionsWhenREMOVEIsEqualToTheUnlinkAction()
    external
    whenTheCallerIsNotTheOwner
    givenThatCallerHasActiveWorkflows
  {
    // it should remove the workflows and unlink the owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = not owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.REMOVE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");

    wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 0, "There should be 0 workflows for the owner");
  }

  function test_unlinkOwner_preUnlinkActionsWhenPAUSEIsEqualToTheUnlinkAction()
    external
    whenTheCallerIsNotTheOwner
    givenThatCallerHasActiveWorkflows
  {
    // it should pause the workflows and unlink the owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    WorkflowRegistry.WorkflowMetadata[] memory wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 5, "There should be 5 workflows for the owner");

    vm.prank(caller); // caller = owner
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, sig, WorkflowRegistry.PreUnlinkAction.PAUSE_WORKFLOWS);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");

    wrs = wr.getWorkflowMetadataListByOwner(owner, 0, 100);
    assertEq(wrs.length, 0, "There should be still 5 workflows for the owner");

    for (uint256 i = 0; i < wrs.length; ++i) {
      assertEq(uint8(wrs[i].status), uint8(WorkflowRegistry.WorkflowStatus.PAUSED), "Workflow should be paused");
    }
  }

  function _upsertTestWorklows(WorkflowRegistry.WorkflowStatus status, bool keepAlive) internal {
    // Workflow 1: Price Oracle
    bytes32 workflowId1 = keccak256("workflow1");
    string memory workflowName1 = "Price Oracle";
    string memory tag1 = "oracle-main";
    string memory binaryUrl1 = "https://example.com/binaries/price-oracle.wasm";
    string memory configUrl1 = "https://example.com/configs/price-oracle.json";
    bytes memory attributes1 = abi.encode("Price Oracle v1.0");

    vm.prank(owner);
    wr.upsertWorkflow(
      workflowName1, tag1, workflowId1, status, donLabel, binaryUrl1, configUrl1, attributes1, keepAlive
    );

    // Workflow 2: Weather Data Feeder
    bytes32 workflowId2 = keccak256("workflow2");
    string memory workflowName2 = "Weather Data Feeder";
    string memory tag2 = "weather-feed";
    string memory binaryUrl2 = "https://example.com/binaries/weather-data.wasm";
    string memory configUrl2 = "https://example.com/configs/weather-config.json";
    bytes memory attributes2 = abi.encode("Weather Data v2.1");

    vm.prank(owner);
    wr.upsertWorkflow(
      workflowName2, tag2, workflowId2, status, donLabel, binaryUrl2, configUrl2, attributes2, keepAlive
    );

    // Workflow 3: NFT Metadata Service
    bytes32 workflowId3 = keccak256("workflow3");
    string memory workflowName3 = "NFT Metadata Service";
    string memory tag3 = "nft-meta";
    string memory binaryUrl3 = "https://example.com/binaries/nft-metadata.wasm";
    string memory configUrl3 = "https://example.com/configs/nft-settings.json";
    bytes memory attributes3 = abi.encode("NFT Metadata Service v1.2");

    vm.prank(owner);
    wr.upsertWorkflow(
      workflowName3, tag3, workflowId3, status, donLabel, binaryUrl3, configUrl3, attributes3, keepAlive
    );

    // Workflow 4: Cross-Chain Bridge Monitor
    bytes32 workflowId4 = keccak256("workflow4");
    string memory workflowName4 = "Cross-Chain Bridge Monitor";
    string memory tag4 = "bridge-monitor";
    string memory binaryUrl4 = "https://example.com/binaries/bridge-monitor.wasm";
    string memory configUrl4 = "https://example.com/configs/bridge-config.json";
    bytes memory attributes4 = abi.encode("Bridge Monitor v3.0");

    vm.prank(owner);
    wr.upsertWorkflow(
      workflowName4, tag4, workflowId4, status, donLabel, binaryUrl4, configUrl4, attributes4, keepAlive
    );

    // Workflow 5: Sports Data Feed
    bytes32 workflowId5 = keccak256("workflow5");
    string memory workflowName5 = "Sports Data Feed";
    string memory tag5 = "sports-feed";
    string memory binaryUrl5 = "https://example.com/binaries/sports-data.wasm";
    string memory configUrl5 = "https://example.com/configs/sports-config.json";
    bytes memory attributes5 = abi.encode("Sports Data Feed v1.5");

    vm.prank(owner);
    wr.upsertWorkflow(
      workflowName5, tag5, workflowId5, status, donLabel, binaryUrl5, configUrl5, attributes5, keepAlive
    );
  }
}
