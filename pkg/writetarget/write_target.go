//nolint:gosec,revive // disable G115,revive
package writetarget

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/monitor"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/platform"
	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/retry"

	wt "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"
)

var (
	_ capabilities.TargetCapability = &writeTarget{}
)

type TransactionStatus uint8

// new chain agnostic transmission state types
const (
	TransmissionStateNotAttempted TransactionStatus = iota
	TransmissionStateSucceeded
	TransmissionStateFailed // retry
	TransmissionStateFatal  // don't retry
)

// alter TransmissionState to reference specific types rather than just
// success bool
type TransmissionState struct {
	Status      TransactionStatus
	Transmitter string
	Err         error
}

type TargetStrategy interface {
	// QueryTransmissionState defines how the report should be queried
	// via ChainReader, and how resulting errors should be classified.
	QueryTransmissionState(ctx context.Context, reportID uint16, request capabilities.CapabilityRequest) (*TransmissionState, error)
	// TransmitReport constructs the tx to transmit the report, and defines
	// any specific handling for sending the report via ChainWriter.
	TransmitReport(ctx context.Context, report []byte, reportContext []byte, signatures [][]byte, request capabilities.CapabilityRequest) (string, error)
}

var (
	_ capabilities.TargetCapability = &writeTarget{}
)

// chain-agnostic consts
const (
	CapabilityName = "write"

	// Input keys
	// Is this key chain agnostic?
	KeySignedReport = "signed_report"
)

type writeTarget struct {
	capabilities.CapabilityInfo

	config    Config
	chainInfo monitor.ChainInfo

	lggr logger.Logger
	// Local beholder client, also hosting the protobuf emitter
	beholder *monitor.BeholderClient

	cs               commontypes.ChainService
	cr               commontypes.ContractReader
	cw               commontypes.ContractWriter
	evm              commontypes.EVMService
	configValidateFn func(request capabilities.CapabilityRequest) (string, error)

	nodeAddress      string
	forwarderAddress string

	targetStrategy TargetStrategy
}
type WriteTargetOpts struct {
	ID string

	// toml: [<CHAIN>.WriteTargetCap]
	Config Config
	// ChainInfo contains the chain information (used as execution context)
	// TODO: simplify by passing via ChainService.GetChainStatus fn
	ChainInfo monitor.ChainInfo

	Logger   logger.Logger
	Beholder *monitor.BeholderClient

	ChainService     commontypes.ChainService
	ContractReader   commontypes.ContractReader
	ChainWriter      commontypes.ContractWriter
	EVMService       commontypes.EVMService
	ConfigValidateFn func(request capabilities.CapabilityRequest) (string, error)

	NodeAddress      string
	ForwarderAddress string

	TargetStrategy TargetStrategy
}

// Capability-specific configuration
type ReqConfig struct {
	Address string
}

// NewWriteTargetID returns the capability ID for the write target
func NewWriteTargetID(chainFamilyName, networkName, chainID, version string) (string, error) {
	// Input args should not be empty
	if version == "" {
		return "", fmt.Errorf("version must not be empty")
	}

	// Network ID: network name is optional, if not provided, use the chain ID
	networkID := networkName
	if networkID == "" && chainID == "" {
		return "", fmt.Errorf("invalid input: networkName or chainID must not be empty")
	}
	if networkID == "" || networkID == "unknown" {
		networkID = chainID
	}

	// allow for chain family to be empty
	if chainFamilyName == "" {
		return fmt.Sprintf("%s_%s@%s", CapabilityName, networkID, version), nil
	}

	return fmt.Sprintf("%s_%s-%s@%s", CapabilityName, chainFamilyName, networkID, version), nil
}

// TODO: opts.Config input is not validated for sanity
func NewWriteTarget(opts WriteTargetOpts) capabilities.TargetCapability {
	capInfo := capabilities.MustNewCapabilityInfo(opts.ID, capabilities.CapabilityTypeTarget, CapabilityName)

	return &writeTarget{
		capInfo,
		opts.Config,
		opts.ChainInfo,
		opts.Logger,
		opts.Beholder,
		opts.ChainService,
		opts.ContractReader,
		opts.ChainWriter,
		opts.EVMService,
		opts.ConfigValidateFn,
		opts.NodeAddress,
		opts.ForwarderAddress,
		opts.TargetStrategy,
	}
}

func success() capabilities.CapabilityResponse {
	return capabilities.CapabilityResponse{}
}

func (c *writeTarget) Execute(ctx context.Context, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
	// Take the local timestamp
	tsStart := time.Now().UnixMilli()

	// Trace the execution
	attrs := c.traceAttributes(request.Metadata.WorkflowExecutionID)
	ctx, span := c.beholder.Tracer.Start(ctx, "Execute", trace.WithAttributes(attrs...))
	defer span.End()

	// Notice: error skipped as implementation always returns nil
	capInfo, _ := c.Info(ctx)

	c.lggr.Debugw("Execute", "request", request, "capInfo", capInfo)

	// Helper to keep track of the request info
	info := &requestInfo{
		tsStart:   tsStart,
		node:      c.nodeAddress,
		forwarder: c.forwarderAddress,
		receiver:  "N/A",
		request:   request,
		reportInfo: &reportInfo{
			reportContext: nil,
			report:        nil,
			signersNum:    0, // N/A
			reportID:      0, // N/A
		},
		reportTransmissionState: nil,
	}
	// Helper to build monitoring (Beholder) messages
	builder := NewMessageBuilder(c.chainInfo, capInfo)

	// Validate the config
	receiver, err := c.configValidateFn(request)
	if err != nil {
		msg := builder.buildWriteError(info, 0, "failed to validate config", err.Error())
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}

	// Source the receiver address from the config
	info.receiver = receiver

	// Source the signed report from the request
	signedReport, ok := request.Inputs.Underlying[KeySignedReport]
	if !ok {
		cause := fmt.Sprintf("input missing required field: '%s'", KeySignedReport)
		msg := builder.buildWriteError(info, 0, "failed to source the signed report", cause)
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}

	// Decode the signed report
	inputs := types.SignedReport{}
	if err = signedReport.UnwrapTo(&inputs); err != nil {
		msg := builder.buildWriteError(info, 0, "failed to parse signed report", err.Error())
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}

	// Source the report ID from the input
	info.reportInfo.reportID = binary.BigEndian.Uint16(inputs.ID)

	// TODO: Not sure if I should be returning the error here or just logging it as I am now.
	err = c.beholder.ProtoEmitter.EmitWithLog(ctx, builder.buildWriteInitiated(info))
	if err != nil {
		c.lggr.Errorw("failed to emit write initiated", "err", err)
	}

	// Check whether the report is valid (e.g., not empty)
	if len(inputs.Report) == 0 {
		// We received any empty report -- this means we should skip transmission.
		err = c.beholder.ProtoEmitter.EmitWithLog(ctx, builder.buildWriteSkipped(info, "empty report"))
		if err != nil {
			c.lggr.Errorw("failed to emit write skipped", "err", err)
		}
		return success(), nil
	}

	// Update the info with the report info
	info.reportInfo = &reportInfo{
		reportID:      info.reportInfo.reportID,
		reportContext: inputs.Context,
		report:        inputs.Report,
		signersNum:    uint32(len(inputs.Signatures)),
	}

	// Decode the report
	reportDecoded, err := platform.Decode(inputs.Report)
	if err != nil {
		msg := builder.buildWriteError(info, 0, "failed to decode the report", err.Error())
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}

	// Validate encoded report is prefixed with workflowID and executionID that match the request meta
	if reportDecoded.ExecutionID != request.Metadata.WorkflowExecutionID {
		msg := builder.buildWriteError(info, 0, "decoded report execution ID does not match the request", "")
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	} else if reportDecoded.WorkflowID != request.Metadata.WorkflowID {
		msg := builder.buildWriteError(info, 0, "decoded report workflow ID does not match the request", "")
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}

	// Fetch the latest head from the chain (timestamp), retry with a default backoff strategy
	ctx = retry.CtxWithID(ctx, info.request.Metadata.WorkflowExecutionID)
	head, err := retry.With(ctx, c.lggr, c.cs.LatestHead)
	if err != nil {
		msg := builder.buildWriteError(info, 0, "failed to fetch the latest head", err.Error())
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}

	c.lggr.Debugw("non-empty valid report",
		"reportID", info.reportInfo.reportID,
		"report", "0x"+hex.EncodeToString(inputs.Report),
		"reportLen", len(inputs.Report),
		"reportDecoded", reportDecoded,
		"reportContext", "0x"+hex.EncodeToString(inputs.Context),
		"reportContextLen", len(inputs.Context),
		"signaturesLen", len(inputs.Signatures),
		"executionID", request.Metadata.WorkflowExecutionID,
	)

	state, err := c.targetStrategy.QueryTransmissionState(ctx, info.reportInfo.reportID, request)

	if err != nil {
		msg := builder.buildWriteError(info, 0, "failed to fetch [TransmissionState]", err.Error())
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}

	switch state.Status {
	case TransmissionStateNotAttempted:
		c.lggr.Debugw("Transmission not attempted yet, retrying", "reportID", info.reportInfo.reportID)
	case TransmissionStateFailed:
		c.lggr.Debugw("Tranmissions previously failed, retrying", "reportID", info.reportInfo.reportID)
	case TransmissionStateFatal:
		msg := builder.buildWriteError(info, 0, "Transmission attempt fatal", state.Err.Error())
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	case TransmissionStateSucceeded:
		// Source the transmitter address from the on-chain state
		info.reportTransmissionState = state

		err = c.beholder.ProtoEmitter.EmitWithLog(ctx, builder.buildWriteConfirmed(info, head))
		if err != nil {
			c.lggr.Errorw("failed to emit write confirmed", "err", err)
		}
		return success(), nil
	}

	c.lggr.Infow("on-chain report check done - attempting to push to txmgr",
		"reportID", info.reportInfo.reportID,
		"reportLen", len(inputs.Report),
		"reportContextLen", len(inputs.Context),
		"signaturesLen", len(inputs.Signatures),
		"executionID", request.Metadata.WorkflowExecutionID,
	)

	txID, err := c.targetStrategy.TransmitReport(ctx, inputs.Report, inputs.Context, inputs.Signatures, request)
	c.lggr.Debugw("Transaction submitted", "request", request, "transaction-id", txID)
	if err != nil {
		msg := builder.buildWriteError(info, 0, "failed to transmit the report", err.Error())
		return capabilities.CapabilityResponse{}, c.asEmittedError(ctx, msg)
	}
	err = c.beholder.ProtoEmitter.EmitWithLog(ctx, builder.buildWriteSent(info, head, txID))
	if err != nil {
		c.lggr.Errorw("failed to emit write sent", "err", err)
	}

	// TODO: implement a background WriteTxConfirmer to periodically source new events/transactions,
	// relevant to this forwarder), and emit write-tx-accepted/confirmed events.

	go c.acceptAndConfirmWrite(ctx, *info, txID)
	return success(), nil
}

func (c *writeTarget) RegisterToWorkflow(ctx context.Context, request capabilities.RegisterToWorkflowRequest) error {
	// TODO: notify the background WriteTxConfirmer (workflow registered)
	return nil
}

func (c *writeTarget) UnregisterFromWorkflow(ctx context.Context, request capabilities.UnregisterFromWorkflowRequest) error {
	// TODO: notify the background WriteTxConfirmer (workflow unregistered)
	return nil
}

// acceptAndConfirmWrite waits (until timeout) for the report to be accepted and (optionally) confirmed on-chain
// Emits Beholder messages:
//   - 'platform.write-target.WriteError'     if not accepted
//   - 'platform.write-target.WriteAccepted'  if accepted (with or without an error)
//   - 'platform.write-target.WriteError'     if accepted (with an error)
//   - 'platform.write-target.WriteConfirmed' if confirmed (until timeout)
func (c *writeTarget) acceptAndConfirmWrite(ctx context.Context, info requestInfo, txID string) {
	attrs := c.traceAttributes(info.request.Metadata.WorkflowExecutionID)
	_, span := c.beholder.Tracer.Start(ctx, "Execute.acceptAndConfirmWrite", trace.WithAttributes(attrs...))
	defer span.End()

	lggr := logger.Named(c.lggr, "write-confirmer")

	// Timeout for the confirmation process
	timeout := c.config.ConfirmerTimeout.Duration()
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), timeout)
	defer cancel()

	// Retry interval for the confirmation process
	interval := c.config.ConfirmerPollPeriod.Duration()
	ticker := services.NewTicker(interval)
	defer ticker.Stop()

	// Helper to build monitoring (Beholder) messages
	// Notice: error skipped as implementation always returns nil
	capInfo, _ := c.Info(ctx)
	builder := NewMessageBuilder(c.chainInfo, capInfo)

	// Fn helpers
	checkAcceptedStatus := func(ctx context.Context) (commontypes.TransactionStatus, bool, error) {
		// Check TXM for status
		status, err := c.cw.GetTransactionStatus(ctx, txID)
		if err != nil {
			return commontypes.Unknown, false, fmt.Errorf("failed to get tx status: %w", err)
		}

		lggr.Debugw("txm - tx status", "txID", txID, "status", status)

		// Check if the transaction was accepted (included in a chain block, not required to be finalized)
		// Notice: 'Unconfirmed' is used by TXM to indicate the transaction is not yet included in a block,
		// while 'Included' (N/A yet) could be used to indicate the transaction is included in a block but not yet finalized.
		if /* status == commontypes.Included || */ status == commontypes.Finalized {
			return status, true, nil
		}

		// false if [Unknown, Pending, Failed, Fatal]
		return status, false, nil
	}

	for {
		select {
		case <-ctx.Done():
			// We (eventually) failed to confirm the report was transmitted
			err := c.beholder.ProtoEmitter.EmitWithLog(ctx, builder.buildWriteError(&info, 0, "write confirmation - failed", "timed out"))
			if err != nil {
				lggr.Errorw("failed to emit write error", "err", err)
			}
			return
		case <-ticker.C:
			// Fetch the latest head from the chain (timestamp)
			head, err := c.cs.LatestHead(ctx)
			if err != nil {
				lggr.Errorw("failed to fetch the latest head", "txID", txID, "err", err)
				continue
			}

			// Check acceptance status
			status, accepted, statusErr := checkAcceptedStatus(ctx)
			if statusErr != nil {
				lggr.Errorw("failed to check accepted status", "txID", txID, "err", statusErr)
				continue
			}

			if !accepted {
				lggr.Infow("not accepted yet", "txID", txID, "status", status)
				continue
			}

			lggr.Infow("accepted", "txID", txID, "status", status)
			// Notice: report write confirmation is only possible after a tx is accepted without an error
			// TODO: [Beholder] Emit 'platform.write-target.WriteAccepted' (useful to source tx hash, block number, and tx status/error)

			// TODO: check if accepted with an error (e.g., on-chain revert)
			// Notice: this functionality is not available in the current CW/TXM API

			// Check confirmation status (transmission state)
			state, err := c.targetStrategy.QueryTransmissionState(ctx, info.reportInfo.reportID, info.request)
			if err != nil {
				lggr.Errorw("failed to check confirmed status", "txID", txID, "err", err)
				continue
			}

			if state == nil {
				lggr.Infow("not confirmed yet - transmission state NOT visible", "txID", txID)
				continue
			}

			// We (eventually) confirmed the report was transmitted
			// Emit the confirmation message and return
			lggr.Infow("confirmed - transmission state visible", "txID", txID)

			// Source the transmitter address from the on-chain state
			info.reportTransmissionState = state

			err = c.beholder.ProtoEmitter.EmitWithLog(ctx, builder.buildWriteConfirmed(&info, head))
			if err != nil {
				lggr.Errorw("failed to emit write confirmed", "err", err)
			}
			return
		}
	}
}

// traceAttributes returns the attributes to be used for tracing
func (c *writeTarget) traceAttributes(workflowExecutionID string) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("capability_id", c.ID),
		attribute.String("capability_type", string(c.CapabilityType)),
		attribute.String("workflow_execution_id", workflowExecutionID),
	}
}

// asEmittedError returns the WriteError message as an (Go) error, after emitting it first
func (c *writeTarget) asEmittedError(ctx context.Context, e *wt.WriteError, attrKVs ...any) error {
	// Notice: we always want to log the error
	err := c.beholder.ProtoEmitter.EmitWithLog(ctx, e, attrKVs...)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to emit error: %+w", err), e)
	}
	return e
}
