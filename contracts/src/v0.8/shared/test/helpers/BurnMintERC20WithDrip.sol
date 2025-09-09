// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {BurnMintERC20} from "../../token/ERC20/BurnMintERC20.sol";

contract BurnMintERC20WithDrip is BurnMintERC20 {
  constructor(string memory name, string memory symbol) BurnMintERC20(name, symbol, 18, 0, 0) {}

  // Gives one full token to any given address.
  function drip(
    address to
  ) external {
    _mint(to, 1e18);
  }
}
