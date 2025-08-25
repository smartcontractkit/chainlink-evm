// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_setDONLimit is WorkflowRegistrySetup {
  function test_setDONLimit_WhenTheCallerIsNOTTheContractOwner() external {
    // it should revert with Ownable2StepMsgSender: caller is not the owner
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setDONLimit(s_donFamily, 100, true);
  }

  // whenTheCallerISTheContractOwner whenEnabledIsTrue
  function test_setDONLimit_WhenNoPreviousLimitExistsForDonLabel() external {
    // it should set s_cfg.donLimit[donHash], append an event record, and emit DONLimitSet
    uint32 newLimit = 100;
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.DONLimitSet(s_donFamily, newLimit);

    s_registry.setDONLimit(s_donFamily, newLimit, true);
    uint32 donLimit = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, newLimit);

    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 1);
  }

  //   whenTheCallerISTheContractOwner
  //   whenEnabledIsTrue
  //   whenAPreviousLimitExistsForDonLabel
  function test_setDONLimit_WhenNewLimitDoesNotEqualExistingLimit() external {
    // it should overwrite s_cfg.donLimit[donHash] with the new value, append an event record, and emit DONLimitSet

    vm.startPrank(s_owner);
    // set a limit first
    s_registry.setDONLimit(s_donFamily, 100, true);
    assertEq(s_registry.getMaxWorkflowsPerDON(s_donFamily), 100);

    // set a different limit again
    uint32 newLimit = 200;
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.DONLimitSet(s_donFamily, newLimit);
    s_registry.setDONLimit(s_donFamily, newLimit, true);
    assertEq(s_registry.getMaxWorkflowsPerDON(s_donFamily), newLimit);

    // there should now be two event records for each capacity set
    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 2);

    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner
  // whenEnabledIsTrue
  // whenAPreviousLimitExistsForDonLabel
  function test_setDONLimit_WhenNewLimitIsEqualToExistingLimit() external {
    // it should do nothing

    vm.startPrank(s_owner);
    // set a limit first
    s_registry.setDONLimit(s_donFamily, 100, true);

    // set the same limit again
    vm.recordLogs();
    s_registry.setDONLimit(s_donFamily, 100, true);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("DONLimitSet(string,uint32)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("DONLimitSet was emitted when it should not have been");
        fail();
      }
    }

    assertEq(s_registry.getMaxWorkflowsPerDON(s_donFamily), 100);

    // only one event from when the limit was initially set, and no second one
    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 1);

    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner whenEnabledIsFalse
  function test_setDONLimit_WhenPreviousLimitExistsForDonLabel() external {
    // it should delete s_cfg.donLimit[donHash], append an event record with capacity set to 0, and emit DONLimitSet
    vm.startPrank(s_owner);
    // set a limit first
    s_registry.setDONLimit(s_donFamily, 100, true);
    assertEq(s_registry.getMaxWorkflowsPerDON(s_donFamily), 100);

    // remove the limit by passing in false to the last argument of setDONLimit
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.DONLimitSet(s_donFamily, 100);
    s_registry.setDONLimit(s_donFamily, 100, false);
    assertEq(s_registry.getMaxWorkflowsPerDON(s_donFamily), 0);

    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 2);

    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner whenEnabledIsFalse
  function test_setDONLimit_WhenNooPreviousLimitExistsForDonLabel() external {
    // it should do nothing
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, false);
    assertEq(s_registry.getMaxWorkflowsPerDON(s_donFamily), 0);

    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 0);
  }
}
