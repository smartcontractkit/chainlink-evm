// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_getEvents is WorkflowRegistrySetup {
  function test_getEvents_WhenNoEventsHaveBeenRecorded() external view {
    // it should returns an empty array
    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(0, 100);
    assertEq(events.length, 0);
  }

  modifier whenThereIsMoreThan1Event() {
    vm.startPrank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, true);
    s_registry.setDONLimit(s_donFamily, 200, true);
    s_registry.setDONLimit(s_donFamily, 150, true); // should push three events to event records
    vm.stopPrank();
    _;
  }

  // when start ≥ N
  function test_getEvents_WhenStartIsGreaterThanNumberOfEvents() external whenThereIsMoreThan1Event {
    // it should return an empty array
    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(3, 100);
    assertEq(events.length, 0);
  }

  // when start < N
  modifier whenStartIsLessThanNumberOfEventsN() {
    _;
  }

  // when limit = 0
  function test_getEvents_WhenLimitIs0() external whenThereIsMoreThan1Event whenStartIsLessThanNumberOfEventsN {
    // it should return an empty array
    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(1, 0);
    assertEq(events.length, 0);
  }

  // when 0 < limit < (N – start)
  function test_getEvents_When0IsLessThanLimitWhichIsLessThanNMinusStart()
    external
    whenThereIsMoreThan1Event
    whenStartIsLessThanNumberOfEventsN
  {
    // it should return limit - start number of events
    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(1, 3);
    assertEq(events.length, 2);
  }

  function test_getEvents_WhenLimitIsGreaterThanNMinusStart()
    external
    whenThereIsMoreThan1Event
    whenStartIsLessThanNumberOfEventsN
  {
    // it should return the last N – start events beginning at index start
    WorkflowRegistry.EventRecord[] memory events = s_registry.getEvents(1, 20);
    assertEq(events.length, 2);
  }
}
