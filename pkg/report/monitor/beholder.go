//nolint:revive,staticcheck // disable revive, staticcheck
package monitor

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

const (
	// Helper keys to avoid duplicating attributes
	CtxKeySkipAppendAttrs = "skip_append_attrs"
)

// BeholderClient is a Beholder client extension with a custom ProtoEmitter
type BeholderClient struct {
	*beholder.Client
	ProtoEmitter ProtoEmitter
}

// ProtoEmitter is an interface for emitting protobuf messages
// TODO: this should be moved to chainlink-common
type ProtoEmitter interface {
	// Sends message with bytes and attributes to OTel Collector
	Emit(ctx context.Context, m proto.Message, attrKVs ...any) error
	EmitWithLog(ctx context.Context, m proto.Message, attrKVs ...any) error
}

// ProtoProcessor is an interface for processing emitted protobuf messages
type ProtoProcessor interface {
	Process(ctx context.Context, m proto.Message, attrKVs ...any) error
}

func NewProtoEmitter(lggr logger.Logger, client *beholder.Client, schemaBasePath string) ProtoEmitter {
	return &protoEmitter{lggr, client, schemaBasePath}
}

// protoEmitter is a ProtoEmitter implementation
var _ ProtoEmitter = (*protoEmitter)(nil)

type protoEmitter struct {
	lggr           logger.Logger
	client         *beholder.Client
	schemaBasePath string
}

func (e *protoEmitter) Emit(ctx context.Context, m proto.Message, attrKVs ...any) error {
	payload, err := proto.Marshal(m)
	if err != nil {
		// Notice: we log here because emit errors are usually not critical and swallowed by the caller
		e.lggr.Errorw("[Beholder] Failed to marshal", "err", err)
		return err
	}

	// Skip appending attributes if the context says it's already done that
	if skip, ok := ctx.Value(CtxKeySkipAppendAttrs).(bool); !ok || !skip {
		attrKVs = e.appendAttrsRequired(ctx, m, attrKVs)
	}

	// Emit the message with attributes
	err = e.client.Emitter.Emit(ctx, payload, attrKVs...)
	if err != nil {
		// Notice: we log here because emit errors are usually not critical and swallowed by the caller
		e.lggr.Errorw("[Beholder] Failed to client.Emitter.Emit", "err", err)
		return err
	}

	return nil
}

// EmitWithLog emits a protobuf message with attributes and logs the emitted message
func (e *protoEmitter) EmitWithLog(ctx context.Context, m proto.Message, attrKVs ...any) error {
	attrKVs = e.appendAttrsRequired(ctx, m, attrKVs)
	// attach a bool switch to ctx to avoid duplicating common attrs
	ctx = context.WithValue(ctx, CtxKeySkipAppendAttrs, true)

	// Marshal the message as JSON and log before emitting
	// https://protobuf.dev/programming-guides/json/
	mStr := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Format(m)
	e.lggr.Infow("[Beholder.emit]", "message", mStr, "attributes", attrKVs)

	return e.Emit(ctx, m, attrKVs...)
}

// appendAttrsRequired appends required attributes to the attribute key-value list
func (e *protoEmitter) appendAttrsRequired(ctx context.Context, m proto.Message, attrKVs []any) []any {
	attrKVs = appendRequiredAttrDataSchema(m, attrKVs, e.schemaBasePath)
	attrKVs = appendRequiredAttrEntity(m, attrKVs)
	attrKVs = appendRequiredAttrDomain(m, attrKVs)
	return attrKVs
}
