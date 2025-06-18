// Code generated — DO NOT EDIT.

package {{.Package}}

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-evm/pkg/bindings"
	evmcappb "github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/chain-capabilities/evm"
	"github.com/smartcontractkit/chainlink-common/pkg/chains/evm"
	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk/v2"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

{{range $contract := .Contracts}}

var {{$contract.Type}}MetaData = &bind.MetaData{
	ABI: "{{.InputABI}}",
	{{- if .InputBin}}
	Bin: "0x{{.InputBin}}",
	{{- end}}
}

type {{$contract.Type}}Codec interface {
	{{- range $call := $contract.Calls}}
	{{- if or $call.Original.Constant (eq $call.Original.StateMutability "view")}}
	Encode{{$call.Normalized.Name}}MethodCall(in {{$call.Normalized.Name}}Input) ([]byte, error)
	Decode{{$call.Normalized.Name}}MethodOutput(data []byte) ({{with index $call.Normalized.Outputs 0}}{{bindtype .Type $.Structs}}{{end}}, error)
	{{- end}}
	{{- end}}

	{{- range $.Structs}}
	Encode{{.Name}}Struct(in {{.Name}}Input) ([]byte, error)
	{{- end}}

	{{- range $event := .Events}}
	{{.Normalized.Name}}LogHash() []byte
	Decode{{.Normalized.Name}}(log *evm.Log) (*{{.Normalized.Name}}, error)
	{{- end}}
}

type {{decapitalise $contract.Type}}CodecImpl struct {
	abi *abi.ABI
}

func New{{$contract.Type}}Codec() ({{$contract.Type}}Codec, error) {
	parsed, err := abi.JSON(strings.NewReader({{$contract.Type}}MetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &{{decapitalise $contract.Type}}CodecImpl{abi: &parsed}, nil
}

{{range $call := $contract.Calls}}
{{- if or $call.Original.Constant (eq $call.Original.StateMutability "view")}}
func (c *{{decapitalise $contract.Type}}CodecImpl) Encode{{$call.Normalized.Name}}MethodCall(in {{$call.Normalized.Name}}Input) ([]byte, error) {
	return c.abi.Pack("{{$call.Original.Name}}"{{range .Normalized.Inputs}}, in.{{capitalise .Name}}{{end}})
}
func (c *{{decapitalise $contract.Type}}CodecImpl) Decode{{$call.Normalized.Name}}MethodOutput(data []byte) ({{with index $call.Normalized.Outputs 0}}{{bindtype .Type $.Structs}}{{end}}, error) {
	vals, err := c.abi.Methods["{{$call.Original.Name}}"].Outputs.Unpack(data)
	if err != nil {
		return {{with index $call.Normalized.Outputs 0}}*new({{bindtype .Type $.Structs}}){{end}}, err
	}
	return vals[0].({{bindtype (index $call.Normalized.Outputs 0).Type $.Structs}}), nil
}
{{- end}}
{{end}}

{{range $.Structs}}
func (c *{{decapitalise $contract.Type}}CodecImpl) Encode{{.Name}}Struct(in {{.Name}}Input) ([]byte, error) {
	return c.abi.Pack("{{decapitalise .Name}}", in)
}
{{end}}

{{range $event := $contract.Events}}
func (c *{{decapitalise $contract.Type}}CodecImpl) {{.Normalized.Name}}LogHash() []byte {
	return c.abi.Events["{{.Original.Name}}"].ID.Bytes()
}

// Decode{{.Normalized.Name}} decodes a log into a {{.Normalized.Name}} struct.
func (c *{{decapitalise $contract.Type}}CodecImpl) Decode{{.Normalized.Name}}(log *evm.Log) (*{{.Normalized.Name}}, error) {
	event := new({{.Normalized.Name}})
	if err := c.abi.UnpackIntoInterface(event, "{{.Original.Name}}", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["{{.Original.Name}}"].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}
{{end}}

type {{$contract.Type}} struct {
	Address   []byte
	Options   *bindings.ContractInitOptions
	ABI       *abi.ABI
	evmClient evmcappb.Client
	codec     {{$contract.Type}}Codec
}

func New{{$contract.Type}}(
	client evmcappb.Client,
	address []byte,
	options *bindings.ContractInitOptions,
) (*{{$contract.Type}}, error) {
	parsed, err := abi.JSON(strings.NewReader({{$contract.Type}}MetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := New{{$contract.Type}}Codec()
	if err != nil {
		return nil, err
	}
	return &{{$contract.Type}}{
		Address:   address,
		Options:   options,
		ABI:       &parsed,
		evmClient: client,
		codec:     codec,
	}, nil
}

{{range $call := $contract.Calls}}
{{- if or $call.Original.Constant (eq $call.Original.StateMutability "view")}}

type {{$call.Normalized.Name}}Input struct {
	{{- range $param := $call.Normalized.Inputs}}
	{{capitalise $param.Name}} {{bindtype .Type $.Structs}}
	{{- end}}
}

func (c {{$contract.Type}}) {{$call.Normalized.Name}}(
	runtime sdk.Runtime,
	args {{$call.Normalized.Name}}Input,
	options *bindings.ReadOptions,
) ({{with index $call.Normalized.Outputs 0}}{{bindtype .Type $.Structs}}{{end}}, error) {
	calldata, err := c.codec.Encode{{$call.Normalized.Name}}MethodCall(args)
	if err != nil {
		return {{with index $call.Normalized.Outputs 0}}*new({{bindtype .Type $.Structs}}){{end}}, err
	}
	if options == nil {
		options = &bindings.ReadOptions{BlockNumber: nil}
	}
	promise := c.evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call:        &evm.CallMsg{To: c.Address, Data: calldata},
		BlockNumber: toPbBigInt(options.BlockNumber),
	})
	reply, err := promise.Await()
	if err != nil {
		return {{with index $call.Normalized.Outputs 0}}*new({{bindtype .Type $.Structs}}){{end}}, err
	}
	return c.codec.Decode{{$call.Normalized.Name}}MethodOutput(reply.Data)
}
{{- end}}
{{end}}

{{range $.Structs}}

type {{.Name}}Input struct {
	{{- range .Fields}}
	{{capitalise .Name}} {{.Type}}
	{{- end}}
}

func (c {{$contract.Type}}) WriteReport{{.Name}}(
	runtime sdk.Runtime,
	input {{.Name}}Input,
	gasConfig *evmcappb.GasConfig,
) (*big.Int, error) {
	encoded, err := c.codec.Encode{{.Name}}Struct(input)
	if err != nil {
		return nil, err
	}
	report := bindings.GenerateReport(getChainID(c.evmClient), encoded)
	writeReportReplyPromise := c.evmClient.WriteReport(runtime, &evmcappb.WriteReportRequest{
		Receiver: c.Address,
		Report: &evmcappb.SignedReport{
			RawReport:     report.RawReport,
			ReportContext: report.ReportContext,
			Signatures:    report.Signatures,
			Id:            report.ID,
		},
		GasConfig: gasConfig,
	})
	reply, err := writeReportReplyPromise.Await()
	if err != nil {
		return nil, err
	}
	if reply.TxStatus == evm.TxStatus_TX_FATAL {
		return nil, &bindings.TxFatalError{
			Message: "Fatal tx execution",
		}
	}
	for {
		txByHashPromise := c.evmClient.GetTransactionByHash(runtime, &evm.GetTransactionByHashRequest{
			Hash: reply.TxHash,
		})
		getTxResult, err := txByHashPromise.Await()
		if err != nil {
			return nil, err
		}
		if getTxResult.Transaction.IsFinalized {
			if *reply.ReceiverContractExecutionStatus == evmcappb.ReceiverContractExecutionStatus_REVERTED {
				return nil, &bindings.ReceiverContractError{
					Message: "Transaction finalized but receiver contract failed to execute",
					TxHash:  reply.TxHash,
				}
			}
			return reply.TxHash, nil
		}
	}
}
{{end}}

{{range $error := $contract.Errors}}

// {{.Normalized.Name}} represents the {{.Original.Name}} error raised by the {{$contract.Type}} contract.
type {{.Normalized.Name}} struct {
	{{- range .Normalized.Inputs}}
	{{capitalise .Name}} {{bindtype .Type $.Structs}}
	{{- end}}
}

// TODO: possibly clean this up
// Decode{{.Normalized.Name}}Error decodes a {{.Original.Name}} error from revert data.
func (c *{{$contract.Type}}) Decode{{.Normalized.Name}}Error(data []byte) (*{{.Normalized.Name}}, error) {
	args := c.ABI.Errors["{{.Original.Name}}"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != {{len .Normalized.Inputs}} {
		return nil, fmt.Errorf("expected {{len .Normalized.Inputs}} values, got %%d", len(values))
	}

	{{$err := .}} {{/* capture outer context */}}

	{{range $i, $param := $err.Normalized.Inputs}}
	{{$param.Name}}, ok{{$i}} := values[{{$i}}].({{bindtype $param.Type $.Structs}})
	if !ok{{$i}} {
		return nil, fmt.Errorf("unexpected type for {{$param.Name}} in {{$err.Normalized.Name}} error")
	}
	{{end}}

	return &{{$err.Normalized.Name}}{
		{{- range $i, $param := $err.Normalized.Inputs}}
		{{capitalise $param.Name}}: {{$param.Name}},
		{{- end}}
	}, nil
}


func (c *{{$contract.Type}}) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	{{range $error := $contract.Errors}}
	case common.Bytes2Hex(c.ABI.Errors["{{$error.Original.Name}}"].ID.Bytes()):
		return c.Decode{{$error.Normalized.Name}}Error(data)
	{{end}}
	default:
		return nil, errors.New("unknown error selector")
	}
}

// Error implements the error interface for {{.Normalized.Name}}.
func (e *{{.Normalized.Name}}) Error() string {
	return fmt.Sprintf("{{.Normalized.Name}} error:{{range .Normalized.Inputs}} {{.Name}}=%%v;{{end}}"{{range .Normalized.Inputs}}, e.{{capitalise .Name}}{{end}})
}

{{end}}


{{range $event := $contract.Events}}

// {{.Normalized.Name}} represents a {{.Original.Name}} event raised by the {{$contract.Type}} contract.
type {{.Normalized.Name}} struct {
	{{- range .Normalized.Inputs}}
	{{capitalise .Name}} {{if .Indexed}}{{bindtopictype .Type $.Structs}}{{else}}{{bindtype .Type $.Structs}}{{end}}
	{{- end}}
}

func (c *{{$contract.Type}}) RegisterLogTracking{{.Normalized.Name}}(runtime sdk.Runtime, options *bindings.LogTrackingOptions) {
	//TODO use log tracking options if set
	c.evmClient.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:      "{{.Normalized.Name}}-" + common.Bytes2Hex(c.Address),
			Addresses: [][]byte{c.Address},
			EventSigs: [][]byte{c.codec.{{.Normalized.Name}}LogHash()},
		},
	})
}

func (c *{{$contract.Type}}) UnregisterLogTracking{{.Normalized.Name}}(runtime sdk.Runtime) {
	c.evmClient.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "{{.Normalized.Name}}-" + common.Bytes2Hex(c.Address),
	})
}

func (c *{{$contract.Type}}) QueryTrackedLogs{{.Normalized.Name}}(runtime sdk.Runtime, options *bindings.QueryTrackedLogsOptions) ([]bindings.ParsedLog[{{.Normalized.Name}}], any) {
	promise := c.evmClient.QueryTrackedLogs(runtime, &evm.QueryTrackedLogsRequest{
		Expression: []*evm.Expression{
			//TODO add proper expression
			&evm.Expression{Evaluator: &evm.Expression_BooleanExpression{&evm.BooleanExpression{
				Expression: []*evm.Expression{},
			}}},
		},
	})
	reply, err := promise.Await()
	if err != nil {
		return nil, fmt.Errorf("failed to query tracked logs: %w", err)
	}
	logs := reply.Logs
	parsedLogs := make([]bindings.ParsedLog[{{.Normalized.Name}}], len(logs))
	for i, log := range logs {
		decodedLog, err := c.codec.Decode{{.Normalized.Name}}(log)
		if err != nil {
			return nil, fmt.Errorf("failed to decode {{.Normalized.Name}} log: %w", err)
		}
		parsedLogs[i] = bindings.ParsedLog[{{.Normalized.Name}}]{
			LogData: *decodedLog,
			RawLog: log,
		}
	}
	
	return parsedLogs, nil
}

func (c *{{$contract.Type}}) FilterLogs{{.Normalized.Name}}(runtime sdk.Runtime, options *bindings.FilterOptions) ([]bindings.ParsedLog[{{.Normalized.Name}}], error) {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: "finalized", //TODO we need a enum / constant
		}
	}
	filterLogsReplyPromise := c.evmClient.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address},
			Topics:    []*evm.Topics{
				{Topic:[][]byte{c.codec.{{.Normalized.Name}}LogHash()}},
			},			BlockHash: options.BlockHash,
			FromBlock: toPbBigInt(options.FromBlock),
			ToBlock:   toPbBigInt(options.ToBlock),
		},
	})
	reply, err := filterLogsReplyPromise.Await()
	if err != nil {
		return nil, err
	}
	logs := reply.Logs
	parsedLogs := make([]bindings.ParsedLog[{{.Normalized.Name}}], len(logs))
	for i, log := range logs {
		decodedLog, err := c.codec.Decode{{.Normalized.Name}}(log)
		if err != nil {
			return nil, fmt.Errorf("failed to decode {{.Normalized.Name}} log: %w", err)
		}
		parsedLogs[i] = bindings.ParsedLog[{{.Normalized.Name}}]{
			LogData: *decodedLog,
			RawLog: log,
		}
	}
	
	return parsedLogs, nil
}
{{end}}

{{end}}

func toPbBigInt(i *big.Int) *pb.BigInt   { panic("unimplemented") }
func getChainID(e evmcappb.Client) uint32 { panic("unimplemented") }
