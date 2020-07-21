// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrage

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbitrageABI is the input ABI used to generate the binding from.
const ArbitrageABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractKyberERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAI_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KYBERSWAP_PROXY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_1_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeB\",\"outputs\":[{\"internalType\":\"contractIUniswapExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractKyberERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenBought\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"execSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingPool\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractKyberERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"makeArbitrage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractKyberNetworkProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"sellEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniSwap1Exchange\",\"outputs\":[{\"internalType\":\"contractIUniswapExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapOneFactory\",\"outputs\":[{\"internalType\":\"contractIUniswapFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ArbitrageBin is the compiled bytecode used for deploying new contracts.
var ArbitrageBin = "0x60806040523480156200001157600080fd5b507324a42fd28c976a61df5d00d0599c34c4f90748c8600062000039620003ec60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505073c0a47dfe034b400b47bdad5fecda2621de6c4d95600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306f2bf62736b175474e89094c44da98b954eedeac495271d0f6040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156200022457600080fd5b505afa15801562000239573d6000803e3d6000fd5b505050506040513d60208110156200025057600080fd5b8101908080519060200190929190505050905080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073818e6fecd516ecc3849daf6845e3ec868087b755600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200036457600080fd5b505afa15801562000379573d6000803e3d6000fd5b505050506040513d60208110156200039057600080fd5b8101908080519060200190929190505050905080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620003f4565b600033905090565b61250380620004046000396000f3fe60806040526004361061010d5760003560e01c80638b698f9011610095578063b5607acf11610064578063b5607acf1461055b578063c72c4d10146105b2578063ec55688914610609578063ee87255814610660578063f2fde38b1461071a57610114565b80638b698f90146103ff5780638da5cb5b146104565780639d76ea58146104ad578063a59a99731461050457610114565b806340032b9c116100dc57806340032b9c1461027557806351cff8d9146102ce578063575023e81461031f5780635c3c60681461037a578063715018a6146103e857610114565b80630d764fbe1461011957806311ff70621461017057806320fee44b146101c75780632a4c0a1a1461021e57610114565b3661011457005b600080fd5b34801561012557600080fd5b5061012e61076b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017c57600080fd5b50610185610791565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101d357600080fd5b506101dc6107b7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561022a57600080fd5b506102336107dd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561028157600080fd5b506102b86004803603604081101561029857600080fd5b8101908080359060200190929190803590602001909291905050506107f5565b6040518082815260200191505060405180910390f35b3480156102da57600080fd5b5061031d600480360360208110156102f157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a35565b005b34801561032b57600080fd5b506103786004803603604081101561034257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cef565b005b6103e66004803603606081101561039057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611152565b005b3480156103f457600080fd5b506103fd6114f2565b005b34801561040b57600080fd5b5061041461167a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046257600080fd5b5061046b611692565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104b957600080fd5b506104c26116bb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561051057600080fd5b506105196116e1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561056757600080fd5b50610570611707565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105be57600080fd5b506105c761171f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561061557600080fd5b5061061e611745565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561066c57600080fd5b506107186004803603608081101561068357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001906401000000008111156106d457600080fd5b8201836020820111156106e657600080fd5b8035906020019184600183028401116401000000008311171561070857600080fd5b909192939192939050505061176b565b005b34801561072657600080fd5b506107696004803603602081101561073d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611bab565b005b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b736b175474e89094c44da98b954eedeac495271d0f81565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156108c257600080fd5b505af11580156108d6573d6000803e3d6000fd5b505050506040513d60208110156108ec57600080fd5b810190808051906020019092919050505061096f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f436f756c64206e6f7420617070726f7665204574682073656c6c00000000000081525060200191505060405180910390fd5b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f39b5b9b85856040518363ffffffff1660e01b81526004018083815260200182815260200192505050602060405180830381600087803b1580156109ee57600080fd5b505af1158015610a02573d6000803e3d6000fd5b505050506040513d6020811015610a1857600080fd5b810190808051906020019092919050505090508091505092915050565b610a3d611db8565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610afe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ba05760003090508073ffffffffffffffffffffffffffffffffffffffff163191503373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610b99573d6000803e3d6000fd5b5050610c86565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610c1d57600080fd5b505afa158015610c31573d6000803e3d6000fd5b505050506040513d6020811015610c4757600080fd5b81019080805190602001909291905050509050610c8533828473ffffffffffffffffffffffffffffffffffffffff16611dc09092919063ffffffff16565b5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099836040518082815260200191505060405180910390a35050565b610cf7611db8565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610db8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606060405180602001604052806000815250905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306f2bf62600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610f1157600080fd5b505afa158015610f25573d6000803e3d6000fd5b505050506040513d6020811015610f3b57600080fd5b8101908080519060200190929190505050905080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635cffe9de3073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee87866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561109e578082015181840152602081019050611083565b50505050905090810190601f1680156110cb5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b1580156110ed57600080fd5b505af1158015611101573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561114b573d6000803e3d6000fd5b5050505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558373eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee346040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b15801561124257600080fd5b505afa158015611256573d6000803e3d6000fd5b505050506040513d604081101561126c57600080fd5b81019080805190602001909291908051906020019092919050505050809150506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633bba21dc8686856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b15801561133f57600080fd5b505af1158015611353573d6000803e3d6000fd5b505050506040513d602081101561136957600080fd5b81019080805190602001909291905050509050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561142557600080fd5b505af1158015611439573d6000803e3d6000fd5b505050506040513d602081101561144f57600080fd5b810190808051906020019092919050505061146957600080fd5b3373ffffffffffffffffffffffffffffffffffffffff167fea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca68683604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a25050505050565b6114fa611db8565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146115bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b73818e6fecd516ecc3849daf6845e3ec868087b75581565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73c0a47dfe034b400b47bdad5fecda2621de6c4d9581565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000611775611e78565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16876040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156118a557600080fd5b505af11580156118b9573d6000803e3d6000fd5b505050506040513d60208110156118cf57600080fd5b8101908080519060200190929190505050611952576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f436f756c64206e6f7420617070726f7665204574682073656c6c00000000000081525060200191505060405180910390fd5b600061195e86836107f5565b9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611a2b57600080fd5b505af1158015611a3f573d6000803e3d6000fd5b505050506040513d6020811015611a5557600080fd5b8101908080519060200190929190505050611ad8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f436f756c64206e6f7420617070726f766520746f6b656e2073656c6c0000000081525060200191505060405180910390fd5b611b05600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168230611152565b60004790506000611b1f8789611e8490919063ffffffff16565b9050808211611b96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f446964206e6f742070726f66697400000000000000000000000000000000000081525060200191505060405180910390fd5b611ba08982611f0c565b505050505050505050565b611bb3611db8565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611c74576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611cfa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061247e6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b611e738363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611fc3565b505050565b6000610bb84201905090565b600080828401905083811015611f02576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed6ff7606040518163ffffffff1660e01b815260040160206040518083038186803b158015611f7657600080fd5b505afa158015611f8a573d6000803e3d6000fd5b505050506040513d6020811015611fa057600080fd5b81019080805190602001909291905050509050611fbe8184846120b2565b505050565b6060612025826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166122149092919063ffffffff16565b90506000815111156120ad5780806020019051602081101561204657600080fd5b81019080805190602001909291905050506120ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806124a4602a913960400191505060405180910390fd5b5b505050565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156121e35760008373ffffffffffffffffffffffffffffffffffffffff168260405180600001905060006040518083038185875af1925050503d806000811461215a576040519150601f19603f3d011682016040523d82523d6000602084013e61215f565b606091505b5050905060011515811515146121dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f436f756c646e2774207472616e7366657220455448000000000000000000000081525060200191505060405180910390fd5b5061220f565b61220e83828473ffffffffffffffffffffffffffffffffffffffff16611dc09092919063ffffffff16565b5b505050565b6060612223848460008561222c565b90509392505050565b606061223785612432565b6122a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b602083106122f957805182526020820191506020810190506020830392506122d6565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d806000811461235b576040519150601f19603f3d011682016040523d82523d6000602084013e612360565b606091505b5091509150811561237557809250505061242a565b6000815111156123885780518082602001fd5b836040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156123ef5780820151818401526020810190506123d4565b50505050905090810190601f16801561241c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b949350505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561247457506000801b8214155b9250505091905056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a264697066735822122070eb5e22fa75ac77b3e08f587877b5d363deee471f005f3e4f19e909b0e6710c64736f6c634300060a0033"

// DeployArbitrage deploys a new Ethereum contract, binding an instance of Arbitrage to it.
func DeployArbitrage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Arbitrage, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbitrageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Arbitrage{ArbitrageCaller: ArbitrageCaller{contract: contract}, ArbitrageTransactor: ArbitrageTransactor{contract: contract}, ArbitrageFilterer: ArbitrageFilterer{contract: contract}}, nil
}

// Arbitrage is an auto generated Go binding around an Ethereum contract.
type Arbitrage struct {
	ArbitrageCaller     // Read-only binding to the contract
	ArbitrageTransactor // Write-only binding to the contract
	ArbitrageFilterer   // Log filterer for contract events
}

// ArbitrageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbitrageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbitrageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbitrageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbitrageSession struct {
	Contract     *Arbitrage        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbitrageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbitrageCallerSession struct {
	Contract *ArbitrageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArbitrageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbitrageTransactorSession struct {
	Contract     *ArbitrageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArbitrageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbitrageRaw struct {
	Contract *Arbitrage // Generic contract binding to access the raw methods on
}

// ArbitrageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbitrageCallerRaw struct {
	Contract *ArbitrageCaller // Generic read-only contract binding to access the raw methods on
}

// ArbitrageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbitrageTransactorRaw struct {
	Contract *ArbitrageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbitrage creates a new instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrage(address common.Address, backend bind.ContractBackend) (*Arbitrage, error) {
	contract, err := bindArbitrage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arbitrage{ArbitrageCaller: ArbitrageCaller{contract: contract}, ArbitrageTransactor: ArbitrageTransactor{contract: contract}, ArbitrageFilterer: ArbitrageFilterer{contract: contract}}, nil
}

// NewArbitrageCaller creates a new read-only instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageCaller(address common.Address, caller bind.ContractCaller) (*ArbitrageCaller, error) {
	contract, err := bindArbitrage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageCaller{contract: contract}, nil
}

// NewArbitrageTransactor creates a new write-only instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrageTransactor, error) {
	contract, err := bindArbitrage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageTransactor{contract: contract}, nil
}

// NewArbitrageFilterer creates a new log filterer instance of Arbitrage, bound to a specific deployed contract.
func NewArbitrageFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrageFilterer, error) {
	contract, err := bindArbitrage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrageFilterer{contract: contract}, nil
}

// bindArbitrage binds a generic wrapper to an already deployed contract.
func bindArbitrage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrage *ArbitrageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Arbitrage.Contract.ArbitrageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrage *ArbitrageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.Contract.ArbitrageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrage *ArbitrageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrage.Contract.ArbitrageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrage *ArbitrageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Arbitrage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrage *ArbitrageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrage *ArbitrageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrage.Contract.contract.Transact(opts, method, params...)
}

// DAIADDRESS is a free data retrieval call binding the contract method 0x2a4c0a1a.
//
// Solidity: function DAI_ADDRESS() view returns(address)
func (_Arbitrage *ArbitrageCaller) DAIADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "DAI_ADDRESS")
	return *ret0, err
}

// DAIADDRESS is a free data retrieval call binding the contract method 0x2a4c0a1a.
//
// Solidity: function DAI_ADDRESS() view returns(address)
func (_Arbitrage *ArbitrageSession) DAIADDRESS() (common.Address, error) {
	return _Arbitrage.Contract.DAIADDRESS(&_Arbitrage.CallOpts)
}

// DAIADDRESS is a free data retrieval call binding the contract method 0x2a4c0a1a.
//
// Solidity: function DAI_ADDRESS() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) DAIADDRESS() (common.Address, error) {
	return _Arbitrage.Contract.DAIADDRESS(&_Arbitrage.CallOpts)
}

// KYBERSWAPPROXY is a free data retrieval call binding the contract method 0x8b698f90.
//
// Solidity: function KYBERSWAP_PROXY() view returns(address)
func (_Arbitrage *ArbitrageCaller) KYBERSWAPPROXY(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "KYBERSWAP_PROXY")
	return *ret0, err
}

// KYBERSWAPPROXY is a free data retrieval call binding the contract method 0x8b698f90.
//
// Solidity: function KYBERSWAP_PROXY() view returns(address)
func (_Arbitrage *ArbitrageSession) KYBERSWAPPROXY() (common.Address, error) {
	return _Arbitrage.Contract.KYBERSWAPPROXY(&_Arbitrage.CallOpts)
}

// KYBERSWAPPROXY is a free data retrieval call binding the contract method 0x8b698f90.
//
// Solidity: function KYBERSWAP_PROXY() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) KYBERSWAPPROXY() (common.Address, error) {
	return _Arbitrage.Contract.KYBERSWAPPROXY(&_Arbitrage.CallOpts)
}

// UNISWAP1FACTORY is a free data retrieval call binding the contract method 0xb5607acf.
//
// Solidity: function UNISWAP_1_FACTORY() view returns(address)
func (_Arbitrage *ArbitrageCaller) UNISWAP1FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "UNISWAP_1_FACTORY")
	return *ret0, err
}

// UNISWAP1FACTORY is a free data retrieval call binding the contract method 0xb5607acf.
//
// Solidity: function UNISWAP_1_FACTORY() view returns(address)
func (_Arbitrage *ArbitrageSession) UNISWAP1FACTORY() (common.Address, error) {
	return _Arbitrage.Contract.UNISWAP1FACTORY(&_Arbitrage.CallOpts)
}

// UNISWAP1FACTORY is a free data retrieval call binding the contract method 0xb5607acf.
//
// Solidity: function UNISWAP_1_FACTORY() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) UNISWAP1FACTORY() (common.Address, error) {
	return _Arbitrage.Contract.UNISWAP1FACTORY(&_Arbitrage.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Arbitrage *ArbitrageCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "addressesProvider")
	return *ret0, err
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Arbitrage *ArbitrageSession) AddressesProvider() (common.Address, error) {
	return _Arbitrage.Contract.AddressesProvider(&_Arbitrage.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) AddressesProvider() (common.Address, error) {
	return _Arbitrage.Contract.AddressesProvider(&_Arbitrage.CallOpts)
}

// ExchangeB is a free data retrieval call binding the contract method 0x20fee44b.
//
// Solidity: function exchangeB() view returns(address)
func (_Arbitrage *ArbitrageCaller) ExchangeB(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "exchangeB")
	return *ret0, err
}

// ExchangeB is a free data retrieval call binding the contract method 0x20fee44b.
//
// Solidity: function exchangeB() view returns(address)
func (_Arbitrage *ArbitrageSession) ExchangeB() (common.Address, error) {
	return _Arbitrage.Contract.ExchangeB(&_Arbitrage.CallOpts)
}

// ExchangeB is a free data retrieval call binding the contract method 0x20fee44b.
//
// Solidity: function exchangeB() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) ExchangeB() (common.Address, error) {
	return _Arbitrage.Contract.ExchangeB(&_Arbitrage.CallOpts)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_Arbitrage *ArbitrageCaller) LendingPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "lendingPool")
	return *ret0, err
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_Arbitrage *ArbitrageSession) LendingPool() (common.Address, error) {
	return _Arbitrage.Contract.LendingPool(&_Arbitrage.CallOpts)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) LendingPool() (common.Address, error) {
	return _Arbitrage.Contract.LendingPool(&_Arbitrage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arbitrage *ArbitrageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arbitrage *ArbitrageSession) Owner() (common.Address, error) {
	return _Arbitrage.Contract.Owner(&_Arbitrage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Owner() (common.Address, error) {
	return _Arbitrage.Contract.Owner(&_Arbitrage.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Arbitrage *ArbitrageCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Arbitrage *ArbitrageSession) Proxy() (common.Address, error) {
	return _Arbitrage.Contract.Proxy(&_Arbitrage.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) Proxy() (common.Address, error) {
	return _Arbitrage.Contract.Proxy(&_Arbitrage.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Arbitrage *ArbitrageCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Arbitrage *ArbitrageSession) TokenAddress() (common.Address, error) {
	return _Arbitrage.Contract.TokenAddress(&_Arbitrage.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) TokenAddress() (common.Address, error) {
	return _Arbitrage.Contract.TokenAddress(&_Arbitrage.CallOpts)
}

// UniSwap1Exchange is a free data retrieval call binding the contract method 0x0d764fbe.
//
// Solidity: function uniSwap1Exchange() view returns(address)
func (_Arbitrage *ArbitrageCaller) UniSwap1Exchange(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "uniSwap1Exchange")
	return *ret0, err
}

// UniSwap1Exchange is a free data retrieval call binding the contract method 0x0d764fbe.
//
// Solidity: function uniSwap1Exchange() view returns(address)
func (_Arbitrage *ArbitrageSession) UniSwap1Exchange() (common.Address, error) {
	return _Arbitrage.Contract.UniSwap1Exchange(&_Arbitrage.CallOpts)
}

// UniSwap1Exchange is a free data retrieval call binding the contract method 0x0d764fbe.
//
// Solidity: function uniSwap1Exchange() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) UniSwap1Exchange() (common.Address, error) {
	return _Arbitrage.Contract.UniSwap1Exchange(&_Arbitrage.CallOpts)
}

// UniswapOneFactory is a free data retrieval call binding the contract method 0x11ff7062.
//
// Solidity: function uniswapOneFactory() view returns(address)
func (_Arbitrage *ArbitrageCaller) UniswapOneFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Arbitrage.contract.Call(opts, out, "uniswapOneFactory")
	return *ret0, err
}

// UniswapOneFactory is a free data retrieval call binding the contract method 0x11ff7062.
//
// Solidity: function uniswapOneFactory() view returns(address)
func (_Arbitrage *ArbitrageSession) UniswapOneFactory() (common.Address, error) {
	return _Arbitrage.Contract.UniswapOneFactory(&_Arbitrage.CallOpts)
}

// UniswapOneFactory is a free data retrieval call binding the contract method 0x11ff7062.
//
// Solidity: function uniswapOneFactory() view returns(address)
func (_Arbitrage *ArbitrageCallerSession) UniswapOneFactory() (common.Address, error) {
	return _Arbitrage.Contract.UniswapOneFactory(&_Arbitrage.CallOpts)
}

// ExecSwap is a paid mutator transaction binding the contract method 0x5c3c6068.
//
// Solidity: function execSwap(address _token, uint256 tokenBought, address destAddress) payable returns()
func (_Arbitrage *ArbitrageTransactor) ExecSwap(opts *bind.TransactOpts, _token common.Address, tokenBought *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "execSwap", _token, tokenBought, destAddress)
}

// ExecSwap is a paid mutator transaction binding the contract method 0x5c3c6068.
//
// Solidity: function execSwap(address _token, uint256 tokenBought, address destAddress) payable returns()
func (_Arbitrage *ArbitrageSession) ExecSwap(_token common.Address, tokenBought *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.ExecSwap(&_Arbitrage.TransactOpts, _token, tokenBought, destAddress)
}

// ExecSwap is a paid mutator transaction binding the contract method 0x5c3c6068.
//
// Solidity: function execSwap(address _token, uint256 tokenBought, address destAddress) payable returns()
func (_Arbitrage *ArbitrageTransactorSession) ExecSwap(_token common.Address, tokenBought *big.Int, destAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.ExecSwap(&_Arbitrage.TransactOpts, _token, tokenBought, destAddress)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_Arbitrage *ArbitrageTransactor) ExecuteOperation(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "executeOperation", _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_Arbitrage *ArbitrageSession) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _Arbitrage.Contract.ExecuteOperation(&_Arbitrage.TransactOpts, _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) returns()
func (_Arbitrage *ArbitrageTransactorSession) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _Arbitrage.Contract.ExecuteOperation(&_Arbitrage.TransactOpts, _reserve, _amount, _fee, _params)
}

// MakeArbitrage is a paid mutator transaction binding the contract method 0x575023e8.
//
// Solidity: function makeArbitrage(uint256 amount, address _token) returns()
func (_Arbitrage *ArbitrageTransactor) MakeArbitrage(opts *bind.TransactOpts, amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "makeArbitrage", amount, _token)
}

// MakeArbitrage is a paid mutator transaction binding the contract method 0x575023e8.
//
// Solidity: function makeArbitrage(uint256 amount, address _token) returns()
func (_Arbitrage *ArbitrageSession) MakeArbitrage(amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.MakeArbitrage(&_Arbitrage.TransactOpts, amount, _token)
}

// MakeArbitrage is a paid mutator transaction binding the contract method 0x575023e8.
//
// Solidity: function makeArbitrage(uint256 amount, address _token) returns()
func (_Arbitrage *ArbitrageTransactorSession) MakeArbitrage(amount *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.MakeArbitrage(&_Arbitrage.TransactOpts, amount, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arbitrage *ArbitrageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arbitrage *ArbitrageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arbitrage.Contract.RenounceOwnership(&_Arbitrage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arbitrage *ArbitrageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arbitrage.Contract.RenounceOwnership(&_Arbitrage.TransactOpts)
}

// SellEth is a paid mutator transaction binding the contract method 0x40032b9c.
//
// Solidity: function sellEth(uint256 _amount, uint256 deadline) returns(uint256)
func (_Arbitrage *ArbitrageTransactor) SellEth(opts *bind.TransactOpts, _amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "sellEth", _amount, deadline)
}

// SellEth is a paid mutator transaction binding the contract method 0x40032b9c.
//
// Solidity: function sellEth(uint256 _amount, uint256 deadline) returns(uint256)
func (_Arbitrage *ArbitrageSession) SellEth(_amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.SellEth(&_Arbitrage.TransactOpts, _amount, deadline)
}

// SellEth is a paid mutator transaction binding the contract method 0x40032b9c.
//
// Solidity: function sellEth(uint256 _amount, uint256 deadline) returns(uint256)
func (_Arbitrage *ArbitrageTransactorSession) SellEth(_amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Arbitrage.Contract.SellEth(&_Arbitrage.TransactOpts, _amount, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arbitrage *ArbitrageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arbitrage *ArbitrageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.TransferOwnership(&_Arbitrage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arbitrage *ArbitrageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.TransferOwnership(&_Arbitrage.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_Arbitrage *ArbitrageTransactor) Withdraw(opts *bind.TransactOpts, _assetAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "withdraw", _assetAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_Arbitrage *ArbitrageSession) Withdraw(_assetAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.Withdraw(&_Arbitrage.TransactOpts, _assetAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_Arbitrage *ArbitrageTransactorSession) Withdraw(_assetAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.Withdraw(&_Arbitrage.TransactOpts, _assetAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrage *ArbitrageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrage *ArbitrageSession) Receive() (*types.Transaction, error) {
	return _Arbitrage.Contract.Receive(&_Arbitrage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrage *ArbitrageTransactorSession) Receive() (*types.Transaction, error) {
	return _Arbitrage.Contract.Receive(&_Arbitrage.TransactOpts)
}

// ArbitrageLogWithdrawIterator is returned from FilterLogWithdraw and is used to iterate over the raw logs and unpacked data for LogWithdraw events raised by the Arbitrage contract.
type ArbitrageLogWithdrawIterator struct {
	Event *ArbitrageLogWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbitrageLogWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrageLogWithdraw)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbitrageLogWithdraw)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbitrageLogWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrageLogWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrageLogWithdraw represents a LogWithdraw event raised by the Arbitrage contract.
type ArbitrageLogWithdraw struct {
	From         common.Address
	AssetAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogWithdraw is a free log retrieval operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed _from, address indexed _assetAddress, uint256 amount)
func (_Arbitrage *ArbitrageFilterer) FilterLogWithdraw(opts *bind.FilterOpts, _from []common.Address, _assetAddress []common.Address) (*ArbitrageLogWithdrawIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _Arbitrage.contract.FilterLogs(opts, "LogWithdraw", _fromRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrageLogWithdrawIterator{contract: _Arbitrage.contract, event: "LogWithdraw", logs: logs, sub: sub}, nil
}

// WatchLogWithdraw is a free log subscription operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed _from, address indexed _assetAddress, uint256 amount)
func (_Arbitrage *ArbitrageFilterer) WatchLogWithdraw(opts *bind.WatchOpts, sink chan<- *ArbitrageLogWithdraw, _from []common.Address, _assetAddress []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _Arbitrage.contract.WatchLogs(opts, "LogWithdraw", _fromRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrageLogWithdraw)
				if err := _Arbitrage.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
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

// ParseLogWithdraw is a log parse operation binding the contract event 0x9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099.
//
// Solidity: event LogWithdraw(address indexed _from, address indexed _assetAddress, uint256 amount)
func (_Arbitrage *ArbitrageFilterer) ParseLogWithdraw(log types.Log) (*ArbitrageLogWithdraw, error) {
	event := new(ArbitrageLogWithdraw)
	if err := _Arbitrage.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbitrageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Arbitrage contract.
type ArbitrageOwnershipTransferredIterator struct {
	Event *ArbitrageOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbitrageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrageOwnershipTransferred)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbitrageOwnershipTransferred)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbitrageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrageOwnershipTransferred represents a OwnershipTransferred event raised by the Arbitrage contract.
type ArbitrageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arbitrage *ArbitrageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbitrageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arbitrage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrageOwnershipTransferredIterator{contract: _Arbitrage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arbitrage *ArbitrageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arbitrage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrageOwnershipTransferred)
				if err := _Arbitrage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arbitrage *ArbitrageFilterer) ParseOwnershipTransferred(log types.Log) (*ArbitrageOwnershipTransferred, error) {
	event := new(ArbitrageOwnershipTransferred)
	if err := _Arbitrage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ArbitrageSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Arbitrage contract.
type ArbitrageSwapIterator struct {
	Event *ArbitrageSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbitrageSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrageSwap)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbitrageSwap)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbitrageSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrageSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrageSwap represents a Swap event raised by the Arbitrage contract.
type ArbitrageSwap struct {
	Sender    common.Address
	DestToken common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed sender, address destToken, uint256 amount)
func (_Arbitrage *ArbitrageFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address) (*ArbitrageSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Arbitrage.contract.FilterLogs(opts, "Swap", senderRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrageSwapIterator{contract: _Arbitrage.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed sender, address destToken, uint256 amount)
func (_Arbitrage *ArbitrageFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ArbitrageSwap, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Arbitrage.contract.WatchLogs(opts, "Swap", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrageSwap)
				if err := _Arbitrage.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed sender, address destToken, uint256 amount)
func (_Arbitrage *ArbitrageFilterer) ParseSwap(log types.Log) (*ArbitrageSwap, error) {
	event := new(ArbitrageSwap)
	if err := _Arbitrage.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	return event, nil
}
