// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2Step} from "../../shared/access/Ownable2Step.sol";
import {DonIDClaimer} from "../DonIDClaimer.sol";
import {MockCapabilitiesRegistry} from "./mocks/MockCapabilitiesRegistry.sol";
import {Test} from "forge-std/Test.sol";

contract DonIDClaimerTest is Test {
  uint32 internal constant INITIAL_CLAIM_ID = 100;
  DonIDClaimer private s_donIDClaimer;
  MockCapabilitiesRegistry private s_mockRegistry;
  address private s_owner = address(0x1);
  address private s_deployer = address(0x2);
  address private s_unauthorized = address(0x3);

  function setUp() public {
    vm.startPrank(s_owner);
    s_mockRegistry = new MockCapabilitiesRegistry(INITIAL_CLAIM_ID);
    s_donIDClaimer = new DonIDClaimer(address(s_mockRegistry));
    s_donIDClaimer.setAuthorizedDeployer(s_deployer, true);
    vm.stopPrank();
  }

  function test_Constructor() public {
    // Check the revert if constructor is called with a zero address
    vm.expectRevert(abi.encodeWithSelector(DonIDClaimer.ZeroAddressNotAllowed.selector));
    new DonIDClaimer(address(0));

    // Now test the normal constructor behavior with a valid address
    DonIDClaimer validDonIDClaimer = new DonIDClaimer(address(s_mockRegistry));
    assertEq(
      validDonIDClaimer.getNextDONId(), INITIAL_CLAIM_ID, "Initial DON ID should be set correctly from the registry"
    );
  }

  function test_ClaimNextDONId() public {
    vm.prank(s_deployer);
    vm.expectEmit();
    emit DonIDClaimer.DonIDClaimed(s_deployer, INITIAL_CLAIM_ID);

    uint32 claimedId = s_donIDClaimer.claimNextDONId();
    assertEq(claimedId, INITIAL_CLAIM_ID, "Claimed DON ID should be 100");
    assertEq(s_donIDClaimer.getNextDONId(), INITIAL_CLAIM_ID + 1, "Next DON ID should be incremented to 101");
  }

  function test_SyncNextDONIdWithOffset() public {
    vm.expectEmit();
    emit DonIDClaimer.DonIDSynced(INITIAL_CLAIM_ID + 10);

    vm.prank(s_deployer);
    s_donIDClaimer.syncNextDONIdWithOffset(10);
    assertEq(s_donIDClaimer.getNextDONId(), INITIAL_CLAIM_ID + 10, "Next DON ID should be 110 after offset");
  }

  function test_SetAuthorizedDeployer() public {
    vm.expectEmit();
    emit DonIDClaimer.AuthorizedDeployerSet(s_unauthorized, true);

    vm.prank(s_owner);
    s_donIDClaimer.setAuthorizedDeployer(s_unauthorized, true);
    assertTrue(s_donIDClaimer.isAuthorizedDeployer(s_unauthorized), "Address should be authorized");
  }

  // Reverts
  function test_RevertWhen_UnauthorizedSenderClaimReverts() public {
    vm.expectRevert(abi.encodeWithSelector(DonIDClaimer.AccessForbidden.selector, s_unauthorized));
    vm.prank(s_unauthorized);
    s_donIDClaimer.claimNextDONId();
  }

  function test_RevertWhen_UnauthorizedSetAuthorizedDeployer() public {
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    vm.prank(s_unauthorized);
    s_donIDClaimer.setAuthorizedDeployer(s_unauthorized, true);
  }
}
