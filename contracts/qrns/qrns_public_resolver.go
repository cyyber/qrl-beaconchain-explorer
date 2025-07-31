// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qrns

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

// ZNSPublicResolverMetaData contains all meta data concerning the ZNSPublicResolver contract.
var ZNSPublicResolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractZNS\",\"name\":\"_zns\",\"type\":\"address\"},{\"internalType\":\"contractINameWrapper\",\"name\":\"wrapperAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedZONDController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedReverseRegistrar\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"record\",\"type\":\"bytes\"}],\"name\":\"DNSRecordChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"DNSRecordDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lastzonehash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zonehash\",\"type\":\"bytes\"}],\"name\":\"DNSZonehashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVersion\",\"type\":\"uint64\"}],\"name\":\"VersionChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"clearRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"hasDNSRecords\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"interfaceImplementer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"isApprovedFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodehash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicallWithNodeCheck\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"recordVersions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDNSRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setZonehash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"zonehash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZNSPublicResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use ZNSPublicResolverMetaData.ABI instead.
var ZNSPublicResolverABI = ZNSPublicResolverMetaData.ABI

// ZNSPublicResolver is an auto generated Go binding around a QRL contract.
type ZNSPublicResolver struct {
	ZNSPublicResolverCaller     // Read-only binding to the contract
	ZNSPublicResolverTransactor // Write-only binding to the contract
	ZNSPublicResolverFilterer   // Log filterer for contract events
}

// ZNSPublicResolverCaller is an auto generated read-only Go binding around a QRL contract.
type ZNSPublicResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSPublicResolverTransactor is an auto generated write-only Go binding around a QRL contract.
type ZNSPublicResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSPublicResolverFilterer is an auto generated log filtering Go binding around a QRL contract events.
type ZNSPublicResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSPublicResolverSession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type ZNSPublicResolverSession struct {
	Contract     *ZNSPublicResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZNSPublicResolverCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type ZNSPublicResolverCallerSession struct {
	Contract *ZNSPublicResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZNSPublicResolverTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type ZNSPublicResolverTransactorSession struct {
	Contract     *ZNSPublicResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZNSPublicResolverRaw is an auto generated low-level Go binding around a QRL contract.
type ZNSPublicResolverRaw struct {
	Contract *ZNSPublicResolver // Generic contract binding to access the raw methods on
}

// ZNSPublicResolverCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type ZNSPublicResolverCallerRaw struct {
	Contract *ZNSPublicResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ZNSPublicResolverTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type ZNSPublicResolverTransactorRaw struct {
	Contract *ZNSPublicResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZNSPublicResolver creates a new instance of ZNSPublicResolver, bound to a specific deployed contract.
func NewZNSPublicResolver(address common.Address, backend bind.ContractBackend) (*ZNSPublicResolver, error) {
	contract, err := bindZNSPublicResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolver{ZNSPublicResolverCaller: ZNSPublicResolverCaller{contract: contract}, ZNSPublicResolverTransactor: ZNSPublicResolverTransactor{contract: contract}, ZNSPublicResolverFilterer: ZNSPublicResolverFilterer{contract: contract}}, nil
}

// NewZNSPublicResolverCaller creates a new read-only instance of ZNSPublicResolver, bound to a specific deployed contract.
func NewZNSPublicResolverCaller(address common.Address, caller bind.ContractCaller) (*ZNSPublicResolverCaller, error) {
	contract, err := bindZNSPublicResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverCaller{contract: contract}, nil
}

// NewZNSPublicResolverTransactor creates a new write-only instance of ZNSPublicResolver, bound to a specific deployed contract.
func NewZNSPublicResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ZNSPublicResolverTransactor, error) {
	contract, err := bindZNSPublicResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverTransactor{contract: contract}, nil
}

// NewZNSPublicResolverFilterer creates a new log filterer instance of ZNSPublicResolver, bound to a specific deployed contract.
func NewZNSPublicResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ZNSPublicResolverFilterer, error) {
	contract, err := bindZNSPublicResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverFilterer{contract: contract}, nil
}

// bindZNSPublicResolver binds a generic wrapper to an already deployed contract.
func bindZNSPublicResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZNSPublicResolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSPublicResolver *ZNSPublicResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSPublicResolver.Contract.ZNSPublicResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSPublicResolver *ZNSPublicResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.ZNSPublicResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSPublicResolver *ZNSPublicResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.ZNSPublicResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSPublicResolver *ZNSPublicResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSPublicResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSPublicResolver *ZNSPublicResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSPublicResolver *ZNSPublicResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "ABI", node, contentTypes)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ZNSPublicResolver *ZNSPublicResolverSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _ZNSPublicResolver.Contract.ABI(&_ZNSPublicResolver.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _ZNSPublicResolver.Contract.ABI(&_ZNSPublicResolver.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _ZNSPublicResolver.Contract.Addr(&_ZNSPublicResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _ZNSPublicResolver.Contract.Addr(&_ZNSPublicResolver.CallOpts, node)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) Addr0(opts *bind.CallOpts, node [32]byte, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "addr0", node, coinType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _ZNSPublicResolver.Contract.Addr0(&_ZNSPublicResolver.CallOpts, node, coinType)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _ZNSPublicResolver.Contract.Addr0(&_ZNSPublicResolver.CallOpts, node, coinType)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) Contenthash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "contenthash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Contenthash(node [32]byte) ([]byte, error) {
	return _ZNSPublicResolver.Contract.Contenthash(&_ZNSPublicResolver.CallOpts, node)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) Contenthash(node [32]byte) ([]byte, error) {
	return _ZNSPublicResolver.Contract.Contenthash(&_ZNSPublicResolver.CallOpts, node)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) DnsRecord(opts *bind.CallOpts, node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "dnsRecord", node, name, resource)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _ZNSPublicResolver.Contract.DnsRecord(&_ZNSPublicResolver.CallOpts, node, name, resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _ZNSPublicResolver.Contract.DnsRecord(&_ZNSPublicResolver.CallOpts, node, name, resource)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) HasDNSRecords(opts *bind.CallOpts, node [32]byte, name [32]byte) (bool, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "hasDNSRecords", node, name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverSession) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _ZNSPublicResolver.Contract.HasDNSRecords(&_ZNSPublicResolver.CallOpts, node, name)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _ZNSPublicResolver.Contract.HasDNSRecords(&_ZNSPublicResolver.CallOpts, node, name)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) InterfaceImplementer(opts *bind.CallOpts, node [32]byte, interfaceID [4]byte) (common.Address, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "interfaceImplementer", node, interfaceID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ZNSPublicResolver *ZNSPublicResolverSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _ZNSPublicResolver.Contract.InterfaceImplementer(&_ZNSPublicResolver.CallOpts, node, interfaceID)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _ZNSPublicResolver.Contract.InterfaceImplementer(&_ZNSPublicResolver.CallOpts, node, interfaceID)
}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) IsApprovedFor(opts *bind.CallOpts, owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "isApprovedFor", owner, node, delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverSession) IsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	return _ZNSPublicResolver.Contract.IsApprovedFor(&_ZNSPublicResolver.CallOpts, owner, node, delegate)
}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) IsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	return _ZNSPublicResolver.Contract.IsApprovedFor(&_ZNSPublicResolver.CallOpts, owner, node, delegate)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ZNSPublicResolver.Contract.IsApprovedForAll(&_ZNSPublicResolver.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ZNSPublicResolver.Contract.IsApprovedForAll(&_ZNSPublicResolver.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Name(node [32]byte) (string, error) {
	return _ZNSPublicResolver.Contract.Name(&_ZNSPublicResolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) Name(node [32]byte) (string, error) {
	return _ZNSPublicResolver.Contract.Name(&_ZNSPublicResolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "pubkey", node)

	outstruct := new(struct {
		X [32]byte
		Y [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ZNSPublicResolver.Contract.Pubkey(&_ZNSPublicResolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ZNSPublicResolver.Contract.Pubkey(&_ZNSPublicResolver.CallOpts, node)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) RecordVersions(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "recordVersions", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_ZNSPublicResolver *ZNSPublicResolverSession) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _ZNSPublicResolver.Contract.RecordVersions(&_ZNSPublicResolver.CallOpts, arg0)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _ZNSPublicResolver.Contract.RecordVersions(&_ZNSPublicResolver.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSPublicResolver.Contract.SupportsInterface(&_ZNSPublicResolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSPublicResolver.Contract.SupportsInterface(&_ZNSPublicResolver.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Text(node [32]byte, key string) (string, error) {
	return _ZNSPublicResolver.Contract.Text(&_ZNSPublicResolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _ZNSPublicResolver.Contract.Text(&_ZNSPublicResolver.CallOpts, node, key)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCaller) Zonehash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ZNSPublicResolver.contract.Call(opts, &out, "zonehash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Zonehash(node [32]byte) ([]byte, error) {
	return _ZNSPublicResolver.Contract.Zonehash(&_ZNSPublicResolver.CallOpts, node)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_ZNSPublicResolver *ZNSPublicResolverCallerSession) Zonehash(node [32]byte) ([]byte, error) {
	return _ZNSPublicResolver.Contract.Zonehash(&_ZNSPublicResolver.CallOpts, node)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) Approve(opts *bind.TransactOpts, node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "approve", node, delegate, approved)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) Approve(node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.Approve(&_ZNSPublicResolver.TransactOpts, node, delegate, approved)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) Approve(node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.Approve(&_ZNSPublicResolver.TransactOpts, node, delegate, approved)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) ClearRecords(opts *bind.TransactOpts, node [32]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "clearRecords", node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) ClearRecords(node [32]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.ClearRecords(&_ZNSPublicResolver.TransactOpts, node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) ClearRecords(node [32]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.ClearRecords(&_ZNSPublicResolver.TransactOpts, node)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ZNSPublicResolver *ZNSPublicResolverSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.Multicall(&_ZNSPublicResolver.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.Multicall(&_ZNSPublicResolver.TransactOpts, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) MulticallWithNodeCheck(opts *bind.TransactOpts, nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "multicallWithNodeCheck", nodehash, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_ZNSPublicResolver *ZNSPublicResolverSession) MulticallWithNodeCheck(nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.MulticallWithNodeCheck(&_ZNSPublicResolver.TransactOpts, nodehash, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) MulticallWithNodeCheck(nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.MulticallWithNodeCheck(&_ZNSPublicResolver.TransactOpts, nodehash, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetABI(&_ZNSPublicResolver.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetABI(&_ZNSPublicResolver.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setAddr", node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetAddr(&_ZNSPublicResolver.TransactOpts, node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetAddr(&_ZNSPublicResolver.TransactOpts, node, coinType, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetAddr0(opts *bind.TransactOpts, node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setAddr0", node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetAddr0(&_ZNSPublicResolver.TransactOpts, node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetAddr0(&_ZNSPublicResolver.TransactOpts, node, a)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetApprovalForAll(&_ZNSPublicResolver.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetApprovalForAll(&_ZNSPublicResolver.TransactOpts, operator, approved)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetContenthash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setContenthash", node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetContenthash(&_ZNSPublicResolver.TransactOpts, node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetContenthash(&_ZNSPublicResolver.TransactOpts, node, hash)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetDNSRecords(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setDNSRecords", node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetDNSRecords(&_ZNSPublicResolver.TransactOpts, node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetDNSRecords(&_ZNSPublicResolver.TransactOpts, node, data)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetInterface(opts *bind.TransactOpts, node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setInterface", node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetInterface(&_ZNSPublicResolver.TransactOpts, node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetInterface(&_ZNSPublicResolver.TransactOpts, node, interfaceID, implementer)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetName(opts *bind.TransactOpts, node [32]byte, newName string) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setName", node, newName)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetName(node [32]byte, newName string) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetName(&_ZNSPublicResolver.TransactOpts, node, newName)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetName(node [32]byte, newName string) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetName(&_ZNSPublicResolver.TransactOpts, node, newName)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetPubkey(&_ZNSPublicResolver.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetPubkey(&_ZNSPublicResolver.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetText(&_ZNSPublicResolver.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetText(&_ZNSPublicResolver.TransactOpts, node, key, value)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactor) SetZonehash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.contract.Transact(opts, "setZonehash", node, hash)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (_ZNSPublicResolver *ZNSPublicResolverSession) SetZonehash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetZonehash(&_ZNSPublicResolver.TransactOpts, node, hash)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (_ZNSPublicResolver *ZNSPublicResolverTransactorSession) SetZonehash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ZNSPublicResolver.Contract.SetZonehash(&_ZNSPublicResolver.TransactOpts, node, hash)
}

// ZNSPublicResolverABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverABIChangedIterator struct {
	Event *ZNSPublicResolverABIChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverABIChanged)
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
		it.Event = new(ZNSPublicResolverABIChanged)
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
func (it *ZNSPublicResolverABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverABIChanged represents a ABIChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*ZNSPublicResolverABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverABIChangedIterator{contract: _ZNSPublicResolver.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverABIChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ParseABIChanged is a log parse operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseABIChanged(log types.Log) (*ZNSPublicResolverABIChanged, error) {
	event := new(ZNSPublicResolverABIChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverAddrChangedIterator struct {
	Event *ZNSPublicResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverAddrChanged)
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
		it.Event = new(ZNSPublicResolverAddrChanged)
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
func (it *ZNSPublicResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverAddrChanged represents a AddrChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverAddrChangedIterator{contract: _ZNSPublicResolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverAddrChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseAddrChanged(log types.Log) (*ZNSPublicResolverAddrChanged, error) {
	event := new(ZNSPublicResolverAddrChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverAddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverAddressChangedIterator struct {
	Event *ZNSPublicResolverAddressChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverAddressChanged)
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
		it.Event = new(ZNSPublicResolverAddressChanged)
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
func (it *ZNSPublicResolverAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverAddressChanged represents a AddressChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverAddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterAddressChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverAddressChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverAddressChangedIterator{contract: _ZNSPublicResolver.contract, event: "AddressChanged", logs: logs, sub: sub}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverAddressChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverAddressChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "AddressChanged", log); err != nil {
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

// ParseAddressChanged is a log parse operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseAddressChanged(log types.Log) (*ZNSPublicResolverAddressChanged, error) {
	event := new(ZNSPublicResolverAddressChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverApprovalForAllIterator struct {
	Event *ZNSPublicResolverApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverApprovalForAll)
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
		it.Event = new(ZNSPublicResolverApprovalForAll)
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
func (it *ZNSPublicResolverApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverApprovalForAll represents a ApprovalForAll event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ZNSPublicResolverApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverApprovalForAllIterator{contract: _ZNSPublicResolver.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverApprovalForAll)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseApprovalForAll(log types.Log) (*ZNSPublicResolverApprovalForAll, error) {
	event := new(ZNSPublicResolverApprovalForAll)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverApprovedIterator struct {
	Event *ZNSPublicResolverApproved // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverApproved)
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
		it.Event = new(ZNSPublicResolverApproved)
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
func (it *ZNSPublicResolverApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverApproved represents a Approved event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverApproved struct {
	Owner    common.Address
	Node     [32]byte
	Delegate common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Solidity: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterApproved(opts *bind.FilterOpts, node [][32]byte, delegate []common.Address, approved []bool) (*ZNSPublicResolverApprovedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "Approved", nodeRule, delegateRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverApprovedIterator{contract: _ZNSPublicResolver.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Solidity: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverApproved, node [][32]byte, delegate []common.Address, approved []bool) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "Approved", nodeRule, delegateRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverApproved)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Solidity: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseApproved(log types.Log) (*ZNSPublicResolverApproved, error) {
	event := new(ZNSPublicResolverApproved)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverContenthashChangedIterator is returned from FilterContenthashChanged and is used to iterate over the raw logs and unpacked data for ContenthashChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverContenthashChangedIterator struct {
	Event *ZNSPublicResolverContenthashChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverContenthashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverContenthashChanged)
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
		it.Event = new(ZNSPublicResolverContenthashChanged)
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
func (it *ZNSPublicResolverContenthashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverContenthashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverContenthashChanged represents a ContenthashChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContenthashChanged is a free log retrieval operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterContenthashChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverContenthashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverContenthashChangedIterator{contract: _ZNSPublicResolver.contract, event: "ContenthashChanged", logs: logs, sub: sub}, nil
}

// WatchContenthashChanged is a free log subscription operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchContenthashChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverContenthashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverContenthashChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
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

// ParseContenthashChanged is a log parse operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseContenthashChanged(log types.Log) (*ZNSPublicResolverContenthashChanged, error) {
	event := new(ZNSPublicResolverContenthashChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverDNSRecordChangedIterator is returned from FilterDNSRecordChanged and is used to iterate over the raw logs and unpacked data for DNSRecordChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverDNSRecordChangedIterator struct {
	Event *ZNSPublicResolverDNSRecordChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverDNSRecordChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverDNSRecordChanged)
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
		it.Event = new(ZNSPublicResolverDNSRecordChanged)
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
func (it *ZNSPublicResolverDNSRecordChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverDNSRecordChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverDNSRecordChanged represents a DNSRecordChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverDNSRecordChanged struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Record   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordChanged is a free log retrieval operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterDNSRecordChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverDNSRecordChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverDNSRecordChangedIterator{contract: _ZNSPublicResolver.contract, event: "DNSRecordChanged", logs: logs, sub: sub}, nil
}

// WatchDNSRecordChanged is a free log subscription operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchDNSRecordChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverDNSRecordChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverDNSRecordChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
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

// ParseDNSRecordChanged is a log parse operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseDNSRecordChanged(log types.Log) (*ZNSPublicResolverDNSRecordChanged, error) {
	event := new(ZNSPublicResolverDNSRecordChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverDNSRecordDeletedIterator is returned from FilterDNSRecordDeleted and is used to iterate over the raw logs and unpacked data for DNSRecordDeleted events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverDNSRecordDeletedIterator struct {
	Event *ZNSPublicResolverDNSRecordDeleted // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverDNSRecordDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverDNSRecordDeleted)
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
		it.Event = new(ZNSPublicResolverDNSRecordDeleted)
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
func (it *ZNSPublicResolverDNSRecordDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverDNSRecordDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverDNSRecordDeleted represents a DNSRecordDeleted event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverDNSRecordDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordDeleted is a free log retrieval operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterDNSRecordDeleted(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverDNSRecordDeletedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverDNSRecordDeletedIterator{contract: _ZNSPublicResolver.contract, event: "DNSRecordDeleted", logs: logs, sub: sub}, nil
}

// WatchDNSRecordDeleted is a free log subscription operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchDNSRecordDeleted(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverDNSRecordDeleted, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverDNSRecordDeleted)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
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

// ParseDNSRecordDeleted is a log parse operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseDNSRecordDeleted(log types.Log) (*ZNSPublicResolverDNSRecordDeleted, error) {
	event := new(ZNSPublicResolverDNSRecordDeleted)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverDNSZonehashChangedIterator is returned from FilterDNSZonehashChanged and is used to iterate over the raw logs and unpacked data for DNSZonehashChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverDNSZonehashChangedIterator struct {
	Event *ZNSPublicResolverDNSZonehashChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverDNSZonehashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverDNSZonehashChanged)
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
		it.Event = new(ZNSPublicResolverDNSZonehashChanged)
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
func (it *ZNSPublicResolverDNSZonehashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverDNSZonehashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverDNSZonehashChanged represents a DNSZonehashChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverDNSZonehashChanged struct {
	Node         [32]byte
	Lastzonehash []byte
	Zonehash     []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDNSZonehashChanged is a free log retrieval operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterDNSZonehashChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverDNSZonehashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverDNSZonehashChangedIterator{contract: _ZNSPublicResolver.contract, event: "DNSZonehashChanged", logs: logs, sub: sub}, nil
}

// WatchDNSZonehashChanged is a free log subscription operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchDNSZonehashChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverDNSZonehashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverDNSZonehashChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
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

// ParseDNSZonehashChanged is a log parse operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseDNSZonehashChanged(log types.Log) (*ZNSPublicResolverDNSZonehashChanged, error) {
	event := new(ZNSPublicResolverDNSZonehashChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverInterfaceChangedIterator is returned from FilterInterfaceChanged and is used to iterate over the raw logs and unpacked data for InterfaceChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverInterfaceChangedIterator struct {
	Event *ZNSPublicResolverInterfaceChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverInterfaceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverInterfaceChanged)
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
		it.Event = new(ZNSPublicResolverInterfaceChanged)
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
func (it *ZNSPublicResolverInterfaceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverInterfaceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverInterfaceChanged represents a InterfaceChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverInterfaceChanged struct {
	Node        [32]byte
	InterfaceID [4]byte
	Implementer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterfaceChanged is a free log retrieval operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterInterfaceChanged(opts *bind.FilterOpts, node [][32]byte, interfaceID [][4]byte) (*ZNSPublicResolverInterfaceChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverInterfaceChangedIterator{contract: _ZNSPublicResolver.contract, event: "InterfaceChanged", logs: logs, sub: sub}, nil
}

// WatchInterfaceChanged is a free log subscription operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchInterfaceChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverInterfaceChanged, node [][32]byte, interfaceID [][4]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverInterfaceChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
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

// ParseInterfaceChanged is a log parse operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseInterfaceChanged(log types.Log) (*ZNSPublicResolverInterfaceChanged, error) {
	event := new(ZNSPublicResolverInterfaceChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverNameChangedIterator struct {
	Event *ZNSPublicResolverNameChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverNameChanged)
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
		it.Event = new(ZNSPublicResolverNameChanged)
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
func (it *ZNSPublicResolverNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverNameChanged represents a NameChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverNameChangedIterator{contract: _ZNSPublicResolver.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverNameChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseNameChanged(log types.Log) (*ZNSPublicResolverNameChanged, error) {
	event := new(ZNSPublicResolverNameChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverPubkeyChangedIterator struct {
	Event *ZNSPublicResolverPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverPubkeyChanged)
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
		it.Event = new(ZNSPublicResolverPubkeyChanged)
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
func (it *ZNSPublicResolverPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverPubkeyChanged represents a PubkeyChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverPubkeyChangedIterator{contract: _ZNSPublicResolver.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverPubkeyChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ParsePubkeyChanged is a log parse operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParsePubkeyChanged(log types.Log) (*ZNSPublicResolverPubkeyChanged, error) {
	event := new(ZNSPublicResolverPubkeyChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverTextChangedIterator struct {
	Event *ZNSPublicResolverTextChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverTextChanged)
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
		it.Event = new(ZNSPublicResolverTextChanged)
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
func (it *ZNSPublicResolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverTextChanged represents a TextChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Value      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ZNSPublicResolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverTextChangedIterator{contract: _ZNSPublicResolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverTextChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseTextChanged(log types.Log) (*ZNSPublicResolverTextChanged, error) {
	event := new(ZNSPublicResolverTextChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSPublicResolverVersionChangedIterator is returned from FilterVersionChanged and is used to iterate over the raw logs and unpacked data for VersionChanged events raised by the ZNSPublicResolver contract.
type ZNSPublicResolverVersionChangedIterator struct {
	Event *ZNSPublicResolverVersionChanged // Event containing the contract specifics and raw log

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
func (it *ZNSPublicResolverVersionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSPublicResolverVersionChanged)
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
		it.Event = new(ZNSPublicResolverVersionChanged)
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
func (it *ZNSPublicResolverVersionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSPublicResolverVersionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSPublicResolverVersionChanged represents a VersionChanged event raised by the ZNSPublicResolver contract.
type ZNSPublicResolverVersionChanged struct {
	Node       [32]byte
	NewVersion uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVersionChanged is a free log retrieval operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) FilterVersionChanged(opts *bind.FilterOpts, node [][32]byte) (*ZNSPublicResolverVersionChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.FilterLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ZNSPublicResolverVersionChangedIterator{contract: _ZNSPublicResolver.contract, event: "VersionChanged", logs: logs, sub: sub}, nil
}

// WatchVersionChanged is a free log subscription operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) WatchVersionChanged(opts *bind.WatchOpts, sink chan<- *ZNSPublicResolverVersionChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ZNSPublicResolver.contract.WatchLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSPublicResolverVersionChanged)
				if err := _ZNSPublicResolver.contract.UnpackLog(event, "VersionChanged", log); err != nil {
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

// ParseVersionChanged is a log parse operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_ZNSPublicResolver *ZNSPublicResolverFilterer) ParseVersionChanged(log types.Log) (*ZNSPublicResolverVersionChanged, error) {
	event := new(ZNSPublicResolverVersionChanged)
	if err := _ZNSPublicResolver.contract.UnpackLog(event, "VersionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
