// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/console.sol";

import {BundleAggregatorProxy} from "../BundleAggregatorProxy.sol";
import {DataFeedsCache} from "../DataFeedsCache.sol";

import {BaseTest} from "./BaseTest.t.sol";
import {DataFeedsLegacyAggregatorProxy} from "./helpers/DataFeedsLegacyAggregatorProxy.sol";

// solhint-disable-next-line max-states-count
contract DataFeedsSetupBatchGasTestLocal is BaseTest {
  struct ReceivedBundleReport {
    bytes32 dataId;
    uint32 timestamp;
    bytes bundle;
  }

  DataFeedsLegacyAggregatorProxy internal s_dataFeedsLegacyAggregatorProxy;
  BundleAggregatorProxy internal s_dataFeedsAggregatorProxy;
  DataFeedsCache internal s_dataFeedsCache;


  address internal s_reportSender = address(10001);

  bytes32 internal s_workflowId = hex"6d795f6964000000000000000000000000000000000000000000000000000000";
  bytes2 internal s_reportId = hex"0001";

  address[] internal s_senders = [s_reportSender, s_reportSender];
  address[] internal s_workflowOwners = [address(10002), address(10003)];
  bytes10[] internal s_workflowNames = [bytes10("abc"), bytes10("xyz")];

  DataFeedsCache.WorkflowMetadata internal s_workflowMetadata1 = DataFeedsCache.WorkflowMetadata({
    allowedSender: s_senders[0],
    allowedWorkflowOwner: s_workflowOwners[0],
    allowedWorkflowName: s_workflowNames[0]
  });

  DataFeedsCache.WorkflowMetadata internal s_workflowMetadata2 = DataFeedsCache.WorkflowMetadata({
    allowedSender: s_senders[1],
    allowedWorkflowOwner: s_workflowOwners[1],
    allowedWorkflowName: s_workflowNames[1]
  });

  DataFeedsCache.WorkflowMetadata[] internal s_workflowMetadata;

  string internal s_description = "description";

  bytes internal s_metadata;

  uint256 constant internal i_batchSize = 20;

  function setUp() public virtual override {
    BaseTest.setUp();

    vm.stopPrank();
    vm.startPrank(OWNER);

    s_dataFeedsCache = new DataFeedsCache();
    s_dataFeedsLegacyAggregatorProxy = new DataFeedsLegacyAggregatorProxy(address(s_dataFeedsCache));
    s_dataFeedsAggregatorProxy = new BundleAggregatorProxy(address(s_dataFeedsCache), OWNER);

    s_metadata = abi.encodePacked(s_workflowId, s_workflowNames[0], s_workflowOwners[0], s_reportId);

    s_workflowMetadata.push(s_workflowMetadata1);
    s_workflowMetadata.push(s_workflowMetadata2);

    vm.stopPrank();
  }
}
