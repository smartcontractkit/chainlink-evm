// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../../token/ERC20/BaseERC20.sol";
import {BurnMintERC20} from "../../../../token/ERC20/BurnMintERC20.sol";
import {BurnMintERC20Setup} from "./BurnMintERC20Setup.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract BurnMintERC20_mint is BurnMintERC20Setup {
  function test_mint() public {
    uint256 balancePre = s_burnMintERC20.balanceOf(OWNER);

    s_burnMintERC20.grantMintAndBurnRoles(OWNER);

    vm.expectEmit();
    emit IERC20.Transfer(address(0), OWNER, s_amount);

    s_burnMintERC20.mint(OWNER, s_amount);

    assertEq(balancePre + s_amount, s_burnMintERC20.balanceOf(OWNER));
  }

  // Revert

  function test_mint_RevertWhen_SenderNotMinter() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(
      abi.encodeWithSelector(
        IAccessControl.AccessControlUnauthorizedAccount.selector, STRANGER, s_burnMintERC20.MINTER_ROLE()
      )
    );

    s_burnMintERC20.mint(STRANGER, 1e18);
  }

  function test_mint_RevertWhen_MaxSupplyExceeded() public {
    changePrank(s_mockPool);

    // Mint max supply
    s_burnMintERC20.mint(OWNER, s_burnMintERC20.maxSupply());

    vm.expectRevert(abi.encodeWithSelector(BurnMintERC20.MaxSupplyExceeded.selector, s_burnMintERC20.maxSupply() + 1));

    // Attempt to mint 1 more than max supply
    s_burnMintERC20.mint(OWNER, 1);
  }

  function test_mint_RevertWhen_InvalidRecipient() public {
    s_burnMintERC20.grantMintAndBurnRoles(OWNER);

    vm.expectRevert(abi.encodeWithSelector(BaseERC20.InvalidRecipient.selector, address(s_burnMintERC20)));
    s_burnMintERC20.mint(address(s_burnMintERC20), 1e18);
  }
}
