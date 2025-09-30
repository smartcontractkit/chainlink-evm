// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";

import {Test} from "forge-std/Test.sol";

contract WorkflowRegistry_getAllowedSigners is Test {
  address internal s_owner;
  WorkflowRegistry internal s_registry;

  function setUp() public virtual {
    s_owner = makeAddr("owner");

    vm.startPrank(s_owner);
    s_registry = new WorkflowRegistry();
    address[] memory signers = new address[](5);
    signers[0] = address(0x1111);
    signers[1] = address(0x2222);
    signers[2] = address(0x3333);
    signers[3] = address(0x4444);
    signers[4] = address(0x5555);

    s_registry.updateAllowedSigners(signers, true);

    assertEq(5, s_registry.totalAllowedSigners(), "Total number of signers should be 5");
    vm.stopPrank();
  }

  function test_getAllowedSigners_WhenNoAllowedSignersAreConfigured() external {
    // it should return an empty array
  }

  modifier whenSomeAllowedSignersAreConfigured() {
    _;
  }

  function test_getAllowedSigners_WhenStartIs0AndLimitIsGreaterThanTheNumberOfSigners()
    external
    view
    whenSomeAllowedSignersAreConfigured
  {
    // it should return all of the signers
    address[] memory signers = s_registry.getAllowedSigners(0, 10);
    assertEq(signers.length, 5, "Should return all 5 signers");
    assertEq(signers[0], address(0x1111), "Signer 1 should match");
    assertEq(signers[1], address(0x2222), "Signer 2 should match");
    assertEq(signers[2], address(0x3333), "Signer 3 should match");
    assertEq(signers[3], address(0x4444), "Signer 4 should match");
    assertEq(signers[4], address(0x5555), "Signer 5 should match");
  }

  function test_getAllowedSigners_WhenAPageIsRequested() external view whenSomeAllowedSignersAreConfigured {
    // it should return a limited number of signers
    address[] memory signers = s_registry.getAllowedSigners(0, 2);
    assertEq(signers.length, 2, "Should return 2 signers");
    assertEq(signers[0], address(0x1111), "Signer 1 should match");
    assertEq(signers[1], address(0x2222), "Signer 2 should match");

    signers = s_registry.getAllowedSigners(2, 2);
    assertEq(signers.length, 2, "Should return 2 signers");
    assertEq(signers[0], address(0x3333), "Signer 3 should match");
    assertEq(signers[1], address(0x4444), "Signer 4 should match");

    signers = s_registry.getAllowedSigners(4, 2);
    assertEq(signers.length, 1, "Should return 1 signer");
    assertEq(signers[0], address(0x5555), "Signer 5 should match");

    signers = s_registry.getAllowedSigners(5, 2);
    assertEq(signers.length, 0, "Should return 0 signers");
  }
}
