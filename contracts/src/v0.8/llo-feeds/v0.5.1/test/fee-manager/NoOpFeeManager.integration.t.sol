// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import {Common} from "../../../libraries/Common.sol";
import {NoOpFeeManager} from "../../NoOpFeeManager.sol";
import {BaseTestWithConfiguredVerifierAndFeeManager} from "../verifier/BaseVerifierTest.t.sol";

/// @notice Extended interface matching what real clients call on FeeManager/NoOpFeeManager
interface IFeeManagerWithGetters {
  function getFeeAndReward(
    address subscriber,
    bytes memory report,
    address quoteAddress
  ) external returns (Common.Asset memory, Common.Asset memory, uint256);
  function i_linkAddress() external view returns (address);
  function i_nativeAddress() external view returns (address);
  function i_rewardManager() external view returns (address);
}

/**
 * @title NoOpFeeManagerIntegrationTest
 * @notice Tests that NoOpFeeManager is compatible with the Data Streams verification workflow when billing is enabled.
 * @dev Uses a single _runVerificationWorkflow() function to ensure identical code paths are tested
 *      with both FeeManager and NoOpFeeManager.
 */
contract NoOpFeeManagerIntegrationTestV051 is BaseTestWithConfiguredVerifierAndFeeManager {
  NoOpFeeManager internal noOpFeeManager;

  uint64 internal constant PERCENTAGE_SCALAR = 1e18;
  uint256 internal constant DEFAULT_LINK_MINT = 100 ether;

  /// @notice Result of running the verification workflow
  struct VerificationResult {
    bytes verifiedReport;
    uint256 feeAmount;
    uint256 rewardAmount;
    uint256 discount;
    uint256 linkBalanceBefore;
    uint256 linkBalanceAfter;
    int192 decodedPrice;
    bytes32 decodedFeedId;
  }

  function setUp() public override {
    BaseTestWithConfiguredVerifierAndFeeManager.setUp();

    // Deploy NoOpFeeManager with same addresses as FeeManager
    noOpFeeManager = new NoOpFeeManager(address(link), address(native), address(rewardManager));

    // Mint tokens for users
    link.mint(USER, DEFAULT_LINK_MINT);
    vm.deal(USER, DEFAULT_LINK_MINT);
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // Main Test: Verification Workflow Compatibility
  // ═══════════════════════════════════════════════════════════════════════════

  function test_verificationWorkflowCompatibility() public {
    // Generate report once - used for both workflows
    V3Report memory report = _generateV3Report();
    Signer[] memory signers = _getSigners(FAULT_TOLERANCE + 1);
    bytes32[3] memory reportContext = _generateReportContext(v3ConfigDigest);
    bytes memory signedReport = _generateV3EncodedBlob(report, reportContext, signers);
    bytes memory reportData = _encodeReport(report);

    // ─── Run 1: FeeManager (billing enabled) ─────────────────────────────────
    VerificationResult memory result1 = _runVerificationWorkflow(signedReport, reportData, USER);

    // Validate FeeManager was invoked: fee charged, LINK transferred
    assertEq(result1.feeAmount, DEFAULT_REPORT_LINK_FEE, "Fee should match report linkFee");
    assertEq(result1.rewardAmount, DEFAULT_REPORT_LINK_FEE, "Reward should match report linkFee");
    assertEq(result1.discount, 0, "No discount configured");
    assertLt(result1.linkBalanceAfter, result1.linkBalanceBefore, "LINK should be transferred");
    assertEq(
      result1.linkBalanceBefore - result1.linkBalanceAfter,
      DEFAULT_REPORT_LINK_FEE,
      "Transferred amount should match fee"
    );

    // Validate report decoded correctly
    assertEq(result1.decodedPrice, MEDIAN, "Decoded price should match");
    assertEq(result1.decodedFeedId, FEED_ID_V3, "Decoded feedId should match");

    // ─── Swap to NoOpFeeManager ──────────────────────────────────────────────
    changePrank(ADMIN);
    s_verifierProxy.setFeeManager(noOpFeeManager);

    // ─── Run 2: NoOpFeeManager (billing deactivated) ─────────────────────────
    VerificationResult memory result2 = _runVerificationWorkflow(signedReport, reportData, USER);

    // Validate NoOpFeeManager: no fees, no transfer
    assertEq(result2.feeAmount, 0, "NoOp fee should be 0");
    assertEq(result2.rewardAmount, 0, "NoOp reward should be 0");
    assertEq(result2.discount, PERCENTAGE_SCALAR, "NoOp discount should be 100%");
    assertEq(result2.linkBalanceAfter, result2.linkBalanceBefore, "No LINK should be transferred");

    // Validate same report decoded - proves verification still works
    assertEq(result2.decodedPrice, result1.decodedPrice, "Same price decoded after swap");
    assertEq(result2.decodedFeedId, result1.decodedFeedId, "Same feedId decoded after swap");
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // Verification Workflow
  // ═══════════════════════════════════════════════════════════════════════════

  /**
   * @notice Runs the Data Streams verification workflow when billing is enabled
   * @dev This function mirrors the integration pattern from the docs:
   *      1. Query s_feeManager() from proxy to get active fee manager
   *      2. Call i_linkAddress() to get fee token
   *      3. Call i_rewardManager() to get active reward manager
   *      4. Call getFeeAndReward() to calculate fees
   *      5. Approve RewardManager to collect the fee amount
   *      6. Call verify() on VerifierProxy
   * @param signedReport The signed report blob to verify
   * @param reportData The encoded report data for fee calculation
   * @param sender The address to run the workflow as
   * @return result The verification result containing all workflow outputs
   */
  function _runVerificationWorkflow(
    bytes memory signedReport,
    bytes memory reportData,
    address sender
  ) internal returns (VerificationResult memory result) {
    changePrank(sender);

    // 1. Get active fee manager from proxy (exactly like real integration)
    IFeeManagerWithGetters activeFeeManager = IFeeManagerWithGetters(address(s_verifierProxy.s_feeManager()));

    // 2. Query fee token address
    address feeToken = activeFeeManager.i_linkAddress();

    // 3. Query reward manager for approval
    address rewardMgr = address(activeFeeManager.i_rewardManager());

    // 4. Calculate fee
    (Common.Asset memory fee, Common.Asset memory reward, uint256 discount) =
      activeFeeManager.getFeeAndReward(sender, reportData, feeToken);

    result.feeAmount = fee.amount;
    result.rewardAmount = reward.amount;
    result.discount = discount;

    // 5. Record balance before and approve
    result.linkBalanceBefore = link.balanceOf(sender);
    if (fee.amount > 0) {
      link.approve(rewardMgr, fee.amount);
    }

    // 6. Verify through proxy
    result.verifiedReport = s_verifierProxy.verify(signedReport, abi.encode(feeToken));

    // Record balance after
    result.linkBalanceAfter = link.balanceOf(sender);

    // Decode verified report
    V3Report memory decoded = abi.decode(result.verifiedReport, (V3Report));
    result.decodedPrice = decoded.benchmarkPrice;
    result.decodedFeedId = decoded.feedId;

    return result;
  }
}
