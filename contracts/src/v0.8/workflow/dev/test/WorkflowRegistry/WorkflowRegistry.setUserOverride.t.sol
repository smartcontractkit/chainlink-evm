// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.24;

contract WorkflowRegistry_setUserOverride {
  function test_WhenCallerIsNOTTheOwner() external {
    // it should revert with OnlyCallableByOwner
  }

  modifier whenCallerISTheOwner() {
    _;
  }

  function test_WhenThereAreNoOverridesYet() external whenCallerISTheOwner {
    // it is the default
    //     getMaxWorkflowsPerUser(user) returns the default (e.g. 200)
  }

  function test_WhenLimitIsSetTo0() external whenCallerISTheOwner {
    // it correctly sets the limit to 0
    //     call setUserOverride(user, 0)
    //     getMaxWorkflowsPerUser(user) returns default (override cleared)
  }

  function test_WhenLimitIsSetToANormalPositiveValue() external whenCallerISTheOwner {
    // it correctly sets the limit
    //     call setUserOverride(user, 42)
    //     getMaxWorkflowsPerUser(user) returns 42
  }

  function test_WhenLimitIsSetToUint32Max4294967295() external whenCallerISTheOwner {
    // it correctly sets the limit
    //     call setUserOverride(user, 4294967295)
    //     getMaxWorkflowsPerUser(user) returns 4294967295
  }

  function test_WhenTheFunctionIsCalledMultipleTimes() external whenCallerISTheOwner {
    // it correctly sets the latest value
    //     call setUserOverride(user, 10)
    //     getMaxWorkflowsPerUser(user) returns 10
    //     call setUserOverride(user, 20)
    //     getMaxWorkflowsPerUser(user) returns 20 (updated)
  }

  function test_WhenThereAreMultipleUsers() external whenCallerISTheOwner {
    // it correctly sets it for each user
    //     setUserOverride(userA, 99)
    //     getMaxWorkflowsPerUser(userA) returns 99
    //     getMaxWorkflowsPerUser(userB) returns default (e.g. 200)
  }
}
