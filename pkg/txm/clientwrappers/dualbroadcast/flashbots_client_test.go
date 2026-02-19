package dualbroadcast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseURLParams(t *testing.T) {
	tests := []struct {
		name           string
		params         string
		wantPrivacy    Privacy
		wantRefund     RefundConfig
		wantErr        bool
		wantErrContain string
	}{
		{
			name:        "empty params",
			params:      "",
			wantPrivacy: Privacy{},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "auctionTimeout",
			params:      "auctionTimeout=60",
			wantPrivacy: Privacy{AuctionTimeout: 60},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "auctionTimeout invalid ignored",
			params:      "auctionTimeout=notanint",
			wantPrivacy: Privacy{},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "single builder",
			params:      "builder=test_builder",
			wantPrivacy: Privacy{Builders: []string{"test_builder"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "multiple builders",
			params:      "builder=test_builder_1&builder=test_builder_2",
			wantPrivacy: Privacy{Builders: []string{"test_builder_1", "test_builder_2"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "single hint",
			params:      "hint=calldata",
			wantPrivacy: Privacy{Hints: []string{"calldata"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "multiple hints",
			params:      "hint=calldata&hint=hash",
			wantPrivacy: Privacy{Hints: []string{"calldata", "hash"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "refund valid",
			params:      "refund=0xRefundAddr:50",
			wantPrivacy: Privacy{WantRefund: 50},
			wantRefund:  RefundConfig{Address: "0xRefundAddr", Percent: 100},
		},
		{
			name:           "refund invalid percent",
			params:         "refund=0xRefundAddr:bad",
			wantErr:        true,
			wantErrContain: "unable to parse percentage",
		},
		{
			name:        "refund single part ignored",
			params:      "refund=0xRefundAddr",
			wantPrivacy: Privacy{},
			wantRefund:  RefundConfig{},
		},
		{
			name:           "invalid query",
			params:         "%",
			wantErr:        true,
			wantErrContain: "unable to parse params",
		},
		{
			name:        "combined params",
			params:      "auctionTimeout=120&builder=test_builder_1&builder=test_builder_2&hint=h1&refund=0xR:75",
			wantPrivacy: Privacy{
				AuctionTimeout: 120,
				Builders:       []string{"test_builder_1", "test_builder_2"},
				Hints:          []string{"h1"},
				WantRefund:     75,
			},
			wantRefund: RefundConfig{Address: "0xR", Percent: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privacy, refund, err := parseURLParams(tt.params)
			if tt.wantErr {
				require.Error(t, err)
				if tt.wantErrContain != "" {
					assert.Contains(t, err.Error(), tt.wantErrContain)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantPrivacy, privacy)
			assert.Equal(t, tt.wantRefund, refund)
		})
	}
}
