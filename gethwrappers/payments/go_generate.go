// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

// Payments

//go:generate go run ../generation/wrap.go payments PaymentTokenOnRamp payment_token_on_ramp latest
