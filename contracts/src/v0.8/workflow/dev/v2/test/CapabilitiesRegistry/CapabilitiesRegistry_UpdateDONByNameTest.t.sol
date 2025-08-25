// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {ICapabilityConfiguration} from "../../interfaces/ICapabilityConfiguration.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_UpdateDONByNameTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    bytes32[] memory donNodes = new bytes32[](2);
    donNodes[0] = P2P_ID;
    donNodes[1] = P2P_ID_TWO;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: donNodes,
      capabilityConfigurations: capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: TEST_DON_CONFIG
    });
    s_CapabilitiesRegistry.addDONs(newDONs);
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

  function test_RevertWhen_DONDoesNotExist() public {
    string memory nonExistentDONName = "non-existent-don-name";
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONWithNameDoesNotExist.selector, nonExistentDONName));
    s_CapabilitiesRegistry.updateDONByName(
      nonExistentDONName,
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

  function test_UpdatesDONByName() public {
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

    CapabilitiesRegistry.DONInfo memory oldDONInfo = s_CapabilitiesRegistry.getDONByName(TEST_DON_NAME_ONE);

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
    s_CapabilitiesRegistry.updateDONByName(
      TEST_DON_NAME_ONE,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: nodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: expectedDONIsPublic,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: TEST_DON_CONFIG
      })
    );

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDONByName(TEST_DON_NAME_ONE);
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
