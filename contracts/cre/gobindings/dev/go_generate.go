// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// CRE Dev

//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate dev ShardConfig shard_config v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate dev MessageEmitter message_emitter v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate dev MockKeystoneForwarder mock_forwarder v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate dev ReserveManager reserve_manager v1_0_0

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/dev -abi=../../abi/dev
