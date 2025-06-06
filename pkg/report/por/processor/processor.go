package por

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/data-feeds/on-chain/registry"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"

	wt "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform"
)

func NewPORProcessor(metrics *registry.Metrics, emitter beholder.ProtoEmitter) *registry.Processor {
	return registry.NewProcessor(metrics, emitter,
		GetPORSchema(),
		Decode,
	)
}

func GetPORSchema() abi.Arguments {
	return registry.GetSchema("tuple(bytes32,uint32,bytes)[]", "", []abi.ArgumentMarshaling{
		{Name: "dataId", Type: "bytes32"},
		{Name: "timestamp", Type: "uint32"},
		{Name: "bundle", Type: "bytes"},
	})
}

// Decode is made available to external users (i.e. mercury server)
func Decode(m *wt.WriteConfirmed, data []byte, schema abi.Arguments) ([]*registry.FeedUpdated, error) {
	values, err := schema.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	var decoded registry.PORReports
	if err = schema.Copy(&decoded, values); err != nil {
		return nil, fmt.Errorf("failed to copy report values to struct: %w", err)
	}

	// Allocate space for the messages (event per updated feed)
	msgs := make([]*registry.FeedUpdated, 0, len(decoded))

	// Iterate over the underlying reports
	for _, rf := range decoded {
		feedID := datafeeds.FeedID(rf.DataID)

		// Notice: uses a placeholder for the benchmark price
		msgs = append(msgs, registry.NewFeedUpdated(m, feedID, rf.Timestamp, big.NewInt(0), rf.Bundle, []byte{}, true))
	}

	return msgs, nil
}
