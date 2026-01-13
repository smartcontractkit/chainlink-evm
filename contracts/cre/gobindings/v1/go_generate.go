// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// CRE V1: Keystone/ST3

//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate v1 WorkflowRegistry workflow_registry_wrapper v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate v1 CapabilitiesRegistry capabilities_registry_wrapper v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate v1 BalanceReader balance_reader v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate v1 KeystoneFeedsConsumer feeds_consumer v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate v1 KeystoneForwarder forwarder v1_0_0
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate v1 OCR3Capability ocr3_capability v1_0_0

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/v1 -abi=../../abi/v1
