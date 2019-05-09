// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibooleanresult

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IBooleanResultABI is the input ABI used to generate the binding from.
const IBooleanResultABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"}],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"}],\"name\":\"getResult\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IBooleanResultBin is the compiled bytecode used for deploying new contracts.
const IBooleanResultBin = `0x`

// DeployIBooleanResult deploys a new Ethereum contract, binding an instance of IBooleanResult to it.
func DeployIBooleanResult(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IBooleanResult, error) {
	parsed, err := abi.JSON(strings.NewReader(IBooleanResultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IBooleanResultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IBooleanResult{IBooleanResultCaller: IBooleanResultCaller{contract: contract}, IBooleanResultTransactor: IBooleanResultTransactor{contract: contract}, IBooleanResultFilterer: IBooleanResultFilterer{contract: contract}}, nil
}

// IBooleanResult is an auto generated Go binding around an Ethereum contract.
type IBooleanResult struct {
	IBooleanResultCaller     // Read-only binding to the contract
	IBooleanResultTransactor // Write-only binding to the contract
	IBooleanResultFilterer   // Log filterer for contract events
}

// IBooleanResultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBooleanResultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBooleanResultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBooleanResultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBooleanResultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBooleanResultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBooleanResultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBooleanResultSession struct {
	Contract     *IBooleanResult   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBooleanResultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBooleanResultCallerSession struct {
	Contract *IBooleanResultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IBooleanResultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBooleanResultTransactorSession struct {
	Contract     *IBooleanResultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IBooleanResultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBooleanResultRaw struct {
	Contract *IBooleanResult // Generic contract binding to access the raw methods on
}

// IBooleanResultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBooleanResultCallerRaw struct {
	Contract *IBooleanResultCaller // Generic read-only contract binding to access the raw methods on
}

// IBooleanResultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBooleanResultTransactorRaw struct {
	Contract *IBooleanResultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBooleanResult creates a new instance of IBooleanResult, bound to a specific deployed contract.
func NewIBooleanResult(address common.Address, backend bind.ContractBackend) (*IBooleanResult, error) {
	contract, err := bindIBooleanResult(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBooleanResult{IBooleanResultCaller: IBooleanResultCaller{contract: contract}, IBooleanResultTransactor: IBooleanResultTransactor{contract: contract}, IBooleanResultFilterer: IBooleanResultFilterer{contract: contract}}, nil
}

// NewIBooleanResultCaller creates a new read-only instance of IBooleanResult, bound to a specific deployed contract.
func NewIBooleanResultCaller(address common.Address, caller bind.ContractCaller) (*IBooleanResultCaller, error) {
	contract, err := bindIBooleanResult(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBooleanResultCaller{contract: contract}, nil
}

// NewIBooleanResultTransactor creates a new write-only instance of IBooleanResult, bound to a specific deployed contract.
func NewIBooleanResultTransactor(address common.Address, transactor bind.ContractTransactor) (*IBooleanResultTransactor, error) {
	contract, err := bindIBooleanResult(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBooleanResultTransactor{contract: contract}, nil
}

// NewIBooleanResultFilterer creates a new log filterer instance of IBooleanResult, bound to a specific deployed contract.
func NewIBooleanResultFilterer(address common.Address, filterer bind.ContractFilterer) (*IBooleanResultFilterer, error) {
	contract, err := bindIBooleanResult(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBooleanResultFilterer{contract: contract}, nil
}

// bindIBooleanResult binds a generic wrapper to an already deployed contract.
func bindIBooleanResult(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBooleanResultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBooleanResult *IBooleanResultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBooleanResult.Contract.IBooleanResultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBooleanResult *IBooleanResultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBooleanResult.Contract.IBooleanResultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBooleanResult *IBooleanResultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBooleanResult.Contract.IBooleanResultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBooleanResult *IBooleanResultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBooleanResult.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBooleanResult *IBooleanResultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBooleanResult.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBooleanResult *IBooleanResultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBooleanResult.Contract.contract.Transact(opts, method, params...)
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(bool)
func (_IBooleanResult *IBooleanResultCaller) GetResult(opts *bind.CallOpts, _query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IBooleanResult.contract.Call(opts, out, "getResult", _query)
	return *ret0, err
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(bool)
func (_IBooleanResult *IBooleanResultSession) GetResult(_query []byte) (bool, error) {
	return _IBooleanResult.Contract.GetResult(&_IBooleanResult.CallOpts, _query)
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(bool)
func (_IBooleanResult *IBooleanResultCallerSession) GetResult(_query []byte) (bool, error) {
	return _IBooleanResult.Contract.GetResult(&_IBooleanResult.CallOpts, _query)
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_IBooleanResult *IBooleanResultCaller) IsFinalized(opts *bind.CallOpts, _query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IBooleanResult.contract.Call(opts, out, "isFinalized", _query)
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_IBooleanResult *IBooleanResultSession) IsFinalized(_query []byte) (bool, error) {
	return _IBooleanResult.Contract.IsFinalized(&_IBooleanResult.CallOpts, _query)
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_IBooleanResult *IBooleanResultCallerSession) IsFinalized(_query []byte) (bool, error) {
	return _IBooleanResult.Contract.IsFinalized(&_IBooleanResult.CallOpts, _query)
}
