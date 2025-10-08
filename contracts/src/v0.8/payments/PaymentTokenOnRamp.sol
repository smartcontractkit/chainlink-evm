// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {ITypeAndVersion} from "../shared/interfaces/ITypeAndVersion.sol";
import {IFeeWithdrawer} from "./interfaces/IFeeWithdrawer.sol";

import {EmergencyWithdrawer} from "./EmergencyWithdrawer.sol";
import {PausableWithAccessControl} from "./PausableWithAccessControl.sol";
import {Common} from "./libraries/Common.sol";
import {Errors} from "./libraries/Errors.sol";
import {Roles} from "./libraries/Roles.sol";

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts@5.0.2/utils/cryptography/MessageHashUtils.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @notice The PaymentTokenOnRamp contract handles payment requests from an off-chain system by verifying signatures
/// and processing token transfers.
contract PaymentTokenOnRamp is IERC165, ITypeAndVersion, IFeeWithdrawer, EmergencyWithdrawer {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  /// @notice This event is emitted when the chain selector is set.
  /// @param chainSelector The chain selector of the chain.
  event ChainSelectorSet(uint256 chainSelector);
  /// @notice This event is emitted when a new fee aggregator is set.
  /// @param feeAggregator The address of the fee aggregator.
  event FeeAggregatorSet(address feeAggregator);
  /// @notice This event is emitted when a fee token is withdrawn to the fee aggregator.
  /// @param feeAggregator The fee aggregator address.
  /// @param feeToken The fee token address.
  /// @param amount The amount of fee token withdrawn.
  event FeeTokenWithdrawn(address indexed feeAggregator, address indexed feeToken, uint256 amount);
  /// @notice This event is emitted when a payment is requested.
  /// @param requestId The ID of the request (obtained off-chain).
  /// @param tokenAmounts The tokens and amounts to fund.
  event PaymentRequestSubmitted(bytes32 indexed requestId, Common.AssetAmount[] tokenAmounts);

  /// @notice This error is thrown when trying to set the chain selector to zero.
  error InvalidChainSelector();
  /// @notice This error is thrown when an expired payment request is submitted.
  error ExpiredRequest(bytes32 requestId, uint256 deadline);
  /// @notice This error is thrown when a request is resubmitted after having already been submitted.
  error RequestAlreadySubmitted(bytes32 requestId);
  /// @notice This error is thrown when the signature attached to the payment request is invalid.
  error InvalidSignature(address signer);

  /// @notice Parameters to instantiate the contract in the constructor.
  struct ConstructorParams {
    address admin; // ─────────────────╮ The initial contract admin.
    uint48 adminRoleTransferDelay; // ─╯ The minimum seconds before the admin address can be transferred.
    address feeAggregator; // ─────────╮ The fee aggregator.
    uint96 chainSelector; // ──────────╯ The chain selector of the chain.
    address[] paymentRequestSigners; // The payment requests' signers.
  }

  /// @notice This struct contains the parameters to submit a payment.
  struct PaymentRequest {
    bytes32 requestId; // The request ID, obtained off-chain.
    address fundingAddress; // ─╮ The address funding the request.
    uint40 deadline; // ────────╯ The expiration timestamp of the request (inclusive).
    bytes32 r; // First 32 bytes of the ECDSA signature.
    bytes32 s; // Second 32 bytes of the ECDSA signature.
    Common.AssetAmount[] tokenAmounts; // The tokens and amounts of to transfer.
  }

  /// @inheritdoc ITypeAndVersion
  string public constant override typeAndVersion = "PaymentTokenOnRamp 1.0.0";

  /// @dev The off-chain system only generates sigs with v=27; making this constant allows us to save gas by not
  /// transmitting v.
  /// @dev Any valid ECDSA sig (r, s, v) can be "flipped" into (r, s*, v*) without knowing the private key (where v=27
  /// or 28 for secp256k1)
  uint8 public constant ECDSA_RECOVERY_V = 27;

  /// @notice The chain selector of the chain
  uint256 public immutable i_chainSelector;

  /// @notice The fee tokens recipient.
  address private s_feeAggregator;

  /// @notice Mapping of submitted payment request ids.
  EnumerableSet.Bytes32Set private s_submittedPaymentRequests;

  constructor(
    ConstructorParams memory params
  ) EmergencyWithdrawer(params.adminRoleTransferDelay, params.admin) {
    if (params.chainSelector == 0) {
      revert InvalidChainSelector();
    }

    i_chainSelector = params.chainSelector;

    emit ChainSelectorSet(params.chainSelector);

    for (uint256 i; i < params.paymentRequestSigners.length; ++i) {
      if (params.paymentRequestSigners[i] == address(0)) {
        revert Errors.InvalidZeroAddress();
      }
      _grantRole(Roles.PAYMENT_VALIDATOR_ROLE, params.paymentRequestSigners[i]);
    }

    _setFeeAggregator(params.feeAggregator);
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) public view override(PausableWithAccessControl, IERC165) returns (bool) {
    return (interfaceId == type(IFeeWithdrawer).interfaceId || PausableWithAccessControl.supportsInterface(interfaceId));
  }

  // ================================================================================================
  // │                                       Payment Requests                                       │
  // ================================================================================================

  /// @notice Submits signed payment requests.
  /// @dev precondition - The payment request list must be greater than zero.
  /// @dev precondition - The payment request funding address must not be the zero address.
  /// @dev precondition - The payment request must not have expired.
  /// @dev precondition - The payment request must not have already been submitted.
  /// @dev precondition - The list of tokens to transfer must not be empty.
  /// @dev precondition - The payment request must have been signed by an address with the PAYMENT_VALIDATOR_ROLE.
  /// @dev precondition - The amount of tokens to transfer must be greater than zero.
  /// @param paymentRequests The payment requests to submit.
  function submitPaymentRequests(
    PaymentRequest[] calldata paymentRequests
  ) external whenNotPaused {
    if (paymentRequests.length == 0) {
      revert Errors.EmptyList();
    }

    for (uint256 i; i < paymentRequests.length; ++i) {
      PaymentRequest memory paymentRequest = paymentRequests[i];

      if (paymentRequest.fundingAddress == address(0)) {
        revert Errors.InvalidZeroAddress();
      }
      if (paymentRequest.deadline < block.timestamp) {
        revert ExpiredRequest(paymentRequest.requestId, paymentRequest.deadline);
      }
      if (!s_submittedPaymentRequests.add(paymentRequest.requestId)) {
        revert RequestAlreadySubmitted(paymentRequest.requestId);
      }

      uint256 tokenAmountsLen = paymentRequest.tokenAmounts.length;

      if (tokenAmountsLen == 0) revert Errors.EmptyList();

      bytes32 digest = MessageHashUtils.toEthSignedMessageHash(
        keccak256(
          abi.encode(
            typeAndVersion,
            i_chainSelector,
            address(this),
            paymentRequest.requestId,
            paymentRequest.deadline,
            paymentRequest.fundingAddress,
            paymentRequest.tokenAmounts
          )
        )
      );

      address signer = ecrecover(digest, ECDSA_RECOVERY_V, paymentRequest.r, paymentRequest.s);

      if (!hasRole(Roles.PAYMENT_VALIDATOR_ROLE, signer)) {
        revert InvalidSignature(signer);
      }

      for (uint256 j; j < tokenAmountsLen; ++j) {
        Common.AssetAmount memory tokenAmount = paymentRequest.tokenAmounts[j];

        if (tokenAmount.amount == 0) {
          revert Errors.InvalidZeroAmount();
        }

        IERC20(tokenAmount.asset).safeTransferFrom(paymentRequest.fundingAddress, address(this), tokenAmount.amount);
      }

      emit PaymentRequestSubmitted(paymentRequest.requestId, paymentRequest.tokenAmounts);
    }
  }

  // ================================================================================================
  // │                                         Withdrawals                                          │
  // ================================================================================================

  /// @inheritdoc IFeeWithdrawer
  /// @dev This function can be permissionless as it only transfers tokens to the fee aggregator which is a trusted
  /// address.
  /// @dev precondition - The contract must not be paused
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external whenNotPaused {
    if (feeTokens.length == 0) {
      revert Errors.EmptyList();
    }

    address feeAggregator = s_feeAggregator;

    for (uint256 i = 0; i < feeTokens.length; ++i) {
      IERC20 feeToken = IERC20(feeTokens[i]);
      uint256 feeTokenBalance = feeToken.balanceOf(address(this));

      if (feeTokenBalance > 0) {
        feeToken.safeTransfer(feeAggregator, feeTokenBalance);

        emit FeeTokenWithdrawn(feeAggregator, address(feeToken), feeTokenBalance);
      }
    }
  }

  // ================================================================================================
  // ||                                           Config                                           ||
  // ================================================================================================

  /// @notice Sets the fee aggregator.
  /// @dev precondition The caller must have the DEFAULT_ADMIN_ROLE.
  /// @dev precondition The new fee aggregator address must not be the zero address.
  /// @dev precondition The new fee aggregator address must be different from the already configured fee aggregator.
  function setFeeAggregator(
    address feeAggregator
  ) external onlyRole(DEFAULT_ADMIN_ROLE) {
    _setFeeAggregator(feeAggregator);
  }

  /// @notice Sets the fee aggregator.
  /// @param feeAggregator The new fee aggregator.
  function _setFeeAggregator(
    address feeAggregator
  ) internal {
    if (feeAggregator == address(0)) {
      revert Errors.InvalidZeroAddress();
    }
    if (s_feeAggregator == feeAggregator) {
      revert Errors.ValueNotUpdated();
    }

    s_feeAggregator = feeAggregator;

    emit FeeAggregatorSet(feeAggregator);
  }

  // ================================================================================================
  // │                                           Getters                                            │
  // ================================================================================================

  /// @notice Getter function to retrieve the configured fee aggregator.
  /// @return feeAggregator The configured fee aggregator.
  function getFeeAggregator() external view returns (address feeAggregator) {
    return s_feeAggregator;
  }

  /// @notice Checks whether a payment request has already been submitted or not.
  /// @param requestId The payment request ID.
  /// @return isSubmitted True if the payment request has already been submitted, false if not.
  function isPaymentRequestSubmitted(
    bytes32 requestId
  ) external view returns (bool isSubmitted) {
    return s_submittedPaymentRequests.contains(requestId);
  }

  /// @notice Getter function to retrieve a paginated list of submitted payment requests.
  /// @param start Zero-based index of the first owner to include in the result.
  /// @param limit Maximum number of owners to return (clamped by a sensible internal cap).
  /// @return submittedRequests The list of submitted payment requests.
  function getSubmittedPaymentRequests(
    uint256 start,
    uint256 limit
  ) external view returns (bytes32[] memory submittedRequests) {
    uint256 total = s_submittedPaymentRequests.length();

    if (start >= total) return new bytes32[](0);

    uint256 remaining = total - start;
    uint256 count = start + limit > total ? remaining : limit;

    submittedRequests = new bytes32[](count);

    for (uint256 i; i < count; ++i) {
      submittedRequests[i] = s_submittedPaymentRequests.at(start + i);
    }

    return submittedRequests;
  }
}
