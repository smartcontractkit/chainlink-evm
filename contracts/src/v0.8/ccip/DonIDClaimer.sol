// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable2StepMsgSender} from "../shared/access/Ownable2StepMsgSender.sol";
import {ITypeAndVersion} from "../shared/interfaces/ITypeAndVersion.sol";
import {EnumerableSet} from "../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

interface ICapabilitiesRegistry {
  /// @notice Gets the next available DON ID from the CapabilitiesRegistry
  /// @return uint32 The next available DON ID
  function getNextDONId() external view returns (uint32);
}

/// @notice DonIDClaimer contract is used to claim and manage DON IDs. It interacts with
/// the CapabilitiesRegistry to fetch the next available DON ID and allows
/// for synchronization of the DON ID with an optional offset to skip certain
/// DON IDs as needed. The contract provides functionality for claiming,
/// retrieving, and syncing DON IDs, ensuring that multiple workflows can
/// manage DON IDs without conflict or accidental reuse.
/// @dev The contract maintains its own internal counter for DON IDs and ensures
/// the next available ID is claimed and tracked by the contract. The sync function
/// allows for alignment with the CapabilitiesRegistry.
contract DonIDClaimer is ITypeAndVersion, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  error ZeroAddressNotAllowed();
  error AccessForbidden(address sender);

  event AuthorizedDeployerSet(address indexed senderAddress, bool allowed);
  event DonIDClaimed(address indexed claimer, uint32 donId);
  event DonIDSynced(uint32 newDONId);

  string public constant override typeAndVersion = "DonIDClaimer 1.6.1";
  /// @notice The next available DON ID that is claimed and incremented
  uint32 private s_nextDONId;

  /// @notice The address of the CapabilitiesRegistry contract used to fetch the next DON ID
  ICapabilitiesRegistry private immutable i_capabilitiesRegistry;

  /// @notice List to track authorized deployed keys
  EnumerableSet.AddressSet private s_authorizedDeployers;

  /// @notice Initializes the contract with the CapabilitiesRegistry address
  /// @param _capabilitiesRegistry The address of the CapabilitiesRegistry contract
  constructor(
    address _capabilitiesRegistry
  ) {
    if (_capabilitiesRegistry == address(0)) revert ZeroAddressNotAllowed();
    i_capabilitiesRegistry = ICapabilitiesRegistry(_capabilitiesRegistry);

    // Initializing the deployer authorization (owner can be the initial deployer)
    s_authorizedDeployers.add(msg.sender);

    // Sync the initial s_nextDONId from the CapabilitiesRegistry contract
    s_nextDONId = i_capabilitiesRegistry.getNextDONId();
  }

  /// @notice Modifier to check if the caller is an authorized deployer
  modifier onlyAuthorizedDeployer() {
    if (!s_authorizedDeployers.contains(msg.sender)) {
      revert AccessForbidden(msg.sender);
    }
    _;
  }

  /// @notice Claims the next available DON ID and increments the internal counter
  /// @dev The function increments s_nextDONId after returning the current value
  /// @return uint32 The DON ID that was claimed
  function claimNextDONId() external onlyAuthorizedDeployer returns (uint32) {
    emit DonIDClaimed(msg.sender, s_nextDONId);

    return s_nextDONId++;
  }

  /// @notice Synchronizes the next donID with the CapabilitiesRegistry and applies an offset
  /// @param offset The offset to adjust the donID (useful when certain DON IDs are dropped)
  /// @dev This can be used to synchronize with the CapabilitiesRegistry after some actions have occurred
  function syncNextDONIdWithOffset(
    uint32 offset
  ) external onlyAuthorizedDeployer {
    s_nextDONId = i_capabilitiesRegistry.getNextDONId() + offset;

    emit DonIDSynced(s_nextDONId);
  }

  /// @notice Sets authorization status for a deployer address
  /// @param senderAddress The address to be added or removed as an authorized deployer
  /// @param allowed Boolean indicating whether the address is authorized (true) or revoked (false)
  /// @dev Can only be called by an existing authorized deployer
  function setAuthorizedDeployer(address senderAddress, bool allowed) external onlyOwner {
    if (senderAddress == address(0)) revert ZeroAddressNotAllowed();

    if (allowed) {
      s_authorizedDeployers.add(senderAddress);
    } else {
      s_authorizedDeployers.remove(senderAddress);
    }

    emit AuthorizedDeployerSet(senderAddress, allowed);
  }

  /// @notice Returns the next available donID
  /// @return uint32 The next available donID to be claimed
  function getNextDONId() external view returns (uint32) {
    return s_nextDONId;
  }

  /// @notice Checks if an address is an authorized deployer
  /// @param senderAddress The address to check for authorization
  /// @return bool True if the address is an authorized deployer, false otherwise
  function isAuthorizedDeployer(
    address senderAddress
  ) external view returns (bool) {
    return s_authorizedDeployers.contains(senderAddress);
  }
}
