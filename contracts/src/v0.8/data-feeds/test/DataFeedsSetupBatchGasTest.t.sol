// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/console.sol";

import {BundleAggregatorProxy} from "../BundleAggregatorProxy.sol";
import {DataFeedsCache} from "../DataFeedsCache.sol";

import {BaseTest} from "./BaseTest.t.sol";
import {DataFeedsLegacyAggregatorProxy} from "./helpers/DataFeedsLegacyAggregatorProxy.sol";

// solhint-disable-next-line max-states-count
contract DataFeedsSetupBatchGasTest is BaseTest {
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

  uint256 constant internal i_batchSize = 3;

  bytes16[] internal s_dataIds = new bytes16[](i_batchSize);
  uint32[] internal s_timestamps = new uint32[](i_batchSize);
  uint256[] internal s_prices = new uint256[](i_batchSize);
  string[] internal s_descriptions = new string[](i_batchSize);

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

    s_dataFeedsCache.setFeedAdmin(OWNER, true);

    for (uint256 i = 0; i < i_batchSize; i++) {
        s_dataIds[i] = bytes16(uint128(i + 1));
        s_timestamps[i] = uint32((i_batchSize * 10) - i);
        s_prices[i] = uint256(1000000 + i);
        s_descriptions[i] = s_description;
    }

    s_dataFeedsCache.setDecimalFeedConfigs(s_dataIds, s_descriptions, s_workflowMetadata);

    // Manually construct encoded payload
    bytes memory payload = abi.encodePacked(
        bytes32(uint256(32)),          // Offset (points to data start)
        bytes32(uint256(i_batchSize))    // Length
    );

    for (uint256 i = 0; i < i_batchSize; i++) {
        payload = bytes.concat(
            payload,
            abi.encodePacked(
              bytes32(s_dataIds[i]),
              abi.encode(s_timestamps[i]),
              abi.encode(s_prices[i])
            )
        );
    }

    vm.stopPrank();
    vm.startPrank(s_reportSender);

    s_dataFeedsCache.onReport(
      s_metadata,
      payload
    );

    vm.stopPrank();
  }
}
