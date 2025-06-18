// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Ownable2Step} from "../../../../../shared/access/Ownable2Step.sol";
import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";
import {Vm} from "forge-std/Vm.sol";

contract CapabilitiesRegistry_SetDONFamilyTest is BaseTest {
  function setUp() public override {
    BaseTest.setUp();

    s_CapabilitiesRegistry.addCapabilities(s_capabilities);
    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());
    s_CapabilitiesRegistry.addNodes(s_paramsForTwoNodes);
    s_CapabilitiesRegistry.addDONs(s_paramsForTwoDONs);

    vm.startPrank(ADMIN);
  }

  function test_RevertWhen_CalledByNonOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(Ownable2Step.OnlyCallableByOwner.selector));
    string[] memory addToFamilies = new string[](1);
    addToFamilies[0] = TEST_DON_FAMILY_ONE;
    string[] memory removeFromFamilies = new string[](0);
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);
  }

  function test_RevertWhen_DONDoesNotExist() public {
    uint32 nonExistentDONId = 999;
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONDoesNotExist.selector, nonExistentDONId));
    string[] memory addToFamilies = new string[](1);
    addToFamilies[0] = TEST_DON_FAMILY_ONE;
    string[] memory removeFromFamilies = new string[](0);
    s_CapabilitiesRegistry.setDONFamilies(nonExistentDONId, addToFamilies, removeFromFamilies);
  }

  function test_SetDONFamily_EmitsEvent() public {
    vm.expectEmit(true, true, false, true);
    emit CapabilitiesRegistry.DONAddedToFamily(DON_ID, TEST_DON_FAMILY_TWO);
    vm.expectEmit(true, true, false, true);
    emit CapabilitiesRegistry.DONRemovedFromFamily(DON_ID, TEST_DON_FAMILY_ONE);

    string[] memory addToFamilies = new string[](1);
    addToFamilies[0] = TEST_DON_FAMILY_TWO;
    string[] memory removeFromFamilies = new string[](1);
    removeFromFamilies[0] = TEST_DON_FAMILY_ONE;
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);
  }

  function test_SetDONFamily_MultipleDONsInSameFamily() public {
    string[] memory addToFamilies = new string[](1);
    addToFamilies[0] = TEST_DON_FAMILY_ONE;
    string[] memory removeFromFamilies = new string[](0);
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);
    s_CapabilitiesRegistry.setDONFamilies(DON_ID_TWO, addToFamilies, removeFromFamilies);

    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(familyDONs.length, 2);

    // DONs could be in any order, so check both are present
    bool foundDON1 = false;
    bool foundDON2 = false;
    for (uint256 i = 0; i < familyDONs.length; i++) {
      if (familyDONs[i] == DON_ID) foundDON1 = true;
      if (familyDONs[i] == DON_ID_TWO) foundDON2 = true;
    }
    assertTrue(foundDON1);
    assertTrue(foundDON2);
  }

  function test_SetDONFamily_MoveDONBetweenFamilies() public {
    string[] memory addToFamilies = new string[](1);
    addToFamilies[0] = TEST_DON_FAMILY_TWO;
    string[] memory removeFromFamilies = new string[](1);
    removeFromFamilies[0] = TEST_DON_FAMILY_ONE;
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);

    uint256[] memory family1DONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(family1DONs.length, 1, "Expected 1 DON in TEST_DON_FAMILY_ONE");
    assertEq(family1DONs[0], DON_ID_TWO, "Expected DON_ID_TWO to be in TEST_DON_FAMILY_ONE");

    addToFamilies[0] = TEST_DON_FAMILY_TWO;
    removeFromFamilies = new string[](1);
    removeFromFamilies[0] = TEST_DON_FAMILY_ONE;
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);

    uint256[] memory family2DONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_TWO);
    assertEq(family2DONs.length, 1, "Expected 1 DON in TEST_DON_FAMILY_TWO");
    assertEq(family2DONs[0], DON_ID, "Expected DON_ID to be in TEST_DON_FAMILY_TWO");

    family1DONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(family1DONs.length, 1, "Expected 1 DON in TEST_DON_FAMILY_ONE");
    assertEq(family1DONs[0], DON_ID_TWO, "Expected DON_ID_TWO to be in TEST_DON_FAMILY_ONE");

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamilies.length, 1, "Expected 1 family");
    assertEq(donInfo.donFamilies[0], TEST_DON_FAMILY_TWO, "Expected family to be TEST_DON_FAMILY_TWO");
  }

  function test_SetDONFamily_RemoveFromFamily() public {
    vm.expectEmit(true, true, false, true);
    emit CapabilitiesRegistry.DONRemovedFromFamily(DON_ID, TEST_DON_FAMILY_ONE);
    string[] memory addToFamilies = new string[](0);
    string[] memory removeFromFamilies = new string[](1);
    removeFromFamilies[0] = TEST_DON_FAMILY_ONE;
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamilies.length, 0, "Expected DON_ID to not be in any family");

    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(familyDONs.length, 1, "Expected 1 DON in TEST_DON_FAMILY_ONE");
    assertEq(familyDONs[0], DON_ID_TWO, "Expected DON_ID_TWO to be in TEST_DON_FAMILY_ONE");
  }

  function test_SetDONFamily_SameFamilyNoOp() public {
    string[] memory addToFamilies = new string[](1);
    addToFamilies[0] = TEST_DON_FAMILY_ONE;
    string[] memory removeFromFamilies = new string[](0);
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);

    // Try to set to the same family again - should be a no-op
    // The function should return early without emitting an event
    vm.recordLogs();
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);

    // Check that no DONAddedToFamily event was emitted for the second call
    Vm.Log[] memory entries = vm.getRecordedLogs();
    // Should have no logs since it's a no-op
    assertEq(entries.length, 0, "Expected no event logs");

    // Verify DON is still in the family
    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(familyDONs.length, 2, "Expected 2 DONs in family");
    assertEq(familyDONs[0], DON_ID, "Expected DON_ID to be in family");
    assertEq(familyDONs[1], DON_ID_TWO, "Expected DON_ID_TWO to be in family");
  }

  function test_GetDONsInFamily_EmptyFamily() public view {
    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily("non-existent-family");
    assertEq(familyDONs.length, 0);
  }

  function test_SetDONFamilies_RemoveFromAllFamilies() public {
    string[] memory addToFamilies = new string[](0);
    string[] memory removeFromFamilies = new string[](1);
    removeFromFamilies[0] = TEST_DON_FAMILY_ONE;
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamilies.length, 0, "Expected DON_ID to not be in any families");
  }

  function test_SetDONFamilies_AddToMultipleFamilies() public {
    string[] memory addToFamilies = new string[](2);
    addToFamilies[0] = TEST_DON_FAMILY_ONE;
    addToFamilies[1] = TEST_DON_FAMILY_TWO;
    string[] memory removeFromFamilies = new string[](0);
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamilies.length, 2, "Expected DON_ID to be in 2 families");
    assertEq(donInfo.donFamilies[0], TEST_DON_FAMILY_ONE, "Expected DON_ID to be in TEST_DON_FAMILY_ONE");
    assertEq(donInfo.donFamilies[1], TEST_DON_FAMILY_TWO, "Expected DON_ID to be in TEST_DON_FAMILY_TWO");
  }

  function test_DONInfo_IncludesFamilyInformation() public view {
    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamilies.length, 1, "Expected 1 family");
    assertEq(donInfo.donFamilies[0], TEST_DON_FAMILY_ONE, "Expected family to be TEST_DON_FAMILY_ONE");
  }

  function test_FamilyCleanupOnDONRemoval() public {
    string[] memory addToFamilies = new string[](1);
    addToFamilies[0] = TEST_DON_FAMILY_ONE;
    string[] memory removeFromFamilies = new string[](0);
    s_CapabilitiesRegistry.setDONFamilies(DON_ID, addToFamilies, removeFromFamilies);
    s_CapabilitiesRegistry.setDONFamilies(DON_ID_TWO, addToFamilies, removeFromFamilies);

    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(familyDONs.length, 2);

    uint32[] memory donsToRemove = new uint32[](1);
    donsToRemove[0] = DON_ID;
    s_CapabilitiesRegistry.removeDONs(donsToRemove);

    familyDONs = s_CapabilitiesRegistry.getDONsInFamily(TEST_DON_FAMILY_ONE);
    assertEq(familyDONs.length, 1);
    assertEq(familyDONs[0], DON_ID_TWO);
  }
}
