// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

//go:generate go run ../generation/wrap.go shared BurnMintERC677 burn_mint_erc677 latest
//go:generate go run ../generation/wrap.go shared ERC677 erc677 latest
//go:generate go run ../generation/wrap.go shared LinkToken link_token latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20 burn_mint_erc20 latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20WithDrip burn_mint_erc20_with_drip latest
//go:generate go run ../generation/wrap.go shared WERC20Mock werc20_mock latest
//go:generate go run ../generation/wrap.go shared ChainReaderTester chain_reader_tester latest
//go:generate go run ../generation/wrap.go shared AggregatorV3Interface aggregator_v3_interface latest
//go:generate go run ../generation/wrap.go shared MockV3Aggregator mock_v3_aggregator_contract latest
//go:generate go run ../generation/wrap.go shared LogEmitter log_emitter latest
//go:generate go run ../generation/wrap.go shared VRFLogEmitter vrf_log_emitter latest
//go:generate go run ../generation/wrap.go shared ITypeAndVersion type_and_version latest
//go:generate go run ../generation/wrap.go shared WETH9ZKSync weth9_zksync latest

//go:generate go run ../generation/wrap.go shared ERC20 erc20 latest
//go:generate go run ../generation/wrap.go shared Multicall3 multicall3 latest
//go:generate go run ../generation/wrap.go shared WETH9 weth9 latest

//go:generate go run ../generation/wrap.go shared BurnMintERC20PausableFreezableTransparent burn_mint_erc20_pausable_freezable_transparent latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20PausableFreezableUUPS burn_mint_erc20_pausable_freezable_uups latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20PausableTransparent burn_mint_erc20_pausable_transparent latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20PausableUUPS burn_mint_erc20_pausable_uups latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20Transparent burn_mint_erc20_transparent latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20UUPS burn_mint_erc20_uups latest
//go:generate go run ../generation/wrap.go shared IBurnMintERC20Upgradeable i_burn_mint_erc20_upgradeable latest
