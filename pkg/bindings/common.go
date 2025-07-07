package bindings

import (
	"math/big"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/sdk"
	"google.golang.org/protobuf/types/known/emptypb"
)

// This function is not EVM specific, it's generic and should be provided by CRE
// TO BE REPLACED BY SDK FUNCTION
func GenerateReport(chainID uint32, userData []byte) evmcappb.SignedReport {
	return evmcappb.SignedReport{}
}

// Minimal Chain Capabilities SDK client interface.
type EVMClient interface {
	CallContract(sdk.Runtime, *evm.CallContractRequest) sdk.Promise[*evm.CallContractReply]
	WriteReport(sdk.Runtime, *evmcappb.WriteReportRequest) sdk.Promise[*evmcappb.WriteReportReply]
	RegisterLogTracking(sdk.Runtime, *evm.RegisterLogTrackingRequest)
	UnregisterLogTracking(sdk.Runtime, *evm.UnregisterLogTrackingRequest)
	FilterLogs(sdk.Runtime, *evm.FilterLogsRequest) sdk.Promise[*evm.FilterLogsReply]
	LogTrigger(config *evm.FilterLogTriggerRequest) sdk.Trigger[*evm.Log, *evm.Log]
	LatestAndFinalizedHead(runtime sdk.Runtime, input *emptypb.Empty) sdk.Promise[*evm.LatestAndFinalizedHeadReply]
}

type ContractInitOptions struct {
	GasConfig *evm.GasConfig
}

type ReadOptions struct {
	BlockNumber *big.Int
}

type WriteOptions struct {
	GasConfig  *evmcappb.GasConfig
	BlockDepth uint16 // 0 means finalized, 1 confirmed, positive numbers block depth - TODO to be defined together with all other operations
}

type LogTrackingOptions struct {
	MaxLogsKept   uint64   `protobuf:"varint,1,opt,name=max_logs_kept,json=maxLogsKept,proto3" json:"max_logs_kept,omitempty"`     // maximum number of logs to retain ( 0 = unlimited )
	RetentionTime int64    `protobuf:"varint,2,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time,omitempty"` // maximum amount of time to retain logs in seconds
	LogsPerBlock  uint64   `protobuf:"varint,3,opt,name=logs_per_block,json=logsPerBlock,proto3" json:"logs_per_block,omitempty"`  // rate limit ( maximum # of logs per block, 0 = unlimited )
	Topic2        [][]byte `protobuf:"bytes,7,rep,name=topic2,proto3" json:"topic2,omitempty"`                                     // list of possible values for topic2
	Topic3        [][]byte `protobuf:"bytes,8,rep,name=topic3,proto3" json:"topic3,omitempty"`                                     // list of possible values for topic3
	Topic4        [][]byte `protobuf:"bytes,9,rep,name=topic4,proto3" json:"topic4,omitempty"`                                     // list of possible values for topic4
}

type FilterOptions struct {
	BlockHash []byte
	FromBlock *big.Int
	ToBlock   *big.Int
}

func ValidateLogTrackingOptions(opts *LogTrackingOptions) {
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
