// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {INodeInfoProvider} from "../../interfaces/INodeInfoProvider.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_UpdateNodesTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    vm.stopPrank();
    vm.startPrank(ADMIN);
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](2);
    capabilities[0] = s_basicCapability;
    capabilities[1] = s_capabilityWithConfigurationContract;

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addCapabilities(capabilities);

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

    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    s_CapabilitiesRegistry.addNodes(nodes);

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_TWO_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: capabilityIds
    });

    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_TWO_ADMIN);
    s_CapabilitiesRegistry.addNodes(nodes);
  }

  function test_RevertWhen_CalledByNonNodeOperatorAdminAndNonOwner() public {
    vm.startPrank(STRANGER);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.AccessForbidden.selector, STRANGER));
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_CalledByAnotherNodeOperatorAdmin() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_TWO_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_TWO_ID,
      p2pId: P2P_ID,
      signer: NEW_NODE_SIGNER,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.AccessForbidden.selector, NODE_OPERATOR_TWO_ADMIN));
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_NodeDoesNotExist() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: INVALID_P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(abi.encodeWithSelector(INodeInfoProvider.NodeDoesNotExist.selector, INVALID_P2P_ID));
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_P2PIDEmpty() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
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

    vm.expectRevert(abi.encodeWithSelector(INodeInfoProvider.NodeDoesNotExist.selector, bytes32("")));
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_SignerAddressEmpty() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
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
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_EncryptionPublicKeyEmpty() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
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
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_CSAKeyEmpty() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
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
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_NodeSignerAlreadyAssignedToAnotherNode() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);

    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: capabilityIds
    });

    vm.expectRevert(CapabilitiesRegistry.InvalidNodeSigner.selector);
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_UpdatingNodeWithoutCapabilities() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
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
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_AddingNodeWithInvalidCapability() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);
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
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_RemovingCapabilityRequiredByWorkflowDON() public {
    // SETUP: addDON
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    bytes32[] memory nodeIds = new bytes32[](2);
    nodeIds[0] = P2P_ID;
    nodeIds[1] = P2P_ID_TWO;

    // SETUP: updateNodes
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
    string[] memory capabilityIds = new string[](1);
    // DON requires s_basicCapabilityId but we are swapping for
    // s_capabilityWithConfigurationContractId
    capabilityIds[0] = s_capabilityWithConfigurationContractId;
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });
    uint32 workflowDonId = 1;

    // Operations
    vm.startPrank(ADMIN);
    CapabilitiesRegistry.NewDONParams[] memory newDONs = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodeIds,
      capabilityConfigurations: capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: true,
      f: 1,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: bytes("")
    });
    s_CapabilitiesRegistry.addDONs(newDONs);

    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.CapabilityRequiredByDON.selector, s_basicCapabilityId, workflowDonId)
    );
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_RevertWhen_RemovingCapabilityRequiredByCapabilityDON() public {
    // SETUP: addDON
    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] =
      CapabilitiesRegistry.CapabilityConfiguration({capabilityId: s_basicCapabilityId, config: BASIC_CAPABILITY_CONFIG});
    bytes32[] memory nodeIds = new bytes32[](2);
    nodeIds[0] = P2P_ID;
    nodeIds[1] = P2P_ID_TWO;

    // SETUP: updateNodes
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
    string[] memory capabilityIds = new string[](1);
    // DON requires s_basicCapabilityId but we are swapping for
    // s_capabilityWithConfigurationContractId
    capabilityIds[0] = s_capabilityWithConfigurationContractId;
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });
    uint32 capabilitiesDonId = 1;

    // Operations
    vm.stopPrank();
    vm.startPrank(ADMIN);
    CapabilitiesRegistry.NewDONParams[] memory newDONs2 = new CapabilitiesRegistry.NewDONParams[](1);
    newDONs2[0] = CapabilitiesRegistry.NewDONParams({
      nodes: nodeIds,
      capabilityConfigurations: capabilityConfigs,
      isPublic: true,
      acceptsWorkflows: false,
      f: 1,
      name: TEST_DON_NAME_ONE,
      donFamilies: new string[](0),
      config: bytes("")
    });
    s_CapabilitiesRegistry.addDONs(newDONs2);

    vm.expectRevert(
      abi.encodeWithSelector(
        CapabilitiesRegistry.CapabilityRequiredByDON.selector, s_basicCapabilityId, capabilitiesDonId
      )
    );
    s_CapabilitiesRegistry.updateNodes(nodes);
  }

  function test_CanUpdateParamsIfNodeSignerAddressNoLongerUsed() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    // Set node one's signer to another address
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: bytes32(abi.encodePacked(address(6666))),
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    s_CapabilitiesRegistry.updateNodes(nodes);

    // Set node two's signer to node one's signer
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_TWO_ADMIN);
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_TWO_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      csaKey: TEST_CSA_KEY_TWO,
      capabilityIds: capabilityIds
    });
    s_CapabilitiesRegistry.updateNodes(nodes);

    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID_TWO);
    assertEq(node.signer, NODE_OPERATOR_ONE_SIGNER_ADDRESS);
  }

  function test_UpdatesNodeParams() public {
    vm.stopPrank();
    vm.startPrank(NODE_OPERATOR_ONE_ADMIN);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NEW_NODE_SIGNER,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectEmit(address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeUpdated(P2P_ID, TEST_NODE_OPERATOR_ONE_ID, NEW_NODE_SIGNER);
    s_CapabilitiesRegistry.updateNodes(nodes);

    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(node.p2pId, P2P_ID);
    assertEq(node.signer, NEW_NODE_SIGNER);
    assertEq(node.capabilityIds.length, 1);
    assertEq(node.capabilityIds[0], s_basicCapabilityId);
    assertEq(node.configCount, 2);
  }

  function test_OwnerCanUpdateNodes() public {
    vm.stopPrank();
    vm.startPrank(ADMIN);

    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](1);
    string[] memory capabilityIds = new string[](1);
    capabilityIds[0] = s_basicCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NEW_NODE_SIGNER,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    vm.expectEmit(address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeUpdated(P2P_ID, TEST_NODE_OPERATOR_ONE_ID, NEW_NODE_SIGNER);
    s_CapabilitiesRegistry.updateNodes(nodes);

    CapabilitiesRegistry.NodeInfo memory node = s_CapabilitiesRegistry.getNode(P2P_ID);
    assertEq(node.nodeOperatorId, TEST_NODE_OPERATOR_ONE_ID);
    assertEq(node.p2pId, P2P_ID);
    assertEq(node.signer, NEW_NODE_SIGNER);
    assertEq(node.capabilityIds.length, 1);
    assertEq(node.capabilityIds[0], s_basicCapabilityId);
    assertEq(node.configCount, 2);
  }
}
