package gethwrappers

//go:generate ../../contracts/scripts/zksync_compile_all

//go:generate go run ../generation/zksync/wrap.go shared LinkToken link_token latest
//go:generate go run ../generation/zksync/wrap.go shared BurnMintERC677 burn_mint_erc677 latest
//go:generate go run ../generation/zksync/wrap.go shared Multicall3 multicall3 latest
//go:generate go run ../generation/zksync/wrap.go shared WETH9ZKSync weth9_zksync latest
//go:generate go run ../generation/zksync/wrap.go shared MockV3Aggregator mock_v3_aggregator_contract latest

//go:generate go run ../generation/zksync/wrap.go automation MockETHUSDAggregator mock_ethusd_aggregator_wrapper latest
