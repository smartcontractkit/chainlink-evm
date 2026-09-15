package toml

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	commonassets "github.com/smartcontractkit/chainlink-common/pkg/assets"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

const (
	testKeyA = "0x2a3e23c6f242F5345320814aC8a1b4E58707D292"
	testKeyB = "0x538aAaB4ea120b2bC2fe5D296852D948F07D849e"
)

// chainA returns a Chain with every settable field non-nil. KeySpecific is left empty so that it can be
// exercised separately - unlike the other fields it merges rather than overwrites.
func chainA() Chain {
	return Chain{
		AutoCreateKey:                new(true),
		BlockBackfillDepth:           new(uint32(1)),
		BlockBackfillSkip:            new(true),
		ChainType:                    chaintype.NewConfig("optimismBedrock"),
		FinalityDepth:                new(uint32(2)),
		SafeDepth:                    new(uint32(3)),
		FinalityTagEnabled:           new(true),
		SafeTagSupported:             new(true),
		FlagsContractAddress:         new(types.MustEIP55Address(testKeyA)),
		LinkContractAddress:          new(types.MustEIP55Address(testKeyA)),
		LogBackfillBatchSize:         new(uint32(4)),
		LogPollInterval:              commonconfig.MustNewDuration(time.Second),
		LogPollerSkipEmptyBlocks:     new(true),
		LogKeepBlocksDepth:           new(uint32(5)),
		LogPrunePageSize:             new(uint32(6)),
		BackupLogPollerBlockDelay:    new(uint64(7)),
		MinIncomingConfirmations:     new(uint32(8)),
		MinContractPayment:           commonassets.NewLinkFromJuels(9),
		NonceAutoSync:                new(true),
		NoNewHeadsThreshold:          commonconfig.MustNewDuration(2 * time.Second),
		OperatorFactoryAddress:       new(types.MustEIP55Address(testKeyA)),
		LogBroadcasterEnabled:        new(true),
		RPCDefaultBatchSize:          new(uint32(10)),
		RPCBlockQueryDelay:           new(uint16(11)),
		FinalizedBlockOffset:         new(uint32(12)),
		NoNewFinalizedHeadsThreshold: commonconfig.MustNewDuration(3 * time.Second),

		Transactions:   Transactions{MaxInFlight: new(uint32(13))},
		BalanceMonitor: BalanceMonitor{Enabled: new(true)},
		GasEstimator:   GasEstimator{Mode: new("FixedPrice")},
		HeadTracker:    HeadTracker{HistoryDepth: new(uint32(14))},
		NodePool:       NodePool{SelectionMode: new("HighestHead")},
		OCR:            OCR{ContractConfirmations: new(uint16(15))},
		Workflow:       Workflow{GasLimitDefault: new(uint64(16))},
	}
}

// chainB mirrors chainA with distinct values, so that an override is observable for every field.
func chainB() Chain {
	return Chain{
		AutoCreateKey:                new(false),
		BlockBackfillDepth:           new(uint32(101)),
		BlockBackfillSkip:            new(false),
		ChainType:                    chaintype.NewConfig("arbitrum"),
		FinalityDepth:                new(uint32(102)),
		SafeDepth:                    new(uint32(103)),
		FinalityTagEnabled:           new(false),
		SafeTagSupported:             new(false),
		FlagsContractAddress:         new(types.MustEIP55Address(testKeyB)),
		LinkContractAddress:          new(types.MustEIP55Address(testKeyB)),
		LogBackfillBatchSize:         new(uint32(104)),
		LogPollInterval:              commonconfig.MustNewDuration(time.Minute),
		LogPollerSkipEmptyBlocks:     new(false),
		LogKeepBlocksDepth:           new(uint32(105)),
		LogPrunePageSize:             new(uint32(106)),
		BackupLogPollerBlockDelay:    new(uint64(107)),
		MinIncomingConfirmations:     new(uint32(108)),
		MinContractPayment:           commonassets.NewLinkFromJuels(109),
		NonceAutoSync:                new(false),
		NoNewHeadsThreshold:          commonconfig.MustNewDuration(2 * time.Minute),
		OperatorFactoryAddress:       new(types.MustEIP55Address(testKeyB)),
		LogBroadcasterEnabled:        new(false),
		RPCDefaultBatchSize:          new(uint32(110)),
		RPCBlockQueryDelay:           new(uint16(111)),
		FinalizedBlockOffset:         new(uint32(112)),
		NoNewFinalizedHeadsThreshold: commonconfig.MustNewDuration(3 * time.Minute),

		Transactions:   Transactions{MaxInFlight: new(uint32(113))},
		BalanceMonitor: BalanceMonitor{Enabled: new(false)},
		GasEstimator:   GasEstimator{Mode: new("BlockHistory")},
		HeadTracker:    HeadTracker{HistoryDepth: new(uint32(114))},
		NodePool:       NodePool{SelectionMode: new("RoundRobin")},
		OCR:            OCR{ContractConfirmations: new(uint16(115))},
		Workflow:       Workflow{GasLimitDefault: new(uint64(116))},
	}
}

func TestChain_SetFrom(t *testing.T) {
	t.Run("sets every field on an empty target", func(t *testing.T) {
		var c Chain
		f := chainA()

		c.SetFrom(&f)

		require.Equal(t, chainA(), c)
	})

	t.Run("nil source fields leave the target untouched", func(t *testing.T) {
		c := chainA()

		c.SetFrom(&Chain{})

		require.Equal(t, chainA(), c)
	})

	t.Run("non-nil source fields override the target", func(t *testing.T) {
		c := chainA()
		f := chainB()

		c.SetFrom(&f)

		require.Equal(t, chainB(), c)
	})

	t.Run("nested configs are merged field by field", func(t *testing.T) {
		c := Chain{
			Transactions:   Transactions{MaxInFlight: new(uint32(1)), MaxQueued: new(uint32(2))},
			BalanceMonitor: BalanceMonitor{Enabled: new(true)},
			GasEstimator:   GasEstimator{Mode: new("FixedPrice"), PriceMax: assets.NewWeiI(3)},
			HeadTracker:    HeadTracker{HistoryDepth: new(uint32(4)), MaxBufferSize: new(uint32(5))},
			NodePool:       NodePool{SelectionMode: new("HighestHead"), SyncThreshold: new(uint32(6))},
			OCR:            OCR{ContractConfirmations: new(uint16(7))},
			Workflow:       Workflow{GasLimitDefault: new(uint64(8)), TxAcceptanceState: new(commontypes.Unconfirmed)},
		}
		f := Chain{
			Transactions:   Transactions{MaxQueued: new(uint32(20))},
			BalanceMonitor: BalanceMonitor{Enabled: new(false)},
			GasEstimator:   GasEstimator{PriceMax: assets.NewWeiI(30)},
			HeadTracker:    HeadTracker{MaxBufferSize: new(uint32(50))},
			NodePool:       NodePool{SyncThreshold: new(uint32(60))},
			OCR:            OCR{DatabaseTimeout: commonconfig.MustNewDuration(time.Second)},
			Workflow:       Workflow{TxAcceptanceState: new(commontypes.Finalized)},
		}

		c.SetFrom(&f)

		// overridden
		require.Equal(t, uint32(20), *c.Transactions.MaxQueued)
		require.False(t, *c.BalanceMonitor.Enabled)
		require.Equal(t, assets.NewWeiI(30), c.GasEstimator.PriceMax)
		require.Equal(t, uint32(50), *c.HeadTracker.MaxBufferSize)
		require.Equal(t, uint32(60), *c.NodePool.SyncThreshold)
		require.Equal(t, commonconfig.MustNewDuration(time.Second), c.OCR.DatabaseTimeout)
		require.Equal(t, commontypes.Finalized, *c.Workflow.TxAcceptanceState)

		// untouched
		require.Equal(t, uint32(1), *c.Transactions.MaxInFlight)
		require.Equal(t, "FixedPrice", *c.GasEstimator.Mode)
		require.Equal(t, uint32(4), *c.HeadTracker.HistoryDepth)
		require.Equal(t, "HighestHead", *c.NodePool.SelectionMode)
		require.Equal(t, uint16(7), *c.OCR.ContractConfirmations)
		require.Equal(t, uint64(8), *c.Workflow.GasLimitDefault)
	})

	t.Run("KeySpecific", func(t *testing.T) {
		keySpecific := func(key *types.EIP55Address, priceMax *assets.Wei) KeySpecific {
			return KeySpecific{Key: key, GasEstimator: KeySpecificGasEstimator{PriceMax: priceMax}}
		}

		t.Run("nil source keeps the target list", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(1)),
			}}

			c.SetFrom(&Chain{})

			require.Len(t, c.KeySpecific, 1)
			require.Equal(t, assets.NewWeiI(1), c.KeySpecific[0].GasEstimator.PriceMax)
		})

		t.Run("appends to an empty target list", func(t *testing.T) {
			var c Chain
			f := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(1)),
			}}

			c.SetFrom(&f)

			require.Equal(t, f.KeySpecific, c.KeySpecific)
		})

		t.Run("appends an entry whose key is not present", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(1)),
			}}
			f := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyB)), assets.NewWeiI(2)),
			}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 2)
			require.Equal(t, assets.NewWeiI(1), c.KeySpecific[0].GasEstimator.PriceMax)
			require.Equal(t, assets.NewWeiI(2), c.KeySpecific[1].GasEstimator.PriceMax)
		})

		t.Run("merges entries with equal keys held by distinct pointers", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(1)),
			}}
			f := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(2)),
			}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 1)
			require.Equal(t, assets.NewWeiI(2), c.KeySpecific[0].GasEstimator.PriceMax)
		})

		t.Run("merges entries that share a key pointer", func(t *testing.T) {
			key := new(types.MustEIP55Address(testKeyA))
			c := Chain{KeySpecific: KeySpecificConfig{keySpecific(key, assets.NewWeiI(1))}}
			f := Chain{KeySpecific: KeySpecificConfig{keySpecific(key, assets.NewWeiI(2))}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 1)
			require.Equal(t, assets.NewWeiI(2), c.KeySpecific[0].GasEstimator.PriceMax)
		})

		t.Run("merges entries that both have a nil key", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{keySpecific(nil, assets.NewWeiI(1))}}
			f := Chain{KeySpecific: KeySpecificConfig{keySpecific(nil, assets.NewWeiI(2))}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 1)
			require.Equal(t, assets.NewWeiI(2), c.KeySpecific[0].GasEstimator.PriceMax)
		})

		t.Run("appends when only the source key is nil", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(1)),
			}}
			f := Chain{KeySpecific: KeySpecificConfig{keySpecific(nil, assets.NewWeiI(2))}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 2)
			require.Equal(t, assets.NewWeiI(1), c.KeySpecific[0].GasEstimator.PriceMax)
			require.Nil(t, c.KeySpecific[1].Key)
		})

		t.Run("appends when only the target key is nil", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{keySpecific(nil, assets.NewWeiI(1))}}
			f := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(2)),
			}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 2)
			require.Nil(t, c.KeySpecific[0].Key)
			require.Equal(t, assets.NewWeiI(1), c.KeySpecific[0].GasEstimator.PriceMax)
			require.Equal(t, assets.NewWeiI(2), c.KeySpecific[1].GasEstimator.PriceMax)
		})

		t.Run("a nil source PriceMax leaves the merged entry untouched", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(1)),
			}}
			f := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), nil),
			}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 1)
			require.Equal(t, assets.NewWeiI(1), c.KeySpecific[0].GasEstimator.PriceMax)
		})

		t.Run("merges and appends across multiple source entries", func(t *testing.T) {
			c := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(1)),
			}}
			f := Chain{KeySpecific: KeySpecificConfig{
				keySpecific(new(types.MustEIP55Address(testKeyA)), assets.NewWeiI(2)),
				keySpecific(new(types.MustEIP55Address(testKeyB)), assets.NewWeiI(3)),
			}}

			c.SetFrom(&f)

			require.Len(t, c.KeySpecific, 2)
			require.Equal(t, types.MustEIP55Address(testKeyA), *c.KeySpecific[0].Key)
			require.Equal(t, assets.NewWeiI(2), c.KeySpecific[0].GasEstimator.PriceMax)
			require.Equal(t, types.MustEIP55Address(testKeyB), *c.KeySpecific[1].Key)
			require.Equal(t, assets.NewWeiI(3), c.KeySpecific[1].GasEstimator.PriceMax)
		})
	})
}
