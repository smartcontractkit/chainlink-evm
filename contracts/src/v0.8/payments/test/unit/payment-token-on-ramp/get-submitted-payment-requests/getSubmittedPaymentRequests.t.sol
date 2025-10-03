// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {PaymentTokenOnRamp} from "../../../../PaymentTokenOnRamp.sol";
import {Common} from "../../../../libraries/Common.sol";
import {BaseUnitTest} from "../../BaseUnitTest.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.0.2/interfaces/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts@5.0.2/utils/cryptography/MessageHashUtils.sol";

contract PaymentTokenOnRamp_getSubmittedPaymentRequestsUnitTest is BaseUnitTest {
  uint256 private constant TOTAL_REQUESTS = 5;

  bytes32[] private s_submittedPaymentRequests;

  function setUp() public {
    vm.mockCall(i_token1, abi.encodeWithSelector(IERC20.approve.selector), abi.encode(true));
    vm.mockCall(i_token1, abi.encodeWithSelector(IERC20.transferFrom.selector), abi.encode(true));

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](TOTAL_REQUESTS);

    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](1);
    tokenAmounts[0] = Common.AssetAmount({asset: i_token1, amount: 1 ether});

    for (uint256 i = 0; i < TOTAL_REQUESTS; ++i) {
      bytes32 requestId = keccak256(abi.encode(i));
      s_submittedPaymentRequests.push(requestId);

      paymentRequests[i].requestId = requestId;
      paymentRequests[i].fundingAddress = i_owner;
      paymentRequests[i].deadline = uint40(block.timestamp);
      paymentRequests[i].tokenAmounts = tokenAmounts;

      bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
        keccak256(
          abi.encode(
            s_paymentTokenOnRamp.typeAndVersion(),
            CHAIN_SELECTOR,
            address(s_paymentTokenOnRamp),
            requestId,
            uint40(block.timestamp),
            i_owner,
            tokenAmounts
          )
        )
      );

      (bytes32 r, bytes32 s) = _sign(s_authorityPk, digest);

      paymentRequests[i].r = r;
      paymentRequests[i].s = s;
    }

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);
  }

  function test_getSubmittedPaymentRequests_StartGeLimit() external view {
    bytes32[] memory submittedPaymentRequests = s_paymentTokenOnRamp.getSubmittedPaymentRequests(10, 5);

    assertEq(submittedPaymentRequests.length, 0);
  }

  function test_getSubmittedPaymentRequests_StartPlusLimitGtTotalRequests() external view {
    bytes32[] memory submittedPaymentRequests = s_paymentTokenOnRamp.getSubmittedPaymentRequests(3, 5);

    assertEq(submittedPaymentRequests.length, 2);

    for (uint256 i = 0; i < 2; ++i) {
      assertEq(submittedPaymentRequests[i], s_submittedPaymentRequests[i + 3]);
    }
  }

  function test_getSubmittedPaymentRequests_StartPlusLimitLteTotalRequests() external view {
    bytes32[] memory submittedPaymentRequests = s_paymentTokenOnRamp.getSubmittedPaymentRequests(1, 3);

    assertEq(submittedPaymentRequests.length, 3);

    for (uint256 i = 0; i < 3; ++i) {
      assertEq(submittedPaymentRequests[i], s_submittedPaymentRequests[i + 1]);
    }
  }

  function test_getSubmittedPaymentRequests_StartPlusLimitEqTotalRequests() external view {
    bytes32[] memory submittedPaymentRequests = s_paymentTokenOnRamp.getSubmittedPaymentRequests(0, 5);

    assertEq(submittedPaymentRequests, submittedPaymentRequests);
  }

  function test_getSubmittedPaymentRequests_LimitEqZero() external view {
    bytes32[] memory submittedPaymentRequests = s_paymentTokenOnRamp.getSubmittedPaymentRequests(1, 0);

    assertEq(submittedPaymentRequests.length, 0);
  }
}
