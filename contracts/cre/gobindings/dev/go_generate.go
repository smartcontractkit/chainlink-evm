// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// CRE Dev

//go:generate go run ../wrap dev ShardConfig shard_config
//go:generate go run ../wrap dev MessageEmitter message_emitter
//go:generate go run ../wrap dev MockKeystoneForwarder mock_forwarder
//go:generate go run ../wrap dev ReserveManager reserve_manager

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/dev -abi=../../abi/dev
