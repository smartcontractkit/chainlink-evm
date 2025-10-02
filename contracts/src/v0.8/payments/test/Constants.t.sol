// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

abstract contract Constants {
  bytes32 internal constant DEFAULT_ADMIN_ROLE = 0x00;
  uint48 internal constant DEFAULT_ADMIN_TRANSFER_DELAY = 0;
  uint96 internal constant CHAIN_SELECTOR = 5_009_297_550_715_157_269;
  bytes32 public constant TEST_ROLE = keccak256("TEST_ROLE");
}
