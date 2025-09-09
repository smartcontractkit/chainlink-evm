pragma solidity 0.8.16;

import "../KeeperBase.sol";
import "../interfaces/KeeperCompatibleInterface.sol";

contract KeeperConsumer is KeeperCompatibleInterface, KeeperBase {
  uint256 public counter;
  uint256 public immutable interval;
  uint256 public lastTimeStamp;

  constructor(
    uint256 updateInterval
  ) public {
    interval = updateInterval;
    lastTimeStamp = block.timestamp;
    counter = 0;
  }

  function checkUpkeep(
    bytes calldata checkData
  ) external view override cannotExecute returns (bool upkeepNeeded, bytes memory performData) {
    return (true, checkData);
  }

  function performUpkeep(
    bytes calldata performData
  ) external override {
    counter = counter + 1;
  }
}
