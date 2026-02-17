// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../../../token/ERC20/IBurnMintERC20.sol";

import {BurnMintERC20Setup} from "./BurnMintERC20Setup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@4.8.3/access/IAccessControl.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/utils/introspection/IERC165.sol";

contract BurnMintERC20_supportsInterface is BurnMintERC20Setup {
  function test_SupportsInterface() public view {
    // BaseERC20 interfaces
    assertTrue(s_burnMintERC20.supportsInterface(type(IERC20).interfaceId));
    assertTrue(s_burnMintERC20.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_burnMintERC20.supportsInterface(type(IAccessControl).interfaceId));
    // BurnMintERC20 specific interface
    assertTrue(s_burnMintERC20.supportsInterface(type(IBurnMintERC20).interfaceId));
  }
}
