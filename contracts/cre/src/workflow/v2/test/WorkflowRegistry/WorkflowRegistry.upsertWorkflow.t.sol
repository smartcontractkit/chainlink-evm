// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {WorkflowRegistrySetup} from "./WorkflowRegistrySetup.t.sol";

contract WorkflowRegistry_upsertWorkflow is WorkflowRegistrySetup {
  // whenMsgSenderNotLinked
  function test_upsertWorkflow_WhenMsgSenderNotALinkedOwner() external {
    // it should revert with OwnershipLinkDoesNotExist
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.OwnershipLinkDoesNotExist.selector, s_user));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  modifier whenMsgSenderIsALinkedOwner() {
    _linkOwner(s_user);
    _;
  }

  // whenNoExistingRecord
  function test_upsertWorkflow_WhenWorkflowNameLengthIsZero() external whenMsgSenderIsALinkedOwner {
    // it should revert with WorkflowNameRequired
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowNameRequired.selector));
    s_registry.upsertWorkflow(
      "",
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  // ================================================================
  // |                       CREATE PATH                            |
  // ================================================================
  // whenNoExistingRecordForOwnerNameTagCombo
  // whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenWorkflowNameLengthGreaterThanAllowed() external whenMsgSenderIsALinkedOwner {
    // it should revert with WorkflowNameTooLong
    uint8 cap = s_registry.getConfig().maxNameLen;
    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.WorkflowNameTooLong.selector, bytes(s_invalidLongString).length, cap)
    );
    s_registry.upsertWorkflow(
      s_invalidLongString,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenTagLengthIsZero() external whenMsgSenderIsALinkedOwner {
    // it should revert with WorkflowTagRequired
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowTagRequired.selector));
    s_registry.upsertWorkflow(
      s_workflowName,
      "",
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenTagLengthGreaterThanAllowed() external whenMsgSenderIsALinkedOwner {
    // it should revert with WorkflowTagTooLong
    uint8 cap = s_registry.getConfig().maxTagLen;
    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.WorkflowTagTooLong.selector, bytes(s_invalidLongString).length, cap)
    );
    s_registry.upsertWorkflow(
      s_workflowName,
      s_invalidLongString,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenWorkflowIdIsZero() external whenMsgSenderIsALinkedOwner {
    // it should revert with ZeroWorkflowIDNotAllowed
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.ZeroWorkflowIDNotAllowed.selector));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      bytes32(0),
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  //   whenNoExistingRecordForOwnerNameTagCombo
  //   whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenWorkflowIdAlreadyExists() external whenMsgSenderIsALinkedOwner {
    // it should revert with WorkflowIDAlreadyExists
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.PAUSED, false, s_user);
    // upser the same workflow again
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowIDAlreadyExists.selector, s_workflowId));
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.PAUSED, false, s_user);
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenBinaryUrlLengthGreaterThanAllowed() external whenMsgSenderIsALinkedOwner {
    // it should revert with URLTooLong
    uint8 cap = s_registry.getConfig().maxUrlLen;
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.URLTooLong.selector, bytes(s_invalidURL).length, cap));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_invalidURL,
      s_configUrl,
      s_attributes,
      true
    );
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenBinaryUrlIsMissing() external whenMsgSenderIsALinkedOwner {
    // it should revert with BinaryURLRequired
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.BinaryURLRequired.selector));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      "",
      s_configUrl,
      s_attributes,
      true
    );
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenConfigUrlLengthGreaterThanAllowed() external whenMsgSenderIsALinkedOwner {
    // it should revert with URLTooLong
    uint8 cap = s_registry.getConfig().maxUrlLen;
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.URLTooLong.selector, bytes(s_invalidURL).length, cap));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_invalidURL,
      s_attributes,
      true
    );
  }

  //   whenNoExistingRecordForOwnerNameTagCombo
  //   whenThereAreInvalidMetadataInputs
  function test_upsertWorkflow_WhenAttributesLengthGreaterThanAllowed() external whenMsgSenderIsALinkedOwner {
    // it should revert with AttributesTooLong
    bytes memory tooBigAttrs = new bytes(1025);
    for (uint256 i = 0; i < tooBigAttrs.length; i++) {
      tooBigAttrs[i] = bytes1(uint8(i % 256));
    }
    uint16 cap = s_registry.getConfig().maxAttrLen;
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.AttributesTooLong.selector, bytes(tooBigAttrs).length, cap));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      tooBigAttrs,
      true
    );
  }

  modifier whenDONLimtisAreSet() {
    _setDONLimit();
    _;
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenAllMetadataInputsAreValid
  function test_upsertWorkflow_WhenKeepAliveIsTrue() external whenMsgSenderIsALinkedOwner whenDONLimtisAreSet {
    // it should not pause other active versions
    // deploy a workflow first
    bytes32 workflowId1 = keccak256("workflow1");
    string memory workflowName1 = "Price Oracle";
    string memory tag1 = "oracle-main";
    string memory binaryUrl1 = "https://example.com/binaries/price-oracle.wasm";
    string memory configUrl1 = "https://example.com/configs/price-oracle.json";
    bytes memory attributes1 = abi.encode("Price Oracle v1.0");

    vm.startPrank(s_user);
    s_registry.upsertWorkflow(
      workflowName1,
      tag1,
      workflowId1,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 1, "There should be 1 workflow for the s_user");

    // deploy another workflow with the same name but a different tag
    bytes32 workflowId2 = keccak256("workflow2");
    string memory tag2 = "weather-feed";

    s_registry.upsertWorkflow(
      workflowName1,
      tag2,
      workflowId2,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
    wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 2, "There should be 2 workflows for the s_user");

    // now assert both are ACTIVE
    for (uint256 i = 0; i < wrs.length; i++) {
      assertEq(uint8(wrs[i].status), uint8(WorkflowRegistry.WorkflowStatus.ACTIVE), "workflow should be ACTIVE");
    }

    vm.stopPrank();
  }

  function test_upsertWorkflow_RevertWhenDonFamilyIsChangedOnUpdateWorkflow()
    external
    whenMsgSenderIsALinkedOwner
    whenDONLimtisAreSet
  {
    // it should revert because this is not allowed

    // deploy a workflow first
    bytes32 workflowId1 = keccak256("workflow1");
    string memory workflowName1 = "Price Oracle";
    string memory tag1 = "oracle-main";
    string memory binaryUrl1 = "https://example.com/binaries/price-oracle.wasm";
    string memory configUrl1 = "https://example.com/configs/price-oracle.json";
    bytes memory attributes1 = abi.encode("Price Oracle v1.0");

    vm.prank(s_user);
    s_registry.upsertWorkflow(
      workflowName1,
      tag1,
      workflowId1,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 1, "There should be 1 workflow for the s_user");

    vm.prank(s_owner);
    string memory differentDonFamily = "different-don";
    s_registry.setDONLimit(differentDonFamily, 10, 5);

    bytes32 workflowId2 = keccak256("workflow2");
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.CannotChangeDONFamilyOnUpdate.selector,
        workflowId2,
        s_user,
        workflowName1,
        tag1,
        differentDonFamily
      )
    );
    vm.prank(s_user);
    s_registry.upsertWorkflow(
      workflowName1,
      tag1,
      workflowId2,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      differentDonFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
  }

  function test_upsertWorkflow_RevertWhenStatusIsChangedOnUpdateWorkflow()
    external
    whenMsgSenderIsALinkedOwner
    whenDONLimtisAreSet
  {
    // it should revert because this is not allowed

    // deploy a workflow first
    bytes32 workflowId1 = keccak256("workflow1");
    string memory workflowName1 = "Price Oracle";
    string memory tag1 = "oracle-main";
    string memory binaryUrl1 = "https://example.com/binaries/price-oracle.wasm";
    string memory configUrl1 = "https://example.com/configs/price-oracle.json";
    bytes memory attributes1 = abi.encode("Price Oracle v1.0");

    vm.prank(s_user);
    s_registry.upsertWorkflow(
      workflowName1,
      tag1,
      workflowId1,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 1, "There should be 1 workflow for the s_user");

    bytes32 workflowId2 = keccak256("workflow2");
    vm.expectRevert(
      abi.encodeWithSelector(
        WorkflowRegistry.CannotChangeStatusOnUpdate.selector,
        workflowId2,
        s_user,
        workflowName1,
        tag1,
        WorkflowRegistry.WorkflowStatus.PAUSED
      )
    );
    vm.prank(s_user);
    s_registry.upsertWorkflow(
      workflowName1,
      tag1,
      workflowId2,
      WorkflowRegistry.WorkflowStatus.PAUSED,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
  }

  //   modifier whenKeepAliveIsFalse() {
  //     _;
  //   }

  //   whenNoExistingRecordForOwnerNameTagCombo
  //   whenAllMetadataInputsAreValid
  //   whenKeepAliveIsFalse
  function test_upsertWorkflow_WhenThereAreMoreThanOneActiveWorkflowThatSharesTheKey()
    external
    whenMsgSenderIsALinkedOwner
    whenDONLimtisAreSet
  {
    // it should pause each before continuing
    // deploy a workflow first
    bytes32 workflowId1 = keccak256("workflow1");
    string memory workflowName1 = "Price Oracle";
    string memory tag1 = "oracle-main";
    string memory binaryUrl1 = "https://example.com/binaries/price-oracle.wasm";
    string memory configUrl1 = "https://example.com/configs/price-oracle.json";
    bytes memory attributes1 = abi.encode("Price Oracle v1.0");

    vm.startPrank(s_user);
    s_registry.upsertWorkflow(
      workflowName1,
      tag1,
      workflowId1,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 1, "There should be 1 workflow for the s_user");

    // deploy another workflow with the same name but a different tag
    bytes32 workflowId2 = keccak256("workflow2");
    string memory tag2 = "weather-feed";

    s_registry.upsertWorkflow(
      workflowName1,
      tag2,
      workflowId2,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      false
    );
    wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 2, "There should be 2 workflows for the s_user");

    WorkflowRegistry.WorkflowMetadataView memory wf1 = s_registry.getWorkflowById(workflowId1);
    WorkflowRegistry.WorkflowMetadataView memory wf2 = s_registry.getWorkflowById(workflowId2);
    // the first one should be paused now, the second one should be active
    assertTrue(wf1.status == WorkflowRegistry.WorkflowStatus.PAUSED);
    assertTrue(wf2.status == WorkflowRegistry.WorkflowStatus.ACTIVE);

    vm.stopPrank();
  }

  //   whenNoExistingRecordForOwnerNameTagCombo
  //   whenAllMetadataInputsAreValid
  //   whenKeepAliveIsFalse
  function test_upsertWorkflow_WhenThereAreNoActiveWorkflowsWithTheSameKey()
    external
    whenMsgSenderIsALinkedOwner
    whenDONLimtisAreSet
  {
    // it does not pause any other workflows
    // deploy a workflow first
    bytes32 workflowId1 = keccak256("workflow1");
    string memory workflowName1 = "Price Oracle";
    string memory tag1 = "oracle-main";
    string memory binaryUrl1 = "https://example.com/binaries/price-oracle.wasm";
    string memory configUrl1 = "https://example.com/configs/price-oracle.json";
    bytes memory attributes1 = abi.encode("Price Oracle v1.0");

    vm.startPrank(s_user);
    s_registry.upsertWorkflow(
      workflowName1,
      tag1,
      workflowId1,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      true
    );
    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 1, "There should be 1 workflow for the s_user");

    // deploy another workflow with different name
    bytes32 workflowId2 = keccak256("workflow2");
    string memory workflowName2 = "Weather Feed";

    s_registry.upsertWorkflow(
      workflowName2,
      tag1,
      workflowId2,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      binaryUrl1,
      configUrl1,
      attributes1,
      false
    );
    wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 2, "There should be 2 workflows for the s_user");

    WorkflowRegistry.WorkflowMetadataView memory wf1 = s_registry.getWorkflowById(workflowId1);
    WorkflowRegistry.WorkflowMetadataView memory wf2 = s_registry.getWorkflowById(workflowId2);
    // they should both be active as they have different names
    assertTrue(wf1.status == WorkflowRegistry.WorkflowStatus.ACTIVE);
    assertTrue(wf2.status == WorkflowRegistry.WorkflowStatus.ACTIVE);

    vm.stopPrank();
  }

  modifier whenTheNewWorkflowStatusIsActive() {
    _;
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenAllMetadataInputsAreValid
  // whenTheNewWorkflowStatusIsActive
  function test_upsertWorkflow_WhenTheDonHasNoGlobalCapSet() external whenMsgSenderIsALinkedOwner {
    // it should revert with DonLimitNotSet
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.DonLimitNotSet.selector, s_donFamily));
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenAllMetadataInputsAreValid
  // whenTheNewWorkflowStatusIsActive
  function test_upsertWorkflow_WhenOwnerWouldExceedTheirEffectiveCap() external whenMsgSenderIsALinkedOwner {
    // it should revert with MaxWorkflowsPerUserDONExceeded
    // set the DON limit to 1 and default user limit to 1, but DON limit has to be respected first
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 1, 1);

    // first upsert will work
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);

    WorkflowRegistry.WorkflowMetadataView[] memory wrs = s_registry.getWorkflowListByOwner(s_user, 0, 100);
    assertEq(wrs.length, 1, "There should be 1 workflow for the s_user");

    // try upsert another workflow in the same donFamily, it will revert saying that the global DON limit has been
    // exceeded (regardless of the user limit)
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "second-workflow",
      s_tag,
      keccak256("workflow2"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // set the DON limit to 2 and leave the default user limit to 1
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 2, 1);

    // try upsert another workflow in the same donFamily, it will revert saying that the user default limit for this DON
    // has been exceeded
    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerUserDONExceeded.selector, s_user, s_donFamily)
    );
    s_registry.upsertWorkflow(
      "second-workflow",
      s_tag,
      keccak256("workflow2"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // now apply the user limit override and allow the user to create 2 workflows in this DON family
    // the global limit must still be respected (currently set at 2)
    vm.prank(s_owner);
    s_registry.setUserDONOverride(s_user, s_donFamily, 2, true);

    // second upsert finally succeeds
    vm.prank(s_user);
    s_registry.upsertWorkflow(
      "second-workflow",
      s_tag,
      keccak256("workflow2"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // but now the third upsert will fail as the global DON limit is reached (user override must be ignored)
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerDONExceeded.selector, s_donFamily));
    s_registry.upsertWorkflow(
      "third-workflow",
      s_tag,
      keccak256("workflow3"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // set the DON limit to 5 and leave the default user limit to 1
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 5, 1);

    // try upsert another workflow in the same donFamily, it will revert saying that the user limit for this DON has
    // been exceeded (limit is 2 because override is respected)
    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerUserDONExceeded.selector, s_user, s_donFamily)
    );
    s_registry.upsertWorkflow(
      "third-workflow",
      s_tag,
      keccak256("workflow3"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // now apply the user limit override and allow the user to create 4 workflows in this DON family
    // the global limit must still be respected (currently set at 5)
    vm.prank(s_owner);
    s_registry.setUserDONOverride(s_user, s_donFamily, 4, true);

    // third upsert finally succeeds
    vm.prank(s_user);
    s_registry.upsertWorkflow(
      "third-workflow",
      s_tag,
      keccak256("workflow3"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    // now remove the user override
    vm.prank(s_owner);
    s_registry.setUserDONOverride(s_user, s_donFamily, 4, false);

    // even though the user override was previously set at 4, now it has to fallback to the default user limit,
    // currently set at 1 workflow, which means the next upsert will fail
    vm.prank(s_user);
    vm.expectRevert(
      abi.encodeWithSelector(WorkflowRegistry.MaxWorkflowsPerUserDONExceeded.selector, s_user, s_donFamily)
    );
    s_registry.upsertWorkflow(
      "fourth-workflow",
      s_tag,
      keccak256("workflow4"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );
  }

  // whenNoExistingRecordForOwnerNameTagCombo
  // whenAllMetadataInputsAreValid
  function test_upsertWorkflow_WhenThereAreNoFailures() external whenMsgSenderIsALinkedOwner whenDONLimtisAreSet {
    // it should write the new record update all indices and emit WorkflowRegistered
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.WorkflowRegistered(
      s_workflowId, s_user, s_donFamily, WorkflowRegistry.WorkflowStatus.ACTIVE, s_workflowName
    );
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);
  }

  // ================================================================
  // |                       UPDATE PATH                            |
  // ================================================================
  modifier whenAnExistingRecordExistsAtRid() {
    _linkOwner(s_user);
    _setDONLimit();
    _upsertTestWorklow(WorkflowRegistry.WorkflowStatus.ACTIVE, false, s_user);
    _;
  }

  //   modifier whenThereAreValidationFailures() {
  //     _;
  //   }

  // whenAnExistingRecordExistsAtRid
  // whenThereAreValidationFailures
  function test_upsertWorkflow_WhenNewWorkflowIdIsZero() external whenAnExistingRecordExistsAtRid {
    // it should revert with ZeroWorkflowIDNotAllowed
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.ZeroWorkflowIDNotAllowed.selector));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      "",
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
  }

  // whenAnExistingRecordExistsAtRid
  // whenThereAreValidationFailures
  function test_upsertWorkflow_WhenNewWorkflowIdAlreadyExists() external whenAnExistingRecordExistsAtRid {
    // it should revert with WorkflowIDAlreadyExists
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.WorkflowIDAlreadyExists.selector, s_workflowId));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      s_workflowId,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      false
    );
  }

  function test_upsertWorkflow_WhenNewBinaryUrlOrNewConfigUrlLengthGreaterThanAllowed()
    external
    whenAnExistingRecordExistsAtRid
  {
    // it should revert with URLTooLong
    uint8 cap = s_registry.getConfig().maxUrlLen;
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.URLTooLong.selector, bytes(s_invalidURL).length, cap));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      keccak256("workflow2"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_invalidURL,
      s_configUrl,
      s_attributes,
      false
    );
  }

  function test_upsertWorkflow_WhenNewBinaryUrlIsMissing() external whenAnExistingRecordExistsAtRid {
    // it should revert with BinaryURLRequired
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.BinaryURLRequired.selector));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      keccak256("workflow2"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      "",
      s_configUrl,
      s_attributes,
      false
    );
  }

  function test_upsertWorkflow_WhenNewAttributesLengthGreaterThanAllowed() external whenAnExistingRecordExistsAtRid {
    // it should revert with AttributesTooLong
    bytes memory tooBigAttrs = new bytes(1025);
    for (uint256 i = 0; i < tooBigAttrs.length; i++) {
      tooBigAttrs[i] = bytes1(uint8(i % 256));
    }

    uint16 cap = s_registry.getConfig().maxAttrLen;
    vm.prank(s_user);
    vm.expectRevert(abi.encodeWithSelector(WorkflowRegistry.AttributesTooLong.selector, bytes(tooBigAttrs).length, cap));
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      keccak256("workflow2"),
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      tooBigAttrs,
      true
    );
  }

  function test_upsertWorkflow_WhenThereAreNoValidationFailures() external whenAnExistingRecordExistsAtRid {
    // it should remap id to rid with the new workflowId
    // it should patch mutable fields
    // it should emit WorkflowUpdated
    bytes32 newWorkflowName = keccak256("workflow2");
    vm.prank(s_user);
    vm.expectEmit(true, true, true, false);
    emit WorkflowRegistry.WorkflowUpdated(s_workflowId, newWorkflowName, s_user, s_donFamily, s_workflowName);
    s_registry.upsertWorkflow(
      s_workflowName,
      s_tag,
      newWorkflowName,
      WorkflowRegistry.WorkflowStatus.ACTIVE,
      s_donFamily,
      s_binaryUrl,
      s_configUrl,
      s_attributes,
      true
    );

    WorkflowRegistry.WorkflowMetadataView memory wf = s_registry.getWorkflowById(newWorkflowName);
    assertEq(wf.workflowName, s_workflowName);
  }
}
