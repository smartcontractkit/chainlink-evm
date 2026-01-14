// Package gobindings provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gobindings

// Make sure solidity compiler artifacts are up-to-date. Only output stdout on failure.
//go:generate ../scripts/compile_all_keystone
//go:generate ../scripts/compile_all_workflow
//go:generate ../scripts/compile_all_shardconfig

//go:generate go generate ./keystone
//go:generate go generate ./workflow
//go:generate go generate ./shardconfig
