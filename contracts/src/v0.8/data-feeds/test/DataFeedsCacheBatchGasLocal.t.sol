// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {DataFeedsSetupBatchGasTestLocal} from "./DataFeedsSetupBatchGasTestLocal.t.sol";

contract DataFeedsCacheBatchGasTestLocal is DataFeedsSetupBatchGasTestLocal {

  function setUp() public virtual override {
    DataFeedsSetupBatchGasTestLocal.setUp();

    for (uint256 i = 0; i < i_batchSize; i++) {
      uint256 currBatchSize = i + 1;

      bytes16[] memory tmp_dataIds = new bytes16[](currBatchSize);
      string[] memory tmp_descriptions = new string[](currBatchSize);

      uint256 totalPreviousDataIds = i*(i+1)/2;

      uint32 timestamp = uint32(totalPreviousDataIds + 100);
      uint256 price = uint256(totalPreviousDataIds + 1000);

      for (uint256 j = 0; j < currBatchSize; j++) {
        // Unique DataIds to avoid warm writes.
        uint128 dataId = uint128(totalPreviousDataIds + j + 1);
        tmp_dataIds[j] = bytes16(dataId);
        tmp_descriptions[j] = s_description;
      }

      vm.startPrank(OWNER);

      s_dataFeedsCache.setFeedAdmin(OWNER, true);

      s_dataFeedsCache.setDecimalFeedConfigs(tmp_dataIds, tmp_descriptions, s_workflowMetadata);
      vm.stopPrank();

      // Manually construct encoded payload
      bytes memory payload = abi.encodePacked(
          bytes32(uint256(32)),             // Offset (points to data start)
          bytes32(uint256(currBatchSize))   // Length
      );

      for (uint256 j = 0; j < currBatchSize; j++) {
          payload = bytes.concat(
              payload,
              abi.encodePacked(
                bytes32(tmp_dataIds[j]),
                abi.encode(timestamp),
                abi.encode(price)
              )
          );
      }

      vm.startPrank(s_reportSender);

      //vm.store(address(s_dataFeedsCache), ,)

      s_dataFeedsCache.onReport(
        s_metadata,
        payload
      );

      vm.stopPrank();
    }
  }

  function test_write_onReport_prices_batch_gas() public {
    string memory testBaseName = "test_write_onReport_prices_batch_gas_";

    for (uint256 i = 0; i < i_batchSize; i++) {
      uint256 currBatchSize = i + 1;

      bytes16[] memory tmp_dataIds = new bytes16[](currBatchSize);

      uint256 totalPreviousDataIds = i*(i+1)/2;

      uint32 timestamp = uint32(totalPreviousDataIds + 100);
      uint256 price = uint256(totalPreviousDataIds + 1000);

      for (uint256 j = 0; j < currBatchSize; j++) {
        // Unique DataIds to avoid warm writes.
        uint128 dataId = uint128(totalPreviousDataIds + j + 1);
        tmp_dataIds[j] = bytes16(dataId);
      }

      // Manually construct encoded payload
      bytes memory writePayload = abi.encodePacked(
          bytes32(uint256(32)),            // Offset (points to data start)
          bytes32(uint256(currBatchSize))  // Length
      );

      for (uint256 j = 0; j < currBatchSize; j++) {
        // Increment timestamp and price to write different values each time
        writePayload = bytes.concat(
            writePayload,
            bytes32(tmp_dataIds[j]),
            abi.encode(timestamp),
            abi.encode(price)
        );
      }

      string memory testName = string.concat(testBaseName, vm.toString(currBatchSize));
      
      vm.startPrank(s_reportSender);

      vm.startSnapshotGas(testName);
      s_dataFeedsCache.onReport(s_metadata, writePayload);
      vm.stopSnapshotGas(testName);

      vm.stopPrank();
    }
  }
}
