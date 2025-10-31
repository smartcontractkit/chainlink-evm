// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ReserveManager} from "../../../workflow/dev/ReserveManager.sol";
import {Test} from "forge-std/Test.sol";

// Minimal local interfaces so we can compute interfaceIds without chasing paths.
interface IERC165Local {
  function supportsInterface(
    bytes4 interfaceId
  ) external view returns (bool);
}

interface IReceiverLocal is IERC165Local {
  function onReport(bytes calldata metadata, bytes calldata report) external;
}

contract ReserveManagerTest is Test {
  // Re-declare the event signature so vm.expectEmit can match it.
  event RequestReserveUpdate(ReserveManager.UpdateReserves u);

  ReserveManager internal s_manager;

  function setUp() public {
    s_manager = new ReserveManager();
  }

  function testOnReportUpdatesStateAndEmitsEvent() public {
    // Arrange
    ReserveManager.UpdateReserves memory u = ReserveManager.UpdateReserves({totalMinted: 123, totalReserve: 456});
    bytes memory report = abi.encode(u);

    // Expect the exact event (no indexed args, just data) from the contract address
    vm.expectEmit(false, false, false, true, address(s_manager));
    emit RequestReserveUpdate(u);

    // Act (metadata is ignored by the contract)
    s_manager.onReport(bytes("ignored"), report);

    // Assert storage updated
    assertEq(s_manager.lastTotalMinted(), 123);
    assertEq(s_manager.lastTotalReserve(), 456);
  }

  function testSupportsInterface() public view {
    // ERC165 ID (0x01ffc9a7) via local interface
    bytes4 erc165Id = type(IERC165Local).interfaceId;
    assertTrue(s_manager.supportsInterface(erc165Id), "should support IERC165");

    // IReceiver-like interfaceId (same shape as your IReceiver)
    bytes4 iReceiverId = type(IReceiverLocal).interfaceId;
    assertTrue(s_manager.supportsInterface(iReceiverId), "should support IReceiver-like interface");

    // Random interface should be false
    assertFalse(s_manager.supportsInterface(0x12345678), "should not support random interface");
  }
}
