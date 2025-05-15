// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "../../../shared/access/Ownable2StepMsgSender.sol";

contract WorkflowRegistry is Ownable2StepMsgSender, ITypeAndVersion {
  string public constant override typeAndVersion = "WorkflowRegistry 2.0.0-dev";

  enum WorkflowStatus {
    ACTIVE,
    PAUSED
  }

  struct WorkflowMetadata {
    bytes32 workflowID; //     Unique identifier from hash of owner address, WASM binary content, config content and secrets URL.
    bytes32 donLabel; //       Label for the DON that is used when distributing the workflow across DONs.
    address owner; // ─────────╮ Workflow owner.
    uint64 created_at; //      │ block.timestamp when the workflow was first registered.
    WorkflowStatus status; // ─╯ Current status of the workflow (active, paused).
    string workflowName; //    Human readable string capped at 64 characters length.
    string binaryURL; //       URL to the WASM binary.
    string configURL; //       URL to the config.
    string secretsURL; //      URL to the encrypted secrets. Workflow DON applies a default refresh period (e.g. daily).
  }
}
