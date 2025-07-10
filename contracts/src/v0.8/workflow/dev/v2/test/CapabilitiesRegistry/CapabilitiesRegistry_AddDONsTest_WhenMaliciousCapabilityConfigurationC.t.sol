// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

import {MaliciousConfigurationContract} from "./mocks/MaliciousConfigurationContract.sol";

contract CapabilitiesRegistry_AddDONsTest_WhenMaliciousCapabilityConfigurationConfigured is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    MaliciousConfigurationContract maliciousConfigurationContract =
      new MaliciousConfigurationContract(s_capabilityWithConfigurationContractId);

    address maliciousConfigContractAddr = address(maliciousConfigurationContract);
    s_basicCapability.configurationContract = maliciousConfigContractAddr;

    bytes memory config = maliciousConfigurationContract.getCapabilityConfiguration(DON_ID);
    assertEq(config, bytes(""));

    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = _getNodeOperators();
    nodeOperators[0].admin = maliciousConfigContractAddr;
    nodeOperators[1].admin = maliciousConfigContractAddr;
    nodeOperators[2].admin = maliciousConfigContractAddr;

    // Set the configuration contract to the malicious contract
    s_capabilities[0].configurationContract = maliciousConfigContractAddr;

    s_CapabilitiesRegistry.addNodeOperators(nodeOperators);
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);

    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    vm.startPrank(ADMIN);
  }

  function test_RevertWhen_MaliciousCapabilitiesConfigContractTriesToRemoveCapabilitiesFromDONNodes() public {
    CapabilitiesRegistry.NodeParams[] memory nodeParams = new CapabilitiesRegistry.NodeParams[](1);
    nodeParams[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_THREE_ID,
      p2pId: P2P_ID_THREE,
      signer: NODE_OPERATOR_THREE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_THREE,
      csaKey: TEST_CSA_KEY_THREE,
      capabilityIds: s_twoCapabilitiesArray
    });
    s_CapabilitiesRegistry.addNodes(nodeParams);

    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_THREE;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodes,
      capabilityConfigurations: capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: bytes("")
    });

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.CapabilityRequiredByDON.selector, s_basicCapabilityId, DON_ID)
    );
    s_CapabilitiesRegistry.addDONs(newDONs);
  }
}
