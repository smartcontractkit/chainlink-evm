// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../DataFeedsCache.sol";


contract DataFeedsBatchWriteTest is Script {
    DataFeedsCache internal s_dataFeedsCache;
    uint256 constant internal i_maxBatchSize = 60;
    uint256 constant internal i_everySizeUpToN = 60;
    uint256 constant internal i_batchIncreaseSize = 10;

    bytes internal payload;

    function run() external {
        address reportSender = 0x82EEd7C3a79Cb05578272ACF032d9935C759A898;

        bytes32 workflowId = hex"6d795f6964000000000000000000000000000000000000000000000000000000";
        bytes2 reportId = hex"0001";

        address[] memory senders = new address[](2);
        senders[0] = reportSender;
        senders[1] = reportSender;
        address[] memory workflowOwners = new address[](2);
        workflowOwners[0] = 0x82EEd7C3a79Cb05578272ACF032d9935C759A898;
        workflowOwners[1] = 0x82EEd7C3a79Cb05578272ACF032d9935C759A898;
        bytes10[] memory workflowNames = new bytes10[](2);
        workflowNames[0] = bytes10("abc");
        workflowNames[1] = bytes10("xyz");

        DataFeedsCache.WorkflowMetadata memory workflowMetadata1 = DataFeedsCache.WorkflowMetadata({
            allowedSender: senders[0],
            allowedWorkflowOwner: workflowOwners[0],
            allowedWorkflowName: workflowNames[0]
        });

        DataFeedsCache.WorkflowMetadata memory workflowMetadata2 = DataFeedsCache.WorkflowMetadata({
            allowedSender: senders[1],
            allowedWorkflowOwner: workflowOwners[1],
            allowedWorkflowName: workflowNames[1]
        });

        DataFeedsCache.WorkflowMetadata[] memory workflowMetadata = new DataFeedsCache.WorkflowMetadata[](2);

        workflowMetadata[0] = workflowMetadata1;
        workflowMetadata[1] = workflowMetadata2;

        string memory description = "description";

        bytes memory metadata;

        metadata = abi.encodePacked(workflowId, workflowNames[0], workflowOwners[0], reportId);

        bytes16[] memory dataIds = new bytes16[](i_maxBatchSize);
        string[] memory descriptions = new string[](i_maxBatchSize);

        // Manually construct encoded initiaPpayload
        bytes memory initialPayload = abi.encodePacked(
            bytes32(uint256(32)),             // Offset (points to data start)
            bytes32(uint256(i_maxBatchSize))   // Length
        );

        for (uint256 i = 0; i < i_maxBatchSize; i++) {
            dataIds[i] = bytes16(uint128(i + 1));
            descriptions[i] = description;

            initialPayload = bytes.concat(
                initialPayload,
                abi.encodePacked(
                    bytes32(dataIds[i]),
                    abi.encode(5),
                    abi.encode(50)
                )
            );
        }

        vm.startBroadcast();

        s_dataFeedsCache = new DataFeedsCache();

        s_dataFeedsCache.setFeedAdmin(0x82EEd7C3a79Cb05578272ACF032d9935C759A898, true);

        // Setup i_maxBatchSize feed configs
        s_dataFeedsCache.setDecimalFeedConfigs(dataIds, descriptions, workflowMetadata);

        // Write i_maxBatchSize initial feed values to remove zero to non zero write gas cost from testing
        s_dataFeedsCache.onReport(
            metadata,
            initialPayload
        );

        vm.stopBroadcast();

        uint256 currentBatchSize = 1;

        while (currentBatchSize <= i_maxBatchSize) {
            // Manually construct encoded payload
            payload = abi.encodePacked(
                bytes32(uint256(32)),             // Offset (points to data start)
                bytes32(uint256(currentBatchSize))   // Length
            );

            for (uint256 j = 0; j < currentBatchSize; j++) {
                payload = bytes.concat(
                    payload,
                    abi.encodePacked(
                        bytes32(dataIds[j]),
                        abi.encode(currentBatchSize * 100),
                        abi.encode(currentBatchSize * 10000)
                    )
                );
            }

            vm.startBroadcast();

            s_dataFeedsCache.onReport(
                metadata,
                payload
            );

            vm.stopBroadcast();

            if (currentBatchSize < i_everySizeUpToN) {
                currentBatchSize++;
            } else {
                currentBatchSize+= i_batchIncreaseSize;
            }
        }
    }
}