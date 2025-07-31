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

// DNSSECRRSetWithSignature is an auto generated low-level Go binding around an user-defined struct.
type DNSSECRRSetWithSignature struct {
	Rrset []byte
	Sig   []byte
}

// ZNSDNSRegistrarMetaData contains all meta data concerning the ZNSDNSRegistrar contract.
var ZNSDNSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousRegistrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"},{\"internalType\":\"contractDNSSEC\",\"name\":\"_dnssec\",\"type\":\"address\"},{\"internalType\":\"contractPublicSuffixList\",\"name\":\"_suffixes\",\"type\":\"address\"},{\"internalType\":\"contractZNS\",\"name\":\"_zns\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"InvalidPublicSuffix\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOwnerRecordFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"OffsetOutOfBoundsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreconditionNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleProof\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dnsname\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"inception\",\"type\":\"uint32\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"suffixes\",\"type\":\"address\"}],\"name\":\"NewPublicSuffixList\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"domain\",\"type\":\"bytes\"}],\"name\":\"enableNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zns\",\"outputs\":[{\"internalType\":\"contractZNS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"inceptions\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractDNSSEC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousRegistrar\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rrset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structDNSSEC.RRSetWithSignature[]\",\"name\":\"input\",\"type\":\"tuple[]\"}],\"name\":\"proveAndClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rrset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structDNSSEC.RRSetWithSignature[]\",\"name\":\"input\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"proveAndClaimWithResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPublicSuffixList\",\"name\":\"_suffixes\",\"type\":\"address\"}],\"name\":\"setPublicSuffixList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suffixes\",\"outputs\":[{\"internalType\":\"contractPublicSuffixList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ZNSDNSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use ZNSDNSRegistrarMetaData.ABI instead.
var ZNSDNSRegistrarABI = ZNSDNSRegistrarMetaData.ABI

// ZNSDNSRegistrar is an auto generated Go binding around a QRL contract.
type ZNSDNSRegistrar struct {
	ZNSDNSRegistrarCaller     // Read-only binding to the contract
	ZNSDNSRegistrarTransactor // Write-only binding to the contract
	ZNSDNSRegistrarFilterer   // Log filterer for contract events
}

// ZNSDNSRegistrarCaller is an auto generated read-only Go binding around a QRL contract.
type ZNSDNSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSDNSRegistrarTransactor is an auto generated write-only Go binding around a QRL contract.
type ZNSDNSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSDNSRegistrarFilterer is an auto generated log filtering Go binding around a QRL contract events.
type ZNSDNSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSDNSRegistrarSession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type ZNSDNSRegistrarSession struct {
	Contract     *ZNSDNSRegistrar  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZNSDNSRegistrarCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type ZNSDNSRegistrarCallerSession struct {
	Contract *ZNSDNSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ZNSDNSRegistrarTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type ZNSDNSRegistrarTransactorSession struct {
	Contract     *ZNSDNSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ZNSDNSRegistrarRaw is an auto generated low-level Go binding around a QRL contract.
type ZNSDNSRegistrarRaw struct {
	Contract *ZNSDNSRegistrar // Generic contract binding to access the raw methods on
}

// ZNSDNSRegistrarCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type ZNSDNSRegistrarCallerRaw struct {
	Contract *ZNSDNSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// ZNSDNSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type ZNSDNSRegistrarTransactorRaw struct {
	Contract *ZNSDNSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZNSDNSRegistrar creates a new instance of ZNSDNSRegistrar, bound to a specific deployed contract.
func NewZNSDNSRegistrar(address common.Address, backend bind.ContractBackend) (*ZNSDNSRegistrar, error) {
	contract, err := bindZNSDNSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZNSDNSRegistrar{ZNSDNSRegistrarCaller: ZNSDNSRegistrarCaller{contract: contract}, ZNSDNSRegistrarTransactor: ZNSDNSRegistrarTransactor{contract: contract}, ZNSDNSRegistrarFilterer: ZNSDNSRegistrarFilterer{contract: contract}}, nil
}

// NewZNSDNSRegistrarCaller creates a new read-only instance of ZNSDNSRegistrar, bound to a specific deployed contract.
func NewZNSDNSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*ZNSDNSRegistrarCaller, error) {
	contract, err := bindZNSDNSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSDNSRegistrarCaller{contract: contract}, nil
}

// NewZNSDNSRegistrarTransactor creates a new write-only instance of ZNSDNSRegistrar, bound to a specific deployed contract.
func NewZNSDNSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*ZNSDNSRegistrarTransactor, error) {
	contract, err := bindZNSDNSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSDNSRegistrarTransactor{contract: contract}, nil
}

// NewZNSDNSRegistrarFilterer creates a new log filterer instance of ZNSDNSRegistrar, bound to a specific deployed contract.
func NewZNSDNSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*ZNSDNSRegistrarFilterer, error) {
	contract, err := bindZNSDNSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZNSDNSRegistrarFilterer{contract: contract}, nil
}

// bindZNSDNSRegistrar binds a generic wrapper to an already deployed contract.
func bindZNSDNSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZNSDNSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSDNSRegistrar *ZNSDNSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSDNSRegistrar.Contract.ZNSDNSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSDNSRegistrar *ZNSDNSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.ZNSDNSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSDNSRegistrar *ZNSDNSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.ZNSDNSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSDNSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Zns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function zns() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCaller) Zns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSDNSRegistrar.contract.Call(opts, &out, "zns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Zns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function zns() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) Zns() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Zns(&_ZNSDNSRegistrar.CallOpts)
}

// Zns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function zns() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerSession) Zns() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Zns(&_ZNSDNSRegistrar.CallOpts)
}

// Inceptions is a free data retrieval call binding the contract method 0x25916d41.
//
// Solidity: function inceptions(bytes32 ) view returns(uint32)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCaller) Inceptions(opts *bind.CallOpts, arg0 [32]byte) (uint32, error) {
	var out []interface{}
	err := _ZNSDNSRegistrar.contract.Call(opts, &out, "inceptions", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Inceptions is a free data retrieval call binding the contract method 0x25916d41.
//
// Solidity: function inceptions(bytes32 ) view returns(uint32)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) Inceptions(arg0 [32]byte) (uint32, error) {
	return _ZNSDNSRegistrar.Contract.Inceptions(&_ZNSDNSRegistrar.CallOpts, arg0)
}

// Inceptions is a free data retrieval call binding the contract method 0x25916d41.
//
// Solidity: function inceptions(bytes32 ) view returns(uint32)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerSession) Inceptions(arg0 [32]byte) (uint32, error) {
	return _ZNSDNSRegistrar.Contract.Inceptions(&_ZNSDNSRegistrar.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSDNSRegistrar.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) Oracle() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Oracle(&_ZNSDNSRegistrar.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerSession) Oracle() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Oracle(&_ZNSDNSRegistrar.CallOpts)
}

// PreviousRegistrar is a free data retrieval call binding the contract method 0xab14ec59.
//
// Solidity: function previousRegistrar() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCaller) PreviousRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSDNSRegistrar.contract.Call(opts, &out, "previousRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousRegistrar is a free data retrieval call binding the contract method 0xab14ec59.
//
// Solidity: function previousRegistrar() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) PreviousRegistrar() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.PreviousRegistrar(&_ZNSDNSRegistrar.CallOpts)
}

// PreviousRegistrar is a free data retrieval call binding the contract method 0xab14ec59.
//
// Solidity: function previousRegistrar() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerSession) PreviousRegistrar() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.PreviousRegistrar(&_ZNSDNSRegistrar.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSDNSRegistrar.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) Resolver() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Resolver(&_ZNSDNSRegistrar.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerSession) Resolver() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Resolver(&_ZNSDNSRegistrar.CallOpts)
}

// Suffixes is a free data retrieval call binding the contract method 0x30349ebe.
//
// Solidity: function suffixes() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCaller) Suffixes(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSDNSRegistrar.contract.Call(opts, &out, "suffixes")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Suffixes is a free data retrieval call binding the contract method 0x30349ebe.
//
// Solidity: function suffixes() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) Suffixes() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Suffixes(&_ZNSDNSRegistrar.CallOpts)
}

// Suffixes is a free data retrieval call binding the contract method 0x30349ebe.
//
// Solidity: function suffixes() view returns(address)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerSession) Suffixes() (common.Address, error) {
	return _ZNSDNSRegistrar.Contract.Suffixes(&_ZNSDNSRegistrar.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ZNSDNSRegistrar.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSDNSRegistrar.Contract.SupportsInterface(&_ZNSDNSRegistrar.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSDNSRegistrar.Contract.SupportsInterface(&_ZNSDNSRegistrar.CallOpts, interfaceID)
}

// EnableNode is a paid mutator transaction binding the contract method 0x6f951221.
//
// Solidity: function enableNode(bytes domain) returns(bytes32 node)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactor) EnableNode(opts *bind.TransactOpts, domain []byte) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.contract.Transact(opts, "enableNode", domain)
}

// EnableNode is a paid mutator transaction binding the contract method 0x6f951221.
//
// Solidity: function enableNode(bytes domain) returns(bytes32 node)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) EnableNode(domain []byte) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.EnableNode(&_ZNSDNSRegistrar.TransactOpts, domain)
}

// EnableNode is a paid mutator transaction binding the contract method 0x6f951221.
//
// Solidity: function enableNode(bytes domain) returns(bytes32 node)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactorSession) EnableNode(domain []byte) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.EnableNode(&_ZNSDNSRegistrar.TransactOpts, domain)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0x29d56630.
//
// Solidity: function proveAndClaim(bytes name, (bytes,bytes)[] input) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactor) ProveAndClaim(opts *bind.TransactOpts, name []byte, input []DNSSECRRSetWithSignature) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.contract.Transact(opts, "proveAndClaim", name, input)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0x29d56630.
//
// Solidity: function proveAndClaim(bytes name, (bytes,bytes)[] input) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) ProveAndClaim(name []byte, input []DNSSECRRSetWithSignature) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.ProveAndClaim(&_ZNSDNSRegistrar.TransactOpts, name, input)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0x29d56630.
//
// Solidity: function proveAndClaim(bytes name, (bytes,bytes)[] input) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactorSession) ProveAndClaim(name []byte, input []DNSSECRRSetWithSignature) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.ProveAndClaim(&_ZNSDNSRegistrar.TransactOpts, name, input)
}

// ProveAndClaimWithResolver is a paid mutator transaction binding the contract method 0x06963218.
//
// Solidity: function proveAndClaimWithResolver(bytes name, (bytes,bytes)[] input, address resolver, address addr) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactor) ProveAndClaimWithResolver(opts *bind.TransactOpts, name []byte, input []DNSSECRRSetWithSignature, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.contract.Transact(opts, "proveAndClaimWithResolver", name, input, resolver, addr)
}

// ProveAndClaimWithResolver is a paid mutator transaction binding the contract method 0x06963218.
//
// Solidity: function proveAndClaimWithResolver(bytes name, (bytes,bytes)[] input, address resolver, address addr) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) ProveAndClaimWithResolver(name []byte, input []DNSSECRRSetWithSignature, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.ProveAndClaimWithResolver(&_ZNSDNSRegistrar.TransactOpts, name, input, resolver, addr)
}

// ProveAndClaimWithResolver is a paid mutator transaction binding the contract method 0x06963218.
//
// Solidity: function proveAndClaimWithResolver(bytes name, (bytes,bytes)[] input, address resolver, address addr) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactorSession) ProveAndClaimWithResolver(name []byte, input []DNSSECRRSetWithSignature, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.ProveAndClaimWithResolver(&_ZNSDNSRegistrar.TransactOpts, name, input, resolver, addr)
}

// SetPublicSuffixList is a paid mutator transaction binding the contract method 0x1ecfc411.
//
// Solidity: function setPublicSuffixList(address _suffixes) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactor) SetPublicSuffixList(opts *bind.TransactOpts, _suffixes common.Address) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.contract.Transact(opts, "setPublicSuffixList", _suffixes)
}

// SetPublicSuffixList is a paid mutator transaction binding the contract method 0x1ecfc411.
//
// Solidity: function setPublicSuffixList(address _suffixes) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarSession) SetPublicSuffixList(_suffixes common.Address) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.SetPublicSuffixList(&_ZNSDNSRegistrar.TransactOpts, _suffixes)
}

// SetPublicSuffixList is a paid mutator transaction binding the contract method 0x1ecfc411.
//
// Solidity: function setPublicSuffixList(address _suffixes) returns()
func (_ZNSDNSRegistrar *ZNSDNSRegistrarTransactorSession) SetPublicSuffixList(_suffixes common.Address) (*types.Transaction, error) {
	return _ZNSDNSRegistrar.Contract.SetPublicSuffixList(&_ZNSDNSRegistrar.TransactOpts, _suffixes)
}

// ZNSDNSRegistrarClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the ZNSDNSRegistrar contract.
type ZNSDNSRegistrarClaimIterator struct {
	Event *ZNSDNSRegistrarClaim // Event containing the contract specifics and raw log

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
func (it *ZNSDNSRegistrarClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSDNSRegistrarClaim)
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
		it.Event = new(ZNSDNSRegistrarClaim)
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
func (it *ZNSDNSRegistrarClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSDNSRegistrarClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSDNSRegistrarClaim represents a Claim event raised by the ZNSDNSRegistrar contract.
type ZNSDNSRegistrarClaim struct {
	Node      [32]byte
	Owner     common.Address
	Dnsname   []byte
	Inception uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x87db02a0e483e2818060eddcbb3488ce44e35aff49a70d92c2aa6c8046cf01e2.
//
// Solidity: event Claim(bytes32 indexed node, address indexed owner, bytes dnsname, uint32 inception)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarFilterer) FilterClaim(opts *bind.FilterOpts, node [][32]byte, owner []common.Address) (*ZNSDNSRegistrarClaimIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZNSDNSRegistrar.contract.FilterLogs(opts, "Claim", nodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ZNSDNSRegistrarClaimIterator{contract: _ZNSDNSRegistrar.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x87db02a0e483e2818060eddcbb3488ce44e35aff49a70d92c2aa6c8046cf01e2.
//
// Solidity: event Claim(bytes32 indexed node, address indexed owner, bytes dnsname, uint32 inception)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *ZNSDNSRegistrarClaim, node [][32]byte, owner []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZNSDNSRegistrar.contract.WatchLogs(opts, "Claim", nodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSDNSRegistrarClaim)
				if err := _ZNSDNSRegistrar.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x87db02a0e483e2818060eddcbb3488ce44e35aff49a70d92c2aa6c8046cf01e2.
//
// Solidity: event Claim(bytes32 indexed node, address indexed owner, bytes dnsname, uint32 inception)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarFilterer) ParseClaim(log types.Log) (*ZNSDNSRegistrarClaim, error) {
	event := new(ZNSDNSRegistrarClaim)
	if err := _ZNSDNSRegistrar.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSDNSRegistrarNewPublicSuffixListIterator is returned from FilterNewPublicSuffixList and is used to iterate over the raw logs and unpacked data for NewPublicSuffixList events raised by the ZNSDNSRegistrar contract.
type ZNSDNSRegistrarNewPublicSuffixListIterator struct {
	Event *ZNSDNSRegistrarNewPublicSuffixList // Event containing the contract specifics and raw log

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
func (it *ZNSDNSRegistrarNewPublicSuffixListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSDNSRegistrarNewPublicSuffixList)
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
		it.Event = new(ZNSDNSRegistrarNewPublicSuffixList)
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
func (it *ZNSDNSRegistrarNewPublicSuffixListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSDNSRegistrarNewPublicSuffixListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSDNSRegistrarNewPublicSuffixList represents a NewPublicSuffixList event raised by the ZNSDNSRegistrar contract.
type ZNSDNSRegistrarNewPublicSuffixList struct {
	Suffixes common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPublicSuffixList is a free log retrieval operation binding the contract event 0x9176b7f47e4504df5e5516c99d90d82ac7cbd49cc77e7f22ba2ac2f2e3a3eba8.
//
// Solidity: event NewPublicSuffixList(address suffixes)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarFilterer) FilterNewPublicSuffixList(opts *bind.FilterOpts) (*ZNSDNSRegistrarNewPublicSuffixListIterator, error) {

	logs, sub, err := _ZNSDNSRegistrar.contract.FilterLogs(opts, "NewPublicSuffixList")
	if err != nil {
		return nil, err
	}
	return &ZNSDNSRegistrarNewPublicSuffixListIterator{contract: _ZNSDNSRegistrar.contract, event: "NewPublicSuffixList", logs: logs, sub: sub}, nil
}

// WatchNewPublicSuffixList is a free log subscription operation binding the contract event 0x9176b7f47e4504df5e5516c99d90d82ac7cbd49cc77e7f22ba2ac2f2e3a3eba8.
//
// Solidity: event NewPublicSuffixList(address suffixes)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarFilterer) WatchNewPublicSuffixList(opts *bind.WatchOpts, sink chan<- *ZNSDNSRegistrarNewPublicSuffixList) (event.Subscription, error) {

	logs, sub, err := _ZNSDNSRegistrar.contract.WatchLogs(opts, "NewPublicSuffixList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSDNSRegistrarNewPublicSuffixList)
				if err := _ZNSDNSRegistrar.contract.UnpackLog(event, "NewPublicSuffixList", log); err != nil {
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

// ParseNewPublicSuffixList is a log parse operation binding the contract event 0x9176b7f47e4504df5e5516c99d90d82ac7cbd49cc77e7f22ba2ac2f2e3a3eba8.
//
// Solidity: event NewPublicSuffixList(address suffixes)
func (_ZNSDNSRegistrar *ZNSDNSRegistrarFilterer) ParseNewPublicSuffixList(log types.Log) (*ZNSDNSRegistrarNewPublicSuffixList, error) {
	event := new(ZNSDNSRegistrarNewPublicSuffixList)
	if err := _ZNSDNSRegistrar.contract.UnpackLog(event, "NewPublicSuffixList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
