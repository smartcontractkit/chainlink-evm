// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {Constants} from "./Constants.t.sol";
import {CapabilityConfigurationContract} from "./mocks/CapabilityConfigurationContract.sol";
import {Test} from "forge-std/Test.sol";

contract BaseTest is Test, Constants {
  CapabilitiesRegistry internal s_CapabilitiesRegistry;
  CapabilityConfigurationContract internal s_capabilityConfigurationContract;
  CapabilitiesRegistry.Capability internal s_basicCapability;
  CapabilitiesRegistry.Capability internal s_capabilityWithConfigurationContract;
  string[] internal s_oneCapabilityArray;
  string[] internal s_twoCapabilitiesArray;
  string internal s_basicCapabilityId;
  string internal s_capabilityWithConfigurationContractId;
  string internal s_nonExistentCapabilityId;
  CapabilitiesRegistry.Capability[] internal s_capabilities;
  CapabilitiesRegistry.CapabilityConfiguration[] internal s_capabilityConfigs;
  CapabilitiesRegistry.NodeParams[] internal s_paramsForTwoNodes;
  CapabilitiesRegistry.NewDONParams[] internal s_paramsForTwoDONs;
  bytes32[] internal s_nodeIds;

  function setUp() public virtual {
    vm.startPrank(ADMIN);
    s_CapabilitiesRegistry =
      new CapabilitiesRegistry(CapabilitiesRegistry.ConstructorParams({canAddOneNodeDONs: false}));
    s_capabilityConfigurationContract = new CapabilityConfigurationContract();

    s_basicCapability = CapabilitiesRegistry.Capability({
      capabilityId: "data-streams-reports@1.0.0",
      configurationContract: address(0),
      metadata: TEST_CAPABILITY_METADATA
    });
    s_capabilityWithConfigurationContract = CapabilitiesRegistry.Capability({
      capabilityId: "read-ethereum-mainnet-gas-price@1.0.2",
      configurationContract: address(s_capabilityConfigurationContract),
      metadata: bytes("")
    });

    s_basicCapabilityId = s_basicCapability.capabilityId;
    s_capabilityWithConfigurationContractId = s_capabilityWithConfigurationContract.capabilityId;
    s_nonExistentCapabilityId = "non-existent-capability@1.0.0";

    s_oneCapabilityArray = new string[](1);
    s_oneCapabilityArray[0] = s_basicCapabilityId;

    s_twoCapabilitiesArray = new string[](2);
    s_twoCapabilitiesArray[0] = s_basicCapabilityId;
    s_twoCapabilitiesArray[1] = s_capabilityWithConfigurationContractId;

    s_capabilities = new CapabilitiesRegistry.Capability[](2);
    s_capabilities[0] = s_basicCapability;
    s_capabilities[1] = s_capabilityWithConfigurationContract;

    s_capabilityConfigs = new CapabilitiesRegistry.CapabilityConfiguration[](1);
    s_capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    s_paramsForTwoNodes = new CapabilitiesRegistry.NodeParams[](2);
    s_paramsForTwoNodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: s_twoCapabilitiesArray
    });

    s_paramsForTwoNodes[1] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: s_twoCapabilitiesArray
    });

    s_nodeIds = new bytes32[](2);
    s_nodeIds[0] = s_paramsForTwoNodes[0].p2pId;
    s_nodeIds[1] = s_paramsForTwoNodes[1].p2pId;

    string[] memory donFamilies = new string[](1);
    donFamilies[0] = TEST_DON_FAMILY_ONE;

    s_paramsForTwoDONs = new CapabilitiesRegistry.NewDONParams[](2);
    s_paramsForTwoDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: s_nodeIds,
      capabilityConfigurations: s_capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: 1,
      name: TEST_DON_NAME_ONE,
      donFamilies: donFamilies,
      config: bytes("")
    });

    s_paramsForTwoDONs[1] = CapabilitiesRegistry.NewDONParams({
      nodes: s_nodeIds,
      capabilityConfigurations: s_capabilityConfigs,
      isPublic: false,
      acceptsWorkflows: false,
      f: 1,
      name: TEST_DON_NAME_TWO,
      donFamilies: donFamilies,
      config: TEST_DON_CONFIG
    });
  }

  function _getNodeOperators() internal pure returns (CapabilitiesRegistry.NodeOperator[] memory) {
    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](3);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: NODE_OPERATOR_ONE_ADMIN, name: NODE_OPERATOR_ONE_NAME});
    nodeOperators[1] = CapabilitiesRegistry.NodeOperator({admin: NODE_OPERATOR_TWO_ADMIN, name: NODE_OPERATOR_TWO_NAME});
    nodeOperators[2] = CapabilitiesRegistry.NodeOperator({admin: NODE_OPERATOR_THREE, name: NODE_OPERATOR_THREE_NAME});
    return nodeOperators;
  }
}
