// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// Keystone

//go:generate go run ../wrap keystone BalanceReader balance_reader
//go:generate go run ../wrap keystone CapabilitiesRegistry capabilities_registry
//go:generate go run ../wrap keystone KeystoneFeedsConsumer feeds_consumer
//go:generate go run ../wrap keystone KeystoneForwarder forwarder
//go:generate go run ../wrap keystone OCR3Capability ocr3_capability

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/keystone -abi=../../abi/keystone
