// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package workflow_registry_wrapper_v2_dev

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

type WorkflowRegistryConfig struct {
	MaxNameLen   uint8
	MaxTagLen    uint8
	MaxUrlLen    uint8
	MaxAttrLen   uint16
	MaxExpiryLen uint32
}

type WorkflowRegistryDonConfigView struct {
	DonHash          [32]byte
	Family           string
	DonLimit         uint32
	DefaultUserLimit uint32
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
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlistRequest\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowlistedRequests\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"allowlistedRequests\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.OwnerAllowlistedRequest[]\",\"components\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilitiesRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.Config\",\"components\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxExpiryLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonConfigs\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.DonConfigView[]\",\"components\":[{\"name\":\"donHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"family\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultUserLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"maxWorkflows\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultUserLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserDONOverrides\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.UserOverrideView[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflow\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowById\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwnerAndName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCapabilitiesRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"nameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"urlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"attrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"expiryLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"userDefaultLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"userLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsOnDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAllowlistedRequests\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalEvents\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWorkflowDONFamily\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilitiesRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigUpdated\",\"inputs\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxExpiryLen\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"userDefaultLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDonFamilyUpdated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowOwnerConfigUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BinaryURLRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUpdateDONFamilyForPausedWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DonLimitNotSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidExpiryTimestamp\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxAllowed\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerDONExceeded\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONDefaultLimitExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWorkflowIDNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60803460a8573315609757600180546001600160a01b0319163317905560a081016001600160401b03811182821017608157604090815280825260208083015260c882820152610400606083015262093a80608090920191909152600280546001600160481b03191667093a800400c82040179055516152a490816100ae8239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fdfe610140604052600436101561001357600080fd5b60003560e01c806301b7690514612c645780630987294c14612c2057806317e0edfc14612b14578063181f5a7714612af85780631c08b00a14612a295780631c71682c14612826578063245b8e4e146127e8578063274e00e01461278a578063289bd108146126be5780632afc4130146126665780632c50a9551461260957806339d68c6a146124a457806339e43234146124225780633c17181b146123cf5780633c54b50b1461237c5780634b6d2e5b146121da578063530979d61461214b578063556dbd0d146120d45780636104319314612037578063695e134014611fe657806370ae264014611fb057806379ba509714611f17578063865ec9e014611ede5780638b42a96d14611d8b5780638c42ffc514611cd75780638da5cb5b14611cb057806394ea0da614611ad5578063952bb98414611a8a578063a0b8a4fe14611a6c578063a4089016146119c9578063a6008f20146118a7578063a7d0185814611800578063afbb2401146117e9578063b377bfc514610960578063b668435f14610900578063ba870686146108e2578063bae5c29a14610869578063bdf6b4ff146107fd578063be674333146107d6578063bf2a0d9314610783578063c3f909d4146106d2578063cabb9e7a14610693578063d8b80738146105b1578063d8e4a72414610457578063dc101969146103b9578063de49b95f1461039b578063e690f33214610317578063ea32308b146102dd5763f2fde38b1461023857600080fd5b346102d85760203660031901126102d8576001600160a01b03610259612e72565b610261614631565b163381146102ae57806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102d85760203660031901126102d8576001600160a01b036102fe612e72565b16600052600b6020526020604060002054604051908152f35b346102d85760203660031901126102d85761033f336000526005602052604060002054151590565b1561038657600435600052600a60205260406000205461035f8133614876565b600160ff8183015460e01c1661037481612fa0565b0361037b57005b61038491614694565b005b63c2dda3f960e01b6000523360045260246000fd5b346102d85760003660031901126102d8576020601454604051908152f35b346102d85760603660031901126102d85760443560243567ffffffffffffffff82116102d8576103fd6103f26001933690600401612f72565b908360043533613a23565b3360005260066020528060406000205561041633614f4e565b5080600052600760205260406000208260ff19825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346102d85760403660031901126102d85760043567ffffffffffffffff81116102d85761048890369060040161311c565b602435918215158093036102d85761049e614631565b60ff831660005b83811061052f57505060405191806040840160408552526060830191906000905b8082106104fc577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b9091928335906001600160a01b03821682036102d857602080916001600160a01b036001941681520194019201906104c6565b6001600160a01b0361054a610545838787613589565b614475565b161561058757806001600160a01b036105696105456001948888613589565b16600052600360205260406000208360ff19825416179055016104a5565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102d85760203660031901126102d85760043567ffffffffffffffff81116102d8576105e290369060040161311c565b8015610669576105ff336000526005602052604060002054151590565b156103865760005b81811061061057005b8061061e6001928486613589565b35600052600a6020526040600020546106378133614876565b8360ff8183015460e01c1661064b81612fa0565b03610659575b505001610607565b61066291614694565b8480610651565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346102d85760203660031901126102d8576001600160a01b036106b4612e72565b166000526003602052602060ff604060002054166040519015158152f35b346102d85760003660031901126102d857600060806040516106f381612edc565b828152826020820152826040820152826060820152015260a060405161071881612edc565b63ffffffff60025461ffff60ff82169384815260ff60208201818560081c168152816040840191818760101c1683528760806060870196888a60181c168852019760281c16875260405198895251166020880152511660408601525116606084015251166080820152f35b346102d85760603660031901126102d85760043567ffffffffffffffff81116102d8576107b7610384913690600401612f72565b6107bf6131f4565b906107c8613207565b926107d1614631565b6141d8565b346102d85760206107ef6107e9366131b4565b91613ff6565b63ffffffff60405191168152f35b346102d85760203660031901126102d85760043567ffffffffffffffff81116102d857610831610838913690600401612f72565b3691613326565b6020815191012060005260166020526040600181600020015463ffffffff825191818116835260201c166020820152f35b346102d85760403660031901126102d857610882612e72565b6040516108bf816108b160208201946024359086602090939291936001600160a01b0360408201951681520152565b03601f198101835282612f15565b5190206000526013602052602063ffffffff604060002054166040519042108152f35b346102d85760003660031901126102d8576020601b54604051908152f35b346102d85760803660031901126102d857610919612e72565b60243567ffffffffffffffff81116102d857610939903690600401612f72565b610941613207565b906064359283151584036102d8576103849461095b614631565b613dd0565b346102d8576101203660031901126102d85760043567ffffffffffffffff81116102d857610992903690600401612f72565b610100526101205260243567ffffffffffffffff81116102d8576109ba903690600401612f72565b90600260643510156102d85760843567ffffffffffffffff81116102d8576109e6903690600401612f72565b9060e0529060a43567ffffffffffffffff81116102d857610a0b903690600401612f72565b9160c43567ffffffffffffffff81116102d857610a2c903690600401612f72565b9560e43567ffffffffffffffff81116102d857610a4d903690600401612f72565b929093610104359586151587036102d857610a75336000526005602052604060002054151590565b1561038657604435156117bf57604435600052600a60205260406000205461178f5760025460ff8160101c1689156117655780611728575b5061ffff8160181c168015158061171f575b6116ee575084156116c45760ff8160081c16801515806116bb575b61168a575061010051156116605760ff1680151580611654575b61162057506040516020810190610b19816108b1888761010051610120513389613c34565b5190209788600052600860205260406000206001600160a01b03600182015416801560001461129e575050610b5e610b58366101005161012051613326565b336144e4565b9a610b6c368c60e051613326565b602081519101209815611212575b610b85606435612fa0565b606435156111ea575b604051978861012081011067ffffffffffffffff6101208b01111761101357610c3b899695610c2c8897610c4a956101208a016040526044358a5260a0610c2060208c019a338c5260408d0160c05267ffffffffffffffff421660c0515260608d019a610bfc606435612fa0565b6064358c526080610c14366101005161012051613326565b9e019d8e523691613326565b9b019a8b523691613326565b9860c08c01998a523691613326565b9760e08a019889523691613326565b610100880160a05260a051528860005260086020526040600020965187556001600160a01b038060018901935116166001600160a01b031983541617825560c051517fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b8085549360a01b1616911617825551610cd181612fa0565b610cda81612fa0565b60ff60e01b197cff0000000000000000000000000000000000000000000000000000000083549260e01b16911617905560028501905180519067ffffffffffffffff821161101357610d3682610d30855461323f565b856139dc565b602090601f831160011461118357610d67929160009183611178575b50508160011b916000199060031b1c19161790565b90555b518051600385019167ffffffffffffffff821161101357610d8f82610d30855461323f565b602090601f831160011461111157610dbf9291600091836111065750508160011b916000199060031b1c19161790565b90555b518051600484019167ffffffffffffffff821161101357610de782610d30855461323f565b602090601f831160011461109f57610e179291600091836110295750508160011b916000199060031b1c19161790565b90555b518051600583019167ffffffffffffffff821161101357610e3f82610d30855461323f565b602090601f8311600114611034579180610e749260069695946000926110295750508160011b916000199060031b1c19161790565b90555b019360a0515194855167ffffffffffffffff811161101357610ea381610e9d845461323f565b846139dc565b6020601f8211600114610fa9579080610ed792610f2d9899600092610f9e5750508160011b916000199060031b1c19161790565b90555b6000526009602052610ef082604060002061501f565b50604435600052600a60205281604060002055600052600d602052610f1981604060002061501f565b5033600052600e602052604060002061501f565b50610f466040519160608352606083019060e05161353c565b610f51606435612fa0565b606435602083015281810360408301527f74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a09339280610f9960443594610100516101205161353c565b0390a3005b015190508980610d52565b601f1982169783600052816000209860005b818110610ffb575091610f2d989991846001959410610fe2575b505050811b019055610eda565b015160001960f88460031b161c19169055888080610fd5565b838301518b556001909a019960209384019301610fbb565b634e487b7160e01b600052604160045260246000fd5b015190508a80610d52565b90601f1983169184600052816000209260005b81811061108757509160019391856006989796941061106e575b505050811b019055610e77565b015160001960f88460031b161c19169055898080611061565b92936020600181928786015181550195019301611047565b90601f1983169184600052816000209260005b8181106110ee57509084600195949392106110d5575b505050811b019055610e1a565b015160001960f88460031b161c191690558980806110c8565b929360206001819287860151815501950193016110b2565b015190508b80610d52565b90601f1983169184600052816000209260005b8181106111605750908460019594939210611147575b505050811b019055610dc2565b015160001960f88460031b161c191690558a808061113a565b92936020600181928786015181550195019301611124565b015190508c80610d52565b90601f1983169184600052816000209260005b8181106111d257509084600195949392106111b9575b505050811b019055610d6a565b015160001960f88460031b161c191690558b80806111ac565b92936020600181928786015181550195019301611196565b6112016111fa368d60e051613326565b8a3361491e565b61120d8c8a338d614cf2565b610b8e565b9a999897969594939291908a600052600f60205260406000206080525b60805154801561128d5780600019810111611277576112576112729160001901608051614db8565b90549060031b1c806000526008602052604060002090614694565b61122f565b634e487b7160e01b600052601160045260246000fd5b50909192939495969798999a610b7a565b9294969850969a9450979150330361160b578454600052600a60205260006040812055604435600052600a602052604060002055835495604435855560038501916040516112f7816112f08187613279565b0382612f15565b60208151910120611309368484613326565b6020815191012003611544575b505050600483019160405161132f816112f08187613279565b60208151910120611341368484613326565b6020815191012003611474575b5050506006019067ffffffffffffffff81116110135761137281610e9d845461323f565b6000601f82116001146114105781906113a4939495966000926114055750508160011b916000199060031b1c19161790565b90555b6113bf6040519260408452604084019060e05161353c565b9082820360208401527f03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e32233938061140060443595610100516101205161353c565b0390a4005b013590508680610d52565b601f198216958382526020822091805b88811061145c57508360019596979810611442575b505050811b0190556113a7565b0135600019600384901b60f8161c19169055858080611435565b90926020600181928686013581550194019101611420565b67ffffffffffffffff82116110135761149182610d30855461323f565b600090601f83116001146114dc5791806114c69260069695946000926114d15750508160011b916000199060031b1c19161790565b90555b90868061134e565b013590508a80610d52565b8382526020822091601f198416815b81811061152c575091600193918560069897969410611512575b505050811b0190556114c9565b0135600019600384901b60f8161c19169055898080611505565b919360206001819287870135815501950192016114eb565b67ffffffffffffffff82116110135761156182610d30855461323f565b600090601f83116001146115a75761159192916000918361159c5750508160011b916000199060031b1c19161790565b90555b878080611316565b013590508b80610d52565b8382526020822091601f198416815b8181106115f357509084600195949392106115d9575b505050811b019055611594565b0135600019600384901b60f8161c191690558a80806115cc565b919360206001819287870135815501950192016115b6565b6331ee6dc760e01b6000523360045260246000fd5b7f36a7c503000000000000000000000000000000000000000000000000000000006000526101005160045260245260446000fd5b50806101005111610af4565b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b857f436f97540000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808611610ada565b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b867f354f25140000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808711610abf565b808a1161174e57808c1115610aad578b6219aad560e31b60005260045260245260446000fd5b896219aad560e31b60005260045260245260446000fd5b7f9cd963cf0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260443560045260246000fd5b7f315de7450000000000000000000000000000000000000000000000000000000060005260046000fd5b346102d8576103846117fa36613184565b91613c69565b346102d85760203660031901126102d85760043567ffffffffffffffff81116102d85761183190369060040161311c565b80156106695760005b81811061184357005b806118516001928486613589565b3561185a614631565b600052600a602052604060002054806000526008602052604060002060ff8482015460e01c1661188981612fa0565b15611897575b50500161183a565b6118a091614694565b848061188f565b346102d85760a03660031901126102d85760043560ff81168091036102d8576024359060ff8216908183036102d8576044359160ff8316908184036102d8576064359161ffff8316918284036102d8576084359663ffffffff8816948589036102d85764ffff00000060a09862ff00007f9c1a161a4cdd9b19a46f9660eee21b6394dc5aa70fc9e093dbb36d2c1786d7739b611941614631565b89608060405161195081612edc565b8d81528960208201528a60408201528b6060820152015268ffffffff00000000006002549160281b169561ff0068ffffffff000000000019928d87199162ffffff191617169160081b1617169160101b16179160181b1617176002556040519485526020850152604084015260608301526080820152a1005b346102d85760603660031901126102d8576119e2612e72565b60243567ffffffffffffffff81116102d857611a02903690600401612f72565b916044359267ffffffffffffffff84116102d857611a4c611a54936108b1611a31611a68973690600401612f72565b90611a3a61336d565b50604051958694602086019889613c34565b519020614503565b604051918291602083526020830190612faa565b0390f35b346102d85760003660031901126102d8576020600454604051908152f35b346102d85760803660031901126102d857611aa3612e72565b6064359067ffffffffffffffff82116102d857611ac7610384923690600401612f72565b916044359060243590613a23565b346102d85760403660031901126102d857600435611af16131f4565b63ffffffff8060025460281c16911690428211801590611c94575b611c605750611b28336000526005602052604060002054151590565b1561038657604080513360208201908152918101849052611b4c81606081016108b1565b519020600052601360205260406000208163ffffffff19825416179055604051611b7581612ea4565b828152602081019033825260408101908382526014546801000000000000000081101561101357806001611bac920160145561340b565b611c4a5760016001600160a01b039291839251815501935116166001600160a01b0319835416178255517fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff000000000000000000000000000000000000000083549260a01b1691161790556040519081527ff69135e4f80a25991d2f877c365c191c51ec3c0063ecb9299d314cd9da4880d160203392a3005b634e487b7160e01b600052600060045260246000fd5b917f7ffd3b8f0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b508015158015611b0c575080611caa428461352f565b11611b0c565b346102d85760003660031901126102d85760206001600160a01b0360015416604051908152f35b346102d85760603660031901126102d857611cf0612e72565b6001600160a01b0360243591169081600052600e602052611d1960443582604060002054614489565b90611d23826133bb565b9260005b838110611d3c5760405180611a688782613075565b60019082600052600e602052611d6f611d636040600020611d5d8488613232565b90614db8565b90549060031b1c614503565b611d7982886132fc565b52611d8481876132fc565b5001611d27565b346102d8576001600160a01b03611da1366131b4565b929091611dac614631565b169081600052601560205260406000209267ffffffffffffffff811161101357611de081611dda865461323f565b866139dc565b600093601f8211600114611e5c57611e3282807f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f139697600091611e51575b508160011b916000199060031b1c19161790565b90555b611e4c60405192839260208452602084019161353c565b0390a2005b905085013588611e1e565b80855260208520601f19831695805b878110611ec65750837f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f13969710611eac575b5050600182811b019055611e35565b840135600019600385901b60f8161c191690558580611e9d565b90916020600181928589013581550193019101611e6b565b346102d85760003660031901126102d8576040601a5467ffffffffffffffff8251916001600160a01b038116835260a01c166020820152f35b346102d85760003660031901126102d8576000546001600160a01b0381163303611f86576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102d85760203660031901126102d857611fc961336d565b50600435600052600a602052611a68611a54604060002054614503565b346102d85760203660031901126102d85761200e336000526005602052604060002054151590565b1561038657600435600052600a6020526103846040600020546120318133614876565b90614b5b565b346102d85761207461205761204b3661314d565b94929093913691613326565b602081519101209283600052600d60205282604060002054614489565b9061207e826133bb565b9260005b8381106120975760405180611a688782613075565b60019082600052600d6020526120b8611d636040600020611d5d8488613232565b6120c282886132fc565b526120cd81876132fc565b5001612082565b346102d85760203660031901126102d85760043567ffffffffffffffff81116102d857612108612110913690600401612f72565b610831614631565b60208151910120600052600c60205260406000205b80548015610384576000198101908111611277576112576121469183614db8565b612125565b346102d85761215936613184565b9091612172336000526005602052604060002054151590565b1561038657600052600a60205260406000205461218f8133614876565b9060ff600183015460e01c166121a481612fa0565b6121aa57005b6121d56121ce610384956121bf368783613326565b60208151910120953691613326565b843361491e565b614a60565b346102d8576121f66121eb3661314d565b929093913691613326565b60208151910120806000526019602052612217604060002092848454614489565b916122218361321a565b9361222f6040519586612f15565b838552601f1961223e8561321a565b0160005b81811061235757505060005b8481106122b6578560405180916020820160208352815180915260206040840192019060005b818110612282575050500390f35b825180516001600160a01b0316855260209081015163ffffffff168186015286955060409094019390920191600101612274565b806001600160a01b036122d46122ce60019486613232565b86614db8565b90549060031b1c168560005260166020526040600281600020016000906001600160a01b0384168252602052206040519061230e82612ec0565b5490602060ff63ffffffff841693848452821c1615159101526040519161233483612ec0565b8252602082015261234582896132fc565b5261235081886132fc565b500161224e565b60209060405161236681612ec0565b6000815260008382015282828a01015201612242565b346102d85760203660031901126102d85760043567ffffffffffffffff81116102d8576108316123b0913690600401612f72565b60208151910120600052600c6020526020604060002054604051908152f35b346102d85760203660031901126102d8576123e8614631565b600435600052600a602052604060002054806000526008602052604060002060ff600182015460e01c1661241b81612fa0565b1561037b57005b346102d85761243961243336612e1b565b9061385c565b60405180916020820160208352815180915260206040840192019060005b818110612465575050500390f35b919350916020606060019263ffffffff60408851805184526001600160a01b038682015116868501520151166040820152019401910191849392612457565b346102d8576124b2366130d5565b9282919242116125e4576001600160a01b038116936124de856000526005602052604060002054151590565b156125cf576124ec85614eaf565b936001600160a01b036125028383888888614795565b16600052600360205260ff60406000205416156125ae57858581600052600e6020526040600020905b8154801561256c5760001981019081116112775761254c6125679184614db8565b90549060031b1c806000526008602052604060002090614b5b565b61252b565b6000828580835260066020528260408120556125878161507b565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b60405163335d4ce160e01b81529485946125cb946004870161355d565b0390fd5b8463c2dda3f960e01b60005260045260246000fd5b6001600160a01b0390631ec5288b60e11b600052166004524260245260445260646000fd5b346102d85760403660031901126102d85760043567ffffffffffffffff81116102d85761263a90369060040161311c565b6024359167ffffffffffffffff83116102d85761265e610384933690600401612f72565b929091613599565b346102d85760203660031901126102d8576001600160a01b03612687612e72565b166000526015602052611a686112f06126aa604060002060405192838092613279565b604051918291602083526020830190612e31565b346102d8576126cc366130d5565b8293924211612764576001600160a01b0383166126f6816000526005602052604060002054151590565b156127505761270490614eaf565b916001600160a01b0361271a8383868989614795565b16600052600360205260ff604060002054161561273357005b6125cb9260405195869563335d4ce160e01b87526004870161355d565b63c2dda3f960e01b60005260045260246000fd5b836001600160a01b0384631ec5288b60e11b600052166004524260245260445260646000fd5b346102d85760203660031901126102d8576001600160a01b036127ab612e72565b6127b3614631565b16600052600b60205260406000205b80548015610384576000198101908111611277576112576127e39183614db8565b6127c2565b346102d85760403660031901126102d857612801612e72565b60243567ffffffffffffffff811681036102d85761038491612821614631565b61342a565b346102d85761284361283736612e1b565b81601b93929354614489565b9061284d8261321a565b9161285b6040519384612f15565b808352601f1961286a8261321a565b0160005b8181106129fd575050601b54909160005b83811061292b578460405160208101916020825280518093526040820192602060408260051b85010192019060005b8181106128bb5784840385f35b909192603f1985820301865283519081519160038310156129155761290682606060406020959460019787965263ffffffff8682015116868501520151918160408201520190612e31565b950196019101949190946128ae565b634e487b7160e01b600052602160045260246000fd5b6129358183613232565b6000848210156129e957601b90526040519060011b7f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc101600061297783612ea4565b81549060ff82169060038210156129d55750835260081c63ffffffff16602083015260405160019392916129b49082906112f09082908801613279565b60408201526129c382886132fc565b526129ce81876132fc565b500161287f565b80634e487b7160e01b602492526021600452fd5b80634e487b7160e01b602492526032600452fd5b602090604051612a0c81612ea4565b60008152600083820152606060408201528282880101520161286e565b346102d85760803660031901126102d857612a42612e72565b60243567ffffffffffffffff81116102d857612a7691612a69612a7c923690600401612f72565b9390604435943691613326565b906144e4565b90816000526009602052612a9860643582604060002054614489565b90612aa2826133bb565b9260005b838110612abb5760405180611a688782613075565b600190826000526009602052612adc611d636040600020611d5d8488613232565b612ae682886132fc565b52612af181876132fc565b5001612aa6565b346102d85760003660031901126102d857611a686126aa612f37565b346102d857612b31612b2536612e1b565b81600493929354614489565b612b3a8161321a565b91612b486040519384612f15565b818352612b548261321a565b602084019290601f19013684376004549160005b828110612bbd5784866040519182916020830190602084525180915260408301919060005b818110612b9b575050500390f35b82516001600160a01b0316845285945060209384019390920191600101612b8d565b612bc78183613232565b6000858210156129e957600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b015460008190526006602052600191906001600160a01b0316612c1982896132fc565b5201612b68565b346102d85760203660031901126102d8576020612c5a6001600160a01b03612c46612e72565b166000526005602052604060002054151590565b6040519015158152f35b346102d857612c81612c7536612e1b565b81601793929354614489565b612c8a8161321a565b91612c986040519384612f15565b818352601f19612ca78361321a565b0160005b818110612de8575050601754919060005b828110612d5a57846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210612cfc57505050500390f35b919360019193955060208091603f1989820301855287519081518152606063ffffffff81612d37868601516080888701526080860190612e31565b948260408201511660408601520151169101529601920192018594939192612ced565b612d648183613232565b6000858210156129e9579060208260176001959452200163ffffffff604060009254928381526016602052206112f0612db6868301549260405195612da887612e88565b865260405192838092613279565b6020840152818116604084015260201c166060820152612dd682886132fc565b52612de181876132fc565b5001612cbc565b602090604051612df781612e88565b60008152606083820152600060408201526000606082015282828801015201612cab565b60409060031901126102d8576004359060243590565b919082519283825260005b848110612e5d575050826000602080949584010152601f8019910116010190565b80602080928401015182828601015201612e3c565b600435906001600160a01b03821682036102d857565b6080810190811067ffffffffffffffff82111761101357604052565b6060810190811067ffffffffffffffff82111761101357604052565b6040810190811067ffffffffffffffff82111761101357604052565b60a0810190811067ffffffffffffffff82111761101357604052565b610140810190811067ffffffffffffffff82111761101357604052565b90601f8019910116810190811067ffffffffffffffff82111761101357604052565b60405190612f46604083612f15565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b9181601f840112156102d85782359167ffffffffffffffff83116102d857602083818601950101116102d857565b6002111561291557565b61307291815181526001600160a01b03602083015116602082015267ffffffffffffffff60408301511660408201526060820151612fe781612fa0565b606082015261012061306061304c61303a61302861301660808801516101406080890152610140880190612e31565b60a088015187820360a0890152612e31565b60c087015186820360c0880152612e31565b60e086015185820360e0870152612e31565b610100850151848203610100860152612e31565b92015190610120818403910152612e31565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106130a857505050505090565b90919293946020806130c6600193603f198682030187528951612faa565b97019301930191939290613099565b60606003198201126102d8576004356001600160a01b03811681036102d85791602435916044359067ffffffffffffffff82116102d85761311891600401612f72565b9091565b9181601f840112156102d85782359167ffffffffffffffff83116102d8576020808501948460051b0101116102d857565b60606003198201126102d8576004359067ffffffffffffffff82116102d85761317891600401612f72565b90916024359060443590565b9060406003198301126102d857600435916024359067ffffffffffffffff82116102d85761311891600401612f72565b9060406003198301126102d8576004356001600160a01b03811681036102d857916024359067ffffffffffffffff82116102d85761311891600401612f72565b6024359063ffffffff821682036102d857565b6044359063ffffffff821682036102d857565b67ffffffffffffffff81116110135760051b60200190565b9190820180921161127757565b90600182811c9216801561326f575b602083101461325957565b634e487b7160e01b600052602260045260246000fd5b91607f169161324e565b600092918154916132898361323f565b80835292600181169081156132df57506001146132a557505050565b60009081526020812093945091925b8383106132c5575060209250010190565b6001816020929493945483858701015201910191906132b4565b915050602093945060ff929192191683830152151560051b010190565b80518210156133105760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b92919267ffffffffffffffff82116110135760405191613350601f8201601f191660200184612f15565b8294818452818301116102d8578281602093846000960137010152565b6040519061337a82612ef8565b6060610120836000815260006020820152600060408201526000838201528260808201528260a08201528260c08201528260e0820152826101008201520152565b906133c58261321a565b6133d26040519182612f15565b82815280926133e3601f199161321a565b019060005b8281106133f457505050565b6020906133ff61336d565b828285010152016133e8565b60145481101561331057601460005260206000209060011b0190600090565b90601a54906001600160a01b038216906001600160a01b0367ffffffffffffffff8460a01c169416938285149485809661351c575b61351457806080957fc0c3ee74e6d6070ee9c493e8b4f0477d2e66600f22997a4e073288d38d65933b97156134ff575b505067ffffffffffffffff8316928284036134be575b50604051938452602084015260408301526060820152a1565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b601a549260a01b16911617601a55386134a5565b6001600160a01b03191617601a55803861348f565b505050505050565b508167ffffffffffffffff84161461345f565b9190820391821161127757565b908060209392818452848401376000828201840152601f01601f1916010190565b90926080926001600160a01b03613072979516835260208301526040820152816060820152019161353c565b91908110156133105760051b0190565b92918115610669576135b8336000526005602052604060002054151590565b15610386576135c8368483613326565b6020815191012092600091825b848110613794575063ffffffff831615613514576135f4913691613326565b90836000526016602052604060002090600182015463ffffffff811680156137725786600052601160205263ffffffff613635848260406000205416614904565b16116137505763ffffffff60029160201c1692016001600160a01b0333166000528060205260ff60406000205460201c16613730575b5063ffffffff61369881923360005260106020526040600020886000526020528260406000205416614904565b9216911611613712575060005b8181106136b25750505050565b806136c06001928487613589565b35600052600a6020526040600020548060005260086020528460406000208460ff8183015460e01c166136f281612fa0565b14613701575b505050016136a5565b61370a92614a60565b3884816136f8565b6125cb9060405191829163038857ff60e01b835233600484016144c4565b336000908152602091909152604090205463ffffffff908116925061366b565b60405163b993868760e01b815260206004820152806125cb6024820187612e31565b60405163393f328760e11b815260206004820152806125cb6024820188612e31565b926137a0848689613589565b35600052600a60205260ff60016137bc60406000205433614876565b015460e01c166137cb81612fa0565b156137ee5763ffffffff1663ffffffff8114611277576001809101935b016135d5565b926001906137e8565b906138018261321a565b61380e6040519182612f15565b828152809261381f601f199161321a565b019060005b82811061383057505050565b60209060405161383f81612ea4565b600081526000838201526000604082015282828501015201613824565b9061386a9082601454614489565b9081156139615761387a826137f7565b91600091825b8281106138d75750508110613893575090565b61389c816137f7565b9160005b8281106138ad5750505090565b806138ba600192846132fc565b516138c582876132fc565b526138d081866132fc565b50016138a0565b6138e96138e48284613232565b61340b565b50600181015463ffffffff8160a01c169042821161390d575b505050600101613880565b6001600160a01b03906040979497519361392685612ea4565b548452166020830152604082015261393e82876132fc565b5261394981866132fc565b50600019811461127757600180910193903880613902565b50506040516000613973602083612f15565b81526000805b81811061398557505090565b60209060405161399481612ea4565b600081526000838201526000604082015282828601015201613979565b80634e487b7160e01b602492526041600452fd5b8181106139d0575050565b600081556001016139c5565b9190601f81116139eb57505050565b613a17926000526020600020906020601f840160051c83019310613a19575b601f0160051c01906139c5565b565b9091508190613a0a565b92909391844211613bf5576001600160a01b038416613a4f816000526005602052604060002054151590565b613bc85781600052600760205260ff60406000205416613b9857600090613aab613ac5613a7a612f37565b6040519283916020830195878752604084015246606084015230608084015260e060a0840152610100830190612e31565b8a60c08301528660e083015203601f198101835282612f15565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613b06603c8220613b00368787613326565b90614f12565b90919260048310156129d55782613b60575050506001600160a01b031660009081526003602052604090205460ff1615613b41575050505050565b906125cb929160405195869563335d4ce160e01b87526004870161355d565b5060405163d36ab6b960e01b81526060600482015291829160ff613b88606485018a8a61353c565b9216602484015260448301520390fd5b7f77a338580000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7fd9a5f5ca0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b846001600160a01b03857f502d038700000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b939161307295936001600160a01b03613c5b9316865260606020870152606086019161353c565b92604081850391015261353c565b9091613c82336000526005602052604060002054151590565b156103865781600052600a60205260406000205490613ca18233614876565b60ff600182015460e01c16613cb581612fa0565b613d8157613cea9083600052601260205260406000205492836000526016602052613cf1604060002060405194858092613279565b0384612f15565b613cfc368289613326565b60208151910120809414613d7857613d4f613d63948387613d41613d73967f9b5361a5258ef6ac8039fd2d2ac276734695219cfd870711e7922c236e5db16d9a614694565b6121d56121ce36878e613326565b604051938493604085526040850190612e31565b908382036020850152339761353c565b0390a3565b50505050505050565b7fd74915a80000000000000000000000000000000000000000000000000000000060005260046000fd5b91613dc960209263ffffffff9296959660408652604086019161353c565b9416910152565b92613ddc368484613326565b60208151910120948560005260166020526040600020600181019063ffffffff82541615613fd357600201604060006001600160a01b0389168152826020522092600014613f4557505463ffffffff83811691168111613f1b57815460ff8160201c1615600014613ebb5750815464ffffffffff1916176401000000001790556000948552601960205260409094206001600160a01b03909316937f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c8232493613eb69190613ea990879061501f565b5060405193849384613dab565b0390a2565b819392959694975063ffffffff161415600014613514577f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c82324946001600160a01b0394613eb69363ffffffff19825416179055604051948594169684613dab565b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b959692505060ff909392935460201c1615613fcc577f6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d849936001600160a01b038616600052602052600060408120556000526019602052613fb46001600160a01b03604060002095168095615145565b50613eb660405192839260208452602084019161353c565b5050505050565b60405163393f328760e11b815260206004820152806125cb60248201898961353c565b9190614003913691613326565b6020815191012060005260166020526001600160a01b036040600020911660005260028101602052604060002060206040519161403f83612ec0565b549160ff63ffffffff841693848352831c161515918291015261406d57506001015460201c63ffffffff1690565b905090565b90601b54680100000000000000008110156110135760018101601b556000601b548210156129e957601b90526000929060011b7f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc101815160038110156141c45760409160019160ff825491168060ff19831617835564ffffffff00602087015160081b169164ffffffffff191617178155019101519283519067ffffffffffffffff82116139b15761412882610d30855461323f565b602090601f831160011461416657906141579394958361415b5750508160011b916000199060031b1c19161790565b9055565b015190503880610d52565b90601f198316848352818320925b8181106141ac57509583600195969710614193575b505050811b019055565b015160001960f88460031b161c19169055388080614189565b9192602060018192868b015181550194019201614174565b602485634e487b7160e01b81526021600452fd5b919392936141e7368385613326565b60208151910120948560005260166020526040600020956001870163ffffffff81549416938463ffffffff821614908161445f575b506144555763ffffffff83169284841161442b57885461423b8161323f565b15614304575b507fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d979850906142dd92918563ffffffff1983541617825567ffffffff0000000082549160201b169067ffffffff000000001916179055604051906020820152836040820152604081526142b6606082612f15565b604051906142c382612ea4565b6000825263ffffffff421660208301526040820152614072565b6142f460405194859460608652606086019161353c565b91602084015260408301520390a1565b67ffffffffffffffff871161101357866143206143269261323f565b8b6139dc565b600098601f87116001146143a157906142dd93929161437e88807fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d9c9d60009161439657508160011b916000199060031b1c19161790565b90555b61438a83614fc5565b50909192899850614241565b90508b013538611e1e565b808a5260208a20601f1988168b5b8181106144135750907fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d9a9b896142dd9796959493106143f9575b5050600188811b019055614381565b8a013560001960038b901b60f8161c1916905538806143ea565b8a8d0135835560209c8d019c600190930192016143af565b7fec623c5f0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050505050509050565b905063ffffffff8085169160201c16143861421c565b356001600160a01b03811681036102d85790565b808210156144bc5782816144a06130729585613232565b11156144ac575061352f565b6144b7915082613232565b61352f565b505050600090565b6040906001600160a01b0361307294931681528160208201520190612e31565b906144fd6108b1916040519283916020830195866144c4565b51902090565b61450b61336d565b81600052600860205260406000206001810154926001600160a01b038416928315614629575091600661460e836112f09567ffffffffffffffff61462096549860ff8160e01c1692600052601260205260406000205460005260166020526040600020966040519a61457c8c612ef8565b8b5260208b015260a01c16604089015261459581612fa0565b60608801526040516145ae816112f08160028601613279565b60808801526040516145c7816112f08160038601613279565b60a08801526040516145e0816112f08160048601613279565b60c08801526040516145f9816112f08160058601613279565b60e08801526112f06040518094819301613279565b61010085015260405192838092613279565b61012082015290565b935050505090565b6001600160a01b0360015416330361464557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b909161468661307293604084526040840190613279565b916020818403910152613279565b7ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e48868960018301916001600160a01b038354947c010000000000000000000000000000000000000000000000000000000060ff60e01b19871617855582600052601260205261472a8260406000205497169387600284019561472460405161471e816112f0818c613279565b826144e4565b92614de9565b61477481546040519088602083015260408201526040815261474d606082612f15565b6040519061475a82612ea4565b6002825263ffffffff421660208301526040820152614072565b549354169360005260166020526040600020613d736040519283928361466f565b916147fb906147e2926147a6612f37565b916040519485936001600160a01b03602086019860018a5216604086015246606086015230608086015260e060a0860152610100850190612e31565b9160c084015260e083015203601f198101835282612f15565b5190207f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52614838603c600020613b00368585613326565b60048295939510156129155781614850575050505090565b60ff613b8860405195869563d36ab6b960e01b875260606004880152606487019161353c565b9060005260086020526040600020906001600160a01b036001830154169081156148c2576001600160a01b03168091036148ae575090565b6331ee6dc760e01b60005260045260246000fd5b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b63ffffffff60019116019063ffffffff821161127757565b9063ffffffff8091169116019063ffffffff821161127757565b90806000526016602052604060002090600182015463ffffffff81168015614a3e5782600052601160205263ffffffff61496060018260406000205416614904565b1611614a1c5763ffffffff60029160201c1692016001600160a01b0384166000528060205260ff60406000205460201c166149f6575b506001600160a01b038316600052601060205260406000209060005260205263ffffffff806149cd60018260406000205416614904565b92169116116149da575050565b6125cb60405192839263038857ff60e01b8452600484016144c4565b9091506001600160a01b03831660005260205263ffffffff604060002054169038614996565b60405163b993868760e01b815260206004820152806125cb6024820188612e31565b60405163393f328760e11b815260206004820152806125cb6024820189612e31565b7f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a906001600160a01b036001840193614ab78286541693876002840195614ab160405161471e816112f0818c613279565b92614cf2565b60ff60e01b198554168555614774815460405190886020830152604082015260408152614ae5606082612f15565b60405190614af282612ea4565b6001825263ffffffff421660208301526040820152614072565b614b16815461323f565b9081614b20575050565b81601f60009311600114614b32575055565b81835260208320614b4e91601f0160051c8101906001016139c5565b8082528160208120915555565b9061471e907f48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd94614c7a614be9614c886001600160a01b036001860195865460ff8a848316614bb76112f0600287019d8e60405192838092613279565b938491836000526012602052604060002054948592836000526016602052614bf060406000206040519d8e8092613279565b038d612f15565b60e01c16614bfd81612fa0565b15614ce0575b50505050600052600d602052614c1d8a6040600020615145565b5082885416600052600e602052614c388a6040600020615145565b506000526009602052614c4f896040600020615145565b508054600052600a602052600060408120555495541695604051938493604085526040850190612e31565b908382036020850152613279565b0390a36000526008602052600660406000206000815560006001820155614cb160028201614b0c565b614cbd60038201614b0c565b614cc960048201614b0c565b614cd560058201614b0c565b01614b16815461323f565b614ce993614de9565b8a828238614c03565b9192906001600160a01b0316806000526010602052604060002084600052602052604060002063ffffffff614d29818354166148ec565b1663ffffffff19825416179055836000526011602052604060002063ffffffff614d55818354166148ec565b1663ffffffff1982541617905583600052600c602052614d7983604060002061501f565b50600052600b602052614d9082604060002061501f565b50600052600f602052614da781604060002061501f565b506000526012602052604060002055565b80548210156133105760005260206000200190600090565b63ffffffff6000199116019063ffffffff821161127757565b9291906001600160a01b031680600052600b602052614e0c846040600020615145565b5081600052600c602052614e24846040600020615145565b506000526010602052604060002081600052602052604060002063ffffffff614e4f81835416614dd0565b1663ffffffff198254161790556000526011602052604060002063ffffffff614e7a81835416614dd0565b1663ffffffff19825416179055600052600f602052614e9d816040600020615145565b50600052601260205260006040812055565b80600052600660205260406000205490811580614efc575b614ecf575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5080600052600560205260406000205415614ec7565b8151919060418303614f4357614f3c92506020820151906060604084015193015160001a90615202565b9192909190565b505060009160029190565b80600052600560205260406000205415600014614fbf576004546801000000000000000081101561101357614fa6614f8f8260018594016004556004614db8565b819391549060031b91821b91600019901b19161790565b9055600454906000526005602052604060002055600190565b50600090565b80600052601860205260406000205415600014614fbf576017546801000000000000000081101561101357615006614f8f8260018594016017556017614db8565b9055601754906000526018602052604060002055600190565b60008281526001820160205260409020546150745780549068010000000000000000821015611013578261505d614f8f846001809601855584614db8565b905580549260005201602052604060002055600190565b5050600090565b6000818152600560205260409020548015615074576000198101818111611277576004546000198101919082116112775781810361510b575b50505060045480156150f557600019016150cf816004614db8565b8154906000199060031b1b19169055600455600052600560205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61512d61511c614f8f936004614db8565b90549060031b1c9283926004614db8565b905560005260056020526040600020553880806150b4565b90600182019181600052826020526040600020548015156000146151f9576000198101818111611277578254600019810191908211611277578181036151c2575b505050805480156150f55760001901906151a08282614db8565b8154906000199060031b1b191690555560005260205260006040812055600190565b6151e26151d2614f8f9386614db8565b90549060031b1c92839286614db8565b905560005283602052604060002055388080615186565b50505050600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161528b579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa1561527f576000516001600160a01b038116156152735790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b5050506000916003919056fea164736f6c634300081a000a",
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetConfig(opts *bind.CallOpts) (WorkflowRegistryConfig, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(WorkflowRegistryConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryConfig)).(*WorkflowRegistryConfig)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetConfig() (WorkflowRegistryConfig, error) {
	return _WorkflowRegistry.Contract.GetConfig(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetConfig() (WorkflowRegistryConfig, error) {
	return _WorkflowRegistry.Contract.GetConfig(&_WorkflowRegistry.CallOpts)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (GetMaxWorkflowsPerDON,

	error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerDON", donFamily)

	outstruct := new(GetMaxWorkflowsPerDON)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxWorkflows = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DefaultUserLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerDON(donFamily string) (GetMaxWorkflowsPerDON,

	error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerDON(donFamily string) (GetMaxWorkflowsPerDON,

	error) {
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetConfig(opts *bind.TransactOpts, nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setConfig", nameLen, tagLen, urlLen, attrLen, expiryLen)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetConfig(nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetConfig(&_WorkflowRegistry.TransactOpts, nameLen, tagLen, urlLen, attrLen, expiryLen)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetConfig(nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetConfig(&_WorkflowRegistry.TransactOpts, nameLen, tagLen, urlLen, attrLen, expiryLen)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONLimit(opts *bind.TransactOpts, donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONLimit", donFamily, donLimit, userDefaultLimit)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONLimit(donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donFamily, donLimit, userDefaultLimit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONLimit(donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donFamily, donLimit, userDefaultLimit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setUserDONOverride", user, donFamily, userLimit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetUserDONOverride(user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donFamily, userLimit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetUserDONOverride(user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donFamily, userLimit, enabled)
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

type WorkflowRegistryConfigUpdatedIterator struct {
	Event *WorkflowRegistryConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryConfigUpdated)
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
		it.Event = new(WorkflowRegistryConfigUpdated)
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

func (it *WorkflowRegistryConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryConfigUpdated struct {
	MaxNameLen   uint8
	MaxTagLen    uint8
	MaxUrlLen    uint8
	MaxAttrLen   uint16
	MaxExpiryLen uint32
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterConfigUpdated(opts *bind.FilterOpts) (*WorkflowRegistryConfigUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryConfigUpdatedIterator{contract: _WorkflowRegistry.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryConfigUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseConfigUpdated(log types.Log) (*WorkflowRegistryConfigUpdated, error) {
	event := new(WorkflowRegistryConfigUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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
	DonFamily        string
	DonLimit         uint32
	UserDefaultLimit uint32
	Raw              types.Log
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

type GetMaxWorkflowsPerDON struct {
	MaxWorkflows     uint32
	DefaultUserLimit uint32
}

func (_WorkflowRegistry *WorkflowRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WorkflowRegistry.abi.Events["AllowedSignersUpdated"].ID:
		return _WorkflowRegistry.ParseAllowedSignersUpdated(log)
	case _WorkflowRegistry.abi.Events["CapabilitiesRegistryUpdated"].ID:
		return _WorkflowRegistry.ParseCapabilitiesRegistryUpdated(log)
	case _WorkflowRegistry.abi.Events["ConfigUpdated"].ID:
		return _WorkflowRegistry.ParseConfigUpdated(log)
	case _WorkflowRegistry.abi.Events["DONLimitSet"].ID:
		return _WorkflowRegistry.ParseDONLimitSet(log)
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

func (WorkflowRegistryConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x9c1a161a4cdd9b19a46f9660eee21b6394dc5aa70fc9e093dbb36d2c1786d773")
}

func (WorkflowRegistryDONLimitSet) Topic() common.Hash {
	return common.HexToHash("0xe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d")
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

	GetConfig(opts *bind.CallOpts) (WorkflowRegistryConfig, error)

	GetDonConfigs(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error)

	GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (GetMaxWorkflowsPerDON,

		error)

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

	SetConfig(opts *bind.TransactOpts, nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error)

	SetDONLimit(opts *bind.TransactOpts, donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error)

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

	FilterConfigUpdated(opts *bind.FilterOpts) (*WorkflowRegistryConfigUpdatedIterator, error)

	WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryConfigUpdated) (event.Subscription, error)

	ParseConfigUpdated(log types.Log) (*WorkflowRegistryConfigUpdated, error)

	FilterDONLimitSet(opts *bind.FilterOpts) (*WorkflowRegistryDONLimitSetIterator, error)

	WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet) (event.Subscription, error)

	ParseDONLimitSet(log types.Log) (*WorkflowRegistryDONLimitSet, error)

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
