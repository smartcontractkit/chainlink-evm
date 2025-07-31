package bindings_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
	"github.com/smartcontractkit/cre-sdk-go/cre/testutils"
	consensusmock "github.com/smartcontractkit/cre-sdk-go/internal_testing/capabilities/consensus/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	ocr3types "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
	pb2 "github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk/v2/pb"
	datastorage "github.com/smartcontractkit/chainlink-evm/pkg/bindings/testdata"
)

const anyChainSelector = uint64(1337)

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
	client := &evm.Client{ChainSelector: anyChainSelector}
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	evmCap, err := evmmock.NewClientCapability(anyChainSelector, t)
	require.NoError(t, err, "Failed to create EVM client capability")

	evmCap.HeaderByNumber = func(_ context.Context, input *evm.HeaderByNumberRequest) (*evm.HeaderByNumberReply, error) {
		header := &evm.HeaderByNumberReply{
			Header: &evm.Header{
				BlockNumber: pb2.NewBigIntFromInt(big.NewInt(123456)),
			},
		}
		return header, nil
	}

	evmCap.CallContract = func(_ context.Context, input *evm.CallContractRequest) (*evm.CallContractReply, error) {
		// Simulate a successful call with dummy data
		reply := &evm.CallContractReply{
			Data: []byte{0x01, 0x02, 0x03, 0x04}, // Example data
		}
		return reply, nil
	}

	runtime, _ := testutils.NewRuntimeAndEnv(t, "", map[string]string{})
	reply := ds.ReadData(runtime, datastorage.ReadDataInput{
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
	client := &evm.Client{ChainSelector: anyChainSelector}
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	report := ocr3types.Metadata{
		Version:          1,
		ExecutionID:      "1234567890123456789012345678901234567890123456789012345678901234",
		Timestamp:        1620000000,
		DONID:            1,
		DONConfigVersion: 1,
		WorkflowID:       "1234567890123456789012345678901234567890123456789012345678901234",
		WorkflowName:     "12",
		WorkflowOwner:    "1234567890123456789012345678901234567890",
		ReportID:         "1234",
	}

	rawReport, err := report.Encode()
	require.NoError(t, err)

	consensusCap, err := consensusmock.NewConsensusCapability(t)
	require.NoError(t, err, "Failed to create Consensus capability")
	consensusCap.Report = func(_ context.Context, input *pb.ReportRequest) (*pb.ReportResponse, error) {
		return &pb.ReportResponse{
			RawReport: rawReport,
		}, nil
	}

	evmCap, err := evmmock.NewClientCapability(anyChainSelector, t)
	require.NoError(t, err, "Failed to create EVM client capability")
	evmCap.WriteReport = func(_ context.Context, req *evm.WriteReportRequest) (*evm.WriteReportReply, error) {
		require.Equal(t, rawReport, req.Report.RawReport)
		return &evm.WriteReportReply{
			TxStatus: evm.TxStatus_TX_STATUS_SUCCESS,
			TxHash:   []byte{0x01, 0x02, 0x03, 0x04},
		}, nil
	}

	runtime, _ := testutils.NewRuntimeAndEnv(t, "", map[string]string{})

	reply := ds.WriteReportDataStorageUserData(runtime, datastorage.DataStorageUserData{
		Key:   "testKey",
		Value: "testValue",
	}, nil)
	require.NoError(t, err, "WriteReportDataStorageUserData should not return an error")
	response, err := reply.Await()
	require.NoError(t, err, "Awaiting WriteReportDataStorageUserData reply should not return an error")
	require.NotNil(t, response, "Response from WriteReportDataStorageUserData should not be nil")
	require.True(t, proto.Equal(&evm.WriteReportReply{
		TxStatus: evm.TxStatus_TX_STATUS_SUCCESS,
		TxHash:   []byte{0x01, 0x02, 0x03, 0x04},
	}, response), "Response should match expected WriteReportReply")
}

func TestEncodeStruct(t *testing.T) {
	ds := newDataStorage(t)

	str := datastorage.DataStorageUpdateReserves{
		TotalMinted:  big.NewInt(100),
		TotalReserve: big.NewInt(200),
	}

	encoded, err := ds.Codec.EncodeDataStorageUpdateReservesStruct(str)
	require.NoError(t, err, "Encoding DataStorageUpdateReserves should not return an error")
	require.NotNil(t, encoded, "Encoded data should not be nil")
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
	client := &evm.Client{ChainSelector: anyChainSelector}
	anyAddress := common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2").Bytes()
	ds, err := datastorage.NewDataStorage(client, anyAddress, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	evmCap, err := evmmock.NewClientCapability(anyChainSelector, t)
	require.NoError(t, err, "Failed to create EVM client capability")
	evmCap.RegisterLogTracking = func(_ context.Context, req *evm.RegisterLogTrackingRequest) (*emptypb.Empty, error) {
		require.Equal(t, req.Filter.Name, "AccessLogged-"+common.Bytes2Hex(ds.Address))
		require.ElementsMatch(t, [][]byte{ds.Address}, req.Filter.Addresses)
		require.ElementsMatch(t, [][]byte{ds.Codec.AccessLoggedLogHash()}, req.Filter.EventSigs)
		require.Len(t, req.Filter.Topic2, 1)
		require.Equal(t, req.Filter.Topic2[0], common.HexToHash("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2").Bytes())
		return &emptypb.Empty{}, nil
	}

	evmCap.UnregisterLogTracking = func(ctx context.Context, req *evm.UnregisterLogTrackingRequest) (*emptypb.Empty, error) {
		require.Equal(t, req.FilterName, "AccessLogged-"+common.Bytes2Hex(ds.Address))
		return &emptypb.Empty{}, nil
	}

	runtime, _ := testutils.NewRuntimeAndEnv(t, "", map[string]string{})

	register := ds.RegisterLogTrackingAccessLogged(runtime, &bindings.LogTrackingOptions[datastorage.AccessLogged]{
		Filters: []datastorage.AccessLogged{
			{
				Caller: common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			},
		},
	})
	_, err = register.Await()
	require.NoError(t, err)

	_, err = ds.UnregisterLogTrackingAccessLogged(runtime).Await()
	require.NoError(t, err)
}

func TestFilterLogs(t *testing.T) {
	client := &evm.Client{ChainSelector: anyChainSelector}
	anyAddress := common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2").Bytes()
	ds, err := datastorage.NewDataStorage(client, anyAddress, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")

	bh := []byte{0x01, 0x02, 0x03, 0x04}
	fb := big.NewInt(100)
	tb := big.NewInt(200)

	evmCap, err := evmmock.NewClientCapability(anyChainSelector, t)
	require.NoError(t, err, "Failed to create EVM client capability")
	evmCap.FilterLogs = func(_ context.Context, req *evm.FilterLogsRequest) (*evm.FilterLogsReply, error) {
		require.Equal(t, [][]byte{ds.Address}, req.FilterQuery.Addresses, "Filter should contain the correct address")
		require.Equal(t, bh, req.FilterQuery.BlockHash, "Filter should contain the correct block hash")
		require.Equal(t, fb.Bytes(), req.FilterQuery.FromBlock.GetAbsVal(), "Filter should contain the correct from block")
		require.Equal(t, tb.Bytes(), req.FilterQuery.ToBlock.GetAbsVal(), "Filter should contain the correct to block")
		logs := []*evm.Log{
			{
				Address: ds.Address,
				Topics:  [][]byte{ds.Codec.AccessLoggedLogHash()},
				Data:    []byte("test log data"),
			},
		}
		return &evm.FilterLogsReply{Logs: logs}, nil
	}

	runtime, _ := testutils.NewRuntimeAndEnv(t, "", map[string]string{})

	reply := ds.FilterLogsAccessLogged(runtime, &bindings.FilterOptions{
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
	client := &evm.Client{ChainSelector: anyChainSelector}
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")
	t.Run("simple event", func(t *testing.T) {
		ev := ds.ABI.Events["DataStored"]
		events := []datastorage.DataStored{
			{
				Sender: common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				Key:    "testKey",
				Value:  "testValue",
			},
			{
				Sender: common.HexToAddress("0xBb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
				Key:    "testKey",
				Value:  "testValue",
			},
		}

		encoded, err := ds.Codec.EncodeDataStoredTopics(ev, events)
		require.NoError(t, err, "Encoding DataStored topics should not return an error")

		require.Equal(t, ds.Codec.DataStoredLogHash(), encoded[0].Values[0], "First topic value should be AccessLogged log hash")
		require.Len(t, encoded[1].Values, 2, "Second topic should have two values")
		expected1, err := abi.Arguments{ev.Inputs[0]}.Pack(events[0].Sender)
		require.NoError(t, err)
		require.Equal(t, expected1, encoded[1].Values[0])
		expected2, err := abi.Arguments{ev.Inputs[0]}.Pack(events[1].Sender)
		require.NoError(t, err)
		require.Equal(t, expected2, encoded[1].Values[1])

		trigger, err := ds.LogTriggerDataStoredLog(1, evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED, events)
		require.NotNil(t, trigger)
		require.NoError(t, err)
	})
	t.Run("dynamic event", func(t *testing.T) {
		ev := ds.ABI.Events["DynamicEvent"]
		events := []datastorage.DynamicEvent{
			{
				Key: "testKey1",
				UserData: datastorage.DataStorageUserData{
					Key:   "userKey1",
					Value: "userValue1",
				},
				Sender:   "testSender1",
				Metadata: common.HexToHash("metadata1"),
				MetadataArray: [][]byte{
					[]byte("meta1"),
					[]byte("meta2"),
				},
			},
			{
				Key: "testKey2",
				UserData: datastorage.DataStorageUserData{
					Key:   "userKey2",
					Value: "userValue2",
				},
				Sender:   "testSender2",
				Metadata: common.HexToHash("metadata2"),
				MetadataArray: [][]byte{
					[]byte("meta3"),
					[]byte("meta4"),
				},
			},
		}

		encoded, err := ds.Codec.EncodeDynamicEventTopics(ev, events)
		require.NoError(t, err, "Encoding DynamicEvent topics should not return an error")

		require.Len(t, encoded, 4, "Trigger should have four topics")
		require.Equal(t, ds.Codec.DynamicEventLogHash(), encoded[0].Values[0], "First topic value should be DynamicEvent log hash")
		require.Len(t, encoded[1].Values, 2, "Second topic should have two values")
		packed1, err := abi.Arguments{ev.Inputs[1]}.Pack(events[0].UserData)

		expected1 := crypto.Keccak256(packed1)
		require.NoError(t, err)
		require.Equal(t, expected1, encoded[1].Values[0])
		packed2, err := abi.Arguments{ev.Inputs[1]}.Pack(events[1].UserData)

		expected2 := crypto.Keccak256(packed2)
		require.NoError(t, err)
		require.Equal(t, expected2, encoded[1].Values[1])

		expected3 := events[0].Metadata.Bytes()
		require.Equal(t, expected3, encoded[2].Values[0])

		expected4 := events[1].Metadata.Bytes()
		require.Equal(t, expected4, encoded[2].Values[1])

		packed3, err := abi.Arguments{ev.Inputs[4]}.Pack(events[0].MetadataArray)
		expected5 := crypto.Keccak256(packed3)
		require.NoError(t, err)
		require.Equal(t, expected5, encoded[3].Values[0])

		packed4, err := abi.Arguments{ev.Inputs[4]}.Pack(events[1].MetadataArray)
		require.NoError(t, err)
		expected6 := crypto.Keccak256(packed4)
		require.Equal(t, expected6, encoded[3].Values[1])

		trigger, err := ds.LogTriggerDynamicEventLog(1, evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED, events)
		require.NotNil(t, trigger)
		require.NoError(t, err)
	})
}

func newDataStorage(t *testing.T) *datastorage.DataStorage {
	client := &evm.Client{ChainSelector: anyChainSelector}
	ds, err := datastorage.NewDataStorage(client, nil, &bindings.ContractInitOptions{})
	require.NoError(t, err, "Failed to create DataStorage instance")
	return ds
}
