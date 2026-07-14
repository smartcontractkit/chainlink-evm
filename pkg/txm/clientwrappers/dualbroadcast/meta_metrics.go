package dualbroadcast

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	pb "github.com/smartcontractkit/chainlink-protos/svr/v1"
)

// MetaMetrics handles all Meta-related metrics via OTEL
type MetaMetrics struct {
	chainID           string
	statusCodeCounter metric.Int64Counter
	latencyHistogram  metric.Int64Histogram
	bidHistogram      metric.Int64Histogram
	errorCounter      metric.Int64Counter
	emitter           beholder.Emitter
	lggr              logger.SugaredLogger
}

// NewMetaMetrics creates a new MetaMetrics instance
func NewMetaMetrics(chainID string, lggr logger.Logger) (*MetaMetrics, error) {
	statusCodeCounter, err := beholder.GetMeter().Int64Counter("meta_endpoint_status_codes")
	if err != nil {
		return nil, err
	}

	latencyHistogram, err := beholder.GetMeter().Int64Histogram("meta_endpoint_latency",
		metric.WithUnit("ms"),
		metric.WithDescription("Latency of Meta auction endpoint requests"),
		metric.WithExplicitBucketBoundaries(500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 7500, 10000),
	)
	if err != nil {
		return nil, err
	}

	bidHistogram, err := beholder.GetMeter().Int64Histogram("meta_bids_per_transaction")
	if err != nil {
		return nil, err
	}

	errorCounter, err := beholder.GetMeter().Int64Counter("meta_errors")
	if err != nil {
		return nil, err
	}

	return &MetaMetrics{
		chainID:           chainID,
		statusCodeCounter: statusCodeCounter,
		latencyHistogram:  latencyHistogram,
		bidHistogram:      bidHistogram,
		errorCounter:      errorCounter,
		emitter:           beholder.GetEmitter(),
		lggr:              logger.Sugared(logger.Named(lggr, "Txm.MetaClient.MetaMetrics")),
	}, nil
}

// RecordStatusCode records the HTTP status code from Meta endpoint
func (m *MetaMetrics) RecordStatusCode(ctx context.Context, statusCode int, feedAddress string) {
	m.statusCodeCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("statusCode", strconv.Itoa(statusCode)),
			attribute.String("contract_address", feedAddress),
		),
	)
}

// RecordLatency records the latency of Meta endpoint requests
func (m *MetaMetrics) RecordLatency(ctx context.Context, duration time.Duration, feedAddress string) {
	m.latencyHistogram.Record(ctx, duration.Milliseconds(),
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("contract_address", feedAddress),
		),
	)
}

// RecordBidsReceived records the distribution of bids per transaction
func (m *MetaMetrics) RecordBidsReceived(ctx context.Context, bidCount int, feedAddress string) {
	m.bidHistogram.Record(ctx, int64(bidCount),
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("contract_address", feedAddress),
		),
	)
}

// RecordSendRequestError records errors from SendRequest method
func (m *MetaMetrics) RecordSendRequestError(ctx context.Context, feedAddress string) {
	m.errorCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("errorType", "send_request"),
			attribute.String("contract_address", feedAddress),
		),
	)
}

// RecordSendOperationError records errors from SendOperation method
func (m *MetaMetrics) RecordSendOperationError(ctx context.Context, feedAddress string) {
	m.errorCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("errorType", "send_operation"),
			attribute.String("contract_address", feedAddress),
		),
	)
}

// emitAtlasError emits an OTel event to track FastLane Atlas errors
func (m *MetaMetrics) emitAtlasError(ctx context.Context, errType string, customURL *url.URL, cause error, tx *types.Transaction) {
	var nonce string
	if tx.Nonce != nil {
		nonce = fmt.Sprintf("%d", *tx.Nonce)
	}

	meta, err := tx.GetMeta()
	if err != nil {
		m.lggr.Errorw(fmt.Sprintf("Failed to get meta for tx. Error to emit was: %v", cause), "txId", tx.ID, "err", err)
		return
	}

	var destAddress string
	if meta != nil && meta.FwdrDestAddress != nil {
		destAddress = meta.FwdrDestAddress.String()
	}

	msg := &pb.FastLaneAtlasError{
		ChainId:       m.chainID,
		FromAddress:   tx.FromAddress.Hex(),
		ToAddress:     tx.ToAddress.Hex(),
		FeedAddress:   destAddress,
		Nonce:         nonce,
		ErrorType:     errType,
		ErrorMessage:  cause.Error(),
		TransactionId: tx.ID,
		AtlasUrl:      customURL.String(),
		CreatedAt:     time.Now().UnixMicro(),
	}

	messageBytes, err := proto.Marshal(msg)
	if err != nil {
		m.lggr.Errorw("Failed to marshal Atlas error event", "err", err)
		return
	}

	attrKVs := []any{
		"beholder_domain", "svr",
		"beholder_entity", "svr.v1.FastLaneAtlasError",
		"beholder_data_schema", "/fastlane-atlas-error/versions/1",
	}

	mStr := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Format(msg)
	m.lggr.Debugw("[Beholder.emit]", "message", mStr, "attributes", attrKVs)

	if emitErr := m.emitter.Emit(ctx, messageBytes, attrKVs...); emitErr != nil {
		m.lggr.Errorw("Failed to emit Atlas error event", "err", emitErr)
	}
	m.lggr.Debugw("Successfully emitted Atlas error event to Beholder", "message", mStr, "attributes", attrKVs)
}

func (m *MetaMetrics) emitAtlasUserOp(ctx context.Context, userOpHash common.Hash, tx *types.Transaction, requestSentAt, responseReceivedAt int64) {
	var nonce string
	if tx.Nonce != nil {
		nonce = strconv.FormatUint(*tx.Nonce, 10)
	}

	meta, err := tx.GetMeta()
	if err != nil {
		m.lggr.Errorw("Failed to get meta for tx", "txId", tx.ID, "userOpHash", userOpHash.Hex(), "err", err)
		return
	}

	var destAddress string
	var transactionLifecycleID string
	if meta != nil {
		if meta.FwdrDestAddress != nil {
			destAddress = meta.FwdrDestAddress.String()
		}
		if meta.TransactionLifecycleID != nil {
			transactionLifecycleID = *meta.TransactionLifecycleID
		}
	}

	msg := &pb.FastLaneAtlasUserOp{
		ChainId:                m.chainID,
		FromAddress:            tx.FromAddress.Hex(),
		ToAddress:              tx.ToAddress.Hex(),
		FeedAddress:            destAddress,
		Nonce:                  nonce,
		UserOpHash:             userOpHash.Hex(),
		TransactionLifecycleId: transactionLifecycleID,
		RequestSentAt:          requestSentAt,
		ResponseReceivedAt:     responseReceivedAt,
	}

	messageBytes, err := proto.Marshal(msg)
	if err != nil {
		m.lggr.Errorw("Failed to marshal Atlas user op event", "err", err)
		return
	}

	attrKVs := []any{
		"beholder_domain", "svr",
		"beholder_entity", "svr.v1.FastLaneAtlasUserOp",
		"beholder_data_schema", "/fastlane-atlas-user-op/versions/1",
	}

	mStr := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Format(msg)
	m.lggr.Debugw("[Beholder.emit]", "message", mStr, "attributes", attrKVs)

	if emitErr := m.emitter.Emit(ctx, messageBytes, attrKVs...); emitErr != nil {
		m.lggr.Errorw("Failed to emit Atlas user op event", "err", emitErr)
		return
	}
	m.lggr.Debugw("Successfully emitted Atlas user op event to Beholder", "message", mStr, "attributes", attrKVs)
}
