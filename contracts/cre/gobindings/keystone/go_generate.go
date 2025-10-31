// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// Keystone

//go:generate go run ../../../../gethwrappers/generation/wrap.go keystone BalanceReader balance_reader
//go:generate go run ../../../../gethwrappers/generation/wrap.go keystone CapabilitiesRegistry capabilities_registry
//go:generate go run ../../../../gethwrappers/generation/wrap.go keystone KeystoneFeedsConsumer feeds_consumer
//go:generate go run ../../../../gethwrappers/generation/wrap.go keystone KeystoneForwarder forwarder
//go:generate go run ../../../../gethwrappers/generation/wrap.go keystone OCR3Capability ocr3_capability
