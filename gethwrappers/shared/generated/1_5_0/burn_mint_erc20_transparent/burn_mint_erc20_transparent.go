// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_transparent

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

var BurnMintERC20TransparentMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20Transparent__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20Transparent__MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b612c32806100d96000396000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c806384ef8ffc11610160578063a9059cbb116100d8578063d53913931161008c578063d5abeb0111610071578063d5abeb0114610736578063d602b9fd1461075d578063dd62ed3e1461076557600080fd5b8063d5391393146106fc578063d547741f1461072357600080fd5b8063cc8463c8116100bd578063cc8463c814610681578063cefc142914610689578063cf6eefb71461069157600080fd5b8063a9059cbb1461065b578063c630948d1461066e57600080fd5b806395d89b411161012f578063a1eda53c11610114578063a1eda53c14610619578063a217fddf14610640578063a8fa343c1461064857600080fd5b806395d89b41146105fe5780639dc29fac1461060657600080fd5b806384ef8ffc146104f65780638da5cb5b146105545780638fd6a6ac1461055c57806391d148541461059957600080fd5b80632cd77a5a1161020e57806342966c68116101c2578063649a5ec7116101a7578063649a5ec71461047b57806370a082311461048e57806379cc6790146104e357600080fd5b806342966c6814610455578063634e93da1461046857600080fd5b8063313ce567116101f3578063313ce567146103e057806336568abe1461042f57806340c10f191461044257600080fd5b80632cd77a5a146103ba5780632f2ff15d146103cd57600080fd5b80630aa6220b1161026557806323b872dd1161024a57806323b872dd1461033e578063248a9ca314610351578063282c51f31461039357600080fd5b80630aa6220b1461030357806318160ddd1461030d57600080fd5b806301ffc9a714610297578063022d63fb146102bf57806306fdde03146102db578063095ea7b3146102f0575b600080fd5b6102aa6102a5366004612665565b6107ca565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016102b6565b6102e3610947565b6040516102b691906126a7565b6102aa6102fe36600461273c565b610a1c565b61030b610a34565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b6040519081526020016102b6565b6102aa61034c366004612766565b610a4a565b61033061035f3660046127a3565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6103307f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b61030b6103c83660046128b7565b610a6e565b61030b6103db36600461295b565b610d0b565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489005474010000000000000000000000000000000000000000900460ff1660405160ff90911681526020016102b6565b61030b61043d36600461295b565b610d50565b61030b61045036600461273c565b610eb7565b61030b6104633660046127a3565b610fbe565b61030b610476366004612987565b610ff1565b61030b6104893660046129a2565b611005565b61033061049c366004612987565b73ffffffffffffffffffffffffffffffffffffffff1660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b61030b6104f136600461273c565b611019565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102b6565b61052f61104d565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489005473ffffffffffffffffffffffffffffffffffffffff1661052f565b6102aa6105a736600461295b565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6102e3611092565b61030b61061436600461273c565b6110e3565b6106216110ed565b6040805165ffffffffffff9384168152929091166020830152016102b6565b610330600081565b61030b610656366004612987565b6111ac565b6102aa61066936600461273c565b61124e565b61030b61067c366004612987565b61125c565b6102c46112b0565b61030b611391565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102b6565b6103307f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b61030b61073136600461295b565b61140c565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff07904890154610330565b61030b61144d565b6103306107733660046129ca565b73ffffffffffffffffffffffffffffffffffffffff91821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000148061085d57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b806108a957507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b806108f557507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b8061094157507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac00000000000000000000000000000000000000000000000000000000145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610998906129f4565b80601f01602080910402602001604051908101604052809291908181526020018280546109c4906129f4565b8015610a115780601f106109e657610100808354040283529160200191610a11565b820191906000526020600020905b8154815290600101906020018083116109f457829003601f168201915b505050505091505090565b600033610a2a818585611460565b5060019392505050565b6000610a3f8161146d565b610a47611477565b50565b600033610a58858285611484565b610a63858585611572565b506001949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610ab95750825b905060008267ffffffffffffffff166001148015610ad65750303b155b905081158015610ae4575080155b15610b1b576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610b7c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610b868b8b61161d565b610b8e61162f565b610b9661162f565b60007fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489008054600182018b90557fffffffffffffffffffffff000000000000000000000000000000000000000000167401000000000000000000000000000000000000000060ff8d16027fffffffffffffffffffffffff0000000000000000000000000000000000000000161773ffffffffffffffffffffffffffffffffffffffff891617815590508715610c905788881115610c86576040517f25cc7152000000000000000000000000000000000000000000000000000000008152600481018990526024015b60405180910390fd5b610c908789611637565b610c9b600088611693565b50508315610cfe5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b81610d42576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d4c828261179d565b5050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840082158015610db857507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8381169116145b15610ea8577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610e2b575065ffffffffffff8116155b80610e3e57504265ffffffffffff821610155b15610e7f576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610c7d565b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b610eb283836117e1565b505050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6610ee18161146d565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff079048901547fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff079048900906000610f4f7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b90508115801590610f68575081610f668683612a76565b115b15610fac57610f778582612a76565b6040517f25cc7152000000000000000000000000000000000000000000000000000000008152600401610c7d91815260200190565b610fb68686611637565b505050505050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848610fe88161146d565b610d4c8261183a565b6000610ffc8161146d565b610d4c82611844565b60006110108161146d565b610d4c826118c4565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8486110438161146d565b610eb28383611934565b600061108d7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610998906129f4565b610d4c8282611019565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400811580159061117057504265ffffffffffff831610155b61117c576000806111a3565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b60006111b78161146d565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff07904890080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff858116918217845560405192169182907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a350505050565b600033610a2a818585611572565b6112867f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610d0b565b610a477f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610d0b565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801580159061133357504265ffffffffffff8216105b6113645781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff1661138a565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff16338114611404576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610c7d565b610a47611949565b81611443576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d4c8282611a7a565b60006114588161146d565b610a47611abe565b610eb28383836001611ac9565b610a478133611b3c565b611482600080611be3565b565b73ffffffffffffffffffffffffffffffffffffffff83811660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461156c578181101561155d576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610c7d565b61156c84848484036000611ac9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166115c2576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610c7d565b73ffffffffffffffffffffffffffffffffffffffff8216611612576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610c7d565b610eb2838383611d7c565b611625611dee565b610d4c8282611e55565b611482611dee565b73ffffffffffffffffffffffffffffffffffffffff8216611687576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610c7d565b610d4c60008383611d7c565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008361178b5760006116fb7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614611748576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85161790555b6117958484611eb8565b949350505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546117d78161146d565b61156c8383611693565b73ffffffffffffffffffffffffffffffffffffffff81163314611830576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610eb28282611fd9565b610a47338261207d565b600061184e6112b0565b611857426120d9565b6118619190612a89565b905061186d8282612129565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b60006118cf826121e4565b6118d8426120d9565b6118e29190612a89565b90506118ee8282611be3565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b61193f823383611484565b610d4c828261207d565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400805473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff168015806119b957504265ffffffffffff821610155b156119fa576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610c7d565b611a426000611a3d7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b611fd9565b50611a4e600083611693565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611ab48161146d565b61156c8383611fd9565b611482600080612129565b3073ffffffffffffffffffffffffffffffffffffffff841603611b30576040517f54e16bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610c7d565b61156c84848484612233565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610d4c576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610c7d565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611cf6574265ffffffffffff82161015611ccc576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255611cf6565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600101805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b3073ffffffffffffffffffffffffffffffffffffffff831603611de3576040517f54e16bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c7d565b610eb28383836123a0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611482576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611e5d611dee565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03611ea98482612aee565b506004810161156c8382612aee565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16611fcf5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611f6b3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610941565b6000915050610941565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561204357507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8481169116145b15612073576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6117958484612571565b73ffffffffffffffffffffffffffffffffffffffff82166120cd576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610c7d565b610d4c82600083611d7c565b600065ffffffffffff821115612125576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610c7d565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617178455910416801561156c576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a150505050565b6000806121ef6112b0565b90508065ffffffffffff168365ffffffffffff1611612217576122128382612c07565b61222c565b61222c65ffffffffffff84166206978061264f565b9392505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff85166122a4576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610c7d565b73ffffffffffffffffffffffffffffffffffffffff84166122f4576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610c7d565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260018301602090815260408083209388168352929052208390558115612399578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161239091815260200190565b60405180910390a35b5050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff84166123fb57818160020160008282546123f09190612a76565b909155506124ad9050565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020829052604090205482811015612481576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610c7d565b73ffffffffffffffffffffffffffffffffffffffff851660009081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff83166124d8576002810180548390039055612504565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161256391815260200190565b60405180910390a350505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615611fcf5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610941565b600081831061265e578161222c565b5090919050565b60006020828403121561267757600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461222c57600080fd5b602081526000825180602084015260005b818110156126d557602081860181015160408684010152016126b8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461273757600080fd5b919050565b6000806040838503121561274f57600080fd5b61275883612713565b946020939093013593505050565b60008060006060848603121561277b57600080fd5b61278484612713565b925061279260208501612713565b929592945050506040919091013590565b6000602082840312156127b557600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126127fc57600080fd5b813567ffffffffffffffff811115612816576128166127bc565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715612882576128826127bc565b60405281815283820160200185101561289a57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c087890312156128d057600080fd5b863567ffffffffffffffff8111156128e757600080fd5b6128f389828a016127eb565b965050602087013567ffffffffffffffff81111561291057600080fd5b61291c89828a016127eb565b955050604087013560ff8116811461293357600080fd5b9350606087013592506080870135915061294f60a08801612713565b90509295509295509295565b6000806040838503121561296e57600080fd5b8235915061297e60208401612713565b90509250929050565b60006020828403121561299957600080fd5b61222c82612713565b6000602082840312156129b457600080fd5b813565ffffffffffff8116811461222c57600080fd5b600080604083850312156129dd57600080fd5b6129e683612713565b915061297e60208401612713565b600181811c90821680612a0857607f821691505b602082108103612a41577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561094157610941612a47565b65ffffffffffff818116838216019081111561094157610941612a47565b601f821115610eb257806000526020600020601f840160051c81016020851015612ace5750805b601f840160051c820191505b818110156123995760008155600101612ada565b815167ffffffffffffffff811115612b0857612b086127bc565b612b1c81612b1684546129f4565b84612aa7565b6020601f821160018114612b6e5760008315612b385750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455612399565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015612bbc5787850151825560209485019460019092019101612b9c565b5084821015612bf857868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b65ffffffffffff828116828216039081111561094157610941612a4756fea164736f6c634300081a000a",
}

var BurnMintERC20TransparentABI = BurnMintERC20TransparentMetaData.ABI

var BurnMintERC20TransparentBin = BurnMintERC20TransparentMetaData.Bin

func DeployBurnMintERC20Transparent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnMintERC20Transparent, error) {
	parsed, err := BurnMintERC20TransparentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20TransparentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20Transparent{address: address, abi: *parsed, BurnMintERC20TransparentCaller: BurnMintERC20TransparentCaller{contract: contract}, BurnMintERC20TransparentTransactor: BurnMintERC20TransparentTransactor{contract: contract}, BurnMintERC20TransparentFilterer: BurnMintERC20TransparentFilterer{contract: contract}}, nil
}

type BurnMintERC20Transparent struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20TransparentCaller
	BurnMintERC20TransparentTransactor
	BurnMintERC20TransparentFilterer
}

type BurnMintERC20TransparentCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC20TransparentTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20TransparentFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20TransparentSession struct {
	Contract     *BurnMintERC20Transparent
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20TransparentCallerSession struct {
	Contract *BurnMintERC20TransparentCaller
	CallOpts bind.CallOpts
}

type BurnMintERC20TransparentTransactorSession struct {
	Contract     *BurnMintERC20TransparentTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20TransparentRaw struct {
	Contract *BurnMintERC20Transparent
}

type BurnMintERC20TransparentCallerRaw struct {
	Contract *BurnMintERC20TransparentCaller
}

type BurnMintERC20TransparentTransactorRaw struct {
	Contract *BurnMintERC20TransparentTransactor
}

func NewBurnMintERC20Transparent(address common.Address, backend bind.ContractBackend) (*BurnMintERC20Transparent, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20TransparentABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20Transparent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20Transparent{address: address, abi: abi, BurnMintERC20TransparentCaller: BurnMintERC20TransparentCaller{contract: contract}, BurnMintERC20TransparentTransactor: BurnMintERC20TransparentTransactor{contract: contract}, BurnMintERC20TransparentFilterer: BurnMintERC20TransparentFilterer{contract: contract}}, nil
}

func NewBurnMintERC20TransparentCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20TransparentCaller, error) {
	contract, err := bindBurnMintERC20Transparent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentCaller{contract: contract}, nil
}

func NewBurnMintERC20TransparentTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20TransparentTransactor, error) {
	contract, err := bindBurnMintERC20Transparent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentTransactor{contract: contract}, nil
}

func NewBurnMintERC20TransparentFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20TransparentFilterer, error) {
	contract, err := bindBurnMintERC20Transparent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentFilterer{contract: contract}, nil
}

func bindBurnMintERC20Transparent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20TransparentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20Transparent.Contract.BurnMintERC20TransparentCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.BurnMintERC20TransparentTransactor.contract.Transfer(opts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.BurnMintERC20TransparentTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20Transparent.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.BURNERROLE(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.BURNERROLE(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.DEFAULTADMINROLE(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.DEFAULTADMINROLE(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.MINTERROLE(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.MINTERROLE(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.Allowance(&_BurnMintERC20Transparent.CallOpts, owner, spender)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.Allowance(&_BurnMintERC20Transparent.CallOpts, owner, spender)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.BalanceOf(&_BurnMintERC20Transparent.CallOpts, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.BalanceOf(&_BurnMintERC20Transparent.CallOpts, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Decimals() (uint8, error) {
	return _BurnMintERC20Transparent.Contract.Decimals(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20Transparent.Contract.Decimals(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20Transparent.Contract.DefaultAdmin(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20Transparent.Contract.DefaultAdmin(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.DefaultAdminDelay(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.DefaultAdminDelay(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20Transparent.Contract.GetCCIPAdmin(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20Transparent.Contract.GetCCIPAdmin(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.GetRoleAdmin(&_BurnMintERC20Transparent.CallOpts, role)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20Transparent.Contract.GetRoleAdmin(&_BurnMintERC20Transparent.CallOpts, role)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20Transparent.Contract.HasRole(&_BurnMintERC20Transparent.CallOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20Transparent.Contract.HasRole(&_BurnMintERC20Transparent.CallOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.MaxSupply(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.MaxSupply(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Name() (string, error) {
	return _BurnMintERC20Transparent.Contract.Name(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) Name() (string, error) {
	return _BurnMintERC20Transparent.Contract.Name(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Owner() (common.Address, error) {
	return _BurnMintERC20Transparent.Contract.Owner(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC20Transparent.Contract.Owner(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20Transparent.Contract.PendingDefaultAdmin(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20Transparent.Contract.PendingDefaultAdmin(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20Transparent.Contract.PendingDefaultAdminDelay(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20Transparent.Contract.PendingDefaultAdminDelay(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20Transparent.Contract.SupportsInterface(&_BurnMintERC20Transparent.CallOpts, interfaceId)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20Transparent.Contract.SupportsInterface(&_BurnMintERC20Transparent.CallOpts, interfaceId)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Symbol() (string, error) {
	return _BurnMintERC20Transparent.Contract.Symbol(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) Symbol() (string, error) {
	return _BurnMintERC20Transparent.Contract.Symbol(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20Transparent.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.TotalSupply(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20Transparent.Contract.TotalSupply(&_BurnMintERC20Transparent.CallOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20Transparent.TransactOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20Transparent.TransactOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Approve(&_BurnMintERC20Transparent.TransactOpts, spender, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Approve(&_BurnMintERC20Transparent.TransactOpts, spender, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20Transparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20Transparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Burn(&_BurnMintERC20Transparent.TransactOpts, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Burn(&_BurnMintERC20Transparent.TransactOpts, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Burn0(&_BurnMintERC20Transparent.TransactOpts, account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Burn0(&_BurnMintERC20Transparent.TransactOpts, account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.BurnFrom(&_BurnMintERC20Transparent.TransactOpts, account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.BurnFrom(&_BurnMintERC20Transparent.TransactOpts, account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20Transparent.TransactOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20Transparent.TransactOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20Transparent.TransactOpts, newDelay)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20Transparent.TransactOpts, newDelay)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.GrantMintAndBurnRoles(&_BurnMintERC20Transparent.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.GrantMintAndBurnRoles(&_BurnMintERC20Transparent.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.GrantRole(&_BurnMintERC20Transparent.TransactOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.GrantRole(&_BurnMintERC20Transparent.TransactOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "initialize", name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Initialize(&_BurnMintERC20Transparent.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Initialize(&_BurnMintERC20Transparent.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Mint(&_BurnMintERC20Transparent.TransactOpts, account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Mint(&_BurnMintERC20Transparent.TransactOpts, account, amount)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.RenounceRole(&_BurnMintERC20Transparent.TransactOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.RenounceRole(&_BurnMintERC20Transparent.TransactOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.RevokeRole(&_BurnMintERC20Transparent.TransactOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.RevokeRole(&_BurnMintERC20Transparent.TransactOpts, role, account)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20Transparent.TransactOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20Transparent.TransactOpts)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.SetCCIPAdmin(&_BurnMintERC20Transparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.SetCCIPAdmin(&_BurnMintERC20Transparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Transfer(&_BurnMintERC20Transparent.TransactOpts, to, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.Transfer(&_BurnMintERC20Transparent.TransactOpts, to, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.TransferFrom(&_BurnMintERC20Transparent.TransactOpts, from, to, value)
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20Transparent.Contract.TransferFrom(&_BurnMintERC20Transparent.TransactOpts, from, to, value)
}

type BurnMintERC20TransparentApprovalIterator struct {
	Event *BurnMintERC20TransparentApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentApproval)
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
		it.Event = new(BurnMintERC20TransparentApproval)
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

func (it *BurnMintERC20TransparentApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20TransparentApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentApprovalIterator{contract: _BurnMintERC20Transparent.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentApproval)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseApproval(log types.Log) (*BurnMintERC20TransparentApproval, error) {
	event := new(BurnMintERC20TransparentApproval)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentCCIPAdminTransferredIterator struct {
	Event *BurnMintERC20TransparentCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentCCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20TransparentCCIPAdminTransferred)
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

func (it *BurnMintERC20TransparentCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20TransparentCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentCCIPAdminTransferredIterator{contract: _BurnMintERC20Transparent.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentCCIPAdminTransferred)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20TransparentCCIPAdminTransferred, error) {
	event := new(BurnMintERC20TransparentCCIPAdminTransferred)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentDefaultAdminDelayChangeCanceledIterator struct {
	Event *BurnMintERC20TransparentDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BurnMintERC20TransparentDefaultAdminDelayChangeCanceled)
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

func (it *BurnMintERC20TransparentDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20TransparentDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentDefaultAdminDelayChangeCanceledIterator{contract: _BurnMintERC20Transparent.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentDefaultAdminDelayChangeCanceled)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20TransparentDefaultAdminDelayChangeCanceled, error) {
	event := new(BurnMintERC20TransparentDefaultAdminDelayChangeCanceled)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentDefaultAdminDelayChangeScheduledIterator struct {
	Event *BurnMintERC20TransparentDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BurnMintERC20TransparentDefaultAdminDelayChangeScheduled)
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

func (it *BurnMintERC20TransparentDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20TransparentDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentDefaultAdminDelayChangeScheduledIterator{contract: _BurnMintERC20Transparent.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentDefaultAdminDelayChangeScheduled)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20TransparentDefaultAdminDelayChangeScheduled, error) {
	event := new(BurnMintERC20TransparentDefaultAdminDelayChangeScheduled)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentDefaultAdminTransferCanceledIterator struct {
	Event *BurnMintERC20TransparentDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentDefaultAdminTransferCanceled)
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
		it.Event = new(BurnMintERC20TransparentDefaultAdminTransferCanceled)
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

func (it *BurnMintERC20TransparentDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20TransparentDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentDefaultAdminTransferCanceledIterator{contract: _BurnMintERC20Transparent.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentDefaultAdminTransferCanceled)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20TransparentDefaultAdminTransferCanceled, error) {
	event := new(BurnMintERC20TransparentDefaultAdminTransferCanceled)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentDefaultAdminTransferScheduledIterator struct {
	Event *BurnMintERC20TransparentDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentDefaultAdminTransferScheduled)
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
		it.Event = new(BurnMintERC20TransparentDefaultAdminTransferScheduled)
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

func (it *BurnMintERC20TransparentDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20TransparentDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentDefaultAdminTransferScheduledIterator{contract: _BurnMintERC20Transparent.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentDefaultAdminTransferScheduled)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20TransparentDefaultAdminTransferScheduled, error) {
	event := new(BurnMintERC20TransparentDefaultAdminTransferScheduled)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentInitializedIterator struct {
	Event *BurnMintERC20TransparentInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentInitialized)
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
		it.Event = new(BurnMintERC20TransparentInitialized)
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

func (it *BurnMintERC20TransparentInitializedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentInitialized struct {
	Version uint64
	Raw     types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20TransparentInitializedIterator, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentInitializedIterator{contract: _BurnMintERC20Transparent.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentInitialized) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentInitialized)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseInitialized(log types.Log) (*BurnMintERC20TransparentInitialized, error) {
	event := new(BurnMintERC20TransparentInitialized)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentRoleAdminChangedIterator struct {
	Event *BurnMintERC20TransparentRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentRoleAdminChanged)
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
		it.Event = new(BurnMintERC20TransparentRoleAdminChanged)
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

func (it *BurnMintERC20TransparentRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20TransparentRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentRoleAdminChangedIterator{contract: _BurnMintERC20Transparent.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentRoleAdminChanged)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20TransparentRoleAdminChanged, error) {
	event := new(BurnMintERC20TransparentRoleAdminChanged)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentRoleGrantedIterator struct {
	Event *BurnMintERC20TransparentRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentRoleGranted)
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
		it.Event = new(BurnMintERC20TransparentRoleGranted)
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

func (it *BurnMintERC20TransparentRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20TransparentRoleGrantedIterator, error) {

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

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentRoleGrantedIterator{contract: _BurnMintERC20Transparent.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentRoleGranted)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseRoleGranted(log types.Log) (*BurnMintERC20TransparentRoleGranted, error) {
	event := new(BurnMintERC20TransparentRoleGranted)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentRoleRevokedIterator struct {
	Event *BurnMintERC20TransparentRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentRoleRevoked)
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
		it.Event = new(BurnMintERC20TransparentRoleRevoked)
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

func (it *BurnMintERC20TransparentRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20TransparentRoleRevokedIterator, error) {

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

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentRoleRevokedIterator{contract: _BurnMintERC20Transparent.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentRoleRevoked)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20TransparentRoleRevoked, error) {
	event := new(BurnMintERC20TransparentRoleRevoked)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransparentTransferIterator struct {
	Event *BurnMintERC20TransparentTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransparentTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20TransparentTransfer)
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
		it.Event = new(BurnMintERC20TransparentTransfer)
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

func (it *BurnMintERC20TransparentTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransparentTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20TransparentTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20TransparentTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransparentTransferIterator{contract: _BurnMintERC20Transparent.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20Transparent.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20TransparentTransfer)
				if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20TransparentFilterer) ParseTransfer(log types.Log) (*BurnMintERC20TransparentTransfer, error) {
	event := new(BurnMintERC20TransparentTransfer)
	if err := _BurnMintERC20Transparent.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20Transparent *BurnMintERC20Transparent) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC20Transparent.abi.Events["Approval"].ID:
		return _BurnMintERC20Transparent.ParseApproval(log)
	case _BurnMintERC20Transparent.abi.Events["CCIPAdminTransferred"].ID:
		return _BurnMintERC20Transparent.ParseCCIPAdminTransferred(log)
	case _BurnMintERC20Transparent.abi.Events["DefaultAdminDelayChangeCanceled"].ID:
		return _BurnMintERC20Transparent.ParseDefaultAdminDelayChangeCanceled(log)
	case _BurnMintERC20Transparent.abi.Events["DefaultAdminDelayChangeScheduled"].ID:
		return _BurnMintERC20Transparent.ParseDefaultAdminDelayChangeScheduled(log)
	case _BurnMintERC20Transparent.abi.Events["DefaultAdminTransferCanceled"].ID:
		return _BurnMintERC20Transparent.ParseDefaultAdminTransferCanceled(log)
	case _BurnMintERC20Transparent.abi.Events["DefaultAdminTransferScheduled"].ID:
		return _BurnMintERC20Transparent.ParseDefaultAdminTransferScheduled(log)
	case _BurnMintERC20Transparent.abi.Events["Initialized"].ID:
		return _BurnMintERC20Transparent.ParseInitialized(log)
	case _BurnMintERC20Transparent.abi.Events["RoleAdminChanged"].ID:
		return _BurnMintERC20Transparent.ParseRoleAdminChanged(log)
	case _BurnMintERC20Transparent.abi.Events["RoleGranted"].ID:
		return _BurnMintERC20Transparent.ParseRoleGranted(log)
	case _BurnMintERC20Transparent.abi.Events["RoleRevoked"].ID:
		return _BurnMintERC20Transparent.ParseRoleRevoked(log)
	case _BurnMintERC20Transparent.abi.Events["Transfer"].ID:
		return _BurnMintERC20Transparent.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC20TransparentApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20TransparentCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20TransparentDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (BurnMintERC20TransparentDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (BurnMintERC20TransparentDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (BurnMintERC20TransparentDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (BurnMintERC20TransparentInitialized) Topic() common.Hash {
	return common.HexToHash("0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2")
}

func (BurnMintERC20TransparentRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20TransparentRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20TransparentRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20TransparentTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_BurnMintERC20Transparent *BurnMintERC20Transparent) Address() common.Address {
	return _BurnMintERC20Transparent.address
}

type BurnMintERC20TransparentInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

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

	Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20TransparentApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20TransparentApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20TransparentCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20TransparentCCIPAdminTransferred, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20TransparentDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20TransparentDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20TransparentDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20TransparentDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20TransparentDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20TransparentDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20TransparentDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20TransparentDefaultAdminTransferScheduled, error)

	FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20TransparentInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*BurnMintERC20TransparentInitialized, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20TransparentRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20TransparentRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20TransparentRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20TransparentRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20TransparentRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20TransparentRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20TransparentTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20TransparentTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20TransparentTransfer, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
