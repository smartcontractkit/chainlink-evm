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

type WorkflowRegistryWorkflowMetadata struct {
	WorkflowId   [32]byte
	Owner        common.Address
	CreatedAt    uint64
	Status       uint8
	DonLabel     [32]byte
	WorkflowName string
	BinaryUrl    string
	ConfigUrl    string
	Tag          string
	Attributes   []byte
}

var WorkflowRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByUser\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadata\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwnerName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUnlinkWithActiveWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerDONExceeded\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowAlreadyInDesiredStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561502490816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c8063034a6ed5146131755780630987294c1461310657806317e0edfc14612fc5578063181f5a7714612f775780631c387fd614612eb55780631c71682c14612c9257806323f381bb14612c3f5780633c17181b14612bc557806347d1ed8314612b9357806352453a4414612a82578063695e1340146129605780636ee80b44146127ff5780636f3517711461275257806379ba509714612669578063809b23cc14612619578063810967671461243a5780638da5cb5b146123e85780639c5b2304146123105780639f2b2cbd1461108f578063a0b8a4fe14611053578063a7d0185814610f8e578063b87a019414610dc0578063cabb9e7a14610d56578063cb647e8214610cc0578063d8e056de14610bd2578063d8e4a72414610a08578063dc10196914610974578063dfcb0b311461095a578063e690f33214610906578063e9df655314610536578063eec860b714610459578063f23e37ff1461027c5763f2fde38b1461018757600080fd5b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775773ffffffffffffffffffffffffffffffffffffffff6101d3613466565b6101db613de9565b1633811461024d57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102775760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610277576102b3613466565b6024359060443563ffffffff8116809103610277576064358015158103610277576102dc613de9565b156103ff5782600052601060205263ffffffff6040600020541681116103d55760207f2a666be2166cbc432b468f0c0273b4c7009e2fc0168952f8adc3c3b0aecf8c619160405161032c816134f6565b81815273ffffffffffffffffffffffffffffffffffffffff838201956001875216948560005260118452604060002087600052845263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff00000000835492861b169116179055604051908152a3005b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b5073ffffffffffffffffffffffffffffffffffffffff16806000526011602052604060002082600052602052600060408120557fb991c2ee405ccdec48f880c2f2edd1b4161ceefb047f499c9a024502779a039a600080a3005b346102775760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027757610490613466565b6024359061049c613de9565b73ffffffffffffffffffffffffffffffffffffffff6012541682602073ffffffffffffffffffffffffffffffffffffffff604051946104da866134f6565b16938481520152817fffffffffffffffffffffffff00000000000000000000000000000000000000006012541617601255826013557fc60b6bc6904a32fe864ae9ad697fd792f40e1d9e9fee5da0a10d0fb26b7380d6600080a4005b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043567ffffffffffffffff8111610277576105859036906004016136c3565b80156108dc57610594336142c4565b61059d8161388d565b6105aa6040519182613512565b8181527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06105d78361388d565b013660208301376105e78261388d565b906105f56040519283613512565b8282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06106228461388d565b01366020840137600090815b8481106107d1575060005b8281106106a757858560005b81811061064e57005b8061065c6001928486613a3f565b35600052600960205260406000205461067581336145c9565b8360ff8183015460e01c1661068981613301565b14610697575b505001610645565b6106a091614648565b848061068f565b6106b181836138b2565b5163ffffffff6106c183876138b2565b511681600052600b6020526106db816040600020546138a5565b82600052601060205263ffffffff60406000205416106107a35733600052600f60205260406000208260005260205263ffffffff604060002054160163ffffffff81116107745763ffffffff8061073284336138f5565b169116116107435750600101610639565b7f6721d3de000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b507fcad00a8d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6107dc818688613a3f565b3560005260096020526107f4604060002054336145c9565b60ff600182015460e01c1661080881613301565b156108d357600201546000805b85811061087d575b501561082f575b506001905b0161062e565b61083c84849593956138b2565b52600161084982866138b2565b527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107745760018091019290610824565b8261088882876138b2565b511461089657600101610815565b905063ffffffff6108a782886138b2565b51169063ffffffff82146107745760016108c663ffffffff92896138b2565b920116905260018861081d565b50600190610829565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775761093e336142c4565b60043560005260096020526109586040600020543361459c565b005b346102775761095861096b366136f4565b92919091613ae8565b34610277576001610990610987366136f4565b91809493613ae8565b336000526005602052806040600020556109a933614ec0565b508060005260066020526040600020827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346102775760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043567ffffffffffffffff811161027757610a579036906004016136c3565b6024359182151580930361027757610a6d613de9565b60ff831660005b838110610b1857505060405191806040840160408552526060830191906000905b808210610acb577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b90919283359073ffffffffffffffffffffffffffffffffffffffff82168203610277576020809173ffffffffffffffffffffffffffffffffffffffff600194168152019401920190610a95565b73ffffffffffffffffffffffffffffffffffffffff610b40610b3b838787613a3f565b613a4f565b1615610ba8578073ffffffffffffffffffffffffffffffffffffffff610b6c610b3b6001948888613a3f565b1660005260026020526040600020837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905501610a74565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027757610c09613de9565b600435600052600b60205260406000205b80548015610958577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161077457610c5690826147d0565b90549060031b1c6000526007602052604060002054610c73613de9565b6000526009602052604060002054806000526007602052604060002060ff600182015460e01c16610ca381613301565b15610cb0575b5050610c1a565b610cb991613e6b565b8180610ca9565b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043567ffffffffffffffff811161027757610d0f9036906004016136c3565b80156108dc57610d1e336142c4565b60005b818110610d2a57005b80610d386001928486613a3f565b356000526009602052610d506040600020543361459c565b01610d21565b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775773ffffffffffffffffffffffffffffffffffffffff610da2613466565b166000526002602052602060ff604060002054166040519015158152f35b346102775760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775773ffffffffffffffffffffffffffffffffffffffff610e0c613466565b1680600052600d602052610e2a604435602435604060002054613d8e565b90610e34826139d2565b9260005b838110610e515760405180610e4d8782613645565b0390f35b60019082600052600d602052610e756040600020610e6f83876138a5565b906147d0565b90549060031b1c600052600760205260406000206007610f6c60405192610e9b846134bd565b8054845260ff8682015473ffffffffffffffffffffffffffffffffffffffff8116602087015267ffffffffffffffff8160a01c16604087015260e01c16610ee181613301565b606085015260028101546080850152604051610f0b81610f0481600386016137eb565b0382613512565b60a0850152604051610f2481610f0481600486016137eb565b60c0850152604051610f3d81610f0481600586016137eb565b60e0850152604051610f5681610f0481600686016137eb565b610100850152610f0460405180948193016137eb565b610120820152610f7c82886138b2565b52610f8781876138b2565b5001610e38565b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043567ffffffffffffffff811161027757610fdd9036906004016136c3565b80156108dc5760005b818110610fef57005b80610ffd6001928486613a3f565b35611006613de9565b6000526009602052604060002054806000526007602052604060002060ff8482015460e01c1661103581613301565b15611043575b505001610fe6565b61104c91613e6b565b848061103b565b346102775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610277576020600354604051908152f35b34610277576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043567ffffffffffffffff8111610277576110df903690600401613598565b60243567ffffffffffffffff8111610277576110ff903690600401613598565b90600260643510156102775760a43567ffffffffffffffff81116102775761112b903690600401613598565b909160c43567ffffffffffffffff81116102775761114d903690600401613598565b94909160e43567ffffffffffffffff811161027757611170903690600401613598565b92909361010435978815988915036102775761118d368b8d61396d565b6111f061121c61119e36878961396d565b6040519283916111c0602084019633885260606040860152608085019061333a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084830301606085015261333a565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282613512565b5190209788600052600760205273ffffffffffffffffffffffffffffffffffffffff6001604060002001541615600014611e9157611259336142c4565b8a15611e675760408b11611e34578315611e0a5760208411611dd757611280604435614ac7565b61128a8289614b10565b61129386614b55565b6112a76112a1368d8f61396d565b33614283565b996112b3606435613301565b60643515611dc6575b611bca575b604051976112ce896134bd565b60443589523360208a015267ffffffffffffffff421660408a01526112f4606435613301565b60643560608a015260843560808a015261130f368d8f61396d565b60a08a0152369061131f9261396d565b60c0880152369061132f9261396d565b60e0860152369061133f9261396d565b61010084015236906113509261396d565b61012082015281600052600760205260406000208151815560018101602083015173ffffffffffffffffffffffffffffffffffffffff1681547fffffffffffffffffffffffff000000000000000000000000000000000000000016178155604083015181549060a01b7bffffffffffffffff000000000000000000000000000000000000000016907fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff16178155606083015161140b81613301565b61141481613301565b81549060e01b7cff0000000000000000000000000000000000000000000000000000000016907fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff16179055608082015160028201556003810160a083015180519067ffffffffffffffff82116118f257611498826114928554613798565b8561433d565b602090601f8311600114611b27576114e69291600091836119d6575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6004810160c083015180519067ffffffffffffffff82116118f257611512826114928554613798565b602090601f8311600114611a845761155f9291600091836119d65750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6005810160e083015180519067ffffffffffffffff82116118f25761158b826114928554613798565b602090601f83116001146119e1576115d89291600091836119d65750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6006810161010083015180519067ffffffffffffffff82116118f257611605826114928554613798565b602090601f831160011461192c578261012095936007959361165a936000926119215750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b0191015180519067ffffffffffffffff82116118f257611681826114928554613798565b602090601f831160011461184f576116ce9291600091836118445750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b8160005260086020526116e8816040600020614f20565b50604435600052600960205280604060002055608435600052600c602052611714816040600020614f20565b5033600052600d60205261172c816040600020614f20565b50611738606435613301565b6064351561179f575b50507f03bdc9e1a7b65f76fa1e989e76c28e6ac9f556262cfed31a11a9461e77ad2ff56040516084358152611777606435613301565b6064356020820152606060408201528061179a3395604435956060840191613a70565b0390a3005b61183c91608435600052600b6020526117bc826040600020614f20565b5033600052600a6020526117d4826040600020614f20565b5033600052600f6020526040600020608435600052602052604060002063ffffffff61180281835416613d76565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055600052600e6020526040600020614f20565b508280611741565b0151905088806114b4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b8181106118da57509084600195949392106118a3575b505050811b0190556116d1565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055878080611896565b92936020600181928786015181550195019301611880565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b015190508c806114b4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b8181106119be575092600192859261012098966007989610611987575b505050811b01905561165d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558b808061197a565b9293602060018192878601518155019501930161195d565b015190508a806114b4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611a6c5750908460019594939210611a35575b505050811b0190556115db565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611a28565b92936020600181928786015181550195019301611a12565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611b0f5750908460019594939210611ad8575b505050811b019055611562565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611acb565b92936020600181928786015181550195019301611ab5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611bb25750908460019594939210611b7b575b505050811b0190556114e9565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611b6e565b92936020600181928786015181550195019301611b58565b989a979593918b60009b989694929b52600e60205260406000209a5b8b546001811115611db4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161077457611c25908d6147d0565b90549060031b1c8060005260076020528d604060002091611d53600184019173ffffffffffffffffffffffffffffffffffffffff83547c01000000000000000000000000000000000000000000000000000000007fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff821617855516600052600f60205260406000209360028601948554600052602052604060002063ffffffff611cd181835416613e34565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000082541617905573ffffffffffffffffffffffffffffffffffffffff845416600052600a602052611d27826040600020614d8c565b508454600052600b602052611d40826040600020614d8c565b50600052600e6020526040600020614d8c565b507f52613f4f9f39f66629197514b47e67cab5e5c40b83fad8fccec907d71a8a3978611dac73ffffffffffffffffffffffffffffffffffffffff85549354169354946040519182916020835260036020840191016137eb565b0390a4611be6565b50919395979a50919395979a986112c1565b611dd260843533614a1e565b6112bc565b837f436f975400000000000000000000000000000000000000000000000000000000600052600452602060245260446000fd5b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b8a7f36a7c50300000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b925092509293949650611ea3336142c4565b611eae604435614ac7565b611eb88288614b10565b611ec184614b55565b611ecb86336145c9565b958654600052600960205260006040812055604435600052600960205260406000205585549660443587556004870191611f20604051611f0f81610f0481886137eb565b611f1a36858561396d565b90614b94565b156121f0575b5050506005850191611f42604051611f0f81610f0481886137eb565b156120db575b505050600783019167ffffffffffffffff82116118f257611f6d826114928554613798565b600090601f831160011461201657827f67057592e7b546efc47d8093ece5f7d8837b769a512710d66a892555875bc05e959360029593611fe09360009261200b5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b015460405190815260406020820152806120063396604435966040840191613a70565b0390a4005b013590508a806114b4565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b8181106120c357509260019285927f67057592e7b546efc47d8093ece5f7d8837b769a512710d66a892555875bc05e9896600298961061208b575b505050811b019055611fe3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c1991013516905589808061207e565b91936020600181928787013581550195019201612043565b67ffffffffffffffff82116118f2576120f8826114928554613798565b600090601f83116001146121505761214592916000918361200b5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b868080611f48565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b8181106121d857509084600195949392106121a0575b505050811b019055612148565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055898080612193565b9193602060018192878701358155019501920161217d565b67ffffffffffffffff82116118f25761220d826114928554613798565b600090601f83116001146122705761225a9291600091836122655750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b888080611f26565b013590508c806114b4565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b8181106122f857509084600195949392106122c0575b505050811b01905561225d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558b80806122b3565b9193602060018192878701358155019501920161229d565b346102775760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043580600052600c60205261235f604435602435604060002054613d8e565b90612369826139d2565b9260005b8381106123825760405180610e4d8782613645565b60019082600052600c6020526123a06040600020610e6f83876138a5565b90549060031b1c6000526007602052604060002060076123c660405192610e9b846134bd565b6101208201526123d682886138b2565b526123e181876138b2565b500161236d565b346102775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102775760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043560243563ffffffff811680910361027757604435801515810361027757612491613de9565b156125925760207f3800848414d9fa9c555c8b830833c55ea97e0743a1de02f19a0dd76e9578ebab916040516124c6816134f6565b81815282810160018152856000526010845263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff00000000835492861b169116179055612589604051858482015282604082015260408152612563606082613512565b60405190612570826134da565b6000825263ffffffff4216858301526040820152614384565b604051908152a2005b50806000526010602052600060408120556125ec60405182602082015260006040820152604081526125c5606082613512565b604051906125d2826134da565b6000825263ffffffff421660208301526040820152614384565b7f3800848414d9fa9c555c8b830833c55ea97e0743a1de02f19a0dd76e9578ebab602060405160008152a2005b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610277576004356000526010602052602063ffffffff60406000205416604051908152f35b346102775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760005473ffffffffffffffffffffffffffffffffffffffff81163303612728577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775761278a336142c4565b60043560005260096020526040600020546127a581336145c9565b60ff600182015460e01c166127b981613301565b156127d557610958916127d0600283015433614a1e565b614648565b7f6f861db10000000000000000000000000000000000000000000000000000000060005260046000fd5b346102775761282b612810366135c6565b93949161281c8561358e565b84159384612952575b8661408b565b6128348261358e565b1561289b575b600073ffffffffffffffffffffffffffffffffffffffff831661285c816149bb565b90808352600560205282604081205561287481614bfd565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b73ffffffffffffffffffffffffffffffffffffffff8216600052600a60205260406000205b80548015612935577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111610774576128fc90826147d0565b90549060031b1c61290d81856145c9565b6129168461358e565b6002840361292c5761292791613e6b565b6128c0565b61292791614870565b50505073ffffffffffffffffffffffffffffffffffffffff61283a565b61295b87614037565b612825565b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775760043560005260096020526109586040600020546129ad81336145c9565b9060ff600183015460e01c166129c281613301565b6148705760028201546129e26040516112a181610f0481600389016137eb565b9033600052600a6020526129fa836040600020614d8c565b5080600052600b602052612a12836040600020614d8c565b5033600052600f602052604060002090600052602052604060002063ffffffff612a3e81835416613e34565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055600052600e602052612a7c816040600020614d8c565b50614870565b346102775760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027757612ab9613466565b6024359067ffffffffffffffff821161027757612ae7612ae0612aed933690600401613598565b369161396d565b90614283565b806000526008602052612b0a606435604435604060002054613d8e565b90612b14826139d2565b9260005b838110612b2d5760405180610e4d8782613645565b600190826000526008602052612b4b6040600020610e6f83876138a5565b90549060031b1c600052600760205260406000206007612b7160405192610e9b846134bd565b610120820152612b8182886138b2565b52612b8c81876138b2565b5001612b18565b3461027757610958612ba4366135c6565b909391929190612bb38161358e565b61408b57612bc081614037565b61408b565b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027757612bfc613de9565b6004356000526009602052604060002054806000526007602052604060002060ff600182015460e01c16612c2f81613301565b15612c3657005b61095891613e6b565b346102775760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610277576020612c84612c7b613466565b602435906138f5565b63ffffffff60405191168152f35b3461027757612cac612ca336613489565b90601454613d8e565b612cb58161388d565b91612cc36040519384613512565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612cf08361388d565b0160005b818110612e89575050601454919060005b828110612dba57846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210612d4557505050500390f35b91936020612daa827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652606060408a518051612d8a8161358e565b845263ffffffff868201511686850152015191816040820152019061333a565b9601920192018594939192612d36565b612dc481836138a5565b600085821015612e5c5760149052604051600192918390612e3b90821b7fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec01612e0c846134da565b63ffffffff815460ff8116612e208161358e565b865260081c166020850152610f0460405180948193016137eb565b6040820152612e4a82886138b2565b52612e5581876138b2565b5001612d05565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b602090604051612e98816134da565b600081526000838201526060604082015282828801015201612cf4565b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102775773ffffffffffffffffffffffffffffffffffffffff612f01613466565b612f09613de9565b16600052600a60205260406000205b80548015610958577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161077457612f57612f7291836147d0565b90549060031b1c806000526007602052604060002090613e6b565b612f18565b346102775760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027757610e4d612fb1613553565b60405191829160208352602083019061333a565b3461027757612fdf612fd636613489565b90600354613d8e565b612fe88161388d565b91612ff66040519384613512565b8183526130028261388d565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020850193013684376003549160005b8281106130965784866040519182916020830190602084525180915260408301919060005b818110613067575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101613059565b6130a081836138a5565b600085821015612e5c57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0154600081905260056020526001919073ffffffffffffffffffffffffffffffffffffffff166130ff82896138b2565b5201613034565b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027757602061316b73ffffffffffffffffffffffffffffffffffffffff613157613466565b166000526004602052604060002054151590565b6040519015158152f35b346102775760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610277576131ac613749565b50600435600052600960205260406000205480156132d7576000526007602052604060002073ffffffffffffffffffffffffffffffffffffffff604051916131f3836134bd565b8054835260076132b260018301549260ff6020870194868116865267ffffffffffffffff8160a01c16604089015260e01c1661322e81613301565b60608701526002810154608087015260405161325181610f0481600386016137eb565b60a087015260405161326a81610f0481600486016137eb565b60c087015260405161328381610f0481600586016137eb565b60e087015260405161329c81610f0481600686016137eb565b610100870152610f0460405180948193016137eb565b6101208401525116156132d757610e4d90604051918291602083526020830190613399565b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b6002111561330b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b919082519283825260005b8481106133845750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613345565b613463918151815273ffffffffffffffffffffffffffffffffffffffff602083015116602082015267ffffffffffffffff604083015116604082015260608201516133e381613301565b60608201526080820151608082015261012061345161343d61342b61341960a087015161014060a088015261014087019061333a565b60c087015186820360c088015261333a565b60e086015185820360e087015261333a565b61010085015184820361010086015261333a565b9201519061012081840391015261333a565b90565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361027757565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112610277576004359060243590565b610140810190811067ffffffffffffffff8211176118f257604052565b6060810190811067ffffffffffffffff8211176118f257604052565b6040810190811067ffffffffffffffff8211176118f257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176118f257604052565b60405190613562604083613512565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b6003111561330b57565b9181601f840112156102775782359167ffffffffffffffff8311610277576020838186019501011161027757565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102775760043573ffffffffffffffffffffffffffffffffffffffff811681036102775791602435916044359067ffffffffffffffff82116102775761363491600401613598565b909160643560038110156102775790565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061367857505050505090565b90919293946020806136b4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951613399565b97019301930191939290613669565b9181601f840112156102775782359167ffffffffffffffff8311610277576020808501948460051b01011161027757565b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102775760043591602435916044359067ffffffffffffffff82116102775761374591600401613598565b9091565b60405190613756826134bd565b606061012083600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201528260e0820152826101008201520152565b90600182811c921680156137e1575b60208310146137b257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916137a7565b600092918154916137fb83613798565b8083529260018116908115613851575060011461381757505050565b60009081526020812093945091925b838310613837575060209250010190565b600181602092949394548385870101520191019190613826565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b67ffffffffffffffff81116118f25760051b60200190565b9190820180921161077457565b80518210156138c65760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff1660005260116020526040600020816000526020526040600020602060405191613933836134f6565b549160ff63ffffffff841693848352831c16151591829101526139685750600052601060205263ffffffff6040600020541690565b905090565b92919267ffffffffffffffff82116118f257604051916139b5601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184613512565b829481845281830111610277578281602093846000960137010152565b906139dc8261388d565b6139e96040519182613512565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613a17829461388d565b019060005b828110613a2857505050565b602090613a33613749565b82828501015201613a1c565b91908110156138c65760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036102775790565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b909260809273ffffffffffffffffffffffffffffffffffffffff6134639795168352602083015260408201528160608201520191613a70565b91929092824211613d4057613b0a336000526004602052604060002054151590565b613d125783600052600660205260ff60406000205416613ce0576000613b2e613553565b604051613b9c81613b64602082019486865233604084015246606084015230608084015260e060a084015261010083019061333a565b8860c08301528960e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282613512565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613bdd603c8220613bd736868661396d565b90614bc1565b9091926004831015613cb35782613c605750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1615613c245750505050565b90613c5c916040519485947f335d4ce10000000000000000000000000000000000000000000000000000000086523360048701613aaf565b0390fd5b8593505060ff613ca36040519586957fd36ab6b9000000000000000000000000000000000000000000000000000000008752606060048801526064870191613a70565b9216602484015260448301520390fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526021600452fd5b837f77a33858000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b7fd9a5f5ca000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b827f502d038700000000000000000000000000000000000000000000000000000000600052336004524260245260445260646000fd5b63ffffffff60019116019063ffffffff821161077457565b929183821015613de25780158015613dd8575b613dd0575b613db090826138a5565b92808411613dc8575b50808303928311610774579190565b925038613db9565b506064613da6565b5060648111613da1565b5050600090565b73ffffffffffffffffffffffffffffffffffffffff600154163303613e0a57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b63ffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9116019063ffffffff821161077457565b906001810173ffffffffffffffffffffffffffffffffffffffff815416600052600f6020527f52613f4f9f39f66629197514b47e67cab5e5c40b83fad8fccec907d71a8a397861403273ffffffffffffffffffffffffffffffffffffffff60406000209460028101958654600052602052604060002063ffffffff613ef281835416613e34565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008254161790558554600052600b602052613f32876040600020614d8c565b5081855416600052600a602052613f4d876040600020614d8c565b50613f8f8286541697613f7d613f71600385019a613f788c604051938480926137eb565b0383613512565b614283565b600052600e6020526040600020614d8c565b508454947c01000000000000000000000000000000000000000000000000000000007fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff871617905554931693549461401e604051613ff7816111f0858b8b8b602086016147e8565b60405190614004826134da565b6002825263ffffffff421660208301526040820152614384565b6040519182916020835260208301906137eb565b0390a4565b73ffffffffffffffffffffffffffffffffffffffff16600052600a60205260406000205461406157565b7f61bc2e180000000000000000000000000000000000000000000000000000000060005260046000fd5b919290928342116142375761409f836142c4565b73ffffffffffffffffffffffffffffffffffffffff83166140bf816149bb565b9060009061410361413b6140d1613553565b604051928391602083019560018752604084015246606084015230608084015260e060a084015261010083019061333a565b8a60c08301528660e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282613512565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52614176603c8220613bd736878761396d565b9091926004831015613cb357826141f65750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff16156141be575050505050565b90613c5c92916040519586957f335d4ce100000000000000000000000000000000000000000000000000000000875260048701613aaf565b506040517fd36ab6b90000000000000000000000000000000000000000000000000000000081526060600482015291829160ff613ca3606485018a8a613a70565b8373ffffffffffffffffffffffffffffffffffffffff847f3d8a511600000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b906142be6111f09160405192839173ffffffffffffffffffffffffffffffffffffffff6020840196168652604080840152606083019061333a565b51902090565b73ffffffffffffffffffffffffffffffffffffffff166142f1816000526004602052604060002054151590565b156142f95750565b7fc2dda3f90000000000000000000000000000000000000000000000000000000060005260045260246000fd5b818110614331575050565b60008155600101614326565b9190601f811161434c57505050565b614378926000526020600020906020601f840160051c8301931061437a575b601f0160051c0190614326565b565b909150819061436b565b90601454680100000000000000008110156118f257600181016014556000601454821015612e5c57601490526000929060011b7fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec01600160409183516143e98161358e565b6143f28161358e565b60ff82549116807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083161783557fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ffffffff00602088015160081b16921617178155019101519283519067ffffffffffffffff821161456f5761447a826114928554613798565b602090601f83116001146144d557906144c6939495836144ca5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0151905038806114b4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316848352818320925b81811061455757509583600195969710614520575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080614516565b9192602060018192868b015181550194019201614501565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b816145a6916145c9565b600160ff8183015460e01c166145bb81613301565b146127d55761437891613e6b565b90600052600760205260406000209073ffffffffffffffffffffffffffffffffffffffff6001830154169081156132d75773ffffffffffffffffffffffffffffffffffffffff1680910361461b575090565b7f31ee6dc70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b906001810173ffffffffffffffffffffffffffffffffffffffff815416600052600f6020527f3131ee488b1b459e313b249599f9ebc23dfc5dbfcbd082dd503d171c7436fe8261403273ffffffffffffffffffffffffffffffffffffffff60406000209460028101958654600052602052604060002063ffffffff6146cf81835416613d76565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008254161790558554600052600b60205261470f876040600020614f20565b5081855416600052600a60205261472a876040600020614f20565b50614760828654169761474e613f71600385019a613f788c604051938480926137eb565b600052600e6020526040600020614f20565b508454947fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff8616905554931693549461401e6040516147a9816111f0858b8b8b602086016147e8565b604051906147b6826134da565b6001825263ffffffff421660208301526040820152614384565b80548210156138c65760005260206000200190600090565b909273ffffffffffffffffffffffffffffffffffffffff60809361346396958452166020830152604082015281606082015201906137eb565b61482b8154613798565b9081614835575050565b81601f60009311600114614847575055565b8183526020832061486391601f0160051c810190600101614326565b8082528160208120915555565b90600281018054600052600c60205261488d836040600020614d8c565b506001820173ffffffffffffffffffffffffffffffffffffffff815416600052600d6020526148c0846040600020614d8c565b507fbe9d8e7e7313f6e3020af5157a64045cca8bcadce6e6497ca12fe56dbb4fdeb961495c73ffffffffffffffffffffffffffffffffffffffff8084541695614919600382019760405190613f7882613f71818d6137eb565b600052600860205261492f886040600020614d8c565b508054600052600960205260006040812055549354169354946040519182916020835260208301906137eb565b0390a460005260076020526007604060002060008155600060018201556000600282015561498c60038201614821565b61499860048201614821565b6149a460058201614821565b6149b060068201614821565b0161482b8154613798565b80600052600560205260406000205490811580614a08575b6149db575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b50806000526004602052604060002054156149d3565b81600052600b60205260406000205482600052601060205263ffffffff6040600020541611156107a35773ffffffffffffffffffffffffffffffffffffffff81169081600052600f60205260406000208360005260205263ffffffff614a8c848260406000205416936138f5565b161115614a97575050565b7f6721d3de0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b8015610ba857806000526009602052604060002054614ae35750565b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60c88111614b24575060c88111614b245750565b7ecd56a80000000000000000000000000000000000000000000000000000000060005260045260c860245260446000fd5b6102008111614b615750565b7f354f25140000000000000000000000000000000000000000000000000000000060005260045261020060245260446000fd5b9081518151908181149384614bab575b5050505090565b6020929394508201209201201438808080614ba4565b8151919060418303614bf257614beb92506020820151906060604084015193015160001a90614f75565b9192909190565b505060009160029190565b6000818152600460205260409020548015613de2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161077457600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161077457818103614d1d575b5050506003548015614cee577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614cab8160036147d0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b614d74614d2e614d3f9360036147d0565b90549060031b1c92839260036147d0565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526004602052604060002055388080614c72565b9060018201918160005282602052604060002054801515600014614eb7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610774578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161077457818103614e80575b50505080548015614cee577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190614e4182826147d0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b614ea0614e90614d3f93866147d0565b90549060031b1c928392866147d0565b905560005283602052604060002055388080614e09565b50505050600090565b80600052600460205260406000205415600014614f1a57600354680100000000000000008110156118f257614f01614d3f82600185940160035560036147d0565b9055600354906000526004602052604060002055600190565b50600090565b6000828152600182016020526040902054613de257805490680100000000000000008210156118f25782614f5e614d3f8460018096018555846147d0565b905580549260005201602052604060002055600190565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161500b579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa15614fff5760005173ffffffffffffffffffffffffffffffffffffffff811615614ff35790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b5050506000916003919056fea164736f6c634300081a000a",
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel [32]byte) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerDON", donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerDON(donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerDON(donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel [32]byte) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerUserDON", user, donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donLabel)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByDON(opts *bind.CallOpts, donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByDON", donLabel, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByDON(donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByDON(&_WorkflowRegistry.CallOpts, donLabel, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByDON(donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByDON(&_WorkflowRegistry.CallOpts, donLabel, start, limit)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByOwnerName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByOwnerName", owner, workflowName, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByOwnerName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwnerName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByOwnerName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwnerName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "activateWorkflow", workflowId)
}

func (_WorkflowRegistry *WorkflowRegistrySession) ActivateWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) ActivateWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseAllByDON(opts *bind.TransactOpts, donLabel [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseAllByDON", donLabel)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseAllByDON(donLabel [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByDON(&_WorkflowRegistry.TransactOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseAllByDON(donLabel [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByDON(&_WorkflowRegistry.TransactOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseAllByUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseAllByUser", user)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseAllByUser(user common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByUser(&_WorkflowRegistry.TransactOpts, user)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseAllByUser(user common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByUser(&_WorkflowRegistry.TransactOpts, user)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) PauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "pauseWorkflows", workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistrySession) PauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) PauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONLimit(opts *bind.TransactOpts, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONLimit", donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONLimit(donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONLimit(donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONRegistry", registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONRegistry(registry common.Address, chainSelector *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONRegistry(registry common.Address, chainSelector *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setUserDONOverride", user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetUserDONOverride(user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetUserDONOverride(user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donLabel, limit, enabled)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donLabel [32]byte, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "upsertWorkflow", workflowName, tag, workflowId, status, donLabel, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donLabel [32]byte, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowId, status, donLabel, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donLabel [32]byte, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowId, status, donLabel, binaryUrl, configUrl, attributes, keepAlive)
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
	DonLabel [32]byte
	Limit    uint32
	Raw      types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterDONLimitSet(opts *bind.FilterOpts, donLabel [][32]byte) (*WorkflowRegistryDONLimitSetIterator, error) {

	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "DONLimitSet", donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryDONLimitSetIterator{contract: _WorkflowRegistry.contract, event: "DONLimitSet", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet, donLabel [][32]byte) (event.Subscription, error) {

	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "DONLimitSet", donLabelRule)
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
	OldAddr       common.Address
	NewAddr       common.Address
	ChainSelector *big.Int
	Raw           types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterDONRegistryUpdated(opts *bind.FilterOpts, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (*WorkflowRegistryDONRegistryUpdatedIterator, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}
	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "DONRegistryUpdated", oldAddrRule, newAddrRule, chainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryDONRegistryUpdatedIterator{contract: _WorkflowRegistry.contract, event: "DONRegistryUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchDONRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONRegistryUpdated, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (event.Subscription, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}
	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "DONRegistryUpdated", oldAddrRule, newAddrRule, chainSelectorRule)
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
	User     common.Address
	DonLabel [32]byte
	Limit    uint32
	Raw      types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterUserDONLimitSet(opts *bind.FilterOpts, user []common.Address, donLabel [][32]byte) (*WorkflowRegistryUserDONLimitSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "UserDONLimitSet", userRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryUserDONLimitSetIterator{contract: _WorkflowRegistry.contract, event: "UserDONLimitSet", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchUserDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitSet, user []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "UserDONLimitSet", userRule, donLabelRule)
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
	User     common.Address
	DonLabel [32]byte
	Raw      types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterUserDONLimitUnset(opts *bind.FilterOpts, user []common.Address, donLabel [][32]byte) (*WorkflowRegistryUserDONLimitUnsetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "UserDONLimitUnset", userRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryUserDONLimitUnsetIterator{contract: _WorkflowRegistry.contract, event: "UserDONLimitUnset", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchUserDONLimitUnset(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitUnset, user []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "UserDONLimitUnset", userRule, donLabelRule)
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
	DonLabel     [32]byte
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowActivated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowActivatedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowActivated", workflowIdRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowActivatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowActivated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowActivated", workflowIdRule, ownerRule, donLabelRule)
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
	DonLabel     [32]byte
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowDeleted(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowDeletedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowDeleted", workflowIdRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowDeletedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowDeleted", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowDeleted", workflowIdRule, ownerRule, donLabelRule)
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
	DonLabel     [32]byte
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowPaused(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowPausedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowPaused", workflowIdRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowPausedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowPaused", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowPaused", workflowIdRule, ownerRule, donLabelRule)
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
	DonLabel     [32]byte
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
	DonLabel      [32]byte
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
	return common.HexToHash("0x3800848414d9fa9c555c8b830833c55ea97e0743a1de02f19a0dd76e9578ebab")
}

func (WorkflowRegistryDONRegistryUpdated) Topic() common.Hash {
	return common.HexToHash("0xc60b6bc6904a32fe864ae9ad697fd792f40e1d9e9fee5da0a10d0fb26b7380d6")
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
	return common.HexToHash("0x2a666be2166cbc432b468f0c0273b4c7009e2fc0168952f8adc3c3b0aecf8c61")
}

func (WorkflowRegistryUserDONLimitUnset) Topic() common.Hash {
	return common.HexToHash("0xb991c2ee405ccdec48f880c2f2edd1b4161ceefb047f499c9a024502779a039a")
}

func (WorkflowRegistryWorkflowActivated) Topic() common.Hash {
	return common.HexToHash("0x3131ee488b1b459e313b249599f9ebc23dfc5dbfcbd082dd503d171c7436fe82")
}

func (WorkflowRegistryWorkflowDeleted) Topic() common.Hash {
	return common.HexToHash("0xbe9d8e7e7313f6e3020af5157a64045cca8bcadce6e6497ca12fe56dbb4fdeb9")
}

func (WorkflowRegistryWorkflowPaused) Topic() common.Hash {
	return common.HexToHash("0x52613f4f9f39f66629197514b47e67cab5e5c40b83fad8fccec907d71a8a3978")
}

func (WorkflowRegistryWorkflowRegistered) Topic() common.Hash {
	return common.HexToHash("0x03bdc9e1a7b65f76fa1e989e76c28e6ac9f556262cfed31a11a9461e77ad2ff5")
}

func (WorkflowRegistryWorkflowUpdated) Topic() common.Hash {
	return common.HexToHash("0x67057592e7b546efc47d8093ece5f7d8837b769a512710d66a892555875bc05e")
}

func (_WorkflowRegistry *WorkflowRegistry) Address() common.Address {
	return _WorkflowRegistry.address
}

type WorkflowRegistryInterface interface {
	CanLinkOwner(opts *bind.CallOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) error

	CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error

	GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel [32]byte) (uint32, error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel [32]byte) (uint32, error)

	GetWorkflowMetadata(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByDON(opts *bind.CallOpts, donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByOwnerName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error)

	IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	AdminBatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	AdminPauseAllByDON(opts *bind.TransactOpts, donLabel [32]byte) (*types.Transaction, error)

	AdminPauseAllByUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error)

	AdminPauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	DeleteWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error)

	PauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	PauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	SetDONLimit(opts *bind.TransactOpts, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error)

	SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector *big.Int) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donLabel [32]byte, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error)

	FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error)

	WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error)

	ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error)

	FilterDONLimitSet(opts *bind.FilterOpts, donLabel [][32]byte) (*WorkflowRegistryDONLimitSetIterator, error)

	WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet, donLabel [][32]byte) (event.Subscription, error)

	ParseDONLimitSet(log types.Log) (*WorkflowRegistryDONLimitSet, error)

	FilterDONRegistryUpdated(opts *bind.FilterOpts, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (*WorkflowRegistryDONRegistryUpdatedIterator, error)

	WatchDONRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONRegistryUpdated, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (event.Subscription, error)

	ParseDONRegistryUpdated(log types.Log) (*WorkflowRegistryDONRegistryUpdated, error)

	FilterOwnershipLinkUpdated(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedIterator, error)

	WatchOwnershipLinkUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdated, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error)

	ParseOwnershipLinkUpdated(log types.Log) (*WorkflowRegistryOwnershipLinkUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error)

	FilterUserDONLimitSet(opts *bind.FilterOpts, user []common.Address, donLabel [][32]byte) (*WorkflowRegistryUserDONLimitSetIterator, error)

	WatchUserDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitSet, user []common.Address, donLabel [][32]byte) (event.Subscription, error)

	ParseUserDONLimitSet(log types.Log) (*WorkflowRegistryUserDONLimitSet, error)

	FilterUserDONLimitUnset(opts *bind.FilterOpts, user []common.Address, donLabel [][32]byte) (*WorkflowRegistryUserDONLimitUnsetIterator, error)

	WatchUserDONLimitUnset(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitUnset, user []common.Address, donLabel [][32]byte) (event.Subscription, error)

	ParseUserDONLimitUnset(log types.Log) (*WorkflowRegistryUserDONLimitUnset, error)

	FilterWorkflowActivated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowActivatedIterator, error)

	WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error)

	ParseWorkflowActivated(log types.Log) (*WorkflowRegistryWorkflowActivated, error)

	FilterWorkflowDeleted(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowDeletedIterator, error)

	WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error)

	ParseWorkflowDeleted(log types.Log) (*WorkflowRegistryWorkflowDeleted, error)

	FilterWorkflowPaused(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowPausedIterator, error)

	WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowId [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error)

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
