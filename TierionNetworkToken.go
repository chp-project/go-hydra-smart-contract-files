// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethcontracts

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

// TNTABI is the input ABI used to generate the binding from.
const TNTABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"coreLastMintedAtBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodeLastMintedAtBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_nodes\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_blockHeight\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"UsagePurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodes\",\"type\":\"address[]\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"signature1\",\"type\":\"bytes\"},{\"name\":\"signature2\",\"type\":\"bytes\"},{\"name\":\"signature3\",\"type\":\"bytes\"},{\"name\":\"signature4\",\"type\":\"bytes\"},{\"name\":\"signature5\",\"type\":\"bytes\"},{\"name\":\"signature6\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"purchaseUsage\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setChainpointRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"toEthSignedMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TNT is an auto generated Go binding around an Ethereum contract.
type TNT struct {
	TNTCaller     // Read-only binding to the contract
	TNTTransactor // Write-only binding to the contract
	TNTFilterer   // Log filterer for contract events
}

// TNTCaller is an auto generated read-only Go binding around an Ethereum contract.
type TNTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TNTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TNTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TNTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TNTSession struct {
	Contract     *TNT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TNTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TNTCallerSession struct {
	Contract *TNTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TNTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TNTTransactorSession struct {
	Contract     *TNTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TNTRaw is an auto generated low-level Go binding around an Ethereum contract.
type TNTRaw struct {
	Contract *TNT // Generic contract binding to access the raw methods on
}

// TNTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TNTCallerRaw struct {
	Contract *TNTCaller // Generic read-only contract binding to access the raw methods on
}

// TNTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TNTTransactorRaw struct {
	Contract *TNTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTNT creates a new instance of TNT, bound to a specific deployed contract.
func NewTNT(address common.Address, backend bind.ContractBackend) (*TNT, error) {
	contract, err := bindTNT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TNT{TNTCaller: TNTCaller{contract: contract}, TNTTransactor: TNTTransactor{contract: contract}, TNTFilterer: TNTFilterer{contract: contract}}, nil
}

// NewTNTCaller creates a new read-only instance of TNT, bound to a specific deployed contract.
func NewTNTCaller(address common.Address, caller bind.ContractCaller) (*TNTCaller, error) {
	contract, err := bindTNT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TNTCaller{contract: contract}, nil
}

// NewTNTTransactor creates a new write-only instance of TNT, bound to a specific deployed contract.
func NewTNTTransactor(address common.Address, transactor bind.ContractTransactor) (*TNTTransactor, error) {
	contract, err := bindTNT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TNTTransactor{contract: contract}, nil
}

// NewTNTFilterer creates a new log filterer instance of TNT, bound to a specific deployed contract.
func NewTNTFilterer(address common.Address, filterer bind.ContractFilterer) (*TNTFilterer, error) {
	contract, err := bindTNT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TNTFilterer{contract: contract}, nil
}

// bindTNT binds a generic wrapper to an already deployed contract.
func bindTNT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TNTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT *TNTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TNT.Contract.TNTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT *TNTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT.Contract.TNTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT *TNTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT.Contract.TNTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TNT *TNTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TNT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TNT *TNTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TNT *TNTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TNT.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_TNT *TNTCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_TNT *TNTSession) INITIALSUPPLY() (*big.Int, error) {
	return _TNT.Contract.INITIALSUPPLY(&_TNT.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_TNT *TNTCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _TNT.Contract.INITIALSUPPLY(&_TNT.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_TNT *TNTCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_TNT *TNTSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TNT.Contract.Allowance(&_TNT.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_TNT *TNTCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TNT.Contract.Allowance(&_TNT.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_TNT *TNTCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_TNT *TNTSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TNT.Contract.Allowed(&_TNT.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_TNT *TNTCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TNT.Contract.Allowed(&_TNT.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_TNT *TNTCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_TNT *TNTSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TNT.Contract.BalanceOf(&_TNT.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_TNT *TNTCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TNT.Contract.BalanceOf(&_TNT.CallOpts, _owner)
}

// CoreLastMintedAtBlock is a free data retrieval call binding the contract method 0x4c5f25dd.
//
// Solidity: function coreLastMintedAtBlock() constant returns(uint256)
func (_TNT *TNTCaller) CoreLastMintedAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "coreLastMintedAtBlock")
	return *ret0, err
}

// CoreLastMintedAtBlock is a free data retrieval call binding the contract method 0x4c5f25dd.
//
// Solidity: function coreLastMintedAtBlock() constant returns(uint256)
func (_TNT *TNTSession) CoreLastMintedAtBlock() (*big.Int, error) {
	return _TNT.Contract.CoreLastMintedAtBlock(&_TNT.CallOpts)
}

// CoreLastMintedAtBlock is a free data retrieval call binding the contract method 0x4c5f25dd.
//
// Solidity: function coreLastMintedAtBlock() constant returns(uint256)
func (_TNT *TNTCallerSession) CoreLastMintedAtBlock() (*big.Int, error) {
	return _TNT.Contract.CoreLastMintedAtBlock(&_TNT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TNT *TNTCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TNT *TNTSession) Decimals() (uint8, error) {
	return _TNT.Contract.Decimals(&_TNT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TNT *TNTCallerSession) Decimals() (uint8, error) {
	return _TNT.Contract.Decimals(&_TNT.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_TNT *TNTCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_TNT *TNTSession) IsOwner() (bool, error) {
	return _TNT.Contract.IsOwner(&_TNT.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_TNT *TNTCallerSession) IsOwner() (bool, error) {
	return _TNT.Contract.IsOwner(&_TNT.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_TNT *TNTCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_TNT *TNTSession) IsPauser(account common.Address) (bool, error) {
	return _TNT.Contract.IsPauser(&_TNT.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_TNT *TNTCallerSession) IsPauser(account common.Address) (bool, error) {
	return _TNT.Contract.IsPauser(&_TNT.CallOpts, account)
}

// MintAmount is a free data retrieval call binding the contract method 0x5a2bcc18.
//
// Solidity: function mintAmount() constant returns(uint256)
func (_TNT *TNTCaller) MintAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "mintAmount")
	return *ret0, err
}

// MintAmount is a free data retrieval call binding the contract method 0x5a2bcc18.
//
// Solidity: function mintAmount() constant returns(uint256)
func (_TNT *TNTSession) MintAmount() (*big.Int, error) {
	return _TNT.Contract.MintAmount(&_TNT.CallOpts)
}

// MintAmount is a free data retrieval call binding the contract method 0x5a2bcc18.
//
// Solidity: function mintAmount() constant returns(uint256)
func (_TNT *TNTCallerSession) MintAmount() (*big.Int, error) {
	return _TNT.Contract.MintAmount(&_TNT.CallOpts)
}

// MintingInterval is a free data retrieval call binding the contract method 0xa3059ef0.
//
// Solidity: function mintingInterval() constant returns(uint256)
func (_TNT *TNTCaller) MintingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "mintingInterval")
	return *ret0, err
}

// MintingInterval is a free data retrieval call binding the contract method 0xa3059ef0.
//
// Solidity: function mintingInterval() constant returns(uint256)
func (_TNT *TNTSession) MintingInterval() (*big.Int, error) {
	return _TNT.Contract.MintingInterval(&_TNT.CallOpts)
}

// MintingInterval is a free data retrieval call binding the contract method 0xa3059ef0.
//
// Solidity: function mintingInterval() constant returns(uint256)
func (_TNT *TNTCallerSession) MintingInterval() (*big.Int, error) {
	return _TNT.Contract.MintingInterval(&_TNT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TNT *TNTCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TNT *TNTSession) Name() (string, error) {
	return _TNT.Contract.Name(&_TNT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TNT *TNTCallerSession) Name() (string, error) {
	return _TNT.Contract.Name(&_TNT.CallOpts)
}

// NodeLastMintedAtBlock is a free data retrieval call binding the contract method 0x5a8ef28b.
//
// Solidity: function nodeLastMintedAtBlock() constant returns(uint256)
func (_TNT *TNTCaller) NodeLastMintedAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "nodeLastMintedAtBlock")
	return *ret0, err
}

// NodeLastMintedAtBlock is a free data retrieval call binding the contract method 0x5a8ef28b.
//
// Solidity: function nodeLastMintedAtBlock() constant returns(uint256)
func (_TNT *TNTSession) NodeLastMintedAtBlock() (*big.Int, error) {
	return _TNT.Contract.NodeLastMintedAtBlock(&_TNT.CallOpts)
}

// NodeLastMintedAtBlock is a free data retrieval call binding the contract method 0x5a8ef28b.
//
// Solidity: function nodeLastMintedAtBlock() constant returns(uint256)
func (_TNT *TNTCallerSession) NodeLastMintedAtBlock() (*big.Int, error) {
	return _TNT.Contract.NodeLastMintedAtBlock(&_TNT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TNT *TNTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TNT *TNTSession) Owner() (common.Address, error) {
	return _TNT.Contract.Owner(&_TNT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TNT *TNTCallerSession) Owner() (common.Address, error) {
	return _TNT.Contract.Owner(&_TNT.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_TNT *TNTCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_TNT *TNTSession) Paused() (bool, error) {
	return _TNT.Contract.Paused(&_TNT.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_TNT *TNTCallerSession) Paused() (bool, error) {
	return _TNT.Contract.Paused(&_TNT.CallOpts)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) constant returns(address)
func (_TNT *TNTCaller) Recover(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "recover", hash, signature)
	return *ret0, err
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) constant returns(address)
func (_TNT *TNTSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _TNT.Contract.Recover(&_TNT.CallOpts, hash, signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) constant returns(address)
func (_TNT *TNTCallerSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _TNT.Contract.Recover(&_TNT.CallOpts, hash, signature)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TNT *TNTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TNT *TNTSession) Symbol() (string, error) {
	return _TNT.Contract.Symbol(&_TNT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TNT *TNTCallerSession) Symbol() (string, error) {
	return _TNT.Contract.Symbol(&_TNT.CallOpts)
}

// ToEthSignedMessageHash is a free data retrieval call binding the contract method 0x918a15cf.
//
// Solidity: function toEthSignedMessageHash(bytes32 hash) constant returns(bytes32)
func (_TNT *TNTCaller) ToEthSignedMessageHash(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "toEthSignedMessageHash", hash)
	return *ret0, err
}

// ToEthSignedMessageHash is a free data retrieval call binding the contract method 0x918a15cf.
//
// Solidity: function toEthSignedMessageHash(bytes32 hash) constant returns(bytes32)
func (_TNT *TNTSession) ToEthSignedMessageHash(hash [32]byte) ([32]byte, error) {
	return _TNT.Contract.ToEthSignedMessageHash(&_TNT.CallOpts, hash)
}

// ToEthSignedMessageHash is a free data retrieval call binding the contract method 0x918a15cf.
//
// Solidity: function toEthSignedMessageHash(bytes32 hash) constant returns(bytes32)
func (_TNT *TNTCallerSession) ToEthSignedMessageHash(hash [32]byte) ([32]byte, error) {
	return _TNT.Contract.ToEthSignedMessageHash(&_TNT.CallOpts, hash)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TNT *TNTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TNT.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TNT *TNTSession) TotalSupply() (*big.Int, error) {
	return _TNT.Contract.TotalSupply(&_TNT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TNT *TNTCallerSession) TotalSupply() (*big.Int, error) {
	return _TNT.Contract.TotalSupply(&_TNT.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_TNT *TNTTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_TNT *TNTSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _TNT.Contract.AddPauser(&_TNT.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_TNT *TNTTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _TNT.Contract.AddPauser(&_TNT.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_TNT *TNTTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_TNT *TNTSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.Approve(&_TNT.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_TNT *TNTTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.Approve(&_TNT.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x382e6fdb.
//
// Solidity: function mint(address[] _nodes, bytes32 _hash, bytes signature1, bytes signature2, bytes signature3, bytes signature4, bytes signature5, bytes signature6) returns(bool)
func (_TNT *TNTTransactor) Mint(opts *bind.TransactOpts, _nodes []common.Address, _hash [32]byte, signature1 []byte, signature2 []byte, signature3 []byte, signature4 []byte, signature5 []byte, signature6 []byte) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "mint", _nodes, _hash, signature1, signature2, signature3, signature4, signature5, signature6)
}

// Mint is a paid mutator transaction binding the contract method 0x382e6fdb.
//
// Solidity: function mint(address[] _nodes, bytes32 _hash, bytes signature1, bytes signature2, bytes signature3, bytes signature4, bytes signature5, bytes signature6) returns(bool)
func (_TNT *TNTSession) Mint(_nodes []common.Address, _hash [32]byte, signature1 []byte, signature2 []byte, signature3 []byte, signature4 []byte, signature5 []byte, signature6 []byte) (*types.Transaction, error) {
	return _TNT.Contract.Mint(&_TNT.TransactOpts, _nodes, _hash, signature1, signature2, signature3, signature4, signature5, signature6)
}

// Mint is a paid mutator transaction binding the contract method 0x382e6fdb.
//
// Solidity: function mint(address[] _nodes, bytes32 _hash, bytes signature1, bytes signature2, bytes signature3, bytes signature4, bytes signature5, bytes signature6) returns(bool)
func (_TNT *TNTTransactorSession) Mint(_nodes []common.Address, _hash [32]byte, signature1 []byte, signature2 []byte, signature3 []byte, signature4 []byte, signature5 []byte, signature6 []byte) (*types.Transaction, error) {
	return _TNT.Contract.Mint(&_TNT.TransactOpts, _nodes, _hash, signature1, signature2, signature3, signature4, signature5, signature6)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TNT *TNTTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TNT *TNTSession) Pause() (*types.Transaction, error) {
	return _TNT.Contract.Pause(&_TNT.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TNT *TNTTransactorSession) Pause() (*types.Transaction, error) {
	return _TNT.Contract.Pause(&_TNT.TransactOpts)
}

// PurchaseUsage is a paid mutator transaction binding the contract method 0xb55968ec.
//
// Solidity: function purchaseUsage(uint256 _value) returns(bool)
func (_TNT *TNTTransactor) PurchaseUsage(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "purchaseUsage", _value)
}

// PurchaseUsage is a paid mutator transaction binding the contract method 0xb55968ec.
//
// Solidity: function purchaseUsage(uint256 _value) returns(bool)
func (_TNT *TNTSession) PurchaseUsage(_value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.PurchaseUsage(&_TNT.TransactOpts, _value)
}

// PurchaseUsage is a paid mutator transaction binding the contract method 0xb55968ec.
//
// Solidity: function purchaseUsage(uint256 _value) returns(bool)
func (_TNT *TNTTransactorSession) PurchaseUsage(_value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.PurchaseUsage(&_TNT.TransactOpts, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TNT *TNTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TNT *TNTSession) RenounceOwnership() (*types.Transaction, error) {
	return _TNT.Contract.RenounceOwnership(&_TNT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TNT *TNTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TNT.Contract.RenounceOwnership(&_TNT.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_TNT *TNTTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_TNT *TNTSession) RenouncePauser() (*types.Transaction, error) {
	return _TNT.Contract.RenouncePauser(&_TNT.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_TNT *TNTTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _TNT.Contract.RenouncePauser(&_TNT.TransactOpts)
}

// SetChainpointRegistry is a paid mutator transaction binding the contract method 0x44ea8d44.
//
// Solidity: function setChainpointRegistry(address _addr) returns(bool)
func (_TNT *TNTTransactor) SetChainpointRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "setChainpointRegistry", _addr)
}

// SetChainpointRegistry is a paid mutator transaction binding the contract method 0x44ea8d44.
//
// Solidity: function setChainpointRegistry(address _addr) returns(bool)
func (_TNT *TNTSession) SetChainpointRegistry(_addr common.Address) (*types.Transaction, error) {
	return _TNT.Contract.SetChainpointRegistry(&_TNT.TransactOpts, _addr)
}

// SetChainpointRegistry is a paid mutator transaction binding the contract method 0x44ea8d44.
//
// Solidity: function setChainpointRegistry(address _addr) returns(bool)
func (_TNT *TNTTransactorSession) SetChainpointRegistry(_addr common.Address) (*types.Transaction, error) {
	return _TNT.Contract.SetChainpointRegistry(&_TNT.TransactOpts, _addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_TNT *TNTTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_TNT *TNTSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.Transfer(&_TNT.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_TNT *TNTTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.Transfer(&_TNT.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_TNT *TNTTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_TNT *TNTSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.TransferFrom(&_TNT.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_TNT *TNTTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TNT.Contract.TransferFrom(&_TNT.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TNT *TNTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TNT *TNTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TNT.Contract.TransferOwnership(&_TNT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TNT *TNTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TNT.Contract.TransferOwnership(&_TNT.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TNT *TNTTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TNT.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TNT *TNTSession) Unpause() (*types.Transaction, error) {
	return _TNT.Contract.Unpause(&_TNT.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TNT *TNTTransactorSession) Unpause() (*types.Transaction, error) {
	return _TNT.Contract.Unpause(&_TNT.TransactOpts)
}

// TNTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TNT contract.
type TNTApprovalIterator struct {
	Event *TNTApproval // Event containing the contract specifics and raw log

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
func (it *TNTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTApproval)
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
		it.Event = new(TNTApproval)
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
func (it *TNTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTApproval represents a Approval event raised by the TNT contract.
type TNTApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TNT *TNTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TNTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TNT.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TNTApprovalIterator{contract: _TNT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TNT *TNTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TNTApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TNT.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTApproval)
				if err := _TNT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TNTMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the TNT contract.
type TNTMintIterator struct {
	Event *TNTMint // Event containing the contract specifics and raw log

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
func (it *TNTMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTMint)
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
		it.Event = new(TNTMint)
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
func (it *TNTMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTMint represents a Mint event raised by the TNT contract.
type TNTMint struct {
	Nodes       []common.Address
	MintAmount  *big.Int
	BlockHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x35cc581d6103df194c8c20f437d4b065c9107097eb2ba6f9702489434e03b64c.
//
// Solidity: event Mint(address[] _nodes, uint256 _mintAmount, uint256 _blockHeight)
func (_TNT *TNTFilterer) FilterMint(opts *bind.FilterOpts) (*TNTMintIterator, error) {

	logs, sub, err := _TNT.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &TNTMintIterator{contract: _TNT.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x35cc581d6103df194c8c20f437d4b065c9107097eb2ba6f9702489434e03b64c.
//
// Solidity: event Mint(address[] _nodes, uint256 _mintAmount, uint256 _blockHeight)
func (_TNT *TNTFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TNTMint) (event.Subscription, error) {

	logs, sub, err := _TNT.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTMint)
				if err := _TNT.contract.UnpackLog(event, "Mint", log); err != nil {
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

// TNTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TNT contract.
type TNTOwnershipTransferredIterator struct {
	Event *TNTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TNTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTOwnershipTransferred)
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
		it.Event = new(TNTOwnershipTransferred)
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
func (it *TNTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTOwnershipTransferred represents a OwnershipTransferred event raised by the TNT contract.
type TNTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TNT *TNTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TNTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TNT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TNTOwnershipTransferredIterator{contract: _TNT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TNT *TNTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TNTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TNT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTOwnershipTransferred)
				if err := _TNT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TNTPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TNT contract.
type TNTPausedIterator struct {
	Event *TNTPaused // Event containing the contract specifics and raw log

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
func (it *TNTPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTPaused)
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
		it.Event = new(TNTPaused)
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
func (it *TNTPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTPaused represents a Paused event raised by the TNT contract.
type TNTPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TNT *TNTFilterer) FilterPaused(opts *bind.FilterOpts) (*TNTPausedIterator, error) {

	logs, sub, err := _TNT.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TNTPausedIterator{contract: _TNT.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TNT *TNTFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TNTPaused) (event.Subscription, error) {

	logs, sub, err := _TNT.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTPaused)
				if err := _TNT.contract.UnpackLog(event, "Paused", log); err != nil {
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

// TNTPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the TNT contract.
type TNTPauserAddedIterator struct {
	Event *TNTPauserAdded // Event containing the contract specifics and raw log

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
func (it *TNTPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTPauserAdded)
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
		it.Event = new(TNTPauserAdded)
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
func (it *TNTPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTPauserAdded represents a PauserAdded event raised by the TNT contract.
type TNTPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_TNT *TNTFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*TNTPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TNT.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &TNTPauserAddedIterator{contract: _TNT.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_TNT *TNTFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *TNTPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TNT.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTPauserAdded)
				if err := _TNT.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// TNTPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the TNT contract.
type TNTPauserRemovedIterator struct {
	Event *TNTPauserRemoved // Event containing the contract specifics and raw log

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
func (it *TNTPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTPauserRemoved)
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
		it.Event = new(TNTPauserRemoved)
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
func (it *TNTPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTPauserRemoved represents a PauserRemoved event raised by the TNT contract.
type TNTPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_TNT *TNTFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*TNTPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TNT.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &TNTPauserRemovedIterator{contract: _TNT.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_TNT *TNTFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *TNTPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TNT.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTPauserRemoved)
				if err := _TNT.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// TNTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TNT contract.
type TNTTransferIterator struct {
	Event *TNTTransfer // Event containing the contract specifics and raw log

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
func (it *TNTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTTransfer)
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
		it.Event = new(TNTTransfer)
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
func (it *TNTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTTransfer represents a Transfer event raised by the TNT contract.
type TNTTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TNT *TNTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TNTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TNT.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TNTTransferIterator{contract: _TNT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TNT *TNTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TNTTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TNT.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTTransfer)
				if err := _TNT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TNTUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TNT contract.
type TNTUnpausedIterator struct {
	Event *TNTUnpaused // Event containing the contract specifics and raw log

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
func (it *TNTUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTUnpaused)
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
		it.Event = new(TNTUnpaused)
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
func (it *TNTUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTUnpaused represents a Unpaused event raised by the TNT contract.
type TNTUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TNT *TNTFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TNTUnpausedIterator, error) {

	logs, sub, err := _TNT.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TNTUnpausedIterator{contract: _TNT.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TNT *TNTFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TNTUnpaused) (event.Subscription, error) {

	logs, sub, err := _TNT.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTUnpaused)
				if err := _TNT.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// TNTUsagePurchasedIterator is returned from FilterUsagePurchased and is used to iterate over the raw logs and unpacked data for UsagePurchased events raised by the TNT contract.
type TNTUsagePurchasedIterator struct {
	Event *TNTUsagePurchased // Event containing the contract specifics and raw log

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
func (it *TNTUsagePurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TNTUsagePurchased)
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
		it.Event = new(TNTUsagePurchased)
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
func (it *TNTUsagePurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TNTUsagePurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TNTUsagePurchased represents a UsagePurchased event raised by the TNT contract.
type TNTUsagePurchased struct {
	Buyer common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUsagePurchased is a free log retrieval operation binding the contract event 0x0651f693d262f8b7c2f4a40577914247c33a4eb3fa57d735c8ac49be550021cb.
//
// Solidity: event UsagePurchased(address _buyer, uint256 _value)
func (_TNT *TNTFilterer) FilterUsagePurchased(opts *bind.FilterOpts) (*TNTUsagePurchasedIterator, error) {

	logs, sub, err := _TNT.contract.FilterLogs(opts, "UsagePurchased")
	if err != nil {
		return nil, err
	}
	return &TNTUsagePurchasedIterator{contract: _TNT.contract, event: "UsagePurchased", logs: logs, sub: sub}, nil
}

// WatchUsagePurchased is a free log subscription operation binding the contract event 0x0651f693d262f8b7c2f4a40577914247c33a4eb3fa57d735c8ac49be550021cb.
//
// Solidity: event UsagePurchased(address _buyer, uint256 _value)
func (_TNT *TNTFilterer) WatchUsagePurchased(opts *bind.WatchOpts, sink chan<- *TNTUsagePurchased) (event.Subscription, error) {

	logs, sub, err := _TNT.contract.WatchLogs(opts, "UsagePurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TNTUsagePurchased)
				if err := _TNT.contract.UnpackLog(event, "UsagePurchased", log); err != nil {
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
