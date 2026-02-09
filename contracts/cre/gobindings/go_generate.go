// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// Make sure solidity compiler artifacts are up-to-date. Only output stdout on failure.
//go:generate ../scripts/compile_all_dev
//go:generate ../scripts/compile_all_v1
//go:generate ../scripts/compile_all_v2

//go:generate go generate ./dev
//go:generate go generate ./v1
//go:generate go generate ./v2
