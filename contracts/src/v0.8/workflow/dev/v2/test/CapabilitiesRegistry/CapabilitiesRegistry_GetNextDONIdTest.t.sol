// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {ICapabilityConfiguration} from "../../interfaces/ICapabilityConfiguration.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_GetNextDONIdTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    changePrank(ADMIN);
  }

  function test_CorrectlyFetchesNextDONId() public {
    uint32 nextDONId = s_CapabilitiesRegistry.getNextDONId();
    assertEq(nextDONId, 1); // Expecting the first DON ID since no DONs have been added yet

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
      new CapabilitiesRegistry.CapabilityConfiguration[](2);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    capabilityConfigs[1] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_capabilityWithConfigurationContractId,
      config: CONFIG_CAPABILITY_CONFIG
    });

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.ConfigSet(DON_ID, 1);
    vm.expectCall(
      address(s_capabilityConfigurationContract),
      abi.encodeWithSelector(
        ICapabilityConfiguration.beforeCapabilityConfigSet.selector, nodes, CONFIG_CAPABILITY_CONFIG, 1, DON_ID
      ),
      1
    );
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
    s_CapabilitiesRegistry.addDONs(newDONs);

    nextDONId = s_CapabilitiesRegistry.getNextDONId();
    assertEq(nextDONId, 2); // After adding one DON, the next DON ID should be 2
  }
}
