// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_RemoveNodeOperatorsTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();
    changePrank(ADMIN);
    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
  }

  function test_RevertWhen_CalledByNonOwner() public {
    changePrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    uint32[] memory nodeOperatorsToRemove = new uint32[](2);
    nodeOperatorsToRemove[1] = 1;
    s_CapabilitiesRegistry.removeNodeOperators(nodeOperatorsToRemove);
  }

  function test_RemovesNodeOperator() public {
    changePrank(ADMIN);

    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorRemoved(TEST_NODE_OPERATOR_ONE_ID);
    vm.expectEmit(true, true, true, true, address(s_CapabilitiesRegistry));
    emit CapabilitiesRegistry.NodeOperatorRemoved(TEST_NODE_OPERATOR_TWO_ID);
    uint32[] memory nodeOperatorsToRemove = new uint32[](2);
    nodeOperatorsToRemove[0] = TEST_NODE_OPERATOR_ONE_ID;
    nodeOperatorsToRemove[1] = TEST_NODE_OPERATOR_TWO_ID;
    s_CapabilitiesRegistry.removeNodeOperators(nodeOperatorsToRemove);

    CapabilitiesRegistry.NodeOperator memory nodeOperatorOne =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_ONE_ID);
    assertEq(nodeOperatorOne.admin, address(0));
    assertEq(nodeOperatorOne.name, "");

    CapabilitiesRegistry.NodeOperator memory nodeOperatorTwo =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_TWO_ID);
    assertEq(nodeOperatorTwo.admin, address(0));
    assertEq(nodeOperatorTwo.name, "");
  }

  function test_RemovesNodeOperator_UnblocksReAdding() public {
    changePrank(ADMIN);

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

    CapabilitiesRegistry.NodeOperator memory nodeOperatorOne =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_ONE_ID);
    assertEq(nodeOperatorOne.admin, address(0));
    assertEq(nodeOperatorOne.name, "");

    CapabilitiesRegistry.NodeOperator memory nodeOperatorTwo =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_TWO_ID);
    assertEq(nodeOperatorTwo.admin, address(0));
    assertEq(nodeOperatorTwo.name, "");

    CapabilitiesRegistry.NodeOperator memory nodeOperatorThree =
      s_CapabilitiesRegistry.getNodeOperator(TEST_NODE_OPERATOR_THREE_ID);
    assertEq(nodeOperatorThree.admin, address(0));
    assertEq(nodeOperatorThree.name, "");

    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
  }
}
