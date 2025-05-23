// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

import {Ownable2Step} from "../../shared/access/Ownable2StepMsgSender.sol";
import {OwnershipProof} from "../dev/OwnershipProof.sol";
import "forge-std/Test.sol";

// ,inimal inheriting contract for testing
contract OwnershipProofTestable is OwnershipProof {}

contract OwnershipProofTest is Test {
  OwnershipProofTestable op;
  address owner = address(0xabcd);
  uint256 allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
  address allowedSigner;
  uint96 validityTimestamp;
  bytes32 proof = keccak256("test-proof");

  function setUp() public {
    // hardcode the signer's private key into test environment (so that vm.sign can be used)
    allowedSigner = vm.addr(allowedSignerPrivateKey);
    assertEq(allowedSigner, address(0x86f2cE81640Fd86e68CF3EB25c2801D6E1C62bd0));

    vm.startPrank(owner);
    op = new OwnershipProofTestable();
    address[] memory signers = new address[](1);
    signers[0] = allowedSigner;
    op.updateAllowedSigners(signers, true);
    validityTimestamp = uint96(block.timestamp + 1 hours);
    vm.stopPrank();
  }

  function testUpdateAllowedSigners() public {
    vm.prank(owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0xbeef);
    op.updateAllowedSigners(signers, true);
    assertTrue(op.isAllowedSigner(address(0xbeef)));
  }

  function testSubmitOwnershipProof() public {
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);

    assertTrue(op.isProofSubmitted(owner));
  }

  function testRevertIfProofAlreadySubmitted() public {
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);

    vm.expectRevert(abi.encodeWithSelector(OwnershipProof.ProofAlreadySubmitted.selector, owner));
    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);
  }

  function testRevertSubmitProofIfRequestExpired() public {
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.warp(block.timestamp + 24 hours);

    vm.expectRevert(abi.encodeWithSelector(OwnershipProof.RequestExpired.selector, owner, validityTimestamp));
    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);
  }

  function testRevertRevokeProofIfRequestExpired() public {
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.warp(block.timestamp + 24 hours);

    vm.expectRevert(abi.encodeWithSelector(OwnershipProof.RequestExpired.selector, owner, validityTimestamp));
    vm.prank(owner);
    op.revokeOwnershipProof(validityTimestamp, proof, sig);
  }

  function testRevokeOwnershipProof() public {
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);

    vm.prank(owner);
    op.revokeOwnershipProof(validityTimestamp, proof, sig);

    assertFalse(op.isProofSubmitted(owner));
  }

  function testGetOwnersWithSubmittedProofSingle() public {
    // Submit proof for owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);

    address[] memory owners = op.getOwnersWithSubmittedProof(0, 10);
    assertEq(owners.length, 1);
    assertEq(owners[0], owner);

    // no owners from index 1 onwards
    for (uint256 i = 1; i < 10; i++) {
      owners = op.getOwnersWithSubmittedProof(i, 10);
      assertEq(owners.length, 0);
    }
  }

  function testGetOwnersWithSubmittedProofMultiple() public {
    address owner2 = address(0x2222);
    address owner3 = address(0x3333);

    // Submit proofs for multiple owners
    (uint8 v1, bytes32 r1, bytes32 s1) =
      vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig1 = abi.encodePacked(r1, s1, v1);

    (uint8 v2, bytes32 r2, bytes32 s2) =
      vm.sign(allowedSignerPrivateKey, _getMessageHash(owner2, validityTimestamp, proof));
    bytes memory sig2 = abi.encodePacked(r2, s2, v2);

    (uint8 v3, bytes32 r3, bytes32 s3) =
      vm.sign(allowedSignerPrivateKey, _getMessageHash(owner3, validityTimestamp, proof));
    bytes memory sig3 = abi.encodePacked(r3, s3, v3);

    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig1);
    vm.prank(owner2);
    op.submitOwnershipProof(validityTimestamp, proof, sig2);
    vm.prank(owner3);
    op.submitOwnershipProof(validityTimestamp, proof, sig3);

    // Batch size larger than total
    address[] memory owners = op.getOwnersWithSubmittedProof(0, 10);
    assertEq(owners.length, 3);
    assertEq(owners[0], owner);
    assertEq(owners[1], owner2);
    assertEq(owners[2], owner3);

    // Batch size 2, start 0
    owners = op.getOwnersWithSubmittedProof(0, 2);
    assertEq(owners.length, 2);
    assertEq(owners[0], owner);
    assertEq(owners[1], owner2);

    // Batch size 2, start 1
    owners = op.getOwnersWithSubmittedProof(1, 2);
    assertEq(owners.length, 2);
    assertEq(owners[0], owner2);
    assertEq(owners[1], owner3);

    // Batch size 2, start 2
    owners = op.getOwnersWithSubmittedProof(2, 2);
    assertEq(owners.length, 1);
    assertEq(owners[0], owner3);

    // Start > total
    owners = op.getOwnersWithSubmittedProof(3, 2);
    assertEq(owners.length, 0);
  }

  function testGetOwnersWithSubmittedProofEmpty() public {
    address[] memory owners = op.getOwnersWithSubmittedProof(0, 10);
    assertEq(owners.length, 0);
  }

  function testAdminRevokeOwnershipProof() public {
    // Submit proof for owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);

    // Admin revokes proof
    vm.prank(owner);
    vm.expectEmit(true, true, false, true);
    emit OwnershipProof.OwnershipProofRevokedV1(owner, bytes32(0));
    op.adminRevokeOwnershipProof(owner);

    assertFalse(op.isProofSubmitted(owner));
  }

  function testAdminRevokeOwnershipProofRevertsIfNoProof() public {
    vm.prank(owner);
    vm.expectRevert(abi.encodeWithSelector(OwnershipProof.ProofNotSubmitted.selector, owner));
    op.adminRevokeOwnershipProof(owner);
  }

  function testAdminRevokeOwnershipProofOnlyOwner() public {
    address notOwner = address(0xBEEF);

    // Submit proof for owner
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(allowedSignerPrivateKey, _getMessageHash(owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    op.submitOwnershipProof(validityTimestamp, proof, sig);

    // Try to revoke as not owner
    vm.prank(notOwner);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    op.adminRevokeOwnershipProof(owner);
  }

  // Helper to get the EIP-191 message hash
  function _getMessageHash(address _owner, uint96 _validityTimestamp, bytes32 _proof) public pure returns (bytes32) {
    bytes32 messageHash = keccak256(abi.encodePacked(_owner, _validityTimestamp, _proof));
    return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash));
  }
}
