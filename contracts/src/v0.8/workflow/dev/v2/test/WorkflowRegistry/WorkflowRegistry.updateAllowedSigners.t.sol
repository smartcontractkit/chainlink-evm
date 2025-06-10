// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {Test} from "forge-std/Test.sol";

contract WorkflowRegistry_updateAllowedSigners is Test {
  WorkflowRegistry public wr;
  address public owner = address(0xabcd);

  function setUp() external {
    vm.prank(owner);
    wr = new WorkflowRegistry();

    vm.prank(owner);
    address[] memory signers = new address[](3);
    signers[0] = address(0x1111);
    signers[1] = address(0x2222);
    signers[2] = address(0x3333);

    wr.updateAllowedSigners(signers, true);
    assertTrue(wr.isAllowedSigner(address(0x1111)), "Signer 1 should be added");
    assertTrue(wr.isAllowedSigner(address(0x2222)), "Signer 2 should be added");
    assertTrue(wr.isAllowedSigner(address(0x3333)), "Signer 3 should be added");
  }

  function test_updateAllowedSigners_ShouldOnlyBeCalledByTheContractOwner() external {
    // it should only be called by the contract owner
    address notOwner = address(0x1234);
    vm.prank(notOwner);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    wr.updateAllowedSigners(new address[](0), true);
  }

  modifier whenANewSignerIsAdded() {
    _;
  }

  function test_updateAllowedSigners_GivenSignerIsNotAlreadyAdded() external whenANewSignerIsAdded {
    // it should update the allowed signers
    vm.prank(owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0xaaaa);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdatedV1(signers, true);

    wr.updateAllowedSigners(signers, true);
    assertTrue(wr.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
    assertTrue(wr.isAllowedSigner(address(0xaaaa)), "New signer should be added");
  }

  function test_updateAllowedSigners_GivenTheSignerIsAlreadyAdded() external whenANewSignerIsAdded {
    // it should not have any effect
    vm.prank(owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0x2222);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdatedV1(signers, true);

    wr.updateAllowedSigners(signers, true);
    assertTrue(wr.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
  }

  modifier whenAnExistingSignerIsRemoved() {
    _;
  }

  function test_updateAllowedSigners_GivenTheSignerIsNotAlreadyRemoved() external whenAnExistingSignerIsRemoved {
    // it should update the allowed signers
    vm.prank(owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0x2222);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdatedV1(signers, false);

    wr.updateAllowedSigners(signers, false);
    assertTrue(wr.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertFalse(wr.isAllowedSigner(address(0x2222)), "Signer 2 should be removed");
    assertTrue(wr.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
  }

  function test_updateAllowedSigners_GivenTheSignerIsAlreadyRemoved() external whenAnExistingSignerIsRemoved {
    // it should not have any effect
    vm.prank(owner);
    address[] memory signers = new address[](1);
    // this signer was never added in the first place
    signers[0] = address(0x5555);

    vm.expectEmit(true, false, false, true);
    emit WorkflowRegistry.AllowedSignersUpdatedV1(signers, false);

    wr.updateAllowedSigners(signers, false);
    assertTrue(wr.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
    assertFalse(wr.isAllowedSigner(address(0x5555)), "New signer should not be on the list");
  }

  function test_updateAllowedSigners_WhenTheSignerIsTheZeroAddress() external {
    // it should revert with an error
    vm.prank(owner);
    address[] memory signers = new address[](1);
    signers[0] = address(0x0);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.ZeroAddressNotAllowed.selector));
    wr.updateAllowedSigners(signers, true);
    assertTrue(wr.isAllowedSigner(address(0x1111)), "Signer 1 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x2222)), "Signer 2 should be still here");
    assertTrue(wr.isAllowedSigner(address(0x3333)), "Signer 3 should be still here");
  }
}
