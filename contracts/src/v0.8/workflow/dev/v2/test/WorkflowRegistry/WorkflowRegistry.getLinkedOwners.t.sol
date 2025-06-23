// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getLinkedOwners is WorkflowRegistrySetup {
  address private s_owner1 = makeAddr("1");
  address private s_owner2 = makeAddr("2");
  address private s_owner3 = makeAddr("3");
  address private s_owner4 = makeAddr("4");
  address private s_owner5 = makeAddr("5");

  function test_getLinkedOwners_WhenThereAreNoLinkedOwners() external view {
    // it should return an empty result
    address[] memory owners = s_registry.getLinkedOwners(0, 10);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = s_registry.getLinkedOwners(0, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = s_registry.getLinkedOwners(0, 0);
    assertEq(owners.length, 0, "Expected no linked owners");
  }

  modifier whenThereAreLinkedOwners() {
    _linkOwner(s_owner1);
    _linkOwner(s_owner2);
    _linkOwner(s_owner3);
    _linkOwner(s_owner4);
    _linkOwner(s_owner5);
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
    address[] memory owners = s_registry.getLinkedOwners(0, 1);
    assertEq(owners.length, 1, "Expected one linked owner");
    assertEq(owners[0], s_owner1, "Expected first linked owner to be s_owner1");

    owners = s_registry.getLinkedOwners(0, 2);
    assertEq(owners.length, 2, "Expected two linked owners");
    assertEq(owners[0], s_owner1, "Expected first linked owner to be s_owner1");
    assertEq(owners[1], s_owner2, "Expected second linked owner to be s_owner2");
  }

  function test_getLinkedOwners_GivenThatBatchSizeIsEqualToTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return all linked owners
    address[] memory owners = s_registry.getLinkedOwners(0, 5);
    assertEq(owners.length, 5, "Expected five linked owners");
    assertEq(owners[0], s_owner1, "Expected first linked owner to be s_owner1");
    assertEq(owners[1], s_owner2, "Expected second linked owner to be s_owner2");
    assertEq(owners[2], s_owner3, "Expected third linked owner to be s_owner3");
    assertEq(owners[3], s_owner4, "Expected fourth linked owner to be s_owner4");
    assertEq(owners[4], s_owner5, "Expected fifth linked owner to be s_owner5");
  }

  function test_getLinkedOwners_GivenThatBatchSizeIsGreaterThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    givenThatStartIndexIsZero
  {
    // it should return the list of all linked owners
    address[] memory owners = s_registry.getLinkedOwners(0, 10);
    assertEq(owners.length, 5, "Expected five linked owners");
    assertEq(owners[0], s_owner1, "Expected first linked owner to be s_owner1");
    assertEq(owners[1], s_owner2, "Expected second linked owner to be s_owner2");
    assertEq(owners[2], s_owner3, "Expected third linked owner to be s_owner3");
    assertEq(owners[3], s_owner4, "Expected fourth linked owner to be s_owner4");
    assertEq(owners[4], s_owner5, "Expected fifth linked owner to be s_owner5");
  }

  modifier whenTheStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners() {
    _;
  }

  function test_getLinkedOwners_WhenBatchSizeIsLessThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    whenTheStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return some linked owners
    address[] memory owners = s_registry.getLinkedOwners(1, 2);
    assertEq(owners.length, 2, "Expected two linked owners");
    assertEq(owners[0], s_owner2, "Expected first linked owner to be s_owner2");
    assertEq(owners[1], s_owner3, "Expected second linked owner to be s_owner3");

    owners = s_registry.getLinkedOwners(2, 3);
    assertEq(owners.length, 3, "Expected three linked owners");
    assertEq(owners[0], s_owner3, "Expected first linked owner to be s_owner3");
    assertEq(owners[1], s_owner4, "Expected second linked owner to be s_owner4");
    assertEq(owners[2], s_owner5, "Expected third linked owner to be s_owner5");
  }

  function test_getLinkedOwners_WhenBatchSizeIsEqualToTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    whenTheStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return complete list of linked owners
    address[] memory owners = s_registry.getLinkedOwners(1, 5);
    assertEq(owners.length, 4, "Expected four linked owners");
    assertEq(owners[0], s_owner2, "Expected first linked owner to be s_owner2");
    assertEq(owners[1], s_owner3, "Expected second linked owner to be s_owner3");
    assertEq(owners[2], s_owner4, "Expected third linked owner to be s_owner4");
    assertEq(owners[3], s_owner5, "Expected fourth linked owner to be s_owner5");
  }

  function test_getLinkedOwners_WhenBatchSizeIsGreaterThanTotalLinkedOwners()
    external
    whenThereAreLinkedOwners
    whenTheStartIndexIsGreaterThanZeroAndLessThanTotalLinkedOwners
  {
    // it should return entire list of linked owners
    address[] memory owners = s_registry.getLinkedOwners(1, 10);
    assertEq(owners.length, 4, "Expected four linked owners");
    assertEq(owners[0], s_owner2, "Expected first linked owner to be s_owner2");
    assertEq(owners[1], s_owner3, "Expected second linked owner to be s_owner3");
    assertEq(owners[2], s_owner4, "Expected third linked owner to be s_owner4");
    assertEq(owners[3], s_owner5, "Expected fourth linked owner to be s_owner5");
  }

  function test_getLinkedOwners_GivenThatStartIndexIsEqualToTotalLinkedOwners() external whenThereAreLinkedOwners {
    // it should return an empty array
    address[] memory owners = s_registry.getLinkedOwners(5, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = s_registry.getLinkedOwners(5, 10);
    assertEq(owners.length, 0, "Expected no linked owners");
  }

  function test_getLinkedOwners_GivenThatStartIndexIsGreaterThanTotalLinkedOwners() external whenThereAreLinkedOwners {
    // it should return an empty list
    address[] memory owners = s_registry.getLinkedOwners(6, 1);
    assertEq(owners.length, 0, "Expected no linked owners");

    owners = s_registry.getLinkedOwners(10, 10);
    assertEq(owners.length, 0, "Expected no linked owners");
  }
}
