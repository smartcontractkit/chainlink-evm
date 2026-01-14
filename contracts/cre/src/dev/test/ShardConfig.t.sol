// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Test} from "forge-std/Test.sol";
import {ShardConfig} from "../../dev/ShardConfig.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract ShardConfigTest is Test {
  ShardConfig internal s_shardConfig;

  address internal constant NON_OWNER_ADDRESS = address(0xDEADBEEF);
  uint256 internal constant INITIAL_SHARD_COUNT = 10;

  function setUp() public virtual {
    s_shardConfig = new ShardConfig(INITIAL_SHARD_COUNT);
  }

  function testTypeAndVersion() public view {
    assertEq(s_shardConfig.typeAndVersion(), "ShardConfig 1.0.0-dev");
  }

  function testConstructor_Success() public view {
    assertEq(s_shardConfig.desiredShardCount(), INITIAL_SHARD_COUNT);
    assertEq(s_shardConfig.owner(), address(this));
  }

  function testConstructor_RevertZeroShardCount() public {
    vm.expectRevert("Shard count must be greater than 0");
    new ShardConfig(0);
  }

  function testSetDesiredShardCount_Success() public {
    uint256 newCount = 20;

    s_shardConfig.setDesiredShardCount(newCount);

    assertEq(s_shardConfig.desiredShardCount(), newCount);
    assertEq(s_shardConfig.getDesiredShardCount(), newCount);
  }

  function testSetDesiredShardCount_EmitsEvent() public {
    uint256 newCount = 20;

    vm.expectEmit(true, false, false, false);
    emit ShardConfig.ShardCountUpdated(newCount);

    s_shardConfig.setDesiredShardCount(newCount);
  }

  function testSetDesiredShardCount_RevertNotOwner() public {
    vm.prank(NON_OWNER_ADDRESS);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_shardConfig.setDesiredShardCount(20);
  }

  function testSetDesiredShardCount_RevertZeroCount() public {
    vm.expectRevert("Shard count must be greater than 0");
    s_shardConfig.setDesiredShardCount(0);
  }

  function testGetDesiredShardCount() public view {
    assertEq(s_shardConfig.getDesiredShardCount(), INITIAL_SHARD_COUNT);
  }

  function testOwnershipTransfer() public {
    address newOwner = address(0x1234);

    // Step 1: Current owner initiates transfer
    s_shardConfig.transferOwnership(newOwner);

    // Owner hasn't changed yet
    assertEq(s_shardConfig.owner(), address(this));

    // Step 2: New owner accepts
    vm.prank(newOwner);
    s_shardConfig.acceptOwnership();

    // Now ownership has transferred
    assertEq(s_shardConfig.owner(), newOwner);
  }
}
