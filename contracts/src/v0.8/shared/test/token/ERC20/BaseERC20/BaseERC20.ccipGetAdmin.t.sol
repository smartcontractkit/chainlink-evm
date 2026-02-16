// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../../token/ERC20/BaseERC20.sol";
import {BaseERC20Setup} from "./BaseERC20Setup.t.sol";

contract BaseERC20_getCCIPAdmin is BaseERC20Setup {
  function test_getCCIPAdmin() public view {
    assertEq(OWNER, s_baseERC20.getCCIPAdmin());
  }

  function test_setCCIPAdmin() public {
    address newAdmin = makeAddr("newAdmin");

    vm.expectEmit();
    emit BaseERC20.CCIPAdminTransferred(OWNER, newAdmin);

    s_baseERC20.setCCIPAdmin(newAdmin);

    assertEq(newAdmin, s_baseERC20.getCCIPAdmin());
  }
}
