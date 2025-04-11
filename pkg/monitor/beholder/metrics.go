package monitor

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	beholdercommon "github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/beholder"
)

// TODO: move to another directory / repo
// Define a new gauge metric for account balance
type GaugeAccBalance struct {
	// account_balance
	gauge metric.Float64Gauge
}

func NewGaugeAccBalance(unitStr string) (*GaugeAccBalance, error) {
	name := "account_balance"
	description := "Balance for configured WT account"
	gauge, err := beholdercommon.GetMeter().Float64Gauge(name, metric.WithUnit(unitStr), metric.WithDescription(description))
	if err != nil {
		return nil, fmt.Errorf("failed to create new gauge %s: %+w", name, err)
	}
	return &GaugeAccBalance{gauge}, nil
}

func (g *GaugeAccBalance) Record(ctx context.Context, balance float64, account string, chainInfo ChainInfo) {
	oAttrs := metric.WithAttributeSet(g.GetAttributes(account, chainInfo))
	g.gauge.Record(ctx, balance, oAttrs)
}

func (g *GaugeAccBalance) GetAttributes(account string, chainInfo ChainInfo) attribute.Set {
	return attribute.NewSet(
		attribute.String("account", account),

		// Execution Context - Source
		attribute.String("source_id", beholder.ValOrUnknown(account)), // reusing account as source_id
		// Execution Context - Chain
		attribute.String("chain_family_name", beholder.ValOrUnknown(chainInfo.ChainFamilyName)),
		attribute.String("chain_id", beholder.ValOrUnknown(chainInfo.ChainID)),
		attribute.String("network_name", beholder.ValOrUnknown(chainInfo.NetworkName)),
		attribute.String("network_name_full", beholder.ValOrUnknown(chainInfo.NetworkNameFull)),
	)
}
