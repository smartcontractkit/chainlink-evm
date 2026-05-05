// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

//go:generate go run ../wrap l2ep ArbitrumValidator arbitrum_validator
//go:generate go run ../wrap l2ep ArbitrumSequencerUptimeFeed arbitrum_sequencer_uptime_feed
//go:generate go run ../wrap l2ep ArbitrumCrossDomainForwarder arbitrum_cross_domain_forwarder
//go:generate go run ../wrap l2ep ArbitrumCrossDomainGovernor arbitrum_cross_domain_governor

//go:generate go run ../wrap l2ep OptimismValidator optimism_validator
//go:generate go run ../wrap l2ep OptimismSequencerUptimeFeed optimism_sequencer_uptime_feed
//go:generate go run ../wrap l2ep OptimismCrossDomainForwarder optimism_cross_domain_forwarder
//go:generate go run ../wrap l2ep OptimismCrossDomainGovernor optimism_cross_domain_governor

//go:generate go run ../wrap l2ep ScrollValidator scroll_validator
//go:generate go run ../wrap l2ep ScrollSequencerUptimeFeed scroll_sequencer_uptime_feed
//go:generate go run ../wrap l2ep ScrollCrossDomainForwarder scroll_cross_domain_forwarder
//go:generate go run ../wrap l2ep ScrollCrossDomainGovernor scroll_cross_domain_governor
