// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

/// @notice MessageEmitter is used to emit custom messages from a contract.
/// @dev Sender may only emit a message once per block timestamp.
contract MessageEmitter is ITypeAndVersion {
  string public constant override typeAndVersion = "MessageEmitter 1.0.0-dev";

  event MessageEmitted(address indexed emitter, uint256 indexed timestamp, string message);

  mapping(bytes32 key => string message) private s_messages;
  mapping(address emitter => string message) private s_lastMessage;

  function emitMessage(
    string calldata message
  ) public {
    // solhint-disable-next-line gas-custom-errors
    require(bytes(message).length > 0, "Message cannot be empty");
    bytes32 key = _hashKey(msg.sender, block.timestamp);
    // solhint-disable-next-line gas-custom-errors
    require(bytes(s_messages[key]).length == 0, "Message already exists for the same sender and block timestamp");
    s_messages[key] = message;
    s_lastMessage[msg.sender] = message;
    emit MessageEmitted(msg.sender, block.timestamp, message);
  }

  function getMessage(address emitter, uint256 timestamp) public view returns (string memory) {
    bytes32 key = _hashKey(emitter, timestamp);
    // solhint-disable-next-line gas-custom-errors
    require(bytes(s_messages[key]).length > 0, "Message does not exist for the given sender and timestamp");
    return s_messages[key];
  }

  function getLastMessage(
    address emitter
  ) public view returns (string memory) {
    // solhint-disable-next-line gas-custom-errors
    require(bytes(s_lastMessage[emitter]).length > 0, "No last message for the given sender");
    return s_lastMessage[emitter];
  }

  function _hashKey(address emitter, uint256 timestamp) internal pure returns (bytes32) {
    return keccak256(abi.encode(emitter, timestamp));
  }
}
