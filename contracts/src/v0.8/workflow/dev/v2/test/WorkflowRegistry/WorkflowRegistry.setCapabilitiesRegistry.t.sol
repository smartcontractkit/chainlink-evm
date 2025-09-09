// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_setCapabilitiesRegistry is WorkflowRegistrySetup {
  function test_setCapabilitiesRegistry_WhenTheCallerIsNOTTheContractOwner() external {
    // It should revert with caller is not the owner
    vm.prank(s_stranger);
    address capReg = makeAddr("cap-registry-address");
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setCapabilitiesRegistry(capReg, 123_456);
  }

  // whenTheCallerISTheContractOwner
  function test_setCapabilitiesRegistry_WhenThereAreNoExistingRegistries() external {
    // It should write to s_capabilitiesRegistry with the pair and emit CapabilitiesRegistryUpdated
    vm.prank(s_owner);
    address capReg = makeAddr("cap-registry-address");
    uint64 chainSel = 123_456;

    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.CapabilitiesRegistryUpdated(address(0), capReg, uint64(0), chainSel);
    s_registry.setCapabilitiesRegistry(capReg, chainSel);

    (address capRegValue, uint64 chainSelValue) = s_registry.getCapabilitiesRegistry();
    assertEq(chainSelValue, chainSel);
    assertEq(capRegValue, capReg);
  }

  // whenTheCallerISTheContractOwner
  function test_setCapabilitiesRegistry_WhenBothRegistryAndChainSelectorDifferFromTheCurrentValues() external {
    // It should overwrite s_capabilitiesRegistry with the new pair and emit CapabilitiesRegistryUpdated

    vm.startPrank(s_owner);
    // set the capabilities registry
    address capReg = makeAddr("cap-registry-address");
    uint64 chainSel = 123_456;

    s_registry.setCapabilitiesRegistry(capReg, chainSel);

    // set it with different values
    address newCapReg = makeAddr("cap-registry-address-2");
    uint64 newChainSel = 678_910;

    s_registry.setCapabilitiesRegistry(newCapReg, newChainSel);

    (address capRegValue, uint64 chainSelValue) = s_registry.getCapabilitiesRegistry();
    assertEq(chainSelValue, newChainSel);
    assertEq(capRegValue, newCapReg);
    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner
  function test_setCapabilitiesRegistry_WhenBothRegistryAndChainSelectorAreTheSameAsCurrent() external {
    // It should do nothing

    vm.startPrank(s_owner);
    // set the capabilities registry
    address capReg = makeAddr("cap-registry-address");
    uint64 chainSel = 123_456;
    s_registry.setCapabilitiesRegistry(capReg, chainSel);

    // set the same registry again
    vm.recordLogs();
    s_registry.setCapabilitiesRegistry(capReg, chainSel);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("CapabilitiesRegistryUpdated(address,address,uint64,uint64)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("CapabilitiesRegistryUpdated was emitted when it should not have been");
        fail();
      }
    }

    (address capRegValue, uint64 chainSelValue) = s_registry.getCapabilitiesRegistry();
    assertEq(chainSelValue, chainSel);
    assertEq(capRegValue, capReg);

    vm.stopPrank();
  }
}
