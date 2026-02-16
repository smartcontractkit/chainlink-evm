// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../../token/ERC20/BaseERC20.sol";
import {BaseERC20Setup} from "./BaseERC20Setup.t.sol";

contract BaseERC20_transfer is BaseERC20Setup {
  function test_transfer() public {
    uint256 balancePre = s_baseERC20.balanceOf(STRANGER);
    uint256 sendingAmount = s_amount / 2;

    s_baseERC20.transfer(STRANGER, sendingAmount);

    assertEq(sendingAmount + balancePre, s_baseERC20.balanceOf(STRANGER));
  }

  // Reverts

  function test_transfer_RevertWhen_InvalidAddress() public {
    vm.expectRevert(abi.encodeWithSelector(BaseERC20.InvalidRecipient.selector, address(s_baseERC20)));

    s_baseERC20.transfer(address(s_baseERC20), s_amount);
  }
}
