// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// Keystone

//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate core/v1 BalanceReader balance_reader
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate core/v1 CapabilitiesRegistry capabilities_registry_v1
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate core/v1 KeystoneFeedsConsumer feeds_consumer
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate core/v1 KeystoneForwarder forwarder
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate core/v1 OCR3Capability ocr3_capability
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate core/v2 CapabilitiesRegistry capabilities_registry_wrapper_v2

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/core -abi=../../abi/core
