// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_updateAllowedSigners is WorkflowRegistrySetup {
  function setUp() public override {
    super.setUp();
    vm.prank(s_owner);
    address[] memory signers = new address[](3);
    signers[0] = address(0x1111);
    signers[1] = address(0x2222);
    signers[2] = address(0x3333);

    s_registry.updateAllowedSigners(signers, true);
    assertTrue(s_registry.isAllowedSigner(address(0x1111)), "Signer 1 should be added");
    assertTrue(s_registry.isAllowedSigner(address(0x2222)), "Signer 2 should be added");
    assertTrue(s_registry.isAllowedSigner(address(0x3333)), "Signer 3 should be added");
  }

  function test_updateAllowedSigners_ShouldOnlyBeCalledByTheContractOwner() external {
    // it should only be called by the contract s_owner
    vm.prank(s_stranger);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    s_registry.updateAllowedSigners(new address[](0), true);
  }

  modifier whenANewSignerIsAdded() {
    _;
  }

  function test_updateAllowedSigners_GivenSignerIsNotAlreadyAdded() external whenANewSignerIsAdded {
    // it should update the allowed signers
    vm.prank(s_owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0xaaaa);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdated(signers, true);

    s_registry.updateAllowedSigners(signers, true);
    assertTrue(s_registry.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0xaaaa)), "New signer should be added");
  }

  function test_updateAllowedSigners_GivenTheSignerIsAlreadyAdded() external whenANewSignerIsAdded {
    // it should not have any effect
    vm.prank(s_owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0x2222);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdated(signers, true);

    s_registry.updateAllowedSigners(signers, true);
    assertTrue(s_registry.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
  }

  modifier whenAnExistingSignerIsRemoved() {
    _;
  }

  function test_updateAllowedSigners_GivenTheSignerIsNotAlreadyRemoved() external whenAnExistingSignerIsRemoved {
    // it should update the allowed signers
    vm.prank(s_owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0x2222);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdated(signers, false);

    s_registry.updateAllowedSigners(signers, false);
    assertTrue(s_registry.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertFalse(s_registry.isAllowedSigner(address(0x2222)), "Signer 2 should be removed");
    assertTrue(s_registry.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
  }

  function test_updateAllowedSigners_GivenTheSignerIsAlreadyRemoved() external whenAnExistingSignerIsRemoved {
    // it should not have any effect
    vm.prank(s_owner);
    address[] memory signers = new address[](1);
    // this signer was never added in the first place
    signers[0] = address(0x5555);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdated(signers, false);

    s_registry.updateAllowedSigners(signers, false);
    assertTrue(s_registry.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
    assertFalse(s_registry.isAllowedSigner(address(0x5555)), "New signer should not be on the list");
  }

  function test_updateAllowedSigners_WhenTheSignerIsTheZeroAddress() external {
    // it should revert with an error
    vm.prank(s_owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0x0);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.ZeroAddressNotAllowed.selector));
    s_registry.updateAllowedSigners(signers, true);
    assertTrue(s_registry.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(s_registry.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
  }
}
