// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {CapabilitiesRegistry} from "../../CapabilitiesRegistry.sol";
import {BaseTest} from "./BaseTest.t.sol";
import "forge-std/Vm.sol";

contract CapabilitiesRegistry_SetDONFamilyTest is BaseTest {
  string internal constant FAMILY_NAME_ONE = "production-mainnet";
  string internal constant FAMILY_NAME_TWO = "production-testnet";

  function setUp() public override {
    BaseTest.setUp();

    // Add capabilities
    CapabilitiesRegistry.Capability[] memory capabilities = new CapabilitiesRegistry.Capability[](1);
    capabilities[0] = s_basicCapability;
    s_CapabilitiesRegistry.addCapabilities(capabilities);

    // Add node operators
    s_CapabilitiesRegistry.addNodeOperators(_getNodeOperators());

    // Add nodes
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](3);
    bytes32[] memory capabilityIds = new bytes32[](1);
    capabilityIds[0] = s_basicHashedCapabilityId;

    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      hashedCapabilityIds: capabilityIds
    });

    nodes[1] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_TWO_ID,
      p2pId: P2P_ID_TWO,
      signer: NODE_OPERATOR_TWO_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_TWO,
      hashedCapabilityIds: capabilityIds
    });

    nodes[2] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_THREE_ID,
      p2pId: P2P_ID_THREE,
      signer: NODE_OPERATOR_THREE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_THREE,
      hashedCapabilityIds: capabilityIds
    });

    s_CapabilitiesRegistry.addNodes(nodes);

    // Add DONs
    bytes32[] memory don1Nodes = new bytes32[](2);
    don1Nodes[0] = P2P_ID;
    don1Nodes[1] = P2P_ID_TWO;

    bytes32[] memory don2Nodes = new bytes32[](2);
    don2Nodes[0] = P2P_ID_TWO;
    don2Nodes[1] = P2P_ID_THREE;

    CapabilitiesRegistry.CapabilityConfiguration[] memory capabilityConfigs =
      new CapabilitiesRegistry.CapabilityConfiguration[](1);
    capabilityConfigs[0] = CapabilitiesRegistry.CapabilityConfiguration({
      capabilityId: s_basicHashedCapabilityId,
      config: BASIC_CAPABILITY_CONFIG
    });

    s_CapabilitiesRegistry.addDON(don1Nodes, capabilityConfigs, true, false, F_VALUE);
    s_CapabilitiesRegistry.addDON(don2Nodes, capabilityConfigs, true, false, F_VALUE);

    vm.startPrank(ADMIN);
  }

  function test_RevertWhen_CalledByNonOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);
  }

  function test_RevertWhen_DONDoesNotExist() public {
    uint32 nonExistentDONId = 999;
    vm.expectRevert(abi.encodeWithSelector(CapabilitiesRegistry.DONDoesNotExist.selector, nonExistentDONId));
    s_CapabilitiesRegistry.setDONFamily(nonExistentDONId, FAMILY_NAME_ONE);
  }

  function test_SetDONFamily_EmitsEvent() public {
    vm.expectEmit(true, true, false, true);
    emit CapabilitiesRegistry.DONFamilySet(DON_ID, FAMILY_NAME_ONE);

    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);
  }

  function test_SetDONFamily_MultipleDONsInSameFamily() public {
    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);
    s_CapabilitiesRegistry.setDONFamily(DON_ID_TWO, FAMILY_NAME_ONE);

    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_ONE);
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
    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);

    uint256[] memory family1DONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_ONE);
    assertEq(family1DONs.length, 1);
    assertEq(family1DONs[0], DON_ID);

    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_TWO);

    uint256[] memory family2DONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_TWO);
    assertEq(family2DONs.length, 1);
    assertEq(family2DONs[0], DON_ID);

    family1DONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_ONE);
    assertEq(family1DONs.length, 0);

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamily, FAMILY_NAME_TWO);
  }

  function test_SetDONFamily_RemoveFromFamily() public {
    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);

    vm.expectEmit(true, true, false, true);
    emit CapabilitiesRegistry.DONFamilySet(DON_ID, "");
    s_CapabilitiesRegistry.setDONFamily(DON_ID, "");

    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_ONE);
    assertEq(familyDONs.length, 0);

    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamily, "");
  }

  function test_SetDONFamily_SameFamilyNoOp() public {
    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);

    // Try to set to the same family again - should be a no-op
    // The function should return early without emitting an event
    vm.recordLogs();
    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);

    // Check that no DONFamilySet event was emitted for the second call
    Vm.Log[] memory entries = vm.getRecordedLogs();
    // Should have no logs since it's a no-op
    assertEq(entries.length, 0);

    // Verify DON is still in the family
    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_ONE);
    assertEq(familyDONs.length, 1);
    assertEq(familyDONs[0], DON_ID);
  }

  function test_GetDONsInFamily_EmptyFamily() public view {
    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily("non-existent-family");
    assertEq(familyDONs.length, 0);
  }

  function test_DONInfo_IncludesFamilyInformation() public {
    CapabilitiesRegistry.DONInfo memory donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamily, "");

    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);

    donInfo = s_CapabilitiesRegistry.getDON(DON_ID);
    assertEq(donInfo.donFamily, FAMILY_NAME_ONE);
  }

  function test_FamilyCleanupOnDONRemoval() public {
    s_CapabilitiesRegistry.setDONFamily(DON_ID, FAMILY_NAME_ONE);
    s_CapabilitiesRegistry.setDONFamily(DON_ID_TWO, FAMILY_NAME_ONE);

    uint256[] memory familyDONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_ONE);
    assertEq(familyDONs.length, 2);

    uint32[] memory donsToRemove = new uint32[](1);
    donsToRemove[0] = DON_ID;
    s_CapabilitiesRegistry.removeDONs(donsToRemove);

    familyDONs = s_CapabilitiesRegistry.getDONsInFamily(FAMILY_NAME_ONE);
    assertEq(familyDONs.length, 1);
    assertEq(familyDONs[0], DON_ID_TWO);
  }
}
