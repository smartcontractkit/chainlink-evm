// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {INodeInfoProvider} from "../../interfaces/INodeInfoProvider.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_GetNodesTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    changePrank(ADMIN);

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);

    changePrank(NODE_OPERATOR_ONE_ADMIN);

    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);
  }

  function test_CorrectlyFetchesNodes() public view {
    CapabilitiesRegistry.NodeInfo[] memory nodes = s_CapabilitiesRegistry.getNodes();
    assertEq(nodes.length, 2);

    assertEq(nodes[0].nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(nodes[0].signer, NODE_OPERATOR_ONE_SIGNER_ADDRESS);
    assertEq(nodes[0].p2pId, P2P_ID);
    assertEq(nodes[0].capabilityIds.length, 2);
    assertEq(nodes[0].capabilityIds[0], s_basicCapabilityId);
    assertEq(nodes[0].capabilityIds[1], s_capabilityWithConfigurationContractId);
    assertEq(nodes[0].configCount, 1);

    assertEq(nodes[1].nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(nodes[1].signer, NODE_OPERATOR_TWO_SIGNER_ADDRESS);
    assertEq(nodes[1].p2pId, P2P_ID_TWO);
    assertEq(nodes[1].capabilityIds.length, 2);
    assertEq(nodes[1].capabilityIds[0], s_basicCapabilityId);
    assertEq(nodes[1].capabilityIds[1], s_capabilityWithConfigurationContractId);
    assertEq(nodes[1].configCount, 1);
  }

  function test_CorrectlyFetchesSpecificNodes() public view {
    bytes32[] memory p2pIds = new bytes32[](1);
    p2pIds[0] = P2P_ID;

    CapabilitiesRegistry.NodeInfo[] memory nodes = s_CapabilitiesRegistry.getNodesByP2PIds(p2pIds);
    assertEq(nodes.length, 1);
    assertEq(nodes[0].p2pId, P2P_ID);
    assertEq(nodes[0].signer, NODE_OPERATOR_ONE_SIGNER_ADDRESS);
  }

  function test_GetNodesByP2PIdsInvalidNode_Reverts() public {
    bytes32[] memory p2pIds = new bytes32[](1);
    p2pIds[0] = keccak256(abi.encodePacked("invalid"));

    vm.expectRevert(abi.encodeWithSelector(INodeInfoProvider.NodeDoesNotExist.selector, p2pIds[0]));
    s_CapabilitiesRegistry.getNodesByP2PIds(p2pIds);
  }

  function test_DoesNotIncludeRemovedNodes() public {
    changePrank(ADMIN);
    bytes32[] memory nodesToRemove = new bytes32[](1);
    nodesToRemove[0] = P2P_ID_TWO;
    s_CapabilitiesRegistry.removeNodes(nodesToRemove);

    CapabilitiesRegistry.NodeInfo[] memory nodes = s_CapabilitiesRegistry.getNodes();
    assertEq(nodes.length, 1);

    assertEq(nodes[0].nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(nodes[0].signer, NODE_OPERATOR_ONE_SIGNER_ADDRESS);
    assertEq(nodes[0].p2pId, P2P_ID);
    assertEq(nodes[0].capabilityIds.length, 2);
    assertEq(nodes[0].capabilityIds[0], s_basicCapabilityId);
    assertEq(nodes[0].capabilityIds[1], s_capabilityWithConfigurationContractId);
    assertEq(nodes[0].configCount, 1);
  }
}
