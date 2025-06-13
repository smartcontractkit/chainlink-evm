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

  function setUp() public {
    // hardcode the signer's private key into test environment (so that vm.sign can be used)
    allowedSigner = vm.addr(allowedSignerPrivateKey);
    assertEq(allowedSigner, address(0x86f2cE81640Fd86e68CF3EB25c2801D6E1C62bd0));

    vm.startPrank(owner);
    wr = new WorkflowRegistry();
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
    wr.unlinkOwner(owner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
    wr.unlinkOwner(unlinkedOwner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
    emit WorkflowRegistry.OwnershipLinkUpdatedV1(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
    wr.unlinkOwner(owner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
      abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkProofDoesNotMatch.selector, owner, differentProof, proof)
    );
    wr.unlinkOwner(owner, validityTimestamp, differentProof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
    wr.unlinkOwner(owner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
    wr.unlinkOwner(unlinkedOwner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
    emit WorkflowRegistry.OwnershipLinkUpdatedV1(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
    wr.unlinkOwner(owner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
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
      abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkProofDoesNotMatch.selector, owner, differentProof, proof)
    );
    wr.unlinkOwner(owner, validityTimestamp, differentProof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }
}
