// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {ICapabilityConfiguration} from "../../interfaces/ICapabilityConfiguration.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_UpdateDONTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);
    s_CapabilitiesRegistry.addDONs(s_paramsForTwoDONs);
  }

  function test_RevertWhen_CalledByNonAdmin() public {
    changePrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    bytes32[] memory nodes = new bytes32[](1);
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);

    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );
  }

  function test_RevertWhen_NodeDoesNotSupportCapability() public {
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_capabilityWithConfigurationContractId,
      config: CONFIG_CAPABILITY_CONFIG
    });

    s_paramsForTwoNodes[1].capabilityIds = s_oneCapabilityArray;
    s_CapabilitiesRegistry.updateNodes(s_paramsForTwoNodes);
    vm.expectRevert(
      abi.encodeWithSelector(
        CapabilitiesRegistry.NodeDoesNotSupportCapability.selector, P2P_ID_TWO, s_capabilityWithConfigurationContractId
      )
    );
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: s_nodeIds,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );
  }

  function test_RevertWhen_DONDoesNotExist() public {
    uint32 nonExistentDONId = 10;
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONDoesNotExist.selector, nonExistentDONId));
    s_CapabilitiesRegistry.updateDON(
      nonExistentDONId,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );
  }

  function test_RevertWhen_CapabilityDoesNotExist() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_nonExistentCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });
    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.CapabilityDoesNotExist.selector, s_nonExistentCapabilityId)
    );
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );
  }

  function test_RevertWhen_DuplicateCapabilityAdded() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](2);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    capabilityConfigs[1] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.DuplicateDONCapability.selector, 1, s_basicCapabilityId)
    );
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );
  }

  function test_RevertWhen_DeprecatedCapabilityAdded() public {
    string[] memory deprecatedCapabilities = new string[](1);
    deprecatedCapabilities[0] = s_basicCapabilityId;
    s_CapabilitiesRegistry.deprecateCapabilities(deprecatedCapabilities);

    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.CapabilityIsDeprecated.selector, s_basicCapabilityId));
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );
  }

  function test_RevertWhen_DuplicateNodeAdded() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DuplicateDONNode.selector, 1, P2P_ID));
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );
  }

  function test_UpdatesDON() public {
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

    CapabilitiesRegistry.DONInfo memory oldDONInfo = s_CapabilitiesRegistry.getDON(DON_ID);

    bool expectedDONIsPublic = false;
    uint32 expectedConfigCount = oldDONInfo.configCount + 1;

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.ConfigSet(DON_ID, expectedConfigCount);
    vm.expectCall(
      address(s_capabilityConfigurationContract),
      abi.encodeWithSelector(
        ICapabilityConfiguration.beforeCapabilityConfigSet.selector,
        nodes,
        CONFIG_CAPABILITY_CONFIG,
        expectedConfigCount,
        DON_ID
      ),
      1
    );
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: expectedDONIsPublic,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.id, DON_ID);
    assertEq(donInfo.configCount, expectedConfigCount);
    assertEq(donInfo.isPublic, false);
    assertEq(donInfo.capabilityConfigurations.length, capabilityConfigs.length);
    assertEq(donInfo.capabilityConfigurations[0].capabilityId, s_basicCapabilityId);

    (bytes memory CapabilitiesRegistryDONConfig, bytes memory capabilityConfigContractConfig) =
      s_CapabilitiesRegistry.getCapabilityConfigs(DON_ID, s_basicCapabilityId);
    assertEq(CapabilitiesRegistryDONConfig, BASIC_CAPABILITY_CONFIG);
    assertEq(capabilityConfigContractConfig, bytes(""));

    assertEq(donInfo.nodeP2PIds.length, nodes.length);
    assertEq(donInfo.nodeP2PIds[0], P2P_ID);
    assertEq(donInfo.nodeP2PIds[1], P2P_ID_THREE);
  }
}
