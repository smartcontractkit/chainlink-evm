package evm

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/pb/data-feeds/on-chain/registry"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/platform"
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
		feedID := datafeeds.FeedID(rf.FeedID)

		// TODO: unsure if r.Data is correct for Report
		msgs = append(msgs, registry.NewFeedUpdated(m, feedID, rf.Timestamp, rf.Price, r.Data, false))
	}

	return msgs, nil
}
