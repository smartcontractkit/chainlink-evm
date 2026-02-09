package integrationtests

import (
	"crypto/ecdsa"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	tomldecode "github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

const (
	envPath    = "env.toml"
	configPath = "configs.toml"
)

// This wrapper is required because of the way Gas Estimator components expect configs.
// Instead of passing down a struct with values, we need to implement an interface with
// the required methods.
type AppConfig struct {
	BlockTimeF          time.Duration
	EIP1559DynamicFeesF bool
	BumpPercentF        uint16
	BumpThresholdF      uint64
	BumpTxDepthF        uint32
	BumpMinF            *assets.Wei
	FeeCapDefaultF      *assets.Wei
	LimitDefaultF       uint64
	LimitMaxF           uint64
	LimitMultiplierF    float32
	LimitTransferF      uint64
	PriceDefaultF       *assets.Wei
	TipCapDefaultF      *assets.Wei
	TipCapMinF          *assets.Wei
	PriceMaxF           *assets.Wei
	PriceMinF           *assets.Wei
	ModeF               string
	EstimateLimitF      bool
	SenderAddressF      *types.EIP55Address
	FeeHistoryF         *FeeHistory
	BlockHistoryF       *BlockHistory
}

func (a AppConfig) BlockTime() time.Duration {
	return a.BlockTimeF
}

func (a AppConfig) PriceMaxKey(common.Address) *assets.Wei {
	return a.PriceMaxF
}

func (a AppConfig) EIP1559DynamicFees() bool {
	return a.EIP1559DynamicFeesF
}

func (a AppConfig) BumpPercent() uint16 {
	return a.BumpPercentF
}

func (a AppConfig) BumpThreshold() uint64 {
	return a.BumpThresholdF
}

func (a AppConfig) BumpTxDepth() uint32 {
	return a.BumpTxDepthF
}
func (a AppConfig) BumpMin() *assets.Wei {
	return a.BumpMinF
}

func (a AppConfig) TipCapMin() *assets.Wei {
	return a.TipCapMinF
}

func (a AppConfig) PriceMax() *assets.Wei {
	return a.PriceMaxF
}

func (a AppConfig) PriceMin() *assets.Wei {
	return a.PriceMinF
}

func (a AppConfig) Mode() string {
	return a.ModeF
}

func (a AppConfig) PriceDefault() *assets.Wei {
	return a.PriceDefaultF
}

func (a AppConfig) TipCapDefault() *assets.Wei {
	return a.TipCapDefaultF
}

func (a AppConfig) FeeCapDefault() *assets.Wei {
	return a.FeeCapDefaultF
}

func (a AppConfig) LimitDefault() uint64 {
	return a.LimitDefaultF
}

func (a AppConfig) LimitMax() uint64 {
	return a.LimitMaxF
}

func (a AppConfig) LimitMultiplier() float32 {
	return a.LimitMultiplierF
}

func (a AppConfig) LimitTransfer() uint64 {
	return a.LimitTransferF
}

func (a AppConfig) EstimateLimit() bool {
	return a.EstimateLimitF
}

func (a AppConfig) SenderAddress() *types.EIP55Address {
	return a.SenderAddressF
}

// -------------------------------
func (a AppConfig) DAOracle() config.DAOracle {
	return &DAOracle{}
}

type DAOracle struct {
	OracleTypeF             *toml.DAOracleType
	OracleAddressF          *types.EIP55Address
	CustomGasPriceCalldataF *string
}

func (o DAOracle) OracleType() *toml.DAOracleType {
	return o.OracleTypeF
}

func (o DAOracle) OracleAddress() *types.EIP55Address {
	return o.OracleAddressF
}

func (o DAOracle) CustomGasPriceCalldata() *string {
	return o.CustomGasPriceCalldataF
}

// -------------------------------
func (a AppConfig) LimitJobType() config.LimitJobType {
	return nil
}

// -------------------------------
func (a AppConfig) FeeHistory() config.FeeHistory {
	return a.FeeHistoryF
}

type FeeHistory struct {
	CacheTimeoutF time.Duration
}

func (b FeeHistory) CacheTimeout() time.Duration {
	return b.CacheTimeoutF
}

// -------------------------------
func (a AppConfig) BlockHistory() config.BlockHistory {
	return a.BlockHistoryF
}

type BlockHistory struct {
	BatchSizeF                 uint32
	BlockHistorySizeF          uint16
	BlockDelayF                uint16
	CheckInclusionBlocksF      uint16
	CheckInclusionPercentileF  uint16
	EIP1559FeeCapBufferBlocksF uint16
	TransactionPercentileF     uint16
}

func (b BlockHistory) BatchSize() uint32 {
	return b.BatchSizeF
}

func (b BlockHistory) BlockHistorySize() uint16 {
	return b.BlockHistorySizeF
}

func (b BlockHistory) BlockDelay() uint16 {
	return b.BlockDelayF
}

func (b BlockHistory) CheckInclusionBlocks() uint16 {
	return b.CheckInclusionBlocksF
}

func (b BlockHistory) CheckInclusionPercentile() uint16 {
	return b.CheckInclusionPercentileF
}

func (b BlockHistory) EIP1559FeeCapBufferBlocks() uint16 {
	return b.EIP1559FeeCapBufferBlocksF
}

func (b BlockHistory) TransactionPercentile() uint16 {
	return b.TransactionPercentileF
}

// EnvVariables holds the environment variables.
type EnvVariables struct {
	RPC         string `toml:"RPC"`
	PrivateKey  string `toml:"PrivateKey"`
	FromAddress string
}

// LoadEnvVariables loads the environment variables from a config file.
// The default config file is "env.toml"
func LoadEnvVariables(t *testing.T) *EnvVariables {
	data, err := os.ReadFile(envPath)
	require.NoError(t, err)

	envs := &EnvVariables{}
	require.NoError(t, tomldecode.Unmarshal(data, envs), "failed to parse TOML config")
	privateKey, err := crypto.HexToECDSA(envs.PrivateKey)
	require.NoError(t, err)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok, "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	envs.FromAddress = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return envs
}

// LoadConfigVariables loads the configs variables from a config file.
// The default config file is "config.toml"
// We optimistically expect to load the configs that affect the tests and we load the default
// values for the configs that don't to minimize the friction for the tester.
func LoadConfigVariablesWithDefaults(t *testing.T) AppConfig {
	data, err := os.ReadFile(configPath)
	require.NoError(t, err)

	// First unmarshal into a temporary struct that uses int64 for BlockTime
	type tempConfig struct {
		BlockTimeSeconds    int64  `toml:"BlockTime"`
		EIP1559DynamicFeesF bool   `toml:"EIP1559DynamicFees"`
		BumpThresholdF      uint64 `toml:"BumpThreshold"`
	}
	temp := &tempConfig{}
	require.NoError(t, tomldecode.Unmarshal(data, temp), "failed to parse TOML config")

	// Convert BlockTime from seconds to time.Duration
	configs := &AppConfig{
		BlockTimeF:          time.Duration(temp.BlockTimeSeconds) * time.Second,
		EIP1559DynamicFeesF: temp.EIP1559DynamicFeesF,
		BumpThresholdF:      temp.BumpThresholdF,
	}

	// Add default values to the configs that don't affect the tests
	configs.BumpPercentF = 20
	configs.LimitDefaultF = 30000
	configs.LimitTransferF = 21000
	configs.LimitMultiplierF = 1
	configs.PriceMaxF = assets.GWei(1000)
	configs.ModeF = "FeeHistory"
	configs.FeeHistoryF = &FeeHistory{
		CacheTimeoutF: configs.BlockTimeF,
	}
	configs.BlockHistoryF = &BlockHistory{
		BlockHistorySizeF:      8,
		TransactionPercentileF: 55,
	}
	return *configs
}

// defaultConfigs returns the default configurations for Ethereum chains (Mainnet and Testnet).
func defaultConfigs() AppConfig {
	return AppConfig{
		BlockTimeF:          testutils.TestInterval,
		EIP1559DynamicFeesF: true,
		BumpPercentF:        20,
		BumpThresholdF:      3,
		LimitDefaultF:       30000,
		LimitTransferF:      21000,
		LimitMultiplierF:    1,
		PriceMaxF:           assets.GWei(1000),
		ModeF:               "FeeHistory",
		FeeHistoryF: &FeeHistory{
			CacheTimeoutF: time.Second,
		},
		BlockHistoryF: &BlockHistory{
			BlockHistorySizeF:      14,
			TransactionPercentileF: 55,
		},
	}
}
