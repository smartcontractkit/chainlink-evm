// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

interface IChannelConfigStore is IERC165 {
  function setChannelDefinitions(
    uint32 donId,
    string calldata url,
    bytes32 sha
  ) external;
  function addChannelDefinitions(
    uint256 donId,
    uint32 channelAdderId,
    string calldata url,
    bytes32 sha
  ) external;
  function setChannelAdderAddress(
    uint32 channelAdderId,
    address adderAddress
  ) external;
  function setChannelAdder(
    uint256 donId,
    uint32 channelAdderId,
    bool allowed
  ) external;
  function getChannelAdderAddress(
    uint32 channelAdderId
  ) external view returns (address);
  function isChannelAdderAllowed(
    uint256 donId,
    uint32 channelAdderId
  ) external view returns (bool);
  function getAllowedChannelAdders(
    uint256 donId
  ) external view returns (uint256[] memory);
}
