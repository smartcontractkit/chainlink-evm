// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IRouter} from "../../../keystone/interfaces/IRouter.sol";
import {MockKeystoneForwarder} from "../../../workflow/dev/MockKeystoneForwarder.sol";
import {Test} from "forge-std/Test.sol";

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

  function test_TypeAndVersion() public view {
    assertEq(s_forwarder.typeAndVersion(), "MockKeystoneForwarder 1.0.0-dev");
  }

  function test_IsForwarder_SelfTrue_AddRemove() public {
    // constructor marks itself
    assertTrue(s_forwarder.isForwarder(address(s_forwarder)));

    address other = address(0xBEEF);
    s_forwarder.addForwarder(other);
    assertTrue(s_forwarder.isForwarder(other));

    s_forwarder.removeForwarder(other);
    assertFalse(s_forwarder.isForwarder(other));

    // random not marked
    assertFalse(s_forwarder.isForwarder(address(1234)));
  }

  // ---------------------------
  //     getTransmissionInfo / getTransmitter / getTransmissionId
  // ---------------------------

  function test_GetTransmissionInfo_NotAttempted_And_GetTransmitter_Zero() public view {
    // Before any route/report, NOT_ATTEMPTED branch should be taken
    bytes32 execId = keccak256("no-attempt");
    bytes2 repId = 0x0001;

    IRouter.TransmissionInfo memory info = s_forwarder.getTransmissionInfo(address(s_good), execId, repId);

    assertEq(info.transmitter, address(0), "expected zero transmitter");
    assertFalse(info.success, "success should be false by default");
    assertEq(info.gasLimit, 0, "gasLimit should be zero before attempt");
    assertEq(uint8(info.state), uint8(IRouter.TransmissionState.NOT_ATTEMPTED), "state mismatch");

    assertEq(s_forwarder.getTransmitter(address(s_good), execId, repId), address(0), "expected zero transmitter");
  }

  function test_GetTransmissionId_MatchesManual() public view {
    address receiver = address(0xA11CE);
    bytes32 execId = keccak256("manual");
    bytes2 repId = 0x1234;

    // Contract helper
    bytes32 got = s_forwarder.getTransmissionId(receiver, execId, repId);

    // Manual computation: keccak256(bytes20(receiver) ++ execId ++ repId)
    bytes32 expected = keccak256(bytes.concat(bytes20(uint160(receiver)), execId, repId));

    assertEq(got, expected, "getTransmissionId mismatch");
  }

  // ---------------------------
  //              route()
  // ---------------------------

  function test_Route_Direct_Success() public {
    bytes32 execId = keccak256("route-ok");
    bytes2 repId = 0x0ACE;
    bytes32 tid = s_forwarder.getTransmissionId(address(s_good), execId, repId);

    bool success = s_forwarder.route(tid, address(this), address(s_good), bytes("m"), bytes("v"));
    assertTrue(success, "route should succeed");

    IRouter.TransmissionInfo memory info = s_forwarder.getTransmissionInfo(address(s_good), execId, repId);

    assertEq(info.transmitter, address(this), "transmitter mismatch");
    assertTrue(info.success, "expected success");
    assertGt(info.gasLimit, 0, "gas limit should be recorded");
    assertEq(uint8(info.state), uint8(IRouter.TransmissionState.SUCCEEDED), "state mismatch");
  }

  function test_Route_Direct_Failure() public {
    bytes32 execId = keccak256("route-fail");
    bytes2 repId = 0x0BAD;
    bytes32 tid = s_forwarder.getTransmissionId(address(s_bad), execId, repId);

    bool success = s_forwarder.route(tid, address(this), address(s_bad), bytes("m"), bytes("v"));
    assertFalse(success, "route should fail");

    IRouter.TransmissionInfo memory info = s_forwarder.getTransmissionInfo(address(s_bad), execId, repId);

    assertEq(info.transmitter, address(this), "transmitter mismatch");
    assertFalse(info.success, "expected failure");
    assertGt(info.gasLimit, 0, "gas limit should be recorded");
    assertEq(uint8(info.state), uint8(IRouter.TransmissionState.FAILED), "state mismatch");
  }
}
