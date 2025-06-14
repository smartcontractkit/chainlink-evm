// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";

import {Test} from "forge-std/Test.sol";

contract WorkflowRegistry_getLinkedOwners is Test {
  WorkflowRegistry public wr;
  address public owner = address(0xdddd);
  uint256 public allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
  address public allowedSigner;
  uint256 public validityTimestamp = uint256(block.timestamp + 1 hours);
  bytes32 public proofSeed = keccak256("test-proof");
  address public owner1 = address(0xabcd);
  address public owner2 = address(0x1234);
  address public owner3 = address(0x5678);
  address public owner4 = address(0xdef0);
  address public owner5 = address(0x1111);

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

  function test_getLinkedOwners_WhenThereAreNoLinkedOwners() external view {
    // it should return an empty result
    address[] memory owners = wr.getLinkedOwners(0, 10);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = wr.getLinkedOwners(0, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = wr.getLinkedOwners(0, 0);
    assertEq(owners.length, 0, "Expected no linked owners");
  }

  modifier whenThereAreLinkedOwners() {
    linkOwner(owner1);
    linkOwner(owner2);
    linkOwner(owner3);
    linkOwner(owner4);
    linkOwner(owner5);
    _;
  }

  modifier givenThatStartIndexIsZero() {
    _;
  }

  function test_getLinkedOwners_GivenThatBatchSizeIsLessThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return the first batch of linked owners
    address[] memory owners = wr.getLinkedOwners(0, 1);
    assertEq(owners.length, 1, "Expected one linked owner");
    assertEq(owners[0], owner1, "Expected first linked owner to be owner1");

    owners = wr.getLinkedOwners(0, 2);
    assertEq(owners.length, 2, "Expected two linked owners");
    assertEq(owners[0], owner1, "Expected first linked owner to be owner1");
    assertEq(owners[1], owner2, "Expected second linked owner to be owner2");
  }

  function test_getLinkedOwners_GivenThatBatchSizeIsEqualToTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return all linked owners
    address[] memory owners = wr.getLinkedOwners(0, 5);
    assertEq(owners.length, 5, "Expected five linked owners");
    assertEq(owners[0], owner1, "Expected first linked owner to be owner1");
    assertEq(owners[1], owner2, "Expected second linked owner to be owner2");
    assertEq(owners[2], owner3, "Expected third linked owner to be owner3");
    assertEq(owners[3], owner4, "Expected fourth linked owner to be owner4");
    assertEq(owners[4], owner5, "Expected fifth linked owner to be owner5");
  }

  function test_getLinkedOwners_GivenThatBatchSizeIsGreaterThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return the list of all linked owners
    address[] memory owners = wr.getLinkedOwners(0, 10);
    assertEq(owners.length, 5, "Expected five linked owners");
    assertEq(owners[0], owner1, "Expected first linked owner to be owner1");
    assertEq(owners[1], owner2, "Expected second linked owner to be owner2");
    assertEq(owners[2], owner3, "Expected third linked owner to be owner3");
    assertEq(owners[3], owner4, "Expected fourth linked owner to be owner4");
    assertEq(owners[4], owner5, "Expected fifth linked owner to be owner5");
  }

  modifier givenThatStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners() {
    _;
  }

  function test_getLinkedOwners_WhenBatchSizeIsLessThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return some linked owners
    address[] memory owners = wr.getLinkedOwners(1, 2);
    assertEq(owners.length, 2, "Expected two linked owners");
    assertEq(owners[0], owner2, "Expected first linked owner to be owner2");
    assertEq(owners[1], owner3, "Expected second linked owner to be owner3");

    owners = wr.getLinkedOwners(2, 3);
    assertEq(owners.length, 3, "Expected three linked owners");
    assertEq(owners[0], owner3, "Expected first linked owner to be owner3");
    assertEq(owners[1], owner4, "Expected second linked owner to be owner4");
    assertEq(owners[2], owner5, "Expected third linked owner to be owner5");
  }

  function test_getLinkedOwners_WhenBatchSizeIsEqualToTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return complete list of linked owners
    address[] memory owners = wr.getLinkedOwners(1, 5);
    assertEq(owners.length, 4, "Expected four linked owners");
    assertEq(owners[0], owner2, "Expected first linked owner to be owner2");
    assertEq(owners[1], owner3, "Expected second linked owner to be owner3");
    assertEq(owners[2], owner4, "Expected third linked owner to be owner4");
    assertEq(owners[3], owner5, "Expected fourth linked owner to be owner5");
  }

  function test_getLinkedOwners_WhenBatchSizeIsGreaterThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return entire list of linked owners
    address[] memory owners = wr.getLinkedOwners(1, 10);
    assertEq(owners.length, 4, "Expected four linked owners");
    assertEq(owners[0], owner2, "Expected first linked owner to be owner2");
    assertEq(owners[1], owner3, "Expected second linked owner to be owner3");
    assertEq(owners[2], owner4, "Expected third linked owner to be owner4");
    assertEq(owners[3], owner5, "Expected fourth linked owner to be owner5");
  }

  function test_getLinkedOwners_GivenThatStartIndexIsEqualToTotalLinkedOwners() external whenThereAreLinkedOwners {
    // it should return an empty array
    address[] memory owners = wr.getLinkedOwners(5, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = wr.getLinkedOwners(5, 10);
    assertEq(owners.length, 0, "Expected no linked owners");
  }

  function test_getLinkedOwners_GivenThatStartIndexIsGreaterThanTotalLinkedOwners() external whenThereAreLinkedOwners {
    // it should return an empty list
    address[] memory owners = wr.getLinkedOwners(6, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = wr.getLinkedOwners(10, 10);
    assertEq(owners.length, 0, "Expected no linked owners");
  }

  // Helper to link an owner
  function linkOwner(
    address newOwner
  ) public {
    bytes32 ownerProof = keccak256(abi.encode(proofSeed, newOwner));
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(LinkingUtils.REQUEST_TYPE_LINK, address(wr), newOwner, validityTimestamp, ownerProof)
    );
    bytes memory sig = abi.encodePacked(r, s, v);
    vm.prank(newOwner);
    wr.linkOwner(validityTimestamp, ownerProof, sig);
  }
}
