// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_limitsScenario is WorkflowRegistrySetup {
  // Helper to build deterministic IDs
  function _wfId(string memory prefix, uint256 i) internal pure returns (bytes32) {
    return keccak256(abi.encodePacked(prefix, "-", i));
  }

  // Scenario:
  // Global DON limit = 10
  // Default per-user = 2 (applies to userA)
  // userB override = 5
  // userC override = 3
  // sum of (2 + 5 + 3) = global (10)
  // Ensure global cap is enforced regardless of individual per-user limits.
  function test_limitsScenario_whenGlobalLimitEqualsSumOfPerUserLimits() external {
    // Users
    address userA = s_user; // Will use default per-user limit
    address userB = makeAddr("userB"); // Has override
    address userC = makeAddr("userC"); // Has override

    // Link all as owners
    _linkOwner(userA);
    _linkOwner(userB);
    _linkOwner(userC);

    // Set DON config:
    //   Global DON active workflow cap = 10
    //   Default per-user cap           = 2
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 10, 2);

    // Set user overrides:
    // userB -> 5, userC -> 3 (total theoretical = 2 + 5 + 3 = 10 == global)
    vm.startPrank(s_owner);
    s_registry.setUserDONOverride(userB, s_donFamily, 5, true);
    s_registry.setUserDONOverride(userC, s_donFamily, 3, true);
    vm.stopPrank();

    // -----------------------------
    // User A (default limit = 2)
    // -----------------------------
    vm.startPrank(userA);
    for (uint256 i = 0; i < 2; ++i) {
      s_registry.upsertWorkflow(
        string.concat("A-", vm.toString(i)),
        s_tag,
        _wfId("A", i),
        WorkflowRegistry.WorkflowStatus.ACTIVE,
        s_donFamily,
        s_binaryUrl,
        s_configUrl,
        s_attributes,
        false
      );
    }

    // Third ACTIVE should exceed per-user default limit
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerUserDONExceeded.selector, userA, s_donFamily)
    );
    s_registry.upsertWorkflow(
      "A-2",
      s_tag,
      _wfId("A", 2),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.stopPrank();

    // -----------------------------
    // User B (override = 5)
    // -----------------------------
    vm.startPrank(userB);
    for (uint256 i = 0; i < 5; ++i) {
      s_registry.upsertWorkflow(
        string.concat("B-", vm.toString(i)),
        s_tag,
        _wfId("B", i),
        WorkflowRegistry.WorkflowStatus.ACTIVE,
        s_donFamily,
        s_binaryUrl,
        s_configUrl,
        s_attributes,
        false
      );
    }
    // Sixth should exceed override
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerUserDONExceeded.selector, userB, s_donFamily)
    );
    s_registry.upsertWorkflow(
      "B-5",
      s_tag,
      _wfId("B", 5),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.stopPrank();

    // -----------------------------
    // User C (override = 3)
    // -----------------------------
    vm.startPrank(userC);
    for (uint256 i = 0; i < 3; ++i) {
      s_registry.upsertWorkflow(
        string.concat("C-", vm.toString(i)),
        s_tag,
        _wfId("C", i),
        WorkflowRegistry.WorkflowStatus.ACTIVE,
        s_donFamily,
        s_binaryUrl,
        s_configUrl,
        s_attributes,
        false
      );
    }
    // Fourth should exceed override
    // Fourth attempt hits GLOBAL limit first (2 + 5 + 3 = 10 already)
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "C-3",
      s_tag,
      _wfId("C", 3),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.stopPrank();

    // At this point total ACTIVE = 2 (A) + 5 (B) + 3 (C) = 10 (global cap reached)

    // Any attempt by any user to add another ACTIVE (even if per-user still allows, hypothetically)
    // should revert with the global cap error.
    // If your contract uses a different selector/name for the global cap, adjust below.
    vm.startPrank(userB);
    // First, add a PAUSED workflow (should NOT count toward global active cap)
    s_registry.upsertWorkflow(
      "B-paused-extra",
      s_tag,
      _wfId("B", 100),
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    // Now attempt to insert another ACTIVE (should hit global cap)
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "B-overflow",
      s_tag,
      _wfId("B", 101),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.stopPrank();

    // Free capacity by pausing one ACTIVE workflow (B-0), then activate the previously PAUSED one (B-paused-extra),
    // using dedicated pause/activate functions instead of re-upserting.
    vm.startPrank(userB);

    // Pause existing ACTIVE workflow B-0
    s_registry.pauseWorkflow(_wfId("B", 0));

    // Activate previously PAUSED workflow B-paused-extra
    s_registry.activateWorkflow(_wfId("B", 100), s_donFamily);

    // Capacity is full again; another ACTIVE should revert (global cap)
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "B-overflow-2",
      s_tag,
      _wfId("B", 102),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    vm.stopPrank();
  }

  // Scenario:
  // Global DON limit = 8
  // Default per-user = 2 (applies to userA)
  // userB override = 6
  // userC override = 6
  // sum of (2 + 6 + 6) > global (8)
  // Ensure global cap is enforced even when individual overrides would allow more.
  function test_limitsScenario_whenGlobalLimitLowerThanSumOfPerUserLimits() external {
    // Users
    address userA = s_user;
    address userB = makeAddr("userB");
    address userC = makeAddr("userC");

    _linkOwner(userA);
    _linkOwner(userB);
    _linkOwner(userC);

    // Set global + default
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 8, 2);

    // Set large overrides (each could exceed remaining capacity alone)
    vm.startPrank(s_owner);
    s_registry.setUserDONOverride(userB, s_donFamily, 6, true);
    s_registry.setUserDONOverride(userC, s_donFamily, 6, true);
    vm.stopPrank();

    // -----------------------------
    // userA uses default (2)
    // -----------------------------
    vm.startPrank(userA);
    for (uint256 i = 0; i < 2; ++i) {
      s_registry.upsertWorkflow(
        string.concat("A-", vm.toString(i)),
        s_tag,
        _wfId("A", i),
        WorkflowRegistry.WorkflowStatus.ACTIVE,
        s_donFamily,
        s_binaryUrl,
        s_configUrl,
        s_attributes,
        false
      );
    }
    vm.stopPrank();
    // Active total: 2

    // -----------------------------
    // userB tries to fill up to override (6) but global will cap overall at 8
    // -----------------------------
    vm.startPrank(userB);
    for (uint256 i = 0; i < 6; ++i) {
      // First 6 from userB should all succeed while total <= 8
      // i = 0..5 => adds 6 more -> total becomes 8 exactly after i=5
      s_registry.upsertWorkflow(
        string.concat("B-", vm.toString(i)),
        s_tag,
        _wfId("B", i),
        WorkflowRegistry.WorkflowStatus.ACTIVE,
        s_donFamily,
        s_binaryUrl,
        s_configUrl,
        s_attributes,
        false
      );
    }
    // Active total: 2 + 6 = 8 (global full)
    vm.stopPrank();

    // -----------------------------
    // userC: Even first ACTIVE should now fail (global cap reached)
    // -----------------------------
    vm.startPrank(userC);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "C-0",
      s_tag,
      _wfId("C", 0),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    // But userC can still register PAUSED (should not count toward active cap)
    s_registry.upsertWorkflow(
      "C-paused",
      s_tag,
      _wfId("C", 100),
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.stopPrank();

    // -----------------------------
    // Free one slot: pause one of userB's ACTIVE workflows
    // -----------------------------
    vm.prank(userB);
    s_registry.pauseWorkflow(_wfId("B", 0)); // Active total: 7 now

    // Now userC can activate (was paused) exactly one workflow
    vm.prank(userC);
    s_registry.activateWorkflow(_wfId("C", 100), s_donFamily); // Active total: 8 again

    // Trying to activate/create another ACTIVE for userC should revert (global full),
    // even though its override (6) is far from reached (userC has only 1 active).
    vm.startPrank(userC);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "C-1",
      s_tag,
      _wfId("C", 1),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );

    // Also attempting to activate another paused workflow should fail (create another paused first)
    s_registry.upsertWorkflow(
      "C-paused-2",
      s_tag,
      _wfId("C", 101),
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.activateWorkflow(_wfId("C", 101), s_donFamily);
    vm.stopPrank();

    // If userB tries to add another ACTIVE, should also fail (global full)
    vm.startPrank(userB);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "B-6",
      s_tag,
      _wfId("B", 6),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.stopPrank();

    // If userA tries to add another ACTIVE, should also fail (global full)
    vm.startPrank(userA);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "A-2",
      s_tag,
      _wfId("A", 2),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
    vm.stopPrank();
  }
}
