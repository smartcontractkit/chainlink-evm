package client

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-framework/multinode"

	"github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
)

func TestRPCClient_MakeLogsValid(t *testing.T) {
	chainTypes := []struct {
		Name      string
		ChainType chaintype.ChainType
	}{
		{Name: "Sei", ChainType: chaintype.ChainSei},
		{Name: "Hedera", ChainType: chaintype.ChainHedera},
		{Name: "Rootstock", ChainType: chaintype.ChainRootstock},
		{Name: "Pharos", ChainType: chaintype.ChainPharos},
	}

	testCases := []struct {
		Name             string
		TxIndex          uint
		LogIndex         uint
		ExpectedLogIndex uint
		ExpectedError    error
	}{
		{
			Name:             "TxIndex = 0 LogIndex = 0",
			TxIndex:          0,
			LogIndex:         0,
			ExpectedLogIndex: 0,
			ExpectedError:    nil,
		},
		{
			Name:             "TxIndex = 0 LogIndex = 1",
			TxIndex:          0,
			LogIndex:         1,
			ExpectedLogIndex: 1,
			ExpectedError:    nil,
		},
		{
			Name:             "TxIndex = 0 LogIndex = MaxUint32",
			TxIndex:          0,
			LogIndex:         math.MaxUint32,
			ExpectedLogIndex: math.MaxUint32,
			ExpectedError:    nil,
		},
		{
			Name:             "LogIndex = MaxUint32 + 1 => returns an error",
			TxIndex:          0,
			LogIndex:         math.MaxUint32 + 1,
			ExpectedLogIndex: 0,
			ExpectedError:    errors.New("log's index 4294967296 of tx 0x0000000000000000000000000000000000000000000000000000000000000000 exceeds max supported value of 4294967295"),
		},
		{
			Name:             "TxIndex = 1 LogIndex = 0",
			TxIndex:          1,
			LogIndex:         0,
			ExpectedLogIndex: math.MaxUint32 + 1,
			ExpectedError:    nil,
		},
		{
			Name:             "TxIndex = MaxUint32 LogIndex = MaxUint32",
			TxIndex:          math.MaxUint32,
			LogIndex:         math.MaxUint32,
			ExpectedLogIndex: math.MaxUint64,
			ExpectedError:    nil,
		},
		{
			Name:             "TxIndex = MaxUint32 + 1 => returns an error",
			TxIndex:          math.MaxUint32 + 1,
			LogIndex:         0,
			ExpectedLogIndex: 0,
			ExpectedError:    errors.New("TxIndex of tx 0x0000000000000000000000000000000000000000000000000000000000000000 exceeds max supported value of 4294967295"),
		},
	}

	for _, ct := range chainTypes {
		t.Run(ct.Name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.Name, func(t *testing.T) {
					rpc := NewTestRPCClient(t, RPCClientOpts{ChainType: ct.ChainType})
					log, err := rpc.makeLogValid(ethtypes.Log{TxIndex: tc.TxIndex, Index: tc.LogIndex})
					if tc.ExpectedError != nil {
						require.EqualError(t, err, tc.ExpectedError.Error())
						return
					}
					require.Equal(t, tc.ExpectedLogIndex, log.Index)
					require.Equal(t, tc.TxIndex, log.TxIndex)
				})
			}
		})
	}

	t.Run("Other chains", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.Name, func(t *testing.T) {
				rpc := NewTestRPCClient(t, RPCClientOpts{})
				log, err := rpc.makeLogValid(ethtypes.Log{TxIndex: tc.TxIndex, Index: tc.LogIndex})
				// other chains should return as is
				require.NoError(t, err)
				require.Equal(t, tc.TxIndex, log.TxIndex)
				require.Equal(t, tc.LogIndex, log.Index)
			})
		}
	})
}

func NewDialedTestRPCClient(t *testing.T, opts RPCClientOpts) *RPCClient {
	rpcClient := NewTestRPCClient(t, opts)
	t.Cleanup(func() {
		rpcClient.Close()
	})
	err := rpcClient.Dial(t.Context())
	require.NoError(t, err)
	return rpcClient
}

func MakeHeadMsgForNumber(number uint64) string {
	return fmt.Sprintf(`{"number":"%s","author":"0x1687736326c9fea17e25fc5287613693c912909c","baseFeePerGas":"0x3b9aca00","difficulty":"0x0","extraData":"0x","gasLimit":"0xe4e1c0","gasUsed":"0x0","hash":"0x62f03413681948b06882e7d9f91c4949bc39ded98d36336ab03faea038ec8e3d","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x1687736326c9fea17e25fc5287613693c912909c","nonce":"0x0000000000000000","parentHash":"0x43f504afdc639cbb8daf5fd5328a37762164b73f9c70ed54e1928c1fca6d8f23","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x200","stateRoot":"0x0cb938d51ad83bdf401e3f5f7f989e60df64fdea620d394af41a3e72629f7495","timestamp":"0x61bd8d1a","totalDifficulty":"0x0","transactions":[],"transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","uncles":[]}`,
		hexutil.EncodeUint64(number))
}

func TestRPCClient_doWithConfidence(t *testing.T) {
	t.Parallel()
	chainID := big.NewInt(1234567)

	testCases := []struct {
		Name string
		// input
		CallBlockNumber *big.Int
		Confidence      primitives.ConfidenceLevel
		// returned values
		EthCallResult       string
		EthCallError        string
		ExpectedTag         string
		BlockByNumberResult string
		BlockByNumberError  string
		// expectations
		ExpectedError  string
		ExpectedResult string
	}{
		{
			Name:          "Returns error if block number is nil",
			Confidence:    primitives.Safe,
			ExpectedError: "blockNumber must be non nil and fit into int64. Got: <nil>",
		},
		{
			Name:            "Returns error if confidence level if not supported",
			Confidence:      primitives.Unconfirmed,
			CallBlockNumber: big.NewInt(1),
			ExpectedError:   "confidence level unconfirmed not supported",
		},
		{
			Name:                "Failed to make user's call",
			Confidence:          primitives.Finalized,
			CallBlockNumber:     big.NewInt(5),
			EthCallResult:       "0x00",
			EthCallError:        "call_contract failed",
			ExpectedTag:         "finalized",
			BlockByNumberResult: MakeHeadMsgForNumber(4),
			ExpectedError:       "RPC call failed: caller request failed: call_contract failed",
		},
		{
			Name:                "Failed to make reference call",
			Confidence:          primitives.Finalized,
			CallBlockNumber:     big.NewInt(5),
			EthCallResult:       "0x00",
			ExpectedTag:         "finalized",
			BlockByNumberResult: MakeHeadMsgForNumber(4),
			BlockByNumberError:  "failed to get finalized block number",
			ExpectedError:       "RPC call failed: referenced block request failed: failed to get finalized block number",
		},
		{
			Name:                "Finalized block number if smaller than request",
			Confidence:          primitives.Finalized,
			CallBlockNumber:     big.NewInt(5),
			EthCallResult:       "0x00",
			ExpectedTag:         "finalized",
			BlockByNumberResult: MakeHeadMsgForNumber(4),
			ExpectedError:       "data was requested at block 5 while max available height with confidence level finalized is 4",
		},
		{
			Name:                "Safe block number if smaller than request",
			Confidence:          primitives.Safe,
			CallBlockNumber:     big.NewInt(9),
			EthCallResult:       "0x00",
			ExpectedTag:         "safe",
			BlockByNumberResult: MakeHeadMsgForNumber(8),
			ExpectedError:       "data was requested at block 9 while max available height with confidence level safe is 8",
		},
		{
			Name:                "Returns error if requested tag is not supported",
			Confidence:          primitives.Safe,
			CallBlockNumber:     big.NewInt(9),
			EthCallResult:       "0x00",
			ExpectedTag:         "safe",
			BlockByNumberResult: "null",
			ExpectedError:       "referenced block request returned nil. RPC is unhealthy or chain does not support specified tag",
		},
		{
			Name:                "Happy path",
			Confidence:          primitives.Safe,
			CallBlockNumber:     big.NewInt(8),
			EthCallResult:       "0x" + hex.EncodeToString([]byte("happy path result")),
			ExpectedTag:         "safe",
			BlockByNumberResult: MakeHeadMsgForNumber(8),
			ExpectedResult:      "happy path result",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			wsURL := testutils.NewWSServer(t, chainID, func(method string, params gjson.Result) (resp testutils.JSONRPCResponse) {
				switch method {
				case "eth_call":
					resp.Result = fmt.Sprintf(`"%s"`, tc.EthCallResult)
					resp.Error.Message = tc.EthCallError
				case "eth_getBlockByNumber":
					require.True(t, params.IsArray())
					require.Equal(t, params.Array()[0].String(), tc.ExpectedTag)
					resp.Result = tc.BlockByNumberResult
					resp.Error.Message = tc.BlockByNumberError
				default:
					require.Fail(t, "unexpected method: "+method)
				}
				return
			}).WSURL()
			rpcClient := NewDialedTestRPCClient(t, RPCClientOpts{HTTP: wsURL, FinalityTagsEnabled: true, ChainID: chainID})
			var result hexutil.Bytes
			err := rpcClient.doWithConfidence(t.Context(), rpc.BatchElem{Method: "eth_call", Result: &result}, tc.CallBlockNumber, tc.Confidence)
			if tc.ExpectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedResult, string(result))
			}
		})
	}
}

func TestRPCClient_confidenceToBlockNumber(t *testing.T) {
	testCases := []struct {
		Name               string
		ConfidenceLevel    primitives.ConfidenceLevel
		FinalityTagEnabled bool
		ExpectedBlock      rpc.BlockNumber
		ExpectedError      string
	}{
		{
			Name:               "Finalized confidence level with finality tags enabled",
			ConfidenceLevel:    primitives.Finalized,
			FinalityTagEnabled: true,
			ExpectedBlock:      rpc.FinalizedBlockNumber,
		},
		{
			Name:               "Safe confidence level with finality tags enabled",
			ConfidenceLevel:    primitives.Safe,
			FinalityTagEnabled: true,
			ExpectedBlock:      rpc.SafeBlockNumber,
		},
		{
			Name:               "Unconfirmed confidence level with finality tags enabled",
			ConfidenceLevel:    primitives.Unconfirmed,
			FinalityTagEnabled: true,
			ExpectedError:      "confidence level unconfirmed not supported",
		},
		{
			Name:               "Finalized confidence level with finality tags disabled",
			ConfidenceLevel:    primitives.Finalized,
			FinalityTagEnabled: false,
			ExpectedBlock:      rpc.LatestBlockNumber,
		},
		{
			Name:               "Safe confidence level with finality tags disabled",
			ConfidenceLevel:    primitives.Safe,
			FinalityTagEnabled: false,
			ExpectedBlock:      rpc.LatestBlockNumber,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rpcClient := NewTestRPCClient(t, RPCClientOpts{FinalityTagsEnabled: tc.FinalityTagEnabled})
			block, err := rpcClient.confidenceToBlockNumber(tc.ConfidenceLevel)
			if tc.ExpectedError != "" {
				require.ErrorContains(t, err, tc.ExpectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedBlock, block)
			}
		})
	}
}

func TestRPCClient_referenceHeadToMaxAvailableHeight(t *testing.T) {
	testCases := []struct {
		Name                string
		FinalityTagEnabled  bool
		FinalityDepth       uint32
		SafeDepth           uint32
		ReferenceHeadHeight int64
		ConfidenceLevel     primitives.ConfidenceLevel
		ExpectedBlock       int64
		ExpectedError       string
	}{
		{
			Name:                "Finalized confidence level with finality tags enabled",
			FinalityTagEnabled:  true,
			ReferenceHeadHeight: 100,
			ConfidenceLevel:     primitives.Finalized,
			ExpectedBlock:       100,
		},
		{
			Name:                "Safe confidence level with finality tags enabled",
			FinalityTagEnabled:  true,
			ReferenceHeadHeight: 100,
			ConfidenceLevel:     primitives.Safe,
			ExpectedBlock:       100,
		},
		{
			Name:                "Finalized confidence level with finality tags disabled",
			FinalityTagEnabled:  false,
			FinalityDepth:       10,
			ReferenceHeadHeight: 100,
			ConfidenceLevel:     primitives.Finalized,
			ExpectedBlock:       90,
		},
		{
			Name:                "Safe confidence level with finality tags disabled",
			FinalityTagEnabled:  false,
			ReferenceHeadHeight: 100,
			SafeDepth:           16,
			ConfidenceLevel:     primitives.Safe,
			ExpectedBlock:       84,
		},
		{
			Name:            "Returns error on not supported confidence level",
			ConfidenceLevel: primitives.Unconfirmed,
			ExpectedError:   "confidence level unconfirmed not supported",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rpcClient := NewTestRPCClient(t, RPCClientOpts{
				FinalityTagsEnabled: tc.FinalityTagEnabled,
				FinalityDepth:       tc.FinalityDepth,
				SafeDepth:           tc.SafeDepth,
			})

			block, err := rpcClient.referenceHeadToMaxAvailableHeight(tc.ConfidenceLevel, tc.ReferenceHeadHeight)
			if tc.ExpectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedBlock, block)
			}
		})
	}
}

type RPCClientOpts struct {
	Cfg                            config.NodePool
	Lggr                           logger.Logger
	WS                             *url.URL
	HTTP                           *url.URL
	Name                           string
	ID                             int
	ChainID                        *big.Int
	Tier                           multinode.NodeTier
	LargePayloadRPCTimeout         *time.Duration
	RPCTimeout                     *time.Duration
	ChainType                      chaintype.ChainType
	FinalityTagsEnabled            bool
	FinalityDepth                  uint32
	SafeDepth                      uint32
	ExternalRequestMaxResponseSize uint32
}

func NewTestRPCClient(t *testing.T, opts RPCClientOpts) *RPCClient {
	if opts.Lggr == nil {
		opts.Lggr = logger.Test(t)
	}

	if opts.ChainID == nil {
		opts.ChainID = testutils.FixtureChainID
	}

	if opts.Cfg == nil {
		opts.Cfg = &TestNodePoolConfig{
			NodeNewHeadsPollInterval:       1 * time.Second,
			NodeFinalizedBlockPollInterval: 1 * time.Second,
		}
	}

	if opts.RPCTimeout == nil {
		opts.RPCTimeout = ptr(QueryTimeout)
	}

	if opts.LargePayloadRPCTimeout == nil {
		opts.LargePayloadRPCTimeout = ptr(QueryTimeout)
	}

	if opts.Name == "" {
		opts.Name = "rpc"
	}

	return NewRPCClient(opts.Cfg, opts.Lggr, opts.WS, opts.HTTP, opts.Name, opts.ID, opts.ChainID, opts.Tier,
		*opts.LargePayloadRPCTimeout, *opts.RPCTimeout, opts.ChainType, opts.FinalityTagsEnabled, opts.FinalityDepth,
		opts.SafeDepth, opts.ExternalRequestMaxResponseSize)
}

func ptr[T any](v T) *T {
	return &v
}
