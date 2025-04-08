// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

contract MockCapabilitiesRegistry {
  uint32 private s_nextDonId;

  constructor(
    uint32 _initialDonId
  ) {
    s_nextDonId = _initialDonId;
  }

  function getNextDONId() external view returns (uint32) {
    return s_nextDonId;
  }
}
