package writetarget

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/monitor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestNewWriteTargetID(t *testing.T) {
	tests := []struct {
		name            string
		chainFamilyName string
		networkName     string
		chainID         string
		version         string
		expected        string
		expectError     bool
	}{
		{
			name:            "Valid input with network name",
			chainFamilyName: "aptos",
			networkName:     "mainnet",
			chainID:         "1",
			version:         "1.0.0",
			expected:        "write_aptos-mainnet@1.0.0",
			expectError:     false,
		},
		{
			name:            "Valid input without network name",
			chainFamilyName: "aptos",
			networkName:     "",
			chainID:         "1",
			version:         "1.0.0",
			expected:        "write_aptos-1@1.0.0",
			expectError:     false,
		},
		{
			name:            "Valid input with empty chainFamilyName",
			chainFamilyName: "",
			networkName:     "ethereum-mainnet",
			chainID:         "1",
			version:         "1.0.0",
			expected:        "write_ethereum-mainnet@1.0.0",
			expectError:     false,
		},
		{
			name:            "Invalid input with empty version",
			chainFamilyName: "aptos",
			networkName:     "mainnet",
			chainID:         "1",
			version:         "",
			expected:        "",
			expectError:     true,
		},
		{
			name:            "Invalid input with empty networkName and chainID",
			chainFamilyName: "aptos",
			networkName:     "",
			chainID:         "",
			version:         "2.0.0",
			expected:        "",
			expectError:     true,
		},
		{
			name:            "Valid input with unknown network name",
			chainFamilyName: "aptos",
			networkName:     "unknown",
			chainID:         "1",
			version:         "2.0.1",
			expected:        "write_aptos-1@2.0.1",
			expectError:     false,
		},
		{
			name:            "Valid input with network name (testnet)",
			chainFamilyName: "aptos",
			networkName:     "testnet",
			chainID:         "2",
			version:         "1.0.3",
			expected:        "write_aptos-testnet@1.0.3",
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewWriteTargetID(tt.chainFamilyName, tt.networkName, tt.chainID, tt.version)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}

// ReportParams defines the parameters for constructing a report
type ReportParams struct {
	ExecutionID   string // 32 bytes
	Timestamp     uint32 // 4 bytes
	DonID         uint32 // 4 bytes
	ConfigVersion uint32 // 4 bytes
	WorkflowID    string // 32 bytes
	WorkflowName  string // 10 bytes
	WorkflowOwner string // 20 bytes
	ReportID      uint16 // 2 bytes
}

// constructReportBytes creates a report byte array from human-readable inputs
func constructReportBytes(params ReportParams) ([]byte, error) {
	reportBytes := []byte{0x01} // version

	// Convert execution ID to bytes
	executionIDBytes, err := hex.DecodeString(params.ExecutionID)
	if err != nil {
		return nil, fmt.Errorf("invalid execution ID: %w", err)
	}
	if len(executionIDBytes) != 32 {
		return nil, fmt.Errorf("execution ID must be 32 bytes, got %d", len(executionIDBytes))
	}
	reportBytes = append(reportBytes, executionIDBytes...)

	// Add timestamp (4 bytes)
	timestampBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(timestampBytes, params.Timestamp)
	reportBytes = append(reportBytes, timestampBytes...)

	// Add don_id (4 bytes)
	donIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(donIDBytes, params.DonID)
	reportBytes = append(reportBytes, donIDBytes...)

	// Add config_version (4 bytes)
	configVersionBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(configVersionBytes, params.ConfigVersion)
	reportBytes = append(reportBytes, configVersionBytes...)

	// Convert workflow ID to bytes
	workflowIDBytes, err := hex.DecodeString(params.WorkflowID)
	if err != nil {
		return nil, fmt.Errorf("invalid workflow ID: %w", err)
	}
	if len(workflowIDBytes) != 32 {
		return nil, fmt.Errorf("workflow ID must be 32 bytes, got %d", len(workflowIDBytes))
	}
	reportBytes = append(reportBytes, workflowIDBytes...)

	// Add workflow name (10 bytes)
	if len(params.WorkflowName) > 10 {
		return nil, fmt.Errorf("workflow name must be 10 bytes or less, got %d", len(params.WorkflowName))
	}
	workflowNameBytes := make([]byte, 10)
	copy(workflowNameBytes, params.WorkflowName)
	reportBytes = append(reportBytes, workflowNameBytes...)

	// Add workflow owner (20 bytes)
	workflowOwnerBytes, err := hex.DecodeString(params.WorkflowOwner)
	if err != nil {
		return nil, fmt.Errorf("invalid workflow owner: %w", err)
	}
	if len(workflowOwnerBytes) != 20 {
		return nil, fmt.Errorf("workflow owner must be 20 bytes, got %d", len(workflowOwnerBytes))
	}
	reportBytes = append(reportBytes, workflowOwnerBytes...)

	// Add report ID (2 bytes)
	reportIDBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(reportIDBytes, params.ReportID)
	reportBytes = append(reportBytes, reportIDBytes...)

	return reportBytes, nil
}

// mockContractWriter implements types.ContractWriter for testing
type mockContractWriter struct {
	estimateFee types.EstimateFee
	err         error
}

func (m *mockContractWriter) GetEstimateFee(ctx context.Context, contract, method string, args any, toAddress string, meta *types.TxMeta, val *big.Int) (types.EstimateFee, error) {
	return m.estimateFee, m.err
}

func (m *mockContractWriter) SubmitTransaction(ctx context.Context, contractName, method string, args any, transactionID types.IdempotencyKey, toAddress string, meta *types.TxMeta, value *big.Int) error {
	return nil
}

func (m *mockContractWriter) GetTransactionStatus(ctx context.Context, transactionID types.IdempotencyKey) (types.TransactionStatus, error) {
	return types.Unknown, nil
}

func (m *mockContractWriter) GetFeeComponents(ctx context.Context) (*types.ChainFeeComponents, error) {
	return nil, nil
}

func (m *mockContractWriter) Close() error {
	return nil
}

func (m *mockContractWriter) HealthReport() map[string]error {
	return nil
}

func (m *mockContractWriter) Name() string {
	return "mock-contract-writer"
}

func (m *mockContractWriter) Ready() error {
	return nil
}

func (m *mockContractWriter) Start(ctx context.Context) error {
	return nil
}

// mockProtoEmitter implements monitor.ProtoEmitter for testing
type mockProtoEmitter struct{}

func (m *mockProtoEmitter) Emit(ctx context.Context, msg proto.Message, attrKVs ...any) error {
	return nil
}

func (m *mockProtoEmitter) EmitWithLog(ctx context.Context, msg proto.Message, attrKVs ...any) error {
	return nil
}

// mockTargetStrategy implements TargetStrategy for testing
type mockTargetStrategy struct{}

func (m *mockTargetStrategy) QueryTransmissionState(ctx context.Context, reportID uint16, request capabilities.CapabilityRequest) (*TransmissionState, error) {
	return &TransmissionState{
		Status: TransmissionStateNotAttempted,
	}, nil
}

func (m *mockTargetStrategy) TransmitReport(ctx context.Context, report []byte, reportContext []byte, signatures [][]byte, request capabilities.CapabilityRequest) (string, error) {
	return "test-tx-id", nil
}

// mockChainService implements types.ChainService for testing
type mockChainService struct{}

func (m *mockChainService) LatestHead(ctx context.Context) (types.Head, error) {
	return types.Head{}, nil // return a dummy head
}

func (m *mockChainService) Close() error {
	return nil
}

func (m *mockChainService) GetChainStatus(ctx context.Context) (types.ChainStatus, error) {
	return types.ChainStatus{}, nil // return a dummy status
}

func (m *mockChainService) HealthReport() map[string]error {
	return map[string]error{}
}

func (m *mockChainService) ListNodeStatuses(ctx context.Context, pageSize int32, pageToken string) (stats []types.NodeStatus, nextPageToken string, total int, err error) {
	return nil, "", 0, nil
}

func (m *mockChainService) Name() string {
	return "mock-chain-service"
}

func (m *mockChainService) Ready() error {
	return nil
}

func (m *mockChainService) Replay(ctx context.Context, fromBlock string, args map[string]interface{}) error {
	return nil
}

func (m *mockChainService) Start(ctx context.Context) error {
	return nil
}

func (m *mockChainService) Transact(ctx context.Context, from string, to string, amount *big.Int, balanceCheck bool) error {
	return nil
}

func TestExecuteGasEstimation(t *testing.T) {
	tests := []struct {
		name          string
		spendLimit    string // in ETH
		estimateFee   types.EstimateFee
		estimateErr   error
		expectErr     bool
		expectErrMsg  string
		expectedSpend string // expected spend value in response
	}{
		{
			name:       "successful estimation within limit",
			spendLimit: "0.1", // 0.1 ETH
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(50000000000000000), // 0.05 ETH in wei
				Decimals: 18,
			},
			expectErr:     false,
			expectedSpend: "0.050000000000000000",
		},
		{
			name:       "fee exceeds limit",
			spendLimit: "0.1", // 0.1 ETH
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(150000000000000000), // 0.15 ETH in wei
				Decimals: 18,
			},
			expectErr:    true,
			expectErrMsg: "estimated gas fee 150000000000000000 exceeds spend limit 100000000000000000",
		},
		{
			name:       "invalid spend limit format",
			spendLimit: "not-a-number",
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(50000000000000000),
				Decimals: 18,
			},
			expectErr:    true,
			expectErrMsg: "invalid gas spend limit format: not-a-number",
		},
		{
			name:         "estimate fee error",
			spendLimit:   "0.1",
			estimateErr:  assert.AnError,
			expectErr:    true,
			expectErrMsg: "failed to get gas estimate: assert.AnError general error for testing",
		},
		{
			name:       "zero spend limit",
			spendLimit: "0",
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(50000000000000000),
				Decimals: 18,
			},
			expectErr:    true,
			expectErrMsg: "estimated gas fee 50000000000000000 exceeds spend limit 0",
		},
		{
			name:       "very small spend limit",
			spendLimit: "0.000000000000000001", // 1 wei in ETH
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(2), // 2 wei
				Decimals: 18,
			},
			expectErr:    true,
			expectErrMsg: "estimated gas fee 2 exceeds spend limit 1",
		},
		{
			name:       "very large spend limit",
			spendLimit: "1000000", // 1M ETH
			estimateFee: types.EstimateFee{
				Fee:      func() *big.Int { b, _ := new(big.Int).SetString("1000000000000000000000", 10); return b }(), // 1000 ETH in wei
				Decimals: 18,
			},
			expectErr:     false,
			expectedSpend: "1000.000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock contract writer
			mockCW := &mockContractWriter{
				estimateFee: tt.estimateFee,
				err:         tt.estimateErr,
			}

			// Create write target with mock
			wt := &writeTarget{
				CapabilityInfo: capabilities.CapabilityInfo{
					ID:             "test-write-target",
					CapabilityType: capabilities.CapabilityTypeTarget,
				},
				chainInfo: monitor.ChainInfo{
					ChainID: "1",
				},
				cw:   mockCW,
				lggr: logger.Nop(),
				beholder: &monitor.BeholderClient{
					Client:       beholder.GetClient(),
					ProtoEmitter: &mockProtoEmitter{},
				},
				cs: &mockChainService{},
				configValidateFn: func(request capabilities.CapabilityRequest) (string, error) {
					return "0x123", nil
				},
				nodeAddress:      "0x456",
				forwarderAddress: "0x789",
				targetStrategy:   &mockTargetStrategy{},
			}

			// Use a 32-byte execution ID for both the report and the request
			executionID := "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
			workflowID := "2122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f40"

			// Construct report bytes using the helper function
			reportBytes, err := constructReportBytes(ReportParams{
				ExecutionID:   executionID,
				Timestamp:     0x66f5bf69,
				DonID:         1,
				ConfigVersion: 1,
				WorkflowID:    workflowID,
				WorkflowName:  "0000FOOBAR",
				WorkflowOwner: "00000000000000000000000000000000000000aa",
				ReportID:      1,
			})
			require.NoError(t, err)

			inputsMap, err := values.NewMap(map[string]any{
				KeySignedReport: map[string]any{
					"ID":         []byte{0, 1}, // reportID = 1
					"Report":     reportBytes,
					"Context":    []byte{4, 5, 6},
					"Signatures": [][]byte{{7, 8, 9}},
				},
			})
			require.NoError(t, err)
			request := capabilities.CapabilityRequest{
				Metadata: capabilities.RequestMetadata{
					SpendLimits: []capabilities.SpendLimit{
						{
							SpendType: "GAS.1",
							Limit:     tt.spendLimit,
						},
					},
					WorkflowExecutionID: executionID,
					WorkflowID:          workflowID,
				},
				Inputs: inputsMap,
			}

			// Test Execute
			response, err := wt.Execute(context.Background(), request)
			if tt.expectErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectErrMsg)
				return
			}

			require.NoError(t, err)
			require.Len(t, response.Metadata.Metering, 1)
			assert.Equal(t, "GAS.1", response.Metadata.Metering[0].SpendUnit)
			assert.Equal(t, tt.expectedSpend, response.Metadata.Metering[0].SpendValue)
		})
	}
}

func TestMeteredSpendReporting(t *testing.T) {
	tests := []struct {
		name          string
		estimateFee   types.EstimateFee
		expectedSpend string
	}{
		{
			name: "exact wei amount",
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(1000000000000000000), // 1 ETH in wei
				Decimals: 18,
			},
			expectedSpend: "1.000000000000000000",
		},
		{
			name: "fractional wei amount",
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(123456789000000000), // 0.123456789 ETH in wei
				Decimals: 18,
			},
			expectedSpend: "0.123456789000000000",
		},
		{
			name: "very small amount",
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(1), // 1 wei
				Decimals: 18,
			},
			expectedSpend: "0.000000000000000001",
		},
		{
			name: "very large amount",
			estimateFee: types.EstimateFee{
				Fee:      func() *big.Int { b, _ := new(big.Int).SetString("1000000000000000000000", 10); return b }(), // 1000 ETH in wei
				Decimals: 18,
			},
			expectedSpend: "1000.000000000000000000",
		},
		{
			name: "zero amount",
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(0),
				Decimals: 18,
			},
			expectedSpend: "0.000000000000000000",
		},
		{
			name: "different decimals",
			estimateFee: types.EstimateFee{
				Fee:      big.NewInt(1000000), // 1.0 with 6 decimals
				Decimals: 6,
			},
			expectedSpend: "1.000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock contract writer
			mockCW := &mockContractWriter{
				estimateFee: tt.estimateFee,
			}

			// Create write target with mock
			wt := &writeTarget{
				CapabilityInfo: capabilities.CapabilityInfo{
					ID:             "test-write-target",
					CapabilityType: capabilities.CapabilityTypeTarget,
				},
				chainInfo: monitor.ChainInfo{
					ChainID: "1",
				},
				cw:   mockCW,
				lggr: logger.Nop(),
				beholder: &monitor.BeholderClient{
					Client:       beholder.GetClient(),
					ProtoEmitter: &mockProtoEmitter{},
				},
				cs: &mockChainService{},
				configValidateFn: func(request capabilities.CapabilityRequest) (string, error) {
					return "0x123", nil
				},
				nodeAddress:      "0x456",
				forwarderAddress: "0x789",
				targetStrategy:   &mockTargetStrategy{},
			}

			// Create request with high spend limit to ensure it passes
			executionID := "816d80fa4bb8b350cacc3a2e395236bcc6b813b0568b61eea0bd3e6ba7218dd3"
			workflowID := "bc06f300e797d5a8575637a14aae13e3f8508008d1fc54f4c4611fff17a68cb0"

			// Construct report bytes using the helper function
			reportBytes, err := constructReportBytes(ReportParams{
				ExecutionID:   executionID,
				Timestamp:     0x66f5bf69,
				DonID:         1,
				ConfigVersion: 1,
				WorkflowID:    workflowID,
				WorkflowName:  "0000FOOBAR",
				WorkflowOwner: "00000000000000000000000000000000000000aa",
				ReportID:      1,
			})
			require.NoError(t, err)

			inputsMap, err := values.NewMap(map[string]any{
				KeySignedReport: map[string]any{
					"ID":         []byte{0, 1}, // reportID = 1
					"Report":     reportBytes,
					"Context":    []byte{4, 5, 6},
					"Signatures": [][]byte{{7, 8, 9}},
				},
			})
			require.NoError(t, err)
			request := capabilities.CapabilityRequest{
				Metadata: capabilities.RequestMetadata{
					SpendLimits: []capabilities.SpendLimit{
						{
							SpendType: "GAS.1",
							Limit:     "1000000", // 1M ETH
						},
					},
					WorkflowExecutionID: executionID,
					WorkflowID:          workflowID,
				},
				Inputs: inputsMap,
			}

			// Test Execute
			response, err := wt.Execute(context.Background(), request)
			require.NoError(t, err)

			// Verify metering details
			require.Len(t, response.Metadata.Metering, 1)
			metering := response.Metadata.Metering[0]
			assert.Equal(t, "GAS.1", metering.SpendUnit)
			assert.Equal(t, tt.expectedSpend, metering.SpendValue)
			assert.Empty(t, metering.Peer2PeerID) // Should be unset as per requirements
		})
	}
}
