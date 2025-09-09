// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC20Metadata as IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

interface IWrappedNative is IERC20 {
  function deposit() external payable;

  function withdraw(
    uint256 wad
  ) external;
}
