// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IReceiver} from "../../interfaces/IReceiver.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

contract Receiver is IReceiver {
  event MessageReceived(bytes metadata, bytes[] mercuryReports);

  bytes public latestReport;

  constructor() {}

  function onReport(bytes calldata metadata, bytes calldata rawReport) external {
    latestReport = rawReport;

    // parse actual report
    bytes[] memory mercuryReports = abi.decode(rawReport, (bytes[]));
    emit MessageReceived(metadata, mercuryReports);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) public pure override returns (bool) {
    return interfaceId == type(IReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
