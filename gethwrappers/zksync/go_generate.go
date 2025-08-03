package gethwrappers

//go:generate ../../contracts/scripts/zksync_compile_all

//go:generate go run ../generation/zksync/wrap.go shared LinkToken link_token
//go:generate go run ../generation/zksync/wrap.go shared BurnMintERC677 burn_mint_erc677
//go:generate go run ../generation/zksync/wrap.go shared Multicall3 multicall3
//go:generate go run ../generation/zksync/wrap.go shared WETH9ZKSync weth9_zksync
//go:generate go run ../generation/zksync/wrap.go shared MockV3Aggregator mock_v3_aggregator_contract

//go:generate go run ../generation/zksync/wrap.go automation MockETHUSDAggregator mock_ethusd_aggregator_wrapper
