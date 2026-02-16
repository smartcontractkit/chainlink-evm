// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BurnMintERC20} from "../../../../token/ERC20/BurnMintERC20.sol";
import {BurnMintERC20Setup} from "./BurnMintERC20Setup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {IERC20Errors} from "@openzeppelin/contracts@5.3.0/interfaces/draft-IERC6093.sol";

contract BurnMintERC20_burn is BurnMintERC20Setup {
  function test_BasicBurn() public {
    s_burnMintERC20.grantRole(s_burnMintERC20.BURNER_ROLE(), OWNER);
    deal(address(s_burnMintERC20), OWNER, s_amount);

    vm.expectEmit();
    emit IERC20.Transfer(OWNER, address(0), s_amount);

    s_burnMintERC20.burn(s_amount);

    assertEq(0, s_burnMintERC20.balanceOf(OWNER));
  }

  // Revert

  function test_burn_RevertWhen_SenderNotBurner() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        IAccessControl.AccessControlUnauthorizedAccount.selector,
        OWNER,
        s_burnMintERC20.BURNER_ROLE()
      )
    );

    s_burnMintERC20.burnFrom(STRANGER, s_amount);
  }

  function test_burn_RevertWhen_ExceedsBalance() public {
    changePrank(s_mockPool);

    vm.expectRevert(
      abi.encodeWithSelector(
        IERC20Errors.ERC20InsufficientBalance.selector,
        s_mockPool,
        0,
        s_amount * 2
      )
    );

    s_burnMintERC20.burn(s_amount * 2);
  }

  function test_burn_RevertWhen_BurnFromZeroAddress() public {
    s_burnMintERC20.grantRole(s_burnMintERC20.BURNER_ROLE(), address(0));
    changePrank(address(0));

    vm.expectRevert(
      abi.encodeWithSelector(IERC20Errors.ERC20InvalidSender.selector, address(0))
    );

    s_burnMintERC20.burn(0);
  }
}
