// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_AddNodesTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    changePrank(ADMIN);
    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(capabilities);
  }

  function test_RevertWhen_CalledByNonNodeOperatorAdminAndNonOwner() public {
    changePrank(STRANGER);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.AccessForbidden.selector, STRANGER));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_AddingNodeWithInvalidNodeOperator() public {
    changePrank(ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    uint32 invalidNodeOperatorId = 10_000;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: invalidNodeOperatorId, // Invalid NOP
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.NodeOperatorDoesNotExist.selector, invalidNodeOperatorId)
    );
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_SignerAddressEmpty() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: bytes32(""),
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidNodeSigner.selector));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_EncryptionPublicKeyEmpty() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: bytes32(""),
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidNodeEncryptionPublicKey.selector, bytes32("")));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_CSAKeyEmpty() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: bytes32(""),
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidNodeCSAKey.selector, bytes32("")));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_SignerAddressNotUnique() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    s_CapabilitiesRegistry.addNodes(nodes);

    changePrank(NODE_OPERATOR_TWO_ADMIN);

    // Try adding another node with the same signer address
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_TWO_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: capabilityIds
    });
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidNodeSigner.selector));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_AddingDuplicateP2PId() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    s_CapabilitiesRegistry.addNodes(nodes);

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.NodeAlreadyExists.selector, P2P_ID));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_P2PIDEmpty() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: bytes32(""),
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidNodeP2PId.selector, bytes32("")));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_AddingNodeWithoutCapabilities() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](0);

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidNodeCapabilities.selector, capabilityIds));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_AddingNodeWithInvalidCapability() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_nonExistentCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.InvalidNodeCapabilities.selector, capabilityIds));
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_AddsNodeParams() public {
    changePrank(NODE_OPERATOR_ONE_ADMIN);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
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

    vm.expectEmit(address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeAdded(P2P_ID, TEST_NODE_OPERATOR_ONE_ID, NODE_OPERATOR_ONE_SIGNER_ADDRESS);
    s_CapabilitiesRegistry.addNodes(nodes);

    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(node.p2pId, P2P_ID);
    assertEq(node.capabilityIds.length, 2);
    assertEq(node.capabilityIds[0], s_basicCapabilityId);
    assertEq(node.capabilityIds[1], s_capabilityWithConfigurationContractId);
    assertEq(node.configCount, 1);
  }

  function test_OwnerCanAddNodes() public {
    changePrank(ADMIN);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
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

    vm.expectEmit(address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeAdded(P2P_ID, TEST_NODE_OPERATOR_ONE_ID, NODE_OPERATOR_ONE_SIGNER_ADDRESS);
    s_CapabilitiesRegistry.addNodes(nodes);

    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(node.p2pId, P2P_ID);
    assertEq(node.capabilityIds.length, 2);
    assertEq(node.capabilityIds[0], s_basicCapabilityId);
    assertEq(node.capabilityIds[1], s_capabilityWithConfigurationContractId);
    assertEq(node.configCount, 1);
  }
}
