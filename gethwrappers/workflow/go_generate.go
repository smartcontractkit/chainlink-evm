// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

// Workflow

//go:generate go run ../generation/wrap.go workflow/v1 WorkflowRegistry workflow_registry_wrapper_v1
//go:generate go run ../generation/wrap.go workflow/v2 WorkflowRegistry workflow_registry_wrapper_v2
//go:generate go run ../generation/wrap.go workflow/v2 CapabilitiesRegistry capabilities_registry_wrapper_v2
//go:generate go run ../generation/wrap.go workflow/dev MessageEmitter message_emitter
//go:generate go run ../generation/wrap.go workflow/dev MockKeystoneForwarder mock_forwarder
//go:generate go run ../generation/wrap.go workflow/dev ReserveManager reserve_manager
