// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";

contract CapabilitiesRegistry_GetHistoricalDONInfoTest is BaseTest {
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
    s_CapabilitiesRegistry.addDON(nodeIds, s_capabilityConfigs, true, true, F_VALUE, s_testDONParams);
    // Remove the DON name to test the historical DON info
    s_CapabilitiesRegistry.updateDON(DON_ID, nodeIds, s_capabilityConfigs, false, F_VALUE, s_emptyOptionalDONParams);
  }

  function test_RevertWhen_DONDoesNotExist() public {
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONDoesNotExist.selector, 999));
    s_CapabilitiesRegistry.getHistoricalDONInfo(999, 1);
  }

  function test_RevertWhen_DONConfigDoesNotExist() public {
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONConfigDoesNotExist.selector, DON_ID, 2, 10));
    s_CapabilitiesRegistry.getHistoricalDONInfo(DON_ID, 10);
  }

  function test_CorrectlyFetchesHistoricalDONInfo() public view {
    CapabilitiesRegistry.DONInfo memory don = s_CapabilitiesRegistry.getHistoricalDONInfo(DON_ID, 1);
    assertEq(don.id, DON_ID);
    assertEq(don.configCount, 1);
    assertEq(don.isPublic, true);
    assertEq(don.acceptsWorkflows, true);
    assertEq(don.f, 1);
    assertEq(don.capabilityConfigurations.length, s_capabilityConfigs.length);
    assertEq(don.capabilityConfigurations[0].capabilityId, s_basicHashedCapabilityId);
    assertEq(don.name, "test-name");
    assertEq(don.config, bytes("abc"));

    don = s_CapabilitiesRegistry.getHistoricalDONInfo(DON_ID, 2);
    assertEq(don.id, DON_ID);
    assertEq(don.configCount, 2);
    assertEq(don.isPublic, false);
    assertEq(don.acceptsWorkflows, true);
    assertEq(don.f, 1);
    assertEq(don.capabilityConfigurations.length, s_capabilityConfigs.length);
    assertEq(don.capabilityConfigurations[0].capabilityId, s_basicHashedCapabilityId);
    assertEq(don.name, "");
    assertEq(don.config, bytes(""));
  }
}
