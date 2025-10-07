// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_setDONLimit is WorkflowRegistrySetup {
  function test_setDONLimit_WhenTheCallerIsNOTTheContractOwner() external {
    // it should revert with Ownable2StepMsgSender: caller is not the owner
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setDONLimit(s_donFamily, 100, 10);
  }

  // whenTheCallerISTheContractOwner whenEnabledIsTrue
  function test_setDONLimit_WhenNoPreviousLimitExistsForDonLabel() external {
    // it should set s_cfg.donLimit[donHash], append an event record, and emit DONLimitSet
    uint32 newLimit = 100;
    uint32 newDefaultUserLimit = 10;
    vm.prank(s_owner);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.DONLimitSet(s_donFamily, newLimit, newDefaultUserLimit);

    s_registry.setDONLimit(s_donFamily, newLimit, newDefaultUserLimit);
    (uint32 donLimit, uint32 defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, newLimit);
    assertEq(defaultUserLimit, newDefaultUserLimit);

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
    s_registry.setDONLimit(s_donFamily, 100, 10);
    (uint32 donLimit, uint32 defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, 100);
    assertEq(defaultUserLimit, 10);

    // set a different limit again
    uint32 newLimit = 200;
    uint32 newDefaultUserLimit = 20;
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.DONLimitSet(s_donFamily, newLimit, newDefaultUserLimit);
    s_registry.setDONLimit(s_donFamily, newLimit, newDefaultUserLimit);
    (donLimit, defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, newLimit);
    assertEq(defaultUserLimit, newDefaultUserLimit);

    // check that default user limit can't be higher than DON limit
    vm.expectRevert(WorkflowRegistry.UserDONDefaultLimitExceedsDONLimit.selector);
    s_registry.setDONLimit(s_donFamily, 250, 251);

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
    s_registry.setDONLimit(s_donFamily, 100, 10);

    // set the same limit again
    vm.recordLogs();
    s_registry.setDONLimit(s_donFamily, 100, 10);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("DONLimitSet(string,uint32,uint32)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("DONLimitSet was emitted when it should not have been");
        fail();
      }
    }

    (uint32 donLimit, uint32 defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, 100);
    assertEq(defaultUserLimit, 10);

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
    s_registry.setDONLimit(s_donFamily, 100, 10);
    (uint32 donLimit, uint32 defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, 100);
    assertEq(defaultUserLimit, 10);

    // remove the limit by setting it to zero
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.DONLimitSet(s_donFamily, 0, 0);
    // adding new workflows should be disabled now
    s_registry.setDONLimit(s_donFamily, 0, 0);

    (donLimit, defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, 0);
    assertEq(defaultUserLimit, 0);

    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 2);

    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner whenEnabledIsFalse
  function test_setDONLimit_WhenNooPreviousLimitExistsForDonLabel() external {
    // it should do nothing
    vm.prank(s_owner);
    // set a global limit to zero to disable adding new workflows to the DON
    // but default user limit can't be greater than DON limit so it has to be disabled as well
    vm.expectRevert(WorkflowRegistry.UserDONDefaultLimitExceedsDONLimit.selector);
    s_registry.setDONLimit(s_donFamily, 0, 10);

    // set a global limit to zero to disable adding new workflows to the DON
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 0, 0);
    (uint32 donLimit, uint32 defaultUserLimit) = s_registry.getMaxWorkflowsPerDON(s_donFamily);
    assertEq(donLimit, 0);
    assertEq(defaultUserLimit, 0);

    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 0);
  }
}
