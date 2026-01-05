// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

/// @title ShardConfig
/// @notice Centralized onchain configuration contract for managing desired shard count
/// @dev This contract stores the desired number of shards and integrates with MCMS for secure configuration changes
contract ShardConfig is ITypeAndVersion {
    string public constant override typeAndVersion = "ShardConfig 1.0.0";

    /// @notice The desired number of shards
    uint256 public desiredShardCount;

    /// @notice MCMS contract address authorized to update shard config
    address public mcmsAddress;

    /// @notice Emitted when the desired shard count is updated
    /// @param newCount The new desired shard count
    event ShardCountUpdated(uint256 indexed newCount);

    /// @notice Emitted when the MCMS address is updated
    /// @param newMcmsAddress The new MCMS address
    event MCMSAddressUpdated(address indexed newMcmsAddress);

    /// @notice Initialize the contract with initial shard count and MCMS address
    /// @param _desiredShardCount Initial desired shard count
    /// @param _mcmsAddress Address of the MCMS contract
    constructor(uint256 _desiredShardCount, address _mcmsAddress) {
        require(_mcmsAddress != address(0), "Invalid MCMS address");
        require(_desiredShardCount > 0, "Shard count must be greater than 0");

        desiredShardCount = _desiredShardCount;
        mcmsAddress = _mcmsAddress;

        emit ShardCountUpdated(_desiredShardCount);
        emit MCMSAddressUpdated(_mcmsAddress);
    }

    /// @notice Update the desired shard count (callable only by MCMS)
    /// @param _newCount The new desired shard count
    function setDesiredShardCount(uint256 _newCount) external {
        require(msg.sender == mcmsAddress, "Only MCMS can update shard count");
        require(_newCount > 0, "Shard count must be greater than 0");

        desiredShardCount = _newCount;
        emit ShardCountUpdated(_newCount);
    }

    /// @notice Update the MCMS address (callable only by current MCMS)
    /// @param _newMcmsAddress The new MCMS address
    function setMCMSAddress(address _newMcmsAddress) external {
        require(msg.sender == mcmsAddress, "Only MCMS can update its own address");
        require(_newMcmsAddress != address(0), "Invalid MCMS address");

        mcmsAddress = _newMcmsAddress;
        emit MCMSAddressUpdated(_newMcmsAddress);
    }

    /// @notice Get the current desired shard count
    /// @return The current desired shard count
    function getDesiredShardCount() external view returns (uint256) {
        return desiredShardCount;
    }
}