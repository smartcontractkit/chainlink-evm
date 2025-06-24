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

type WorkflowRegistryMetadataConfig struct {
	MaxWorkflowNameLength uint8
	MaxWorkflowTagLength  uint8
	MaxUrlLength          uint8
	MaxAttributesLength   uint16
}

type WorkflowRegistryWorkflowMetadata struct {
	WorkflowId   [32]byte
	Owner        common.Address
	CreatedAt    uint64
	Status       uint8
	WorkflowName string
	BinaryUrl    string
	ConfigUrl    string
	Tag          string
	Attributes   []byte
}

var WorkflowRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDONRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"maxWorkflows\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.MetadataConfig\",\"components\":[{\"name\":\"maxWorkflowNameLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxWorkflowTagLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxUrlLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxAttributesLength\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadata\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwnerAndName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataConfig\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.MetadataConfig\",\"components\":[{\"name\":\"maxWorkflowNameLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxWorkflowTagLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxUrlLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxAttributesLength\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWorkflowDONLabel\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataConfigUpdated\",\"inputs\":[{\"name\":\"maxWorkflowNameLength\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"maxWorkflowTagLength\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"maxUrlLength\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"maxAttributesLength\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDonFamilyUpdated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowOwnerConfigUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUnlinkWithActiveWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DonLimitNotSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x608034609d573315608c57600180546001600160a01b03191633179055608081016001600160401b03811182821017607657610400916060916040526040815260208082015260c860408201520152640400c8000061204064ffffffffff19601254161717601255604051615d3c90816100a38239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fdfe610120604052600436101561001357600080fd5b60003560e01c8063034a6ed5146133335780630987294c146132c457806317e0edfc14613180578063181f5a77146131465780631c71682c14612f20578063274e00e014612e795780632afc413014612df65780633514125114612d795780633c17181b14612cff57806342aa3a8414612bdf57806347d1ed8314612b84578063530979d614612b6d5780635a1ac5ad14612b09578063695e134014612a8b5780636ee80b441461287b57806379ba509714612792578063867c0721146127735780638b42a96d146125ba5780638da5cb5b1461256857806395be176e1461250c57806396372428146123f65780639c5b23041461231d578063a0b8a4fe146122e1578063a7d018581461221c578063b377bfc514610ff0578063b668435f14610f6a578063b87a019414610da4578063bdf6b4ff14610d23578063be67433314610cfc578063cabb9e7a14610c92578063d8b8073814610b94578063d8e056de14610aed578063d8e4a72414610923578063dc1019691461088f578063dfcb0b3114610875578063e086ec96146107e7578063e690f332146107ab578063e9df6553146102bf5763f2fde38b146101ca57600080fd5b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5773ffffffffffffffffffffffffffffffffffffffff6102166135e3565b61021e6149f6565b1633811461029057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043567ffffffffffffffff81116102ba5761030e9036906004016138fd565b80156107815761032b336000526004602052604060002054151590565b156107535761033981613abb565b610346604051918261368f565b8181527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061037383613abb565b0136602083013761038382613abb565b610390604051918261368f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06103bd84613abb565b01366020830137600091825b8481106105bc5750339060005b84811061044857868660005b8181106103eb57005b806103f96001928486614293565b35600052600960205260406000205480600052600760205260406000208360ff8183015460e01c1661042a8161348a565b14610438575b5050016103e2565b61044191615104565b8480610430565b6104528183613ae0565b5163ffffffff6104628387613ae0565b511690806000526014602052604060002091610480600184016145a5565b6020810151156105815763ffffffff905116916104aa6040600089815260028701602052206145a5565b6020810151610572575b5033600052600f60205260406000209060005260205263ffffffff60406000205416019063ffffffff82116105435763ffffffff8091169116116104fb57506001016103d6565b61053f906040519182917f038857ff000000000000000000000000000000000000000000000000000000008352336004840152604060248401526044830190613a19565b0390fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5163ffffffff1692508b6104b4565b6040517f727e650e000000000000000000000000000000000000000000000000000000008152602060048201528061053f6024820187613a19565b90926105c9828688614293565b35600052600960205260406000205480600052600760205260016040600020015473ffffffffffffffffffffffffffffffffffffffff811680156107295733036106fb5760e01c60ff1661061c8161348a565b156106f15760005260106020526040600020546000805b83811061069b575b5015610651575b50906001905b019290926103c9565b61065b8286613ae0565b5260016106688285613ae0565b527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461054357600180910190610642565b826106a68289613ae0565b51146106b457600101610633565b905063ffffffff6106c58287613ae0565b51169063ffffffff82146105435760016106e463ffffffff9288613ae0565b920116905260018861063b565b5090600190610648565b7f31ee6dc7000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc2dda3f9000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba576107e560043561492f565b005b346102ba5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba576000606060405161082681613673565b828152826020820152826040820152015260806108416142a3565b61ffff60606040519260ff815116845260ff602082015116602085015260ff60408201511660408501520151166060820152f35b346102ba576107e56108863661392e565b929190916146a5565b346102ba5760016108ab6108a23661392e565b918094936146a5565b336000526005602052806040600020556108c4336158d2565b508060005260066020526040600020827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346102ba5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043567ffffffffffffffff81116102ba576109729036906004016138fd565b602435918215158093036102ba576109886149f6565b60ff831660005b838110610a3357505060405191806040840160408552526060830191906000905b8082106109e6577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b90919283359073ffffffffffffffffffffffffffffffffffffffff821682036102ba576020809173ffffffffffffffffffffffffffffffffffffffff6001941681520194019201906109b0565b73ffffffffffffffffffffffffffffffffffffffff610a5b610a56838787614293565b61464b565b1615610ac3578073ffffffffffffffffffffffffffffffffffffffff610a87610a566001948888614293565b1660005260026020526040600020837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790550161098f565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57610b246149f6565b600435600052600b60205260406000205b805480156107e5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161054357610b74610b8f918361550d565b90549060031b1c806000526007602052604060002090614a41565b610b35565b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043567ffffffffffffffff81116102ba57610be39036906004016138fd565b801561078157610c00336000526004602052604060002054151590565b156107535760005b818110610c1157005b610c1c818385614293565b3560005260096020526040600020549081600052600760205260406000209160018301549273ffffffffffffffffffffffffffffffffffffffff841680156107295733036106fb57600160ff819560e01c16610c778161348a565b14610c8b57610c8591614a41565b01610c08565b5050610c85565b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5773ffffffffffffffffffffffffffffffffffffffff610cde6135e3565b166000526002602052602060ff604060002054166040519015158152f35b346102ba576020610d15610d0f36613814565b916145dd565b63ffffffff60405191168152f35b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043567ffffffffffffffff81116102ba57610d75610d7c913690600401613715565b3691613b3f565b602081519101206000526014602052602063ffffffff60016040600020015416604051908152f35b346102ba5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57610ddb6135e3565b73ffffffffffffffffffffffffffffffffffffffff60243591169081600052600d602052610e11604435826040600020546149b3565b90610e1b82614226565b9260005b838110610e385760405180610e34878261387f565b0390f35b60019082600052600d602052610e5c6040600020610e568387613ad3565b9061550d565b90549060031b1c600052600760205260406000206006610f4860405192610e828461363a565b8054845260ff8682015473ffffffffffffffffffffffffffffffffffffffff8116602087015267ffffffffffffffff8160a01c16604087015260e01c16610ec88161348a565b6060850152604051610ee881610ee18160028601613a19565b038261368f565b6080850152604051610f0181610ee18160038601613a19565b60a0850152604051610f1a81610ee18160048601613a19565b60c0850152604051610f3381610ee18160058601613a19565b60e0850152610ee16040518094819301613a19565b610100820152610f588288613ae0565b52610f638187613ae0565b5001610e1f565b346102ba5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57610fa16135e3565b60243567ffffffffffffffff81116102ba57610fc1903690600401613715565b6044359063ffffffff821682036102ba576064359283151584036102ba576107e594610feb6149f6565b614305565b346102ba576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043567ffffffffffffffff81116102ba57611040903690600401613715565b60e0526101005260243567ffffffffffffffff81116102ba57611067903690600401613715565b90600260643510156102ba5760843567ffffffffffffffff81116102ba57611093903690600401613715565b9060c05260a43567ffffffffffffffff81116102ba576110b7903690600401613715565b9060c43567ffffffffffffffff81116102ba576110d8903690600401613715565b9560e43567ffffffffffffffff81116102ba576110f9903690600401613715565b929097610104359485151586036102ba57611121336000526004602052604060002054151590565b156107535760443515610ac35760443560005260096020526040600020546121ec5761114b6142a3565b6040810160ff81511680612173575b505061ffff6060820151168015158061216a575b6121395750841561210f5760ff60208201511680151580612106575b6120d5575060e051156120ab575160ff16801515806120a0575b61206d57506111b93660e05161010051613b3f565b61121c6112486111ca36888e613b3f565b6040519283916111ec60208401963388526060604086015260808501906134c3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160608501526134c3565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261368f565b51902096876000526007602052604060002073ffffffffffffffffffffffffffffffffffffffff6001820154168015600014611bfd5750506112996112933660e05161010051613b3f565b33615422565b996112a7368b60c051613b3f565b602081519101209715611b63575b8a9b611365611374928c8c9d9e8c6113569d9e6112d360643561348a565b60643515611b38575b505050506040519a8b9889986112f18a61363a565b6044358a5260a061134a60208c019a338c5260408d01835267ffffffffffffffff421683515260608d019a61132760643561348a565b6064358c52608061133e3660e05161010051613b3f565b9e019d8e523691613b3f565b9b019a8b523691613b3f565b9860c08c01998a523691613b3f565b9760e08a019889523691613b3f565b61010088016080526080515288600052600760205260406000209651875573ffffffffffffffffffffffffffffffffffffffff600188019251167fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825560a051517fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000008085549360a01b161691161782555161142f8161348a565b6114388161348a565b7fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff7cff0000000000000000000000000000000000000000000000000000000083549260e01b16911617905560028501905180519067ffffffffffffffff8211611858576114af826114a985546139c6565b85613bbb565b602090601f8311600114611a95576114fd929160009183611a8a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b518051600385019167ffffffffffffffff821161185857611525826114a985546139c6565b602090601f83116001146119e7576115729291600091836119dc5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b518051600484019167ffffffffffffffff82116118585761159a826114a985546139c6565b602090601f8311600114611939576115e79291600091836118875750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b518051600583019167ffffffffffffffff82116118585761160f826114a985546139c6565b602090601f83116001146118925791806116619260069695946000926118875750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b01936080515194855167ffffffffffffffff8111611858576116908161168a84546139c6565b84613bbb565b6020601f82116001146117b25790806116e19261173798996000926117a75750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b60005260086020526116fa826040600020615967565b50604435600052600960205281604060002055600052600c602052611723816040600020615967565b5033600052600d6020526040600020615967565b506117506040519160608352606083019060c0516140a7565b61175b60643561348a565b606435602083015281810360408301527f74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a093392806117a26044359460e051610100516140a7565b0390a3005b0151905089806114cb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169783600052816000209860005b818110611840575091611737989991846001959410611809575b505050811b0190556116e4565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558880806117fc565b838301518b556001909a0199602093840193016117e2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b015190508a806114cb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b8181106119215750916001939185600698979694106118ea575b505050811b019055611664565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558980806118dd565b929360206001819287860151815501950193016118c3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b8181106119c4575090846001959493921061198d575b505050811b0190556115ea565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611980565b9293602060018192878601518155019501930161196a565b015190508b806114cb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611a725750908460019594939210611a3b575b505050811b019055611575565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558a8080611a2e565b92936020600181928786015181550195019301611a18565b015190508c806114cb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611b205750908460019594939210611ae9575b505050811b019055611500565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558b8080611adc565b92936020600181928786015181550195019301611ac6565b611b53611b4c611b5a95369060c051613b3f565b8333614fdb565b3390615441565b8d8f8d8f6112dc565b9998979695949392919089600052600e60205260406000209b5b8c548015611be357807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81011161054357610b748e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff611bde93019061550d565b611b7d565b509b509899979896979596949593949293919290916112b5565b92949698939550965097985033036106fb578454600052600960205260006040812055604435600052600960205260406000205583549560443585556003850191604051611c4f81610ee18187613a19565b60208151910120611c61368484613b3f565b6020815191012003611f4d575b5050506004830191604051611c8781610ee18187613a19565b60208151910120611c99368484613b3f565b6020815191012003611e24575b5050506006019067ffffffffffffffff811161185857611cca8161168a84546139c6565b6000601f8211600114611d84578190611d1993949596600092611d795750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b611d346040519260408452604084019060c0516140a7565b9082820360208401527f03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e322339380611d746044359560e051610100516140a7565b0390a4005b0135905086806114cb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216958382526020822091805b888110611e0c57508360019596979810611dd4575b505050811b019055611d1c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055858080611dc7565b90926020600181928686013581550194019101611db2565b67ffffffffffffffff821161185857611e41826114a985546139c6565b600090601f8311600114611ea9579180611e93926006969594600092611e9e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b908680611ca6565b013590508a806114cb565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b818110611f35575091600193918560069897969410611efd575b505050811b019055611e96565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055898080611ef0565b91936020600181928787013581550195019201611ed6565b67ffffffffffffffff821161185857611f6a826114a985546139c6565b600090601f8311600114611fcd57611fb7929160009183611fc25750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b878080611c6e565b013590508b806114cb565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b818110612055575090846001959493921061201d575b505050811b019055611fba565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558a8080612010565b91936020600181928787013581550195019201611ffa565b7f36a7c5030000000000000000000000000000000000000000000000000000000060005260e05160045260245260446000fd5b508060e051116111a4565b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b857f436f97540000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5080861161118a565b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b867f354f25140000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5080871161116e565b808a116121bc57505160ff1680841161218c578061115a565b837ecd56a80000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b897ecd56a80000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260443560045260246000fd5b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043567ffffffffffffffff81116102ba5761226b9036906004016138fd565b80156107815760005b81811061227d57005b8061228b6001928486614293565b356122946149f6565b6000526009602052604060002054806000526007602052604060002060ff8482015460e01c166122c38161348a565b156122d1575b505001612274565b6122da91614a41565b84806122c9565b346102ba5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba576020600354604051908152f35b346102ba5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043560243581600052600c60205261236d604435826040600020546149b3565b9061237782614226565b9260005b8381106123905760405180610e34878261387f565b60019082600052600c6020526123ae6040600020610e568387613ad3565b90549060031b1c6000526007602052604060002060066123d460405192610e828461363a565b6101008201526123e48288613ae0565b526123ef8187613ae0565b500161237b565b346102ba5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5761242d6135e3565b60243567ffffffffffffffff81116102ba5761246191612454612467923690600401613715565b9390604435943691613b3f565b90615422565b90816000526008602052612483606435826040600020546149b3565b9061248d82614226565b9260005b8381106124a65760405180610e34878261387f565b6001908260005260086020526124c46040600020610e568387613ad3565b90549060031b1c6000526007602052604060002060066124ea60405192610e828461363a565b6101008201526124fa8288613ae0565b526125058187613ae0565b5001612491565b346102ba5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba576125436135e3565b60243567ffffffffffffffff811681036102ba576107e5916125636149f6565b6140e6565b346102ba5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102ba5773ffffffffffffffffffffffffffffffffffffffff6125dd36613814565b9290916125e86149f6565b169081600052601360205260406000209267ffffffffffffffff81116118585761261c8161261686546139c6565b86613bbb565b600093601f82116001146126b55761268b82807f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f1396976000916126aa575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6126a56040519283926020845260208401916140a7565b0390a2005b90508501358861265a565b808552602085207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0831695805b87811061275b5750837f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f13969710612723575b5050600182811b01905561268e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19908501351690558580612714565b909160206001819285890135815501930191016126e2565b346102ba576107e5612784366137c2565b9161278d6149f6565b613ff6565b346102ba5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760005473ffffffffffffffffffffffffffffffffffffffff81163303612851577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ba5761288936613743565b92906128948461370b565b8315928380612a5e575b612a34576128ac9286614d76565b6128b58261370b565b1561291c575b600073ffffffffffffffffffffffffffffffffffffffff83166128dd8161586f565b9080835260056020528260408120556128f581615aa1565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b9073ffffffffffffffffffffffffffffffffffffffff81169081600052600a6020526040600020905b81548015612a13577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161054357612981908361550d565b90549060031b1c806000526007602052604060002073ffffffffffffffffffffffffffffffffffffffff60018201541680156107295785036129e5576129c68661370b565b600286036129dc576129d791614a41565b612945565b6129d7916156cb565b847f31ee6dc70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b50925073ffffffffffffffffffffffffffffffffffffffff91506128bb9050565b7f61bc2e180000000000000000000000000000000000000000000000000000000060005260046000fd5b5073ffffffffffffffffffffffffffffffffffffffff8616600052600a602052604060002054151561289e565b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba576004356000526009602052604060002054806000526007602052604060002073ffffffffffffffffffffffffffffffffffffffff60018201541680156107295733036106fb576107e5916156cb565b346102ba5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57604060155467ffffffffffffffff82519173ffffffffffffffffffffffffffffffffffffffff8116835260a01c166020820152f35b346102ba576107e5612b7e366137c2565b91613f4a565b346102ba57612b9236613743565b909290612b9e8161370b565b1580612bb2575b612a34576107e593614d76565b5073ffffffffffffffffffffffffffffffffffffffff8416600052600a6020526040600020541515612ba5565b346102ba5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57612c166149f6565b60ff612c20613f09565b166012549061ff00612c30613f19565b60081b167fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffff62ff0000612c61613f29565b60101b16927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000064ffff000000612c95613f39565b60181b1695161716171717601255612cab613f09565b612cb3613f19565b90612cbc613f29565b7fa254b4f533572ce0e775da29b44e0114b25637b1072bbeaf3fe051fb55aafedc602060ff8080612ceb613f39565b9561ffff60405197168752169616941692a4005b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57612d366149f6565b6004356000526009602052604060002054806000526007602052604060002060ff600182015460e01c16612d698161348a565b15612d7057005b6107e591614a41565b346102ba5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5760043567ffffffffffffffff81116102ba57612dc8903690600401613715565b60243563ffffffff811681036102ba576044359182151583036102ba576107e593612df16149f6565b613c02565b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5773ffffffffffffffffffffffffffffffffffffffff612e426135e3565b166000526013602052610e34610ee1612e65604060002060405192838092613a19565b6040519182916020835260208301906134c3565b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5773ffffffffffffffffffffffffffffffffffffffff612ec56135e3565b612ecd6149f6565b16600052600a60205260406000205b805480156107e5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161054357610b74612f1b918361550d565b612edc565b346102ba57612f3d612f3136613606565b816016939293546149b3565b612f4681613abb565b91612f54604051938461368f565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612f8183613abb565b0160005b81811061311a575050601654919060005b82811061304b57846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210612fd657505050500390f35b9193602061303b827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652606060408a51805161301b8161370b565b845263ffffffff86820151168685015201519181604082015201906134c3565b9601920192018594939192612fc7565b6130558183613ad3565b6000858210156130ed57601690526040516001929183906130cc90821b7fd833147d7dc355ba459fc788f669e58cfaf9dc25ddcd0702e87d69c7b51242890161309d84613657565b63ffffffff815460ff81166130b18161370b565b865260081c166020850152610ee16040518094819301613a19565b60408201526130db8288613ae0565b526130e68187613ae0565b5001612f96565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b60209060405161312981613657565b600081526000838201526060604082015282828801015201612f85565b346102ba5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57610e34612e656136d0565b346102ba5761319d61319136613606565b816003939293546149b3565b6131a681613abb565b916131b4604051938461368f565b8183526131c082613abb565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020850193013684376003549160005b8281106132545784866040519182916020830190602084525180915260408301919060005b818110613225575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101613217565b61325e8183613ad3565b6000858210156130ed57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0154600081905260056020526001919073ffffffffffffffffffffffffffffffffffffffff166132bd8289613ae0565b52016131f2565b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba57602061332973ffffffffffffffffffffffffffffffffffffffff6133156135e3565b166000526004602052604060002054151590565b6040519015158152f35b346102ba5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ba5761336a61397f565b5060043560005260096020526040600020548015610729576000526007602052604060002073ffffffffffffffffffffffffffffffffffffffff604051916133b18361363a565b80548352600661346560018301549260ff6020870194868116865267ffffffffffffffff8160a01c16604089015260e01c166133ec8161348a565b606087015260405161340581610ee18160028601613a19565b608087015260405161341e81610ee18160038601613a19565b60a087015260405161343781610ee18160048601613a19565b60c087015260405161345081610ee18160058601613a19565b60e0870152610ee16040518094819301613a19565b61010084015251161561072957610e3490604051918291602083526020830190613522565b6002111561349457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b919082519283825260005b84811061350d5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016134ce565b6135e0918151815273ffffffffffffffffffffffffffffffffffffffff602083015116602082015267ffffffffffffffff6040830151166040820152606082015161356c8161348a565b60608201526101006135ce6135bc6135aa613598608087015161012060808801526101208701906134c3565b60a087015186820360a08801526134c3565b60c086015185820360c08701526134c3565b60e085015184820360e08601526134c3565b920151906101008184039101526134c3565b90565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036102ba57565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60409101126102ba576004359060243590565b610120810190811067ffffffffffffffff82111761185857604052565b6060810190811067ffffffffffffffff82111761185857604052565b6080810190811067ffffffffffffffff82111761185857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761185857604052565b604051906136df60408361368f565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b6003111561349457565b9181601f840112156102ba5782359167ffffffffffffffff83116102ba57602083818601950101116102ba57565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102ba5760043573ffffffffffffffffffffffffffffffffffffffff811681036102ba5791602435916044359067ffffffffffffffff82116102ba576137b191600401613715565b909160643560038110156102ba5790565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126102ba57600435916024359067ffffffffffffffff82116102ba5761381091600401613715565b9091565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126102ba5760043573ffffffffffffffffffffffffffffffffffffffff811681036102ba57916024359067ffffffffffffffff82116102ba5761381091600401613715565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106138b257505050505090565b90919293946020806138ee837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951613522565b970193019301919392906138a3565b9181601f840112156102ba5782359167ffffffffffffffff83116102ba576020808501948460051b0101116102ba57565b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102ba5760043591602435916044359067ffffffffffffffff82116102ba5761381091600401613715565b6040519061398c8261363a565b6060610100836000815260006020820152600060408201526000838201528260808201528260a08201528260c08201528260e08201520152565b90600182811c92168015613a0f575b60208310146139e057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916139d5565b60009291815491613a29836139c6565b8083529260018116908115613a7f5750600114613a4557505050565b60009081526020812093945091925b838310613a65575060209250010190565b600181602092949394548385870101520191019190613a54565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b67ffffffffffffffff81116118585760051b60200190565b9190820180921161054357565b8051821015613af45760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8054821015613af45760005260206000209060011b0190600090565b92919267ffffffffffffffff82116118585760405191613b87601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0166020018461368f565b8294818452818301116102ba578281602093846000960137010152565b818110613baf575050565b60008155600101613ba4565b9190601f8111613bca57505050565b613bf6926000526020600020906020601f840160051c83019310613bf8575b601f0160051c0190613ba4565b565b9091508190613be9565b9091613c0f368484613b3f565b60208151910120600052601460205260406000208054613c2e816139c6565b15613de0575b50600101938015613d7b57845460ff8160201c1615600014613d10575084547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000001663ffffffff831617640100000000179094557f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b731935b15613d0557613d00905b613cf4604051613ccd8161121c858989602085016142e0565b60405190613cda82613657565b6000825263ffffffff421660208301526040820152614b81565b604051938493846142e0565b0390a1565b50613d006000613cb4565b63ffffffff838116911614613d74577f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b7319463ffffffff83167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055613caa565b5050505050565b84549460ff8660201c1615613dd8577fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169094557f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b73193613caa565b505050505050565b67ffffffffffffffff85116118585784613dfc613e02926139c6565b83613bbb565b60009084601f8111600114613e655780613e5191600194600091613e5a57507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b81555b90613c34565b90508601353861265a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169181815260208120908681945b808610613eeb576001955010613eb3575b50508185811b018155613e54565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19908601351690553880613ea5565b87830135845560209586019560019094019390920191889150613e94565b60043560ff811681036102ba5790565b60243560ff811681036102ba5790565b60443560ff811681036102ba5790565b60643561ffff811681036102ba5790565b613f61336000526004602052604060002054151590565b15610753576000526009602052604060002054806000526007602052604060002091600183015473ffffffffffffffffffffffffffffffffffffffff811680156107295733036106fb5760e01c60ff16613fba8161348a565b15613ff057613bf693613fe4613feb92613fd5368483613b3f565b60208151910120923691613b3f565b9033614fdb565b615104565b50505050565b61400d336000526004602052604060002054151590565b15610753576000526009602052604060002054806000526007602052604060002091600183015473ffffffffffffffffffffffffffffffffffffffff811680156107295733036106fb5760e01c60ff166140668161348a565b613ff057816000526010602052604060002054614084368684613b3f565b60208151910120809114613d7457613fe4613feb92613bf696610d758787614a41565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b906015549073ffffffffffffffffffffffffffffffffffffffff82169073ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff8460a01c1694169382851494858096614213575b613dd857806080957fa7a2a5335a8d1f8f1f7ef8a58332be349ac9fdc25b62512290a91ac4555430a597156141e6575b505067ffffffffffffffff831692828403614194575b50604051938452602084015260408301526060820152a1565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000006015549260a01b169116176015553861417b565b7fffffffffffffffffffffffff000000000000000000000000000000000000000016176015558038614165565b508167ffffffffffffffff841614614135565b9061423082613abb565b61423d604051918261368f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061426b8294613abb565b019060005b82811061427c57505050565b60209061428761397f565b82828501015201614270565b9190811015613af45760051b0190565b604051906142b082613673565b81606061ffff60125460ff8116845260ff8160081c16602085015260ff8160101c16604085015260181c16910152565b916142fe60209263ffffffff929695966040865260408601916140a7565b9416910152565b91929093614314368587613b3f565b6020815191012060005260146020526040600020600181019060ff825460201c161561456957604090600260009173ffffffffffffffffffffffffffffffffffffffff881683520160205220916000146144b8575463ffffffff8381169116811161448e57815460ff8160201c161560001461440c575081547fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000161764010000000017905560405173ffffffffffffffffffffffffffffffffffffffff92909216937f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c8232493839261440792909190846142e0565b0390a2565b919392909163ffffffff168214613dd8577f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c823249484614407937fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000073ffffffffffffffffffffffffffffffffffffffff9754161790556040519485941696846142e0565b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b50809291509392935460ff8160201c1615613d74577f6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d8499373ffffffffffffffffffffffffffffffffffffffff937fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff614407931681557fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000815416905560405193849360208552169560208401916140a7565b6040517f727e650e000000000000000000000000000000000000000000000000000000008152602060048201528061053f60248201898b6140a7565b906040516040810181811067ffffffffffffffff82111761185857604052602060ff82945463ffffffff81168452821c161515910152565b906145ef9063ffffffff933691613b3f565b60208151910120600052601460205273ffffffffffffffffffffffffffffffffffffffff604060002091166000526002810160205261463160406000206145a5565b906020820151614645576001915001541690565b50511690565b3573ffffffffffffffffffffffffffffffffffffffff811681036102ba5790565b909260809273ffffffffffffffffffffffffffffffffffffffff6135e097951683526020830152604082015281606082015201916140a7565b919290928242116148f9576146c7336000526004602052604060002054151590565b6148cb5783600052600660205260ff604060002054166148995760006146eb6136d0565b60405161475981614721602082019486865233604084015246606084015230608084015260e060a08401526101008301906134c3565b8860c08301528960e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261368f565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c5261479a603c8220614794368686613b3f565b906159c3565b909192600483101561486c57826148195750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff16156147e15750505050565b9061053f916040519485947f335d4ce1000000000000000000000000000000000000000000000000000000008652336004870161466c565b8593505060ff61485c6040519586957fd36ab6b90000000000000000000000000000000000000000000000000000000087526060600488015260648701916140a7565b9216602484015260448301520390fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526021600452fd5b837f77a33858000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b7fd9a5f5ca000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b827f502d038700000000000000000000000000000000000000000000000000000000600052336004524260245260445260646000fd5b614946336000526004602052604060002054151590565b156107535760005260096020526040600020548060005260076020526040600020600181015473ffffffffffffffffffffffffffffffffffffffff811680156107295733036106fb5760ff60019160e01c166149a18161348a565b146149af57613bf691614a41565b5050565b9091818310156149ee57816149c88285613ad3565b11156149dd5750905b81039081116105435790565b6149e8915082613ad3565b906149d1565b505050600090565b73ffffffffffffffffffffffffffffffffffffffff600154163303614a1757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b907ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e488689614af99273ffffffffffffffffffffffffffffffffffffffff6001840193614b278554937c01000000000000000000000000000000000000000000000000000000007fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff86161787558060005260106020528360406000205495866000526014602052614b0060406000206040519b8c8092613a19565b038b61368f565b16946002840195614b21604051614b1b81610ee1818c613a19565b82615422565b92615525565b5493541693614b6d604051614b468161121c86868b8b60208601615611565b60405190614b5382613657565b6002825263ffffffff421660208301526040820152614b81565b614b7c60405192839283615657565b0390a3565b906016546801000000000000000081101561185857806001614ba892016016556016613b23565b614d475760016040918451614bbc8161370b565b614bc58161370b565b60ff82549116807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083161783557fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ffffffff00602089015160081b169216171781550192015191825167ffffffffffffffff811161185857614c4c8161168a84546139c6565b6020601f8211600114614ca9578190614c9a939495600092614c9e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0151905038806114cb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110614d2f57509583600195969710614cf8575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080614cee565b9192602060018192868b015181550194019201614cd9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b91929092834211614f625773ffffffffffffffffffffffffffffffffffffffff8316614daf816000526004602052604060002054151590565b15614f3557614dbd8161586f565b90600090614e01614e39614dcf6136d0565b604051928391602083019560018752604084015246606084015230608084015260e060a08401526101008301906134c3565b8a60c08301528660e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261368f565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52614e74603c8220614794368787613b3f565b909192600483101561486c5782614ef45750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1615614ebc575050505050565b9061053f92916040519586957f335d4ce10000000000000000000000000000000000000000000000000000000087526004870161466c565b506040517fd36ab6b90000000000000000000000000000000000000000000000000000000081526060600482015291829160ff61485c606485018a8a6140a7565b7fc2dda3f90000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8373ffffffffffffffffffffffffffffffffffffffff847f3d8a511600000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b60409073ffffffffffffffffffffffffffffffffffffffff6135e0949316815281602082015201906134c3565b90806000526014602052604060002090614ff7600183016145a5565b6020810151156150c95763ffffffff60029151169273ffffffffffffffffffffffffffffffffffffffff85166000520160205261503760406000206145a5565b60208101516150ba575b5073ffffffffffffffffffffffffffffffffffffffff8316600052600f60205260406000209060005260205263ffffffff8060406000205416911610615085575050565b61053f6040519283927f038857ff00000000000000000000000000000000000000000000000000000000845260048401614fae565b5163ffffffff16915038615041565b6040517f727e650e000000000000000000000000000000000000000000000000000000008152602060048201528061053f60248201886134c3565b9061514c929180600052601060205273ffffffffffffffffffffffffffffffffffffffff60406000205492836000526014602052615153604060002060405197888092613a19565b038761368f565b61517f600182019383855416956002840196615179604051614b1b81610ee1818d613a19565b92615441565b8254927fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff8416905554911691604051946151ef866151c38484888860208601615611565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810188528761368f565b6040516151fb81613657565b600181526020810163ffffffff4216815260408201978852601654680100000000000000008110156118585780600161523992016016556016613b23565b919091614d4757600192519061524e8261370b565b6152578261370b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ffffffff0060ff8554941692837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0086161786555160081b1692161717815501955195865167ffffffffffffffff8111611858576152da8161168a84546139c6565b6020601f821160011461535c57908061534b927f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a9798999a600092614c9e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055614b7c60405192839283615657565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169883600052816000209960005b81811061540a5750917f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a9798999a918460019594106153d3575b505050811b019055614b6d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806153c6565b838301518c556001909b019a6020938401930161538c565b9061543b61121c91604051928391602083019586614fae565b51902090565b91929073ffffffffffffffffffffffffffffffffffffffff1680600052600f6020526040600020846000526020526040600020600163ffffffff8254160163ffffffff81116105435763ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000082541617905583600052600b6020526154ce836040600020615967565b50600052600a6020526154e5826040600020615967565b50600052600e6020526154fc816040600020615967565b506000526010602052604060002055565b8054821015613af45760005260206000200190600090565b92919073ffffffffffffffffffffffffffffffffffffffff1680600052600a602052615555846040600020615bfb565b5081600052600b60205261556d846040600020615bfb565b50600052600f60205260406000209060005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff8254160163ffffffff81116105435763ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055600052600e6020526155ff816040600020615bfb565b50600052601060205260006040812055565b9273ffffffffffffffffffffffffffffffffffffffff6135e095936156499386521660208501526080604085015260808401906134c3565b916060818403910152613a19565b909161566e6135e0936040845260408401906134c3565b916020818403910152613a19565b61568681546139c6565b9081615690575050565b81601f600093116001146156a2575055565b818352602083206156be91601f0160051c810190600101613ba4565b8082528160208120915555565b90614b1b907f48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd946157f761576661580573ffffffffffffffffffffffffffffffffffffffff6001860195865460ff8a848316615734610ee1600287019d8e60405192838092613a19565b93849183600052601060205260406000205494859283600052601460205261576d60406000206040519d8e8092613a19565b038d61368f565b60e01c1661577a8161348a565b1561585d575b50505050600052600c60205261579a8a6040600020615bfb565b5082885416600052600d6020526157b58a6040600020615bfb565b5060005260086020526157cc896040600020615bfb565b50805460005260096020526000604081205554955416956040519384936040855260408501906134c3565b908382036020850152613a19565b0390a3600052600760205260066040600020600081556000600182015561582e6002820161567c565b61583a6003820161567c565b6158466004820161567c565b6158526005820161567c565b0161568681546139c6565b61586693615525565b8a828238615780565b806000526005602052604060002054908115806158bc575b61588f575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5080600052600460205260406000205415615887565b80600052600460205260406000205415600014615961576003546801000000000000000081101561185857615948615913826001859401600355600361550d565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600354906000526004602052604060002055600190565b50600090565b60008281526001820160205260409020546159bc578054906801000000000000000082101561185857826159a561591384600180960185558461550d565b905580549260005201602052604060002055600190565b5050600090565b81519190604183036159f4576159ed92506020820151906060604084015193015160001a906159ff565b9192909190565b505060009160029190565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411615a95579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa15615a895760005173ffffffffffffffffffffffffffffffffffffffff811615615a7d5790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b50505060009160039190565b60008181526004602052604090205480156159bc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161054357600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161054357818103615bc1575b5050506003548015615b92577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615b4f81600361550d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615be3615bd261591393600361550d565b90549060031b1c928392600361550d565b90556000526004602052604060002055388080615b16565b9060018201918160005282602052604060002054801515600014615d26577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610543578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161054357818103615cef575b50505080548015615b92577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615cb0828261550d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b615d0f615cff615913938661550d565b90549060031b1c9283928661550d565b905560005283602052604060002055388080615c78565b5050505060009056fea164736f6c634300081a000a",
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

func (_WorkflowRegistry *WorkflowRegistryCaller) CanLinkOwner(opts *bind.CallOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canLinkOwner", validityTimestamp, proof, signature)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanLinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanLinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, validityTimestamp, proof, signature)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMetadataConfig(opts *bind.CallOpts) (WorkflowRegistryMetadataConfig, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMetadataConfig")

	if err != nil {
		return *new(WorkflowRegistryMetadataConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryMetadataConfig)).(*WorkflowRegistryMetadataConfig)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMetadataConfig() (WorkflowRegistryMetadataConfig, error) {
	return _WorkflowRegistry.Contract.GetMetadataConfig(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMetadataConfig() (WorkflowRegistryMetadataConfig, error) {
	return _WorkflowRegistry.Contract.GetMetadataConfig(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadata(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadata", workflowId)

	if err != nil {
		return *new(WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryWorkflowMetadata)).(*WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadata(workflowId [32]byte) (WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadata(&_WorkflowRegistry.CallOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadata(workflowId [32]byte) (WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadata(&_WorkflowRegistry.CallOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByDON(opts *bind.CallOpts, donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByDON", donFamily, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByDON(donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByDON(donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByOwner", owner, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByOwnerAndName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByOwnerAndName", owner, workflowName, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByOwnerAndName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwnerAndName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByOwnerAndName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwnerAndName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseAllByDON(opts *bind.TransactOpts, donFamily [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseAllByDON", donFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseAllByDON(donFamily [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByDON(&_WorkflowRegistry.TransactOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseAllByDON(donFamily [32]byte) (*types.Transaction, error) {
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "batchActivateWorkflows", workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistrySession) BatchActivateWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) BatchActivateWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetMetadataConfig(opts *bind.TransactOpts, config WorkflowRegistryMetadataConfig) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setMetadataConfig", config)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetMetadataConfig(config WorkflowRegistryMetadataConfig) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetMetadataConfig(&_WorkflowRegistry.TransactOpts, config)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetMetadataConfig(config WorkflowRegistryMetadataConfig) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetMetadataConfig(&_WorkflowRegistry.TransactOpts, config)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpdateWorkflowDONLabel(opts *bind.TransactOpts, workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "updateWorkflowDONLabel", workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpdateWorkflowDONLabel(workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateWorkflowDONLabel(&_WorkflowRegistry.TransactOpts, workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpdateWorkflowDONLabel(workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateWorkflowDONLabel(&_WorkflowRegistry.TransactOpts, workflowId, newDonFamily)
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
	MaxWorkflowNameLength uint8
	MaxWorkflowTagLength  uint8
	MaxUrlLength          uint8
	MaxAttributesLength   uint16
	Raw                   types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterMetadataConfigUpdated(opts *bind.FilterOpts, maxWorkflowNameLength []uint8, maxWorkflowTagLength []uint8, maxUrlLength []uint8) (*WorkflowRegistryMetadataConfigUpdatedIterator, error) {

	var maxWorkflowNameLengthRule []interface{}
	for _, maxWorkflowNameLengthItem := range maxWorkflowNameLength {
		maxWorkflowNameLengthRule = append(maxWorkflowNameLengthRule, maxWorkflowNameLengthItem)
	}
	var maxWorkflowTagLengthRule []interface{}
	for _, maxWorkflowTagLengthItem := range maxWorkflowTagLength {
		maxWorkflowTagLengthRule = append(maxWorkflowTagLengthRule, maxWorkflowTagLengthItem)
	}
	var maxUrlLengthRule []interface{}
	for _, maxUrlLengthItem := range maxUrlLength {
		maxUrlLengthRule = append(maxUrlLengthRule, maxUrlLengthItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "MetadataConfigUpdated", maxWorkflowNameLengthRule, maxWorkflowTagLengthRule, maxUrlLengthRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryMetadataConfigUpdatedIterator{contract: _WorkflowRegistry.contract, event: "MetadataConfigUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchMetadataConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryMetadataConfigUpdated, maxWorkflowNameLength []uint8, maxWorkflowTagLength []uint8, maxUrlLength []uint8) (event.Subscription, error) {

	var maxWorkflowNameLengthRule []interface{}
	for _, maxWorkflowNameLengthItem := range maxWorkflowNameLength {
		maxWorkflowNameLengthRule = append(maxWorkflowNameLengthRule, maxWorkflowNameLengthItem)
	}
	var maxWorkflowTagLengthRule []interface{}
	for _, maxWorkflowTagLengthItem := range maxWorkflowTagLength {
		maxWorkflowTagLengthRule = append(maxWorkflowTagLengthRule, maxWorkflowTagLengthItem)
	}
	var maxUrlLengthRule []interface{}
	for _, maxUrlLengthItem := range maxUrlLength {
		maxUrlLengthRule = append(maxUrlLengthRule, maxUrlLengthItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "MetadataConfigUpdated", maxWorkflowNameLengthRule, maxWorkflowTagLengthRule, maxUrlLengthRule)
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
	CanLinkOwner(opts *bind.CallOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) error

	CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error

	GetDONRegistry(opts *bind.CallOpts) (common.Address, uint64, error)

	GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (uint32, error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donFamily string) (uint32, error)

	GetMetadataConfig(opts *bind.CallOpts) (WorkflowRegistryMetadataConfig, error)

	GetWorkflowMetadata(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByDON(opts *bind.CallOpts, donFamily [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByOwnerAndName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowOwnerConfig(opts *bind.CallOpts, owner common.Address) ([]byte, error)

	IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error)

	IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte, donFamily string) (*types.Transaction, error)

	AdminBatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	AdminPauseAllByDON(opts *bind.TransactOpts, donFamily [32]byte) (*types.Transaction, error)

	AdminPauseAllByOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error)

	AdminPauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	BatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	DeleteWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error)

	PauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	SetDONLimit(opts *bind.TransactOpts, donFamily string, limit uint32, enabled bool) (*types.Transaction, error)

	SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector uint64) (*types.Transaction, error)

	SetMetadataConfig(opts *bind.TransactOpts, config WorkflowRegistryMetadataConfig) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, limit uint32, enabled bool) (*types.Transaction, error)

	SetWorkflowOwnerConfig(opts *bind.TransactOpts, owner common.Address, config []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	UpdateWorkflowDONLabel(opts *bind.TransactOpts, workflowId [32]byte, newDonFamily string) (*types.Transaction, error)

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

	FilterMetadataConfigUpdated(opts *bind.FilterOpts, maxWorkflowNameLength []uint8, maxWorkflowTagLength []uint8, maxUrlLength []uint8) (*WorkflowRegistryMetadataConfigUpdatedIterator, error)

	WatchMetadataConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryMetadataConfigUpdated, maxWorkflowNameLength []uint8, maxWorkflowTagLength []uint8, maxUrlLength []uint8) (event.Subscription, error)

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
