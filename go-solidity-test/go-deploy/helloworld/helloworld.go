// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HelloworldMetaData contains all meta data concerning the Helloworld contract.
var HelloworldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"say\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e3d670d7": "balance(address)",
		"5f515226": "checkBalance(address)",
		"b6b55f25": "deposit(uint256)",
		"954ab4b2": "say()",
		"2e1a7d4d": "withdraw(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610249806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632e1a7d4d1461005c5780635f5152261461007b578063954ab4b2146100b3578063b6b55f2514610130578063e3d670d71461014d575b600080fd5b6100796004803603602081101561007257600080fd5b5035610173565b005b6100a16004803603602081101561009157600080fd5b50356001600160a01b03166101a4565b60408051918252519081900360200190f35b6100bb6101bf565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f55781810151838201526020016100dd565b50505050905090810190601f1680156101225780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100796004803603602081101561014657600080fd5b50356101e9565b6100a16004803603602081101561016357600080fd5b50356001600160a01b0316610201565b336000908152602081905260409020548110156101a157336000908152602081905260409020805482900390555b50565b6001600160a01b031660009081526020819052604090205490565b60408051808201909152601081526f1a195b1b1bc8195d1a195c9ddbdc9b1960821b602082015290565b33600090815260208190526040902080549091019055565b6000602081905290815260409020548156fea2646970667358221220eeeb397f7fefd0bed7a7388fdfdb21495d61526ed39306e1f9e365e30487dc0a64736f6c63430007060033",
}

// HelloworldABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloworldMetaData.ABI instead.
var HelloworldABI = HelloworldMetaData.ABI

// Deprecated: Use HelloworldMetaData.Sigs instead.
// HelloworldFuncSigs maps the 4-byte function signature to its string representation.
var HelloworldFuncSigs = HelloworldMetaData.Sigs

// HelloworldBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloworldMetaData.Bin instead.
var HelloworldBin = HelloworldMetaData.Bin

// DeployHelloworld deploys a new Ethereum contract, binding an instance of Helloworld to it.
func DeployHelloworld(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Helloworld, error) {
	parsed, err := HelloworldMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloworldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// Helloworld is an auto generated Go binding around an Ethereum contract.
type Helloworld struct {
	HelloworldCaller     // Read-only binding to the contract
	HelloworldTransactor // Write-only binding to the contract
	HelloworldFilterer   // Log filterer for contract events
}

// HelloworldCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloworldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloworldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloworldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloworldSession struct {
	Contract     *Helloworld       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloworldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloworldCallerSession struct {
	Contract *HelloworldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloworldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloworldTransactorSession struct {
	Contract     *HelloworldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloworldRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloworldRaw struct {
	Contract *Helloworld // Generic contract binding to access the raw methods on
}

// HelloworldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloworldCallerRaw struct {
	Contract *HelloworldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloworldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloworldTransactorRaw struct {
	Contract *HelloworldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloworld creates a new instance of Helloworld, bound to a specific deployed contract.
func NewHelloworld(address common.Address, backend bind.ContractBackend) (*Helloworld, error) {
	contract, err := bindHelloworld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// NewHelloworldCaller creates a new read-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldCaller(address common.Address, caller bind.ContractCaller) (*HelloworldCaller, error) {
	contract, err := bindHelloworld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldCaller{contract: contract}, nil
}

// NewHelloworldTransactor creates a new write-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloworldTransactor, error) {
	contract, err := bindHelloworld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldTransactor{contract: contract}, nil
}

// NewHelloworldFilterer creates a new log filterer instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloworldFilterer, error) {
	contract, err := bindHelloworld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloworldFilterer{contract: contract}, nil
}

// bindHelloworld binds a generic wrapper to an already deployed contract.
func bindHelloworld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloworldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.HelloworldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_Helloworld *HelloworldCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "balance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_Helloworld *HelloworldSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _Helloworld.Contract.Balance(&_Helloworld.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_Helloworld *HelloworldCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _Helloworld.Contract.Balance(&_Helloworld.CallOpts, arg0)
}

// CheckBalance is a free data retrieval call binding the contract method 0x5f515226.
//
// Solidity: function checkBalance(address account) view returns(uint256)
func (_Helloworld *HelloworldCaller) CheckBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "checkBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckBalance is a free data retrieval call binding the contract method 0x5f515226.
//
// Solidity: function checkBalance(address account) view returns(uint256)
func (_Helloworld *HelloworldSession) CheckBalance(account common.Address) (*big.Int, error) {
	return _Helloworld.Contract.CheckBalance(&_Helloworld.CallOpts, account)
}

// CheckBalance is a free data retrieval call binding the contract method 0x5f515226.
//
// Solidity: function checkBalance(address account) view returns(uint256)
func (_Helloworld *HelloworldCallerSession) CheckBalance(account common.Address) (*big.Int, error) {
	return _Helloworld.Contract.CheckBalance(&_Helloworld.CallOpts, account)
}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Helloworld *HelloworldCaller) Say(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "say")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Helloworld *HelloworldSession) Say() (string, error) {
	return _Helloworld.Contract.Say(&_Helloworld.CallOpts)
}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Helloworld *HelloworldCallerSession) Say() (string, error) {
	return _Helloworld.Contract.Say(&_Helloworld.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Helloworld *HelloworldTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Helloworld *HelloworldSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Helloworld.Contract.Deposit(&_Helloworld.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Helloworld *HelloworldTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Helloworld.Contract.Deposit(&_Helloworld.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Helloworld *HelloworldTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Helloworld *HelloworldSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Helloworld.Contract.Withdraw(&_Helloworld.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Helloworld *HelloworldTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Helloworld.Contract.Withdraw(&_Helloworld.TransactOpts, amount)
}
