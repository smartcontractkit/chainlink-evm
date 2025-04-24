// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.24;

contract WorkflowRegistry_setDONOverride {
  function test_WhenCallerIsNOTTheOwner() external {
    // it should revert with OnlyCallableByOwner
  }

  modifier whenCallerISTheOwner() {
    _;
  }

  function test_WhenThereAreNoOverridesYet() external whenCallerISTheOwner {
    // it is the default value
    //     call getMaxWorkflowsPerDon(donLabelA) returns the default (e.g. 500)
  }

  function test_WhenLimitIsSetTo0() external whenCallerISTheOwner {
    // it correctly sets the limit to 0
    //     call setDONOverride(donLabelA, 0)
    //     getMaxWorkflowsPerDon(donLabelA) returns default (override cleared)
  }

  function test_WhenLimitIsSetToANormalPositiveValue() external whenCallerISTheOwner {
    // it correctly sets the limit
    //     call setDONOverride(donLabelA, 123)
    //     getMaxWorkflowsPerDon(donLabelA) returns 123
  }

  function test_WhenLimitIsSetToUint32Max() external whenCallerISTheOwner {
    // it correctly sets the limit
    //     call setDONOverride(donLabelA, 4294967295)
    //     getMaxWorkflowsPerDon(donLabelA) returns 4294967295
  }

  function test_WhenCallingTheFunctionMultipleTimes() external whenCallerISTheOwner {
    // it correctly sets the latest value
    //     call setDONOverride(donLabelA, 10)
    //     getMaxWorkflowsPerDon(donLabelA) returns 10
    //     call setDONOverride(donLabelA, 20)
    //     getMaxWorkflowsPerDon(donLabelA) returns 20 (updated)
  }

  function test_WhenThereAreMultipleDONs() external whenCallerISTheOwner {
    // it should set the correct value for each DON
    //     setDONOverride(donLabelA, 77)
    //     getMaxWorkflowsPerDon(donLabelA) returns 77
    //     getMaxWorkflowsPerDon(donLabelB) returns default (500)
  }
}
