// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

interface IChannelConfigStore is IERC165 {
  type ChannelAdderId is uint32;

  function setChannelDefinitions(uint32 donId, string calldata url, bytes32 sha) external;
  function addChannelDefinitions(
    uint32 donId,
    ChannelAdderId channelAdderId,
    string calldata url,
    bytes32 sha
  ) external;
  function setChannelAdderAddress(ChannelAdderId channelAdderId, address adderAddress) external;
  function setChannelAdder(uint32 donId, ChannelAdderId channelAdderId, bool allowed) external;
  function getChannelAdderAddress(
    ChannelAdderId channelAdderId
  ) external view returns (address);
  function isChannelAdderAllowed(uint32 donId, ChannelAdderId channelAdderId) external view returns (bool);
  function getAllowedChannelAdders(
    uint32 donId
  ) external view returns (ChannelAdderId[] memory);
}
