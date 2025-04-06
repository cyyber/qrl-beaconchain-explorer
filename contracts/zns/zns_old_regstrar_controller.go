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

// ZNSOldRegistrarControllerMetaData contains all meta data concerning the ZNSOldRegistrarController contract.
var ZNSOldRegistrarControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrar\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"makeCommitmentWithConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerWithConfig\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"}],\"name\":\"setCommitmentAges\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZNSOldRegistrarControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ZNSOldRegistrarControllerMetaData.ABI instead.
var ZNSOldRegistrarControllerABI = ZNSOldRegistrarControllerMetaData.ABI

// ZNSOldRegistrarController is an auto generated Go binding around a Zond contract.
type ZNSOldRegistrarController struct {
	ZNSOldRegistrarControllerCaller     // Read-only binding to the contract
	ZNSOldRegistrarControllerTransactor // Write-only binding to the contract
	ZNSOldRegistrarControllerFilterer   // Log filterer for contract events
}

// ZNSOldRegistrarControllerCaller is an auto generated read-only Go binding around a Zond contract.
type ZNSOldRegistrarControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSOldRegistrarControllerTransactor is an auto generated write-only Go binding around a Zond contract.
type ZNSOldRegistrarControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSOldRegistrarControllerFilterer is an auto generated log filtering Go binding around a Zond contract events.
type ZNSOldRegistrarControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSOldRegistrarControllerSession is an auto generated Go binding around a Zond contract,
// with pre-set call and transact options.
type ZNSOldRegistrarControllerSession struct {
	Contract     *ZNSOldRegistrarController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ZNSOldRegistrarControllerCallerSession is an auto generated read-only Go binding around a Zond contract,
// with pre-set call options.
type ZNSOldRegistrarControllerCallerSession struct {
	Contract *ZNSOldRegistrarControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ZNSOldRegistrarControllerTransactorSession is an auto generated write-only Go binding around a Zond contract,
// with pre-set transact options.
type ZNSOldRegistrarControllerTransactorSession struct {
	Contract     *ZNSOldRegistrarControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ZNSOldRegistrarControllerRaw is an auto generated low-level Go binding around a Zond contract.
type ZNSOldRegistrarControllerRaw struct {
	Contract *ZNSOldRegistrarController // Generic contract binding to access the raw methods on
}

// ZNSOldRegistrarControllerCallerRaw is an auto generated low-level read-only Go binding around a Zond contract.
type ZNSOldRegistrarControllerCallerRaw struct {
	Contract *ZNSOldRegistrarControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ZNSOldRegistrarControllerTransactorRaw is an auto generated low-level write-only Go binding around a Zond contract.
type ZNSOldRegistrarControllerTransactorRaw struct {
	Contract *ZNSOldRegistrarControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZNSOldRegistrarController creates a new instance of ZNSOldRegistrarController, bound to a specific deployed contract.
func NewZNSOldRegistrarController(address common.Address, backend bind.ContractBackend) (*ZNSOldRegistrarController, error) {
	contract, err := bindZNSOldRegistrarController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarController{ZNSOldRegistrarControllerCaller: ZNSOldRegistrarControllerCaller{contract: contract}, ZNSOldRegistrarControllerTransactor: ZNSOldRegistrarControllerTransactor{contract: contract}, ZNSOldRegistrarControllerFilterer: ZNSOldRegistrarControllerFilterer{contract: contract}}, nil
}

// NewZNSOldRegistrarControllerCaller creates a new read-only instance of ZNSOldRegistrarController, bound to a specific deployed contract.
func NewZNSOldRegistrarControllerCaller(address common.Address, caller bind.ContractCaller) (*ZNSOldRegistrarControllerCaller, error) {
	contract, err := bindZNSOldRegistrarController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarControllerCaller{contract: contract}, nil
}

// NewZNSOldRegistrarControllerTransactor creates a new write-only instance of ZNSOldRegistrarController, bound to a specific deployed contract.
func NewZNSOldRegistrarControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ZNSOldRegistrarControllerTransactor, error) {
	contract, err := bindZNSOldRegistrarController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarControllerTransactor{contract: contract}, nil
}

// NewZNSOldRegistrarControllerFilterer creates a new log filterer instance of ZNSOldRegistrarController, bound to a specific deployed contract.
func NewZNSOldRegistrarControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ZNSOldRegistrarControllerFilterer, error) {
	contract, err := bindZNSOldRegistrarController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarControllerFilterer{contract: contract}, nil
}

// bindZNSOldRegistrarController binds a generic wrapper to an already deployed contract.
func bindZNSOldRegistrarController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZNSOldRegistrarControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSOldRegistrarController.Contract.ZNSOldRegistrarControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.ZNSOldRegistrarControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.ZNSOldRegistrarControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSOldRegistrarController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.contract.Transact(opts, method, params...)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.MINREGISTRATIONDURATION(&_ZNSOldRegistrarController.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.MINREGISTRATIONDURATION(&_ZNSOldRegistrarController.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Available(name string) (bool, error) {
	return _ZNSOldRegistrarController.Contract.Available(&_ZNSOldRegistrarController.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) Available(name string) (bool, error) {
	return _ZNSOldRegistrarController.Contract.Available(&_ZNSOldRegistrarController.CallOpts, name)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.Commitments(&_ZNSOldRegistrarController.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.Commitments(&_ZNSOldRegistrarController.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) IsOwner() (bool, error) {
	return _ZNSOldRegistrarController.Contract.IsOwner(&_ZNSOldRegistrarController.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) IsOwner() (bool, error) {
	return _ZNSOldRegistrarController.Contract.IsOwner(&_ZNSOldRegistrarController.CallOpts)
}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) MakeCommitment(opts *bind.CallOpts, name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "makeCommitment", name, owner, secret)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) MakeCommitment(name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	return _ZNSOldRegistrarController.Contract.MakeCommitment(&_ZNSOldRegistrarController.CallOpts, name, owner, secret)
}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) MakeCommitment(name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	return _ZNSOldRegistrarController.Contract.MakeCommitment(&_ZNSOldRegistrarController.CallOpts, name, owner, secret)
}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) MakeCommitmentWithConfig(opts *bind.CallOpts, name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "makeCommitmentWithConfig", name, owner, secret, resolver, addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) MakeCommitmentWithConfig(name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	return _ZNSOldRegistrarController.Contract.MakeCommitmentWithConfig(&_ZNSOldRegistrarController.CallOpts, name, owner, secret, resolver, addr)
}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) MakeCommitmentWithConfig(name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	return _ZNSOldRegistrarController.Contract.MakeCommitmentWithConfig(&_ZNSOldRegistrarController.CallOpts, name, owner, secret, resolver, addr)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) MaxCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "maxCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) MaxCommitmentAge() (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.MaxCommitmentAge(&_ZNSOldRegistrarController.CallOpts)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) MaxCommitmentAge() (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.MaxCommitmentAge(&_ZNSOldRegistrarController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) MinCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "minCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) MinCommitmentAge() (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.MinCommitmentAge(&_ZNSOldRegistrarController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) MinCommitmentAge() (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.MinCommitmentAge(&_ZNSOldRegistrarController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Owner() (common.Address, error) {
	return _ZNSOldRegistrarController.Contract.Owner(&_ZNSOldRegistrarController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) Owner() (common.Address, error) {
	return _ZNSOldRegistrarController.Contract.Owner(&_ZNSOldRegistrarController.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "rentPrice", name, duration)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) RentPrice(name string, duration *big.Int) (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.RentPrice(&_ZNSOldRegistrarController.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) RentPrice(name string, duration *big.Int) (*big.Int, error) {
	return _ZNSOldRegistrarController.Contract.RentPrice(&_ZNSOldRegistrarController.CallOpts, name, duration)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSOldRegistrarController.Contract.SupportsInterface(&_ZNSOldRegistrarController.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSOldRegistrarController.Contract.SupportsInterface(&_ZNSOldRegistrarController.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCaller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ZNSOldRegistrarController.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Valid(name string) (bool, error) {
	return _ZNSOldRegistrarController.Contract.Valid(&_ZNSOldRegistrarController.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerCallerSession) Valid(name string) (bool, error) {
	return _ZNSOldRegistrarController.Contract.Valid(&_ZNSOldRegistrarController.CallOpts, name)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) Commit(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Commit(&_ZNSOldRegistrarController.TransactOpts, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Commit(&_ZNSOldRegistrarController.TransactOpts, commitment)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "register", name, owner, duration, secret)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Register(&_ZNSOldRegistrarController.TransactOpts, name, owner, duration, secret)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Register(&_ZNSOldRegistrarController.TransactOpts, name, owner, duration, secret)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) RegisterWithConfig(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "registerWithConfig", name, owner, duration, secret, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) RegisterWithConfig(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.RegisterWithConfig(&_ZNSOldRegistrarController.TransactOpts, name, owner, duration, secret, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) RegisterWithConfig(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.RegisterWithConfig(&_ZNSOldRegistrarController.TransactOpts, name, owner, duration, secret, resolver, addr)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Renew(&_ZNSOldRegistrarController.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Renew(&_ZNSOldRegistrarController.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.RenounceOwnership(&_ZNSOldRegistrarController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.RenounceOwnership(&_ZNSOldRegistrarController.TransactOpts)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) SetCommitmentAges(opts *bind.TransactOpts, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "setCommitmentAges", _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) SetCommitmentAges(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.SetCommitmentAges(&_ZNSOldRegistrarController.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) SetCommitmentAges(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.SetCommitmentAges(&_ZNSOldRegistrarController.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) SetPriceOracle(opts *bind.TransactOpts, _prices common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "setPriceOracle", _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) SetPriceOracle(_prices common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.SetPriceOracle(&_ZNSOldRegistrarController.TransactOpts, _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) SetPriceOracle(_prices common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.SetPriceOracle(&_ZNSOldRegistrarController.TransactOpts, _prices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.TransferOwnership(&_ZNSOldRegistrarController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.TransferOwnership(&_ZNSOldRegistrarController.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSOldRegistrarController.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerSession) Withdraw() (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Withdraw(&_ZNSOldRegistrarController.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ZNSOldRegistrarController.Contract.Withdraw(&_ZNSOldRegistrarController.TransactOpts)
}

// ZNSOldRegistrarControllerNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerNameRegisteredIterator struct {
	Event *ZNSOldRegistrarControllerNameRegistered // Event containing the contract specifics and raw log

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
func (it *ZNSOldRegistrarControllerNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSOldRegistrarControllerNameRegistered)
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
		it.Event = new(ZNSOldRegistrarControllerNameRegistered)
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
func (it *ZNSOldRegistrarControllerNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSOldRegistrarControllerNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSOldRegistrarControllerNameRegistered represents a NameRegistered event raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerNameRegistered struct {
	Name    string
	Label   [32]byte
	Owner   common.Address
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*ZNSOldRegistrarControllerNameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarControllerNameRegisteredIterator{contract: _ZNSOldRegistrarController.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *ZNSOldRegistrarControllerNameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSOldRegistrarControllerNameRegistered)
				if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) ParseNameRegistered(log types.Log) (*ZNSOldRegistrarControllerNameRegistered, error) {
	event := new(ZNSOldRegistrarControllerNameRegistered)
	if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSOldRegistrarControllerNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerNameRenewedIterator struct {
	Event *ZNSOldRegistrarControllerNameRenewed // Event containing the contract specifics and raw log

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
func (it *ZNSOldRegistrarControllerNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSOldRegistrarControllerNameRenewed)
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
		it.Event = new(ZNSOldRegistrarControllerNameRenewed)
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
func (it *ZNSOldRegistrarControllerNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSOldRegistrarControllerNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSOldRegistrarControllerNameRenewed represents a NameRenewed event raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerNameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*ZNSOldRegistrarControllerNameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarControllerNameRenewedIterator{contract: _ZNSOldRegistrarController.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *ZNSOldRegistrarControllerNameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSOldRegistrarControllerNameRenewed)
				if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) ParseNameRenewed(log types.Log) (*ZNSOldRegistrarControllerNameRenewed, error) {
	event := new(ZNSOldRegistrarControllerNameRenewed)
	if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSOldRegistrarControllerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerNewPriceOracleIterator struct {
	Event *ZNSOldRegistrarControllerNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *ZNSOldRegistrarControllerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSOldRegistrarControllerNewPriceOracle)
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
		it.Event = new(ZNSOldRegistrarControllerNewPriceOracle)
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
func (it *ZNSOldRegistrarControllerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSOldRegistrarControllerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSOldRegistrarControllerNewPriceOracle represents a NewPriceOracle event raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerNewPriceOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts, oracle []common.Address) (*ZNSOldRegistrarControllerNewPriceOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.FilterLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarControllerNewPriceOracleIterator{contract: _ZNSOldRegistrarController.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ZNSOldRegistrarControllerNewPriceOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.WatchLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSOldRegistrarControllerNewPriceOracle)
				if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) ParseNewPriceOracle(log types.Log) (*ZNSOldRegistrarControllerNewPriceOracle, error) {
	event := new(ZNSOldRegistrarControllerNewPriceOracle)
	if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSOldRegistrarControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerOwnershipTransferredIterator struct {
	Event *ZNSOldRegistrarControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZNSOldRegistrarControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSOldRegistrarControllerOwnershipTransferred)
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
		it.Event = new(ZNSOldRegistrarControllerOwnershipTransferred)
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
func (it *ZNSOldRegistrarControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSOldRegistrarControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSOldRegistrarControllerOwnershipTransferred represents a OwnershipTransferred event raised by the ZNSOldRegistrarController contract.
type ZNSOldRegistrarControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZNSOldRegistrarControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZNSOldRegistrarControllerOwnershipTransferredIterator{contract: _ZNSOldRegistrarController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZNSOldRegistrarControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZNSOldRegistrarController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSOldRegistrarControllerOwnershipTransferred)
				if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZNSOldRegistrarController *ZNSOldRegistrarControllerFilterer) ParseOwnershipTransferred(log types.Log) (*ZNSOldRegistrarControllerOwnershipTransferred, error) {
	event := new(ZNSOldRegistrarControllerOwnershipTransferred)
	if err := _ZNSOldRegistrarController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
