// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Test} from "forge-std/Test.sol";
import {ShardConfig} from "../ShardConfig.sol";

contract ShardConfigTest is Test {
  ShardConfig internal s_shardConfig;

  address internal constant MCMS_ADDRESS = address(0x1234567890123456789012345678901234567890);
  address internal constant NEW_MCMS_ADDRESS = address(0x0987654321098765432109876543210987654321);
  address internal constant NON_MCMS_ADDRESS = address(0xDEADBEEF);
  uint256 internal constant INITIAL_SHARD_COUNT = 10;

  function setUp() public virtual {
    s_shardConfig = new ShardConfig(INITIAL_SHARD_COUNT, MCMS_ADDRESS);
  }

  function testTypeAndVersion() public view {
    assertEq(s_shardConfig.typeAndVersion(), "ShardConfig 1.0.0");
  }

  function testConstructor_Success() public view {
    assertEq(s_shardConfig.desiredShardCount(), INITIAL_SHARD_COUNT);
    assertEq(s_shardConfig.mcmsAddress(), MCMS_ADDRESS);
  }

  function testConstructor_RevertZeroMCMS() public {
    vm.expectRevert("Invalid MCMS address");
    new ShardConfig(INITIAL_SHARD_COUNT, address(0));
  }

  function testConstructor_RevertZeroShardCount() public {
    vm.expectRevert("Shard count must be greater than 0");
    new ShardConfig(0, MCMS_ADDRESS);
  }

  function testSetDesiredShardCount_Success() public {
    uint256 newCount = 20;

    vm.prank(MCMS_ADDRESS);
    s_shardConfig.setDesiredShardCount(newCount);

    assertEq(s_shardConfig.desiredShardCount(), newCount);
    assertEq(s_shardConfig.getDesiredShardCount(), newCount);
  }

  function testSetDesiredShardCount_EmitsEvent() public {
    uint256 newCount = 20;

    vm.expectEmit(true, false, false, false);
    emit ShardConfig.ShardCountUpdated(newCount);

    vm.prank(MCMS_ADDRESS);
    s_shardConfig.setDesiredShardCount(newCount);
  }

  function testSetDesiredShardCount_RevertNotMCMS() public {
    vm.prank(NON_MCMS_ADDRESS);
    vm.expectRevert("Only MCMS can update shard count");
    s_shardConfig.setDesiredShardCount(20);
  }

  function testSetDesiredShardCount_RevertZeroCount() public {
    vm.prank(MCMS_ADDRESS);
    vm.expectRevert("Shard count must be greater than 0");
    s_shardConfig.setDesiredShardCount(0);
  }

  function testSetMCMSAddress_Success() public {
    vm.prank(MCMS_ADDRESS);
    s_shardConfig.setMCMSAddress(NEW_MCMS_ADDRESS);

    assertEq(s_shardConfig.mcmsAddress(), NEW_MCMS_ADDRESS);
  }

  function testSetMCMSAddress_EmitsEvent() public {
    vm.expectEmit(true, false, false, false);
    emit ShardConfig.MCMSAddressUpdated(NEW_MCMS_ADDRESS);

    vm.prank(MCMS_ADDRESS);
    s_shardConfig.setMCMSAddress(NEW_MCMS_ADDRESS);
  }

  function testSetMCMSAddress_RevertNotMCMS() public {
    vm.prank(NON_MCMS_ADDRESS);
    vm.expectRevert("Only MCMS can update its own address");
    s_shardConfig.setMCMSAddress(NEW_MCMS_ADDRESS);
  }

  function testSetMCMSAddress_RevertZeroAddress() public {
    vm.prank(MCMS_ADDRESS);
    vm.expectRevert("Invalid MCMS address");
    s_shardConfig.setMCMSAddress(address(0));
  }

  function testGetDesiredShardCount() public view {
    assertEq(s_shardConfig.getDesiredShardCount(), INITIAL_SHARD_COUNT);
  }
}