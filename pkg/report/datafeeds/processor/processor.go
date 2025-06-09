package processor

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"

	df "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/data-feeds/on-chain/registry"
	wt "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform"
)

func NewDataFeedsProcessor(metrics *df.Metrics, emitter beholder.ProtoEmitter) *datafeeds.Processor {
	return datafeeds.NewProcessor(metrics, emitter,
		GetDataFeedsSchema(),
		Decode,
	)
}

func NewCCIPDataFeedsProcessor(metrics *df.Metrics, emitter beholder.ProtoEmitter) *datafeeds.Processor {
	return datafeeds.NewProcessor(metrics, emitter,
		GetCCIPDataFeedsSchema(),
		DecodeCCIP,
	)
}

func GetDataFeedsSchema() abi.Arguments {
	return datafeeds.GetSchema("tuple(bytes32,uint32,uint224)[]", "", []abi.ArgumentMarshaling{
		{Name: "feedID", Type: "bytes32"},
		{Name: "timestamp", Type: "uint32"},
		{Name: "price", Type: "uint224"},
	})
}

func GetCCIPDataFeedsSchema() abi.Arguments {
	return datafeeds.GetSchema("tuple(bytes32,uint224,uint32)[]", "", []abi.ArgumentMarshaling{
		{Name: "feedID", Type: "bytes32"},
		{Name: "price", Type: "uint224"},
		{Name: "timestamp", Type: "uint32"},
	})
}

// Decode is made available to external users (i.e. mercury server)
func Decode(m *wt.WriteConfirmed, data []byte, schema abi.Arguments) ([]*df.FeedUpdated, error) {
	values, err := schema.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	var decoded datafeeds.Reports
	if err = schema.Copy(&decoded, values); err != nil {
		return nil, fmt.Errorf("failed to copy report values to struct: %w", err)
	}

	// Allocate space for the messages (event per updated feed)
	msgs := make([]*df.FeedUpdated, 0, len(decoded))

	// Iterate over the underlying Mercury reports
	for _, rf := range decoded {
		feedID := datafeeds.FeedID(rf.FeedID)

		// Notice: this encoding of a DF report doesn't provide a raw underlying report
		msgs = append(msgs, datafeeds.NewFeedUpdated(m, feedID, rf.Timestamp, rf.Price, []byte{}, []byte{}, true))
	}

	return msgs, nil
}

func DecodeCCIP(m *wt.WriteConfirmed, data []byte, schema abi.Arguments) ([]*df.FeedUpdated, error) {
	values, err := schema.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	var decoded datafeeds.CCIPReports
	if err = schema.Copy(&decoded, values); err != nil {
		return nil, fmt.Errorf("failed to copy report values to struct: %w", err)
	}

	// Allocate space for the messages (event per updated feed)
	msgs := make([]*df.FeedUpdated, 0, len(decoded))

	// Iterate over the underlying Mercury reports
	for _, rf := range decoded {
		feedID := datafeeds.FeedID(rf.FeedID)

		// Notice: this encoding of a DF report doesn't provide a raw underlying report
		msgs = append(msgs, datafeeds.NewFeedUpdated(m, feedID, rf.Timestamp, rf.Price, []byte{}, []byte{}, true))
	}

	return msgs, nil
}
