// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

contract WorkflowRegistry_setDefaults {
  function test_WhenTheCallerIsNOTTheOwner() external {
    // it should revert with OnlyCallableByOwner
  }

  modifier whenTheCallerISTheOwner() {
    _;
  }

  function test_WhenThereAreNoCallsMadeYet() external whenTheCallerISTheOwner {
    // it should be the constructor defaults (500, 200)
  }

  function test_WhenItIsCalledWithATypicalValidUpdate() external whenTheCallerISTheOwner {
    // it should correctly set the updated values
    //     call setDefaults(100, 200)
    //         getMaxWorkflowsPerDon(...) returns 100
    //         getMaxWorkflowsPerUserDon(...) returns 200
    //     other mappings/overrides remain untouched
  }

  function test_WhenAllValuesAreZero() external whenTheCallerISTheOwner {
    // it should set 0 for all two values
    //     call setDefaults(0, 0)
    //         getMaxWorkflowsPerDon(...) returns 0
    //         getMaxWorkflowsPerUserDon(...) returns 0
  }

  function test_WhenAllValuesAreAtUint32Max() external whenTheCallerISTheOwner {
    // it should set uint32 max for all two values
    //     call setDefaults(4294967295, 4294967295)
    //         getMaxWorkflowsPerDon(...) returns 4294967295
    //         getMaxWorkflowsPerUserDon(...) returns 4294967295
  }

  function test_WhenCalledMultipleTimesInSequence() external whenTheCallerISTheOwner {
    // it should set to the most recent values
    //     call setDefaults(A, B)
    //         getters reflect (A, B)
    //     call setDefaults(C, D)
    //         getters now reflect (C, D) (overwriting previous)
  }
}
