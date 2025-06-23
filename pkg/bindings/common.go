package bindings

import (
	"math/big"

	evmcappb "github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/chain-capabilities/evm"
	"github.com/smartcontractkit/chainlink-common/pkg/chains/evm"
	chain_common "github.com/smartcontractkit/chainlink-common/pkg/loop/chain-common"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk/v2"
)

// This function is not EVM specific, it's generic and should be provided by CRE
func GenerateReport(chainID uint32, userData []byte) commonReport {
	return commonReport{}
}

// Minimal Chain Capabilities SDK client interface.
type EVMClient interface {
	CallContract(sdk.Runtime, *evm.CallContractRequest) sdk.Promise[*evm.CallContractReply]
	WriteReport(sdk.Runtime, *evmcappb.WriteReportRequest) sdk.Promise[*evmcappb.WriteReportReply]
	RegisterLogTracking(sdk.Runtime, *evm.RegisterLogTrackingRequest)
	UnregisterLogTracking(sdk.Runtime, *evm.UnregisterLogTrackingRequest)
	FilterLogs(sdk.Runtime, *evm.FilterLogsRequest) sdk.Promise[*evm.FilterLogsReply]
}

// This is not EVM specific, it's generic
type commonReport struct {
	RawReport     []byte
	ReportContext []byte
	Signatures    [][]byte
	ID            []byte
}

// Define a custom error type
type TxFatalError struct {
	Message string
}

// Implement the error interface
func (e *TxFatalError) Error() string {
	return ("Error " + e.Message)
}

// Define a custom error type
type ReceiverContractError struct {
	Message string
	TxHash  []byte
}

// Implement the error interface
func (e *ReceiverContractError) Error() string {
	return ("Error " + e.Message)
}

type ContractInitOptions struct {
	GasConfig *evmcappb.GasConfig
}

type ReadOptions struct {
	BlockNumber *big.Int
}

type WriteOptions struct {
	GasConfig  *evmcappb.GasConfig
	BlockDepth uint16 // 0 means finalized, 1 confirmed, positive numbers block depth - TODO to be defined together with all other operations
}

// Logs support
const FINALIZED = 0
const CONFIRMED = 1

type LogTrackingOptions struct {
	MaxLogsKept   uint64 `protobuf:"varint,1,opt,name=max_logs_kept,json=maxLogsKept,proto3" json:"max_logs_kept,omitempty"`     // maximum number of logs to retain ( 0 = unlimited )
	RetentionTime int64  `protobuf:"varint,2,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time,omitempty"` // maximum amount of time to retain logs in seconds
	LogsPerBlock  uint64 `protobuf:"varint,3,opt,name=logs_per_block,json=logsPerBlock,proto3" json:"logs_per_block,omitempty"`  // rate limit ( maximum # of logs per block, 0 = unlimited )
	// TODO could this be actual values for the indexed log fields instead of hashes?
	Topic2 [][]byte `protobuf:"bytes,7,rep,name=topic2,proto3" json:"topic2,omitempty"` // list of possible values for topic2
	Topic3 [][]byte `protobuf:"bytes,8,rep,name=topic3,proto3" json:"topic3,omitempty"` // list of possible values for topic3
	Topic4 [][]byte `protobuf:"bytes,9,rep,name=topic4,proto3" json:"topic4,omitempty"` // list of possible values for topic4
}

type FilterLogTrigger struct {
	Confidence chain_common.Confidence
	BlockDepth uint64
}

type ParsedLog[T any] struct {
	LogData T
	RawLog  *evm.Log
}

type FilterOptions struct {
	BlockHash []byte
	FromBlock *big.Int
	ToBlock   *big.Int
}

func ValidateLogTrackingOptions(opts *LogTrackingOptions) {
	// TODO: set defaults for opts
	if opts.MaxLogsKept == 0 {
		opts.MaxLogsKept = 1000
	}
	if opts.RetentionTime == 0 {
		opts.RetentionTime = 86400
	}
	if opts.LogsPerBlock == 0 {
		opts.LogsPerBlock = 100
	}
}
