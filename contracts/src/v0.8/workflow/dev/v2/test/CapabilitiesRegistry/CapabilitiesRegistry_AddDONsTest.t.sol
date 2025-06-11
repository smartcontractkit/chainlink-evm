// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {ICapabilityConfiguration} from "../../interfaces/ICapabilityConfiguration.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_AddDONsTest is BaseTest {
  CapabilitiesRegistry.NewDONParams[] private s_DEFAULT_NEW_DON_PARAMS;

  function setUp() public override {
    BaseTest.setUp();
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](3);
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

    bytes32[] memory newNodes = new bytes32[](2);
    newNodes[0] = P2P_ID;
    newNodes[1] = P2P_ID_TWO;

    CapabilitiesRegistry.CapabilityConfiguration[] memory defaultCapabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    defaultCapabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_basicHashedCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });

    s_DEFAULT_NEW_DON_PARAMS = new CapabilitiesRegistry.NewDONParams[](1);
    s_DEFAULT_NEW_DON_PARAMS[0] = CapabilitiesRegistry.NewDONParams({
      nodes: newNodes,
      capabilityConfigurations: defaultCapabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: "",
      config: bytes("")
    });

    vm.startPrank(ADMIN);
  }

  function test_RevertWhen_CalledByNonAdmin() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_NodeDoesNotSupportCapability() public {
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_capabilityWithConfigurationContractId, // This capability is not supported by the nodes
      config: CONFIG_CAPABILITY_CONFIG
    });

    s_DEFAULT_NEW_DON_PARAMS[0].capabilityConfigurations = capabilityConfigs;

    vm.expectRevert(
      abi.encodeWithSelector(
        CapabilitiesRegistry.NodeDoesNotSupportCapability.selector, P2P_ID_TWO, s_capabilityWithConfigurationContractId
      )
    );
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_CapabilityDoesNotExist() public {
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_nonExistentHashedCapabilityId, // This capability does not exist
      config: BASIC_CAPABILITY_CONFIG
    });

    s_DEFAULT_NEW_DON_PARAMS[0].capabilityConfigurations = capabilityConfigs;

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.CapabilityDoesNotExist.selector, s_nonExistentHashedCapabilityId)
    );
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_FaultToleranceIsZero() public {
    s_DEFAULT_NEW_DON_PARAMS[0].f = 0;

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidFaultTolerance.selector, 0, 2));
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_DuplicateCapabilityAdded() public {
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](2);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_basicHashedCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });
    capabilityConfigs[1] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_basicHashedCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });

    s_DEFAULT_NEW_DON_PARAMS[0].capabilityConfigurations = capabilityConfigs;

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.DuplicateDONCapability.selector, 1, s_basicHashedCapabilityId)
    );
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_DeprecatedCapabilityAdded() public {
    bytes32 capabilityId = s_basicHashedCapabilityId;
    bytes32[] memory deprecatedCapabilities = new bytes32[](1);
    deprecatedCapabilities[0] = capabilityId;
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.CapabilityIsDeprecated.selector, capabilityId));
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_DuplicateNodeAdded() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID;

    s_DEFAULT_NEW_DON_PARAMS[0].nodes = nodes;

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DuplicateDONNode.selector, 1, P2P_ID));
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_NodeAlreadyBelongsToWorkflowDON() public {
    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](2);
    newDONs[0] = s_DEFAULT_NEW_DON_PARAMS[0];
    newDONs[0].acceptsWorkflows = true; // This DON accepts workflows
    newDONs[1] = newDONs[0]; // Make a copy of the first DON

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.NodePartOfWorkflowDON.selector, 2, P2P_ID));
    s_CapabilitiesRegistry.addDONs(newDONs);
  }

  function test_RevertWhen_DONNameAlreadyTaken() public {
    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](2);
    newDONs[0] = s_DEFAULT_NEW_DON_PARAMS[0];
    newDONs[0].name = "test";
    newDONs[1] = newDONs[0]; // Make a copy of the first DON
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONNameAlreadyTaken.selector, "test"));
    s_CapabilitiesRegistry.addDONs(newDONs);
  }

  function test_AddDONs() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_THREE;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](2);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_basicHashedCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });
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
      name: "test-name",
      config: bytes("abc")
    });

    s_CapabilitiesRegistry.addDONs(newDONs);

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.id, DON_ID);
    assertEq(donInfo.configCount, 1);
    assertEq(donInfo.isPublic, true);
    assertEq(donInfo.capabilityConfigurations.length, capabilityConfigs.length);
    assertEq(donInfo.capabilityConfigurations[0].capabilityId, s_basicHashedCapabilityId);
    assertEq(donInfo.name, "test-name");
    assertEq(donInfo.config, bytes("abc"));

    (bytes memory CapabilitiesRegistryDONConfig, bytes memory capabilityConfigContractConfig) =
      s_CapabilitiesRegistry.getCapabilityConfigs(DON_ID, s_basicHashedCapabilityId);
    assertEq(CapabilitiesRegistryDONConfig, BASIC_CAPABILITY_CONFIG);
    assertEq(capabilityConfigContractConfig, bytes(""));

    (bytes memory CapabilitiesRegistryDONConfigTwo, bytes memory capabilityConfigContractConfigTwo) =
      s_CapabilitiesRegistry.getCapabilityConfigs(DON_ID, s_capabilityWithConfigurationContractId);
    assertEq(CapabilitiesRegistryDONConfigTwo, CONFIG_CAPABILITY_CONFIG);
    assertEq(capabilityConfigContractConfigTwo, CONFIG_CAPABILITY_CONFIG);

    assertEq(donInfo.nodeP2PIds.length, nodes.length);
    assertEq(donInfo.nodeP2PIds[0], P2P_ID);
    assertEq(donInfo.nodeP2PIds[1], P2P_ID_THREE);
  }
}
