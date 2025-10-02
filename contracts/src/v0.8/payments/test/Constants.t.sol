// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

abstract contract Constants {
  bytes32 internal constant DEFAULT_ADMIN_ROLE = 0x00;
  uint48 internal constant DEFAULT_ADMIN_TRANSFER_DELAY = 0;

  bytes internal constant SENDER_1 = bytes("SENDER_1");
  bytes internal constant SENDER_2 = bytes("SENDER_2");
  bytes internal constant RECEIVER_1 = bytes("RECEIVER_1");
  bytes internal constant RECEIVER_2 = bytes("RECEIVER_2");
  bytes internal constant RECEIVER_3 = bytes("RECEIVER_3");
  uint64 internal constant SOURCE_CHAIN_1 = 2;
  uint64 internal constant SOURCE_CHAIN_2 = 3;
  uint64 internal constant DESTINATION_CHAIN_1 = 2;
  uint64 internal constant DESTINATION_CHAIN_2 = 3;
  uint64 internal constant INVALID_SOURCE_CHAIN = 123;
  uint64 internal constant INVALID_DESTINATION_CHAIN = 123;

  uint32 internal constant STALENESS_THRESHOLD = (1 days);
  bytes internal constant ASSET_1_SWAP_PATH = bytes("ASSET_1_SWAP_PATH");
  bytes internal constant ASSET_2_SWAP_PATH = bytes("ASSET_2_SWAP_PATH");
  bytes internal constant EMPTY_SWAP_PATH = bytes("");
  uint32 internal constant SWAP_INTERVAL = 1 hours;
  uint64 internal constant DESTINATION_CHAIN_SELECTOR = 4949039107694359620;
  uint32 internal constant DESTINATION_CHAIN_GAS_LIMIT = 500_000;
  uint16 internal constant MAX_SLIPPAGE = 200;
  uint128 internal constant MIN_SWAP_SIZE = 1_000e8;
  uint128 internal constant MAX_SWAP_SIZE = 100_000e8;
  uint64 internal constant MAX_GAS_PRICE = 100 gwei;
  uint16 internal constant MAX_PRICE_DEVIATION = 200;
  uint16 internal constant MAX_PRICE_DEVIATION_INVARIANTS = 200;
  uint24 internal constant UNI_POOL_FEE = 3000;
  uint96 internal constant DEADLINE_DELAY = 1 minutes;
  uint96 internal constant MIN_DEADLINE_DELAY = 1 minutes;
  uint96 internal constant MAX_DEADLINE_DELAY = 1 hours;
  uint96 internal constant FEE_RESERVE_INITIAL_LINK_BALANCE = 10 ether;
  uint256 internal constant MAX_PERFORM_DATA_SIZE = 2000;
  uint96 internal constant CHAIN_SELECTOR = 5009297550715157269;

  bytes32 public constant TEST_ROLE = keccak256("TEST_ROLE");
}
