// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.24;

import {Ownable2StepMsgSender} from "../../shared/access/Ownable2StepMsgSender.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @title OwnershipProof
/// @notice This contract allows the transaction sender to submit or revoke ownership proof.
/// THIS IS A WORK IN PROGRESS AND IS NOT READY FOR REVIEWS YET.
abstract contract OwnershipProof is Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  mapping(address signer => bool allowed) private s_allowedSigners;
  mapping(address owner => bytes32 proof) private s_ownersProofs;
  EnumerableSet.AddressSet private s_ownersProofsSet;

  event AllowedSignersUpdatedV1(address[] signers, bool allowed);
  event OwnershipProofSubmittedV1(address indexed owner, bytes32 indexed proof);
  event OwnershipProofRevokedV1(address indexed owner, bytes32 indexed proof);

  error RequestExpired(address caller, uint96 validityTimestamp);
  error OwnershipProofAlreadySubmitted(address owner);
  error OwnershipProofNotSubmitted(address owner);
  error InvalidSignatureLength(bytes signature);
  error InvalidSValue(bytes signature, bytes32 s);
  error InvalidVValue(bytes signature, uint8 v);
  error InvalidOwnershipProof(address owner, uint96 validityTimestamp, bytes32 proof, bytes signature);

  /// @notice Modifier to check if the ownership proof is submitted by the transaction sender.
  /// @dev This modifier can be used to restrict access to functions in the derived contract that require
  // the ownership proof to be submitted. If the ownership proof is not submitted, the transaction will revert.
  modifier hasOwnershipProofSubmitted() {
    if (s_ownersProofs[msg.sender] == bytes32(0)) {
      revert OwnershipProofNotSubmitted(msg.sender);
    }
    _;
  }

  /// @notice Sets the allowed signers for ownership proofs.
  /// @param signers The addresses of the signers
  /// @param allowed The boolean value indicating whether the signer is allowed or not
  /// @dev Ownership proofs can only be signed by approved group of signers.
  /// When submitting signed proof to this contract, if recovered signature doesn't match any of the signers,
  /// it will be rejected.
  function updateAllowedSigners(address[] calldata signers, bool allowed) external onlyOwner {
    for (uint256 i = 0; i < signers.length; i++) {
      s_allowedSigners[signers[i]] = allowed;
    }
    emit AllowedSignersUpdatedV1(signers, allowed);
  }

  /// @notice Returns the allowed signer for ownership proofs.
  /// @param signer The address of the signer
  /// @return The boolean value indicating whether the signer is allowed to sign ownership proofs or not
  function isAllowedSigner(
    address signer
  ) external view returns (bool) {
    return s_allowedSigners[signer];
  }

  /// @notice Transaction sender submits ownership proof for verification and approval.
  /// @param validityTimestamp Validity of the ownership proof.
  /// @param proof The ownership proof to be submitted.
  /// @param signature The signature of the ownership proof metadata.
  /// @dev The ownership proof is a hash of the ownership proof metadata, which is signed by the allowed signer.
  /// The ownership proof metadata is a combination of the claimed owner address, validity timestamp, and the proof hash.
  /// Request will be rejected if the validity timestamp has expired or if the proof has already been submitted (by this
  /// transaction sender or a different one).
  /// If the transaction sender is truly the owner of the claimed address, the proof will be acccepted if the
  /// signature is valid and recoverable signer matches any of the allowed signers.
  function submitOwnershipProof(uint96 validityTimestamp, bytes32 proof, bytes calldata signature) external {
    if (uint96(block.timestamp) > validityTimestamp) {
      revert RequestExpired(msg.sender, validityTimestamp);
    }

    if (s_ownersProofs[msg.sender] != bytes32(0)) {
      revert OwnershipProofAlreadySubmitted(msg.sender);
    }

    address signer = _recoverSigner(validityTimestamp, proof, signature);
    if (!s_allowedSigners[signer]) {
      revert InvalidOwnershipProof(msg.sender, validityTimestamp, proof, signature);
    }

    s_ownersProofs[msg.sender] = proof;
    s_ownersProofsSet.add(msg.sender);
    emit OwnershipProofSubmittedV1(msg.sender, proof);
  }

  /// @notice Transaction sender submits ownership proof for verification and revocation.
  /// @param validityTimestamp Validity of the ownership proof.
  /// @param proof The ownership proof to be submitted.
  /// @param signature The signature of the ownership proof metadata.
  /// @dev The ownership proof is a hash of the ownership proof metadata, which is signed by the allowed signer.
  /// The ownership proof metadata is a combination of the claimed owner address, validity timestamp, and the proof hash.
  /// Request will be rejected if the validity timestamp has expired or if the proof hasn't been submitted by this
  /// transaction sender (can't revoke someone else's proof).
  /// If the transaction sender is truly the owner of the claimed address, the proof will be revoked if the
  /// signature is valid and recoverable signer matches any of the allowed signers.
  function revokeOwnershipProof(uint96 validityTimestamp, bytes32 proof, bytes calldata signature) external {
    if (uint96(block.timestamp) > validityTimestamp) {
      revert RequestExpired(msg.sender, validityTimestamp);
    }

    if (s_ownersProofs[msg.sender] != proof) {
      revert OwnershipProofNotSubmitted(msg.sender);
    }

    address signer = _recoverSigner(validityTimestamp, proof, signature);
    if (!s_allowedSigners[signer]) {
      revert InvalidOwnershipProof(msg.sender, validityTimestamp, proof, signature);
    }

    s_ownersProofs[msg.sender] = bytes32(0);
    s_ownersProofsSet.remove(msg.sender);
    emit OwnershipProofRevokedV1(msg.sender, proof);
  }

  /// @notice Admin revokes ownership proof for the given owner.
  /// @param owner The address of the owner
  /// @dev This function can only be called by the contract owner. It will remove the ownership proof for the given
  /// owner and emit an revocation event. It should only be used in case of emergency if the owner has lost
  /// access to the private key and can't revoke the proof themselves.
  function adminRevokeOwnershipProof(
    address owner
  ) external onlyOwner {
    if (s_ownersProofs[owner] == bytes32(0)) {
      revert OwnershipProofNotSubmitted(owner);
    }

    s_ownersProofs[owner] = bytes32(0);
    s_ownersProofsSet.remove(owner);
    emit OwnershipProofRevokedV1(owner, bytes32(0));
  }

  //// @notice Returns if the ownership proof is submitted for the given owner.
  /// @param owner The address of the owner
  /// @return True if the ownership proof is submitted, false otherwise
  function isProofSubmitted(
    address owner
  ) external view returns (bool) {
    return s_ownersProofs[owner] != bytes32(0);
  }

  /// @notice Returns a list of ownership proofs for the given range of stored owners.
  /// @param start The starting index of the range
  /// @param batchSize The size of the batch to return (be vary of using a reasonable value here)
  /// @return A list of ownership proofs for the given range
  /// @dev The function returns a list of ownership proofs for the given range of stored owners.
  /// The function will return a list of addresses of owners who have submitted ownership proofs in the
  /// order in which they were submitted.
  /// The function will return an empty list if the start index is greater than the total number of owners.
  /// @notice List of owners may change between subsequent calls to this function. It is recommend to call this
  /// function by anchoring a specific block number to ensure the list of owners is immutable between subsequent calls.
  function getOwnersWithSubmittedProof(uint256 start, uint256 batchSize) external view returns (address[] memory) {
    uint256 total = s_ownersProofsSet.length();
    if (start >= total) {
      return new address[](0);
    }

    uint256 end = (start + batchSize > total) ? total : start + batchSize;
    uint256 length = end - start;
    address[] memory owners = new address[](length);

    for (uint256 i = 0; i < length; ++i) {
      owners[i] = s_ownersProofsSet.at(start + i);
    }

    return owners;
  }

  /// @notice Returns the signer of the recovered signature or revert.
  /// @param validityTimestamp The validity timestamp of the ownership proof.
  /// @param proof The ownership proof.
  /// @param signature The signature of the ownership proof metadata.
  /// @return The signer of the recovered signature.
  /// @dev The function tries to re-generate the message digest based on the provided parameters and by following
  /// EIP-191. The it will try to recover the signer address. The function will revert if the signature is invalid.
  function _recoverSigner(
    uint96 validityTimestamp,
    bytes32 proof,
    bytes calldata signature
  ) internal view returns (address) {
    // Follow EIP-191 for recoverable signatures
    bytes32 messageHash = keccak256(abi.encodePacked(msg.sender, block.chainid, validityTimestamp, proof));
    bytes32 prefixedMessageHash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", messageHash));

    if (signature.length != 65) {
      revert InvalidSignatureLength(signature);
    }

    bytes32 r;
    bytes32 s;
    uint8 v;

    assembly {
      // Load r from signature (offset 0)
      r := calldataload(signature.offset)
      // Load s from signature (offset 32)
      s := calldataload(add(signature.offset, 32))
      // Load v from signature (offset 64) and mask it to a byte
      v := byte(0, calldataload(add(signature.offset, 64)))
    }

    // Normalize v: some libraries produce signatures where v is 0 or 1, while Ethereum expects v to be 27 or 28 for ecrecover.
    if (v < 27) {
      v += 27;
    }

    // EIP-2 still allows signature malleability for ecrecover(), remove this possibility and make the signature unique
    if (uint256(s) > 0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff) {
      revert InvalidSValue(signature, s);
    }

    if (v != 27 && v != 28) {
      revert InvalidVValue(signature, v);
    }

    return ecrecover(prefixedMessageHash, v, r, s);
  }
}
