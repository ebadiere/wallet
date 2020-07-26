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
const ArbitrageABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractKyberERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractKyberERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAI_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KYBERSWAP_PROXY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_1_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeB\",\"outputs\":[{\"internalType\":\"contractIUniswapExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractKyberERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractKyberERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"execSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingPool\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractKyberERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"makeArbitrage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractKyberNetworkProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniSwap1Exchange\",\"outputs\":[{\"internalType\":\"contractIUniswapExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapOneFactory\",\"outputs\":[{\"internalType\":\"contractIUniswapFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ArbitrageBin is the compiled bytecode used for deploying new contracts.
var ArbitrageBin = "0x60806040523480156200001157600080fd5b507324a42fd28c976a61df5d00d0599c34c4f90748c8600062000039620003ec60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505073c0a47dfe034b400b47bdad5fecda2621de6c4d95600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306f2bf62736b175474e89094c44da98b954eedeac495271d0f6040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156200022457600080fd5b505afa15801562000239573d6000803e3d6000fd5b505050506040513d60208110156200025057600080fd5b8101908080519060200190929190505050905080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073818e6fecd516ecc3849daf6845e3ec868087b755600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200036457600080fd5b505afa15801562000379573d6000803e3d6000fd5b505050506040513d60208110156200039057600080fd5b8101908080519060200190929190505050905080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620003f4565b600033905090565b61268280620004046000396000f3fe6080604052600436106101025760003560e01c80638b698f9011610095578063b5607acf11610064578063b5607acf14610524578063c72c4d101461057b578063ec556889146105d2578063ee87255814610629578063f2fde38b146106e357610109565b80638b698f90146103c85780638da5cb5b1461041f5780639d76ea5814610476578063a59a9973146104cd57610109565b806351cff8d9116100d157806351cff8d91461026a578063575023e8146102bb5780635d47135c14610316578063715018a6146103b157610109565b80630d764fbe1461010e57806311ff70621461016557806320fee44b146101bc5780632a4c0a1a1461021357610109565b3661010957005b600080fd5b34801561011a57600080fd5b50610123610734565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017157600080fd5b5061017a61075a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101c857600080fd5b506101d1610780565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561021f57600080fd5b506102286107a6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027657600080fd5b506102b96004803603602081101561028d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107be565b005b3480156102c757600080fd5b50610314600480360360408110156102de57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a78565b005b34801561032257600080fd5b506103af6004803603608081101561033957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eff565b005b3480156103bd57600080fd5b506103c66115c0565b005b3480156103d457600080fd5b506103dd611748565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561042b57600080fd5b50610434611760565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561048257600080fd5b5061048b611789565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104d957600080fd5b506104e26117af565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561053057600080fd5b506105396117d5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561058757600080fd5b506105906117ed565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105de57600080fd5b506105e7611813565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561063557600080fd5b506106e16004803603608081101561064c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561069d57600080fd5b8201836020820111156106af57600080fd5b803590602001918460018302840111640100000000831117156106d157600080fd5b9091929391929390505050611839565b005b3480156106ef57600080fd5b506107326004803603602081101561070657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d08565b005b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b736b175474e89094c44da98b954eedeac495271d0f81565b6107c6611f15565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610887576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156109295760003090508073ffffffffffffffffffffffffffffffffffffffff163191503373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610922573d6000803e3d6000fd5b5050610a0f565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156109a657600080fd5b505afa1580156109ba573d6000803e3d6000fd5b505050506040513d60208110156109d057600080fd5b81019080805190602001909291905050509050610a0e33828473ffffffffffffffffffffffffffffffffffffffff16611f1d9092919063ffffffff16565b5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9207361cc2a04b9c7a06691df1eb87c6a63957ae88bf01d0d18c81e3d1272099836040518082815260200191505060405180910390a35050565b610a80611f15565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b41576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b606060405180602001604052806000815250905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000736b175474e89094c44da98b954eedeac495271d0f9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635cffe9de30736b175474e89094c44da98b954eedeac495271d0f87866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610cbe578082015181840152602081019050610ca3565b50505050905090810190601f168015610ceb5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610d0d57600080fd5b505af1158015610d21573d6000803e3d6000fd5b5050505060008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610da457600080fd5b505afa158015610db8573d6000803e3d6000fd5b505050506040513d6020811015610dce57600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610e6857600080fd5b505af1158015610e7c573d6000803e3d6000fd5b505050506040513d6020811015610e9257600080fd5b8101908080519060200190929190505050610ef8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806126016022913960400191505060405180910390fd5b5050505050565b60008473ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610fbc57600080fd5b505af1158015610fd0573d6000803e3d6000fd5b505050506040513d6020811015610fe657600080fd5b810190808051906020019092919050505061100057600080fd5b8473ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156110aa57600080fd5b505af11580156110be573d6000803e3d6000fd5b505050506040513d60208110156110d457600080fd5b81019080805190602001909291905050506110ee57600080fd5b8473ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16866040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561119757600080fd5b505af11580156111ab573d6000803e3d6000fd5b505050506040513d60208110156111c157600080fd5b81019080805190602001909291905050506111db57600080fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b1580156112c957600080fd5b505afa1580156112dd573d6000803e3d6000fd5b505050506040513d60408110156112f357600080fd5b81019080805190602001909291908051906020019092919050505050809150506000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637409e2eb878787866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050602060405180830381600087803b1580156113fa57600080fd5b505af115801561140e573d6000803e3d6000fd5b505050506040513d602081101561142457600080fd5b810190808051906020019092919050505090508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156114be57600080fd5b505af11580156114d2573d6000803e3d6000fd5b505050506040513d60208110156114e857600080fd5b810190808051906020019092919050505061150257600080fd5b3373ffffffffffffffffffffffffffffffffffffffff167fffebebfb273923089a3ed6bac0fd4686ac740307859becadeb82f998e30db614878684604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a2505050505050565b6115c8611f15565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611689576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b73818e6fecd516ecc3849daf6845e3ec868087b75581565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73c0a47dfe034b400b47bdad5fecda2621de6c4d9581565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000611843611fd5565b90506000736b175474e89094c44da98b954eedeac495271d0f90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508173ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16896040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561192e57600080fd5b505af1158015611942573d6000803e3d6000fd5b505050506040513d602081101561195857600080fd5b81019080805190602001909291905050506119db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f436f756c64206e6f7420617070726f7665204441492073656c6c00000000000081525060200191505060405180910390fd5b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ddf7e1a78960018088600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b8152600401808681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200195505050505050602060405180830381600087803b158015611ac157600080fd5b505af1158015611ad5573d6000803e3d6000fd5b505050506040513d6020811015611aeb57600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663095ea7b3600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611ba757600080fd5b505af1158015611bbb573d6000803e3d6000fd5b505050506040513d6020811015611bd157600080fd5b8101908080519060200190929190505050611c54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f436f756c64206e6f7420617070726f766520746f6b656e2073656c6c0000000081525060200191505060405180910390fd5b611c6082828530610eff565b60004790506000611c7a898b611fe190919063ffffffff16565b9050808211611cf1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f446964206e6f742070726f66697400000000000000000000000000000000000081525060200191505060405180910390fd5b611cfb8b82612069565b5050505050505050505050565b611d10611f15565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611dd1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611e57576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806125db6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b611fd08363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612120565b505050565b6000610bb84201905090565b60008082840190508381101561205f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed6ff7606040518163ffffffff1660e01b815260040160206040518083038186803b1580156120d357600080fd5b505afa1580156120e7573d6000803e3d6000fd5b505050506040513d60208110156120fd57600080fd5b8101908080519060200190929190505050905061211b81848461220f565b505050565b6060612182826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166123719092919063ffffffff16565b905060008151111561220a578080602001905160208110156121a357600080fd5b8101908080519060200190929190505050612209576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612623602a913960400191505060405180910390fd5b5b505050565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156123405760008373ffffffffffffffffffffffffffffffffffffffff168260405180600001905060006040518083038185875af1925050503d80600081146122b7576040519150601f19603f3d011682016040523d82523d6000602084013e6122bc565b606091505b50509050600115158115151461233a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f436f756c646e2774207472616e7366657220455448000000000000000000000081525060200191505060405180910390fd5b5061236c565b61236b83828473ffffffffffffffffffffffffffffffffffffffff16611f1d9092919063ffffffff16565b5b505050565b60606123808484600085612389565b90509392505050565b60606123948561258f565b612406576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b602083106124565780518252602082019150602081019050602083039250612433565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146124b8576040519150601f19603f3d011682016040523d82523d6000602084013e6124bd565b606091505b509150915081156124d2578092505050612587565b6000815111156124e55780518082602001fd5b836040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561254c578082015181840152602081019050612531565b50505050905090810190601f1680156125795780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b949350505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f91508082141580156125d157506000801b8214155b9250505091905056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373436f756c64206e6f74207472616e73666572206261636b207468652070726f6669745361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220c75a6e80c226a204bdc823bcf9a1cd2a9e609d5b72d3f0094c3fbf17ba31b79964736f6c634300060a0033"

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

// ExecSwap is a paid mutator transaction binding the contract method 0x5d47135c.
//
// Solidity: function execSwap(address srcToken, uint256 srcQty, address destToken, address destAddress) returns()
func (_Arbitrage *ArbitrageTransactor) ExecSwap(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address, destAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "execSwap", srcToken, srcQty, destToken, destAddress)
}

// ExecSwap is a paid mutator transaction binding the contract method 0x5d47135c.
//
// Solidity: function execSwap(address srcToken, uint256 srcQty, address destToken, address destAddress) returns()
func (_Arbitrage *ArbitrageSession) ExecSwap(srcToken common.Address, srcQty *big.Int, destToken common.Address, destAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.ExecSwap(&_Arbitrage.TransactOpts, srcToken, srcQty, destToken, destAddress)
}

// ExecSwap is a paid mutator transaction binding the contract method 0x5d47135c.
//
// Solidity: function execSwap(address srcToken, uint256 srcQty, address destToken, address destAddress) returns()
func (_Arbitrage *ArbitrageTransactorSession) ExecSwap(srcToken common.Address, srcQty *big.Int, destToken common.Address, destAddress common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.ExecSwap(&_Arbitrage.TransactOpts, srcToken, srcQty, destToken, destAddress)
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
// Solidity: function makeArbitrage(uint256 amount, address token) returns()
func (_Arbitrage *ArbitrageTransactor) MakeArbitrage(opts *bind.TransactOpts, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Arbitrage.contract.Transact(opts, "makeArbitrage", amount, token)
}

// MakeArbitrage is a paid mutator transaction binding the contract method 0x575023e8.
//
// Solidity: function makeArbitrage(uint256 amount, address token) returns()
func (_Arbitrage *ArbitrageSession) MakeArbitrage(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.MakeArbitrage(&_Arbitrage.TransactOpts, amount, token)
}

// MakeArbitrage is a paid mutator transaction binding the contract method 0x575023e8.
//
// Solidity: function makeArbitrage(uint256 amount, address token) returns()
func (_Arbitrage *ArbitrageTransactorSession) MakeArbitrage(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Arbitrage.Contract.MakeArbitrage(&_Arbitrage.TransactOpts, amount, token)
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
	SrcToken  common.Address
	DestToken common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xffebebfb273923089a3ed6bac0fd4686ac740307859becadeb82f998e30db614.
//
// Solidity: event Swap(address indexed sender, address srcToken, address destToken, uint256 amount)
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

// WatchSwap is a free log subscription operation binding the contract event 0xffebebfb273923089a3ed6bac0fd4686ac740307859becadeb82f998e30db614.
//
// Solidity: event Swap(address indexed sender, address srcToken, address destToken, uint256 amount)
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

// ParseSwap is a log parse operation binding the contract event 0xffebebfb273923089a3ed6bac0fd4686ac740307859becadeb82f998e30db614.
//
// Solidity: event Swap(address indexed sender, address srcToken, address destToken, uint256 amount)
func (_Arbitrage *ArbitrageFilterer) ParseSwap(log types.Log) (*ArbitrageSwap, error) {
	event := new(ArbitrageSwap)
	if err := _Arbitrage.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	return event, nil
}
