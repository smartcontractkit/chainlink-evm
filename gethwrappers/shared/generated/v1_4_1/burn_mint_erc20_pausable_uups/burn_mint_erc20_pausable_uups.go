// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_pausable_uups

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

var BurnMintERC20PausableUUPSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultUpgrader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20UUPS__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20UUPS__MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161392a6100fd60003960008181611cc201528181611ceb0152611ef3015261392a6000f3fe6080604052600436106102fd5760003560e01c80638456cb591161018f578063ad3cb1cc116100e1578063d547741f1161008a578063dd62ed3e11610064578063dd62ed3e14610a4e578063e63ab1e914610ac0578063f72c0d8b14610af457600080fd5b8063d547741f146109e5578063d5abeb0114610a05578063d602b9fd14610a3957600080fd5b8063cefc1429116100bb578063cefc142914610924578063cf6eefb714610939578063d5391393146109b157600080fd5b8063ad3cb1cc146108a6578063c630948d146108ef578063cc8463c81461090f57600080fd5b806395d89b4111610143578063a217fddf1161011d578063a217fddf14610851578063a8fa343c14610866578063a9059cbb1461088657600080fd5b806395d89b41146107e85780639dc29fac146107fd578063a1eda53c1461081d57600080fd5b80638da5cb5b116101745780638da5cb5b1461074c5780638fd6a6ac1461076157806391d148541461077657600080fd5b80638456cb59146106cc57806384ef8ffc146106e157600080fd5b806336568abe11610253578063561cf2ab116101fc578063649a5ec7116101d6578063649a5ec71461062a57806370a082311461064a57806379cc6790146106ac57600080fd5b8063561cf2ab146105b35780635c975abb146105d3578063634e93da1461060a57600080fd5b806342966c681161022d57806342966c681461056b5780634f1ef2861461058b57806352d1902d1461059e57600080fd5b806336568abe146105165780633f4ba83a1461053657806340c10f191461054b57600080fd5b806318160ddd116102b5578063282c51f31161028f578063282c51f3146104665780632f2ff15d1461049a578063313ce567146104ba57600080fd5b806318160ddd146103b957806323b872dd146103f7578063248a9ca31461041757600080fd5b806306fdde03116102e657806306fdde0314610360578063095ea7b3146103825780630aa6220b146103a257600080fd5b806301ffc9a714610302578063022d63fb14610337575b600080fd5b34801561030e57600080fd5b5061032261031d3660046132bf565b610b28565b60405190151581526020015b60405180910390f35b34801561034357600080fd5b50620697805b60405165ffffffffffff909116815260200161032e565b34801561036c57600080fd5b50610375610cf1565b60405161032e9190613325565b34801561038e57600080fd5b5061032261039d36600461339f565b610dc6565b3480156103ae57600080fd5b506103b7610dde565b005b3480156103c557600080fd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161032e565b34801561040357600080fd5b506103226104123660046133c9565b610df4565b34801561042357600080fd5b506103e9610432366004613406565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561047257600080fd5b506103e97f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b3480156104a657600080fd5b506103b76104b536600461341f565b610e1a565b3480156104c657600080fd5b507fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff005474010000000000000000000000000000000000000000900460ff1660405160ff909116815260200161032e565b34801561052257600080fd5b506103b761053136600461341f565b610e5f565b34801561054257600080fd5b506103b7610fcb565b34801561055757600080fd5b506103b761056636600461339f565b611015565b34801561057757600080fd5b506103b7610586366004613406565b6110f9565b6103b7610599366004613512565b61112c565b3480156105aa57600080fd5b506103e9611147565b3480156105bf57600080fd5b506103b76105ce366004613594565b611176565b3480156105df57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610322565b34801561061657600080fd5b506103b761062536600461364a565b611442565b34801561063657600080fd5b506103b7610645366004613665565b611456565b34801561065657600080fd5b506103e961066536600461364a565b73ffffffffffffffffffffffffffffffffffffffff1660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b3480156106b857600080fd5b506103b76106c736600461339f565b61146a565b3480156106d857600080fd5b506103b761149e565b3480156106ed57600080fd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161032e565b34801561075857600080fd5b50610727611500565b34801561076d57600080fd5b50610727611545565b34801561078257600080fd5b5061032261079136600461341f565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156107f457600080fd5b50610375611585565b34801561080957600080fd5b506103b761081836600461339f565b6115d6565b34801561082957600080fd5b506108326115e0565b6040805165ffffffffffff93841681529290911660208301520161032e565b34801561085d57600080fd5b506103e9600081565b34801561087257600080fd5b506103b761088136600461364a565b61169f565b34801561089257600080fd5b506103226108a136600461339f565b611741565b3480156108b257600080fd5b506103756040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156108fb57600080fd5b506103b761090a36600461364a565b61174f565b34801561091b57600080fd5b506103496117a3565b34801561093057600080fd5b506103b7611884565b34801561094557600080fd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161032e565b3480156109bd57600080fd5b506103e97f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b3480156109f157600080fd5b506103b7610a0036600461341f565b6118ff565b348015610a1157600080fd5b507fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff01546103e9565b348015610a4557600080fd5b506103b7611940565b348015610a5a57600080fd5b506103e9610a6936600461368d565b73ffffffffffffffffffffffffffffffffffffffff91821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610acc57600080fd5b506103e97f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b348015610b0057600080fd5b506103e97f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b07000000000000000000000000000000000000000000000000000000001480610bbb57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b80610c0757507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b80610c5357507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b80610c9f57507fffffffff0000000000000000000000000000000000000000000000000000000082167f52d1902d00000000000000000000000000000000000000000000000000000000145b80610ceb57507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac00000000000000000000000000000000000000000000000000000000145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610d42906136b7565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6e906136b7565b8015610dbb5780601f10610d9057610100808354040283529160200191610dbb565b820191906000526020600020905b815481529060010190602001808311610d9e57829003601f168201915b505050505091505090565b600033610dd4818585611953565b5060019392505050565b6000610de981611960565b610df161196a565b50565b600033610e02858285611977565b610e0d858585611a65565b60019150505b9392505050565b81610e51576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e5b8282611b10565b5050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840082158015610ec757507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8381169116145b15610fbc577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610f3a575065ffffffffffff8116155b80610f4d57504265ffffffffffff821610155b15610f93576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b610fc68383611b54565b505050565b6000610fd681611960565b610fde611bad565b6040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a150565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a661103f81611960565b7fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff0154600061108b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b905081158015906110a45750816110a28583613739565b115b156110e8576110b38482613739565b6040517f193e2453000000000000000000000000000000000000000000000000000000008152600401610f8a91815260200190565b6110f28585611c44565b5050505050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861112381611960565b610e5b82611ca0565b611134611caa565b61113d82611d78565b610e5b8282611da2565b6000611151611edb565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156111c15750825b905060008267ffffffffffffffff1660011480156111de5750303b155b9050811580156111ec575080155b15611223576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156112845784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b61128e8c8c611f4a565b611296611f5c565b61129e611f5c565b6112a6611f5c565b60007fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff008054600182018c90557fffffffffffffffffffffff000000000000000000000000000000000000000000167401000000000000000000000000000000000000000060ff8e16027fffffffffffffffffffffffff0000000000000000000000000000000000000000161773ffffffffffffffffffffffffffffffffffffffff8a161781559050881561139b5789891115611391576040517f193e2453000000000000000000000000000000000000000000000000000000008152600481018a9052602401610f8a565b61139b888a611c44565b6113a6600089611f64565b506113d17f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e388611f64565b505083156114345784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b600061144d81611960565b610e5b8261206e565b600061146181611960565b610e5b826120ee565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861149481611960565b610fc6838361215e565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6114c881611960565b6114d0612173565b6040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589060200161100a565b60006115407feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b60007fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff005b5473ffffffffffffffffffffffffffffffffffffffff16919050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610d42906136b7565b610e5b828261146a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400811580159061166357504265ffffffffffff831610155b61166f57600080611696565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b60006116aa81611960565b7fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff0080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff858116918217845560405192169182907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a350505050565b600033610dd4818585611a65565b6117797f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610e1a565b610df17f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610e1a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801580159061182657504265ffffffffffff8216105b6118575781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff1661187d565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff163381146118f7576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610f8a565b610df16121ec565b81611936576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e5b828261231d565b600061194b81611960565b610df1612361565b610fc6838383600161236c565b610df18133612380565b611975600080612427565b565b73ffffffffffffffffffffffffffffffffffffffff83811660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611a5f5781811015611a50576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610f8a565b611a5f8484848403600061236c565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611ab5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610f8a565b73ffffffffffffffffffffffffffffffffffffffff8216611b05576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610f8a565b610fc68383836125c0565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611b4a81611960565b611a5f8383611f64565b73ffffffffffffffffffffffffffffffffffffffff81163314611ba3576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fc682826125d3565b611bb5612677565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161100a565b73ffffffffffffffffffffffffffffffffffffffff8216611c94576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610f8a565b610e5b600083836125c0565b610df133826126d2565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480611d4157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611d2861272e565b73ffffffffffffffffffffffffffffffffffffffff1614155b15611975576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610e5b81611960565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611e27575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611e249181019061374c565b60015b611e75576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610f8a565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611ed1576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610f8a565b610fc68383612756565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611975576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f526127b9565b610e5b8282612820565b6119756127b9565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008361205c576000611fcc7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612019576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85161790555b6120668484612883565b949350505050565b60006120786117a3565b612081426129a4565b61208b9190613765565b905061209782826129f4565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b60006120f982612aaf565b612102426129a4565b61210c9190613765565b90506121188282612427565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b612169823383611977565b610e5b82826126d2565b61217b612af7565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611c1f565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400805473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061225c57504265ffffffffffff821610155b1561229d576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610f8a565b6122e560006122e07feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b6125d3565b506122f1600083611f64565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461235781611960565b611a5f83836125d3565b6119756000806129f4565b612374612af7565b611a5f84848484612b53565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610e5b576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610f8a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801561253a574265ffffffffffff82161015612510576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a0100000000000000000000000000000000000000000000000000000217825561253a565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600101805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b6125c8612af7565b610fc6838383612bc6565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561263d57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8481169116145b1561266d576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6120668484612c38565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611975576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216612722576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610f8a565b610e5b826000836125c0565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc611569565b61275f82612d16565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156127b157610fc68282612de5565b610e5b612e68565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611975576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128286127b9565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0361287484826137ca565b5060048101611a5f83826137ca565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1661299a5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556129363390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610ceb565b6000915050610ceb565b600065ffffffffffff8211156129f0576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610f8a565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff8816171784559104168015611a5f576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a150505050565b600080612aba6117a3565b90508065ffffffffffff168365ffffffffffff1611612ae257612add83826138e3565b610e13565b610e1365ffffffffffff841662069780612ea0565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611975576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff841603612bba576040517f99817ca200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610f8a565b611a5f84848484612eb6565b3073ffffffffffffffffffffffffffffffffffffffff831603612c2d576040517f99817ca200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610f8a565b610fc6838383613022565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff161561299a5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610ceb565b8073ffffffffffffffffffffffffffffffffffffffff163b600003612d7f576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610f8a565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60606000808473ffffffffffffffffffffffffffffffffffffffff1684604051612e0f9190613901565b600060405180830381855af49150503d8060008114612e4a576040519150601f19603f3d011682016040523d82523d6000602084013e612e4f565b606091505b5091509150612e5f8583836131f3565b95945050505050565b3415611975576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000818310612eaf5781610e13565b5090919050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516612f27576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610f8a565b73ffffffffffffffffffffffffffffffffffffffff8416612f77576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610f8a565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600183016020908152604080832093881683529290522083905581156110f2578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161301391815260200190565b60405180910390a35050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff841661307d57818160020160008282546130729190613739565b9091555061312f9050565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020829052604090205482811015613103576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610f8a565b73ffffffffffffffffffffffffffffffffffffffff851660009081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff831661315a576002810180548390039055613186565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516131e591815260200190565b60405180910390a350505050565b60608261320357612add8261327d565b8151158015613227575073ffffffffffffffffffffffffffffffffffffffff84163b155b15613276576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610f8a565b5080610e13565b80511561328d5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156132d157600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610e1357600080fd5b60005b8381101561331c578181015183820152602001613304565b50506000910152565b6020815260008251806020840152613344816040850160208701613301565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461339a57600080fd5b919050565b600080604083850312156133b257600080fd5b6133bb83613376565b946020939093013593505050565b6000806000606084860312156133de57600080fd5b6133e784613376565b92506133f560208501613376565b929592945050506040919091013590565b60006020828403121561341857600080fd5b5035919050565b6000806040838503121561343257600080fd5b8235915061344260208401613376565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156134955761349561344b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156134e2576134e261344b565b6040528381529050808284018510156134fa57600080fd5b83836020830137600060208583010152509392505050565b6000806040838503121561352557600080fd5b61352e83613376565b9150602083013567ffffffffffffffff81111561354a57600080fd5b8301601f8101851361355b57600080fd5b61356a8582356020840161347a565b9150509250929050565b600082601f83011261358557600080fd5b610e138383356020850161347a565b600080600080600080600060e0888a0312156135af57600080fd5b873567ffffffffffffffff8111156135c657600080fd5b6135d28a828b01613574565b975050602088013567ffffffffffffffff8111156135ef57600080fd5b6135fb8a828b01613574565b965050604088013560ff8116811461361257600080fd5b9450606088013593506080880135925061362e60a08901613376565b915061363c60c08901613376565b905092959891949750929550565b60006020828403121561365c57600080fd5b610e1382613376565b60006020828403121561367757600080fd5b813565ffffffffffff81168114610e1357600080fd5b600080604083850312156136a057600080fd5b6136a983613376565b915061344260208401613376565b600181811c908216806136cb57607f821691505b602082108103613704577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610ceb57610ceb61370a565b60006020828403121561375e57600080fd5b5051919050565b65ffffffffffff8181168382160190811115610ceb57610ceb61370a565b601f821115610fc657806000526020600020601f840160051c810160208510156137aa5750805b601f840160051c820191505b818110156110f257600081556001016137b6565b815167ffffffffffffffff8111156137e4576137e461344b565b6137f8816137f284546136b7565b84613783565b6020601f82116001811461384a57600083156138145750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556110f2565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156138985787850151825560209485019460019092019101613878565b50848210156138d457868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b65ffffffffffff8281168282160390811115610ceb57610ceb61370a565b60008251613913818460208701613301565b919091019291505056fea164736f6c634300081a000a",
}

var BurnMintERC20PausableUUPSABI = BurnMintERC20PausableUUPSMetaData.ABI

var BurnMintERC20PausableUUPSBin = BurnMintERC20PausableUUPSMetaData.Bin

func DeployBurnMintERC20PausableUUPS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnMintERC20PausableUUPS, error) {
	parsed, err := BurnMintERC20PausableUUPSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20PausableUUPSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20PausableUUPS{address: address, abi: *parsed, BurnMintERC20PausableUUPSCaller: BurnMintERC20PausableUUPSCaller{contract: contract}, BurnMintERC20PausableUUPSTransactor: BurnMintERC20PausableUUPSTransactor{contract: contract}, BurnMintERC20PausableUUPSFilterer: BurnMintERC20PausableUUPSFilterer{contract: contract}}, nil
}

type BurnMintERC20PausableUUPS struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20PausableUUPSCaller
	BurnMintERC20PausableUUPSTransactor
	BurnMintERC20PausableUUPSFilterer
}

type BurnMintERC20PausableUUPSCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableUUPSTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableUUPSFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableUUPSSession struct {
	Contract     *BurnMintERC20PausableUUPS
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableUUPSCallerSession struct {
	Contract *BurnMintERC20PausableUUPSCaller
	CallOpts bind.CallOpts
}

type BurnMintERC20PausableUUPSTransactorSession struct {
	Contract     *BurnMintERC20PausableUUPSTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableUUPSRaw struct {
	Contract *BurnMintERC20PausableUUPS
}

type BurnMintERC20PausableUUPSCallerRaw struct {
	Contract *BurnMintERC20PausableUUPSCaller
}

type BurnMintERC20PausableUUPSTransactorRaw struct {
	Contract *BurnMintERC20PausableUUPSTransactor
}

func NewBurnMintERC20PausableUUPS(address common.Address, backend bind.ContractBackend) (*BurnMintERC20PausableUUPS, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20PausableUUPSABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20PausableUUPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPS{address: address, abi: abi, BurnMintERC20PausableUUPSCaller: BurnMintERC20PausableUUPSCaller{contract: contract}, BurnMintERC20PausableUUPSTransactor: BurnMintERC20PausableUUPSTransactor{contract: contract}, BurnMintERC20PausableUUPSFilterer: BurnMintERC20PausableUUPSFilterer{contract: contract}}, nil
}

func NewBurnMintERC20PausableUUPSCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20PausableUUPSCaller, error) {
	contract, err := bindBurnMintERC20PausableUUPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSCaller{contract: contract}, nil
}

func NewBurnMintERC20PausableUUPSTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20PausableUUPSTransactor, error) {
	contract, err := bindBurnMintERC20PausableUUPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSTransactor{contract: contract}, nil
}

func NewBurnMintERC20PausableUUPSFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20PausableUUPSFilterer, error) {
	contract, err := bindBurnMintERC20PausableUUPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSFilterer{contract: contract}, nil
}

func bindBurnMintERC20PausableUUPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20PausableUUPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableUUPS.Contract.BurnMintERC20PausableUUPSCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.BurnMintERC20PausableUUPSTransactor.contract.Transfer(opts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.BurnMintERC20PausableUUPSTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableUUPS.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.BURNERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.BURNERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.MINTERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.MINTERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.PAUSERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.PAUSERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) UPGRADERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.UPGRADERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.UPGRADERROLE(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BurnMintERC20PausableUUPS.Contract.UPGRADEINTERFACEVERSION(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BurnMintERC20PausableUUPS.Contract.UPGRADEINTERFACEVERSION(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.Allowance(&_BurnMintERC20PausableUUPS.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.Allowance(&_BurnMintERC20PausableUUPS.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.BalanceOf(&_BurnMintERC20PausableUUPS.CallOpts, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.BalanceOf(&_BurnMintERC20PausableUUPS.CallOpts, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableUUPS.Contract.Decimals(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableUUPS.Contract.Decimals(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableUUPS.Contract.DefaultAdmin(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableUUPS.Contract.DefaultAdmin(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.DefaultAdminDelay(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.DefaultAdminDelay(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableUUPS.Contract.GetCCIPAdmin(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableUUPS.Contract.GetCCIPAdmin(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.GetRoleAdmin(&_BurnMintERC20PausableUUPS.CallOpts, role)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.GetRoleAdmin(&_BurnMintERC20PausableUUPS.CallOpts, role)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableUUPS.Contract.HasRole(&_BurnMintERC20PausableUUPS.CallOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableUUPS.Contract.HasRole(&_BurnMintERC20PausableUUPS.CallOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.MaxSupply(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.MaxSupply(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Name() (string, error) {
	return _BurnMintERC20PausableUUPS.Contract.Name(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) Name() (string, error) {
	return _BurnMintERC20PausableUUPS.Contract.Name(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableUUPS.Contract.Owner(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableUUPS.Contract.Owner(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Paused() (bool, error) {
	return _BurnMintERC20PausableUUPS.Contract.Paused(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) Paused() (bool, error) {
	return _BurnMintERC20PausableUUPS.Contract.Paused(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableUUPS.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableUUPS.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableUUPS.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableUUPS.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) ProxiableUUID() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.ProxiableUUID(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BurnMintERC20PausableUUPS.Contract.ProxiableUUID(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableUUPS.Contract.SupportsInterface(&_BurnMintERC20PausableUUPS.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableUUPS.Contract.SupportsInterface(&_BurnMintERC20PausableUUPS.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Symbol() (string, error) {
	return _BurnMintERC20PausableUUPS.Contract.Symbol(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) Symbol() (string, error) {
	return _BurnMintERC20PausableUUPS.Contract.Symbol(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableUUPS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.TotalSupply(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableUUPS.Contract.TotalSupply(&_BurnMintERC20PausableUUPS.CallOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Approve(&_BurnMintERC20PausableUUPS.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Approve(&_BurnMintERC20PausableUUPS.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Burn(&_BurnMintERC20PausableUUPS.TransactOpts, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Burn(&_BurnMintERC20PausableUUPS.TransactOpts, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Burn0(&_BurnMintERC20PausableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Burn0(&_BurnMintERC20PausableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.BurnFrom(&_BurnMintERC20PausableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.BurnFrom(&_BurnMintERC20PausableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableUUPS.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableUUPS.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableUUPS.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableUUPS.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.GrantRole(&_BurnMintERC20PausableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.GrantRole(&_BurnMintERC20PausableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "initialize", name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Initialize(&_BurnMintERC20PausableUUPS.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Initialize(&_BurnMintERC20PausableUUPS.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Mint(&_BurnMintERC20PausableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Mint(&_BurnMintERC20PausableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "pause")
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Pause(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Pause(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.RenounceRole(&_BurnMintERC20PausableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.RenounceRole(&_BurnMintERC20PausableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.RevokeRole(&_BurnMintERC20PausableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.RevokeRole(&_BurnMintERC20PausableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.SetCCIPAdmin(&_BurnMintERC20PausableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.SetCCIPAdmin(&_BurnMintERC20PausableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Transfer(&_BurnMintERC20PausableUUPS.TransactOpts, to, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Transfer(&_BurnMintERC20PausableUUPS.TransactOpts, to, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.TransferFrom(&_BurnMintERC20PausableUUPS.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.TransferFrom(&_BurnMintERC20PausableUUPS.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "unpause")
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Unpause(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.Unpause(&_BurnMintERC20PausableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.UpgradeToAndCall(&_BurnMintERC20PausableUUPS.TransactOpts, newImplementation, data)
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20PausableUUPS.Contract.UpgradeToAndCall(&_BurnMintERC20PausableUUPS.TransactOpts, newImplementation, data)
}

type BurnMintERC20PausableUUPSApprovalIterator struct {
	Event *BurnMintERC20PausableUUPSApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSApproval)
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
		it.Event = new(BurnMintERC20PausableUUPSApproval)
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

func (it *BurnMintERC20PausableUUPSApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableUUPSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSApprovalIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSApproval)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseApproval(log types.Log) (*BurnMintERC20PausableUUPSApproval, error) {
	event := new(BurnMintERC20PausableUUPSApproval)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSCCIPAdminTransferredIterator struct {
	Event *BurnMintERC20PausableUUPSCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSCCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20PausableUUPSCCIPAdminTransferred)
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

func (it *BurnMintERC20PausableUUPSCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableUUPSCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSCCIPAdminTransferredIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSCCIPAdminTransferred)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableUUPSCCIPAdminTransferred, error) {
	event := new(BurnMintERC20PausableUUPSCCIPAdminTransferred)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceledIterator struct {
	Event *BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled)
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

func (it *BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceledIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled, error) {
	event := new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduledIterator struct {
	Event *BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled)
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

func (it *BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduledIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled, error) {
	event := new(BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSDefaultAdminTransferCanceledIterator struct {
	Event *BurnMintERC20PausableUUPSDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSDefaultAdminTransferCanceled)
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
		it.Event = new(BurnMintERC20PausableUUPSDefaultAdminTransferCanceled)
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

func (it *BurnMintERC20PausableUUPSDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSDefaultAdminTransferCanceledIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSDefaultAdminTransferCanceled)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminTransferCanceled, error) {
	event := new(BurnMintERC20PausableUUPSDefaultAdminTransferCanceled)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSDefaultAdminTransferScheduledIterator struct {
	Event *BurnMintERC20PausableUUPSDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSDefaultAdminTransferScheduled)
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
		it.Event = new(BurnMintERC20PausableUUPSDefaultAdminTransferScheduled)
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

func (it *BurnMintERC20PausableUUPSDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableUUPSDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSDefaultAdminTransferScheduledIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSDefaultAdminTransferScheduled)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminTransferScheduled, error) {
	event := new(BurnMintERC20PausableUUPSDefaultAdminTransferScheduled)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSInitializedIterator struct {
	Event *BurnMintERC20PausableUUPSInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSInitialized)
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
		it.Event = new(BurnMintERC20PausableUUPSInitialized)
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

func (it *BurnMintERC20PausableUUPSInitializedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSInitialized struct {
	Version uint64
	Raw     types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSInitializedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSInitializedIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSInitialized) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSInitialized)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseInitialized(log types.Log) (*BurnMintERC20PausableUUPSInitialized, error) {
	event := new(BurnMintERC20PausableUUPSInitialized)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSPausedIterator struct {
	Event *BurnMintERC20PausableUUPSPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSPaused)
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
		it.Event = new(BurnMintERC20PausableUUPSPaused)
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

func (it *BurnMintERC20PausableUUPSPausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSPausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSPausedIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSPaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSPaused)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParsePaused(log types.Log) (*BurnMintERC20PausableUUPSPaused, error) {
	event := new(BurnMintERC20PausableUUPSPaused)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSRoleAdminChangedIterator struct {
	Event *BurnMintERC20PausableUUPSRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSRoleAdminChanged)
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
		it.Event = new(BurnMintERC20PausableUUPSRoleAdminChanged)
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

func (it *BurnMintERC20PausableUUPSRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableUUPSRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSRoleAdminChangedIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSRoleAdminChanged)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableUUPSRoleAdminChanged, error) {
	event := new(BurnMintERC20PausableUUPSRoleAdminChanged)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSRoleGrantedIterator struct {
	Event *BurnMintERC20PausableUUPSRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSRoleGranted)
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
		it.Event = new(BurnMintERC20PausableUUPSRoleGranted)
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

func (it *BurnMintERC20PausableUUPSRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableUUPSRoleGrantedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSRoleGrantedIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSRoleGranted)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseRoleGranted(log types.Log) (*BurnMintERC20PausableUUPSRoleGranted, error) {
	event := new(BurnMintERC20PausableUUPSRoleGranted)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSRoleRevokedIterator struct {
	Event *BurnMintERC20PausableUUPSRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSRoleRevoked)
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
		it.Event = new(BurnMintERC20PausableUUPSRoleRevoked)
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

func (it *BurnMintERC20PausableUUPSRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableUUPSRoleRevokedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSRoleRevokedIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSRoleRevoked)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableUUPSRoleRevoked, error) {
	event := new(BurnMintERC20PausableUUPSRoleRevoked)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSTransferIterator struct {
	Event *BurnMintERC20PausableUUPSTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSTransfer)
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
		it.Event = new(BurnMintERC20PausableUUPSTransfer)
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

func (it *BurnMintERC20PausableUUPSTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableUUPSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSTransferIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSTransfer)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseTransfer(log types.Log) (*BurnMintERC20PausableUUPSTransfer, error) {
	event := new(BurnMintERC20PausableUUPSTransfer)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSUnpausedIterator struct {
	Event *BurnMintERC20PausableUUPSUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSUnpaused)
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
		it.Event = new(BurnMintERC20PausableUUPSUnpaused)
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

func (it *BurnMintERC20PausableUUPSUnpausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSUnpausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSUnpausedIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSUnpaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSUnpaused)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseUnpaused(log types.Log) (*BurnMintERC20PausableUUPSUnpaused, error) {
	event := new(BurnMintERC20PausableUUPSUnpaused)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableUUPSUpgradedIterator struct {
	Event *BurnMintERC20PausableUUPSUpgraded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableUUPSUpgradedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableUUPSUpgraded)
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
		it.Event = new(BurnMintERC20PausableUUPSUpgraded)
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

func (it *BurnMintERC20PausableUUPSUpgradedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableUUPSUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableUUPSUpgraded struct {
	Implementation common.Address
	Raw            types.Log
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BurnMintERC20PausableUUPSUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableUUPSUpgradedIterator{contract: _BurnMintERC20PausableUUPS.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BurnMintERC20PausableUUPS.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableUUPSUpgraded)
				if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPSFilterer) ParseUpgraded(log types.Log) (*BurnMintERC20PausableUUPSUpgraded, error) {
	event := new(BurnMintERC20PausableUUPSUpgraded)
	if err := _BurnMintERC20PausableUUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPS) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC20PausableUUPS.abi.Events["Approval"].ID:
		return _BurnMintERC20PausableUUPS.ParseApproval(log)
	case _BurnMintERC20PausableUUPS.abi.Events["CCIPAdminTransferred"].ID:
		return _BurnMintERC20PausableUUPS.ParseCCIPAdminTransferred(log)
	case _BurnMintERC20PausableUUPS.abi.Events["DefaultAdminDelayChangeCanceled"].ID:
		return _BurnMintERC20PausableUUPS.ParseDefaultAdminDelayChangeCanceled(log)
	case _BurnMintERC20PausableUUPS.abi.Events["DefaultAdminDelayChangeScheduled"].ID:
		return _BurnMintERC20PausableUUPS.ParseDefaultAdminDelayChangeScheduled(log)
	case _BurnMintERC20PausableUUPS.abi.Events["DefaultAdminTransferCanceled"].ID:
		return _BurnMintERC20PausableUUPS.ParseDefaultAdminTransferCanceled(log)
	case _BurnMintERC20PausableUUPS.abi.Events["DefaultAdminTransferScheduled"].ID:
		return _BurnMintERC20PausableUUPS.ParseDefaultAdminTransferScheduled(log)
	case _BurnMintERC20PausableUUPS.abi.Events["Initialized"].ID:
		return _BurnMintERC20PausableUUPS.ParseInitialized(log)
	case _BurnMintERC20PausableUUPS.abi.Events["Paused"].ID:
		return _BurnMintERC20PausableUUPS.ParsePaused(log)
	case _BurnMintERC20PausableUUPS.abi.Events["RoleAdminChanged"].ID:
		return _BurnMintERC20PausableUUPS.ParseRoleAdminChanged(log)
	case _BurnMintERC20PausableUUPS.abi.Events["RoleGranted"].ID:
		return _BurnMintERC20PausableUUPS.ParseRoleGranted(log)
	case _BurnMintERC20PausableUUPS.abi.Events["RoleRevoked"].ID:
		return _BurnMintERC20PausableUUPS.ParseRoleRevoked(log)
	case _BurnMintERC20PausableUUPS.abi.Events["Transfer"].ID:
		return _BurnMintERC20PausableUUPS.ParseTransfer(log)
	case _BurnMintERC20PausableUUPS.abi.Events["Unpaused"].ID:
		return _BurnMintERC20PausableUUPS.ParseUnpaused(log)
	case _BurnMintERC20PausableUUPS.abi.Events["Upgraded"].ID:
		return _BurnMintERC20PausableUUPS.ParseUpgraded(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC20PausableUUPSApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20PausableUUPSCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (BurnMintERC20PausableUUPSDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (BurnMintERC20PausableUUPSDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (BurnMintERC20PausableUUPSInitialized) Topic() common.Hash {
	return common.HexToHash("0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2")
}

func (BurnMintERC20PausableUUPSPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (BurnMintERC20PausableUUPSRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20PausableUUPSRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20PausableUUPSRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20PausableUUPSTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (BurnMintERC20PausableUUPSUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (BurnMintERC20PausableUUPSUpgraded) Topic() common.Hash {
	return common.HexToHash("0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b")
}

func (_BurnMintERC20PausableUUPS *BurnMintERC20PausableUUPS) Address() common.Address {
	return _BurnMintERC20PausableUUPS.address
}

type BurnMintERC20PausableUUPSInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

	PAUSERROLE(opts *bind.CallOpts) ([32]byte, error)

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

	Paused(opts *bind.CallOpts) (bool, error)

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

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableUUPSApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20PausableUUPSApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableUUPSCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableUUPSCCIPAdminTransferred, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableUUPSDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableUUPSDefaultAdminTransferScheduled, error)

	FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*BurnMintERC20PausableUUPSInitialized, error)

	FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*BurnMintERC20PausableUUPSPaused, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableUUPSRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableUUPSRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableUUPSRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20PausableUUPSRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableUUPSRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableUUPSRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableUUPSTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20PausableUUPSTransfer, error)

	FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableUUPSUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*BurnMintERC20PausableUUPSUnpaused, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BurnMintERC20PausableUUPSUpgradedIterator, error)

	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableUUPSUpgraded, implementation []common.Address) (event.Subscription, error)

	ParseUpgraded(log types.Log) (*BurnMintERC20PausableUUPSUpgraded, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
