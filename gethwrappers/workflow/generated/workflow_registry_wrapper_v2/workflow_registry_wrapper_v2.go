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
	DonLabel     string
	WorkflowName string
	BinaryUrl    string
	ConfigUrl    string
	Tag          string
	Attributes   []byte
}

var WorkflowRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByUser\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadata\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwnerName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUnlinkWithActiveWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerDONExceeded\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowAlreadyInDesiredStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561549c90816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe610100604052600436101561001357600080fd5b60003560e01c8063034a6ed5146130565780630987294c14612fe757806317e0edfc14612ea6578063181f5a7714612e585780631c387fd614612d965780631c71682c14612b735780633514125114612af65780633c17181b14612a7c57806347d1ed8314612a4a57806352453a4414612940578063695e1340146128985780636ee80b44146127375780636f3517711461265957806379ba5097146125705780638da5cb5b1461251e5780639c5b230414612446578063a0b8a4fe1461240a578063a7d0185814612345578063b377bfc514611048578063b668435f14610fc2578063b87a019414610dec578063bdf6b4ff14610d6e578063be67433314610cfd578063cabb9e7a14610c93578063cb647e8214610bfd578063d8e056de14610b0f578063d8e4a72414610945578063dc101969146108b1578063dfcb0b3114610897578063e690f33214610843578063e9df65531461035a578063eec860b71461027d5763f2fde38b1461018857600080fd5b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785773ffffffffffffffffffffffffffffffffffffffff6101d4613361565b6101dc614193565b1633811461024e57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102785760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610278576102b4613361565b602435906102c0614193565b73ffffffffffffffffffffffffffffffffffffffff6012541682602073ffffffffffffffffffffffffffffffffffffffff604051946102fe866133f1565b16938481520152817fffffffffffffffffffffffff00000000000000000000000000000000000000006012541617601255826013557fc60b6bc6904a32fe864ae9ad697fd792f40e1d9e9fee5da0a10d0fb26b7380d6600080a4005b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043567ffffffffffffffff8111610278576103a99036906004016135be565b908115610819576103b9336147e3565b6103c282613787565b916103d0604051938461340d565b8083527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06103fd82613787565b0136602085013761040d81613787565b61041a604051918261340d565b8181527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061044783613787565b0136602083013761045782613787565b610464604051918261340d565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061049184613787565b0160005b818110610808575050600091825b8481106106b3575060005b83811061051c57858560005b8181106104c357005b806104d16001928486613b04565b3560005260096020526040600020546104ea8133614872565b8360ff8183015460e01c166104fe816131f1565b1461050c575b5050016104ba565b610515916148f1565b8480610504565b61052681886137ac565b5163ffffffff61053683856137ac565b51169061054383866137ac565b519181600052601060205263ffffffff60406000205416918233600052601160205260406000208260005260205260ff60406000205460201c1661068a575b81600052600b60205261059a8360406000205461379f565b1161064f5733600052600f60205260406000209060005260205263ffffffff60406000205416019063ffffffff82116106205763ffffffff8091169116116105e557506001016104ae565b61061c906040519182917f038857ff000000000000000000000000000000000000000000000000000000008352336004840161410b565b0390fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6040517fb9938687000000000000000000000000000000000000000000000000000000008152602060048201528061061c602482018761322a565b925033600052601160205260406000208160005260205263ffffffff6040600020541692610582565b6106be818688613b04565b3560005260096020526106d660406000205433614872565b60ff600182015460e01c166106ea816131f1565b156107ff5760020160405161070a8161070381856136e5565b038261340d565b602081519101206000805b8781106107a9575b5015610730575b50506001905b016104a3565b6107039161075291610745888c9996996137ac565b52604051928380926136e5565b61075c82856137ac565b5261076781846137ac565b50600161077482846137ac565b527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610620576001809101939088610724565b826107b4828d6137ac565b51146107c257600101610715565b905063ffffffff6107d382876137ac565b51169063ffffffff82146106205760016107f263ffffffff92886137ac565b920116905260018a61071d565b5060019061072a565b806060602080938601015201610495565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785761087b336147e3565b600435600052600960205261089560406000205433614845565b005b34610278576108956108a8366135ef565b92919091613e81565b346102785760016108cd6108c4366135ef565b91809493613e81565b336000526005602052806040600020556108e633615075565b508060005260066020526040600020827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346102785760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043567ffffffffffffffff8111610278576109949036906004016135be565b60243591821515809303610278576109aa614193565b60ff831660005b838110610a5557505060405191806040840160408552526060830191906000905b808210610a08577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b90919283359073ffffffffffffffffffffffffffffffffffffffff82168203610278576020809173ffffffffffffffffffffffffffffffffffffffff6001941681520194019201906109d2565b73ffffffffffffffffffffffffffffffffffffffff610a7d610a78838787613b04565b613e27565b1615610ae5578073ffffffffffffffffffffffffffffffffffffffff610aa9610a786001948888613b04565b1660005260026020526040600020837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055016109b1565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857610b46614193565b600435600052600b60205260406000205b80548015610895577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161062057610b9390826149de565b90549060031b1c6000526007602052604060002054610bb0614193565b6000526009602052604060002054806000526007602052604060002060ff600182015460e01c16610be0816131f1565b15610bed575b5050610b57565b610bf6916141de565b8180610be6565b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043567ffffffffffffffff811161027857610c4c9036906004016135be565b801561081957610c5b336147e3565b60005b818110610c6757005b80610c756001928486613b04565b356000526009602052610c8d60406000205433614845565b01610c5e565b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785773ffffffffffffffffffffffffffffffffffffffff610cdf613361565b166000526002602052602060ff604060002054166040519015158152f35b346102785760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857610d34613361565b60243567ffffffffffffffff811161027857602091610d5a610d60923690600401613493565b91613d9a565b63ffffffff60405191168152f35b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043567ffffffffffffffff811161027857610dc0610dc7913690600401613493565b36916137ef565b602081519101206000526010602052602063ffffffff60406000205416604051908152f35b346102785760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785773ffffffffffffffffffffffffffffffffffffffff610e38613361565b1680600052600d602052610e56604435602435604060002054614138565b90610e6082613a97565b9260005b838110610e7d5760405180610e798782613540565b0390f35b60019082600052600d602052610ea16040600020610e9b838761379f565b906149de565b90549060031b1c600052600760205260406000206007610fa060405192610ec7846133b8565b8054845260ff8682015473ffffffffffffffffffffffffffffffffffffffff8116602087015267ffffffffffffffff8160a01c16604087015260e01c16610f0d816131f1565b6060850152604051610f268161070381600286016136e5565b6080850152604051610f3f8161070381600386016136e5565b60a0850152604051610f588161070381600486016136e5565b60c0850152604051610f718161070381600586016136e5565b60e0850152604051610f8a8161070381600686016136e5565b61010085015261070360405180948193016136e5565b610120820152610fb082886137ac565b52610fbb81876137ac565b5001610e64565b346102785760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857610ff9613361565b60243567ffffffffffffffff811161027857611019903690600401613493565b6044359063ffffffff82168203610278576064359283151584036102785761089594611043614193565b613b78565b34610278576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043567ffffffffffffffff811161027857611098903690600401613493565b60c05260a05260243567ffffffffffffffff8111610278576110be903690600401613493565b90600260643510156102785760843567ffffffffffffffff8111610278576110ea903690600401613493565b9060e0529160a43567ffffffffffffffff81116102785761110f903690600401613493565b91909260c43567ffffffffffffffff811161027857611132903690600401613493565b92909160e43567ffffffffffffffff811161027857611155903690600401613493565b92909361010435958615158703610278576111753660c05160a0516137ef565b6111d86112046111863687896137ef565b6040519283916111a8602084019633885260606040860152608085019061322a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084830301606085015261322a565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261340d565b5190209889600052600760205273ffffffffffffffffffffffffffffffffffffffff6001604060002001541615600014611eca57611241336147e3565b60c05115611ea057604060c05111611e6b578315611e415760208411611e0e5761126c604435614e83565b611276828a614ecc565b61127f86614f11565b6112976112913660c05160a0516137ef565b336147c4565b986112a5368d60e0516137ef565b60208151910120986112b86064356131f1565b8c60643515611df0575b5015611c9e575b926113729492611356611364936113819a9998968f611333906040516080526112f36080516133b8565b60443560805152336020608051015267ffffffffffffffff42166040608051015261131f6064356131f1565b60643560606080510152369060e0516137ef565b6080805101526113483660c05160a0516137ef565b60a0608051015236916137ef565b60c0608051015236916137ef565b60e0608051015236916137ef565b610100608051015236916137ef565b610120608051015282600052600760205260406000206080515181556001810173ffffffffffffffffffffffffffffffffffffffff60206080510151167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416178155604060805101517fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000083549260a01b16911617815560606080510151611446816131f1565b61144f816131f1565b7fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff7cff0000000000000000000000000000000000000000000000000000000083549260e01b1691161790556002810160808051015180519067ffffffffffffffff8211611926576114ca826114c48554613692565b85614319565b602090601f8311600114611bfb57611518929160009183611a07575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6003810160a0608051015180519067ffffffffffffffff821161192657611546826114c48554613692565b602090601f8311600114611b5857611593929160009183611a075750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6004810160c0608051015180519067ffffffffffffffff8211611926576115c1826114c48554613692565b602090601f8311600114611ab55761160e929160009183611a075750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6005810160e0608051015180519067ffffffffffffffff82116119265761163c826114c48554613692565b602090601f8311600114611a1257611689929160009183611a075750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b60068101610100608051015180519067ffffffffffffffff8211611926576116b8826114c48554613692565b602090601f831160011461196057918061170a9260079695946000926119555750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b01610120608051015180519067ffffffffffffffff821161192657611736826114c48554613692565b602090601f8311600114611883576117839291600091836118785750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b81600052600860205261179d83604060002061510a565b5060443560005260096020528260406000205580600052600c6020526117c783604060002061510a565b5033600052600d6020526117df83604060002061510a565b506117eb6064356131f1565b60643515611865575b50505061180f6040519160608352606083019060e051613b14565b61181a6064356131f1565b606435602083015281810360408301527f74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a093392806118606044359460c05160a051613b14565b0390a3005b611870923390614f7d565b8180806117f4565b0151905088806114e6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b81811061190e57509084600195949392106118d7575b505050811b019055611786565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558780806118ca565b929360206001819287860151815501950193016118b4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b015190508a806114e6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b8181106119ef5750916001939185600798979694106119b8575b505050811b01905561170d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558980806119ab565b92936020600181928786015181550195019301611991565b0151905089806114e6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611a9d5750908460019594939210611a66575b505050811b01905561168c565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055888080611a59565b92936020600181928786015181550195019301611a43565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611b405750908460019594939210611b09575b505050811b019055611611565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055888080611afc565b92936020600181928786015181550195019301611ae6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611be35750908460019594939210611bac575b505050811b019055611596565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055888080611b9f565b92936020600181928786015181550195019301611b89565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611c865750908460019594939210611c4f575b505050811b01905561151b565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055888080611c42565b92936020600181928786015181550195019301611c2c565b98999593918a60009c98969492999c52600e6020526040600020985b89546001811115611dd8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111610620578c611d7d8f92611cff908e6149de565b90549060031b1c91826000526007602052604060002093600185019373ffffffffffffffffffffffffffffffffffffffff85547c01000000000000000000000000000000000000000000000000000000007fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff821617875516906149f6565b7ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e488689611dd073ffffffffffffffffffffffffffffffffffffffff845493541693604051918160026003859401910183614b1a565b0390a3611cba565b50979b969a99969850949693959294919390926112c9565b611e01611e0891369060e0516137ef565b8b33614d54565b8c6112c2565b837f436f975400000000000000000000000000000000000000000000000000000000600052600452602060245260446000fd5b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b7f36a7c5030000000000000000000000000000000000000000000000000000000060005260c051600452604060245260446000fd5b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b935093955093975050611edc336147e3565b611ee7604435614e83565b611ef18286614ecc565b611efa84614f11565b611f048633614872565b958654600052600960205260006040812055604435600052600960205260406000205585549460443587556004870191611f59604051611f488161070381886136e5565b611f533685856137ef565b90614f50565b15612225575b5050506005850191611f7b604051611f488161070381886136e5565b15612105575b505050600783019067ffffffffffffffff811161192657611fac81611fa68454613692565b84614319565b6000601f8211600114612065578190611ffb9394959660009261205a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b612016604051926040845260026040850191016136e5565b9082820360208401527f03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e3223393806120556044359560c05160a051613b14565b0390a4005b0135905086806114e6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216958382526020822091805b8881106120ed575083600195969798106120b5575b505050811b019055611ffe565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558580806120a8565b90926020600181928686013581550194019101612093565b67ffffffffffffffff821161192657612122826114c48554613692565b600090601f83116001146121855761216f92916000918361217a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b848080611f81565b0135905088806114e6565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b81811061220d57509084600195949392106121d5575b505050811b019055612172565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558780806121c8565b919360206001819287870135815501950192016121b2565b67ffffffffffffffff821161192657612242826114c48554613692565b600090601f83116001146122a55761228f92916000918361229a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b868080611f5f565b013590508a806114e6565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b81811061232d57509084600195949392106122f5575b505050811b019055612292565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558980806122e8565b919360206001819287870135815501950192016122d2565b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043567ffffffffffffffff8111610278576123949036906004016135be565b80156108195760005b8181106123a657005b806123b46001928486613b04565b356123bd614193565b6000526009602052604060002054806000526007602052604060002060ff8482015460e01c166123ec816131f1565b156123fa575b50500161239d565b612403916141de565b84806123f2565b346102785760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610278576020600354604051908152f35b346102785760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043580600052600c602052612495604435602435604060002054614138565b9061249f82613a97565b9260005b8381106124b85760405180610e798782613540565b60019082600052600c6020526124d66040600020610e9b838761379f565b90549060031b1c6000526007602052604060002060076124fc60405192610ec7846133b8565b61012082015261250c82886137ac565b5261251781876137ac565b50016124a3565b346102785760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102785760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760005473ffffffffffffffffffffffffffffffffffffffff8116330361262f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857612691336147e3565b60043560005260096020526040600020546126ac8133614872565b60ff600182015460e01c166126c0816131f1565b1561270d57610895916127086126fb600284016127026040516126e78161070381866136e5565b6020815191012091604051938480926136e5565b038361340d565b33614d54565b6148f1565b7f6f861db10000000000000000000000000000000000000000000000000000000060005260046000fd5b3461027857612763612748366134c1565b93949161275485613489565b8415938461288a575b866145cc565b61276c82613489565b156127d3575b600073ffffffffffffffffffffffffffffffffffffffff831661279481614cf1565b9080835260056020528260408120556127ac81615201565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b73ffffffffffffffffffffffffffffffffffffffff8216600052600a60205260406000205b8054801561286d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019081116106205761283490826149de565b90549060031b1c6128458185614872565b61284e84613489565b600284036128645761285f916141de565b6127f8565b61285f91614b8e565b50505073ffffffffffffffffffffffffffffffffffffffff612772565b61289387614578565b61275d565b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043560005260096020526108956040600020546128e58133614872565b9060ff600183015460e01c166128fa816131f1565b614b8e5761293b6040516129158161070381600288016136e5565b60208151910120612933604051611291816107038160038a016136e5565b9033846149f6565b614b8e565b346102785760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857612977613361565b6024359067ffffffffffffffff82116102785761299e610dc06129a4933690600401613493565b906147c4565b8060005260086020526129c1606435604435604060002054614138565b906129cb82613a97565b9260005b8381106129e45760405180610e798782613540565b600190826000526008602052612a026040600020610e9b838761379f565b90549060031b1c600052600760205260406000206007612a2860405192610ec7846133b8565b610120820152612a3882886137ac565b52612a4381876137ac565b50016129cf565b3461027857610895612a5b366134c1565b909391929190612a6a81613489565b6145cc57612a7781614578565b6145cc565b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857612ab3614193565b6004356000526009602052604060002054806000526007602052604060002060ff600182015460e01c16612ae6816131f1565b15612aed57005b610895916141de565b346102785760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785760043567ffffffffffffffff811161027857612b45903690600401613493565b60243563ffffffff81168103610278576044359182151583036102785761089593612b6e614193565b613854565b3461027857612b8d612b8436613384565b90601454614138565b612b9681613787565b91612ba4604051938461340d565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612bd183613787565b0160005b818110612d6a575050601454919060005b828110612c9b57846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210612c2657505050500390f35b91936020612c8b827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652606060408a518051612c6b81613489565b845263ffffffff868201511686850152015191816040820152019061322a565b9601920192018594939192612c17565b612ca5818361379f565b600085821015612d3d5760149052604051600192918390612d1c90821b7fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec01612ced846133d5565b63ffffffff815460ff8116612d0181613489565b865260081c16602085015261070360405180948193016136e5565b6040820152612d2b82886137ac565b52612d3681876137ac565b5001612be6565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b602090604051612d79816133d5565b600081526000838201526060604082015282828801015201612bd5565b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785773ffffffffffffffffffffffffffffffffffffffff612de2613361565b612dea614193565b16600052600a60205260406000205b80548015610895577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161062057612e38612e5391836149de565b90549060031b1c8060005260076020526040600020906141de565b612df9565b346102785760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857610e79612e9261344e565b60405191829160208352602083019061322a565b3461027857612ec0612eb736613384565b90600354614138565b612ec981613787565b91612ed7604051938461340d565b818352612ee382613787565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020850193013684376003549160005b828110612f775784866040519182916020830190602084525180915260408301919060005b818110612f48575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101612f3a565b612f81818361379f565b600085821015612d3d57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0154600081905260056020526001919073ffffffffffffffffffffffffffffffffffffffff16612fe082896137ac565b5201612f15565b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027857602061304c73ffffffffffffffffffffffffffffffffffffffff613038613361565b166000526004602052604060002054151590565b6040519015158152f35b346102785760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102785761308d613644565b50600435600052600960205260406000205480156131c7576000526007602052604060002073ffffffffffffffffffffffffffffffffffffffff604051916130d4836133b8565b8054835260076131a260018301549260ff6020870194868116865267ffffffffffffffff8160a01c16604089015260e01c1661310f816131f1565b60608701526040516131288161070381600286016136e5565b60808701526040516131418161070381600386016136e5565b60a087015260405161315a8161070381600486016136e5565b60c08701526040516131738161070381600586016136e5565b60e087015260405161318c8161070381600686016136e5565b61010087015261070360405180948193016136e5565b6101208401525116156131c757610e7990604051918291602083526020830190613289565b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b600211156131fb57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b919082519283825260005b8481106132745750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613235565b61335e918151815273ffffffffffffffffffffffffffffffffffffffff602083015116602082015267ffffffffffffffff604083015116604082015260608201516132d3816131f1565b606082015261012061334c6133386133266133146133026080880151610140608089015261014088019061322a565b60a088015187820360a089015261322a565b60c087015186820360c088015261322a565b60e086015185820360e087015261322a565b61010085015184820361010086015261322a565b9201519061012081840391015261322a565b90565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361027857565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112610278576004359060243590565b610140810190811067ffffffffffffffff82111761192657604052565b6060810190811067ffffffffffffffff82111761192657604052565b6040810190811067ffffffffffffffff82111761192657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761192657604052565b6040519061345d60408361340d565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b600311156131fb57565b9181601f840112156102785782359167ffffffffffffffff8311610278576020838186019501011161027857565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102785760043573ffffffffffffffffffffffffffffffffffffffff811681036102785791602435916044359067ffffffffffffffff82116102785761352f91600401613493565b909160643560038110156102785790565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061357357505050505090565b90919293946020806135af837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951613289565b97019301930191939290613564565b9181601f840112156102785782359167ffffffffffffffff8311610278576020808501948460051b01011161027857565b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102785760043591602435916044359067ffffffffffffffff82116102785761364091600401613493565b9091565b60405190613651826133b8565b6060610120836000815260006020820152600060408201526000838201528260808201528260a08201528260c08201528260e0820152826101008201520152565b90600182811c921680156136db575b60208310146136ac57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916136a1565b600092918154916136f583613692565b808352926001811690811561374b575060011461371157505050565b60009081526020812093945091925b838310613731575060209250010190565b600181602092949394548385870101520191019190613720565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b67ffffffffffffffff81116119265760051b60200190565b9190820180921161062057565b80518210156137c05760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b92919267ffffffffffffffff82116119265760405191613837601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0166020018461340d565b829481845281830111610278578281602093846000960137010152565b9091926138623684846137ef565b602081519101209081600052601060205260406000209060405191613886836133f1565b549063ffffffff8216835260ff602084019260201c161590811583526000146139e0575051151590816139c9575b506139c3577f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b731936139be91604051906138ec826133f1565b63ffffffff83168252602082019060018252600052601060205263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b1691161790556139b260405161398b816111d885898960208501613b53565b60405190613998826133d5565b6000825263ffffffff421660208301526040820152614360565b60405193849384613b53565b0390a1565b50505050565b63ffffffff9150511663ffffffff851614386138b4565b9295505050919091613a92577f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b73192600052601060205260006040812055613a6f6040516040602082015261398b81613a3c606082018688613b14565b60006040830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261340d565b613a86604051928392604084526040840191613b14565b600060208301520390a1565b505050565b90613aa182613787565b613aae604051918261340d565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613adc8294613787565b019060005b828110613aed57505050565b602090613af8613644565b82828501015201613ae1565b91908110156137c05760051b0190565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91613b7160209263ffffffff92969596604086526040860191613b14565b9416910152565b93929373ffffffffffffffffffffffffffffffffffffffff613b9b3685856137ef565b60208151910120911694856000526011602052604060002082600052602052604060002060405190613bcc826133f1565b549563ffffffff8716825260ff602083019760201c16159081158852600014613d39575082600052601060205263ffffffff604060002054169563ffffffff8316968711613d0f5786905115159182613cfe575b5050613cf6577f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c8232494613cf19260405191613c59836133f1565b825260208201906001825288600052601160205260406000209060005260205263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b16911617905560405193849384613b53565b0390a2565b505050505050565b5163ffffffff161490508538613c20565b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b92955050509190916139c3577f6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d8499284600052601160205260406000209060005260205260006040812055613cf1604051928392602084526020840191613b14565b9173ffffffffffffffffffffffffffffffffffffffff91613dbc9136916137ef565b60208151910120911660005260116020526040600020816000526020526040600020602060405191613ded836133f1565b549160ff63ffffffff841693848352831c1615159182910152613e225750600052601060205263ffffffff6040600020541690565b905090565b3573ffffffffffffffffffffffffffffffffffffffff811681036102785790565b909260809273ffffffffffffffffffffffffffffffffffffffff61335e9795168352602083015260408201528160608201520191613b14565b919290928242116140d557613ea3336000526004602052604060002054151590565b6140a75783600052600660205260ff60406000205416614075576000613ec761344e565b604051613f3581613efd602082019486865233604084015246606084015230608084015260e060a084015261010083019061322a565b8860c08301528960e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261340d565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613f76603c8220613f703686866137ef565b90615039565b90919260048310156140485782613ff55750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1615613fbd5750505050565b9061061c916040519485947f335d4ce10000000000000000000000000000000000000000000000000000000086523360048701613e48565b8593505060ff6140386040519586957fd36ab6b9000000000000000000000000000000000000000000000000000000008752606060048801526064870191613b14565b9216602484015260448301520390fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526021600452fd5b837f77a33858000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b7fd9a5f5ca000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b827f502d038700000000000000000000000000000000000000000000000000000000600052336004524260245260445260646000fd5b60409073ffffffffffffffffffffffffffffffffffffffff61335e9493168152816020820152019061322a565b92918382101561418c5780158015614182575b61417a575b61415a908261379f565b92808411614172575b50808303928311610620579190565b925038614163565b506064614150565b506064811161414b565b5050600090565b73ffffffffffffffffffffffffffffffffffffffff6001541633036141b457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b907ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e488689600182019173ffffffffffffffffffffffffffffffffffffffff8084547c01000000000000000000000000000000000000000000000000000000007fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff821617865516916142a8600282019660405161427c81610703818c6136e5565b6020815191012060038401956142a260405161429c81610703818c6136e5565b826147c4565b926149f6565b54935416936142ee6040516142c7816111d886868b8b60208601614ad4565b604051906142d4826133d5565b6002825263ffffffff421660208301526040820152614360565b6142fd60405192839283614b1a565b0390a3565b81811061430d575050565b60008155600101614302565b9190601f811161432857505050565b614354926000526020600020906020601f840160051c83019310614356575b601f0160051c0190614302565b565b9091508190614347565b906014546801000000000000000081101561192657600181016014556000601454821015612d3d57601490526000929060011b7fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec01600160409183516143c581613489565b6143ce81613489565b60ff82549116807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083161783557fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ffffffff00602088015160081b16921617178155019101519283519067ffffffffffffffff821161454b57614456826114c48554613692565b602090601f83116001146144b157906144a2939495836144a65750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0151905038806114e6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316848352818320925b818110614533575095836001959697106144fc575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806144f2565b9192602060018192868b0151815501940192016144dd565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b73ffffffffffffffffffffffffffffffffffffffff16600052600a6020526040600020546145a257565b7f61bc2e180000000000000000000000000000000000000000000000000000000060005260046000fd5b91929092834211614778576145e0836147e3565b73ffffffffffffffffffffffffffffffffffffffff831661460081614cf1565b9060009061464461467c61461261344e565b604051928391602083019560018752604084015246606084015230608084015260e060a084015261010083019061322a565b8a60c08301528660e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261340d565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c526146b7603c8220613f703687876137ef565b909192600483101561404857826147375750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff16156146ff575050505050565b9061061c92916040519586957f335d4ce100000000000000000000000000000000000000000000000000000000875260048701613e48565b506040517fd36ab6b90000000000000000000000000000000000000000000000000000000081526060600482015291829160ff614038606485018a8a613b14565b8373ffffffffffffffffffffffffffffffffffffffff847f3d8a511600000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b906147dd6111d89160405192839160208301958661410b565b51902090565b73ffffffffffffffffffffffffffffffffffffffff16614810816000526004602052604060002054151590565b156148185750565b7fc2dda3f90000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8161484f91614872565b600160ff8183015460e01c16614864816131f1565b1461270d57614354916141de565b90600052600760205260406000209073ffffffffffffffffffffffffffffffffffffffff6001830154169081156131c75773ffffffffffffffffffffffffffffffffffffffff168091036148c4575090565b7f31ee6dc70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b907f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a600182019173ffffffffffffffffffffffffffffffffffffffff8084541691614971600282019660405161494b81610703818c6136e5565b60208151910120600384019561496b60405161429c81610703818c6136e5565b92614f7d565b8454947fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff86169055549316936142ee6040516149b7816111d886868b8b60208601614ad4565b604051906149c4826133d5565b6001825263ffffffff421660208301526040820152614360565b80548210156137c05760005260206000200190600090565b919073ffffffffffffffffffffffffffffffffffffffff1680600052600a602052614a2583604060002061535b565b5081600052600b602052614a3d83604060002061535b565b50600052600f6020526040600020906000526020526040600020917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff845416019263ffffffff84116106205763ffffffff614ad194167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055600052600e602052604060002061535b565b50565b9273ffffffffffffffffffffffffffffffffffffffff61335e9593614b0c9386521660208501526080604085015260808401906136e5565b9160608184039101526136e5565b9091614b3161335e936040845260408401906136e5565b9160208184039101526136e5565b614b498154613692565b9081614b53575050565b81601f60009311600114614b65575055565b81835260208320614b8191601f0160051c810190600101614302565b8082528160208120915555565b906002810190604051614ba58161070381866136e5565b60208151910120600052600c602052614bc283604060002061535b565b507f48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd94600182019173ffffffffffffffffffffffffffffffffffffffff835416600052600d602052614c1785604060002061535b565b5073ffffffffffffffffffffffffffffffffffffffff8084541691614c51600382019360405190614c4c826126fb81896136e5565b6147c4565b6000526008602052614c6787604060002061535b565b5080546000526009602052600060408120555493541693614c8d60405192839283614b1a565b0390a36000526007602052600760406000206000815560006001820155614cb660028201614b3f565b614cc260038201614b3f565b614cce60048201614b3f565b614cda60058201614b3f565b614ce660068201614b3f565b01614b498154613692565b80600052600560205260406000205490811580614d3e575b614d11575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5080600052600460205260406000205415614d09565b9080600052601060205263ffffffff60406000205416908173ffffffffffffffffffffffffffffffffffffffff84169081600052601160205260406000208360005260205260ff60406000205460201c16614e5a575b82600052600b60205260406000205411614e1f57600052600f60205260406000209060005260205263ffffffff8060406000205416911610614dea575050565b61061c6040519283927f038857ff0000000000000000000000000000000000000000000000000000000084526004840161410b565b6040517fb9938687000000000000000000000000000000000000000000000000000000008152602060048201528061061c602482018861322a565b925080600052601160205260406000208260005260205263ffffffff6040600020541692614daa565b8015610ae557806000526009602052604060002054614e9f5750565b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60c88111614ee0575060c88111614ee05750565b7ecd56a80000000000000000000000000000000000000000000000000000000060005260045260c860245260446000fd5b6102008111614f1d5750565b7f354f25140000000000000000000000000000000000000000000000000000000060005260045261020060245260446000fd5b9081518151908181149384614f67575b5050505090565b6020929394508201209201201438808080614f60565b91929073ffffffffffffffffffffffffffffffffffffffff1680600052600f602052604060002084600052602052604060002093600163ffffffff865416019463ffffffff86116106205763ffffffff614ad196167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055600052600b60205261500f83604060002061510a565b50600052600a60205261502682604060002061510a565b50600052600e602052604060002061510a565b815191906041830361506a5761506392506020820151906060604084015193015160001a9061515f565b9192909190565b505060009160029190565b806000526004602052604060002054156000146151045760035468010000000000000000811015611926576150eb6150b682600185940160035560036149de565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600354906000526004602052604060002055600190565b50600090565b600082815260018201602052604090205461418c578054906801000000000000000082101561192657826151486150b68460018096018555846149de565b905580549260005201602052604060002055600190565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116151f5579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa156151e95760005173ffffffffffffffffffffffffffffffffffffffff8116156151dd5790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b50505060009160039190565b600081815260046020526040902054801561418c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161062057600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161062057818103615321575b50505060035480156152f2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016152af8160036149de565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6153436153326150b69360036149de565b90549060031b1c92839260036149de565b90556000526004602052604060002055388080615276565b9060018201918160005282602052604060002054801515600014615486577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610620578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116106205781810361544f575b505050805480156152f2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061541082826149de565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b61546f61545f6150b693866149de565b90549060031b1c928392866149de565b9055600052836020526040600020553880806153d8565b5050505060009056fea164736f6c634300081a000a",
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel string) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerDON", donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerDON(donLabel string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerDON(donLabel string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel string) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerUserDON", user, donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel string) (uint32, error) {
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONLimit(opts *bind.TransactOpts, donLabel string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONLimit", donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONLimit(donLabel string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONLimit(donLabel string, limit uint32, enabled bool) (*types.Transaction, error) {
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setUserDONOverride", user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetUserDONOverride(user common.Address, donLabel string, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetUserDONOverride(user common.Address, donLabel string, limit uint32, enabled bool) (*types.Transaction, error) {
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donLabel string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "upsertWorkflow", workflowName, tag, workflowId, status, donLabel, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donLabel string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowId, status, donLabel, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donLabel string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
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
	DonLabel string
	Limit    uint32
	Raw      types.Log
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
	DonLabel string
	Limit    uint32
	Raw      types.Log
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
	User     common.Address
	DonLabel string
	Raw      types.Log
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
	DonLabel     string
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
	DonLabel     string
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
	DonLabel     string
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
	DonLabel     string
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
	DonLabel      string
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
	return common.HexToHash("0x4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b731")
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

	GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel string) (uint32, error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel string) (uint32, error)

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

	SetDONLimit(opts *bind.TransactOpts, donLabel string, limit uint32, enabled bool) (*types.Transaction, error)

	SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector *big.Int) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel string, limit uint32, enabled bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donLabel string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error)

	FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error)

	WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error)

	ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error)

	FilterDONLimitSet(opts *bind.FilterOpts) (*WorkflowRegistryDONLimitSetIterator, error)

	WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet) (event.Subscription, error)

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
