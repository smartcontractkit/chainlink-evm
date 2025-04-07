// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_pool_factory

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

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type TokenPoolFactoryRemoteChainConfig struct {
	RemotePoolFactory   common.Address
	RemoteRouter        common.Address
	RemoteRMNProxy      common.Address
	RemoteTokenDecimals uint8
}

type TokenPoolFactoryRemoteTokenPoolInfo struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	RemotePoolInitCode  []byte
	RemoteChainConfig   TokenPoolFactoryRemoteChainConfig
	PoolType            uint8
	RemoteTokenAddress  []byte
	RemoteTokenInitCode []byte
	RateLimiterConfig   RateLimiterConfig
}

var TokenPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contractITokenAdminRegistry\"},{\"name\":\"tokenAdminModule\",\"type\":\"address\",\"internalType\":\"contractRegistryModuleOwnerCustom\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenAndTokenPool\",\"inputs\":[{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"structTokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enumTokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenPoolWithExistingToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"structTokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enumTokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enumTokenPoolFactory.PoolType\"}],\"outputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RemoteChainConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101003461013757601f6120cc38819003918201601f19168301916001600160401b0383118484101761013c578084926080946040528339810103126101375780516001600160a01b0381169190828103610137576020820151906001600160a01b03821680830361013757610083606061007c60408701610152565b9501610152565b941590811561012e575b50801561011d575b801561010c575b6100fb5760805260a05260c05260e052604051611f6590816101678239608051816109f7015260a05181610973015260c0518181816106c80152818161153501526118c1015260e0518181816106a601528181611513015261189f0152f35b63f6b2911f60e01b60005260046000fd5b506001600160a01b0384161561009c565b506001600160a01b03831615610095565b9050153861008d565b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101375756fe60a080604052600436101561001357600080fd5b600090813560e01c908163169ed05814610ee657508063181f5a7714610e675763eb03cac11461004257600080fd5b34610bc15760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610bc15760043567ffffffffffffffff8111610bce57610091903690600401611987565b9061009a611969565b9160443567ffffffffffffffff8111610e63576100bb903690600401611b2d565b60643567ffffffffffffffff8111610e5f576100de9094929436906004016119b8565b93909261016360405160208101906101598161012d3360843586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282611a86565b5190209384611d2f565b9061016d87611ba2565b9561017b6040519788611a86565b8787527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06101a889611ba2565b01895b818110610e145750506040516101c081611a4d565b89815260606020820152606060408201526040516101dd81611a6a565b8a81528a60208201528a60408201528a60608201526060820152896080820152606060a0820152606060c0820152610213611bba565b60e08201525088956040965b8981101561066e578060051b8601357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe618736030181121561066a5786018036036101a081126106665789519161027483611a4d565b803567ffffffffffffffff811681036105f4578352602081013567ffffffffffffffff81116105f4576102aa9036908301611b2d565b60208401528a81013567ffffffffffffffff81116105f4576102cf9036908301611b2d565b8b84015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0830112610662578a5161030881611a6a565b61031460608301611943565b815261032260808301611943565b602082015261033360a08301611943565b8c82015261034360c08301611979565b6060820152606084015260e081013560028110156105f457608084015261010081013567ffffffffffffffff81116105f4576103829036908301611b2d565b60a084015261012081013567ffffffffffffffff81116105f4577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec06103cc60609236908501611b2d565b9360c086019485520112610662578a51906103e6826119e6565b610140810135801515810361065e576104179161018091845261040c6101608201611bd9565b602085015201611bd9565b8b82015260e083015260a082015151156105f8575b50602081015151156104fa575b88516104458a82611a86565b600181528c5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08b0181106104e957509060019291602082015161048882611bf6565b5261049281611bf6565b5067ffffffffffffffff8251169160e060a0820151910151918c51936104b785611a31565b845260208401528b83015280606083015260808201526104d7828c611c32565b526104e2818b611c32565b500161021f565b80606060208093850101520161044b565b8881015160608201519060a08301516020818051810103126105f457602001519173ffffffffffffffffffffffffffffffffffffffff831683036105f45760808401519260028410156105c65791610590939173ffffffffffffffffffffffffffffffffffffffff61056d941691611dff565b73ffffffffffffffffffffffffffffffffffffffff60608401515116908a611dc9565b73ffffffffffffffffffffffffffffffffffffffff8a5191166020820152602081526105bc8a82611a86565b6020820152610439565b5060248f7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b8e80fd5b61062790516020815191012073ffffffffffffffffffffffffffffffffffffffff60608401515116908a611dc9565b73ffffffffffffffffffffffffffffffffffffffff8a5191166020820152602081526106538a82611a86565b60a08201523861042c565b8f80fd5b8d80fd5b8c80fd5b8b80fd5b8885898d8761076d888d6107688a60208096895161068c8382611a86565b898152600036813761071a8c6106ee8d51948592878401957f0000000000000000000000000000000000000000000000000000000000000000927f00000000000000000000000000000000000000000000000000000000000000009288611ce3565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283611a86565b61073a878c5198899686880137850191848301938c855251938491611ac7565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283611a86565b611d2f565b94835161077a8382611a86565b838152600036813773ffffffffffffffffffffffffffffffffffffffff87163b15610e10579091839085519384917fe8a1da170000000000000000000000000000000000000000000000000000000083526044830188600485015285518091528160648501960190855b818110610dec575050507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8386030160248401528151908186528086019581808460051b83010194019686915b848310610c7d5750505050508192935003818373ffffffffffffffffffffffffffffffffffffffff8a165af18015610bc457908291610c68575b505073ffffffffffffffffffffffffffffffffffffffff84163b15610bc15781517ff2fde38b00000000000000000000000000000000000000000000000000000000815233600482015281816024818373ffffffffffffffffffffffffffffffffffffffff8a165af18015610bc457908291610c53575b505073ffffffffffffffffffffffffffffffffffffffff8316803b15610bce5782517fc630948d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152828160248183865af18015610be757908391610c3e575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610bf15782809160248651809481937f96ea2f7a0000000000000000000000000000000000000000000000000000000083528760048401525af18015610be757908391610c29575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610bf15783517f156194da000000000000000000000000000000000000000000000000000000008152826004820152838160248183865af18015610c0a57908491610c14575b5050803b15610bf15783517f4e847fc700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015287166024820152838160448183865af18015610c0a57908491610bf5575b5050803b15610bf15783517fddadfa8e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201523360248201529083908290604490829084905af18015610be757908391610bd2575b5050803b15610bce5781809160248551809481937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af18015610bc457610bac575b50505173ffffffffffffffffffffffffffffffffffffffff918216815291166020820152604090f35b0390f35b610bb7828092611a86565b610bc15780610b7f565b80fd5b83513d84823e3d90fd5b5080fd5b81610bdc91611a86565b610bce578186610b35565b84513d85823e3d90fd5b8280fd5b81610bff91611a86565b610bf1578287610ac9565b85513d86823e3d90fd5b81610c1e91611a86565b610bf1578287610a60565b81610c3391611a86565b610bce5781866109de565b81610c4891611a86565b610bce57818661095a565b81610c5d91611a86565b610bc15780856108e2565b81610c7291611a86565b610bc157808561086b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe091949750809396508592950301835284875191610120810167ffffffffffffffff84511682528c8c8486015192610120868601528351809152610140850190866101408260051b8801019501925b818110610d9d5750505050839260c06080610d19889585610d8a9660019b01519086830390870152611aea565b94610d56606082015160608601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01519101906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b9801930193018895929693889592610831565b92959680919450610dd9867ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08a60019699030188528851611aea565b960194019101908f928b96959492610cec565b825167ffffffffffffffff16885296830196899650889550918301916001016107e4565b8380fd5b60209060409a989a51610e2681611a31565b8c815260608382015260606040820152610e3e611bba565b6060820152610e4b611bba565b608082015282828c010152019896986101ab565b8580fd5b8480fd5b5034610bc157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610bc15750610ba8604051610ea8604082611a86565b601681527f546f6b656e506f6f6c466163746f727920312e352e31000000000000000000006020820152604051918291602083526020830190611aea565b905034610bce5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610bce576004359073ffffffffffffffffffffffffffffffffffffffff82168203610bf157610f40611969565b9060443567ffffffffffffffff8111610e6357610f61903690600401611987565b9160643567ffffffffffffffff811161193f57610f829036906004016119b8565b9091600260a435101561193b57608435602082019081523360601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016604083015294969490610fd5816054810161012d565b51902091610fe287611ba2565b95610ff06040519788611a86565b8787527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061101d89611ba2565b01895b8181106118f057505060405161103581611a4d565b898152606060208201526060604082015260405161105281611a6a565b8a81528a60208201528a60408201528a60608201526060820152896080820152606060a0820152606060c0820152611088611bba565b60e08201525088956040965b898110156114d0578060051b8701357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe618836030181121561066a576101a081890136031261066a578851906110e882611a4d565b8089013567ffffffffffffffff811681036106625782526020818a01013567ffffffffffffffff811161066257611124903690838c0101611b2d565b602083015289818a01013567ffffffffffffffff81116106625761114d903690838c0101611b2d565b8a83015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0828b013603011261066657895161118a81611a6a565b6111986060838c0101611943565b81526111a86080838c0101611943565b60208201526111bb60a0838c0101611943565b8b8201526111cd60c0838c0101611979565b6060820152606083015260e0818a0101359060028210156106625760808301918252610100818b01013567ffffffffffffffff81116105f457611215903690838d0101611b2d565b60a0840152610120818b01013567ffffffffffffffff81116105f457611240903690838d0101611b2d565b9060c0840191825260607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec0828d01360301126105f4578b5190611282826119e6565b610140818d01013580151581036114cb5782526112b990610180906112ac8e820161016001611bd9565b60208501528d0101611bd9565b8c82015260e084015260a08301515115611465575b506020820151511561139d575b5088516112e88a82611a86565b600181528c5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08b01811061138c57509060019291602082015161132b82611bf6565b5261133581611bf6565b5067ffffffffffffffff8251169160e060a0820151910151918c519361135a85611a31565b845260208401528b830152806060830152608082015261137a828c611c32565b52611385818b611c32565b5001611094565b8060606020809385010152016112ee565b89820151606083015160a084015160208180518101031261065e57602001519273ffffffffffffffffffffffffffffffffffffffff8416840361065e57519260028410156105c6579161142e939173ffffffffffffffffffffffffffffffffffffffff61140b941691611dff565b73ffffffffffffffffffffffffffffffffffffffff606084015151169089611dc9565b73ffffffffffffffffffffffffffffffffffffffff8a51911660208201526020815261145a8a82611a86565b6020820152386112db565b61149490516020815191012073ffffffffffffffffffffffffffffffffffffffff60608501515116908a611dc9565b73ffffffffffffffffffffffffffffffffffffffff8b5191166020820152602081526114c08b82611a86565b60a0830152386112ce565b508f80fd5b50899550916020806107689361157e979560609160a435156000146118695761155d91925061012d8c51916115058684611a86565b8c83528c3681378d519485937f0000000000000000000000000000000000000000000000000000000000000000927f000000000000000000000000000000000000000000000000000000000000000092898701611ce3565b61073a878b51988996858801378501918383018b8152815194859201611ac7565b6020608081905283519194906115949083611a86565b8282528236813773ffffffffffffffffffffffffffffffffffffffff85163b15610bf15790829084519283917fe8a1da1700000000000000000000000000000000000000000000000000000000835260448301876004850152815180915260648401916080510190855b818110611842575050508281037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0160248401528151808252608051600582901b830181019381019286918101905b83831061173e575050505050819003818373ffffffffffffffffffffffffffffffffffffffff89165af18015610bc457829061172e575b505073ffffffffffffffffffffffffffffffffffffffff83163b15610bc15781517ff2fde38b00000000000000000000000000000000000000000000000000000000815233600482015281816024818373ffffffffffffffffffffffffffffffffffffffff89165af18015610bc457611719575b825173ffffffffffffffffffffffffffffffffffffffff8516815260805190f35b611724828092611a86565b610bc157806116f8565b61173791611a86565b8381611684565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0919497508093965085929503018352855190610120810167ffffffffffffffff83511682528a60805184015191610120608051850152825180915261014084016101408260051b8601019360805101918d5b8181106117f457505050509160c06080610d196117dd9486889760019901519086830390870152611aea565b95608051019260805101930187959387959261164d565b919350919361182d817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec088600194030186528651611aea565b946080510193608051019101918e93926117b1565b825167ffffffffffffffff16845260805189975088965093840193909201916001016115fe565b600160a4351461187b575b505061155d565b6118e991925061012d8c51916118918684611a86565b8c83528c3681378d519485937f0000000000000000000000000000000000000000000000000000000000000000927f000000000000000000000000000000000000000000000000000000000000000092898701611c90565b8b80611874565b60209060409a989a5161190281611a31565b8c81526060838201526060604082015261191a611bba565b6060820152611927611bba565b608082015282828c01015201989698611020565b8780fd5b8680fd5b359073ffffffffffffffffffffffffffffffffffffffff8216820361196457565b600080fd5b6024359060ff8216820361196457565b359060ff8216820361196457565b9181601f840112156119645782359167ffffffffffffffff8311611964576020808501948460051b01011161196457565b9181601f840112156119645782359167ffffffffffffffff8311611964576020838186019501011161196457565b6060810190811067ffffffffffffffff821117611a0257604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff821117611a0257604052565b610100810190811067ffffffffffffffff821117611a0257604052565b6080810190811067ffffffffffffffff821117611a0257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611a0257604052565b60005b838110611ada5750506000910152565b8181015183820152602001611aca565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093611b2681518092818752878088019101611ac7565b0116010190565b81601f820112156119645780359067ffffffffffffffff8211611a025760405192611b80601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185611a86565b8284526020838301011161196457816000926020809301838601378301015290565b67ffffffffffffffff8111611a025760051b60200190565b60405190611bc7826119e6565b60006040838281528260208201520152565b35906fffffffffffffffffffffffffffffffff8216820361196457565b805115611c035760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015611c035760209160051b010190565b906020808351928381520192019060005b818110611c645750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611c57565b93959490611ccf60a09460ff73ffffffffffffffffffffffffffffffffffffffff9586809516895216602088015260c0604088015260c0870190611c46565b961660608501526001608085015216910152565b93959490611d2260809460ff73ffffffffffffffffffffffffffffffffffffffff9586809516895216602088015260a0604088015260a0870190611c46565b9616606085015216910152565b90805115611d9f576020815191016000f5903d1519821516611d935773ffffffffffffffffffffffffffffffffffffffff821615611d6957565b7fb06ebf3d0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b60559173ffffffffffffffffffffffffffffffffffffffff93600b92604051926040840152602083015281520160ff8153201690565b926002811015611f2957611ecb578060ff6060611ec593015116611e6b6020916106ee60405194611e308587611a86565b60008652600036813773ffffffffffffffffffffffffffffffffffffffff858160408401511692015116906040519687948786019a8b611ce3565b604051938492611e9783611e88818701998a815193849201611ac7565b85019151809385840190611ac7565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282611a86565b51902090565b8060ff6060611ec593015116611e6b6020916106ee60405194611eee8587611a86565b60008652600036813773ffffffffffffffffffffffffffffffffffffffff858160408401511692015116906040519687948786019a8b611c90565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c634300081a000a",
}

var TokenPoolFactoryABI = TokenPoolFactoryMetaData.ABI

var TokenPoolFactoryBin = TokenPoolFactoryMetaData.Bin

func DeployTokenPoolFactory(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address, tokenAdminModule common.Address, rmnProxy common.Address, ccipRouter common.Address) (common.Address, *types.Transaction, *TokenPoolFactory, error) {
	parsed, err := TokenPoolFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenPoolFactoryBin), backend, tokenAdminRegistry, tokenAdminModule, rmnProxy, ccipRouter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPoolFactory{address: address, abi: *parsed, TokenPoolFactoryCaller: TokenPoolFactoryCaller{contract: contract}, TokenPoolFactoryTransactor: TokenPoolFactoryTransactor{contract: contract}, TokenPoolFactoryFilterer: TokenPoolFactoryFilterer{contract: contract}}, nil
}

type TokenPoolFactory struct {
	address common.Address
	abi     abi.ABI
	TokenPoolFactoryCaller
	TokenPoolFactoryTransactor
	TokenPoolFactoryFilterer
}

type TokenPoolFactoryCaller struct {
	contract *bind.BoundContract
}

type TokenPoolFactoryTransactor struct {
	contract *bind.BoundContract
}

type TokenPoolFactoryFilterer struct {
	contract *bind.BoundContract
}

type TokenPoolFactorySession struct {
	Contract     *TokenPoolFactory
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TokenPoolFactoryCallerSession struct {
	Contract *TokenPoolFactoryCaller
	CallOpts bind.CallOpts
}

type TokenPoolFactoryTransactorSession struct {
	Contract     *TokenPoolFactoryTransactor
	TransactOpts bind.TransactOpts
}

type TokenPoolFactoryRaw struct {
	Contract *TokenPoolFactory
}

type TokenPoolFactoryCallerRaw struct {
	Contract *TokenPoolFactoryCaller
}

type TokenPoolFactoryTransactorRaw struct {
	Contract *TokenPoolFactoryTransactor
}

func NewTokenPoolFactory(address common.Address, backend bind.ContractBackend) (*TokenPoolFactory, error) {
	abi, err := abi.JSON(strings.NewReader(TokenPoolFactoryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTokenPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactory{address: address, abi: abi, TokenPoolFactoryCaller: TokenPoolFactoryCaller{contract: contract}, TokenPoolFactoryTransactor: TokenPoolFactoryTransactor{contract: contract}, TokenPoolFactoryFilterer: TokenPoolFactoryFilterer{contract: contract}}, nil
}

func NewTokenPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenPoolFactoryCaller, error) {
	contract, err := bindTokenPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryCaller{contract: contract}, nil
}

func NewTokenPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPoolFactoryTransactor, error) {
	contract, err := bindTokenPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryTransactor{contract: contract}, nil
}

func NewTokenPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPoolFactoryFilterer, error) {
	contract, err := bindTokenPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryFilterer{contract: contract}, nil
}

func bindTokenPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenPoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolFactory.Contract.TokenPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.TokenPoolFactoryTransactor.contract.Transfer(opts)
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.TokenPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.contract.Transfer(opts)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.contract.Transact(opts, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenPoolFactory.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TokenPoolFactory *TokenPoolFactorySession) TypeAndVersion() (string, error) {
	return _TokenPoolFactory.Contract.TypeAndVersion(&_TokenPoolFactory.CallOpts)
}

func (_TokenPoolFactory *TokenPoolFactoryCallerSession) TypeAndVersion() (string, error) {
	return _TokenPoolFactory.Contract.TypeAndVersion(&_TokenPoolFactory.CallOpts)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenAndTokenPool", remoteTokenPools, localTokenDecimals, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, tokenInitCode, tokenPoolInitCode, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte, poolType uint8) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenPoolWithExistingToken", token, localTokenDecimals, remoteTokenPools, tokenPoolInitCode, salt, poolType)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte, poolType uint8) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, remoteTokenPools, tokenPoolInitCode, salt, poolType)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte, poolType uint8) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, remoteTokenPools, tokenPoolInitCode, salt, poolType)
}

type TokenPoolFactoryRemoteChainConfigUpdatedIterator struct {
	Event *TokenPoolFactoryRemoteChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolFactoryRemoteChainConfigUpdated)
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
		it.Event = new(TokenPoolFactoryRemoteChainConfigUpdated)
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

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolFactoryRemoteChainConfigUpdated struct {
	RemoteChainSelector uint64
	RemoteChainConfig   TokenPoolFactoryRemoteChainConfig
	Raw                 types.Log
}

func (_TokenPoolFactory *TokenPoolFactoryFilterer) FilterRemoteChainConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolFactoryRemoteChainConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolFactory.contract.FilterLogs(opts, "RemoteChainConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryRemoteChainConfigUpdatedIterator{contract: _TokenPoolFactory.contract, event: "RemoteChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_TokenPoolFactory *TokenPoolFactoryFilterer) WatchRemoteChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolFactoryRemoteChainConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolFactory.contract.WatchLogs(opts, "RemoteChainConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolFactoryRemoteChainConfigUpdated)
				if err := _TokenPoolFactory.contract.UnpackLog(event, "RemoteChainConfigUpdated", log); err != nil {
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

func (_TokenPoolFactory *TokenPoolFactoryFilterer) ParseRemoteChainConfigUpdated(log types.Log) (*TokenPoolFactoryRemoteChainConfigUpdated, error) {
	event := new(TokenPoolFactoryRemoteChainConfigUpdated)
	if err := _TokenPoolFactory.contract.UnpackLog(event, "RemoteChainConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_TokenPoolFactory *TokenPoolFactory) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _TokenPoolFactory.abi.Events["RemoteChainConfigUpdated"].ID:
		return _TokenPoolFactory.ParseRemoteChainConfigUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (TokenPoolFactoryRemoteChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xe3606343290c8e3853a7f92686979dfe24a0f29594186b61177236cb145126f0")
}

func (_TokenPoolFactory *TokenPoolFactory) Address() common.Address {
	return _TokenPoolFactory.address
}

type TokenPoolFactoryInterface interface {
	TypeAndVersion(opts *bind.CallOpts) (string, error)

	DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, tokenInitCode []byte, tokenPoolInitCode []byte, salt [32]byte) (*types.Transaction, error)

	DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, salt [32]byte, poolType uint8) (*types.Transaction, error)

	FilterRemoteChainConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolFactoryRemoteChainConfigUpdatedIterator, error)

	WatchRemoteChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolFactoryRemoteChainConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigUpdated(log types.Log) (*TokenPoolFactoryRemoteChainConfigUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
