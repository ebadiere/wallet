// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap2

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

// Uniswap2ABI is the input ABI used to generate the binding from.
const Uniswap2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockTimestampLast\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"consult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0Average\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1Average\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Uniswap2Bin is the compiled bytecode used for deploying new contracts.
var Uniswap2Bin = "0x60e06040523480156200001157600080fd5b506040516200187838038062001878833981810160405260608110156200003757600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050600062000076848484620004b060201b6200082c1760201c565b90508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015620000f657600080fd5b505afa1580156200010b573d6000803e3d6000fd5b505050506040513d60208110156200012257600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b158015620001b057600080fd5b505afa158015620001c5573d6000803e3d6000fd5b505050506040513d6020811015620001dc57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff16635909c0d56040518163ffffffff1660e01b815260040160206040518083038186803b1580156200026a57600080fd5b505afa1580156200027f573d6000803e3d6000fd5b505050506040513d60208110156200029657600080fd5b81019080805190602001909291905050506000819055508073ffffffffffffffffffffffffffffffffffffffff16635a3d54936040518163ffffffff1660e01b815260040160206040518083038186803b158015620002f457600080fd5b505afa15801562000309573d6000803e3d6000fd5b505050506040513d60208110156200032057600080fd5b81019080805190602001909291905050506001819055506000808273ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b1580156200038157600080fd5b505afa15801562000396573d6000803e3d6000fd5b505050506040513d6060811015620003ad57600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050600260008291906101000a81548163ffffffff021916908363ffffffff1602179055508193508294505050506000826dffffffffffffffffffffffffffff16141580156200043157506000816dffffffffffffffffffffffffffff1614155b620004a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4578616d706c654f7261636c6553696d706c653a204e4f5f524553455256455381525060200191505060405180910390fd5b5050505050506200078f565b6000806000620004c785856200061360201b60201c565b91509150858282604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012060405160200180807fff000000000000000000000000000000000000000000000000000000000000008152506001018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401828152602001807f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f815250602001925050506040516020818303038152906040528051906020012060001c925050509392505050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156200069d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180620018536025913960400191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610620006d9578284620006dc565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000788576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f556e697377617056324c6962726172793a205a45524f5f41444452455353000081525060200191505060405180910390fd5b9250929050565b60805160601c60a05160601c60c05160601c611084620007cf60003980610407528061080a5250806102dd52806103035250806105a552506110846000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063a2e6204511610066578063a2e62045146101e5578063a6bb4539146101ef578063b4d1d79514610249578063c5700a0214610267578063d21220a7146102915761009e565b80630dfe1681146100a35780633ddac953146100ed5780635909c0d51461014f5780635a3d54931461016d5780635e6aaf2c1461018b575b600080fd5b6100ab6102db565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101396004803603604081101561010357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506102ff565b6040518082815260200191505060405180910390f35b61015761055b565b6040518082815260200191505060405180910390f35b610175610561565b6040518082815260200191505060405180910390f35b610193610567565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101ed61059b565b005b6101f76107b7565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102516107eb565b6040518082815260200191505060405180910390f35b61026f6107f2565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b610299610808565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610405576103ea6103e58360036040518060200160405290816000820160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152505061098790919063ffffffff16565b610a5d565b71ffffffffffffffffffffffffffffffffffff169050610555565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146104a9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610fe36022913960400191505060405180910390fd5b61053e6105398360046040518060200160405290816000820160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152505061098790919063ffffffff16565b610a5d565b71ffffffffffffffffffffffffffffffffffff1690505b92915050565b60005481565b60015481565b60048060000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905081565b60008060006105c97f0000000000000000000000000000000000000000000000000000000000000000610a72565b9250925092506000600260009054906101000a900463ffffffff1682039050620151808163ffffffff16101561064a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806110056027913960400191505060405180910390fd5b60405180602001604052808263ffffffff1660005487038161066857fe5b047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815250600360008201518160000160006101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16021790555090505060405180602001604052808263ffffffff1660015486038161070457fe5b047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815250600460008201518160000160006101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550905050836000819055508260018190555081600260006101000a81548163ffffffff021916908363ffffffff16021790555050505050565b60038060000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905081565b6201518081565b600260009054906101000a900463ffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600080600061083b8585610cbd565b91509150858282604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012060405160200180807fff000000000000000000000000000000000000000000000000000000000000008152506001018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401828152602001807f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f815250602001925050506040516020818303038152906040528051906020012060001c925050509392505050565b61098f610f79565b6000808314806109f0575083600001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16838486600001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1602925082816109ed57fe5b04145b610a45576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061102c6023913960400191505060405180910390fd5b60405180602001604052808281525091505092915050565b6000607060ff168260000151901c9050919050565b6000806000610a7f610e34565b90508373ffffffffffffffffffffffffffffffffffffffff16635909c0d56040518163ffffffff1660e01b815260040160206040518083038186803b158015610ac757600080fd5b505afa158015610adb573d6000803e3d6000fd5b505050506040513d6020811015610af157600080fd5b810190808051906020019092919050505092508373ffffffffffffffffffffffffffffffffffffffff16635a3d54936040518163ffffffff1660e01b815260040160206040518083038186803b158015610b4a57600080fd5b505afa158015610b5e573d6000803e3d6000fd5b505050506040513d6020811015610b7457600080fd5b8101908080519060200190929190505050915060008060008673ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015610bd257600080fd5b505afa158015610be6573d6000803e3d6000fd5b505050506040513d6060811015610bfc57600080fd5b810190808051906020019092919080519060200190929190805190602001909291905050509250925092508363ffffffff168163ffffffff1614610cb357600081850390508063ffffffff16610c528486610e4a565b600001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1602870196508063ffffffff16610c8a8585610e4a565b600001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff160286019550505b5050509193909250565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610d45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610fbe6025913960400191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610d7f578284610d82565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f556e697377617056324c6962726172793a205a45524f5f41444452455353000081525060200191505060405180910390fd5b9250929050565b60006401000000004281610e4457fe5b06905090565b610e52610f8c565b6000826dffffffffffffffffffffffffffff1611610ed8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4669786564506f696e743a204449565f42595f5a45524f00000000000000000081525060200191505060405180910390fd5b6040518060200160405280836dffffffffffffffffffffffffffff16607060ff16866dffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16901b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681610f4f57fe5b047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815250905092915050565b6040518060200160405280600081525090565b604051806020016040528060007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152509056fe556e697377617056324c6962726172793a204944454e544943414c5f4144445245535345534578616d706c654f7261636c6553696d706c653a20494e56414c49445f544f4b454e4578616d706c654f7261636c6553696d706c653a20504552494f445f4e4f545f454c41505345444669786564506f696e743a204d554c5449504c49434154494f4e5f4f564552464c4f57a2646970667358221220c00bea712636c9ee4f1cb5777ae88dc2499b38113f5708e9ddeead7725f3556664736f6c634300060a0033556e697377617056324c6962726172793a204944454e544943414c5f414444524553534553"

// DeployUniswap2 deploys a new Ethereum contract, binding an instance of Uniswap2 to it.
func DeployUniswap2(auth *bind.TransactOpts, backend bind.ContractBackend, factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, *types.Transaction, *Uniswap2, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswap2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Uniswap2Bin), backend, factory, tokenA, tokenB)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Uniswap2{Uniswap2Caller: Uniswap2Caller{contract: contract}, Uniswap2Transactor: Uniswap2Transactor{contract: contract}, Uniswap2Filterer: Uniswap2Filterer{contract: contract}}, nil
}

// Uniswap2 is an auto generated Go binding around an Ethereum contract.
type Uniswap2 struct {
	Uniswap2Caller     // Read-only binding to the contract
	Uniswap2Transactor // Write-only binding to the contract
	Uniswap2Filterer   // Log filterer for contract events
}

// Uniswap2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswap2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswap2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswap2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswap2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswap2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswap2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswap2Session struct {
	Contract     *Uniswap2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswap2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswap2CallerSession struct {
	Contract *Uniswap2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Uniswap2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswap2TransactorSession struct {
	Contract     *Uniswap2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Uniswap2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswap2Raw struct {
	Contract *Uniswap2 // Generic contract binding to access the raw methods on
}

// Uniswap2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswap2CallerRaw struct {
	Contract *Uniswap2Caller // Generic read-only contract binding to access the raw methods on
}

// Uniswap2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswap2TransactorRaw struct {
	Contract *Uniswap2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswap2 creates a new instance of Uniswap2, bound to a specific deployed contract.
func NewUniswap2(address common.Address, backend bind.ContractBackend) (*Uniswap2, error) {
	contract, err := bindUniswap2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswap2{Uniswap2Caller: Uniswap2Caller{contract: contract}, Uniswap2Transactor: Uniswap2Transactor{contract: contract}, Uniswap2Filterer: Uniswap2Filterer{contract: contract}}, nil
}

// NewUniswap2Caller creates a new read-only instance of Uniswap2, bound to a specific deployed contract.
func NewUniswap2Caller(address common.Address, caller bind.ContractCaller) (*Uniswap2Caller, error) {
	contract, err := bindUniswap2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswap2Caller{contract: contract}, nil
}

// NewUniswap2Transactor creates a new write-only instance of Uniswap2, bound to a specific deployed contract.
func NewUniswap2Transactor(address common.Address, transactor bind.ContractTransactor) (*Uniswap2Transactor, error) {
	contract, err := bindUniswap2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswap2Transactor{contract: contract}, nil
}

// NewUniswap2Filterer creates a new log filterer instance of Uniswap2, bound to a specific deployed contract.
func NewUniswap2Filterer(address common.Address, filterer bind.ContractFilterer) (*Uniswap2Filterer, error) {
	contract, err := bindUniswap2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswap2Filterer{contract: contract}, nil
}

// bindUniswap2 binds a generic wrapper to an already deployed contract.
func bindUniswap2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswap2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap2 *Uniswap2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswap2.Contract.Uniswap2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap2 *Uniswap2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap2.Contract.Uniswap2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap2 *Uniswap2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap2.Contract.Uniswap2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap2 *Uniswap2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Uniswap2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap2 *Uniswap2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap2 *Uniswap2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap2.Contract.contract.Transact(opts, method, params...)
}

// PERIOD is a free data retrieval call binding the contract method 0xb4d1d795.
//
// Solidity: function PERIOD() view returns(uint256)
func (_Uniswap2 *Uniswap2Caller) PERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "PERIOD")
	return *ret0, err
}

// PERIOD is a free data retrieval call binding the contract method 0xb4d1d795.
//
// Solidity: function PERIOD() view returns(uint256)
func (_Uniswap2 *Uniswap2Session) PERIOD() (*big.Int, error) {
	return _Uniswap2.Contract.PERIOD(&_Uniswap2.CallOpts)
}

// PERIOD is a free data retrieval call binding the contract method 0xb4d1d795.
//
// Solidity: function PERIOD() view returns(uint256)
func (_Uniswap2 *Uniswap2CallerSession) PERIOD() (*big.Int, error) {
	return _Uniswap2.Contract.PERIOD(&_Uniswap2.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_Uniswap2 *Uniswap2Caller) BlockTimestampLast(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "blockTimestampLast")
	return *ret0, err
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_Uniswap2 *Uniswap2Session) BlockTimestampLast() (uint32, error) {
	return _Uniswap2.Contract.BlockTimestampLast(&_Uniswap2.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_Uniswap2 *Uniswap2CallerSession) BlockTimestampLast() (uint32, error) {
	return _Uniswap2.Contract.BlockTimestampLast(&_Uniswap2.CallOpts)
}

// Consult is a free data retrieval call binding the contract method 0x3ddac953.
//
// Solidity: function consult(address token, uint256 amountIn) view returns(uint256 amountOut)
func (_Uniswap2 *Uniswap2Caller) Consult(opts *bind.CallOpts, token common.Address, amountIn *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "consult", token, amountIn)
	return *ret0, err
}

// Consult is a free data retrieval call binding the contract method 0x3ddac953.
//
// Solidity: function consult(address token, uint256 amountIn) view returns(uint256 amountOut)
func (_Uniswap2 *Uniswap2Session) Consult(token common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Uniswap2.Contract.Consult(&_Uniswap2.CallOpts, token, amountIn)
}

// Consult is a free data retrieval call binding the contract method 0x3ddac953.
//
// Solidity: function consult(address token, uint256 amountIn) view returns(uint256 amountOut)
func (_Uniswap2 *Uniswap2CallerSession) Consult(token common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Uniswap2.Contract.Consult(&_Uniswap2.CallOpts, token, amountIn)
}

// Price0Average is a free data retrieval call binding the contract method 0xa6bb4539.
//
// Solidity: function price0Average() view returns(uint224 _x)
func (_Uniswap2 *Uniswap2Caller) Price0Average(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "price0Average")
	return *ret0, err
}

// Price0Average is a free data retrieval call binding the contract method 0xa6bb4539.
//
// Solidity: function price0Average() view returns(uint224 _x)
func (_Uniswap2 *Uniswap2Session) Price0Average() (*big.Int, error) {
	return _Uniswap2.Contract.Price0Average(&_Uniswap2.CallOpts)
}

// Price0Average is a free data retrieval call binding the contract method 0xa6bb4539.
//
// Solidity: function price0Average() view returns(uint224 _x)
func (_Uniswap2 *Uniswap2CallerSession) Price0Average() (*big.Int, error) {
	return _Uniswap2.Contract.Price0Average(&_Uniswap2.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Uniswap2 *Uniswap2Caller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "price0CumulativeLast")
	return *ret0, err
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Uniswap2 *Uniswap2Session) Price0CumulativeLast() (*big.Int, error) {
	return _Uniswap2.Contract.Price0CumulativeLast(&_Uniswap2.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Uniswap2 *Uniswap2CallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Uniswap2.Contract.Price0CumulativeLast(&_Uniswap2.CallOpts)
}

// Price1Average is a free data retrieval call binding the contract method 0x5e6aaf2c.
//
// Solidity: function price1Average() view returns(uint224 _x)
func (_Uniswap2 *Uniswap2Caller) Price1Average(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "price1Average")
	return *ret0, err
}

// Price1Average is a free data retrieval call binding the contract method 0x5e6aaf2c.
//
// Solidity: function price1Average() view returns(uint224 _x)
func (_Uniswap2 *Uniswap2Session) Price1Average() (*big.Int, error) {
	return _Uniswap2.Contract.Price1Average(&_Uniswap2.CallOpts)
}

// Price1Average is a free data retrieval call binding the contract method 0x5e6aaf2c.
//
// Solidity: function price1Average() view returns(uint224 _x)
func (_Uniswap2 *Uniswap2CallerSession) Price1Average() (*big.Int, error) {
	return _Uniswap2.Contract.Price1Average(&_Uniswap2.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Uniswap2 *Uniswap2Caller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "price1CumulativeLast")
	return *ret0, err
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Uniswap2 *Uniswap2Session) Price1CumulativeLast() (*big.Int, error) {
	return _Uniswap2.Contract.Price1CumulativeLast(&_Uniswap2.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Uniswap2 *Uniswap2CallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Uniswap2.Contract.Price1CumulativeLast(&_Uniswap2.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uniswap2 *Uniswap2Caller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "token0")
	return *ret0, err
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uniswap2 *Uniswap2Session) Token0() (common.Address, error) {
	return _Uniswap2.Contract.Token0(&_Uniswap2.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uniswap2 *Uniswap2CallerSession) Token0() (common.Address, error) {
	return _Uniswap2.Contract.Token0(&_Uniswap2.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uniswap2 *Uniswap2Caller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Uniswap2.contract.Call(opts, out, "token1")
	return *ret0, err
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uniswap2 *Uniswap2Session) Token1() (common.Address, error) {
	return _Uniswap2.Contract.Token1(&_Uniswap2.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uniswap2 *Uniswap2CallerSession) Token1() (common.Address, error) {
	return _Uniswap2.Contract.Token1(&_Uniswap2.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_Uniswap2 *Uniswap2Transactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap2.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_Uniswap2 *Uniswap2Session) Update() (*types.Transaction, error) {
	return _Uniswap2.Contract.Update(&_Uniswap2.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_Uniswap2 *Uniswap2TransactorSession) Update() (*types.Transaction, error) {
	return _Uniswap2.Contract.Update(&_Uniswap2.TransactOpts)
}
