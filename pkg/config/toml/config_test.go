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
	"github.com/smartcontractkit/chainlink-common/pkg/config/configtest"
	"github.com/smartcontractkit/chainlink-framework/multinode"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils/big"
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
	unknown.Workflow.GasLimitDefault = ptr(uint64(400000))
	unknown.Transactions.TransactionManagerV2.BlockTime = new(config.Duration)
	unknown.Transactions.TransactionManagerV2.CustomURL = new(config.URL)
	unknown.Transactions.TransactionManagerV2.DualBroadcast = ptr(false)
	unknown.Transactions.AutoPurge.Threshold = ptr(uint32(0))
	unknown.Transactions.AutoPurge.MinAttempts = ptr(uint32(0))
	unknown.Transactions.AutoPurge.DetectionApiUrl = new(config.URL)
	unknown.GasEstimator.BlockHistory.EIP1559FeeCapBufferBlocks = ptr[uint16](10)
	oracleType := DAOracleOPStack
	unknown.GasEstimator.DAOracle.OracleType = &oracleType
	unknown.GasEstimator.DAOracle.OracleAddress = new(types.EIP55Address)
	unknown.GasEstimator.DAOracle.CustomGasPriceCalldata = new(string)
	unknown.GasEstimator.LimitJobType = GasLimitJobType{
		OCR:    ptr[uint32](7),
		OCR2:   ptr[uint32](13),
		DR:     ptr[uint32](25),
		VRF:    ptr[uint32](37),
		FM:     ptr[uint32](42),
		Keeper: ptr[uint32](51),
	}
	unknown.GasEstimator.BumpTxDepth = ptr[uint32](15)
	unknown.NodePool.Errors = ClientErrors{
		NonceTooLow:                       ptr("too-low"),
		NonceTooHigh:                      ptr("too-high"),
		ReplacementTransactionUnderpriced: ptr("under"),
		LimitReached:                      ptr("limit"),
		TransactionAlreadyInMempool:       ptr("already"),
		TerminallyUnderpriced:             ptr("terminal"),
		InsufficientEth:                   ptr("insufficient"),
		TxFeeExceedsCap:                   ptr("exceeds"),
		L2FeeTooLow:                       ptr("low-fee"),
		L2FeeTooHigh:                      ptr("high-fee"),
		L2Full:                            ptr("full"),
		TransactionAlreadyMined:           ptr("mined"),
		Fatal:                             ptr("fatal"),
		ServiceUnavailable:                ptr("unavailable"),
		TooManyResults:                    ptr("too-many"),
	}

	configtest.AssertFieldsNotNil(t, unknown)
}

func TestDocs(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		configtest.AssertDocsTOMLComplete[EVMConfig](t, docsTOML)
	})

	t.Run("aligned", func(t *testing.T) {
		var docDefaults EVMConfig
		require.NoError(t, configtest.DocDefaultsOnly(strings.NewReader(docsTOML), &docDefaults, config.DecodeTOML))

		require.Equal(t, chaintype.ChainType(""), docDefaults.ChainType.ChainType())
		docDefaults.ChainType = nil

		// clean up KeySpecific as a special case
		require.Len(t, docDefaults.KeySpecific, 1)
		ks := KeySpecific{Key: new(types.EIP55Address),
			GasEstimator: KeySpecificGasEstimator{PriceMax: new(assets.Wei)}}
		require.Equal(t, ks, docDefaults.KeySpecific[0])
		docDefaults.KeySpecific = nil

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
		docDefaults.NodePool.Errors = ClientErrors{}

		// Transactions.AutoPurge configs are only set if the feature is enabled
		docDefaults.Transactions.AutoPurge.DetectionApiUrl = nil
		docDefaults.Transactions.AutoPurge.Threshold = nil
		docDefaults.Transactions.AutoPurge.MinAttempts = nil

		// TransactionManagerV2 configs are only set if the feature is enabled
		docDefaults.Transactions.TransactionManagerV2.BlockTime = nil
		docDefaults.Transactions.TransactionManagerV2.CustomURL = nil
		docDefaults.Transactions.TransactionManagerV2.DualBroadcast = nil

		// Fallback DA oracle is not set
		docDefaults.GasEstimator.DAOracle = DAOracle{}

		fallbackDefaults := Defaults(nil)
		assertTOML(t, fallbackDefaults, docDefaults.Chain)
	})
}

//go:embed testdata/config-full.toml
var fullTOML string

var fullConfig = EVMConfig{
	ChainID: big.NewI(42),
	Enabled: ptr(false),
	Chain: Chain{
		AutoCreateKey: ptr(false),
		BalanceMonitor: BalanceMonitor{
			Enabled: ptr(true),
		},
		BlockBackfillDepth:   ptr[uint32](100),
		BlockBackfillSkip:    ptr(true),
		ChainType:            chaintype.NewConfig("Optimism"),
		FinalityDepth:        ptr[uint32](42),
		FinalityTagEnabled:   ptr[bool](true),
		FlagsContractAddress: ptr(types.MustEIP55Address("0xae4E781a6218A8031764928E88d457937A954fC3")),
		FinalizedBlockOffset: ptr[uint32](16),

		GasEstimator: GasEstimator{
			Mode:               ptr("SuggestedPrice"),
			EIP1559DynamicFees: ptr(true),
			BumpPercent:        ptr[uint16](10),
			BumpThreshold:      ptr[uint32](6),
			BumpTxDepth:        ptr[uint32](6),
			BumpMin:            assets.NewWeiI(100),
			FeeCapDefault:      assets.NewWeiI(math.MaxInt64),
			LimitDefault:       ptr[uint64](12),
			LimitMax:           ptr[uint64](17),
			LimitMultiplier:    ptr(decimal.RequireFromString("1.234")),
			LimitTransfer:      ptr[uint64](100),
			EstimateLimit:      ptr(false),
			TipCapDefault:      assets.NewWeiI(2),
			TipCapMin:          assets.NewWeiI(1),
			PriceDefault:       assets.NewWeiI(math.MaxInt64),
			PriceMax:           assets.NewWei(new(stdbig.Int).SetBytes([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})),
			PriceMin:           assets.NewWeiI(13),

			DAOracle: DAOracle{
				OracleType:             ptr(DAOracleOPStack),
				OracleAddress:          ptr(types.MustEIP55Address("0xae4E781a6218A8031764928E88d457937A954fC3")),
				CustomGasPriceCalldata: ptr("0x1234asdf"),
			},

			LimitJobType: GasLimitJobType{
				OCR:    ptr[uint32](1001),
				DR:     ptr[uint32](1002),
				VRF:    ptr[uint32](1003),
				FM:     ptr[uint32](1004),
				Keeper: ptr[uint32](1005),
				OCR2:   ptr[uint32](1006),
			},

			BlockHistory: BlockHistoryEstimator{
				BatchSize:                 ptr[uint32](17),
				BlockHistorySize:          ptr[uint16](12),
				CheckInclusionBlocks:      ptr[uint16](18),
				CheckInclusionPercentile:  ptr[uint16](19),
				EIP1559FeeCapBufferBlocks: ptr[uint16](13),
				TransactionPercentile:     ptr[uint16](15),
			},
			FeeHistory: FeeHistoryEstimator{
				CacheTimeout: config.MustNewDuration(time.Second),
			},
		},

		KeySpecific: []KeySpecific{
			{
				Key: ptr(types.MustEIP55Address("0x2a3e23c6f242F5345320814aC8a1b4E58707D292")),
				GasEstimator: KeySpecificGasEstimator{
					PriceMax: assets.NewWei(new(stdbig.Int).SetBytes([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})),
				},
			},
		},

		LinkContractAddress:          ptr(types.MustEIP55Address("0x538aAaB4ea120b2bC2fe5D296852D948F07D849e")),
		LogBackfillBatchSize:         ptr[uint32](17),
		LogPollInterval:              config.MustNewDuration(time.Minute),
		LogKeepBlocksDepth:           ptr[uint32](100000),
		LogPrunePageSize:             ptr[uint32](0),
		BackupLogPollerBlockDelay:    ptr[uint64](532),
		MinContractPayment:           commonassets.NewLinkFromJuels(math.MaxInt64),
		MinIncomingConfirmations:     ptr[uint32](13),
		NonceAutoSync:                ptr(true),
		NoNewHeadsThreshold:          config.MustNewDuration(time.Minute),
		OperatorFactoryAddress:       ptr(types.MustEIP55Address("0xa5B85635Be42F21f94F28034B7DA440EeFF0F418")),
		LogBroadcasterEnabled:        ptr(true),
		RPCDefaultBatchSize:          ptr[uint32](17),
		RPCBlockQueryDelay:           ptr[uint16](10),
		NoNewFinalizedHeadsThreshold: config.MustNewDuration(time.Hour),

		Transactions: Transactions{
			Enabled:              ptr(true),
			MaxInFlight:          ptr[uint32](19),
			MaxQueued:            ptr[uint32](99),
			ReaperInterval:       config.MustNewDuration(time.Minute),
			ReaperThreshold:      config.MustNewDuration(time.Minute),
			ResendAfterThreshold: config.MustNewDuration(time.Hour),
			ForwardersEnabled:    ptr(true),
			AutoPurge: AutoPurgeConfig{
				Enabled:         ptr(false),
				Threshold:       ptr[uint32](42),
				MinAttempts:     ptr[uint32](13),
				DetectionApiUrl: config.MustParseURL("http://example.net"),
			},
			TransactionManagerV2: TransactionManagerV2Config{
				Enabled:       ptr(false),
				DualBroadcast: ptr(true),
				BlockTime:     config.MustNewDuration(42 * time.Second),
				CustomURL:     config.MustParseURL("http://txs.org"),
			},
		},

		HeadTracker: HeadTracker{
			HistoryDepth:            ptr[uint32](15),
			MaxBufferSize:           ptr[uint32](17),
			SamplingInterval:        config.MustNewDuration(time.Hour),
			FinalityTagBypass:       ptr[bool](false),
			MaxAllowedFinalityDepth: ptr[uint32](1500),
			PersistenceEnabled:      ptr(false),
		},

		NodePool: NodePool{
			PollFailureThreshold:       ptr[uint32](5),
			PollInterval:               config.MustNewDuration(time.Minute),
			SelectionMode:              ptr(multinode.NodeSelectionModeHighestHead),
			SyncThreshold:              ptr[uint32](13),
			LeaseDuration:              config.MustNewDuration(0),
			NodeIsSyncingEnabled:       ptr(true),
			FinalizedBlockPollInterval: config.MustNewDuration(time.Second),
			EnforceRepeatableRead:      ptr(true),
			DeathDeclarationDelay:      config.MustNewDuration(time.Minute),
			VerifyChainID:              ptr(true),
			NewHeadsPollInterval:       config.MustNewDuration(0),
			Errors: ClientErrors{
				NonceTooLow:                       ptr[string]("(: |^)nonce too low"),
				NonceTooHigh:                      ptr[string]("(: |^)nonce too high"),
				ReplacementTransactionUnderpriced: ptr[string]("(: |^)replacement transaction underpriced"),
				LimitReached:                      ptr[string]("(: |^)limit reached"),
				TransactionAlreadyInMempool:       ptr[string]("(: |^)transaction already in mempool"),
				TerminallyUnderpriced:             ptr[string]("(: |^)terminally underpriced"),
				InsufficientEth:                   ptr[string]("(: |^)insufficient eth"),
				TxFeeExceedsCap:                   ptr[string]("(: |^)tx fee exceeds cap"),
				L2FeeTooLow:                       ptr[string]("(: |^)l2 fee too low"),
				L2FeeTooHigh:                      ptr[string]("(: |^)l2 fee too high"),
				L2Full:                            ptr[string]("(: |^)l2 full"),
				TransactionAlreadyMined:           ptr[string]("(: |^)transaction already mined"),
				Fatal:                             ptr[string]("(: |^)fatal"),
				ServiceUnavailable:                ptr[string]("(: |^)service unavailable"),
				TooManyResults:                    ptr[string]("(: |^)too many results"),
			},
		},
		OCR: OCR{
			ContractConfirmations:              ptr[uint16](11),
			ContractTransmitterTransmitTimeout: config.MustNewDuration(time.Minute),
			DatabaseTimeout:                    config.MustNewDuration(time.Second),
			DeltaCOverride:                     config.MustNewDuration(time.Hour),
			DeltaCJitterOverride:               config.MustNewDuration(time.Second),
			ObservationGracePeriod:             config.MustNewDuration(time.Second),
		},
		OCR2: OCR2{
			Automation: Automation{
				GasLimit: ptr[uint32](540),
			},
		},
		Workflow: Workflow{
			FromAddress:      ptr(types.MustEIP55Address("0x627306090abaB3A6e1400e9345bC60c78a8BEf57")),
			ForwarderAddress: ptr(types.MustEIP55Address("0x9FBDa871d559710256a2502A2517b794B482Db40")),
			GasLimitDefault:  ptr[uint64](400000),
		},
	},
	Nodes: EVMNodes{
		{
			Name:              ptr("foo"),
			HTTPURL:           config.MustParseURL("https://foo.web"),
			WSURL:             config.MustParseURL("wss://web.socket/test/foo"),
			HTTPURLExtraWrite: config.MustParseURL("https://foo.web/extra"),
			SendOnly:          ptr(false),
			Order:             ptr[int32](0),
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

func ptr[T any](t T) *T {
	return &t
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
