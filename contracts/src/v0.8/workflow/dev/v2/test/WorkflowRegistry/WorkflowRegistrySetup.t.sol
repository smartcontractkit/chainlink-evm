// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";
import {Test} from "forge-std/Test.sol";

contract WorkflowRegistrySetup is Test {
  WorkflowRegistry internal s_registry;
  address internal s_owner;
  address internal s_stranger;
  address internal s_user;
  uint256 internal s_allowedSignerPrivateKey;
  address internal s_allowedSigner;
  uint256 internal s_validityTimestamp;
  bytes32 internal s_proof;
  bytes32 internal s_proofSeed;
  bytes internal s_signature;
  string internal s_donLabel;
  string internal s_binaryUrl;
  string internal s_configUrl;
  string internal s_tag;
  string internal s_workflowName;
  bytes32 internal s_workflowId;
  bytes internal s_attributes;
  string internal s_invalidLongString;
  string internal s_invalidURL;

  function setUp() public virtual {
    s_owner = makeAddr("owner");
    s_stranger = makeAddr("nonOwner");
    s_allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
    s_validityTimestamp = uint256(block.timestamp + 1 hours);
    s_proof = keccak256("test-proof");
    s_allowedSigner = vm.addr(s_allowedSignerPrivateKey);
    assertEq(s_allowedSigner, address(0x86f2cE81640Fd86e68CF3EB25c2801D6E1C62bd0));

    s_user = makeAddr("user");
    s_donLabel = "DON-A";
    s_binaryUrl = "ipfs://bin";
    s_configUrl = "ipfs://cfg";
    s_tag = "alpha";

    s_workflowName = "my-workflow";
    s_workflowId = keccak256("workflow1");
    s_attributes = hex"11223344556677889900aabbccddeeff";
    s_invalidLongString =
      "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcd";
    s_invalidURL =
      "https://www.example.com/this/is/a/very/long/url/that/keeps/going/on/and/on/to/ensure/that/it/exceeds/two/hundred/and/one/characters/in/length/for/testing/purposes/and/it/should/be/sufficiently/long/to/meet/your/requirements/for/this/test";

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

    s_signature = abi.encodePacked(r, s, v);
    vm.prank(newOwner);
    s_registry.linkOwner(s_validityTimestamp, ownerProof, s_signature);
  }
}
