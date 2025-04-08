package heads_test

import (
	"context"
	"errors"
	"math/big"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"golang.org/x/exp/maps"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox/mailboxtest"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/configtest"

	"github.com/smartcontractkit/chainlink-framework/chains/heads"

	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	evmheads "github.com/smartcontractkit/chainlink-evm/pkg/heads"
	"github.com/smartcontractkit/chainlink-evm/pkg/heads/headstest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	ubig "github.com/smartcontractkit/chainlink-evm/pkg/utils/big"
)

func firstHead(t *testing.T, db *sqlx.DB) *evmtypes.Head {
	h := new(evmtypes.Head)
	if err := db.Get(h, `SELECT * FROM evm.heads ORDER BY number ASC LIMIT 1`); err != nil {
		t.Fatal(err)
	}
	return h
}

func TestHeadTracker_New(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	ethClient := clienttest.NewClientWithDefaultChainID(t)
	ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(testutils.Head(0), nil)
	// finalized
	ethClient.On("HeadByNumber", mock.Anything, big.NewInt(0)).Return(testutils.Head(0), nil)
	mockEth := &clienttest.MockEth{
		EthClient: ethClient,
	}
	ethClient.On("SubscribeToHeads", mock.Anything, mock.Anything).
		Maybe().
		Return(nil, mockEth.NewSub(t), nil)

	orm := evmheads.NewORM(*testutils.FixtureChainID, db)
	require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), testutils.Head(1)))
	last := testutils.Head(16)
	require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), last))
	require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), testutils.Head(10)))

	evmcfg := configtest.NewChainScopedConfig(t, nil)
	ht := createHeadTracker(t, ethClient, evmcfg.EVM(), evmcfg.EVM().HeadTracker(), orm)
	ht.Start(t)

	tests.AssertEventually(t, func() bool {
		latest := ht.headSaver.LatestChain()
		return latest != nil && last.Number == latest.Number
	})
}

func TestHeadTracker_MarkFinalized_MarksAndTrimsTable(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	config := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
		c.HeadTracker.HistoryDepth = ptr[uint32](100)
	})

	ethClient := clienttest.NewClientWithDefaultChainID(t)
	orm := evmheads.NewORM(*testutils.FixtureChainID, db)

	for idx := 0; idx < 200; idx++ {
		require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), testutils.Head(idx)))
	}

	latest := testutils.Head(201)
	require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), latest))

	ht := createHeadTracker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm)
	_, err := ht.headSaver.Load(tests.Context(t), latest.Number)
	require.NoError(t, err)
	require.NoError(t, ht.headSaver.MarkFinalized(tests.Context(t), latest))
	assert.Equal(t, big.NewInt(201), ht.headSaver.LatestChain().ToInt())

	firstHead := firstHead(t, db)
	assert.Equal(t, big.NewInt(101), firstHead.ToInt())

	lastHead, err := orm.LatestHead(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, int64(201), lastHead.Number)
}

func TestHeadTracker_Get(t *testing.T) {
	t.Parallel()

	start := testutils.Head(5)

	cases := []struct {
		name    string
		initial *evmtypes.Head
		toSave  *evmtypes.Head
		want    *big.Int
	}{
		{"greater", start, testutils.Head(6), big.NewInt(6)},
		{"less than", start, testutils.Head(1), big.NewInt(5)},
		{"zero", start, testutils.Head(0), big.NewInt(5)},
		{"nil", start, nil, big.NewInt(5)},
		{"nil no initial", nil, nil, big.NewInt(0)},
	}

	for i := range cases {
		test := cases[i]
		t.Run(test.name, func(t *testing.T) {
			db := testutils.NewSqlxDB(t)
			config := configtest.NewChainScopedConfig(t, nil)
			orm := evmheads.NewORM(*testutils.FixtureChainID, db)

			ethClient := clienttest.NewClientWithDefaultChainID(t)
			chStarted := make(chan struct{})
			mockEth := &clienttest.MockEth{
				EthClient: ethClient,
			}
			ethClient.On("SubscribeToHeads", mock.Anything).
				Maybe().
				Return(
					func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
						defer close(chStarted)
						return make(<-chan *evmtypes.Head), mockEth.NewSub(t), nil
					},
				)
			ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(testutils.Head(0), nil).Maybe()

			fnCall := ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Maybe()
			fnCall.RunFn = func(args mock.Arguments) {
				num := args.Get(1).(*big.Int)
				fnCall.ReturnArguments = mock.Arguments{testutils.Head(num.Int64()), nil}
			}

			if test.initial != nil {
				require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), test.initial))
			}

			ht := createHeadTracker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm)
			ht.Start(t)

			if test.toSave != nil {
				err := ht.headSaver.Save(tests.Context(t), test.toSave)
				require.NoError(t, err)
			}

			tests.AssertEventually(t, func() bool {
				latest := ht.headSaver.LatestChain().ToInt()
				return latest != nil && test.want.Cmp(latest) == 0
			})
		})
	}
}

func TestHeadTracker_Start_NewHeads(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	config := configtest.NewChainScopedConfig(t, nil)
	orm := evmheads.NewORM(*testutils.FixtureChainID, db)

	ethClient := clienttest.NewClientWithDefaultChainID(t)
	chStarted := make(chan struct{})
	mockEth := &clienttest.MockEth{EthClient: ethClient}
	sub := mockEth.NewSub(t)
	// for initial load
	ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(testutils.Head(0), nil).Once()
	ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(testutils.Head(0), nil).Once()
	// for backfill
	ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(testutils.Head(0), nil).Maybe()
	ch := make(chan *evmtypes.Head)
	ethClient.On("SubscribeToHeads", mock.Anything).
		Run(func(mock.Arguments) {
			close(chStarted)
		}).
		Return((<-chan *evmtypes.Head)(ch), sub, nil)

	ht := createHeadTracker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm)
	ht.Start(t)

	<-chStarted
}

func TestHeadTracker_NewHeads_FinalityViolations(t *testing.T) {
	t.Parallel()

	t.Run("Finality violation on block hash mismatch", func(t *testing.T) {
		db := testutils.NewSqlxDB(t)
		config := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
			c.FinalityTagEnabled = ptr(true)
			c.FinalityDepth = ptr(uint32(1))
		})
		orm := evmheads.NewORM(*testutils.FixtureChainID, db)

		ethClient := clienttest.NewClientWithDefaultChainID(t)
		chStarted := make(chan struct{})
		mockEth := &clienttest.MockEth{EthClient: ethClient}
		sub := mockEth.NewSub(t)

		h0 := testutils.Head(0)
		h0.IsFinalized.Store(true)

		// for initial load
		ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h0, nil).Once()
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(h0, nil).Once()
		// for backfill
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(h0, nil).Maybe()
		ethClient.On("HeadByHash", mock.Anything, mock.Anything).Return(h0, nil).Maybe()
		ch := make(chan *evmtypes.Head)
		ethClient.On("SubscribeToHeads", mock.Anything).
			Run(func(mock.Arguments) {
				close(chStarted)
			}).
			Return((<-chan *evmtypes.Head)(ch), sub, nil)

		ht := createHeadTracker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm)
		ht.Start(t)
		<-chStarted

		ch <- h0

		invalid0 := testutils.Head(0)
		invalid0.IsFinalized.Store(true)

		invalid1 := testutils.Head(1)
		invalid1.ParentHash = invalid0.Hash
		invalid1.Parent.Store(invalid0)

		ch <- invalid1 // Deliver head with finalized block hash mismatch compared to h0

		g := gomega.NewWithT(t)
		g.Eventually(func() bool {
			report := ht.headTracker.HealthReport()
			return slices.ContainsFunc(maps.Values(report), func(e error) bool {
				return errors.Is(e, types.ErrFinalityViolated)
			})
		}, 5*time.Second, tests.TestInterval).Should(gomega.BeTrue())
	})

	t.Run("Finality violation on old block", func(t *testing.T) {
		db := testutils.NewSqlxDB(t)
		config := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
			c.FinalityTagEnabled = ptr(true)
			c.FinalityDepth = ptr(uint32(0))
		})
		orm := evmheads.NewORM(*testutils.FixtureChainID, db)

		ethClient := clienttest.NewClientWithDefaultChainID(t)
		chStarted := make(chan struct{})
		mockEth := &clienttest.MockEth{EthClient: ethClient}
		sub := mockEth.NewSub(t)

		h0 := testutils.Head(0)
		h0.IsFinalized.Store(true)

		// for initial load
		ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h0, nil).Once()
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(h0, nil).Once()
		// for backfill
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(h0, nil).Maybe()
		ethClient.On("HeadByHash", mock.Anything, mock.Anything).Return(h0, nil).Maybe()
		ch := make(chan *evmtypes.Head)
		ethClient.On("SubscribeToHeads", mock.Anything).
			Run(func(mock.Arguments) {
				close(chStarted)
			}).
			Return((<-chan *evmtypes.Head)(ch), sub, nil)

		ht := createHeadTracker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm)
		ht.Start(t)
		<-chStarted

		h1 := testutils.Head(1)
		h1.IsFinalized.Store(true)
		h1.ParentHash = h0.Hash
		h1.Parent.Store(h0)
		ch <- h1

		h2 := testutils.Head(2)
		h2.IsFinalized.Store(true)
		h2.ParentHash = h1.Hash
		h2.Parent.Store(h1)
		ch <- h2
		require.NoError(t, ht.headSaver.Save(tests.Context(t), h2))

		// Send old head
		ch <- h0

		g := gomega.NewWithT(t)
		g.Eventually(func() bool {
			report := ht.headTracker.HealthReport()
			return slices.ContainsFunc(maps.Values(report), func(e error) bool {
				return errors.Is(e, types.ErrFinalityViolated) && strings.Contains(e.Error(), "got very old block")
			})
		}, 5*time.Second, tests.TestInterval).Should(gomega.BeTrue())
	})
}

func TestHeadTracker_Start(t *testing.T) {
	t.Parallel()

	const historyDepth = 100
	const finalityDepth = 50
	type opts struct {
		FinalityTagEnable       *bool
		MaxAllowedFinalityDepth *uint32
		FinalityTagBypass       *bool
		ORM                     evmheads.ORM
	}
	newHeadTracker := func(t *testing.T, opts opts) *headTrackerUniverse {
		config := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
			if opts.FinalityTagEnable != nil {
				c.FinalityTagEnabled = opts.FinalityTagEnable
			}
			c.HeadTracker.HistoryDepth = ptr[uint32](historyDepth)
			c.FinalityDepth = ptr[uint32](finalityDepth)
			if opts.MaxAllowedFinalityDepth != nil {
				c.HeadTracker.MaxAllowedFinalityDepth = opts.MaxAllowedFinalityDepth
			}

			if opts.FinalityTagBypass != nil {
				c.HeadTracker.FinalityTagBypass = opts.FinalityTagBypass
			}
		})
		if opts.ORM == nil {
			db := testutils.NewSqlxDB(t)
			opts.ORM = evmheads.NewORM(*testutils.FixtureChainID, db)
		}
		ethClient := clienttest.NewClientWithDefaultChainID(t)
		mockEth := &clienttest.MockEth{EthClient: ethClient}
		sub := mockEth.NewSub(t)
		ethClient.On("SubscribeToHeads", mock.Anything, mock.Anything).Return(nil, sub, nil).Maybe()
		return createHeadTracker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), opts.ORM)
	}
	t.Run("Starts even if failed to get initialHead", func(t *testing.T) {
		ht := newHeadTracker(t, opts{})
		ht.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(testutils.Head(0), errors.New("failed to get init head"))
		ht.Start(t)
		tests.AssertLogEventually(t, ht.observer, "Error handling initial head")
	})
	t.Run("Starts even if received invalid head", func(t *testing.T) {
		ht := newHeadTracker(t, opts{})
		ht.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(nil, nil)
		ht.Start(t)
		tests.AssertLogEventually(t, ht.observer, "Got nil initial head")
	})
	t.Run("Starts even if fails to get finalizedHead", func(t *testing.T) {
		ht := newHeadTracker(t, opts{FinalityTagEnable: ptr(true), FinalityTagBypass: ptr(false)})
		head := testutils.Head(1000)
		ht.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(head, nil).Once()
		ht.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, errors.New("failed to load latest finalized")).Once()
		ht.Start(t)
		tests.AssertLogEventually(t, ht.observer, "Error handling initial head")
	})
	t.Run("Starts even if latest finalizedHead is nil", func(t *testing.T) {
		ht := newHeadTracker(t, opts{FinalityTagEnable: ptr(true), FinalityTagBypass: ptr(false)})
		head := testutils.Head(1000)
		ht.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(head, nil).Once()
		ht.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, nil).Once()
		ht.ethClient.On("SubscribeToHeads", mock.Anything, mock.Anything).Return(nil, nil, errors.New("failed to connect")).Maybe()
		ht.Start(t)
		tests.AssertLogEventually(t, ht.observer, "Error handling initial head")
	})
	happyPathFT := func(t *testing.T, opts opts) {
		head := testutils.Head(1000)
		ht := newHeadTracker(t, opts)
		ctx := tests.Context(t)
		require.NoError(t, ht.orm.IdempotentInsertHead(ctx, testutils.Head(799)))
		ht.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(head, nil).Once()
		finalizedHead := testutils.Head(800)
		// on start
		ht.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(finalizedHead, nil).Once()
		// on backfill
		ht.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, errors.New("backfill call to finalized failed")).Maybe()
		ht.ethClient.On("SubscribeToHeads", mock.Anything, mock.Anything).Return(nil, nil, errors.New("failed to connect")).Maybe()
		ht.Start(t)
		tests.AssertLogEventually(t, ht.observer, "Received new head")
		tests.AssertEventually(t, func() bool {
			latest := ht.headTracker.LatestChain()
			return latest != nil && latest.Number == head.Number
		})
	}
	happyPathFD := func(t *testing.T, opts opts) {
		head := testutils.Head(1000)
		ht := newHeadTracker(t, opts)
		ht.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(head, nil).Once()
		finalizedHead := testutils.Head(head.Number - finalityDepth)
		ht.ethClient.On("HeadByNumber", mock.Anything, big.NewInt(finalizedHead.Number)).Return(finalizedHead, nil).Once()
		ctx := tests.Context(t)
		require.NoError(t, ht.orm.IdempotentInsertHead(ctx, testutils.Head(finalizedHead.Number-1)))
		// on backfill
		ht.ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(nil, errors.New("backfill call to finalized failed")).Maybe()
		ht.ethClient.On("SubscribeToHeads", mock.Anything, mock.Anything).Return(nil, nil, errors.New("failed to connect")).Maybe()
		ht.Start(t)
		tests.AssertLogEventually(t, ht.observer, "Received new head")
		tests.AssertEventually(t, func() bool {
			latest := ht.headTracker.LatestChain()
			return latest != nil && latest.Number == head.Number
		})
	}
	testCases := []struct {
		Name string
		Opts opts
		Run  func(t *testing.T, opts opts)
	}{
		{
			Name: "Happy path (Chain FT is disabled & Tracker's FT is disabled)",
			Opts: opts{FinalityTagEnable: ptr(false), FinalityTagBypass: ptr(true)},
			Run:  happyPathFD,
		},
		{
			Name: "Happy path (Chain FT is disabled & Tracker's FT is enabled, but ignored)",
			Opts: opts{FinalityTagEnable: ptr(false), FinalityTagBypass: ptr(false)},
			Run:  happyPathFD,
		},
		{
			Name: "Happy path (Chain FT is enabled & Tracker's FT is disabled)",
			Opts: opts{FinalityTagEnable: ptr(true), FinalityTagBypass: ptr(true)},
			Run:  happyPathFD,
		},
		{
			Name: "Happy path (Chain FT is enabled)",
			Opts: opts{FinalityTagEnable: ptr(true), FinalityTagBypass: ptr(false)},
			Run:  happyPathFT,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Run(t, tc.Opts)
		})
		t.Run("Disabled Persistence "+tc.Name, func(t *testing.T) {
			opts := tc.Opts
			opts.ORM = evmheads.NewNullORM()
			tc.Run(t, opts)
		})
	}
}

func TestHeadTracker_CallsHeadTrackableCallbacks(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	config := configtest.NewChainScopedConfig(t, nil)
	orm := evmheads.NewORM(*testutils.FixtureChainID, db)

	ethClient := clienttest.NewClientWithDefaultChainID(t)

	chchHeaders := make(chan testutils.RawSub[*evmtypes.Head], 1)
	mockEth := &clienttest.MockEth{EthClient: ethClient}
	chHead := make(chan *evmtypes.Head)
	ethClient.On("SubscribeToHeads", mock.Anything).
		Return(
			func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
				sub := mockEth.NewSub(t)
				chchHeaders <- testutils.NewRawSub(chHead, sub.Err())
				return chHead, sub, nil
			},
		)
	ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(testutils.Head(0), nil)
	ethClient.On("HeadByHash", mock.Anything, mock.Anything).Return(testutils.Head(0), nil).Maybe()

	checker := &CountingHeadTrackable{}
	ht := createHeadTrackerWithChecker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm, checker)

	ht.Start(t)
	assert.Equal(t, int32(0), checker.OnNewLongestChainCount())

	headers := <-chchHeaders
	headers.TrySend(&evmtypes.Head{Number: 1, Hash: utils.NewHash(), EVMChainID: ubig.New(testutils.FixtureChainID)})
	tests.AssertEventually(t, func() bool { return checker.OnNewLongestChainCount() == 1 })

	ht.Stop(t)
	assert.Equal(t, int32(1), checker.OnNewLongestChainCount())
}

func TestHeadTracker_ReconnectOnError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	db := testutils.NewSqlxDB(t)
	config := configtest.NewChainScopedConfig(t, nil)
	orm := evmheads.NewORM(*testutils.FixtureChainID, db)

	ethClient := clienttest.NewClientWithDefaultChainID(t)
	mockEth := &clienttest.MockEth{EthClient: ethClient}
	chHead := make(chan *evmtypes.Head)
	ethClient.On("SubscribeToHeads", mock.Anything).
		Return(
			func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
				return chHead, mockEth.NewSub(t), nil
			},
		)
	ethClient.On("SubscribeToHeads", mock.Anything).Return((<-chan *evmtypes.Head)(chHead), nil, errors.New("cannot reconnect"))
	ethClient.On("SubscribeToHeads", mock.Anything).
		Return(
			func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
				return chHead, mockEth.NewSub(t), nil
			},
		)
	ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(testutils.Head(0), nil)
	checker := &CountingHeadTrackable{}
	ht := createHeadTrackerWithChecker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm, checker)

	// connect
	ht.Start(t)
	assert.Equal(t, int32(0), checker.OnNewLongestChainCount())

	// trigger reconnect loop
	mockEth.SubsErr(errors.New("test error to force reconnect"))
	g.Eventually(checker.OnNewLongestChainCount, 5*time.Second, tests.TestInterval).Should(gomega.Equal(int32(1)))
}

func TestHeadTracker_ResubscribeOnSubscriptionError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	db := testutils.NewSqlxDB(t)
	config := configtest.NewChainScopedConfig(t, nil)
	orm := evmheads.NewORM(*testutils.FixtureChainID, db)

	ethClient := clienttest.NewClientWithDefaultChainID(t)

	ch := make(chan *evmtypes.Head)
	chchHeaders := make(chan testutils.RawSub[*evmtypes.Head], 1)
	mockEth := &clienttest.MockEth{EthClient: ethClient}
	ethClient.On("SubscribeToHeads", mock.Anything).
		Return(
			func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
				sub := mockEth.NewSub(t)
				chchHeaders <- testutils.NewRawSub(ch, sub.Err())
				return ch, sub, nil
			},
		)
	ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(testutils.Head(0), nil)
	ethClient.On("HeadByHash", mock.Anything, mock.Anything).Return(testutils.Head(0), nil).Maybe()

	checker := &CountingHeadTrackable{}
	ht := createHeadTrackerWithChecker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm, checker)

	ht.Start(t)
	assert.Equal(t, int32(0), checker.OnNewLongestChainCount())

	headers := <-chchHeaders
	go func() {
		headers.TrySend(testutils.Head(1))
	}()

	g.Eventually(func() bool {
		report := ht.headTracker.HealthReport()
		return !slices.ContainsFunc(maps.Values(report), func(e error) bool { return e != nil })
	}, 5*time.Second, tests.TestInterval).Should(gomega.BeTrue())

	// trigger reconnect loop
	headers.CloseCh()

	// wait for full disconnect and a new subscription
	g.Eventually(checker.OnNewLongestChainCount, 5*time.Second, tests.TestInterval).Should(gomega.Equal(int32(1)))
}

func TestHeadTracker_Start_LoadsLatestChain(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	config := configtest.NewChainScopedConfig(t, nil)
	ethClient := clienttest.NewClientWithDefaultChainID(t)

	heads := []*evmtypes.Head{
		testutils.Head(0),
		testutils.Head(1),
		testutils.Head(2),
		testutils.Head(3),
	}
	var parentHash common.Hash
	for i := 0; i < len(heads); i++ {
		if parentHash != (common.Hash{}) {
			heads[i].ParentHash = parentHash
		}
		parentHash = heads[i].Hash
	}
	ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(heads[3], nil).Maybe()
	ethClient.On("HeadByNumber", mock.Anything, big.NewInt(0)).Return(heads[0], nil).Maybe()
	ethClient.On("HeadByHash", mock.Anything, heads[2].Hash).Return(heads[2], nil).Maybe()
	ethClient.On("HeadByHash", mock.Anything, heads[1].Hash).Return(heads[1], nil).Maybe()
	ethClient.On("HeadByHash", mock.Anything, heads[0].Hash).Return(heads[0], nil).Maybe()

	chchHeaders := make(chan testutils.RawSub[*evmtypes.Head], 1)
	mockEth := &clienttest.MockEth{EthClient: ethClient}
	ch := make(chan *evmtypes.Head)
	ethClient.On("SubscribeToHeads", mock.Anything).
		Return(
			func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
				sub := mockEth.NewSub(t)
				chchHeaders <- testutils.NewRawSub(ch, sub.Err())
				return ch, sub, nil
			},
		)

	orm := evmheads.NewORM(*testutils.FixtureChainID, db)
	trackable := &CountingHeadTrackable{}
	ht := createHeadTrackerWithChecker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm, trackable)

	require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), heads[2]))

	ht.Start(t)

	assert.Equal(t, int32(0), trackable.OnNewLongestChainCount())

	headers := <-chchHeaders
	go func() {
		headers.TrySend(testutils.Head(1))
	}()

	require.Eventually(t, func() bool {
		report := ht.headTracker.HealthReport()
		services.CopyHealth(report, ht.headBroadcaster.HealthReport())
		return !slices.ContainsFunc(maps.Values(report), func(e error) bool { return e != nil })
	}, 5*time.Second, tests.TestInterval)

	h, err := orm.LatestHead(tests.Context(t))
	require.NoError(t, err)
	require.NotNil(t, h)
	assert.Equal(t, int64(3), h.Number)
}

func TestHeadTracker_SwitchesToLongestChainWithHeadSamplingEnabled(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)

	config := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
		c.FinalityDepth = ptr[uint32](50)
		// Need to set the buffer to something large since we inject a lot of heads at once and otherwise they will be dropped
		c.HeadTracker.MaxBufferSize = ptr[uint32](100)
		c.HeadTracker.SamplingInterval = commonconfig.MustNewDuration(2500 * time.Millisecond)
	})

	ethClient := clienttest.NewClientWithDefaultChainID(t)

	checker := headstest.NewTrackable[*evmtypes.Head, common.Hash](t)
	orm := evmheads.NewORM(*config.EVM().ChainID(), db)
	ht := createHeadTrackerWithChecker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm, checker)

	chchHeaders := make(chan testutils.RawSub[*evmtypes.Head], 1)
	mockEth := &clienttest.MockEth{EthClient: ethClient}
	chHead := make(chan *evmtypes.Head)
	ethClient.On("SubscribeToHeads", mock.Anything).
		Return(
			func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
				sub := mockEth.NewSub(t)
				chchHeaders <- testutils.NewRawSub(chHead, sub.Err())
				return chHead, sub, nil
			},
		)

	// ---------------------
	blocks := NewBlocks(t, 10)

	head0 := blocks.Head(0)
	// Initial query
	ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(head0, nil)
	// backfill query
	ethClient.On("HeadByNumber", mock.Anything, big.NewInt(0)).Return(head0, nil)
	ht.Start(t)

	headSeq := NewHeadBuffer(t)
	headSeq.Append(blocks.Head(0))
	headSeq.Append(blocks.Head(1))

	// Blocks 2 and 3 are out of order
	headSeq.Append(blocks.Head(3))
	headSeq.Append(blocks.Head(2))

	// Block 4 comes in
	headSeq.Append(blocks.Head(4))

	// Another block at level 4 comes in, that will be uncled
	headSeq.Append(blocks.NewHead(4))

	// Reorg happened forking from block 2
	blocksForked := blocks.ForkAt(t, 2, 5)
	headSeq.Append(blocksForked.Head(2))
	headSeq.Append(blocksForked.Head(3))
	headSeq.Append(blocksForked.Head(4))
	headSeq.Append(blocksForked.Head(5)) // Now the new chain is longer

	lastLongestChainAwaiter := testutils.NewAwaiter()

	// the callback is only called for head number 5 because of head sampling
	checker.On("OnNewLongestChain", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			h := args.Get(1).(*evmtypes.Head)
			// This is the new longest chain [0, 5], check that it came with its parents
			assert.Equal(t, uint32(6), h.ChainLength())
			assertChainWithParents(t, blocksForked, 5, 1, h)

			lastLongestChainAwaiter.ItHappened()
		}).Return().Once()

	headers := <-chchHeaders

	// This grotesque construction is the only way to do dynamic return values using
	// the mock package.  We need dynamic returns because we're simulating reorgs.
	latestHeadByHash := make(map[common.Hash]*evmtypes.Head)
	latestHeadByHashMu := new(sync.Mutex)

	fnCall := ethClient.On("HeadByHash", mock.Anything, mock.Anything).Maybe()
	fnCall.RunFn = func(args mock.Arguments) {
		latestHeadByHashMu.Lock()
		defer latestHeadByHashMu.Unlock()
		hash := args.Get(1).(common.Hash)
		head := latestHeadByHash[hash]
		fnCall.ReturnArguments = mock.Arguments{head, nil}
	}

	for _, h := range headSeq.Heads {
		latestHeadByHashMu.Lock()
		latestHeadByHash[h.Hash] = h
		latestHeadByHashMu.Unlock()
		headers.TrySend(h)
	}

	// default 10s may not be sufficient, so using tests.WaitTimeout(t)
	lastLongestChainAwaiter.AwaitOrFail(t, tests.WaitTimeout(t))
	ht.Stop(t)
	assert.Equal(t, int64(5), ht.headSaver.LatestChain().Number)

	for _, h := range headSeq.Heads {
		c := ht.headSaver.Chain(h.Hash)
		require.NotNil(t, c)
		assert.Equal(t, c.ParentHash, h.ParentHash)
		assert.Equal(t, c.Timestamp.Unix(), h.Timestamp.Unix())
		assert.Equal(t, c.Number, h.Number)
	}
}

func assertChainWithParents(t testing.TB, blocks *blocks, startBN, endBN uint64, h *evmtypes.Head) {
	for blockNumber := startBN; blockNumber >= endBN; blockNumber-- {
		require.NotNil(t, h)
		assert.Equal(t, blockNumber, uint64(h.Number))
		assert.Equal(t, blocks.Head(blockNumber).Hash, h.Hash)
		// move to parent
		h = h.Parent.Load()
	}
}

func TestHeadTracker_SwitchesToLongestChainWithHeadSamplingDisabled(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)

	config := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
		c.FinalityDepth = ptr[uint32](50)
		// Need to set the buffer to something large since we inject a lot of heads at once and otherwise they will be dropped
		c.HeadTracker.MaxBufferSize = ptr[uint32](100)
		c.HeadTracker.SamplingInterval = commonconfig.MustNewDuration(0)
	})

	ethClient := clienttest.NewClientWithDefaultChainID(t)

	checker := headstest.NewTrackable[*evmtypes.Head, common.Hash](t)
	orm := evmheads.NewORM(*testutils.FixtureChainID, db)
	ht := createHeadTrackerWithChecker(t, ethClient, config.EVM(), config.EVM().HeadTracker(), orm, checker)

	chchHeaders := make(chan testutils.RawSub[*evmtypes.Head], 1)
	mockEth := &clienttest.MockEth{EthClient: ethClient}
	chHead := make(chan *evmtypes.Head)
	ethClient.On("SubscribeToHeads", mock.Anything).
		Return(
			func(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
				sub := mockEth.NewSub(t)
				chchHeaders <- testutils.NewRawSub(chHead, sub.Err())
				return chHead, sub, nil
			},
		)

	// ---------------------
	blocks := NewBlocks(t, 10)

	head0 := blocks.Head(0) // evmtypes.Head{Number: 0, Hash: utils.NewHash(), ParentHash: utils.NewHash(), Timestamp: time.Unix(0, 0)}
	// Initial query
	ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(head0, nil)
	// backfill
	ethClient.On("HeadByNumber", mock.Anything, big.NewInt(0)).Return(head0, nil)

	headSeq := NewHeadBuffer(t)
	headSeq.Append(blocks.Head(0))
	headSeq.Append(blocks.Head(1))

	// Blocks 2 and 3 are out of order
	headSeq.Append(blocks.Head(3))
	headSeq.Append(blocks.Head(2))

	// Block 4 comes in
	headSeq.Append(blocks.Head(4))

	// Another block at level 4 comes in, that will be uncled
	headSeq.Append(blocks.NewHead(4))

	// Reorg happened forking from block 2
	blocksForked := blocks.ForkAt(t, 2, 5)
	headSeq.Append(blocksForked.Head(2))
	headSeq.Append(blocksForked.Head(3))
	headSeq.Append(blocksForked.Head(4))
	headSeq.Append(blocksForked.Head(5)) // Now the new chain is longer

	lastLongestChainAwaiter := testutils.NewAwaiter()

	checker.On("OnNewLongestChain", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			h := args.Get(1).(*evmtypes.Head)
			require.Equal(t, int64(0), h.Number)
			require.Equal(t, blocks.Head(0).Hash, h.Hash)
		}).Return().Once()

	checker.On("OnNewLongestChain", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			h := args.Get(1).(*evmtypes.Head)
			require.Equal(t, int64(1), h.Number)
			require.Equal(t, blocks.Head(1).Hash, h.Hash)
		}).Return().Once()

	checker.On("OnNewLongestChain", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			h := args.Get(1).(*evmtypes.Head)
			require.Equal(t, int64(3), h.Number)
			require.Equal(t, blocks.Head(3).Hash, h.Hash)
		}).Return().Once()

	checker.On("OnNewLongestChain", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			h := args.Get(1).(*evmtypes.Head)
			assertChainWithParents(t, blocks, 4, 1, h)
		}).Return().Once()

	checker.On("OnNewLongestChain", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			h := args.Get(1).(*evmtypes.Head)
			assertChainWithParents(t, blocksForked, 5, 1, h)
			lastLongestChainAwaiter.ItHappened()
		}).Return().Once()

	ht.Start(t)

	headers := <-chchHeaders

	// This grotesque construction is the only way to do dynamic return values using
	// the mock package.  We need dynamic returns because we're simulating reorgs.
	latestHeadByHash := make(map[common.Hash]*evmtypes.Head)
	latestHeadByHashMu := new(sync.Mutex)

	fnCall := ethClient.On("HeadByHash", mock.Anything, mock.Anything).Maybe()
	fnCall.RunFn = func(args mock.Arguments) {
		latestHeadByHashMu.Lock()
		defer latestHeadByHashMu.Unlock()
		hash := args.Get(1).(common.Hash)
		head := latestHeadByHash[hash]
		fnCall.ReturnArguments = mock.Arguments{head, nil}
	}

	for _, h := range headSeq.Heads {
		latestHeadByHashMu.Lock()
		latestHeadByHash[h.Hash] = h
		latestHeadByHashMu.Unlock()
		headers.TrySend(h)
		time.Sleep(tests.TestInterval)
	}

	// default 10s may not be sufficient, so using tests.WaitTimeout(t)
	lastLongestChainAwaiter.AwaitOrFail(t, tests.WaitTimeout(t))
	ht.Stop(t)
	assert.Equal(t, int64(5), ht.headSaver.LatestChain().Number)

	for _, h := range headSeq.Heads {
		c := ht.headSaver.Chain(h.Hash)
		require.NotNil(t, c)
		assert.Equal(t, c.ParentHash, h.ParentHash)
		assert.Equal(t, c.Timestamp.Unix(), h.Timestamp.Unix())
		assert.Equal(t, c.Number, h.Number)
	}
}

func TestHeadTracker_Backfill(t *testing.T) {
	t.Parallel()
	t.Run("Enabled Persistence", func(t *testing.T) {
		testHeadTrackerBackfill(t, func(t *testing.T) evmheads.ORM {
			db := testutils.NewSqlxDB(t)
			return evmheads.NewORM(*testutils.FixtureChainID, db)
		})
	})
	t.Run("Disabled Persistence", func(t *testing.T) {
		testHeadTrackerBackfill(t, func(t *testing.T) evmheads.ORM {
			return evmheads.NewNullORM()
		})
	})
}

func testHeadTrackerBackfill(t *testing.T, newORM func(t *testing.T) evmheads.ORM) {
	// Heads are arranged as follows:
	// headN indicates an unpersisted ethereum header
	// hN indicates a persisted head record
	//
	// (1)->(H0)
	//
	//       (14Orphaned)-+
	//                    +->(13)->(12)->(11)->(H10)->(9)->(H8)
	// (15)->(14)---------+

	head0 := evmtypes.NewHead(big.NewInt(0), utils.NewHash(), common.BigToHash(big.NewInt(0)), ubig.New(testutils.FixtureChainID))

	h1 := testutils.Head(1)
	h1.ParentHash = head0.Hash

	head8 := evmtypes.NewHead(big.NewInt(8), utils.NewHash(), utils.NewHash(), ubig.New(testutils.FixtureChainID))

	h9 := testutils.Head(9)
	h9.ParentHash = head8.Hash

	head10 := evmtypes.NewHead(big.NewInt(10), utils.NewHash(), h9.Hash, ubig.New(testutils.FixtureChainID))

	h11 := testutils.Head(11)
	h11.ParentHash = head10.Hash

	h12 := testutils.Head(12)
	h12.ParentHash = h11.Hash

	h13 := testutils.Head(13)
	h13.ParentHash = h12.Hash

	h14Orphaned := testutils.Head(14)
	h14Orphaned.ParentHash = h13.Hash

	h14 := testutils.Head(14)
	h14.ParentHash = h13.Hash

	h15 := testutils.Head(15)
	h15.ParentHash = h14.Hash

	hs := []*evmtypes.Head{
		h9,
		h11,
		h12,
		h13,
		h14Orphaned,
		h14,
		h15,
	}

	ctx := tests.Context(t)

	type opts struct {
		Heads                   []*evmtypes.Head
		FinalityTagEnabled      bool
		FinalizedBlockOffset    uint32
		FinalityDepth           uint32
		MaxAllowedFinalityDepth uint32
	}
	newHeadTrackerUniverse := func(t *testing.T, opts opts) *headTrackerUniverse {
		evmcfg := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
			c.FinalityTagEnabled = ptr(opts.FinalityTagEnabled)
			c.FinalizedBlockOffset = ptr(opts.FinalizedBlockOffset)
			c.FinalityDepth = ptr(opts.FinalityDepth)
			c.HeadTracker.FinalityTagBypass = ptr(false)
			if opts.MaxAllowedFinalityDepth > 0 {
				c.HeadTracker.MaxAllowedFinalityDepth = ptr(opts.MaxAllowedFinalityDepth)
			}
		})

		ethClient := clienttest.NewClient(t)
		ethClient.On("ConfiguredChainID", mock.Anything).Return(evmcfg.EVM().ChainID(), nil)
		ht := createHeadTracker(t, ethClient, evmcfg.EVM(), evmcfg.EVM().HeadTracker(), newORM(t))
		for i := range opts.Heads {
			require.NoError(t, ht.headSaver.Save(tests.Context(t), opts.Heads[i]))
		}
		_, err := ht.headSaver.Load(tests.Context(t), 0)
		require.NoError(t, err)
		return ht
	}

	t.Run("returns error if failed to get latestFinalized block", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		const expectedError = "failed to fetch latest finalized block"
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, errors.New(expectedError)).Once()

		err := htu.Backfill(ctx, h12)
		require.ErrorContains(t, err, expectedError)
	})
	t.Run("returns error if latestFinalized is not valid", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, nil).Once()

		err := htu.Backfill(ctx, h12)
		require.EqualError(t, err, "failed to calculate finalized block: failed to get valid latest finalized block")
	})
	t.Run("Returns error if finality gap is too big", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true, MaxAllowedFinalityDepth: 2})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h9, nil).Once()

		err := htu.Backfill(ctx, h12)
		require.EqualError(t, err, "gap between latest finalized block (9) and current head (12) is too large (> 2)")
	})
	t.Run("Returns error if finalized head is ahead of canonical", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h14Orphaned, nil).Once()

		err := htu.Backfill(ctx, h12)
		require.EqualError(t, err, "expected head of canonical chain to be ahead of the latestFinalized, but this may be normal on chains with fast finality due to fetch timing")
	})
	t.Run("Returns error if finalizedHead is not present in the canonical chain", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: hs, FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h14Orphaned, nil)

		err := htu.Backfill(ctx, h15)
		require.ErrorAs(t, err, &heads.FinalizedMissingError[common.Hash]{})
	})
	t.Run("Marks all blocks in chain that are older than finalized", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: hs, FinalityTagEnabled: true})

		assertFinalized := func(expectedFinalized bool, msg string, heads ...*evmtypes.Head) {
			for _, h := range heads {
				storedHead := htu.headSaver.Chain(h.Hash)
				assert.Equal(t, expectedFinalized, storedHead != nil && storedHead.IsFinalized.Load(), msg, "block_number", h.Number)
			}
		}

		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h14, nil).Once()
		err := htu.Backfill(ctx, h15)
		require.NoError(t, err)
		assertFinalized(true, "expected heads to be marked as finalized after backfill", h14, h13, h12, h11)
		assertFinalized(false, "expected heads to remain unfinalized", h15, &head10)
	})

	t.Run("fetches a missing head", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: hs, FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h9, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, head10.Hash).
			Return(&head10, nil)

		err := htu.Backfill(ctx, h12)
		require.NoError(t, err)

		h := htu.headSaver.Chain(h12.Hash)

		for expectedBlockNumber := int64(12); expectedBlockNumber >= 9; expectedBlockNumber-- {
			require.NotNil(t, h)
			assert.Equal(t, expectedBlockNumber, h.Number)
			h = h.Parent.Load()
		}

		writtenHead := htu.headSaver.Chain(head10.Hash)
		require.NoError(t, err)
		assert.Equal(t, int64(10), writtenHead.Number)
	})
	t.Run("fetches only heads that are missing", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: hs, FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(&head8, nil).Once()

		htu.ethClient.On("HeadByHash", mock.Anything, head10.Hash).
			Return(&head10, nil)
		htu.ethClient.On("HeadByHash", mock.Anything, head8.Hash).
			Return(&head8, nil)

		err := htu.Backfill(ctx, h15)
		require.NoError(t, err)

		h := htu.headSaver.Chain(h15.Hash)

		require.Equal(t, uint32(8), h.ChainLength())
		earliestInChain := h.EarliestInChain()
		assert.Equal(t, head8.Number, earliestInChain.BlockNumber())
		assert.Equal(t, head8.Hash, earliestInChain.BlockHash())
	})

	t.Run("abandons backfill and returns error if the eth node returns not found", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: hs, FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(&head8, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, head10.Hash).
			Return(&head10, nil).
			Once()
		htu.ethClient.On("HeadByHash", mock.Anything, head8.Hash).
			Return(nil, ethereum.NotFound).
			Once()

		err := htu.Backfill(ctx, h12)
		require.Error(t, err)
		require.ErrorContains(t, err, "fetchAndSaveHead failed: not found")

		h := htu.headSaver.Chain(h12.Hash)

		// Should contain 12, 11, 10, 9
		assert.Equal(t, 4, int(h.ChainLength()))
		assert.Equal(t, int64(9), h.EarliestInChain().BlockNumber())
	})

	t.Run("abandons backfill and returns error if the context time budget is exceeded", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: hs, FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(&head8, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, head10.Hash).
			Return(&head10, nil)
		lctx, cancel := context.WithCancel(ctx)
		htu.ethClient.On("HeadByHash", mock.Anything, head8.Hash).
			Return(nil, context.DeadlineExceeded).Run(func(args mock.Arguments) {
			cancel()
		})

		err := htu.headTracker.Backfill(lctx, h12, nil)
		require.Error(t, err)
		require.ErrorContains(t, err, "fetchAndSaveHead failed: context canceled")

		h := htu.headSaver.Chain(h12.Hash)

		// Should contain 12, 11, 10, 9
		assert.Equal(t, 4, int(h.ChainLength()))
		assert.Equal(t, int64(9), h.EarliestInChain().BlockNumber())
	})
	t.Run("abandons backfill and returns error when fetching a block by hash fails, indicating a reorg", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h11, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h14.Hash).Return(h14, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h13.Hash).Return(h13, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h12.Hash).Return(nil, errors.New("not found")).Once()

		err := htu.Backfill(ctx, h15)

		require.Error(t, err)
		require.ErrorContains(t, err, "fetchAndSaveHead failed: not found")

		h := htu.headSaver.Chain(h14.Hash)

		// Should contain 14, 13 (15 was never added). When trying to get the parent of h13 by hash, a reorg happened and backfill exited.
		assert.Equal(t, 2, int(h.ChainLength()))
		assert.Equal(t, int64(13), h.EarliestInChain().BlockNumber())
	})
	t.Run("marks head as finalized, if latestHead = finalizedHead (0 finality depth)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: []*evmtypes.Head{h15}, FinalityTagEnabled: true})
		finalizedH15 := h15 // copy h15 to have different addresses
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(finalizedH15, nil).Once()
		err := htu.Backfill(ctx, h15)
		require.NoError(t, err)

		h := htu.headSaver.LatestChain()

		// Should contain 14, 13 (15 was never added). When trying to get the parent of h13 by hash, a reorg happened and backfill exited.
		assert.Equal(t, 1, int(h.ChainLength()))
		assert.True(t, h.IsFinalized.Load())
		assert.Equal(t, h15.BlockNumber(), h.BlockNumber())
		assert.Equal(t, h15.Hash, h.Hash)
	})
	t.Run("marks block as finalized according to FinalizedBlockOffset (finality tag)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: []*evmtypes.Head{h15}, FinalityTagEnabled: true, FinalizedBlockOffset: 2})
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h14, nil).Once()
		// calculateLatestFinalizedBlock fetches blocks at LatestFinalized - FinalizedBlockOffset
		htu.ethClient.On("HeadByNumber", mock.Anything, big.NewInt(h12.Number)).Return(h12, nil).Once()
		// backfill from 15 to 12
		htu.ethClient.On("HeadByHash", mock.Anything, h12.Hash).Return(h12, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h13.Hash).Return(h13, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h14.Hash).Return(h14, nil).Once()
		err := htu.Backfill(ctx, h15)
		require.NoError(t, err)

		h := htu.headSaver.LatestChain()
		// h - must contain 15, 14, 13, 12 and only 12 is finalized
		assert.Equal(t, 4, int(h.ChainLength()))
		for ; h.Hash != h12.Hash; h = h.Parent.Load() {
			assert.False(t, h.IsFinalized.Load())
		}

		assert.True(t, h.IsFinalized.Load())
		assert.Equal(t, h12.BlockNumber(), h.BlockNumber())
		assert.Equal(t, h12.Hash, h.Hash)
	})
	t.Run("marks block as finalized according to FinalizedBlockOffset (finality depth)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: []*evmtypes.Head{h15}, FinalityDepth: 1, FinalizedBlockOffset: 2})
		htu.ethClient.On("HeadByNumber", mock.Anything, big.NewInt(12)).Return(h12, nil).Once()

		// backfill from 15 to 12
		htu.ethClient.On("HeadByHash", mock.Anything, h14.Hash).Return(h14, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h13.Hash).Return(h13, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h12.Hash).Return(h12, nil).Once()
		err := htu.Backfill(ctx, h15)
		require.NoError(t, err)

		h := htu.headSaver.LatestChain()
		// h - must contain 15, 14, 13, 12 and only 12 is finalized
		assert.Equal(t, 4, int(h.ChainLength()))
		for ; h.Hash != h12.Hash; h = h.Parent.Load() {
			assert.False(t, h.IsFinalized.Load())
		}

		assert.True(t, h.IsFinalized.Load())
		assert.Equal(t, h12.BlockNumber(), h.BlockNumber())
		assert.Equal(t, h12.Hash, h.Hash)
	})
	t.Run("marks block as finalized according to FinalizedBlockOffset even with instant finality", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: []*evmtypes.Head{h15}, FinalityDepth: 0, FinalizedBlockOffset: 2})
		htu.ethClient.On("HeadByNumber", mock.Anything, big.NewInt(13)).Return(h13, nil).Once()

		// backfill from 15 to 13
		htu.ethClient.On("HeadByHash", mock.Anything, h14.Hash).Return(h14, nil).Once()
		htu.ethClient.On("HeadByHash", mock.Anything, h13.Hash).Return(h13, nil).Once()
		err := htu.Backfill(ctx, h15)
		require.NoError(t, err)

		h := htu.headSaver.LatestChain()
		// h - must contain 15, 14, 13, only 13 is finalized
		assert.Equal(t, 3, int(h.ChainLength()))
		for ; h.Hash != h13.Hash; h = h.Parent.Load() {
			assert.False(t, h.IsFinalized.Load())
		}

		assert.True(t, h.IsFinalized.Load())
		assert.Equal(t, h13.BlockNumber(), h.BlockNumber())
		assert.Equal(t, h13.Hash, h.Hash)
	})

	t.Run("finality violation error on finalized block hash mismatch", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{Heads: []*evmtypes.Head{h15}, FinalityTagEnabled: true, FinalizedBlockOffset: 2})
		htu.ethClient.On("HeadByNumber", mock.Anything, big.NewInt(12)).Return(h12, nil).Maybe()
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h14, nil)
		htu.Start(t)

		finalized12 := testutils.Head(12)
		finalized12.IsFinalized.Store(true)
		finalized12.Parent.Store(h11)
		finalized12.ParentHash = h11.Hash

		// Invalid chain with block mismatch
		invalid12 := testutils.Head(12)
		invalid12.IsFinalized.Store(true)
		invalid12.Parent.Store(h1) // Mismatch with incorrect parent
		invalid12.ParentHash = h1.Hash

		invalid13 := testutils.Head(12)
		invalid13.Parent.Store(invalid12)
		invalid13.ParentHash = invalid12.Hash

		err := htu.headTracker.Backfill(ctx, invalid13, finalized12)
		require.ErrorIs(t, err, types.ErrFinalityViolated)

		g := gomega.NewWithT(t)
		g.Eventually(func() bool {
			report := htu.headTracker.HealthReport()
			return slices.ContainsFunc(maps.Values(report), func(e error) bool {
				return errors.Is(e, types.ErrFinalityViolated)
			})
		}, 5*time.Second, tests.TestInterval).Should(gomega.BeTrue())
	})
}

func TestHeadTracker_LatestAndFinalizedBlock(t *testing.T) {
	t.Parallel()

	ctx := tests.Context(t)

	h11 := testutils.Head(11)
	h11.ParentHash = utils.NewHash()

	h12 := testutils.Head(12)
	h12.ParentHash = h11.Hash

	h13 := testutils.Head(13)
	h13.ParentHash = h12.Hash

	type opts struct {
		Heads                []*evmtypes.Head
		FinalityTagEnabled   bool
		FinalizedBlockOffset uint32
		FinalityDepth        uint32
	}

	newHeadTrackerUniverse := func(t *testing.T, opts opts) *headTrackerUniverse {
		evmcfg := configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
			c.FinalityTagEnabled = ptr(opts.FinalityTagEnabled)
			c.FinalizedBlockOffset = ptr(opts.FinalizedBlockOffset)
			c.FinalityDepth = ptr(opts.FinalityDepth)
		})

		db := testutils.NewSqlxDB(t)
		orm := evmheads.NewORM(*testutils.FixtureChainID, db)
		for i := range opts.Heads {
			require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), opts.Heads[i]))
		}
		ethClient := clienttest.NewClient(t)
		ethClient.On("ConfiguredChainID", mock.Anything).Return(testutils.FixtureChainID, nil)
		ht := createHeadTracker(t, ethClient, evmcfg.EVM(), evmcfg.EVM().HeadTracker(), orm)
		_, err := ht.headSaver.Load(tests.Context(t), 0)
		require.NoError(t, err)
		return ht
	}
	t.Run("returns error if failed to get latest block", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		const expectedError = "failed to fetch latest block"
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(nil, errors.New(expectedError)).Once()

		_, _, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.ErrorContains(t, err, expectedError)
	})
	t.Run("returns error if latest block is invalid", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(nil, nil).Once()

		_, _, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.ErrorContains(t, err, "expected latest block to be valid")
	})
	t.Run("returns error if failed to get latest finalized (finality tag)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()
		const expectedError = "failed to get latest finalized block"
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, errors.New(expectedError)).Once()

		_, _, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.ErrorContains(t, err, expectedError)
	})
	t.Run("returns error if latest finalized block is not valid (finality tag)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, nil).Once()

		_, _, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.ErrorContains(t, err, "failed to get valid latest finalized block")
	})
	t.Run("returns latest finalized block as is if FinalizedBlockOffset is 0 (finality tag)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h11, nil).Once()

		actualL, actualLF, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.NoError(t, err)
		assert.Equal(t, actualL, h13)
		assert.Equal(t, actualLF, h11)
	})
	t.Run("returns latest finalized block with offset from cache (finality tag)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true, FinalizedBlockOffset: 1, Heads: []*evmtypes.Head{h13, h12, h11}})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h12, nil).Once()

		actualL, actualLF, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.NoError(t, err)
		assert.Equal(t, actualL.Number, h13.Number)
		assert.Equal(t, actualLF.Number, h11.Number)
	})
	t.Run("returns latest finalized block with offset from RPC (finality tag)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityTagEnabled: true, FinalizedBlockOffset: 2, Heads: []*evmtypes.Head{h13, h12, h11}})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()
		htu.ethClient.On("LatestFinalizedBlock", mock.Anything).Return(h12, nil).Once()
		h10 := testutils.Head(10)
		htu.ethClient.On("HeadByNumber", mock.Anything, big.NewInt(10)).Return(h10, nil).Once()

		actualL, actualLF, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.NoError(t, err)
		assert.Equal(t, actualL.Number, h13.Number)
		assert.Equal(t, actualLF.Number, h10.Number)
	})
	t.Run("returns current head for both latest and finalized for FD = 0 (finality depth)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()

		actualL, actualLF, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.NoError(t, err)
		assert.Equal(t, actualL.Number, h13.Number)
		assert.Equal(t, actualLF.Number, h13.Number)
	})
	t.Run("returns latest finalized block with offset from cache (finality depth)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityDepth: 1, FinalizedBlockOffset: 1, Heads: []*evmtypes.Head{h13, h12, h11}})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()

		actualL, actualLF, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.NoError(t, err)
		assert.Equal(t, actualL.Number, h13.Number)
		assert.Equal(t, actualLF.Number, h11.Number)
	})
	t.Run("returns latest finalized block with offset from RPC (finality depth)", func(t *testing.T) {
		htu := newHeadTrackerUniverse(t, opts{FinalityDepth: 1, FinalizedBlockOffset: 2, Heads: []*evmtypes.Head{h13, h12, h11}})
		htu.ethClient.On("HeadByNumber", mock.Anything, (*big.Int)(nil)).Return(h13, nil).Once()
		h10 := testutils.Head(10)
		htu.ethClient.On("HeadByNumber", mock.Anything, big.NewInt(10)).Return(h10, nil).Once()

		actualL, actualLF, err := htu.headTracker.LatestAndFinalizedBlock(ctx)
		require.NoError(t, err)
		assert.Equal(t, actualL.Number, h13.Number)
		assert.Equal(t, actualLF.Number, h10.Number)
	})
}

func createHeadTracker(t testing.TB, ethClient *clienttest.Client, config heads.ChainConfig, htConfig heads.TrackerConfig, orm evmheads.ORM) *headTrackerUniverse {
	lggr, ob := logger.TestObserved(t, zap.DebugLevel)
	hb := evmheads.NewBroadcaster(lggr)
	hs := evmheads.NewSaver(lggr, orm, config, htConfig)
	mailMon := mailboxtest.NewMonitor(t)
	return &headTrackerUniverse{
		mu:              new(sync.Mutex),
		headTracker:     evmheads.NewTracker(lggr, ethClient, config, htConfig, hb, hs, mailMon),
		headBroadcaster: hb,
		headSaver:       hs,
		mailMon:         mailMon,
		observer:        ob,
		orm:             orm,
		ethClient:       ethClient,
	}
}

func createHeadTrackerWithChecker(t *testing.T, ethClient *clienttest.Client, config heads.ChainConfig, htConfig heads.TrackerConfig, orm evmheads.ORM, checker evmheads.Trackable) *headTrackerUniverse {
	lggr, ob := logger.TestObserved(t, zap.DebugLevel)
	hb := evmheads.NewBroadcaster(lggr)
	hs := evmheads.NewSaver(lggr, orm, config, htConfig)
	hb.Subscribe(checker)
	mailMon := mailboxtest.NewMonitor(t)
	ht := evmheads.NewTracker(lggr, ethClient, config, htConfig, hb, hs, mailMon)
	return &headTrackerUniverse{
		mu:              new(sync.Mutex),
		headTracker:     ht,
		headBroadcaster: hb,
		headSaver:       hs,
		mailMon:         mailMon,
		observer:        ob,
		orm:             orm,
		ethClient:       ethClient,
	}
}

type headTrackerUniverse struct {
	mu              *sync.Mutex
	stopped         bool
	headTracker     evmheads.Tracker
	headBroadcaster evmheads.Broadcaster
	headSaver       evmheads.HeadSaver
	mailMon         *mailbox.Monitor
	observer        *observer.ObservedLogs
	orm             evmheads.ORM
	ethClient       *clienttest.Client
}

func (u *headTrackerUniverse) Backfill(ctx context.Context, head *evmtypes.Head) error {
	return u.headTracker.Backfill(ctx, head, head) // Passing head as prevHead should always verify hashes correctly
}

func (u *headTrackerUniverse) Start(t *testing.T) {
	u.mu.Lock()
	defer u.mu.Unlock()
	ctx := tests.Context(t)
	require.NoError(t, u.headBroadcaster.Start(ctx))
	require.NoError(t, u.headTracker.Start(ctx))
	require.NoError(t, u.mailMon.Start(ctx))

	g := gomega.NewWithT(t)
	g.Eventually(func() bool {
		report := u.headBroadcaster.HealthReport()
		return !slices.ContainsFunc(maps.Values(report), func(e error) bool { return e != nil })
	}, 5*time.Second, tests.TestInterval).Should(gomega.BeTrue())

	t.Cleanup(func() {
		u.Stop(t)
	})
}

func (u *headTrackerUniverse) Stop(t *testing.T) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.stopped {
		return
	}
	u.stopped = true
	require.NoError(t, u.headBroadcaster.Close())
	require.NoError(t, u.headTracker.Close())
	require.NoError(t, u.mailMon.Close())
}

func ptr[T any](t T) *T { return &t }

// headBuffer - stores heads in sequence, with increasing timestamps
type headBuffer struct {
	t     *testing.T
	Heads []*evmtypes.Head
}

func NewHeadBuffer(t *testing.T) *headBuffer {
	return &headBuffer{
		t:     t,
		Heads: make([]*evmtypes.Head, 0),
	}
}

func (hb *headBuffer) Append(head *evmtypes.Head) {
	cloned := &evmtypes.Head{
		Number:     head.Number,
		Hash:       head.Hash,
		ParentHash: head.ParentHash,
		Timestamp:  head.Timestamp,
		EVMChainID: head.EVMChainID,
	}
	cloned.Parent.Store(head.Parent.Load())
	hb.Heads = append(hb.Heads, cloned)
}

type blocks struct {
	t     testing.TB
	Heads map[int64]*evmtypes.Head
}

func (b *blocks) Head(number uint64) *evmtypes.Head {
	return b.Heads[int64(number)]
}

func NewBlocks(t testing.TB, numHashes int) *blocks {
	b := &blocks{
		t:     t,
		Heads: make(map[int64]*evmtypes.Head, numHashes),
	}

	if numHashes == 0 {
		return b
	}

	now := time.Now()
	b.Heads[0] = &evmtypes.Head{Hash: testutils.NewHash(), Number: 0, Timestamp: now, EVMChainID: ubig.New(testutils.FixtureChainID)}
	for i := 1; i < numHashes; i++ {
		//nolint:gosec // G115
		head := b.NewHead(uint64(i))
		b.Heads[head.Number] = head
	}

	return b
}

func (b *blocks) ForkAt(t *testing.T, blockNum int64, numHashes int) *blocks {
	forked := NewBlocks(t, len(b.Heads)+numHashes)
	if _, exists := forked.Heads[blockNum]; !exists {
		t.Fatalf("Not enough length for block num: %v", blockNum)
	}

	for i := int64(0); i < blockNum; i++ {
		forked.Heads[i] = b.Heads[i]
	}

	forked.Heads[blockNum].ParentHash = b.Heads[blockNum].ParentHash
	forked.Heads[blockNum].Parent.Store(b.Heads[blockNum].Parent.Load())
	return forked
}

func (b *blocks) NewHead(number uint64) *evmtypes.Head {
	parentNumber := number - 1
	parent, ok := b.Heads[int64(parentNumber)]
	if !ok {
		b.t.Fatalf("Can't find parent block at index: %v", parentNumber)
	}
	head := &evmtypes.Head{
		Number:     parent.Number + 1,
		Hash:       testutils.NewHash(),
		ParentHash: parent.Hash,
		Timestamp:  parent.Timestamp.Add(time.Second),
		EVMChainID: ubig.New(testutils.FixtureChainID),
	}
	head.Parent.Store(parent)
	return head
}

// CountingHeadTrackable allows you to count callbacks
type CountingHeadTrackable struct {
	onNewHeadCount atomic.Int32
}

// OnNewLongestChain increases the OnNewLongestChainCount count by one
func (m *CountingHeadTrackable) OnNewLongestChain(context.Context, *evmtypes.Head) {
	m.onNewHeadCount.Add(1)
}

// OnNewLongestChainCount returns the count of new heads, safely.
func (m *CountingHeadTrackable) OnNewLongestChainCount() int32 {
	return m.onNewHeadCount.Load()
}
