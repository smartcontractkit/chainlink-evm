// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import {ConfirmedOwner} from "../../../shared/access/ConfirmedOwner.sol";
import {EnumerableSet} from "@openzeppelin/contracts@4.8.3/utils/structs/EnumerableSet.sol";

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";
import {IChannelConfigStore} from "./interfaces/IChannelConfigStore.sol";

contract ChannelConfigStore is ConfirmedOwner, IChannelConfigStore, ITypeAndVersion {
  // This contract uses uint32 for donIds when they are function arguments and
  // uint256 elsewhere (e.g. in storage or event params). This inconsistency is
  // ugly, but we maintain it for backwards compatibility.

  using EnumerableSet for EnumerableSet.UintSet;

  event NewChannelDefinition(uint256 indexed donId, uint32 version, string url, bytes32 sha);
  event ChannelDefinitionAdded(uint256 indexed donId, ChannelAdderId indexed channelAdderId, string url, bytes32 sha);
  event ChannelAdderSet(uint256 indexed donId, ChannelAdderId indexed channelAdderId, bool allowed);
  event ChannelAdderAddressSet(ChannelAdderId indexed channelAdderId, address adderAddress);

  /// @notice Thrown when a caller is not authorized to add channel definitions.
  error UnauthorizedChannelAdder();

  /// @notice Thrown when a ChannelAdderId is reserved.
  error ReservedChannelAdderId();

  // We reserve the ChannelAdderIds 0 through 999. 1 is used by the offchain code to internally
  // represent the owner. The others are reserved for future use.
  ChannelAdderId internal constant MIN_CHANNEL_ADDER_ID = ChannelAdderId.wrap(1000);

  constructor() ConfirmedOwner(msg.sender) {}

  /// @notice The version of a channel definition keyed by DON ID.
  // Increments by 1 on every update.
  mapping(uint256 => uint256) internal s_channelDefinitionVersions;

  /// @notice Mapping from channel adder ID to its corresponding address
  mapping(ChannelAdderId => address) internal s_channelAdderAddresses;

  /// @notice Mapping from DON ID to the set of allowed channel adder IDs
  mapping(uint256 => EnumerableSet.UintSet) internal s_allowedChannelAdders;

  /// @notice Allows the owner to arbitrarily set channel definitions to the specified DON.
  /// Unlike the channel adder, the owner can not only add, but also modify and delete
  /// channel definitions. The DON enforces (in its consensus rules), that the channel
  /// definitions provided by the owner are well-formed.
  /// @param donId The DON ID
  /// @param url The URL of the channel definition
  /// @param sha The SHA hash of the channel definition
  function setChannelDefinitions(uint32 donId, string calldata url, bytes32 sha) external onlyOwner {
    uint32 newVersion = uint32(++s_channelDefinitionVersions[uint256(donId)]);
    emit NewChannelDefinition(donId, newVersion, url, sha);
  }

  /// @notice Allows a channel adder to add channel definitions to the specified DON.
  /// The DON enforces (in its consensus rules), that the channel definitions provided
  /// by the channel adder are well-formed, purely additive, and do not overload the DON.
  /// @param donId The DON ID
  /// @param channelAdderId The channel adder ID
  /// @param url The URL of the channel definition
  /// @param sha The SHA hash of the channel definition
  function addChannelDefinitions(
    uint32 donId,
    ChannelAdderId channelAdderId,
    string calldata url,
    bytes32 sha
  ) external {
    if (msg.sender != s_channelAdderAddresses[channelAdderId]) {
      revert UnauthorizedChannelAdder();
    }
    if (!s_allowedChannelAdders[donId].contains(ChannelAdderId.unwrap(channelAdderId))) {
      revert UnauthorizedChannelAdder();
    }
    emit ChannelDefinitionAdded(donId, channelAdderId, url, sha);
  }

  /// @notice Sets the address for a channel adder ID
  /// @param channelAdderId The channel adder ID
  /// @param adderAddress The address to associate with the channel adder ID.
  /// Set this to the zero address (or some other address that cannot make
  /// calls) to disable the channel adder.
  function setChannelAdderAddress(ChannelAdderId channelAdderId, address adderAddress) external onlyOwner {
    if (ChannelAdderId.unwrap(channelAdderId) < ChannelAdderId.unwrap(MIN_CHANNEL_ADDER_ID)) {
      revert ReservedChannelAdderId();
    }
    s_channelAdderAddresses[channelAdderId] = adderAddress;
    emit ChannelAdderAddressSet(channelAdderId, adderAddress);
  }

  /// @notice Sets whether a channel adder ID is allowed for a DON
  /// @param donId The DON ID
  /// @param channelAdderId The channel adder ID
  /// @param allowed Whether the channel adder should be allowed or removed
  function setChannelAdder(uint32 donId, ChannelAdderId channelAdderId, bool allowed) external onlyOwner {
    if (ChannelAdderId.unwrap(channelAdderId) < ChannelAdderId.unwrap(MIN_CHANNEL_ADDER_ID)) {
      revert ReservedChannelAdderId();
    }
    if (allowed) {
      s_allowedChannelAdders[donId].add(ChannelAdderId.unwrap(channelAdderId));
    } else {
      s_allowedChannelAdders[donId].remove(ChannelAdderId.unwrap(channelAdderId));
    }
    emit ChannelAdderSet(donId, channelAdderId, allowed);
  }

  /// @notice Gets the address associated with a channel adder ID
  /// @param channelAdderId The channel adder ID
  /// @return The address associated with the channel adder ID
  function getChannelAdderAddress(
    ChannelAdderId channelAdderId
  ) external view returns (address) {
    return s_channelAdderAddresses[channelAdderId];
  }

  /// @notice Checks if a channel adder is allowed for a DON
  /// @param donId The DON ID
  /// @param channelAdderId The channel adder ID
  /// @return True if the channel adder is allowed for the DON
  function isChannelAdderAllowed(uint32 donId, ChannelAdderId channelAdderId) external view returns (bool) {
    return s_allowedChannelAdders[donId].contains(ChannelAdderId.unwrap(channelAdderId));
  }

  /// @notice Gets all allowed channel adder IDs for a DON
  /// @param donId The DON ID
  /// @return allowedChannelAdderIds An array of allowed channel adder IDs
  function getAllowedChannelAdders(
    uint32 donId
  ) external view returns (ChannelAdderId[] memory allowedChannelAdderIds) {
    // Not very gas efficient, but we don't expect this function to be called
    // from onchain anyways.
    uint256[] memory values = s_allowedChannelAdders[donId].values();
    allowedChannelAdderIds = new ChannelAdderId[](values.length);
    for (uint256 i = 0; i < values.length; i++) {
      allowedChannelAdderIds[i] = ChannelAdderId.wrap(uint32(values[i]));
    }
    return allowedChannelAdderIds;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "ChannelConfigStore 1.0.0";
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(IChannelConfigStore).interfaceId;
  }
}
