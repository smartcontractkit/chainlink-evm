// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

//go:generate go run ../generation/wrap.go shared BurnMintERC677 burn_mint_erc677 latest
//go:generate go run ../generation/wrap.go shared ERC677 erc677 latest
//go:generate go run ../generation/wrap.go shared LinkToken link_token latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20 burn_mint_erc20 latest
//go:generate go run ../generation/wrap.go shared BurnMintERC20 burn_mint_erc20_with_drip latest
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
