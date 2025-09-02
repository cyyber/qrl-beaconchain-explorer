// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sqrctf1

import (
	"math/big"
	"strings"

	qrl "github.com/theQRL/go-zond"
	"github.com/theQRL/go-zond/accounts/abi"
	"github.com/theQRL/go-zond/accounts/abi/bind"
	"github.com/theQRL/go-zond/common"
	"github.com/theQRL/go-zond/core/types"
	"github.com/theQRL/go-zond/event"
)

// TODO(now.youtrack.cloud/issue/TZB-3)
// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = qrl.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SqrcTf1ABI is the input ABI used to generate the binding from.
const SqrcTf1ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"name\":\"_tokenName\",\"type\":\"string\"},{\"name\":\"_decimalUnits\",\"type\":\"uint8\"},{\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"type\":\"constructor\"},{\"payable\":false,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SqrcTf1 is an auto generated Go binding around a QRL contract.
type SqrcTf1 struct {
	SqrcTf1Caller     // Read-only binding to the contract
	SqrcTf1Transactor // Write-only binding to the contract
	SqrcTf1Filterer   // Log filterer for contract events
}

// SqrcTf1Caller is an auto generated read-only Go binding around a QRL contract.
type SqrcTf1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqrcTf1Transactor is an auto generated write-only Go binding around a QRL contract.
type SqrcTf1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqrcTf1Filterer is an auto generated log filtering Go binding around a QRL contract events.
type SqrcTf1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqrcTf1Session is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type SqrcTf1Session struct {
	Contract     *SqrcTf1          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SqrcTf1CallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type SqrcTf1CallerSession struct {
	Contract *SqrcTf1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SqrcTf1TransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type SqrcTf1TransactorSession struct {
	Contract     *SqrcTf1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SqrcTf1Raw is an auto generated low-level Go binding around a QRL contract.
type SqrcTf1Raw struct {
	Contract *SqrcTf1 // Generic contract binding to access the raw methods on
}

// SqrcTf1CallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type SqrcTf1CallerRaw struct {
	Contract *SqrcTf1Caller // Generic read-only contract binding to access the raw methods on
}

// SqrcTf1TransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type SqrcTf1TransactorRaw struct {
	Contract *SqrcTf1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSqrcTf1 creates a new instance of SqrcTf1, bound to a specific deployed contract.
func NewSqrcTf1(address common.Address, backend bind.ContractBackend) (*SqrcTf1, error) {
	contract, err := bindSqrcTf1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SqrcTf1{SqrcTf1Caller: SqrcTf1Caller{contract: contract}, SqrcTf1Transactor: SqrcTf1Transactor{contract: contract}, SqrcTf1Filterer: SqrcTf1Filterer{contract: contract}}, nil
}

// NewSqrcTf1Caller creates a new read-only instance of SqrcTf1, bound to a specific deployed contract.
func NewSqrcTf1Caller(address common.Address, caller bind.ContractCaller) (*SqrcTf1Caller, error) {
	contract, err := bindSqrcTf1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SqrcTf1Caller{contract: contract}, nil
}

// NewSqrcTf1Transactor creates a new write-only instance of SqrcTf1, bound to a specific deployed contract.
func NewSqrcTf1Transactor(address common.Address, transactor bind.ContractTransactor) (*SqrcTf1Transactor, error) {
	contract, err := bindSqrcTf1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SqrcTf1Transactor{contract: contract}, nil
}

// NewSqrcTf1Filterer creates a new log filterer instance of SqrcTf1, bound to a specific deployed contract.
func NewSqrcTf1Filterer(address common.Address, filterer bind.ContractFilterer) (*SqrcTf1Filterer, error) {
	contract, err := bindSqrcTf1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SqrcTf1Filterer{contract: contract}, nil
}

// bindSqrcTf1 binds a generic wrapper to an already deployed contract.
func bindSqrcTf1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SqrcTf1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SqrcTf1 *SqrcTf1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SqrcTf1.Contract.SqrcTf1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SqrcTf1 *SqrcTf1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SqrcTf1.Contract.SqrcTf1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SqrcTf1 *SqrcTf1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SqrcTf1.Contract.SqrcTf1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SqrcTf1 *SqrcTf1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SqrcTf1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SqrcTf1 *SqrcTf1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SqrcTf1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SqrcTf1 *SqrcTf1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SqrcTf1.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_SqrcTf1 *SqrcTf1Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SqrcTf1.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_SqrcTf1 *SqrcTf1Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SqrcTf1.Contract.Allowance(&_SqrcTf1.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 remaining)
func (_SqrcTf1 *SqrcTf1CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SqrcTf1.Contract.Allowance(&_SqrcTf1.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_SqrcTf1 *SqrcTf1Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SqrcTf1.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_SqrcTf1 *SqrcTf1Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SqrcTf1.Contract.BalanceOf(&_SqrcTf1.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_SqrcTf1 *SqrcTf1CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SqrcTf1.Contract.BalanceOf(&_SqrcTf1.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_SqrcTf1 *SqrcTf1Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SqrcTf1.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_SqrcTf1 *SqrcTf1Session) Decimals() (uint8, error) {
	return _SqrcTf1.Contract.Decimals(&_SqrcTf1.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_SqrcTf1 *SqrcTf1CallerSession) Decimals() (uint8, error) {
	return _SqrcTf1.Contract.Decimals(&_SqrcTf1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_SqrcTf1 *SqrcTf1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SqrcTf1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_SqrcTf1 *SqrcTf1Session) Name() (string, error) {
	return _SqrcTf1.Contract.Name(&_SqrcTf1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_SqrcTf1 *SqrcTf1CallerSession) Name() (string, error) {
	return _SqrcTf1.Contract.Name(&_SqrcTf1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_SqrcTf1 *SqrcTf1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SqrcTf1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_SqrcTf1 *SqrcTf1Session) Symbol() (string, error) {
	return _SqrcTf1.Contract.Symbol(&_SqrcTf1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_SqrcTf1 *SqrcTf1CallerSession) Symbol() (string, error) {
	return _SqrcTf1.Contract.Symbol(&_SqrcTf1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_SqrcTf1 *SqrcTf1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SqrcTf1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_SqrcTf1 *SqrcTf1Session) TotalSupply() (*big.Int, error) {
	return _SqrcTf1.Contract.TotalSupply(&_SqrcTf1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_SqrcTf1 *SqrcTf1CallerSession) TotalSupply() (*big.Int, error) {
	return _SqrcTf1.Contract.TotalSupply(&_SqrcTf1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_SqrcTf1 *SqrcTf1Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SqrcTf1.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_SqrcTf1 *SqrcTf1Session) Version() (string, error) {
	return _SqrcTf1.Contract.Version(&_SqrcTf1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_SqrcTf1 *SqrcTf1CallerSession) Version() (string, error) {
	return _SqrcTf1.Contract.Version(&_SqrcTf1.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.Contract.Approve(&_SqrcTf1.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.Contract.Approve(&_SqrcTf1.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool success)
func (_SqrcTf1 *SqrcTf1Transactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _SqrcTf1.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool success)
func (_SqrcTf1 *SqrcTf1Session) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _SqrcTf1.Contract.ApproveAndCall(&_SqrcTf1.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool success)
func (_SqrcTf1 *SqrcTf1TransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _SqrcTf1.Contract.ApproveAndCall(&_SqrcTf1.TransactOpts, _spender, _value, _extraData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.Contract.Transfer(&_SqrcTf1.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.Contract.Transfer(&_SqrcTf1.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.Contract.TransferFrom(&_SqrcTf1.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_SqrcTf1 *SqrcTf1TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SqrcTf1.Contract.TransferFrom(&_SqrcTf1.TransactOpts, _from, _to, _value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SqrcTf1 *SqrcTf1Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SqrcTf1.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SqrcTf1 *SqrcTf1Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SqrcTf1.Contract.Fallback(&_SqrcTf1.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SqrcTf1 *SqrcTf1TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SqrcTf1.Contract.Fallback(&_SqrcTf1.TransactOpts, calldata)
}

// SqrcTf1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SqrcTf1 contract.
type SqrcTf1ApprovalIterator struct {
	Event *SqrcTf1Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log   // Log channel receiving the found contract events
	sub  qrl.Subscription // Subscription for errors, completion and termination
	done bool             // Whether the subscription completed delivering logs
	fail error            // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SqrcTf1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqrcTf1Approval)
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
		it.Event = new(SqrcTf1Approval)
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
func (it *SqrcTf1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqrcTf1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqrcTf1Approval represents a Approval event raised by the SqrcTf1 contract.
type SqrcTf1Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_SqrcTf1 *SqrcTf1Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*SqrcTf1ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _SqrcTf1.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &SqrcTf1ApprovalIterator{contract: _SqrcTf1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_SqrcTf1 *SqrcTf1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SqrcTf1Approval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _SqrcTf1.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqrcTf1Approval)
				if err := _SqrcTf1.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_SqrcTf1 *SqrcTf1Filterer) ParseApproval(log types.Log) (*SqrcTf1Approval, error) {
	event := new(SqrcTf1Approval)
	if err := _SqrcTf1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SqrcTf1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SqrcTf1 contract.
type SqrcTf1TransferIterator struct {
	Event *SqrcTf1Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log   // Log channel receiving the found contract events
	sub  qrl.Subscription // Subscription for errors, completion and termination
	done bool             // Whether the subscription completed delivering logs
	fail error            // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SqrcTf1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqrcTf1Transfer)
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
		it.Event = new(SqrcTf1Transfer)
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
func (it *SqrcTf1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqrcTf1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqrcTf1Transfer represents a Transfer event raised by the SqrcTf1 contract.
type SqrcTf1Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_SqrcTf1 *SqrcTf1Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*SqrcTf1TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SqrcTf1.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &SqrcTf1TransferIterator{contract: _SqrcTf1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_SqrcTf1 *SqrcTf1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SqrcTf1Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SqrcTf1.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqrcTf1Transfer)
				if err := _SqrcTf1.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_SqrcTf1 *SqrcTf1Filterer) ParseTransfer(log types.Log) (*SqrcTf1Transfer, error) {
	event := new(SqrcTf1Transfer)
	if err := _SqrcTf1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
