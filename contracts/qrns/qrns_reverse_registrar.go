// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qrns

import (
	"errors"
	"math/big"
	"strings"

	qrl "github.com/theQRL/go-zond"
	"github.com/theQRL/go-zond/accounts/abi"
	"github.com/theQRL/go-zond/accounts/abi/bind"
	"github.com/theQRL/go-zond/common"
	"github.com/theQRL/go-zond/core/types"
	"github.com/theQRL/go-zond/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = qrl.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// QRNSReverseRegistrarMetaData contains all meta data concerning the QRNSReverseRegistrar contract.
var QRNSReverseRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"qrnsAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ControllerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractNameResolver\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"DefaultResolverChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ReverseClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimForAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimWithResolver\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultResolver\",\"outputs\":[{\"internalType\":\"contractNameResolver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qrns\",\"outputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"node\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setDefaultResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setNameForAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QRNSReverseRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use QRNSReverseRegistrarMetaData.ABI instead.
var QRNSReverseRegistrarABI = QRNSReverseRegistrarMetaData.ABI

// QRNSReverseRegistrar is an auto generated Go binding around a QRL contract.
type QRNSReverseRegistrar struct {
	QRNSReverseRegistrarCaller     // Read-only binding to the contract
	QRNSReverseRegistrarTransactor // Write-only binding to the contract
	QRNSReverseRegistrarFilterer   // Log filterer for contract events
}

// QRNSReverseRegistrarCaller is an auto generated read-only Go binding around a QRL contract.
type QRNSReverseRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSReverseRegistrarTransactor is an auto generated write-only Go binding around a QRL contract.
type QRNSReverseRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSReverseRegistrarFilterer is an auto generated log filtering Go binding around a QRL contract events.
type QRNSReverseRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSReverseRegistrarSession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type QRNSReverseRegistrarSession struct {
	Contract     *QRNSReverseRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// QRNSReverseRegistrarCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type QRNSReverseRegistrarCallerSession struct {
	Contract *QRNSReverseRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// QRNSReverseRegistrarTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type QRNSReverseRegistrarTransactorSession struct {
	Contract     *QRNSReverseRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// QRNSReverseRegistrarRaw is an auto generated low-level Go binding around a QRL contract.
type QRNSReverseRegistrarRaw struct {
	Contract *QRNSReverseRegistrar // Generic contract binding to access the raw methods on
}

// QRNSReverseRegistrarCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type QRNSReverseRegistrarCallerRaw struct {
	Contract *QRNSReverseRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// QRNSReverseRegistrarTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type QRNSReverseRegistrarTransactorRaw struct {
	Contract *QRNSReverseRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQRNSReverseRegistrar creates a new instance of QRNSReverseRegistrar, bound to a specific deployed contract.
func NewQRNSReverseRegistrar(address common.Address, backend bind.ContractBackend) (*QRNSReverseRegistrar, error) {
	contract, err := bindQRNSReverseRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrar{QRNSReverseRegistrarCaller: QRNSReverseRegistrarCaller{contract: contract}, QRNSReverseRegistrarTransactor: QRNSReverseRegistrarTransactor{contract: contract}, QRNSReverseRegistrarFilterer: QRNSReverseRegistrarFilterer{contract: contract}}, nil
}

// NewQRNSReverseRegistrarCaller creates a new read-only instance of QRNSReverseRegistrar, bound to a specific deployed contract.
func NewQRNSReverseRegistrarCaller(address common.Address, caller bind.ContractCaller) (*QRNSReverseRegistrarCaller, error) {
	contract, err := bindQRNSReverseRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrarCaller{contract: contract}, nil
}

// NewQRNSReverseRegistrarTransactor creates a new write-only instance of QRNSReverseRegistrar, bound to a specific deployed contract.
func NewQRNSReverseRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*QRNSReverseRegistrarTransactor, error) {
	contract, err := bindQRNSReverseRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrarTransactor{contract: contract}, nil
}

// NewQRNSReverseRegistrarFilterer creates a new log filterer instance of QRNSReverseRegistrar, bound to a specific deployed contract.
func NewQRNSReverseRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*QRNSReverseRegistrarFilterer, error) {
	contract, err := bindQRNSReverseRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrarFilterer{contract: contract}, nil
}

// bindQRNSReverseRegistrar binds a generic wrapper to an already deployed contract.
func bindQRNSReverseRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QRNSReverseRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSReverseRegistrar *QRNSReverseRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSReverseRegistrar.Contract.QRNSReverseRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSReverseRegistrar *QRNSReverseRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.QRNSReverseRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSReverseRegistrar *QRNSReverseRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.QRNSReverseRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSReverseRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSReverseRegistrar.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) Controllers(arg0 common.Address) (bool, error) {
	return _QRNSReverseRegistrar.Contract.Controllers(&_QRNSReverseRegistrar.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _QRNSReverseRegistrar.Contract.Controllers(&_QRNSReverseRegistrar.CallOpts, arg0)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCaller) DefaultResolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSReverseRegistrar.contract.Call(opts, &out, "defaultResolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) DefaultResolver() (common.Address, error) {
	return _QRNSReverseRegistrar.Contract.DefaultResolver(&_QRNSReverseRegistrar.CallOpts)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCallerSession) DefaultResolver() (common.Address, error) {
	return _QRNSReverseRegistrar.Contract.DefaultResolver(&_QRNSReverseRegistrar.CallOpts)
}

// Qrns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function qrns() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCaller) Qrns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSReverseRegistrar.contract.Call(opts, &out, "qrns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Qrns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function qrns() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) Qrns() (common.Address, error) {
	return _QRNSReverseRegistrar.Contract.Qrns(&_QRNSReverseRegistrar.CallOpts)
}

// Qrns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function qrns() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCallerSession) Qrns() (common.Address, error) {
	return _QRNSReverseRegistrar.Contract.Qrns(&_QRNSReverseRegistrar.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCaller) Node(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _QRNSReverseRegistrar.contract.Call(opts, &out, "node", addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) Node(addr common.Address) ([32]byte, error) {
	return _QRNSReverseRegistrar.Contract.Node(&_QRNSReverseRegistrar.CallOpts, addr)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCallerSession) Node(addr common.Address) ([32]byte, error) {
	return _QRNSReverseRegistrar.Contract.Node(&_QRNSReverseRegistrar.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSReverseRegistrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) Owner() (common.Address, error) {
	return _QRNSReverseRegistrar.Contract.Owner(&_QRNSReverseRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarCallerSession) Owner() (common.Address, error) {
	return _QRNSReverseRegistrar.Contract.Owner(&_QRNSReverseRegistrar.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) Claim(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "claim", owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.Claim(&_QRNSReverseRegistrar.TransactOpts, owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.Claim(&_QRNSReverseRegistrar.TransactOpts, owner)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) ClaimForAddr(opts *bind.TransactOpts, addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "claimForAddr", addr, owner, resolver)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) ClaimForAddr(addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.ClaimForAddr(&_QRNSReverseRegistrar.TransactOpts, addr, owner, resolver)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) ClaimForAddr(addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.ClaimForAddr(&_QRNSReverseRegistrar.TransactOpts, addr, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) ClaimWithResolver(opts *bind.TransactOpts, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "claimWithResolver", owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.ClaimWithResolver(&_QRNSReverseRegistrar.TransactOpts, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.ClaimWithResolver(&_QRNSReverseRegistrar.TransactOpts, owner, resolver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.RenounceOwnership(&_QRNSReverseRegistrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.RenounceOwnership(&_QRNSReverseRegistrar.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) SetController(opts *bind.TransactOpts, controller common.Address, enabled bool) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "setController", controller, enabled)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) SetController(controller common.Address, enabled bool) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetController(&_QRNSReverseRegistrar.TransactOpts, controller, enabled)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) SetController(controller common.Address, enabled bool) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetController(&_QRNSReverseRegistrar.TransactOpts, controller, enabled)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) SetDefaultResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "setDefaultResolver", resolver)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) SetDefaultResolver(resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetDefaultResolver(&_QRNSReverseRegistrar.TransactOpts, resolver)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) SetDefaultResolver(resolver common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetDefaultResolver(&_QRNSReverseRegistrar.TransactOpts, resolver)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) SetName(name string) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetName(&_QRNSReverseRegistrar.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetName(&_QRNSReverseRegistrar.TransactOpts, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) SetNameForAddr(opts *bind.TransactOpts, addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "setNameForAddr", addr, owner, resolver, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) SetNameForAddr(addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetNameForAddr(&_QRNSReverseRegistrar.TransactOpts, addr, owner, resolver, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) SetNameForAddr(addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.SetNameForAddr(&_QRNSReverseRegistrar.TransactOpts, addr, owner, resolver, name)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.TransferOwnership(&_QRNSReverseRegistrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSReverseRegistrar *QRNSReverseRegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QRNSReverseRegistrar.Contract.TransferOwnership(&_QRNSReverseRegistrar.TransactOpts, newOwner)
}

// QRNSReverseRegistrarControllerChangedIterator is returned from FilterControllerChanged and is used to iterate over the raw logs and unpacked data for ControllerChanged events raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarControllerChangedIterator struct {
	Event *QRNSReverseRegistrarControllerChanged // Event containing the contract specifics and raw log

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
func (it *QRNSReverseRegistrarControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSReverseRegistrarControllerChanged)
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
		it.Event = new(QRNSReverseRegistrarControllerChanged)
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
func (it *QRNSReverseRegistrarControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSReverseRegistrarControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSReverseRegistrarControllerChanged represents a ControllerChanged event raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarControllerChanged struct {
	Controller common.Address
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerChanged is a free log retrieval operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) FilterControllerChanged(opts *bind.FilterOpts, controller []common.Address) (*QRNSReverseRegistrarControllerChangedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.FilterLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrarControllerChangedIterator{contract: _QRNSReverseRegistrar.contract, event: "ControllerChanged", logs: logs, sub: sub}, nil
}

// WatchControllerChanged is a free log subscription operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) WatchControllerChanged(opts *bind.WatchOpts, sink chan<- *QRNSReverseRegistrarControllerChanged, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.WatchLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSReverseRegistrarControllerChanged)
				if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
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

// ParseControllerChanged is a log parse operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) ParseControllerChanged(log types.Log) (*QRNSReverseRegistrarControllerChanged, error) {
	event := new(QRNSReverseRegistrarControllerChanged)
	if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSReverseRegistrarDefaultResolverChangedIterator is returned from FilterDefaultResolverChanged and is used to iterate over the raw logs and unpacked data for DefaultResolverChanged events raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarDefaultResolverChangedIterator struct {
	Event *QRNSReverseRegistrarDefaultResolverChanged // Event containing the contract specifics and raw log

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
func (it *QRNSReverseRegistrarDefaultResolverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSReverseRegistrarDefaultResolverChanged)
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
		it.Event = new(QRNSReverseRegistrarDefaultResolverChanged)
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
func (it *QRNSReverseRegistrarDefaultResolverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSReverseRegistrarDefaultResolverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSReverseRegistrarDefaultResolverChanged represents a DefaultResolverChanged event raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarDefaultResolverChanged struct {
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDefaultResolverChanged is a free log retrieval operation binding the contract event 0xeae17a84d9eb83d8c8eb317f9e7d64857bc363fa51674d996c023f4340c577cf.
//
// Solidity: event DefaultResolverChanged(address indexed resolver)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) FilterDefaultResolverChanged(opts *bind.FilterOpts, resolver []common.Address) (*QRNSReverseRegistrarDefaultResolverChangedIterator, error) {

	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.FilterLogs(opts, "DefaultResolverChanged", resolverRule)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrarDefaultResolverChangedIterator{contract: _QRNSReverseRegistrar.contract, event: "DefaultResolverChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultResolverChanged is a free log subscription operation binding the contract event 0xeae17a84d9eb83d8c8eb317f9e7d64857bc363fa51674d996c023f4340c577cf.
//
// Solidity: event DefaultResolverChanged(address indexed resolver)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) WatchDefaultResolverChanged(opts *bind.WatchOpts, sink chan<- *QRNSReverseRegistrarDefaultResolverChanged, resolver []common.Address) (event.Subscription, error) {

	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.WatchLogs(opts, "DefaultResolverChanged", resolverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSReverseRegistrarDefaultResolverChanged)
				if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "DefaultResolverChanged", log); err != nil {
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

// ParseDefaultResolverChanged is a log parse operation binding the contract event 0xeae17a84d9eb83d8c8eb317f9e7d64857bc363fa51674d996c023f4340c577cf.
//
// Solidity: event DefaultResolverChanged(address indexed resolver)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) ParseDefaultResolverChanged(log types.Log) (*QRNSReverseRegistrarDefaultResolverChanged, error) {
	event := new(QRNSReverseRegistrarDefaultResolverChanged)
	if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "DefaultResolverChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSReverseRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarOwnershipTransferredIterator struct {
	Event *QRNSReverseRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QRNSReverseRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSReverseRegistrarOwnershipTransferred)
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
		it.Event = new(QRNSReverseRegistrarOwnershipTransferred)
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
func (it *QRNSReverseRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSReverseRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSReverseRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QRNSReverseRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrarOwnershipTransferredIterator{contract: _QRNSReverseRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QRNSReverseRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSReverseRegistrarOwnershipTransferred)
				if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*QRNSReverseRegistrarOwnershipTransferred, error) {
	event := new(QRNSReverseRegistrarOwnershipTransferred)
	if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSReverseRegistrarReverseClaimedIterator is returned from FilterReverseClaimed and is used to iterate over the raw logs and unpacked data for ReverseClaimed events raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarReverseClaimedIterator struct {
	Event *QRNSReverseRegistrarReverseClaimed // Event containing the contract specifics and raw log

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
func (it *QRNSReverseRegistrarReverseClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSReverseRegistrarReverseClaimed)
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
		it.Event = new(QRNSReverseRegistrarReverseClaimed)
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
func (it *QRNSReverseRegistrarReverseClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSReverseRegistrarReverseClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSReverseRegistrarReverseClaimed represents a ReverseClaimed event raised by the QRNSReverseRegistrar contract.
type QRNSReverseRegistrarReverseClaimed struct {
	Addr common.Address
	Node [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReverseClaimed is a free log retrieval operation binding the contract event 0x6ada868dd3058cf77a48a74489fd7963688e5464b2b0fa957ace976243270e92.
//
// Solidity: event ReverseClaimed(address indexed addr, bytes32 indexed node)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) FilterReverseClaimed(opts *bind.FilterOpts, addr []common.Address, node [][32]byte) (*QRNSReverseRegistrarReverseClaimedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.FilterLogs(opts, "ReverseClaimed", addrRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSReverseRegistrarReverseClaimedIterator{contract: _QRNSReverseRegistrar.contract, event: "ReverseClaimed", logs: logs, sub: sub}, nil
}

// WatchReverseClaimed is a free log subscription operation binding the contract event 0x6ada868dd3058cf77a48a74489fd7963688e5464b2b0fa957ace976243270e92.
//
// Solidity: event ReverseClaimed(address indexed addr, bytes32 indexed node)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) WatchReverseClaimed(opts *bind.WatchOpts, sink chan<- *QRNSReverseRegistrarReverseClaimed, addr []common.Address, node [][32]byte) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSReverseRegistrar.contract.WatchLogs(opts, "ReverseClaimed", addrRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSReverseRegistrarReverseClaimed)
				if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "ReverseClaimed", log); err != nil {
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

// ParseReverseClaimed is a log parse operation binding the contract event 0x6ada868dd3058cf77a48a74489fd7963688e5464b2b0fa957ace976243270e92.
//
// Solidity: event ReverseClaimed(address indexed addr, bytes32 indexed node)
func (_QRNSReverseRegistrar *QRNSReverseRegistrarFilterer) ParseReverseClaimed(log types.Log) (*QRNSReverseRegistrarReverseClaimed, error) {
	event := new(QRNSReverseRegistrarReverseClaimed)
	if err := _QRNSReverseRegistrar.contract.UnpackLog(event, "ReverseClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
