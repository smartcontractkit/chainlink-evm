// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// ShardConfig

//go:generate go run ../wrap shardconfig ShardConfig shard_config

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/shardconfig -abi=../../abi/shardconfig
