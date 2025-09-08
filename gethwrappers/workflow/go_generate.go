// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

// Workflow

//go:generate go run ../generation/wrap.go workflow/v1 WorkflowRegistry workflow_registry_wrapper_v1
/// [FROZEN] go:generate go run ../generation/wrap.go workflow/dev/v2 WorkflowRegistry workflow_registry_wrapper_v2
//go:generate go run ../generation/wrap.go workflow/dev/v2 WorkflowRegistry workflow_registry_wrapper_v2_dev
/// [FROZEN] go:generate go run ../generation/wrap.go workflow/dev/v2 CapabilitiesRegistry capabilities_registry_wrapper_v2
//go:generate go run ../generation/wrap.go workflow/dev/v2 CapabilitiesRegistry capabilities_registry_wrapper_v2_dev
