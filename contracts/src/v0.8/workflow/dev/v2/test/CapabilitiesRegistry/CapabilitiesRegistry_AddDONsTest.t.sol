// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
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
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    bytes32[] memory newNodes = new bytes32[](2);
    newNodes[0] = P2P_ID;
    newNodes[1] = P2P_ID_TWO;

    CapabilitiesRegistry.CapabilityConfiguration[] memory defaultCapabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    defaultCapabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    string[] memory donFamilies = new string[](0);
    s_DEFAULT_NEW_DON_PARAMS = new CapabilitiesRegistry.NewDONParams[](1);
    s_DEFAULT_NEW_DON_PARAMS[0] = CapabilitiesRegistry.NewDONParams({
      nodes: newNodes,
      capabilityConfigurations: defaultCapabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: donFamilies,
      config: bytes("")
    });

    vm.startPrank(ADMIN);
  }

  function test_RevertWhen_CalledByNonAdmin() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
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

    s_paramsForTwoNodes[1].capabilityIds = s_oneCapabilityArray;
    s_CapabilitiesRegistry.updateNodes(s_paramsForTwoNodes);

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
      capabilityId: s_nonExistentCapabilityId, // This capability does not exist
      config: BASIC_CAPABILITY_CONFIG
    });

    s_DEFAULT_NEW_DON_PARAMS[0].capabilityConfigurations = capabilityConfigs;

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.CapabilityDoesNotExist.selector, s_nonExistentCapabilityId)
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
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    capabilityConfigs[1] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    s_DEFAULT_NEW_DON_PARAMS[0].capabilityConfigurations = capabilityConfigs;

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.DuplicateDONCapability.selector, 1, s_basicCapabilityId)
    );
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
  }

  function test_RevertWhen_DeprecatedCapabilityAdded() public {
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_basicCapabilityId;
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.CapabilityIsDeprecated.selector, s_basicCapabilityId));
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
    newDONs[1] = CapabilitiesRegistry.NewDONParams({
      nodes: s_nodeIds,
      capabilityConfigurations: s_capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_TWO,
      donFamilies: new string[](0),
      config: bytes("")
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.NodePartOfWorkflowDON.selector, 2, P2P_ID));
    s_CapabilitiesRegistry.addDONs(newDONs);
  }

  function test_RevertWhen_DONNameCannotBeEmpty() public {
    s_DEFAULT_NEW_DON_PARAMS[0].name = "";
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONNameCannotBeEmpty.selector, 1));
    s_CapabilitiesRegistry.addDONs(s_DEFAULT_NEW_DON_PARAMS);
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
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigsForTwoCapabilities =
      new CapabilitiesRegistry.CapabilityConfiguration[](2);
    capabilityConfigsForTwoCapabilities[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    capabilityConfigsForTwoCapabilities[1] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_capabilityWithConfigurationContractId,
      config: CONFIG_CAPABILITY_CONFIG
    });

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.ConfigSet(DON_ID, 1);
    vm.expectCall(
      address(s_capabilityConfigurationContract),
      abi.encodeWithSelector(
        ICapabilityConfiguration.beforeCapabilityConfigSet.selector, s_nodeIds, CONFIG_CAPABILITY_CONFIG, 1, DON_ID
      ),
      1
    );
    string[] memory donFamilies = new string[](2);
    donFamilies[0] = "basic-family";
    donFamilies[1] = "test-family";
    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: s_nodeIds,
      capabilityConfigurations: capabilityConfigsForTwoCapabilities,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: donFamilies,
      config: bytes("abc")
    });

    s_CapabilitiesRegistry.addDONs(newDONs);

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.id, DON_ID);
    assertEq(donInfo.configCount, 1);
    assertEq(donInfo.isPublic, true);
    assertEq(
      donInfo.capabilityConfigurations.length,
      capabilityConfigsForTwoCapabilities.length,
      "Capability configs length mismatch"
    );
    assertEq(donInfo.capabilityConfigurations[0].capabilityId, s_basicCapabilityId);
    assertEq(donInfo.capabilityConfigurations[1].capabilityId, s_capabilityWithConfigurationContractId);
    assertEq(donInfo.name, TEST_DON_NAME_ONE);
    assertEq(donInfo.config, bytes("abc"));

    (bytes memory CapabilitiesRegistryDONConfig, bytes memory capabilityConfigContractConfig) =
      s_CapabilitiesRegistry.getCapabilityConfigs(DON_ID, s_basicCapabilityId);
    assertEq(CapabilitiesRegistryDONConfig, BASIC_CAPABILITY_CONFIG);
    assertEq(capabilityConfigContractConfig, bytes(""));

    (bytes memory CapabilitiesRegistryDONConfigTwo, bytes memory capabilityConfigContractConfigTwo) =
      s_CapabilitiesRegistry.getCapabilityConfigs(DON_ID, s_capabilityWithConfigurationContractId);
    assertEq(CapabilitiesRegistryDONConfigTwo, CONFIG_CAPABILITY_CONFIG);
    assertEq(capabilityConfigContractConfigTwo, CONFIG_CAPABILITY_CONFIG);

    assertEq(donInfo.nodeP2PIds.length, s_nodeIds.length);
    assertEq(donInfo.nodeP2PIds[0], P2P_ID);
    assertEq(donInfo.nodeP2PIds[1], P2P_ID_TWO);

    assertEq(donInfo.donFamilies.length, 2);
    assertEq(donInfo.donFamilies[0], "basic-family");
    assertEq(donInfo.donFamilies[1], "test-family");
  }

  function test_AddDONs_OneNodeDON() public {
    s_CapabilitiesRegistry = new CapabilitiesRegistry(CapabilitiesRegistry.ConstructorParams({canAddOneNodeDONs: true}));

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    bytes32[] memory nodes = new bytes32[](1);
    nodes[0] = P2P_ID;

    CapabilitiesRegistry.NewDONParams[] memory oneNodeDONs = new CapabilitiesRegistry.NewDONParams[](1);
    oneNodeDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodes,
      capabilityConfigurations: s_capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: 0,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: bytes("abc")
    });

    s_CapabilitiesRegistry.addDONs(oneNodeDONs);
  }

  function test_AddDONs_EmptyConfigDON() public {
    s_CapabilitiesRegistry =
      new CapabilitiesRegistry(CapabilitiesRegistry.ConstructorParams({canAddOneNodeDONs: false}));

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    CapabilitiesRegistry.CapabilityConfiguration[] memory emptyCapabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](0);

    string[] memory donFamilies = new string[](1);
    donFamilies[0] = "basic-family";

    CapabilitiesRegistry.NewDONParams[] memory emptyConfigDONs = new CapabilitiesRegistry.NewDONParams[](1);
    emptyConfigDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: s_nodeIds,
      capabilityConfigurations: emptyCapabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: donFamilies,
      config: bytes("abc")
    });

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.ConfigSet(DON_ID, 1);
    vm.expectCall(
      address(s_capabilityConfigurationContract),
      abi.encodeWithSelector(
        ICapabilityConfiguration.beforeCapabilityConfigSet.selector, s_nodeIds, emptyConfigDONs, 1, DON_ID
      ),
      0
    );
    s_CapabilitiesRegistry.addDONs(emptyConfigDONs);

    // DON ID 1 has been used, next ID should be 2
    uint32 donID = s_CapabilitiesRegistry.getNextDONId();
    assertEq(donID, DON_ID + 1);

    CapabilitiesRegistry.DONInfo[] memory dons = s_CapabilitiesRegistry.getDONs();
    assertEq(dons.length, 1);
    assertEq(dons[0].id, 1);
    assertEq(dons[0].name, TEST_DON_NAME_ONE);
    assertEq(dons[0].config, bytes("abc"));
    assertEq(dons[0].configCount, 1);
    assertEq(dons[0].capabilityConfigurations.length, emptyCapabilityConfigs.length);
    assertEq(dons[0].donFamilies.length, 1);
    assertEq(dons[0].donFamilies[0], "basic-family");
  }
}
