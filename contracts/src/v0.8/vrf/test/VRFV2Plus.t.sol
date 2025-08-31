pragma solidity 0.8.19;

import {MockLinkToken} from "../../functions/tests/v1_X/testhelpers/MockLinkToken.sol";
import {MockV3Aggregator} from "../../shared/mocks/MockV3Aggregator.sol";
import {VRF} from "../VRF.sol";

import {VRFTypes} from "../VRFTypes.sol";
import {BlockhashStore} from "../dev/BlockhashStore.sol";
import {SubscriptionAPI} from "../dev/SubscriptionAPI.sol";
import {VRFCoordinatorV2_5} from "../dev/VRFCoordinatorV2_5.sol";

import {VRFV2PlusClient} from "../dev/libraries/VRFV2PlusClient.sol";
import {ExposedVRFCoordinatorV2_5} from "../dev/testhelpers/ExposedVRFCoordinatorV2_5.sol";
import {VRFV2PlusConsumerExample} from "../dev/testhelpers/VRFV2PlusConsumerExample.sol";

import {VRFV2PlusLoadTestWithMetrics} from "../dev/testhelpers/VRFV2PlusLoadTestWithMetrics.sol";
import "./BaseTest.t.sol";
import "@openzeppelin/contracts@4.9.6/utils/math/Math.sol";
import {VmSafe} from "forge-std/Vm.sol";
import {console} from "forge-std/console.sol"; // for Math.ceilDiv

/*
 * USAGE INSTRUCTIONS:
 * To add new tests/proofs, uncomment the "console.sol" import from foundry, and gather key fields
 * from your VRF request.
 * Then, pass your request info into the generate-proof-v2-plus script command
 * located in /core/scripts/vrfv2/testnet/proofs.go to generate a proof that can be tested on-chain.
 **/

contract VRFV2Plus is BaseTest {
  address internal constant LINK_WHALE = 0xD883a6A1C22fC4AbFE938a5aDF9B2Cc31b1BF18B;
  uint64 internal constant GAS_LANE_MAX_GAS = 5000 gwei;
  uint16 internal constant MIN_CONFIRMATIONS = 0;
  uint32 internal constant CALLBACK_GAS_LIMIT = 1_000_000;
  uint32 internal constant NUM_WORDS = 1;

  // Bytecode for a VRFV2PlusConsumerExample contract.
  // to calculate: console.logBytes(type(VRFV2PlusConsumerExample).creationCode);
  bytes constant initializeCode =
    hex"60806040523480156200001157600080fd5b5060405162001377380380620013778339810160408190526200003491620001cc565b8133806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000103565b5050600280546001600160a01b03199081166001600160a01b0394851617909155600580548216958416959095179094555060038054909316911617905562000204565b6001600160a01b0381163314156200015e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620001c757600080fd5b919050565b60008060408385031215620001e057600080fd5b620001eb83620001af565b9150620001fb60208401620001af565b90509250929050565b61116380620002146000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638098004311610097578063cf62c8ab11610066578063cf62c8ab14610242578063de367c8e14610255578063eff2701714610268578063f2fde38b1461027b57600080fd5b806380980043146101ab5780638da5cb5b146101be5780638ea98117146101cf578063a168fa89146101e257600080fd5b80635d7d53e3116100d35780635d7d53e314610166578063706da1ca1461016f5780637725135b1461017857806379ba5097146101a357600080fd5b80631fe543e31461010557806329e5d8311461011a5780632fa4e4421461014057806336bfffed14610153575b600080fd5b610118610113366004610e4e565b61028e565b005b61012d610128366004610ef2565b6102fa565b6040519081526020015b60405180910390f35b61011861014e366004610f7f565b610410565b610118610161366004610d5b565b6104bc565b61012d60045481565b61012d60065481565b60035461018b906001600160a01b031681565b6040516001600160a01b039091168152602001610137565b6101186105c0565b6101186101b9366004610e1c565b600655565b6000546001600160a01b031661018b565b6101186101dd366004610d39565b61067e565b61021d6101f0366004610e1c565b6007602052600090815260409020805460019091015460ff82169161010090046001600160a01b03169083565b6040805193151584526001600160a01b03909216602084015290820152606001610137565b610118610250366004610f7f565b61073d565b60055461018b906001600160a01b031681565b610118610276366004610f14565b610880565b610118610289366004610d39565b610a51565b6002546001600160a01b031633146102ec576002546040517f1cf993f40000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b0390911660248201526044015b60405180910390fd5b6102f68282610a65565b5050565b60008281526007602090815260408083208151608081018352815460ff81161515825261010090046001600160a01b0316818501526001820154818401526002820180548451818702810187019095528085528695929460608601939092919083018282801561038957602002820191906000526020600020905b815481526020019060010190808311610375575b50505050508152505090508060400151600014156103e95760405162461bcd60e51b815260206004820152601760248201527f7265717565737420494420697320696e636f727265637400000000000000000060448201526064016102e3565b806060015183815181106103ff576103ff61111c565b602002602001015191505092915050565b6003546002546006546040805160208101929092526001600160a01b0393841693634000aea09316918591015b6040516020818303038152906040526040518463ffffffff1660e01b815260040161046a93929190610ffa565b602060405180830381600087803b15801561048457600080fd5b505af1158015610498573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f69190610dff565b60065461050b5760405162461bcd60e51b815260206004820152600d60248201527f7375624944206e6f74207365740000000000000000000000000000000000000060448201526064016102e3565b60005b81518110156102f65760055460065483516001600160a01b039092169163bec4c08c91908590859081106105445761054461111c565b60200260200101516040518363ffffffff1660e01b815260040161057b9291909182526001600160a01b0316602082015260400190565b600060405180830381600087803b15801561059557600080fd5b505af11580156105a9573d6000803e3d6000fd5b5050505080806105b8906110f3565b91505061050e565b6001546001600160a01b0316331461061a5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016102e3565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6000546001600160a01b031633148015906106a457506002546001600160a01b03163314155b1561070e57336106bc6000546001600160a01b031690565b6002546040517f061db9c10000000000000000000000000000000000000000000000000000000081526001600160a01b03938416600482015291831660248301529190911660448201526064016102e3565b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60065461041057600560009054906101000a90046001600160a01b03166001600160a01b031663a21a23e46040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561079457600080fd5b505af11580156107a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cc9190610e35565b60068190556005546040517fbec4c08c00000000000000000000000000000000000000000000000000000000815260048101929092523060248301526001600160a01b03169063bec4c08c90604401600060405180830381600087803b15801561083557600080fd5b505af1158015610849573d6000803e3d6000fd5b505050506003546002546006546040516001600160a01b0393841693634000aea0931691859161043d919060200190815260200190565b60006040518060c0016040528084815260200160065481526020018661ffff1681526020018763ffffffff1681526020018563ffffffff1681526020016108d66040518060200160405280861515815250610af8565b90526002546040517f9b1c385e0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b0390911690639b1c385e90610927908590600401611039565b602060405180830381600087803b15801561094157600080fd5b505af1158015610955573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109799190610e35565b604080516080810182526000808252336020808401918252838501868152855184815280830187526060860190815287855260078352959093208451815493517fffffffffffffffffffffff0000000000000000000000000000000000000000009094169015157fffffffffffffffffffffff0000000000000000000000000000000000000000ff16176101006001600160a01b039094169390930292909217825591516001820155925180519495509193849392610a3f926002850192910190610ca9565b50505060049190915550505050505050565b610a59610b96565b610a6281610bf2565b50565b6004548214610ab65760405162461bcd60e51b815260206004820152601760248201527f7265717565737420494420697320696e636f727265637400000000000000000060448201526064016102e3565b60008281526007602090815260409091208251610adb92600290920191840190610ca9565b50506000908152600760205260409020805460ff19166001179055565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa82604051602401610b3191511515815260200190565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915292915050565b6000546001600160a01b03163314610bf05760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102e3565b565b6001600160a01b038116331415610c4b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102e3565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215610ce4579160200282015b82811115610ce4578251825591602001919060010190610cc9565b50610cf0929150610cf4565b5090565b5b80821115610cf05760008155600101610cf5565b80356001600160a01b0381168114610d2057600080fd5b919050565b803563ffffffff81168114610d2057600080fd5b600060208284031215610d4b57600080fd5b610d5482610d09565b9392505050565b60006020808385031215610d6e57600080fd5b823567ffffffffffffffff811115610d8557600080fd5b8301601f81018513610d9657600080fd5b8035610da9610da4826110cf565b61109e565b80828252848201915084840188868560051b8701011115610dc957600080fd5b600094505b83851015610df357610ddf81610d09565b835260019490940193918501918501610dce565b50979650505050505050565b600060208284031215610e1157600080fd5b8151610d5481611148565b600060208284031215610e2e57600080fd5b5035919050565b600060208284031215610e4757600080fd5b5051919050565b60008060408385031215610e6157600080fd5b8235915060208084013567ffffffffffffffff811115610e8057600080fd5b8401601f81018613610e9157600080fd5b8035610e9f610da4826110cf565b80828252848201915084840189868560051b8701011115610ebf57600080fd5b600094505b83851015610ee2578035835260019490940193918501918501610ec4565b5080955050505050509250929050565b60008060408385031215610f0557600080fd5b50508035926020909101359150565b600080600080600060a08688031215610f2c57600080fd5b610f3586610d25565b9450602086013561ffff81168114610f4c57600080fd5b9350610f5a60408701610d25565b9250606086013591506080860135610f7181611148565b809150509295509295909350565b600060208284031215610f9157600080fd5b81356bffffffffffffffffffffffff81168114610d5457600080fd5b6000815180845260005b81811015610fd357602081850181015186830182015201610fb7565b81811115610fe5576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b03841681526bffffffffffffffffffffffff831660208201526060604082015260006110306060830184610fad565b95945050505050565b60208152815160208201526020820151604082015261ffff60408301511660608201526000606083015163ffffffff80821660808501528060808601511660a0850152505060a083015160c08084015261109660e0840182610fad565b949350505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156110c7576110c7611132565b604052919050565b600067ffffffffffffffff8211156110e9576110e9611132565b5060051b60200190565b600060001982141561111557634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b8015158114610a6257600080fdfea164736f6c6343000806000a";

  BlockhashStore s_bhs;
  ExposedVRFCoordinatorV2_5 s_testCoordinator;
  ExposedVRFCoordinatorV2_5 s_testCoordinator_noLink;
  VRFV2PlusConsumerExample s_testConsumer;
  MockLinkToken s_linkToken;
  MockV3Aggregator s_linkNativeFeed;

  // VRF KeyV2 generated from a node; not sensitive information.
  // The secret key used to generate this key is: 10.
  bytes vrfUncompressedPublicKey =
    hex"a0434d9e47f3c86235477c7b1ae6ae5d3442d49b1943c2b752a68e2a47e247c7893aba425419bc27a3b6c7e693a24c696f794c2ed877a1593cbee53b037368d7";
  bytes vrfCompressedPublicKey = hex"a0434d9e47f3c86235477c7b1ae6ae5d3442d49b1943c2b752a68e2a47e247c701";
  bytes32 vrfKeyHash = hex"9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528";

  function setUp() public override {
    BaseTest.setUp();

    // Fund our users.
    vm.roll(1);
    vm.deal(LINK_WHALE, 10_000 ether);
    changePrank(LINK_WHALE);

    vm.txGasPrice(100 gwei);

    // Instantiate BHS.
    s_bhs = new BlockhashStore();

    // Deploy coordinator and consumer.
    // Note: adding contract deployments to this section will require the VRF proofs be regenerated.
    s_testCoordinator = new ExposedVRFCoordinatorV2_5(address(s_bhs));
    s_linkToken = new MockLinkToken();
    s_linkNativeFeed = new MockV3Aggregator(18, 500_000_000_000_000_000); // .5 ETH (good for testing)

    // Use create2 to deploy our consumer, so that its address is always the same
    // and surrounding changes do not alter our generated proofs.
    bytes memory consumerInitCode =
      bytes.concat(initializeCode, abi.encode(address(s_testCoordinator), address(s_linkToken)));
    bytes32 abiEncodedOwnerAddress = bytes32(uint256(uint160(LINK_WHALE)) << 96);
    address consumerCreate2Address;
    assembly {
      consumerCreate2Address :=
        create2(
          0, // value - left at zero here
          add(0x20, consumerInitCode), // initialization bytecode (excluding first memory slot which contains its length)
          mload(consumerInitCode), // length of initialization bytecode
          abiEncodedOwnerAddress // user-defined nonce to ensure unique SCA addresses
        )
    }
    s_testConsumer = VRFV2PlusConsumerExample(consumerCreate2Address);

    s_testCoordinator_noLink = new ExposedVRFCoordinatorV2_5(address(s_bhs));

    // Configure the coordinator.
    s_testCoordinator.setLINKAndLINKNativeFeed(address(s_linkToken), address(s_linkNativeFeed));
  }

  function setConfig() internal {
    s_testCoordinator.setConfig(
      0, // minRequestConfirmations
      2_500_000, // maxGasLimit
      1, // stalenessSeconds
      50_000, // gasAfterPaymentCalculation
      50_000_000_000_000_000, // fallbackWeiPerUnitLink
      500_000, // fulfillmentFlatFeeNativePPM
      100_000, // fulfillmentFlatFeeLinkDiscountPPM
      15, // nativePremiumPercentage
      10 // linkPremiumPercentage
    );
  }

  function testSetConfig() public {
    // Should setConfig successfully.
    setConfig();

    // Test that setting requestConfirmations above MAX_REQUEST_CONFIRMATIONS reverts.
    vm.expectRevert(abi.encodeWithSelector(VRFCoordinatorV2_5.InvalidRequestConfirmations.selector, 500, 500, 200));
    s_testCoordinator.setConfig(
      500,
      2_500_000,
      1,
      50_000,
      50_000_000_000_000_000,
      500_000, // fulfillmentFlatFeeNativePPM
      100_000, // fulfillmentFlatFeeLinkDiscountPPM
      15, // nativePremiumPercentage
      10 // linkPremiumPercentage
    );

    // Test that setting fallbackWeiPerUnitLink to zero reverts.
    vm.expectRevert(abi.encodeWithSelector(VRFCoordinatorV2_5.InvalidLinkWeiPrice.selector, 0));

    s_testCoordinator.setConfig(
      0,
      2_500_000,
      1,
      50_000,
      0,
      500_000, // fulfillmentFlatFeeNativePPM
      100_000, // fulfillmentFlatFeeLinkDiscountPPM
      15, // nativePremiumPercentage
      10 // linkPremiumPercentage
    );

    // Test that setting link discount flat fee higher than native flat fee reverts
    vm.expectRevert(abi.encodeWithSelector(VRFCoordinatorV2_5.LinkDiscountTooHigh.selector, uint32(501), uint32(500)));

    s_testCoordinator.setConfig(
      0,
      2_500_000,
      1,
      50_000,
      500,
      500, // fulfillmentFlatFeeNativePPM
      501, // fulfillmentFlatFeeLinkDiscountPPM
      15, // nativePremiumPercentage
      10 // linkPremiumPercentage
    );

    // // Test that setting link discount flat fee equal to native flat fee does not revert
    s_testCoordinator.setConfig(
      0,
      2_500_000,
      1,
      50_000,
      500,
      450, // fulfillmentFlatFeeNativePPM
      450, // fulfillmentFlatFeeLinkDiscountPPM
      15, // nativePremiumPercentage
      10 // linkPremiumPercentage
    );

    // Test that setting native premium percentage higher than 155 will revert
    vm.expectRevert(
      abi.encodeWithSelector(VRFCoordinatorV2_5.InvalidPremiumPercentage.selector, uint8(156), uint8(155))
    );

    s_testCoordinator.setConfig(
      0,
      2_500_000,
      1,
      50_000,
      500,
      500_000, // fulfillmentFlatFeeNativePPM
      100_000, // fulfillmentFlatFeeLinkDiscountPPM
      156, // nativePremiumPercentage
      10 // linkPremiumPercentage
    );

    // Test that setting LINK premium percentage higher than 155 will revert
    vm.expectRevert(
      abi.encodeWithSelector(VRFCoordinatorV2_5.InvalidPremiumPercentage.selector, uint8(202), uint8(155))
    );

    s_testCoordinator.setConfig(
      0,
      2_500_000,
      1,
      50_000,
      500,
      500_000, // fulfillmentFlatFeeNativePPM
      100_000, // fulfillmentFlatFeeLinkDiscountPPM
      15, // nativePremiumPercentage
      202 // linkPremiumPercentage
    );
  }

  function testRegisterProvingKey() public {
    // Should set the proving key successfully.
    registerProvingKey();

    // Should revert when already registered.
    uint256[2] memory uncompressedKeyParts = this.getProvingKeyParts(vrfUncompressedPublicKey);
    vm.expectRevert(abi.encodeWithSelector(VRFCoordinatorV2_5.ProvingKeyAlreadyRegistered.selector, vrfKeyHash));
    s_testCoordinator.registerProvingKey(uncompressedKeyParts, GAS_LANE_MAX_GAS);
  }

  event ProvingKeyRegistered(bytes32 keyHash, uint64 maxGas);
  event ProvingKeyDeregistered(bytes32 keyHash, uint64 maxGas);

  function registerProvingKey() public {
    uint256[2] memory uncompressedKeyParts = this.getProvingKeyParts(vrfUncompressedPublicKey);
    bytes32 keyHash = keccak256(abi.encode(uncompressedKeyParts));
    vm.expectEmit(
      false, // no indexed args to check for
      false, // no indexed args to check for
      false, // no indexed args to check for
      true
    ); // check data fields: keyHash and maxGas
    emit ProvingKeyRegistered(keyHash, GAS_LANE_MAX_GAS);
    s_testCoordinator.registerProvingKey(uncompressedKeyParts, GAS_LANE_MAX_GAS);
    (bool exists, uint64 maxGas) = s_testCoordinator.s_provingKeys(keyHash);
    assertTrue(exists);
    assertEq(GAS_LANE_MAX_GAS, maxGas);
    assertEq(s_testCoordinator.s_provingKeyHashes(0), keyHash);
    assertEq(keyHash, vrfKeyHash);
  }

  function testDeregisterProvingKey() public {
    // Should set the proving key successfully.
    registerProvingKey();

    bytes memory unregisteredPubKey =
      hex"6d919e4ed6add6c34b2af77eb6b2d2f5d27db11ba004e70734b23bd4321ea234ff8577a063314bead6d88c1b01849289a5542767a5138924f38fed551a7773db";

    // Should revert when given pubkey is not registered
    uint256[2] memory unregisteredKeyParts = this.getProvingKeyParts(unregisteredPubKey);
    bytes32 unregisterdKeyHash = keccak256(abi.encode(unregisteredKeyParts));
    vm.expectRevert(abi.encodeWithSelector(VRFCoordinatorV2_5.NoSuchProvingKey.selector, unregisterdKeyHash));
    s_testCoordinator.deregisterProvingKey(unregisteredKeyParts);

    // correctly deregister pubkey
    uint256[2] memory uncompressedKeyParts = this.getProvingKeyParts(vrfUncompressedPublicKey);
    bytes32 keyHash = keccak256(abi.encode(uncompressedKeyParts));
    vm.expectEmit(
      false, // no indexed args to check for
      false, // no indexed args to check for
      false, // no indexed args to check for
      true
    ); // check data fields: keyHash and maxGas
    emit ProvingKeyDeregistered(keyHash, GAS_LANE_MAX_GAS);
    s_testCoordinator.deregisterProvingKey(uncompressedKeyParts);
    (bool exists, uint64 maxGas) = s_testCoordinator.s_provingKeys(keyHash);
    assertFalse(exists);
    assertEq(0, maxGas);
  }

  // note: Call this function via this.getProvingKeyParts to be able to pass memory as calldata and
  // index over the byte array.
  function getProvingKeyParts(
    bytes calldata uncompressedKey
  ) public pure returns (uint256[2] memory) {
    uint256 keyPart1 = uint256(bytes32(uncompressedKey[0:32]));
    uint256 keyPart2 = uint256(bytes32(uncompressedKey[32:64]));
    return [keyPart1, keyPart2];
  }

  function testCreateSubscription() public {
    uint256 subId = s_testCoordinator.createSubscription();
    s_testCoordinator.fundSubscriptionWithNative{value: 10 ether}(subId);
  }

  function testCancelSubWithNoLink() public {
    uint256 subId = s_testCoordinator_noLink.createSubscription();
    s_testCoordinator_noLink.fundSubscriptionWithNative{value: 1000 ether}(subId);

    assertEq(LINK_WHALE.balance, 9000 ether);
    s_testCoordinator_noLink.cancelSubscription(subId, LINK_WHALE);
    assertEq(LINK_WHALE.balance, 10_000 ether);

    vm.expectRevert(SubscriptionAPI.InvalidSubscription.selector);
    s_testCoordinator_noLink.getSubscription(subId);
  }

  function testGetActiveSubscriptionIds() public {
    uint256 numSubs = 40;
    for (uint256 i = 0; i < numSubs; i++) {
      s_testCoordinator.createSubscription();
    }
    // get all subscriptions, assert length is correct
    uint256[] memory allSubs = s_testCoordinator.getActiveSubscriptionIds(0, 0);
    assertEq(allSubs.length, s_testCoordinator.getActiveSubscriptionIdsLength());

    // paginate through subscriptions, batching by 10.
    // we should eventually get all the subscriptions this way.
    uint256[][] memory subIds = paginateSubscriptions(s_testCoordinator, 10);
    // check that all subscriptions were returned
    uint256 actualNumSubs = 0;
    for (uint256 batchIdx = 0; batchIdx < subIds.length; batchIdx++) {
      for (uint256 subIdx = 0; subIdx < subIds[batchIdx].length; subIdx++) {
        s_testCoordinator.getSubscription(subIds[batchIdx][subIdx]);
        actualNumSubs++;
      }
    }
    assertEq(actualNumSubs, s_testCoordinator.getActiveSubscriptionIdsLength());

    // cancel a bunch of subscriptions, assert that they are not returned
    uint256[] memory subsToCancel = new uint256[](3);
    for (uint256 i = 0; i < 3; i++) {
      subsToCancel[i] = subIds[0][i];
    }
    for (uint256 i = 0; i < subsToCancel.length; i++) {
      s_testCoordinator.cancelSubscription(subsToCancel[i], LINK_WHALE);
    }
    uint256[][] memory newSubIds = paginateSubscriptions(s_testCoordinator, 10);
    // check that all subscriptions were returned
    // and assert that none of the canceled subscriptions are returned
    actualNumSubs = 0;
    for (uint256 batchIdx = 0; batchIdx < newSubIds.length; batchIdx++) {
      for (uint256 subIdx = 0; subIdx < newSubIds[batchIdx].length; subIdx++) {
        for (uint256 i = 0; i < subsToCancel.length; i++) {
          assertFalse(newSubIds[batchIdx][subIdx] == subsToCancel[i]);
        }
        s_testCoordinator.getSubscription(newSubIds[batchIdx][subIdx]);
        actualNumSubs++;
      }
    }
    assertEq(actualNumSubs, s_testCoordinator.getActiveSubscriptionIdsLength());
  }

  function paginateSubscriptions(
    ExposedVRFCoordinatorV2_5 coordinator,
    uint256 batchSize
  ) internal view returns (uint256[][] memory) {
    uint256 arrIndex = 0;
    uint256 startIndex = 0;
    uint256 numSubs = coordinator.getActiveSubscriptionIdsLength();
    uint256[][] memory subIds = new uint256[][](Math.ceilDiv(numSubs, batchSize));
    while (startIndex < numSubs) {
      subIds[arrIndex] = coordinator.getActiveSubscriptionIds(startIndex, batchSize);
      startIndex += batchSize;
      arrIndex++;
    }
    return subIds;
  }

  event RandomWordsRequested(
    bytes32 indexed keyHash,
    uint256 requestId,
    uint256 preSeed,
    uint256 indexed subId,
    uint16 minimumRequestConfirmations,
    uint32 callbackGasLimit,
    uint32 numWords,
    bytes extraArgs,
    address indexed sender
  );
  event RandomWordsFulfilled(
    uint256 indexed requestId, uint256 outputSeed, uint256 indexed subID, uint96 payment, bytes extraArgs, bool success
  );
  event FallbackWeiPerUnitLinkUsed(uint256 requestId, int256 fallbackWeiPerUnitLink);

  function testRequestAndFulfillRandomWordsNative() public {
    (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc, uint256 subId, uint256 requestId) =
      setupSubAndRequestRandomnessNativePayment();
    (, uint96 nativeBalanceBefore,,,) = s_testCoordinator.getSubscription(subId);

    uint256 outputSeed = s_testCoordinator.getRandomnessFromProofExternal(proof, rc).randomness;
    vm.recordLogs();
    uint96 payment = s_testCoordinator.fulfillRandomWords(proof, rc, false);
    VmSafe.Log[] memory entries = vm.getRecordedLogs();
    assertEq(entries[0].topics[1], bytes32(uint256(requestId)));
    assertEq(entries[0].topics[2], bytes32(uint256(subId)));
    (uint256 loggedOutputSeed,,, bool loggedSuccess) = abi.decode(entries[0].data, (uint256, uint256, bool, bool));
    assertEq(loggedOutputSeed, outputSeed);
    assertEq(loggedSuccess, true);

    (bool fulfilled,,) = s_testConsumer.s_requests(requestId);
    assertEq(fulfilled, true);

    // The cost of fulfillRandomWords is approximately 70_000 gas.
    // gasAfterPaymentCalculation is 50_000.
    //
    // The cost of the VRF fulfillment charged to the user is:
    // baseFeeWei = weiPerUnitGas * (gasAfterPaymentCalculation + startGas - gasleft())
    // baseFeeWei = 1e11 * (50_000 + 70_000)
    // baseFeeWei = 1.2e16
    // flatFeeWei = 1e12 * (fulfillmentFlatFeeNativePPM)
    // flatFeeWei = 1e12 * 500_000 = 5e17
    // ...
    // billed_fee = baseFeeWei * (100 + linkPremiumPercentage / 100) + 5e17
    // billed_fee = 1.2e16 * 1.15 + 5e17
    // billed_fee = 5.138e+17
    (, uint96 nativeBalanceAfter,,,) = s_testCoordinator.getSubscription(subId);
    // 1e15 is less than 1 percent discrepancy
    assertApproxEqAbs(payment, 5.138 * 1e17, 1e15);
    assertApproxEqAbs(nativeBalanceAfter, nativeBalanceBefore - 5.138 * 1e17, 1e15);
    assertFalse(s_testCoordinator.pendingRequestExists(subId));
  }

  function testRequestAndFulfillRandomWordsLINK() public {
    (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc, uint256 subId, uint256 requestId) =
      setupSubAndRequestRandomnessLINKPayment();
    (uint96 linkBalanceBefore,,,,) = s_testCoordinator.getSubscription(subId);

    uint256 outputSeed = s_testCoordinator.getRandomnessFromProofExternal(proof, rc).randomness;
    vm.recordLogs();
    uint96 payment = s_testCoordinator.fulfillRandomWords(proof, rc, false);

    VmSafe.Log[] memory entries = vm.getRecordedLogs();
    assertEq(entries[0].topics[1], bytes32(uint256(requestId)));
    assertEq(entries[0].topics[2], bytes32(uint256(subId)));
    (uint256 loggedOutputSeed,,, bool loggedSuccess) = abi.decode(entries[0].data, (uint256, uint256, bool, bool));
    assertEq(loggedOutputSeed, outputSeed);
    assertEq(loggedSuccess, true);

    (bool fulfilled,,) = s_testConsumer.s_requests(requestId);
    assertEq(fulfilled, true);

    // The cost of fulfillRandomWords is approximately 86_000 gas.
    // gasAfterPaymentCalculation is 50_000.
    //
    // The cost of the VRF fulfillment charged to the user is:
    // paymentNoFee = (weiPerUnitGas * (gasAfterPaymentCalculation + startGas - gasleft() + l1CostWei) /
    // link_native_ratio)
    // paymentNoFee = (1e11 * (50_000 + 86_000 + 0)) / .5
    // paymentNoFee = 2.72e16
    // flatFeeWei = 1e12 * (fulfillmentFlatFeeNativePPM - fulfillmentFlatFeeLinkDiscountPPM)
    // flatFeeWei = 1e12 * (500_000 - 100_000)
    // flatFeeJuels = 1e18 * flatFeeWei / link_native_ratio
    // flatFeeJuels = 4e17 / 0.5 = 8e17
    // billed_fee = paymentNoFee * ((100 + 10) / 100) + 8e17
    // billed_fee = 2.72e16 * 1.1 + 8e17
    // billed_fee = 2.992e16 + 8e17 = 8.2992e17
    // note: delta is doubled from the native test to account for more variance due to the link/native ratio
    (uint96 linkBalanceAfter,,,,) = s_testCoordinator.getSubscription(subId);
    // 1e15 is less than 1 percent discrepancy
    assertApproxEqAbs(payment, 8.2992 * 1e17, 1e15);
    assertApproxEqAbs(linkBalanceAfter, linkBalanceBefore - 8.2992 * 1e17, 1e15);
    assertFalse(s_testCoordinator.pendingRequestExists(subId));
  }

  function testRequestAndFulfillRandomWordsLINK_FallbackWeiPerUnitLinkUsed() public {
    (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc,, uint256 requestId) =
      setupSubAndRequestRandomnessLINKPayment();

    (,,, uint32 stalenessSeconds,,,,,) = s_testCoordinator.s_config();
    int256 fallbackWeiPerUnitLink = s_testCoordinator.s_fallbackWeiPerUnitLink();

    // Set the link feed to be stale.
    (uint80 roundId, int256 answer, uint256 startedAt,,) = s_linkNativeFeed.latestRoundData();
    uint256 timestamp = block.timestamp - stalenessSeconds - 1;
    s_linkNativeFeed.updateRoundData(roundId, answer, timestamp, startedAt);

    vm.expectEmit(false, false, false, true, address(s_testCoordinator));
    emit FallbackWeiPerUnitLinkUsed(requestId, fallbackWeiPerUnitLink);
    s_testCoordinator.fulfillRandomWords(proof, rc, false);
  }

  function setupSubAndRequestRandomnessLINKPayment()
    internal
    returns (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc, uint256 subId, uint256 requestId)
  {
    uint32 requestBlock = 20;
    vm.roll(requestBlock);
    s_linkToken.transfer(address(s_testConsumer), 10 ether);
    s_testConsumer.createSubscriptionAndFund(10 ether);
    subId = s_testConsumer.s_subId();

    // Apply basic configs to contract.
    setConfig();
    registerProvingKey();

    // Request random words.
    vm.expectEmit(true, true, false, true);
    uint256 preSeed;
    (requestId, preSeed) = s_testCoordinator.computeRequestIdExternal(vrfKeyHash, address(s_testConsumer), subId, 1);
    emit RandomWordsRequested(
      vrfKeyHash,
      requestId,
      preSeed,
      subId,
      MIN_CONFIRMATIONS,
      CALLBACK_GAS_LIMIT,
      NUM_WORDS,
      VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: false})), // nativePayment, //
        // nativePayment
      address(s_testConsumer) // requester
    );
    s_testConsumer.requestRandomWords(CALLBACK_GAS_LIMIT, MIN_CONFIRMATIONS, NUM_WORDS, vrfKeyHash, false);
    (bool fulfilled,,) = s_testConsumer.s_requests(requestId);
    assertEq(fulfilled, false);
    assertTrue(s_testCoordinator.pendingRequestExists(subId));

    // Move on to the next block.
    // Store the previous block's blockhash, and assert that it is as expected.
    vm.roll(requestBlock + 1);
    s_bhs.store(requestBlock);
    assertEq(hex"731dc163f73d31d8c68f9917ce4ff967753939f70432973c04fd2c2a48148607", s_bhs.getBlockhash(requestBlock));

    // Fulfill the request.
    // Proof generated via the generate-proof-v2-plus script command.
    // 1st step: Uncomment these 3 console logs to see info about the request and run the test to get output:
    // console.log("requestId: ", requestId);
    // console.log("preSeed: ", preSeed);
    // console.log("sender: ", address(s_testConsumer));
    // 2nd step: Update pre-seed in the command commented out below with new value printed in console logs.
    // 3rd step: export the following environment variables to run the generate-proof-v2-plus script.
    // export ETH_URL=https://ethereum-sepolia-rpc.publicnode.com # or any other RPC provider you prefer
    // export ETH_CHAIN_ID=11155111 # or switch to any other chain
    // export ACCOUNT_KEY=<your test EOA private key>
    // 4th step: run the command and copy the command output in the proof section below.
    /*
        Run from this folder: chainlink/core/scripts/vrfv2plus/testnet
        go run . generate-proof-v2-plus \
        -key-hash 0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528 \
        -pre-seed 77134414723242246520332717536018735794426514244521954002798799849127623496871 \
        -block-hash 0x731dc163f73d31d8c68f9917ce4ff967753939f70432973c04fd2c2a48148607 \
        -block-num 20 \
        -sender 0x90A8820424CC8a819d14cBdE54D12fD3fbFa9bb2 \
        -native-payment false
    */
    proof = VRF.Proof({
      pk: [
        72_488_970_228_380_509_287_422_715_226_575_535_698_893_157_273_063_074_627_791_787_432_852_706_183_111,
        62_070_622_898_698_443_831_883_535_403_436_258_712_770_888_294_397_026_493_185_421_712_108_624_767_191
      ],
      gamma: [
        103_927_982_338_770_370_318_312_316_555_080_928_288_985_522_873_495_041_111_817_988_974_598_585_393_796,
        56_789_421_278_806_198_480_964_888_112_155_620_425_048_056_183_534_931_202_752_833_185_923_411_715_624
      ],
      c: 23_645_475_075_665_525_321_781_505_993_434_124_657_388_421_977_074_956_645_288_621_921_391_376_468_128,
      s: 106_817_081_950_846_808_215_350_231_311_242_951_539_230_271_757_396_902_089_035_477_907_017_240_898_689,
      seed: 77_134_414_723_242_246_520_332_717_536_018_735_794_426_514_244_521_954_002_798_799_849_127_623_496_871,
      uWitness: 0xD6899602060d574DE03FE1cf76fDf66afE12d549,
      cGammaWitness: [
        9_892_458_071_712_426_452_033_749_279_561_067_220_589_549_155_902_380_165_087_951_541_202_159_693_388,
        61_235_995_320_721_681_444_549_354_910_430_438_435_754_757_626_312_862_714_628_885_100_042_911_955_139
      ],
      sHashWitness: [
        101_478_618_362_722_903_511_580_105_256_015_180_591_690_884_037_598_276_249_676_652_094_434_483_808_775,
        82_512_235_485_399_822_034_680_598_942_438_982_472_006_937_353_405_384_896_956_013_889_074_719_896_188
      ],
      zInv: 82_281_039_329_215_616_805_111_360_985_152_709_712_368_762_415_186_906_218_863_971_780_664_103_705_723
    });
    rc = VRFTypes.RequestCommitmentV2Plus({
      blockNum: requestBlock,
      subId: subId,
      callbackGasLimit: 1_000_000,
      numWords: 1,
      sender: address(s_testConsumer),
      extraArgs: VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: false}))
    });
    return (proof, rc, subId, requestId);
  }

  function setupSubAndRequestRandomnessNativePayment()
    internal
    returns (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc, uint256 subId, uint256 requestId)
  {
    uint32 requestBlock = 10;
    vm.roll(requestBlock);
    s_testConsumer.createSubscriptionAndFund(0);
    subId = s_testConsumer.s_subId();
    s_testCoordinator.fundSubscriptionWithNative{value: 10 ether}(subId);

    // Apply basic configs to contract.
    setConfig();
    registerProvingKey();

    // Request random words.
    vm.expectEmit(true, true, true, true);
    uint256 preSeed;
    (requestId, preSeed) = s_testCoordinator.computeRequestIdExternal(vrfKeyHash, address(s_testConsumer), subId, 1);
    emit RandomWordsRequested(
      vrfKeyHash,
      requestId,
      preSeed,
      subId,
      MIN_CONFIRMATIONS,
      CALLBACK_GAS_LIMIT,
      NUM_WORDS,
      VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: true})), // nativePayment
      address(s_testConsumer) // requester
    );
    s_testConsumer.requestRandomWords(CALLBACK_GAS_LIMIT, MIN_CONFIRMATIONS, NUM_WORDS, vrfKeyHash, true);
    (bool fulfilled,,) = s_testConsumer.s_requests(requestId);
    assertEq(fulfilled, false);
    assertTrue(s_testCoordinator.pendingRequestExists(subId));

    // Move on to the next block.
    // Store the previous block's blockhash, and assert that it is as expected.
    vm.roll(requestBlock + 1);
    s_bhs.store(requestBlock);
    assertEq(hex"1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac", s_bhs.getBlockhash(requestBlock));

    // Fulfill the request.
    // Proof generated via the generate-proof-v2-plus script command.
    // 1st step: Uncomment these 3 console logs to see info about the request and run the test to get output:
    // console.log("requestId: ", requestId);
    // console.log("preSeed: ", preSeed);
    // console.log("sender: ", address(s_testConsumer));
    // 2nd step: Update pre-seed in the command commented out below with new value printed in console logs.
    // 3rd step: export the following environment variables to run the generate-proof-v2-plus script.
    // export ETH_URL=https://ethereum-sepolia-rpc.publicnode.com # or any other RPC provider you prefer
    // export ETH_CHAIN_ID=11155111 # or switch to any other chain
    // export ACCOUNT_KEY=<your test EOA private key>
    // 4th step: run the command and copy the command output in the proof section below.
    /*
       Run from this folder: chainlink/core/scripts/vrfv2plus/testnet
       go run . generate-proof-v2-plus \
        -key-hash 0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528 \
        -pre-seed 88177119495082281213609405072572269421661478022189589823108119237563684383163 \
        -block-hash 0x1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac \
        -block-num 10 \
        -sender 0x90A8820424CC8a819d14cBdE54D12fD3fbFa9bb2 \
        -native-payment true
        */
    proof = VRF.Proof({
      pk: [
        72_488_970_228_380_509_287_422_715_226_575_535_698_893_157_273_063_074_627_791_787_432_852_706_183_111,
        62_070_622_898_698_443_831_883_535_403_436_258_712_770_888_294_397_026_493_185_421_712_108_624_767_191
      ],
      gamma: [
        102_142_782_721_757_938_350_759_722_545_721_736_888_276_217_484_353_597_703_162_772_276_193_136_052_353,
        87_167_280_284_008_869_627_768_921_028_415_708_350_806_510_214_000_539_818_296_353_518_495_698_939_660
      ],
      c: 78_738_462_581_063_211_677_832_865_654_743_924_688_552_792_392_007_862_664_964_608_134_754_001_810_280,
      s: 97_066_881_804_257_970_453_329_086_439_696_419_448_135_613_089_654_606_517_271_688_187_030_953_014_593,
      seed: 88_177_119_495_082_281_213_609_405_072_572_269_421_661_478_022_189_589_823_108_119_237_563_684_383_163,
      uWitness: 0xa335ea8dF652d5331a276B60b16c9733435D4f73,
      cGammaWitness: [
        114_435_126_227_922_602_743_444_254_494_036_972_095_649_501_991_695_809_092_954_325_430_947_992_864_624,
        63_032_211_040_463_927_862_594_425_238_691_911_311_087_931_119_674_607_521_158_894_139_074_063_158_678
      ],
      sHashWitness: [
        105_043_781_471_073_183_057_173_130_563_345_930_784_924_139_079_040_814_418_442_661_347_864_735_908_726,
        68_696_469_914_696_211_053_833_437_482_938_344_908_217_760_552_761_185_546_164_836_556_562_945_431_554
      ],
      zInv: 73_325_637_847_357_165_955_904_789_471_972_164_751_975_373_195_750_497_508_525_598_331_798_833_112_175
    });
    rc = VRFTypes.RequestCommitmentV2Plus({
      blockNum: requestBlock,
      subId: subId,
      callbackGasLimit: CALLBACK_GAS_LIMIT,
      numWords: 1,
      sender: address(s_testConsumer),
      extraArgs: VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: true}))
    });

    return (proof, rc, subId, requestId);
  }

  function testRequestAndFulfillRandomWords_NetworkGasPriceExceedsGasLane() public {
    (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc,,) = setupSubAndRequestRandomnessNativePayment();

    // network gas is higher than gas lane max gas
    uint256 networkGasPrice = GAS_LANE_MAX_GAS + 1;
    vm.txGasPrice(networkGasPrice);
    vm.expectRevert(
      abi.encodeWithSelector(VRFCoordinatorV2_5.GasPriceExceeded.selector, networkGasPrice, GAS_LANE_MAX_GAS)
    );
    s_testCoordinator.fulfillRandomWords(proof, rc, false);
  }

  function testRequestAndFulfillRandomWords_OnlyPremium_NativePayment() public {
    (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc, uint256 subId, uint256 requestId) =
      setupSubAndRequestRandomnessNativePayment();
    (, uint96 nativeBalanceBefore,,,) = s_testCoordinator.getSubscription(subId);

    // network gas is twice the gas lane max gas
    uint256 networkGasPrice = GAS_LANE_MAX_GAS * 2;
    vm.txGasPrice(networkGasPrice);

    uint256 outputSeed = s_testCoordinator.getRandomnessFromProofExternal(proof, rc).randomness;
    vm.recordLogs();
    uint96 payment = s_testCoordinator.fulfillRandomWords(proof, rc, true /* onlyPremium */ );
    VmSafe.Log[] memory entries = vm.getRecordedLogs();
    assertEq(entries[0].topics[1], bytes32(uint256(requestId)));
    assertEq(entries[0].topics[2], bytes32(uint256(subId)));
    (uint256 loggedOutputSeed,,, bool loggedSuccess) = abi.decode(entries[0].data, (uint256, uint256, bool, bool));
    assertEq(loggedOutputSeed, outputSeed);
    assertEq(loggedSuccess, true);

    (bool fulfilled,,) = s_testConsumer.s_requests(requestId);
    assertEq(fulfilled, true);

    // The cost of fulfillRandomWords is approximately 72_100 gas.
    // gasAfterPaymentCalculation is 50_000.
    //
    // The cost of the VRF fulfillment charged to the user is:
    // baseFeeWei = weiPerUnitGas * (gasAfterPaymentCalculation + startGas - gasleft())
    // network gas price is capped at gas lane max gas (5000 gwei)
    // baseFeeWei = 5e12 * (50_000 + 72_100)
    // baseFeeWei = 6.11e17
    // flatFeeWei = 1e12 * (fulfillmentFlatFeeNativePPM)
    // flatFeeWei = 1e12 * 500_000 = 5e17
    // ...
    // billed_fee = baseFeeWei * (linkPremiumPercentage / 100) + 5e17
    // billed_fee = 6.11e17 * 0.15 + 5e17
    // billed_fee = 5.9157e+17
    (, uint96 nativeBalanceAfter,,,) = s_testCoordinator.getSubscription(subId);
    // 1e15 is less than 1 percent discrepancy
    assertApproxEqAbs(payment, 5.9157 * 1e17, 1e15);
    assertApproxEqAbs(nativeBalanceAfter, nativeBalanceBefore - 5.9157 * 1e17, 1e15);
    assertFalse(s_testCoordinator.pendingRequestExists(subId));
  }

  function testRequestAndFulfillRandomWords_OnlyPremium_LinkPayment() public {
    (VRF.Proof memory proof, VRFTypes.RequestCommitmentV2Plus memory rc, uint256 subId, uint256 requestId) =
      setupSubAndRequestRandomnessLINKPayment();
    (uint96 linkBalanceBefore,,,,) = s_testCoordinator.getSubscription(subId);

    // network gas is twice the gas lane max gas
    uint256 networkGasPrice = GAS_LANE_MAX_GAS * 5;
    vm.txGasPrice(networkGasPrice);

    uint256 outputSeed = s_testCoordinator.getRandomnessFromProofExternal(proof, rc).randomness;
    vm.recordLogs();
    uint96 payment = s_testCoordinator.fulfillRandomWords(proof, rc, true /* onlyPremium */ );

    VmSafe.Log[] memory entries = vm.getRecordedLogs();
    assertEq(entries[0].topics[1], bytes32(uint256(requestId)));
    assertEq(entries[0].topics[2], bytes32(uint256(subId)));
    (uint256 loggedOutputSeed,,, bool loggedSuccess) = abi.decode(entries[0].data, (uint256, uint256, bool, bool));
    assertEq(loggedOutputSeed, outputSeed);
    assertEq(loggedSuccess, true);

    (bool fulfilled,,) = s_testConsumer.s_requests(requestId);
    assertEq(fulfilled, true);

    // The cost of fulfillRandomWords is approximately 89_100 gas.
    // gasAfterPaymentCalculation is 50_000.
    //
    // The cost of the VRF fulfillment charged to the user is:
    // paymentNoFee = (weiPerUnitGas * (gasAfterPaymentCalculation + startGas - gasleft() + l1CostWei) /
    // link_native_ratio)
    // network gas price is capped at gas lane max gas (5000 gwei)
    // paymentNoFee = (5e12 * (50_000 + 89_100 + 0)) / .5
    // paymentNoFee = 1.391e+18
    // flatFeeWei = 1e12 * (fulfillmentFlatFeeNativePPM - fulfillmentFlatFeeLinkDiscountPPM)
    // flatFeeWei = 1e12 * (500_000 - 100_000)
    // flatFeeJuels = 1e18 * flatFeeWei / link_native_ratio
    // flatFeeJuels = 4e17 / 0.5 = 8e17
    // billed_fee = paymentNoFee * (10 / 100) + 8e17
    // billed_fee = 1.391e+18 * 0.1 + 8e17
    // billed_fee = 9.391e+17
    // note: delta is doubled from the native test to account for more variance due to the link/native ratio
    (uint96 linkBalanceAfter,,,,) = s_testCoordinator.getSubscription(subId);
    // 1e15 is less than 1 percent discrepancy
    assertApproxEqAbs(payment, 9.391 * 1e17, 1e15);
    assertApproxEqAbs(linkBalanceAfter, linkBalanceBefore - 9.391 * 1e17, 1e15);
    assertFalse(s_testCoordinator.pendingRequestExists(subId));
  }

  function testRequestRandomWords_InvalidConsumer() public {
    address subOwner = makeAddr("subOwner");
    changePrank(subOwner);
    uint256 subId = s_testCoordinator.createSubscription();
    VRFV2PlusLoadTestWithMetrics consumer = new VRFV2PlusLoadTestWithMetrics(address(s_testCoordinator));

    // consumer is not added to the subscription
    vm.expectRevert(abi.encodeWithSelector(SubscriptionAPI.InvalidConsumer.selector, subId, address(consumer)));
    consumer.requestRandomWords(
      subId, MIN_CONFIRMATIONS, vrfKeyHash, CALLBACK_GAS_LIMIT, true, NUM_WORDS, 1 /* requestCount */
    );
    assertFalse(s_testCoordinator.pendingRequestExists(subId));
  }

  function testRequestRandomWords_ReAddConsumer_AssertRequestID() public {
    // 1. setup consumer and subscription
    setConfig();
    registerProvingKey();
    address subOwner = makeAddr("subOwner");
    changePrank(subOwner);
    uint256 subId = s_testCoordinator.createSubscription();
    VRFV2PlusLoadTestWithMetrics consumer = createAndAddLoadTestWithMetricsConsumer(subId);
    uint32 requestBlock = 10;
    vm.roll(requestBlock);
    changePrank(LINK_WHALE);
    s_testCoordinator.fundSubscriptionWithNative{value: 10 ether}(subId);

    // 2. Request random words.
    changePrank(subOwner);
    vm.expectEmit(true, true, false, true);
    uint256 requestId;
    uint256 preSeed;
    (requestId, preSeed) = s_testCoordinator.computeRequestIdExternal(vrfKeyHash, address(consumer), subId, 1);
    emit RandomWordsRequested(
      vrfKeyHash,
      requestId,
      preSeed,
      subId,
      MIN_CONFIRMATIONS,
      CALLBACK_GAS_LIMIT,
      NUM_WORDS,
      VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: true})),
      address(consumer) // requester
    );
    consumer.requestRandomWords(
      subId,
      MIN_CONFIRMATIONS,
      vrfKeyHash,
      CALLBACK_GAS_LIMIT,
      true, /* nativePayment */
      NUM_WORDS,
      1 /* requestCount */
    );
    assertTrue(s_testCoordinator.pendingRequestExists(subId));

    // Move on to the next block.
    // Store the previous block's blockhash, and assert that it is as expected.
    vm.roll(requestBlock + 1);
    s_bhs.store(requestBlock);
    assertEq(hex"1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac", s_bhs.getBlockhash(requestBlock));

    // 3. Fulfill the request above
    // Proof generated via the generate-proof-v2-plus script command.
    // 1st step: Uncomment these 3 console logs to see info about the request and run the test to get output:
    // console.log("requestId: ", requestId);
    // console.log("preSeed: ", preSeed);
    // console.log("sender: ", address(s_testConsumer));
    // 2nd step: Update pre-seed in the command commented out below with new value printed in console logs.
    // 3rd step: export the following environment variables to run the generate-proof-v2-plus script.
    // export ETH_URL=https://ethereum-sepolia-rpc.publicnode.com # or any other RPC provider you prefer
    // export ETH_CHAIN_ID=11155111 # or switch to any other chain
    // export ACCOUNT_KEY=<your test EOA private key>
    // 4th step: run the command and copy the command output in the proof section below.
    /*
      Run from this folder: chainlink/core/scripts/vrfv2plus/testnet
      go run . generate-proof-v2-plus \
      -key-hash 0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528 \
      -pre-seed 78857362017365444144484359594634073685493503942324326290718892836953423263381 \
      -block-hash 0x1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac \
      -block-num 10 \
      -sender 0x44CAfC03154A0708F9DCf988681821f648dA74aF \
      -native-payment true
    */
    VRF.Proof memory proof = VRF.Proof({
      pk: [
        72_488_970_228_380_509_287_422_715_226_575_535_698_893_157_273_063_074_627_791_787_432_852_706_183_111,
        62_070_622_898_698_443_831_883_535_403_436_258_712_770_888_294_397_026_493_185_421_712_108_624_767_191
      ],
      gamma: [
        65_913_937_398_148_449_626_792_563_067_325_648_649_534_055_460_473_988_721_938_103_219_381_973_178_278,
        63_156_327_344_180_203_180_831_822_252_171_874_192_175_272_818_200_597_638_000_091_892_096_122_362_120
      ],
      c: 96_524_997_218_413_735_279_221_574_381_819_903_278_651_909_890_109_201_564_980_667_824_986_706_861_580,
      s: 32_941_032_142_956_097_592_442_894_642_111_025_677_491_308_239_274_769_364_799_856_748_447_418_202_313,
      seed: 78_857_362_017_365_444_144_484_359_594_634_073_685_493_503_942_324_326_290_718_892_836_953_423_263_381,
      uWitness: 0xda613621Dc2347d9A6670a1cBA812d52A7ec3A3A,
      cGammaWitness: [
        6_776_842_114_900_054_689_355_891_239_487_365_968_068_230_823_400_902_903_493_665_825_747_641_410_781,
        753_482_930_067_864_853_610_521_010_650_481_816_782_338_376_846_697_006_021_590_704_037_205_560_592
      ],
      sHashWitness: [
        76_619_528_582_417_858_778_905_184_311_764_104_068_650_968_652_636_772_643_050_945_629_834_129_417_915,
        27_947_566_794_040_118_487_986_033_070_014_357_750_801_611_688_958_204_148_187_927_873_566_412_002_355
      ],
      zInv: 77_351_076_831_418_813_780_936_064_446_565_588_198_113_457_019_145_030_499_544_500_588_309_236_458_362
    });
    VRFTypes.RequestCommitmentV2Plus memory rc = VRFTypes.RequestCommitmentV2Plus({
      blockNum: requestBlock,
      subId: subId,
      callbackGasLimit: CALLBACK_GAS_LIMIT,
      numWords: NUM_WORDS,
      sender: address(consumer),
      extraArgs: VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: true}))
    });
    s_testCoordinator.fulfillRandomWords(proof, rc, true /* onlyPremium */ );
    assertFalse(s_testCoordinator.pendingRequestExists(subId));

    // 4. remove consumer and verify request random words doesn't work
    s_testCoordinator.removeConsumer(subId, address(consumer));
    vm.expectRevert(abi.encodeWithSelector(SubscriptionAPI.InvalidConsumer.selector, subId, address(consumer)));
    consumer.requestRandomWords(
      subId,
      MIN_CONFIRMATIONS,
      vrfKeyHash,
      CALLBACK_GAS_LIMIT,
      false, /* nativePayment */
      NUM_WORDS,
      1 /* requestCount */
    );

    // 5. re-add consumer and assert requestID nonce starts from 2 (nonce 1 was used before consumer removal)
    s_testCoordinator.addConsumer(subId, address(consumer));
    vm.expectEmit(true, true, false, true);
    uint256 requestId2;
    uint256 preSeed2;
    (requestId2, preSeed2) = s_testCoordinator.computeRequestIdExternal(vrfKeyHash, address(consumer), subId, 2);
    emit RandomWordsRequested(
      vrfKeyHash,
      requestId2,
      preSeed2,
      subId,
      MIN_CONFIRMATIONS,
      CALLBACK_GAS_LIMIT,
      NUM_WORDS,
      VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: false})), // nativePayment, //
        // nativePayment
      address(consumer) // requester
    );
    consumer.requestRandomWords(
      subId,
      MIN_CONFIRMATIONS,
      vrfKeyHash,
      CALLBACK_GAS_LIMIT,
      false, /* nativePayment */
      NUM_WORDS,
      1 /* requestCount */
    );
    assertNotEq(requestId, requestId2);
    assertNotEq(preSeed, preSeed2);
    assertTrue(s_testCoordinator.pendingRequestExists(subId));
  }

  function testRequestRandomWords_MultipleConsumers_PendingRequestExists() public {
    // 1. setup consumer and subscription
    setConfig();
    registerProvingKey();
    address subOwner = makeAddr("subOwner");
    changePrank(subOwner);
    uint256 subId = s_testCoordinator.createSubscription();
    VRFV2PlusLoadTestWithMetrics consumer1 = createAndAddLoadTestWithMetricsConsumer(subId);
    VRFV2PlusLoadTestWithMetrics consumer2 = createAndAddLoadTestWithMetricsConsumer(subId);
    uint32 requestBlock = 10;
    vm.roll(requestBlock);
    changePrank(LINK_WHALE);
    s_testCoordinator.fundSubscriptionWithNative{value: 10 ether}(subId);

    // 2. Request random words.
    changePrank(subOwner);
    (uint256 requestId1, uint256 preSeed1) =
      s_testCoordinator.computeRequestIdExternal(vrfKeyHash, address(consumer1), subId, 1);
    (uint256 requestId2, uint256 preSeed2) =
      s_testCoordinator.computeRequestIdExternal(vrfKeyHash, address(consumer2), subId, 1);
    assertNotEq(requestId1, requestId2);
    assertNotEq(preSeed1, preSeed2);
    consumer1.requestRandomWords(
      subId,
      MIN_CONFIRMATIONS,
      vrfKeyHash,
      CALLBACK_GAS_LIMIT,
      true, /* nativePayment */
      NUM_WORDS,
      1 /* requestCount */
    );
    consumer2.requestRandomWords(
      subId,
      MIN_CONFIRMATIONS,
      vrfKeyHash,
      CALLBACK_GAS_LIMIT,
      true, /* nativePayment */
      NUM_WORDS,
      1 /* requestCount */
    );
    assertTrue(s_testCoordinator.pendingRequestExists(subId));

    // Move on to the next block.
    // Store the previous block's blockhash, and assert that it is as expected.
    vm.roll(requestBlock + 1);
    s_bhs.store(requestBlock);
    assertEq(hex"1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac", s_bhs.getBlockhash(requestBlock));

    // 3. Fulfill the 1st request above
    // Proof generated via the generate-proof-v2-plus script command.
    // 1st step: Uncomment these 3 console logs to see info about the request and run the test to get output:
    // console.log("requestId: ", requestId);
    // console.log("preSeed: ", preSeed);
    // console.log("sender: ", address(s_testConsumer));
    // 2nd step: Update pre-seed in the command commented out below with new value printed in console logs.
    // 3rd step: export the following environment variables to run the generate-proof-v2-plus script.
    // export ETH_URL=https://ethereum-sepolia-rpc.publicnode.com # or any other RPC provider you prefer
    // export ETH_CHAIN_ID=11155111 # or switch to any other chain
    // export ACCOUNT_KEY=<your test EOA private key>
    // 4th step: run the command and copy the command output in the proof section below.
    /*
      Run from this folder: chainlink/core/scripts/vrfv2plus/testnet
      go run . generate-proof-v2-plus \
      -key-hash 0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528 \
      -pre-seed 78857362017365444144484359594634073685493503942324326290718892836953423263381 \
      -block-hash 0x1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac \
      -block-num 10 \
      -sender 0x44CAfC03154A0708F9DCf988681821f648dA74aF \
      -native-payment true
    */
    VRF.Proof memory proof = VRF.Proof({
      pk: [
        72_488_970_228_380_509_287_422_715_226_575_535_698_893_157_273_063_074_627_791_787_432_852_706_183_111,
        62_070_622_898_698_443_831_883_535_403_436_258_712_770_888_294_397_026_493_185_421_712_108_624_767_191
      ],
      gamma: [
        65_913_937_398_148_449_626_792_563_067_325_648_649_534_055_460_473_988_721_938_103_219_381_973_178_278,
        63_156_327_344_180_203_180_831_822_252_171_874_192_175_272_818_200_597_638_000_091_892_096_122_362_120
      ],
      c: 103_296_526_941_774_692_908_067_234_360_350_834_482_645_116_475_454_593_803_823_148_315_342_533_216_203,
      s: 50_291_245_814_080_656_739_779_812_653_411_869_801_334_231_723_444_391_096_753_849_942_661_931_376_590,
      seed: 78_857_362_017_365_444_144_484_359_594_634_073_685_493_503_942_324_326_290_718_892_836_953_423_263_381,
      uWitness: 0x38500711AdcB471ac1A566c4b915759eb9cBCE2F,
      cGammaWitness: [
        56_476_970_720_509_547_210_740_928_951_846_471_668_018_949_971_632_948_991_136_782_499_758_110_143_588,
        44_326_075_300_781_389_077_656_415_325_167_171_692_706_436_527_877_070_415_603_658_305_817_367_373_598
      ],
      sHashWitness: [
        109_524_696_164_787_283_409_393_383_708_118_913_934_136_014_139_634_321_235_031_691_839_206_768_278_439,
        52_690_039_857_779_635_909_051_684_567_562_068_782_378_693_408_005_554_345_469_129_234_366_171_822_741
      ],
      zInv: 108_537_983_043_800_425_266_290_112_227_943_788_107_669_768_716_438_017_124_275_578_856_644_517_258_573
    });
    VRFTypes.RequestCommitmentV2Plus memory rc = VRFTypes.RequestCommitmentV2Plus({
      blockNum: requestBlock,
      subId: subId,
      callbackGasLimit: CALLBACK_GAS_LIMIT,
      numWords: NUM_WORDS,
      sender: address(consumer1),
      extraArgs: VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: true}))
    });
    s_testCoordinator.fulfillRandomWords(proof, rc, true /* onlyPremium */ );
    assertTrue(s_testCoordinator.pendingRequestExists(subId));

    // 4. Fulfill the 2nd request
    // Proof generated via the generate-proof-v2-plus script command.
    // 1st step: Uncomment these 3 console logs to see info about the request and run the test to get output:
    // console.log("requestId: ", requestId);
    // console.log("preSeed: ", preSeed);
    // console.log("sender: ", address(s_testConsumer));
    // 2nd step: Update pre-seed in the command commented out below with new value printed in console logs.
    // 3rd step: export the following environment variables to run the generate-proof-v2-plus script.
    // export ETH_URL=https://ethereum-sepolia-rpc.publicnode.com # or any other RPC provider you prefer
    // export ETH_CHAIN_ID=11155111 # or switch to any other chain
    // export ACCOUNT_KEY=<your test EOA private key>
    // 4th step: run the command and copy the command output in the proof section below.
    /*
      Run from this folder: chainlink/core/scripts/vrfv2plus/testnet
      go run . generate-proof-v2-plus \
      -key-hash 0x9f2353bde94264dbc3d554a94cceba2d7d2b4fdce4304d3e09a1fea9fbeb1528 \
      -pre-seed 53330100288105770463016865504321558518073051667771993294213115153676065708950 \
      -block-hash 0x1a192fabce13988b84994d4296e6cdc418d55e2f1d7f942188d4040b94fc57ac \
      -block-num 10 \
      -sender 0xf5a165378E120f93784395aDF1E08a437e902865 \
      -native-payment true
    */
    proof = VRF.Proof({
      pk: [
        72_488_970_228_380_509_287_422_715_226_575_535_698_893_157_273_063_074_627_791_787_432_852_706_183_111,
        62_070_622_898_698_443_831_883_535_403_436_258_712_770_888_294_397_026_493_185_421_712_108_624_767_191
      ],
      gamma: [
        7_260_273_098_301_741_284_457_725_182_313_945_178_888_499_328_441_106_869_722_941_415_453_613_782_770,
        91_648_498_042_618_923_465_107_471_165_504_200_585_847_250_228_048_015_102_713_552_756_245_653_299_952
      ],
      c: 64_987_886_290_696_558_870_328_339_791_409_334_400_119_338_012_796_549_091_587_853_494_368_167_422_332,
      s: 69_469_162_696_695_326_295_567_645_789_624_554_797_683_340_898_724_555_794_078_876_350_372_084_267_572,
      seed: 53_330_100_288_105_770_463_016_865_504_321_558_518_073_051_667_771_993_294_213_115_153_676_065_708_950,
      uWitness: 0xa6ce21aD47eC5E90Ac7a2c6152D9710234Afe8ab,
      cGammaWitness: [
        57_318_358_662_553_647_785_891_634_403_735_348_577_492_991_113_152_343_207_139_729_697_842_283_565_417,
        57_942_043_484_796_308_689_103_390_068_712_967_247_519_265_087_617_809_262_260_051_163_954_389_512_396
      ],
      sHashWitness: [
        113_345_999_157_319_332_195_230_171_660_555_736_547_709_417_795_439_282_230_372_737_104_445_523_493_539,
        113_358_219_039_155_973_560_933_190_466_797_830_695_088_313_506_343_976_960_055_230_355_894_888_727_567
      ],
      zInv: 68_349_552_569_605_209_428_774_574_139_615_352_877_146_713_490_794_995_768_725_549_089_572_297_658_255
    });
    rc = VRFTypes.RequestCommitmentV2Plus({
      blockNum: requestBlock,
      subId: subId,
      callbackGasLimit: CALLBACK_GAS_LIMIT,
      numWords: NUM_WORDS,
      sender: address(consumer2),
      extraArgs: VRFV2PlusClient._argsToBytes(VRFV2PlusClient.ExtraArgsV1({nativePayment: true}))
    });
    s_testCoordinator.fulfillRandomWords(proof, rc, true /* onlyPremium */ );
    assertFalse(s_testCoordinator.pendingRequestExists(subId));
  }

  function createAndAddLoadTestWithMetricsConsumer(
    uint256 subId
  ) internal returns (VRFV2PlusLoadTestWithMetrics) {
    VRFV2PlusLoadTestWithMetrics consumer = new VRFV2PlusLoadTestWithMetrics(address(s_testCoordinator));
    s_testCoordinator.addConsumer(subId, address(consumer));
    return consumer;
  }

  function test_RemoveConsumer() public {
    uint256 subId = s_testCoordinator.createSubscription();
    uint256 consumersLength = s_testCoordinator.MAX_CONSUMERS();
    address[] memory consumers = getRandomAddresses(consumersLength);
    for (uint256 i = 0; i < consumersLength; ++i) {
      s_testCoordinator.addConsumer(subId, consumers[i]);
    }

    // test remove consumers from multiple positions to have better gas distribution
    address earlyConsumerAddress = consumers[0];
    s_testCoordinator.removeConsumer(subId, earlyConsumerAddress);
    (,,,, consumers) = s_testCoordinator.getSubscription(subId);
    assertEq(consumers.length, consumersLength - 1);
    assertFalse(addressIsIn(earlyConsumerAddress, consumers));

    consumersLength = consumers.length;
    address middleConsumerAddress = consumers[consumersLength / 2];
    s_testCoordinator.removeConsumer(subId, middleConsumerAddress);
    (,,,, consumers) = s_testCoordinator.getSubscription(subId);
    assertEq(consumers.length, consumersLength - 1);
    assertFalse(addressIsIn(middleConsumerAddress, consumers));

    consumersLength = consumers.length;
    address lateConsumerAddress = consumers[consumersLength - 1];
    s_testCoordinator.removeConsumer(subId, lateConsumerAddress);
    (,,,, consumers) = s_testCoordinator.getSubscription(subId);
    assertEq(consumers.length, consumersLength - 1);
    assertFalse(addressIsIn(lateConsumerAddress, consumers));
  }
}
