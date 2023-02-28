// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.ConvertType
)

// POGTokenMetaData contains all meta data concerning the POGToken contract.
var POGTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"LogViewerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"LogViewerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_is_subscribed\",\"type\":\"bool\"}],\"name\":\"addViewer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeViewer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"viewers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"is_subscribed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// POGTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use POGTokenMetaData.ABI instead.
var POGTokenABI = POGTokenMetaData.ABI

// POGToken is an auto generated Go binding around an Ethereum contract.
type POGToken struct {
	POGTokenCaller     // Read-only binding to the contract
	POGTokenTransactor // Write-only binding to the contract
	POGTokenFilterer   // Log filterer for contract events
}

// POGTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type POGTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POGTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type POGTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POGTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type POGTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// POGTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type POGTokenSession struct {
	Contract     *POGToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// POGTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type POGTokenCallerSession struct {
	Contract *POGTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// POGTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type POGTokenTransactorSession struct {
	Contract     *POGTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// POGTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type POGTokenRaw struct {
	Contract *POGToken // Generic contract binding to access the raw methods on
}

// POGTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type POGTokenCallerRaw struct {
	Contract *POGTokenCaller // Generic read-only contract binding to access the raw methods on
}

// POGTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type POGTokenTransactorRaw struct {
	Contract *POGTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPOGToken creates a new instance of POGToken, bound to a specific deployed contract.
func NewPOGToken(address common.Address, backend bind.ContractBackend) (*POGToken, error) {
	contract, err := bindPOGToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &POGToken{POGTokenCaller: POGTokenCaller{contract: contract}, POGTokenTransactor: POGTokenTransactor{contract: contract}, POGTokenFilterer: POGTokenFilterer{contract: contract}}, nil
}

// NewPOGTokenCaller creates a new read-only instance of POGToken, bound to a specific deployed contract.
func NewPOGTokenCaller(address common.Address, caller bind.ContractCaller) (*POGTokenCaller, error) {
	contract, err := bindPOGToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &POGTokenCaller{contract: contract}, nil
}

// NewPOGTokenTransactor creates a new write-only instance of POGToken, bound to a specific deployed contract.
func NewPOGTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*POGTokenTransactor, error) {
	contract, err := bindPOGToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &POGTokenTransactor{contract: contract}, nil
}

// NewPOGTokenFilterer creates a new log filterer instance of POGToken, bound to a specific deployed contract.
func NewPOGTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*POGTokenFilterer, error) {
	contract, err := bindPOGToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &POGTokenFilterer{contract: contract}, nil
}

// bindPOGToken binds a generic wrapper to an already deployed contract.
func bindPOGToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := POGTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POGToken *POGTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _POGToken.Contract.POGTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POGToken *POGTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POGToken.Contract.POGTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POGToken *POGTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POGToken.Contract.POGTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_POGToken *POGTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _POGToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_POGToken *POGTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POGToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_POGToken *POGTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _POGToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_POGToken *POGTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_POGToken *POGTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _POGToken.Contract.Allowance(&_POGToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_POGToken *POGTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _POGToken.Contract.Allowance(&_POGToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_POGToken *POGTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_POGToken *POGTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _POGToken.Contract.BalanceOf(&_POGToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_POGToken *POGTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _POGToken.Contract.BalanceOf(&_POGToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_POGToken *POGTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_POGToken *POGTokenSession) Decimals() (uint8, error) {
	return _POGToken.Contract.Decimals(&_POGToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_POGToken *POGTokenCallerSession) Decimals() (uint8, error) {
	return _POGToken.Contract.Decimals(&_POGToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_POGToken *POGTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_POGToken *POGTokenSession) Name() (string, error) {
	return _POGToken.Contract.Name(&_POGToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_POGToken *POGTokenCallerSession) Name() (string, error) {
	return _POGToken.Contract.Name(&_POGToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_POGToken *POGTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_POGToken *POGTokenSession) Owner() (common.Address, error) {
	return _POGToken.Contract.Owner(&_POGToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_POGToken *POGTokenCallerSession) Owner() (common.Address, error) {
	return _POGToken.Contract.Owner(&_POGToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_POGToken *POGTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_POGToken *POGTokenSession) Symbol() (string, error) {
	return _POGToken.Contract.Symbol(&_POGToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_POGToken *POGTokenCallerSession) Symbol() (string, error) {
	return _POGToken.Contract.Symbol(&_POGToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_POGToken *POGTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_POGToken *POGTokenSession) TotalSupply() (*big.Int, error) {
	return _POGToken.Contract.TotalSupply(&_POGToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_POGToken *POGTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _POGToken.Contract.TotalSupply(&_POGToken.CallOpts)
}

// Viewers is a free data retrieval call binding the contract method 0xa69439aa.
//
// Solidity: function viewers(address ) view returns(address addr, string username, bool is_subscribed)
func (_POGToken *POGTokenCaller) Viewers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr         common.Address
	Username     string
	IsSubscribed bool
}, error) {
	var out []interface{}
	err := _POGToken.contract.Call(opts, &out, "viewers", arg0)

	outstruct := new(struct {
		Addr         common.Address
		Username     string
		IsSubscribed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Username = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IsSubscribed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Viewers is a free data retrieval call binding the contract method 0xa69439aa.
//
// Solidity: function viewers(address ) view returns(address addr, string username, bool is_subscribed)
func (_POGToken *POGTokenSession) Viewers(arg0 common.Address) (struct {
	Addr         common.Address
	Username     string
	IsSubscribed bool
}, error) {
	return _POGToken.Contract.Viewers(&_POGToken.CallOpts, arg0)
}

// Viewers is a free data retrieval call binding the contract method 0xa69439aa.
//
// Solidity: function viewers(address ) view returns(address addr, string username, bool is_subscribed)
func (_POGToken *POGTokenCallerSession) Viewers(arg0 common.Address) (struct {
	Addr         common.Address
	Username     string
	IsSubscribed bool
}, error) {
	return _POGToken.Contract.Viewers(&_POGToken.CallOpts, arg0)
}

// AddViewer is a paid mutator transaction binding the contract method 0x19cbf0c1.
//
// Solidity: function addViewer(address _addr, string _username, bool _is_subscribed) returns()
func (_POGToken *POGTokenTransactor) AddViewer(opts *bind.TransactOpts, _addr common.Address, _username string, _is_subscribed bool) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "addViewer", _addr, _username, _is_subscribed)
}

// AddViewer is a paid mutator transaction binding the contract method 0x19cbf0c1.
//
// Solidity: function addViewer(address _addr, string _username, bool _is_subscribed) returns()
func (_POGToken *POGTokenSession) AddViewer(_addr common.Address, _username string, _is_subscribed bool) (*types.Transaction, error) {
	return _POGToken.Contract.AddViewer(&_POGToken.TransactOpts, _addr, _username, _is_subscribed)
}

// AddViewer is a paid mutator transaction binding the contract method 0x19cbf0c1.
//
// Solidity: function addViewer(address _addr, string _username, bool _is_subscribed) returns()
func (_POGToken *POGTokenTransactorSession) AddViewer(_addr common.Address, _username string, _is_subscribed bool) (*types.Transaction, error) {
	return _POGToken.Contract.AddViewer(&_POGToken.TransactOpts, _addr, _username, _is_subscribed)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_POGToken *POGTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_POGToken *POGTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Approve(&_POGToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_POGToken *POGTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Approve(&_POGToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_POGToken *POGTokenTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_POGToken *POGTokenSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Burn(&_POGToken.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_POGToken *POGTokenTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Burn(&_POGToken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_POGToken *POGTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_POGToken *POGTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.DecreaseAllowance(&_POGToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_POGToken *POGTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.DecreaseAllowance(&_POGToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_POGToken *POGTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_POGToken *POGTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.IncreaseAllowance(&_POGToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_POGToken *POGTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.IncreaseAllowance(&_POGToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_POGToken *POGTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_POGToken *POGTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Mint(&_POGToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_POGToken *POGTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Mint(&_POGToken.TransactOpts, account, amount)
}

// RemoveViewer is a paid mutator transaction binding the contract method 0xd1d98914.
//
// Solidity: function removeViewer(address _addr) returns()
func (_POGToken *POGTokenTransactor) RemoveViewer(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "removeViewer", _addr)
}

// RemoveViewer is a paid mutator transaction binding the contract method 0xd1d98914.
//
// Solidity: function removeViewer(address _addr) returns()
func (_POGToken *POGTokenSession) RemoveViewer(_addr common.Address) (*types.Transaction, error) {
	return _POGToken.Contract.RemoveViewer(&_POGToken.TransactOpts, _addr)
}

// RemoveViewer is a paid mutator transaction binding the contract method 0xd1d98914.
//
// Solidity: function removeViewer(address _addr) returns()
func (_POGToken *POGTokenTransactorSession) RemoveViewer(_addr common.Address) (*types.Transaction, error) {
	return _POGToken.Contract.RemoveViewer(&_POGToken.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_POGToken *POGTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_POGToken *POGTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _POGToken.Contract.RenounceOwnership(&_POGToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_POGToken *POGTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _POGToken.Contract.RenounceOwnership(&_POGToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_POGToken *POGTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_POGToken *POGTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Transfer(&_POGToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_POGToken *POGTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.Transfer(&_POGToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_POGToken *POGTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_POGToken *POGTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.TransferFrom(&_POGToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_POGToken *POGTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _POGToken.Contract.TransferFrom(&_POGToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_POGToken *POGTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _POGToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_POGToken *POGTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _POGToken.Contract.TransferOwnership(&_POGToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_POGToken *POGTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _POGToken.Contract.TransferOwnership(&_POGToken.TransactOpts, newOwner)
}

// POGTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the POGToken contract.
type POGTokenApprovalIterator struct {
	Event *POGTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POGTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POGTokenApproval)
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
		it.Event = new(POGTokenApproval)
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
func (it *POGTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POGTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POGTokenApproval represents a Approval event raised by the POGToken contract.
type POGTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_POGToken *POGTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*POGTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _POGToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &POGTokenApprovalIterator{contract: _POGToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_POGToken *POGTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *POGTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _POGToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POGTokenApproval)
				if err := _POGToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_POGToken *POGTokenFilterer) ParseApproval(log types.Log) (*POGTokenApproval, error) {
	event := new(POGTokenApproval)
	if err := _POGToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POGTokenLogViewerAddedIterator is returned from FilterLogViewerAdded and is used to iterate over the raw logs and unpacked data for LogViewerAdded events raised by the POGToken contract.
type POGTokenLogViewerAddedIterator struct {
	Event *POGTokenLogViewerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POGTokenLogViewerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POGTokenLogViewerAdded)
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
		it.Event = new(POGTokenLogViewerAdded)
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
func (it *POGTokenLogViewerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POGTokenLogViewerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POGTokenLogViewerAdded represents a LogViewerAdded event raised by the POGToken contract.
type POGTokenLogViewerAdded struct {
	Addr     common.Address
	Username string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogViewerAdded is a free log retrieval operation binding the contract event 0x9cf180f8bde59d8ad4dd0ee968a65a9a38ac25086bb22167b071c1ac73fd3aa1.
//
// Solidity: event LogViewerAdded(address addr, string username)
func (_POGToken *POGTokenFilterer) FilterLogViewerAdded(opts *bind.FilterOpts) (*POGTokenLogViewerAddedIterator, error) {

	logs, sub, err := _POGToken.contract.FilterLogs(opts, "LogViewerAdded")
	if err != nil {
		return nil, err
	}
	return &POGTokenLogViewerAddedIterator{contract: _POGToken.contract, event: "LogViewerAdded", logs: logs, sub: sub}, nil
}

// WatchLogViewerAdded is a free log subscription operation binding the contract event 0x9cf180f8bde59d8ad4dd0ee968a65a9a38ac25086bb22167b071c1ac73fd3aa1.
//
// Solidity: event LogViewerAdded(address addr, string username)
func (_POGToken *POGTokenFilterer) WatchLogViewerAdded(opts *bind.WatchOpts, sink chan<- *POGTokenLogViewerAdded) (event.Subscription, error) {

	logs, sub, err := _POGToken.contract.WatchLogs(opts, "LogViewerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POGTokenLogViewerAdded)
				if err := _POGToken.contract.UnpackLog(event, "LogViewerAdded", log); err != nil {
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

// ParseLogViewerAdded is a log parse operation binding the contract event 0x9cf180f8bde59d8ad4dd0ee968a65a9a38ac25086bb22167b071c1ac73fd3aa1.
//
// Solidity: event LogViewerAdded(address addr, string username)
func (_POGToken *POGTokenFilterer) ParseLogViewerAdded(log types.Log) (*POGTokenLogViewerAdded, error) {
	event := new(POGTokenLogViewerAdded)
	if err := _POGToken.contract.UnpackLog(event, "LogViewerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POGTokenLogViewerRemovedIterator is returned from FilterLogViewerRemoved and is used to iterate over the raw logs and unpacked data for LogViewerRemoved events raised by the POGToken contract.
type POGTokenLogViewerRemovedIterator struct {
	Event *POGTokenLogViewerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POGTokenLogViewerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POGTokenLogViewerRemoved)
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
		it.Event = new(POGTokenLogViewerRemoved)
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
func (it *POGTokenLogViewerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POGTokenLogViewerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POGTokenLogViewerRemoved represents a LogViewerRemoved event raised by the POGToken contract.
type POGTokenLogViewerRemoved struct {
	Addr     common.Address
	Username string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogViewerRemoved is a free log retrieval operation binding the contract event 0x3b5575bb746832578287d16b048923c64a99b5d29b97d2356cf8e636bba10cd4.
//
// Solidity: event LogViewerRemoved(address addr, string username)
func (_POGToken *POGTokenFilterer) FilterLogViewerRemoved(opts *bind.FilterOpts) (*POGTokenLogViewerRemovedIterator, error) {

	logs, sub, err := _POGToken.contract.FilterLogs(opts, "LogViewerRemoved")
	if err != nil {
		return nil, err
	}
	return &POGTokenLogViewerRemovedIterator{contract: _POGToken.contract, event: "LogViewerRemoved", logs: logs, sub: sub}, nil
}

// WatchLogViewerRemoved is a free log subscription operation binding the contract event 0x3b5575bb746832578287d16b048923c64a99b5d29b97d2356cf8e636bba10cd4.
//
// Solidity: event LogViewerRemoved(address addr, string username)
func (_POGToken *POGTokenFilterer) WatchLogViewerRemoved(opts *bind.WatchOpts, sink chan<- *POGTokenLogViewerRemoved) (event.Subscription, error) {

	logs, sub, err := _POGToken.contract.WatchLogs(opts, "LogViewerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POGTokenLogViewerRemoved)
				if err := _POGToken.contract.UnpackLog(event, "LogViewerRemoved", log); err != nil {
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

// ParseLogViewerRemoved is a log parse operation binding the contract event 0x3b5575bb746832578287d16b048923c64a99b5d29b97d2356cf8e636bba10cd4.
//
// Solidity: event LogViewerRemoved(address addr, string username)
func (_POGToken *POGTokenFilterer) ParseLogViewerRemoved(log types.Log) (*POGTokenLogViewerRemoved, error) {
	event := new(POGTokenLogViewerRemoved)
	if err := _POGToken.contract.UnpackLog(event, "LogViewerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POGTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the POGToken contract.
type POGTokenOwnershipTransferredIterator struct {
	Event *POGTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POGTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POGTokenOwnershipTransferred)
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
		it.Event = new(POGTokenOwnershipTransferred)
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
func (it *POGTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POGTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POGTokenOwnershipTransferred represents a OwnershipTransferred event raised by the POGToken contract.
type POGTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_POGToken *POGTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*POGTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _POGToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &POGTokenOwnershipTransferredIterator{contract: _POGToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_POGToken *POGTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *POGTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _POGToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POGTokenOwnershipTransferred)
				if err := _POGToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_POGToken *POGTokenFilterer) ParseOwnershipTransferred(log types.Log) (*POGTokenOwnershipTransferred, error) {
	event := new(POGTokenOwnershipTransferred)
	if err := _POGToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// POGTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the POGToken contract.
type POGTokenTransferIterator struct {
	Event *POGTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *POGTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(POGTokenTransfer)
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
		it.Event = new(POGTokenTransfer)
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
func (it *POGTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *POGTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// POGTokenTransfer represents a Transfer event raised by the POGToken contract.
type POGTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_POGToken *POGTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*POGTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _POGToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &POGTokenTransferIterator{contract: _POGToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_POGToken *POGTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *POGTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _POGToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(POGTokenTransfer)
				if err := _POGToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_POGToken *POGTokenFilterer) ParseTransfer(log types.Log) (*POGTokenTransfer, error) {
	event := new(POGTokenTransfer)
	if err := _POGToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
