// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../../token/ERC20/BaseERC20.sol";
import {BaseTest} from "../../../BaseTest.t.sol";

contract BaseERC20Setup is BaseTest {
  BaseERC20 internal s_baseERC20;

  uint256 internal s_amount = 1e18;

  address internal s_alice;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_alice = makeAddr("alice");

    s_baseERC20 = new BaseERC20("Chainlink Token", "LINK", 18, 1e27, 0, address(0));
    deal(address(s_baseERC20), OWNER, s_amount);
  }
}
