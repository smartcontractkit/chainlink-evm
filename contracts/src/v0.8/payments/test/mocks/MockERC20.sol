// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
  uint8 public immutable i_decimals;

  constructor(string memory name, string memory symbol, uint8 decimals) ERC20(name, symbol) {
    i_decimals = decimals;
  }
}
