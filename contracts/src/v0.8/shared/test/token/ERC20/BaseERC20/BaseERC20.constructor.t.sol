// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../../token/ERC20/BaseERC20.sol";
import {BaseERC20Setup} from "./BaseERC20Setup.t.sol";

contract BaseERC20_constructor is BaseERC20Setup {
  function test_Constructor() public {
    vm.startPrank(s_alice);

    string memory name = "Chainlink token v2";
    string memory symbol = "LINK2";
    uint8 decimals = 19;
    uint256 maxSupply = 1e33;

    s_baseERC20 = new BaseERC20(name, symbol, decimals, maxSupply, 1e18, address(0));

    assertEq(name, s_baseERC20.name());
    assertEq(symbol, s_baseERC20.symbol());
    assertEq(decimals, s_baseERC20.decimals());
    assertEq(maxSupply, s_baseERC20.maxSupply());

    assertTrue(s_baseERC20.hasRole(s_baseERC20.DEFAULT_ADMIN_ROLE(), s_alice));
    assertEq(s_baseERC20.balanceOf(s_alice), 1e18);
    assertEq(s_baseERC20.totalSupply(), 1e18);
  }
}
