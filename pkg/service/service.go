package service

import (
	"context"
	"fmt"
	"github.com/smartcontractkit/chainlink-evm/pkg/logpoller"
	"math/big"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	evmcap "github.com/smartcontractkit/chainlink-common/pkg/loop/chain-capabilities/evm"
	evmcapserver "github.com/smartcontractkit/chainlink-common/pkg/loop/chain-capabilities/evm/server"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	evmservicetypes "github.com/smartcontractkit/chainlink-common/pkg/types/chains/evm"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
)

const (
	CapabilityName = "evm"
)

type evmCapability struct {
	capabilities.CapabilityInfo
	types.EVMService
	LP logpoller.LogPoller
	// TODO ?
	//evmcapserver.UnimplementedEVMCapabilityServer
	// TODO
	//config Config
}

func (e evmCapability) RegisterLogTrigger(ctx context.Context, metadata capabilities.RequestMetadata, input *evmcap.FilterLogsRequest) (<-chan capabilities.TriggerAndId[*evmcap.FilterLogsReply], error) {

	addresses := input.GetFilterQuery().GetAddresses()
	if len(addresses) == 0 {
		return nil, fmt.Errorf("no addresses provided")
	}
	panic("implement me")
}

func (e evmCapability) UnregisterLogTrigger(ctx context.Context, metadata capabilities.RequestMetadata, input *evmcap.FilterLogsRequest) error {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) CallContract(ctx context.Context, _ capabilities.RequestMetadata, input *evmcap.CallContractRequest) (*evmcap.CallContractReply, error) {
	callMsg, err := parseCallMsg(input.Call)
	if err != nil {
		return &evmcap.CallContractReply{}, err
	}

	blockNumber := &big.Int{}
	blockNumber, isOK := blockNumber.SetString(input.BlockNumber.String(), 10)
	if !isOK {
		return &evmcap.CallContractReply{}, fmt.Errorf("failed to parse block number: %s", input.BlockNumber.String())
	}

	response, err := e.EVMService.CallContract(ctx, &callMsg, blockNumber)
	if err != nil {
		return &evmcap.CallContractReply{}, err
	}

	return &evmcap.CallContractReply{Data: &evmcap.ABIPayload{Abi: response}}, nil
}

func (e evmCapability) FilterLogs(ctx context.Context, _ capabilities.RequestMetadata, input *evmcap.FilterLogsRequest) (*evmcap.FilterLogsReply, error) {
	evmCapFq := input.GetFilterQuery()

	fromBlock := &big.Int{}
	fromBlock, isOK := fromBlock.SetString(evmCapFq.FromBlock.String(), 10)
	if !isOK {
		return nil, fmt.Errorf("failed to parse from block number: %s", evmCapFq.FromBlock.String())
	}

	toBlock := &big.Int{}
	toBlock, isOK = toBlock.SetString(evmCapFq.ToBlock.String(), 10)
	if !isOK {
		return nil, fmt.Errorf("failed to parse to block number: %s", evmCapFq.ToBlock.String())
	}

	fq := evmservicetypes.FilterQuery{
		BlockHash: evmservicetypes.Hash(evmCapFq.BlockHash.Hash),
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	}

	for _, address := range evmCapFq.Addresses {
		fq.Addresses = append(fq.Addresses, evmservicetypes.Address(address.Address))
	}

	// TODO fix topic type
	//for _, topic := range evmCapFq.Topics {
	//	fq.Addresses = append(fq.Addresses, evmservicetypes.Address(topic.Topic))
	//}

	logs, err := e.EVMService.FilterLogs(ctx, fq)
	if err != nil {
		return nil, err
	}

	return &evmcap.FilterLogsReply{Logs: parseLogs(logs)}, nil
}

func (e evmCapability) BalanceAt(ctx context.Context, _ capabilities.RequestMetadata, input *evmcap.BalanceAtRequest) (*evmcap.BalanceAtReply, error) {

	blockNumber := &big.Int{}
	blockNumber, isOK := blockNumber.SetString(input.BlockNumber.String(), 10)
	if !isOK {
		return nil, fmt.Errorf("failed to parse to block number: %s", input.BlockNumber.String())
	}

	balance, err := e.EVMService.BalanceAt(ctx, evmservicetypes.Address(input.Account.Address), blockNumber)
	if err != nil {
		return nil, err
	}

	return &evmcap.BalanceAtReply{Balance: &pb.BigInt{AbsVal: balance.Bytes(), Sign: int64(balance.Sign())}}, nil
}

func (e evmCapability) EstimateGas(ctx context.Context, _ capabilities.RequestMetadata, input *evmcap.EstimateGasRequest) (*evmcap.EstimateGasReply, error) {
	callMsg, err := parseCallMsg(input.Msg)
	if err != nil {
		return &evmcap.EstimateGasReply{}, err
	}

	estimate, err := e.EVMService.EstimateGas(ctx, &callMsg)
	if err != nil {
		return &evmcap.EstimateGasReply{}, err
	}

	return &evmcap.EstimateGasReply{Gas: estimate}, nil
}

func (e evmCapability) GetTransactionByHash(ctx context.Context, _ capabilities.RequestMetadata, input *evmcap.GetTransactionByHashRequest) (*evmcap.GetTransactionByHashReply, error) {
	tx, err := e.EVMService.TransactionByHash(ctx, evmservicetypes.Hash(input.Hash.Hash))
	if err != nil {
		return nil, err
	}

	return &evmcap.GetTransactionByHashReply{Transaction: &evmcap.Transaction{
		Nonce:    tx.Nonce,
		Gas:      tx.Gas,
		To:       &evmcap.Address{Address: tx.To[:]},
		Data:     &evmcap.ABIPayload{Abi: tx.Data},
		Value:    &pb.BigInt{AbsVal: tx.Value.Bytes(), Sign: int64(tx.Value.Sign())},
		GasPrice: &pb.BigInt{AbsVal: tx.GasPrice.Bytes(), Sign: int64(tx.GasPrice.Sign())},
		Hash:     &evmcap.Hash{Hash: tx.Hash[:]},
	}}, nil
}

func (e evmCapability) GetTransactionReceipt(ctx context.Context, _ capabilities.RequestMetadata, input *evmcap.GetReceiptRequest) (*evmcap.GetReceiptReply, error) {
	receipt, err := e.EVMService.TransactionReceipt(ctx, evmservicetypes.Hash(input.Hash.Hash))
	if err != nil {
		return nil, err
	}

	reply := &evmcap.GetReceiptReply{Receipt: &evmcap.Receipt{
		Status:          receipt.Status,
		Logs:            parseLogs(receipt.Logs),
		TxHash:          &evmcap.Hash{Hash: receipt.TxHash[:]},
		ContractAddress: &evmcap.Address{Address: receipt.ContractAddress[:]},
		GasUsed:         receipt.GasUsed,
		BlockHash:       &evmcap.Hash{Hash: receipt.BlockHash[:]},
		BlockNumber:     &pb.BigInt{AbsVal: receipt.BlockNumber.Bytes(), Sign: int64(receipt.BlockNumber.Sign())},
		// TODO add tx index
		//TransactionIndex:  tx.TransactionIndex,
		EffectiveGasPrice: &pb.BigInt{AbsVal: receipt.EffectiveGasPrice.Bytes(), Sign: int64(receipt.EffectiveGasPrice.Sign())},
	}}

	return reply, nil
}

func (e evmCapability) LatestAndFinalizedHead(ctx context.Context, _ capabilities.RequestMetadata, _ *emptypb.Empty) (*evmcap.LatestAndFinalizedHeadReply, error) {
	latest, finalized, err := e.EVMService.LatestAndFinalizedHead(ctx)
	if err != nil {
		return nil, err
	}

	return &evmcap.LatestAndFinalizedHeadReply{
		Latest: &evmcap.Head{
			Timestamp:   latest.Timestamp,
			BlockNumber: &pb.BigInt{AbsVal: latest.Number.Bytes(), Sign: int64(latest.Number.Sign())},
			Hash:        &evmcap.Hash{Hash: latest.Hash[:]},
			ParentHash:  &evmcap.Hash{Hash: latest.ParentHash[:]},
		},
		Finalized: &evmcap.Head{
			Timestamp:   finalized.Timestamp,
			BlockNumber: &pb.BigInt{AbsVal: finalized.Number.Bytes(), Sign: int64(finalized.Number.Sign())},
			Hash:        &evmcap.Hash{Hash: finalized.Hash[:]},
			ParentHash:  &evmcap.Hash{Hash: finalized.ParentHash[:]},
		},
	}, nil
}

func (e evmCapability) QueryTrackedLogs(_ context.Context, _ capabilities.RequestMetadata, _ *evmcap.QueryTrackedLogsRequest) (*evmcap.QueryTrackedLogsReply, error) {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) RegisterLogTracking(_ context.Context, _ capabilities.RequestMetadata, _ *evmcap.RegisterLogTrackingRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) UnregisterLogTracking(_ context.Context, _ capabilities.RequestMetadata, _ *evmcap.UnregisterLogTrackingRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) RegisterTrigger(_ context.Context, _ capabilities.TriggerRegistrationRequest) (<-chan capabilities.TriggerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) UnregisterTrigger(_ context.Context, _ capabilities.TriggerRegistrationRequest) error {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) RegisterToWorkflow(_ context.Context, _ capabilities.RegisterToWorkflowRequest) error {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) UnregisterFromWorkflow(_ context.Context, _ capabilities.UnregisterFromWorkflowRequest) error {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) Execute(_ context.Context, _ capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e evmCapability) Start(_ context.Context) error {
	//TODO Relayer already started the EVM Relayer, this is not needed atm afaik
	return nil
}

func (e evmCapability) Close() error {
	//TODO Relayer closes the EVM Relayer, this is not needed atm afaik
	return nil
}

func (e evmCapability) HealthReport() map[string]error {
	// TODO
	return nil
}

func (e evmCapability) Name() string {
	return CapabilityName
}

func (e evmCapability) Description() string {
	return "Contains EVM chain functionalities"
}

func (e evmCapability) Ready() error {
	// TODO
	return nil
}

func (e evmCapability) Initialise(_ context.Context, _ string, _ core.TelemetryService, _ core.KeyValueStore, _ core.ErrorLog, _ core.PipelineRunnerService, _ core.RelayerSet, _ core.OracleFactory) error {
	// TODO no-op for now, figure out if we want to run this as a standard capability
	return nil
}

type EVMCapabilityOpts struct {
	ID string
	types.EVMService
	LP logpoller.LogPoller
	// TODO a lot of things like config are handled by the Relayer indirectly, needs another passover for these detailks
}

func NewEVMCapablity(opts EVMCapabilityOpts) capabilities.ExecutableAndTriggerCapability {
	capInfo := capabilities.MustNewCapabilityInfo(opts.ID, capabilities.CapabilityTypeTarget, CapabilityName)

	//log poller should reach here?
	return evmCapability{CapabilityInfo: capInfo, EVMService: opts.EVMService, LP: opts.LP}
}

var _ evmcapserver.EVMChainCapability = &evmCapability{}

func parseCallMsg(callMsg *evmcap.CallMsg) (evmservicetypes.CallMsg, error) {
	parsed := evmservicetypes.CallMsg{
		To:   evmservicetypes.Address(callMsg.To.Address),
		From: evmservicetypes.Address(callMsg.From.Address),
		Data: callMsg.Data.Abi,
	}

	return parsed, nil
}

func parseLogs(logs []*evmservicetypes.Log) []*evmcap.Log {
	parsed := make([]*evmcap.Log, 0)
	for _, log := range logs {
		parsed = append(parsed, &evmcap.Log{
			Address:     &evmcap.Address{Address: log.Address[:]},
			Topics:      make([]*evmcap.Hash, 0),
			TxHash:      &evmcap.Hash{Hash: log.TxHash[:]},
			BlockHash:   &evmcap.Hash{Hash: log.BlockHash[:]},
			Data:        &evmcap.ABIPayload{Abi: log.Data},
			EventSig:    &evmcap.Hash{Hash: log.EventSig[:]},
			BlockNumber: &pb.BigInt{AbsVal: log.BlockNumber.Bytes(), Sign: int64(log.BlockNumber.Sign())},
			// TODO missing tx index from evm log
			//TxIndex: log.TxIndex,
			Index:   log.LogIndex,
			Removed: log.Removed,
		})
	}

	return parsed
}
