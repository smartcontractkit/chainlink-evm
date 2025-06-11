// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

import {MaliciousConfigurationContract} from "./mocks/MaliciousConfigurationContract.sol";

contract CapabilitiesRegistry_AddDONsTest_WhenMaliciousCapabilityConfigurationConfigured is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    MaliciousConfigurationContract maliciousConfigurationContract =
      new MaliciousConfigurationContract(s_capabilityWithConfigurationContractId);

    address maliciousConfigContractAddr = address(maliciousConfigurationContract);
    s_basicCapability.configurationContract = maliciousConfigContractAddr;
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    bytes memory config = maliciousConfigurationContract.getCapabilityConfiguration(DON_ID);
    assertEq(config, bytes(""));

    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = _getNodeOperators();
    nodeOperators[0].admin = maliciousConfigContractAddr;
    nodeOperators[1].admin = maliciousConfigContractAddr;
    nodeOperators[2].admin = maliciousConfigContractAddr;

    s_CapabilitiesRegistry.addNodeOperators(nodeOperators);
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](3);
    bytes32[] memory capabilityIds = new bytes32[](1);
    capabilityIds[0] = s_basicHashedCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      hashedCapabilityIds: capabilityIds
    });

    bytes32[] memory nodeTwoCapabilityIds = new bytes32[](1);
    nodeTwoCapabilityIds[0] = s_basicHashedCapabilityId;

    nodes[1] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_TWO_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      hashedCapabilityIds: nodeTwoCapabilityIds
    });

    nodes[2] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_THREE_ID,
      p2pId: P2P_ID_THREE,
      signer: NODE_OPERATOR_THREE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_THREE,
      hashedCapabilityIds: capabilityIds
    });

    s_CapabilitiesRegistry.addNodes(nodes);

    changePrank(ADMIN);
  }

  function test_RevertWhen_MaliciousCapabilitiesConfigContractTriesToRemoveCapabilitiesFromDONNodes() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_THREE;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_basicHashedCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });

    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodes,
      capabilityConfigurations: capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: "",
      config: bytes("")
    });

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.CapabilityRequiredByDON.selector, s_basicHashedCapabilityId, DON_ID)
    );
    s_CapabilitiesRegistry.addDONs(newDONs);
  }
}
