// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../KeystoneForwarder.sol";
import "../DataFeedsCache.sol";


contract DataFeedsKeystoneBatchWriteTest is Script {
    KeystoneForwarder internal s_keystoneForwarder;
    DataFeedsCache internal s_dataFeedsCache;
    uint256 constant internal i_maxBatchSize = 500;
    uint256 constant internal i_everySizeUpToN = 500;
    uint256 constant internal i_batchIncreaseSize = 1;

    uint256 internal constant MAX_ORACLES = 31;
    uint32 internal DON_ID = 0x01020304;
    uint8 internal F = 5;
    uint32 internal CONFIG_VERSION = 1;

    struct Signer {
        uint256 mockPrivateKey;
        address signerAddress;
    }

    Signer[MAX_ORACLES] internal s_signers;

    uint256 internal requiredSignaturesNum = F + 1;
    bytes[] internal signatures;
    bytes internal reportContext = new bytes(96);

    bytes internal payload;

    bytes32 internal executionId = hex"6d795f657865637574696f6e5f69640000000000000000000000000000000000";
    uint8 internal version = 1;

    bytes internal header;
    bytes internal metadata;
    bytes internal report;

    function run() external {
        uint32 timestamp = 99;

        vm.startBroadcast();

        s_keystoneForwarder = new KeystoneForwarder();

        s_dataFeedsCache = new DataFeedsCache();

        vm.stopBroadcast();

        uint256 seed = 0;

        for (uint256 i; i < MAX_ORACLES; ++i) {
            uint256 mockPK = seed + i + 1;
            s_signers[i].mockPrivateKey = mockPK;
            s_signers[i].signerAddress = vm.addr(mockPK);
        }

        address reportSender = address(s_keystoneForwarder);

        bytes32 workflowId = hex"6d795f6964000000000000000000000000000000000000000000000000000000";
        bytes2 reportId = bytes2(uint16(1));

        address[] memory senders = new address[](2);
        senders[0] = reportSender;
        senders[1] = reportSender;
        address[] memory workflowOwners = new address[](2);
        workflowOwners[0] = address(s_keystoneForwarder);
        workflowOwners[1] = address(s_keystoneForwarder);
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

        bytes16[] memory dataIds = new bytes16[](i_maxBatchSize);
        string[] memory descriptions = new string[](i_maxBatchSize);

        address[] memory signers = _getSignerAddresses(16);

        for (uint256 i = 0; i < i_maxBatchSize; i++) {
            dataIds[i] = bytes16(uint128(i + 1));
            descriptions[i] = description;
        }

        vm.startBroadcast();

        s_keystoneForwarder.addForwarder(address(s_keystoneForwarder));

        s_keystoneForwarder.setConfig(DON_ID, CONFIG_VERSION, F, signers);

        s_dataFeedsCache.setFeedAdmin(0x82EEd7C3a79Cb05578272ACF032d9935C759A898, true);

        // Setup i_maxBatchSize feed configs
        s_dataFeedsCache.setDecimalFeedConfigs{gas: 1000000000000}(dataIds, descriptions, workflowMetadata);

        vm.stopBroadcast();

        // Manually construct encoded initiaPpayload
        bytes memory initialPayload = abi.encodePacked(
            bytes32(uint256(32)),             // Offset (points to data start)
            bytes32(uint256(i_maxBatchSize))  // Length
        );

        for (uint256 i = 0; i < i_maxBatchSize; i++) {
            initialPayload = bytes.concat(
                initialPayload,
                abi.encodePacked(
                    bytes32(dataIds[i]),
                    abi.encode(99),
                    abi.encode(9999)
                )
            );
        }
        
        timestamp = uint32(99);
        reportId = bytes2(uint16(reportId) + 1);
        bytes memory metadata = abi.encodePacked(workflowId, workflowNames[0], workflowOwners[0], reportId);
        header = abi.encodePacked(version, executionId, timestamp, DON_ID, CONFIG_VERSION, metadata);
        report = abi.encodePacked(header, initialPayload);

        signatures = _signReport(report, reportContext, requiredSignaturesNum);

        vm.startBroadcast();

        s_keystoneForwarder.report{gas: 1000000000000}(address(s_dataFeedsCache), report, reportContext, signatures);

        vm.stopBroadcast();

        uint256 currentBatchSize = 1;

        while (currentBatchSize <= i_maxBatchSize) {
            // Manually construct encoded payload
            payload = abi.encodePacked(
                bytes32(uint256(32)),               // Offset (points to data start)
                bytes32(uint256(currentBatchSize))  // Length
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

            timestamp = uint32(currentBatchSize * 100);
            reportId = bytes2(uint16(reportId) + 1);
            metadata = abi.encodePacked(workflowId, workflowNames[0], workflowOwners[0], reportId);
            header = abi.encodePacked(version, executionId, timestamp, DON_ID, CONFIG_VERSION, metadata);
            report = abi.encodePacked(header, payload);

            signatures = _signReport(report, reportContext, requiredSignaturesNum);

            vm.startBroadcast();

            s_keystoneForwarder.report{gas: 100_000_000}(address(s_dataFeedsCache), report, reportContext, signatures);

            vm.stopBroadcast();

            if (currentBatchSize < i_everySizeUpToN) {
                currentBatchSize++;
            } else {
                currentBatchSize+= i_batchIncreaseSize;
            }
        }
    }

    function _getSignerAddresses() internal view returns (address[] memory) {
        address[] memory signerAddrs = new address[](s_signers.length);
        for (uint256 i = 0; i < signerAddrs.length; ++i) {
            signerAddrs[i] = s_signers[i].signerAddress;
        }
        return signerAddrs;
    }

    function _getSignerAddresses(uint256 limit) internal view returns (address[] memory) {
        address[] memory signerAddrs = new address[](limit);
        for (uint256 i = 0; i < limit; ++i) {
            signerAddrs[i] = s_signers[i].signerAddress;
        }
        return signerAddrs;
    }

    function _signReport(
        bytes memory report,
        bytes memory reportContext,
        uint256 requiredSignatures
    ) internal view returns (bytes[] memory signatures) {
        signatures = new bytes[](requiredSignatures);
        for (uint256 i = 0; i < requiredSignatures; ++i) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(
                s_signers[i].mockPrivateKey,
                keccak256(abi.encodePacked(keccak256(report), reportContext))
            );
            signatures[i] = bytes.concat(r, s, bytes1(v - 27));
        }
        return signatures;
    }
}