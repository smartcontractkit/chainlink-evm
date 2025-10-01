// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {MockKeystoneForwarder} from "../MockKeystoneForwarder.sol";
import {IRouter} from "../interfaces/IRouter.sol";
import "forge-std/Test.sol";

// Minimal receivers to exercise success/failure paths.
contract GoodReceiver {
  event OnReport(bytes metadata, bytes validatedReport);

  function onReport(bytes calldata metadata, bytes calldata validatedReport) external {
    emit OnReport(metadata, validatedReport);
  }
}

contract BadReceiver {
  error Oops();

  function onReport(bytes calldata, bytes calldata) external pure {
    revert Oops();
  }
}

contract MockKeystoneForwarder_BasicTest is Test {
  event ReportProcessed(
    address indexed receiver, bytes32 indexed workflowExecutionId, bytes2 indexed reportId, bool result
  );

  MockKeystoneForwarder internal s_forwarder;
  GoodReceiver internal s_good;
  BadReceiver internal s_bad;

  function setUp() public {
    s_forwarder = new MockKeystoneForwarder();
    s_good = new GoodReceiver();
    s_bad = new BadReceiver();
  }

  // ---------------------------
  //          Basics
  // ---------------------------

  function test_TypeAndVersion() public {
    assertEq(s_forwarder.typeAndVersion(), "MockKeystoneForwarder 1.0.0");
  }

  function test_IsForwarder_SelfTrue_AddRemove() public {
    // Constructor marks itself
    assertTrue(s_forwarder.isForwarder(address(s_forwarder)));

    address other = address(0xBEEF);
    s_forwarder.addForwarder(other);
    assertTrue(s_forwarder.isForwarder(other));

    s_forwarder.removeForwarder(other);
    assertFalse(s_forwarder.isForwarder(other));
  }
}
