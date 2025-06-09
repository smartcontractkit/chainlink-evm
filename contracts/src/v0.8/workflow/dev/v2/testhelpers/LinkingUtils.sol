// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

library LinkingUtils {
  uint8 constant REQUEST_TYPE_LINK = 0;
  uint8 constant REQUEST_TYPE_UNLINK = 1;

  // Helper to get the EIP-191 message hash
  function getMessageHash(
    uint8 requestType,
    address linkContract,
    address owner,
    uint256 validityTimestamp,
    bytes32 proof
  ) public view returns (bytes32) {
    bytes32 messageHash =
      keccak256(abi.encode(requestType, owner, block.chainid, linkContract, validityTimestamp, proof));
    return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash));
  }
}
