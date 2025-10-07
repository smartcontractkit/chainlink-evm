// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IReceiver} from "../../keystone/interfaces/IReceiver.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/interfaces/IERC165.sol";

contract ReserveManager is IReceiver {
  uint256 public lastTotalMinted;
  uint256 public lastTotalReserve;
  uint256 private s_requestIdCounter;

  event RequestReserveUpdate(UpdateReserves u);

  struct UpdateReserves {
    uint256 totalMinted;
    uint256 totalReserve;
  }

  function onReport(bytes calldata, bytes calldata report) external override {
    UpdateReserves memory updateReservesData = abi.decode(report, (UpdateReserves));
    lastTotalMinted = updateReservesData.totalMinted;
    lastTotalReserve = updateReservesData.totalReserve;

    s_requestIdCounter++;
    emit RequestReserveUpdate(updateReservesData);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override returns (bool) {
    return interfaceId == type(IReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
