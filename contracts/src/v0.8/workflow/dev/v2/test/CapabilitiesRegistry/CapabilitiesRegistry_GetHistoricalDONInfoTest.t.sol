// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";

contract CapabilitiesRegistry_GetHistoricalDONInfoTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](2);
    string[] memory capabilityIds = new string[](2);
    capabilityIds[0] = s_basicCapabilityId;
    capabilityIds[1] = s_capabilityWithConfigurationContractId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    string[] memory nodeTwoCapabilityIds = new string[](1);
    nodeTwoCapabilityIds[0] = s_basicCapabilityId;

    nodes[1] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: nodeTwoCapabilityIds
    });

    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    s_CapabilitiesRegistry.addNodes(nodes);

    bytes32[] memory nodeIds = new bytes32[](2);
    nodeIds[0] = P2P_ID;
    nodeIds[1] = P2P_ID_TWO;

    vm.stopPrank();
    vm.startPrank(ADMIN);
    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodeIds,
      capabilityConfigurations: s_capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: TEST_DON_CONFIG
    });
    s_CapabilitiesRegistry.addDONs(newDONs);
    // Remove the DON name to test the historical DON info
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodeIds,
        capabilityConfigurations: s_capabilityConfigs,
        isPublic: false,
        f: F_VALUE,
        name: TEST_DON_NAME_TWO,
        config: bytes("")
      })
    );
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
    assertEq(don.capabilityConfigurations[0].capabilityId, s_basicCapabilityId);
    assertEq(don.name, TEST_DON_NAME_ONE, "DON name mismatch");
    assertEq(don.config, TEST_DON_CONFIG);

    don = s_CapabilitiesRegistry.getHistoricalDONInfo(DON_ID, 2);
    assertEq(don.id, DON_ID);
    assertEq(don.configCount, 2);
    assertEq(don.isPublic, false);
    assertEq(don.acceptsWorkflows, true);
    assertEq(don.f, 1);
    assertEq(don.capabilityConfigurations.length, s_capabilityConfigs.length);
    assertEq(don.capabilityConfigurations[0].capabilityId, s_basicCapabilityId);
    assertEq(don.name, TEST_DON_NAME_TWO, "DON name mismatch");
    assertEq(don.config, bytes(""));
  }
}
