// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../shared/access/Ownable2Step.sol";
import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_RemoveNodeOperatorsTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
  }

  function test_RevertWhen_CalledByNonOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    uint32[] memory nodeOperatorsToRemove = new uint32[](2);
    nodeOperatorsToRemove[1] = 1;
    s_CapabilitiesRegistry.removeNodeOperators(nodeOperatorsToRemove);
  }

  function test_RevertWhen_InUseOnNode() public {
    // Setup
    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);

    // Test
    vm.expectRevert(
      abi.encodeWithSelector(CapabilitiesRegistry.NodeOperatorHasNodes.selector, TEST_NODE_OPERATOR_ONE_ID)
    );
    uint32[] memory nodeOperatorsToRemove = new uint32[](1);
    nodeOperatorsToRemove[0] = TEST_NODE_OPERATOR_ONE_ID;
    s_CapabilitiesRegistry.removeNodeOperators(nodeOperatorsToRemove);
  }

  function test_RemovesNodeOperator() public {
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorRemoved(TEST_NODE_OPERATOR_ONE_ID);
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorRemoved(TEST_NODE_OPERATOR_TWO_ID);
    uint32[] memory nodeOperatorsToRemove = new uint32[](2);
    nodeOperatorsToRemove[0] = TEST_NODE_OPERATOR_ONE_ID;
    nodeOperatorsToRemove[1] = TEST_NODE_OPERATOR_TWO_ID;
    s_CapabilitiesRegistry.removeNodeOperators(nodeOperatorsToRemove);

    CapabilitiesRegistry.NodeOperatorInfo memory nodeOperatorOne =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_ONE_ID);
    assertEq(nodeOperatorOne.admin, address(0));
    assertEq(nodeOperatorOne.name, "");

    CapabilitiesRegistry.NodeOperatorInfo memory nodeOperatorTwo =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_TWO_ID);
    assertEq(nodeOperatorTwo.admin, address(0));
    assertEq(nodeOperatorTwo.name, "");
  }

  function test_RemovesNodeOperator_UnblocksReAdding() public {
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorRemoved(TEST_NODE_OPERATOR_ONE_ID);
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorRemoved(TEST_NODE_OPERATOR_TWO_ID);
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorRemoved(TEST_NODE_OPERATOR_THREE_ID);
    uint32[] memory nodeOperatorsToRemove = new uint32[](3);
    nodeOperatorsToRemove[0] = TEST_NODE_OPERATOR_ONE_ID;
    nodeOperatorsToRemove[1] = TEST_NODE_OPERATOR_TWO_ID;
    nodeOperatorsToRemove[2] = TEST_NODE_OPERATOR_THREE_ID;
    s_CapabilitiesRegistry.removeNodeOperators(nodeOperatorsToRemove);

    CapabilitiesRegistry.NodeOperatorInfo memory nodeOperatorOne =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_ONE_ID);
    assertEq(nodeOperatorOne.admin, address(0));
    assertEq(nodeOperatorOne.name, "");

    CapabilitiesRegistry.NodeOperatorInfo memory nodeOperatorTwo =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_TWO_ID);
    assertEq(nodeOperatorTwo.admin, address(0));
    assertEq(nodeOperatorTwo.name, "");

    CapabilitiesRegistry.NodeOperatorInfo memory nodeOperatorThree =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_THREE_ID);
    assertEq(nodeOperatorThree.admin, address(0));
    assertEq(nodeOperatorThree.name, "");

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
  }
}
