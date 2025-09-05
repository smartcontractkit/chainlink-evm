// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_setDONRegistry is WorkflowRegistrySetup {
  function test_setDONRegistry_WhenTheCallerIsNOTTheContractOwner() external {
    // It should revert with caller is not the owner
    vm.prank(s_stranger);
    address donReg = makeAddr("don-registry-address");
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setDONRegistry(donReg, 123456);
  }

  // whenTheCallerISTheContractOwner
  function test_setDONRegistry_WhenThereAreNoExistingRegistries() external {
    // It should write to s_donRegistry with the pair and emit DONRegistryUpdated
    vm.prank(s_owner);
    address donReg = makeAddr("don-registry-address");
    uint64 chainSel = 123456;

    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.DONRegistryUpdated(address(0), donReg, uint64(0), chainSel);
    s_registry.setDONRegistry(donReg, chainSel);

    (address donRegValue, uint64 chainSelValue) = s_registry.getDONRegistry();
    assertEq(chainSelValue, chainSel);
    assertEq(donRegValue, donReg);
  }

  // whenTheCallerISTheContractOwner
  function test_setDONRegistry_WhenBothRegistryAndChainSelectorDifferFromTheCurrentValues() external {
    // It should overwrite s_donRegistry with the new pair and emit DONRegistryUpdated

    vm.startPrank(s_owner);
    // set the don registry
    address donReg = makeAddr("don-registry-address");
    uint64 chainSel = 123456;

    s_registry.setDONRegistry(donReg, chainSel);

    // set it with different values
    address newDonReg = makeAddr("don-registry-address-2");
    uint64 newChainSel = 678910;

    s_registry.setDONRegistry(newDonReg, newChainSel);

    (address donRegValue, uint64 chainSelValue) = s_registry.getDONRegistry();
    assertEq(chainSelValue, newChainSel);
    assertEq(donRegValue, newDonReg);
    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner
  function test_setDONRegistry_WhenBothRegistryAndChainSelectorAreTheSameAsCurrent() external {
    // It should do nothing

    vm.startPrank(s_owner);
    // set the don registry
    address donReg = makeAddr("don-registry-address");
    uint64 chainSel = 123456;
    s_registry.setDONRegistry(donReg, chainSel);

    // set the same registry again
    vm.recordLogs();
    s_registry.setDONRegistry(donReg, chainSel);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("DONRegistryUpdated(address,address,uint64,uint64)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("DONLimitSet was emitted when it should not have been");
        fail();
      }
    }

    (address donRegValue, uint64 chainSelValue) = s_registry.getDONRegistry();
    assertEq(chainSelValue, chainSel);
    assertEq(donRegValue, donReg);

    vm.stopPrank();
  }
}
