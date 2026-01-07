// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import {ERC20Mock} from "../../../../shared/mocks/ERC20Mock.sol";
import {WERC20Mock} from "../../../../shared/mocks/WERC20Mock.sol";

import {Common} from "../../../libraries/Common.sol";
import {FeeManager} from "../../FeeManager.sol";
import {NoOpFeeManager} from "../../NoOpFeeManager.sol";
import {RewardManager} from "../../RewardManager.sol";
import {IVerifierFeeManager} from "../../interfaces/IVerifierFeeManager.sol";
import {FeeManagerProxy} from "../mocks/FeeManagerProxy.sol";
import {Test} from "forge-std/Test.sol";

contract NoOpFeeManagerTest is Test {
  NoOpFeeManager internal noOpFeeManager;
  FeeManager internal feeManager;
  RewardManager internal rewardManager;
  FeeManagerProxy internal feeManagerProxy;

  ERC20Mock internal link;
  WERC20Mock internal native;

  uint64 internal constant PERCENTAGE_SCALAR = 1e18;
  address internal constant ADMIN = address(uint160(uint256(keccak256("ADMIN"))));
  address internal constant SUBSCRIBER = address(uint160(uint256(keccak256("SUBSCRIBER"))));

  function setUp() public {
    vm.startPrank(ADMIN);

    link = new ERC20Mock(18);
    native = new WERC20Mock();
    noOpFeeManager = new NoOpFeeManager();

    // Deploy real FeeManager for comparison tests
    feeManagerProxy = new FeeManagerProxy();
    rewardManager = new RewardManager(address(link));
    feeManager = new FeeManager(address(link), address(native), address(feeManagerProxy), address(rewardManager));
    feeManagerProxy.setFeeManager(feeManager);
    rewardManager.setFeeManager(address(feeManager));

    vm.stopPrank();
  }

  function test_typeAndVersion() public view {
    assertEq(noOpFeeManager.typeAndVersion(), "NoOpFeeManager 0.5.0");
  }

  // VerifierProxy checks these interface support before accepting a fee manager
  function test_supportsInterface() public view {
    assertTrue(noOpFeeManager.supportsInterface(IVerifierFeeManager.processFee.selector));
    assertTrue(noOpFeeManager.supportsInterface(IVerifierFeeManager.processFeeBulk.selector));
    assertFalse(noOpFeeManager.supportsInterface(bytes4(0xdeadbeef)));
  }

  // Some code queries these to check discount status - must always return 100%
  function test_discountGettersReturn100Percent() public view {
    bytes32 feedId = keccak256("ETH-USD");

    assertEq(noOpFeeManager.s_globalDiscounts(SUBSCRIBER, address(link)), PERCENTAGE_SCALAR);
    assertEq(noOpFeeManager.s_subscriberDiscounts(SUBSCRIBER, feedId, address(link)), PERCENTAGE_SCALAR);
  }

  // Ensures our view functions match FeeManager's public mapping getter signatures
  // so code can swap between implementations without breaking changes
  function test_discountGettersMatchFeeManagerSignature() public view {
    bytes32 feedId = keccak256("ETH-USD");

    // FeeManager public mappings return 0 by default, NoOpFeeManager returns 100%
    assertEq(feeManager.s_globalDiscounts(SUBSCRIBER, address(link)), 0);
    assertEq(noOpFeeManager.s_globalDiscounts(SUBSCRIBER, address(link)), PERCENTAGE_SCALAR);

    assertEq(feeManager.s_subscriberDiscounts(SUBSCRIBER, feedId, address(link)), 0);
    assertEq(noOpFeeManager.s_subscriberDiscounts(SUBSCRIBER, feedId, address(link)), PERCENTAGE_SCALAR);
  }

  // Zero fees, zero rewards, 100% discount indicates no fees are charged
  function test_getFeeAndReward() public view {
    bytes memory report = abi.encode(bytes32(0), uint32(0), int192(0));
    (Common.Asset memory fee, Common.Asset memory reward, uint256 discount) =
      noOpFeeManager.getFeeAndReward(SUBSCRIBER, report, address(link));

    assertEq(fee.amount, 0);
    assertEq(reward.amount, 0);
    assertEq(discount, PERCENTAGE_SCALAR);
  }

  function test_linkAvailableForPayment() public view {
    assertEq(noOpFeeManager.linkAvailableForPayment(), 0);
  }

  // Admin functions are no-ops but must not revert to maintain interface compatibility
  function test_adminFunctionsDoNotRevert() public {
    bytes32 feedId = keccak256("ETH-USD");
    Common.AddressAndWeight[] memory recipients = new Common.AddressAndWeight[](0);

    noOpFeeManager.setNativeSurcharge(uint64(PERCENTAGE_SCALAR));
    noOpFeeManager.updateSubscriberDiscount(SUBSCRIBER, feedId, address(link), uint64(PERCENTAGE_SCALAR));
    noOpFeeManager.updateSubscriberGlobalDiscount(SUBSCRIBER, address(link), uint64(PERCENTAGE_SCALAR));
    noOpFeeManager.setFeeRecipients(feedId, recipients);
    noOpFeeManager.withdraw(address(link), ADMIN, 1000);
    noOpFeeManager.payLinkDeficit(feedId);
  }

  // ETH sent during verification must be refunded since no fees are collected
  function test_processFeeRefundsETH() public {
    uint256 ethAmount = 1 ether;
    vm.deal(address(this), ethAmount);

    uint256 balanceBefore = SUBSCRIBER.balance;
    noOpFeeManager.processFee{value: ethAmount}("", "", SUBSCRIBER);

    assertEq(SUBSCRIBER.balance - balanceBefore, ethAmount);
  }

  function test_processFeeBulkRefundsETH() public {
    uint256 ethAmount = 1 ether;
    vm.deal(address(this), ethAmount);
    bytes[] memory payloads = new bytes[](2);

    uint256 balanceBefore = SUBSCRIBER.balance;
    noOpFeeManager.processFeeBulk{value: ethAmount}(payloads, "", SUBSCRIBER);

    assertEq(SUBSCRIBER.balance - balanceBefore, ethAmount);
  }
}
