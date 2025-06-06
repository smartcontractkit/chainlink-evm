// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {OwnershipLink} from "../../OwnershipLink.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";

import {Test} from "forge-std/Test.sol";

contract WorkflowRegistryUnlinkOwner is Test {
  address public owner = address(0xabcd);
  WorkflowRegistry public wr;
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

  modifier whenTheOwnerIsLinked() {
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(wr), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit OwnershipLink.OwnershipLinkUpdatedV1(owner, proof, true);
    wr.linkOwner(validityTimestamp, proof, sig);
    assertTrue(wr.isOwnerLinked(owner), "Owner should be linked");
    _;
  }

  modifier givenThatTheOwnerSelectsNonePreAction() {
    _;
  }

  function test_GivenThatTheOwnerDoesNotHaveActiveWorkflows()
    external
    whenTheOwnerIsLinked
    givenThatTheOwnerSelectsNonePreAction
  {
    // it should unlink the owner
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(wr), owner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);

    vm.prank(owner);
    vm.expectEmit(true, true, true, false);
    emit OwnershipLink.OwnershipLinkUpdatedV1(owner, proof, false);
    wr.unlinkOwner(owner, validityTimestamp, proof, sig, WorkflowRegistry.PreUnlinkAction.NONE);
    assertFalse(wr.isOwnerLinked(owner), "Owner should be unlinked");
  }

  function test_GivenThatTheOwnerHasActiveWorkflows()
    external
    whenTheOwnerIsLinked
    givenThatTheOwnerSelectsNonePreAction
  {
    // it should revert with active workflows error
  }

  function test_GivenThatTheOwnerSelectsPausingWorkflows() external whenTheOwnerIsLinked {
    // it should pause all workflows and unlink the owner
  }

  function test_GivenThatTheOwnerSelectsRemovingWorkflows() external whenTheOwnerIsLinked {
    // it should remove all workflows and unlink the owner
  }

  function test_WhenTheOwnerIsNotLinked() external {
    // it should revert with not linked error
  }
}
