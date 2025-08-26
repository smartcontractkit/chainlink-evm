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

type WorkflowRegistryDonConfigView struct {
	DonHash      [32]byte
	Family       string
	Limit        uint32
	LimitEnabled bool
}

type WorkflowRegistryEventRecord struct {
	EventType uint8
	Timestamp uint32
	Payload   []byte
}

type WorkflowRegistryOwnerAllowlistedRequest struct {
	RequestDigest   [32]byte
	Owner           common.Address
	ExpiryTimestamp uint32
}

type WorkflowRegistryUserOverrideView struct {
	User  common.Address
	Limit uint32
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
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlistRequest\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowlistedRequests\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"allowlistedRequests\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.OwnerAllowlistedRequest[]\",\"components\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilitiesRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonConfigs\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.DonConfigView[]\",\"components\":[{\"name\":\"donHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"family\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"limitEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxAttrLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxNameLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxTagLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxUrlLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"maxWorkflows\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserDONOverrides\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.UserOverrideView[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflow\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowById\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwnerAndName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCapabilitiesRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataConfig\",\"inputs\":[{\"name\":\"nameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"urlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"attrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsOnDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAllowlistedRequests\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalEvents\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWorkflowDONFamily\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilitiesRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataConfigUpdated\",\"inputs\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDonFamilyUpdated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowOwnerConfigUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BinaryURLRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUpdateDONFamilyForPausedWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DonLimitNotSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RequestExpired\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWorkflowIDNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561510690816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe610140604052600436101561001357600080fd5b60003560e01c806301b7690514612c1d5780630987294c14612bd957806317e0edfc14612acd578063181f5a7714612ab15780631c08b00a146129e25780631c71682c146127df5780631f77b2b6146127b6578063245b8e4e14612778578063274e00e01461271a578063289bd1081461264e5780632afc4130146125f65780632c50a95514612599578063351412511461254257806339d68c6a146123dd57806339e432341461235b5780633c17181b146123085780633c54b50b146122b55780634b6d2e5b14612113578063530979d614612084578063556dbd0d1461200d5780636104319314611f70578063695e134014611f1f5780636f19d3b814611ef957806370ae264014611ec357806379ba509714611e2a578063865ec9e014611df15780638b42a96d14611c9e5780638c42ffc514611bea5780638da5cb5b14611bc357806394ea0da614611a19578063952bb984146119ce57806396fbd7f0146118f1578063a0b8a4fe146118d3578063a408901614611830578063a7d0185814611789578063afbb240114611772578063b377bfc5146108cf578063b668435f14610867578063ba87068614610849578063bae5c29a146107d0578063bdf6b4ff1461076d578063be67433314610746578063cabb9e7a14610707578063d8b8073814610625578063d8e4a724146104cb578063dc1019691461042d578063de49b95f1461040f578063e627fc79146103e6578063e690f33214610362578063ea32308b14610328578063eb3da941146102fe5763f2fde38b1461025957600080fd5b346102f95760203660031901126102f9576001600160a01b0361027a612e2e565b610282614551565b163381146102cf57806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102f95760003660031901126102f957602061ffff61031c61453f565b60181c16604051908152f35b346102f95760203660031901126102f9576001600160a01b03610349612e2e565b16600052600b6020526020604060002054604051908152f35b346102f95760203660031901126102f95761038a336000526005602052604060002054151590565b156103d157600435600052600a6020526040600020546103aa8133614796565b600160ff8183015460e01c166103bf81612f40565b036103c657005b6103cf916145b4565b005b63c2dda3f960e01b6000523360045260246000fd5b346102f95760003660031901126102f957602060ff61040361453f565b60081c16604051908152f35b346102f95760003660031901126102f9576020601354604051908152f35b346102f95760603660031901126102f95760443560243567ffffffffffffffff82116102f9576104716104666001933690600401612f12565b908360043533613d54565b3360005260066020528060406000205561048a33614db0565b5080600052600760205260406000208260ff19825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346102f95760403660031901126102f95760043567ffffffffffffffff81116102f9576104fc9036906004016130bc565b602435918215158093036102f957610512614551565b60ff831660005b8381106105a357505060405191806040840160408552526060830191906000905b808210610570577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b9091928335906001600160a01b03821682036102f957602080916001600160a01b0360019416815201940192019061053a565b6001600160a01b036105be6105b9838787613509565b61437b565b16156105fb57806001600160a01b036105dd6105b96001948888613509565b16600052600360205260406000208360ff1982541617905501610519565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102f95760203660031901126102f95760043567ffffffffffffffff81116102f9576106569036906004016130bc565b80156106dd57610673336000526005602052604060002054151590565b156103d15760005b81811061068457005b806106926001928486613509565b35600052600a6020526040600020546106ab8133614796565b8360ff8183015460e01c166106bf81612f40565b036106cd575b50500161067b565b6106d6916145b4565b84806106c5565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346102f95760203660031901126102f9576001600160a01b03610728612e2e565b166000526003602052602060ff604060002054166040519015158152f35b346102f957602061075f61075936613167565b91614302565b63ffffffff60405191168152f35b346102f95760203660031901126102f95760043567ffffffffffffffff81116102f9576107a16107a8913690600401612f12565b36916132b3565b602081519101206000526015602052602063ffffffff60016040600020015416604051908152f35b346102f95760403660031901126102f9576107e9612e2e565b6040516108268161081860208201946024359086602090939291936001600160a01b0360408201951681520152565b03601f198101835282612eb5565b5190206000526012602052602063ffffffff604060002054166040519042108152f35b346102f95760003660031901126102f9576020601a54604051908152f35b346102f95760803660031901126102f957610880612e2e565b60243567ffffffffffffffff81116102f9576108a0903690600401612f12565b6044359063ffffffff821682036102f9576064359283151584036102f9576103cf946108ca614551565b6140dc565b346102f9576101203660031901126102f95760043567ffffffffffffffff81116102f957610901903690600401612f12565b610100526101205260243567ffffffffffffffff81116102f957610929903690600401612f12565b90600260643510156102f95760843567ffffffffffffffff81116102f957610955903690600401612f12565b9060e0529060a43567ffffffffffffffff81116102f95761097a903690600401612f12565b9160c43567ffffffffffffffff81116102f95761099b903690600401612f12565b9560e43567ffffffffffffffff81116102f9576109bc903690600401612f12565b929093610104359586151587036102f9576109e4336000526005602052604060002054151590565b156103d1576044351561174857604435600052600a6020526040600020546117185760ff610a1061453f565b60101c1688156116ee57806116b1575b5061ffff610a2c61453f565b60181c16801515806116a8575b6116775750831561164d5760ff610a4e61453f565b60081c1680151580611644575b611613575061010051156115e95760ff610a7361453f565b16801515806115dd575b6115a957506040516020810190610aa281610818888761010051610120513389613f65565b5190209788600052600860205260406000206001600160a01b036001820154168015600014611227575050610ae7610ae13661010051610120516132b3565b336143f2565b9a610af5368c60e0516132b3565b60208151910120981561119b575b610b0e606435612f40565b60643515611173575b604051978861012081011067ffffffffffffffff6101208b011117610f9c57610bc4899695610bb58897610bd3956101208a016040526044358a5260a0610ba960208c019a338c5260408d0160c05267ffffffffffffffff421660c0515260608d019a610b85606435612f40565b6064358c526080610b9d3661010051610120516132b3565b9e019d8e5236916132b3565b9b019a8b5236916132b3565b9860c08c01998a5236916132b3565b9760e08a0198895236916132b3565b610100880160a05260a051528860005260086020526040600020965187556001600160a01b038060018901935116166001600160a01b031983541617825560c051517fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b8085549360a01b1616911617825551610c5a81612f40565b610c6381612f40565b60ff60e01b197cff0000000000000000000000000000000000000000000000000000000083549260e01b16911617905560028501905180519067ffffffffffffffff8211610f9c57610cbf82610cb985546131cc565b8561374b565b602090601f831160011461110c57610cf0929160009183611101575b50508160011b916000199060031b1c19161790565b90555b518051600385019167ffffffffffffffff8211610f9c57610d1882610cb985546131cc565b602090601f831160011461109a57610d4892916000918361108f5750508160011b916000199060031b1c19161790565b90555b518051600484019167ffffffffffffffff8211610f9c57610d7082610cb985546131cc565b602090601f831160011461102857610da0929160009183610fb25750508160011b916000199060031b1c19161790565b90555b518051600583019167ffffffffffffffff8211610f9c57610dc882610cb985546131cc565b602090601f8311600114610fbd579180610dfd926006969594600092610fb25750508160011b916000199060031b1c19161790565b90555b019360a0515194855167ffffffffffffffff8111610f9c57610e2c81610e2684546131cc565b8461374b565b6020601f8211600114610f32579080610e6092610eb69899600092610f275750508160011b916000199060031b1c19161790565b90555b6000526009602052610e79826040600020614e81565b50604435600052600a60205281604060002055600052600d602052610ea2816040600020614e81565b5033600052600e6020526040600020614e81565b50610ecf6040519160608352606083019060e0516134bc565b610eda606435612f40565b606435602083015281810360408301527f74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a09339280610f226044359461010051610120516134bc565b0390a3005b015190508980610cdb565b601f1982169783600052816000209860005b818110610f84575091610eb6989991846001959410610f6b575b505050811b019055610e63565b015160001960f88460031b161c19169055888080610f5e565b838301518b556001909a019960209384019301610f44565b634e487b7160e01b600052604160045260246000fd5b015190508a80610cdb565b90601f1983169184600052816000209260005b818110611010575091600193918560069897969410610ff7575b505050811b019055610e00565b015160001960f88460031b161c19169055898080610fea565b92936020600181928786015181550195019301610fd0565b90601f1983169184600052816000209260005b818110611077575090846001959493921061105e575b505050811b019055610da3565b015160001960f88460031b161c19169055898080611051565b9293602060018192878601518155019501930161103b565b015190508b80610cdb565b90601f1983169184600052816000209260005b8181106110e957509084600195949392106110d0575b505050811b019055610d4b565b015160001960f88460031b161c191690558a80806110c3565b929360206001819287860151815501950193016110ad565b015190508c80610cdb565b90601f1983169184600052816000209260005b81811061115b5750908460019594939210611142575b505050811b019055610cf3565b015160001960f88460031b161c191690558b8080611135565b9293602060018192878601518155019501930161111f565b61118a611183368d60e0516132b3565b8a33614826565b6111968c8a338d614baf565b610b17565b9a999897969594939291908a600052600f60205260406000206080525b6080515480156112165780600019810111611200576111e06111fb9160001901608051614c53565b90549060031b1c8060005260086020526040600020906145b4565b6111b8565b634e487b7160e01b600052601160045260246000fd5b50909192939495969798999a610b03565b9294969850969a94509791503303611594578454600052600a60205260006040812055604435600052600a60205260406000205583549560443585556003850191604051611280816112798187613206565b0382612eb5565b602081519101206112923684846132b3565b60208151910120036114cd575b50505060048301916040516112b8816112798187613206565b602081519101206112ca3684846132b3565b60208151910120036113fd575b5050506006019067ffffffffffffffff8111610f9c576112fb81610e2684546131cc565b6000601f821160011461139957819061132d9394959660009261138e5750508160011b916000199060031b1c19161790565b90555b6113486040519260408452604084019060e0516134bc565b9082820360208401527f03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e3223393806113896044359561010051610120516134bc565b0390a4005b013590508680610cdb565b601f198216958382526020822091805b8881106113e5575083600195969798106113cb575b505050811b019055611330565b0135600019600384901b60f8161c191690558580806113be565b909260206001819286860135815501940191016113a9565b67ffffffffffffffff8211610f9c5761141a82610cb985546131cc565b600090601f831160011461146557918061144f92600696959460009261145a5750508160011b916000199060031b1c19161790565b90555b9086806112d7565b013590508a80610cdb565b8382526020822091601f198416815b8181106114b557509160019391856006989796941061149b575b505050811b019055611452565b0135600019600384901b60f8161c1916905589808061148e565b91936020600181928787013581550195019201611474565b67ffffffffffffffff8211610f9c576114ea82610cb985546131cc565b600090601f83116001146115305761151a9291600091836115255750508160011b916000199060031b1c19161790565b90555b87808061129f565b013590508b80610cdb565b8382526020822091601f198416815b81811061157c5750908460019594939210611562575b505050811b01905561151d565b0135600019600384901b60f8161c191690558a8080611555565b9193602060018192878701358155019501920161153f565b6331ee6dc760e01b6000523360045260246000fd5b7f36a7c503000000000000000000000000000000000000000000000000000000006000526101005160045260245260446000fd5b50806101005111610a7d565b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b847f436f97540000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808511610a5b565b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b857f354f25140000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808611610a39565b8089116116d757808b1115610a20578a6219aad560e31b60005260045260245260446000fd5b886219aad560e31b60005260045260245260446000fd5b7f9cd963cf0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260443560045260246000fd5b7f315de7450000000000000000000000000000000000000000000000000000000060005260046000fd5b346102f9576103cf61178336613137565b91613f9a565b346102f95760203660031901126102f95760043567ffffffffffffffff81116102f9576117ba9036906004016130bc565b80156106dd5760005b8181106117cc57005b806117da6001928486613509565b356117e3614551565b600052600a602052604060002054806000526008602052604060002060ff8482015460e01c1661181281612f40565b15611820575b5050016117c3565b611829916145b4565b8480611818565b346102f95760603660031901126102f957611849612e2e565b60243567ffffffffffffffff81116102f957611869903690600401612f12565b916044359267ffffffffffffffff84116102f9576118b36118bb936108186118986118cf973690600401612f12565b906118a16132fa565b50604051958694602086019889613f65565b519020614411565b604051918291602083526020830190612f4a565b0390f35b346102f95760003660031901126102f9576020600454604051908152f35b346102f95760803660031901126102f95760043560ff81168091036102f9576024359060ff8216908183036102f95760443560ff81168082036102f9576064359361ffff8516928386036102f9577fa254b4f533572ce0e775da29b44e0114b25637b1072bbeaf3fe051fb55aafedc9661ff0062ff000064ffff000000608099611979614551565b60181b169360101b169160081b1686171717640400c82040811480156119c6575b156119be575060006002555b604051938452602084015260408301526060820152a1005b6002556119a6565b50801561199a565b346102f95760803660031901126102f9576119e7612e2e565b6064359067ffffffffffffffff82116102f957611a0b6103cf923690600401612f12565b916044359060243590613d54565b346102f95760403660031901126102f95760043563ffffffff611a3a6130ed565b1642811115611b9257611a5a336000526005602052604060002054151590565b156103d157604080513360208201908152918101849052611a7e8160608101610818565b519020600052601260205260406000208163ffffffff19825416179055604051611aa781612e60565b8281526020810190338252604081019083825260135468010000000000000000811015610f9c57806001611ade9201601355613398565b611b7c5760016001600160a01b039291839251815501935116166001600160a01b0319835416178255517fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff000000000000000000000000000000000000000083549260a01b1691161790556040519081527ff69135e4f80a25991d2f877c365c191c51ec3c0063ecb9299d314cd9da4880d160203392a3005b634e487b7160e01b600052600060045260246000fd5b907fa9a6ead20000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b346102f95760003660031901126102f95760206001600160a01b0360015416604051908152f35b346102f95760603660031901126102f957611c03612e2e565b6001600160a01b0360243591169081600052600e602052611c2c6044358260406000205461438f565b90611c3682613348565b9260005b838110611c4f57604051806118cf8782613015565b60019082600052600e602052611c82611c766040600020611c7084886131bf565b90614c53565b90549060031b1c614411565b611c8c8288613289565b52611c978187613289565b5001611c3a565b346102f9576001600160a01b03611cb436613167565b929091611cbf614551565b169081600052601460205260406000209267ffffffffffffffff8111610f9c57611cf381611ced86546131cc565b8661374b565b600093601f8211600114611d6f57611d4582807f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f139697600091611d64575b508160011b916000199060031b1c19161790565b90555b611d5f6040519283926020845260208401916134bc565b0390a2005b905085013588611d31565b80855260208520601f19831695805b878110611dd95750837f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f13969710611dbf575b5050600182811b019055611d48565b840135600019600385901b60f8161c191690558580611db0565b90916020600181928589013581550193019101611d7e565b346102f95760003660031901126102f957604060195467ffffffffffffffff8251916001600160a01b038116835260a01c166020820152f35b346102f95760003660031901126102f9576000546001600160a01b0381163303611e99576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102f95760203660031901126102f957611edc6132fa565b50600435600052600a6020526118cf6118bb604060002054614411565b346102f95760003660031901126102f957602060ff611f1661453f565b16604051908152f35b346102f95760203660031901126102f957611f47336000526005602052604060002054151590565b156103d157600435600052600a6020526103cf604060002054611f6a8133614796565b90614a18565b346102f957611fad611f90611f8436613100565b949290939136916132b3565b602081519101209283600052600d6020528260406000205461438f565b90611fb782613348565b9260005b838110611fd057604051806118cf8782613015565b60019082600052600d602052611ff1611c766040600020611c7084886131bf565b611ffb8288613289565b526120068187613289565b5001611fbb565b346102f95760203660031901126102f95760043567ffffffffffffffff81116102f957612041612049913690600401612f12565b6107a1614551565b60208151910120600052600c60205260406000205b805480156103cf576000198101908111611200576111e061207f9183614c53565b61205e565b346102f95761209236613137565b90916120ab336000526005602052604060002054151590565b156103d157600052600a6020526040600020546120c88133614796565b9060ff600183015460e01c166120dd81612f40565b6120e357005b61210e6121076103cf956120f83687836132b3565b602081519101209536916132b3565b8433614826565b61491d565b346102f95761212f61212436613100565b9290939136916132b3565b6020815191012080600052601860205261215060406000209284845461438f565b9161215a836131a7565b936121686040519586612eb5565b838552601f19612177856131a7565b0160005b81811061229057505060005b8481106121ef578560405180916020820160208352815180915260206040840192019060005b8181106121bb575050500390f35b825180516001600160a01b0316855260209081015163ffffffff1681860152869550604090940193909201916001016121ad565b806001600160a01b0361220d612207600194866131bf565b86614c53565b90549060031b1c168560005260156020526040600281600020016000906001600160a01b0384168252602052206040519061224782612e7c565b5490602060ff63ffffffff841693848452821c1615159101526040519161226d83612e7c565b8252602082015261227e8289613289565b526122898188613289565b5001612187565b60209060405161229f81612e7c565b6000815260008382015282828a0101520161217b565b346102f95760203660031901126102f95760043567ffffffffffffffff81116102f9576107a16122e9913690600401612f12565b60208151910120600052600c6020526020604060002054604051908152f35b346102f95760203660031901126102f957612321614551565b600435600052600a602052604060002054806000526008602052604060002060ff600182015460e01c1661235481612f40565b156103c657005b346102f95761237261236c36612dd7565b90613bff565b60405180916020820160208352815180915260206040840192019060005b81811061239e575050500390f35b919350916020606060019263ffffffff60408851805184526001600160a01b038682015116868501520151166040820152019401910191849392612390565b346102f9576123eb36613075565b92829192421161251d576001600160a01b03811693612417856000526005602052604060002054151590565b156125085761242585614d11565b936001600160a01b0361243b83838888886146b5565b16600052600360205260ff60406000205416156124e757858581600052600e6020526040600020905b815480156124a5576000198101908111611200576124856124a09184614c53565b90549060031b1c806000526008602052604060002090614a18565b612464565b6000828580835260066020528260408120556124c081614edd565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b60405163335d4ce160e01b815294859461250494600487016134dd565b0390fd5b8463c2dda3f960e01b60005260045260246000fd5b6001600160a01b0390631ec5288b60e11b600052166004524260245260445260646000fd5b346102f95760603660031901126102f95760043567ffffffffffffffff81116102f957612573903690600401612f12565b61257b6130ed565b6044359182151583036102f9576103cf93612594614551565b613931565b346102f95760403660031901126102f95760043567ffffffffffffffff81116102f9576125ca9036906004016130bc565b6024359167ffffffffffffffff83116102f9576125ee6103cf933690600401612f12565b929091613519565b346102f95760203660031901126102f9576001600160a01b03612617612e2e565b1660005260146020526118cf61127961263a604060002060405192838092613206565b604051918291602083526020830190612ded565b346102f95761265c36613075565b82939242116126f4576001600160a01b038316612686816000526005602052604060002054151590565b156126e05761269490614d11565b916001600160a01b036126aa83838689896146b5565b16600052600360205260ff60406000205416156126c357005b6125049260405195869563335d4ce160e01b8752600487016134dd565b63c2dda3f960e01b60005260045260246000fd5b836001600160a01b0384631ec5288b60e11b600052166004524260245260445260646000fd5b346102f95760203660031901126102f9576001600160a01b0361273b612e2e565b612743614551565b16600052600b60205260406000205b805480156103cf576000198101908111611200576111e06127739183614c53565b612752565b346102f95760403660031901126102f957612791612e2e565b60243567ffffffffffffffff811681036102f9576103cf916127b1614551565b6133b7565b346102f95760003660031901126102f957602060ff6127d361453f565b60101c16604051908152f35b346102f9576127fc6127f036612dd7565b81601a9392935461438f565b90612806826131a7565b916128146040519384612eb5565b808352601f19612823826131a7565b0160005b8181106129b6575050601a54909160005b8381106128e4578460405160208101916020825280518093526040820192602060408260051b85010192019060005b8181106128745784840385f35b909192603f1985820301865283519081519160038310156128ce576128bf82606060406020959460019787965263ffffffff8682015116868501520151918160408201520190612ded565b95019601910194919094612867565b634e487b7160e01b600052602160045260246000fd5b6128ee81836131bf565b6000848210156129a257601a90526040519060011b7f057c384a7d1c54f3a1b2e5e67b2617b8224fdfd1ea7234eea573a6ff665ff63e01600061293083612e60565b81549060ff821690600382101561298e5750835260081c63ffffffff166020830152604051600193929161296d9082906112799082908801613206565b604082015261297c8288613289565b526129878187613289565b5001612838565b80634e487b7160e01b602492526021600452fd5b80634e487b7160e01b602492526032600452fd5b6020906040516129c581612e60565b600081526000838201526060604082015282828801015201612827565b346102f95760803660031901126102f9576129fb612e2e565b60243567ffffffffffffffff81116102f957612a2f91612a22612a35923690600401612f12565b93906044359436916132b3565b906143f2565b90816000526009602052612a516064358260406000205461438f565b90612a5b82613348565b9260005b838110612a7457604051806118cf8782613015565b600190826000526009602052612a95611c766040600020611c7084886131bf565b612a9f8288613289565b52612aaa8187613289565b5001612a5f565b346102f95760003660031901126102f9576118cf61263a612ed7565b346102f957612aea612ade36612dd7565b8160049392935461438f565b612af3816131a7565b91612b016040519384612eb5565b818352612b0d826131a7565b602084019290601f19013684376004549160005b828110612b765784866040519182916020830190602084525180915260408301919060005b818110612b54575050500390f35b82516001600160a01b0316845285945060209384019390920191600101612b46565b612b8081836131bf565b6000858210156129a257600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b015460008190526006602052600191906001600160a01b0316612bd28289613289565b5201612b21565b346102f95760203660031901126102f9576020612c136001600160a01b03612bff612e2e565b166000526005602052604060002054151590565b6040519015158152f35b346102f957612c3a612c2e36612dd7565b8160169392935461438f565b612c43816131a7565b91612c516040519384612eb5565b818352601f19612c60836131a7565b0160005b818110612da4575050601654919060005b828110612d1357846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210612cb557505050500390f35b919360019193955060208091603f1989820301855287519081518152606080612ceb858501516080878601526080850190612ded565b9363ffffffff6040820151166040850152015115159101529601920192018594939192612ca6565b612d1d81836131bf565b6000858210156129a2579060208260166001959452200160ff60406000925492838152601560205220611279612d6c868301549260405195612d5e87612e44565b865260405192838092613206565b602084015263ffffffff8116604084015260201c1615156060820152612d928288613289565b52612d9d8187613289565b5001612c75565b602090604051612db381612e44565b60008152606083820152600060408201526000606082015282828801015201612c64565b60409060031901126102f9576004359060243590565b919082519283825260005b848110612e19575050826000602080949584010152601f8019910116010190565b80602080928401015182828601015201612df8565b600435906001600160a01b03821682036102f957565b6080810190811067ffffffffffffffff821117610f9c57604052565b6060810190811067ffffffffffffffff821117610f9c57604052565b6040810190811067ffffffffffffffff821117610f9c57604052565b610140810190811067ffffffffffffffff821117610f9c57604052565b90601f8019910116810190811067ffffffffffffffff821117610f9c57604052565b60405190612ee6604083612eb5565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b9181601f840112156102f95782359167ffffffffffffffff83116102f957602083818601950101116102f957565b600211156128ce57565b61301291815181526001600160a01b03602083015116602082015267ffffffffffffffff60408301511660408201526060820151612f8781612f40565b6060820152610120613000612fec612fda612fc8612fb660808801516101406080890152610140880190612ded565b60a088015187820360a0890152612ded565b60c087015186820360c0880152612ded565b60e086015185820360e0870152612ded565b610100850151848203610100860152612ded565b92015190610120818403910152612ded565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061304857505050505090565b9091929394602080613066600193603f198682030187528951612f4a565b97019301930191939290613039565b60606003198201126102f9576004356001600160a01b03811681036102f95791602435916044359067ffffffffffffffff82116102f9576130b891600401612f12565b9091565b9181601f840112156102f95782359167ffffffffffffffff83116102f9576020808501948460051b0101116102f957565b6024359063ffffffff821682036102f957565b60606003198201126102f9576004359067ffffffffffffffff82116102f95761312b91600401612f12565b90916024359060443590565b9060406003198301126102f957600435916024359067ffffffffffffffff82116102f9576130b891600401612f12565b9060406003198301126102f9576004356001600160a01b03811681036102f957916024359067ffffffffffffffff82116102f9576130b891600401612f12565b67ffffffffffffffff8111610f9c5760051b60200190565b9190820180921161120057565b90600182811c921680156131fc575b60208310146131e657565b634e487b7160e01b600052602260045260246000fd5b91607f16916131db565b60009291815491613216836131cc565b808352926001811690811561326c575060011461323257505050565b60009081526020812093945091925b838310613252575060209250010190565b600181602092949394548385870101520191019190613241565b915050602093945060ff929192191683830152151560051b010190565b805182101561329d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b92919267ffffffffffffffff8211610f9c57604051916132dd601f8201601f191660200184612eb5565b8294818452818301116102f9578281602093846000960137010152565b6040519061330782612e98565b6060610120836000815260006020820152600060408201526000838201528260808201528260a08201528260c08201528260e0820152826101008201520152565b90613352826131a7565b61335f6040519182612eb5565b8281528092613370601f19916131a7565b019060005b82811061338157505050565b60209061338c6132fa565b82828501015201613375565b60135481101561329d57601360005260206000209060011b0190600090565b90601954906001600160a01b038216906001600160a01b0367ffffffffffffffff8460a01c16941693828514948580966134a9575b6134a157806080957fc0c3ee74e6d6070ee9c493e8b4f0477d2e66600f22997a4e073288d38d65933b971561348c575b505067ffffffffffffffff83169282840361344b575b50604051938452602084015260408301526060820152a1565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b6019549260a01b1691161760195538613432565b6001600160a01b03191617601955803861341c565b505050505050565b508167ffffffffffffffff8416146133ec565b908060209392818452848401376000828201840152601f01601f1916010190565b90926080926001600160a01b0361301297951683526020830152604082015281606082015201916134bc565b919081101561329d5760051b0190565b929181156106dd57613538336000526005602052604060002054151590565b156103d1576135483684836132b3565b6020815191012092600091825b8481106136d1575063ffffffff8316156134a1576135749136916132b3565b90836000526015602052604060002060018101549060ff8260201c16156136af57600201906001600160a01b0333166000528160205260ff60406000205460201c1660001461369b57506001600160a01b03331660005260205263ffffffff806136038160406000205416935b336000526010602052604060002088600052602052826040600020541661480c565b921691161161367d575060005b81811061361d5750505050565b8061362b6001928487613509565b35600052600a6020526040600020548060005260086020528460406000208460ff8183015460e01c1661365d81612f40565b1461366c575b50505001613610565b6136759261491d565b388481613663565b6125049060405191829163038857ff60e01b835233600484016143d2565b63ffffffff915061360382809216936135e1565b60405163393f328760e11b815260206004820152806125046024820187612ded565b926136dd848689613509565b35600052600a60205260ff60016136f960406000205433614796565b015460e01c1661370881612f40565b1561372b5763ffffffff1663ffffffff8114611200576001809101935b01613555565b92600190613725565b81811061373f575050565b60008155600101613734565b9190601f811161375a57505050565b613786926000526020600020906020601f840160051c83019310613788575b601f0160051c0190613734565b565b9091508190613779565b90601a5468010000000000000000811015610f9c5760018101601a556000601a548210156129a257601a90526000929060011b7f057c384a7d1c54f3a1b2e5e67b2617b8224fdfd1ea7234eea573a6ff665ff63e01815160038110156138f85760409160019160ff825491168060ff19831617835564ffffffff00602087015160081b169164ffffffffff191617178155019101519283519067ffffffffffffffff82116138e45761384882610cb985546131cc565b602090601f831160011461388657906138779394958361387b5750508160011b916000199060031b1c19161790565b9055565b015190503880610cdb565b90601f198316848352818320925b8181106138cc575095836001959697106138b3575b505050811b019055565b015160001960f88460031b161c191690553880806138a9565b9192602060018192868b015181550194019201613894565b80634e487b7160e01b602492526041600452fd5b602485634e487b7160e01b81526021600452fd5b9161392a60209263ffffffff929695966040865260408601916134bc565b9416910152565b90919392936139413684846132b3565b6020815191012090600086600014613b925750905b80600052601560205260406000209560018701908154901515908160ff8260201c161515149081613b7f575b50613b75578754613992816131cc565b15613a43575b50815464ffffffffff1916602091821b64ff00000000161763ffffffff85169081179092556040805191820193909352808301919091529081529394507f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b73193613a3e9190613a3290613a0b606082612eb5565b60405190613a1882612e60565b6000825263ffffffff421660208301526040820152613792565b6040519384938461390c565b0390a1565b67ffffffffffffffff8711610f9c5786613a5f613a65926131cc565b8a61374b565b600097601f8711600114613ae65791613a3e949391613a3293613ac189807f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b7319c9d600091613adb57508160011b916000199060031b1c19161790565b90555b613acd83614e27565b509193945091889750613998565b90508a013538611d31565b8089526020892098601f198816815b818110613b5d5750917f4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b731999a613a3e979694928a613a32979510613b43575b5050600189811b019055613ac4565b89013560001960038c901b60f8161c191690553880613b34565b888301358c556001909b019a60209283019201613af5565b5050505050509050565b905063ffffffff80861691161438613982565b905090613956565b90613ba4826131a7565b613bb16040519182612eb5565b8281528092613bc2601f19916131a7565b019060005b828110613bd357505050565b602090604051613be281612e60565b600081526000838201526000604082015282828501015201613bc7565b90613c0d908260135461438f565b908115613d0457613c1d82613b9a565b91600091825b828110613c7a5750508110613c36575090565b613c3f81613b9a565b9160005b828110613c505750505090565b80613c5d60019284613289565b51613c688287613289565b52613c738186613289565b5001613c43565b613c8c613c8782846131bf565b613398565b50600181015463ffffffff8160a01c1690428211613cb0575b505050600101613c23565b6001600160a01b039060409794975193613cc985612e60565b5484521660208301526040820152613ce18287613289565b52613cec8186613289565b50600019811461120057600180910193903880613ca5565b50506040516000613d16602083612eb5565b81526000805b818110613d2857505090565b602090604051613d3781612e60565b600081526000838201526000604082015282828601015201613d1c565b92909391844211613f26576001600160a01b038416613d80816000526005602052604060002054151590565b613ef95781600052600760205260ff60406000205416613ec957600090613ddc613df6613dab612ed7565b6040519283916020830195878752604084015246606084015230608084015260e060a0840152610100830190612ded565b8a60c08301528660e083015203601f198101835282612eb5565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613e37603c8220613e313687876132b3565b90614d74565b909192600483101561298e5782613e91575050506001600160a01b031660009081526003602052604090205460ff1615613e72575050505050565b90612504929160405195869563335d4ce160e01b8752600487016134dd565b5060405163d36ab6b960e01b81526060600482015291829160ff613eb9606485018a8a6134bc565b9216602484015260448301520390fd5b7f77a338580000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7fd9a5f5ca0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b846001600160a01b03857f502d038700000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b939161301295936001600160a01b03613f8c931686526060602087015260608601916134bc565b9260408185039101526134bc565b9091613fb3336000526005602052604060002054151590565b156103d15781600052600a60205260406000205490613fd28233614796565b60ff600182015460e01c16613fe681612f40565b6140b25761401b9083600052601160205260406000205492836000526015602052614022604060002060405194858092613206565b0384612eb5565b61402d3682896132b3565b602081519101208094146140a9576140806140949483876140726140a4967f9b5361a5258ef6ac8039fd2d2ac276734695219cfd870711e7922c236e5db16d9a6145b4565b61210e61210736878e6132b3565b604051938493604085526040850190612ded565b90838203602085015233976134bc565b0390a3565b50505050505050565b7fd74915a80000000000000000000000000000000000000000000000000000000060005260046000fd5b926140e83684846132b3565b60208151910120948560005260156020526040600020600181019060ff825460201c16156142df57600201604060006001600160a01b038916815282602052209260001461425157505463ffffffff8381169116811161422757815460ff8160201c16156000146141c75750815464ffffffffff1916176401000000001790556000948552601860205260409094206001600160a01b03909316937f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c82324936141c291906141b5908790614e81565b506040519384938461390c565b0390a2565b819392959694975063ffffffff1614156000146134a1577f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c82324946001600160a01b03946141c29363ffffffff1982541617905560405194859416968461390c565b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b959692505060ff909392935460201c16156142d8577f6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d849936001600160a01b0386166000526020526000604081205560005260186020526142c06001600160a01b03604060002095168095614fa7565b506141c26040519283926020845260208401916134bc565b5050505050565b60405163393f328760e11b815260206004820152806125046024820189896134bc565b919061430f9136916132b3565b6020815191012060005260156020526001600160a01b036040600020911660005260028101602052604060002060206040519161434b83612e7c565b549160ff63ffffffff841693848352831c161515918291015261437657506001015463ffffffff1690565b905090565b356001600160a01b03811681036102f95790565b9091818310156143ca57816143a482856131bf565b11156143b95750905b81039081116112005790565b6143c49150826131bf565b906143ad565b505050600090565b6040906001600160a01b0361301294931681528160208201520190612ded565b9061440b610818916040519283916020830195866143d2565b51902090565b6144196132fa565b81600052600860205260406000206001810154926001600160a01b038416928315614537575091600661451c836112799567ffffffffffffffff61452e96549860ff8160e01c1692600052601160205260406000205460005260156020526040600020966040519a61448a8c612e98565b8b5260208b015260a01c1660408901526144a381612f40565b60608801526040516144bc816112798160028601613206565b60808801526040516144d5816112798160038601613206565b60a08801526040516144ee816112798160048601613206565b60c0880152604051614507816112798160058601613206565b60e08801526112796040518094819301613206565b61010085015260405192838092613206565b61012082015290565b935050505090565b600254806130125750640400c8204090565b6001600160a01b0360015416330361456557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90916145a661301293604084526040840190613206565b916020818403910152613206565b7ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e48868960018301916001600160a01b038354947c010000000000000000000000000000000000000000000000000000000060ff60e01b19871617855582600052601160205261464a8260406000205497169387600284019561464460405161463e81611279818c613206565b826143f2565b92614c6b565b61469481546040519088602083015260408201526040815261466d606082612eb5565b6040519061467a82612e60565b6002825263ffffffff421660208301526040820152613792565b5493541693600052601560205260406000206140a46040519283928361458f565b9161471b90614702926146c6612ed7565b916040519485936001600160a01b03602086019860018a5216604086015246606086015230608086015260e060a0860152610100850190612ded565b9160c084015260e083015203601f198101835282612eb5565b5190207f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52614758603c600020613e313685856132b3565b60048295939510156128ce5781614770575050505090565b60ff613eb960405195869563d36ab6b960e01b87526060600488015260648701916134bc565b9060005260086020526040600020906001600160a01b036001830154169081156147e2576001600160a01b03168091036147ce575090565b6331ee6dc760e01b60005260045260246000fd5b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b9063ffffffff8091169116019063ffffffff821161120057565b90806000526015602052604060002060018101549060ff8260201c16156148fb57600201906001600160a01b0384166000528160205260ff60406000205460201c166000146148ed57506001600160a01b03831660005260205263ffffffff60406000205416905b6001600160a01b038316600052601060205260406000209060005260205263ffffffff806148c46001826040600020541661480c565b92169116116148d1575050565b61250460405192839263038857ff60e01b8452600484016143d2565b63ffffffff9150169061488e565b60405163393f328760e11b815260206004820152806125046024820188612ded565b7f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a906001600160a01b036001840193614974828654169387600284019561496e60405161463e81611279818c613206565b92614baf565b60ff60e01b1985541685556146948154604051908860208301526040820152604081526149a2606082612eb5565b604051906149af82612e60565b6001825263ffffffff421660208301526040820152613792565b6149d381546131cc565b90816149dd575050565b81601f600093116001146149ef575055565b81835260208320614a0b91601f0160051c810190600101613734565b8082528160208120915555565b9061463e907f48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd94614b37614aa6614b456001600160a01b036001860195865460ff8a848316614a74611279600287019d8e60405192838092613206565b938491836000526011602052604060002054948592836000526015602052614aad60406000206040519d8e8092613206565b038d612eb5565b60e01c16614aba81612f40565b15614b9d575b50505050600052600d602052614ada8a6040600020614fa7565b5082885416600052600e602052614af58a6040600020614fa7565b506000526009602052614b0c896040600020614fa7565b508054600052600a602052600060408120555495541695604051938493604085526040850190612ded565b908382036020850152613206565b0390a36000526008602052600660406000206000815560006001820155614b6e600282016149c9565b614b7a600382016149c9565b614b86600482016149c9565b614b92600582016149c9565b016149d381546131cc565b614ba693614c6b565b8a828238614ac0565b9192906001600160a01b03168060005260106020526040600020846000526020526040600020600163ffffffff8254160163ffffffff81116112005763ffffffff1663ffffffff1982541617905583600052600c602052614c14836040600020614e81565b50600052600b602052614c2b826040600020614e81565b50600052600f602052614c42816040600020614e81565b506000526011602052604060002055565b805482101561329d5760005260206000200190600090565b9291906001600160a01b031680600052600b602052614c8e846040600020614fa7565b5081600052600c602052614ca6846040600020614fa7565b506000526010602052604060002090600052602052604060002060001963ffffffff8254160163ffffffff81116112005763ffffffff1663ffffffff19825416179055600052600f602052614cff816040600020614fa7565b50600052601160205260006040812055565b80600052600660205260406000205490811580614d5e575b614d31575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5080600052600560205260406000205415614d29565b8151919060418303614da557614d9e92506020820151906060604084015193015160001a90615064565b9192909190565b505060009160029190565b80600052600560205260406000205415600014614e215760045468010000000000000000811015610f9c57614e08614df18260018594016004556004614c53565b819391549060031b91821b91600019901b19161790565b9055600454906000526005602052604060002055600190565b50600090565b80600052601760205260406000205415600014614e215760165468010000000000000000811015610f9c57614e68614df18260018594016016556016614c53565b9055601654906000526017602052604060002055600190565b6000828152600182016020526040902054614ed65780549068010000000000000000821015610f9c5782614ebf614df1846001809601855584614c53565b905580549260005201602052604060002055600190565b5050600090565b6000818152600560205260409020548015614ed65760001981018181116112005760045460001981019190821161120057818103614f6d575b5050506004548015614f575760001901614f31816004614c53565b8154906000199060031b1b19169055600455600052600560205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b614f8f614f7e614df1936004614c53565b90549060031b1c9283926004614c53565b90556000526005602052604060002055388080614f16565b906001820191816000528260205260406000205480151560001461505b57600019810181811161120057825460001981019190821161120057818103615024575b50505080548015614f575760001901906150028282614c53565b8154906000199060031b1b191690555560005260205260006040812055600190565b615044615034614df19386614c53565b90549060031b1c92839286614c53565b905560005283602052604060002055388080614fe8565b50505050600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116150ed579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa156150e1576000516001600160a01b038116156150d55790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b5050506000916003919056fea164736f6c634300081a000a",
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

func (_WorkflowRegistry *WorkflowRegistryCaller) CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canUnlinkOwner", owner, validityTimestamp, signature)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetAllowlistedRequests(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getAllowlistedRequests", start, limit)

	if err != nil {
		return *new([]WorkflowRegistryOwnerAllowlistedRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryOwnerAllowlistedRequest)).(*[]WorkflowRegistryOwnerAllowlistedRequest)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetAllowlistedRequests(start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error) {
	return _WorkflowRegistry.Contract.GetAllowlistedRequests(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetAllowlistedRequests(start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error) {
	return _WorkflowRegistry.Contract.GetAllowlistedRequests(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetCapabilitiesRegistry(opts *bind.CallOpts) (common.Address, uint64, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getCapabilitiesRegistry")

	if err != nil {
		return *new(common.Address), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetCapabilitiesRegistry() (common.Address, uint64, error) {
	return _WorkflowRegistry.Contract.GetCapabilitiesRegistry(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetCapabilitiesRegistry() (common.Address, uint64, error) {
	return _WorkflowRegistry.Contract.GetCapabilitiesRegistry(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetDonConfigs(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getDonConfigs", start, limit)

	if err != nil {
		return *new([]WorkflowRegistryDonConfigView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryDonConfigView)).(*[]WorkflowRegistryDonConfigView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetDonConfigs(start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error) {
	return _WorkflowRegistry.Contract.GetDonConfigs(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetDonConfigs(start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error) {
	return _WorkflowRegistry.Contract.GetDonConfigs(&_WorkflowRegistry.CallOpts, start, limit)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetUserDONOverrides(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getUserDONOverrides", donFamily, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryUserOverrideView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryUserOverrideView)).(*[]WorkflowRegistryUserOverrideView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetUserDONOverrides(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error) {
	return _WorkflowRegistry.Contract.GetUserDONOverrides(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetUserDONOverrides(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error) {
	return _WorkflowRegistry.Contract.GetUserDONOverrides(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowListByDON(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowListByDON", donFamily, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowListByDON(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowListByDON(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
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

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalAllowlistedRequests(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalAllowlistedRequests")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalAllowlistedRequests() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalAllowlistedRequests(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalAllowlistedRequests() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalAllowlistedRequests(&_WorkflowRegistry.CallOpts)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) AllowlistRequest(opts *bind.TransactOpts, requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "allowlistRequest", requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AllowlistRequest(requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AllowlistRequest(&_WorkflowRegistry.TransactOpts, requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AllowlistRequest(requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error) {
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetCapabilitiesRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setCapabilitiesRegistry", registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetCapabilitiesRegistry(registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetCapabilitiesRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetCapabilitiesRegistry(registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetCapabilitiesRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "unlinkOwner", owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature)
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

type WorkflowRegistryCapabilitiesRegistryUpdatedIterator struct {
	Event *WorkflowRegistryCapabilitiesRegistryUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryCapabilitiesRegistryUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryCapabilitiesRegistryUpdated)
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
		it.Event = new(WorkflowRegistryCapabilitiesRegistryUpdated)
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

func (it *WorkflowRegistryCapabilitiesRegistryUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryCapabilitiesRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryCapabilitiesRegistryUpdated struct {
	OldAddr          common.Address
	NewAddr          common.Address
	OldChainSelector uint64
	NewChainSelector uint64
	Raw              types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterCapabilitiesRegistryUpdated(opts *bind.FilterOpts) (*WorkflowRegistryCapabilitiesRegistryUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "CapabilitiesRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryCapabilitiesRegistryUpdatedIterator{contract: _WorkflowRegistry.contract, event: "CapabilitiesRegistryUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchCapabilitiesRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryCapabilitiesRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "CapabilitiesRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryCapabilitiesRegistryUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "CapabilitiesRegistryUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseCapabilitiesRegistryUpdated(log types.Log) (*WorkflowRegistryCapabilitiesRegistryUpdated, error) {
	event := new(WorkflowRegistryCapabilitiesRegistryUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "CapabilitiesRegistryUpdated", log); err != nil {
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
	ExpiryTimestamp uint32
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
	case _WorkflowRegistry.abi.Events["CapabilitiesRegistryUpdated"].ID:
		return _WorkflowRegistry.ParseCapabilitiesRegistryUpdated(log)
	case _WorkflowRegistry.abi.Events["DONLimitSet"].ID:
		return _WorkflowRegistry.ParseDONLimitSet(log)
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

func (WorkflowRegistryCapabilitiesRegistryUpdated) Topic() common.Hash {
	return common.HexToHash("0xc0c3ee74e6d6070ee9c493e8b4f0477d2e66600f22997a4e073288d38d65933b")
}

func (WorkflowRegistryDONLimitSet) Topic() common.Hash {
	return common.HexToHash("0x4495c6e60e22d8ab3b00a16160c4954fb24d2d51d48044a0d292dddc3161b731")
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
	return common.HexToHash("0xf69135e4f80a25991d2f877c365c191c51ec3c0063ecb9299d314cd9da4880d1")
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

	CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) error

	GetAllowlistedRequests(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error)

	GetCapabilitiesRegistry(opts *bind.CallOpts) (common.Address, uint64, error)

	GetDonConfigs(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error)

	GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetMaxAttrLen(opts *bind.CallOpts) (uint16, error)

	GetMaxNameLen(opts *bind.CallOpts) (uint8, error)

	GetMaxTagLen(opts *bind.CallOpts) (uint8, error)

	GetMaxUrlLen(opts *bind.CallOpts) (uint8, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (uint32, error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donFamily string) (uint32, error)

	GetUserDONOverrides(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error)

	GetWorkflow(opts *bind.CallOpts, owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowById(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByDON(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByOwnerAndName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowOwnerConfig(opts *bind.CallOpts, owner common.Address) ([]byte, error)

	IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error)

	IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error)

	IsRequestAllowlisted(opts *bind.CallOpts, owner common.Address, requestDigest [32]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TotalActiveWorkflowsByOwner(opts *bind.CallOpts, owner common.Address) (*big.Int, error)

	TotalActiveWorkflowsOnDON(opts *bind.CallOpts, donFamily string) (*big.Int, error)

	TotalAllowlistedRequests(opts *bind.CallOpts) (*big.Int, error)

	TotalEvents(opts *bind.CallOpts) (*big.Int, error)

	TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte, donFamily string) (*types.Transaction, error)

	AdminBatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	AdminPauseAllByDON(opts *bind.TransactOpts, donFamily string) (*types.Transaction, error)

	AdminPauseAllByOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error)

	AdminPauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	AllowlistRequest(opts *bind.TransactOpts, requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error)

	BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte, donFamily string) (*types.Transaction, error)

	BatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	DeleteWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error)

	PauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	SetCapabilitiesRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector uint64) (*types.Transaction, error)

	SetDONLimit(opts *bind.TransactOpts, donFamily string, limit uint32, enabled bool) (*types.Transaction, error)

	SetMetadataConfig(opts *bind.TransactOpts, nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, limit uint32, enabled bool) (*types.Transaction, error)

	SetWorkflowOwnerConfig(opts *bind.TransactOpts, owner common.Address, config []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	UpdateWorkflowDONFamily(opts *bind.TransactOpts, workflowId [32]byte, newDonFamily string) (*types.Transaction, error)

	UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error)

	FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error)

	WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error)

	ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error)

	FilterCapabilitiesRegistryUpdated(opts *bind.FilterOpts) (*WorkflowRegistryCapabilitiesRegistryUpdatedIterator, error)

	WatchCapabilitiesRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryCapabilitiesRegistryUpdated) (event.Subscription, error)

	ParseCapabilitiesRegistryUpdated(log types.Log) (*WorkflowRegistryCapabilitiesRegistryUpdated, error)

	FilterDONLimitSet(opts *bind.FilterOpts) (*WorkflowRegistryDONLimitSetIterator, error)

	WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet) (event.Subscription, error)

	ParseDONLimitSet(log types.Log) (*WorkflowRegistryDONLimitSet, error)

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
