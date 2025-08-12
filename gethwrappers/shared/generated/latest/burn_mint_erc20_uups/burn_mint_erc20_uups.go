// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_uups

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

var BurnMintERC20UUPSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultUpgrader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20UUPS__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20UUPS__MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516135cf6100fd60003960008181611abe01528181611ae70152611cef01526135cf6000f3fe6080604052600436106102d15760003560e01c806384ef8ffc11610179578063ad3cb1cc116100d6578063d53913931161008a578063d602b9fd11610064578063d602b9fd146109ac578063dd62ed3e146109c1578063f72c0d8b14610a3357600080fd5b8063d539139314610924578063d547741f14610958578063d5abeb011461097857600080fd5b8063cc8463c8116100bb578063cc8463c814610882578063cefc142914610897578063cf6eefb7146108ac57600080fd5b8063ad3cb1cc14610819578063c630948d1461086257600080fd5b80639dc29fac1161012d578063a217fddf11610112578063a217fddf146107c4578063a8fa343c146107d9578063a9059cbb146107f957600080fd5b80639dc29fac14610770578063a1eda53c1461079057600080fd5b80638fd6a6ac1161015e5780638fd6a6ac146106d457806391d14854146106e957806395d89b411461075b57600080fd5b806384ef8ffc146106545780638da5cb5b146106bf57600080fd5b8063313ce5671161023257806352d1902d116101e6578063649a5ec7116101c0578063649a5ec7146105b257806370a08231146105d257806379cc67901461063457600080fd5b806352d1902d1461055d578063561cf2ab14610572578063634e93da1461059257600080fd5b806340c10f191161021757806340c10f191461050a57806342966c681461052a5780634f1ef2861461054a57600080fd5b8063313ce5671461048e57806336568abe146104ea57600080fd5b806318160ddd11610289578063248a9ca31161026e578063248a9ca3146103eb578063282c51f31461043a5780632f2ff15d1461046e57600080fd5b806318160ddd1461038d57806323b872dd146103cb57600080fd5b806306fdde03116102ba57806306fdde0314610334578063095ea7b3146103565780630aa6220b1461037657600080fd5b806301ffc9a7146102d6578063022d63fb1461030b575b600080fd5b3480156102e257600080fd5b506102f66102f1366004612f64565b610a67565b60405190151581526020015b60405180910390f35b34801561031757600080fd5b50620697805b60405165ffffffffffff9091168152602001610302565b34801561034057600080fd5b50610349610c30565b6040516103029190612fca565b34801561036257600080fd5b506102f6610371366004613044565b610d05565b34801561038257600080fd5b5061038b610d1d565b005b34801561039957600080fd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b604051908152602001610302565b3480156103d757600080fd5b506102f66103e636600461306e565b610d33565b3480156103f757600080fd5b506103bd6104063660046130ab565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561044657600080fd5b506103bd7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b34801561047a57600080fd5b5061038b6104893660046130c4565b610d59565b34801561049a57600080fd5b507fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff005474010000000000000000000000000000000000000000900460ff1660405160ff9091168152602001610302565b3480156104f657600080fd5b5061038b6105053660046130c4565b610d9e565b34801561051657600080fd5b5061038b610525366004613044565b610f0a565b34801561053657600080fd5b5061038b6105453660046130ab565b610fee565b61038b6105583660046131b7565b611021565b34801561056957600080fd5b506103bd61103c565b34801561057e57600080fd5b5061038b61058d366004613239565b61106b565b34801561059e57600080fd5b5061038b6105ad3660046132ef565b611337565b3480156105be57600080fd5b5061038b6105cd36600461330a565b61134b565b3480156105de57600080fd5b506103bd6105ed3660046132ef565b73ffffffffffffffffffffffffffffffffffffffff1660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b34801561064057600080fd5b5061038b61064f366004613044565b61135f565b34801561066057600080fd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610302565b3480156106cb57600080fd5b5061069a611393565b3480156106e057600080fd5b5061069a6113d8565b3480156106f557600080fd5b506102f66107043660046130c4565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561076757600080fd5b50610349611418565b34801561077c57600080fd5b5061038b61078b366004613044565b611469565b34801561079c57600080fd5b506107a5611473565b6040805165ffffffffffff938416815292909116602083015201610302565b3480156107d057600080fd5b506103bd600081565b3480156107e557600080fd5b5061038b6107f43660046132ef565b611532565b34801561080557600080fd5b506102f6610814366004613044565b6115d4565b34801561082557600080fd5b506103496040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561086e57600080fd5b5061038b61087d3660046132ef565b6115e2565b34801561088e57600080fd5b5061031d611636565b3480156108a357600080fd5b5061038b611717565b3480156108b857600080fd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610302565b34801561093057600080fd5b506103bd7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b34801561096457600080fd5b5061038b6109733660046130c4565b611792565b34801561098457600080fd5b507fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff01546103bd565b3480156109b857600080fd5b5061038b6117d3565b3480156109cd57600080fd5b506103bd6109dc366004613332565b73ffffffffffffffffffffffffffffffffffffffff91821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610a3f57600080fd5b506103bd7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b07000000000000000000000000000000000000000000000000000000001480610afa57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b80610b4657507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b80610b9257507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b80610bde57507fffffffff0000000000000000000000000000000000000000000000000000000082167f52d1902d00000000000000000000000000000000000000000000000000000000145b80610c2a57507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac00000000000000000000000000000000000000000000000000000000145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610c819061335c565b80601f0160208091040260200160405190810160405280929190818152602001828054610cad9061335c565b8015610cfa5780601f10610ccf57610100808354040283529160200191610cfa565b820191906000526020600020905b815481529060010190602001808311610cdd57829003601f168201915b505050505091505090565b600033610d138185856117e6565b5060019392505050565b6000610d28816117f3565b610d306117fd565b50565b600033610d4185828561180a565b610d4c8585856118f8565b60019150505b9392505050565b81610d90576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d9a82826119a3565b5050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840082158015610e0657507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8381169116145b15610efb577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610e79575065ffffffffffff8116155b80610e8c57504265ffffffffffff821610155b15610ed2576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b610f0583836119e7565b505050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6610f34816117f3565b7fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff01546000610f807f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90508115801590610f99575081610f9785836133de565b115b15610fdd57610fa884826133de565b6040517f193e2453000000000000000000000000000000000000000000000000000000008152600401610ec991815260200190565b610fe78585611a40565b5050505050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848611018816117f3565b610d9a82611a9c565b611029611aa6565b61103282611b74565b610d9a8282611b9e565b6000611046611cd7565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156110b65750825b905060008267ffffffffffffffff1660011480156110d35750303b155b9050811580156110e1575080155b15611118576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156111795784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6111838c8c611d46565b61118b611d58565b611193611d58565b61119b611d58565b60007fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff008054600182018c90557fffffffffffffffffffffff000000000000000000000000000000000000000000167401000000000000000000000000000000000000000060ff8e16027fffffffffffffffffffffffff0000000000000000000000000000000000000000161773ffffffffffffffffffffffffffffffffffffffff8a16178155905088156112905789891115611286576040517f193e2453000000000000000000000000000000000000000000000000000000008152600481018a9052602401610ec9565b611290888a611a40565b61129b600089611d60565b506112c67f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e388611d60565b505083156113295784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b6000611342816117f3565b610d9a82611e6a565b6000611356816117f3565b610d9a82611eea565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848611389816117f3565b610f058383611f5a565b60006113d37feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b60007fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff005b5473ffffffffffffffffffffffffffffffffffffffff16919050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610c819061335c565b610d9a828261135f565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840081158015906114f657504265ffffffffffff831610155b61150257600080611529565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b600061153d816117f3565b7fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff0080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff858116918217845560405192169182907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a350505050565b600033610d138185856118f8565b61160c7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610d59565b610d307f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610d59565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680158015906116b957504265ffffffffffff8216105b6116ea5781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611710565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff1633811461178a576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610ec9565b610d30611f6f565b816117c9576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d9a82826120a0565b60006117de816117f3565b610d306120e4565b610f0583838360016120ef565b610d308133612162565b611808600080612209565b565b73ffffffffffffffffffffffffffffffffffffffff83811660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146118f257818110156118e3576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610ec9565b6118f2848484840360006120ef565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611948576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610ec9565b73ffffffffffffffffffffffffffffffffffffffff8216611998576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610ec9565b610f058383836123a2565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546119dd816117f3565b6118f28383611d60565b73ffffffffffffffffffffffffffffffffffffffff81163314611a36576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f058282612414565b73ffffffffffffffffffffffffffffffffffffffff8216611a90576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610ec9565b610d9a600083836123a2565b610d3033826124b8565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480611b3d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611b24612514565b73ffffffffffffffffffffffffffffffffffffffff1614155b15611808576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610d9a816117f3565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611c23575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611c20918101906133f1565b60015b611c71576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610ec9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611ccd576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610ec9565b610f05838361253c565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611808576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d4e61259f565b610d9a8282612606565b61180861259f565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083611e58576000611dc87feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614611e15576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85161790555b611e628484612669565b949350505050565b6000611e74611636565b611e7d4261278a565b611e87919061340a565b9050611e9382826127da565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b6000611ef582612895565b611efe4261278a565b611f08919061340a565b9050611f148282612209565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b611f6582338361180a565b610d9a82826124b8565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400805473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580611fdf57504265ffffffffffff821610155b15612020576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610ec9565b61206860006120637feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b612414565b50612074600083611d60565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546120da816117f3565b6118f28383612414565b6118086000806127da565b3073ffffffffffffffffffffffffffffffffffffffff841603612156576040517f99817ca200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610ec9565b6118f2848484846128dd565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610d9a576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610ec9565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801561231c574265ffffffffffff821610156122f2576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a0100000000000000000000000000000000000000000000000000000217825561231c565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600101805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b3073ffffffffffffffffffffffffffffffffffffffff831603612409576040517f99817ca200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610ec9565b610f05838383612a49565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561247e57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8481169116145b156124ae576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b611e628484612c1a565b73ffffffffffffffffffffffffffffffffffffffff8216612508576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610ec9565b610d9a826000836123a2565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6113fc565b61254582612cf8565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561259757610f058282612dc7565b610d9a612e4a565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611808576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61260e61259f565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0361265a848261346f565b50600481016118f2838261346f565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166127805760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561271c3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610c2a565b6000915050610c2a565b600065ffffffffffff8211156127d6576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610ec9565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff88161717845591041680156118f2576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a150505050565b6000806128a0611636565b90508065ffffffffffff168365ffffffffffff16116128c8576128c38382613588565b610d52565b610d5265ffffffffffff841662069780612e82565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff851661294e576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610ec9565b73ffffffffffffffffffffffffffffffffffffffff841661299e576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610ec9565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260018301602090815260408083209388168352929052208390558115610fe7578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051612a3a91815260200190565b60405180910390a35050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8416612aa45781816002016000828254612a9991906133de565b90915550612b569050565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020829052604090205482811015612b2a576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610ec9565b73ffffffffffffffffffffffffffffffffffffffff851660009081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316612b81576002810180548390039055612bad565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051612c0c91815260200190565b60405180910390a350505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156127805760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610c2a565b8073ffffffffffffffffffffffffffffffffffffffff163b600003612d61576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610ec9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60606000808473ffffffffffffffffffffffffffffffffffffffff1684604051612df191906135a6565b600060405180830381855af49150503d8060008114612e2c576040519150601f19603f3d011682016040523d82523d6000602084013e612e31565b606091505b5091509150612e41858383612e98565b95945050505050565b3415611808576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000818310612e915781610d52565b5090919050565b606082612ea8576128c382612f22565b8151158015612ecc575073ffffffffffffffffffffffffffffffffffffffff84163b155b15612f1b576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610ec9565b5080610d52565b805115612f325780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215612f7657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610d5257600080fd5b60005b83811015612fc1578181015183820152602001612fa9565b50506000910152565b6020815260008251806020840152612fe9816040850160208701612fa6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461303f57600080fd5b919050565b6000806040838503121561305757600080fd5b6130608361301b565b946020939093013593505050565b60008060006060848603121561308357600080fd5b61308c8461301b565b925061309a6020850161301b565b929592945050506040919091013590565b6000602082840312156130bd57600080fd5b5035919050565b600080604083850312156130d757600080fd5b823591506130e76020840161301b565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff84111561313a5761313a6130f0565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715613187576131876130f0565b60405283815290508082840185101561319f57600080fd5b83836020830137600060208583010152509392505050565b600080604083850312156131ca57600080fd5b6131d38361301b565b9150602083013567ffffffffffffffff8111156131ef57600080fd5b8301601f8101851361320057600080fd5b61320f8582356020840161311f565b9150509250929050565b600082601f83011261322a57600080fd5b610d528383356020850161311f565b600080600080600080600060e0888a03121561325457600080fd5b873567ffffffffffffffff81111561326b57600080fd5b6132778a828b01613219565b975050602088013567ffffffffffffffff81111561329457600080fd5b6132a08a828b01613219565b965050604088013560ff811681146132b757600080fd5b945060608801359350608088013592506132d360a0890161301b565b91506132e160c0890161301b565b905092959891949750929550565b60006020828403121561330157600080fd5b610d528261301b565b60006020828403121561331c57600080fd5b813565ffffffffffff81168114610d5257600080fd5b6000806040838503121561334557600080fd5b61334e8361301b565b91506130e76020840161301b565b600181811c9082168061337057607f821691505b6020821081036133a9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610c2a57610c2a6133af565b60006020828403121561340357600080fd5b5051919050565b65ffffffffffff8181168382160190811115610c2a57610c2a6133af565b601f821115610f0557806000526020600020601f840160051c8101602085101561344f5750805b601f840160051c820191505b81811015610fe7576000815560010161345b565b815167ffffffffffffffff811115613489576134896130f0565b61349d81613497845461335c565b84613428565b6020601f8211600181146134ef57600083156134b95750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455610fe7565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561353d578785015182556020948501946001909201910161351d565b508482101561357957868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b65ffffffffffff8281168282160390811115610c2a57610c2a6133af565b600082516135b8818460208701612fa6565b919091019291505056fea164736f6c634300081a000a",
}

var BurnMintERC20UUPSABI = BurnMintERC20UUPSMetaData.ABI

var BurnMintERC20UUPSBin = BurnMintERC20UUPSMetaData.Bin

func DeployBurnMintERC20UUPS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnMintERC20UUPS, error) {
	parsed, err := BurnMintERC20UUPSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20UUPSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20UUPS{address: address, abi: *parsed, BurnMintERC20UUPSCaller: BurnMintERC20UUPSCaller{contract: contract}, BurnMintERC20UUPSTransactor: BurnMintERC20UUPSTransactor{contract: contract}, BurnMintERC20UUPSFilterer: BurnMintERC20UUPSFilterer{contract: contract}}, nil
}

type BurnMintERC20UUPS struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20UUPSCaller
	BurnMintERC20UUPSTransactor
	BurnMintERC20UUPSFilterer
}

type BurnMintERC20UUPSCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC20UUPSTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20UUPSFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20UUPSSession struct {
	Contract     *BurnMintERC20UUPS
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20UUPSCallerSession struct {
	Contract *BurnMintERC20UUPSCaller
	CallOpts bind.CallOpts
}

type BurnMintERC20UUPSTransactorSession struct {
	Contract     *BurnMintERC20UUPSTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20UUPSRaw struct {
	Contract *BurnMintERC20UUPS
}

type BurnMintERC20UUPSCallerRaw struct {
	Contract *BurnMintERC20UUPSCaller
}

type BurnMintERC20UUPSTransactorRaw struct {
	Contract *BurnMintERC20UUPSTransactor
}

func NewBurnMintERC20UUPS(address common.Address, backend bind.ContractBackend) (*BurnMintERC20UUPS, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20UUPSABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20UUPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPS{address: address, abi: abi, BurnMintERC20UUPSCaller: BurnMintERC20UUPSCaller{contract: contract}, BurnMintERC20UUPSTransactor: BurnMintERC20UUPSTransactor{contract: contract}, BurnMintERC20UUPSFilterer: BurnMintERC20UUPSFilterer{contract: contract}}, nil
}

func NewBurnMintERC20UUPSCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20UUPSCaller, error) {
	contract, err := bindBurnMintERC20UUPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSCaller{contract: contract}, nil
}

func NewBurnMintERC20UUPSTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20UUPSTransactor, error) {
	contract, err := bindBurnMintERC20UUPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSTransactor{contract: contract}, nil
}

func NewBurnMintERC20UUPSFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20UUPSFilterer, error) {
	contract, err := bindBurnMintERC20UUPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSFilterer{contract: contract}, nil
}

func bindBurnMintERC20UUPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20UUPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20UUPS.Contract.BurnMintERC20UUPSCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.BurnMintERC20UUPSTransactor.contract.Transfer(opts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.BurnMintERC20UUPSTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20UUPS.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.BURNERROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.BURNERROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.DEFAULTADMINROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.DEFAULTADMINROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.MINTERROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.MINTERROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) UPGRADERROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.UPGRADERROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.UPGRADERROLE(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BurnMintERC20UUPS.Contract.UPGRADEINTERFACEVERSION(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BurnMintERC20UUPS.Contract.UPGRADEINTERFACEVERSION(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.Allowance(&_BurnMintERC20UUPS.CallOpts, owner, spender)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.Allowance(&_BurnMintERC20UUPS.CallOpts, owner, spender)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.BalanceOf(&_BurnMintERC20UUPS.CallOpts, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.BalanceOf(&_BurnMintERC20UUPS.CallOpts, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Decimals() (uint8, error) {
	return _BurnMintERC20UUPS.Contract.Decimals(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20UUPS.Contract.Decimals(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20UUPS.Contract.DefaultAdmin(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20UUPS.Contract.DefaultAdmin(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.DefaultAdminDelay(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.DefaultAdminDelay(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20UUPS.Contract.GetCCIPAdmin(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20UUPS.Contract.GetCCIPAdmin(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.GetRoleAdmin(&_BurnMintERC20UUPS.CallOpts, role)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.GetRoleAdmin(&_BurnMintERC20UUPS.CallOpts, role)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20UUPS.Contract.HasRole(&_BurnMintERC20UUPS.CallOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20UUPS.Contract.HasRole(&_BurnMintERC20UUPS.CallOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.MaxSupply(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.MaxSupply(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Name() (string, error) {
	return _BurnMintERC20UUPS.Contract.Name(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) Name() (string, error) {
	return _BurnMintERC20UUPS.Contract.Name(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Owner() (common.Address, error) {
	return _BurnMintERC20UUPS.Contract.Owner(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC20UUPS.Contract.Owner(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20UUPS.Contract.PendingDefaultAdmin(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20UUPS.Contract.PendingDefaultAdmin(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20UUPS.Contract.PendingDefaultAdminDelay(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20UUPS.Contract.PendingDefaultAdminDelay(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) ProxiableUUID() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.ProxiableUUID(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BurnMintERC20UUPS.Contract.ProxiableUUID(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20UUPS.Contract.SupportsInterface(&_BurnMintERC20UUPS.CallOpts, interfaceId)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20UUPS.Contract.SupportsInterface(&_BurnMintERC20UUPS.CallOpts, interfaceId)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Symbol() (string, error) {
	return _BurnMintERC20UUPS.Contract.Symbol(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) Symbol() (string, error) {
	return _BurnMintERC20UUPS.Contract.Symbol(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20UUPS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.TotalSupply(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20UUPS.Contract.TotalSupply(&_BurnMintERC20UUPS.CallOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20UUPS.TransactOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20UUPS.TransactOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Approve(&_BurnMintERC20UUPS.TransactOpts, spender, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Approve(&_BurnMintERC20UUPS.TransactOpts, spender, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20UUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20UUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Burn(&_BurnMintERC20UUPS.TransactOpts, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Burn(&_BurnMintERC20UUPS.TransactOpts, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Burn0(&_BurnMintERC20UUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Burn0(&_BurnMintERC20UUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.BurnFrom(&_BurnMintERC20UUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.BurnFrom(&_BurnMintERC20UUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20UUPS.TransactOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20UUPS.TransactOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20UUPS.TransactOpts, newDelay)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20UUPS.TransactOpts, newDelay)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.GrantMintAndBurnRoles(&_BurnMintERC20UUPS.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.GrantMintAndBurnRoles(&_BurnMintERC20UUPS.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.GrantRole(&_BurnMintERC20UUPS.TransactOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.GrantRole(&_BurnMintERC20UUPS.TransactOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "initialize", name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Initialize(&_BurnMintERC20UUPS.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Initialize(&_BurnMintERC20UUPS.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Mint(&_BurnMintERC20UUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Mint(&_BurnMintERC20UUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.RenounceRole(&_BurnMintERC20UUPS.TransactOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.RenounceRole(&_BurnMintERC20UUPS.TransactOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.RevokeRole(&_BurnMintERC20UUPS.TransactOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.RevokeRole(&_BurnMintERC20UUPS.TransactOpts, role, account)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20UUPS.TransactOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20UUPS.TransactOpts)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.SetCCIPAdmin(&_BurnMintERC20UUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.SetCCIPAdmin(&_BurnMintERC20UUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Transfer(&_BurnMintERC20UUPS.TransactOpts, to, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.Transfer(&_BurnMintERC20UUPS.TransactOpts, to, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.TransferFrom(&_BurnMintERC20UUPS.TransactOpts, from, to, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.TransferFrom(&_BurnMintERC20UUPS.TransactOpts, from, to, value)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.UpgradeToAndCall(&_BurnMintERC20UUPS.TransactOpts, newImplementation, data)
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20UUPS.Contract.UpgradeToAndCall(&_BurnMintERC20UUPS.TransactOpts, newImplementation, data)
}

type BurnMintERC20UUPSApprovalIterator struct {
	Event *BurnMintERC20UUPSApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSApproval)
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
		it.Event = new(BurnMintERC20UUPSApproval)
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

func (it *BurnMintERC20UUPSApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20UUPSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSApprovalIterator{contract: _BurnMintERC20UUPS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSApproval)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseApproval(log types.Log) (*BurnMintERC20UUPSApproval, error) {
	event := new(BurnMintERC20UUPSApproval)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSCCIPAdminTransferredIterator struct {
	Event *BurnMintERC20UUPSCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSCCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20UUPSCCIPAdminTransferred)
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

func (it *BurnMintERC20UUPSCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20UUPSCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSCCIPAdminTransferredIterator{contract: _BurnMintERC20UUPS.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSCCIPAdminTransferred)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20UUPSCCIPAdminTransferred, error) {
	event := new(BurnMintERC20UUPSCCIPAdminTransferred)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSDefaultAdminDelayChangeCanceledIterator struct {
	Event *BurnMintERC20UUPSDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BurnMintERC20UUPSDefaultAdminDelayChangeCanceled)
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

func (it *BurnMintERC20UUPSDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20UUPSDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSDefaultAdminDelayChangeCanceledIterator{contract: _BurnMintERC20UUPS.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSDefaultAdminDelayChangeCanceled)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20UUPSDefaultAdminDelayChangeCanceled, error) {
	event := new(BurnMintERC20UUPSDefaultAdminDelayChangeCanceled)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSDefaultAdminDelayChangeScheduledIterator struct {
	Event *BurnMintERC20UUPSDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BurnMintERC20UUPSDefaultAdminDelayChangeScheduled)
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

func (it *BurnMintERC20UUPSDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20UUPSDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSDefaultAdminDelayChangeScheduledIterator{contract: _BurnMintERC20UUPS.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSDefaultAdminDelayChangeScheduled)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20UUPSDefaultAdminDelayChangeScheduled, error) {
	event := new(BurnMintERC20UUPSDefaultAdminDelayChangeScheduled)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSDefaultAdminTransferCanceledIterator struct {
	Event *BurnMintERC20UUPSDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSDefaultAdminTransferCanceled)
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
		it.Event = new(BurnMintERC20UUPSDefaultAdminTransferCanceled)
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

func (it *BurnMintERC20UUPSDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20UUPSDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSDefaultAdminTransferCanceledIterator{contract: _BurnMintERC20UUPS.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSDefaultAdminTransferCanceled)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20UUPSDefaultAdminTransferCanceled, error) {
	event := new(BurnMintERC20UUPSDefaultAdminTransferCanceled)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSDefaultAdminTransferScheduledIterator struct {
	Event *BurnMintERC20UUPSDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSDefaultAdminTransferScheduled)
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
		it.Event = new(BurnMintERC20UUPSDefaultAdminTransferScheduled)
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

func (it *BurnMintERC20UUPSDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20UUPSDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSDefaultAdminTransferScheduledIterator{contract: _BurnMintERC20UUPS.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSDefaultAdminTransferScheduled)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20UUPSDefaultAdminTransferScheduled, error) {
	event := new(BurnMintERC20UUPSDefaultAdminTransferScheduled)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSInitializedIterator struct {
	Event *BurnMintERC20UUPSInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSInitialized)
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
		it.Event = new(BurnMintERC20UUPSInitialized)
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

func (it *BurnMintERC20UUPSInitializedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSInitialized struct {
	Version uint64
	Raw     types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20UUPSInitializedIterator, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSInitializedIterator{contract: _BurnMintERC20UUPS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSInitialized) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSInitialized)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseInitialized(log types.Log) (*BurnMintERC20UUPSInitialized, error) {
	event := new(BurnMintERC20UUPSInitialized)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSRoleAdminChangedIterator struct {
	Event *BurnMintERC20UUPSRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSRoleAdminChanged)
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
		it.Event = new(BurnMintERC20UUPSRoleAdminChanged)
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

func (it *BurnMintERC20UUPSRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20UUPSRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSRoleAdminChangedIterator{contract: _BurnMintERC20UUPS.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSRoleAdminChanged)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20UUPSRoleAdminChanged, error) {
	event := new(BurnMintERC20UUPSRoleAdminChanged)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSRoleGrantedIterator struct {
	Event *BurnMintERC20UUPSRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSRoleGranted)
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
		it.Event = new(BurnMintERC20UUPSRoleGranted)
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

func (it *BurnMintERC20UUPSRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20UUPSRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSRoleGrantedIterator{contract: _BurnMintERC20UUPS.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSRoleGranted)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseRoleGranted(log types.Log) (*BurnMintERC20UUPSRoleGranted, error) {
	event := new(BurnMintERC20UUPSRoleGranted)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSRoleRevokedIterator struct {
	Event *BurnMintERC20UUPSRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSRoleRevoked)
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
		it.Event = new(BurnMintERC20UUPSRoleRevoked)
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

func (it *BurnMintERC20UUPSRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20UUPSRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSRoleRevokedIterator{contract: _BurnMintERC20UUPS.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSRoleRevoked)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20UUPSRoleRevoked, error) {
	event := new(BurnMintERC20UUPSRoleRevoked)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSTransferIterator struct {
	Event *BurnMintERC20UUPSTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSTransfer)
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
		it.Event = new(BurnMintERC20UUPSTransfer)
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

func (it *BurnMintERC20UUPSTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20UUPSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSTransferIterator{contract: _BurnMintERC20UUPS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSTransfer)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseTransfer(log types.Log) (*BurnMintERC20UUPSTransfer, error) {
	event := new(BurnMintERC20UUPSTransfer)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20UUPSUpgradedIterator struct {
	Event *BurnMintERC20UUPSUpgraded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20UUPSUpgradedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20UUPSUpgraded)
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
		it.Event = new(BurnMintERC20UUPSUpgraded)
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

func (it *BurnMintERC20UUPSUpgradedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20UUPSUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20UUPSUpgraded struct {
	Implementation common.Address
	Raw            types.Log
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BurnMintERC20UUPSUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20UUPSUpgradedIterator{contract: _BurnMintERC20UUPS.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BurnMintERC20UUPS.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20UUPSUpgraded)
				if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

func (_BurnMintERC20UUPS *BurnMintERC20UUPSFilterer) ParseUpgraded(log types.Log) (*BurnMintERC20UUPSUpgraded, error) {
	event := new(BurnMintERC20UUPSUpgraded)
	if err := _BurnMintERC20UUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPS) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC20UUPS.abi.Events["Approval"].ID:
		return _BurnMintERC20UUPS.ParseApproval(log)
	case _BurnMintERC20UUPS.abi.Events["CCIPAdminTransferred"].ID:
		return _BurnMintERC20UUPS.ParseCCIPAdminTransferred(log)
	case _BurnMintERC20UUPS.abi.Events["DefaultAdminDelayChangeCanceled"].ID:
		return _BurnMintERC20UUPS.ParseDefaultAdminDelayChangeCanceled(log)
	case _BurnMintERC20UUPS.abi.Events["DefaultAdminDelayChangeScheduled"].ID:
		return _BurnMintERC20UUPS.ParseDefaultAdminDelayChangeScheduled(log)
	case _BurnMintERC20UUPS.abi.Events["DefaultAdminTransferCanceled"].ID:
		return _BurnMintERC20UUPS.ParseDefaultAdminTransferCanceled(log)
	case _BurnMintERC20UUPS.abi.Events["DefaultAdminTransferScheduled"].ID:
		return _BurnMintERC20UUPS.ParseDefaultAdminTransferScheduled(log)
	case _BurnMintERC20UUPS.abi.Events["Initialized"].ID:
		return _BurnMintERC20UUPS.ParseInitialized(log)
	case _BurnMintERC20UUPS.abi.Events["RoleAdminChanged"].ID:
		return _BurnMintERC20UUPS.ParseRoleAdminChanged(log)
	case _BurnMintERC20UUPS.abi.Events["RoleGranted"].ID:
		return _BurnMintERC20UUPS.ParseRoleGranted(log)
	case _BurnMintERC20UUPS.abi.Events["RoleRevoked"].ID:
		return _BurnMintERC20UUPS.ParseRoleRevoked(log)
	case _BurnMintERC20UUPS.abi.Events["Transfer"].ID:
		return _BurnMintERC20UUPS.ParseTransfer(log)
	case _BurnMintERC20UUPS.abi.Events["Upgraded"].ID:
		return _BurnMintERC20UUPS.ParseUpgraded(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC20UUPSApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20UUPSCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20UUPSDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (BurnMintERC20UUPSDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (BurnMintERC20UUPSDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (BurnMintERC20UUPSDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (BurnMintERC20UUPSInitialized) Topic() common.Hash {
	return common.HexToHash("0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2")
}

func (BurnMintERC20UUPSRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20UUPSRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20UUPSRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20UUPSTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (BurnMintERC20UUPSUpgraded) Topic() common.Hash {
	return common.HexToHash("0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b")
}

func (_BurnMintERC20UUPS *BurnMintERC20UUPS) Address() common.Address {
	return _BurnMintERC20UUPS.address
}

type BurnMintERC20UUPSInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

	UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error)

	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

		error)

	PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

		error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20UUPSApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20UUPSApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20UUPSCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20UUPSCCIPAdminTransferred, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20UUPSDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20UUPSDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20UUPSDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20UUPSDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20UUPSDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20UUPSDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20UUPSDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20UUPSDefaultAdminTransferScheduled, error)

	FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20UUPSInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*BurnMintERC20UUPSInitialized, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20UUPSRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20UUPSRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20UUPSRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20UUPSRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20UUPSRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20UUPSRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20UUPSTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20UUPSTransfer, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BurnMintERC20UUPSUpgradedIterator, error)

	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BurnMintERC20UUPSUpgraded, implementation []common.Address) (event.Subscription, error)

	ParseUpgraded(log types.Log) (*BurnMintERC20UUPSUpgraded, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
