// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {LinkingUtils} from "../../testutils/LinkingUtils.sol";
import {OwnershipLinkTestable} from "../../testutils/OwnershipLinkTestable.sol";

import "forge-std/Test.sol";

contract OwnershipLinkGetLinkedOwners is Test {
  OwnershipLinkTestable op;
  address owner = address(0xdddd);
  uint256 allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
  address allowedSigner;
  uint256 validityTimestamp = uint256(block.timestamp + 1 hours);
  bytes32 proof = keccak256("test-proof");
  address owner1 = address(0xabcd);
  address owner2 = address(0x1234);
  address owner3 = address(0x5678);
  address owner4 = address(0xdef0);
  address owner5 = address(0x1111);

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

  function test_WhenThereAreNoLinkedOwners() external {
    // it should return an empty result
    address[] memory owners = op.getLinkedOwners(0, 10);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = op.getLinkedOwners(0, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = op.getLinkedOwners(0, 0);
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

  function test_GivenThatBatchSizeIsLessThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return the first batch of linked owners
    address[] memory owners = op.getLinkedOwners(0, 1);
    assertEq(owners.length, 1, "Expected one linked owner");
    assertEq(owners[0], owner1, "Expected first linked owner to be owner1");

    owners = op.getLinkedOwners(0, 2);
    assertEq(owners.length, 2, "Expected two linked owners");
    assertEq(owners[0], owner1, "Expected first linked owner to be owner1");
    assertEq(owners[1], owner2, "Expected second linked owner to be owner2");
  }

  function test_GivenThatBatchSizeIsEqualToTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return all linked owners
    address[] memory owners = op.getLinkedOwners(0, 5);
    assertEq(owners.length, 5, "Expected five linked owners");
    assertEq(owners[0], owner1, "Expected first linked owner to be owner1");
    assertEq(owners[1], owner2, "Expected second linked owner to be owner2");
    assertEq(owners[2], owner3, "Expected third linked owner to be owner3");
    assertEq(owners[3], owner4, "Expected fourth linked owner to be owner4");
    assertEq(owners[4], owner5, "Expected fifth linked owner to be owner5");
  }

  function test_GivenThatBatchSizeIsGreaterThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return the list of all linked owners
    address[] memory owners = op.getLinkedOwners(0, 10);
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

  function test_WhenBatchSizeIsLessThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return some linked owners
    address[] memory owners = op.getLinkedOwners(1, 2);
    assertEq(owners.length, 2, "Expected two linked owners");
    assertEq(owners[0], owner2, "Expected first linked owner to be owner2");
    assertEq(owners[1], owner3, "Expected second linked owner to be owner3");

    owners = op.getLinkedOwners(2, 3);
    assertEq(owners.length, 3, "Expected three linked owners");
    assertEq(owners[0], owner3, "Expected first linked owner to be owner3");
    assertEq(owners[1], owner4, "Expected second linked owner to be owner4");
    assertEq(owners[2], owner5, "Expected third linked owner to be owner5");
  }

  function test_WhenBatchSizeIsEqualToTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return complete list of linked owners
    address[] memory owners = op.getLinkedOwners(1, 5);
    assertEq(owners.length, 4, "Expected four linked owners");
    assertEq(owners[0], owner2, "Expected first linked owner to be owner2");
    assertEq(owners[1], owner3, "Expected second linked owner to be owner3");
    assertEq(owners[2], owner4, "Expected third linked owner to be owner4");
    assertEq(owners[3], owner5, "Expected fourth linked owner to be owner5");
  }

  function test_WhenBatchSizeIsGreaterThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return entire list of linked owners
    address[] memory owners = op.getLinkedOwners(1, 10);
    assertEq(owners.length, 4, "Expected four linked owners");
    assertEq(owners[0], owner2, "Expected first linked owner to be owner2");
    assertEq(owners[1], owner3, "Expected second linked owner to be owner3");
    assertEq(owners[2], owner4, "Expected third linked owner to be owner4");
    assertEq(owners[3], owner5, "Expected fourth linked owner to be owner5");
  }

  function test_GivenThatStartIndexIsEqualToTotalLinkedOwners() external whenThereAreLinkedOwners {
    // it should return an empty array
    address[] memory owners = op.getLinkedOwners(5, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = op.getLinkedOwners(5, 10);
    assertEq(owners.length, 0, "Expected no linked owners");
  }

  function test_GivenThatStartIndexIsGreaterThanTotalLinkedOwners() external whenThereAreLinkedOwners {
    // it should return an empty list
    address[] memory owners = op.getLinkedOwners(6, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = op.getLinkedOwners(10, 10);
    assertEq(owners.length, 0, "Expected no linked owners");
  }

  // Helper to link an owner
  function linkOwner(
    address newOwner
  ) public {
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(allowedSignerPrivateKey, LinkingUtils.getMessageHash(address(op), newOwner, validityTimestamp, proof));
    bytes memory sig = abi.encodePacked(r, s, v);
    vm.prank(newOwner);
    op.linkOwner(validityTimestamp, proof, sig);
  }
}
