package aptos

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/pb/data-feeds/on-chain/registry"
	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/platform"

	mercury_vX "github.com/smartcontractkit/chainlink-evm/pkg/report/mercury/common"
	mercury_v3 "github.com/smartcontractkit/chainlink-evm/pkg/report/mercury/v3"
	mercury_v4 "github.com/smartcontractkit/chainlink-evm/pkg/report/mercury/v4"
)

func DecodeAsFeedUpdated(m *wt_msg.WriteConfirmed) ([]*registry.FeedUpdated, error) {
	// Decode the confirmed report (WT -> DF contract event)
	r, err := platform.Decode(m.Report)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	// Decode the underlying Data Feeds reports
	reports, err := Decode(r.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Data Feeds report: %w", err)
	}

	// Allocate space for the messages (event per updated feed)
	msgs := make([]*registry.FeedUpdated, 0, len(*reports))

	// Iterate over the underlying Mercury reports
	for _, rf := range *reports {
		// Decode the common Mercury report and get report type
		rmCommon, err := mercury_vX.Decode(rf.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode Mercury report: %w", err)
		}

		// Parse the report type from the common header
		t := mercury_vX.GetReportType(rmCommon.FeedID)
		feedID := datafeeds.FeedID(rf.FeedID)

		switch t {
		case uint16(3):
			rm, err := mercury_v3.Decode(rf.Data)
			if err != nil {
				return nil, fmt.Errorf("failed to decode Mercury v%d report: %w", t, err)
			}
			// For Mercury v3, include TxSender and TxReceiver
			msgs = append(msgs, registry.NewFeedUpdated(m, feedID, rm.ObservationsTimestamp, rm.BenchmarkPrice, rf.Data, true))
		case uint16(4):
			rm, err := mercury_v4.Decode(rf.Data)
			if err != nil {
				return nil, fmt.Errorf("failed to decode Mercury v%d report: %w", t, err)
			}
			// For Mercury v4, skip TxSender and TxReceiver (if not applicable)
			msgs = append(msgs, registry.NewFeedUpdated(m, feedID, rm.ObservationsTimestamp, rm.BenchmarkPrice, rf.Data, false))
		default:
			return nil, fmt.Errorf("unsupported Mercury report type: %d", t)
		}
	}

	return msgs, nil
}
