// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

//go:generate go run ../wrap shared BurnMintERC677 burn_mint_erc677
//go:generate go run ../wrap shared ERC677 erc677
//go:generate go run ../wrap shared LinkToken link_token
//go:generate go run ../wrap shared BurnMintERC20 burn_mint_erc20
//go:generate go run ../wrap shared BurnMintERC20WithDrip burn_mint_erc20_with_drip
//go:generate go run ../wrap shared WERC20Mock werc20_mock
//go:generate go run ../wrap shared ChainReaderTester chain_reader_tester
//go:generate go run ../wrap shared AggregatorV3Interface aggregator_v3_interface
//go:generate go run ../wrap shared MockV3Aggregator mock_v3_aggregator_contract
//go:generate go run ../wrap shared LogEmitter log_emitter
//go:generate go run ../wrap shared VRFLogEmitter vrf_log_emitter
//go:generate go run ../wrap shared ITypeAndVersion type_and_version
//go:generate go run ../wrap shared WETH9ZKSync weth9_zksync

//go:generate go run ../wrap shared ERC20 erc20
//go:generate go run ../wrap shared Multicall3 multicall3
//go:generate go run ../wrap shared WETH9 weth9

//go:generate go run ../wrap shared BurnMintERC20PausableFreezableTransparent burn_mint_erc20_pausable_freezable_transparent
//go:generate go run ../wrap shared BurnMintERC20PausableFreezableUUPS burn_mint_erc20_pausable_freezable_uups
//go:generate go run ../wrap shared BurnMintERC20PausableTransparent burn_mint_erc20_pausable_transparent
//go:generate go run ../wrap shared BurnMintERC20PausableUUPS burn_mint_erc20_pausable_uups
//go:generate go run ../wrap shared BurnMintERC20Transparent burn_mint_erc20_transparent
//go:generate go run ../wrap shared BurnMintERC20UUPS burn_mint_erc20_uups
//go:generate go run ../wrap shared IBurnMintERC20Upgradeable i_burn_mint_erc20_upgradeable
