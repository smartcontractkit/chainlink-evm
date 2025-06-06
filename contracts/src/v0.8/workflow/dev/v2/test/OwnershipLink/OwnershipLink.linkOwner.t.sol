// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {OwnershipLink} from "../../OwnershipLink.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";
import {OwnershipLinkTestable} from "../../testhelpers/OwnershipLinkTestable.sol";

import {ECDSA} from "../../../../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/cryptography/ECDSA.sol";

import {Test} from "forge-std/Test.sol";

contract OwnershipLinkLinkOwner is Test {
  OwnershipLinkTestable public op;
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
    op = new OwnershipLinkTestable();
    address[] memory signers = new address[](1);
    signers[0] = allowedSigner;
    op.updateAllowedSigners(signers, true);
    vm.stopPrank();
  }

  modifier whenTheOwnerIsNotAlreadyLinked() {
    _;
  }

  modifier whenTheTimestampHasNotExpired() {
    _;
  }

  function test_WhenProofIsValid() external whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired {
    // it should link the owner
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(op), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit OwnershipLink.OwnershipLinkUpdatedV1(owner, proof, true);
    op.linkOwner(validityTimestamp, proof, sig);
    assertTrue(op.isOwnerLinked(owner), "Owner should be linked");
  }

  function test_WhenTheProofIsNotSignedByAnAllowedSigner()
    external
    whenTheOwnerIsNotAlreadyLinked
    whenTheTimestampHasNotExpired
  {
    // it should revert with signature error
    uint256 unknownSignerPrivateKey = 0xffc0c927f94d71f7c5c21a865d7c47d050a34f1583ba93576edf67cf2fa32da7;
    address unknownSigner = vm.addr(unknownSignerPrivateKey);
    assertEq(unknownSigner, address(0xfF2B8E43743892d9a8416254711A473b8B70DDe4));

    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(unknownSignerPrivateKey, LinkingUtils.getMessageHash(address(op), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(OwnershipLink.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    op.linkOwner(validityTimestamp, proof, sig);
    assertFalse(op.isOwnerLinked(owner), "Owner should not be linked");
  }

  function test_WhenTheProofContainsInvalidData() external whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired {
    // it should revert with invalid signature error
    address invalidOwner = address(0x1234);
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(op), invalidOwner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(OwnershipLink.InvalidOwnershipLink.selector, owner, validityTimestamp, proof, sig)
    );
    op.linkOwner(validityTimestamp, proof, sig);
    assertFalse(op.isOwnerLinked(owner), "Owner should not be linked");
  }

  function test_WhenTheSignatureIsNotValid() external whenTheOwnerIsNotAlreadyLinked whenTheTimestampHasNotExpired {
    // it should revert with internal signature error
    bytes memory invalidSignature = "invalid-signature";

    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(
        OwnershipLink.InvalidSignature.selector, invalidSignature, ECDSA.RecoverError.InvalidSignatureLength, 0x11
      )
    );
    op.linkOwner(validityTimestamp, proof, invalidSignature);
    assertFalse(op.isOwnerLinked(owner), "Owner should not be linked");
  }

  function test_WhenTheTimestampHasExpired() external whenTheOwnerIsNotAlreadyLinked {
    // it should revert with expiration error
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(op), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(OwnershipLink.LinkOwnerRequestExpired.selector, owner, block.timestamp, validityTimestamp)
    );
    op.linkOwner(validityTimestamp, proof, sig);
    assertFalse(op.isOwnerLinked(owner), "Owner should not be linked");
  }

  modifier whenTheOwnerIsAlreadyLinked() {
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(op), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit OwnershipLink.OwnershipLinkUpdatedV1(owner, proof, true);
    op.linkOwner(validityTimestamp, proof, sig);
    assertTrue(op.isOwnerLinked(owner), "Owner should be linked");
    _;
  }

  function test_WhenTheTimestampIsStillValid() external whenTheOwnerIsAlreadyLinked {
    // it should revert with already linked error
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(op), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectRevert(abi.encodeWithSelector(OwnershipLink.OwnershipLinkAlreadyExists.selector, owner));
    op.linkOwner(validityTimestamp, proof, sig);
    assertTrue(op.isOwnerLinked(owner), "Owner should be already linked");
  }

  function test_WhenTheTimestampIsExpired() external whenTheOwnerIsAlreadyLinked {
    // it should revert with expired error
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(op), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    // block time has advanced by 24 hours so the validity timestamp is in the past
    vm.warp(block.timestamp + 24 hours);
    vm.prank(owner);
    vm.expectRevert(
      abi.encodeWithSelector(OwnershipLink.LinkOwnerRequestExpired.selector, owner, block.timestamp, validityTimestamp)
    );
    op.linkOwner(validityTimestamp, proof, sig);
    assertTrue(op.isOwnerLinked(owner), "Owner should be already linked");
  }
}
