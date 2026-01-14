// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// CRE V1: Mainline CRE

//go:generate go run ../wrap v2 WorkflowRegistry workflow_registry_wrapper v2_0_0
//go:generate go run ../wrap v2 CapabilitiesRegistry capabilities_registry_wrapper v2_0_0

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/v2 -abi=../../abi/v2
