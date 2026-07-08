package config

import (
	"net/url"
	"time"

	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
)

type transactionsConfig struct {
	c toml.Transactions
}

func (t *transactionsConfig) Enabled() bool {
	return *t.c.Enabled
}

func (t *transactionsConfig) ForwardersEnabled() bool {
	return *t.c.ForwardersEnabled
}

func (t *transactionsConfig) ReaperInterval() time.Duration {
	return t.c.ReaperInterval.Duration()
}

func (t *transactionsConfig) ReaperThreshold() time.Duration {
	return t.c.ReaperThreshold.Duration()
}

func (t *transactionsConfig) ResendAfterThreshold() time.Duration {
	return t.c.ResendAfterThreshold.Duration()
}

func (t *transactionsConfig) MaxInFlight() uint32 {
	return *t.c.MaxInFlight
}

func (t *transactionsConfig) MaxQueued() uint64 {
	return uint64(*t.c.MaxQueued)
}

func (t *transactionsConfig) TransactionManagerV2() TransactionManagerV2 {
	return &transactionManagerV2Config{c: t.c.TransactionManagerV2}
}

type transactionManagerV2Config struct {
	c toml.TransactionManagerV2Config
}

func (t *transactionManagerV2Config) Enabled() bool {
	return *t.c.Enabled
}

func (t *transactionManagerV2Config) BlockTime() *time.Duration {
	d := t.c.BlockTime.Duration()
	return &d
}

func (t *transactionManagerV2Config) CustomURLs() []*url.URL {
	if len(t.c.CustomURLs) > 0 {
		out := make([]*url.URL, 0, len(t.c.CustomURLs))
		for _, u := range t.c.CustomURLs {
			if u != nil {
				out = append(out, u.URL())
			}
		}
		return out
	}
	// fall back to deprecated CustomURL if no CustomURLs are configured
	if t.c.CustomURL != nil {
		return []*url.URL{t.c.CustomURL.URL()}
	}
	return nil
}

func (t *transactionManagerV2Config) DualBroadcast() *bool {
	return t.c.DualBroadcast
}

func (t *transactionManagerV2Config) ReadRequestsToMultipleNodes() *bool {
	return t.c.ReadRequestsToMultipleNodes
}

func (t *transactionManagerV2Config) Bundles() *bool {
	return t.c.Bundles
}

func (t *transactionManagerV2Config) FastlaneAuctionRequestTimeout() *time.Duration {
	if t.c.FastlaneAuctionRequestTimeout == nil {
		return nil
	}
	d := t.c.FastlaneAuctionRequestTimeout.Duration()
	return &d
}

func (t *transactionManagerV2Config) FeeBoost() bool {
	return t.c.FeeBoost != nil && *t.c.FeeBoost
}

func (t *transactionsConfig) AutoPurge() AutoPurgeConfig {
	return &autoPurgeConfig{c: t.c.AutoPurge}
}

func (t *transactionsConfig) HederaSequencePollTimeout() *time.Duration {
	if t.c.HederaBroadcastValidation.SequencePollTimeout == nil {
		return nil
	}
	d := t.c.HederaBroadcastValidation.SequencePollTimeout.Duration()
	return &d
}

func (t *transactionsConfig) HederaSequencePollInterval() *time.Duration {
	if t.c.HederaBroadcastValidation.SequencePollInterval == nil {
		return nil
	}
	d := t.c.HederaBroadcastValidation.SequencePollInterval.Duration()
	return &d
}

type autoPurgeConfig struct {
	c toml.AutoPurgeConfig
}

func (a *autoPurgeConfig) Enabled() bool {
	return *a.c.Enabled
}

func (a *autoPurgeConfig) Threshold() *uint32 {
	return a.c.Threshold
}

func (a *autoPurgeConfig) MinAttempts() *uint32 {
	return a.c.MinAttempts
}

func (a *autoPurgeConfig) DetectionApiUrl() *url.URL {
	return a.c.DetectionApiUrl.URL()
}
