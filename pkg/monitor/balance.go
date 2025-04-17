package monitor

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	pkgerrors "github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-framework/chains/heads"
	"github.com/smartcontractkit/chainlink-framework/metrics"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type (
	HeadTrackable = heads.Trackable[*evmtypes.Head, common.Hash]
	// BalanceMonitor checks the balance for each key on every new head
	BalanceMonitor interface {
		HeadTrackable
		GetEthBalance(common.Address) *assets.Eth
		services.Service
	}

	balanceMonitor struct {
		services.Service
		eng *services.Engine

		ethClient      evmclient.Client
		chainIDStr     string
		ethKeyStore    keys.AddressLister
		ethBalances    map[common.Address]*assets.Eth
		ethBalancesMtx sync.RWMutex
		sleeperTask    *utils.SleeperTask
	}

	NullBalanceMonitor struct{}
)

var _ BalanceMonitor = (*balanceMonitor)(nil)

// NewBalanceMonitor returns a new balanceMonitor
func NewBalanceMonitor(ethClient evmclient.Client, ethKeyStore keys.AddressLister, lggr logger.Logger) *balanceMonitor {
	bm := &balanceMonitor{
		ethClient:   ethClient,
		chainIDStr:  ethClient.ConfiguredChainID().String(),
		ethKeyStore: ethKeyStore,
		ethBalances: make(map[common.Address]*assets.Eth),
	}
	bm.Service, bm.eng = services.Config{
		Name:  "BalanceMonitor",
		Start: bm.start,
		Close: bm.close,
	}.NewServiceEngine(lggr)
	bm.sleeperTask = utils.NewSleeperTaskCtx(&worker{bm: bm})
	return bm
}

func (bm *balanceMonitor) start(ctx context.Context) error {
	// Always query latest balance on start
	(&worker{bm}).Work(ctx)
	return nil
}

// Close shuts down the BalanceMonitor, should not be used after this
func (bm *balanceMonitor) close() error {
	return bm.sleeperTask.Stop()
}

// OnNewLongestChain checks the balance for each key
func (bm *balanceMonitor) OnNewLongestChain(_ context.Context, _ *evmtypes.Head) {
	bm.eng.Debugw("BalanceMonitor: signalling balance worker")
	ok := bm.sleeperTask.WakeUpIfStarted()
	if !ok {
		bm.eng.Debugw("BalanceMonitor: ignoring OnNewLongestChain call, balance monitor is not started", "state", bm.sleeperTask.State())
	}
}

func (bm *balanceMonitor) updateBalance(ethBal assets.Eth, address common.Address) {
	bm.promUpdateEthBalance(&ethBal, address)

	bm.ethBalancesMtx.Lock()
	oldBal := bm.ethBalances[address]
	bm.ethBalances[address] = &ethBal
	bm.ethBalancesMtx.Unlock()

	lgr := logger.Named(bm.eng, "BalanceLog")
	lgr = logger.With(lgr,
		"address", address.Hex(),
		"ethBalance", ethBal.String(),
		"weiBalance", ethBal.ToInt())

	if oldBal == nil {
		lgr.Infof("ETH balance for %s: %s", address.Hex(), ethBal.String())
		return
	}

	if ethBal.Cmp(oldBal) != 0 {
		lgr.Infof("New ETH balance for %s: %s", address.Hex(), ethBal.String())
	}
}

func (bm *balanceMonitor) GetEthBalance(address common.Address) *assets.Eth {
	bm.ethBalancesMtx.RLock()
	defer bm.ethBalancesMtx.RUnlock()
	return bm.ethBalances[address]
}

// Deprecated: use github.com/smartcontractkit/chainlink-framework/metrics.AccountBalance instead.
var promETHBalance = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "eth_balance",
		Help: "Each Ethereum account's balance",
	},
	[]string{"account", "evmChainID"},
)

func (bm *balanceMonitor) promUpdateEthBalance(balance *assets.Eth, from common.Address) {
	balanceFloat, err := ApproximateFloat64(balance)

	if err != nil {
		bm.eng.Error(fmt.Errorf("updatePrometheusEthBalance: %w", err))
		return
	}

	metrics.NodeBalance.WithLabelValues(from.Hex(), bm.chainIDStr, metrics.EVM).Set(balanceFloat)
	// TODO: Remove deprecated metric
	promETHBalance.WithLabelValues(from.Hex(), bm.chainIDStr).Set(balanceFloat)
}

type worker struct {
	bm *balanceMonitor
}

func (*worker) Name() string {
	return "BalanceMonitorWorker"
}

func (w *worker) Work(ctx context.Context) {
	enabledAddresses, err := w.bm.ethKeyStore.EnabledAddresses(ctx)
	if err != nil {
		w.bm.eng.Error("BalanceMonitor: error getting keys", err)
	}

	var wg sync.WaitGroup

	wg.Add(len(enabledAddresses))
	for _, address := range enabledAddresses {
		go func(k common.Address) {
			defer wg.Done()
			w.checkAccountBalance(ctx, k)
		}(address)
	}
	wg.Wait()
}

func (w *worker) checkAccountBalance(ctx context.Context, address common.Address) {
	bal, err := w.bm.ethClient.BalanceAt(ctx, address, nil)
	if err != nil {
		w.bm.eng.Errorw("BalanceMonitor: error getting balance for key "+address.Hex(),
			"err", err,
			"address", address,
		)
	} else if bal == nil {
		w.bm.eng.Errorw(fmt.Sprintf("BalanceMonitor: error getting balance for key %s: invariant violation, bal may not be nil", address.Hex()),
			"err", err,
			"address", address,
		)
	} else {
		ethBal := assets.Eth(*bal)
		w.bm.updateBalance(ethBal, address)
	}
}

func (*NullBalanceMonitor) GetEthBalance(common.Address) *assets.Eth {
	return nil
}

// Start does noop for NullBalanceMonitor.
func (*NullBalanceMonitor) Start(context.Context) error                                { return nil }
func (*NullBalanceMonitor) Close() error                                               { return nil }
func (*NullBalanceMonitor) Ready() error                                               { return nil }
func (*NullBalanceMonitor) OnNewLongestChain(ctx context.Context, head *evmtypes.Head) {}

func ApproximateFloat64(e *assets.Eth) (float64, error) {
	ef := new(big.Float).SetInt(e.ToInt())
	weif := new(big.Float).SetInt(evmtypes.WeiPerEth)
	bf := new(big.Float).Quo(ef, weif)
	f64, _ := bf.Float64()
	if f64 == math.Inf(1) || f64 == math.Inf(-1) {
		return math.Inf(1), pkgerrors.New("assets.Eth.Float64: Could not approximate Eth value into float")
	}
	return f64, nil
}
