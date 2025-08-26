// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

import {ECDSA} from "@openzeppelin/contracts@5.1.0/utils/cryptography/ECDSA.sol";

contract WorkflowRegistry_linkOwner is WorkflowRegistrySetup {
  // whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired
  function test_linkOwner_WhenProofIsValid() external {
    // it should link the s_owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(s_owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, s_proof, true);
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
  }

  // whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired
  function test_linkOwner_WhenTheProofIsNotSignedByAnAllowedSigner() external {
    // it should revert with signature error
    uint256 unknownSignerPrivateKey = 0xffc0c927f94d71f7c5c21a865d7c47d050a34f1583ba93576edf67cf2fa32da7;
    address unknownSigner = vm.addr(unknownSignerPrivateKey);
    assertEq(unknownSigner, address(0xfF2B8E43743892d9a8416254711A473b8B70DDe4));

    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      unknownSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, s_owner, s_validityTimestamp, s_proof, sig)
    );
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should not be linked");
  }

  // whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired
  function test_linkOwner_WhenTheProofContainsInvalidData() external {
    // it should revert with invalid signature error
    address invalidOwner = address(0x1234);
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), invalidOwner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, s_owner, s_validityTimestamp, s_proof, sig)
    );
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should not be linked");
  }

  // whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired
  function test_linkOwner_WhenTheSignatureIsNotValid() external {
    // it should revert with internal signature error
    bytes memory invalidSignature = "invalid-signature";

    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidSignature.selector, invalidSignature, ECDSA.RecoverError.InvalidSignatureLength, 0x11
      )
    );
    s_registry.linkOwner(s_validityTimestamp, s_proof, invalidSignature);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should not be linked");
  }

  // whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired
  function test_WhenTheProofWasPreviouslyUsed() external {
    // it should revert with already used s_proof error
    (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory linkSignature = abi.encodePacked(r1, s1, v1);

    // link the s_owner using a unique s_proof
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, s_proof, true);
    s_registry.linkOwner(s_validityTimestamp, s_proof, linkSignature);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");

    (uint8 v2, bytes32 r2, bytes32 s2) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory unlinkSignature = abi.encodePacked(r2, s2, v2);

    // now unlink the s_owner from the registry
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, s_proof, false);
    s_registry.unlinkOwner(s_owner, s_validityTimestamp, unlinkSignature);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should be unlinked");

    // next, attempt to link the s_owner again using the same s_proof (this should fail because s_proof can't be reused)
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipProofAlreadyUsed.selector, s_owner, s_proof));
    s_registry.linkOwner(s_validityTimestamp, s_proof, linkSignature);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should be still unlinked");

    address newOwner = address(0x5678);
    (uint8 v3, bytes32 r3, bytes32 s3) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), newOwner, s_validityTimestamp, s_proof
      )
    );
    bytes memory newLinkSignature = abi.encodePacked(r3, s3, v3);

    // now try to link a different s_owner with the same s_proof as before (this should also fail)
    vm.prank(newOwner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipProofAlreadyUsed.selector, newOwner, s_proof));
    s_registry.linkOwner(s_validityTimestamp, s_proof, newLinkSignature);
    assertFalse(s_registry.isOwnerLinked(newOwner), "Owner should be still unlinked");
  }

  // whenTheOwnerIsNotAlreadyLinked
  function test_linkOwner_WhenTheTimestampHasExpired() external {
    // it should revert with expiration error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.LinkOwnerRequestExpired.selector, s_owner, block.timestamp, s_validityTimestamp
      )
    );
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig);
    assertFalse(s_registry.isOwnerLinked(s_owner), "Owner should not be linked");
  }

  modifier whenTheOwnerIsAlreadyLinked() {
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(s_owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.OwnershipLinkUpdated(s_owner, s_proof, true);
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be linked");
    _;
  }

  function test_linkOwner_WhenTheTimestampIsStillValid() external whenTheOwnerIsAlreadyLinked {
    // it should revert with already linked error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkAlreadyExists.selector, s_owner));
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be already linked");
  }

  function test_linkOwner_WhenTheTimestampIsExpired() external whenTheOwnerIsAlreadyLinked {
    // it should revert with expired error
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.LinkOwnerRequestExpired.selector, s_owner, block.timestamp, s_validityTimestamp
      )
    );
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig);
    assertTrue(s_registry.isOwnerLinked(s_owner), "Owner should be already linked");
  }
}
