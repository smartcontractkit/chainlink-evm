// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {INodeInfoProvider} from "../../interfaces/INodeInfoProvider.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_RemoveNodesTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](3);
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: s_twoCapabilitiesArray
    });

    nodes[1] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_TWO_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: s_twoCapabilitiesArray
    });

    nodes[2] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_THREE_ID,
      p2pId: P2P_ID_THREE,
      signer: NODE_OPERATOR_THREE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_THREE,
      csaKey: TEST_CSA_KEY_THREE,
      capabilityIds: s_twoCapabilitiesArray
    });

    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_CalledByNonNodeOperatorAdminAndNonOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);
    bytes32[] memory nodes = new bytes32[](1);
    nodes[0] = P2P_ID;

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.AccessForbidden.selector, STRANGER));
    s_CapabilitiesRegistry.removeNodes(nodes);
  }

  function test_RevertWhen_NodeDoesNotExist() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    bytes32[] memory nodes = new bytes32[](1);
    nodes[0] = INVALID_P2P_ID;

    vm.expectRevert(abi.encodeWithSelector(INodeInfoProvider.NodeDoesNotExist.selector, INVALID_P2P_ID));
    s_CapabilitiesRegistry.removeNodes(nodes);
  }

  function test_RevertWhen_P2PIDEmpty() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    bytes32[] memory nodes = new bytes32[](1);
    nodes[0] = bytes32("");

    vm.expectRevert(abi.encodeWithSelector(INodeInfoProvider.NodeDoesNotExist.selector, bytes32("")));
    s_CapabilitiesRegistry.removeNodes(nodes);
  }

  function test_RevertWhen_NodePartOfCapabilitiesDON() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodes,
      capabilityConfigurations: capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: false,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: bytes("")
    });
    s_CapabilitiesRegistry.addDONs(newDONs);

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.NodePartOfCapabilitiesDON.selector, 1, P2P_ID));
    s_CapabilitiesRegistry.removeNodes(nodes);
  }

  function test_CanRemoveWhenDONDeleted() public {
    bytes32[] memory nodes = new bytes32[](2);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    // Add DON
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

    // Try remove nodes
    bytes32[] memory removedNodes = new bytes32[](1);
    removedNodes[0] = P2P_ID;
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.NodePartOfWorkflowDON.selector, 1, P2P_ID));
    s_CapabilitiesRegistry.removeNodes(removedNodes);

    // Remove DON
    uint32[] memory donIds = new uint32[](1);
    donIds[0] = DON_ID;
    s_CapabilitiesRegistry.removeDONs(donIds);

    // Remove node
    s_CapabilitiesRegistry.removeNodes(removedNodes);
    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, 0);
    assertEq(node.p2pId, bytes32(""));
    assertEq(node.signer, bytes32(""));
    assertEq(node.capabilityIds.length, 0, "Capabilities in the removed node should be cleared.");
    assertEq(node.configCount, 2, "Config cound should be incremented.");
  }

  function test_CanRemoveWhenNodeNoLongerPartOfDON() public {
    bytes32[] memory nodes = new bytes32[](3);
    nodes[0] = P2P_ID;
    nodes[1] = P2P_ID_TWO;
    nodes[2] = P2P_ID_THREE;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});

    // Add DON
    CapabilitiesRegistry.NewDONParams[] memory newDONs2 = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs2[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodes,
      capabilityConfigurations: capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: F_VALUE,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: bytes("")
    });
    s_CapabilitiesRegistry.addDONs(newDONs2);

    // Try remove nodes
    bytes32[] memory removedNodes = new bytes32[](1);
    removedNodes[0] = P2P_ID_TWO;
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.NodePartOfWorkflowDON.selector, 1, P2P_ID_TWO));
    s_CapabilitiesRegistry.removeNodes(removedNodes);

    // Update nodes in DON
    bytes32[] memory updatedNodes = new bytes32[](2);
    updatedNodes[0] = P2P_ID;
    updatedNodes[1] = P2P_ID_THREE;
    s_CapabilitiesRegistry.updateDON(
      DON_ID,
      CapabilitiesRegistry.UpdateDONParams({
        nodes: updatedNodes,
        capabilityConfigurations: capabilityConfigs,
        isPublic: true,
        f: F_VALUE,
        name: TEST_DON_NAME_ONE,
        config: bytes("")
      })
    );

    // Remove node
    s_CapabilitiesRegistry.removeNodes(removedNodes);
    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID_TWO);
    assertEq(node.nodeOperatorId, 0);
    assertEq(node.p2pId, bytes32(""));
    assertEq(node.signer, bytes32(""));
    assertEq(node.capabilityIds.length, 0, "Capabilities in the removed node should be cleared.");
    assertEq(node.configCount, 2, "Config cound should be incremented.");
  }

  function test_RemovesNode() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);

    bytes32[] memory nodesP2PIds = new bytes32[](1);
    nodesP2PIds[0] = P2P_ID;

    vm.expectEmit(address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeRemoved(P2P_ID);
    s_CapabilitiesRegistry.removeNodes(nodesP2PIds);

    CapabilitiesRegistry.NodeInfo memory removedNode = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(removedNode.nodeOperatorId, 0);
    assertEq(removedNode.p2pId, bytes32(""));
    assertEq(removedNode.signer, bytes32(""));
    assertEq(removedNode.capabilityIds.length, 0, "Capabilities in the removed node should be cleared.");
    assertEq(removedNode.configCount, 2, "Config cound should be incremented.");

    // 2nd part of the test to check that node capability configs are correctly cleared and don't show up
    // when the node is re-added.

    // Remove one of the capabilities so only one capability is readded
    s_twoCapabilitiesArray.pop();

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      // Add a node with a single capability. Previously this node had two capabilities.
      capabilityIds: s_twoCapabilitiesArray
    });

    s_CapabilitiesRegistry.addNodes(nodes);

    // Ensure that old capabilities do not show up
    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID, "Node should have the right node operator ID");
    assertEq(node.signer, NODE_OPERATOR_ONE_SIGNER_ADDRESS, "Node should have the right signer address");
    assertEq(node.p2pId, P2P_ID, "Node should have the right P2P ID");
    // Config count should be 3 because the node was added with two capabilities first, then it was removed and re-added
    // with one capability.
    assertEq(node.configCount, 3, "Node should have the right config count.");
    assertEq(node.capabilityIds.length, 1, "Node should have one capability");
    assertEq(node.capabilityIds[0], s_basicCapabilityId, "Node should have the right capability");
  }

  function test_CanAddNodeWithSameSignerAddressAfterRemoving() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);

    bytes32[] memory nodes = new bytes32[](1);
    nodes[0] = P2P_ID;

    s_CapabilitiesRegistry.removeNodes(nodes);

    CapabilitiesRegistry.NodeParams[] memory nodeParams = new CapabilitiesRegistry.NodeParams[](1);
    string[] memory capabilityIds = new string[](2);
    capabilityIds[0] = s_basicCapabilityId;
    capabilityIds[1] = s_capabilityWithConfigurationContractId;

    nodeParams[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    s_CapabilitiesRegistry.addNodes(nodeParams);

    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(node.p2pId, P2P_ID);
    assertEq(node.capabilityIds.length, 2);
    assertEq(node.capabilityIds[0], s_basicCapabilityId);
    assertEq(node.capabilityIds[1], s_capabilityWithConfigurationContractId);
    assertEq(node.configCount, 3, "Config count should be 2 because the node was added, removed and re-added.");
  }

  function test_OwnerCanRemoveNodes() public {
    bytes32[] memory nodes = new bytes32[](1);
    nodes[0] = P2P_ID;

    vm.expectEmit(address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeRemoved(P2P_ID);
    s_CapabilitiesRegistry.removeNodes(nodes);

    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, 0);
    assertEq(node.p2pId, bytes32(""));
    assertEq(node.signer, bytes32(""));
    assertEq(node.capabilityIds.length, 0, "Capabilities in the removed node should be cleared.");
    assertEq(node.configCount, 2, "Config count should be incremented.");
  }
}
