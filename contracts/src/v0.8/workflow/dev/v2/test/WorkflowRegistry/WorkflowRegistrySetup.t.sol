// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";
import {Test} from "forge-std/Test.sol";

contract WorkflowRegistrySetup is Test {
  WorkflowRegistry internal s_registry;
  address internal s_owner;
  address internal s_stranger;
  uint256 internal s_allowedSignerPrivateKey;
  address internal s_allowedSigner;
  uint256 internal s_validityTimestamp;
  bytes32 internal s_proof;
  bytes32 public s_proofSeed;

  bytes32 internal s_donLabel;
  string internal s_binaryURL;
  string internal s_configURL;
  string internal s_tag;

  function setUp() public virtual {
    s_owner = makeAddr("owner");
    s_stranger = makeAddr("nonOwner");
    s_allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
    s_validityTimestamp = uint256(block.timestamp + 1 hours);
    s_proof = keccak256("test-proof");
    s_allowedSigner = vm.addr(s_allowedSignerPrivateKey);
    assertEq(s_allowedSigner, address(0x86f2cE81640Fd86e68CF3EB25c2801D6E1C62bd0));

    s_donLabel = bytes32("DON-A");
    s_binaryURL = "ipfs://bin";
    s_configURL = "ipfs://cfg";
    s_tag = "alpha";

    vm.startPrank(s_owner);
    s_registry = new WorkflowRegistry();
    address[] memory signers = new address[](1);
    signers[0] = s_allowedSigner;
    s_registry.updateAllowedSigners(signers, true);
    vm.stopPrank();
  }

  // Helper to link an owner
  function _linkOwner(
    address newOwner
  ) public {
    bytes32 ownerProof = keccak256(abi.encode(s_proof, newOwner));
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), newOwner, s_validityTimestamp, ownerProof
      )
    );
    bytes memory sig = abi.encodePacked(r, s, v);
    vm.prank(newOwner);
    s_registry.linkOwner(s_validityTimestamp, ownerProof, sig);
  }
}
