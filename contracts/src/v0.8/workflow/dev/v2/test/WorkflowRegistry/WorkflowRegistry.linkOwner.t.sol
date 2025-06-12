// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";

import {ECDSA} from "@openzeppelin/contracts@5.1.0/utils/cryptography/ECDSA.sol";

import {Test} from "forge-std/Test.sol";

contract WorkflowRegistry_linkOwner is Test {
  WorkflowRegistry public wr;
  address public owner = address(0xabcd);
  uint256 public allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
  address public allowedSigner;
  uint256 public validityTimestamp = uint256(block.timestamp + 1 hours);
  bytes32 public proof = keccak256("test-proof");

  function setUp() public {
    // hardcode the signer's private key into test environment (so that vm.sign can be used)
    allowedSigner = vm.addr(allowedSignerPrivateKey);
    assertEq(allowedSigner, address(0x86f2cE81640Fd86e68CF3EB25c2801D6E1C62bd0));

    vm.startPrank(owner);
    wr = new WorkflowRegistry();
    address[] memory signers = new address[](1);
    signers[0] = allowedSigner;
    wr.updateAllowedSigners(signers, true);
    vm.stopPrank();
  }

  modifier whenTheOwnerIsNotAlreadyLinked() {
    _;
  }

  modifier whenTheTimestampHasNotExpired() {
    _;
  }

  function test_linkOwner_WhenProofIsValid() external whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired {
    // it should link the owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, true);
    wr.linkOwner(validityTimestamp, proof, sig);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
  }

  function test_linkOwner_WhenTheProofIsNotSignedByAnAllowedSigner()
    external
    whenTheOwnerIsNotAlreadyLinked
    whenTheTimestampHasNotExpired
  {
    // it should revert with signature error
    uint256 unknownSignerPrivateKey = 0xffc0c927f94d71f7c5c21a865d7c47d050a34f1583ba93576edf67cf2fa32da7;
    address unknownSigner = vm.addr(unknownSignerPrivateKey);
    assertEq(unknownSigner, address(0xfF2B8E43743892d9a8416254711A473b8B70DDe4));

    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      unknownSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    wr.linkOwner(validityTimestamp, proof, sig);
    assertFalse(wr.isOwnerLinked(owner), "Owner should not be linked");
  }

  function test_linkOwner_WhenTheProofContainsInvalidData()
    external
    whenTheOwnerIsNotAlreadyLinked
    whenTheTimestampHasNotExpired
  {
    // it should revert with invalid signature error
    address invalidOwner = address(0x1234);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), invalidOwner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    wr.linkOwner(validityTimestamp, proof, sig);
    assertFalse(wr.isOwnerLinked(owner), "Owner should not be linked");
  }

  function test_linkOwner_WhenTheSignatureIsNotValid()
    external
    whenTheOwnerIsNotAlreadyLinked
    whenTheTimestampHasNotExpired
  {
    // it should revert with internal signature error
    bytes memory invalidSignature = "invalid-signature";

    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidSignature.selector, invalidSignature, ECDSA.RecoverError.InvalidSignatureLength, 0x11
      )
    );
    wr.linkOwner(validityTimestamp, proof, invalidSignature);
    assertFalse(wr.isOwnerLinked(owner), "Owner should not be linked");
  }

  function test_WhenTheProofWasPreviouslyUsed() external whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired {
    // it should revert with already used proof error
    (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory linkSignature = abi.encodePacked(r1, s1, v1);

    // link the owner using a unique proof
    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, true);
    wr.linkOwner(validityTimestamp, proof, linkSignature);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");

    (uint8 v2, bytes32 r2, bytes32 s2) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_UNLINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory unlinkSignature = abi.encodePacked(r2, s2, v2);

    // now unlink the owner from the registry
    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, unlinkSignature, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");

    // next, attempt to link the owner again using the same proof (this should fail because proof can't be reused)
    vm.prank(owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipProofAlreadyUsed.selector, owner, proof));
    wr.linkOwner(validityTimestamp, proof, linkSignature);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be still unlinked");

    address newOwner = address(0x5678);
    (uint8 v3, bytes32 r3, bytes32 s3) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), newOwner, validityTimestamp, proof)
    );
    bytes memory newLinkSignature = abi.encodePacked(r3, s3, v3);

    // now try to link a different owner with the same proof as before (this should also fail)
    vm.prank(newOwner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipProofAlreadyUsed.selector, newOwner, proof));
    wr.linkOwner(validityTimestamp, proof, newLinkSignature);
    assertFalse(wr.isOwnerLinked(newOwner), "Owner should be still unlinked");
  }

  function test_linkOwner_WhenTheTimestampHasExpired() external whenTheOwnerIsNotAlreadyLinked {
    // it should revert with expiration error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.LinkOwnerRequestExpired.selector, owner, block.timestamp, validityTimestamp
      )
    );
    wr.linkOwner(validityTimestamp, proof, sig);
    assertFalse(wr.isOwnerLinked(owner), "Owner should not be linked");
  }

  modifier whenTheOwnerIsAlreadyLinked() {
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(owner, proof, true);
    wr.linkOwner(validityTimestamp, proof, sig);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
    _;
  }

  function test_linkOwner_WhenTheTimestampIsStillValid() external whenTheOwnerIsAlreadyLinked {
    // it should revert with already linked error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkAlreadyExists.selector, owner));
    wr.linkOwner(validityTimestamp, proof, sig);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be already linked");
  }

  function test_linkOwner_WhenTheTimestampIsExpired() external whenTheOwnerIsAlreadyLinked {
    // it should revert with expired error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), owner, validityTimestamp, proof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.LinkOwnerRequestExpired.selector, owner, block.timestamp, validityTimestamp
      )
    );
    wr.linkOwner(validityTimestamp, proof, sig);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be already linked");
  }
}
