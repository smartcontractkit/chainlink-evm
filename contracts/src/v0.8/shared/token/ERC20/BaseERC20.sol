// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IGetCCIPAdmin} from "../../../shared/interfaces/IGetCCIPAdmin.sol";
import {ITypeAndVersion} from "../../interfaces/ITypeAndVersion.sol";

import {AccessControl} from "@openzeppelin/contracts@5.3.0/access/AccessControl.sol";
import {IAccessControl} from "@openzeppelin/contracts@5.3.0/access/IAccessControl.sol";
import {ERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

/// @notice A basic ERC20 compatible token contract with burn and minting roles.
/// @dev The total supply can be limited during deployment.
contract BaseERC20 is IGetCCIPAdmin, IERC165, ERC20, AccessControl, ITypeAndVersion {
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "BasicERC20 2.0.0-dev";
  }

  error InvalidRecipient(address recipient);

  event CCIPAdminTransferred(address indexed previousAdmin, address indexed newAdmin);

  /// @dev The number of decimals for the token
  uint8 internal immutable i_decimals;

  /// @dev The maximum supply of the token, 0 if unlimited
  uint256 internal immutable i_maxSupply;

  /// @dev the CCIPAdmin can be used to register with the CCIP token admin registry, but has no other special powers,
  /// and can only be transferred by the owner.
  address internal s_ccipAdmin;

  /// @dev the underscores in parameter names are used to suppress compiler warnings about shadowing ERC20 functions
  /// @param name The name of the token
  /// @param symbol The symbol of the token
  /// @param decimals_ The number of decimals the token uses
  /// @param maxSupply_ The maximum supply of the token, 0 if unlimited
  /// @param preMint The amount of tokens to mint to the deployer upon construction. NOTE: the base version of this
  /// contract does not support minting additional tokens after deployment, so this should be set to the full supply.
  /// @param newOwner The address to grant the DEFAULT_ADMIN_ROLE to, which can be used to transfer the CCIPAdmin role.
  /// If address(0) is passed, the deployer will be used as the owner. This address will receive the preMinted tokens,
  /// if any.
  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals_,
    uint256 maxSupply_,
    uint256 preMint,
    address newOwner
  ) ERC20(name, symbol) {
    i_decimals = decimals_;
    i_maxSupply = maxSupply_;

    address owner = newOwner == address(0) ? msg.sender : newOwner;

    s_ccipAdmin = owner;

    // Mint the initial supply to the new Owner, saving gas by not calling if the mint amount is zero.
    if (preMint != 0) _mint(owner, preMint);

    // Set up the owner as the initial minter and burner.
    _grantRole(DEFAULT_ADMIN_ROLE, owner);
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(AccessControl, IERC165) returns (bool) {
    return interfaceId == type(IERC20).interfaceId || interfaceId == type(IERC165).interfaceId
      || interfaceId == type(IAccessControl).interfaceId || interfaceId == type(IGetCCIPAdmin).interfaceId;
  }

  // ================================================================
  // │                            ERC20                             │
  // ================================================================

  /// @dev Returns the number of decimals used in its user representation.
  function decimals() public view virtual override returns (uint8) {
    return i_decimals;
  }

  /// @dev Returns the max supply of the token, 0 if unlimited.
  function maxSupply() public view virtual returns (uint256) {
    return i_maxSupply;
  }

  /// @dev Uses OZ ERC20 _approve to disallow approving for address(0).
  /// @dev Disallows approving for address(this).
  function _approve(address owner, address spender, uint256 value, bool emitEvent) internal virtual override {
    if (spender == address(this)) revert InvalidRecipient(spender);

    super._approve(owner, spender, value, emitEvent);
  }

  /// @dev This check applies to transfer, minting, and burning.
  /// @dev Disallows approving for address(this).
  function _update(address from, address to, uint256 value) internal virtual override {
    if (to == address(this)) revert InvalidRecipient(to);

    super._update(from, to, value);
  }

  // ================================================================
  // │                            Roles                             │
  // ================================================================

  /// @notice Returns the current CCIPAdmin.
  function getCCIPAdmin() external view virtual returns (address) {
    return s_ccipAdmin;
  }

  /// @notice Transfers the CCIPAdmin role to a new address
  /// @dev only the owner can call this function, NOT the current ccipAdmin, and 1-step ownership transfer is used.
  /// @param newAdmin The address to transfer the CCIPAdmin role to. Setting to address(0) is a valid way to revoke
  /// the role.
  function setCCIPAdmin(
    address newAdmin
  ) external virtual onlyRole(DEFAULT_ADMIN_ROLE) {
    address currentAdmin = s_ccipAdmin;

    s_ccipAdmin = newAdmin;

    emit CCIPAdminTransferred(currentAdmin, newAdmin);
  }
}
