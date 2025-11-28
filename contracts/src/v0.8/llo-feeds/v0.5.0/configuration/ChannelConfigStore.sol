// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import {ConfirmedOwner} from "../../../shared/access/ConfirmedOwner.sol";
import {EnumerableSet} from "@openzeppelin/contracts@4.8.3/utils/structs/EnumerableSet.sol";

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";
import {IChannelConfigStore} from "./interfaces/IChannelConfigStore.sol";

contract ChannelConfigStore is ConfirmedOwner, IChannelConfigStore, ITypeAndVersion {
  using EnumerableSet for EnumerableSet.UintSet;

  event NewChannelDefinition(uint256 indexed donId, uint32 version, string url, bytes32 sha);
  event ChannelDefinitionAdded(uint256 indexed donId, uint32 indexed channelAdderId, string url, bytes32 sha);
  event ChannelAdderSet(uint256 indexed donId, uint32 indexed channelAdderId, bool allowed);
  event ChannelAdderAddressSet(uint32 indexed channelAdderId, address adderAddress);

  error UnauthorizedChannelAdder();

  constructor() ConfirmedOwner(msg.sender) {}

  /// @notice The version of a channel definition keyed by DON ID
  // Increments by 1 on every update
  mapping(uint256 => uint256) internal s_channelDefinitionVersions;

  /// @notice Mapping from channel adder ID to its corresponding address
  mapping(uint32 => address) internal s_channelAdderAddresses;

  /// @notice Mapping from DON ID to the set of allowed channel adder IDs
  mapping(uint256 => EnumerableSet.UintSet) internal s_allowedChannelAdders;

  function setChannelDefinitions(
    uint32 donId,
    string calldata url,
    bytes32 sha
  ) external onlyOwner {
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
    uint256 donId,
    uint32 channelAdderId,
    string calldata url,
    bytes32 sha
  ) external {
    if (msg.sender != s_channelAdderAddresses[channelAdderId]) {
      revert UnauthorizedChannelAdder();
    }
    if (!s_allowedChannelAdders[donId].contains(channelAdderId)) {
      revert UnauthorizedChannelAdder();
    }
    emit ChannelDefinitionAdded(donId, channelAdderId, url, sha);
  }

  /// @notice Sets the address for a channel adder ID
  /// @param channelAdderId The channel adder ID
  /// @param adderAddress The address to associate with the channel adder ID.
  /// Set this to the zero address (or some other address that cannot make
  /// calls) to disable the channel adder.
  function setChannelAdderAddress(
    uint32 channelAdderId,
    address adderAddress
  ) external onlyOwner {
    s_channelAdderAddresses[channelAdderId] = adderAddress;
    emit ChannelAdderAddressSet(channelAdderId, adderAddress);
  }

  /// @notice Sets whether a channel adder ID is allowed for a DON
  /// @param donId The DON ID
  /// @param channelAdderId The channel adder ID
  /// @param allowed Whether the channel adder should be allowed or removed
  function setChannelAdder(
    uint256 donId,
    uint32 channelAdderId,
    bool allowed
  ) external onlyOwner {
    if (allowed) {
      s_allowedChannelAdders[donId].add(channelAdderId);
    } else {
      s_allowedChannelAdders[donId].remove(channelAdderId);
    }
    emit ChannelAdderSet(donId, channelAdderId, allowed);
  }

  /// @notice Gets the address associated with a channel adder ID
  /// @param channelAdderId The channel adder ID
  /// @return The address associated with the channel adder ID
  function getChannelAdderAddress(
    uint32 channelAdderId
  ) external view returns (address) {
    return s_channelAdderAddresses[channelAdderId];
  }

  /// @notice Checks if a channel adder is allowed for a DON
  /// @param donId The DON ID
  /// @param channelAdderId The channel adder ID
  /// @return True if the channel adder is allowed for the DON
  function isChannelAdderAllowed(
    uint256 donId,
    uint32 channelAdderId
  ) external view returns (bool) {
    return s_allowedChannelAdders[donId].contains(channelAdderId);
  }

  /// @notice Gets all allowed channel adder IDs for a DON
  /// @param donId The DON ID
  /// @return An array of allowed channel adder IDs
  function getAllowedChannelAdders(
    uint256 donId
  ) external view returns (uint256[] memory) {
    return s_allowedChannelAdders[donId].values();
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "ChannelConfigStore 0.0.1";
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(IChannelConfigStore).interfaceId;
  }
}
