// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

contract WorkflowRegistry_setUserDONOverride {
  function test_WhenCallerIsNOTTheOwner() external {
    // it reverts with OnlyCallableByOwner
  }

  modifier whenCallerISTheOwner() {
    _;
  }

  function test_WhenThereAreNoOverridesYet() external whenCallerISTheOwner {
    // it is the default value
    //     getMaxWorkflowsPerUserDon(user, donLabel) returns the default (e.g. 200)
  }

  function test_WhenLimitIsSetTo0() external whenCallerISTheOwner {
    // it correctly sets the limit to 0
    //     call setUserDONOverride(user, donLabel, 0)
    //     getMaxWorkflowsPerUserDon(user, donLabel) returns default (override cleared)
  }

  function test_WhenLimitIsSetToANormalPositiveValue() external whenCallerISTheOwner {
    // it correctly sets the limit
    //     call setUserDONOverride(user, donLabel, 42)
    //     getMaxWorkflowsPerUserDon(user, donLabel) returns 42
  }

  function test_WhenLimitIsUint32Max4294967295() external whenCallerISTheOwner {
    // it correctly sets the limit
    //     call setUserDONOverride(user, donLabel, 4294967295)
    //     getMaxWorkflowsPerUserDon(user, donLabel) returns 4294967295
  }

  function test_WhenItIsCalledMultipleTimes() external whenCallerISTheOwner {
    // it correctly sets the latest value
    //     call setUserDONOverride(user, donLabel, 10)
    //     getMaxWorkflowsPerUserDon(user, donLabel) returns 10
    //     call setUserDONOverride(user, donLabel, 20)
    //     getMaxWorkflowsPerUserDon(user, donLabel) returns 20 (updated)
  }

  function test_WhenThereAreMultipleUsersAndMultipleDONs() external whenCallerISTheOwner {
    // it correctly sets the limit for the specific user and DON
    //     setUserDONOverride(userA, donLabelA, 33)
    //     getMaxWorkflowsPerUserDon(userA, donLabelA) returns 33
    //     getMaxWorkflowsPerUserDon(userA, donLabelB) returns default (200)
    //     getMaxWorkflowsPerUserDon(userB, donLabelA) returns default (200)
  }
}
