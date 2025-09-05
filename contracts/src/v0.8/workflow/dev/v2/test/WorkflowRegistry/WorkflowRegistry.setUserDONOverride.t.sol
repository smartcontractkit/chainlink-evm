// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";
import {Vm} from "forge-std/Test.sol";

contract WorkflowRegistry_setUserDONOverride is WorkflowRegistrySetup {
  function test_setUserDONOverride_WhenTheCallerIsNOTTheContractOwner() external {
    // it should revert with caller is not the owner
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector, s_stranger));
    s_registry.setUserDONOverride(s_stranger, s_donFamily, 100, true);
  }

  // whenTheCallerISTheContractOwner
  // whenEnabledIsTrue
  function test_setUserDONOverride_WhenGlobalDONLimitIsNotSet() external {
    // It should revert with DonLimitNotSet
    vm.startPrank(s_owner);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.DonLimitNotSet.selector, s_donFamily));
    s_registry.setUserDONOverride(s_user, s_donFamily, 5, true);
  }

  // whenTheCallerISTheContractOwner
  // whenEnabledIsTrue
  // whenLimitLessThanOrEqualToGlobalDONLimit
  function test_setUserDONOverride_WhenNoPriorOverrideExistsForUserDonLabel() external {
    // It should set s_cfg.userDONOverride[user][donHash] = ConfigValue(limit, true) and emit UserDONLimitSet
    vm.startPrank(s_owner);

    // set a DON limit first, otherwise the global limit is 0
    s_registry.setDONLimit(s_donFamily, 100, true);

    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.UserDONLimitSet(s_user, s_donFamily, 5);
    s_registry.setUserDONOverride(s_user, s_donFamily, 5, true);

    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 5);
    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner
  // whenEnabledIsTrue
  // whenLimitLessThanOrEqualToGlobalDONLimit
  // whenAPriorOverrideExists
  function test_setUserDONOverride_WhenNewLimitDoesNotEqualExistingOverrideValue() external {
    // It should overwrite the override and emit UserDONLimitSet

    vm.startPrank(s_owner);
    // set a DON limit first, otherwise the global limit is 0
    s_registry.setDONLimit(s_donFamily, 100, true);
    // set a limit first
    s_registry.setUserDONOverride(s_user, s_donFamily, 5, true);
    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 5);

    // set a different limit again
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.UserDONLimitSet(s_user, s_donFamily, 20);
    s_registry.setUserDONOverride(s_user, s_donFamily, 20, true);
    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 20);

    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner
  // whenEnabledIsTrue
  // whenLimitLessThanOrEqualToGlobalDONLimit
  // whenAPriorOverrideExists
  function test_setUserDONOverride_WhenNewLimitIsEqualToExistingOverrideValue() external {
    // It should do nothing
    vm.startPrank(s_owner);
    // set a DON limit first, otherwise the global limit is 0
    s_registry.setDONLimit(s_donFamily, 100, true);
    // set a limit first
    s_registry.setUserDONOverride(s_user, s_donFamily, 5, true);
    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 5);

    // set the same limit again
    // start recording all logs
    vm.recordLogs();

    s_registry.setUserDONOverride(s_user, s_donFamily, 5, true);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("UserDONLimitSet(address,string,uint32)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("UserDONLimitSet was emitted when it should not have been");
        fail();
      }
    }

    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 5);

    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner whenEnabledIsTrue
  function test_setUserDONOverride_WhenLimitIsGreaterThanGlobalDONLimit() external {
    // It should revert with UserDONOverrideExceedsDONLimit

    vm.startPrank(s_owner);
    // set a DON limit first, otherwise the global limit is 0
    s_registry.setDONLimit(s_donFamily, 100, true);
    // set a limit greater than DON
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.UserDONOverrideExceedsDONLimit.selector));
    s_registry.setUserDONOverride(s_user, s_donFamily, 200, true);

    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner whenEnabledIsFalse
  function test_setUserDONOverride_WhenAPriorOverrideExistsForUserDonLabel() external {
    // It should delete s_cfg.userDONOverride[user][donHash] and emit UserDONLimitUnset
    vm.startPrank(s_owner);
    // set a DON limit first, otherwise the global limit is 0
    s_registry.setDONLimit(s_donFamily, 100, true);
    // set a limit
    s_registry.setUserDONOverride(s_user, s_donFamily, 5, true);
    // remove the limit
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.UserDONLimitUnset(s_user, s_donFamily);
    s_registry.setUserDONOverride(s_user, s_donFamily, 5, false);

    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 100); // global don value
    vm.stopPrank();
  }

  // whenTheCallerISTheContractOwner whenEnabledIsFalse
  function test_setUserDONOverride_WhenNoPriorOverrideExists() external {
    // It should do nothing
    vm.startPrank(s_owner);
    // set a DON limit first, otherwise the global limit is 0
    s_registry.setDONLimit(s_donFamily, 100, true);
    // remove the limit for a user that doesn't have a limit
    // start recording all logs
    vm.recordLogs();

    s_registry.setUserDONOverride(s_user, s_donFamily, 5, false);

    Vm.Log[] memory entries = vm.getRecordedLogs();
    bytes32 sig = keccak256("UserDONLimitUnset(address,string)");
    for (uint256 i = 0; i < entries.length; i++) {
      if (entries[i].topics[0] == sig) {
        emit log("UserDONLimitSet was emitted when it should not have been");
        fail();
      }
    }

    assertEq(s_registry.getMaxWorkflowsPerUserDON(s_user, s_donFamily), 100); // global don value

    vm.stopPrank();
  }
}
