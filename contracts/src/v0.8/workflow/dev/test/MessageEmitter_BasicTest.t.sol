// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {MessageEmitter} from "../../../workflow/dev/MessageEmitter.sol";
import {Test} from "forge-std/Test.sol";

contract MessageEmitter_BasicTest is Test {
  // Mirror the contract event for expectEmit
  event MessageEmitted(address indexed emitter, uint256 indexed timestamp, string message);

  MessageEmitter internal s_emitter;

  address internal s_ALICE = address(0xA11CE);
  address internal s_BOB = address(0xB0B);

  function setUp() public {
    s_emitter = new MessageEmitter();
  }

  function test_TypeAndVersion() public view {
    assertEq(s_emitter.typeAndVersion(), "MessageEmitter 1.0.0-dev");
  }

  function test_RevertWhen_EmptyMessage() public {
    vm.startPrank(s_ALICE);
    vm.expectRevert(bytes("Message cannot be empty"));
    s_emitter.emitMessage("");
    vm.stopPrank();
  }

  function test_EmitMessage_EmitsEventAndUpdatesLastMessage() public {
    uint256 t = 1_000_000;
    vm.warp(t);

    vm.startPrank(s_ALICE);

    vm.expectEmit(address(s_emitter));
    emit MessageEmitted(s_ALICE, t, "hello");

    s_emitter.emitMessage("hello");

    // getLastMessage should return what we just set
    string memory last = s_emitter.getLastMessage(s_ALICE);
    assertEq(last, "hello", "last message mismatch");

    vm.stopPrank();
  }

  function test_RevertWhen_DuplicateInSameBlockTimestamp() public {
    uint256 t = 1_234_567;
    vm.warp(t);

    vm.startPrank(s_ALICE);
    s_emitter.emitMessage("first in block");

    // Same sender, same block timestamp -> must revert
    vm.expectRevert(bytes("Message already exists for the same sender and block timestamp"));
    s_emitter.emitMessage("second in same block");
    vm.stopPrank();
  }

  function test_EmitMessage_SucceedsInNextBlock() public {
    uint256 t = 42;
    vm.warp(t);

    vm.prank(s_ALICE);
    s_emitter.emitMessage("first");

    // Different block timestamp -> allowed
    vm.warp(t + 1);
    vm.prank(s_ALICE);
    s_emitter.emitMessage("second");

    // Last message should now be "second"
    string memory last = s_emitter.getLastMessage(s_ALICE);
    assertEq(last, "second");
  }

  function test_EmitMessage_MultipleAddressesSameBlock() public {
    uint256 t = 999;
    vm.warp(t);

    vm.prank(s_ALICE);
    s_emitter.emitMessage("alice says hi");

    vm.prank(s_BOB);
    s_emitter.emitMessage("bob says hi");

    assertEq(s_emitter.getLastMessage(s_ALICE), "alice says hi");
    assertEq(s_emitter.getLastMessage(s_BOB), "bob says hi");
  }

  function test_RevertWhen_GetLastMessage_NoMessage() public {
    vm.expectRevert(bytes("No last message for the given sender"));
    s_emitter.getLastMessage(s_ALICE);
  }

  function test_GetMessage_Succeeds_WhenExists() public {
    uint256 t = 111_222;
    vm.warp(t);
    vm.prank(s_ALICE);
    s_emitter.emitMessage("stored");
    string memory got = s_emitter.getMessage(s_ALICE, t);
    assertEq(got, "stored");
  }

  function test_RevertWhen_GetMessage_NotExists() public {
    vm.expectRevert(bytes("Message does not exist for the given sender and timestamp"));
    s_emitter.getMessage(s_ALICE, 555);
  }
}
