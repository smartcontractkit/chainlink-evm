// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// Keystone

//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate keystone BalanceReader balance_reader
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate keystone CapabilitiesRegistry capabilities_registry
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate keystone KeystoneFeedsConsumer feeds_consumer
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate keystone KeystoneForwarder forwarder
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate keystone OCR3Capability ocr3_capability
