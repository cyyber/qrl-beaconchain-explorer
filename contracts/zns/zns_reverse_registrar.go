// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zns

import (
	"errors"
	"math/big"
	"strings"

	zond "github.com/theQRL/go-zond"
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
	_ = zond.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ZNSReverseRegistrarMetaData contains all meta data concerning the ZNSReverseRegistrar contract.
var ZNSReverseRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractZNS\",\"name\":\"znsAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ControllerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractNameResolver\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"DefaultResolverChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ReverseClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimForAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimWithResolver\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultResolver\",\"outputs\":[{\"internalType\":\"contractNameResolver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zns\",\"outputs\":[{\"internalType\":\"contractZNS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"node\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setDefaultResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setNameForAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZNSReverseRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use ZNSReverseRegistrarMetaData.ABI instead.
var ZNSReverseRegistrarABI = ZNSReverseRegistrarMetaData.ABI

// ZNSReverseRegistrar is an auto generated Go binding around an Zond contract.
type ZNSReverseRegistrar struct {
	ZNSReverseRegistrarCaller     // Read-only binding to the contract
	ZNSReverseRegistrarTransactor // Write-only binding to the contract
	ZNSReverseRegistrarFilterer   // Log filterer for contract events
}

// ZNSReverseRegistrarCaller is an auto generated read-only Go binding around an Zond contract.
type ZNSReverseRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSReverseRegistrarTransactor is an auto generated write-only Go binding around an Zond contract.
type ZNSReverseRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSReverseRegistrarFilterer is an auto generated log filtering Go binding around an Zond contract events.
type ZNSReverseRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSReverseRegistrarSession is an auto generated Go binding around an Zond contract,
// with pre-set call and transact options.
type ZNSReverseRegistrarSession struct {
	Contract     *ZNSReverseRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZNSReverseRegistrarCallerSession is an auto generated read-only Go binding around an Zond contract,
// with pre-set call options.
type ZNSReverseRegistrarCallerSession struct {
	Contract *ZNSReverseRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ZNSReverseRegistrarTransactorSession is an auto generated write-only Go binding around an Zond contract,
// with pre-set transact options.
type ZNSReverseRegistrarTransactorSession struct {
	Contract     *ZNSReverseRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ZNSReverseRegistrarRaw is an auto generated low-level Go binding around an Zond contract.
type ZNSReverseRegistrarRaw struct {
	Contract *ZNSReverseRegistrar // Generic contract binding to access the raw methods on
}

// ZNSReverseRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Zond contract.
type ZNSReverseRegistrarCallerRaw struct {
	Contract *ZNSReverseRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// ZNSReverseRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Zond contract.
type ZNSReverseRegistrarTransactorRaw struct {
	Contract *ZNSReverseRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZNSReverseRegistrar creates a new instance of ZNSReverseRegistrar, bound to a specific deployed contract.
func NewZNSReverseRegistrar(address common.Address, backend bind.ContractBackend) (*ZNSReverseRegistrar, error) {
	contract, err := bindZNSReverseRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrar{ZNSReverseRegistrarCaller: ZNSReverseRegistrarCaller{contract: contract}, ZNSReverseRegistrarTransactor: ZNSReverseRegistrarTransactor{contract: contract}, ZNSReverseRegistrarFilterer: ZNSReverseRegistrarFilterer{contract: contract}}, nil
}

// NewZNSReverseRegistrarCaller creates a new read-only instance of ZNSReverseRegistrar, bound to a specific deployed contract.
func NewZNSReverseRegistrarCaller(address common.Address, caller bind.ContractCaller) (*ZNSReverseRegistrarCaller, error) {
	contract, err := bindZNSReverseRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrarCaller{contract: contract}, nil
}

// NewZNSReverseRegistrarTransactor creates a new write-only instance of ZNSReverseRegistrar, bound to a specific deployed contract.
func NewZNSReverseRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*ZNSReverseRegistrarTransactor, error) {
	contract, err := bindZNSReverseRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrarTransactor{contract: contract}, nil
}

// NewZNSReverseRegistrarFilterer creates a new log filterer instance of ZNSReverseRegistrar, bound to a specific deployed contract.
func NewZNSReverseRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*ZNSReverseRegistrarFilterer, error) {
	contract, err := bindZNSReverseRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrarFilterer{contract: contract}, nil
}

// bindZNSReverseRegistrar binds a generic wrapper to an already deployed contract.
func bindZNSReverseRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZNSReverseRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSReverseRegistrar *ZNSReverseRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSReverseRegistrar.Contract.ZNSReverseRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSReverseRegistrar *ZNSReverseRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.ZNSReverseRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSReverseRegistrar *ZNSReverseRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.ZNSReverseRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSReverseRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ZNSReverseRegistrar.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) Controllers(arg0 common.Address) (bool, error) {
	return _ZNSReverseRegistrar.Contract.Controllers(&_ZNSReverseRegistrar.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _ZNSReverseRegistrar.Contract.Controllers(&_ZNSReverseRegistrar.CallOpts, arg0)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCaller) DefaultResolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSReverseRegistrar.contract.Call(opts, &out, "defaultResolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) DefaultResolver() (common.Address, error) {
	return _ZNSReverseRegistrar.Contract.DefaultResolver(&_ZNSReverseRegistrar.CallOpts)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCallerSession) DefaultResolver() (common.Address, error) {
	return _ZNSReverseRegistrar.Contract.DefaultResolver(&_ZNSReverseRegistrar.CallOpts)
}

// Zns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function zns() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCaller) Zns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSReverseRegistrar.contract.Call(opts, &out, "zns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Zns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function zns() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) Zns() (common.Address, error) {
	return _ZNSReverseRegistrar.Contract.Zns(&_ZNSReverseRegistrar.CallOpts)
}

// Zns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function zns() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCallerSession) Zns() (common.Address, error) {
	return _ZNSReverseRegistrar.Contract.Zns(&_ZNSReverseRegistrar.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCaller) Node(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ZNSReverseRegistrar.contract.Call(opts, &out, "node", addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) Node(addr common.Address) ([32]byte, error) {
	return _ZNSReverseRegistrar.Contract.Node(&_ZNSReverseRegistrar.CallOpts, addr)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCallerSession) Node(addr common.Address) ([32]byte, error) {
	return _ZNSReverseRegistrar.Contract.Node(&_ZNSReverseRegistrar.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSReverseRegistrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) Owner() (common.Address, error) {
	return _ZNSReverseRegistrar.Contract.Owner(&_ZNSReverseRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarCallerSession) Owner() (common.Address, error) {
	return _ZNSReverseRegistrar.Contract.Owner(&_ZNSReverseRegistrar.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) Claim(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "claim", owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.Claim(&_ZNSReverseRegistrar.TransactOpts, owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.Claim(&_ZNSReverseRegistrar.TransactOpts, owner)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) ClaimForAddr(opts *bind.TransactOpts, addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "claimForAddr", addr, owner, resolver)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) ClaimForAddr(addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.ClaimForAddr(&_ZNSReverseRegistrar.TransactOpts, addr, owner, resolver)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) ClaimForAddr(addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.ClaimForAddr(&_ZNSReverseRegistrar.TransactOpts, addr, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) ClaimWithResolver(opts *bind.TransactOpts, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "claimWithResolver", owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.ClaimWithResolver(&_ZNSReverseRegistrar.TransactOpts, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.ClaimWithResolver(&_ZNSReverseRegistrar.TransactOpts, owner, resolver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.RenounceOwnership(&_ZNSReverseRegistrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.RenounceOwnership(&_ZNSReverseRegistrar.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) SetController(opts *bind.TransactOpts, controller common.Address, enabled bool) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "setController", controller, enabled)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) SetController(controller common.Address, enabled bool) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetController(&_ZNSReverseRegistrar.TransactOpts, controller, enabled)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) SetController(controller common.Address, enabled bool) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetController(&_ZNSReverseRegistrar.TransactOpts, controller, enabled)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) SetDefaultResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "setDefaultResolver", resolver)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) SetDefaultResolver(resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetDefaultResolver(&_ZNSReverseRegistrar.TransactOpts, resolver)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) SetDefaultResolver(resolver common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetDefaultResolver(&_ZNSReverseRegistrar.TransactOpts, resolver)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) SetName(name string) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetName(&_ZNSReverseRegistrar.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetName(&_ZNSReverseRegistrar.TransactOpts, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) SetNameForAddr(opts *bind.TransactOpts, addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "setNameForAddr", addr, owner, resolver, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) SetNameForAddr(addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetNameForAddr(&_ZNSReverseRegistrar.TransactOpts, addr, owner, resolver, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) SetNameForAddr(addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.SetNameForAddr(&_ZNSReverseRegistrar.TransactOpts, addr, owner, resolver, name)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.TransferOwnership(&_ZNSReverseRegistrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSReverseRegistrar *ZNSReverseRegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZNSReverseRegistrar.Contract.TransferOwnership(&_ZNSReverseRegistrar.TransactOpts, newOwner)
}

// ZNSReverseRegistrarControllerChangedIterator is returned from FilterControllerChanged and is used to iterate over the raw logs and unpacked data for ControllerChanged events raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarControllerChangedIterator struct {
	Event *ZNSReverseRegistrarControllerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  zond.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZNSReverseRegistrarControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSReverseRegistrarControllerChanged)
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
		it.Event = new(ZNSReverseRegistrarControllerChanged)
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
func (it *ZNSReverseRegistrarControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSReverseRegistrarControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSReverseRegistrarControllerChanged represents a ControllerChanged event raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarControllerChanged struct {
	Controller common.Address
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerChanged is a free log retrieval operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) FilterControllerChanged(opts *bind.FilterOpts, controller []common.Address) (*ZNSReverseRegistrarControllerChangedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.FilterLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrarControllerChangedIterator{contract: _ZNSReverseRegistrar.contract, event: "ControllerChanged", logs: logs, sub: sub}, nil
}

// WatchControllerChanged is a free log subscription operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) WatchControllerChanged(opts *bind.WatchOpts, sink chan<- *ZNSReverseRegistrarControllerChanged, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.WatchLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSReverseRegistrarControllerChanged)
				if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
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
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) ParseControllerChanged(log types.Log) (*ZNSReverseRegistrarControllerChanged, error) {
	event := new(ZNSReverseRegistrarControllerChanged)
	if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSReverseRegistrarDefaultResolverChangedIterator is returned from FilterDefaultResolverChanged and is used to iterate over the raw logs and unpacked data for DefaultResolverChanged events raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarDefaultResolverChangedIterator struct {
	Event *ZNSReverseRegistrarDefaultResolverChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  zond.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZNSReverseRegistrarDefaultResolverChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSReverseRegistrarDefaultResolverChanged)
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
		it.Event = new(ZNSReverseRegistrarDefaultResolverChanged)
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
func (it *ZNSReverseRegistrarDefaultResolverChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSReverseRegistrarDefaultResolverChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSReverseRegistrarDefaultResolverChanged represents a DefaultResolverChanged event raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarDefaultResolverChanged struct {
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDefaultResolverChanged is a free log retrieval operation binding the contract event 0xeae17a84d9eb83d8c8eb317f9e7d64857bc363fa51674d996c023f4340c577cf.
//
// Solidity: event DefaultResolverChanged(address indexed resolver)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) FilterDefaultResolverChanged(opts *bind.FilterOpts, resolver []common.Address) (*ZNSReverseRegistrarDefaultResolverChangedIterator, error) {

	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.FilterLogs(opts, "DefaultResolverChanged", resolverRule)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrarDefaultResolverChangedIterator{contract: _ZNSReverseRegistrar.contract, event: "DefaultResolverChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultResolverChanged is a free log subscription operation binding the contract event 0xeae17a84d9eb83d8c8eb317f9e7d64857bc363fa51674d996c023f4340c577cf.
//
// Solidity: event DefaultResolverChanged(address indexed resolver)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) WatchDefaultResolverChanged(opts *bind.WatchOpts, sink chan<- *ZNSReverseRegistrarDefaultResolverChanged, resolver []common.Address) (event.Subscription, error) {

	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.WatchLogs(opts, "DefaultResolverChanged", resolverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSReverseRegistrarDefaultResolverChanged)
				if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "DefaultResolverChanged", log); err != nil {
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
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) ParseDefaultResolverChanged(log types.Log) (*ZNSReverseRegistrarDefaultResolverChanged, error) {
	event := new(ZNSReverseRegistrarDefaultResolverChanged)
	if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "DefaultResolverChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSReverseRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarOwnershipTransferredIterator struct {
	Event *ZNSReverseRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  zond.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZNSReverseRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSReverseRegistrarOwnershipTransferred)
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
		it.Event = new(ZNSReverseRegistrarOwnershipTransferred)
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
func (it *ZNSReverseRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSReverseRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSReverseRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZNSReverseRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrarOwnershipTransferredIterator{contract: _ZNSReverseRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZNSReverseRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSReverseRegistrarOwnershipTransferred)
				if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*ZNSReverseRegistrarOwnershipTransferred, error) {
	event := new(ZNSReverseRegistrarOwnershipTransferred)
	if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSReverseRegistrarReverseClaimedIterator is returned from FilterReverseClaimed and is used to iterate over the raw logs and unpacked data for ReverseClaimed events raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarReverseClaimedIterator struct {
	Event *ZNSReverseRegistrarReverseClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  zond.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZNSReverseRegistrarReverseClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSReverseRegistrarReverseClaimed)
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
		it.Event = new(ZNSReverseRegistrarReverseClaimed)
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
func (it *ZNSReverseRegistrarReverseClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSReverseRegistrarReverseClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSReverseRegistrarReverseClaimed represents a ReverseClaimed event raised by the ZNSReverseRegistrar contract.
type ZNSReverseRegistrarReverseClaimed struct {
	Addr common.Address
	Node [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReverseClaimed is a free log retrieval operation binding the contract event 0x6ada868dd3058cf77a48a74489fd7963688e5464b2b0fa957ace976243270e92.
//
// Solidity: event ReverseClaimed(address indexed addr, bytes32 indexed node)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) FilterReverseClaimed(opts *bind.FilterOpts, addr []common.Address, node [][32]byte) (*ZNSReverseRegistrarReverseClaimedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.FilterLogs(opts, "ReverseClaimed", addrRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSReverseRegistrarReverseClaimedIterator{contract: _ZNSReverseRegistrar.contract, event: "ReverseClaimed", logs: logs, sub: sub}, nil
}

// WatchReverseClaimed is a free log subscription operation binding the contract event 0x6ada868dd3058cf77a48a74489fd7963688e5464b2b0fa957ace976243270e92.
//
// Solidity: event ReverseClaimed(address indexed addr, bytes32 indexed node)
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) WatchReverseClaimed(opts *bind.WatchOpts, sink chan<- *ZNSReverseRegistrarReverseClaimed, addr []common.Address, node [][32]byte) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSReverseRegistrar.contract.WatchLogs(opts, "ReverseClaimed", addrRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSReverseRegistrarReverseClaimed)
				if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "ReverseClaimed", log); err != nil {
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
func (_ZNSReverseRegistrar *ZNSReverseRegistrarFilterer) ParseReverseClaimed(log types.Log) (*ZNSReverseRegistrarReverseClaimed, error) {
	event := new(ZNSReverseRegistrarReverseClaimed)
	if err := _ZNSReverseRegistrar.contract.UnpackLog(event, "ReverseClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
