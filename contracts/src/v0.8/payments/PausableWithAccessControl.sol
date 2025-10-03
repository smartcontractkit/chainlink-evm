// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {IPausable} from "./interfaces/IPausable.sol";

import {Roles} from "./libraries/Roles.sol";

import {AccessControlDefaultAdminRules} from
  "@openzeppelin/contracts@5.0.2/access/extensions/AccessControlDefaultAdminRules.sol";
import {IAccessControlEnumerable} from "@openzeppelin/contracts@5.0.2/access/extensions/IAccessControlEnumerable.sol";
import {Pausable} from "@openzeppelin/contracts@5.0.2/utils/Pausable.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @notice Base contract that adds pausing and access control functionality.
abstract contract PausableWithAccessControl is
  Pausable,
  AccessControlDefaultAdminRules,
  IPausable,
  IAccessControlEnumerable
{
  using EnumerableSet for EnumerableSet.AddressSet;

  /// @notice The set of members in each role
  mapping(bytes32 role => EnumerableSet.AddressSet) private s_roleMembers;

  constructor(
    uint48 adminRoleTransferDelay,
    address admin
  ) AccessControlDefaultAdminRules(adminRoleTransferDelay, admin) {}

  /// @notice This function pauses the contract
  /// @dev Sets the pause flag to true
  function emergencyPause() external onlyRole(Roles.PAUSER_ROLE) {
    _pause();
  }

  /// @inheritdoc AccessControlDefaultAdminRules
  function supportsInterface(
    bytes4 interfaceId
  ) public view virtual override returns (bool) {
    return interfaceId == type(IAccessControlEnumerable).interfaceId || super.supportsInterface(interfaceId);
  }

  /// @notice This function unpauses the contract
  /// @dev Sets the pause flag to false
  function emergencyUnpause() external onlyRole(Roles.UNPAUSER_ROLE) {
    _unpause();
  }

  /// @inheritdoc IAccessControlEnumerable
  function getRoleMember(bytes32 role, uint256 index) external view override returns (address) {
    return s_roleMembers[role].at(index);
  }

  /// @inheritdoc IAccessControlEnumerable
  function getRoleMemberCount(
    bytes32 role
  ) external view override returns (uint256) {
    return s_roleMembers[role].length();
  }

  /// @notice This function returns the members of a role
  /// @param role The role to get the members of
  /// @return roleMembers members of the role
  function getRoleMembers(
    bytes32 role
  ) public view virtual returns (address[] memory roleMembers) {
    return s_roleMembers[role].values();
  }

  /// @inheritdoc AccessControlDefaultAdminRules
  function _grantRole(bytes32 role, address account) internal virtual override returns (bool) {
    bool granted = super._grantRole(role, account);
    if (granted) {
      s_roleMembers[role].add(account);
    }
    return granted;
  }

  /// @inheritdoc AccessControlDefaultAdminRules
  function _revokeRole(bytes32 role, address account) internal virtual override returns (bool) {
    bool revoked = super._revokeRole(role, account);
    if (revoked) {
      s_roleMembers[role].remove(account);
    }
    return revoked;
  }
}
