// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "../../shared/access/Ownable2StepMsgSender.sol";

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
    s_cfg.defaultsPacked = _packDefaults(200, 500, 200);
  }

  // ================================================================
  // |                        Limits Config                         |
  // ================================================================

  /// @dev Instead of a big struct with mappings, we store
  ///      defaults in a single 32-byte slot, and use nested mappings
  ///      for overrides.
  struct Config {
    // Pack three uint32 defaults into a single 96-bit field:
    // Layout in bits: [maxUser(0..31) | maxDON(32..63) | maxUserDON(64..95)]
    uint96 defaultsPacked;
    // Struct to distinguish between unset and explicitly set zero values
    struct Value {
      uint32 value;
      bool enabled;
    }
    // 1) user-specific overrides: address -> limit
    mapping(address => Value) userOverride;
    // 2) DON-specific overrides: donLabel -> limit
    mapping(bytes32 => Value) donOverride;
    // 3) user+don override: user => (donLabel => limit)
    mapping(address => mapping(bytes32 => Value)) userDONOverride;
  }

  // Our single config instance
  Config private s_cfg;

  // ─────────────────────────────────────────────────────────────────────────
  // Limits Config - External Setters: Owner can set defaults and individual overrides
  // ─────────────────────────────────────────────────────────────────────────

  /// @notice Update the three default limits in a single call.
  function setDefaults(uint32 maxPerUser, uint32 maxPerDON, uint32 maxPerUserDON) external onlyOwner {
    s_cfg.defaultsPacked = _packDefaults(maxPerUser, maxPerDON, maxPerUserDON);
  }

  /// @notice Override the maximum # of workflows a specific user can register.
  function setUserOverride(address user, uint32 limit, bool enabled) external onlyOwner {
    if (enabled) {
      s_cfg.userOverride[user] = Config.Value(limit, true);
    } else {
      delete s_cfg.userOverride[user];
    }
  }

  /// @notice Override the maximum # of workflows allowed for a specific DON label.
  /// @dev donLabel is a bytes32 value of the string, which should not exceed 32 ASCII characters.
  function setDONOverride(bytes32 donLabel, uint32 limit, bool enabled) external onlyOwner {
    if (enabled) {
      s_cfg.donOverride[donLabel] = Config.Value(limit, true);
    } else {
      delete s_cfg.donOverride[donLabel];
    }
  }

  /// @notice Override the max # of workflows for a specific (user, DON) pair.
  function setUserDONOverride(address user, bytes32 donLabel, uint32 limit, bool enabled) external onlyOwner {
    if (enabled) {
      s_cfg.userDONOverride[user][donLabel] = Config.Value(limit, true);
    } else {
      delete s_cfg.userDONOverride[user][donLabel];
    }
  }

  // ─────────────────────────────────────────────────────────────────────────
  // Limits Config -  Public Getters that return the *effective* limit
  //    (override if set, else default)
  // ─────────────────────────────────────────────────────────────────────────

  /// @notice Effective max # of workflows for a particular user.
  function getMaxWorkflowsPerUser(
    address user
  ) public view returns (uint32) {
    Config.Value memory override = s_cfg.userOverride[user];
    if (override.enabled) {
      return override.value;
    }
    // fallback to the default
    (uint32 defUser,,) = _unpackDefaults(s_cfg.defaultsPacked);
    return defUser;
  }

  /// @notice Effective max # of workflows for a particular DON.
  function getMaxWorkflowsPerDON(
    bytes32 donLabel
  ) public view returns (uint32) {
    Config.Value memory override = s_cfg.donOverride[donLabel];
    if (override.enabled) {
      return override.value;
    }
    // fallback to the default
    (, uint32 defDON,) = _unpackDefaults(s_cfg.defaultsPacked);
    return defDON;
  }

  /// @notice Effective max # of workflows for a (user, DON) combo.
  function getMaxWorkflowsPerUserDON(address user, bytes32 donLabel) public view returns (uint32) {
    Config.Value memory override = s_cfg.userDONOverride[user][donLabel];
    if (override.enabled) {
      return override.value;
    }
    // fallback to the default
    (,, uint32 defUserDON) = _unpackDefaults(s_cfg.defaultsPacked);
    return defUserDON;
  }

  /// @notice Returns the three default limits:
  ///         (maxWorkflowsPerUser, maxWorkflowsPerDon, maxWorkflowsPerUserDon)
  function getDefaults() external view returns (uint32 maxPerUser, uint32 maxPerDon, uint32 maxPerUserDon) {
    (maxPerUser, maxPerDon, maxPerUserDon) = _unpackDefaults(s_cfg.defaultsPacked);
    return (maxPerUser, maxPerDon, maxPerUserDon);
  }

  // ─────────────────────────────────────────────────────────────────────────
  // Limits Config - Internal Helpers: set/read defaults (packed into one 96-bit variable)
  // ─────────────────────────────────────────────────────────────────────────

  /// @dev Store 3 uint32 values in a single 96-bit field.
  function _packDefaults(
    uint32 maxPerUser,
    uint32 maxPerDON,
    uint32 maxPerUserDON
  ) internal pure returns (uint96 packed) {
    // lower 32 bits: maxPerUser
    // middle 32 bits: maxPerDON
    // top 32 bits: maxPerUserDON
    packed = uint96(maxPerUser) | (uint96(maxPerDON) << 32) | (uint96(maxPerUserDON) << 64);
    return packed;
  }

  /// @dev Extract the 3 defaults from the packed value.
  function _unpackDefaults(
    uint96 packed
  ) internal pure returns (uint32 maxPerUser, uint32 maxPerDON, uint32 maxPerUserDON) {
    maxPerUser = uint32(packed);
    maxPerDON = uint32(packed >> 32);
    maxPerUserDON = uint32(packed >> 64);
    return (maxPerUser, maxPerDON, maxPerUserDON);
  }
}
