// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";

contract CapabilitiesRegistry_RemoveDONsByNameTest is BaseTest {
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

    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    s_CapabilitiesRegistry.addNodes(nodes);

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_basicHashedCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });

    bytes32[] memory nodeIds = new bytes32[](2);
    nodeIds[0] = P2P_ID;
    nodeIds[1] = P2P_ID_TWO;

    vm.stopPrank();
    vm.startPrank(ADMIN);
    s_CapabilitiesRegistry.addDON(nodeIds, capabilityConfigs, true, true, 1, s_testDONParams);
  }

  function test_RevertWhen_CalledByNonAdmin() public {
    string[] memory donNames = new string[](1);
    donNames[0] = "test-name";
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_CapabilitiesRegistry.removeDONsByName(donNames);
  }

  function test_RemovesDONsByName() public {
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.ConfigSet(DON_ID, 0);

    string[] memory donNames = new string[](1);
    donNames[0] = "test-name";
    s_CapabilitiesRegistry.removeDONsByName(donNames);

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONDoesNotExist.selector, DON_ID));
    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);

    (bytes memory CapabilitiesRegistryDONConfig, bytes memory capabilityConfigContractConfig) =
      s_CapabilitiesRegistry.getCapabilityConfigs(DON_ID, s_basicHashedCapabilityId);

    assertEq(CapabilitiesRegistryDONConfig, bytes(""));
    assertEq(capabilityConfigContractConfig, bytes(""));
    assertEq(donInfo.nodeP2PIds.length, 0);

    assertEq(s_CapabilitiesRegistry.isDONNameTaken("test-name"), false);
  }
}
