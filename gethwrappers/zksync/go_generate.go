package gethwrappers

//go:generate ../../contracts/scripts/zksync_compile_all

//go:generate go run ../generation/zksync/wrap.go shared LinkToken link_token
//go:generate go run ../generation/zksync/wrap.go shared BurnMintERC677 burn_mint_erc677
//go:generate go run ../generation/zksync/wrap.go shared Multicall3 multicall3
//go:generate go run ../generation/zksync/wrap.go keystone CapabilitiesRegistry capabilities_registry
