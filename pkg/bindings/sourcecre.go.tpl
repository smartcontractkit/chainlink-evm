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
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk/v2/pb"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
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
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
)

{{range $contract := .Contracts}}
var {{$contract.Type}}MetaData = &bind.MetaData{
	ABI: "{{.InputABI}}",
	{{- if .InputBin}}
	Bin: "0x{{.InputBin}}",
	{{- end}}
}

// Structs 
{{range $.Structs}}type {{.Name}} struct {
	{{- range .Fields}}
	{{capitalise .Name}} {{.Type}}
	{{- end}}
}

{{end}}

// Contract Method Inputs{{- range $call := $contract.Calls}}
{{- if gt (len $call.Normalized.Inputs) 0 }}
type {{$call.Normalized.Name}}Input struct {
	{{- range $param := $call.Normalized.Inputs}}
	{{capitalise $param.Name}} {{bindtype .Type $.Structs}}
	{{- end}}
}
{{end}}

{{- end}}

// Errors
{{range $error := $contract.Errors}}type {{.Normalized.Name}} struct {
	{{- range .Normalized.Inputs}}
	{{capitalise .Name}} {{bindtype .Type $.Structs}}
	{{- end}}
}

{{end}}

// Events
{{range $event := $contract.Events}}type {{.Normalized.Name}} struct {
	{{- range .Normalized.Inputs}}
	{{capitalise .Name}} {{if .Indexed}}{{bindtopictype .Type $.Structs}}{{else}}{{bindtype .Type $.Structs}}{{end}}
	{{- end}}
}

{{end}}

// Main Binding Type for {{$contract.Type}}
type {{$contract.Type}} struct {
	Address   []byte
	Options   *bindings.ContractInitOptions
	ABI       *abi.ABI
	client *evm.Client
	Codec     {{$contract.Type}}Codec
}

type {{$contract.Type}}Codec interface {
	{{- range $call := $contract.Calls}}
	
	{{- if gt (len $call.Normalized.Inputs) 0 }}
	Encode{{$call.Normalized.Name}}MethodCall(in {{$call.Normalized.Name}}Input) ([]byte, error)
	{{- else }}
	Encode{{$call.Normalized.Name}}MethodCall() ([]byte, error)
	{{- end }}
	{{- if gt (len $call.Normalized.Outputs) 0 }}
	Decode{{$call.Normalized.Name}}MethodOutput(data []byte) ({{with index $call.Normalized.Outputs 0}}{{bindtype .Type $.Structs}}{{end}}, error)
	{{- end }}
	
	{{- end}}

	{{- range $.Structs}}
	Encode{{.Name}}Struct(in {{.Name}}) ([]byte, error)
	{{- end}}

	{{- range $event := .Events}}
	{{.Normalized.Name}}LogHash() []byte
	Encode{{.Normalized.Name}}Topics(evt abi.Event, values []{{.Normalized.Name}}) ([]*evm.TopicValues, error)
	Decode{{.Normalized.Name}}(log *evm.Log) (*{{.Normalized.Name}}, error)
	{{- end}}
}

func New{{$contract.Type}}(
	client *evm.Client,
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
		client: client,
		Codec:     codec,
	}, nil
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

{{- if gt (len $call.Normalized.Inputs) 0 }}

func (c *{{ decapitalise $contract.Type }}CodecImpl) Encode{{ $call.Normalized.Name }}MethodCall(in {{ $call.Normalized.Name }}Input) ([]byte, error) {
	return c.abi.Pack("{{ $call.Original.Name }}"{{- range .Normalized.Inputs }}, in.{{ capitalise .Name }}{{- end }})
}
{{- else }}
func (c *{{ decapitalise $contract.Type }}CodecImpl) Encode{{ $call.Normalized.Name }}MethodCall() ([]byte, error) {
	return c.abi.Pack("{{ $call.Original.Name }}")
}

{{- end }}

{{- if gt (len $call.Normalized.Outputs) 0 }}

func (c *{{ decapitalise $contract.Type }}CodecImpl) Decode{{ $call.Normalized.Name }}MethodOutput(data []byte) ({{ with index $call.Normalized.Outputs 0 }}{{ bindtype .Type $.Structs }}{{ end }}, error) {
	vals, err := c.abi.Methods["{{ $call.Original.Name }}"].Outputs.Unpack(data)
	if err != nil {
		return {{ with index $call.Normalized.Outputs 0 }}*new({{ bindtype .Type $.Structs }}){{ end }}, err
	}
	return vals[0].({{ bindtype (index $call.Normalized.Outputs 0).Type $.Structs }}), nil
}
{{- end }}

{{end}}

{{range $.Structs}}
func (c *{{decapitalise $contract.Type}}CodecImpl) Encode{{.Name}}Struct(in {{.Name}}) ([]byte, error) {
	tupleType, err := abi.NewType(
        "tuple", "",
        []abi.ArgumentMarshaling{
			{{range $f := .Fields}}{Name: "{{ decapitalise $f.Name }}", Type: "{{ $f.SolKind }}"},
			{{end}}
        },
    )
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for {{.Name}}: %w", err)
	}
	args := abi.Arguments{
        {Name: "{{ decapitalise .Name }}", Type: tupleType},
    }

	return args.Pack(in)
}
{{- end }}

{{range $event := $contract.Events}}
func (c *{{decapitalise $contract.Type}}CodecImpl) {{.Normalized.Name}}LogHash() []byte {
	return c.abi.Events["{{.Original.Name}}"].ID.Bytes()
}

func (c *{{decapitalise $contract.Type}}CodecImpl) Encode{{.Normalized.Name}}Topics(
    evt abi.Event,
    values []{{.Normalized.Name}},
) ([]*evm.TopicValues, error) {
    {{- range $idx, $inp := .Normalized.Inputs }}
    {{- if $inp.Indexed }}
    var {{ decapitalise $inp.Name }}Rule []interface{}
    for _, v := range values {
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[{{$idx}}], v.{{capitalise $inp.Name}})
		if err != nil {
			return nil, err
		}
		{{ decapitalise $inp.Name }}Rule = append({{ decapitalise $inp.Name }}Rule, fieldVal)
	}
    {{- end }}
    {{- end }}

    rawTopics, err := abi.MakeTopics(
        {{- range $inp := .Normalized.Inputs }}
        {{- if $inp.Indexed }}
        {{ decapitalise $inp.Name }}Rule,
        {{- end }}
        {{- end }}
    )
    if err != nil {
        return nil, err
    }

	topics := make([]*evm.TopicValues, len(rawTopics)+1)
	topics[0] = &evm.TopicValues{
		Values: [][]byte{evt.ID.Bytes()},
	}
    for i, hashList := range rawTopics {
        bs := make([][]byte, len(hashList))
        for j, h := range hashList {
            bs[j] = h.Bytes()
        }
        topics[i+1] = &evm.TopicValues{Values: bs}
    }
    return topics, nil
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

{{range $call := $contract.Calls}}
  {{- if or $call.Original.Constant (eq $call.Original.StateMutability "view")}}

func (c {{$contract.Type}}) {{$call.Normalized.Name}}(
    runtime cre.Runtime,
    {{- if gt (len $call.Normalized.Inputs) 0}}
    args {{$call.Normalized.Name}}Input,
    {{- end}}
    blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
    {{- if gt (len $call.Normalized.Inputs) 0}}
    calldata, err := c.Codec.Encode{{$call.Normalized.Name}}MethodCall(args)
	{{- else }}
	calldata, err := c.Codec.Encode{{$call.Normalized.Name}}MethodCall()
	{{- end}}
    if err != nil {
        return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
    }

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(-3)), // -3 means latest finalized block
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
            if finalizedBlock == nil || finalizedBlock.Header == nil {
                return nil, errors.New("failed to get finalized block header")
            }
            return finalizedBlock.Header.BlockNumber, nil
        })
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

    return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
        return c.client.CallContract(runtime, &evm.CallContractRequest{
            Call:        &evm.CallMsg{To: c.Address, Data: calldata},
            BlockNumber: bn,
        })
    })

}
  {{- end}}
{{end}}

{{range $.Structs}}

func (c {{$contract.Type}}) WriteReport{{.Name}}(
	runtime cre.Runtime,
	input {{.Name}},
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.Encode{{.Name}}Struct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *pb2.ReportResponse) cre.Promise[*evm.WriteReportReply] {
	    return c.client.WriteReport(runtime, &evm.WriteReportRequest{
    		Receiver: c.Address,
    		Report: report,
    		GasConfig: gasConfig,
    	})
	})
}
{{end}}

{{range $error := $contract.Errors}}

// Decode{{.Normalized.Name}}Error decodes a {{.Original.Name}} error from revert data.
func (c *{{$contract.Type}}) Decode{{.Normalized.Name}}Error(data []byte) (*{{.Normalized.Name}}, error) {
	args := c.ABI.Errors["{{.Original.Name}}"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != {{len .Normalized.Inputs}} {
		return nil, fmt.Errorf("expected {{len .Normalized.Inputs}} values, got %d", len(values))
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

// Error implements the error interface for {{.Normalized.Name}}.
func (e *{{.Normalized.Name}}) Error() string {
	return fmt.Sprintf("{{.Normalized.Name}} error:{{range .Normalized.Inputs}} {{.Name}}=%v;{{end}}"{{range .Normalized.Inputs}}, e.{{capitalise .Name}}{{end}})
}

{{end}}

func (c *{{$contract.Type}}) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	{{range $error := $contract.Errors}}case common.Bytes2Hex(c.ABI.Errors["{{$error.Original.Name}}"].ID.Bytes()[:4]):
		return c.Decode{{$error.Normalized.Name}}Error(data)
	{{end}}default:
		return nil, errors.New("unknown error selector")
	}
}

{{range $event := $contract.Events}}

func (c *{{$contract.Type}}) LogTrigger{{.Normalized.Name}}Log(chainSelector uint64, confidence evm.ConfidenceLevel, filters []{{.Normalized.Name}}) (cre.Trigger[*evm.Log, *evm.Log], error) {
	event := c.ABI.Events["{{.Normalized.Name}}"]
	topics, err := c.Codec.Encode{{.Normalized.Name}}Topics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for {{.Normalized.Name}}: %w", err)
	}

	return evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address},
		Topics:     topics,
		Confidence: confidence,
	}), nil
}

func (c *{{$contract.Type}}) RegisterLogTracking{{.Normalized.Name}}(runtime cre.Runtime, options *bindings.LogTrackingOptions[{{.Normalized.Name}}]) cre.Promise[*emptypb.Empty] {
	bindings.ValidateLogTrackingOptions[{{.Normalized.Name}}](options)
	topics, err := c.Codec.Encode{{.Normalized.Name}}Topics(c.ABI.Events["{{.Normalized.Name}}"], options.Filters)
	if err != nil {
		return cre.PromiseFromResult[*emptypb.Empty](nil, fmt.Errorf("failed to encode topics for {{.Normalized.Name}}: %w", err))
	}
	padded := bindings.PadTopics(topics)
	return c.client.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "{{.Normalized.Name}}-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.{{.Normalized.Name}}LogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        padded[1].Values,
			Topic3:        padded[2].Values,
			Topic4:        padded[3].Values,
		},
	})
}

func (c *{{$contract.Type}}) UnregisterLogTracking{{.Normalized.Name}}(runtime cre.Runtime) cre.Promise[*emptypb.Empty] {
	return c.client.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "{{.Normalized.Name}}-" + common.Bytes2Hex(c.Address),
	})
}

func (c *{{$contract.Type}}) FilterLogs{{.Normalized.Name}}(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address},
			Topics:    []*evm.Topics{
				{Topic:[][]byte{c.Codec.{{.Normalized.Name}}LogHash()}},
			},			
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	})
}
{{end}}

{{end}}
