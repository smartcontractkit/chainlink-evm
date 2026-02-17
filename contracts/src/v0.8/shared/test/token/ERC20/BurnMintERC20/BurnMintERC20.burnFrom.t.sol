// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BurnMintERC20} from "../../../../token/ERC20/BurnMintERC20.sol";
import {BurnMintERC20Setup} from "./BurnMintERC20Setup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {IERC20Errors} from "@openzeppelin/contracts@5.3.0/interfaces/draft-IERC6093.sol";

contract BurnMintERC20_burnFrom is BurnMintERC20Setup {
  function setUp() public virtual override {
    BurnMintERC20Setup.setUp();
  }

  function test_BurnFrom() public {
    s_burnMintERC20.approve(s_mockPool, s_amount);

    changePrank(s_mockPool);

    s_burnMintERC20.burnFrom(OWNER, s_amount);

    assertEq(0, s_burnMintERC20.balanceOf(OWNER));
  }

  // Reverts

  function test_burnFrom_RevertWhen_SenderNotBurner() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        IAccessControl.AccessControlUnauthorizedAccount.selector, OWNER, s_burnMintERC20.BURNER_ROLE()
      )
    );

    s_burnMintERC20.burnFrom(OWNER, s_amount);
  }

  function test_burnFrom_RevertWhen_InsufficientAllowance() public {
    changePrank(s_mockPool);

    vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InsufficientAllowance.selector, s_mockPool, 0, s_amount));

    s_burnMintERC20.burnFrom(OWNER, s_amount);
  }

  function test_burnFrom_RevertWhen_ExceedsBalance() public {
    s_burnMintERC20.approve(s_mockPool, s_amount * 2);

    changePrank(s_mockPool);

    vm.expectRevert(
      abi.encodeWithSelector(IERC20Errors.ERC20InsufficientBalance.selector, OWNER, s_amount, s_amount * 2)
    );

    s_burnMintERC20.burnFrom(OWNER, s_amount * 2);
  }
}
