// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IReceiver} from "../../interfaces/IReceiver.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

contract MaliciousReportReceiver is IReceiver {
  event MessageReceived(bytes metadata, bytes[] mercuryReports);

  bytes public latestReport;

  function onReport(bytes calldata metadata, bytes calldata rawReport) external {
    // Exhaust all gas that was provided
    for (uint256 i = 0; i < 1_000_000_000; ++i) {
      bytes[] memory mercuryReports = abi.decode(rawReport, (bytes[]));
      latestReport = rawReport;
      emit MessageReceived(metadata, mercuryReports);
    }
  }

  function supportsInterface(
    bytes4 interfaceId
  ) public pure override returns (bool) {
    return interfaceId == type(IReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
