package writetarget_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	wt "github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/platform"
	"github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/platform/on-chain/forwarder"
	"github.com/smartcontractkit/chainlink-evm/pkg/writetarget"
	"github.com/smartcontractkit/chainlink-evm/pkg/writetarget/mocks"
	"github.com/test-go/testify/mock"

	"github.com/stretchr/testify/require"
)

func TestWriteTargetMonitor(t *testing.T) {
	processor := mocks.NewProductSpecificProcessor(t)
	processor.On("SetEmitter", mock.Anything).Return(nil).Once()

	monitor, err := writetarget.NewMonitor(logger.Test(t), []writetarget.ProductSpecificProcessor{processor})
	require.NoError(t, err)

	encoded := forwarder.NewTestReport(t)

	msg := &wt.WriteConfirmed{
		BlockHeight:             "10",
		MetaCapabilityProcessor: "test",
		Report:                  encoded,
	}

	t.Run("Uses processor when name equals config", func(t *testing.T) {
		processor.On("GetName").Return("test").Once()
		processor.On("Process", t.Context(), msg, mock.Anything).Return(nil).Once()
		monitor.ProtoEmitter.EmitWithLog(t.Context(), msg)
	})

	t.Run("Errors when config name is not found", func(t *testing.T) {
		processor.On("GetName").Return("other")
		err := monitor.ProtoEmitter.EmitWithLog(t.Context(), msg)
		require.Error(t, err)
	})

	t.Run("Does not use processor when none is configured", func(t *testing.T) {
		// get new processor
		processor = mocks.NewProductSpecificProcessor(t)
		msg.MetaCapabilityProcessor = ""
		processor.AssertNotCalled(t, "Process", mock.Anything, mock.Anything, mock.Anything)

		err := monitor.ProtoEmitter.EmitWithLog(t.Context(), msg)
		require.NoError(t, err)
	})
}
