// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {Ownable2StepMsgSender} from "../../../shared/access/Ownable2StepMsgSender.sol";
import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";
import {ICapabilityConfiguration} from "./interfaces/ICapabilityConfiguration.sol";
import {INodeInfoProvider} from "./interfaces/INodeInfoProvider.sol";

import {ERC165Checker} from "@openzeppelin/contracts@4.8.3/utils/introspection/ERC165Checker.sol";
import {EnumerableSet} from "@openzeppelin/contracts@4.8.3/utils/structs/EnumerableSet.sol";

/// @notice CapabilitiesRegistry is used to manage Nodes (including their links to Node Operators), Capabilities,
/// and DONs (Decentralized Oracle Networks) which are sets of nodes that support those Capabilities.
/// @dev The contract currently stores the entire state of Node Operators, Nodes, Capabilities and DONs in the
/// contract and requires a full state migration if an upgrade is ever required. The team acknowledges this and is
/// fine reconfiguring the upgraded contract in the future so as to not add extra complexity to this current version.
// solhint-disable-next-line max-states-count
contract CapabilitiesRegistry is INodeInfoProvider, Ownable2StepMsgSender, ITypeAndVersion {
  // Add the library methods
  using EnumerableSet for EnumerableSet.Bytes32Set;
  using EnumerableSet for EnumerableSet.UintSet;

  // ================================================================
  // |                         Structs                               |
  // ================================================================

  struct NodeOperator {
    /// @notice The address of the admin that can manage a node operator
    address admin;
    /// @notice Human readable name of a Node Operator managing the node
    /// @dev The contract does not validate the length or characters of the node operator name because
    /// a trusted admin will supply these names. We reduce gas costs by omitting these checks on-chain.
    string name;
  }

  struct NodeParams {
    /// @notice The id of the node operator that manages this node
    uint32 nodeOperatorId;
    /// @notice The signer address for application-layer message verification.
    bytes32 signer;
    /// @notice This is an Ed25519 public key that is used to identify a node. This key is guaranteed to
    /// be unique in the CapabilitiesRegistry. It is used to identify a node in the the P2P network.
    bytes32 p2pId;
    /// @notice Public key used to encrypt secrets for this node
    bytes32 encryptionPublicKey;
    /// @notice CSA (Centralized Server Authentication) public key used as identity to non-P2P networks.
    bytes32 csaKey;
    /// @notice The list of capability IDs supported by the node
    string[] capabilityIds;
  }

  struct Node {
    /// @notice The node's parameters
    /// @notice The id of the node operator that manages this node
    uint32 nodeOperatorId;
    /// @notice The number of times the node's configuration has been updated
    uint32 configCount;
    /// @notice The ID of the Workflow DON that the node belongs to. A node can
    /// only belong to one DON that accepts Workflows.
    uint32 workflowDONId;
    /// @notice The signer address for application-layer message verification.
    /// @dev This key is guaranteed to be unique in the CapabilitiesRegistry as a signer
    /// address can only belong to one node.
    /// @dev This should be the ABI encoded version of the node's address. I.e 0x0000address. The Capability Registry
    /// does not store it as an address so that non EVM chains with addresses greater than 20 bytes can be supported
    /// in the future.
    bytes32 signer;
    /// @notice This is an Ed25519 public key that is used to identify a node. This key is guaranteed
    /// to be unique in the CapabilitiesRegistry. It is used to identify a node in the the P2P network.
    bytes32 p2pId;
    /// @notice Public key used to encrypt secrets for this node
    bytes32 encryptionPublicKey;
    /// @notice CSA (Centralized Server Authentication) public key used as identity to non-P2P networks.
    bytes32 csaKey;
    /// @notice The node's supported capabilities
    /// @dev This is stored as a map so that we can easily update to a set of new capabilities by
    /// incrementing the configCount and creating a new set of supported capability IDs
    mapping(uint32 configCount => EnumerableSet.Bytes32Set capabilityId) supportedHashedCapabilityIds;
    /// @notice The list of capabilities DON Ids supported by the node. A node can belong to multiple
    /// capabilities DONs. This list does not include a Workflow DON id if the node belongs to one.
    EnumerableSet.UintSet capabilitiesDONIds;
  }

  /// @notice Capability is a struct that holds the capability information
  /// It is an input struct for the `addCapability` function.
  struct Capability {
    /// @notice The capability ID
    /// @dev Example: "data-streams-reports@1.0.0"
    string capabilityId;
    /// @notice An address to the capability configuration contract. Having this defined on a capability enforces
    /// consistent configuration across DON instances serving the same capability. Configuration contract MUST implement
    /// CapabilityConfigurationContractInterface.
    ///
    /// @dev The main use cases are:
    /// 1) Sharing capability configuration across DON instances
    /// 2) Inspect and modify on-chain configuration without off-chain capability code.
    ///
    /// It is not recommended to store configuration which requires knowledge of the DON membership.
    address configurationContract;
    /// @notice Metadata for the capability. This is used to store additional information about the capability.
    bytes metadata;
  }

  /// @notice CapabilityInfo is a struct that holds the capability information
  /// It is an output struct for the `getCapabilities` function.
  struct CapabilityInfo {
    /// @notice The capability ID
    /// @dev Example: "data-streams-reports@1.0.0"
    string capabilityId;
    /// @notice An address to the capability configuration contract. Having this defined on a capability enforces
    /// consistent configuration across DON instances serving the same capability. Configuration contract MUST implement
    /// CapabilityConfigurationContractInterface.
    ///
    /// @dev The main use cases are:
    /// 1) Sharing capability configuration across DON instances
    /// 2) Inspect and modify on-chain configuration without off-chain capability code.
    ///
    /// It is not recommended to store configuration which requires knowledge of the DON membership.
    address configurationContract;
    /// @notice True if the capability is deprecated
    bool isDeprecated;
    /// @notice Metadata for the capability. This is used to store additional information about the capability.
    bytes metadata;
  }

  /// @notice CapabilityConfiguration is a struct that holds the capability configuration
  /// for a specific DON
  struct CapabilityConfiguration {
    /// @notice The capability Id
    string capabilityId;
    /// @notice The capability config specific to a DON.  This will be decoded offchain
    bytes config;
  }

  /// @notice MutableDONConfig is a struct that holds the configuration for a
  /// specific DON. It is used to store the configuration for a DON that can be
  /// updated.
  struct MutableDONConfig {
    /// @notice The set of p2pIds of nodes that belong to this DON. A node (the same p2pId) can belong to multiple DONs.
    EnumerableSet.Bytes32Set nodes;
    /// @notice The general config for the DON. This holds general DON config that is not
    /// specific to a capability.
    bytes config;
    /// @notice The set of capabilityIds
    bytes32[] capabilityIds;
    /// @notice True if the DON is public. A public DON means that it accepts
    /// external capability requests
    bool isPublic;
    /// @notice The f value for the DON.  This is the number of faulty nodes
    /// that the DON can tolerate. This can be different from the f value of
    /// the OCR instances that capabilities spawn.
    uint8 f;
    /// @notice The name of the DON. Can be empty. If not empty, must be unique
    /// to the registry.
    string name;
    /// @notice Mapping from capability IDs to configs
    mapping(string capabilityId => bytes config) capabilityConfigs;
  }

  /// @notice DON (Decentralized Oracle Network) is a grouping of nodes that support
  // the same capabilities.
  struct DON {
    /// @notice Computed. Auto-increment.
    uint32 id;
    /// @notice The number of times the DON was configured
    uint32 configCount;
    /// @notice True if the DON accepts Workflows. A DON that accepts Workflows
    /// is called Workflow DON and it can execute Workflows. A Workflow DON can
    /// also support capabilities.
    bool acceptsWorkflows;
    /// @notice Mapping of config counts to configurations
    mapping(uint32 configCount => MutableDONConfig donConfig) config;
  }

  struct DONInfo {
    /// @notice Computed. Auto-increment.
    uint32 id;
    /// @notice The number of times the DON was configured
    uint32 configCount;
    /// @notice The f value for the DON.  This is the number of faulty nodes
    /// that the DON can tolerate. This can be different from the f value of
    /// the OCR instances that capabilities spawn.
    uint8 f;
    /// @notice True if the DON is public.  A public DON means that it accepts
    /// external capability requests
    bool isPublic;
    /// @notice True if the DON accepts Workflows.
    bool acceptsWorkflows;
    /// @notice List of member node P2P Ids
    bytes32[] nodeP2PIds;
    /// @notice The families the DON belongs to. A DON family is a group of DONs
    /// that are connected with each other. Empty string is the default family.
    string[] donFamilies;
    /// @notice The name of the DON. Must be unique to the registry.
    string name;
    /// @notice The config for the DON. This holds general DON config that is not
    /// specific to a capability.
    bytes config;
    /// @notice List of capability configurations
    CapabilityConfiguration[] capabilityConfigurations;
  }

  /// @notice DONParams is a struct that holds the parameters for a DON.
  /// @dev This is needed to avoid "stack too deep" errors in _setDONConfig.
  struct DONParams {
    uint32 id;
    uint32 configCount;
    bool isPublic;
    bool acceptsWorkflows;
    uint8 f;
    string name;
    bytes config;
  }

  /// @notice NewDONParams is a struct that holds the parameters for a new DON.
  struct NewDONParams {
    string name;
    string[] donFamilies;
    bytes config;
    CapabilityConfiguration[] capabilityConfigurations;
    bytes32[] nodes;
    uint8 f;
    bool isPublic;
    bool acceptsWorkflows;
  }

  /// @notice UpdateDONParams is a struct that holds the parameters for updating a DON.
  struct UpdateDONParams {
    string name;
    bytes config;
    CapabilityConfiguration[] capabilityConfigurations;
    bytes32[] nodes;
    uint8 f;
    bool isPublic;
  }

  /// @notice ConstructorParams is a struct that holds the parameters for the constructor
  struct ConstructorParams {
    /// @notice Whether to allow DONs with a single node. Used only for testing.
    bool canAddOneNodeDONs;
  }

  // ================================================================
  // |                         Errors                               |
  // ================================================================

  /// @notice This error is thrown when a caller is not allowed
  /// to execute the transaction
  /// @param sender The address that tried to execute the transaction
  error AccessForbidden(address sender);

  /// @notice This error is thrown when there is a mismatch between
  /// array arguments
  /// @param lengthOne The length of the first array argument
  /// @param lengthTwo The length of the second array argument
  error LengthMismatch(uint256 lengthOne, uint256 lengthTwo);

  /// @notice This error is thrown when trying to set a node operator's
  /// admin address to the zero address
  error InvalidNodeOperatorAdmin();

  /// @notice This error is thrown when trying to add a node with P2P ID that
  /// is empty bytes
  /// @param p2pId The provided P2P ID
  error InvalidNodeP2PId(bytes32 p2pId);

  /// @notice This error is thrown when trying to add a node without
  /// including the encryption public key bytes.
  /// @param encryptionPublicKey The encryption public key bytes
  error InvalidNodeEncryptionPublicKey(bytes32 encryptionPublicKey);

  /// @notice This error is thrown when trying to add a node without
  /// including the CSA public key bytes.
  /// @param csaKey The CSA public key bytes
  error InvalidNodeCSAKey(bytes32 csaKey);

  /// @notice This error is thrown when trying to add a node without
  /// capabilities or with capabilities that do not exist.
  /// @param capabilityIds The IDs of the capabilities that are being added.
  error InvalidNodeCapabilities(string[] capabilityIds);

  /// @notice This error is emitted when a DON does not exist
  /// @param donId The ID of the nonexistent DON
  error DONDoesNotExist(uint32 donId);

  /// @notice This error is emitted when a DON with the given name does not exist
  /// @param donName The name of the nonexistent DON
  error DONWithNameDoesNotExist(string donName);

  /// @notice This error is emitted when trying to set the name of a DON to an empty string
  /// @param donId The ID of the DON
  error DONNameCannotBeEmpty(uint32 donId);

  /// @notice This error is thrown when trying to set the node's
  /// signer address to zero or if the signer address has already
  /// been used by another node
  error InvalidNodeSigner();

  /// @notice This error is thrown when trying to add a capability that already
  /// exists.
  /// @param capabilityId The capability ID of the capability
  /// that already exists
  error CapabilityAlreadyExists(string capabilityId);

  /// @notice This error is thrown when trying to add a node that already
  /// exists.
  /// @param nodeP2PId The P2P ID of the node that already exists
  error NodeAlreadyExists(bytes32 nodeP2PId);

  /// @notice This error is thrown when trying to add a node to a DON where
  /// the node does not support the capability
  /// @param nodeP2PId The P2P ID of the node
  /// @param capabilityId The ID of the capability
  error NodeDoesNotSupportCapability(bytes32 nodeP2PId, string capabilityId);

  /// @notice This error is thrown when trying to add a capability configuration
  /// for a capability that was already configured on a DON
  /// @param donId The ID of the DON that the capability was configured for
  /// @param capabilityId The ID of the capability that was configured
  error DuplicateDONCapability(uint32 donId, string capabilityId);

  /// @notice This error is thrown when trying to add a duplicate node to a DON
  /// @param donId The ID of the DON that the node was added for
  /// @param nodeP2PId The P2P ID of the node
  error DuplicateDONNode(uint32 donId, bytes32 nodeP2PId);

  /// @notice This error is thrown when trying to configure a DON with invalid
  /// fault tolerance value.
  /// @param f The proposed fault tolerance value
  /// @param nodeCount The proposed number of nodes in the DON
  error InvalidFaultTolerance(uint8 f, uint256 nodeCount);

  /// @notice This error is thrown when a capability with the provided ID is
  /// not found.
  /// @param capabilityId The ID used for the lookup.
  error CapabilityDoesNotExist(string capabilityId);

  /// @notice This error is thrown when trying to deprecate a capability that
  /// is deprecated.
  /// @param capabilityId The ID of the capability that is deprecated.
  error CapabilityIsDeprecated(string capabilityId);

  /// @notice This error is thrown when a node operator does not exist
  /// @param nodeOperatorId The ID of the node operator that does not exist
  error NodeOperatorDoesNotExist(uint32 nodeOperatorId);

  /// @notice This error is thrown when a node operator already exists
  /// @param existingNodeOperatorId The ID of the node operator that already exists
  error NodeOperatorAlreadyExists(uint32 existingNodeOperatorId);

  /// @notice This error is thrown when trying to remove a node that is still
  /// part of a capabilities DON
  /// @param donId The Id of the DON the node belongs to
  /// @param nodeP2PId The P2P Id of the node being removed
  error NodePartOfCapabilitiesDON(uint32 donId, bytes32 nodeP2PId);

  /// @notice This error is thrown when attempting to add a node to a second
  /// Workflow DON or when trying to remove a node that belongs to a Workflow
  /// DON
  /// @param donId The Id of the DON the node belongs to
  /// @param nodeP2PId The P2P Id of the node
  error NodePartOfWorkflowDON(uint32 donId, bytes32 nodeP2PId);

  /// @notice This error is thrown when removing a capability from the node
  /// when that capability is still required by one of the DONs the node
  /// belongs to.
  /// @param capabilityId The ID of the capability
  /// @param donId The ID of the DON that requires the capability
  error CapabilityRequiredByDON(string capabilityId, uint32 donId);

  /// @notice This error is thrown when trying to add a capability with a
  /// configuration contract that does not implement the required interface.
  /// @param proposedConfigurationContract The address of the proposed
  /// configuration contract.
  error InvalidCapabilityConfigurationContractInterface(address proposedConfigurationContract);

  /// @notice This error is thrown when trying to add a DON with a name that
  /// is already taken
  /// @param name The name of the DON that is already taken
  error DONNameAlreadyTaken(string name);

  /// @notice This error is thrown when trying to get a historical DON config
  /// that does not exist
  /// @param donId The ID of the DON
  /// @param maxConfigCount The current config count of the DON
  /// @param requestedConfigCount The requested config count
  error DONConfigDoesNotExist(uint32 donId, uint32 maxConfigCount, uint32 requestedConfigCount);

  // ================================================================
  // |                         Events                               |
  // ================================================================

  /// @notice This event is emitted when a new node is added
  /// @param p2pId The P2P ID of the node
  /// @param nodeOperatorId The ID of the node operator that manages this node
  /// @param signer The encoded node's signer address
  event NodeAdded(bytes32 p2pId, uint32 indexed nodeOperatorId, bytes32 signer);

  /// @notice This event is emitted when a node is removed
  /// @param p2pId The P2P ID of the node that was removed
  event NodeRemoved(bytes32 p2pId);

  /// @notice This event is emitted when a node is updated
  /// @param p2pId The P2P ID of the node
  /// @param nodeOperatorId The ID of the node operator that manages this node
  /// @param signer The node's signer address
  event NodeUpdated(bytes32 p2pId, uint32 indexed nodeOperatorId, bytes32 signer);

  /// @notice This event is emitted when a DON's config is set
  /// @param donId The ID of the DON the config was set for
  /// @param configCount The number of times the DON has been
  /// configured
  event ConfigSet(uint32 indexed donId, uint32 configCount);

  /// @notice This event is emitted when a new node operator is added
  /// @param nodeOperatorId The ID of the newly added node operator
  /// @param admin The address of the admin that can manage the node
  /// operator
  /// @param name The human readable name of the node operator
  event NodeOperatorAdded(uint32 indexed nodeOperatorId, address indexed admin, string name);

  /// @notice This event is emitted when a node operator is removed
  /// @param nodeOperatorId The ID of the node operator that was removed
  event NodeOperatorRemoved(uint32 indexed nodeOperatorId);

  /// @notice This event is emitted when a node operator is updated
  /// @param nodeOperatorId The ID of the node operator that was updated
  /// @param admin The address of the node operator's admin
  /// @param name The node operator's human readable name
  event NodeOperatorUpdated(uint32 indexed nodeOperatorId, address indexed admin, string name);

  /// @notice This event is emitted when a new capability is added
  /// @param capabilityId The ID of the newly added capability
  event CapabilityConfigured(string indexed capabilityId);

  /// @notice This event is emitted when a capability is deprecated
  /// @param capabilityId The ID of the deprecated capability
  event CapabilityDeprecated(string indexed capabilityId);

  /// @notice This event is emitted when a DON family is set
  /// @param donId The ID of the DON whose family was set
  /// @param donFamily The family name that was set
  event DONAddedToFamily(uint32 indexed donId, string indexed donFamily);

  /// @notice This event is emitted when a DON is removed from a family
  /// @param donId The ID of the DON that was removed from the family
  /// @param donFamily The family name that the DON was removed from
  event DONRemovedFromFamily(uint32 indexed donId, string indexed donFamily);

  // ================================================================
  // |                 Internal variables                            |
  // ================================================================

  string public constant override typeAndVersion = "CapabilitiesRegistry 2.0.0-dev";

  /// @notice Mapping of DON names to boolean indicating if the name is taken
  mapping(string donName => uint32 donId) private s_donNameToId;

  /// @notice Mapping of capabilities
  mapping(bytes32 hashedCapabilityId => Capability capability) private s_capabilities;

  /// @notice Set of hashed capability IDs.
  /// A hashed ID is created using the `_hash` function.
  EnumerableSet.Bytes32Set private s_hashedCapabilityIds;

  /// @notice Set of deprecated hashed capability IDs.
  /// A hashed ID is created using the `_hash` function.
  EnumerableSet.Bytes32Set private s_deprecatedHashedCapabilityIds;

  /// @notice Encoded node signer addresses
  EnumerableSet.Bytes32Set private s_nodeSigners;

  /// @notice Set of node P2P IDs
  EnumerableSet.Bytes32Set private s_nodeP2PIds;

  /// @notice Set of active DON family names that have at least one DON in the
  /// registry with that family name.
  EnumerableSet.Bytes32Set private s_activeDONFamilyNames;

  /// @notice Mapping of node operators
  mapping(uint32 nodeOperatorId => NodeOperator nodeOperator) private s_nodeOperators;

  /// @notice Mapping of node operator data hash to node operator ID. Used to maintain unique node operators.
  mapping(bytes32 nodeOperatorDataHash => uint32 nodeOperatorId) private s_nodeOperatorDataHashToId;

  /// @notice Mapping of nodes
  mapping(bytes32 p2pId => Node node) private s_nodes;

  /// @notice Mapping of DON IDs to DONs
  mapping(uint32 donId => DON don) private s_dons;

  /// @notice Mapping of DON ID to DON families. A single DON can belong to
  /// multiple families. Empty string is the default family.
  mapping(uint32 donId => EnumerableSet.Bytes32Set donFamilies) private s_donIdToDonFamilyHashes;

  /// @notice Mapping of DON family name hashes to DON IDs
  mapping(bytes32 donFamilyHash => EnumerableSet.UintSet donIds) private s_donFamilyMembers;

  /// @notice Mapping of DON family name hashes to DON family names
  mapping(bytes32 donFamilyHash => string donFamily) private s_donFamilyHashToDonFamily;

  /// @notice Mapping of hash of capability ID to capability ID
  mapping(bytes32 hashedCapabilityId => string capabilityId) private s_hashedCapabilityIdToCapabilityId;

  /// @notice The next ID to assign a new node operator to
  /// @dev Starting with 1 to avoid confusion with the zero value
  /// @dev No getter for this as this is an implementation detail
  uint32 private s_nextNodeOperatorId = 1;

  /// @notice The next ID to assign a new DON to
  /// @dev Starting with 1 to avoid confusion with the zero value
  uint32 private s_nextDONId = 1;

  /// @notice Whether to allow DONs with a single node. Used only for testing.
  bool private immutable i_canAddOneNodeDONs;

  constructor(
    ConstructorParams memory params
  ) {
    i_canAddOneNodeDONs = params.canAddOneNodeDONs;
  }

  // ================================================================
  // |                   External functions                         |
  // ================================================================

  /// @notice Adds a list of node operators
  /// @param nodeOperators List of node operators to add
  function addNodeOperators(
    NodeOperator[] calldata nodeOperators
  ) external onlyOwner {
    for (uint256 i; i < nodeOperators.length; ++i) {
      NodeOperator memory nodeOperator = nodeOperators[i];
      if (nodeOperator.admin == address(0)) revert InvalidNodeOperatorAdmin();
      uint32 nodeOperatorId = s_nextNodeOperatorId;
      s_nodeOperators[nodeOperatorId] = NodeOperator({admin: nodeOperator.admin, name: nodeOperator.name});
      bytes32 nodeOperatorDataHash = _nodeOperatorHash(nodeOperator);
      if (s_nodeOperatorDataHashToId[nodeOperatorDataHash] != 0) {
        revert NodeOperatorAlreadyExists(s_nodeOperatorDataHashToId[nodeOperatorDataHash]);
      }
      s_nodeOperatorDataHashToId[nodeOperatorDataHash] = nodeOperatorId;
      ++s_nextNodeOperatorId;
      emit NodeOperatorAdded(nodeOperatorId, nodeOperator.admin, nodeOperator.name);
    }
  }

  /// @notice Removes a node operator
  /// @param nodeOperatorIds The IDs of the node operators to remove
  function removeNodeOperators(
    uint32[] calldata nodeOperatorIds
  ) external onlyOwner {
    for (uint32 i; i < nodeOperatorIds.length; ++i) {
      uint32 nodeOperatorId = nodeOperatorIds[i];
      NodeOperator memory nodeOperator = s_nodeOperators[nodeOperatorId];
      bytes32 nodeOperatorDataHash = _nodeOperatorHash(nodeOperator);
      s_nodeOperatorDataHashToId[nodeOperatorDataHash] = 0;
      delete s_nodeOperators[nodeOperatorId];
      emit NodeOperatorRemoved(nodeOperatorId);
    }
  }

  /// @notice Updates a node operator
  /// @param nodeOperatorIds The ID of the node operator being updated
  /// @param nodeOperators The updated node operator params
  function updateNodeOperators(uint32[] calldata nodeOperatorIds, NodeOperator[] calldata nodeOperators) external {
    if (nodeOperatorIds.length != nodeOperators.length) {
      revert LengthMismatch(nodeOperatorIds.length, nodeOperators.length);
    }

    address owner = owner();
    for (uint256 i; i < nodeOperatorIds.length; ++i) {
      uint32 nodeOperatorId = nodeOperatorIds[i];

      NodeOperator storage currentNodeOperator = s_nodeOperators[nodeOperatorId];
      if (currentNodeOperator.admin == address(0)) revert NodeOperatorDoesNotExist(nodeOperatorId);

      NodeOperator memory nodeOperator = nodeOperators[i];
      if (nodeOperator.admin == address(0)) revert InvalidNodeOperatorAdmin();
      if (msg.sender != currentNodeOperator.admin && msg.sender != owner) revert AccessForbidden(msg.sender);

      bytes32 nodeOperatorDataHash = _nodeOperatorHash(nodeOperator);
      if (s_nodeOperatorDataHashToId[nodeOperatorDataHash] != 0) {
        revert NodeOperatorAlreadyExists(s_nodeOperatorDataHashToId[nodeOperatorDataHash]);
      }
      bytes32 currentNodeOperatorDataHash = _nodeOperatorHash(currentNodeOperator);
      delete s_nodeOperatorDataHashToId[currentNodeOperatorDataHash];
      s_nodeOperatorDataHashToId[nodeOperatorDataHash] = nodeOperatorId;

      if (
        currentNodeOperator.admin != nodeOperator.admin
          || keccak256(abi.encode(currentNodeOperator.name)) != keccak256(abi.encode(nodeOperator.name))
      ) {
        currentNodeOperator.admin = nodeOperator.admin;
        currentNodeOperator.name = nodeOperator.name;
        emit NodeOperatorUpdated(nodeOperatorId, nodeOperator.admin, nodeOperator.name);
      }
    }
  }

  /// @notice Adds nodes. Nodes can be added with deprecated capabilities to
  /// avoid breaking changes when deprecating capabilities.
  /// @param nodes The nodes to add
  function addNodes(
    NodeParams[] calldata nodes
  ) external {
    bool isOwner = msg.sender == owner();
    for (uint256 i; i < nodes.length; ++i) {
      NodeParams memory node = nodes[i];

      NodeOperator memory nodeOperator = s_nodeOperators[node.nodeOperatorId];
      if (nodeOperator.admin == address(0)) revert NodeOperatorDoesNotExist(node.nodeOperatorId);
      if (!isOwner && msg.sender != nodeOperator.admin) revert AccessForbidden(msg.sender);

      Node storage storedNode = s_nodes[node.p2pId];
      if (storedNode.signer != bytes32("")) revert NodeAlreadyExists(node.p2pId);
      if (node.p2pId == bytes32("")) revert InvalidNodeP2PId(node.p2pId);

      if (node.signer == bytes32("") || s_nodeSigners.contains(node.signer)) revert InvalidNodeSigner();

      if (node.encryptionPublicKey == bytes32("")) {
        revert InvalidNodeEncryptionPublicKey(node.encryptionPublicKey);
      }

      if (node.csaKey == bytes32("")) {
        revert InvalidNodeCSAKey(node.csaKey);
      }

      string[] memory capabilityIds = node.capabilityIds;
      if (capabilityIds.length == 0) revert InvalidNodeCapabilities(capabilityIds);

      ++storedNode.configCount;

      uint32 capabilityConfigCount = storedNode.configCount;
      for (uint256 j; j < capabilityIds.length; ++j) {
        bytes32 hashedCapabilityId = _hash(capabilityIds[j]);
        if (!s_hashedCapabilityIds.contains(hashedCapabilityId)) {
          revert InvalidNodeCapabilities(capabilityIds);
        }
        storedNode.supportedHashedCapabilityIds[capabilityConfigCount].add(hashedCapabilityId);
      }

      storedNode.encryptionPublicKey = node.encryptionPublicKey;
      storedNode.csaKey = node.csaKey;
      storedNode.nodeOperatorId = node.nodeOperatorId;
      storedNode.p2pId = node.p2pId;
      storedNode.signer = node.signer;
      s_nodeSigners.add(node.signer);
      s_nodeP2PIds.add(node.p2pId);
      emit NodeAdded(node.p2pId, node.nodeOperatorId, node.signer);
    }
  }

  /// @notice Removes nodes.  The node operator admin or contract owner
  /// can remove nodes
  /// @param removedNodeP2PIds The P2P Ids of the nodes to remove
  function removeNodes(
    bytes32[] calldata removedNodeP2PIds
  ) external {
    bool isOwner = msg.sender == owner();
    for (uint256 i; i < removedNodeP2PIds.length; ++i) {
      bytes32 p2pId = removedNodeP2PIds[i];

      Node storage node = s_nodes[p2pId];

      if (node.signer == bytes32("")) revert NodeDoesNotExist(p2pId);
      if (node.capabilitiesDONIds.length() > 0) {
        // Showing the first DON ID for the node as the node. Users can fetch
        // node info to get the full list of DONs.
        revert NodePartOfCapabilitiesDON(uint32(node.capabilitiesDONIds.at(0)), p2pId);
      }
      if (node.workflowDONId != 0) revert NodePartOfWorkflowDON(node.workflowDONId, p2pId);

      if (!isOwner && msg.sender != s_nodeOperators[node.nodeOperatorId].admin) {
        revert AccessForbidden(msg.sender);
      }
      s_nodeSigners.remove(node.signer);
      s_nodeP2PIds.remove(node.p2pId);
      delete s_nodes[p2pId];
      emit NodeRemoved(p2pId);
    }
  }

  /// @notice Updates nodes.  The node admin can update the node's signer address
  /// and reconfigure its supported capabilities
  /// @param nodes The nodes to update
  function updateNodes(
    NodeParams[] calldata nodes
  ) external {
    bool isOwner = msg.sender == owner();
    for (uint256 i; i < nodes.length; ++i) {
      NodeParams memory node = nodes[i];
      Node storage storedNode = s_nodes[node.p2pId];
      NodeOperator memory nodeOperator = s_nodeOperators[storedNode.nodeOperatorId];

      if (storedNode.signer == bytes32("")) revert NodeDoesNotExist(node.p2pId);
      if (!isOwner && msg.sender != nodeOperator.admin) revert AccessForbidden(msg.sender);

      if (node.signer == bytes32("")) revert InvalidNodeSigner();

      bytes32 previousSigner = storedNode.signer;
      if (previousSigner != node.signer) {
        if (s_nodeSigners.contains(node.signer)) revert InvalidNodeSigner();
        storedNode.signer = node.signer;
        s_nodeSigners.remove(previousSigner);
        s_nodeSigners.add(node.signer);
      }

      if (node.encryptionPublicKey == bytes32("")) {
        revert InvalidNodeEncryptionPublicKey(node.encryptionPublicKey);
      }

      if (node.csaKey == bytes32("")) {
        revert InvalidNodeCSAKey(node.csaKey);
      }

      string[] memory capabilityIds = node.capabilityIds;
      if (capabilityIds.length == 0) revert InvalidNodeCapabilities(capabilityIds);

      uint32 capabilityConfigCount = ++storedNode.configCount;
      for (uint256 j; j < capabilityIds.length; ++j) {
        if (!s_hashedCapabilityIds.contains(_hash(capabilityIds[j]))) {
          revert InvalidNodeCapabilities(capabilityIds);
        }
        storedNode.supportedHashedCapabilityIds[capabilityConfigCount].add(_hash(capabilityIds[j]));
      }

      // Validate that capabilities required by a Workflow DON are still supported
      uint32 nodeWorkflowDONId = storedNode.workflowDONId;
      if (nodeWorkflowDONId != 0) {
        bytes32[] memory workflowDonCapabilityIds =
          s_dons[nodeWorkflowDONId].config[s_dons[nodeWorkflowDONId].configCount].capabilityIds;

        for (uint256 j; j < workflowDonCapabilityIds.length; ++j) {
          if (!storedNode.supportedHashedCapabilityIds[capabilityConfigCount].contains(workflowDonCapabilityIds[j])) {
            revert CapabilityRequiredByDON(
              s_hashedCapabilityIdToCapabilityId[workflowDonCapabilityIds[j]], nodeWorkflowDONId
            );
          }
        }
      }

      // Validate that capabilities required by capabilities DONs are still supported
      uint256[] memory capabilitiesDONIds = storedNode.capabilitiesDONIds.values();
      for (uint32 j; j < capabilitiesDONIds.length; ++j) {
        uint32 donId = uint32(capabilitiesDONIds[j]);
        bytes32[] memory donCapabilityIds = s_dons[donId].config[s_dons[donId].configCount].capabilityIds;

        for (uint256 k; k < donCapabilityIds.length; ++k) {
          if (!storedNode.supportedHashedCapabilityIds[capabilityConfigCount].contains(donCapabilityIds[k])) {
            revert CapabilityRequiredByDON(s_hashedCapabilityIdToCapabilityId[donCapabilityIds[k]], donId);
          }
        }
      }

      storedNode.nodeOperatorId = node.nodeOperatorId;
      storedNode.p2pId = node.p2pId;
      storedNode.encryptionPublicKey = node.encryptionPublicKey;
      storedNode.csaKey = node.csaKey;

      emit NodeUpdated(node.p2pId, node.nodeOperatorId, node.signer);
    }
  }

  /// @notice Adds a new capability to the capability registry
  /// @param capabilities The capabilities being added
  /// @dev There is no function to update capabilities as this would require
  /// nodes to trust that the capabilities they support are not updated by the
  /// admin
  function addCapabilities(
    Capability[] calldata capabilities
  ) external onlyOwner {
    for (uint256 i; i < capabilities.length; ++i) {
      Capability memory capability = capabilities[i];
      bytes32 hashedCapabilityId = _hash(capability.capabilityId);
      if (!s_hashedCapabilityIds.add(hashedCapabilityId)) {
        revert CapabilityAlreadyExists(capability.capabilityId);
      }
      s_hashedCapabilityIdToCapabilityId[hashedCapabilityId] = capability.capabilityId;

      if (capability.configurationContract != address(0)) {
        /// Check that the configuration contract being assigned
        /// correctly supports the ICapabilityConfiguration interface
        /// by implementing both getCapabilityConfiguration and
        /// beforeCapabilityConfigSet
        if (
          !ERC165Checker.supportsInterface(capability.configurationContract, type(ICapabilityConfiguration).interfaceId)
        ) revert InvalidCapabilityConfigurationContractInterface(capability.configurationContract);
      }
      s_capabilities[hashedCapabilityId] = capability;
      emit CapabilityConfigured(capability.capabilityId);
    }
  }

  /// @notice Deprecates a capability
  /// @param capabilityIds[] The IDs of the capabilities to deprecate
  function deprecateCapabilities(
    string[] calldata capabilityIds
  ) external onlyOwner {
    for (uint256 i; i < capabilityIds.length; ++i) {
      string memory capabilityId = capabilityIds[i];
      bytes32 hashedCapabilityId = _hash(capabilityId);
      if (!s_hashedCapabilityIds.contains(hashedCapabilityId)) revert CapabilityDoesNotExist(capabilityId);
      if (!s_deprecatedHashedCapabilityIds.add(hashedCapabilityId)) revert CapabilityIsDeprecated(capabilityId);

      emit CapabilityDeprecated(capabilityId);
    }
  }

  /// @notice Adds a list of DONs
  /// @param newDONs The list of DONs to add
  /// @dev The DONs are added in the order they are provided in the `newDONs` array
  function addDONs(
    NewDONParams[] calldata newDONs
  ) external onlyOwner {
    if (newDONs.length == 0) return;

    for (uint256 i; i < newDONs.length; ++i) {
      NewDONParams memory newDON = newDONs[i];
      uint32 nextDONId = s_nextDONId++;

      _setDONConfig(
        newDON.nodes,
        newDON.capabilityConfigurations,
        DONParams({
          id: nextDONId,
          configCount: 1,
          isPublic: newDON.isPublic,
          acceptsWorkflows: newDON.acceptsWorkflows,
          f: newDON.f,
          name: newDON.name,
          config: newDON.config
        })
      );

      for (uint256 j; j < newDON.donFamilies.length; ++j) {
        _addDONToFamily(nextDONId, newDON.donFamilies[j]);
      }
    }
  }

  /// @notice Updates a DON's configuration.  This allows
  /// the admin to reconfigure the list of capabilities supported
  /// by the DON, the list of nodes that make up the DON as well
  /// as whether or not the DON can accept external workflows
  /// @param donId The ID of the DON to update
  /// @param updateDONParams The parameters for the DON to update
  function updateDON(uint32 donId, UpdateDONParams calldata updateDONParams) external onlyOwner {
    DON storage don = s_dons[donId];
    uint32 configCount = don.configCount;
    if (configCount == 0) revert DONDoesNotExist(donId);
    _setDONConfig(
      updateDONParams.nodes,
      updateDONParams.capabilityConfigurations,
      DONParams({
        id: donId,
        configCount: ++configCount,
        isPublic: updateDONParams.isPublic,
        acceptsWorkflows: don.acceptsWorkflows,
        f: updateDONParams.f,
        name: updateDONParams.name,
        config: updateDONParams.config
      })
    );
  }

  /// @notice Updates a DON's configuration by its name
  /// @param donName The name of the DON to update
  /// @param updateDONParams The parameters for the DON to update
  function updateDONByName(string calldata donName, UpdateDONParams calldata updateDONParams) external onlyOwner {
    uint32 donId = s_donNameToId[donName];
    if (donId == 0) revert DONWithNameDoesNotExist(donName);

    DON storage don = s_dons[donId];

    _setDONConfig(
      updateDONParams.nodes,
      updateDONParams.capabilityConfigurations,
      DONParams({
        id: donId,
        configCount: ++don.configCount,
        isPublic: updateDONParams.isPublic,
        acceptsWorkflows: don.acceptsWorkflows,
        f: updateDONParams.f,
        name: updateDONParams.name,
        config: updateDONParams.config
      })
    );
  }

  /// @notice Removes DONs from the Capability Registry
  /// @param donIds The IDs of the DON to be removed
  function removeDONs(
    uint32[] calldata donIds
  ) external onlyOwner {
    for (uint256 i; i < donIds.length; ++i) {
      uint32 donId = donIds[i];
      _removeDON(donId);
    }
  }

  /// @notice Removes DONs from the Capability Registry by their name
  /// @param donNames The names of the DONs to be removed
  function removeDONsByName(
    string[] calldata donNames
  ) external onlyOwner {
    for (uint256 i; i < donNames.length; ++i) {
      uint32 donId = s_donNameToId[donNames[i]];
      if (donId == 0) continue;
      _removeDON(donId);
    }
  }

  /// @notice Sets the DON family for a DON
  /// @param donId The ID of the DON to set the family for
  /// @param addToFamilies The families to add the DON to
  /// @param removeFromFamilies The families to remove the DON from
  function setDONFamilies(
    uint32 donId,
    string[] calldata addToFamilies,
    string[] calldata removeFromFamilies
  ) external onlyOwner {
    if (s_dons[donId].configCount == 0) revert DONDoesNotExist(donId);

    for (uint256 i; i < addToFamilies.length; ++i) {
      _addDONToFamily(donId, addToFamilies[i]);
    }

    for (uint256 i; i < removeFromFamilies.length; ++i) {
      bytes32 removeFromFamilyHash = _hash(removeFromFamilies[i]);

      // If the DON is not in the family, do nothing.
      // There is no point in erroring out as this is a no-op and the erroring
      // would not provide any value to the user.
      if (!s_donIdToDonFamilyHashes[donId].contains(removeFromFamilyHash)) return;

      _removeDONFromFamily(donId, removeFromFamilyHash);
    }
  }

  // ================================================================
  // |                      View functions                          |
  // ================================================================

  /// @notice Returns a Capability by its ID.
  function getCapability(
    string memory capabilityId
  ) public view returns (CapabilityInfo memory) {
    bytes32 hashedCapabilityId = _hash(capabilityId);
    return (
      CapabilityInfo({
        capabilityId: capabilityId,
        metadata: s_capabilities[hashedCapabilityId].metadata,
        configurationContract: s_capabilities[hashedCapabilityId].configurationContract,
        isDeprecated: s_deprecatedHashedCapabilityIds.contains(hashedCapabilityId)
      })
    );
  }

  /// @notice Returns all capabilities. This operation will copy capabilities
  /// to memory, which can be quite expensive. This is designed to mostly be
  /// used by view accessors that are queried without any gas fees.
  /// @return CapabilityInfo[] List of capabilities
  function getCapabilities() external view returns (CapabilityInfo[] memory) {
    bytes32[] memory capabilityIds = s_hashedCapabilityIds.values();
    CapabilityInfo[] memory capabilitiesInfo = new CapabilityInfo[](capabilityIds.length);

    for (uint256 i; i < capabilityIds.length; ++i) {
      capabilitiesInfo[i] = getCapability(s_hashedCapabilityIdToCapabilityId[capabilityIds[i]]);
    }
    return capabilitiesInfo;
  }

  /// @notice Returns whether a capability is deprecated
  /// @param capabilityId The ID of the capability to check
  /// @return bool True if the capability is deprecated, false otherwise
  function isCapabilityDeprecated(
    string calldata capabilityId
  ) external view returns (bool) {
    bytes32 hashedCapabilityId = _hash(capabilityId);
    return s_deprecatedHashedCapabilityIds.contains(hashedCapabilityId);
  }

  /// @notice Gets a node operator's data
  /// @param nodeOperatorId The ID of the node operator to query for
  /// @return NodeOperator The node operator data
  function getNodeOperator(
    uint32 nodeOperatorId
  ) external view returns (NodeOperator memory) {
    return s_nodeOperators[nodeOperatorId];
  }

  /// @notice Gets all node operators
  /// @return NodeOperator[] All node operators
  function getNodeOperators() external view returns (NodeOperator[] memory) {
    uint32 nodeOperatorId = s_nextNodeOperatorId;
    /// Minus one to account for s_nextNodeOperatorId starting at index 1
    NodeOperator[] memory nodeOperators = new NodeOperator[](s_nextNodeOperatorId - 1);
    uint256 idx;
    for (uint32 i = 1; i < nodeOperatorId; ++i) {
      if (s_nodeOperators[i].admin != address(0)) {
        nodeOperators[idx] = s_nodeOperators[i];
        ++idx;
      }
    }
    if (idx != s_nextNodeOperatorId - 1) {
      assembly {
        mstore(nodeOperators, idx)
      }
    }
    return nodeOperators;
  }

  /// @notice Gets the next node DON ID
  /// @return uint32 The next node DON ID
  function getNextDONId() external view returns (uint32) {
    return s_nextDONId;
  }

  /// @notice Gets a node's data
  /// @param p2pId The P2P ID of the node to query for
  /// @return nodeInfo NodeInfo The node data
  function getNode(
    bytes32 p2pId
  ) public view returns (NodeInfo memory nodeInfo) {
    bytes32[] memory capabilityIds = s_nodes[p2pId].supportedHashedCapabilityIds[s_nodes[p2pId].configCount].values();
    string[] memory capabilityIdsString = new string[](capabilityIds.length);
    for (uint256 i; i < capabilityIds.length; ++i) {
      capabilityIdsString[i] = s_hashedCapabilityIdToCapabilityId[capabilityIds[i]];
    }
    return (
      NodeInfo({
        nodeOperatorId: s_nodes[p2pId].nodeOperatorId,
        p2pId: s_nodes[p2pId].p2pId,
        signer: s_nodes[p2pId].signer,
        encryptionPublicKey: s_nodes[p2pId].encryptionPublicKey,
        csaKey: s_nodes[p2pId].csaKey,
        capabilityIds: capabilityIdsString,
        configCount: s_nodes[p2pId].configCount,
        workflowDONId: s_nodes[p2pId].workflowDONId,
        capabilitiesDONIds: s_nodes[p2pId].capabilitiesDONIds.values()
      })
    );
  }

  /// @notice Gets all nodes
  /// @return NodeInfo[] All nodes in the capability registry
  function getNodes() external view returns (NodeInfo[] memory) {
    bytes32[] memory p2pIds = s_nodeP2PIds.values();
    NodeInfo[] memory nodesInfo = new NodeInfo[](p2pIds.length);

    for (uint256 i; i < p2pIds.length; ++i) {
      nodesInfo[i] = getNode(p2pIds[i]);
    }
    return nodesInfo;
  }

  /// @notice Gets nodes by their P2P IDs
  /// @param p2pIds The P2P IDs of the nodes to query for
  /// @return NodeInfo[] The nodes data
  function getNodesByP2PIds(
    bytes32[] calldata p2pIds
  ) external view returns (NodeInfo[] memory) {
    NodeInfo[] memory nodesInfo = new NodeInfo[](p2pIds.length);

    for (uint256 i; i < p2pIds.length; ++i) {
      nodesInfo[i] = getNode(p2pIds[i]);
      if (nodesInfo[i].p2pId == bytes32("")) revert NodeDoesNotExist(p2pIds[i]);
    }
    return nodesInfo;
  }

  /// @notice Gets DON's info
  /// @param donId The DON ID
  /// @return DONInfo The DON's parameters

  function getDON(
    uint32 donId
  ) external view returns (DONInfo memory) {
    uint32 configCount = s_dons[donId].configCount;
    if (configCount == 0) revert DONDoesNotExist(donId);
    return _getDON(donId, configCount);
  }

  /// @notice Gets a historical DON's info
  /// @param donId The DON ID
  /// @param configCount The config count of the DON
  /// @return DONInfo The DON's parameters
  function getHistoricalDONInfo(uint32 donId, uint32 configCount) external view returns (DONInfo memory) {
    uint32 donConfigCount = s_dons[donId].configCount;
    if (donConfigCount == 0) revert DONDoesNotExist(donId);
    if (configCount > donConfigCount) {
      revert DONConfigDoesNotExist(donId, donConfigCount, configCount);
    }

    return _getDON(donId, configCount);
  }

  /// @notice Gets DON's info by its name
  /// @param donName The name of the DON
  /// @return DONInfo The DON's parameters
  function getDONByName(
    string calldata donName
  ) external view returns (DONInfo memory) {
    uint32 donId = s_donNameToId[donName];
    if (donId == 0) revert DONWithNameDoesNotExist(donName);
    uint32 configCount = s_dons[donId].configCount;
    return _getDON(donId, configCount);
  }

  /// @notice Returns the list of configured DONs
  /// @return DONInfo[] The list of configured DONs
  function getDONs() external view returns (DONInfo[] memory) {
    /// Minus one to account for s_nextDONId starting at index 1
    uint32 donId = s_nextDONId;
    DONInfo[] memory dons = new DONInfo[](donId - 1);
    uint256 idx;
    ///
    for (uint32 i = 1; i < donId; ++i) {
      if (s_dons[i].id != 0) {
        uint32 configCount = s_dons[i].configCount;
        dons[idx] = _getDON(i, configCount);
        ++idx;
      }
    }
    if (idx != donId - 1) {
      assembly {
        mstore(dons, idx)
      }
    }
    return dons;
  }

  /// @notice Returns the DON specific configuration for a capability
  /// @param donId The DON's ID
  /// @param capabilityId The Capability ID
  /// @return bytes The DON specific configuration for the capability stored on the capability registry
  /// @return bytes The DON specific configuration stored on the capability's configuration contract
  function getCapabilityConfigs(
    uint32 donId,
    string memory capabilityId
  ) external view returns (bytes memory, bytes memory) {
    uint32 configCount = s_dons[donId].configCount;

    bytes32 hashedCapabilityId = _hash(capabilityId);
    bytes memory mutableDONConfig = s_dons[donId].config[configCount].capabilityConfigs[capabilityId];
    bytes memory globalCapabilityConfig;

    if (s_capabilities[hashedCapabilityId].configurationContract != address(0)) {
      globalCapabilityConfig = ICapabilityConfiguration(s_capabilities[hashedCapabilityId].configurationContract)
        .getCapabilityConfiguration(donId);
    }

    return (mutableDONConfig, globalCapabilityConfig);
  }

  /// @notice Gets all DON IDs that belong to a specific family
  /// @param donFamily The family name to query for
  /// @return uint[] Array of DON IDs that belong to the specified family
  function getDONsInFamily(
    string calldata donFamily
  ) external view returns (uint256[] memory) {
    return s_donFamilyMembers[_hash(donFamily)].values();
  }

  /// @notice Checks if a DON name is already taken
  /// @param donName The name of the DON to check
  /// @return bool True if the DON name is already taken, false otherwise
  function isDONNameTaken(
    string calldata donName
  ) external view returns (bool) {
    return s_donNameToId[donName] != 0;
  }

  /// @notice Returns the list of existing DON families including the default
  /// family that is an empty string unless there are no DONs in that family.
  /// @return string[] The list of existing DON families
  function getDONFamilies() external view returns (string[] memory) {
    bytes32[] memory donFamilyHashes = s_activeDONFamilyNames.values();
    string[] memory donFamilies = new string[](donFamilyHashes.length);
    for (uint256 i; i < donFamilyHashes.length; ++i) {
      donFamilies[i] = s_donFamilyHashToDonFamily[donFamilyHashes[i]];
    }
    return donFamilies;
  }

  // ================================================================
  // |                   Internal functions                         |
  // ================================================================

  /// @notice Removes a DON from the Capability Registry
  /// @param donId The ID of the DON to remove
  function _removeDON(
    uint32 donId
  ) internal {
    DON storage don = s_dons[donId];

    uint32 configCount = don.configCount;
    EnumerableSet.Bytes32Set storage nodeP2PIds = don.config[configCount].nodes;

    bool isWorkflowDON = don.acceptsWorkflows;
    for (uint256 j; j < nodeP2PIds.length(); ++j) {
      if (isWorkflowDON) {
        delete s_nodes[nodeP2PIds.at(j)].workflowDONId;
      } else {
        s_nodes[nodeP2PIds.at(j)].capabilitiesDONIds.remove(donId);
      }
    }

    // DON config count starts at index 1
    if (don.configCount == 0) revert DONDoesNotExist(donId);

    for (uint256 i; i < s_donIdToDonFamilyHashes[donId].length(); ++i) {
      _removeDONFromFamily(donId, s_donIdToDonFamilyHashes[donId].at(i));
    }

    // Free up the DON name for reuse
    delete s_donNameToId[don.config[configCount].name];

    delete s_dons[donId];
    emit ConfigSet(donId, 0);
  }

  /// @notice Removes a DON from a family
  /// @param donId The ID of the DON to remove from the family
  function _removeDONFromFamily(uint32 donId, bytes32 donFamilyHash) internal {
    // Remove the family hash from the families the DON belongs to
    s_donIdToDonFamilyHashes[donId].remove(donFamilyHash);
    // Remove the DON ID from the list of DON IDs in the current family
    s_donFamilyMembers[donFamilyHash].remove(donId);
    // If the current family is empty, remove it from the set of family names
    if (s_donFamilyMembers[donFamilyHash].length() == 0) {
      s_activeDONFamilyNames.remove(donFamilyHash);
    }

    emit DONRemovedFromFamily(donId, s_donFamilyHashToDonFamily[donFamilyHash]);
  }

  /// @notice Adds a DON to a family
  /// @param donId The ID of the DON to add to the family
  /// @param donFamily The family name to add the DON to
  function _addDONToFamily(uint32 donId, string memory donFamily) internal {
    bytes32 donFamilyHash = _hash(donFamily);

    // If the DON is already in the family, do nothing.
    // There is no point in erroring out as this is a no-op and the erroring
    // would not provide any value to the user.
    if (s_donIdToDonFamilyHashes[donId].contains(donFamilyHash)) return;

    // Set the DON family name hash to the new family name hash so it can be
    // retrieved by the hash later. We do not need to clean up the old family
    // name hash as it is only used to retrieve the family name by the hash.
    s_donFamilyHashToDonFamily[donFamilyHash] = donFamily;

    // Add the new family name to the set of family names. This operation is
    // idempotent and will not add the family name if it already exists.
    s_activeDONFamilyNames.add(donFamilyHash);

    // Add the DON ID to the list of DON IDs in the new family
    s_donFamilyMembers[donFamilyHash].add(donId);
    // Add the family hash to the list of families the DON belongs to
    s_donIdToDonFamilyHashes[donId].add(donFamilyHash);

    emit DONAddedToFamily(donId, donFamily);
  }

  /// @notice Sets the configuration for a DON
  /// @param nodes The nodes making up the DON
  /// @param capabilityConfigurations The list of configurations for the capabilities supported by the DON
  /// @param donParams The DON's parameters
  function _setDONConfig(
    bytes32[] memory nodes,
    CapabilityConfiguration[] memory capabilityConfigurations,
    DONParams memory donParams
  ) internal {
    DON storage don = s_dons[donParams.id];
    MutableDONConfig storage donConfig = don.config[donParams.configCount];

    // Validate the f value. We are intentionally relaxing the 3f+1 requirement
    // as not all DONs will run OCR instances.
    if ((!i_canAddOneNodeDONs && donParams.f == 0) || donParams.f + 1 > nodes.length) {
      revert InvalidFaultTolerance(donParams.f, nodes.length);
    }

    if (bytes(donParams.name).length == 0) {
      revert DONNameCannotBeEmpty(donParams.id);
    }

    MutableDONConfig storage prevDONConfig = don.config[donParams.configCount - 1];

    // Check if the DON name is changing. If it is, we need to update the mapping.
    if (_hash(prevDONConfig.name) != _hash(donParams.name)) {
      if (s_donNameToId[donParams.name] != 0) {
        revert DONNameAlreadyTaken(donParams.name);
      }

      delete s_donNameToId[donConfig.name];
      s_donNameToId[donParams.name] = donParams.id;
    }

    // Skip removing supported DON Ids from previously configured nodes in DON if
    // we are adding the DON for the first time
    if (donParams.configCount > 1) {
      // We acknowledge that this may result in an out of gas error if the number of configured
      // nodes is large.  This is mitigated by ensuring that there will not be a large number
      // of nodes configured to a DON.
      // We also do not remove the nodes from the previous DON capability config.  This is not
      // needed as the previous config will be overwritten by storing the latest config
      // at configCount
      for (uint256 i; i < prevDONConfig.nodes.length(); ++i) {
        s_nodes[prevDONConfig.nodes.at(i)].capabilitiesDONIds.remove(donParams.id);
        delete s_nodes[prevDONConfig.nodes.at(i)].workflowDONId;
      }
    }

    s_dons[donParams.id].id = donParams.id;
    s_dons[donParams.id].acceptsWorkflows = donParams.acceptsWorkflows;
    s_dons[donParams.id].configCount = donParams.configCount;

    donConfig.config = donParams.config;
    donConfig.name = donParams.name;
    donConfig.isPublic = donParams.isPublic;
    donConfig.f = donParams.f;

    for (uint256 i; i < nodes.length; ++i) {
      if (!donConfig.nodes.add(nodes[i])) revert DuplicateDONNode(donParams.id, nodes[i]);

      if (donParams.acceptsWorkflows) {
        if (s_nodes[nodes[i]].workflowDONId != donParams.id && s_nodes[nodes[i]].workflowDONId != 0) {
          revert NodePartOfWorkflowDON(donParams.id, nodes[i]);
        }
        s_nodes[nodes[i]].workflowDONId = donParams.id;
      } else {
        /// Fine to add a duplicate DON ID to the set of supported DON IDs again as the set
        /// will only store unique DON IDs
        s_nodes[nodes[i]].capabilitiesDONIds.add(donParams.id);
      }
    }

    for (uint256 i; i < capabilityConfigurations.length; ++i) {
      CapabilityConfiguration memory configuration = capabilityConfigurations[i];

      bytes32 hashedCapabilityId = _hash(configuration.capabilityId);
      if (!s_hashedCapabilityIds.contains(hashedCapabilityId)) {
        revert CapabilityDoesNotExist(configuration.capabilityId);
      }
      if (s_deprecatedHashedCapabilityIds.contains(hashedCapabilityId)) {
        revert CapabilityIsDeprecated(configuration.capabilityId);
      }

      if (donConfig.capabilityConfigs[configuration.capabilityId].length > 0) {
        revert DuplicateDONCapability(donParams.id, configuration.capabilityId);
      }

      for (uint256 j; j < nodes.length; ++j) {
        if (!s_nodes[nodes[j]].supportedHashedCapabilityIds[s_nodes[nodes[j]].configCount].contains(hashedCapabilityId))
        {
          revert NodeDoesNotSupportCapability(nodes[j], configuration.capabilityId);
        }
      }

      donConfig.capabilityIds.push(hashedCapabilityId);
      donConfig.capabilityConfigs[configuration.capabilityId] = configuration.config;

      _setCapabilityConfig(donParams.id, donParams.configCount, configuration.capabilityId, nodes, configuration.config);
    }
    emit ConfigSet(donParams.id, donParams.configCount);
  }

  /// @notice Sets the capability's config on the config contract
  /// @param donId The ID of the DON the capability is being configured for
  /// @param configCount The number of times the DON has been configured
  /// @param capabilityId The capability's ID
  /// @param nodes The nodes in the DON
  /// @param config The DON's capability config
  /// @dev Helper function used to resolve stack too deep errors in _setDONConfig
  function _setCapabilityConfig(
    uint32 donId,
    uint32 configCount,
    string memory capabilityId,
    bytes32[] memory nodes,
    bytes memory config
  ) internal {
    bytes32 hashedCapabilityId = _hash(capabilityId);
    if (s_capabilities[hashedCapabilityId].configurationContract != address(0)) {
      ICapabilityConfiguration(s_capabilities[hashedCapabilityId].configurationContract).beforeCapabilityConfigSet(
        nodes, config, configCount, donId
      );
    }
  }

  /// @notice Gets DON's data
  /// @param donId The DON ID
  /// @param configCount The config count of the DON
  /// @return DONInfo The DON's parameters
  function _getDON(uint32 donId, uint32 configCount) internal view returns (DONInfo memory) {
    DON storage don = s_dons[donId];
    MutableDONConfig storage donConfig = don.config[configCount];

    bytes32[] memory capabilityIds = donConfig.capabilityIds;
    CapabilityConfiguration[] memory capabilityConfigurations = new CapabilityConfiguration[](capabilityIds.length);

    for (uint256 i; i < capabilityConfigurations.length; ++i) {
      string memory capabilityId = s_hashedCapabilityIdToCapabilityId[capabilityIds[i]];
      capabilityConfigurations[i] =
        CapabilityConfiguration({capabilityId: capabilityId, config: donConfig.capabilityConfigs[capabilityId]});
    }

    string[] memory donFamilies = new string[](s_donIdToDonFamilyHashes[donId].length());
    for (uint256 i; i < s_donIdToDonFamilyHashes[donId].length(); ++i) {
      donFamilies[i] = s_donFamilyHashToDonFamily[s_donIdToDonFamilyHashes[donId].at(i)];
    }

    return DONInfo({
      id: don.id,
      acceptsWorkflows: don.acceptsWorkflows,
      configCount: configCount,
      donFamilies: donFamilies,
      name: donConfig.name,
      config: donConfig.config,
      f: donConfig.f,
      isPublic: donConfig.isPublic,
      nodeP2PIds: donConfig.nodes.values(),
      capabilityConfigurations: capabilityConfigurations
    });
  }

  /// @notice Hashes a string
  /// @param str The string to hash
  /// @return bytes32 The hash of the string
  function _hash(
    string memory str
  ) internal pure returns (bytes32) {
    return keccak256(bytes(str));
  }

  /// @notice Generates a hash from the node operator's data which uniquely identifies it
  /// @param nodeOperator The node operator's data
  /// @return bytes32 The hash of the node operator's data
  function _nodeOperatorHash(
    NodeOperator memory nodeOperator
  ) internal pure returns (bytes32) {
    return keccak256(abi.encode(nodeOperator.admin, nodeOperator.name));
  }
}
