// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {OwnershipLinkTestable} from "./OwnershipLinkTestable.sol";

import "forge-std/Test.sol";

library LinkingUtils {
  // Helper to get the EIP-191 message hash
  function getMessageHash(
    address linkContract,
    address owner,
    uint256 validityTimestamp,
    bytes32 proof
  ) public view returns (bytes32) {
    bytes32 messageHash = keccak256(abi.encodePacked(owner, block.chainid, linkContract, validityTimestamp, proof));
    return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash));
  }

  // Helper to link an owner
  function linkOwner(
    Vm vm,
    address linkContract,
    uint256 allowedSignerPrivateKey,
    address owner,
    uint256 validityTimestamp,
    bytes32 proof
  ) public {
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, getMessageHash(linkContract, owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);
    vm.prank(owner);
    OwnershipLinkTestable op = OwnershipLinkTestable(linkContract);
    op.linkOwner(validityTimestamp, proof, sig);
  }
}
