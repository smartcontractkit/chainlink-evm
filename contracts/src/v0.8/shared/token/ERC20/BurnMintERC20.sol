// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IBurnMintERC20} from "./IBurnMintERC20.sol";

import {BaseERC20} from "./BaseERC20.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

/// @notice A basic ERC20 compatible token contract with burn and minting roles.
/// @dev The total supply can be limited during deployment.
contract BurnMintERC20 is BaseERC20, IBurnMintERC20  {
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "CCT 2.0.0-dev";
  }

  error MaxSupplyExceeded(uint256 supplyAfterMint);

  bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
  bytes32 public constant BURNER_ROLE = keccak256("BURNER_ROLE");

  /// @dev the underscores in parameter names are used to suppress compiler warnings about shadowing ERC20 functions.
  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals_,
    uint256 maxSupply_,
    uint256 preMint,
    address newOwner
  ) BaseERC20(name, symbol, decimals_, maxSupply_, preMint, newOwner) {}

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override returns (bool) {
    return super.supportsInterface(interfaceId) || interfaceId == type(IBurnMintERC20).interfaceId;
  }

  // ================================================================
  // │                      Burning & minting                       │
  // ================================================================

  /// @dev Uses OZ ERC20 _burn to disallow burning from address(0).
  /// @dev Decreases the total supply.
  function burn(
    uint256 amount
  ) public virtual override onlyRole(BURNER_ROLE) {
    _burn(_msgSender(), amount);
  }

  /// @inheritdoc IBurnMintERC20
  /// @dev Alias for BurnFrom for compatibility with the older naming convention.
  /// @dev Uses burnFrom for all validation & logic.
  function burn(address account, uint256 amount) public virtual override {
    burnFrom(account, amount);
  }

  /// @dev Uses OZ ERC20 _burn to disallow burning from address(0).
  /// @dev Decreases the total supply.
  function burnFrom(address account, uint256 amount) public virtual override onlyRole(BURNER_ROLE) {
    _spendAllowance(account, _msgSender(), amount);
    _burn(account, amount);
  }

  /// @inheritdoc IBurnMintERC20
  /// @dev Uses OZ ERC20 _mint to disallow minting to address(0).
  /// @dev Disallows minting to address(this)
  /// @dev Increases the total supply.
  function mint(address account, uint256 amount) public virtual override onlyRole(MINTER_ROLE) {
    if (account == address(this)) revert InvalidRecipient(account);
    if (i_maxSupply != 0 && totalSupply() + amount > i_maxSupply) revert MaxSupplyExceeded(totalSupply() + amount);

    _mint(account, amount);
  }

  // ================================================================
  // │                            Roles                             │
  // ================================================================

  /// @notice grants both mint and burn roles to `burnAndMinter`.
  /// @dev calls public functions so this function does not require
  /// access controls. This is handled in the inner functions.
  function grantMintAndBurnRoles(
    address burnAndMinter
  ) public virtual {
    grantRole(MINTER_ROLE, burnAndMinter);
    grantRole(BURNER_ROLE, burnAndMinter);
  }
}
