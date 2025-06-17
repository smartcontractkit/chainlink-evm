// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

interface IChannelConfigStore is IERC165 {
  function setChannelDefinitions(uint32 donId, string calldata url, bytes32 sha) external;
}
