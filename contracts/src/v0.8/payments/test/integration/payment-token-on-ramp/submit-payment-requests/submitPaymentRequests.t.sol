// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {PaymentTokenOnRamp} from "../../../../PaymentTokenOnRamp.sol";
import {Common} from "../../../../libraries/Common.sol";
import {Errors} from "../../../../libraries/Errors.sol";
import {BaseIntegrationTest} from "../../BaseIntegrationTest.t.sol";

import {Pausable} from "@openzeppelin/contracts@5.0.2/utils/Pausable.sol";
import {MessageHashUtils} from "@openzeppelin/contracts@5.0.2/utils/cryptography/MessageHashUtils.sol";

contract PaymentTokenOnRamp_submitPaymentRequestsUnitTest is BaseIntegrationTest {
  uint256 private constant ETH_PAYMENT_AMOUNT = 1 ether;
  uint256 private constant USDC_PAYMENT_AMOUNT = 10e6;

  bytes32 private s_requestId = keccak256("payment request");

  function setUp() public {
    deal(address(s_mockWETH), i_owner, ETH_PAYMENT_AMOUNT);
    deal(address(s_mockUSDC), i_owner, USDC_PAYMENT_AMOUNT);
    s_mockWETH.approve(address(s_paymentTokenOnRamp), ETH_PAYMENT_AMOUNT);
    s_mockUSDC.approve(address(s_paymentTokenOnRamp), USDC_PAYMENT_AMOUNT);
  }

  function test_submitPaymentRequests_RevertWhen_ContractIsPaused()
    external
    givenContractIsPaused(address(s_paymentTokenOnRamp))
  {
    vm.expectRevert(Pausable.EnforcedPause.selector);

    s_paymentTokenOnRamp.submitPaymentRequests(new PaymentTokenOnRamp.PaymentRequest[](0));
  }

  function test_submitPaymentRequests_RevertWhen_PaymentRequestLengthIsZero() external {
    vm.expectRevert(Errors.EmptyList.selector);

    s_paymentTokenOnRamp.submitPaymentRequests(new PaymentTokenOnRamp.PaymentRequest[](0));
  }

  function test_submitPaymentRequests_RevertWhen_FundingAddressEqAddressZero() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](1);
    tokenAmounts[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: ETH_PAYMENT_AMOUNT});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);

    vm.expectRevert(Errors.InvalidZeroAddress.selector);

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);
  }

  function test_submitPaymentRequests_RevertWhen_RequestHasExpired() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](1);
    tokenAmounts[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: ETH_PAYMENT_AMOUNT});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);

    skip(1);

    vm.expectRevert(
      abi.encodeWithSelector(PaymentTokenOnRamp.ExpiredRequest.selector, s_requestId, block.timestamp - 1)
    );
    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);
  }

  function test_submitPaymentRequests_RevertWhen_RequestHasAlreadyBeenSubmitted() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](1);
    tokenAmounts[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: ETH_PAYMENT_AMOUNT});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);
    paymentRequests[0].tokenAmounts = tokenAmounts;

    bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          s_requestId,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts
        )
      )
    );

    (bytes32 r, bytes32 s) = _sign(s_authorityPk, digest);

    paymentRequests[0].r = r;
    paymentRequests[0].s = s;

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);

    vm.expectRevert(abi.encodeWithSelector(PaymentTokenOnRamp.RequestAlreadySubmitted.selector, s_requestId));
    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);
  }

  function test_submitPaymentRequests_RevertWhen_SignatureIsNotValid() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](1);
    tokenAmounts[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: ETH_PAYMENT_AMOUNT});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);
    paymentRequests[0].tokenAmounts = tokenAmounts;

    bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          s_requestId,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts
        )
      )
    );

    (address invalidSigner, uint256 invalidSignerPk) = makeAddrAndKey("invalid signer");

    (bytes32 r, bytes32 s) = _sign(invalidSignerPk, digest);

    paymentRequests[0].r = r;
    paymentRequests[0].s = s;

    vm.expectRevert(abi.encodeWithSelector(PaymentTokenOnRamp.InvalidSignature.selector, invalidSigner, s_authority));

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);
  }

  function test_submitPaymentRequests_RevertWhen_TokenAmountsListIsEmpty() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](0);

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);
    paymentRequests[0].tokenAmounts = tokenAmounts;

    bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          s_requestId,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts
        )
      )
    );

    (bytes32 r, bytes32 s) = _sign(s_authorityPk, digest);

    paymentRequests[0].r = r;
    paymentRequests[0].s = s;

    vm.expectRevert(Errors.EmptyList.selector);

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);
  }

  function test_submitPaymentRequests_RevertWhen_AmountEqZero() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](1);
    tokenAmounts[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: 0});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);
    paymentRequests[0].tokenAmounts = tokenAmounts;

    bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          s_requestId,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts
        )
      )
    );

    (bytes32 r, bytes32 s) = _sign(s_authorityPk, digest);

    paymentRequests[0].r = r;
    paymentRequests[0].s = s;

    vm.expectRevert(Errors.InvalidZeroAmount.selector);

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);
  }

  function test_submitPaymentRequests_SingleRequestWithSingleToken() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](1);
    tokenAmounts[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: ETH_PAYMENT_AMOUNT});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);
    paymentRequests[0].tokenAmounts = tokenAmounts;

    bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          s_requestId,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts
        )
      )
    );

    (bytes32 r, bytes32 s) = _sign(s_authorityPk, digest);

    paymentRequests[0].r = r;
    paymentRequests[0].s = s;

    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.PaymentRequestSubmitted(s_requestId, tokenAmounts);

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);

    assertTrue(s_paymentTokenOnRamp.isPaymentRequestSubmitted(s_requestId));
  }

  function test_submitPaymentRequests_SingleRequestWithMultipleTokens() external {
    Common.AssetAmount[] memory tokenAmounts = new Common.AssetAmount[](2);
    tokenAmounts[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: ETH_PAYMENT_AMOUNT});
    tokenAmounts[1] = Common.AssetAmount({asset: address(s_mockUSDC), amount: USDC_PAYMENT_AMOUNT});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](1);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);
    paymentRequests[0].tokenAmounts = tokenAmounts;

    bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          s_requestId,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts
        )
      )
    );

    (bytes32 r, bytes32 s) = _sign(s_authorityPk, digest);

    paymentRequests[0].r = r;
    paymentRequests[0].s = s;

    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.PaymentRequestSubmitted(s_requestId, tokenAmounts);

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);

    assertTrue(s_paymentTokenOnRamp.isPaymentRequestSubmitted(s_requestId));
  }

  function test_submitPaymentRequests_MultipleRequests() external {
    Common.AssetAmount[] memory tokenAmounts1 = new Common.AssetAmount[](1);
    tokenAmounts1[0] = Common.AssetAmount({asset: address(s_mockWETH), amount: ETH_PAYMENT_AMOUNT});
    Common.AssetAmount[] memory tokenAmounts2 = new Common.AssetAmount[](1);
    tokenAmounts2[0] = Common.AssetAmount({asset: address(s_mockUSDC), amount: USDC_PAYMENT_AMOUNT});

    PaymentTokenOnRamp.PaymentRequest[] memory paymentRequests = new PaymentTokenOnRamp.PaymentRequest[](2);
    paymentRequests[0].requestId = s_requestId;
    paymentRequests[0].fundingAddress = i_owner;
    paymentRequests[0].deadline = uint40(block.timestamp);
    paymentRequests[0].tokenAmounts = tokenAmounts1;

    bytes32 requestId2 = keccak256("requestId2");
    paymentRequests[1].requestId = requestId2;
    paymentRequests[1].fundingAddress = i_owner;
    paymentRequests[1].deadline = uint40(block.timestamp);
    paymentRequests[1].tokenAmounts = tokenAmounts2;

    bytes32 request1Digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          s_requestId,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts1
        )
      )
    );

    (bytes32 r, bytes32 s) = _sign(s_authorityPk, request1Digest);

    paymentRequests[0].r = r;
    paymentRequests[0].s = s;

    bytes32 request2Digest = MessageHashUtils.toEthSignedMessageHash(
      keccak256(
        abi.encode(
          s_paymentTokenOnRamp.typeAndVersion(),
          CHAIN_SELECTOR,
          address(s_paymentTokenOnRamp),
          requestId2,
          uint40(block.timestamp),
          i_owner,
          tokenAmounts2
        )
      )
    );
    (r, s) = _sign(s_authorityPk, request2Digest);

    paymentRequests[1].r = r;
    paymentRequests[1].s = s;

    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.PaymentRequestSubmitted(s_requestId, tokenAmounts1);
    vm.expectEmit(address(s_paymentTokenOnRamp));
    emit PaymentTokenOnRamp.PaymentRequestSubmitted(requestId2, tokenAmounts2);

    s_paymentTokenOnRamp.submitPaymentRequests(paymentRequests);

    assertTrue(s_paymentTokenOnRamp.isPaymentRequestSubmitted(s_requestId));
    assertTrue(s_paymentTokenOnRamp.isPaymentRequestSubmitted(requestId2));
  }
}
