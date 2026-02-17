// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseERC20Setup} from "./BaseERC20Setup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract BaseERC20_supportsInterface is BaseERC20Setup {
  function test_SupportsInterface() public view {
    assertTrue(s_baseERC20.supportsInterface(type(IERC20).interfaceId));
    assertTrue(s_baseERC20.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_baseERC20.supportsInterface(type(IAccessControl).interfaceId));
  }
}
