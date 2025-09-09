// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_UpdateNodeOperatorTest is BaseTest {
  uint32 private constant TEST_NODE_OPERATOR_ID = 1;
  address private constant NEW_NODE_OPERATOR_ADMIN = address(3);
  string private constant NEW_NODE_OPERATOR_NAME = "new-node-operator";

  function setUp() public override {
    BaseTest.setUp();
    changePrank(ADMIN);
    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
  }

  function test_RevertWhen_CalledByNonAdminAndNonOwner() public {
    changePrank(STRANGER);

    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](1);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: ADMIN, name: NEW_NODE_OPERATOR_NAME});

    uint32[] memory nodeOperatorIds = new uint32[](1);
    nodeOperatorIds[0] = TEST_NODE_OPERATOR_ID;

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.AccessForbidden.selector, STRANGER));
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);
  }

  function test_RevertWhen_NodeOperatorAdminIsZeroAddress() public {
    changePrank(ADMIN);
    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](1);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: address(0), name: NEW_NODE_OPERATOR_NAME});

    uint32[] memory nodeOperatorIds = new uint32[](1);
    nodeOperatorIds[0] = TEST_NODE_OPERATOR_ID;

    vm.expectRevert(CapabilitiesRegistry.InvalidNodeOperatorAdmin.selector);
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);
  }

  function test_RevertWhen_NodeOperatorIdAndParamLengthsMismatch() public {
    changePrank(ADMIN);
    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](1);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: NEW_NODE_OPERATOR_ADMIN, name: NEW_NODE_OPERATOR_NAME});

    uint32 invalidNodeOperatorId = 10_000;
    uint32[] memory nodeOperatorIds = new uint32[](2);
    nodeOperatorIds[0] = invalidNodeOperatorId;
    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.LengthMismatch.selector, nodeOperatorIds.length, nodeOperators.length)
    );
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);
  }

  function test_RevertWhen_NodeOperatorDoesNotExist() public {
    changePrank(ADMIN);
    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](1);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: NEW_NODE_OPERATOR_ADMIN, name: NEW_NODE_OPERATOR_NAME});

    uint32 invalidNodeOperatorId = 10_000;
    uint32[] memory nodeOperatorIds = new uint32[](1);
    nodeOperatorIds[0] = invalidNodeOperatorId;
    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.NodeOperatorDoesNotExist.selector, invalidNodeOperatorId)
    );
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);
  }

  function test_RevertWhen_DuplicateNodeOperator() public {
    changePrank(ADMIN);

    // Try to make Node Operator 1 the same as Node Operator 2
    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](1);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: NODE_OPERATOR_TWO_ADMIN, name: NODE_OPERATOR_TWO_NAME});

    uint32[] memory nodeOperatorIds = new uint32[](1);
    nodeOperatorIds[0] = TEST_NODE_OPERATOR_ID;

    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.NodeOperatorAlreadyExists.selector, 2));
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);
  }

  function test_UpdatesNodeOperator_UnblocksReAddingOldNodeOperator() public {
    changePrank(ADMIN);
    // First update the node operator
    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](1);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: NEW_NODE_OPERATOR_ADMIN, name: NEW_NODE_OPERATOR_NAME});

    uint32[] memory nodeOperatorIds = new uint32[](1);
    nodeOperatorIds[0] = TEST_NODE_OPERATOR_ID;

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorUpdated(
      TEST_NODE_OPERATOR_ID, NEW_NODE_OPERATOR_ADMIN, NEW_NODE_OPERATOR_NAME
    );
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);

    // Then re-add the old node operator
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: NODE_OPERATOR_ONE_ADMIN, name: NODE_OPERATOR_ONE_NAME});
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorUpdated(
      TEST_NODE_OPERATOR_ID, NODE_OPERATOR_ONE_ADMIN, NODE_OPERATOR_ONE_NAME
    );
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);
  }

  function test_UpdatesNodeOperator() public {
    changePrank(ADMIN);

    CapabilitiesRegistry.NodeOperator[] memory nodeOperators = new CapabilitiesRegistry.NodeOperator[](1);
    nodeOperators[0] = CapabilitiesRegistry.NodeOperator({admin: NEW_NODE_OPERATOR_ADMIN, name: NEW_NODE_OPERATOR_NAME});

    uint32[] memory nodeOperatorIds = new uint32[](1);
    nodeOperatorIds[0] = TEST_NODE_OPERATOR_ID;

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorUpdated(
      TEST_NODE_OPERATOR_ID, NEW_NODE_OPERATOR_ADMIN, NEW_NODE_OPERATOR_NAME
    );
    s_CapabilitiesRegistry.updateNodeOperators(nodeOperatorIds, nodeOperators);

    CapabilitiesRegistry.NodeOperator memory nodeOperator =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_ID);
    assertEq(nodeOperator.admin, NEW_NODE_OPERATOR_ADMIN);
    assertEq(nodeOperator.name, NEW_NODE_OPERATOR_NAME);
  }
}
