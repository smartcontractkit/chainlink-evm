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

type CapabilitiesRegistryConstructorParams struct {
	CanAddOneNodeDONs bool
}

type CapabilitiesRegistryDONInfo struct {
	Id                       uint32
	ConfigCount              uint32
	F                        uint8
	IsPublic                 bool
	AcceptsWorkflows         bool
	NodeP2PIds               [][32]byte
	DonFamilies              []string
	Name                     string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
}

type CapabilitiesRegistryNewDONParams struct {
	Name                     string
	DonFamilies              []string
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.ConstructorParams\",\"components\":[{\"name\":\"canAddOneNodeDONs\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCapabilities\",\"inputs\":[{\"name\":\"capabilities\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.Capability[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addDONs\",\"inputs\":[{\"name\":\"newDONs\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NewDONParams[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodeOperators\",\"inputs\":[{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deprecateCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCapabilities\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapability\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityConfigs\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONFamilies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.DONInfo[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONsInFamily\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHistoricalDONInfo\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextDONId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nodeInfo\",\"type\":\"tuple\",\"internalType\":\"structINodeInfoProvider.NodeInfo\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperator\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodesByP2PIds\",\"inputs\":[{\"name\":\"p2pIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDONNameTaken\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDONs\",\"inputs\":[{\"name\":\"donIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDONsByName\",\"inputs\":[{\"name\":\"donNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodes\",\"inputs\":[{\"name\":\"removedNodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONFamilies\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addToFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"removeFromFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updateDONParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.UpdateDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"updateDONParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.UpdateDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CapabilityConfigured\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONAddedToFamily\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONRemovedFromFamily\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeAdded\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorAdded\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorRemoved\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorUpdated\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeRemoved\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CapabilityAlreadyExists\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityDoesNotExist\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityIsDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityRequiredByDON\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONConfigDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"requestedConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONNameAlreadyTaken\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DONWithNameDoesNotExist\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONCapability\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONNode\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidCapabilityConfigurationContractInterface\",\"inputs\":[{\"name\":\"proposedConfigurationContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidFaultTolerance\",\"inputs\":[{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nodeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeEncryptionPublicKey\",\"inputs\":[{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeOperatorAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNodeP2PId\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"lengthOne\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lengthTwo\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NodeAlreadyExists\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotExist\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotSupportCapability\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfCapabilitiesDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfWorkflowDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60a0604052346100e557604051601f6163df38819003918201601f19168301916001600160401b038311848410176100ea578084926020946040528339810103126100e55760405190600090602083016001600160401b038111848210176100d1576040525180151581036100cd57825233156100be5750600180546001600160a01b03191633179055601580546001600160401b0319166401000000011790555115156080526040516162de90816101018239608051816152360152f35b639b15e16f60e01b8152600490fd5b5080fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c80628375c61461024657806305a5196614610241578063106e9a081461023c578063181f5a77146102375780631d05394c14610232578063214502431461022d57806322bdbcbc146102285780632353740514610223578063275459f21461021e5780632af97674146102195780632c01a1e814610214578063398f37731461020f57806350c946fe1461020a57806353a25dd714610205578063543f40251461020057806359003602146101fb57806359110666146101f657806366acaa33146101f157806379ba5097146101ec57806386fa4246146101e757806388ea09ee146101e257806388eafafb146101dd5780638da5cb5b146101d857806394bbb012146101d357806396ef4fc9146101ce578063a04ab55e146101c9578063a7eb612c146101c4578063b8521761146101bf578063bfa8eef5146101ba578063c9315179146101b5578063cd71fd09146101b0578063ddbe4f82146101ab578063e29581aa146101a6578063f2fde38b146101a15763fcdc8efe1461019c57600080fd5b6132ea565b6131f7565b613106565b613019565b612f2d565b612ea0565b612dad565b612d44565b6127ab565b61271b565b612676565b6125d9565b612587565b612480565b612379565b612015565b611f2c565b611e34565b611d44565b611d0b565b611c92565b611a4a565b6119bd565b6117f2565b6115ff565b61141c565b6112af565b6111f6565b611137565b610fa6565b610d94565b610d17565b61071c565b61066a565b6102ce565b9181601f8401121561027c5782359167ffffffffffffffff831161027c576020808501948460051b01011161027c57565b600080fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261027c576004359067ffffffffffffffff821161027c576102ca9160040161024b565b9091565b3461027c576102dc36610281565b906102e5614a31565b60005b8281106102f157005b6103046102ff82858561335d565b6133a2565b61031381516020815191012090565b61032361031f82615bb4565b1590565b61041c576103458251610340836000526014602052604060002090565b61350f565b60208201805173ffffffffffffffffffffffffffffffffffffffff16806103be575b50508161038d60019493610388610393946000526003602052604060002090565b613636565b516137dd565b7fe671cf109707667795a875c19f031bdbc7ed40a130f6dc18a55615a0e0099fbb600080a2016102e8565b61031f6103ca91614a7c565b6103d45780610367565b517fabb5e3fd0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6000fd5b61045382516040519182917f8f51ece800000000000000000000000000000000000000000000000000000000835260048301610d06565b0390fd5b60005b83811061046a5750506000910152565b818101518382015260200161045a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936104b681518092818752878088019101610457565b0116010190565b9080602083519182815201916020808360051b8301019401926000915b8383106104e957505050505090565b9091929394602080610525837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161047a565b970193019301919392906104da565b906020808351928381520192019060005b8181106105525750505090565b8251845260209384019390920191600101610545565b6105e99163ffffffff825116815261058d6020830151602083019063ffffffff169052565b60408281015163ffffffff1690820152606082015160608201526080820151608082015260a082015160a082015260e06105d860c084015161010060c08501526101008401906104bd565b9201519060e0818403910152610534565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061061f57505050505090565b909192939460208061065b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610568565b97019301930191939290610610565b3461027c5761067836610281565b61068181613869565b9060005b81811061069e576040518061069a85826105ec565b0390f35b6106b26106ac8284876138d6565b3561401b565b6106bc82856138e6565b526106c781846138e6565b5060806106d482856138e6565b510151156106e457600101610685565b6106ee91846138d6565b357fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461027c5761072a36610281565b61076561074c73ffffffffffffffffffffffffffffffffffffffff6001541690565b73ffffffffffffffffffffffffffffffffffffffff1690565b33149060009115925b81831061077757005b61078a6107858484846138fa565b6139b9565b926107ba6107b561079f865163ffffffff1690565b63ffffffff16600052600e602052604060002090565b613af2565b6107db61074c825173ffffffffffffffffffffffffffffffffffffffff1690565b15610b9157859081610b67575b50610b395760408401936108078551600052600f602052604060002090565b9460018601908154610b0b5780518015610ade5750602083019182518015908115610ac0575b50610a9657606084019788518015610a6957506080850151998a5115610a3557966108a461086b610866845463ffffffff9060201c1690565b613b6c565b83547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff1660209190911b67ffffffff0000000016178355565b815460201c63ffffffff1698600483019860005b8d5181101561095b576108da6108cf8f83906138e6565b516020815191012090565b6108f461031f826000526005602052604060002054151590565b610925579061091e6001926109198f8f9063ffffffff16600052602052604060002090565b615cc3565b50016108b8565b6104538f6040519182917fdb68f0c2000000000000000000000000000000000000000000000000000000008352600483016125f8565b509a63ffffffff95919c50610a0c929a50600199507f74becb12a5e8fd0e98077d02dfba8f647c9670c9df177e42c2418cf17a636f059694610a29946109f4925160038201556109e36109b28b5163ffffffff1690565b829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b600284519101558551809155615bef565b506109ff8151615c24565b5051955163ffffffff1690565b915160405193849316958360209093929193604081019481520152565b0390a20191929061076e565b6040517fdb68f0c2000000000000000000000000000000000000000000000000000000008152806104538d600483016125f8565b7f37d897650000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f837731460000000000000000000000000000000000000000000000000000000060005260046000fd5b610ad891506000526009602052604060002054151590565b3861082d565b7f64e2ee920000000000000000000000000000000000000000000000000000000060005260045260246000fd5b517f546184830000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f9473075d000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b51610b88915073ffffffffffffffffffffffffffffffffffffffff1661074c565b331415386107e8565b610418610ba2865163ffffffff1690565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602490565b600091031261027c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117610c2a57604052565b610bdf565b6080810190811067ffffffffffffffff821117610c2a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610c2a57604052565b60405190610c9b604083610c4b565b565b60405190610c9b60e083610c4b565b60405190610c9b61010083610c4b565b60405190610c9b61014083610c4b565b67ffffffffffffffff8111610c2a57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9060206105e992818152019061047a565b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5761069a6040805190610d588183610c4b565b601a82527f4361706162696c6974696573526567697374727920322e302e3000000000000060208301525191829160208352602083019061047a565b3461027c57610da236610281565b90610dab614a31565b60005b828110610db757005b80610dd7610dc860019386866138d6565b35610dd2816110e5565b614c8f565b01610dae565b9080602083519182815201916020808360051b8301019401926000915b838310610e0957505050505090565b9091929394602080610e65837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519083610e55835160408452604084019061047a565b920151908481840391015261047a565b97019301930191939290610dfa565b805163ffffffff1682526105e99160208281015163ffffffff169082015260408281015160ff1690820152606082810151151590820152608082810151151590820152610120610f16610f02610ef0610ede60a087015161014060a0880152610140870190610534565b60c087015186820360c08801526104bd565b60e086015185820360e087015261047a565b61010085015184820361010086015261047a565b92015190610120818403910152610ddd565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310610f5b57505050505090565b9091929394602080610f97837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610e74565b97019301930191939290610f4c565b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5760155460201c63ffffffff16610ffe610ff9610ff083613b92565b63ffffffff1690565b613c2a565b60009163ffffffff811660015b8163ffffffff8216106110455761069a8486611029610ff087613b92565b810361103d575b5060405191829182610f28565b815282611030565b61106f610ff06110658363ffffffff166000526010602052604060002090565b5463ffffffff1690565b611082575b60010163ffffffff1661100b565b9360016110dc63ffffffff926110c16110bb6110ae8a63ffffffff166000526010602052604060002090565b5460201c63ffffffff1690565b89614f12565b6110cb82896138e6565b526110d681886138e6565b50613c97565b95915050611074565b63ffffffff81160361027c57565b90604060206105e99373ffffffffffffffffffffffffffffffffffffffff8151168452015191816020820152019061047a565b9060206105e99281815201906110f3565b3461027c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5763ffffffff600435611177816110e5565b61117f613cc4565b5016600052600e60205261069a604060002060016111d4604051926111a384610c0e565b73ffffffffffffffffffffffffffffffffffffffff81541684526111cd6040518094819301613a35565b0382610c4b565b602082015260405191829182611126565b9060206105e9928181520190610e74565b3461027c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c57600435611231816110e5565b611239613bc9565b5063ffffffff81169081600052601060205263ffffffff60406000205460201c169182156112825761069a61126e8484614f12565b604051918291602083526020830190610e74565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461027c576112bd36610281565b6112c5614a31565b60005b63ffffffff811690828210156113a35763ffffffff6112eb6113489385876138d6565b356112f5816110e5565b1680600052600e6020526001604060002060008155016113158154613424565b908161134d575b50507fa59268ca81d40429e65ccea5385b59cf2d3fc6519371dee92f8eb1dae5107a7a600080a2613b6c565b6112c8565b601f82116001146113655760009055505b388061131c565b61138d61139e926001601f61137f85600052602060002090565b920160051c820191016134b3565b600081815260208120918190559055565b61135e565b005b9291926113b182610ccc565b916113bf6040519384610c4b565b82948184528183011161027c578281602093846000960137010152565b9080601f8301121561027c578160206105e9933591016113a5565b909161140e6105e99360408452604084019061047a565b91602081840391015261047a565b3461027c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c57600435611457816110e5565b60243567ffffffffffffffff811161027c576114779036906004016113dc565b906114ee6114e961149b6110ae8463ffffffff166000526010602052604060002090565b9360066114e26114af836020815191012090565b9660016114cc8863ffffffff166000526010602052604060002090565b019063ffffffff16600052602052604060002090565b0190613d04565b613ad7565b9060609273ffffffffffffffffffffffffffffffffffffffff61153b6001611520846000526003602052604060002090565b015473ffffffffffffffffffffffffffffffffffffffff1690565b16611552575b505061069a604051928392836113f7565b6115bb9293509061157a61074c61074c60016115206000966000526003602052604060002090565b60405180809581947f8318ed5d0000000000000000000000000000000000000000000000000000000083526004830191909163ffffffff6020820193169052565b03915afa9081156115fa576000916115d7575b50903880611541565b6115f491503d806000833e6115ec8183610c4b565b810190613d2a565b386115ce565b613d89565b3461027c5761160d36610281565b9061163061074c73ffffffffffffffffffffffffffffffffffffffff6001541690565b3314159160005b81811061164057005b61164b8183856138d6565b359061166182600052600f602052604060002090565b600181015480156117c35760058201805461177e5750815463ffffffff604082901c1680611746575087908161170b575b50610b39576001936116d560027f5254e609a97bab37b7cc79fe128f85c097bd6015c6e1624ae0ba392eb9753205946116cd61170295615e4e565b500154615f21565b506116f26116ed82600052600f602052604060002090565b613d95565b6040519081529081906020820190565b0390a101611637565b61173d915061172361079f61074c9263ffffffff1690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b33141538611692565b7f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602485905260446000fd5b8461178e610ff061041893615d94565b7f60a6d8980000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602452604490565b7fd82f6adb00000000000000000000000000000000000000000000000000000000600052600484905260246000fd5b3461027c5761180036610281565b90611809614a31565b60005b82811061181557005b611828611823828585613de6565b613e26565b9061184a61074c835173ffffffffffffffffffffffffffffffffffffffff1690565b156119935760019161186160155463ffffffff1690565b907f78e94ca80be2c30abc061b99e7eb8583b1254781734b1e3ce339abb57da2fe8e73ffffffffffffffffffffffffffffffffffffffff61196f6118b9845173ffffffffffffffffffffffffffffffffffffffff1690565b9361190f602082019586516118eb6118cf610c8c565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015261190a8863ffffffff16600052600e602052604060002090565b613e71565b61195561192461086660155463ffffffff1690565b63ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000006015541617601555565b5173ffffffffffffffffffffffffffffffffffffffff1690565b92519261198a63ffffffff6040519384931696169482610d06565b0390a30161180c565b7feeacd9390000000000000000000000000000000000000000000000000000000060005260046000fd5b3461027c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5761069a6119fa60043561401b565b604051918291602083526020830190610568565b9181601f8401121561027c5782359167ffffffffffffffff831161027c576020838186019501011161027c57565b908160c091031261027c5790565b3461027c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5760043567ffffffffffffffff811161027c57611a99903690600401611a0e565b9060243567ffffffffffffffff811161027c57611aba903690600401611a3c565b91611ac3614a31565b611ad061106582846141b3565b9163ffffffff831615611c14575050611af98163ffffffff166000526010602052604060002090565b611b066060840184614212565b611b136040860186614212565b8454909691949060201c63ffffffff16611b2c90613b6c565b81547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16602082901b67ffffffff000000001617825591611b6f60a08201614270565b915460401c60ff165b611b8460808301614285565b611b8e838061428f565b92909360208101611b9e9161428f565b959096611ba9610c9d565b63ffffffff909c168c5263ffffffff1660208c0152151560408b0152151560608a015260ff1660808901523690611bdf926113a5565b60a08701523690611bef926113a5565b60c08501523690611bff926142e0565b923690611c0b9261432c565b6113a3926151da565b6104536040519283927f4071db54000000000000000000000000000000000000000000000000000000008452600484016141cc565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261027c576004359067ffffffffffffffff821161027c576102ca91600401611a0e565b3461027c57611caa611ca336611c49565b36916113a5565b602081519101206000526012602052611cc6604060002061490d565b60405180916020820160208352815180915260206040840192019060005b818110611cf2575050500390f35b8251845285945060209384019390920191600101611ce4565b3461027c576020611d3a611d21611ca336611c49565b8281519101206000526007602052604060002054151590565b6040519015158152f35b3461027c57611d5236611c49565b611d5a613bc9565b5063ffffffff6040518284823760208184810160028152030190205416918215611c145761069a611daa84806000526010602052611da4604060002063ffffffff905460201c1690565b90614f12565b604051918291826111e5565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611de957505050505090565b9091929394602080611e25837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516110f3565b97019301930191939290611dda565b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5760155463ffffffff16611e80611e7b610ff083613b92565b6143e9565b60009163ffffffff811660015b8163ffffffff821610611ec75761069a8486611eab610ff087613b92565b8103611ebf575b5060405191829182611db6565b815282611eb2565b611ee761074c6117238363ffffffff16600052600e602052604060002090565b611efa575b60010163ffffffff16611e8d565b936001611f2363ffffffff926110c16107b58963ffffffff16600052600e602052604060002090565b95915050611eec565b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5760005473ffffffffffffffffffffffffffffffffffffffff81163303611feb577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461027c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5760043567ffffffffffffffff811161027c5761206490369060040161024b565b60243567ffffffffffffffff811161027c5761208490369060040161024b565b928383036123455760015473ffffffffffffffffffffffffffffffffffffffff169060005b8481106120b257005b6120c56120c08287856138d6565b613b88565b906120e08263ffffffff16600052600e602052604060002090565b916120ff835473ffffffffffffffffffffffffffffffffffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff841680156123115761212b611823858c8b613de6565b9061214d61074c835173ffffffffffffffffffffffffffffffffffffffff1690565b1561199357331415806122f1575b610b395760019473ffffffffffffffffffffffffffffffffffffffff61219861074c845173ffffffffffffffffffffffffffffffffffffffff1690565b911614801590612285575b6121b1575b505050016120a9565b73ffffffffffffffffffffffffffffffffffffffff61225f8261224d61220c7f86f41145bde5dd7f523305452e4aad3685508c181432ec733d5f345009358a28955173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6119556020820195898751910161350f565b92519261227a63ffffffff6040519384931696169482610d06565b0390a33880806121a8565b5060405160208101906122ca8161229e89870185614456565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610c4b565b51902060208201516040516122e78161229e602082019485610d06565b51902014156121a3565b5073ffffffffffffffffffffffffffffffffffffffff871633141561215b565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b7fab8b67c6000000000000000000000000000000000000000000000000000000006000526004839052602484905260446000fd5b3461027c5761238736610281565b90612390614a31565b60005b82811061239c57005b6123aa611ca3828585614467565b805160208201206123cb61031f826000526005602052604060002054151590565b61244c5761031f6123db91615c59565b61241657906123eb6001926137dd565b7fb2553249d353abf34f62139c85f44b5bdeab968ec0ab296a9bf735b75200ed83600080a201612393565b610453906040519182917f88c8a73700000000000000000000000000000000000000000000000000000000835260048301610d06565b6040517f4ff9e48a000000000000000000000000000000000000000000000000000000008152806104538460048301610d06565b3461027c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c576004356124bb816110e5565b60243567ffffffffffffffff811161027c576124db903690600401611a3c565b906124e4614a31565b6124fe8163ffffffff166000526010602052604060002090565b549163ffffffff602084901c169081156125535761251f6060820182614212565b909161252e6040820182614212565b96909461253a90613b6c565b9161254760a08201614270565b9160401c60ff16611b78565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461027c576113a36125ea36610281565b906125f3614a31565b61460c565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061262b57505050505090565b9091929394602080612667837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161047a565b9701930193019193929061261c565b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c576126ad614877565b6126b78151613fb4565b9060005b815181101561270d57806126d1600192846138e6565b5160005260136020526111cd6126f1604060002060405192838092613a35565b6126fb82866138e6565b5261270681856138e6565b50016126bb565b6040518061069a85826125f8565b3461027c5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c57600435612756816110e5565b60243567ffffffffffffffff811161027c5761277690369060040161024b565b916044359267ffffffffffffffff841161027c5761279b6113a394369060040161024b565b9390926127a6614a31565b61476c565b3461027c576127b936610281565b6127db61074c60015473ffffffffffffffffffffffffffffffffffffffff1690565b33149060009115925b8183106127ed57005b6127fb6107858484846138fa565b9260408401936128168551600052600f602052604060002090565b9161282b6107b561079f855463ffffffff1690565b90600184018054928315612d1557899081612ceb575b50610b39576020840192835115610a96578351808203612caf575b5050506060830180518015610a695750608084015198895115612c7b579561288f610866875463ffffffff9060201c1690565b86547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16602082901b67ffffffff000000001617875597600487019760005b8c51811015612973576128fe61031f6128eb8f846108cf916138e6565b6000526005602052604060002054151590565b61293f57806129388e6129326108cf8f958f60019761292c919063ffffffff16600052602052604060002090565b936138e6565b90615cc3565b50016128ce565b6040517fdb68f0c2000000000000000000000000000000000000000000000000000000008152806104538f600483016125f8565b50929a509897969593949290612991855463ffffffff9060401c1690565b63ffffffff8116612bb5575b50979198996129ae6005860161490d565b9260009a5b845163ffffffff8d1690811015612b0b576129ce90866138e6565b5163ffffffff169c8d6129f18163ffffffff166000526010602052604060002090565b60010190612a0f9063ffffffff166000526010602052604060002090565b5460201c63ffffffff16612a32919063ffffffff16600052602052604060002090565b600301612a3e9061490d565b9b60005b8d51811015612af057612a8f61031f8f8f8f612a72612a789287929063ffffffff16600052602052604060002090565b926138e6565b519060019160005201602052604060002054151590565b612a9b57600101612a42565b8f612abb8f83612aaa916138e6565b516000526014602052604060002090565b6104536040519283927f16c2b7c400000000000000000000000000000000000000000000000000000000845260048401614958565b50939d509b612b0191929b50613b6c565b9a9b9199906129b3565b509b9a5092509460019750612baa9196507f4b5b465e22eea0c3d40c30e936643245b80d19b2dcf75788c0699fe8d8db645b946003612b5363ffffffff965163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8216178355945197886002840155519101555160405193849316958360209093929193604081019481520152565b0390a20191906127e4565b9a91929899612c1b6003612c158e612c016110ae6001612be9849e9c9d9e63ffffffff166000526010602052604060002090565b019263ffffffff166000526010602052604060002090565b63ffffffff16600052602052604060002090565b0161490d565b9960005b8b51811015612c6957612c4f61031f8d612a78848f8f90612a72919063ffffffff16600052602052604060002090565b612c5b57600101612c1f565b612abb612aaa8f928e6138e6565b5093929b50999850949392943861299d565b6040517fdb68f0c2000000000000000000000000000000000000000000000000000000008152806104538c600483016125f8565b612cc6906000526009602052604060002054151590565b610a9657612cd79184519055615e4e565b50612ce28251615bef565b5038808061285c565b51612d0c915073ffffffffffffffffffffffffffffffffffffffff1661074c565b33141538612841565b88517fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461027c57612d5236610281565b612d5a614a31565b60005b818110612d6657005b8063ffffffff6020612d7b6001948688614467565b91908260405193849283378101600281520301902054168015612da757612da190614c8f565b01612d5d565b50612da1565b3461027c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c57600435612de8816110e5565b60243590612df5826110e5565b612dfd613bc9565b5063ffffffff811680600052601060205263ffffffff60406000205460201c168015612e725763ffffffff841691818311612e3f5761069a611daa8686614f12565b7ff3c16e2c0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b507f2b62be9b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461027c57602063ffffffff81612eb636611c49565b91908260405193849283378101600281520301902054161515604051908152f35b6105e9916060612ef0835160808452608084019061047a565b9273ffffffffffffffffffffffffffffffffffffffff6020820151166020840152604081015115156040840152015190606081840391015261047a565b3461027c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5760043567ffffffffffffffff811161027c57612f87612f8261069a9236906004016113dc565b6149a2565b604051918291602083526020830190612ed7565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310612fce57505050505090565b909192939460208061300a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951612ed7565b97019301930191939290612fbf565b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c576130506148c2565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613096613080846137fd565b9361308e6040519586610c4b565b8085526137fd565b0160005b8181106130ef57505060005b81518110156130e157806130c5612f826114e9612aaa600195876138e6565b6130cf82866138e6565b526130da81856138e6565b50016130a6565b6040518061069a8582612f9b565b6020906130fa61497c565b8282870101520161309a565b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c57604051600a548082528160208101600a60005260206000209260005b8181106131bd57505061316692500382610c4b565b6131708151613869565b9060005b81518110156131af578061319361318d600193856138e6565b5161401b565b61319d82866138e6565b526131a881856138e6565b5001613174565b6040518061069a85826105ec565b8454835260019485019486945060209093019201613151565b359073ffffffffffffffffffffffffffffffffffffffff8216820361027c57565b3461027c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c5760043573ffffffffffffffffffffffffffffffffffffffff811680910361027c5761324f614a31565b3381146132c057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461027c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261027c57602063ffffffff601554821c16604051908152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561339d5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18136030182121561027c570190565b61332e565b60608136031261027c57604051906060820182811067ffffffffffffffff821117610c2a57604052803567ffffffffffffffff811161027c576133e890369083016113dc565b82526133f6602082016131d6565b602083015260408101359067ffffffffffffffff821161027c5761341c913691016113dc565b604082015290565b90600182811c9216801561346d575b602083101461343e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613433565b916134af918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b8181106134be575050565b600081556001016134b3565b9190601f81116134d957505050565b610c9b926000526020600020906020601f840160051c83019310613505575b601f0160051c01906134b3565b90915081906134f8565b919091825167ffffffffffffffff8111610c2a57613537816135318454613424565b846134ca565b6020601f82116001146135915781906134af939495600092613586575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b015190503880613554565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216906135c484600052602060002090565b9160005b81811061361e575095836001959697106135e7575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806135dd565b9192602060018192868b0151815501940192016135c8565b919091825192835167ffffffffffffffff8111610c2a576136618161365b8554613424565b856134ca565b6020601f821160011461372b57916136b782604093600295610c9b98996000926135865750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b84555b6137226136de602083015173ffffffffffffffffffffffffffffffffffffffff1690565b600186019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b0151910161350f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169561375e85600052602060002090565b9660005b8181106137c5575092610c9b969760029593600193836040971061378e575b505050811b0184556136ba565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613781565b83830151895560019098019760209384019301613762565b6137f590602060405192828480945193849201610457565b810103902090565b67ffffffffffffffff8111610c2a5760051b60200190565b60405190610100820182811067ffffffffffffffff821117610c2a57604052606060e08360008152600060208201526000604082015260008382015260006080820152600060a08201528260c08201520152565b90613873826137fd565b6138806040519182610c4b565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06138ae82946137fd565b019060005b8281106138bf57505050565b6020906138ca613815565b828285010152016138b3565b919081101561339d5760051b0190565b805182101561339d5760209160051b010190565b919081101561339d5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561027c570190565b9080601f8301121561027c578135613951816137fd565b9261395f6040519485610c4b565b81845260208085019260051b8201019183831161027c5760208201905b83821061398b57505050505090565b813567ffffffffffffffff811161027c576020916139ae878480948801016113dc565b81520191019061397c565b60a08136031261027c576040519060a0820182811067ffffffffffffffff821117610c2a5760405280356139ec816110e5565b825260208101356020830152604081013560408301526060810135606083015260808101359067ffffffffffffffff821161027c57613a2d9136910161393a565b608082015290565b60009291815491613a4583613424565b8083529260018116908115613a9b5750600114613a6157505050565b60009081526020812093945091925b838310613a81575060209250010190565b600181602092949394548385870101520191019190613a70565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b90610c9b613aeb9260405193848092613a35565b0383610c4b565b9060016020604051613b0381610c0e565b613b39819573ffffffffffffffffffffffffffffffffffffffff8154168352613b326040518096819301613a35565b0384610c4b565b0152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff1663ffffffff8114613b835760010190565b613b3d565b356105e9816110e5565b63ffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9116019063ffffffff8211613b8357565b60405190610140820182811067ffffffffffffffff821117610c2a57604052606061012083600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201528260e0820152826101008201520152565b90613c34826137fd565b613c416040519182610c4b565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613c6f82946137fd565b019060005b828110613c8057505050565b602090613c8b613bc9565b82828501015201613c74565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114613b835760010190565b60405190613cd182610c0e565b6060602083600081520152565b6020613cf7918160405193828580945193849201610457565b8101600281520301902090565b602090613d1e928260405194838680955193849201610457565b82019081520301902090565b60208183031261027c5780519067ffffffffffffffff821161027c570181601f8201121561027c578051613d5d81610ccc565b92613d6b6040519485610c4b565b8184526020828401011161027c576105e99160208085019101610457565b6040513d6000823e3d90fd5b60059060008155600060018201556000600282015560006003820155018054906000815581613dc2575050565b6000526020600020908101905b818110613dda575050565b60008155600101613dcf565b919081101561339d5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18136030182121561027c570190565b60408136031261027c5760405190613e3d82610c0e565b613e46816131d6565b825260208101359067ffffffffffffffff821161027c57613e69913691016113dc565b602082015290565b6001602091939293613ed573ffffffffffffffffffffffffffffffffffffffff865116829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b0192015191825167ffffffffffffffff8111610c2a57613ef9816135318454613424565b6020601f8211600114613f475781906134af9394956000926135865750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690613f7a84600052602060002090565b9160005b818110613f9c575095836001959697106135e757505050811b019055565b9192602060018192868b015181550194019201613f7e565b90613fbe826137fd565b613fcb6040519182610c4b565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613ff982946137fd565b019060005b82811061400a57505050565b806060602080938501015201613ffe565b90614024613815565b5061405e614059600461404185600052600f602052604060002090565b01612c016110ae86600052600f602052604060002090565b61490d565b6140688151613fb4565b9160005b82518110156140a457806140886114e9612aaa600194876138e6565b61409282876138e6565b5261409d81866138e6565b500161406c565b509290506140bf61106582600052600f602052604060002090565b9160026140d683600052600f602052604060002090565b015460016140ee84600052600f602052604060002090565b015490600361410785600052600f602052604060002090565b0154916141976141246110ae87600052600f602052604060002090565b9561418a6141656005612c1561415461414786600052600f602052604060002090565b5460401c63ffffffff1690565b94600052600f602052604060002090565b9761417d614171610cac565b63ffffffff909b168b52565b63ffffffff1660208a0152565b63ffffffff166040880152565b6060860152608085015260a084015260c083015260e082015290565b6020908260405193849283378101600281520301902090565b90601f836040947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe093602086528160208701528686013760008582860101520116010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561027c570180359067ffffffffffffffff821161027c57602001918160051b3603831361027c57565b8015150361027c57565b356105e981614266565b60ff81160361027c57565b356105e98161427a565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561027c570180359067ffffffffffffffff821161027c5760200191813603831361027c57565b9291906142ec816137fd565b936142fa6040519586610c4b565b602085838152019160051b810192831161027c57905b82821061431c57505050565b8135815260209182019101614310565b92919092614339846137fd565b936143476040519586610c4b565b602085828152019060051b82019183831161027c5780915b83831061436d575050505050565b823567ffffffffffffffff811161027c57820160408187031261027c576040519161439783610c0e565b813567ffffffffffffffff811161027c57876143b49184016113dc565b835260208201359267ffffffffffffffff841161027c576143da886020958695016113dc565b8382015281520192019161435f565b906143f3826137fd565b6144006040519182610c4b565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061442e82946137fd565b019060005b82811061443f57505050565b60209061444a613cc4565b82828501015201614433565b9060206105e9928181520190613a35565b9082101561339d576102ca9160051b81019061428f565b919081101561339d5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018136030182121561027c570190565b9080601f8301121561027c578160206105e99335910161432c565b9080601f8301121561027c578160206105e9933591016142e0565b3590610c9b8261427a565b3590610c9b82614266565b6101008136031261027c5761451d610cac565b90803567ffffffffffffffff811161027c5761453c90369083016113dc565b8252602081013567ffffffffffffffff811161027c5761455f903690830161393a565b6020830152604081013567ffffffffffffffff811161027c5761458590369083016113dc565b6040830152606081013567ffffffffffffffff811161027c576145ab90369083016144be565b6060830152608081013567ffffffffffffffff811161027c57614604916145d760e092369083016144d9565b60808501526145e860a082016144f4565b60a08501526145f960c082016144ff565b60c0850152016144ff565b60e082015290565b908015614768579060005b82811061462357505050565b614638614633828585979561447e565b61450a565b9261464c60155463ffffffff9060201c1690565b9261469361465985613b6c565b7fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff000000006015549260201b16911617601555565b6147268460808701516060880151906146af60c08a0151151590565b896147176146c060e0830151151590565b61470d6146d160a085015160ff1690565b9161470460408651960151966146f46146e8610c9d565b63ffffffff909c168c52565b600160208c0152151560408b0152565b15156060890152565b60ff166080870152565b60a085015260c08401526151da565b602060009501945b855180518210156147565790614750614749826001946138e6565b5187615a48565b0161472e565b50509493509150600101919091614617565b5050565b92939163ffffffff6147916110ae8663ffffffff166000526010602052604060002090565b16156148435760005b8181106148235750505060005b8181106147b45750505050565b6147cf6147c5611ca3838588614467565b6020815191012090565b9061480661031f836147f18763ffffffff166000526011602052604060002090565b60019160005201602052604060002054151590565b61481c5761481660019285615b12565b016147a7565b5050505050565b8061483d614837611ca36001948688614467565b87615a48565b0161479a565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff841660045260246000fd5b60405190600c548083528260208101600c60005260206000209260005b8181106148a9575050610c9b92500383610c4b565b8454835260019485019487945060209093019201614894565b604051906004548083528260208101600460005260206000209260005b8181106148f4575050610c9b92500383610c4b565b84548352600194850194879450602090930192016148df565b906040519182815491828252602082019060005260206000209260005b81811061493f575050610c9b92500383610c4b565b845483526001948501948794506020909301920161492a565b9063ffffffff614975602092959495604085526040850190613a35565b9416910152565b6040519061498982610c2f565b6060808381815260006020820152600060408201520152565b6149aa61497c565b50614a298151602083012080600052600360205260026040600020019080600052600360205273ffffffffffffffffffffffffffffffffffffffff614a078160016040600020015416926000526007602052604060002054151590565b9160405195614a1587610c2f565b865216602085015215156040840152613ad7565b606082015290565b73ffffffffffffffffffffffffffffffffffffffff600154163303614a5257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60206000604051828101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a700000000000000000000000000000000000000000000000000000000602482015260248152614adf604482610c4b565b519084617530fa903d6000519083614bda575b5082614bd0575b5081614b35575b81614b09575090565b6105e991507f78bea7210000000000000000000000000000000000000000000000000000000090615d01565b905060206000604051828101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff00000000000000000000000000000000000000000000000000000000602482015260248152614b9a604482610c4b565b519084617530fa6000513d82614bc4575b5081614bba575b501590614b00565b9050151538614bb2565b60201115915038614bab565b1515915038614af9565b60201115925038614af2565b60009291815491614bf683613424565b9260018116908115614c415750600114614c0f57505050565b909192935060005260206000206000905b838210614c2d5750500190565b600181602092548486015201910190614c20565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001683525050811515909102019150565b6020614c849160405192838092614be6565b600281520301902090565b614ca98163ffffffff166000526010602052604060002090565b90815492614cbe8463ffffffff9060201c1690565b906001840190614ceb614ce184849063ffffffff16600052602052604060002090565b9660401c60ff1690565b9260005b8754811015614d75576001908515614d4f57614d49614d21614d11838c615da7565b600052600f602052604060002090565b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff8154169055565b01614cef565b614d6f6005614d61614d11848d615da7565b0163ffffffff8916906160af565b50614d49565b5094549195509293915060201c63ffffffff16156125535760005b614daa8463ffffffff166000526011602052604060002090565b54811015614de75780614de1614ddb600193614dd68863ffffffff166000526011602052604060002090565b615da7565b86615b12565b01614d90565b506005614e0b614e1192614e3994969063ffffffff16600052602052604060002090565b01614c72565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008154169055565b6000614e558263ffffffff166000526010602052604060002090565b557ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c158170365163ffffffff60405192169180614e9281906000602083019252565b0390a2565b90614ea1826137fd565b614eae6040519182610c4b565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0614edc82946137fd565b019060005b828110614eed57505050565b602090604051614efc81610c0e565b6060815260608382015282828501015201614ee1565b9091614f1c613bc9565b50614f378263ffffffff166000526010602052604060002090565b614f5484600183019063ffffffff16600052602052604060002090565b614f606003820161490d565b91614f6b8351614e97565b94600683019460005b8751811015614fce5780614f906114e9612aaa6001948a6138e6565b614fad614f9d8a83613d04565b614fa5610c8c565b928352613ad7565b6020820152614fbc828b6138e6565b52614fc7818a6138e6565b5001614f74565b509350935093949094614ffa614ff48563ffffffff166000526011602052604060002090565b54613fb4565b9560005b6150188663ffffffff166000526011602052604060002090565b5481101561507357806150576114e9615047600194614dd68b63ffffffff166000526011602052604060002090565b6000526013602052604060002090565b615061828b6138e6565b5261506c818a6138e6565b5001614ffe565b5092959194909350549361508a8563ffffffff1690565b9460401c60ff166004840154600881901c60ff169060ff16906150ac8661490d565b936150b5610cbc565b63ffffffff909916895263ffffffff16602089015260ff166040880152151560608701521515608086015260a085015260c08401526150f660058201613ad7565b60e084015260020161510790613ad7565b61010083015261012082015290565b60ff60019116019060ff8211613b8357565b6105e99054613424565b60409063ffffffff6105e99493168152816020820152019061047a565b6040906105e993928152816020820152019061047a565b805482101561339d5760005260206000200190600090565b80549068010000000000000000821015610c2a57816151a59160016134af94018155615166565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b919060016152056151ef845163ffffffff1690565b63ffffffff166000526010602052604060002090565b0190602083019061523361521d835163ffffffff1690565b849063ffffffff16600052602052604060002090565b927f00000000000000000000000000000000000000000000000000000000000000001580615a30575b8015615a0d575b6159c55761528190612c0161527c855163ffffffff1690565b613b92565b916152916147c560058501613ad7565b9260a08601936152a685516020815191012090565b03615939575b60016152bf610ff0845163ffffffff1690565b116158d5575b5060005b8651811015615487576152e961031f6152e2838a6138e6565b5187615cc3565b6154395760608601511561540c57615318614147615307838a6138e6565b51600052600f602052604060002090565b63ffffffff61532e610ff0895163ffffffff1690565b91161415806153ee575b6153a05760019061539a615350885163ffffffff1690565b61535d615307848c6138e6565b907fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff6bffffffff000000000000000083549260401b169116179055565b016152c9565b866153b661041892612a72895163ffffffff1690565b517f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff909116600452602452604490565b5063ffffffff615404614147615307848b6138e6565b161515615338565b8061543360056154216153076001958c6138e6565b01612932610ff08a5163ffffffff1690565b5061539a565b8661544f61041892612a72895163ffffffff1690565b517f636e40570000000000000000000000000000000000000000000000000000000060005263ffffffff909116600452602452604490565b5090956003840195949093600093600682019291600481019160058201916002015b8b5188101561586a576154bc888d6138e6565b519b6154cd8d516020815191012090565b996154e861031f8c6000526005602052604060002054151590565b615833576155038b6000526007602052604060002054151590565b6157fc5761551b6155168f8a9051613d04565b615128565b6157b55760005b8c518110156155a75761555761031f8d8f80612c016110ae615307886004615550615307836147f1996138e6565b01946138e6565b61556357600101615522565b61556e8f918e6138e6565b519051906104536040519283927f4b5786e70000000000000000000000000000000000000000000000000000000084526004840161514f565b509890929c6157aa908c8f8f8b91879f60019860c08e6155cd6155ea946155f19761517e565b6155e260208a01976103408951918c51613d04565b01519061350f565b518961350f565b61563261560160408c0151151590565b8a9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b61567661564360808c015160ff1690565b8a547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1660089190911b61ff0016178a55565b6156d16156878b5163ffffffff1690565b6156a18163ffffffff166000526010602052604060002090565b9063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b61572e6156e160608c0151151590565b6156f26151ef8d5163ffffffff1690565b907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff68ff0000000000000000835492151560401b169116179055565b61578961573f8d5163ffffffff1690565b6157506151ef8d5163ffffffff1690565b907fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff0000000083549260201b169116179055565b895163ffffffff16926157a08d5163ffffffff1690565b90519151936161f5565b0196979a909a6154a9565b8d6157c4885163ffffffff1690565b9051906104536040519283927f368812ac00000000000000000000000000000000000000000000000000000000845260048401615132565b6104538e516040519182917f88c8a73700000000000000000000000000000000000000000000000000000000835260048301610d06565b6104538e516040519182917f4ff9e48a00000000000000000000000000000000000000000000000000000000835260048301610d06565b50505050905063ffffffff9396506158b09195507ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c158170365194506158bb92505163ffffffff1690565b935163ffffffff1690565b60405163ffffffff90911681529216918060208101614e92565b9693916000969391965b885481101561592c578061591560056158fd614d116001958e615da7565b0161590f610ff08b5163ffffffff1690565b906160af565b50615926614d21614d11838d615da7565b016158df565b50919396509194386152c5565b615948614e1160058701614c72565b83518051615957575b506152ac565b61596861106563ffffffff92613cde565b1661598e5761598861597e875163ffffffff1690565b6156a18651613cde565b38615951565b61045384516040519182917f07bf02d600000000000000000000000000000000000000000000000000000000835260048301610d06565b610418866159d7608088015160ff1690565b90517f25b4d6180000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b50615a24615a1f608087015160ff1690565b615116565b60ff8751911611615263565b5060ff615a41608087015160ff1690565b161561525c565b90805160208201209063ffffffff831692836000526011602052615a8083604060002060019160005201602052604060002054151590565b615b0c5782610919615ade92615ae4956000526013602052615aa685604060002061350f565b615aaf83615c8e565b50826000526012602052615ac7876040600020615cc3565b5063ffffffff166000526011602052604060002090565b506137dd565b907fc00ca38a0d4dd24af204fcc9a39d94708b58426bcf57796b94c4b5437919ede2600080a3565b50505050565b63ffffffff1690816000526011602052615b308160406000206160af565b50806000526012602052615b488260406000206160af565b5080600052601260205260406000205415615ba5575b6000526013602052615b7a604060002060405191828092614be6565b039020907f257129637d1e1b80e89cae4f5e49de63c09628e1622724b24dd19b406627de30600080a3565b615bae81615fe8565b50615b5e565b600081815260056020526040902054615be957615bd281600461517e565b600454906000526005602052604060002055600190565b50600090565b600081815260096020526040902054615be957615c0d81600861517e565b600854906000526009602052604060002055600190565b6000818152600b6020526040902054615be957615c4281600a61517e565b600a5490600052600b602052604060002055600190565b600081815260076020526040902054615be957615c7781600661517e565b600654906000526007602052604060002055600190565b6000818152600d6020526040902054615be957615cac81600c61517e565b600c5490600052600d602052604060002055600190565b6000828152600182016020526040902054615cfa5780615ce58360019361517e565b80549260005201602052604060002055600190565b5050600090565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a700000000000000000000000000000000000000000000000000000000845216602482015260248152615d67604482610c4b565b5191617530fa6000513d82615d88575b5081615d81575090565b9050151590565b60201115915038615d77565b80541561339d5760005260206000205490565b90615db191615166565b90549060031b1c90565b80548015615e1f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615df08282615166565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260096020526040902054908115615cfa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211613b8357600854927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411613b83578383615eea9460009603615ef0575b505050615ed96008615dbb565b600990600052602052604060002090565b55600190565b615ed9615f1291615f08615db1615f18956008615166565b9283916008615166565b90613477565b55388080615ecc565b6000818152600b6020526040902054908115615cfa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211613b8357600a54927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411613b83578383615eea9460009603615fbd575b505050615fac600a615dbb565b600b90600052602052604060002090565b615fac615f1291615fd5615db1615fdf95600a615166565b928391600a615166565b55388080615f9f565b6000818152600d6020526040902054908115615cfa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820190828211613b8357600c54927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411613b83578383615eea9460009603616084575b505050616073600c615dbb565b600d90600052602052604060002090565b616073615f129161609c615db16160a695600c615166565b928391600c615166565b55388080616066565b6001810191806000528260205260406000205492831515600014616187577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111613b83578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8501948511613b8357600095858361613f94615eea980361614e575b505050615dbb565b90600052602052604060002090565b61616e615f1291616165615db161617e9588615166565b92839187615166565b8590600052602052604060002090565b55388080616137565b50505050600090565b9294939160808401608085528251809152602060a0860193019060005b8181106161df575050509163ffffffff6161d28360609587849603602089015261047a565b9616604085015216910152565b82518552602094850194909201916001016161ad565b9390916020815191012073ffffffffffffffffffffffffffffffffffffffff80600161622b846000526003602052604060002090565b0154161661623a575050505050565b61074c61074c6001611520616259946000526003602052604060002090565b90813b1561027c576000809461629e604051978896879586947ffba64a7c00000000000000000000000000000000000000000000000000000000865260048601616190565b03925af180156115fa576162b6575b8080808061481c565b806162c560006162cb93610c4b565b80610bd4565b386162ad56fea164736f6c634300081a000a",
}

var CapabilitiesRegistryABI = CapabilitiesRegistryMetaData.ABI

var CapabilitiesRegistryBin = CapabilitiesRegistryMetaData.Bin

func DeployCapabilitiesRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, params CapabilitiesRegistryConstructorParams) (common.Address, *types.Transaction, *CapabilitiesRegistry, error) {
	parsed, err := CapabilitiesRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CapabilitiesRegistryBin), backend, params)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONFamilies(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONFamilies")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONFamilies() ([]string, error) {
	return _CapabilitiesRegistry.Contract.GetDONFamilies(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONFamilies() ([]string, error) {
	return _CapabilitiesRegistry.Contract.GetDONFamilies(&_CapabilitiesRegistry.CallOpts)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) SetDONFamilies(opts *bind.TransactOpts, donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "setDONFamilies", donId, addToFamilies, removeFromFamilies)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) SetDONFamilies(donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamilies(&_CapabilitiesRegistry.TransactOpts, donId, addToFamilies, removeFromFamilies)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) SetDONFamilies(donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamilies(&_CapabilitiesRegistry.TransactOpts, donId, addToFamilies, removeFromFamilies)
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

type CapabilitiesRegistryDONAddedToFamilyIterator struct {
	Event *CapabilitiesRegistryDONAddedToFamily

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryDONAddedToFamilyIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryDONAddedToFamily)
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
		it.Event = new(CapabilitiesRegistryDONAddedToFamily)
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

func (it *CapabilitiesRegistryDONAddedToFamilyIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryDONAddedToFamilyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryDONAddedToFamily struct {
	DonId     uint32
	DonFamily common.Hash
	Raw       types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterDONAddedToFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONAddedToFamilyIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "DONAddedToFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryDONAddedToFamilyIterator{contract: _CapabilitiesRegistry.contract, event: "DONAddedToFamily", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchDONAddedToFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONAddedToFamily, donId []uint32, donFamily []string) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "DONAddedToFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryDONAddedToFamily)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONAddedToFamily", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseDONAddedToFamily(log types.Log) (*CapabilitiesRegistryDONAddedToFamily, error) {
	event := new(CapabilitiesRegistryDONAddedToFamily)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONAddedToFamily", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryDONRemovedFromFamilyIterator struct {
	Event *CapabilitiesRegistryDONRemovedFromFamily

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryDONRemovedFromFamilyIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryDONRemovedFromFamily)
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
		it.Event = new(CapabilitiesRegistryDONRemovedFromFamily)
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

func (it *CapabilitiesRegistryDONRemovedFromFamilyIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryDONRemovedFromFamilyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryDONRemovedFromFamily struct {
	DonId     uint32
	DonFamily common.Hash
	Raw       types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterDONRemovedFromFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONRemovedFromFamilyIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "DONRemovedFromFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryDONRemovedFromFamilyIterator{contract: _CapabilitiesRegistry.contract, event: "DONRemovedFromFamily", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchDONRemovedFromFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONRemovedFromFamily, donId []uint32, donFamily []string) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "DONRemovedFromFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryDONRemovedFromFamily)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONRemovedFromFamily", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseDONRemovedFromFamily(log types.Log) (*CapabilitiesRegistryDONRemovedFromFamily, error) {
	event := new(CapabilitiesRegistryDONRemovedFromFamily)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONRemovedFromFamily", log); err != nil {
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
	case _CapabilitiesRegistry.abi.Events["DONAddedToFamily"].ID:
		return _CapabilitiesRegistry.ParseDONAddedToFamily(log)
	case _CapabilitiesRegistry.abi.Events["DONRemovedFromFamily"].ID:
		return _CapabilitiesRegistry.ParseDONRemovedFromFamily(log)
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

func (CapabilitiesRegistryDONAddedToFamily) Topic() common.Hash {
	return common.HexToHash("0xc00ca38a0d4dd24af204fcc9a39d94708b58426bcf57796b94c4b5437919ede2")
}

func (CapabilitiesRegistryDONRemovedFromFamily) Topic() common.Hash {
	return common.HexToHash("0x257129637d1e1b80e89cae4f5e49de63c09628e1622724b24dd19b406627de30")
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

	GetDONFamilies(opts *bind.CallOpts) ([]string, error)

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

	SetDONFamilies(opts *bind.TransactOpts, donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error)

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

	FilterDONAddedToFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONAddedToFamilyIterator, error)

	WatchDONAddedToFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONAddedToFamily, donId []uint32, donFamily []string) (event.Subscription, error)

	ParseDONAddedToFamily(log types.Log) (*CapabilitiesRegistryDONAddedToFamily, error)

	FilterDONRemovedFromFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONRemovedFromFamilyIterator, error)

	WatchDONRemovedFromFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONRemovedFromFamily, donId []uint32, donFamily []string) (event.Subscription, error)

	ParseDONRemovedFromFamily(log types.Log) (*CapabilitiesRegistryDONRemovedFromFamily, error)

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
