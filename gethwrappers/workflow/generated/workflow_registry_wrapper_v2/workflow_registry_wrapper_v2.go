// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package workflow_registry_wrapper_v2

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type WorkflowRegistryEventRecord struct {
	EventType uint8
	Timestamp uint32
	Payload   []byte
}

type WorkflowRegistryWorkflowMetadataView struct {
	WorkflowId   [32]byte
	Owner        common.Address
	CreatedAt    uint64
	Status       uint8
	WorkflowName string
	BinaryUrl    string
	ConfigUrl    string
	Tag          string
	Attributes   []byte
	DonFamily    string
}

var WorkflowRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlistRequest\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDONRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxAttrLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxNameLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxTagLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxUrlLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"maxWorkflows\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflow\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowById\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwnerAndName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataConfig\",\"inputs\":[{\"name\":\"nameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"urlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"attrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsOnDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalEvents\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWorkflowDONFamily\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataConfigUpdated\",\"inputs\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDonFamilyUpdated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowOwnerConfigUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BinaryURLRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUnlinkWithActiveWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUpdateDONFamilyForPausedWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DonLimitNotSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RequestExpired\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWorkflowIDNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561486f90816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe610140604052600436101561001357600080fd5b60003560e01c806302daf47c146126c65780630987294c1461268257806317e0edfc14612576578063181f5a771461255a5780631c08b00a1461248b5780631c71682c146122b85780631f77b2b61461228f578063274e00e0146122315780632afc4130146121d95780632c0403ef146121235780632c50a955146120c657806335141251146120675780633c17181b146120145780633c54b50b14611fc157806347d1ed8314611fa7578063530979d614611f18578063556dbd0d14611ea15780635a1ac5ad14611e68578063695e134014611e335780636ee80b4414611d345780636f19d3b814611d0e57806370ae264014611cd857806379ba509714611c325780638b42a96d14611adf5780638c42ffc514611a2b5780638da5cb5b14611a04578063952bb984146119b957806395be176e1461197b57806396fbd7f01461189e578063a0b8a4fe14611880578063a4089016146117dd578063a7d0185814611736578063afbb24011461171f578063b377bfc514610863578063b668435f146107fb578063ba870686146107dd578063bae5c29a14610793578063bdf6b4ff14610730578063be67433314610709578063cabb9e7a146106ca578063d8b80738146105e8578063d8e4a7241461048e578063dc101969146103f0578063e627fc79146103c7578063e690f33214610343578063ea32308b14610309578063eb3da941146102df5763f2fde38b1461022d57600080fd5b346102da5760203660031901126102da576001600160a01b0361024e6128e4565b610256613df5565b163381146102b0578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102da5760003660031901126102da57602061ffff6102fd613de3565b60181c16604051908152f35b346102da5760203660031901126102da576001600160a01b0361032a6128e4565b16600052600b6020526020604060002054604051908152f35b346102da5760203660031901126102da5761036b336000526005602052604060002054151590565b156103b257600435600052600a60205260406000205461038b8133613f59565b600160ff8183015460e01c166103a081612758565b036103a757005b6103b091613e58565b005b63c2dda3f960e01b6000523360045260246000fd5b346102da5760003660031901126102da57602060ff6103e4613de3565b60081c16604051908152f35b346102da5760603660031901126102da5760443560243567ffffffffffffffff82116102da5761043461042960019336906004016129a6565b9083600435336135fe565b3360005260066020528060406000205561044d33614573565b5080600052600760205260406000208260ff19825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346102da5760403660031901126102da5760043567ffffffffffffffff81116102da576104bf9036906004016129de565b602435918215158093036102da576104d5613df5565b60ff831660005b83811061056657505060405191806040840160408552526060830191906000905b808210610533577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b9091928335906001600160a01b03821682036102da57602080916001600160a01b036001941681520194019201906104fd565b6001600160a01b0361058161057c838787612cc8565b613c18565b16156105be57806001600160a01b036105a061057c6001948888612cc8565b16600052600360205260406000208360ff19825416179055016104dc565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102da5760203660031901126102da5760043567ffffffffffffffff81116102da576106199036906004016129de565b80156106a057610636336000526005602052604060002054151590565b156103b25760005b81811061064757005b806106556001928486612cc8565b35600052600a60205260406000205461066e8133613f59565b8360ff8183015460e01c1661068281612758565b03610690575b50500161063e565b61069991613e58565b8480610688565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346102da5760203660031901126102da576001600160a01b036106eb6128e4565b166000526003602052602060ff604060002054166040519015158152f35b346102da57602061072261071c36612a97565b91613b8c565b63ffffffff60405191168152f35b346102da5760203660031901126102da5760043567ffffffffffffffff81116102da5761076461076b9136906004016129a6565b3691612bc4565b602081519101206000526014602052602063ffffffff60016040600020015416604051908152f35b346102da5760403660031901126102da576001600160a01b036107b46128e4565b166000526012602052604060002060243560005260205260206040600020546040519042108152f35b346102da5760003660031901126102da576020601654604051908152f35b346102da5760803660031901126102da576108146128e4565b60243567ffffffffffffffff81116102da576108349036906004016129a6565b6044359063ffffffff821682036102da576064359283151584036102da576103b09461085e613df5565b6139a3565b346102da576101203660031901126102da5760043567ffffffffffffffff81116102da576108959036906004016129a6565b610100526101205260243567ffffffffffffffff81116102da576108bd9036906004016129a6565b90600260643510156102da5760843567ffffffffffffffff81116102da576108e99036906004016129a6565b9060e0529060a43567ffffffffffffffff81116102da5761090e9036906004016129a6565b9160c43567ffffffffffffffff81116102da5761092f9036906004016129a6565b9560e43567ffffffffffffffff81116102da576109509036906004016129a6565b929093610104359586151587036102da57610978336000526005602052604060002054151590565b156103b257604435156116f557604435600052600a6020526040600020546116c55760ff6109a4613de3565b60101c16881561169b578061165e575b5061ffff6109c0613de3565b60181c1680151580611655575b611624575083156115fa5760ff6109e2613de3565b60081c16801515806115f1575b6115c0575061010051156115965760ff610a07613de3565b168015158061158a575b61155657506040516020810190610a4481610a3688876101005161012051338961382c565b03601f198101835282612949565b5190209788600052600860205260406000206001600160a01b0360018201541680156000146111d4575050610a89610a83366101005161012051612bc4565b33613dc4565b9a610a97368c60e051612bc4565b602081519101209815611148575b610ab0606435612758565b60643515611120575b604051978861012081011067ffffffffffffffff6101208b011117610f4957610b66899695610b578897610b75956101208a016040526044358a5260a0610b4b60208c019a338c5260408d0160c05267ffffffffffffffff421660c0515260608d019a610b27606435612758565b6064358c526080610b3f366101005161012051612bc4565b9e019d8e523691612bc4565b9b019a8b523691612bc4565b9860c08c01998a523691612bc4565b9760e08a019889523691612bc4565b610100880160a05260a051528860005260086020526040600020965187556001600160a01b036001880192511673ffffffffffffffffffffffffffffffffffffffff1983541617825560c051517fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b8085549360a01b1616911617825551610c0781612758565b610c1081612758565b60ff60e01b197cff0000000000000000000000000000000000000000000000000000000083549260e01b16911617905560028501905180519067ffffffffffffffff8211610f4957610c6c82610c668554612c0b565b85612f16565b602090601f83116001146110b957610c9d9291600091836110ae575b50508160011b916000199060031b1c19161790565b90555b518051600385019167ffffffffffffffff8211610f4957610cc582610c668554612c0b565b602090601f831160011461104757610cf592916000918361103c5750508160011b916000199060031b1c19161790565b90555b518051600484019167ffffffffffffffff8211610f4957610d1d82610c668554612c0b565b602090601f8311600114610fd557610d4d929160009183610f5f5750508160011b916000199060031b1c19161790565b90555b518051600583019167ffffffffffffffff8211610f4957610d7582610c668554612c0b565b602090601f8311600114610f6a579180610daa926006969594600092610f5f5750508160011b916000199060031b1c19161790565b90555b019360a0515194855167ffffffffffffffff8111610f4957610dd981610dd38454612c0b565b84612f16565b6020601f8211600114610edf579080610e0d92610e639899600092610ed45750508160011b916000199060031b1c19161790565b90555b6000526009602052610e268260406000206145ea565b50604435600052600a60205281604060002055600052600d602052610e4f8160406000206145ea565b5033600052600e60205260406000206145ea565b50610e7c6040519160608352606083019060e0516130cb565b610e87606435612758565b606435602083015281810360408301527f74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a09339280610ecf6044359461010051610120516130cb565b0390a3005b015190508980610c88565b601f1982169783600052816000209860005b818110610f31575091610e63989991846001959410610f18575b505050811b019055610e10565b015160001960f88460031b161c19169055888080610f0b565b838301518b556001909a019960209384019301610ef1565b634e487b7160e01b600052604160045260246000fd5b015190508a80610c88565b90601f1983169184600052816000209260005b818110610fbd575091600193918560069897969410610fa4575b505050811b019055610dad565b015160001960f88460031b161c19169055898080610f97565b92936020600181928786015181550195019301610f7d565b90601f1983169184600052816000209260005b818110611024575090846001959493921061100b575b505050811b019055610d50565b015160001960f88460031b161c19169055898080610ffe565b92936020600181928786015181550195019301610fe8565b015190508b80610c88565b90601f1983169184600052816000209260005b818110611096575090846001959493921061107d575b505050811b019055610cf8565b015160001960f88460031b161c191690558a8080611070565b9293602060018192878601518155019501930161105a565b015190508c80610c88565b90601f1983169184600052816000209260005b81811061110857509084600195949392106110ef575b505050811b019055610ca0565b015160001960f88460031b161c191690558b80806110e2565b929360206001819287860151815501950193016110cc565b611137611130368d60e051612bc4565b8a33613fe9565b6111438c8a338d614372565b610ab9565b9a999897969594939291908a600052600f60205260406000206080525b6080515480156111c357806000198101116111ad5761118d6111a89160001901608051614416565b90549060031b1c806000526008602052604060002090613e58565b611165565b634e487b7160e01b600052601160045260246000fd5b50909192939495969798999a610aa5565b9294969850969a94509791503303611541578454600052600a60205260006040812055604435600052600a6020526040600020558354956044358555600385019160405161122d816112268187612c45565b0382612949565b6020815191012061123f368484612bc4565b602081519101200361147a575b5050506004830191604051611265816112268187612c45565b60208151910120611277368484612bc4565b60208151910120036113aa575b5050506006019067ffffffffffffffff8111610f49576112a881610dd38454612c0b565b6000601f82116001146113465781906112da9394959660009261133b5750508160011b916000199060031b1c19161790565b90555b6112f56040519260408452604084019060e0516130cb565b9082820360208401527f03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e3223393806113366044359561010051610120516130cb565b0390a4005b013590508680610c88565b601f198216958382526020822091805b88811061139257508360019596979810611378575b505050811b0190556112dd565b0135600019600384901b60f8161c1916905585808061136b565b90926020600181928686013581550194019101611356565b67ffffffffffffffff8211610f49576113c782610c668554612c0b565b600090601f83116001146114125791806113fc9260069695946000926114075750508160011b916000199060031b1c19161790565b90555b908680611284565b013590508a80610c88565b8382526020822091601f198416815b818110611462575091600193918560069897969410611448575b505050811b0190556113ff565b0135600019600384901b60f8161c1916905589808061143b565b91936020600181928787013581550195019201611421565b67ffffffffffffffff8211610f495761149782610c668554612c0b565b600090601f83116001146114dd576114c79291600091836114d25750508160011b916000199060031b1c19161790565b90555b87808061124c565b013590508b80610c88565b8382526020822091601f198416815b818110611529575090846001959493921061150f575b505050811b0190556114ca565b0135600019600384901b60f8161c191690558a8080611502565b919360206001819287870135815501950192016114ec565b6331ee6dc760e01b6000523360045260246000fd5b7f36a7c503000000000000000000000000000000000000000000000000000000006000526101005160045260245260446000fd5b50806101005111610a11565b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b847f436f97540000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b508085116109ef565b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b857f354f25140000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b508086116109cd565b80891161168457808b11156109b4578a6219aad560e31b60005260045260245260446000fd5b886219aad560e31b60005260045260245260446000fd5b7f9cd963cf0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260443560045260246000fd5b7f315de7450000000000000000000000000000000000000000000000000000000060005260046000fd5b346102da576103b061173036612a63565b91613861565b346102da5760203660031901126102da5760043567ffffffffffffffff81116102da576117679036906004016129de565b80156106a05760005b81811061177957005b806117876001928486612cc8565b35611790613df5565b600052600a602052604060002054806000526008602052604060002060ff8482015460e01c166117bf81612758565b156117cd575b505001611770565b6117d691613e58565b84806117c5565b346102da5760603660031901126102da576117f66128e4565b60243567ffffffffffffffff81116102da576118169036906004016129a6565b916044359267ffffffffffffffff84116102da5761186061186893610a3661184561187c9736906004016129a6565b9061184e612aef565b5060405195869460208601988961382c565b519020613c6f565b6040519182916020835260208301906127b9565b0390f35b346102da5760003660031901126102da576020600454604051908152f35b346102da5760803660031901126102da5760043560ff81168091036102da576024359060ff8216908183036102da5760443560ff81168082036102da576064359361ffff8516928386036102da577fa254b4f533572ce0e775da29b44e0114b25637b1072bbeaf3fe051fb55aafedc9661ff0062ff000064ffff000000608099611926613df5565b60181b169360101b169160081b1686171717640400c8204081148015611973575b1561196b575060006002555b604051938452602084015260408301526060820152a1005b600255611953565b508015611947565b346102da5760403660031901126102da576119946128e4565b60243567ffffffffffffffff811681036102da576103b0916119b4613df5565b613722565b346102da5760803660031901126102da576119d26128e4565b6064359067ffffffffffffffff82116102da576119f66103b09236906004016129a6565b9160443590602435906135fe565b346102da5760003660031901126102da5760206001600160a01b0360015416604051908152f35b346102da5760603660031901126102da57611a446128e4565b6001600160a01b0360243591169081600052600e602052611a6d60443582604060002054613c2c565b90611a7782612b3d565b9260005b838110611a90576040518061187c8782612884565b60019082600052600e602052611ac3611ab76040600020611ab18488612b8d565b90614416565b90549060031b1c613c6f565b611acd8288612b9a565b52611ad88187612b9a565b5001611a7b565b346102da576001600160a01b03611af536612a97565b929091611b00613df5565b169081600052601360205260406000209267ffffffffffffffff8111610f4957611b3481611b2e8654612c0b565b86612f16565b600093601f8211600114611bb057611b8682807f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f139697600091611ba5575b508160011b916000199060031b1c19161790565b90555b611ba06040519283926020845260208401916130cb565b0390a2005b905085013588611b72565b80855260208520601f19831695805b878110611c1a5750837f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f13969710611c00575b5050600182811b019055611b89565b840135600019600385901b60f8161c191690558580611bf1565b90916020600181928589013581550193019101611bbf565b346102da5760003660031901126102da576000546001600160a01b0381163303611cae5773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102da5760203660031901126102da57611cf1612aef565b50600435600052600a60205261187c611868604060002054613c6f565b346102da5760003660031901126102da57602060ff611d2b613de3565b16604051908152f35b346102da57611d51611d4536612a0f565b9394929184928661339c565b611d5a816129d4565b80611db4575b60006001600160a01b038316611d75816144d4565b908083526006602052826040812055611d8d816146db565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b6001600160a01b038216600052600b60205260406000205b80548015611e235760001981019081116111ad57611dea9082614416565b90549060031b1c611dfb8185613f59565b611e04846129d4565b60028403611e1a57611e1591613e58565b611dcc565b611e15916141db565b5050506001600160a01b03611d60565b346102da5760203660031901126102da57600435600052600a6020526103b0604060002054611e628133613f59565b906141db565b346102da5760003660031901126102da57604060155467ffffffffffffffff8251916001600160a01b038116835260a01c166020820152f35b346102da5760203660031901126102da5760043567ffffffffffffffff81116102da57611ed5611edd9136906004016129a6565b610764613df5565b60208151910120600052600c60205260406000205b805480156103b05760001981019081116111ad5761118d611f139183614416565b611ef2565b346102da57611f2636612a63565b9091611f3f336000526005602052604060002054151590565b156103b257600052600a602052604060002054611f5c8133613f59565b9060ff600183015460e01c16611f7181612758565b611f7757005b611fa2611f9b6103b095611f8c368783612bc4565b60208151910120953691612bc4565b8433613fe9565b6140e0565b346102da576103b0611fb836612a0f565b9392909261339c565b346102da5760203660031901126102da5760043567ffffffffffffffff81116102da57610764611ff59136906004016129a6565b60208151910120600052600c6020526020604060002054604051908152f35b346102da5760203660031901126102da5761202d613df5565b600435600052600a602052604060002054806000526008602052604060002060ff600182015460e01c1661206081612758565b156103a757005b346102da5760603660031901126102da5760043567ffffffffffffffff81116102da576120989036906004016129a6565b60243563ffffffff811681036102da576044359182151583036102da576103b0936120c1613df5565b613111565b346102da5760403660031901126102da5760043567ffffffffffffffff81116102da576120f79036906004016129de565b6024359167ffffffffffffffff83116102da5761211b6103b09336906004016129a6565b929091612cd8565b346102da5760403660031901126102da57600435602435428111156121a857612159336000526005602052604060002054151590565b156103b257336000526012602052604060002082600052602052806040600020556040519081527fd2e6cd8f534fad4a57838c5102996a78ebe4bd24ca8443437528c6c9295e4b9660203392a3005b907f2ec4e4750000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b346102da5760203660031901126102da576001600160a01b036121fa6128e4565b16600052601360205261187c61122661221d604060002060405192838092612c45565b604051918291602083526020830190612778565b346102da5760203660031901126102da576001600160a01b036122526128e4565b61225a613df5565b16600052600b60205260406000205b805480156103b05760001981019081116111ad5761118d61228a9183614416565b612269565b346102da5760003660031901126102da57602060ff6122ac613de3565b60101c16604051908152f35b346102da576122d56122c9366128fa565b81601693929354613c2c565b6122de81612ad7565b916122ec6040519384612949565b818352601f196122fb83612ad7565b0160005b81811061245f575050601654919060005b8281106123a957846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b82821061235057505050500390f35b919360019193955060206123998192603f198a8203018652606060408a518051612379816129d4565b845263ffffffff8682015116868501520151918160408201520190612778565b9601920192018594939192612341565b6123b38183612b8d565b60008582101561244b576016905260405160019291839061242a90821b7fd833147d7dc355ba459fc788f669e58cfaf9dc25ddcd0702e87d69c7b5124289016123fb84612910565b63ffffffff815460ff811661240f816129d4565b865260081c1660208501526112266040518094819301612c45565b60408201526124398288612b9a565b526124448187612b9a565b5001612310565b80634e487b7160e01b602492526032600452fd5b60209060405161246e81612910565b6000815260008382015260606040820152828288010152016122ff565b346102da5760803660031901126102da576124a46128e4565b60243567ffffffffffffffff81116102da576124d8916124cb6124de9236906004016129a6565b9390604435943691612bc4565b90613dc4565b908160005260096020526124fa60643582604060002054613c2c565b9061250482612b3d565b9260005b83811061251d576040518061187c8782612884565b60019082600052600960205261253e611ab76040600020611ab18488612b8d565b6125488288612b9a565b526125538187612b9a565b5001612508565b346102da5760003660031901126102da5761187c61221d61296b565b346102da57612593612587366128fa565b81600493929354613c2c565b61259c81612ad7565b916125aa6040519384612949565b8183526125b682612ad7565b602084019290601f19013684376004549160005b82811061261f5784866040519182916020830190602084525180915260408301919060005b8181106125fd575050500390f35b82516001600160a01b03168452859450602093840193909201916001016125ef565b6126298183612b8d565b60008582101561244b57600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b015460008190526006602052600191906001600160a01b031661267b8289612b9a565b52016125ca565b346102da5760203660031901126102da5760206126bc6001600160a01b036126a86128e4565b166000526005602052604060002054151590565b6040519015158152f35b346102da5760603660031901126102da5760043560243581600052600d6020526126f860443582604060002054613c2c565b9061270282612b3d565b9260005b83811061271b576040518061187c8782612884565b60019082600052600d60205261273c611ab76040600020611ab18488612b8d565b6127468288612b9a565b526127518187612b9a565b5001612706565b6002111561276257565b634e487b7160e01b600052602160045260246000fd5b919082519283825260005b8481106127a4575050826000602080949584010152601f8019910116010190565b80602080928401015182828601015201612783565b61288191815181526001600160a01b03602083015116602082015267ffffffffffffffff604083015116604082015260608201516127f681612758565b606082015261012061286f61285b61284961283761282560808801516101406080890152610140880190612778565b60a088015187820360a0890152612778565b60c087015186820360c0880152612778565b60e086015185820360e0870152612778565b610100850151848203610100860152612778565b92015190610120818403910152612778565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106128b757505050505090565b90919293946020806128d5600193603f1986820301875289516127b9565b970193019301919392906128a8565b600435906001600160a01b03821682036102da57565b60409060031901126102da576004359060243590565b6060810190811067ffffffffffffffff821117610f4957604052565b610140810190811067ffffffffffffffff821117610f4957604052565b90601f8019910116810190811067ffffffffffffffff821117610f4957604052565b6040519061297a604083612949565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b9181601f840112156102da5782359167ffffffffffffffff83116102da57602083818601950101116102da57565b6003111561276257565b9181601f840112156102da5782359167ffffffffffffffff83116102da576020808501948460051b0101116102da57565b60806003198201126102da576004356001600160a01b03811681036102da5791602435916044359067ffffffffffffffff82116102da57612a52916004016129a6565b909160643560038110156102da5790565b9060406003198301126102da57600435916024359067ffffffffffffffff82116102da57612a93916004016129a6565b9091565b9060406003198301126102da576004356001600160a01b03811681036102da57916024359067ffffffffffffffff82116102da57612a93916004016129a6565b67ffffffffffffffff8111610f495760051b60200190565b60405190612afc8261292c565b6060610120836000815260006020820152600060408201526000838201528260808201528260a08201528260c08201528260e0820152826101008201520152565b90612b4782612ad7565b612b546040519182612949565b8281528092612b65601f1991612ad7565b019060005b828110612b7657505050565b602090612b81612aef565b82828501015201612b6a565b919082018092116111ad57565b8051821015612bae5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b92919267ffffffffffffffff8211610f495760405191612bee601f8201601f191660200184612949565b8294818452818301116102da578281602093846000960137010152565b90600182811c92168015612c3b575b6020831014612c2557565b634e487b7160e01b600052602260045260246000fd5b91607f1691612c1a565b60009291815491612c5583612c0b565b8083529260018116908115612cab5750600114612c7157505050565b60009081526020812093945091925b838310612c91575060209250010190565b600181602092949394548385870101520191019190612c80565b915050602093945060ff929192191683830152151560051b010190565b9190811015612bae5760051b0190565b929181156106a057612cf7336000526005602052604060002054151590565b156103b257612d07368483612bc4565b6020815191012092600091825b848110612e9c575063ffffffff831615612e9457612d33913691612bc4565b90836000526014602052604060002060018101549060ff8260201c1615612e7257600201906001600160a01b0333166000528160205260ff60406000205460201c16600014612e5e57506001600160a01b03331660005260205263ffffffff80612dc28160406000205416935b3360005260106020526040600020886000526020528260406000205416613fcf565b9216911611612e3c575060005b818110612ddc5750505050565b80612dea6001928487612cc8565b35600052600a6020526040600020548060005260086020528460406000208460ff8183015460e01c16612e1c81612758565b14612e2b575b50505001612dcf565b612e34926140e0565b388481612e22565b612e5a9060405191829163038857ff60e01b83523360048401613da4565b0390fd5b63ffffffff9150612dc28280921693612da0565b60405163393f328760e11b81526020600482015280612e5a6024820187612778565b505050505050565b92612ea8848689612cc8565b35600052600a60205260ff6001612ec460406000205433613f59565b015460e01c16612ed381612758565b15612ef65763ffffffff1663ffffffff81146111ad576001809101935b01612d14565b92600190612ef0565b818110612f0a575050565b60008155600101612eff565b9190601f8111612f2557505050565b612f51926000526020600020906020601f840160051c83019310612f53575b601f0160051c0190612eff565b565b9091508190612f44565b9060165468010000000000000000811015610f495760018101601655600060165482101561244b57601690526000929060011b7fd833147d7dc355ba459fc788f669e58cfaf9dc25ddcd0702e87d69c7b51242890160016040918351612fc2816129d4565b612fcb816129d4565b60ff825491168060ff19831617835564ffffffff00602087015160081b169164ffffffffff191617178155019101519283519067ffffffffffffffff82116130b75761301b82610c668554612c0b565b602090601f8311600114613059579061304a9394958361304e5750508160011b916000199060031b1c19161790565b9055565b015190503880610c88565b90601f198316848352818320925b81811061309f57509583600195969710613086575b505050811b019055565b015160001960f88460031b161c1916905538808061307c565b9192602060018192868b015181550194019201613067565b80634e487b7160e01b602492526041600452fd5b908060209392818452848401376000828201840152601f01601f1916010190565b9161310a60209263ffffffff929695966040865260408601916130cb565b9416910152565b9091939293613121368484612bc4565b60208151910120906000866000146133685750905b80600052601460205260406000209560018701908154901515908160ff8260201c161515149081613355575b5061334b57875461317281612c0b565b15613223575b50815464ffffffffff1916602091821b64ff00000000161763ffffffff85169081179092556040805191820193909352808301919091529081529394507f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b7319361321e9190613212906131eb606082612949565b604051906131f882612910565b6000825263ffffffff421660208301526040820152612f5d565b604051938493846130ec565b0390a1565b67ffffffffffffffff8711610f49578661323f61324592612c0b565b8a612f16565b600097601f87116001146132bc579161321e949391613212936132a189807f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b7319c9d6000916132b157508160011b916000199060031b1c19161790565b90555b9193945091889750613178565b90508a013538611b72565b8089526020892098601f198816815b8181106133335750917f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b731999a61321e979694928a613212979510613319575b5050600189811b0190556132a4565b89013560001960038c901b60f8161c19169055388061330a565b888301358c556001909b019a602092830192016132cb565b5050505050509050565b905063ffffffff80861691161438613162565b905090613136565b90926080926001600160a01b0361288197951683526020830152604082015281606082015201916130cb565b929190936133a9816129d4565b15806135de575b6135b457834211613575576001600160a01b0383166133dc816000526005602052604060002054151590565b15613561576133ea816144d4565b9060009061342e6134486133fc61296b565b604051928391602083019560018752604084015246606084015230608084015260e060a0840152610100830190612778565b8a60c08301528660e083015203601f198101835282612949565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613489603c8220613483368787612bc4565b90614537565b909192600483101561354d57826134fc575050506001600160a01b031660009081526003602052604090205460ff16156134c4575050505050565b90612e5a92916040519586957f335d4ce100000000000000000000000000000000000000000000000000000000875260048701613370565b506040517fd36ab6b90000000000000000000000000000000000000000000000000000000081526060600482015291829160ff61353d606485018a8a6130cb565b9216602484015260448301520390fd5b80634e487b7160e01b602492526021600452fd5b63c2dda3f960e01b60005260045260246000fd5b836001600160a01b03847f3d8a511600000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b7f61bc2e180000000000000000000000000000000000000000000000000000000060005260046000fd5b506001600160a01b038316600052600b60205260406000205415156133b0565b929093918442116136e3576001600160a01b03841661362a816000526005602052604060002054151590565b6136b65781600052600760205260ff604060002054166136865760009061342e61344861365561296b565b6040519283916020830195878752604084015246606084015230608084015260e060a0840152610100830190612778565b7f77a338580000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7fd9a5f5ca0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b846001600160a01b03857f502d038700000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b90601554906001600160a01b038216906001600160a01b0367ffffffffffffffff8460a01c1694169382851494858096613819575b612e9457806080957fa7a2a5335a8d1f8f1f7ef8a58332be349ac9fdc25b62512290a91ac4555430a597156137f7575b505067ffffffffffffffff8316928284036137b6575b50604051938452602084015260408301526060820152a1565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b6015549260a01b169116176015553861379d565b73ffffffffffffffffffffffffffffffffffffffff1916176015558038613787565b508167ffffffffffffffff841614613757565b939161288195936001600160a01b03613853931686526060602087015260608601916130cb565b9260408185039101526130cb565b909161387a336000526005602052604060002054151590565b156103b25781600052600a602052604060002054906138998233613f59565b60ff600182015460e01c166138ad81612758565b613979576138e290836000526011602052604060002054928360005260146020526138e9604060002060405194858092612c45565b0384612949565b6138f4368289612bc4565b602081519101208094146139705761394761395b94838761393961396b967f9b5361a5258ef6ac8039fd2d2ac276734695219cfd870711e7922c236e5db16d9a613e58565b611fa2611f9b36878e612bc4565b604051938493604085526040850190612778565b90838203602085015233976130cb565b0390a3565b50505050505050565b7fd74915a80000000000000000000000000000000000000000000000000000000060005260046000fd5b919290936139b2368587612bc4565b6020815191012060005260146020526040600020600181019060ff825460201c1615613b695760409060026000916001600160a01b0388168352016020522091600014613afa575463ffffffff83811691168111613ad057815460ff8160201c1615600014613a765750815464ffffffffff1916176401000000001790556040516001600160a01b0392909216937f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c82324938392613a7192909190846130ec565b0390a2565b919392909163ffffffff168214612e94577f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c82324946001600160a01b0394613a719363ffffffff198254161790556040519485941696846130ec565b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b50809291509392935460ff8160201c1615613b625764ffffffffff191690915560405160208082526001600160a01b03909216937f6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d8499391928392613a719291840191906130cb565b5050505050565b60405163393f328760e11b81526020600482015280612e5a60248201898b6130cb565b9190613b99913691612bc4565b6020815191012060005260146020526001600160a01b036040600020911660005260028101602052604060002060405190604082019082821067ffffffffffffffff831117610f4957602091604052549160ff63ffffffff841693848352831c1615159182910152613c1357506001015463ffffffff1690565b905090565b356001600160a01b03811681036102da5790565b909181831015613c675781613c418285612b8d565b1115613c565750905b81039081116111ad5790565b613c61915082612b8d565b90613c4a565b505050600090565b613c77612aef565b8160005260086020526040600020906001820154926001600160a01b038416918215613d9c5750613cd392600692613d8d9260005260116020526040600020546000526014602052613cda604060002060405196878092612c45565b0386612949565b67ffffffffffffffff82549660ff8160e01c169260405198613cfb8a61292c565b8952602089015260a01c166040870152613d1481612758565b6060860152604051613d2d816112268160028601612c45565b6080860152604051613d46816112268160038601612c45565b60a0860152604051613d5f816112268160048601612c45565b60c0860152604051613d78816112268160058601612c45565b60e08601526112266040518094819301612c45565b61010083015261012082015290565b935050505090565b6040906001600160a01b0361288194931681528160208201520190612778565b90613ddd610a3691604051928391602083019586613da4565b51902090565b600254806128815750640400c8204090565b6001600160a01b03600154163303613e0957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9091613e4a61288193604084526040840190612c45565b916020818403910152612c45565b7ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e48868960018301916001600160a01b038354947c010000000000000000000000000000000000000000000000000000000060ff60e01b198716178555826000526011602052613eee82604060002054971693876002840195613ee8604051613ee281611226818c612c45565b82613dc4565b9261442e565b613f38815460405190886020830152604082015260408152613f11606082612949565b60405190613f1e82612910565b6002825263ffffffff421660208301526040820152612f5d565b54935416936000526014602052604060002061396b60405192839283613e33565b9060005260086020526040600020906001600160a01b03600183015416908115613fa5576001600160a01b0316809103613f91575090565b6331ee6dc760e01b60005260045260246000fd5b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b9063ffffffff8091169116019063ffffffff82116111ad57565b90806000526014602052604060002060018101549060ff8260201c16156140be57600201906001600160a01b0384166000528160205260ff60406000205460201c166000146140b057506001600160a01b03831660005260205263ffffffff60406000205416905b6001600160a01b038316600052601060205260406000209060005260205263ffffffff8061408760018260406000205416613fcf565b9216911611614094575050565b612e5a60405192839263038857ff60e01b845260048401613da4565b63ffffffff91501690614051565b60405163393f328760e11b81526020600482015280612e5a6024820188612778565b7f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a906001600160a01b0360018401936141378286541693876002840195614131604051613ee281611226818c612c45565b92614372565b60ff60e01b198554168555613f38815460405190886020830152604082015260408152614165606082612949565b6040519061417282612910565b6001825263ffffffff421660208301526040820152612f5d565b6141968154612c0b565b90816141a0575050565b81601f600093116001146141b2575055565b818352602083206141ce91601f0160051c810190600101612eff565b8082528160208120915555565b90613ee2907f48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd946142fa6142696143086001600160a01b036001860195865460ff8a848316614237611226600287019d8e60405192838092612c45565b93849183600052601160205260406000205494859283600052601460205261427060406000206040519d8e8092612c45565b038d612949565b60e01c1661427d81612758565b15614360575b50505050600052600d60205261429d8a60406000206147a5565b5082885416600052600e6020526142b88a60406000206147a5565b5060005260096020526142cf8960406000206147a5565b508054600052600a602052600060408120555495541695604051938493604085526040850190612778565b908382036020850152612c45565b0390a360005260086020526006604060002060008155600060018201556143316002820161418c565b61433d6003820161418c565b6143496004820161418c565b6143556005820161418c565b016141968154612c0b565b6143699361442e565b8a828238614283565b9192906001600160a01b03168060005260106020526040600020846000526020526040600020600163ffffffff8254160163ffffffff81116111ad5763ffffffff1663ffffffff1982541617905583600052600c6020526143d78360406000206145ea565b50600052600b6020526143ee8260406000206145ea565b50600052600f6020526144058160406000206145ea565b506000526011602052604060002055565b8054821015612bae5760005260206000200190600090565b9291906001600160a01b031680600052600b6020526144518460406000206147a5565b5081600052600c6020526144698460406000206147a5565b506000526010602052604060002090600052602052604060002060001963ffffffff8254160163ffffffff81116111ad5763ffffffff1663ffffffff19825416179055600052600f6020526144c28160406000206147a5565b50600052601160205260006040812055565b80600052600660205260406000205490811580614521575b6144f4575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b50806000526005602052604060002054156144ec565b81519190604183036145685761456192506020820151906060604084015193015160001a90614646565b9192909190565b505060009160029190565b806000526005602052604060002054156000146145e45760045468010000000000000000811015610f49576145cb6145b48260018594016004556004614416565b819391549060031b91821b91600019901b19161790565b9055600454906000526005602052604060002055600190565b50600090565b600082815260018201602052604090205461463f5780549068010000000000000000821015610f4957826146286145b4846001809601855584614416565b905580549260005201602052604060002055600190565b5050600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116146cf579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa156146c3576000516001600160a01b038116156146b75790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b50505060009160039190565b600081815260056020526040902054801561463f5760001981018181116111ad576004546000198101919082116111ad5781810361476b575b5050506004548015614755576000190161472f816004614416565b8154906000199060031b1b19169055600455600052600560205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61478d61477c6145b4936004614416565b90549060031b1c9283926004614416565b90556000526005602052604060002055388080614714565b90600182019181600052826020526040600020548015156000146148595760001981018181116111ad5782546000198101919082116111ad57818103614822575b505050805480156147555760001901906148008282614416565b8154906000199060031b1b191690555560005260205260006040812055600190565b6148426148326145b49386614416565b90549060031b1c92839286614416565b9055600052836020526040600020553880806147e6565b5050505060009056fea164736f6c634300081a000a",
}

var WorkflowRegistryABI = WorkflowRegistryMetaData.ABI

var WorkflowRegistryBin = WorkflowRegistryMetaData.Bin

func DeployWorkflowRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WorkflowRegistry, error) {
	parsed, err := WorkflowRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WorkflowRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WorkflowRegistry{address: address, abi: *parsed, WorkflowRegistryCaller: WorkflowRegistryCaller{contract: contract}, WorkflowRegistryTransactor: WorkflowRegistryTransactor{contract: contract}, WorkflowRegistryFilterer: WorkflowRegistryFilterer{contract: contract}}, nil
}

type WorkflowRegistry struct {
	address common.Address
	abi     abi.ABI
	WorkflowRegistryCaller
	WorkflowRegistryTransactor
	WorkflowRegistryFilterer
}

type WorkflowRegistryCaller struct {
	contract *bind.BoundContract
}

type WorkflowRegistryTransactor struct {
	contract *bind.BoundContract
}

type WorkflowRegistryFilterer struct {
	contract *bind.BoundContract
}

type WorkflowRegistrySession struct {
	Contract     *WorkflowRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type WorkflowRegistryCallerSession struct {
	Contract *WorkflowRegistryCaller
	CallOpts bind.CallOpts
}

type WorkflowRegistryTransactorSession struct {
	Contract     *WorkflowRegistryTransactor
	TransactOpts bind.TransactOpts
}

type WorkflowRegistryRaw struct {
	Contract *WorkflowRegistry
}

type WorkflowRegistryCallerRaw struct {
	Contract *WorkflowRegistryCaller
}

type WorkflowRegistryTransactorRaw struct {
	Contract *WorkflowRegistryTransactor
}

func NewWorkflowRegistry(address common.Address, backend bind.ContractBackend) (*WorkflowRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(WorkflowRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindWorkflowRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistry{address: address, abi: abi, WorkflowRegistryCaller: WorkflowRegistryCaller{contract: contract}, WorkflowRegistryTransactor: WorkflowRegistryTransactor{contract: contract}, WorkflowRegistryFilterer: WorkflowRegistryFilterer{contract: contract}}, nil
}

func NewWorkflowRegistryCaller(address common.Address, caller bind.ContractCaller) (*WorkflowRegistryCaller, error) {
	contract, err := bindWorkflowRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryCaller{contract: contract}, nil
}

func NewWorkflowRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkflowRegistryTransactor, error) {
	contract, err := bindWorkflowRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryTransactor{contract: contract}, nil
}

func NewWorkflowRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkflowRegistryFilterer, error) {
	contract, err := bindWorkflowRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryFilterer{contract: contract}, nil
}

func bindWorkflowRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkflowRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowRegistry.Contract.WorkflowRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.WorkflowRegistryTransactor.contract.Transfer(opts)
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.WorkflowRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.contract.Transfer(opts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) CanLinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canLinkOwner", owner, validityTimestamp, proof, signature)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanLinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanLinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canUnlinkOwner", owner, validityTimestamp, signature, action)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetDONRegistry(opts *bind.CallOpts) (common.Address, uint64, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getDONRegistry")

	if err != nil {
		return *new(common.Address), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetDONRegistry() (common.Address, uint64, error) {
	return _WorkflowRegistry.Contract.GetDONRegistry(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetDONRegistry() (common.Address, uint64, error) {
	return _WorkflowRegistry.Contract.GetDONRegistry(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getEvents", start, limit)

	if err != nil {
		return *new([]WorkflowRegistryEventRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryEventRecord)).(*[]WorkflowRegistryEventRecord)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetEvents(start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error) {
	return _WorkflowRegistry.Contract.GetEvents(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetEvents(start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error) {
	return _WorkflowRegistry.Contract.GetEvents(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getLinkedOwners", start, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetLinkedOwners(start *big.Int, limit *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetLinkedOwners(start *big.Int, limit *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxAttrLen(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxAttrLen")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxAttrLen() (uint16, error) {
	return _WorkflowRegistry.Contract.GetMaxAttrLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxAttrLen() (uint16, error) {
	return _WorkflowRegistry.Contract.GetMaxAttrLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxNameLen(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxNameLen")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxNameLen() (uint8, error) {
	return _WorkflowRegistry.Contract.GetMaxNameLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxNameLen() (uint8, error) {
	return _WorkflowRegistry.Contract.GetMaxNameLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxTagLen(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxTagLen")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxTagLen() (uint8, error) {
	return _WorkflowRegistry.Contract.GetMaxTagLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxTagLen() (uint8, error) {
	return _WorkflowRegistry.Contract.GetMaxTagLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxUrlLen(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxUrlLen")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxUrlLen() (uint8, error) {
	return _WorkflowRegistry.Contract.GetMaxUrlLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxUrlLen() (uint8, error) {
	return _WorkflowRegistry.Contract.GetMaxUrlLen(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerDON", donFamily)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerDON(donFamily string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerDON(donFamily string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donFamily string) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerUserDON", user, donFamily)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerUserDON(user common.Address, donFamily string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerUserDON(user common.Address, donFamily string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflow(opts *bind.CallOpts, owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflow", owner, workflowName, tag)

	if err != nil {
		return *new(WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryWorkflowMetadataView)).(*WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflow(owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflow(&_WorkflowRegistry.CallOpts, owner, workflowName, tag)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflow(owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflow(&_WorkflowRegistry.CallOpts, owner, workflowName, tag)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowById(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowById", workflowId)

	if err != nil {
		return *new(WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryWorkflowMetadataView)).(*WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowById(workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowById(&_WorkflowRegistry.CallOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowById(workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowById(&_WorkflowRegistry.CallOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowListByDON(opts *bind.CallOpts, donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowListByDON", donFamily, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowListByDON(donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowListByDON(donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowListByOwner", owner, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowListByOwnerAndName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowListByOwnerAndName", owner, workflowName, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowListByOwnerAndName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwnerAndName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowListByOwnerAndName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwnerAndName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowOwnerConfig(opts *bind.CallOpts, owner common.Address) ([]byte, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowOwnerConfig", owner)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowOwnerConfig(owner common.Address) ([]byte, error) {
	return _WorkflowRegistry.Contract.GetWorkflowOwnerConfig(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowOwnerConfig(owner common.Address) ([]byte, error) {
	return _WorkflowRegistry.Contract.GetWorkflowOwnerConfig(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isAllowedSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsAllowedSigner(signer common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsAllowedSigner(&_WorkflowRegistry.CallOpts, signer)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsAllowedSigner(signer common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsAllowedSigner(&_WorkflowRegistry.CallOpts, signer)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isOwnerLinked", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsOwnerLinked(owner common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsOwnerLinked(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsOwnerLinked(owner common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsOwnerLinked(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) IsRequestAllowlisted(opts *bind.CallOpts, owner common.Address, requestDigest [32]byte) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isRequestAllowlisted", owner, requestDigest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsRequestAllowlisted(owner common.Address, requestDigest [32]byte) (bool, error) {
	return _WorkflowRegistry.Contract.IsRequestAllowlisted(&_WorkflowRegistry.CallOpts, owner, requestDigest)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsRequestAllowlisted(owner common.Address, requestDigest [32]byte) (bool, error) {
	return _WorkflowRegistry.Contract.IsRequestAllowlisted(&_WorkflowRegistry.CallOpts, owner, requestDigest)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) Owner() (common.Address, error) {
	return _WorkflowRegistry.Contract.Owner(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) Owner() (common.Address, error) {
	return _WorkflowRegistry.Contract.Owner(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalActiveWorkflowsByOwner(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalActiveWorkflowsByOwner", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalActiveWorkflowsByOwner(owner common.Address) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsByOwner(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalActiveWorkflowsByOwner(owner common.Address) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsByOwner(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalActiveWorkflowsOnDON(opts *bind.CallOpts, donFamily string) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalActiveWorkflowsOnDON", donFamily)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalActiveWorkflowsOnDON(donFamily string) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsOnDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalActiveWorkflowsOnDON(donFamily string) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsOnDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalEvents(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalEvents")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalEvents() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalEvents(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalEvents() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalEvents(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalLinkedOwners")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalLinkedOwners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalLinkedOwners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalLinkedOwners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalLinkedOwners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TypeAndVersion() (string, error) {
	return _WorkflowRegistry.Contract.TypeAndVersion(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TypeAndVersion() (string, error) {
	return _WorkflowRegistry.Contract.TypeAndVersion(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_WorkflowRegistry *WorkflowRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AcceptOwnership(&_WorkflowRegistry.TransactOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AcceptOwnership(&_WorkflowRegistry.TransactOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "activateWorkflow", workflowId, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) ActivateWorkflow(workflowId [32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowId, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) ActivateWorkflow(workflowId [32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowId, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminBatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminBatchPauseWorkflows", workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminBatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminBatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminBatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminBatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseAllByDON(opts *bind.TransactOpts, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseAllByDON", donFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseAllByDON(donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByDON(&_WorkflowRegistry.TransactOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseAllByDON(donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByDON(&_WorkflowRegistry.TransactOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseAllByOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseAllByOwner", owner)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseAllByOwner(owner common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByOwner(&_WorkflowRegistry.TransactOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseAllByOwner(owner common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByOwner(&_WorkflowRegistry.TransactOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseWorkflow", workflowId)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AllowlistRequest(opts *bind.TransactOpts, requestDigest [32]byte, expiryTimestamp *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "allowlistRequest", requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AllowlistRequest(requestDigest [32]byte, expiryTimestamp *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AllowlistRequest(&_WorkflowRegistry.TransactOpts, requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AllowlistRequest(requestDigest [32]byte, expiryTimestamp *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AllowlistRequest(&_WorkflowRegistry.TransactOpts, requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "batchActivateWorkflows", workflowIds, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) BatchActivateWorkflows(workflowIds [][32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) BatchActivateWorkflows(workflowIds [][32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) BatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "batchPauseWorkflows", workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistrySession) BatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) BatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) DeleteWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "deleteWorkflow", workflowId)
}

func (_WorkflowRegistry *WorkflowRegistrySession) DeleteWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.DeleteWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) DeleteWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.DeleteWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "linkOwner", validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistrySession) LinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.LinkOwner(&_WorkflowRegistry.TransactOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) LinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.LinkOwner(&_WorkflowRegistry.TransactOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) PauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "pauseWorkflow", workflowId)
}

func (_WorkflowRegistry *WorkflowRegistrySession) PauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) PauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONLimit(opts *bind.TransactOpts, donFamily string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONLimit", donFamily, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONLimit(donFamily string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donFamily, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONLimit(donFamily string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donFamily, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONRegistry", registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONRegistry(registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONRegistry(registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetMetadataConfig(opts *bind.TransactOpts, nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setMetadataConfig", nameLen, tagLen, urlLen, attrLen)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetMetadataConfig(nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetMetadataConfig(&_WorkflowRegistry.TransactOpts, nameLen, tagLen, urlLen, attrLen)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetMetadataConfig(nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetMetadataConfig(&_WorkflowRegistry.TransactOpts, nameLen, tagLen, urlLen, attrLen)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setUserDONOverride", user, donFamily, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetUserDONOverride(user common.Address, donFamily string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donFamily, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetUserDONOverride(user common.Address, donFamily string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donFamily, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetWorkflowOwnerConfig(opts *bind.TransactOpts, owner common.Address, config []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setWorkflowOwnerConfig", owner, config)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetWorkflowOwnerConfig(owner common.Address, config []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetWorkflowOwnerConfig(&_WorkflowRegistry.TransactOpts, owner, config)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetWorkflowOwnerConfig(owner common.Address, config []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetWorkflowOwnerConfig(&_WorkflowRegistry.TransactOpts, owner, config)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_WorkflowRegistry *WorkflowRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.TransferOwnership(&_WorkflowRegistry.TransactOpts, to)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.TransferOwnership(&_WorkflowRegistry.TransactOpts, to)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "unlinkOwner", owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "updateAllowedSigners", signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpdateAllowedSigners(signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateAllowedSigners(&_WorkflowRegistry.TransactOpts, signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpdateAllowedSigners(signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateAllowedSigners(&_WorkflowRegistry.TransactOpts, signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpdateWorkflowDONFamily(opts *bind.TransactOpts, workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "updateWorkflowDONFamily", workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpdateWorkflowDONFamily(workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateWorkflowDONFamily(&_WorkflowRegistry.TransactOpts, workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpdateWorkflowDONFamily(workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateWorkflowDONFamily(&_WorkflowRegistry.TransactOpts, workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "upsertWorkflow", workflowName, tag, workflowId, status, donFamily, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowId, status, donFamily, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowId, status, donFamily, binaryUrl, configUrl, attributes, keepAlive)
}

type WorkflowRegistryAllowedSignersUpdatedIterator struct {
	Event *WorkflowRegistryAllowedSignersUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryAllowedSignersUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryAllowedSignersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryAllowedSignersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryAllowedSignersUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryAllowedSignersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryAllowedSignersUpdated struct {
	Signers []common.Address
	Allowed bool
	Raw     types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "AllowedSignersUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryAllowedSignersUpdatedIterator{contract: _WorkflowRegistry.contract, event: "AllowedSignersUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "AllowedSignersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryAllowedSignersUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "AllowedSignersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error) {
	event := new(WorkflowRegistryAllowedSignersUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "AllowedSignersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryDONLimitSetIterator struct {
	Event *WorkflowRegistryDONLimitSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryDONLimitSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryDONLimitSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryDONLimitSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryDONLimitSetIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryDONLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryDONLimitSet struct {
	DonFamily string
	Limit     uint32
	Raw       types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterDONLimitSet(opts *bind.FilterOpts) (*WorkflowRegistryDONLimitSetIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "DONLimitSet")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryDONLimitSetIterator{contract: _WorkflowRegistry.contract, event: "DONLimitSet", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "DONLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryDONLimitSet)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "DONLimitSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseDONLimitSet(log types.Log) (*WorkflowRegistryDONLimitSet, error) {
	event := new(WorkflowRegistryDONLimitSet)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "DONLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryDONRegistryUpdatedIterator struct {
	Event *WorkflowRegistryDONRegistryUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryDONRegistryUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryDONRegistryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryDONRegistryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryDONRegistryUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryDONRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryDONRegistryUpdated struct {
	OldAddr          common.Address
	NewAddr          common.Address
	OldChainSelector uint64
	NewChainSelector uint64
	Raw              types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterDONRegistryUpdated(opts *bind.FilterOpts) (*WorkflowRegistryDONRegistryUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "DONRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryDONRegistryUpdatedIterator{contract: _WorkflowRegistry.contract, event: "DONRegistryUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchDONRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "DONRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryDONRegistryUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "DONRegistryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseDONRegistryUpdated(log types.Log) (*WorkflowRegistryDONRegistryUpdated, error) {
	event := new(WorkflowRegistryDONRegistryUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "DONRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryMetadataConfigUpdatedIterator struct {
	Event *WorkflowRegistryMetadataConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryMetadataConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryMetadataConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryMetadataConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryMetadataConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryMetadataConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryMetadataConfigUpdated struct {
	MaxNameLen uint8
	MaxTagLen  uint8
	MaxUrlLen  uint8
	MaxAttrLen uint16
	Raw        types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterMetadataConfigUpdated(opts *bind.FilterOpts) (*WorkflowRegistryMetadataConfigUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "MetadataConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryMetadataConfigUpdatedIterator{contract: _WorkflowRegistry.contract, event: "MetadataConfigUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchMetadataConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryMetadataConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "MetadataConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryMetadataConfigUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "MetadataConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseMetadataConfigUpdated(log types.Log) (*WorkflowRegistryMetadataConfigUpdated, error) {
	event := new(WorkflowRegistryMetadataConfigUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "MetadataConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipLinkUpdatedIterator struct {
	Event *WorkflowRegistryOwnershipLinkUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipLinkUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipLinkUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryOwnershipLinkUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryOwnershipLinkUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipLinkUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipLinkUpdated struct {
	Owner common.Address
	Proof [32]byte
	Added bool
	Raw   types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipLinkUpdated(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipLinkUpdated", ownerRule, proofRule, addedRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipLinkUpdatedIterator{contract: _WorkflowRegistry.contract, event: "OwnershipLinkUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipLinkUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdated, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipLinkUpdated", ownerRule, proofRule, addedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipLinkUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipLinkUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipLinkUpdated(log types.Log) (*WorkflowRegistryOwnershipLinkUpdated, error) {
	event := new(WorkflowRegistryOwnershipLinkUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipLinkUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipTransferRequestedIterator struct {
	Event *WorkflowRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipTransferRequestedIterator{contract: _WorkflowRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipTransferRequested)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error) {
	event := new(WorkflowRegistryOwnershipTransferRequested)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipTransferredIterator struct {
	Event *WorkflowRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipTransferredIterator{contract: _WorkflowRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipTransferred)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error) {
	event := new(WorkflowRegistryOwnershipTransferred)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryRequestAllowlistedIterator struct {
	Event *WorkflowRegistryRequestAllowlisted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryRequestAllowlistedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryRequestAllowlisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryRequestAllowlisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryRequestAllowlistedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryRequestAllowlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryRequestAllowlisted struct {
	Owner           common.Address
	RequestDigest   [32]byte
	ExpiryTimestamp *big.Int
	Raw             types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterRequestAllowlisted(opts *bind.FilterOpts, owner []common.Address, requestDigest [][32]byte) (*WorkflowRegistryRequestAllowlistedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestDigestRule []interface{}
	for _, requestDigestItem := range requestDigest {
		requestDigestRule = append(requestDigestRule, requestDigestItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "RequestAllowlisted", ownerRule, requestDigestRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryRequestAllowlistedIterator{contract: _WorkflowRegistry.contract, event: "RequestAllowlisted", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchRequestAllowlisted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryRequestAllowlisted, owner []common.Address, requestDigest [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestDigestRule []interface{}
	for _, requestDigestItem := range requestDigest {
		requestDigestRule = append(requestDigestRule, requestDigestItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "RequestAllowlisted", ownerRule, requestDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryRequestAllowlisted)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "RequestAllowlisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseRequestAllowlisted(log types.Log) (*WorkflowRegistryRequestAllowlisted, error) {
	event := new(WorkflowRegistryRequestAllowlisted)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "RequestAllowlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryUserDONLimitSetIterator struct {
	Event *WorkflowRegistryUserDONLimitSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryUserDONLimitSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryUserDONLimitSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryUserDONLimitSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryUserDONLimitSetIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryUserDONLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryUserDONLimitSet struct {
	User      common.Address
	DonFamily string
	Limit     uint32
	Raw       types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterUserDONLimitSet(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "UserDONLimitSet", userRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryUserDONLimitSetIterator{contract: _WorkflowRegistry.contract, event: "UserDONLimitSet", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchUserDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "UserDONLimitSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryUserDONLimitSet)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseUserDONLimitSet(log types.Log) (*WorkflowRegistryUserDONLimitSet, error) {
	event := new(WorkflowRegistryUserDONLimitSet)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryUserDONLimitUnsetIterator struct {
	Event *WorkflowRegistryUserDONLimitUnset

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryUserDONLimitUnsetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryUserDONLimitUnset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryUserDONLimitUnset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryUserDONLimitUnsetIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryUserDONLimitUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryUserDONLimitUnset struct {
	User      common.Address
	DonFamily string
	Raw       types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterUserDONLimitUnset(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitUnsetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "UserDONLimitUnset", userRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryUserDONLimitUnsetIterator{contract: _WorkflowRegistry.contract, event: "UserDONLimitUnset", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchUserDONLimitUnset(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitUnset, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "UserDONLimitUnset", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryUserDONLimitUnset)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitUnset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseUserDONLimitUnset(log types.Log) (*WorkflowRegistryUserDONLimitUnset, error) {
	event := new(WorkflowRegistryUserDONLimitUnset)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowActivatedIterator struct {
	Event *WorkflowRegistryWorkflowActivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowActivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryWorkflowActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryWorkflowActivatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowActivated struct {
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowActivated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowActivatedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowActivated", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowActivatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowActivated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowActivated", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowActivated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowActivated(log types.Log) (*WorkflowRegistryWorkflowActivated, error) {
	event := new(WorkflowRegistryWorkflowActivated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowDeletedIterator struct {
	Event *WorkflowRegistryWorkflowDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryWorkflowDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryWorkflowDeletedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowDeleted struct {
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowDeleted(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDeletedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowDeleted", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowDeletedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowDeleted", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowDeleted", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowDeleted)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowDeleted(log types.Log) (*WorkflowRegistryWorkflowDeleted, error) {
	event := new(WorkflowRegistryWorkflowDeleted)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowDonFamilyUpdatedIterator struct {
	Event *WorkflowRegistryWorkflowDonFamilyUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowDonFamilyUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowDonFamilyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryWorkflowDonFamilyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryWorkflowDonFamilyUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowDonFamilyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowDonFamilyUpdated struct {
	WorkflowId   [32]byte
	Owner        common.Address
	OldDonFamily string
	NewDonFamily string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowDonFamilyUpdated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDonFamilyUpdatedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowDonFamilyUpdated", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowDonFamilyUpdatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowDonFamilyUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowDonFamilyUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDonFamilyUpdated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowDonFamilyUpdated", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowDonFamilyUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDonFamilyUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowDonFamilyUpdated(log types.Log) (*WorkflowRegistryWorkflowDonFamilyUpdated, error) {
	event := new(WorkflowRegistryWorkflowDonFamilyUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDonFamilyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowOwnerConfigUpdatedIterator struct {
	Event *WorkflowRegistryWorkflowOwnerConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowOwnerConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowOwnerConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryWorkflowOwnerConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryWorkflowOwnerConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowOwnerConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowOwnerConfigUpdated struct {
	Owner  common.Address
	Config []byte
	Raw    types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowOwnerConfigUpdated(opts *bind.FilterOpts, owner []common.Address) (*WorkflowRegistryWorkflowOwnerConfigUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowOwnerConfigUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowOwnerConfigUpdatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowOwnerConfigUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowOwnerConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowOwnerConfigUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowOwnerConfigUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowOwnerConfigUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowOwnerConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowOwnerConfigUpdated(log types.Log) (*WorkflowRegistryWorkflowOwnerConfigUpdated, error) {
	event := new(WorkflowRegistryWorkflowOwnerConfigUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowOwnerConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowPausedIterator struct {
	Event *WorkflowRegistryWorkflowPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryWorkflowPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryWorkflowPausedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowPaused struct {
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowPaused(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowPausedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowPaused", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowPausedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowPaused", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowPaused", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowPaused)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowPaused(log types.Log) (*WorkflowRegistryWorkflowPaused, error) {
	event := new(WorkflowRegistryWorkflowPaused)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowRegisteredIterator struct {
	Event *WorkflowRegistryWorkflowRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryWorkflowRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryWorkflowRegisteredIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowRegistered struct {
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	Status       uint8
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowRegistered(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowRegisteredIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowRegistered", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowRegisteredIterator{contract: _WorkflowRegistry.contract, event: "WorkflowRegistered", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowRegistered, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowRegistered", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowRegistered)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowRegistered(log types.Log) (*WorkflowRegistryWorkflowRegistered, error) {
	event := new(WorkflowRegistryWorkflowRegistered)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowUpdatedIterator struct {
	Event *WorkflowRegistryWorkflowUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(WorkflowRegistryWorkflowUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *WorkflowRegistryWorkflowUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowUpdated struct {
	OldWorkflowId [32]byte
	NewWorkflowId [32]byte
	Owner         common.Address
	DonFamily     string
	WorkflowName  string
	Raw           types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowUpdated(opts *bind.FilterOpts, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowUpdatedIterator, error) {

	var oldWorkflowIdRule []interface{}
	for _, oldWorkflowIdItem := range oldWorkflowId {
		oldWorkflowIdRule = append(oldWorkflowIdRule, oldWorkflowIdItem)
	}
	var newWorkflowIdRule []interface{}
	for _, newWorkflowIdItem := range newWorkflowId {
		newWorkflowIdRule = append(newWorkflowIdRule, newWorkflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowUpdated", oldWorkflowIdRule, newWorkflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowUpdatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowUpdated, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var oldWorkflowIdRule []interface{}
	for _, oldWorkflowIdItem := range oldWorkflowId {
		oldWorkflowIdRule = append(oldWorkflowIdRule, oldWorkflowIdItem)
	}
	var newWorkflowIdRule []interface{}
	for _, newWorkflowIdItem := range newWorkflowId {
		newWorkflowIdRule = append(newWorkflowIdRule, newWorkflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowUpdated", oldWorkflowIdRule, newWorkflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowUpdated(log types.Log) (*WorkflowRegistryWorkflowUpdated, error) {
	event := new(WorkflowRegistryWorkflowUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_WorkflowRegistry *WorkflowRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WorkflowRegistry.abi.Events["AllowedSignersUpdated"].ID:
		return _WorkflowRegistry.ParseAllowedSignersUpdated(log)
	case _WorkflowRegistry.abi.Events["DONLimitSet"].ID:
		return _WorkflowRegistry.ParseDONLimitSet(log)
	case _WorkflowRegistry.abi.Events["DONRegistryUpdated"].ID:
		return _WorkflowRegistry.ParseDONRegistryUpdated(log)
	case _WorkflowRegistry.abi.Events["MetadataConfigUpdated"].ID:
		return _WorkflowRegistry.ParseMetadataConfigUpdated(log)
	case _WorkflowRegistry.abi.Events["OwnershipLinkUpdated"].ID:
		return _WorkflowRegistry.ParseOwnershipLinkUpdated(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferRequested(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferred"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferred(log)
	case _WorkflowRegistry.abi.Events["RequestAllowlisted"].ID:
		return _WorkflowRegistry.ParseRequestAllowlisted(log)
	case _WorkflowRegistry.abi.Events["UserDONLimitSet"].ID:
		return _WorkflowRegistry.ParseUserDONLimitSet(log)
	case _WorkflowRegistry.abi.Events["UserDONLimitUnset"].ID:
		return _WorkflowRegistry.ParseUserDONLimitUnset(log)
	case _WorkflowRegistry.abi.Events["WorkflowActivated"].ID:
		return _WorkflowRegistry.ParseWorkflowActivated(log)
	case _WorkflowRegistry.abi.Events["WorkflowDeleted"].ID:
		return _WorkflowRegistry.ParseWorkflowDeleted(log)
	case _WorkflowRegistry.abi.Events["WorkflowDonFamilyUpdated"].ID:
		return _WorkflowRegistry.ParseWorkflowDonFamilyUpdated(log)
	case _WorkflowRegistry.abi.Events["WorkflowOwnerConfigUpdated"].ID:
		return _WorkflowRegistry.ParseWorkflowOwnerConfigUpdated(log)
	case _WorkflowRegistry.abi.Events["WorkflowPaused"].ID:
		return _WorkflowRegistry.ParseWorkflowPaused(log)
	case _WorkflowRegistry.abi.Events["WorkflowRegistered"].ID:
		return _WorkflowRegistry.ParseWorkflowRegistered(log)
	case _WorkflowRegistry.abi.Events["WorkflowUpdated"].ID:
		return _WorkflowRegistry.ParseWorkflowUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (WorkflowRegistryAllowedSignersUpdated) Topic() common.Hash {
	return common.HexToHash("0x861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e1316")
}

func (WorkflowRegistryDONLimitSet) Topic() common.Hash {
	return common.HexToHash("0x4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b731")
}

func (WorkflowRegistryDONRegistryUpdated) Topic() common.Hash {
	return common.HexToHash("0xa7a2a5335a8d1f8f1f7ef8a58332be349ac9fdc25b62512290a91ac4555430a5")
}

func (WorkflowRegistryMetadataConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xa254b4f533572ce0e775da29b44e0114b25637b1072bbeaf3fe051fb55aafedc")
}

func (WorkflowRegistryOwnershipLinkUpdated) Topic() common.Hash {
	return common.HexToHash("0x07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35")
}

func (WorkflowRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (WorkflowRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (WorkflowRegistryRequestAllowlisted) Topic() common.Hash {
	return common.HexToHash("0xd2e6cd8f534fad4a57838c5102996a78ebe4bd24ca8443437528c6c9295e4b96")
}

func (WorkflowRegistryUserDONLimitSet) Topic() common.Hash {
	return common.HexToHash("0x945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c82324")
}

func (WorkflowRegistryUserDONLimitUnset) Topic() common.Hash {
	return common.HexToHash("0x6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d849")
}

func (WorkflowRegistryWorkflowActivated) Topic() common.Hash {
	return common.HexToHash("0x7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a")
}

func (WorkflowRegistryWorkflowDeleted) Topic() common.Hash {
	return common.HexToHash("0x48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd94")
}

func (WorkflowRegistryWorkflowDonFamilyUpdated) Topic() common.Hash {
	return common.HexToHash("0x9b5361a5258ef6ac8039fd2d2ac276734695219cfd870711e7922c236e5db16d")
}

func (WorkflowRegistryWorkflowOwnerConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f13")
}

func (WorkflowRegistryWorkflowPaused) Topic() common.Hash {
	return common.HexToHash("0xf764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e488689")
}

func (WorkflowRegistryWorkflowRegistered) Topic() common.Hash {
	return common.HexToHash("0x74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a09")
}

func (WorkflowRegistryWorkflowUpdated) Topic() common.Hash {
	return common.HexToHash("0x03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e322")
}

func (_WorkflowRegistry *WorkflowRegistry) Address() common.Address {
	return _WorkflowRegistry.address
}

type WorkflowRegistryInterface interface {
	CanLinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error

	CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error

	GetDONRegistry(opts *bind.CallOpts) (common.Address, uint64, error)

	GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetMaxAttrLen(opts *bind.CallOpts) (uint16, error)

	GetMaxNameLen(opts *bind.CallOpts) (uint8, error)

	GetMaxTagLen(opts *bind.CallOpts) (uint8, error)

	GetMaxUrlLen(opts *bind.CallOpts) (uint8, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (uint32, error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donFamily string) (uint32, error)

	GetWorkflow(opts *bind.CallOpts, owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowById(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByDON(opts *bind.CallOpts, donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByOwnerAndName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowOwnerConfig(opts *bind.CallOpts, owner common.Address) ([]byte, error)

	IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error)

	IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error)

	IsRequestAllowlisted(opts *bind.CallOpts, owner common.Address, requestDigest [32]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TotalActiveWorkflowsByOwner(opts *bind.CallOpts, owner common.Address) (*big.Int, error)

	TotalActiveWorkflowsOnDON(opts *bind.CallOpts, donFamily string) (*big.Int, error)

	TotalEvents(opts *bind.CallOpts) (*big.Int, error)

	TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte, donFamily string) (*types.Transaction, error)

	AdminBatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	AdminPauseAllByDON(opts *bind.TransactOpts, donFamily string) (*types.Transaction, error)

	AdminPauseAllByOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error)

	AdminPauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	AllowlistRequest(opts *bind.TransactOpts, requestDigest [32]byte, expiryTimestamp *big.Int) (*types.Transaction, error)

	BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte, donFamily string) (*types.Transaction, error)

	BatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	DeleteWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error)

	PauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	SetDONLimit(opts *bind.TransactOpts, donFamily string, limit uint32, enabled bool) (*types.Transaction, error)

	SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector uint64) (*types.Transaction, error)

	SetMetadataConfig(opts *bind.TransactOpts, nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, limit uint32, enabled bool) (*types.Transaction, error)

	SetWorkflowOwnerConfig(opts *bind.TransactOpts, owner common.Address, config []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	UpdateWorkflowDONFamily(opts *bind.TransactOpts, workflowId [32]byte, newDonFamily string) (*types.Transaction, error)

	UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error)

	FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error)

	WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error)

	ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error)

	FilterDONLimitSet(opts *bind.FilterOpts) (*WorkflowRegistryDONLimitSetIterator, error)

	WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet) (event.Subscription, error)

	ParseDONLimitSet(log types.Log) (*WorkflowRegistryDONLimitSet, error)

	FilterDONRegistryUpdated(opts *bind.FilterOpts) (*WorkflowRegistryDONRegistryUpdatedIterator, error)

	WatchDONRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONRegistryUpdated) (event.Subscription, error)

	ParseDONRegistryUpdated(log types.Log) (*WorkflowRegistryDONRegistryUpdated, error)

	FilterMetadataConfigUpdated(opts *bind.FilterOpts) (*WorkflowRegistryMetadataConfigUpdatedIterator, error)

	WatchMetadataConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryMetadataConfigUpdated) (event.Subscription, error)

	ParseMetadataConfigUpdated(log types.Log) (*WorkflowRegistryMetadataConfigUpdated, error)

	FilterOwnershipLinkUpdated(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedIterator, error)

	WatchOwnershipLinkUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdated, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error)

	ParseOwnershipLinkUpdated(log types.Log) (*WorkflowRegistryOwnershipLinkUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error)

	FilterRequestAllowlisted(opts *bind.FilterOpts, owner []common.Address, requestDigest [][32]byte) (*WorkflowRegistryRequestAllowlistedIterator, error)

	WatchRequestAllowlisted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryRequestAllowlisted, owner []common.Address, requestDigest [][32]byte) (event.Subscription, error)

	ParseRequestAllowlisted(log types.Log) (*WorkflowRegistryRequestAllowlisted, error)

	FilterUserDONLimitSet(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitSetIterator, error)

	WatchUserDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitSet, user []common.Address) (event.Subscription, error)

	ParseUserDONLimitSet(log types.Log) (*WorkflowRegistryUserDONLimitSet, error)

	FilterUserDONLimitUnset(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitUnsetIterator, error)

	WatchUserDONLimitUnset(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitUnset, user []common.Address) (event.Subscription, error)

	ParseUserDONLimitUnset(log types.Log) (*WorkflowRegistryUserDONLimitUnset, error)

	FilterWorkflowActivated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowActivatedIterator, error)

	WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowActivated(log types.Log) (*WorkflowRegistryWorkflowActivated, error)

	FilterWorkflowDeleted(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDeletedIterator, error)

	WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowDeleted(log types.Log) (*WorkflowRegistryWorkflowDeleted, error)

	FilterWorkflowDonFamilyUpdated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDonFamilyUpdatedIterator, error)

	WatchWorkflowDonFamilyUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDonFamilyUpdated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowDonFamilyUpdated(log types.Log) (*WorkflowRegistryWorkflowDonFamilyUpdated, error)

	FilterWorkflowOwnerConfigUpdated(opts *bind.FilterOpts, owner []common.Address) (*WorkflowRegistryWorkflowOwnerConfigUpdatedIterator, error)

	WatchWorkflowOwnerConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowOwnerConfigUpdated, owner []common.Address) (event.Subscription, error)

	ParseWorkflowOwnerConfigUpdated(log types.Log) (*WorkflowRegistryWorkflowOwnerConfigUpdated, error)

	FilterWorkflowPaused(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowPausedIterator, error)

	WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowPaused(log types.Log) (*WorkflowRegistryWorkflowPaused, error)

	FilterWorkflowRegistered(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowRegisteredIterator, error)

	WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowRegistered, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowRegistered(log types.Log) (*WorkflowRegistryWorkflowRegistered, error)

	FilterWorkflowUpdated(opts *bind.FilterOpts, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowUpdatedIterator, error)

	WatchWorkflowUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowUpdated, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowUpdated(log types.Log) (*WorkflowRegistryWorkflowUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
