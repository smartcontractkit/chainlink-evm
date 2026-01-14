// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

// Payments

//go:generate go run ../wrap payments PaymentTokenOnRamp payment_token_on_ramp
