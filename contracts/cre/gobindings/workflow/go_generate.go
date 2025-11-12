// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// Workflow

//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate workflow/v1 WorkflowRegistry workflow_registry_wrapper_v1
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate workflow/v2 WorkflowRegistry workflow_registry_wrapper_v2
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate workflow/v2 CapabilitiesRegistry capabilities_registry_wrapper_v2
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate workflow/dev MessageEmitter message_emitter
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate workflow/dev MockKeystoneForwarder mock_forwarder
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate workflow/dev ReserveManager reserve_manager

// Extract bytecode and ABI from generated wrappers
//go:generate go run github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/extract_bytecode -input=generated -bytecode=../../bytecode/workflow -abi=../../abi/workflow
