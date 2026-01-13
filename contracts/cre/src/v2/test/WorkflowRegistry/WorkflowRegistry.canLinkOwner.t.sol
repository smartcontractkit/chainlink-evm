// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_canLinkOwner is WorkflowRegistrySetup {
  function test_canLinkOwner_WhenBlockTimestampIsGreaterThanValidityTimestamp() external {
    // It should revert with LinkOwnerRequestExpired
    // block time has advanced by 24 hours so the validity timestamp is in the past
    (bytes32 ownerProof, bytes memory sig) = _getLinkProofSignature(s_owner);
    vm.warp(block.timestamp + 24 hours);
    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.LinkOwnerRequestExpired.selector, s_owner, block.timestamp, s_validityTimestamp
      )
    );
    s_registry.canLinkOwner(s_owner, s_validityTimestamp, ownerProof, sig);
  }

  // whenSignatureHasNotExpired
  function test_canLinkOwner_WhenMsgSenderIsAlreadyLinked() external {
    // It should revert with OwnershipLinkAlreadyExists
    _linkOwner(s_owner);
    (bytes32 ownerProof, bytes memory sig) = _getLinkProofSignature(s_owner);
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkAlreadyExists.selector, s_owner));
    s_registry.canLinkOwner(s_owner, s_validityTimestamp, ownerProof, sig);
  }

  // whenSignatureHasNotExpired whenMsgSenderIsNotYetLinked
  function test_canLinkOwner_WhenProofHasAlreadyBeenUsed() external {
    // It should revert with OwnershipProofAlreadyUsed
    address anotherUser = makeAddr("another-user");
    (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), anotherUser, s_validityTimestamp, s_proof
      )
    );

    bytes memory sig1 = abi.encodePacked(r1, s1, v1);
    // link the first user
    vm.prank(anotherUser);
    s_registry.linkOwner(s_validityTimestamp, s_proof, sig1);

    // now call canLinkOwner for s_owner with the same proof
    (uint8 v2, bytes32 r2, bytes32 s2) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );

    bytes memory sig2 = abi.encodePacked(r2, s2, v2);
    vm.prank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipProofAlreadyUsed.selector, s_owner, s_proof));
    s_registry.canLinkOwner(s_owner, s_validityTimestamp, s_proof, sig2);
  }

  // whenBlockTimestampIsLessOrEqualToValidityTimestamp
  // whenMsgSenderIsNotYetLinked
  // whenProofIsUnused
  function test_canLinkOwner_WhenSignatureRecoveryFails() external {
    // It should revert with InvalidSignature
    // build a garbage signature that's too short
    bytes memory badSig = hex"abcd";
    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.InvalidSignature.selector,
        badSig,
        2, // InvalidSignatureLength is the error code from ECDSA.RecoverError with errId = 2
        bytes32(uint256(2)) // bytes32(uint256(2)) is the length of hex"abcd",
      )
    );
    s_registry.canLinkOwner(s_owner, s_validityTimestamp, s_proof, badSig);
  }

  // whenSignatureHasNotExpired whenMsgSenderIsNotYetLinked whenProofIsUnused
  function test_canLinkOwner_WhenSignatureRecoversToASignerNotInAllowedSigners() external {
    // It should revert with InvalidOwnershipLink
    uint256 randomPrivateKey = 0x7f3c2a9b5d4e1f8c0b2d3a4e5f6c7d8e9a0b1c2d3e4f5a6b7c8d9e0f1a2b3c4d;
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      randomPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), s_owner, s_validityTimestamp, s_proof
      )
    );

    bytes memory sig = abi.encodePacked(r, s, v);
    vm.prank(s_owner);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.InvalidOwnershipLink.selector, s_owner, s_validityTimestamp, s_proof, sig)
    );
    s_registry.canLinkOwner(s_owner, s_validityTimestamp, s_proof, sig);
  }

  // whenSignatureHasNotExpired whenMsgSenderIsNotYetLinked whenProofIsUnused
  function test_canLinkOwner_WhenSignatureIsValidAndSignerIsAllowed() external {
    // It should return (no revert)
    (bytes32 ownerProof, bytes memory sig) = _getLinkProofSignature(s_owner);
    vm.prank(s_owner);
    s_registry.canLinkOwner(s_owner, s_validityTimestamp, ownerProof, sig);
  }
}
