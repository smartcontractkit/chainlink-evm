pragma solidity 0.8.19;

import {VRF} from "../VRF.sol";
import {VRFTypes} from "../VRFTypes.sol";
import {BatchVRFCoordinatorV2Plus} from "../dev/BatchVRFCoordinatorV2Plus.sol";

import {VRFCoordinatorV2_5} from "../dev/VRFCoordinatorV2_5.sol";
import {VRFV2PlusClient} from "../dev/libraries/VRFV2PlusClient.sol";
import "./BaseTest.t.sol";
import {FixtureVRFCoordinatorV2_5} from "./FixtureVRFCoordinatorV2_5.t.sol";
import {console} from "forge-std/console.sol";

contract BatchVRFCoordinatorV2PlusTest is FixtureVRFCoordinatorV2_5 {
  BatchVRFCoordinatorV2Plus private s_batchCoordinator;

  event RandomWordsFulfilled(
    uint256 indexed requestId,
    uint256 outputSeed,
    uint256 indexed subId,
    uint96 payment,
    bool nativePayment,
    bool success,
    bool onlyPremium
  );

  function setUp() public override {
    FixtureVRFCoordinatorV2_5.setUp();

    s_batchCoordinator = new BatchVRFCoordinatorV2Plus(address(s_coordinator));
  }

  function test_fulfillRandomWords() public {
    _setUpConfig();
    _setUpProvingKey();
    _setUpSubscription();

    uint32 requestBlock = 10;
    vm.roll(requestBlock);

    vm.startPrank(SUBSCRIPTION_OWNER);
    vm.deal(SUBSCRIPTION_OWNER, 10 ether);
    s_coordinator.fundSubscriptionWithNative{value: 10 ether}(s_subId);

    // Request random words.
    s_consumer.requestRandomWords(CALLBACK_GAS_LIMIT, MIN_CONFIRMATIONS, NUM_WORDS, VRF_KEY_HASH, true);
    vm.stopPrank();

    // Move on to the next block.
    // Store the previous block's blockhash.
    vm.roll(requestBlock + 1);
    s_bhs.store(requestBlock);
    assertEq(hex"1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac", s_bhs.getBlockhash(requestBlock));

    VRFTypes.Proof[] memory proofs = new VRFTypes.Proof[](2);
    VRFTypes.RequestCommitmentV2Plus[] memory rcs = new VRFTypes.RequestCommitmentV2Plus[](2);

    // Proof generated via the generate-proof-v2-plus script command.
    // 1st step: Uncomment the print command below and run the test to print the output.
    // _printGenerateProofV2PlusCommand(address(s_consumer1), 1, requestBlock, false);
    // 2nd step: export the following environment variables to run the generate-proof-v2-plus script.
    // export ETH_URL=https://ethereum-sepolia-rpc.publicnode.com # or any other RPC provider you prefer
    // export ETH_CHAIN_ID=11155111 # or switch to any other chain
    // export ACCOUNT_KEY=<your test EOA private key>
    // 3rd step: copy the output from the 1st step and update the command below, then run the command
    // and copy the command output in the proof section below
    /*
       Run from this folder: chainlink/core/scripts/vrfv2plus/testnet
       go run . generate-proof-v2-plus \
         -key-hash 0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528 \
         -pre-seed 4430852740828987645228960511496023658059009607317025880962658187812299131155 \
         -block-hash 0x1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac \
         -block-num 10 \
         -sender 0xdc90e8ce61c1af8a638b95264037c8e67ee5765c \
         -native-payment true

    */
    proofs[0] = VRFTypes.Proof({
      pk: [
        72_488_970_228_380_509_287_422_715_226_575_535_698_893_157_273_063_074_627_791_787_432_852_706_183_111,
        62_070_622_898_698_443_831_883_535_403_436_258_712_770_888_294_397_026_493_185_421_712_108_624_767_191
      ],
      gamma: [
        26_762_213_923_453_052_192_184_693_334_574_145_607_290_366_984_305_044_804_336_172_347_176_490_943_606,
        70_503_534_560_525_619_072_578_237_689_732_581_746_976_650_376_431_765_635_714_023_643_649_039_207_077
      ],
      c: 10_992_233_996_918_874_905_152_274_435_276_937_088_064_589_467_016_709_044_984_819_613_170_049_539_489,
      s: 79_662_863_379_962_724_455_809_192_044_326_025_082_567_113_176_696_761_949_197_261_107_120_333_769_102,
      seed: 4_430_852_740_828_987_645_228_960_511_496_023_658_059_009_607_317_025_880_962_658_187_812_299_131_155,
      uWitness: 0x421A52Fb797d76Fb610aA1a0c020346fC1Ee2DeB,
      cGammaWitness: [
        50_748_523_246_052_507_241_857_300_891_945_475_679_319_243_536_065_937_584_940_024_494_820_365_165_901,
        85_746_856_994_474_260_612_851_047_426_766_648_416_105_284_284_185_975_301_552_792_881_940_939_754_570
      ],
      sHashWitness: [
        78_637_275_871_978_664_522_379_716_948_105_702_461_748_200_460_627_087_255_706_483_027_519_919_611_423,
        82_219_236_913_923_465_822_780_520_561_305_604_064_850_823_877_720_616_893_986_252_854_976_640_396_959
      ],
      zInv: 60_547_558_497_534_848_069_125_896_511_700_272_238_016_171_243_048_151_035_528_198_622_956_754_542_730
    });
    rcs[0] = VRFTypes.RequestCommitmentV2Plus({
      blockNum: requestBlock,
      subId: s_subId,
      callbackGasLimit: CALLBACK_GAS_LIMIT,
      numWords: 1,
      sender: address(s_consumer),
      extraArgs: VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: true}))
    });

    VRFCoordinatorV2_5.Output memory output =
      s_coordinator.getRandomnessFromProofExternal(abi.decode(abi.encode(proofs[0]), (VRF.Proof)), rcs[0]);

    requestBlock = 20;
    vm.roll(requestBlock);

    vm.startPrank(SUBSCRIPTION_OWNER);
    s_linkToken.setBalance(address(SUBSCRIPTION_OWNER), 10 ether);
    s_linkToken.transferAndCall(address(s_coordinator), 10 ether, abi.encode(s_subId));

    // Request random words.
    s_consumer1.requestRandomWords(CALLBACK_GAS_LIMIT, MIN_CONFIRMATIONS, NUM_WORDS, VRF_KEY_HASH, false);
    vm.stopPrank();

    // Move on to the next block.
    // Store the previous block's blockhash.
    vm.roll(requestBlock + 1);
    s_bhs.store(requestBlock);
    assertEq(hex"731dc163f73d31d8c68f9917ce4ff967753939f70432973c04fd2c2a48148607", s_bhs.getBlockhash(requestBlock));

    // Proof generated via the generate-proof-v2-plus script command.
    // 1st step: Uncomment the print command below and run the test to print the output.
    // _printGenerateProofV2PlusCommand(address(s_consumer1), 1, requestBlock, false);
    // 2nd step: export the following environment variables to run the generate-proof-v2-plus script.
    // export ETH_URL=https://ethereum-sepolia-rpc.publicnode.com # or any other RPC provider you prefer
    // export ETH_CHAIN_ID=11155111 # or switch to any other chain
    // export ACCOUNT_KEY=<your test EOA private key>
    // 3rd step: copy the output from the 1st step and update the command below, then run the command
    // and copy the command output in the proof section below
    /*
       Run from this folder: chainlink/core/scripts/vrfv2plus/testnet
       go run . generate-proof-v2-plus \
         -key-hash 0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528 \
         -pre-seed 14541556911652758131165474365357244907354309169650401973525070879190071151266 \
         -block-hash 0x731dc163f73d31d8c68f9917ce4ff967753939f70432973c04fd2c2a48148607 \
         -block-num 20 \
         -sender 0x2f1c0761d6e4b1e5f01968d6c746f695e5f3e25d \
         -native-payment false
    */
    proofs[1] = VRFTypes.Proof({
      pk: [
        72_488_970_228_380_509_287_422_715_226_575_535_698_893_157_273_063_074_627_791_787_432_852_706_183_111,
        62_070_622_898_698_443_831_883_535_403_436_258_712_770_888_294_397_026_493_185_421_712_108_624_767_191
      ],
      gamma: [
        97_658_842_840_420_719_674_383_370_910_135_023_062_422_561_858_595_941_631_054_490_821_636_116_883_585,
        44_255_438_468_488_339_528_368_406_358_785_988_551_798_314_198_954_634_050_943_346_751_039_644_360_856
      ],
      c: 5_233_652_943_248_967_403_606_766_735_502_925_802_264_855_214_922_758_107_203_237_169_366_748_118_852,
      s: 87_931_642_435_666_855_739_510_477_620_068_257_005_869_145_374_865_238_974_094_299_759_068_218_698_655,
      seed: 14_541_556_911_652_758_131_165_474_365_357_244_907_354_309_169_650_401_973_525_070_879_190_071_151_266,
      uWitness: 0x0A87a9CB71983cE0F2C4bA41D0c1A6Fb1785c46A,
      cGammaWitness: [
        54_062_743_217_909_816_783_918_413_821_204_010_151_082_432_359_411_822_104_552_882_037_459_289_383_418,
        67_491_004_534_731_980_264_926_765_871_774_299_056_809_003_077_448_271_411_776_926_359_153_820_235_981
      ],
      sHashWitness: [
        7_745_933_951_617_569_731_026_754_652_291_310_837_540_252_155_195_826_133_994_719_499_558_406_927_394,
        58_405_861_596_456_412_358_325_504_621_101_233_475_720_292_237_067_230_796_670_629_212_111_423_924_259
      ],
      zInv: 44_253_513_765_558_903_217_330_502_897_662_324_213_800_000_485_156_126_961_643_960_636_269_885_275_795
    });
    rcs[1] = VRFTypes.RequestCommitmentV2Plus({
      blockNum: requestBlock,
      subId: s_subId,
      callbackGasLimit: CALLBACK_GAS_LIMIT,
      numWords: 1,
      sender: address(s_consumer1),
      extraArgs: VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: false}))
    });

    VRFCoordinatorV2_5.Output memory output1 =
      s_coordinator.getRandomnessFromProofExternal(abi.decode(abi.encode(proofs[1]), (VRF.Proof)), rcs[1]);

    // The payments are NOT pre-calculated and simply copied from the actual event.
    // We can assert and ignore the payment field but the code will be considerably longer.
    vm.expectEmit(true, true, false, true, address(s_coordinator));
    emit RandomWordsFulfilled(output.requestId, output.randomness, s_subId, 500_000_000_000_143_261, true, true, false);
    vm.expectEmit(true, true, false, true, address(s_coordinator));
    emit RandomWordsFulfilled(
      output1.requestId, output1.randomness, s_subId, 800_000_000_000_312_358, false, true, false
    );

    // Fulfill the requests.
    s_batchCoordinator.fulfillRandomWords(proofs, rcs);
  }
}
