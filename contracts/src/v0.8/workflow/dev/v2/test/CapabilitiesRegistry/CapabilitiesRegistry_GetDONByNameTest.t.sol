// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";

contract CapabilitiesRegistry_GetDONByNameTest is BaseTest {
  CapabilitiesRegistry.CapabilityConfiguration[] private s_capabilityConfigs;

  function setUp() public override {
    BaseTest.setUp();

    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](2);
    bytes32[] memory capabilityIds = new bytes32[](2);
    capabilityIds[0] = s_basicHashedCapabilityId;
    capabilityIds[1] = s_capabilityWithConfigurationContractId;

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
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      hashedCapabilityIds: nodeTwoCapabilityIds
    });

    changePrank(NODE_OPERATOR_ONE_ADMIN);
    s_CapabilitiesRegistry.addNodes(nodes);

    s_capabilityConfigs.push(
      CapabilitiesRegistry.CapabilityConfiguration({
        capabilityId: s_basicHashedCapabilityId,
        config: BASIC_CAPABILITY_CONFIG
      })
    );

    bytes32[] memory nodeIds = new bytes32[](2);
    nodeIds[0] = P2P_ID;
    nodeIds[1] = P2P_ID_TWO;

    changePrank(ADMIN);
    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodeIds,
      capabilityConfigurations: s_capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: 1,
      name: s_testDONParams.name,
      config: s_testDONParams.config
    });
    s_CapabilitiesRegistry.addDONs(newDONs);
  }

  function test_RevertWhen_DONDoesNotExist() public {
    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.DONWithNameDoesNotExist.selector, "non-existent-don-name")
    );
    s_CapabilitiesRegistry.getDONByName("non-existent-don-name");
  }

  function test_CorrectlyFetchesDONByName() public view {
    CapabilitiesRegistry.DONInfo memory don = s_CapabilitiesRegistry.getDONByName(s_testDONParams.name);
    assertEq(don.id, DON_ID);
    assertEq(don.configCount, 1);
    assertEq(don.isPublic, true);
    assertEq(don.acceptsWorkflows, true);
    assertEq(don.f, 1);
    assertEq(don.capabilityConfigurations.length, s_capabilityConfigs.length);
    assertEq(don.capabilityConfigurations[0].capabilityId, s_basicHashedCapabilityId);
    assertEq(don.name, s_testDONParams.name);
    assertEq(don.config, s_testDONParams.config);
  }
}
