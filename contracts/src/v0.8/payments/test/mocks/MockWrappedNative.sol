// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {MockERC20} from "../mocks/MockERC20.sol";

import {IWERC20} from "../../../shared/interfaces/IWERC20.sol";

// solhint-disable chainlink-solidity/inherited-constructor-args-not-in-contract-definition
contract MockWrappedNative is IWERC20, MockERC20("WETH", "WETH", 18) {
  function deposit() external payable override {
    _mint(msg.sender, msg.value);
  }

  function withdraw(
    uint256 amount
  ) external override {}
}
