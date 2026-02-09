// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @title ShardConfig
/// @notice Centralized onchain configuration contract for managing desired shard count
/// @dev This contract stores the desired number of shards and uses Ownable2StepMsgSender for secure ownership management
contract ShardConfig is ITypeAndVersion, Ownable2StepMsgSender {
    string public constant override typeAndVersion = "ShardConfig 1.0.0-dev";

    /// @notice The desired number of shards
    uint256 public desiredShardCount;

    /// @notice Emitted when the desired shard count is updated
    /// @param newCount The new desired shard count
    event ShardCountUpdated(uint256 indexed newCount);

    /// @notice Initialize the contract with initial shard count
    /// @param _desiredShardCount Initial desired shard count
    constructor(uint256 _desiredShardCount) {
        require(_desiredShardCount > 0, "Shard count must be greater than 0");

        desiredShardCount = _desiredShardCount;

        emit ShardCountUpdated(_desiredShardCount);
    }

    /// @notice Update the desired shard count (callable only by owner)
    /// @param _newCount The new desired shard count
    function setDesiredShardCount(uint256 _newCount) external onlyOwner {
        require(_newCount > 0, "Shard count must be greater than 0");

        desiredShardCount = _newCount;
        emit ShardCountUpdated(_newCount);
    }

    /// @notice Get the current desired shard count
    /// @return The current desired shard count
    function getDesiredShardCount() external view returns (uint256) {
        return desiredShardCount;
    }
}
