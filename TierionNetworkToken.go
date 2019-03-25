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

// EthcontractsABI is the input ABI used to generate the binding from.
const EthcontractsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastMintedAtBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quorumRegisteredBallots\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_nodes\",\"type\":\"address[72]\"},{\"indexed\":false,\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_blockHeight\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodes\",\"type\":\"address[72]\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"signature1\",\"type\":\"bytes\"},{\"name\":\"signature2\",\"type\":\"bytes\"},{\"name\":\"signature3\",\"type\":\"bytes\"},{\"name\":\"signature4\",\"type\":\"bytes\"},{\"name\":\"signature5\",\"type\":\"bytes\"},{\"name\":\"signature6\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setChainpointRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"toEthSignedMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Ethcontracts is an auto generated Go binding around an Ethereum contract.
type Ethcontracts struct {
	EthcontractsCaller     // Read-only binding to the contract
	EthcontractsTransactor // Write-only binding to the contract
	EthcontractsFilterer   // Log filterer for contract events
}

// EthcontractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthcontractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthcontractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthcontractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthcontractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthcontractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthcontractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthcontractsSession struct {
	Contract     *Ethcontracts     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthcontractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthcontractsCallerSession struct {
	Contract *EthcontractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EthcontractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthcontractsTransactorSession struct {
	Contract     *EthcontractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EthcontractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthcontractsRaw struct {
	Contract *Ethcontracts // Generic contract binding to access the raw methods on
}

// EthcontractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthcontractsCallerRaw struct {
	Contract *EthcontractsCaller // Generic read-only contract binding to access the raw methods on
}

// EthcontractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthcontractsTransactorRaw struct {
	Contract *EthcontractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthcontracts creates a new instance of Ethcontracts, bound to a specific deployed contract.
func NewEthcontracts(address common.Address, backend bind.ContractBackend) (*Ethcontracts, error) {
	contract, err := bindEthcontracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethcontracts{EthcontractsCaller: EthcontractsCaller{contract: contract}, EthcontractsTransactor: EthcontractsTransactor{contract: contract}, EthcontractsFilterer: EthcontractsFilterer{contract: contract}}, nil
}

// NewEthcontractsCaller creates a new read-only instance of Ethcontracts, bound to a specific deployed contract.
func NewEthcontractsCaller(address common.Address, caller bind.ContractCaller) (*EthcontractsCaller, error) {
	contract, err := bindEthcontracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthcontractsCaller{contract: contract}, nil
}

// NewEthcontractsTransactor creates a new write-only instance of Ethcontracts, bound to a specific deployed contract.
func NewEthcontractsTransactor(address common.Address, transactor bind.ContractTransactor) (*EthcontractsTransactor, error) {
	contract, err := bindEthcontracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthcontractsTransactor{contract: contract}, nil
}

// NewEthcontractsFilterer creates a new log filterer instance of Ethcontracts, bound to a specific deployed contract.
func NewEthcontractsFilterer(address common.Address, filterer bind.ContractFilterer) (*EthcontractsFilterer, error) {
	contract, err := bindEthcontracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthcontractsFilterer{contract: contract}, nil
}

// bindEthcontracts binds a generic wrapper to an already deployed contract.
func bindEthcontracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthcontractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethcontracts *EthcontractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethcontracts.Contract.EthcontractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethcontracts *EthcontractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.Contract.EthcontractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethcontracts *EthcontractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethcontracts.Contract.EthcontractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethcontracts *EthcontractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethcontracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethcontracts *EthcontractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethcontracts *EthcontractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethcontracts.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) INITIALSUPPLY() (*big.Int, error) {
	return _Ethcontracts.Contract.INITIALSUPPLY(&_Ethcontracts.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _Ethcontracts.Contract.INITIALSUPPLY(&_Ethcontracts.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_Ethcontracts *EthcontractsCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_Ethcontracts *EthcontractsSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Ethcontracts.Contract.Allowance(&_Ethcontracts.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_Ethcontracts *EthcontractsCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Ethcontracts.Contract.Allowance(&_Ethcontracts.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ethcontracts.Contract.Allowed(&_Ethcontracts.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ethcontracts.Contract.Allowed(&_Ethcontracts.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Ethcontracts *EthcontractsCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Ethcontracts *EthcontractsSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ethcontracts.Contract.BalanceOf(&_Ethcontracts.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Ethcontracts *EthcontractsCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Ethcontracts.Contract.BalanceOf(&_Ethcontracts.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Ethcontracts *EthcontractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Ethcontracts *EthcontractsSession) Decimals() (uint8, error) {
	return _Ethcontracts.Contract.Decimals(&_Ethcontracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Ethcontracts *EthcontractsCallerSession) Decimals() (uint8, error) {
	return _Ethcontracts.Contract.Decimals(&_Ethcontracts.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ethcontracts *EthcontractsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ethcontracts *EthcontractsSession) IsOwner() (bool, error) {
	return _Ethcontracts.Contract.IsOwner(&_Ethcontracts.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ethcontracts *EthcontractsCallerSession) IsOwner() (bool, error) {
	return _Ethcontracts.Contract.IsOwner(&_Ethcontracts.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Ethcontracts *EthcontractsCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Ethcontracts *EthcontractsSession) IsPauser(account common.Address) (bool, error) {
	return _Ethcontracts.Contract.IsPauser(&_Ethcontracts.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Ethcontracts *EthcontractsCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Ethcontracts.Contract.IsPauser(&_Ethcontracts.CallOpts, account)
}

// LastMintedAtBlock is a free data retrieval call binding the contract method 0x91bfa2bc.
//
// Solidity: function lastMintedAtBlock() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) LastMintedAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "lastMintedAtBlock")
	return *ret0, err
}

// LastMintedAtBlock is a free data retrieval call binding the contract method 0x91bfa2bc.
//
// Solidity: function lastMintedAtBlock() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) LastMintedAtBlock() (*big.Int, error) {
	return _Ethcontracts.Contract.LastMintedAtBlock(&_Ethcontracts.CallOpts)
}

// LastMintedAtBlock is a free data retrieval call binding the contract method 0x91bfa2bc.
//
// Solidity: function lastMintedAtBlock() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) LastMintedAtBlock() (*big.Int, error) {
	return _Ethcontracts.Contract.LastMintedAtBlock(&_Ethcontracts.CallOpts)
}

// MintAmount is a free data retrieval call binding the contract method 0x5a2bcc18.
//
// Solidity: function mintAmount() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) MintAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "mintAmount")
	return *ret0, err
}

// MintAmount is a free data retrieval call binding the contract method 0x5a2bcc18.
//
// Solidity: function mintAmount() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) MintAmount() (*big.Int, error) {
	return _Ethcontracts.Contract.MintAmount(&_Ethcontracts.CallOpts)
}

// MintAmount is a free data retrieval call binding the contract method 0x5a2bcc18.
//
// Solidity: function mintAmount() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) MintAmount() (*big.Int, error) {
	return _Ethcontracts.Contract.MintAmount(&_Ethcontracts.CallOpts)
}

// MintingInterval is a free data retrieval call binding the contract method 0xa3059ef0.
//
// Solidity: function mintingInterval() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) MintingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "mintingInterval")
	return *ret0, err
}

// MintingInterval is a free data retrieval call binding the contract method 0xa3059ef0.
//
// Solidity: function mintingInterval() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) MintingInterval() (*big.Int, error) {
	return _Ethcontracts.Contract.MintingInterval(&_Ethcontracts.CallOpts)
}

// MintingInterval is a free data retrieval call binding the contract method 0xa3059ef0.
//
// Solidity: function mintingInterval() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) MintingInterval() (*big.Int, error) {
	return _Ethcontracts.Contract.MintingInterval(&_Ethcontracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Ethcontracts *EthcontractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Ethcontracts *EthcontractsSession) Name() (string, error) {
	return _Ethcontracts.Contract.Name(&_Ethcontracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Ethcontracts *EthcontractsCallerSession) Name() (string, error) {
	return _Ethcontracts.Contract.Name(&_Ethcontracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ethcontracts *EthcontractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ethcontracts *EthcontractsSession) Owner() (common.Address, error) {
	return _Ethcontracts.Contract.Owner(&_Ethcontracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ethcontracts *EthcontractsCallerSession) Owner() (common.Address, error) {
	return _Ethcontracts.Contract.Owner(&_Ethcontracts.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Ethcontracts *EthcontractsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Ethcontracts *EthcontractsSession) Paused() (bool, error) {
	return _Ethcontracts.Contract.Paused(&_Ethcontracts.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Ethcontracts *EthcontractsCallerSession) Paused() (bool, error) {
	return _Ethcontracts.Contract.Paused(&_Ethcontracts.CallOpts)
}

// QuorumRegisteredBallots is a free data retrieval call binding the contract method 0xdff11d11.
//
// Solidity: function quorumRegisteredBallots(uint256 ) constant returns(bytes32)
func (_Ethcontracts *EthcontractsCaller) QuorumRegisteredBallots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "quorumRegisteredBallots", arg0)
	return *ret0, err
}

// QuorumRegisteredBallots is a free data retrieval call binding the contract method 0xdff11d11.
//
// Solidity: function quorumRegisteredBallots(uint256 ) constant returns(bytes32)
func (_Ethcontracts *EthcontractsSession) QuorumRegisteredBallots(arg0 *big.Int) ([32]byte, error) {
	return _Ethcontracts.Contract.QuorumRegisteredBallots(&_Ethcontracts.CallOpts, arg0)
}

// QuorumRegisteredBallots is a free data retrieval call binding the contract method 0xdff11d11.
//
// Solidity: function quorumRegisteredBallots(uint256 ) constant returns(bytes32)
func (_Ethcontracts *EthcontractsCallerSession) QuorumRegisteredBallots(arg0 *big.Int) ([32]byte, error) {
	return _Ethcontracts.Contract.QuorumRegisteredBallots(&_Ethcontracts.CallOpts, arg0)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) constant returns(address)
func (_Ethcontracts *EthcontractsCaller) Recover(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "recover", hash, signature)
	return *ret0, err
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) constant returns(address)
func (_Ethcontracts *EthcontractsSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Ethcontracts.Contract.Recover(&_Ethcontracts.CallOpts, hash, signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) constant returns(address)
func (_Ethcontracts *EthcontractsCallerSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Ethcontracts.Contract.Recover(&_Ethcontracts.CallOpts, hash, signature)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Ethcontracts *EthcontractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Ethcontracts *EthcontractsSession) Symbol() (string, error) {
	return _Ethcontracts.Contract.Symbol(&_Ethcontracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Ethcontracts *EthcontractsCallerSession) Symbol() (string, error) {
	return _Ethcontracts.Contract.Symbol(&_Ethcontracts.CallOpts)
}

// ToEthSignedMessageHash is a free data retrieval call binding the contract method 0x918a15cf.
//
// Solidity: function toEthSignedMessageHash(bytes32 hash) constant returns(bytes32)
func (_Ethcontracts *EthcontractsCaller) ToEthSignedMessageHash(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "toEthSignedMessageHash", hash)
	return *ret0, err
}

// ToEthSignedMessageHash is a free data retrieval call binding the contract method 0x918a15cf.
//
// Solidity: function toEthSignedMessageHash(bytes32 hash) constant returns(bytes32)
func (_Ethcontracts *EthcontractsSession) ToEthSignedMessageHash(hash [32]byte) ([32]byte, error) {
	return _Ethcontracts.Contract.ToEthSignedMessageHash(&_Ethcontracts.CallOpts, hash)
}

// ToEthSignedMessageHash is a free data retrieval call binding the contract method 0x918a15cf.
//
// Solidity: function toEthSignedMessageHash(bytes32 hash) constant returns(bytes32)
func (_Ethcontracts *EthcontractsCallerSession) ToEthSignedMessageHash(hash [32]byte) ([32]byte, error) {
	return _Ethcontracts.Contract.ToEthSignedMessageHash(&_Ethcontracts.CallOpts, hash)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) TotalSupply() (*big.Int, error) {
	return _Ethcontracts.Contract.TotalSupply(&_Ethcontracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Ethcontracts.Contract.TotalSupply(&_Ethcontracts.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Ethcontracts *EthcontractsTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Ethcontracts *EthcontractsSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.AddPauser(&_Ethcontracts.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Ethcontracts *EthcontractsTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.AddPauser(&_Ethcontracts.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Approve(&_Ethcontracts.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Approve(&_Ethcontracts.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x5279f455.
//
// Solidity: function mint(address[72] _nodes, bytes32 _hash, bytes signature1, bytes signature2, bytes signature3, bytes signature4, bytes signature5, bytes signature6) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) Mint(opts *bind.TransactOpts, _nodes [72]common.Address, _hash [32]byte, signature1 []byte, signature2 []byte, signature3 []byte, signature4 []byte, signature5 []byte, signature6 []byte) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "mint", _nodes, _hash, signature1, signature2, signature3, signature4, signature5, signature6)
}

// Mint is a paid mutator transaction binding the contract method 0x5279f455.
//
// Solidity: function mint(address[72] _nodes, bytes32 _hash, bytes signature1, bytes signature2, bytes signature3, bytes signature4, bytes signature5, bytes signature6) returns(bool)
func (_Ethcontracts *EthcontractsSession) Mint(_nodes [72]common.Address, _hash [32]byte, signature1 []byte, signature2 []byte, signature3 []byte, signature4 []byte, signature5 []byte, signature6 []byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Mint(&_Ethcontracts.TransactOpts, _nodes, _hash, signature1, signature2, signature3, signature4, signature5, signature6)
}

// Mint is a paid mutator transaction binding the contract method 0x5279f455.
//
// Solidity: function mint(address[72] _nodes, bytes32 _hash, bytes signature1, bytes signature2, bytes signature3, bytes signature4, bytes signature5, bytes signature6) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) Mint(_nodes [72]common.Address, _hash [32]byte, signature1 []byte, signature2 []byte, signature3 []byte, signature4 []byte, signature5 []byte, signature6 []byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Mint(&_Ethcontracts.TransactOpts, _nodes, _hash, signature1, signature2, signature3, signature4, signature5, signature6)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ethcontracts *EthcontractsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ethcontracts *EthcontractsSession) Pause() (*types.Transaction, error) {
	return _Ethcontracts.Contract.Pause(&_Ethcontracts.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ethcontracts *EthcontractsTransactorSession) Pause() (*types.Transaction, error) {
	return _Ethcontracts.Contract.Pause(&_Ethcontracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethcontracts *EthcontractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethcontracts *EthcontractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethcontracts.Contract.RenounceOwnership(&_Ethcontracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethcontracts *EthcontractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethcontracts.Contract.RenounceOwnership(&_Ethcontracts.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Ethcontracts *EthcontractsTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Ethcontracts *EthcontractsSession) RenouncePauser() (*types.Transaction, error) {
	return _Ethcontracts.Contract.RenouncePauser(&_Ethcontracts.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Ethcontracts *EthcontractsTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Ethcontracts.Contract.RenouncePauser(&_Ethcontracts.TransactOpts)
}

// SetChainpointRegistry is a paid mutator transaction binding the contract method 0x44ea8d44.
//
// Solidity: function setChainpointRegistry(address _addr) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) SetChainpointRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "setChainpointRegistry", _addr)
}

// SetChainpointRegistry is a paid mutator transaction binding the contract method 0x44ea8d44.
//
// Solidity: function setChainpointRegistry(address _addr) returns(bool)
func (_Ethcontracts *EthcontractsSession) SetChainpointRegistry(_addr common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.SetChainpointRegistry(&_Ethcontracts.TransactOpts, _addr)
}

// SetChainpointRegistry is a paid mutator transaction binding the contract method 0x44ea8d44.
//
// Solidity: function setChainpointRegistry(address _addr) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) SetChainpointRegistry(_addr common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.SetChainpointRegistry(&_Ethcontracts.TransactOpts, _addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Transfer(&_Ethcontracts.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Transfer(&_Ethcontracts.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.Contract.TransferFrom(&_Ethcontracts.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ethcontracts.Contract.TransferFrom(&_Ethcontracts.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethcontracts *EthcontractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethcontracts *EthcontractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.TransferOwnership(&_Ethcontracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethcontracts *EthcontractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.TransferOwnership(&_Ethcontracts.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ethcontracts *EthcontractsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ethcontracts *EthcontractsSession) Unpause() (*types.Transaction, error) {
	return _Ethcontracts.Contract.Unpause(&_Ethcontracts.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ethcontracts *EthcontractsTransactorSession) Unpause() (*types.Transaction, error) {
	return _Ethcontracts.Contract.Unpause(&_Ethcontracts.TransactOpts)
}

// EthcontractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ethcontracts contract.
type EthcontractsApprovalIterator struct {
	Event *EthcontractsApproval // Event containing the contract specifics and raw log

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
func (it *EthcontractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsApproval)
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
		it.Event = new(EthcontractsApproval)
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
func (it *EthcontractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsApproval represents a Approval event raised by the Ethcontracts contract.
type EthcontractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ethcontracts *EthcontractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EthcontractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsApprovalIterator{contract: _Ethcontracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ethcontracts *EthcontractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EthcontractsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsApproval)
				if err := _Ethcontracts.contract.UnpackLog(event, "Approval", log); err != nil {
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

// EthcontractsMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Ethcontracts contract.
type EthcontractsMintIterator struct {
	Event *EthcontractsMint // Event containing the contract specifics and raw log

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
func (it *EthcontractsMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsMint)
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
		it.Event = new(EthcontractsMint)
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
func (it *EthcontractsMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsMint represents a Mint event raised by the Ethcontracts contract.
type EthcontractsMint struct {
	Nodes       [72]common.Address
	MintAmount  *big.Int
	BlockHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xc03b15d4ff90a71c53eed1e6907bc35129f4e1507337e30a34770ea77963f602.
//
// Solidity: event Mint(address[72] _nodes, uint256 _mintAmount, uint256 _blockHeight)
func (_Ethcontracts *EthcontractsFilterer) FilterMint(opts *bind.FilterOpts) (*EthcontractsMintIterator, error) {

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &EthcontractsMintIterator{contract: _Ethcontracts.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xc03b15d4ff90a71c53eed1e6907bc35129f4e1507337e30a34770ea77963f602.
//
// Solidity: event Mint(address[72] _nodes, uint256 _mintAmount, uint256 _blockHeight)
func (_Ethcontracts *EthcontractsFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *EthcontractsMint) (event.Subscription, error) {

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsMint)
				if err := _Ethcontracts.contract.UnpackLog(event, "Mint", log); err != nil {
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

// EthcontractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ethcontracts contract.
type EthcontractsOwnershipTransferredIterator struct {
	Event *EthcontractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthcontractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsOwnershipTransferred)
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
		it.Event = new(EthcontractsOwnershipTransferred)
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
func (it *EthcontractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsOwnershipTransferred represents a OwnershipTransferred event raised by the Ethcontracts contract.
type EthcontractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethcontracts *EthcontractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthcontractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsOwnershipTransferredIterator{contract: _Ethcontracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethcontracts *EthcontractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthcontractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsOwnershipTransferred)
				if err := _Ethcontracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// EthcontractsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Ethcontracts contract.
type EthcontractsPausedIterator struct {
	Event *EthcontractsPaused // Event containing the contract specifics and raw log

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
func (it *EthcontractsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsPaused)
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
		it.Event = new(EthcontractsPaused)
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
func (it *EthcontractsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsPaused represents a Paused event raised by the Ethcontracts contract.
type EthcontractsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ethcontracts *EthcontractsFilterer) FilterPaused(opts *bind.FilterOpts) (*EthcontractsPausedIterator, error) {

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthcontractsPausedIterator{contract: _Ethcontracts.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ethcontracts *EthcontractsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthcontractsPaused) (event.Subscription, error) {

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsPaused)
				if err := _Ethcontracts.contract.UnpackLog(event, "Paused", log); err != nil {
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

// EthcontractsPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Ethcontracts contract.
type EthcontractsPauserAddedIterator struct {
	Event *EthcontractsPauserAdded // Event containing the contract specifics and raw log

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
func (it *EthcontractsPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsPauserAdded)
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
		it.Event = new(EthcontractsPauserAdded)
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
func (it *EthcontractsPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsPauserAdded represents a PauserAdded event raised by the Ethcontracts contract.
type EthcontractsPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_Ethcontracts *EthcontractsFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*EthcontractsPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsPauserAddedIterator{contract: _Ethcontracts.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_Ethcontracts *EthcontractsFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *EthcontractsPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsPauserAdded)
				if err := _Ethcontracts.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// EthcontractsPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Ethcontracts contract.
type EthcontractsPauserRemovedIterator struct {
	Event *EthcontractsPauserRemoved // Event containing the contract specifics and raw log

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
func (it *EthcontractsPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsPauserRemoved)
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
		it.Event = new(EthcontractsPauserRemoved)
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
func (it *EthcontractsPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsPauserRemoved represents a PauserRemoved event raised by the Ethcontracts contract.
type EthcontractsPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_Ethcontracts *EthcontractsFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*EthcontractsPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsPauserRemovedIterator{contract: _Ethcontracts.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_Ethcontracts *EthcontractsFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *EthcontractsPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsPauserRemoved)
				if err := _Ethcontracts.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// EthcontractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ethcontracts contract.
type EthcontractsTransferIterator struct {
	Event *EthcontractsTransfer // Event containing the contract specifics and raw log

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
func (it *EthcontractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsTransfer)
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
		it.Event = new(EthcontractsTransfer)
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
func (it *EthcontractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsTransfer represents a Transfer event raised by the Ethcontracts contract.
type EthcontractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ethcontracts *EthcontractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EthcontractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsTransferIterator{contract: _Ethcontracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ethcontracts *EthcontractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EthcontractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsTransfer)
				if err := _Ethcontracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// EthcontractsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Ethcontracts contract.
type EthcontractsUnpausedIterator struct {
	Event *EthcontractsUnpaused // Event containing the contract specifics and raw log

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
func (it *EthcontractsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsUnpaused)
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
		it.Event = new(EthcontractsUnpaused)
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
func (it *EthcontractsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsUnpaused represents a Unpaused event raised by the Ethcontracts contract.
type EthcontractsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ethcontracts *EthcontractsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthcontractsUnpausedIterator, error) {

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthcontractsUnpausedIterator{contract: _Ethcontracts.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ethcontracts *EthcontractsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthcontractsUnpaused) (event.Subscription, error) {

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsUnpaused)
				if err := _Ethcontracts.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
