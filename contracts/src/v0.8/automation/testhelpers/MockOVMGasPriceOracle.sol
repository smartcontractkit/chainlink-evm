// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockOVMGasPriceOracle {
  function getL1Fee(
    bytes memory
  ) public pure returns (uint256) {
    return 2_000_000;
  }

  function getL1FeeUpperBound(
    uint256
  ) public pure returns (uint256) {
    return 2_000_000;
  }
}
