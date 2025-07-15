package bindings

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	ocr3types "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/sdk"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk/v2/pb"
)

var _ EVMClient = &evm.Client{}

// Minimal Chain Capabilities SDK client interface.
type EVMClient interface {
	CallContract(sdk.Runtime, *evm.CallContractRequest) sdk.Promise[*evm.CallContractReply]
	WriteReport(sdk.Runtime, *evm.WriteReportRequest) sdk.Promise[*evm.WriteReportReply]
	LatestAndFinalizedHead(runtime sdk.Runtime, input *emptypb.Empty) sdk.Promise[*evm.LatestAndFinalizedHeadReply]
	RegisterLogTracking(sdk.Runtime, *evm.RegisterLogTrackingRequest) sdk.Promise[*emptypb.Empty]
	UnregisterLogTracking(sdk.Runtime, *evm.UnregisterLogTrackingRequest) sdk.Promise[*emptypb.Empty]
	FilterLogs(sdk.Runtime, *evm.FilterLogsRequest) sdk.Promise[*evm.FilterLogsReply]
}

type ContractInitOptions struct {
	GasConfig *evm.GasConfig
}

type ReadOptions struct {
	BlockNumber *big.Int
}

type WriteOptions struct {
	GasConfig  *evm.GasConfig
	BlockDepth uint16 // 0 means finalized, 1 confirmed, positive numbers block depth - TODO to be defined together with all other operations
}

type LogTrackingOptions[T any] struct {
	MaxLogsKept   uint64 `protobuf:"varint,1,opt,name=max_logs_kept,json=maxLogsKept,proto3" json:"max_logs_kept,omitempty"`     // maximum number of logs to retain ( 0 = unlimited )
	RetentionTime int64  `protobuf:"varint,2,opt,name=retention_time,json=retentionTime,proto3" json:"retention_time,omitempty"` // maximum amount of time to retain logs in seconds
	LogsPerBlock  uint64 `protobuf:"varint,3,opt,name=logs_per_block,json=logsPerBlock,proto3" json:"logs_per_block,omitempty"`  // rate limit ( maximum # of logs per block, 0 = unlimited )
	Filters       []T
}

type FilterOptions struct {
	BlockHash []byte
	FromBlock *big.Int
	ToBlock   *big.Int
}

func ValidateLogTrackingOptions[T any](opts *LogTrackingOptions[T]) {
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

func ExtractID(report []byte) ([]byte, error) {
	metadata, _, err := ocr3types.Decode(report)
	if err != nil {
		return nil, err
	}
	return []byte(metadata.ReportID), nil
}

func ExtractSigs(attrSigs []*pb2.AttributedSignature) [][]byte {
	sigs := make([][]byte, len(attrSigs))
	for i, sig := range attrSigs {
		sigs[i] = sig.Signature
	}
	return sigs
}

func PrepareTopicArg(arg abi.Argument, value interface{}) (interface{}, error) {
	t := reflect.TypeOf(value)

	// only pre-hash:
	//  - dynamic slices that aren't []byte
	//  - fixed arrays that aren't [N]byte
	//  - structs (i.e. tuple types)
	if (t.Kind() == reflect.Slice && t.Elem().Kind() != reflect.Uint8) ||
		(t.Kind() == reflect.Array && t.Elem().Kind() != reflect.Uint8) ||
		t.Kind() == reflect.Struct {

		packed, err := abi.Arguments{arg}.Pack(value)
		if err != nil {
			return nil, fmt.Errorf("packing %q for topic: %w", arg.Name, err)
		}
		// hash the packed bytes:
		return crypto.Keccak256Hash(packed), nil
	}

	return value, nil
}

func PadTopics(topics []*evm.TopicValues) []*evm.TopicValues {
	for i := len(topics); i < 4; i++ {
		topics = append(topics, &evm.TopicValues{
			Values: [][]byte{},
		})
	}

	return topics
}
