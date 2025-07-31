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

// QRNSRegistryMetaData contains all meta data concerning the QRNSRegistry contract.
var QRNSRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"_old\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"old\",\"outputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QRNSRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use QRNSRegistryMetaData.ABI instead.
var QRNSRegistryABI = QRNSRegistryMetaData.ABI

// QRNSRegistry is an auto generated Go binding around a QRL contract.
type QRNSRegistry struct {
	QRNSRegistryCaller     // Read-only binding to the contract
	QRNSRegistryTransactor // Write-only binding to the contract
	QRNSRegistryFilterer   // Log filterer for contract events
}

// QRNSRegistryCaller is an auto generated read-only Go binding around a QRL contract.
type QRNSRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSRegistryTransactor is an auto generated write-only Go binding around a QRL contract.
type QRNSRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSRegistryFilterer is an auto generated log filtering Go binding around a QRL contract events.
type QRNSRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSRegistrySession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type QRNSRegistrySession struct {
	Contract     *QRNSRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QRNSRegistryCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type QRNSRegistryCallerSession struct {
	Contract *QRNSRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// QRNSRegistryTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type QRNSRegistryTransactorSession struct {
	Contract     *QRNSRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// QRNSRegistryRaw is an auto generated low-level Go binding around a QRL contract.
type QRNSRegistryRaw struct {
	Contract *QRNSRegistry // Generic contract binding to access the raw methods on
}

// QRNSRegistryCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type QRNSRegistryCallerRaw struct {
	Contract *QRNSRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// QRNSRegistryTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type QRNSRegistryTransactorRaw struct {
	Contract *QRNSRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQRNSRegistry creates a new instance of QRNSRegistry, bound to a specific deployed contract.
func NewQRNSRegistry(address common.Address, backend bind.ContractBackend) (*QRNSRegistry, error) {
	contract, err := bindQRNSRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistry{QRNSRegistryCaller: QRNSRegistryCaller{contract: contract}, QRNSRegistryTransactor: QRNSRegistryTransactor{contract: contract}, QRNSRegistryFilterer: QRNSRegistryFilterer{contract: contract}}, nil
}

// NewQRNSRegistryCaller creates a new read-only instance of QRNSRegistry, bound to a specific deployed contract.
func NewQRNSRegistryCaller(address common.Address, caller bind.ContractCaller) (*QRNSRegistryCaller, error) {
	contract, err := bindQRNSRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryCaller{contract: contract}, nil
}

// NewQRNSRegistryTransactor creates a new write-only instance of QRNSRegistry, bound to a specific deployed contract.
func NewQRNSRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*QRNSRegistryTransactor, error) {
	contract, err := bindQRNSRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryTransactor{contract: contract}, nil
}

// NewQRNSRegistryFilterer creates a new log filterer instance of QRNSRegistry, bound to a specific deployed contract.
func NewQRNSRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*QRNSRegistryFilterer, error) {
	contract, err := bindQRNSRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryFilterer{contract: contract}, nil
}

// bindQRNSRegistry binds a generic wrapper to an already deployed contract.
func bindQRNSRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QRNSRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSRegistry *QRNSRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSRegistry.Contract.QRNSRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSRegistry *QRNSRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.QRNSRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSRegistry *QRNSRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.QRNSRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSRegistry *QRNSRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSRegistry *QRNSRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSRegistry *QRNSRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_QRNSRegistry *QRNSRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_QRNSRegistry *QRNSRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _QRNSRegistry.Contract.IsApprovedForAll(&_QRNSRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_QRNSRegistry *QRNSRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _QRNSRegistry.Contract.IsApprovedForAll(&_QRNSRegistry.CallOpts, owner, operator)
}

// Old is a free data retrieval call binding the contract method 0xb83f8663.
//
// Solidity: function old() view returns(address)
func (_QRNSRegistry *QRNSRegistryCaller) Old(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSRegistry.contract.Call(opts, &out, "old")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Old is a free data retrieval call binding the contract method 0xb83f8663.
//
// Solidity: function old() view returns(address)
func (_QRNSRegistry *QRNSRegistrySession) Old() (common.Address, error) {
	return _QRNSRegistry.Contract.Old(&_QRNSRegistry.CallOpts)
}

// Old is a free data retrieval call binding the contract method 0xb83f8663.
//
// Solidity: function old() view returns(address)
func (_QRNSRegistry *QRNSRegistryCallerSession) Old() (common.Address, error) {
	return _QRNSRegistry.Contract.Old(&_QRNSRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_QRNSRegistry *QRNSRegistryCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _QRNSRegistry.contract.Call(opts, &out, "owner", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_QRNSRegistry *QRNSRegistrySession) Owner(node [32]byte) (common.Address, error) {
	return _QRNSRegistry.Contract.Owner(&_QRNSRegistry.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_QRNSRegistry *QRNSRegistryCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _QRNSRegistry.Contract.Owner(&_QRNSRegistry.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_QRNSRegistry *QRNSRegistryCaller) RecordExists(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _QRNSRegistry.contract.Call(opts, &out, "recordExists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_QRNSRegistry *QRNSRegistrySession) RecordExists(node [32]byte) (bool, error) {
	return _QRNSRegistry.Contract.RecordExists(&_QRNSRegistry.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_QRNSRegistry *QRNSRegistryCallerSession) RecordExists(node [32]byte) (bool, error) {
	return _QRNSRegistry.Contract.RecordExists(&_QRNSRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_QRNSRegistry *QRNSRegistryCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _QRNSRegistry.contract.Call(opts, &out, "resolver", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_QRNSRegistry *QRNSRegistrySession) Resolver(node [32]byte) (common.Address, error) {
	return _QRNSRegistry.Contract.Resolver(&_QRNSRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_QRNSRegistry *QRNSRegistryCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _QRNSRegistry.Contract.Resolver(&_QRNSRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_QRNSRegistry *QRNSRegistryCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var out []interface{}
	err := _QRNSRegistry.contract.Call(opts, &out, "ttl", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_QRNSRegistry *QRNSRegistrySession) Ttl(node [32]byte) (uint64, error) {
	return _QRNSRegistry.Contract.Ttl(&_QRNSRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_QRNSRegistry *QRNSRegistryCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _QRNSRegistry.Contract.Ttl(&_QRNSRegistry.CallOpts, node)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSRegistry *QRNSRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSRegistry *QRNSRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetApprovalForAll(&_QRNSRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSRegistry *QRNSRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetApprovalForAll(&_QRNSRegistry.TransactOpts, operator, approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_QRNSRegistry *QRNSRegistryTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_QRNSRegistry *QRNSRegistrySession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetOwner(&_QRNSRegistry.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_QRNSRegistry *QRNSRegistryTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetOwner(&_QRNSRegistry.TransactOpts, node, owner)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistryTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistrySession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetRecord(&_QRNSRegistry.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistryTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetRecord(&_QRNSRegistry.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_QRNSRegistry *QRNSRegistryTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_QRNSRegistry *QRNSRegistrySession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetResolver(&_QRNSRegistry.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_QRNSRegistry *QRNSRegistryTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetResolver(&_QRNSRegistry.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_QRNSRegistry *QRNSRegistryTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_QRNSRegistry *QRNSRegistrySession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetSubnodeOwner(&_QRNSRegistry.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_QRNSRegistry *QRNSRegistryTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetSubnodeOwner(&_QRNSRegistry.TransactOpts, node, label, owner)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistryTransactor) SetSubnodeRecord(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.contract.Transact(opts, "setSubnodeRecord", node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistrySession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetSubnodeRecord(&_QRNSRegistry.TransactOpts, node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistryTransactorSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetSubnodeRecord(&_QRNSRegistry.TransactOpts, node, label, owner, resolver, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistryTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistrySession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetTTL(&_QRNSRegistry.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_QRNSRegistry *QRNSRegistryTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _QRNSRegistry.Contract.SetTTL(&_QRNSRegistry.TransactOpts, node, ttl)
}

// QRNSRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the QRNSRegistry contract.
type QRNSRegistryApprovalForAllIterator struct {
	Event *QRNSRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *QRNSRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSRegistryApprovalForAll)
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
		it.Event = new(QRNSRegistryApprovalForAll)
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
func (it *QRNSRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSRegistryApprovalForAll represents a ApprovalForAll event raised by the QRNSRegistry contract.
type QRNSRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QRNSRegistry *QRNSRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*QRNSRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QRNSRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryApprovalForAllIterator{contract: _QRNSRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QRNSRegistry *QRNSRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *QRNSRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QRNSRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSRegistryApprovalForAll)
				if err := _QRNSRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QRNSRegistry *QRNSRegistryFilterer) ParseApprovalForAll(log types.Log) (*QRNSRegistryApprovalForAll, error) {
	event := new(QRNSRegistryApprovalForAll)
	if err := _QRNSRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSRegistryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the QRNSRegistry contract.
type QRNSRegistryNewOwnerIterator struct {
	Event *QRNSRegistryNewOwner // Event containing the contract specifics and raw log

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
func (it *QRNSRegistryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSRegistryNewOwner)
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
		it.Event = new(QRNSRegistryNewOwner)
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
func (it *QRNSRegistryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSRegistryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSRegistryNewOwner represents a NewOwner event raised by the QRNSRegistry contract.
type QRNSRegistryNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_QRNSRegistry *QRNSRegistryFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*QRNSRegistryNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _QRNSRegistry.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryNewOwnerIterator{contract: _QRNSRegistry.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_QRNSRegistry *QRNSRegistryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *QRNSRegistryNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _QRNSRegistry.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSRegistryNewOwner)
				if err := _QRNSRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_QRNSRegistry *QRNSRegistryFilterer) ParseNewOwner(log types.Log) (*QRNSRegistryNewOwner, error) {
	event := new(QRNSRegistryNewOwner)
	if err := _QRNSRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSRegistryNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the QRNSRegistry contract.
type QRNSRegistryNewResolverIterator struct {
	Event *QRNSRegistryNewResolver // Event containing the contract specifics and raw log

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
func (it *QRNSRegistryNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSRegistryNewResolver)
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
		it.Event = new(QRNSRegistryNewResolver)
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
func (it *QRNSRegistryNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSRegistryNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSRegistryNewResolver represents a NewResolver event raised by the QRNSRegistry contract.
type QRNSRegistryNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_QRNSRegistry *QRNSRegistryFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*QRNSRegistryNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSRegistry.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryNewResolverIterator{contract: _QRNSRegistry.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_QRNSRegistry *QRNSRegistryFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *QRNSRegistryNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSRegistry.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSRegistryNewResolver)
				if err := _QRNSRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_QRNSRegistry *QRNSRegistryFilterer) ParseNewResolver(log types.Log) (*QRNSRegistryNewResolver, error) {
	event := new(QRNSRegistryNewResolver)
	if err := _QRNSRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSRegistryNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the QRNSRegistry contract.
type QRNSRegistryNewTTLIterator struct {
	Event *QRNSRegistryNewTTL // Event containing the contract specifics and raw log

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
func (it *QRNSRegistryNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSRegistryNewTTL)
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
		it.Event = new(QRNSRegistryNewTTL)
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
func (it *QRNSRegistryNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSRegistryNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSRegistryNewTTL represents a NewTTL event raised by the QRNSRegistry contract.
type QRNSRegistryNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_QRNSRegistry *QRNSRegistryFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*QRNSRegistryNewTTLIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSRegistry.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryNewTTLIterator{contract: _QRNSRegistry.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_QRNSRegistry *QRNSRegistryFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *QRNSRegistryNewTTL, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSRegistry.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSRegistryNewTTL)
				if err := _QRNSRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_QRNSRegistry *QRNSRegistryFilterer) ParseNewTTL(log types.Log) (*QRNSRegistryNewTTL, error) {
	event := new(QRNSRegistryNewTTL)
	if err := _QRNSRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the QRNSRegistry contract.
type QRNSRegistryTransferIterator struct {
	Event *QRNSRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *QRNSRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSRegistryTransfer)
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
		it.Event = new(QRNSRegistryTransfer)
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
func (it *QRNSRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSRegistryTransfer represents a Transfer event raised by the QRNSRegistry contract.
type QRNSRegistryTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_QRNSRegistry *QRNSRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*QRNSRegistryTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSRegistry.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSRegistryTransferIterator{contract: _QRNSRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_QRNSRegistry *QRNSRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *QRNSRegistryTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSRegistry.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSRegistryTransfer)
				if err := _QRNSRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_QRNSRegistry *QRNSRegistryFilterer) ParseTransfer(log types.Log) (*QRNSRegistryTransfer, error) {
	event := new(QRNSRegistryTransfer)
	if err := _QRNSRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
