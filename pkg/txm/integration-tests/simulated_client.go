package integrationtests

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers/dualbroadcast"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type SimulationMode string

const (
	Standard         SimulationMode = "standard"
	Retransmission   SimulationMode = "retransmission"
	StuckTxDetection SimulationMode = "stuck_tx_detection"
	ErrorHandling    SimulationMode = "error_handling"
)

type SimulatedClient interface {
	ChainID(ctx context.Context) (*big.Int, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	CallContext(ctx context.Context, result interface{}, method string, args ...any) error
	HeadByNumber(ctx context.Context, n *big.Int) (*evmtypes.Head, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (feeHistory *ethereum.FeeHistory, err error)
	Commit()
}

type gethSimulatedClient struct {
	*clientwrappers.GethClient
	Mode SimulationMode
}

func NewGethSimulatedClient(client *clientwrappers.GethClient, mode SimulationMode) SimulatedClient {
	return &gethSimulatedClient{
		GethClient: client,
		Mode:       mode,
	}
}

func (s *gethSimulatedClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	switch s.Mode {
	case ErrorHandling:
		// ErrorHandling: Force the TXM to inject transactions with certain error messages so we can test error handling.
		if tx.ID%2 == 0 && attempt.ID%2 == 0 {
			return dualbroadcast.ErrNoBids
		}
		return s.GethClient.SendTransaction(ctx, tx, attempt)
	case StuckTxDetection:
		// StuckTxDetection: Force the TXM to assume some of the attempts were successful while they weren't so we can test stuck tx detection.
		// We fail on every other tx and attempt to inject a more realistic mix of transactions for the TXM to handle.
		if tx.ID%2 == 0 && attempt.ID%2 == 0 {
			return nil
		}
		return s.GethClient.SendTransaction(ctx, tx, attempt)
	case Retransmission:
		// Retransmission: Force the TXM to assume some of the attempts were successful while they weren't so we can test retransmission.
		if attempt.ID%2 == 0 {
			return nil
		}
		return s.GethClient.SendTransaction(ctx, tx, attempt)
	case Standard:
		return s.GethClient.SendTransaction(ctx, tx, attempt)
	default:
		return s.GethClient.SendTransaction(ctx, tx, attempt)
	}
}

func (s *gethSimulatedClient) Commit() {}

type backendSimulatedClient struct {
	*simulated.Backend
	Mode SimulationMode
}

func NewBackendSimulatedClient(client *simulated.Backend, mode SimulationMode) SimulatedClient {
	return &backendSimulatedClient{
		Backend: client,
		Mode:    mode,
	}
}

func (s *backendSimulatedClient) ChainID(ctx context.Context) (*big.Int, error) {
	return s.Backend.Client().ChainID(ctx)
}

func (s *backendSimulatedClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return s.Backend.Client().NonceAt(ctx, account, blockNumber)
}

func (s *backendSimulatedClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return s.Backend.Client().PendingNonceAt(ctx, account)
}

func (s *backendSimulatedClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	switch s.Mode {
	case ErrorHandling:
		// ErrorHandling: Force the TXM to inject transactions with certain error messages so we can test error handling.
		if tx.ID%2 == 0 && attempt.ID%2 == 0 {
			return dualbroadcast.ErrNoBids
		}
		return s.Backend.Client().SendTransaction(ctx, attempt.SignedTransaction)
	case StuckTxDetection:
		// StuckTxDetection: Force the TXM to assume some of the attempts were successful while they weren't so we can test stuck tx detection.
		// We fail on every other tx and attempt to inject a more realistic mix of transactions for the TXM to handle.
		if tx.ID%2 == 0 && attempt.ID%2 == 0 {
			return nil
		}
		return s.Backend.Client().SendTransaction(ctx, attempt.SignedTransaction)
	case Retransmission:
		// Retransmission: Force the TXM to assume some of the attempts were successful while they weren't so we can test retransmission.
		if attempt.ID%2 == 0 {
			return nil
		}
		return s.Backend.Client().SendTransaction(ctx, attempt.SignedTransaction)
	case Standard:
		return s.Backend.Client().SendTransaction(ctx, attempt.SignedTransaction)
	default:
		return s.Backend.Client().SendTransaction(ctx, attempt.SignedTransaction)
	}
}

func (s *backendSimulatedClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return s.Backend.Client().CallContract(ctx, msg, blockNumber)
}

func (s *backendSimulatedClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	panic("not implemented")
}

func (s *backendSimulatedClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	panic("not implemented")
}

func (s *backendSimulatedClient) HeadByNumber(ctx context.Context, n *big.Int) (*evmtypes.Head, error) {
	header, err := s.Backend.Client().HeaderByNumber(ctx, n)
	if err != nil {
		return nil, err
	}
	if header == nil {
		return nil, ethereum.NotFound
	}
	chainID, err := s.Backend.Client().ChainID(ctx)
	if err != nil {
		return nil, err
	}
	head := &evmtypes.Head{EVMChainID: sqlutil.New(chainID)}
	head.SetFromHeader(header)
	return head, nil
}

func (s *backendSimulatedClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return s.Backend.Client().EstimateGas(ctx, call)
}

func (s *backendSimulatedClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return s.Backend.Client().SuggestGasPrice(ctx)
}

func (s *backendSimulatedClient) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	return s.Backend.Client().FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
}

func (s *backendSimulatedClient) Commit() {
	s.Backend.Commit()
}
