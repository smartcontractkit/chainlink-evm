// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {DataFeedsSetupBatchGasTest} from "./DataFeedsSetupBatchGasTest.t.sol";

contract DataFeedsCacheBatchGasTest is DataFeedsSetupBatchGasTest {

  function setUp() public virtual override {
    DataFeedsSetupBatchGasTest.setUp();

  }

  function test_write_onReport_prices_batch_gas() public {
    string memory testBaseName = "test_write_onReport_prices_batch_gas_";
    for (uint256 i = 0; i < i_batchSize; i++) {
      uint256 currBatchSize = i + 1;

      // Manually construct encoded payload
      bytes memory writePayload = abi.encodePacked(
          bytes32(uint256(32)),            // Offset (points to data start)
          bytes32(uint256(currBatchSize))  // Length
      );


      for (uint256 j = 0; j < currBatchSize; j++) {
        // Increment timestamp and price to write different values each time
        s_timestamps[j] += 100;
        s_prices[j] += 100;
        writePayload = bytes.concat(
            writePayload,
            bytes32(s_dataIds[j]),
            abi.encode(s_timestamps[j]),
            abi.encode(s_prices[j])
        );
      }
      string memory testName = string.concat(testBaseName, vm.toString(currBatchSize));
      vm.startPrank(s_reportSender);
      
      vm.stopPrank();
      
      vm.startSnapshotGas(testName);
      s_dataFeedsCache.onReport(s_metadata, writePayload);
      vm.stopSnapshotGas(testName);
    }
  }
}
