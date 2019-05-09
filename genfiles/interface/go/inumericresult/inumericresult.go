// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inumericresult

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

// INumericResultABI is the input ABI used to generate the binding from.
const INumericResultABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"}],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_query\",\"type\":\"bytes\"}],\"name\":\"getResult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// INumericResultBin is the compiled bytecode used for deploying new contracts.
const INumericResultBin = `0x`

// DeployINumericResult deploys a new Ethereum contract, binding an instance of INumericResult to it.
func DeployINumericResult(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *INumericResult, error) {
	parsed, err := abi.JSON(strings.NewReader(INumericResultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(INumericResultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &INumericResult{INumericResultCaller: INumericResultCaller{contract: contract}, INumericResultTransactor: INumericResultTransactor{contract: contract}, INumericResultFilterer: INumericResultFilterer{contract: contract}}, nil
}

// INumericResult is an auto generated Go binding around an Ethereum contract.
type INumericResult struct {
	INumericResultCaller     // Read-only binding to the contract
	INumericResultTransactor // Write-only binding to the contract
	INumericResultFilterer   // Log filterer for contract events
}

// INumericResultCaller is an auto generated read-only Go binding around an Ethereum contract.
type INumericResultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INumericResultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INumericResultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INumericResultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INumericResultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INumericResultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INumericResultSession struct {
	Contract     *INumericResult   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INumericResultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INumericResultCallerSession struct {
	Contract *INumericResultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// INumericResultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INumericResultTransactorSession struct {
	Contract     *INumericResultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// INumericResultRaw is an auto generated low-level Go binding around an Ethereum contract.
type INumericResultRaw struct {
	Contract *INumericResult // Generic contract binding to access the raw methods on
}

// INumericResultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INumericResultCallerRaw struct {
	Contract *INumericResultCaller // Generic read-only contract binding to access the raw methods on
}

// INumericResultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INumericResultTransactorRaw struct {
	Contract *INumericResultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINumericResult creates a new instance of INumericResult, bound to a specific deployed contract.
func NewINumericResult(address common.Address, backend bind.ContractBackend) (*INumericResult, error) {
	contract, err := bindINumericResult(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INumericResult{INumericResultCaller: INumericResultCaller{contract: contract}, INumericResultTransactor: INumericResultTransactor{contract: contract}, INumericResultFilterer: INumericResultFilterer{contract: contract}}, nil
}

// NewINumericResultCaller creates a new read-only instance of INumericResult, bound to a specific deployed contract.
func NewINumericResultCaller(address common.Address, caller bind.ContractCaller) (*INumericResultCaller, error) {
	contract, err := bindINumericResult(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INumericResultCaller{contract: contract}, nil
}

// NewINumericResultTransactor creates a new write-only instance of INumericResult, bound to a specific deployed contract.
func NewINumericResultTransactor(address common.Address, transactor bind.ContractTransactor) (*INumericResultTransactor, error) {
	contract, err := bindINumericResult(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INumericResultTransactor{contract: contract}, nil
}

// NewINumericResultFilterer creates a new log filterer instance of INumericResult, bound to a specific deployed contract.
func NewINumericResultFilterer(address common.Address, filterer bind.ContractFilterer) (*INumericResultFilterer, error) {
	contract, err := bindINumericResult(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INumericResultFilterer{contract: contract}, nil
}

// bindINumericResult binds a generic wrapper to an already deployed contract.
func bindINumericResult(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INumericResultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INumericResult *INumericResultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _INumericResult.Contract.INumericResultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INumericResult *INumericResultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INumericResult.Contract.INumericResultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INumericResult *INumericResultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INumericResult.Contract.INumericResultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INumericResult *INumericResultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _INumericResult.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INumericResult *INumericResultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INumericResult.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INumericResult *INumericResultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INumericResult.Contract.contract.Transact(opts, method, params...)
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(uint256)
func (_INumericResult *INumericResultCaller) GetResult(opts *bind.CallOpts, _query []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _INumericResult.contract.Call(opts, out, "getResult", _query)
	return *ret0, err
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(uint256)
func (_INumericResult *INumericResultSession) GetResult(_query []byte) (*big.Int, error) {
	return _INumericResult.Contract.GetResult(&_INumericResult.CallOpts, _query)
}

// GetResult is a free data retrieval call binding the contract method 0xef4592fb.
//
// Solidity: function getResult(bytes _query) constant returns(uint256)
func (_INumericResult *INumericResultCallerSession) GetResult(_query []byte) (*big.Int, error) {
	return _INumericResult.Contract.GetResult(&_INumericResult.CallOpts, _query)
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_INumericResult *INumericResultCaller) IsFinalized(opts *bind.CallOpts, _query []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _INumericResult.contract.Call(opts, out, "isFinalized", _query)
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_INumericResult *INumericResultSession) IsFinalized(_query []byte) (bool, error) {
	return _INumericResult.Contract.IsFinalized(&_INumericResult.CallOpts, _query)
}

// IsFinalized is a free data retrieval call binding the contract method 0xbcdbda94.
//
// Solidity: function isFinalized(bytes _query) constant returns(bool)
func (_INumericResult *INumericResultCallerSession) IsFinalized(_query []byte) (bool, error) {
	return _INumericResult.Contract.IsFinalized(&_INumericResult.CallOpts, _query)
}
