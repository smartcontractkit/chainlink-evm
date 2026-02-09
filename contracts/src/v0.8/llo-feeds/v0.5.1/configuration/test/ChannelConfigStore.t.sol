// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import {ChannelConfigStore} from "../ChannelConfigStore.sol";
import {IChannelConfigStore} from "../interfaces/IChannelConfigStore.sol";
import {ExposedChannelConfigStore} from "./mocks/ExposedChannelConfigStore.sol";
import {Test} from "forge-std/Test.sol";

/**
 * @title ChannelConfigStoreTest
 * @author samsondav
 * @notice Base class for ChannelConfigStore tests
 */
contract ChannelConfigStoreTest is Test {
  ExposedChannelConfigStore public channelConfigStore;

  event NewChannelDefinition(uint256 indexed donId, uint32 version, string url, bytes32 sha);
  event ChannelDefinitionAdded(
    uint256 indexed donId, IChannelConfigStore.ChannelAdderId indexed channelAdderId, string url, bytes32 sha
  );
  event ChannelAdderSet(uint256 indexed donId, IChannelConfigStore.ChannelAdderId indexed channelAdderId, bool allowed);
  event ChannelAdderAddressSet(IChannelConfigStore.ChannelAdderId indexed channelAdderId, address adderAddress);

  address public constant CHANNEL_ADDER_1 = address(0x1234);
  address public constant CHANNEL_ADDER_2 = address(0x5678);
  IChannelConfigStore.ChannelAdderId public constant CHANNEL_ADDER_ID_1 = IChannelConfigStore.ChannelAdderId.wrap(1000);
  IChannelConfigStore.ChannelAdderId public constant CHANNEL_ADDER_ID_2 = IChannelConfigStore.ChannelAdderId.wrap(1001);
  uint32 public constant DON_ID_1 = 42;
  uint32 public constant DON_ID_2 = 99;

  function setUp() public virtual {
    channelConfigStore = new ExposedChannelConfigStore();
  }

  function testTypeAndVersion() public view {
    assertEq(channelConfigStore.typeAndVersion(), "ChannelConfigStore 1.0.0");
  }

  function testSupportsInterface() public view {
    assertTrue(channelConfigStore.supportsInterface(type(IChannelConfigStore).interfaceId));
  }

  function test_revertsIfCalledByNonOwner() public {
    vm.expectRevert("Only callable by owner");

    vm.startPrank(address(2));
    channelConfigStore.setChannelDefinitions(42, "url", keccak256("sha"));
  }

  function testSetChannelDefinitions() public {
    vm.expectEmit();
    emit NewChannelDefinition(42, 1, "url", keccak256("sha"));
    channelConfigStore.setChannelDefinitions(42, "url", keccak256("sha"));

    vm.expectEmit();
    emit NewChannelDefinition(42, 2, "url2", keccak256("sha2"));
    channelConfigStore.setChannelDefinitions(42, "url2", keccak256("sha2"));

    assertEq(channelConfigStore.exposedReadChannelDefinitionStates(42), uint32(2));
  }
}

/**
 * @title ChannelConfigStoreChannelAdderTest
 * @notice Test suite for channel adder functionality
 */
contract ChannelConfigStoreChannelAdderTest is ChannelConfigStoreTest {
  function testSetChannelAdderAddress() public {
    vm.expectEmit();
    emit ChannelAdderAddressSet(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);

    assertEq(channelConfigStore.getChannelAdderAddress(CHANNEL_ADDER_ID_1), CHANNEL_ADDER_1);
  }

  function testSetChannelAdderAddress_UpdatesExistingAddress() public {
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    assertEq(channelConfigStore.getChannelAdderAddress(CHANNEL_ADDER_ID_1), CHANNEL_ADDER_1);

    vm.expectEmit();
    emit ChannelAdderAddressSet(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_2);
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_2);

    assertEq(channelConfigStore.getChannelAdderAddress(CHANNEL_ADDER_ID_1), CHANNEL_ADDER_2);
  }

  function testSetChannelAdderAddress_CanSetToZeroAddress() public {
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);

    vm.expectEmit();
    emit ChannelAdderAddressSet(CHANNEL_ADDER_ID_1, address(0));
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, address(0));

    assertEq(channelConfigStore.getChannelAdderAddress(CHANNEL_ADDER_ID_1), address(0));
  }

  function testSetChannelAdderAddress_RevertsWhenCalledByNonOwner() public {
    vm.expectRevert("Only callable by owner");
    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
  }

  function testSetChannelAdderAddress_RevertsForReservedChannelAdderId() public {
    // Channel adder IDs 0-999 are reserved
    IChannelConfigStore.ChannelAdderId reservedId = IChannelConfigStore.ChannelAdderId.wrap(0);
    vm.expectRevert(ChannelConfigStore.ReservedChannelAdderId.selector);
    channelConfigStore.setChannelAdderAddress(reservedId, CHANNEL_ADDER_1);

    reservedId = IChannelConfigStore.ChannelAdderId.wrap(999);
    vm.expectRevert(ChannelConfigStore.ReservedChannelAdderId.selector);
    channelConfigStore.setChannelAdderAddress(reservedId, CHANNEL_ADDER_1);
  }

  function testSetChannelAdderAddress_SucceedsAtMinimumAllowedId() public {
    // Channel adder ID 1000 is the minimum allowed
    IChannelConfigStore.ChannelAdderId minAllowedId = IChannelConfigStore.ChannelAdderId.wrap(1000);
    channelConfigStore.setChannelAdderAddress(minAllowedId, CHANNEL_ADDER_1);
    assertEq(channelConfigStore.getChannelAdderAddress(minAllowedId), CHANNEL_ADDER_1);
  }

  function testSetChannelAdder_AllowsChannelAdder() public {
    vm.expectEmit();
    emit ChannelAdderSet(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);

    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_1));
  }

  function testSetChannelAdder_RemovesChannelAdder() public {
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_1));

    vm.expectEmit();
    emit ChannelAdderSet(DON_ID_1, CHANNEL_ADDER_ID_1, false);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, false);

    assertFalse(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_1));
  }

  function testSetChannelAdder_AllowsMultipleAddersPerDon() public {
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_2, true);

    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_1));
    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_2));

    IChannelConfigStore.ChannelAdderId[] memory allowedAdders = channelConfigStore.getAllowedChannelAdders(DON_ID_1);
    assertEq(allowedAdders.length, 2);

    // Assert that allowedAdders contains both CHANNEL_ADDER_ID_1 and CHANNEL_ADDER_ID_2 (order agnostic)
    assertTrue(
      (
        IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[0])
          == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_1)
          && IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[1])
            == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_2)
      )
        || (
          IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[0])
            == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_2)
            && IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[1])
              == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_1)
        )
    );
  }

  function testSetChannelAdder_AdderCanBeAllowedOnMultipleDons() public {
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    channelConfigStore.setChannelAdder(DON_ID_2, CHANNEL_ADDER_ID_1, true);

    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_1));
    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_2, CHANNEL_ADDER_ID_1));
  }

  function testSetChannelAdder_RevertsWhenCalledByNonOwner() public {
    vm.expectRevert("Only callable by owner");
    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
  }

  function testGetAllowedChannelAdders_ReturnsEmptyArrayWhenNone() public view {
    IChannelConfigStore.ChannelAdderId[] memory allowedAdders = channelConfigStore.getAllowedChannelAdders(DON_ID_1);
    assertEq(allowedAdders.length, 0);
  }

  function testGetAllowedChannelAdders_ReturnsAllAllowedAdders() public {
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_2, true);

    IChannelConfigStore.ChannelAdderId[] memory allowedAdders = channelConfigStore.getAllowedChannelAdders(DON_ID_1);
    assertEq(allowedAdders.length, 2);

    // EnumerableSet doesn't guarantee order, so check both combinations
    assertTrue(
      (
        IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[0])
          == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_1)
          && IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[1])
            == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_2)
      )
        || (
          IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[0])
            == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_2)
            && IChannelConfigStore.ChannelAdderId.unwrap(allowedAdders[1])
              == IChannelConfigStore.ChannelAdderId.unwrap(CHANNEL_ADDER_ID_1)
        )
    );
  }

  function testIsChannelAdderAllowed_ReturnsFalseForNonAllowedAdder() public view {
    assertFalse(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_1));
  }

  function testAddChannelDefinitions_SucceedsWithAuthorizedAdder() public {
    // Setup: set address and allow adder
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);

    vm.expectEmit();
    emit ChannelDefinitionAdded(DON_ID_1, CHANNEL_ADDER_ID_1, "url", keccak256("sha"));

    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url", keccak256("sha"));
  }

  function testAddChannelDefinitions_RevertsWhenCallerNotMatchingChannelAdderId() public {
    // Setup: set address and allow adder
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);

    // Try to call from wrong address
    vm.expectRevert(ChannelConfigStore.UnauthorizedChannelAdder.selector);
    vm.prank(CHANNEL_ADDER_2);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url", keccak256("sha"));
  }

  function testAddChannelDefinitions_RevertsWhenAdderNotAllowedForDon() public {
    // Setup: set address but don't allow adder for this DON
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);

    vm.expectRevert(ChannelConfigStore.UnauthorizedChannelAdder.selector);
    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url", keccak256("sha"));
  }

  function testAddChannelDefinitions_RevertsWhenNoAddressSetForChannelAdderId() public {
    // Setup: allow adder but don't set address
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);

    vm.expectRevert(ChannelConfigStore.UnauthorizedChannelAdder.selector);
    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url", keccak256("sha"));
  }

  function testAddChannelDefinitions_RevertsWhenAdderRemovedFromDon() public {
    // Setup: set address and allow adder
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);

    // Verify it works
    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url", keccak256("sha"));

    // Remove adder
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, false);

    // Should now revert
    vm.expectRevert(ChannelConfigStore.UnauthorizedChannelAdder.selector);
    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url2", keccak256("sha2"));
  }

  function testAddChannelDefinitions_MultipleAddersCanAddToDon() public {
    // Setup: set addresses and allow both adders
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_2, CHANNEL_ADDER_2);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_2, true);

    // Both should be able to add
    vm.prank(CHANNEL_ADDER_1);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url1", keccak256("sha1"));

    vm.prank(CHANNEL_ADDER_2);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_2, "url2", keccak256("sha2"));
  }

  function testAddChannelDefinitions_AdderCanAddToMultipleDons() public {
    // Setup: set address and allow adder on multiple DONs
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    channelConfigStore.setChannelAdder(DON_ID_2, CHANNEL_ADDER_ID_1, true);

    // Should work for both DONs
    vm.startPrank(CHANNEL_ADDER_1);
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url1", keccak256("sha1"));
    channelConfigStore.addChannelDefinitions(DON_ID_2, CHANNEL_ADDER_ID_1, "url2", keccak256("sha2"));
    vm.stopPrank();
  }

  function testGetChannelAdderAddress_ReturnsZeroAddressForUnsetId() public view {
    assertEq(channelConfigStore.getChannelAdderAddress(CHANNEL_ADDER_ID_1), address(0));
  }

  function testAddChannelDefinitions_OneAddressCanControlMultipleChannelAdders() public {
    // Setup: one address controls two different channel adder IDs
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_1, CHANNEL_ADDER_1);
    channelConfigStore.setChannelAdderAddress(CHANNEL_ADDER_ID_2, CHANNEL_ADDER_1); // same address

    // Allow both channel adder IDs for the DON
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_1, true);
    channelConfigStore.setChannelAdder(DON_ID_1, CHANNEL_ADDER_ID_2, true);

    // Verify the same address can use both channel adder IDs
    vm.startPrank(CHANNEL_ADDER_1);

    vm.expectEmit();
    emit ChannelDefinitionAdded(DON_ID_1, CHANNEL_ADDER_ID_1, "url1", keccak256("sha1"));
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_1, "url1", keccak256("sha1"));

    vm.expectEmit();
    emit ChannelDefinitionAdded(DON_ID_1, CHANNEL_ADDER_ID_2, "url2", keccak256("sha2"));
    channelConfigStore.addChannelDefinitions(DON_ID_1, CHANNEL_ADDER_ID_2, "url2", keccak256("sha2"));

    vm.stopPrank();

    // Verify both channel adder IDs are allowed
    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_1));
    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, CHANNEL_ADDER_ID_2));
  }

  function testSetChannelAdder_RevertsForReservedChannelAdderId() public {
    // Channel adder IDs 0-999 are reserved
    IChannelConfigStore.ChannelAdderId reservedId = IChannelConfigStore.ChannelAdderId.wrap(0);
    vm.expectRevert(ChannelConfigStore.ReservedChannelAdderId.selector);
    channelConfigStore.setChannelAdder(DON_ID_1, reservedId, true);

    reservedId = IChannelConfigStore.ChannelAdderId.wrap(999);
    vm.expectRevert(ChannelConfigStore.ReservedChannelAdderId.selector);
    channelConfigStore.setChannelAdder(DON_ID_1, reservedId, true);

    // Should also revert when trying to remove a reserved ID
    vm.expectRevert(ChannelConfigStore.ReservedChannelAdderId.selector);
    channelConfigStore.setChannelAdder(DON_ID_1, reservedId, false);
  }

  function testSetChannelAdder_SucceedsAtMinimumAllowedId() public {
    // Channel adder ID 1000 is the minimum allowed
    IChannelConfigStore.ChannelAdderId minAllowedId = IChannelConfigStore.ChannelAdderId.wrap(1000);
    channelConfigStore.setChannelAdder(DON_ID_1, minAllowedId, true);
    assertTrue(channelConfigStore.isChannelAdderAllowed(DON_ID_1, minAllowedId));
  }
}
