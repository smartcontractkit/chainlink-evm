// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package capabilities_registry_wrapper_v2

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

type CapabilitiesRegistryAdditionalDONParams struct {
	Name   string
	Config []byte
}

type CapabilitiesRegistryCapability struct {
	CapabilityId          string
	ConfigurationContract common.Address
	Metadata              []byte
}

type CapabilitiesRegistryCapabilityConfiguration struct {
	CapabilityId string
	Config       []byte
}

type CapabilitiesRegistryCapabilityInfo struct {
	CapabilityId          string
	ConfigurationContract common.Address
	IsDeprecated          bool
	Metadata              []byte
}

type CapabilitiesRegistryDONInfo struct {
	Id                       uint32
	ConfigCount              uint32
	F                        uint8
	IsPublic                 bool
	AcceptsWorkflows         bool
	NodeP2PIds               [][32]byte
	DonFamily                string
	Name                     string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
}

type CapabilitiesRegistryNewDONParams struct {
	Nodes                    [][32]byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
	IsPublic                 bool
	AcceptsWorkflows         bool
	F                        uint8
	Name                     string
	Config                   []byte
}

type CapabilitiesRegistryNodeOperator struct {
	Admin common.Address
	Name  string
}

type CapabilitiesRegistryNodeParams struct {
	NodeOperatorId      uint32
	Signer              [32]byte
	P2pId               [32]byte
	EncryptionPublicKey [32]byte
	CapabilityIds       []string
}

type INodeInfoProviderNodeInfo struct {
	NodeOperatorId      uint32
	ConfigCount         uint32
	WorkflowDONId       uint32
	Signer              [32]byte
	P2pId               [32]byte
	EncryptionPublicKey [32]byte
	CapabilityIds       []string
	CapabilitiesDONIds  []*big.Int
}

var CapabilitiesRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCapabilities\",\"inputs\":[{\"name\":\"capabilities\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.Capability[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addDONs\",\"inputs\":[{\"name\":\"newDONs\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NewDONParams[]\",\"components\":[{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodeOperators\",\"inputs\":[{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deprecateCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCapabilities\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapability\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityConfigs\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.DONInfo[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONsInFamily\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHistoricalDONInfo\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextDONId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nodeInfo\",\"type\":\"tuple\",\"internalType\":\"structINodeInfoProvider.NodeInfo\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperator\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodesByP2PIds\",\"inputs\":[{\"name\":\"p2pIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDONNameTaken\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDONs\",\"inputs\":[{\"name\":\"donIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDONsByName\",\"inputs\":[{\"name\":\"donNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodes\",\"inputs\":[{\"name\":\"removedNodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONFamily\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"additionalParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.AdditionalDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"additionalParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.AdditionalDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CapabilityConfigured\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONFamilySet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeAdded\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorAdded\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorRemoved\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorUpdated\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeRemoved\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CapabilityAlreadyExists\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityDoesNotExist\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityIsDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityRequiredByDON\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONConfigDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"requestedConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONNameAlreadyTaken\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DONWithNameDoesNotExist\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONCapability\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONNode\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidCapabilityConfigurationContractInterface\",\"inputs\":[{\"name\":\"proposedConfigurationContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidFaultTolerance\",\"inputs\":[{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nodeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeEncryptionPublicKey\",\"inputs\":[{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeOperatorAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNodeP2PId\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"lengthOne\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lengthTwo\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyExists\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotExist\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotSupportCapability\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfCapabilitiesDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfWorkflowDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x608080604052346088573315604657600080546001600160a01b03191633179055601180546001600160401b03191664010000000117905560405161602c908161008e8239f35b62461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f00000000000000006044820152606490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80628375c61461023657806305a5196614610231578063106e9a081461022c578063181f5a77146102275780631d05394c14610222578063214502431461021d57806322bdbcbc146102185780632353740514610213578063275459f21461020e5780632af97674146102095780632c01a1e814610204578063398f3773146101ff57806350c946fe146101fa57806351b10006146101f5578063543f4025146101f057806359003602146101eb57806359110666146101e6578063664fe797146101e157806366acaa33146101dc57806379ba5097146101d757806386fa4246146101d257806388ea09ee146101cd5780638da5cb5b146101c8578063a580ef30146101c3578063a7eb612c146101be578063b2246c49146101b9578063b8521761146101b4578063bfa8eef5146101af578063c9315179146101aa578063cd71fd09146101a5578063ddbe4f82146101a0578063e29581aa1461019b578063f2fde38b146101965763fcdc8efe1461019157600080fd5b61305b565b612f1b565b612e07565b612d1a565b612c2e565b612ba1565b612a95565b612a2c565b6129bf565b612477565b612458565b612406565b612312565b611fae565b611e53565b611d5b565b611c10565b611b69565b611b3c565b611ad8565b6119a7565b61190e565b611743565b61154f565b611392565b611286565b6111cd565b61110e565b610f7d565b610d6b565b610cee565b6106d6565b610624565b6102be565b9181601f8401121561026c5782359167ffffffffffffffff831161026c576020808501948460051b01011161026c57565b600080fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261026c576004359067ffffffffffffffff821161026c576102ba9160040161023b565b9091565b3461026c576102cc36610271565b906102d5614933565b60005b8281106102e157005b6102f46102ef8285856130ce565b613113565b6103056103018251614a2b565b1590565b6103d65760208101805173ffffffffffffffffffffffffffffffffffffffff1680610378575b50509061034d826103478161034260019651613195565b613465565b516135b9565b7fe671cf109707667795a875c19f031bdbc7ed40a130f6dc18a55615a0e0099fbb600080a2016102d8565b61030161038491614ab9565b61038e578061032b565b517fabb5e3fd0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6000fd5b61040d90516040519182917f8f51ece800000000000000000000000000000000000000000000000000000000835260048301610cdd565b0390fd5b60005b8381106104245750506000910152565b8181015183820152602001610414565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361047081518092818752878088019101610411565b0116010190565b9080602083519182815201916020808360051b8301019401926000915b8383106104a357505050505090565b90919293946020806104df837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610434565b97019301930191939290610494565b906020808351928381520192019060005b81811061050c5750505090565b82518452602093840193909201916001016104ff565b6105a39163ffffffff82511681526105476020830151602083019063ffffffff169052565b60408281015163ffffffff1690820152606082015160608201526080820151608082015260a082015160a082015260e061059260c084015161010060c0850152610100840190610477565b9201519060e08184039101526104ee565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106105d957505050505090565b9091929394602080610615837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610522565b970193019301919392906105ca565b3461026c5761063236610271565b61063b81613645565b9060005b818110610658576040518061065485826105a6565b0390f35b61066c6106668284876136b2565b35613d4b565b61067682856136c2565b5261068181846136c2565b50608061068e82856136c2565b5101511561069e5760010161063f565b6106a891846136b2565b357fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461026c576106e436610271565b61071f61070673ffffffffffffffffffffffffffffffffffffffff6000541690565b73ffffffffffffffffffffffffffffffffffffffff1690565b33149060009115925b81831061073157005b61074461073f8484846136d6565b613716565b9261077461076f610759865163ffffffff1690565b63ffffffff16600052600c602052604060002090565b6138ae565b610795610706825173ffffffffffffffffffffffffffffffffffffffff1690565b15610b3d57859081610b13575b50610ae55760408401936107c18551600052600d602052604060002090565b9460018601908154610ab75780518015610a8a5750602083019182518015908115610a6c575b50610a4257606084019788518015610a1557506080850151998a51156109e1579661085e610825610820845463ffffffff9060201c1690565b6139a6565b83547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff1660209190911b67ffffffff0000000016178355565b815460201c63ffffffff1698600483019860005b8d51811015610907576108926103018f8361088c916136c2565b51614c23565b6108d157806108ca8f6108c38f948f6001966108bd919063ffffffff16600052602052604060002090565b926136c2565b5190614a89565b5001610872565b61040d8e6040519182917fdb68f0c2000000000000000000000000000000000000000000000000000000008352600483016138f9565b509a63ffffffff95919c506109b8929a50600199507f74becb12a5e8fd0e98077d02dfba8f647c9670c9df177e42c2418cf17a636f0596946109d5946109a09251600382015561098f61095e8b5163ffffffff1690565b829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b600284519101558551809155615a4c565b506109ab8151615ad7565b5051955163ffffffff1690565b915160405193849316958360209093929193604081019481520152565b0390a201919290610728565b6040517fdb68f0c20000000000000000000000000000000000000000000000000000000081528061040d8d600483016138f9565b7f37d897650000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f837731460000000000000000000000000000000000000000000000000000000060005260046000fd5b610a8491506000526009602052604060002054151590565b386107e7565b7f64e2ee920000000000000000000000000000000000000000000000000000000060005260045260246000fd5b517f546184830000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f9473075d000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b51610b34915073ffffffffffffffffffffffffffffffffffffffff16610706565b331415386107a2565b6103d2610b4e865163ffffffff1690565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602490565b600091031261026c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117610bd657604052565b610b8b565b60a0810190811067ffffffffffffffff821117610bd657604052565b6080810190811067ffffffffffffffff821117610bd657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610bd657604052565b60405190610c63604083610c13565b565b60405190610c6361010083610c13565b60405190610c6360c083610c13565b60405190610c6360e083610c13565b60405190610c6361014083610c13565b67ffffffffffffffff8111610bd657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9060206105a3928181520190610434565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c576106546040805190610d2f8183610c13565b601a82527f4361706162696c6974696573526567697374727920312e312e30000000000000602083015251918291602083526020830190610434565b3461026c57610d7936610271565b90610d82614933565b60005b828110610d8e57005b80610dae610d9f60019386866136b2565b35610da9816110bc565b614cf9565b01610d85565b9080602083519182815201916020808360051b8301019401926000915b838310610de057505050505090565b9091929394602080610e3c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519083610e2c8351604084526040840190610434565b9201519084818403910152610434565b97019301930191939290610dd1565b805163ffffffff1682526105a39160208281015163ffffffff169082015260408281015160ff1690820152606082810151151590820152608082810151151590820152610120610eed610ed9610ec7610eb560a087015161014060a08801526101408701906104ee565b60c087015186820360c0880152610434565b60e086015185820360e0870152610434565b610100850151848203610100860152610434565b92015190610120818403910152610db4565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310610f3257505050505090565b9091929394602080610f6e837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610e4b565b97019301930191939290610f23565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760115460201c63ffffffff16610fd5610fd0610fc7836139cc565b63ffffffff1690565b613a64565b60009163ffffffff811660015b8163ffffffff82161061101c576106548486611000610fc7876139cc565b8103611014575b5060405191829182610eff565b815282611007565b611046610fc761103c8363ffffffff16600052600e602052604060002090565b5463ffffffff1690565b611059575b60010163ffffffff16610fe2565b9360016110b363ffffffff926110986110926110858a63ffffffff16600052600e602052604060002090565b5460201c63ffffffff1690565b89614f82565b6110a282896136c2565b526110ad81886136c2565b50613ad1565b9591505061104b565b63ffffffff81160361026c57565b90604060206105a39373ffffffffffffffffffffffffffffffffffffffff81511684520151918160208201520190610434565b9060206105a39281815201906110ca565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5763ffffffff60043561114e816110bc565b611156613afe565b5016600052600c602052610654604060002060016111ab6040519261117a84610bba565b73ffffffffffffffffffffffffffffffffffffffff81541684526111a460405180948193016137f1565b0382610c13565b6020820152604051918291826110fd565b9060206105a3928181520190610e4b565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57600435611208816110bc565b611210613a03565b5063ffffffff81169081600052600e60205263ffffffff60406000205460201c16918215611259576106546112458484614f82565b604051918291602083526020830190610e4b565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461026c5761129436610271565b61129c614933565b60005b63ffffffff811690828210156113195763ffffffff6112c26113149385876136b2565b356112cc816110bc565b1680600052600c6020526112ea600160406000206000815501613b18565b7fa59268ca81d40429e65ccea5385b59cf2d3fc6519371dee92f8eb1dae5107a7a600080a26139a6565b61129f565b005b92919261132782610ca3565b916113356040519384610c13565b82948184528183011161026c578281602093846000960137010152565b9080601f8301121561026c578160206105a39335910161131b565b90916113846105a393604084526040840190610434565b916020818403910152610434565b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c576004356113cd816110bc565b60243567ffffffffffffffff811161026c576113ed903690600401611352565b9061145461144f60066114486114166110858663ffffffff16600052600e602052604060002090565b60016114328763ffffffff16600052600e602052604060002090565b019063ffffffff16600052602052604060002090565b018461322d565b613893565b9060609273ffffffffffffffffffffffffffffffffffffffff611496600161147b84613195565b015473ffffffffffffffffffffffffffffffffffffffff1690565b166114ad575b50506106546040519283928361136d565b61150b929350906114ca610706610706600161147b600096613195565b60405180809581947f8318ed5d0000000000000000000000000000000000000000000000000000000083526004830191909163ffffffff6020820193169052565b03915afa90811561154a57600091611527575b5090388061149c565b61154491503d806000833e61153c8183610c13565b810190613b97565b3861151e565b613bf6565b3461026c5761155d36610271565b9061158061070673ffffffffffffffffffffffffffffffffffffffff6000541690565b3314159160005b81811061159057005b61159b8183856136b2565b35906115b182600052600d602052604060002090565b60018101548015611714576005820180546116ce5750815463ffffffff604082901c1680611696575087908161165b575b50610ae55760019361162560027f5254e609a97bab37b7cc79fe128f85c097bd6015c6e1624ae0ba392eb97532059461161d61165295615c29565b500154615cf6565b5061164261163d82600052600d602052604060002090565b613c02565b6040519081529081906020820190565b0390a101611587565b61168d91506116736107596107069263ffffffff1690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b331415386115e2565b7f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602485905260446000fd5b846116df610fc7866103d294615b82565b7f60a6d8980000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602452604490565b7fd82f6adb00000000000000000000000000000000000000000000000000000000600052600484905260246000fd5b3461026c5761175136610271565b9061175a614933565b60005b82811061176657005b611779611774828585613c53565b613c93565b9061179b610706835173ffffffffffffffffffffffffffffffffffffffff1690565b156118e4576001916117b260115463ffffffff1690565b907f78e94ca80be2c30abc061b99e7eb8583b1254781734b1e3ce339abb57da2fe8e73ffffffffffffffffffffffffffffffffffffffff6118c061180a845173ffffffffffffffffffffffffffffffffffffffff1690565b936118606020820195865161183c611820610c54565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015261185b8863ffffffff16600052600c602052604060002090565b613cde565b6118a661187561082060115463ffffffff1690565b63ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000006011541617601155565b5173ffffffffffffffffffffffffffffffffffffffff1690565b9251926118db63ffffffff6040519384931696169482610cdd565b0390a30161175d565b7feeacd9390000000000000000000000000000000000000000000000000000000060005260046000fd5b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5761065461194b600435613d4b565b604051918291602083526020830190610522565b60643590811515820361026c57565b3590811515820361026c57565b6084359060ff8216820361026c57565b359060ff8216820361026c57565b9081604091031261026c5790565b3461026c5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c576004356119e2816110bc565b60243567ffffffffffffffff811161026c57611a0290369060040161023b565b60449291923567ffffffffffffffff811161026c57611a2590369060040161023b565b611a2d61195f565b91611a3661197b565b9360a4359667ffffffffffffffff881161026c57611a5b611319983690600401611999565b96613e9f565b9181601f8401121561026c5782359167ffffffffffffffff831161026c576020838186019501011161026c57565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261026c576004359067ffffffffffffffff821161026c576102ba91600401611a61565b3461026c57611af7611af2611aec36611a8f565b9061409b565b615e9e565b60405180916020820160208352815180915260206040840192019060005b818110611b23575050500390f35b8251845285945060209384019390920191600101611b15565b3461026c576020611b5e611b59611b5236611a8f565b369161131b565b6131bb565b541515604051908152f35b3461026c57611b7736611a8f565b611b7f613a03565b5063ffffffff6040518284823760208184810160028152030190205416918215611bdb57610654611bcf8480600052600e602052611bc9604060002063ffffffff905460201c1690565b90614f82565b604051918291826111bc565b61040d6040519283927f4071db54000000000000000000000000000000000000000000000000000000008452600484016140cd565b3461026c5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760043567ffffffffffffffff811161026c57611c5f903690600401611a61565b9060243567ffffffffffffffff811161026c57611c8090369060040161023b565b60443567ffffffffffffffff811161026c57611ca090369060040161023b565b90611ca961195f565b92611cb261197b565b9460a4359767ffffffffffffffff891161026c57611cd7611319993690600401611999565b97614113565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611d1057505050505090565b9091929394602080611d4c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516110ca565b97019301930191939290611d01565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760115463ffffffff16611da7611da2610fc7836139cc565b6141b3565b60009163ffffffff811660015b8163ffffffff821610611dee576106548486611dd2610fc7876139cc565b8103611de6575b5060405191829182611cdd565b815282611dd9565b611e0e6107066116738363ffffffff16600052600c602052604060002090565b611e21575b60010163ffffffff16611db4565b936001611e4a63ffffffff9261109861076f8963ffffffff16600052600c602052604060002090565b95915050611e13565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5773ffffffffffffffffffffffffffffffffffffffff600154163303611f505760005473ffffffffffffffffffffffffffffffffffffffff16600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055611f117fffffffffffffffffffffffff000000000000000000000000000000000000000060015416600155565b73ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152fd5b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760043567ffffffffffffffff811161026c57611ffd90369060040161023b565b60243567ffffffffffffffff811161026c5761201d90369060040161023b565b928383036122de5760005473ffffffffffffffffffffffffffffffffffffffff169060005b84811061204b57005b61205e6120598287856136b2565b6139c2565b906120798263ffffffff16600052600c602052604060002090565b91612098835473ffffffffffffffffffffffffffffffffffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff841680156122aa576120c4611774858c8b613c53565b906120e6610706835173ffffffffffffffffffffffffffffffffffffffff1690565b156118e4573314158061228a575b610ae55760019473ffffffffffffffffffffffffffffffffffffffff612131610706845173ffffffffffffffffffffffffffffffffffffffff1690565b91161480159061221e575b61214a575b50505001612042565b73ffffffffffffffffffffffffffffffffffffffff6121f8826121e66121a57f86f41145bde5dd7f523305452e4aad3685508c181432ec733d5f345009358a28955173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6118a66020820195898751910161333e565b92519261221363ffffffff6040519384931696169482610cdd565b0390a3388080612141565b5060405160208101906122638161223789870185614220565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610c13565b519020602082015160405161228081612237602082019485610cdd565b519020141561213c565b5073ffffffffffffffffffffffffffffffffffffffff87163314156120f4565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b7fab8b67c6000000000000000000000000000000000000000000000000000000006000526004839052602484905260446000fd5b3461026c5761232036610271565b90612329614933565b60005b82811061233557005b612343611b52828585614231565b61234f61030182614c23565b6123d05761235f61030182614a60565b61239a579061236f6001926135b9565b7fb2553249d353abf34f62139c85f44b5bdeab968ec0ab296a9bf735b75200ed83600080a20161232c565b61040d906040519182917f88c8a73700000000000000000000000000000000000000000000000000000000835260048301610cdd565b61040d906040519182917f4ff9e48a00000000000000000000000000000000000000000000000000000000835260048301610cdd565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57602073ffffffffffffffffffffffffffffffffffffffff60005416604051908152f35b3461026c5761131961246936610271565b90612472614933565b6143e2565b3461026c5761248536610271565b6124a761070660005473ffffffffffffffffffffffffffffffffffffffff1690565b33149060009115925b8183106124b957005b6124c761073f8484846136d6565b9260408401936124e28551600052600d602052604060002090565b916124f761076f610759855463ffffffff1690565b9060018401805492831561299057899081612966575b50610ae5576020840192835115610a4257835180820361292a575b5050506060830180518015610a1557506080840151988951156128f6579561255b610820875463ffffffff9060201c1690565b86547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16602082901b67ffffffff000000001617875597600487019760005b8c5181101561261a576125b461030161088c8f84906136c2565b6125e657806125df8e6108c36001948f8f906108bd919063ffffffff16600052602052604060002090565b500161259a565b6040517fdb68f0c20000000000000000000000000000000000000000000000000000000081528061040d8f600483016138f9565b50929a509897969593949290612638855463ffffffff9060401c1690565b63ffffffff8116612833575b509791989961265560058601615e9e565b9260009a5b845163ffffffff8d16908110156127895761267590866136c2565b5163ffffffff169c8d6126988163ffffffff16600052600e602052604060002090565b600101906126b69063ffffffff16600052600e602052604060002090565b5460201c63ffffffff166126d9919063ffffffff16600052602052604060002090565b6003016126e59061457e565b9b60005b8d5181101561276e576127206103018f8f8f6108bd6127199287929063ffffffff16600052602052604060002090565b5190614c3f565b61272c576001016126e9565b8f90612738908f6136c2565b5161040d6040519283927f16c2b7c4000000000000000000000000000000000000000000000000000000008452600484016145de565b50939d509b61277f91929b506139a6565b9a9b91999061265a565b509b9a50925094600197506128289196507f4b5b465e22eea0c3d40c30e936643245b80d19b2dcf75788c0699fe8d8db645b9460036127d163ffffffff965163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8216178355945197886002840155519101555160405193849316958360209093929193604081019481520152565b0390a20191906124b0565b9a9192989961289960036128938e61287f6110856001612867849e9c9d9e63ffffffff16600052600e602052604060002090565b019263ffffffff16600052600e602052604060002090565b63ffffffff16600052602052604060002090565b0161457e565b9960005b8b518110156128e4576128cd6103018d612719848f8f906108bd919063ffffffff16600052602052604060002090565b6128d95760010161289d565b6127388e918d6136c2565b5093929b509998509493929438612644565b6040517fdb68f0c20000000000000000000000000000000000000000000000000000000081528061040d8c600483016138f9565b612941906000526009602052604060002054151590565b610a42576129529184519055615c29565b5061295d8251615a4c565b50388080612528565b51612987915073ffffffffffffffffffffffffffffffffffffffff16610706565b3314153861250d565b88517fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c576004356129fa816110bc565b6024359067ffffffffffffffff821161026c57612a1e611319923690600401611a61565b91612a27614933565b614733565b3461026c57612a3a36610271565b612a42614933565b60005b818110612a4e57005b8063ffffffff6020612a636001948688614231565b91908260405193849283378101600281520301902054168015612a8f57612a8990614cf9565b01612a45565b50612a89565b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57600435612ad0816110bc565b602435612adc816110bc565b612ae4613a03565b5063ffffffff821680600052600e60205263ffffffff60406000205460201c16156112595750612b2a610fc76110858463ffffffff16600052600e602052604060002090565b63ffffffff821611612b4357611bcf9061065492614f82565b81612b646110856103d29463ffffffff16600052600e602052604060002090565b7ff3c16e2c0000000000000000000000000000000000000000000000000000000060005263ffffffff918216600452811660245216604452606490565b3461026c57602063ffffffff81612bb736611a8f565b91908260405193849283378101600281520301902054161515604051908152f35b6105a3916060612bf18351608084526080840190610434565b9273ffffffffffffffffffffffffffffffffffffffff60208201511660208401526040810151151560408401520151906060818403910152610434565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760043567ffffffffffffffff811161026c57612c88612c83610654923690600401611352565b6148bb565b604051918291602083526020830190612bd8565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310612ccf57505050505090565b9091929394602080612d0b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951612bd8565b97019301930191939290612cc0565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57612d51614501565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612d97612d81846135d9565b93612d8f6040519586610c13565b8085526135d9565b0160005b818110612df057505060005b8151811015612de25780612dc6612dc0600193856136c2565b516148bb565b612dd082866136c2565b52612ddb81856136c2565b5001612da7565b604051806106548582612c9c565b602090612dfb614895565b82828701015201612d9b565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57604051600a548082528160208101600a60005260206000209260005b818110612ebe575050612e6792500382610c13565b612e718151613645565b9060005b8151811015612eb05780612e94612e8e600193856136c2565b51613d4b565b612e9e82866136c2565b52612ea981856136c2565b5001612e75565b6040518061065485826105a6565b8454835260019485019486945060209093019201612e52565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361026c57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361026c57565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5773ffffffffffffffffffffffffffffffffffffffff612f67612ed7565b612f6f614933565b16338114612ffd57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff612fd560005473ffffffffffffffffffffffffffffffffffffffff1690565b167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152fd5b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57602063ffffffff601154821c16604051908152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561310e5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18136030182121561026c570190565b61309f565b60608136031261026c57604051906060820182811067ffffffffffffffff821117610bd657604052803567ffffffffffffffff811161026c576131599036908301611352565b825261316760208201612efa565b602083015260408101359067ffffffffffffffff821161026c5761318d91369101611352565b604082015290565b60206131ae918160405193828580945193849201610411565b8101600381520301902090565b60206131d4918160405193828580945193849201610411565b8101600781520301902090565b60206131fa918160405193828580945193849201610411565b8101601081520301902090565b6020613220918160405193828580945193849201610411565b8101600281520301902090565b602090613247928260405194838680955193849201610411565b82019081520301902090565b90600182811c9216801561329c575b602083101461326d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613262565b916132de918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b8181106132ed575050565b600081556001016132e2565b9190601f811161330857505050565b610c63926000526020600020906020601f840160051c83019310613334575b601f0160051c01906132e2565b9091508190613327565b919091825167ffffffffffffffff8111610bd657613366816133608454613253565b846132f9565b6020601f82116001146133c05781906132de9394956000926133b5575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b015190503880613383565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216906133f384600052602060002090565b9160005b81811061344d57509583600195969710613416575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538808061340c565b9192602060018192868b0151815501940192016133f7565b600260409193929361347885518261333e565b6134da73ffffffffffffffffffffffffffffffffffffffff602087015116600183019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b0192015191825167ffffffffffffffff8111610bd6576134fe816133608454613253565b6020601f821160011461354c5781906132de9394956000926133b55750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061357f84600052602060002090565b9160005b8181106135a15750958360019596971061341657505050811b019055565b9192602060018192868b015181550194019201613583565b6135d190602060405192828480945193849201610411565b810103902090565b67ffffffffffffffff8111610bd65760051b60200190565b60405190610100820182811067ffffffffffffffff821117610bd657604052606060e08360008152600060208201526000604082015260008382015260006080820152600060a08201528260c08201520152565b9061364f826135d9565b61365c6040519182610c13565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061368a82946135d9565b019060005b82811061369b57505050565b6020906136a66135f1565b8282850101520161368f565b919081101561310e5760051b0190565b805182101561310e5760209160051b010190565b919081101561310e5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561026c570190565b60a08136031261026c576040519061372d82610bdb565b8035613738816110bc565b825260208101356020830152604081013560408301526060810135606083015260808101359067ffffffffffffffff821161026c570136601f8201121561026c578035613784816135d9565b916137926040519384610c13565b81835260208084019260051b8201019036821161026c5760208101925b8284106137c25750505050608082015290565b833567ffffffffffffffff811161026c576020916137e68392833691870101611352565b8152019301926137af565b6000929181549161380183613253565b8083529260018116908115613857575060011461381d57505050565b60009081526020812093945091925b83831061383d575060209250010190565b60018160209294939454838587010152019101919061382c565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b90610c636138a792604051938480926137f1565b0383610c13565b90600160206040516138bf81610bba565b6138f5819573ffffffffffffffffffffffffffffffffffffffff81541683526138ee60405180968193016137f1565b0384610c13565b0152565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061392c57505050505090565b9091929394602080613968837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610434565b9701930193019193929061391d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff1663ffffffff81146139bd5760010190565b613977565b356105a3816110bc565b63ffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9116019063ffffffff82116139bd57565b60405190610140820182811067ffffffffffffffff821117610bd657604052606061012083600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201528260e0820152826101008201520152565b90613a6e826135d9565b613a7b6040519182610c13565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613aa982946135d9565b019060005b828110613aba57505050565b602090613ac5613a03565b82828501015201613aae565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146139bd5760010190565b60405190613b0b82610bba565b6060602083600081520152565b613b228154613253565b9081613b2c575050565b81601f60009311600114613b3e575055565b81835260208320613b5a91601f0160051c8101906001016132e2565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c191617905555565b60208183031261026c5780519067ffffffffffffffff821161026c570181601f8201121561026c578051613bca81610ca3565b92613bd86040519485610c13565b8184526020828401011161026c576105a39160208085019101610411565b6040513d6000823e3d90fd5b60059060008155600060018201556000600282015560006003820155018054906000815581613c2f575050565b6000526020600020908101905b818110613c47575050565b60008155600101613c3c565b919081101561310e5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18136030182121561026c570190565b60408136031261026c5760405190613caa82610bba565b613cb381612efa565b825260208101359067ffffffffffffffff821161026c57613cd691369101611352565b602082015290565b60016020610c6393613d4273ffffffffffffffffffffffffffffffffffffffff825116859073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b0151910161333e565b613d536135f1565b50613d6b61103c82600052600d602052604060002090565b906002613d8282600052600d602052604060002090565b01546001613d9a83600052600d602052604060002090565b0154906003613db384600052600d602052604060002090565b015490613def613dea6004613dd287600052600d602052604060002090565b0161287f61108588600052600d602052604060002090565b61457e565b92613e83613e0a61108587600052600d602052604060002090565b95613e76613e516005613e4b613e3a613e2d86600052600d602052604060002090565b5460401c63ffffffff1690565b94600052600d602052604060002090565b01615e9e565b97613e69613e5d610c65565b63ffffffff909b168b52565b63ffffffff1660208a0152565b63ffffffff166040880152565b6060860152608085015260a084015260c083015260e082015290565b95909396919294613eae614933565b63ffffffff8716600052600e6020526040600020549063ffffffff8260201c1663ffffffff811615613f4e57613ee3906139a6565b9160401c60ff165b91613ef4610c75565b63ffffffff909916895263ffffffff166020890152151560408801521515606087015260ff16608086015236613f2991613f82565b60a08501523690613f3992613fe3565b923690613f459261402f565b610c6392615163565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff891660045260246000fd5b919060408382031261026c5760405190613f9b82610bba565b8193803567ffffffffffffffff811161026c5782613fba918301611352565b835260208101359167ffffffffffffffff831161026c57602092613fde9201611352565b910152565b929190613fef816135d9565b93613ffd6040519586610c13565b602085838152019160051b810192831161026c57905b82821061401f57505050565b8135815260209182019101614013565b92919061403b816135d9565b936140496040519586610c13565b602085838152019160051b81019183831161026c5781905b83821061406f575050505050565b813567ffffffffffffffff811161026c576020916140908784938701613f82565b815201910190614061565b6020908260405193849283378101601081520301902090565b6020908260405193849283378101600281520301902090565b90601f836040947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe093602086528160208701528686013760008582860101520116010190565b9691949792939590614123614933565b61413061103c828a6140b4565b9763ffffffff891615611bdb5750506141598763ffffffff16600052600e602052604060002090565b805490919060201c63ffffffff16614170906139a6565b82547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16602082901b67ffffffff0000000016178355915460401c60ff16613eeb565b906141bd826135d9565b6141ca6040519182610c13565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06141f882946135d9565b019060005b82811061420957505050565b602090614214613afe565b828285010152016141fd565b9060206105a39281815201906137f1565b919081101561310e5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561026c57019081359167ffffffffffffffff831161026c57602001823603811361026c579190565b919081101561310e5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218136030182121561026c570190565b9080601f8301121561026c578160206105a393359101613fe3565b9080601f8301121561026c578160206105a39335910161402f565b60e08136031261026c57614319610c84565b90803567ffffffffffffffff811161026c5761433890369083016142d1565b8252602081013567ffffffffffffffff811161026c5761435b90369083016142ec565b602083015261436c6040820161196e565b604083015261437d6060820161196e565b606083015261438e6080820161198b565b608083015260a081013567ffffffffffffffff811161026c576143b49036908301611352565b60a083015260c08101359067ffffffffffffffff821161026c576143da91369101611352565b60c082015290565b9080156144fd5760005b8181106143f857505050565b806144f761441161440c6001948688614291565b614307565b805160208201519061442c60115463ffffffff9060201c1690565b92614473614439856139a6565b7fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff000000006011549260201b16911617601155565b6144ed6144836040830151151590565b6144e36144936060850151151590565b6144da6144a4608087015160ff1690565b9360c060a08801519701516144b7610c54565b97885260208801526144ca613e5d610c75565b600160208b0152151560408a0152565b15156060880152565b60ff166080860152565b60a0840152615163565b016143ec565b5050565b6004549061450e826135d9565b9161451c6040519384610c13565b808352600460009081527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b602085015b8383106145595750505050565b600160208192604051614570816111a481896137f1565b81520192019201919061454c565b90815461458a816135d9565b926145986040519485610c13565b818452602084019060005260206000206000915b8383106145b95750505050565b6001602081926040516145d0816111a481896137f1565b8152019201920191906145ac565b9063ffffffff6145fb602092959495604085526040850190610434565b9416910152565b90929167ffffffffffffffff8111610bd657614622816133608454613253565b6000601f821160011461467b5781906132de9394956000926146705750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b013590503880613383565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216946146ae84600052602060002090565b91805b8781106147065750836001959697106146ce57505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c1991013516905538808061340c565b909260206001819286860135815501940191016146b1565b81604051928392833781016000815203902090565b9163ffffffff6147566110858563ffffffff16600052600e602052604060002090565b16156148615761477961144f8463ffffffff16600052600f602052604060002090565b908151602083012061478c36838661131b565b602081519101201461485b5763ffffffff92826147d6935161483f575b50816147ff576147d16147cc8663ffffffff16600052600f602052604060002090565b613b18565b61471e565b91167f3143778fe5911daca9ced3c73dde73fdb95e7edae557614c40588f3f1a1b44b8600080a3565b614823828261481e8863ffffffff16600052600f602052604060002090565b614602565b614839614830838361409b565b85871690615b2d565b5061471e565b61484b614854916131e1565b85871690615dbd565b50386147a9565b50505050565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b604051906148a282610bf7565b6060808381815260006020820152600060408201520152565b6148c3614895565b506111a461492b60026148d584613195565b0173ffffffffffffffffffffffffffffffffffffffff60016148f686613195565b015416614902856131bb565b541515906040519561491387610bf7565b865260208601526040850152604051928380926137f1565b606082015290565b73ffffffffffffffffffffffffffffffffffffffff60005416330361495457565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152fd5b805482101561310e5760005260206000200190600090565b805468010000000000000000811015610bd6576149ec916001820181556149b2565b9190916149fc57610c639161333e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b614a3660058261322d565b54614a5a57614a468160046149ca565b614a5460056004549261322d565b55600190565b50600090565b614a6b60078261322d565b54614a5a57614a7b8160066149ca565b614a5460076006549261322d565b90614a948183614c3f565b614ab25780614aa6614a5492846149ca565b6001835493019061322d565b5050600090565b60206000604051828101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a700000000000000000000000000000000000000000000000000000000602482015260248152614b1c604482610c13565b519084617530fa903d6000519083614c17575b5082614c0d575b5081614b72575b81614b46575090565b6105a391507f78bea72100000000000000000000000000000000000000000000000000000000906159b9565b905060206000604051828101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff00000000000000000000000000000000000000000000000000000000602482015260248152614bd7604482610c13565b519084617530fa6000513d82614c01575b5081614bf7575b501590614b3d565b9050151538614bef565b60201115915038614be8565b1515915038614b36565b60201115925038614b2f565b6005614c2e9161322d565b54151590565b6007614c2e9161322d565b906001614c2e92019061322d565b60405190816000825492614c6084613253565b9360018116908115614cc15750600114614c85575b5060209250600281520301902090565b9150506000528160206000206000905b838210614ca9575050602091810138614c75565b60209192508060019154848701520191018391614c95565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168452506020938015150283019150389050614c75565b614d138163ffffffff16600052600e602052604060002090565b90815492614d288463ffffffff9060201c1690565b906001840191614d55614d4b82859063ffffffff16600052602052604060002090565b9660401c60ff1690565b9260005b8754811015614ddf576001908515614db957614db3614d8b614d7b838c615b82565b600052600d602052604060002090565b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff8154169055565b01614d59565b614dd96005614dcb614d7b848d615b82565b0163ffffffff891690615dbd565b50614db3565b5094549195509293915060201c63ffffffff1615614861576005614e19614e4793614e1f939063ffffffff16600052602052604060002090565b01614c4d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008154169055565b614e6461144f8263ffffffff16600052600f602052604060002090565b8051614eca575b506000614e888263ffffffff16600052600e602052604060002090565b557ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c158170365163ffffffff60405192169180614ec581906000602083019252565b0390a2565b614ed6614ee3916131e1565b63ffffffff831690615dbd565b50614f016147cc8263ffffffff16600052600f602052604060002090565b38614e6b565b90614f11826135d9565b614f1e6040519182610c13565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0614f4c82946135d9565b019060005b828110614f5d57505050565b602090604051614f6c81610bba565b6060815260608382015282828501015201614f51565b9091614f8c613a03565b50614fa78263ffffffff16600052600e602052604060002090565b91614fc584600185019063ffffffff16600052602052604060002090565b91614fd26003840161457e565b94614fdd8651614f07565b93600681019360005b86518110156150465780896150256150158961500f85615008600198876136c2565b51956136c2565b5161322d565b61501d610c54565b928352613893565b6020820152615034828a6136c2565b5261503f81896136c2565b5001614fe6565b50929650925093549361505c8563ffffffff1690565b9460401c60ff169161507e9063ffffffff16600052600f602052604060002090565b9160048401546150928160ff9060081c1690565b9060ff16906150a086615e9e565b936150a9610c93565b63ffffffff909916895263ffffffff16602089015260ff166040880152151560608701521515608086015260a08501526150e290613893565b60c08401526150f360058201613893565b60e084015260020161510490613893565b61010083015261012082015290565b60ff60019116019060ff82116139bd57565b6105a39054613253565b60409063ffffffff6105a394931681528160208201520190610434565b6040906105a3939281528160208201520190610434565b9190600161518e615178845163ffffffff1690565b63ffffffff16600052600e602052604060002090565b0190602083016151bb6151a5825163ffffffff1690565b849063ffffffff16600052602052604060002090565b9260808501906151cc825160ff1690565b60ff8116159081156159a2575b5061595d576151f89061287f6151f3855163ffffffff1690565b6139cc565b9161520560058401613893565b602081519101209260a0870193845151602081519101200361589e575b6001615235610fc7845163ffffffff1690565b11615837575b5060005b87518110156154035761525f610301615258838b6136c2565b5188615b2d565b6153b5576060870151156153825761528e613e2d61527d838b6136c2565b51600052600d602052604060002090565b63ffffffff6152a4610fc78a5163ffffffff1690565b9116141580615364575b615316576001906153106152c6895163ffffffff1690565b6152d361527d848d6136c2565b907fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff6bffffffff000000000000000083549260401b169116179055565b0161523f565b8761532c6103d2926108bd8a5163ffffffff1690565b517f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff909116600452602452604490565b5063ffffffff61537a613e2d61527d848c6136c2565b1615156152ae565b806153af600561539761527d6001958d6136c2565b016153a9610fc78b5163ffffffff1690565b90615b2d565b50615310565b876153cb6103d2926108bd8a5163ffffffff1690565b517f636e40570000000000000000000000000000000000000000000000000000000060005263ffffffff909116600452602452604490565b509295949093919560009460068401936003810198600282016004600584019301945b80518a10156157cd576103019c6154496154408c846136c2565b519e8f51614c23565b615796576154578e51614c34565b61575f5761546f61546a8f8b905161322d565b615125565b6157185760005b8c51811015615504576154b48f610301908f8061287f61108561527d8860046154a561527d836154ac996136c2565b01946136c2565b905190614c3f565b6154c057600101615476565b6154cb8f918e6136c2565b5190519061040d6040519283927f4b5786e70000000000000000000000000000000000000000000000000000000084526004840161514c565b50999091809d8d8151615516916149ca565b602081019182518b83519061552a9161322d565b906155349161333e565b8551602001516155439161333e565b845151615550908761333e565b6040890151151561558c90899060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b865160ff1688547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1660089190911b61ff0016178855885163ffffffff166155e48163ffffffff16600052600e602052604060002090565b90615619919063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b60608901511515895163ffffffff166156429063ffffffff16600052600e602052604060002090565b80547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1691151560401b68ff0000000000000000169190911790558a5163ffffffff16895163ffffffff166156a79063ffffffff16600052600e602052604060002090565b906156e591907fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff0000000083549260201b169116179055565b88518d9063ffffffff16928c516156ff9063ffffffff1690565b925190519261570d94615f4e565b600101989b90615426565b8d615727895163ffffffff1690565b90519061040d6040519283927f368812ac0000000000000000000000000000000000000000000000000000000084526004840161512f565b61040d8e516040519182917f88c8a73700000000000000000000000000000000000000000000000000000000835260048301610cdd565b61040d8e516040519182917f4ff9e48a00000000000000000000000000000000000000000000000000000000835260048301610cdd565b50505050505061581d9295507ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c1581703651945063ffffffff935061581291505163ffffffff1690565b935163ffffffff1690565b60405163ffffffff90911681529216918060208101614ec5565b9794929096939160005b895481101561588f5780615889614d8b614d7b838e6158838e61587d610fc76005615871614d7b60019d88615b82565b01925163ffffffff1690565b90615dbd565b50615b82565b01615841565b5091939690929497503861523b565b6158ad614e1f60058801614c4d565b83515180516158bd575b50615222565b6158ce61103c63ffffffff92613207565b166159255761591f6158e4885163ffffffff1690565b6158ef865151613207565b9063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b386158b7565b61040d8451516040519182917f07bf02d600000000000000000000000000000000000000000000000000000000835260048301610cdd565b6103d28761596c845160ff1690565b90517f25b4d6180000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b6159ac9150615113565b60ff8851911611386151d9565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a700000000000000000000000000000000000000000000000000000000845216602482015260248152615a1f604482610c13565b5191617530fa6000513d82615a40575b5081615a39575090565b9050151590565b60201115915038615a2f565b600081815260096020526040902054614a5a5760085468010000000000000000811015610bd657615abe615a8982600185940160085560086149b2565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600854906000526009602052604060002055600190565b6000818152600b6020526040902054614a5a57600a5468010000000000000000811015610bd657615b14615a89826001859401600a55600a6149b2565b9055600a5490600052600b602052604060002055600190565b6000828152600182016020526040902054614ab25780549068010000000000000000821015610bd65782615b6b615a898460018096018555846149b2565b905580549260005201602052604060002055600190565b90615b8c916149b2565b90549060031b1c90565b80548015615bfa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615bcb82826149b2565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260096020526040902054908115614ab2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116139bd57600854927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116139bd578383600095614a549503615cc5575b505050615cb46008615b96565b600990600052602052604060002090565b615cb4615ce791615cdd615b8c615ced9560086149b2565b92839160086149b2565b906132a6565b55388080615ca7565b6000818152600b6020526040902054908115614ab2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116139bd57600a54927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116139bd578383600095614a549503615d92575b505050615d81600a615b96565b600b90600052602052604060002090565b615d81615ce791615daa615b8c615db495600a6149b2565b928391600a6149b2565b55388080615d74565b6001810191806000528260205260406000205492831515600014615e95577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116139bd578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116139bd576000958583614a5497615e4d9503615e5c575b505050615b96565b90600052602052604060002090565b615e7c615ce791615e73615b8c615e8c95886149b2565b928391876149b2565b8590600052602052604060002090565b55388080615e45565b50505050600090565b906040519182815491828252602082019060005260206000209260005b818110615ed0575050610c6392500383610c13565b8454835260019485019487945060209093019201615ebb565b9294939160808401608085528251809152602060a0860193019060005b818110615f38575050509163ffffffff615f2b83606095878496036020890152610434565b9616604085015216910152565b8251855260209485019490920191600101615f06565b93909173ffffffffffffffffffffffffffffffffffffffff6001615f7183613195565b015416615f80575b5050505050565b73ffffffffffffffffffffffffffffffffffffffff806001615fa28294613195565b015416161690813b1561026c5760008094615fec604051978896879586947ffba64a7c00000000000000000000000000000000000000000000000000000000865260048601615ee9565b03925af1801561154a57616004575b80808080615f79565b80616013600061601993610c13565b80610b80565b38615ffb56fea164736f6c634300081a000a",
}

var CapabilitiesRegistryABI = CapabilitiesRegistryMetaData.ABI

var CapabilitiesRegistryBin = CapabilitiesRegistryMetaData.Bin

func DeployCapabilitiesRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CapabilitiesRegistry, error) {
	parsed, err := CapabilitiesRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CapabilitiesRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CapabilitiesRegistry{address: address, abi: *parsed, CapabilitiesRegistryCaller: CapabilitiesRegistryCaller{contract: contract}, CapabilitiesRegistryTransactor: CapabilitiesRegistryTransactor{contract: contract}, CapabilitiesRegistryFilterer: CapabilitiesRegistryFilterer{contract: contract}}, nil
}

type CapabilitiesRegistry struct {
	address common.Address
	abi     abi.ABI
	CapabilitiesRegistryCaller
	CapabilitiesRegistryTransactor
	CapabilitiesRegistryFilterer
}

type CapabilitiesRegistryCaller struct {
	contract *bind.BoundContract
}

type CapabilitiesRegistryTransactor struct {
	contract *bind.BoundContract
}

type CapabilitiesRegistryFilterer struct {
	contract *bind.BoundContract
}

type CapabilitiesRegistrySession struct {
	Contract     *CapabilitiesRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CapabilitiesRegistryCallerSession struct {
	Contract *CapabilitiesRegistryCaller
	CallOpts bind.CallOpts
}

type CapabilitiesRegistryTransactorSession struct {
	Contract     *CapabilitiesRegistryTransactor
	TransactOpts bind.TransactOpts
}

type CapabilitiesRegistryRaw struct {
	Contract *CapabilitiesRegistry
}

type CapabilitiesRegistryCallerRaw struct {
	Contract *CapabilitiesRegistryCaller
}

type CapabilitiesRegistryTransactorRaw struct {
	Contract *CapabilitiesRegistryTransactor
}

func NewCapabilitiesRegistry(address common.Address, backend bind.ContractBackend) (*CapabilitiesRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(CapabilitiesRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCapabilitiesRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistry{address: address, abi: abi, CapabilitiesRegistryCaller: CapabilitiesRegistryCaller{contract: contract}, CapabilitiesRegistryTransactor: CapabilitiesRegistryTransactor{contract: contract}, CapabilitiesRegistryFilterer: CapabilitiesRegistryFilterer{contract: contract}}, nil
}

func NewCapabilitiesRegistryCaller(address common.Address, caller bind.ContractCaller) (*CapabilitiesRegistryCaller, error) {
	contract, err := bindCapabilitiesRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCaller{contract: contract}, nil
}

func NewCapabilitiesRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CapabilitiesRegistryTransactor, error) {
	contract, err := bindCapabilitiesRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryTransactor{contract: contract}, nil
}

func NewCapabilitiesRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CapabilitiesRegistryFilterer, error) {
	contract, err := bindCapabilitiesRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryFilterer{contract: contract}, nil
}

func bindCapabilitiesRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CapabilitiesRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CapabilitiesRegistry.Contract.CapabilitiesRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.CapabilitiesRegistryTransactor.contract.Transfer(opts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.CapabilitiesRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CapabilitiesRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.contract.Transfer(opts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapabilities(opts *bind.CallOpts) ([]CapabilitiesRegistryCapabilityInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapabilities")

	if err != nil {
		return *new([]CapabilitiesRegistryCapabilityInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]CapabilitiesRegistryCapabilityInfo)).(*[]CapabilitiesRegistryCapabilityInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapabilities() ([]CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilities(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapabilities() ([]CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilities(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapability(opts *bind.CallOpts, capabilityId string) (CapabilitiesRegistryCapabilityInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapability", capabilityId)

	if err != nil {
		return *new(CapabilitiesRegistryCapabilityInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryCapabilityInfo)).(*CapabilitiesRegistryCapabilityInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapability(capabilityId string) (CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapability(&_CapabilitiesRegistry.CallOpts, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapability(capabilityId string) (CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapability(&_CapabilitiesRegistry.CallOpts, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapabilityConfigs(opts *bind.CallOpts, donId uint32, capabilityId string) ([]byte, []byte, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapabilityConfigs", donId, capabilityId)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapabilityConfigs(donId uint32, capabilityId string) ([]byte, []byte, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilityConfigs(&_CapabilitiesRegistry.CallOpts, donId, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapabilityConfigs(donId uint32, capabilityId string) ([]byte, []byte, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilityConfigs(&_CapabilitiesRegistry.CallOpts, donId, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDON(opts *bind.CallOpts, donId uint32) (CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDON", donId)

	if err != nil {
		return *new(CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryDONInfo)).(*CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDON(donId uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDON(&_CapabilitiesRegistry.CallOpts, donId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDON(donId uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDON(&_CapabilitiesRegistry.CallOpts, donId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONByName(opts *bind.CallOpts, donName string) (CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONByName", donName)

	if err != nil {
		return *new(CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryDONInfo)).(*CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONByName(donName string) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONByName(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONByName(donName string) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONByName(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONs(opts *bind.CallOpts) ([]CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONs")

	if err != nil {
		return *new([]CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]CapabilitiesRegistryDONInfo)).(*[]CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONs() ([]CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONs(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONs() ([]CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONs(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONsInFamily(opts *bind.CallOpts, donFamily string) ([]*big.Int, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONsInFamily", donFamily)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONsInFamily(donFamily string) ([]*big.Int, error) {
	return _CapabilitiesRegistry.Contract.GetDONsInFamily(&_CapabilitiesRegistry.CallOpts, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONsInFamily(donFamily string) ([]*big.Int, error) {
	return _CapabilitiesRegistry.Contract.GetDONsInFamily(&_CapabilitiesRegistry.CallOpts, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetHistoricalDONInfo(opts *bind.CallOpts, donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getHistoricalDONInfo", donId, configCount)

	if err != nil {
		return *new(CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryDONInfo)).(*CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetHistoricalDONInfo(donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetHistoricalDONInfo(&_CapabilitiesRegistry.CallOpts, donId, configCount)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetHistoricalDONInfo(donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetHistoricalDONInfo(&_CapabilitiesRegistry.CallOpts, donId, configCount)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNextDONId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNextDONId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNextDONId() (uint32, error) {
	return _CapabilitiesRegistry.Contract.GetNextDONId(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNextDONId() (uint32, error) {
	return _CapabilitiesRegistry.Contract.GetNextDONId(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNode(opts *bind.CallOpts, p2pId [32]byte) (INodeInfoProviderNodeInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNode", p2pId)

	if err != nil {
		return *new(INodeInfoProviderNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(INodeInfoProviderNodeInfo)).(*INodeInfoProviderNodeInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNode(p2pId [32]byte) (INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNode(&_CapabilitiesRegistry.CallOpts, p2pId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNode(p2pId [32]byte) (INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNode(&_CapabilitiesRegistry.CallOpts, p2pId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodeOperator(opts *bind.CallOpts, nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodeOperator", nodeOperatorId)

	if err != nil {
		return *new(CapabilitiesRegistryNodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryNodeOperator)).(*CapabilitiesRegistryNodeOperator)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodeOperator(nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperator(&_CapabilitiesRegistry.CallOpts, nodeOperatorId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodeOperator(nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperator(&_CapabilitiesRegistry.CallOpts, nodeOperatorId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodeOperators(opts *bind.CallOpts) ([]CapabilitiesRegistryNodeOperator, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodeOperators")

	if err != nil {
		return *new([]CapabilitiesRegistryNodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([]CapabilitiesRegistryNodeOperator)).(*[]CapabilitiesRegistryNodeOperator)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodeOperators() ([]CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperators(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodeOperators() ([]CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperators(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodes(opts *bind.CallOpts) ([]INodeInfoProviderNodeInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodes")

	if err != nil {
		return *new([]INodeInfoProviderNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]INodeInfoProviderNodeInfo)).(*[]INodeInfoProviderNodeInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodes() ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodes(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodes() ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodes(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodesByP2PIds(opts *bind.CallOpts, p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodesByP2PIds", p2pIds)

	if err != nil {
		return *new([]INodeInfoProviderNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]INodeInfoProviderNodeInfo)).(*[]INodeInfoProviderNodeInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodesByP2PIds(p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodesByP2PIds(&_CapabilitiesRegistry.CallOpts, p2pIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodesByP2PIds(p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodesByP2PIds(&_CapabilitiesRegistry.CallOpts, p2pIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) IsCapabilityDeprecated(opts *bind.CallOpts, capabilityId string) (bool, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "isCapabilityDeprecated", capabilityId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) IsCapabilityDeprecated(capabilityId string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsCapabilityDeprecated(&_CapabilitiesRegistry.CallOpts, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) IsCapabilityDeprecated(capabilityId string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsCapabilityDeprecated(&_CapabilitiesRegistry.CallOpts, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) IsDONNameTaken(opts *bind.CallOpts, donName string) (bool, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "isDONNameTaken", donName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) IsDONNameTaken(donName string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsDONNameTaken(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) IsDONNameTaken(donName string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsDONNameTaken(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) Owner() (common.Address, error) {
	return _CapabilitiesRegistry.Contract.Owner(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) Owner() (common.Address, error) {
	return _CapabilitiesRegistry.Contract.Owner(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) TypeAndVersion() (string, error) {
	return _CapabilitiesRegistry.Contract.TypeAndVersion(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) TypeAndVersion() (string, error) {
	return _CapabilitiesRegistry.Contract.TypeAndVersion(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AcceptOwnership(&_CapabilitiesRegistry.TransactOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AcceptOwnership(&_CapabilitiesRegistry.TransactOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddCapabilities(opts *bind.TransactOpts, capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addCapabilities", capabilities)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddCapabilities(capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilities)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddCapabilities(capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilities)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddDONs(opts *bind.TransactOpts, newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addDONs", newDONs)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddDONs(newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddDONs(&_CapabilitiesRegistry.TransactOpts, newDONs)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddDONs(newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddDONs(&_CapabilitiesRegistry.TransactOpts, newDONs)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddNodeOperators(opts *bind.TransactOpts, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addNodeOperators", nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddNodeOperators(nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddNodeOperators(nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addNodes", nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) DeprecateCapabilities(opts *bind.TransactOpts, capabilityIds []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "deprecateCapabilities", capabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) DeprecateCapabilities(capabilityIds []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.DeprecateCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) DeprecateCapabilities(capabilityIds []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.DeprecateCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveDONs(opts *bind.TransactOpts, donIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeDONs", donIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveDONs(donIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONs(&_CapabilitiesRegistry.TransactOpts, donIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveDONs(donIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONs(&_CapabilitiesRegistry.TransactOpts, donIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveDONsByName(opts *bind.TransactOpts, donNames []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeDONsByName", donNames)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveDONsByName(donNames []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONsByName(&_CapabilitiesRegistry.TransactOpts, donNames)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveDONsByName(donNames []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONsByName(&_CapabilitiesRegistry.TransactOpts, donNames)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeNodeOperators", nodeOperatorIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveNodeOperators(nodeOperatorIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveNodeOperators(nodeOperatorIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveNodes(opts *bind.TransactOpts, removedNodeP2PIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeNodes", removedNodeP2PIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveNodes(removedNodeP2PIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodes(&_CapabilitiesRegistry.TransactOpts, removedNodeP2PIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveNodes(removedNodeP2PIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodes(&_CapabilitiesRegistry.TransactOpts, removedNodeP2PIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) SetDONFamily(opts *bind.TransactOpts, donId uint32, donFamily string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "setDONFamily", donId, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) SetDONFamily(donId uint32, donFamily string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamily(&_CapabilitiesRegistry.TransactOpts, donId, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) SetDONFamily(donId uint32, donFamily string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamily(&_CapabilitiesRegistry.TransactOpts, donId, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.TransferOwnership(&_CapabilitiesRegistry.TransactOpts, to)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.TransferOwnership(&_CapabilitiesRegistry.TransactOpts, to)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateDON(opts *bind.TransactOpts, donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateDON", donId, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateDON(donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateDON(donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateDONByName(opts *bind.TransactOpts, donName string, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateDONByName", donName, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateDONByName(donName string, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDONByName(&_CapabilitiesRegistry.TransactOpts, donName, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateDONByName(donName string, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDONByName(&_CapabilitiesRegistry.TransactOpts, donName, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateNodeOperators", nodeOperatorIds, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateNodeOperators(nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateNodeOperators(nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateNodes", nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

type CapabilitiesRegistryCapabilityConfiguredIterator struct {
	Event *CapabilitiesRegistryCapabilityConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryCapabilityConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryCapabilityConfigured)
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
		it.Event = new(CapabilitiesRegistryCapabilityConfigured)
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

func (it *CapabilitiesRegistryCapabilityConfiguredIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryCapabilityConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryCapabilityConfigured struct {
	CapabilityId common.Hash
	Raw          types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterCapabilityConfigured(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityConfiguredIterator, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "CapabilityConfigured", capabilityIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCapabilityConfiguredIterator{contract: _CapabilitiesRegistry.contract, event: "CapabilityConfigured", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchCapabilityConfigured(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityConfigured, capabilityId []string) (event.Subscription, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "CapabilityConfigured", capabilityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryCapabilityConfigured)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityConfigured", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseCapabilityConfigured(log types.Log) (*CapabilitiesRegistryCapabilityConfigured, error) {
	event := new(CapabilitiesRegistryCapabilityConfigured)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryCapabilityDeprecatedIterator struct {
	Event *CapabilitiesRegistryCapabilityDeprecated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryCapabilityDeprecatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryCapabilityDeprecated)
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
		it.Event = new(CapabilitiesRegistryCapabilityDeprecated)
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

func (it *CapabilitiesRegistryCapabilityDeprecatedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryCapabilityDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryCapabilityDeprecated struct {
	CapabilityId common.Hash
	Raw          types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterCapabilityDeprecated(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityDeprecatedIterator, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "CapabilityDeprecated", capabilityIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCapabilityDeprecatedIterator{contract: _CapabilitiesRegistry.contract, event: "CapabilityDeprecated", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchCapabilityDeprecated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityDeprecated, capabilityId []string) (event.Subscription, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "CapabilityDeprecated", capabilityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryCapabilityDeprecated)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityDeprecated", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseCapabilityDeprecated(log types.Log) (*CapabilitiesRegistryCapabilityDeprecated, error) {
	event := new(CapabilitiesRegistryCapabilityDeprecated)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryConfigSetIterator struct {
	Event *CapabilitiesRegistryConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryConfigSet)
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
		it.Event = new(CapabilitiesRegistryConfigSet)
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

func (it *CapabilitiesRegistryConfigSetIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryConfigSet struct {
	DonId       uint32
	ConfigCount uint32
	Raw         types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterConfigSet(opts *bind.FilterOpts, donId []uint32) (*CapabilitiesRegistryConfigSetIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "ConfigSet", donIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryConfigSetIterator{contract: _CapabilitiesRegistry.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryConfigSet, donId []uint32) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "ConfigSet", donIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryConfigSet)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseConfigSet(log types.Log) (*CapabilitiesRegistryConfigSet, error) {
	event := new(CapabilitiesRegistryConfigSet)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryDONFamilySetIterator struct {
	Event *CapabilitiesRegistryDONFamilySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryDONFamilySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryDONFamilySet)
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
		it.Event = new(CapabilitiesRegistryDONFamilySet)
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

func (it *CapabilitiesRegistryDONFamilySetIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryDONFamilySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryDONFamilySet struct {
	DonId     uint32
	DonFamily common.Hash
	Raw       types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterDONFamilySet(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONFamilySetIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "DONFamilySet", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryDONFamilySetIterator{contract: _CapabilitiesRegistry.contract, event: "DONFamilySet", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchDONFamilySet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONFamilySet, donId []uint32, donFamily []string) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "DONFamilySet", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryDONFamilySet)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONFamilySet", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseDONFamilySet(log types.Log) (*CapabilitiesRegistryDONFamilySet, error) {
	event := new(CapabilitiesRegistryDONFamilySet)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONFamilySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeAddedIterator struct {
	Event *CapabilitiesRegistryNodeAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeAdded)
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
		it.Event = new(CapabilitiesRegistryNodeAdded)
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

func (it *CapabilitiesRegistryNodeAddedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeAdded struct {
	P2pId          [32]byte
	NodeOperatorId uint32
	Signer         [32]byte
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeAdded(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeAddedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeAdded, nodeOperatorId []uint32) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeAdded)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeAdded(log types.Log) (*CapabilitiesRegistryNodeAdded, error) {
	event := new(CapabilitiesRegistryNodeAdded)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeOperatorAddedIterator struct {
	Event *CapabilitiesRegistryNodeOperatorAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeOperatorAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeOperatorAdded)
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
		it.Event = new(CapabilitiesRegistryNodeOperatorAdded)
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

func (it *CapabilitiesRegistryNodeOperatorAddedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeOperatorAdded struct {
	NodeOperatorId uint32
	Admin          common.Address
	Name           string
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeOperatorAddedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorAdded, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeOperatorAdded)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeOperatorAdded(log types.Log) (*CapabilitiesRegistryNodeOperatorAdded, error) {
	event := new(CapabilitiesRegistryNodeOperatorAdded)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeOperatorRemovedIterator struct {
	Event *CapabilitiesRegistryNodeOperatorRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeOperatorRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeOperatorRemoved)
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
		it.Event = new(CapabilitiesRegistryNodeOperatorRemoved)
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

func (it *CapabilitiesRegistryNodeOperatorRemovedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeOperatorRemoved struct {
	NodeOperatorId uint32
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeOperatorRemoved(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeOperatorRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeOperatorRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeOperatorRemovedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeOperatorRemoved", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeOperatorRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorRemoved, nodeOperatorId []uint32) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeOperatorRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeOperatorRemoved)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorRemoved", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeOperatorRemoved(log types.Log) (*CapabilitiesRegistryNodeOperatorRemoved, error) {
	event := new(CapabilitiesRegistryNodeOperatorRemoved)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeOperatorUpdatedIterator struct {
	Event *CapabilitiesRegistryNodeOperatorUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeOperatorUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeOperatorUpdated)
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
		it.Event = new(CapabilitiesRegistryNodeOperatorUpdated)
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

func (it *CapabilitiesRegistryNodeOperatorUpdatedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeOperatorUpdated struct {
	NodeOperatorId uint32
	Admin          common.Address
	Name           string
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeOperatorUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorUpdatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeOperatorUpdated", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeOperatorUpdatedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeOperatorUpdated", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeOperatorUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorUpdated, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeOperatorUpdated", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeOperatorUpdated)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeOperatorUpdated(log types.Log) (*CapabilitiesRegistryNodeOperatorUpdated, error) {
	event := new(CapabilitiesRegistryNodeOperatorUpdated)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeRemovedIterator struct {
	Event *CapabilitiesRegistryNodeRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeRemoved)
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
		it.Event = new(CapabilitiesRegistryNodeRemoved)
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

func (it *CapabilitiesRegistryNodeRemovedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeRemoved struct {
	P2pId [32]byte
	Raw   types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeRemoved(opts *bind.FilterOpts) (*CapabilitiesRegistryNodeRemovedIterator, error) {

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeRemovedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeRemoved) (event.Subscription, error) {

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeRemoved)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeRemoved(log types.Log) (*CapabilitiesRegistryNodeRemoved, error) {
	event := new(CapabilitiesRegistryNodeRemoved)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeUpdatedIterator struct {
	Event *CapabilitiesRegistryNodeUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeUpdated)
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
		it.Event = new(CapabilitiesRegistryNodeUpdated)
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

func (it *CapabilitiesRegistryNodeUpdatedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeUpdated struct {
	P2pId          [32]byte
	NodeOperatorId uint32
	Signer         [32]byte
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeUpdatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeUpdated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeUpdatedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeUpdated", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeUpdated, nodeOperatorId []uint32) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeUpdated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeUpdated)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeUpdated(log types.Log) (*CapabilitiesRegistryNodeUpdated, error) {
	event := new(CapabilitiesRegistryNodeUpdated)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryOwnershipTransferRequestedIterator struct {
	Event *CapabilitiesRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryOwnershipTransferRequested)
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
		it.Event = new(CapabilitiesRegistryOwnershipTransferRequested)
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

func (it *CapabilitiesRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryOwnershipTransferRequestedIterator{contract: _CapabilitiesRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryOwnershipTransferRequested)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*CapabilitiesRegistryOwnershipTransferRequested, error) {
	event := new(CapabilitiesRegistryOwnershipTransferRequested)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryOwnershipTransferredIterator struct {
	Event *CapabilitiesRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryOwnershipTransferred)
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
		it.Event = new(CapabilitiesRegistryOwnershipTransferred)
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

func (it *CapabilitiesRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryOwnershipTransferredIterator{contract: _CapabilitiesRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryOwnershipTransferred)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*CapabilitiesRegistryOwnershipTransferred, error) {
	event := new(CapabilitiesRegistryOwnershipTransferred)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CapabilitiesRegistry.abi.Events["CapabilityConfigured"].ID:
		return _CapabilitiesRegistry.ParseCapabilityConfigured(log)
	case _CapabilitiesRegistry.abi.Events["CapabilityDeprecated"].ID:
		return _CapabilitiesRegistry.ParseCapabilityDeprecated(log)
	case _CapabilitiesRegistry.abi.Events["ConfigSet"].ID:
		return _CapabilitiesRegistry.ParseConfigSet(log)
	case _CapabilitiesRegistry.abi.Events["DONFamilySet"].ID:
		return _CapabilitiesRegistry.ParseDONFamilySet(log)
	case _CapabilitiesRegistry.abi.Events["NodeAdded"].ID:
		return _CapabilitiesRegistry.ParseNodeAdded(log)
	case _CapabilitiesRegistry.abi.Events["NodeOperatorAdded"].ID:
		return _CapabilitiesRegistry.ParseNodeOperatorAdded(log)
	case _CapabilitiesRegistry.abi.Events["NodeOperatorRemoved"].ID:
		return _CapabilitiesRegistry.ParseNodeOperatorRemoved(log)
	case _CapabilitiesRegistry.abi.Events["NodeOperatorUpdated"].ID:
		return _CapabilitiesRegistry.ParseNodeOperatorUpdated(log)
	case _CapabilitiesRegistry.abi.Events["NodeRemoved"].ID:
		return _CapabilitiesRegistry.ParseNodeRemoved(log)
	case _CapabilitiesRegistry.abi.Events["NodeUpdated"].ID:
		return _CapabilitiesRegistry.ParseNodeUpdated(log)
	case _CapabilitiesRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _CapabilitiesRegistry.ParseOwnershipTransferRequested(log)
	case _CapabilitiesRegistry.abi.Events["OwnershipTransferred"].ID:
		return _CapabilitiesRegistry.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CapabilitiesRegistryCapabilityConfigured) Topic() common.Hash {
	return common.HexToHash("0xe671cf109707667795a875c19f031bdbc7ed40a130f6dc18a55615a0e0099fbb")
}

func (CapabilitiesRegistryCapabilityDeprecated) Topic() common.Hash {
	return common.HexToHash("0xb2553249d353abf34f62139c85f44b5bdeab968ec0ab296a9bf735b75200ed83")
}

func (CapabilitiesRegistryConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c1581703651")
}

func (CapabilitiesRegistryDONFamilySet) Topic() common.Hash {
	return common.HexToHash("0x3143778fe5911daca9ced3c73dde73fdb95e7edae557614c40588f3f1a1b44b8")
}

func (CapabilitiesRegistryNodeAdded) Topic() common.Hash {
	return common.HexToHash("0x74becb12a5e8fd0e98077d02dfba8f647c9670c9df177e42c2418cf17a636f05")
}

func (CapabilitiesRegistryNodeOperatorAdded) Topic() common.Hash {
	return common.HexToHash("0x78e94ca80be2c30abc061b99e7eb8583b1254781734b1e3ce339abb57da2fe8e")
}

func (CapabilitiesRegistryNodeOperatorRemoved) Topic() common.Hash {
	return common.HexToHash("0xa59268ca81d40429e65ccea5385b59cf2d3fc6519371dee92f8eb1dae5107a7a")
}

func (CapabilitiesRegistryNodeOperatorUpdated) Topic() common.Hash {
	return common.HexToHash("0x86f41145bde5dd7f523305452e4aad3685508c181432ec733d5f345009358a28")
}

func (CapabilitiesRegistryNodeRemoved) Topic() common.Hash {
	return common.HexToHash("0x5254e609a97bab37b7cc79fe128f85c097bd6015c6e1624ae0ba392eb9753205")
}

func (CapabilitiesRegistryNodeUpdated) Topic() common.Hash {
	return common.HexToHash("0x4b5b465e22eea0c3d40c30e936643245b80d19b2dcf75788c0699fe8d8db645b")
}

func (CapabilitiesRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CapabilitiesRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CapabilitiesRegistry *CapabilitiesRegistry) Address() common.Address {
	return _CapabilitiesRegistry.address
}

type CapabilitiesRegistryInterface interface {
	GetCapabilities(opts *bind.CallOpts) ([]CapabilitiesRegistryCapabilityInfo, error)

	GetCapability(opts *bind.CallOpts, capabilityId string) (CapabilitiesRegistryCapabilityInfo, error)

	GetCapabilityConfigs(opts *bind.CallOpts, donId uint32, capabilityId string) ([]byte, []byte, error)

	GetDON(opts *bind.CallOpts, donId uint32) (CapabilitiesRegistryDONInfo, error)

	GetDONByName(opts *bind.CallOpts, donName string) (CapabilitiesRegistryDONInfo, error)

	GetDONs(opts *bind.CallOpts) ([]CapabilitiesRegistryDONInfo, error)

	GetDONsInFamily(opts *bind.CallOpts, donFamily string) ([]*big.Int, error)

	GetHistoricalDONInfo(opts *bind.CallOpts, donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error)

	GetNextDONId(opts *bind.CallOpts) (uint32, error)

	GetNode(opts *bind.CallOpts, p2pId [32]byte) (INodeInfoProviderNodeInfo, error)

	GetNodeOperator(opts *bind.CallOpts, nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error)

	GetNodeOperators(opts *bind.CallOpts) ([]CapabilitiesRegistryNodeOperator, error)

	GetNodes(opts *bind.CallOpts) ([]INodeInfoProviderNodeInfo, error)

	GetNodesByP2PIds(opts *bind.CallOpts, p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error)

	IsCapabilityDeprecated(opts *bind.CallOpts, capabilityId string) (bool, error)

	IsDONNameTaken(opts *bind.CallOpts, donName string) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddCapabilities(opts *bind.TransactOpts, capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error)

	AddDONs(opts *bind.TransactOpts, newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error)

	AddNodeOperators(opts *bind.TransactOpts, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error)

	AddNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error)

	DeprecateCapabilities(opts *bind.TransactOpts, capabilityIds []string) (*types.Transaction, error)

	RemoveDONs(opts *bind.TransactOpts, donIds []uint32) (*types.Transaction, error)

	RemoveDONsByName(opts *bind.TransactOpts, donNames []string) (*types.Transaction, error)

	RemoveNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32) (*types.Transaction, error)

	RemoveNodes(opts *bind.TransactOpts, removedNodeP2PIds [][32]byte) (*types.Transaction, error)

	SetDONFamily(opts *bind.TransactOpts, donId uint32, donFamily string) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateDON(opts *bind.TransactOpts, donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error)

	UpdateDONByName(opts *bind.TransactOpts, donName string, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error)

	UpdateNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error)

	UpdateNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error)

	FilterCapabilityConfigured(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityConfiguredIterator, error)

	WatchCapabilityConfigured(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityConfigured, capabilityId []string) (event.Subscription, error)

	ParseCapabilityConfigured(log types.Log) (*CapabilitiesRegistryCapabilityConfigured, error)

	FilterCapabilityDeprecated(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityDeprecatedIterator, error)

	WatchCapabilityDeprecated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityDeprecated, capabilityId []string) (event.Subscription, error)

	ParseCapabilityDeprecated(log types.Log) (*CapabilitiesRegistryCapabilityDeprecated, error)

	FilterConfigSet(opts *bind.FilterOpts, donId []uint32) (*CapabilitiesRegistryConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryConfigSet, donId []uint32) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CapabilitiesRegistryConfigSet, error)

	FilterDONFamilySet(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONFamilySetIterator, error)

	WatchDONFamilySet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONFamilySet, donId []uint32, donFamily []string) (event.Subscription, error)

	ParseDONFamilySet(log types.Log) (*CapabilitiesRegistryDONFamilySet, error)

	FilterNodeAdded(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeAddedIterator, error)

	WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeAdded, nodeOperatorId []uint32) (event.Subscription, error)

	ParseNodeAdded(log types.Log) (*CapabilitiesRegistryNodeAdded, error)

	FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorAddedIterator, error)

	WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorAdded, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error)

	ParseNodeOperatorAdded(log types.Log) (*CapabilitiesRegistryNodeOperatorAdded, error)

	FilterNodeOperatorRemoved(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeOperatorRemovedIterator, error)

	WatchNodeOperatorRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorRemoved, nodeOperatorId []uint32) (event.Subscription, error)

	ParseNodeOperatorRemoved(log types.Log) (*CapabilitiesRegistryNodeOperatorRemoved, error)

	FilterNodeOperatorUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorUpdatedIterator, error)

	WatchNodeOperatorUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorUpdated, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error)

	ParseNodeOperatorUpdated(log types.Log) (*CapabilitiesRegistryNodeOperatorUpdated, error)

	FilterNodeRemoved(opts *bind.FilterOpts) (*CapabilitiesRegistryNodeRemovedIterator, error)

	WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeRemoved) (event.Subscription, error)

	ParseNodeRemoved(log types.Log) (*CapabilitiesRegistryNodeRemoved, error)

	FilterNodeUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeUpdatedIterator, error)

	WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeUpdated, nodeOperatorId []uint32) (event.Subscription, error)

	ParseNodeUpdated(log types.Log) (*CapabilitiesRegistryNodeUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CapabilitiesRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CapabilitiesRegistryOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
