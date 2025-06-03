// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "../../../shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "../../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

contract WorkflowRegistry is Ownable2StepMsgSender, ITypeAndVersion {
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.AddressSet;

  string public constant override typeAndVersion = "WorkflowRegistry 2.0.0-dev";
  uint8 private constant MAX_WORKFLOW_NAME_LENGTH = 64;
  uint8 private constant MAX_DON_LABEL_LENGTH = 32;
  uint8 private constant MAX_URL_LENGTH = 200;
  uint8 private constant MAX_PAGINATION_LIMIT = 100;

  enum WorkflowStatus {
    ACTIVE,
    PAUSED
  }

  struct WorkflowMetadata {
    bytes32 workflowID; //       Unique identifier from hash(owner, workflow name, wasm binary, cfg).
    address owner; // ─────────╮ Workflow owner.
    uint64 createdAt; //       │ block.timestamp when the workflow was first registered.
    WorkflowStatus status; // ─╯ Current status of the workflow (active, paused).
    string workflowName; //      Human readable string capped at 64 characters length.
    string binaryURL; //         URL to the wasm binary (64 chars limit).
    string configURL; //         URL to the config (64 chars limit).
    string donLabel; //          Label for the DON (32 chars limit).
    string tag; //               Unique per (owner, workflowName) human readable identifier
    bytes attributes; //         Arbitrary bytes for additional workflow details.
  }

  struct ConfigValue {
    uint32 value;
    bool enabled;
  }

  struct Config {
    uint32 defaultMaxPerDON; // e.g. 500
    uint32 defaultMaxPerUserDON; // e.g. 200
    mapping(bytes32 donLabel => ConfigValue limitValue) donOverride;
    mapping(address user => mapping(bytes32 donLabel => ConfigValue limitValue)) userDONOverride;
  }

  Config private s_cfg;

  // -------------------------------------------------------------------------
  //                         Storage Indices
  // -------------------------------------------------------------------------
  // Primary lookup by workflowID
  mapping(bytes32 workflowID => WorkflowMetadata workflowMetadata) private s_workflows;

  // Secondary indices for iteration / queries
  mapping(address => EnumerableSet.Bytes32Set) private s_ownerWorkflowIDs; // owner -> workflowID set
  mapping(bytes32 => EnumerableSet.Bytes32Set) private s_donWorkflowIDs; // donLabel -> workflowID set
  mapping(bytes32 => bytes32) private s_workflowKeyToLatestID; // keccak(owner,name) → *latest* workflowID
  mapping(bytes32 => EnumerableSet.Bytes32Set) private s_activeIDsByWorkflowKey; // workflowKey → active IDs
  // Fast counters for limit enforcement
  mapping(address => mapping(bytes32 => uint32)) private s_userDONCount; // owner -> (donLabel -> #workflows)
  // Tag uniqueness enforcement: (owner + name + tag) => bool
  mapping(bytes32 => bool) private s_tagsUsed;

  event WorkflowRegisteredV2(
    bytes32 indexed workflowID,
    address indexed owner,
    bytes32 indexed donLabel,
    WorkflowStatus status,
    string workflowName
  );

  event WorkflowUpdatedV2(
    bytes32 indexed oldWorkflowID,
    bytes32 indexed newWorkflowID,
    address indexed owner,
    bytes32 donLabel,
    string workflowName
  );

  event WorkflowPausedV2(
    bytes32 indexed workflowID, address indexed owner, string indexed donLabel, string workflowName
  );
  event WorkflowActivatedV2(
    bytes32 indexed workflowID, address indexed owner, string indexed donLabel, string workflowName
  );
  event WorkflowDeletedV2(
    bytes32 indexed workflowID, address indexed owner, string indexed donLabel, string workflowName
  );

  error CallerNotOwner();
  error DuplicateTag(string tag);
  error InvalidWorkflowID();
  error MaxWorkflowsPerDONExceeded(bytes32 donLabel);
  error MaxWorkflowsPerUserDONExceeded(address owner, bytes32 donLabel);
  error StringTooLong(uint8 expected, uint256 provided);
  error URLTooLong(uint256 provided, uint8 maxAllowed);
  error WorkflowAlreadyInDesiredStatus();
  error WorkflowDoesNotExist();
  error WorkflowIDAlreadyExists();
  error WorkflowNameRequired();
  error WorkflowNameTooLong(uint256 provided, uint8 maxAllowed);

  constructor() {
    s_cfg.defaultMaxPerDON = 500;
    s_cfg.defaultMaxPerUserDON = 200;
  }

  // ================================================================
  //                    Capabilities registry                       |
  // ================================================================
  struct CapabilitiesRegistryConfig {
    address registry;
    uint256 chainID; // EVM Chain where the capabilities registry is deployed
  }

  CapabilitiesRegistryConfig private s_capabilitiesRegistry;

  event CapabilitiesRegistryUpdated(address indexed oldAddr, address indexed newAddr, uint256 indexed chainID);

  /// @notice Owner can set/replace the capabilities registry config.
  function setCapabilitiesRegistry(address registry, uint256 chainID) external onlyOwner {
    address old = s_capabilitiesRegistry.registry;
    s_capabilitiesRegistry = CapabilitiesRegistryConfig({registry: registry, chainID: chainID});
    emit CapabilitiesRegistryUpdated(old, registry, chainID);
  }

  /// @return registry The capabilities registry address.
  /// @return chainID  The chain ID where it's deployed.
  function getCapabilitiesRegistry() external view returns (address registry, uint256 chainID) {
    CapabilitiesRegistryConfig memory cfg = s_capabilitiesRegistry;
    return (cfg.registry, cfg.chainID);
  }

  // ================================================================
  // |                        Limits Config                         |
  // ================================================================
  function setDefaults(uint32 maxPerDON, uint32 maxPerUserDON) external onlyOwner {
    s_cfg.defaultMaxPerDON = maxPerDON;
    s_cfg.defaultMaxPerUserDON = maxPerUserDON;
  }

  function setDONOverride(bytes32 donLabel, uint32 limit, bool enabled) external onlyOwner {
    if (enabled) {
      s_cfg.donOverride[donLabel] = ConfigValue(limit, true);
    } else {
      delete s_cfg.donOverride[donLabel];
    }
  }

  function setUserDONOverride(address user, bytes32 donLabel, uint32 limit, bool enabled) external onlyOwner {
    if (enabled) {
      s_cfg.userDONOverride[user][donLabel] = ConfigValue(limit, true);
    } else {
      delete s_cfg.userDONOverride[user][donLabel];
    }
  }

  function getMaxWorkflowsPerDON(
    bytes32 donLabel
  ) public view returns (uint32) {
    ConfigValue memory cv = s_cfg.donOverride[donLabel];
    return cv.enabled ? cv.value : s_cfg.defaultMaxPerDON;
  }

  function getMaxWorkflowsPerUserDON(address user, bytes32 donLabel) public view returns (uint32) {
    ConfigValue memory cv = s_cfg.userDONOverride[user][donLabel];
    return cv.enabled ? cv.value : s_cfg.defaultMaxPerUserDON;
  }

  function getDefaults() external view returns (uint32 maxPerDON, uint32 maxPerUserDON) {
    return (s_cfg.defaultMaxPerDON, s_cfg.defaultMaxPerUserDON);
  }

  // ================================================================
  // |                       Workflow Management                    |
  // ================================================================
  /**
   * @notice Upserts a new workflow based on workflowName + owner
   * @param workflowName  Human‑readable name (≤64 chars)
   * @param workflowID    Deterministic hash computed off‑chain (must be unique)
   * @param donLabel      Label of the DON (32 chars limit)
   * @param status        Initial status (ACTIVE / PAUSED)
   * @param binaryURL     URL of the wasm binary (required)
   * @param configURL     URL of the config (optional)
   * @param attributes    Arbitrary bytes for additional workflow details (optional)
   */
  function upsertWorkflow(
    string calldata workflowName,
    bytes32 workflowID,
    string calldata donLabel,
    WorkflowStatus status,
    string calldata binaryURL,
    string calldata configURL,
    bytes calldata attributes,
    string calldata tag,
    bool keepAlive
  ) external {
    _validateWorkflowName(bytes(workflowName).length);
    _validateWorkflowURLs(bytes(binaryURL).length, bytes(configURL).length);
    _validateWorkflowID(workflowID);

    // validate tag for uniqueness
    bytes32 tagKey = keccak256(abi.encodePacked(msg.sender, workflowName, tag));
    if (s_tagsUsed[tagKey]) revert DuplicateTag(tag);

    bytes32 dl = _stringToBytes32(donLabel);
    _enforceLimits(msg.sender, dl);

    bytes32 workflowKey = computeHashKey(msg.sender, workflowName);
    bytes32 latestID = s_workflowKeyToLatestID[workflowKey];
    bool isCreate = latestID == bytes32(0);

    if (!keepAlive && !isCreate) {
      pauseAllActiveWorkflowsByWorkflowKey(workflowKey);
    }
    // store metadata
    WorkflowMetadata memory meta = WorkflowMetadata({
      workflowID: workflowID,
      donLabel: donLabel,
      owner: msg.sender,
      createdAt: uint64(block.timestamp),
      status: status,
      workflowName: workflowName,
      binaryURL: binaryURL,
      configURL: configURL,
      tag: tag,
      attributes: attributes
    });

    s_workflows[workflowID] = meta;

    // update secondary indices & counters
    if (status == WorkflowStatus.ACTIVE) {
      s_activeIDsByWorkflowKey[workflowKey].add(workflowID);
    }
    s_ownerWorkflowIDs[msg.sender].add(workflowID);
    s_donWorkflowIDs[dl].add(workflowID);
    s_userDONCount[msg.sender][dl] += 1;
    s_tagsUsed[tagKey] = true;

    if (isCreate) {
      emit WorkflowRegisteredV2(workflowID, msg.sender, dl, status, workflowName);
    } else {
      emit WorkflowUpdatedV2(latestID, workflowID, msg.sender, dl, workflowName);
    }
  }

  function pauseWorkflow(
    bytes32 workflowID
  ) external {
    _updateStatus(workflowID, WorkflowStatus.PAUSED);
  }

  function activateWorkflow(
    bytes32 workflowID
  ) external {
    _updateStatus(workflowID, WorkflowStatus.ACTIVE);
  }

  function deleteWorkflow(
    bytes32 workflowID
  ) external {
    WorkflowMetadata storage meta = s_workflows[workflowID];
    if (meta.owner == address(0)) revert WorkflowDoesNotExist();
    if (meta.owner != msg.sender) revert CallerNotOwner();

    bytes32 dl = _stringToBytes32(meta.donLabel);

    // remove from indices & counters
    s_ownerWorkflowIDs[msg.sender].remove(workflowID);
    s_donWorkflowIDs[dl].remove(workflowID);
    s_userDONCount[msg.sender][dl] -= 1;

    bytes32 tagKey = keccak256(abi.encodePacked(meta.owner, meta.workflowName, meta.tag));
    delete s_tagsUsed[tagKey];

    emit WorkflowDeletedV2(workflowID, msg.sender, meta.donLabel, meta.workflowName);

    delete s_workflows[workflowID];
  }

  function pauseAllActiveWorkflowsByWorkflowKey(
    bytes32 workflowKey
  ) public {
    EnumerableSet.Bytes32Set storage activeSet = s_activeIDsByWorkflowKey[workflowKey];
    uint256 len = activeSet.length();
    if (len <= 1) return; // nothing else to pause

    for (uint256 i = 0; i < len; ++i) {
      bytes32 vID = activeSet.at(i);
      WorkflowMetadata storage vm = s_workflows[vID];
      vm.status = WorkflowStatus.PAUSED;
      bytes32 dl = _stringToBytes32(vm.donLabel);
      s_userDONCount[vm.owner][dl] -= 1;

      emit WorkflowPausedV2(vID, vm.owner, vm.donLabel, vm.workflowName);
      // defer removal to avoid enumerate‑while‑mutate issue.
      activeSet.remove(vID);
      --i;
      len = activeSet.length();
    }
  }

  // ================================================================
  // |                           Queries                            |
  // ================================================================
  function getWorkflowMetadata(
    bytes32 workflowID
  ) external view returns (WorkflowMetadata memory) {
    WorkflowMetadata memory meta = s_workflows[workflowID];
    if (meta.owner == address(0)) revert WorkflowDoesNotExist();
    return meta;
  }

  function getWorkflowMetadataListByOwner(
    address owner,
    uint256 start,
    uint256 limit
  ) external view returns (WorkflowMetadata[] memory list) {
    uint256 total = s_ownerWorkflowIDs[owner].length();
    if (start >= total) return new WorkflowMetadata[](0);
    if (limit == 0 || limit > MAX_PAGINATION_LIMIT) limit = MAX_PAGINATION_LIMIT;

    uint256 end = start + limit > total ? total : start + limit;
    uint256 len = end - start;
    list = new WorkflowMetadata[](len);

    for (uint256 i = 0; i < len; ++i) {
      bytes32 id = s_ownerWorkflowIDs[owner].at(start + i);
      list[i] = s_workflows[id];
    }
    return list;
  }

  function getWorkflowMetadataListByDON(
    bytes32 donLabel,
    uint256 start,
    uint256 limit
  ) external view returns (WorkflowMetadata[] memory list) {
    uint256 total = s_donWorkflowIDs[donLabel].length();
    if (start >= total) return new WorkflowMetadata[](0);
    if (limit == 0 || limit > MAX_PAGINATION_LIMIT) limit = MAX_PAGINATION_LIMIT;

    uint256 end = start + limit > total ? total : start + limit;
    uint256 len = end - start;
    list = new WorkflowMetadata[](len);

    for (uint256 i = 0; i < len; ++i) {
      bytes32 id = s_donWorkflowIDs[donLabel].at(start + i);
      list[i] = s_workflows[id];
    }
    return list;
  }

  /// @notice Returns the latest workflowID for a given (owner, name) combo
  function getLatestWorkflowID(address owner, string calldata name) external view returns (bytes32) {
    return s_workflowKeyToLatestID[computeHashKey(owner, name)];
  }

  /// @notice Returns the latest full WorkflowMetadata for a given (owner, name)
  function getLatestWorkflowMetadata(
    address owner,
    string calldata name
  ) external view returns (WorkflowMetadata memory) {
    bytes32 key = computeHashKey(owner, name);
    bytes32 latestID = s_workflowKeyToLatestID[key];
    if (latestID == bytes32(0)) revert WorkflowDoesNotExist();
    return s_workflows[latestID];
  }

  // ================================================================
  //                        Internal Helpers                        |
  // ================================================================
  function _updateStatus(bytes32 workflowID, WorkflowStatus newStatus) internal {
    WorkflowMetadata storage meta = s_workflows[workflowID];
    if (meta.owner == address(0)) revert WorkflowDoesNotExist();
    if (meta.owner != msg.sender) revert CallerNotOwner();
    if (meta.status == newStatus) revert WorkflowAlreadyInDesiredStatus();

    bytes32 dl = _stringToBytes32(meta.donLabel);

    // Update counters based on status change
    if (newStatus == WorkflowStatus.ACTIVE) {
      // Enforce limits before activation
      _enforceLimits(msg.sender, dl);

      s_userDONCount[meta.owner][dl] += 1;
      s_activeIDsByWorkflowKey[computeHashKey(meta.owner, meta.workflowName)].add(workflowID);
      emit WorkflowActivatedV2(workflowID, meta.owner, meta.donLabel, meta.workflowName);
    } else {
      // Decrement count when transitioning from ACTIVE to PAUSED

      s_userDONCount[meta.owner][dl] -= 1;
      s_activeIDsByWorkflowKey[computeHashKey(meta.owner, meta.workflowName)].remove(workflowID);

      emit WorkflowPausedV2(workflowID, meta.owner, meta.donLabel, meta.workflowName);
    }

    meta.status = newStatus;
  }

  function _enforceLimits(address owner, bytes32 donLabel) internal view {
    if (s_donWorkflowIDs[donLabel].length() >= getMaxWorkflowsPerDON(donLabel)) {
      revert MaxWorkflowsPerDONExceeded(donLabel);
    }
    if (s_userDONCount[owner][donLabel] >= getMaxWorkflowsPerUserDON(owner, donLabel)) {
      revert MaxWorkflowsPerUserDONExceeded(owner, donLabel);
    }
  }

  function _validateWorkflowURLs(uint256 binaryLen, uint256 configLen) internal pure {
    if (binaryLen > MAX_URL_LENGTH) revert URLTooLong(binaryLen, MAX_URL_LENGTH);
    if (configLen > MAX_URL_LENGTH) revert URLTooLong(configLen, MAX_URL_LENGTH);
  }

  function _validateWorkflowName(
    uint256 nameLen
  ) internal pure {
    if (nameLen == 0) revert WorkflowNameRequired();
    if (nameLen > MAX_WORKFLOW_NAME_LENGTH) revert WorkflowNameTooLong(nameLen, MAX_WORKFLOW_NAME_LENGTH);
  }

  /// @notice Ensures the given workflowID is unique and non-zero.
  /// @param workflowID The workflowID to validate and consume.
  function _validateWorkflowID(
    bytes32 workflowID
  ) internal view {
    if (workflowID == bytes32(0)) revert InvalidWorkflowID();

    if (s_workflows[workflowID].owner != address(0)) {
      revert WorkflowIDAlreadyExists();
    }
  }

  /// @dev Converts a ≤32-byte ASCII string to a bytes32 value by zero-padding on the right.
  /// This is used to store donLabel as a fixed-size key in mappings for gas efficiency,
  /// while preserving the original string in storage. Assembly is used for direct
  /// calldata access to avoid unnecessary memory allocation and reduce gas cost.
  function _stringToBytes32(
    string memory s
  ) internal pure returns (bytes32 out) {
    bytes memory b = bytes(s);
    if (b.length > 32) revert StringTooLong(32, b.length);
    assembly ("memory-safe") {
      out := mload(add(b, 32))
    }
    return out;
  }

  /// @notice Generates a `workflowKey` by combining the owner's address with a specific field.
  /// @param owner The address of the owner. Typically used to uniquely associate the field with the owner.
  /// @param field A string field, such as the workflow name or secrets URL, that is used to generate the unique hash.
  /// @return A unique `bytes32` hash computed from the combination of the owner's address and the given field.
  function computeHashKey(address owner, string memory field) public pure returns (bytes32) {
    return keccak256(abi.encodePacked(owner, field));
  }
}
