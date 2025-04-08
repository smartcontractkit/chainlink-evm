package forwarder

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/platform"

	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"
)

// DecodeAsReportProcessed decodes a 'platform.write-target.WriteConfirmed' message
// as a 'keytone.forwarder.ReportProcessed' message
func DecodeAsReportProcessed(m *wt_msg.WriteConfirmed) (*ReportProcessed, error) {
	// Decode the confirmed report (WT -> platform forwarder contract event akka. Keystone)
	r, err := platform.Decode(m.Report)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	return &ReportProcessed{
		// Event data
		Receiver:            m.Receiver,
		WorkflowExecutionId: r.ExecutionID,
		ReportId:            m.ReportId,
		Success:             m.Success,

		// Head data - when was the event produced on-chain
		BlockHash:      m.BlockHash,
		BlockHeight:    m.BlockHeight,
		BlockTimestamp: m.BlockTimestamp,

		// Transaction data - info about the tx that mained the event (optional)
		// Notice: we skip SOME head/tx data here (unknown), as we map from 'platform.write-target.WriteConfirmed'
		// and not from tx/event data (e.g., 'platform.write-target.WriteTxConfirmed')
		TxSender:   m.Transmitter,
		TxReceiver: m.Forwarder,

		// Execution Context - Source
		MetaSourceId: m.MetaSourceId,

		// Execution Context - Chain
		MetaChainFamilyName: m.MetaChainFamilyName,
		MetaChainId:         m.MetaChainId,
		MetaNetworkName:     m.MetaNetworkName,
		MetaNetworkNameFull: m.MetaNetworkNameFull,

		// Execution Context - Workflow (capabilities.RequestMetadata)
		MetaWorkflowId:               m.MetaWorkflowId,
		MetaWorkflowOwner:            m.MetaWorkflowOwner,
		MetaWorkflowExecutionId:      m.MetaWorkflowExecutionId,
		MetaWorkflowName:             m.MetaWorkflowName,
		MetaWorkflowDonId:            m.MetaWorkflowDonId,
		MetaWorkflowDonConfigVersion: m.MetaWorkflowDonConfigVersion,
		MetaReferenceId:              m.MetaReferenceId,

		// Execution Context - Capability
		MetaCapabilityType:           m.MetaCapabilityType,
		MetaCapabilityId:             m.MetaCapabilityId,
		MetaCapabilityTimestampStart: m.MetaCapabilityTimestampStart,
		MetaCapabilityTimestampEmit:  m.MetaCapabilityTimestampEmit,
	}, nil
}
