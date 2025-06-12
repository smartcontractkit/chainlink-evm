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
	Name                     string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
	Nodes                    [][32]byte
	F                        uint8
	IsPublic                 bool
	AcceptsWorkflows         bool
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

type CapabilitiesRegistryUpdateDONParams struct {
	Name                     string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
	Nodes                    [][32]byte
	F                        uint8
	IsPublic                 bool
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
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCapabilities\",\"inputs\":[{\"name\":\"capabilities\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.Capability[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addDONs\",\"inputs\":[{\"name\":\"newDONs\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NewDONParams[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodeOperators\",\"inputs\":[{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deprecateCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCapabilities\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapability\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityConfigs\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.DONInfo[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONsInFamily\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHistoricalDONInfo\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextDONId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nodeInfo\",\"type\":\"tuple\",\"internalType\":\"structINodeInfoProvider.NodeInfo\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperator\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodesByP2PIds\",\"inputs\":[{\"name\":\"p2pIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDONNameTaken\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDONs\",\"inputs\":[{\"name\":\"donIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDONsByName\",\"inputs\":[{\"name\":\"donNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodes\",\"inputs\":[{\"name\":\"removedNodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONFamily\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updateDONParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.UpdateDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"updateDONParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.UpdateDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CapabilityConfigured\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONFamilySet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeAdded\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorAdded\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorRemoved\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorUpdated\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeRemoved\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CapabilityAlreadyExists\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityDoesNotExist\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityIsDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityRequiredByDON\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONConfigDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"requestedConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONNameAlreadyTaken\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DONWithNameDoesNotExist\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONCapability\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONNode\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidCapabilityConfigurationContractInterface\",\"inputs\":[{\"name\":\"proposedConfigurationContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidFaultTolerance\",\"inputs\":[{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nodeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeEncryptionPublicKey\",\"inputs\":[{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeOperatorAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNodeP2PId\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"lengthOne\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lengthTwo\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyExists\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotExist\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotSupportCapability\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfCapabilitiesDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfWorkflowDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x608080604052346088573315604657600080546001600160a01b03191633179055601280546001600160401b031916640100000001179055604051616112908161008e8239f35b62461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f00000000000000006044820152606490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80628375c61461023657806305a5196614610231578063106e9a081461022c578063181f5a77146102275780631d05394c14610222578063214502431461021d57806322bdbcbc146102185780632353740514610213578063275459f21461020e5780632af97674146102095780632c01a1e814610204578063398f3773146101ff57806350c946fe146101fa57806353a25dd7146101f5578063543f4025146101f057806359003602146101eb57806359110666146101e657806366acaa33146101e157806379ba5097146101dc57806386fa4246146101d757806388ea09ee146101d257806388eafafb146101cd5780638da5cb5b146101c8578063a7eb612c146101c3578063b2246c49146101be578063b5e7508b146101b9578063b8521761146101b4578063bfa8eef5146101af578063c9315179146101aa578063cd71fd09146101a5578063ddbe4f82146101a0578063e29581aa1461019b578063f2fde38b146101965763fcdc8efe1461019157600080fd5b61323d565b6130fd565b612fe9565b612efc565b612e10565b612d83565b612c77565b612c0e565b612bef565b612b82565b6125e9565b612597565b612490565b612389565b612025565b611eca565b611dd2565b611ce2565b611ca2565b611c3e565b6119f6565b611969565b61179e565b6115aa565b6113c7565b6112bb565b611202565b611143565b610fb2565b610da0565b610d23565b61070c565b61065a565b6102be565b9181601f8401121561026c5782359167ffffffffffffffff831161026c576020808501948460051b01011161026c57565b600080fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261026c576004359067ffffffffffffffff821161026c576102ba9160040161023b565b9091565b3461026c576102cc36610271565b906102d5614b39565b60005b8281106102e157005b6102f46102ef8285856132b0565b6132f5565b61030381516020815191012090565b61031361030f82615af6565b1590565b61040c576103358251610330836000526011602052604060002090565b613462565b60208201805173ffffffffffffffffffffffffffffffffffffffff16806103ae575b50508161037d60019493610378610383946000526003602052604060002090565b613589565b51613730565b7fe671cf109707667795a875c19f031bdbc7ed40a130f6dc18a55615a0e0099fbb600080a2016102d8565b61030f6103ba91614bb8565b6103c45780610357565b517fabb5e3fd0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6000fd5b61044382516040519182917f8f51ece800000000000000000000000000000000000000000000000000000000835260048301610d12565b0390fd5b60005b83811061045a5750506000910152565b818101518382015260200161044a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936104a681518092818752878088019101610447565b0116010190565b9080602083519182815201916020808360051b8301019401926000915b8383106104d957505050505090565b9091929394602080610515837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161046a565b970193019301919392906104ca565b906020808351928381520192019060005b8181106105425750505090565b8251845260209384019390920191600101610535565b6105d99163ffffffff825116815261057d6020830151602083019063ffffffff169052565b60408281015163ffffffff1690820152606082015160608201526080820151608082015260a082015160a082015260e06105c860c084015161010060c08501526101008401906104ad565b9201519060e0818403910152610524565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061060f57505050505090565b909192939460208061064b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610558565b97019301930191939290610600565b3461026c5761066836610271565b610671816137bc565b9060005b81811061068e576040518061068a85826105dc565b0390f35b6106a261069c828487613829565b35614071565b6106ac8285613839565b526106b78184613839565b5060806106c48285613839565b510151156106d457600101610675565b6106de9184613829565b357fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461026c5761071a36610271565b61075561073c73ffffffffffffffffffffffffffffffffffffffff6000541690565b73ffffffffffffffffffffffffffffffffffffffff1690565b33149060009115925b81831061076757005b61077a61077584848461384d565b61388d565b926107aa6107a561078f865163ffffffff1690565b63ffffffff16600052600c602052604060002090565b613a25565b6107cb61073c825173ffffffffffffffffffffffffffffffffffffffff1690565b15610b8157859081610b57575b50610b295760408401936107f78551600052600d602052604060002090565b9460018601908154610afb5780518015610ace5750602083019182518015908115610ab0575b50610a8657606084019788518015610a5957506080850151998a5115610a25579661089461085b610856845463ffffffff9060201c1690565b613b1d565b83547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff1660209190911b67ffffffff0000000016178355565b815460201c63ffffffff1698600483019860005b8d5181101561094b576108ca6108bf8f8390613839565b516020815191012090565b6108e461030f826000526005602052604060002054151590565b610915579061090e6001926109098f8f9063ffffffff16600052602052604060002090565b615bd0565b50016108a8565b6104438f6040519182917fdb68f0c200000000000000000000000000000000000000000000000000000000835260048301613a70565b509a63ffffffff95919c506109fc929a50600199507f74becb12a5e8fd0e98077d02dfba8f647c9670c9df177e42c2418cf17a636f059694610a19946109e4925160038201556109d36109a28b5163ffffffff1690565b829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b600284519101558551809155615b31565b506109ef8151615b66565b5051955163ffffffff1690565b915160405193849316958360209093929193604081019481520152565b0390a20191929061075e565b6040517fdb68f0c2000000000000000000000000000000000000000000000000000000008152806104438d60048301613a70565b7f37d897650000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f837731460000000000000000000000000000000000000000000000000000000060005260046000fd5b610ac891506000526009602052604060002054151590565b3861081d565b7f64e2ee920000000000000000000000000000000000000000000000000000000060005260045260246000fd5b517f546184830000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f9473075d000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b51610b78915073ffffffffffffffffffffffffffffffffffffffff1661073c565b331415386107d8565b610408610b92865163ffffffff1690565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602490565b600091031261026c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117610c1a57604052565b610bcf565b60a0810190811067ffffffffffffffff821117610c1a57604052565b6080810190811067ffffffffffffffff821117610c1a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610c1a57604052565b60405190610ca7604083610c57565b565b60405190610ca760e083610c57565b60405190610ca761010083610c57565b60405190610ca761014083610c57565b67ffffffffffffffff8111610c1a57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9060206105d992818152019061046a565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5761068a6040805190610d648183610c57565b601a82527f4361706162696c6974696573526567697374727920322e302e3000000000000060208301525191829160208352602083019061046a565b3461026c57610dae36610271565b90610db7614b39565b60005b828110610dc357005b80610de3610dd46001938686613829565b35610dde816110f1565b614dce565b01610dba565b9080602083519182815201916020808360051b8301019401926000915b838310610e1557505050505090565b9091929394602080610e71837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519083610e61835160408452604084019061046a565b920151908481840391015261046a565b97019301930191939290610e06565b805163ffffffff1682526105d99160208281015163ffffffff169082015260408281015160ff1690820152606082810151151590820152608082810151151590820152610120610f22610f0e610efc610eea60a087015161014060a0880152610140870190610524565b60c087015186820360c088015261046a565b60e086015185820360e087015261046a565b61010085015184820361010086015261046a565b92015190610120818403910152610de9565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310610f6757505050505090565b9091929394602080610fa3837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610e80565b97019301930191939290610f58565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760125460201c63ffffffff1661100a611005610ffc83613b43565b63ffffffff1690565b613bdb565b60009163ffffffff811660015b8163ffffffff8216106110515761068a8486611035610ffc87613b43565b8103611049575b5060405191829182610f34565b81528261103c565b61107b610ffc6110718363ffffffff16600052600e602052604060002090565b5463ffffffff1690565b61108e575b60010163ffffffff16611017565b9360016110e863ffffffff926110cd6110c76110ba8a63ffffffff16600052600e602052604060002090565b5460201c63ffffffff1690565b89615057565b6110d78289613839565b526110e28188613839565b50613c48565b95915050611080565b63ffffffff81160361026c57565b90604060206105d99373ffffffffffffffffffffffffffffffffffffffff8151168452015191816020820152019061046a565b9060206105d99281815201906110ff565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5763ffffffff600435611183816110f1565b61118b613c75565b5016600052600c60205261068a604060002060016111e0604051926111af84610bfe565b73ffffffffffffffffffffffffffffffffffffffff81541684526111d96040518094819301613968565b0382610c57565b602082015260405191829182611132565b9060206105d9928181520190610e80565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760043561123d816110f1565b611245613b7a565b5063ffffffff81169081600052600e60205263ffffffff60406000205460201c1691821561128e5761068a61127a8484615057565b604051918291602083526020830190610e80565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461026c576112c936610271565b6112d1614b39565b60005b63ffffffff8116908282101561134e5763ffffffff6112f7611349938587613829565b35611301816110f1565b1680600052600c60205261131f600160406000206000815501613c8f565b7fa59268ca81d40429e65ccea5385b59cf2d3fc6519371dee92f8eb1dae5107a7a600080a2613b1d565b6112d4565b005b92919261135c82610cd8565b9161136a6040519384610c57565b82948184528183011161026c578281602093846000960137010152565b9080601f8301121561026c578160206105d993359101611350565b90916113b96105d99360408452604084019061046a565b91602081840391015261046a565b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57600435611402816110f1565b60243567ffffffffffffffff811161026c57611422903690600401611387565b906114996114946114466110ba8463ffffffff16600052600e602052604060002090565b93600661148d61145a836020815191012090565b9660016114778863ffffffff16600052600e602052604060002090565b019063ffffffff16600052602052604060002090565b0190613d5a565b613a0a565b9060609273ffffffffffffffffffffffffffffffffffffffff6114e660016114cb846000526003602052604060002090565b015473ffffffffffffffffffffffffffffffffffffffff1690565b166114fd575b505061068a604051928392836113a2565b6115669293509061152561073c61073c60016114cb6000966000526003602052604060002090565b60405180809581947f8318ed5d0000000000000000000000000000000000000000000000000000000083526004830191909163ffffffff6020820193169052565b03915afa9081156115a557600091611582575b509038806114ec565b61159f91503d806000833e6115978183610c57565b810190613d80565b38611579565b613ddf565b3461026c576115b836610271565b906115db61073c73ffffffffffffffffffffffffffffffffffffffff6000541690565b3314159160005b8181106115eb57005b6115f6818385613829565b359061160c82600052600d602052604060002090565b6001810154801561176f576005820180546117295750815463ffffffff604082901c16806116f157508790816116b6575b50610b295760019361168060027f5254e609a97bab37b7cc79fe128f85c097bd6015c6e1624ae0ba392eb9753205946116786116ad95615d48565b500154615e1b565b5061169d61169882600052600d602052604060002090565b613deb565b6040519081529081906020820190565b0390a1016115e2565b6116e891506116ce61078f61073c9263ffffffff1690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b3314153861163d565b7f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602485905260446000fd5b8461173a610ffc8661040894615ca1565b7f60a6d8980000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602452604490565b7fd82f6adb00000000000000000000000000000000000000000000000000000000600052600484905260246000fd5b3461026c576117ac36610271565b906117b5614b39565b60005b8281106117c157005b6117d46117cf828585613e3c565b613e7c565b906117f661073c835173ffffffffffffffffffffffffffffffffffffffff1690565b1561193f5760019161180d60125463ffffffff1690565b907f78e94ca80be2c30abc061b99e7eb8583b1254781734b1e3ce339abb57da2fe8e73ffffffffffffffffffffffffffffffffffffffff61191b611865845173ffffffffffffffffffffffffffffffffffffffff1690565b936118bb6020820195865161189761187b610c98565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201526118b68863ffffffff16600052600c602052604060002090565b613ec7565b6119016118d061085660125463ffffffff1690565b63ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000006012541617601255565b5173ffffffffffffffffffffffffffffffffffffffff1690565b92519261193663ffffffff6040519384931696169482610d12565b0390a3016117b8565b7feeacd9390000000000000000000000000000000000000000000000000000000060005260046000fd5b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5761068a6119a6600435614071565b604051918291602083526020830190610558565b9181601f8401121561026c5782359167ffffffffffffffff831161026c576020838186019501011161026c57565b908160c091031261026c5790565b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760043567ffffffffffffffff811161026c57611a459036906004016119ba565b9060243567ffffffffffffffff811161026c57611a669036906004016119e8565b91611a6f614b39565b611a7c6110718284614204565b9163ffffffff831615611bc0575050611aa58163ffffffff16600052600e602052604060002090565b611ab2606084018461427c565b611abf604086018661427c565b8454909691949060201c63ffffffff16611ad890613b1d565b81547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16602082901b67ffffffff000000001617825591611b1b60a082016142da565b915460401c60ff165b611b30608083016142ef565b611b3a83806142f9565b92909360208101611b4a916142f9565b959096611b55610ca9565b63ffffffff909c168c5263ffffffff1660208c0152151560408b0152151560608a015260ff1660808901523690611b8b92611350565b60a08701523690611b9b92611350565b60c08501523690611bab9261434a565b923690611bb792614396565b61134e926152a6565b6104436040519283927f4071db5400000000000000000000000000000000000000000000000000000000845260048401614236565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261026c576004359067ffffffffffffffff821161026c576102ba916004016119ba565b3461026c57611c5d611c58611c5236611bf5565b9061421d565b614533565b60405180916020820160208352815180915260206040840192019060005b818110611c89575050500390f35b8251845285945060209384019390920191600101611c7b565b3461026c576020611cd8611cbf611cb836611bf5565b3691611350565b8281519101206000526007602052604060002054151590565b6040519015158152f35b3461026c57611cf036611bf5565b611cf8613b7a565b5063ffffffff6040518284823760208184810160028152030190205416918215611bc05761068a611d488480600052600e602052611d42604060002063ffffffff905460201c1690565b90615057565b604051918291826111f1565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611d8757505050505090565b9091929394602080611dc3837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516110ff565b97019301930191939290611d78565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760125463ffffffff16611e1e611e19610ffc83613b43565b614453565b60009163ffffffff811660015b8163ffffffff821610611e655761068a8486611e49610ffc87613b43565b8103611e5d575b5060405191829182611d54565b815282611e50565b611e8561073c6116ce8363ffffffff16600052600c602052604060002090565b611e98575b60010163ffffffff16611e2b565b936001611ec163ffffffff926110cd6107a58963ffffffff16600052600c602052604060002090565b95915050611e8a565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5773ffffffffffffffffffffffffffffffffffffffff600154163303611fc75760005473ffffffffffffffffffffffffffffffffffffffff16600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055611f887fffffffffffffffffffffffff000000000000000000000000000000000000000060015416600155565b73ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152fd5b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760043567ffffffffffffffff811161026c5761207490369060040161023b565b60243567ffffffffffffffff811161026c5761209490369060040161023b565b928383036123555760005473ffffffffffffffffffffffffffffffffffffffff169060005b8481106120c257005b6120d56120d0828785613829565b613b39565b906120f08263ffffffff16600052600c602052604060002090565b9161210f835473ffffffffffffffffffffffffffffffffffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff841680156123215761213b6117cf858c8b613e3c565b9061215d61073c835173ffffffffffffffffffffffffffffffffffffffff1690565b1561193f5733141580612301575b610b295760019473ffffffffffffffffffffffffffffffffffffffff6121a861073c845173ffffffffffffffffffffffffffffffffffffffff1690565b911614801590612295575b6121c1575b505050016120b9565b73ffffffffffffffffffffffffffffffffffffffff61226f8261225d61221c7f86f41145bde5dd7f523305452e4aad3685508c181432ec733d5f345009358a28955173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61190160208201958987519101613462565b92519261228a63ffffffff6040519384931696169482610d12565b0390a33880806121b8565b5060405160208101906122da816122ae898701856144c0565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610c57565b51902060208201516040516122f7816122ae602082019485610d12565b51902014156121b3565b5073ffffffffffffffffffffffffffffffffffffffff871633141561216b565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b7fab8b67c6000000000000000000000000000000000000000000000000000000006000526004839052602484905260446000fd5b3461026c5761239736610271565b906123a0614b39565b60005b8281106123ac57005b6123ba611cb88285856144d1565b805160208201206123db61030f826000526005602052604060002054151590565b61245c5761030f6123eb91615b9b565b61242657906123fb600192613730565b7fb2553249d353abf34f62139c85f44b5bdeab968ec0ab296a9bf735b75200ed83600080a2016123a3565b610443906040519182917f88c8a73700000000000000000000000000000000000000000000000000000000835260048301610d12565b6040517f4ff9e48a000000000000000000000000000000000000000000000000000000008152806104438460048301610d12565b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c576004356124cb816110f1565b60243567ffffffffffffffff811161026c576124eb9036906004016119e8565b906124f4614b39565b61250e8163ffffffff16600052600e602052604060002090565b549163ffffffff602084901c169081156125635761252f606082018261427c565b909161253e604082018261427c565b96909461254a90613b1d565b9161255760a082016142da565b9160401c60ff16611b24565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57602073ffffffffffffffffffffffffffffffffffffffff60005416604051908152f35b3461026c576125f736610271565b61261961073c60005473ffffffffffffffffffffffffffffffffffffffff1690565b33149060009115925b81831061262b57005b61263961077584848461384d565b9260408401936126548551600052600d602052604060002090565b916126696107a561078f855463ffffffff1690565b90600184018054928315612b5357899081612b29575b50610b29576020840192835115610a86578351808203612aed575b5050506060830180518015610a595750608084015198895115612ab957956126cd610856875463ffffffff9060201c1690565b86547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16602082901b67ffffffff000000001617875597600487019760005b8c518110156127b15761273c61030f6127298f846108bf91613839565b6000526005602052604060002054151590565b61277d57806127768e6127706108bf8f958f60019761276a919063ffffffff16600052602052604060002090565b93613839565b90615bd0565b500161270c565b6040517fdb68f0c2000000000000000000000000000000000000000000000000000000008152806104438f60048301613a70565b50929a5098979695939492906127cf855463ffffffff9060401c1690565b63ffffffff81166129f3575b50979198996127ec60058601614533565b9260009a5b845163ffffffff8d16908110156129495761280c9086613839565b5163ffffffff169c8d61282f8163ffffffff16600052600e602052604060002090565b6001019061284d9063ffffffff16600052600e602052604060002090565b5460201c63ffffffff16612870919063ffffffff16600052602052604060002090565b60030161287c90614533565b9b60005b8d5181101561292e576128cd61030f8f8f8f6128b06128b69287929063ffffffff16600052602052604060002090565b92613839565b519060019160005201602052604060002054151590565b6128d957600101612880565b8f6128f98f836128e891613839565b516000526011602052604060002090565b6104436040519283927f16c2b7c40000000000000000000000000000000000000000000000000000000084526004840161457e565b50939d509b61293f91929b50613b1d565b9a9b9199906127f1565b509b9a50925094600197506129e89196507f4b5b465e22eea0c3d40c30e936643245b80d19b2dcf75788c0699fe8d8db645b94600361299163ffffffff965163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8216178355945197886002840155519101555160405193849316958360209093929193604081019481520152565b0390a2019190612622565b9a91929899612a596003612a538e612a3f6110ba6001612a27849e9c9d9e63ffffffff16600052600e602052604060002090565b019263ffffffff16600052600e602052604060002090565b63ffffffff16600052602052604060002090565b01614533565b9960005b8b51811015612aa757612a8d61030f8d6128b6848f8f906128b0919063ffffffff16600052602052604060002090565b612a9957600101612a5d565b6128f96128e88f928e613839565b5093929b5099985094939294386127db565b6040517fdb68f0c2000000000000000000000000000000000000000000000000000000008152806104438c60048301613a70565b612b04906000526009602052604060002054151590565b610a8657612b159184519055615d48565b50612b208251615b31565b5038808061269a565b51612b4a915073ffffffffffffffffffffffffffffffffffffffff1661073c565b3314153861267f565b88517fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57600435612bbd816110f1565b6024359067ffffffffffffffff821161026c57612be161134e9236906004016119ba565b91612bea614b39565b6146d3565b3461026c5761134e612c0036610271565b90612c09614b39565b614968565b3461026c57612c1c36610271565b612c24614b39565b60005b818110612c3057005b8063ffffffff6020612c4560019486886144d1565b91908260405193849283378101600281520301902054168015612c7157612c6b90614dce565b01612c27565b50612c6b565b3461026c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57600435612cb2816110f1565b602435612cbe816110f1565b612cc6613b7a565b5063ffffffff821680600052600e60205263ffffffff60406000205460201c161561128e5750612d0c610ffc6110ba8463ffffffff16600052600e602052604060002090565b63ffffffff821611612d2557611d489061068a92615057565b81612d466110ba6104089463ffffffff16600052600e602052604060002090565b7ff3c16e2c0000000000000000000000000000000000000000000000000000000060005263ffffffff918216600452811660245216604452606490565b3461026c57602063ffffffff81612d9936611bf5565b91908260405193849283378101600281520301902054161515604051908152f35b6105d9916060612dd3835160808452608084019061046a565b9273ffffffffffffffffffffffffffffffffffffffff6020820151166020840152604081015115156040840152015190606081840391015261046a565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5760043567ffffffffffffffff811161026c57612e6a612e6561068a923690600401611387565b614aaa565b604051918291602083526020830190612dba565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310612eb157505050505090565b9091929394602080612eed837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951612dba565b97019301930191939290612ea2565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57612f336144e8565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612f79612f6384613750565b93612f716040519586610c57565b808552613750565b0160005b818110612fd257505060005b8151811015612fc45780612fa8612e656114946128e860019587613839565b612fb28286613839565b52612fbd8185613839565b5001612f89565b6040518061068a8582612e7e565b602090612fdd614a84565b82828701015201612f7d565b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57604051600a548082528160208101600a60005260206000209260005b8181106130a057505061304992500382610c57565b61305381516137bc565b9060005b8151811015613092578061307661307060019385613839565b51614071565b6130808286613839565b5261308b8185613839565b5001613057565b6040518061068a85826105dc565b8454835260019485019486945060209093019201613034565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361026c57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361026c57565b3461026c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c5773ffffffffffffffffffffffffffffffffffffffff6131496130b9565b613151614b39565b163381146131df57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff6131b760005473ffffffffffffffffffffffffffffffffffffffff1690565b167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152fd5b3461026c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026c57602063ffffffff601254821c16604051908152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156132f05760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18136030182121561026c570190565b613281565b60608136031261026c57604051906060820182811067ffffffffffffffff821117610c1a57604052803567ffffffffffffffff811161026c5761333b9036908301611387565b8252613349602082016130dc565b602083015260408101359067ffffffffffffffff821161026c5761336f91369101611387565b604082015290565b90600182811c921680156133c0575b602083101461339157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613386565b91613402918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b818110613411575050565b60008155600101613406565b9190601f811161342c57505050565b610ca7926000526020600020906020601f840160051c83019310613458575b601f0160051c0190613406565b909150819061344b565b919091825167ffffffffffffffff8111610c1a5761348a816134848454613377565b8461341d565b6020601f82116001146134e45781906134029394956000926134d9575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b0151905038806134a7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061351784600052602060002090565b9160005b8181106135715750958360019596971061353a575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613530565b9192602060018192868b01518155019401920161351b565b919091825192835167ffffffffffffffff8111610c1a576135b4816135ae8554613377565b8561341d565b6020601f821160011461367e579161360a82604093600295610ca798996000926134d95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b84555b613675613631602083015173ffffffffffffffffffffffffffffffffffffffff1690565b600186019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b01519101613462565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216956136b185600052602060002090565b9660005b818110613718575092610ca796976002959360019383604097106136e1575b505050811b01845561360d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806136d4565b838301518955600190980197602093840193016136b5565b61374890602060405192828480945193849201610447565b810103902090565b67ffffffffffffffff8111610c1a5760051b60200190565b60405190610100820182811067ffffffffffffffff821117610c1a57604052606060e08360008152600060208201526000604082015260008382015260006080820152600060a08201528260c08201520152565b906137c682613750565b6137d36040519182610c57565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06138018294613750565b019060005b82811061381257505050565b60209061381d613768565b82828501015201613806565b91908110156132f05760051b0190565b80518210156132f05760209160051b010190565b91908110156132f05760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561026c570190565b60a08136031261026c57604051906138a482610c1f565b80356138af816110f1565b825260208101356020830152604081013560408301526060810135606083015260808101359067ffffffffffffffff821161026c570136601f8201121561026c5780356138fb81613750565b916139096040519384610c57565b81835260208084019260051b8201019036821161026c5760208101925b8284106139395750505050608082015290565b833567ffffffffffffffff811161026c5760209161395d8392833691870101611387565b815201930192613926565b6000929181549161397883613377565b80835292600181169081156139ce575060011461399457505050565b60009081526020812093945091925b8383106139b4575060209250010190565b6001816020929493945483858701015201910191906139a3565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b90610ca7613a1e9260405193848092613968565b0383610c57565b9060016020604051613a3681610bfe565b613a6c819573ffffffffffffffffffffffffffffffffffffffff8154168352613a656040518096819301613968565b0384610c57565b0152565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310613aa357505050505090565b9091929394602080613adf837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161046a565b97019301930191939290613a94565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff1663ffffffff8114613b345760010190565b613aee565b356105d9816110f1565b63ffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9116019063ffffffff8211613b3457565b60405190610140820182811067ffffffffffffffff821117610c1a57604052606061012083600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201528260e0820152826101008201520152565b90613be582613750565b613bf26040519182610c57565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613c208294613750565b019060005b828110613c3157505050565b602090613c3c613b7a565b82828501015201613c25565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114613b345760010190565b60405190613c8282610bfe565b6060602083600081520152565b613c998154613377565b9081613ca3575050565b81601f60009311600114613cb5575055565b81835260208320613cd191601f0160051c810190600101613406565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c191617905555565b6020613d27918160405193828580945193849201610447565b8101601081520301902090565b6020613d4d918160405193828580945193849201610447565b8101600281520301902090565b602090613d74928260405194838680955193849201610447565b82019081520301902090565b60208183031261026c5780519067ffffffffffffffff821161026c570181601f8201121561026c578051613db381610cd8565b92613dc16040519485610c57565b8184526020828401011161026c576105d99160208085019101610447565b6040513d6000823e3d90fd5b60059060008155600060018201556000600282015560006003820155018054906000815581613e18575050565b6000526020600020908101905b818110613e30575050565b60008155600101613e25565b91908110156132f05760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18136030182121561026c570190565b60408136031261026c5760405190613e9382610bfe565b613e9c816130dc565b825260208101359067ffffffffffffffff821161026c57613ebf91369101611387565b602082015290565b6001602091939293613f2b73ffffffffffffffffffffffffffffffffffffffff865116829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b0192015191825167ffffffffffffffff8111610c1a57613f4f816134848454613377565b6020601f8211600114613f9d5781906134029394956000926134d95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690613fd084600052602060002090565b9160005b818110613ff25750958360019596971061353a57505050811b019055565b9192602060018192868b015181550194019201613fd4565b9061401482613750565b6140216040519182610c57565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061404f8294613750565b019060005b82811061406057505050565b806060602080938501015201614054565b9061407a613768565b506140af611c58600461409785600052600d602052604060002090565b01612a3f6110ba86600052600d602052604060002090565b6140b9815161400a565b9160005b82518110156140f557806140d96114946128e860019487613839565b6140e38287613839565b526140ee8186613839565b50016140bd565b5092905061411061107182600052600d602052604060002090565b91600261412783600052600d602052604060002090565b0154600161413f84600052600d602052604060002090565b015490600361415885600052600d602052604060002090565b0154916141e86141756110ba87600052600d602052604060002090565b956141db6141b66005612a536141a561419886600052600d602052604060002090565b5460401c63ffffffff1690565b94600052600d602052604060002090565b976141ce6141c2610cb8565b63ffffffff909b168b52565b63ffffffff1660208a0152565b63ffffffff166040880152565b6060860152608085015260a084015260c083015260e082015290565b6020908260405193849283378101600281520301902090565b6020908260405193849283378101601081520301902090565b90601f836040947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe093602086528160208701528686013760008582860101520116010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561026c570180359067ffffffffffffffff821161026c57602001918160051b3603831361026c57565b8015150361026c57565b356105d9816142d0565b60ff81160361026c57565b356105d9816142e4565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561026c570180359067ffffffffffffffff821161026c5760200191813603831361026c57565b92919061435681613750565b936143646040519586610c57565b602085838152019160051b810192831161026c57905b82821061438657505050565b813581526020918201910161437a565b929190926143a384613750565b936143b16040519586610c57565b602085828152019060051b82019183831161026c5780915b8383106143d7575050505050565b823567ffffffffffffffff811161026c57820160408187031261026c576040519161440183610bfe565b813567ffffffffffffffff811161026c578761441e918401611387565b835260208201359267ffffffffffffffff841161026c5761444488602095869501611387565b838201528152019201916143c9565b9061445d82613750565b61446a6040519182610c57565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06144988294613750565b019060005b8281106144a957505050565b6020906144b4613c75565b8282850101520161449d565b9060206105d9928181520190613968565b908210156132f0576102ba9160051b8101906142f9565b604051906004548083528260208101600460005260206000209260005b81811061451a575050610ca792500383610c57565b8454835260019485019487945060209093019201614505565b906040519182815491828252602082019060005260206000209260005b818110614565575050610ca792500383610c57565b8454835260019485019487945060209093019201614550565b9063ffffffff61459b602092959495604085526040850190613968565b9416910152565b90929167ffffffffffffffff8111610c1a576145c2816134848454613377565b6000601f821160011461461b5781906134029394956000926146105750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b0135905038806134a7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169461464e84600052602060002090565b91805b8781106146a657508360019596971061466e57505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080613530565b90926020600181928686013581550194019101614651565b81604051928392833781016000815203902090565b9163ffffffff6146f66110ba8563ffffffff16600052600e602052604060002090565b1615612563576147196114948463ffffffff16600052600f602052604060002090565b908151602083012061472c368386611350565b60208151910120146147fb5763ffffffff928261477693516147df575b508161479f5761477161476c8663ffffffff16600052600f602052604060002090565b613c8f565b6146be565b91167f3143778fe5911daca9ced3c73dde73fdb95e7edae557614c40588f3f1a1b44b8600080a3565b6147c382826147be8863ffffffff16600052600f602052604060002090565b6145a2565b6147d96147d0838361421d565b85871690615bd0565b506146be565b6147eb6147f491613d0e565b85871690615ee2565b5038614749565b50505050565b91908110156132f05760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218136030182121561026c570190565b9080601f8301121561026c578160206105d993359101614396565b9080601f8301121561026c578160206105d99335910161434a565b3590610ca7826142e4565b3590610ca7826142d0565b60e08136031261026c5761489f610ca9565b90803567ffffffffffffffff811161026c576148be9036908301611387565b8252602081013567ffffffffffffffff811161026c576148e19036908301611387565b6020830152604081013567ffffffffffffffff811161026c576149079036908301614841565b6040830152606081013567ffffffffffffffff811161026c576149609161493360c0923690830161485c565b606085015261494460808201614877565b608085015261495560a08201614882565b60a085015201614882565b60c082015290565b908015614a805760005b81811061497e57505050565b80614a7a6149976149926001948688614801565b61488d565b60608101516040820151906149b560125463ffffffff9060201c1690565b926149fc6149c285613b1d565b7fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff000000006012549260201b16911617601255565b60a0810151151590614a6b614a1460c0830151151590565b614a61614a25608085015160ff1690565b91614a586020865196015196614a48614a3c610ca9565b63ffffffff909c168c52565b600160208c0152151560408b0152565b15156060890152565b60ff166080870152565b60a085015260c08401526152a6565b01614972565b5050565b60405190614a9182610c3b565b6060808381815260006020820152600060408201520152565b614ab2614a84565b50614b318151602083012080600052600360205260026040600020019080600052600360205273ffffffffffffffffffffffffffffffffffffffff614b0f8160016040600020015416926000526007602052604060002054151590565b9160405195614b1d87610c3b565b865216602085015215156040840152613a0a565b606082015290565b73ffffffffffffffffffffffffffffffffffffffff600054163303614b5a57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152fd5b60206000604051828101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a700000000000000000000000000000000000000000000000000000000602482015260248152614c1b604482610c57565b519084617530fa903d6000519083614d16575b5082614d0c575b5081614c71575b81614c45575090565b6105d991507f78bea7210000000000000000000000000000000000000000000000000000000090615c0e565b905060206000604051828101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff00000000000000000000000000000000000000000000000000000000602482015260248152614cd6604482610c57565b519084617530fa6000513d82614d00575b5081614cf6575b501590614c3c565b9050151538614cee565b60201115915038614ce7565b1515915038614c35565b60201115925038614c2e565b60405190816000825492614d3584613377565b9360018116908115614d965750600114614d5a575b5060209250600281520301902090565b9150506000528160206000206000905b838210614d7e575050602091810138614d4a565b60209192508060019154848701520191018391614d6a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168452506020938015150283019150389050614d4a565b614de88163ffffffff16600052600e602052604060002090565b90815492614dfd8463ffffffff9060201c1690565b906001840191614e2a614e2082859063ffffffff16600052602052604060002090565b9660401c60ff1690565b9260005b8754811015614eb4576001908515614e8e57614e88614e60614e50838c615ca1565b600052600d602052604060002090565b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff8154169055565b01614e2e565b614eae6005614ea0614e50848d615ca1565b0163ffffffff891690615ee2565b50614e88565b5094549195509293915060201c63ffffffff1615612563576005614eee614f1c93614ef4939063ffffffff16600052602052604060002090565b01614d22565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008154169055565b614f396114948263ffffffff16600052600f602052604060002090565b8051614f9f575b506000614f5d8263ffffffff16600052600e602052604060002090565b557ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c158170365163ffffffff60405192169180614f9a81906000602083019252565b0390a2565b614fab614fb891613d0e565b63ffffffff831690615ee2565b50614fd661476c8263ffffffff16600052600f602052604060002090565b38614f40565b90614fe682613750565b614ff36040519182610c57565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06150218294613750565b019060005b82811061503257505050565b60209060405161504181610bfe565b6060815260608382015282828501015201615026565b91615060613b7a565b5061507b8363ffffffff16600052600e602052604060002090565b9161509981600185019063ffffffff16600052602052604060002090565b916150a660038401614533565b916150b18351614fdc565b9360009660068201975b865181101561511557806150d76114946128e86001948a613839565b6150f46150e48c83613d5a565b6150ec610c98565b928352613a0a565b6020820152615103828a613839565b5261510e8189613839565b50016150bb565b50929650925093549361512b8563ffffffff1690565b9460401c60ff169161514d9063ffffffff16600052600f602052604060002090565b9160048401546151618160ff9060081c1690565b9060ff169061516f86614533565b93615178610cc8565b63ffffffff909916895263ffffffff16602089015260ff166040880152151560608701521515608086015260a08501526151b190613a0a565b60c08401526151c260058201613a0a565b60e08401526002016151d390613a0a565b61010083015261012082015290565b60ff60019116019060ff8211613b3457565b6105d99054613377565b60409063ffffffff6105d99493168152816020820152019061046a565b6040906105d993928152816020820152019061046a565b80548210156132f05760005260206000200190600090565b80549068010000000000000000821015610c1a578161527191600161340294018155615232565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b91909260016152d26152bc845163ffffffff1690565b63ffffffff16600052600e602052604060002090565b019360208301906153006152ea835163ffffffff1690565b879063ffffffff16600052602052604060002090565b926080850196615311885160ff1690565b60ff811615908115615adf575b50615a9a5761533d90612a3f615338865163ffffffff1690565b613b43565b9261534a60058501613a0a565b602081519101209360a087019485516020815191012003615a0e575b6001615379610ffc845163ffffffff1690565b116159a7575b5060005b8751811015615541576153a361030f61539c838b613839565b5188615bd0565b6154f3576060870151156154c6576153d26141986153c1838b613839565b51600052600d602052604060002090565b63ffffffff6153e8610ffc8a5163ffffffff1690565b91161415806154a8575b61545a5760019061545461540a895163ffffffff1690565b6154176153c1848d613839565b907fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff6bffffffff000000000000000083549260401b169116179055565b01615383565b87615470610408926128b08a5163ffffffff1690565b517f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff909116600452602452604490565b5063ffffffff6154be6141986153c1848c613839565b1615156153f2565b806154ed60056154db6153c16001958d613839565b01612770610ffc8b5163ffffffff1690565b50615454565b87615509610408926128b08a5163ffffffff1690565b517f636e40570000000000000000000000000000000000000000000000000000000060005263ffffffff909116600452602452604490565b509196909395946000936006880192600389019860028101926004600583019201935b8c5189101561593b578c61557b8a6155899f613839565b519d8e516020815191012090565b9a6155a461030f8d6000526005602052604060002054151590565b615904576155bf8c6000526007602052604060002054151590565b6158cd576155d78f8a6155d29151613d5a565b6151f4565b6158865760005b8d5181101561567a5761562961030f8f8f9080612a3f6110ba6153c188600461560d6153c18361561499613839565b0194613839565b60019160005201602052604060002054151590565b615635576001016155de565b8f90615641908f613839565b519051906104436040519283927f4b5786e70000000000000000000000000000000000000000000000000000000084526004840161521b565b509990929d898f8f8f9193869f6156d960408f948f908d61587b9a6156cb60019e6156ab61570a996156d29661524a565b6156c060208b019b6103308d51918d51613d5a565b60c086015190613462565b5190613462565b0151151590565b8c9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b61574b6157188b5160ff1690565b8c547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1660089190911b61ff0016178c55565b6157a661575c8d5163ffffffff1690565b6157768163ffffffff16600052600e602052604060002090565b9063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6158048c6157c86152bc6157bd6060840151151590565b925163ffffffff1690565b907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff68ff0000000000000000835492151560401b169116179055565b6158558c61581c6152bc6157bd885163ffffffff1690565b907fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff0000000083549260201b169116179055565b6158716158668d5163ffffffff1690565b945163ffffffff1690565b9051915193616028565b0197989b909b615564565b8e615895895163ffffffff1690565b9051906104436040519283927f368812ac000000000000000000000000000000000000000000000000000000008452600484016151fe565b6104438f516040519182917f88c8a73700000000000000000000000000000000000000000000000000000000835260048301610d12565b6104438f516040519182917f4ff9e48a00000000000000000000000000000000000000000000000000000000835260048301610d12565b5050505050905063ffffffff9396506159829195507ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c1581703651945061598d92505163ffffffff1690565b935163ffffffff1690565b60405163ffffffff90911681529216918060208101614f9a565b9794929096939160005b89548110156159ff57806159f9614e60614e50838e6159f38e6159ed610ffc60056159e1614e5060019d88615ca1565b01925163ffffffff1690565b90615ee2565b50615ca1565b016159b1565b5091939690929497503861537f565b615a1d614ef460058801614d22565b84518051615a2c575b50615366565b615a3d61107163ffffffff92613d34565b16615a6357615a5d615a53885163ffffffff1690565b6157768751613d34565b38615a26565b61044385516040519182917f07bf02d600000000000000000000000000000000000000000000000000000000835260048301610d12565b61040887615aa98a5160ff1690565b90517f25b4d6180000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b615ae991506151e2565b60ff88519116113861531e565b600081815260056020526040902054615b2b57615b1481600461524a565b600454906000526005602052604060002055600190565b50600090565b600081815260096020526040902054615b2b57615b4f81600861524a565b600854906000526009602052604060002055600190565b6000818152600b6020526040902054615b2b57615b8481600a61524a565b600a5490600052600b602052604060002055600190565b600081815260076020526040902054615b2b57615bb981600661524a565b600654906000526007602052604060002055600190565b6000828152600182016020526040902054615c075780615bf28360019361524a565b80549260005201602052604060002055600190565b5050600090565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a700000000000000000000000000000000000000000000000000000000845216602482015260248152615c74604482610c57565b5191617530fa6000513d82615c95575b5081615c8e575090565b9050151590565b60201115915038615c84565b90615cab91615232565b90549060031b1c90565b80548015615d19577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615cea8282615232565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260096020526040902054908115615c07577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211613b3457600854927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411613b34578383615de49460009603615dea575b505050615dd36008615cb5565b600990600052602052604060002090565b55600190565b615dd3615e0c91615e02615cab615e12956008615232565b9283916008615232565b906133ca565b55388080615dc6565b6000818152600b6020526040902054908115615c07577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211613b3457600a54927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411613b34578383615de49460009603615eb7575b505050615ea6600a615cb5565b600b90600052602052604060002090565b615ea6615e0c91615ecf615cab615ed995600a615232565b928391600a615232565b55388080615e99565b6001810191806000528260205260406000205492831515600014615fba577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111613b34578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501948511613b34576000958583615f7294615de49803615f81575b505050615cb5565b90600052602052604060002090565b615fa1615e0c91615f98615cab615fb19588615232565b92839187615232565b8590600052602052604060002090565b55388080615f6a565b50505050600090565b9294939160808401608085528251809152602060a0860193019060005b818110616012575050509163ffffffff6160058360609587849603602089015261046a565b9616604085015216910152565b8251855260209485019490920191600101615fe0565b9390916020815191012073ffffffffffffffffffffffffffffffffffffffff80600161605e846000526003602052604060002090565b0154161661606e575b5050505050565b61073c61073c60016114cb61608d946000526003602052604060002090565b90813b1561026c57600080946160d2604051978896879586947ffba64a7c00000000000000000000000000000000000000000000000000000000865260048601615fc3565b03925af180156115a5576160ea575b80808080616067565b806160f960006160ff93610c57565b80610bc4565b386160e156fea164736f6c634300081a000a",
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

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateDON(opts *bind.TransactOpts, donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateDON", donId, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateDON(donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateDON(donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateDONByName(opts *bind.TransactOpts, donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateDONByName", donName, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateDONByName(donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDONByName(&_CapabilitiesRegistry.TransactOpts, donName, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateDONByName(donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDONByName(&_CapabilitiesRegistry.TransactOpts, donName, updateDONParams)
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

	UpdateDON(opts *bind.TransactOpts, donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error)

	UpdateDONByName(opts *bind.TransactOpts, donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error)

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
