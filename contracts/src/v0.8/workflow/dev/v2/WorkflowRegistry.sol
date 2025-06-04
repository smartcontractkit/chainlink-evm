// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "../../../shared/access/Ownable2StepMsgSender.sol";

contract WorkflowRegistry is Ownable2StepMsgSender, ITypeAndVersion {
  string public constant override typeAndVersion = "WorkflowRegistry 2.0.0-dev";

  enum WorkflowStatus {
    ACTIVE,
    PAUSED
  }

  struct WorkflowMetadata {
    bytes32 workflowID; //     Unique identifier from hash of owner address, WASM binary content, config content and secrets URL.
    bytes32 donLabel; //       Label for the DON that is used when distributing the workflow across DONs.
    address owner; // ─────────╮ Workflow owner.
    uint64 created_at; //      │ block.timestamp when the workflow was first registered.
    WorkflowStatus status; // ─╯ Current status of the workflow (active, paused).
    string workflowName; //    Human readable string capped at 64 characters length.
    string binaryURL; //       URL to the WASM binary.
    string configURL; //       URL to the config.
    string secretsURL; //      URL to the encrypted secrets. Workflow DON applies a default refresh period (e.g. daily).
  }

  constructor() {
    // Intialize with default limits for Config.
    s_cfg.defaultMaxPerDON = 500;
    s_cfg.defaultMaxPerUserDON = 200;
  }

  // ================================================================
  // |                        Limits Config                         |
  // ================================================================

  // Struct to distinguish between unset and explicitly set zero values
  struct ConfigValue {
    uint32 value;
    bool enabled;
  }

  /// @dev Instead of a big struct with mappings, we store
  ///      defaults in a single 32-byte slot, and use nested mappings
  ///      for overrides.
  struct Config {
    uint32 defaultMaxPerDON;
    uint32 defaultMaxPerUserDON;
    // 2) DON-specific overrides: donLabel -> limit
    mapping(bytes32 => ConfigValue) donOverride;
    // 3) user+don override: user => (donLabel => limit)
    mapping(address => mapping(bytes32 => ConfigValue)) userDONOverride;
  }

  // Our single config instance
  Config private s_cfg;

  // ─────────────────────────────────────────────────────────────────────────
  // Limits Config - External Setters: Owner can set defaults and individual overrides
  // ─────────────────────────────────────────────────────────────────────────

  /// @notice Update the default limits.
  function setDefaults(uint32 maxPerDON, uint32 maxPerUserDON) external onlyOwner {
    s_cfg.defaultMaxPerDON = maxPerDON;
    s_cfg.defaultMaxPerUserDON = maxPerUserDON;
  }

  /// @notice Override the maximum # of workflows allowed for a specific DON label.
  /// @dev donLabel is a bytes32 value of the string, which should not exceed 32 ASCII characters.
  function setDONOverride(bytes32 donLabel, uint32 limit, bool enabled) external onlyOwner {
    if (enabled) {
      s_cfg.donOverride[donLabel] = ConfigValue(limit, true);
    } else {
      delete s_cfg.donOverride[donLabel];
    }
  }

  /// @notice Override the max # of workflows for a specific (user, DON) pair.
  function setUserDONOverride(address user, bytes32 donLabel, uint32 limit, bool enabled) external onlyOwner {
    if (enabled) {
      s_cfg.userDONOverride[user][donLabel] = ConfigValue(limit, true);
    } else {
      delete s_cfg.userDONOverride[user][donLabel];
    }
  }

  // ─────────────────────────────────────────────────────────────────────────
  // Limits Config -  Public Getters that return the *effective* limit
  //    (override if set, else default)
  // ─────────────────────────────────────────────────────────────────────────

  /// @notice Effective max # of workflows for a particular DON.
  function getMaxWorkflowsPerDON(
    bytes32 donLabel
  ) public view returns (uint32) {
    ConfigValue memory cfgVal = s_cfg.donOverride[donLabel];
    if (cfgVal.enabled) {
      return cfgVal.value;
    }
    return s_cfg.defaultMaxPerDON;
  }

  /// @notice Effective max # of workflows for a (user, DON) combo.
  function getMaxWorkflowsPerUserDON(address user, bytes32 donLabel) public view returns (uint32) {
    ConfigValue memory cfgVal = s_cfg.userDONOverride[user][donLabel];
    if (cfgVal.enabled) {
      return cfgVal.value;
    }
    return s_cfg.defaultMaxPerUserDON;
  }

  /// @notice Returns the default limits:
  ///         (maxWorkflowsPerDON, maxWorkflowsPerUserDON)
  function getDefaults() external view returns (uint32 maxPerDON, uint32 maxPerUserDON) {
    return (s_cfg.defaultMaxPerDON, s_cfg.defaultMaxPerUserDON);
  }
}
