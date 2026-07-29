package toml

import (
	_ "embed"
	"fmt"
	"math"
	stdbig "math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kylelemons/godebug/diff"
	"github.com/pelletier/go-toml/v2"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commonassets "github.com/smartcontractkit/chainlink-common/pkg/assets"
	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/config/configdoc"
	"github.com/smartcontractkit/chainlink-common/pkg/config/configtest"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-framework/multinode"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

func TestEVMConfig_ValidateConfig(t *testing.T) {
	name := "fake"
	for _, id := range DefaultIDs {
		t.Run(fmt.Sprintf("chainID-%s", id), func(t *testing.T) {
			evmCfg := &EVMConfig{
				ChainID: id,
				Chain:   Defaults(id),
				Nodes: EVMNodes{{
					Name:    &name,
					WSURL:   config.MustParseURL("wss://foo.test/ws"),
					HTTPURL: config.MustParseURL("http://foo.test"),
				}},
			}

			assert.NoError(t, config.Validate(evmCfg))
		})
	}
}

func TestEVMConfig_ValidateInvalidConfig(t *testing.T) {
	testCases := []struct {
		Name          string
		MakeInvalid   func(cfg *Chain)
		ExpectedError string
	}{
		{
			Name: "LogBackfillBatchSize must be > 0",
			MakeInvalid: func(cfg *Chain) {
				cfg.LogBackfillBatchSize = new(uint32(0))
			},
			ExpectedError: "LogBackfillBatchSize: invalid value (0): must be greater than 0",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var cfg Chain
			cfg.SetFrom(&fallback)
			tc.MakeInvalid(&cfg)
			err := cfg.ValidateConfig()
			require.Error(t, err)
			require.ErrorContains(t, err, tc.ExpectedError)
		})
	}
}

func TestEVMConfig_ValidateConfig_RPCDefaultBatchSize(t *testing.T) {
	name := "fake"
	id := DefaultIDs[0]
	evmCfg := &EVMConfig{
		ChainID: id,
		Chain:   Defaults(id),
		Nodes: EVMNodes{{
			Name:    &name,
			WSURL:   config.MustParseURL("wss://foo.test/ws"),
			HTTPURL: config.MustParseURL("http://foo.test"),
		}},
	}
	evmCfg.RPCDefaultBatchSize = new(uint32(0))
	err := config.Validate(evmCfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "RPCDefaultBatchSize")
}

func TestDefaults_fieldsNotNil(t *testing.T) {
	unknown := Defaults(nil)

	// exceptional nilable fields
	unknown.ChainType = chaintype.NewConfig("arbitrum")
	unknown.FlagsContractAddress = asEIP55Address(t, "0x1234567890abcdefaC8a1b4E58707D29258707D2")
	unknown.LinkContractAddress = asEIP55Address(t, "0xabcdef1234567890aC8a1b4E58707D29258707D2")
	unknown.OperatorFactoryAddress = asEIP55Address(t, "0xababab12341234aC8a1b4E58707D29258707D292")
	addr, err := types.NewEIP55Address("0x2a3e23c6f242F5345320814aC8a1b4E58707D292")
	require.NoError(t, err)
	unknown.Workflow.FromAddress = &addr
	unknown.Workflow.ForwarderAddress = &addr
	unknown.Workflow.GasLimitDefault = new(uint64(400000))
	unknown.Transactions.TransactionManagerV2.BlockTime = new(config.Duration)
	unknown.Transactions.TransactionManagerV2.CustomURL = new(config.URL)
	unknown.Transactions.TransactionManagerV2.CustomURLs = []*config.URL{new(config.URL)}
	unknown.Transactions.TransactionManagerV2.DualBroadcast = new(false)
	unknown.Transactions.TransactionManagerV2.ReadRequestsToMultipleNodes = new(false)
	unknown.Transactions.TransactionManagerV2.Bundles = new(false)
	unknown.Transactions.TransactionManagerV2.FastlaneAuctionRequestTimeout = new(config.Duration)
	unknown.Transactions.TransactionManagerV2.FeeBoost = new(false)
	unknown.Transactions.AutoPurge.Threshold = new(uint32(0))
	unknown.Transactions.AutoPurge.MinAttempts = new(uint32(0))
	unknown.Transactions.AutoPurge.DetectionApiUrl = new(config.URL)
	unknown.GasEstimator.BlockHistory.EIP1559FeeCapBufferBlocks = new(uint16(10))
	unknown.GasEstimator.SenderAddress = asEIP55Address(t, "0xae4E781a6218A8031764928E88d457937A954fC3")
	oracleType := DAOracleOPStack
	unknown.GasEstimator.DAOracle.OracleType = &oracleType
	unknown.GasEstimator.DAOracle.OracleAddress = new(types.EIP55Address)
	unknown.GasEstimator.DAOracle.CustomGasPriceCalldata = new(string)
	unknown.GasEstimator.LimitJobType = GasLimitJobType{
		OCR:    new(uint32(7)),
		OCR2:   new(uint32(13)),
		DR:     new(uint32(25)),
		VRF:    new(uint32(37)),
		FM:     new(uint32(42)),
		Keeper: new(uint32(51)),
	}
	unknown.GasEstimator.BumpTxDepth = new(uint32(15))
	unknown.NodePool.Errors = ClientErrors{
		NonceTooLow:                       new("too-low"),
		NonceTooHigh:                      new("too-high"),
		ReplacementTransactionUnderpriced: new("under"),
		LimitReached:                      new("limit"),
		TransactionAlreadyInMempool:       new("already"),
		TerminallyUnderpriced:             new("terminal"),
		InsufficientEth:                   new("insufficient"),
		TxFeeExceedsCap:                   new("exceeds"),
		L2FeeTooLow:                       new("low-fee"),
		L2FeeTooHigh:                      new("high-fee"),
		L2Full:                            new("full"),
		TransactionAlreadyMined:           new("mined"),
		Fatal:                             new("fatal"),
		ServiceUnavailable:                new("unavailable"),
		TooManyResults:                    new("too-many"),
		MissingBlocks:                     new("missing"),
		FinalizedStateUnavailable:         new("finalized-unavailable"),
	}

	configtest.AssertFieldsNotNil(t, unknown)
}

func TestDocs(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		configtest.AssertDocsTOMLComplete[EVMConfig](t, docsTOML)
	})

	t.Run("aligned", func(t *testing.T) {
		var docDefaults EVMConfig
		require.NoError(t, configdoc.DefaultsOnly(strings.NewReader(docsTOML), &docDefaults, config.DecodeTOML))

		require.Equal(t, chaintype.ChainType(""), docDefaults.ChainType.ChainType())
		docDefaults.ChainType = nil

		// EVM.GasEstimator.BumpTxDepth doesn't have a constant default - it is derived from another field
		require.Zero(t, *docDefaults.GasEstimator.BumpTxDepth)
		docDefaults.GasEstimator.BumpTxDepth = nil

		// per-job limits are nilable
		require.Zero(t, *docDefaults.GasEstimator.LimitJobType.OCR)
		require.Zero(t, *docDefaults.GasEstimator.LimitJobType.OCR2)
		require.Zero(t, *docDefaults.GasEstimator.LimitJobType.DR)
		require.Zero(t, *docDefaults.GasEstimator.LimitJobType.Keeper)
		require.Zero(t, *docDefaults.GasEstimator.LimitJobType.VRF)
		require.Zero(t, *docDefaults.GasEstimator.LimitJobType.FM)
		docDefaults.GasEstimator.LimitJobType = GasLimitJobType{}

		// EIP1559FeeCapBufferBlocks doesn't have a constant default - it is derived from another field
		require.Zero(t, *docDefaults.GasEstimator.BlockHistory.EIP1559FeeCapBufferBlocks)
		docDefaults.GasEstimator.BlockHistory.EIP1559FeeCapBufferBlocks = nil

		// addresses w/o global values
		require.Zero(t, *docDefaults.FlagsContractAddress)
		require.Zero(t, *docDefaults.LinkContractAddress)
		require.Zero(t, *docDefaults.OperatorFactoryAddress)
		docDefaults.FlagsContractAddress = nil
		docDefaults.LinkContractAddress = nil
		docDefaults.OperatorFactoryAddress = nil
		require.Empty(t, docDefaults.Workflow.FromAddress)
		require.Empty(t, docDefaults.Workflow.ForwarderAddress)
		gasLimitDefault := uint64(400_000)
		require.Equal(t, &gasLimitDefault, docDefaults.Workflow.GasLimitDefault)

		docDefaults.Workflow.FromAddress = nil
		docDefaults.Workflow.ForwarderAddress = nil
		docDefaults.Workflow.GasLimitDefault = &gasLimitDefault
		finalizedStateUnavailable := docDefaults.NodePool.Errors.FinalizedStateUnavailable
		docDefaults.NodePool.Errors = ClientErrors{}
		docDefaults.NodePool.Errors.FinalizedStateUnavailable = finalizedStateUnavailable

		// Transactions.AutoPurge configs are only set if the feature is enabled
		docDefaults.Transactions.AutoPurge.DetectionApiUrl = nil
		docDefaults.Transactions.AutoPurge.Threshold = nil
		docDefaults.Transactions.AutoPurge.MinAttempts = nil

		// TransactionManagerV2 configs are only set if the feature is enabled
		docDefaults.Transactions.TransactionManagerV2.BlockTime = nil
		docDefaults.Transactions.TransactionManagerV2.CustomURL = nil
		docDefaults.Transactions.TransactionManagerV2.CustomURLs = nil
		docDefaults.Transactions.TransactionManagerV2.DualBroadcast = nil
		docDefaults.Transactions.TransactionManagerV2.ReadRequestsToMultipleNodes = nil
		docDefaults.Transactions.TransactionManagerV2.Bundles = nil
		docDefaults.Transactions.TransactionManagerV2.FastlaneAuctionRequestTimeout = nil
		docDefaults.Transactions.TransactionManagerV2.FeeBoost = nil

		// Fallback DA oracle is not set
		docDefaults.GasEstimator.DAOracle = DAOracle{}

		// GasEstimator SendAddress is only set if EstimateLimit is enabled
		docDefaults.GasEstimator.SenderAddress = nil

		// HistoricalBalanceCheckAddress is documented as # Example; fallback.toml supplies the runtime default.
		docDefaults.NodePool.HistoricalBalanceCheckAddress = nil

		fallbackDefaults := Defaults(nil)
		fallbackDefaults.NodePool.HistoricalBalanceCheckAddress = nil
		assertTOML(t, fallbackDefaults, docDefaults.Chain)
	})
}

//go:embed testdata/config-full.toml
var fullTOML string

var fullConfig = EVMConfig{
	ChainID: sqlutil.NewI(42),
	Enabled: new(false),
	Chain: Chain{
		AutoCreateKey: new(false),
		BalanceMonitor: BalanceMonitor{
			Enabled: new(true),
		},
		BlockBackfillDepth:   new(uint32(100)),
		BlockBackfillSkip:    new(true),
		ChainType:            chaintype.NewConfig("Optimism"),
		FinalityDepth:        new(uint32(42)),
		SafeDepth:            new(uint32(10)),
		FinalityTagEnabled:   new(true),
		SafeTagSupported:     new(true),
		FlagsContractAddress: new(types.MustEIP55Address("0xae4E781a6218A8031764928E88d457937A954fC3")),
		FinalizedBlockOffset: new(uint32(16)),

		GasEstimator: GasEstimator{
			Mode:               new("SuggestedPrice"),
			EIP1559DynamicFees: new(true),
			BumpPercent:        new(uint16(10)),
			BumpThreshold:      new(uint32(6)),
			BumpTxDepth:        new(uint32(6)),
			BumpMin:            assets.NewWeiI(100),
			FeeCapDefault:      assets.NewWeiI(math.MaxInt64),
			LimitDefault:       new(uint64(12)),
			LimitMax:           new(uint64(17)),
			LimitMultiplier:    new(decimal.RequireFromString("1.234")),
			LimitTransfer:      new(uint64(100)),
			EstimateLimit:      new(false),
			SenderAddress:      new(types.MustEIP55Address("0xae4E781a6218A8031764928E88d457937A954fC3")),
			TipCapDefault:      assets.NewWeiI(2),
			TipCapMin:          assets.NewWeiI(1),
			PriceDefault:       assets.NewWeiI(math.MaxInt64),
			PriceMax:           assets.NewWei(new(stdbig.Int).SetBytes([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})),
			PriceMin:           assets.NewWeiI(13),

			DAOracle: DAOracle{
				OracleType:             new(DAOracleOPStack),
				OracleAddress:          new(types.MustEIP55Address("0xae4E781a6218A8031764928E88d457937A954fC3")),
				CustomGasPriceCalldata: new("0x1234asdf"),
			},

			LimitJobType: GasLimitJobType{
				OCR:    new(uint32(1001)),
				DR:     new(uint32(1002)),
				VRF:    new(uint32(1003)),
				FM:     new(uint32(1004)),
				Keeper: new(uint32(1005)),
				OCR2:   new(uint32(1006)),
			},

			BlockHistory: BlockHistoryEstimator{
				BatchSize:                 new(uint32(17)),
				BlockHistorySize:          new(uint16(12)),
				CheckInclusionBlocks:      new(uint16(18)),
				CheckInclusionPercentile:  new(uint16(19)),
				EIP1559FeeCapBufferBlocks: new(uint16(13)),
				TransactionPercentile:     new(uint16(15)),
			},
			FeeHistory: FeeHistoryEstimator{
				CacheTimeout: config.MustNewDuration(time.Second),
			},
		},

		KeySpecific: []KeySpecific{
			{
				Key: new(types.MustEIP55Address("0x2a3e23c6f242F5345320814aC8a1b4E58707D292")),
				GasEstimator: KeySpecificGasEstimator{
					PriceMax: assets.NewWei(new(stdbig.Int).SetBytes([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})),
				},
			},
		},

		LinkContractAddress:          new(types.MustEIP55Address("0x538aAaB4ea120b2bC2fe5D296852D948F07D849e")),
		LogBackfillBatchSize:         new(uint32(17)),
		LogPollInterval:              config.MustNewDuration(time.Minute),
		LogPollerSkipEmptyBlocks:     new(false),
		LogKeepBlocksDepth:           new(uint32(100000)),
		LogPrunePageSize:             new(uint32(0)),
		BackupLogPollerBlockDelay:    new(uint64(532)),
		MinContractPayment:           commonassets.NewLinkFromJuels(math.MaxInt64),
		MinIncomingConfirmations:     new(uint32(13)),
		NonceAutoSync:                new(true),
		NoNewHeadsThreshold:          config.MustNewDuration(time.Minute),
		OperatorFactoryAddress:       new(types.MustEIP55Address("0xa5B85635Be42F21f94F28034B7DA440EeFF0F418")),
		LogBroadcasterEnabled:        new(true),
		RPCDefaultBatchSize:          new(uint32(17)),
		RPCBlockQueryDelay:           new(uint16(10)),
		NoNewFinalizedHeadsThreshold: config.MustNewDuration(time.Hour),

		Transactions: Transactions{
			Enabled:              new(true),
			MaxInFlight:          new(uint32(19)),
			MaxQueued:            new(uint32(99)),
			ReaperInterval:       config.MustNewDuration(time.Minute),
			ReaperThreshold:      config.MustNewDuration(time.Minute),
			ResendAfterThreshold: config.MustNewDuration(time.Hour),
			ConfirmationTimeout:  config.MustNewDuration(time.Minute),
			ForwardersEnabled:    new(true),
			AutoPurge: AutoPurgeConfig{
				Enabled:         new(false),
				Threshold:       new(uint32(42)),
				MinAttempts:     new(uint32(13)),
				DetectionApiUrl: config.MustParseURL("http://example.net"),
			},
			TransactionManagerV2: TransactionManagerV2Config{
				Enabled:                       new(false),
				DualBroadcast:                 new(true),
				ReadRequestsToMultipleNodes:   new(false),
				Bundles:                       new(false),
				BlockTime:                     config.MustNewDuration(42 * time.Second),
				CustomURL:                     config.MustParseURL("http://txs.org"),
				CustomURLs:                    []*config.URL{config.MustParseURL("http://txs.org"), config.MustParseURL("http://txs.org/secondary")},
				FastlaneAuctionRequestTimeout: config.MustNewDuration(15 * time.Second),
				FeeBoost:                      new(true),
			},
		},

		HeadTracker: HeadTracker{
			HistoryDepth:            new(uint32(15)),
			MaxBufferSize:           new(uint32(17)),
			SamplingInterval:        config.MustNewDuration(time.Hour),
			FinalityTagBypass:       new(false),
			MaxAllowedFinalityDepth: new(uint32(1500)),
			PersistenceEnabled:      new(false),
			PersistenceBatchSize:    new(int64(100)),
		},

		NodePool: NodePool{
			PollFailureThreshold:                new(uint32(5)),
			PollSuccessThreshold:                new(uint32(0)),
			PollInterval:                        config.MustNewDuration(time.Minute),
			SelectionMode:                       new(multinode.NodeSelectionModeHighestHead),
			SyncThreshold:                       new(uint32(13)),
			LeaseDuration:                       config.MustNewDuration(0),
			NodeIsSyncingEnabled:                new(true),
			FinalizedBlockPollInterval:          config.MustNewDuration(time.Second),
			HistoricalBalanceCheckAddress:       new(types.MustEIP55Address("0x0000000000000000000000000000000000000001")),
			EnforceRepeatableRead:               new(true),
			DeathDeclarationDelay:               config.MustNewDuration(time.Minute),
			VerifyChainID:                       new(true),
			NewHeadsPollInterval:                config.MustNewDuration(0),
			ExternalRequestMaxResponseSize:      new(uint32(10)),
			FinalizedStateCheckFailureThreshold: new(uint32(3)),
			Errors: ClientErrors{
				NonceTooLow:                       new("(: |^)nonce too low"),
				NonceTooHigh:                      new("(: |^)nonce too high"),
				ReplacementTransactionUnderpriced: new("(: |^)replacement transaction underpriced"),
				LimitReached:                      new("(: |^)limit reached"),
				TransactionAlreadyInMempool:       new("(: |^)transaction already in mempool"),
				TerminallyUnderpriced:             new("(: |^)terminally underpriced"),
				InsufficientEth:                   new("(: |^)insufficient eth"),
				TxFeeExceedsCap:                   new("(: |^)tx fee exceeds cap"),
				L2FeeTooLow:                       new("(: |^)l2 fee too low"),
				L2FeeTooHigh:                      new("(: |^)l2 fee too high"),
				L2Full:                            new("(: |^)l2 full"),
				TransactionAlreadyMined:           new("(: |^)transaction already mined"),
				Fatal:                             new("(: |^)fatal"),
				ServiceUnavailable:                new("(: |^)service unavailable"),
				TooManyResults:                    new("(: |^)too many results"),
				MissingBlocks:                     new("(: |^)invalid block range"),
				FinalizedStateUnavailable:         new("(: |^)(missing trie node|state not available|historical state unavailable)"),
			},
		},
		OCR: OCR{
			ContractConfirmations:              new(uint16(11)),
			ContractTransmitterTransmitTimeout: config.MustNewDuration(time.Minute),
			DatabaseTimeout:                    config.MustNewDuration(time.Second),
			DeltaCOverride:                     config.MustNewDuration(time.Hour),
			DeltaCJitterOverride:               config.MustNewDuration(time.Second),
			ObservationGracePeriod:             config.MustNewDuration(time.Second),
		},
		OCR2: OCR2{
			Automation: Automation{
				GasLimit: new(uint32(540)),
			},
		},
		Workflow: Workflow{
			FromAddress:       new(types.MustEIP55Address("0x627306090abaB3A6e1400e9345bC60c78a8BEf57")),
			ForwarderAddress:  new(types.MustEIP55Address("0x9FBDa871d559710256a2502A2517b794B482Db40")),
			GasLimitDefault:   new(uint64(400000)),
			TxAcceptanceState: new(commontypes.Unconfirmed),
			PollPeriod:        config.MustNewDuration(2 * time.Second),
			AcceptanceTimeout: config.MustNewDuration(30 * time.Second),
		},
	},
	Nodes: EVMNodes{
		{
			Name:              new("foo"),
			HTTPURL:           config.MustParseURL("https://foo.web"),
			WSURL:             config.MustParseURL("wss://web.socket/test/foo"),
			HTTPURLExtraWrite: config.MustParseURL("https://foo.web/extra"),
			SendOnly:          new(false),
			Order:             new(int32(0)),
			IsLoadBalancedRPC: new(false),
		},
	},
}

func TestTOMLConfig_FullMarshal(t *testing.T) {
	configtest.AssertFullMarshal(t, fullConfig, fullTOML)
}

func TestTOMLConfig_SetFrom(t *testing.T) {
	var config EVMConfig
	config.SetFrom(&fullConfig)
	require.Equal(t, fullConfig, config)
}

func assertTOML[T any](t *testing.T, fallback, docs T) {
	t.Helper()
	t.Logf("fallback: %#v", fallback)
	t.Logf("docs: %#v", docs)
	fb, err := toml.Marshal(fallback)
	require.NoError(t, err)
	db, err := toml.Marshal(docs)
	require.NoError(t, err)
	fs, ds := string(fb), string(db)
	assert.Equal(t, fs, ds, diff.Diff(fs, ds))
}

func asEIP55Address(t *testing.T, s string) *types.EIP55Address {
	t.Helper()
	if !common.IsHexAddress(s) {
		t.Fatal("invalid address: " + s)
	}
	a := types.EIP55AddressFromAddress(common.HexToAddress(s))
	return &a
}
