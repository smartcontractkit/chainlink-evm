package bindings

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/sdk"
)

// Minimal Chain Capabilities SDK client interface.
type EVMClient interface {
	CallContract(sdk.Runtime, *evm.CallContractRequest) sdk.Promise[*evm.CallContractReply]
	RegisterLogTracking(sdk.Runtime, *evm.RegisterLogTrackingRequest)
	UnregisterLogTracking(sdk.Runtime, *evm.UnregisterLogTrackingRequest)
	FilterLogs(sdk.Runtime, *evm.FilterLogsRequest) sdk.Promise[*evm.FilterLogsReply]
	LogTrigger(config *evm.FilterLogTriggerRequest) sdk.Trigger[*evm.Log, *evm.Log]
}

type ContractInitOptions struct {
	GasConfig *evm.GasConfig
}

type ReadOptions struct {
	BlockNumber *big.Int
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

func EncodeTopics(evt abi.Event, values ...interface{}) ([][]byte, error) {
	var indexed []abi.Argument
	for _, arg := range evt.Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if len(values) != len(indexed) {
		return nil, fmt.Errorf("wrong arg count: %d values vs %d indexed params",
			len(values), len(indexed))
	}
	topics := make([][]byte, 0, len(values)+1)
	topics = append(topics, evt.ID.Bytes())

	for i, arg := range indexed {
		packed, err := abi.Arguments{arg}.Pack(values[i])
		if err != nil {
			return nil, err
		}

		topics = append(topics, common.BytesToHash(crypto.Keccak256(packed)).Bytes())
	}
	return topics, nil
}
