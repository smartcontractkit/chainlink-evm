package bindings_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/sdk"

	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"

	"github.com/smartcontractkit/chainlink-evm/pkg/bindings"
	"github.com/smartcontractkit/chainlink-evm/pkg/bindings/mocks"
	datastorage "github.com/smartcontractkit/chainlink-evm/pkg/bindings/testdata"
)

func TestGenerateBindings(t *testing.T) {
	err := bindings.GenerateBindings(
		"./testdata/DataStorage_combined.json",
		"",
		"bindings",
		"",
		"./testdata/bindings.go",
	)
	require.NoError(t, err, "Failed to generate bindings from combined JSON")
}

func TestGeneratedBindingsCodec(t *testing.T) {
	ds := newDataStorage(t)

	t.Run("encode functions", func(t *testing.T) {
		// structs
		userData := datastorage.DataStorageUserData{
			Key:   "testKey",
			Value: "testValue",
		}

		_, err := ds.Codec.EncodeDataStorageUserDataStruct(userData)
		require.NoError(t, err)

		// inputs
		logAccess := datastorage.LogAccessInput{
			Message: "testMessage",
		}
		_, err = ds.Codec.EncodeLogAccessMethodCall(logAccess)
		require.NoError(t, err)

		onReport := datastorage.OnReportInput{
			Metadata: []byte("testMetadata"),
			Payload:  []byte("testPayload"),
		}
		_, err = ds.Codec.EncodeOnReportMethodCall(onReport)
		require.NoError(t, err)

		readData := datastorage.ReadDataInput{
			User: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			Key:  "testKey",
		}
		_, err = ds.Codec.EncodeReadDataMethodCall(readData)
		require.NoError(t, err)

		storeData := datastorage.StoreDataInput{
			Key:   "testKey",
			Value: "testValue",
		}
		_, err = ds.Codec.EncodeStoreDataMethodCall(storeData)
		require.NoError(t, err)

		storeUserData := datastorage.StoreUserDataInput{
			UserData: userData,
		}
		_, err = ds.Codec.EncodeStoreUserDataMethodCall(storeUserData)
		require.NoError(t, err)

		updateDataInput := datastorage.UpdateDataInput{
			Key:      "testKey",
			NewValue: "newTestValue",
		}
		_, err = ds.Codec.EncodeUpdateDataMethodCall(updateDataInput)
		require.NoError(t, err)
	})
}

func TestDecodeEvents(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ds := newDataStorage(t)

		caller := common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2")
		message := "Test access log"

		ev := ds.ABI.Events["AccessLogged"]

		topics := [][]byte{
			ds.Codec.AccessLoggedLogHash(),
			caller.Bytes(),
		}

		var nonIndexed abi.Arguments
		for _, arg := range ev.Inputs {
			if !arg.Indexed {
				nonIndexed = append(nonIndexed, arg)
			}
		}
		data, err := nonIndexed.Pack(message)
		require.NoError(t, err)

		log := &evm.Log{
			Topics: topics,
			Data:   data,
		}

		out, err := ds.Codec.DecodeAccessLogged(log)
		require.NoError(t, err)
		require.Equal(t, caller, out.Caller)
		require.Equal(t, message, out.Message)
	})
}

func TestReadMethods(t *testing.T) {
	client := mocks.NewEVMClient(t)
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	client.EXPECT().LatestAndFinalizedHead(mock.Anything, mock.Anything).Return(
		sdk.NewBasicPromise(func() (*evm.LatestAndFinalizedHeadReply, error) {
			// Simulate a successful call with dummy data
			reply := &evm.LatestAndFinalizedHeadReply{
				Finalized: &evm.Head{
					BlockNumber: pb.NewBigIntFromInt(big.NewInt(123)),
				},
			}
			return reply, nil
		})).Once()

	client.EXPECT().CallContract(mock.Anything, mock.Anything).Return(
		sdk.NewBasicPromise(func() (*evm.CallContractReply, error) {
			// Simulate a successful call with dummy data
			reply := &evm.CallContractReply{
				Data: []byte{0x01, 0x02, 0x03, 0x04}, // Example data
			}
			return reply, nil
		})).Once()

	reply, err := ds.ReadData(nil, datastorage.ReadDataInput{
		User: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		Key:  "testKey",
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, reply, "ReadData should return a non-nil reply")

	response, err := reply.Await()
	require.NoError(t, err, "Awaiting ReadData reply should not return an error")
	require.NotNil(t, response, "Response from ReadData should not be nil")
	require.Equal(t, []byte{0x01, 0x02, 0x03, 0x04}, response.Data, "Response data should match expected dummy data")
}

func TestWriteReportMethods(t *testing.T) {
	client := mocks.NewEVMClient(t)
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	client.EXPECT().WriteReport(mock.Anything, mock.Anything).Return(
		sdk.NewBasicPromise(func() (*evmcappb.WriteReportReply, error) {
			// Simulate a successful write report
			return &evmcappb.WriteReportReply{
				TxStatus: evmcappb.TxStatus_TX_SUCCESS,
				TxHash:   []byte{0x01, 0x02, 0x03, 0x04},
			}, nil
		})).Once()

	reply, err := ds.WriteReportDataStorageUserData(nil, datastorage.DataStorageUserData{
		Key:   "testKey",
		Value: "testValue",
	}, nil)
	require.NoError(t, err, "WriteReportDataStorageUserData should not return an error")
	response, err := reply.Await()
	require.NoError(t, err, "Awaiting WriteReportDataStorageUserData reply should not return an error")
	require.NotNil(t, response, "Response from WriteReportDataStorageUserData should not be nil")
	require.Equal(t, &evmcappb.WriteReportReply{
		TxStatus: evmcappb.TxStatus_TX_SUCCESS,
		TxHash:   []byte{0x01, 0x02, 0x03, 0x04},
	}, response, "Response should match expected WriteReportReply")
}

func TestErrorHandling(t *testing.T) {
	ds := newDataStorage(t)

	requester := common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2")
	key := "testKey"
	reason := "not found"

	t.Run("valid", func(t *testing.T) {
		errDesc := ds.ABI.Errors["DataNotFound"]
		encoded := errDesc.ID.Bytes()[:4]
		args, err := errDesc.Inputs.Pack(requester, key, reason)
		require.NoError(t, err)
		encoded = append(encoded, args...)

		unpacked, err := ds.UnpackError(encoded)
		require.NoError(t, err)

		result, ok := unpacked.(*datastorage.DataNotFound)
		require.True(t, ok, "Unpacked error should be of type DataNotFoundError")

		require.Equal(t, requester, result.Requester)
		require.Equal(t, key, result.Key)
		require.Equal(t, reason, result.Reason)
	})

	t.Run("invalid", func(t *testing.T) {
		// Simulate an invalid error code
		invalidCode := []byte{0x01, 0x02, 0x03, 0x04}
		_, err := ds.UnpackError(invalidCode)
		require.Error(t, err, "Unpacking an invalid error code should return an error")
		require.Contains(t, err.Error(), "unknown error selector", "Error message should indicate unknown error code")
	})
}

func TestRegisterUnregisterLogTracking(t *testing.T) {
	client := mocks.NewEVMClient(t)
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	client.
		EXPECT().
		RegisterLogTracking(mock.Anything, mock.Anything).
		Run(func(_ sdk.Runtime, req *evm.RegisterLogTrackingRequest) {
			require.Equal(t, req.Filter.Name, "AccessLogged-"+common.Bytes2Hex(ds.Address))
			require.Equal(t, [][]byte{ds.Address}, req.Filter.Addresses)
			require.Equal(t, [][]byte{ds.Codec.AccessLoggedLogHash()}, req.Filter.EventSigs)
		}).
		Return().
		Once()

	client.
		EXPECT().
		UnregisterLogTracking(mock.Anything, mock.Anything).
		Run(func(_ sdk.Runtime, req *evm.UnregisterLogTrackingRequest) {
			require.Equal(t, req.FilterName, "AccessLogged-"+common.Bytes2Hex(ds.Address))
		}).
		Return().
		Once()

	ds.RegisterLogTrackingAccessLogged(mocks.NewRuntime(t), &bindings.LogTrackingOptions{})
	ds.UnregisterLogTrackingAccessLogged(mocks.NewRuntime(t))
}

func TestFilterLogs(t *testing.T) {
	client := mocks.NewEVMClient(t)
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	bh := []byte{0x01, 0x02, 0x03, 0x04}
	fb := big.NewInt(100)
	tb := big.NewInt(200)

	// Mock the client to return a successful response
	client.EXPECT().FilterLogs(mock.Anything, mock.Anything).Run(
		func(_ sdk.Runtime, req *evm.FilterLogsRequest) {
			require.Equal(t, [][]byte{ds.Address}, req.FilterQuery.Addresses, "Filter should contain the correct address")
			require.Equal(t, bh, req.FilterQuery.BlockHash, "Filter should contain the correct block hash")
			require.Equal(t, fb.Bytes(), req.FilterQuery.FromBlock.GetAbsVal(), "Filter should contain the correct from block")
			require.Equal(t, tb.Bytes(), req.FilterQuery.ToBlock.GetAbsVal(), "Filter should contain the correct to block")
		}).Return(
		sdk.NewBasicPromise(func() (*evm.FilterLogsReply, error) {
			logs := []*evm.Log{
				{
					Address: ds.Address,
					Topics:  [][]byte{ds.Codec.AccessLoggedLogHash()},
					Data:    []byte("test log data"),
				},
			}
			return &evm.FilterLogsReply{Logs: logs}, nil
		})).Once()

	reply := ds.FilterLogsAccessLogged(mocks.NewRuntime(t), &bindings.FilterOptions{
		BlockHash: bh,
		FromBlock: fb,
		ToBlock:   tb,
	})
	response, err := reply.Await()
	require.NoError(t, err, "Awaiting FilteredLogsAccessLogged reply should not return an error")
	require.NotNil(t, response, "Response from FilteredLogsAccessLogged should not be nil")
	require.Len(t, response.Logs, 1, "Response should contain one log")
	require.Equal(t, ds.Address, response.Logs[0].Address)
}

func TestLogTrigger(t *testing.T) {
	client := mocks.NewEVMClient(t)
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	ev := ds.ABI.Events["AccessLogged"]
	event := datastorage.AccessLogged{
		Caller:  common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		Message: "Test access log",
	}

	client.EXPECT().LogTrigger(mock.Anything).Run(
		func(req *evm.FilterLogTriggerRequest) {
			require.Equal(t, [][]byte{ds.Address}, req.Addresses)
			require.Len(t, req.Topics, 1, "Trigger should have one topic")
			require.Equal(t, evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED, req.Confidence)
			for _, topic := range req.Topics {
				require.Len(t, topic.Values, 2, "Each topic should have two values")
				require.Equal(t, ds.Codec.AccessLoggedLogHash(), topic.Values[0], "First topic value should be AccessLogged log hash")
				expected, err := abi.Arguments{ev.Inputs[0]}.Pack(event.Caller)
				require.NoError(t, err, "packing caller")
				require.Equal(t, expected, topic.Values[1], "Second topic value should be the caller address")
			}
		}).Return(nil).Once()

	_, err = ds.LogTriggerAccessLoggedLog(evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED, []datastorage.AccessLogged{event})
	require.NoError(t, err)
}

func newDataStorage(t *testing.T) *datastorage.DataStorage {
	client := mocks.NewEVMClient(t)
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")
	return ds
}
