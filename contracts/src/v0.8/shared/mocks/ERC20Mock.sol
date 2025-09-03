// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/ERC20.sol";

contract ERC20Mock is ERC20 {
  uint8 internal immutable i_decimals;

  constructor(
    uint8 decimals_
  ) ERC20("ERC20Mock", "E20M") {
    i_decimals = decimals_;
  }

  function mint(address account, uint256 amount) external {
    _mint(account, amount);
  }

  function burn(address account, uint256 amount) external {
    _burn(account, amount);
  }

  function decimals() public view override returns (uint8) {
    return i_decimals;
  }
}
