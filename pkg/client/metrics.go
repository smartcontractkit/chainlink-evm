package client

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

type rpcClientMetrics struct {
	callsTotal   metric.Int64Counter
	callsSuccess metric.Int64Counter
	callsFailed  metric.Int64Counter
}

func newRPCClientMetrics() (*rpcClientMetrics, error) {
	callsTotal, err := beholder.GetMeter().Int64Counter("evm_pool_rpc_node_calls_total")
	if err != nil {
		return nil, fmt.Errorf("failed to register rpc calls total metric: %w", err)
	}

	callsSuccess, err := beholder.GetMeter().Int64Counter("evm_pool_rpc_node_calls_success")
	if err != nil {
		return nil, fmt.Errorf("failed to register rpc calls success metric: %w", err)
	}

	callsFailed, err := beholder.GetMeter().Int64Counter("evm_pool_rpc_node_calls_failed")
	if err != nil {
		return nil, fmt.Errorf("failed to register rpc calls failed metric: %w", err)
	}

	return &rpcClientMetrics{
		callsTotal:   callsTotal,
		callsSuccess: callsSuccess,
		callsFailed:  callsFailed,
	}, nil
}

func (m *rpcClientMetrics) IncrementTotal(ctx context.Context, chainID, nodeName, rpcDomain, callName string) {
	m.callsTotal.Add(ctx, 1, metric.WithAttributes(
		attribute.String("chainFamily", "EVM"),
		attribute.String("chainID", chainID),
		attribute.String("nodeName", nodeName),
		attribute.String("rpcDomain", rpcDomain),
		attribute.String("callName", callName),
	))
}

func (m *rpcClientMetrics) IncrementSuccess(ctx context.Context, chainID, nodeName, rpcDomain, callName string) {
	m.callsSuccess.Add(ctx, 1, metric.WithAttributes(
		attribute.String("chainFamily", "EVM"),
		attribute.String("chainID", chainID),
		attribute.String("nodeName", nodeName),
		attribute.String("rpcDomain", rpcDomain),
		attribute.String("callName", callName),
	))
}

func (m *rpcClientMetrics) IncrementFailed(ctx context.Context, chainID, nodeName, rpcDomain, callName string) {
	m.callsFailed.Add(ctx, 1, metric.WithAttributes(
		attribute.String("chainFamily", "EVM"),
		attribute.String("chainID", chainID),
		attribute.String("nodeName", nodeName),
		attribute.String("rpcDomain", rpcDomain),
		attribute.String("callName", callName),
	))
}
