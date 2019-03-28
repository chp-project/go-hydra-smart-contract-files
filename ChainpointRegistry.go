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

// ChpRegistryABI is the input ABI used to generate the binding from.
const ChpRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"name\":\"nodeIp\",\"type\":\"bytes32\"},{\"name\":\"nodePublicKey\",\"type\":\"bytes32\"},{\"name\":\"isStaked\",\"type\":\"bool\"},{\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"name\":\"stakeLockedUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedCoresArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CORE_STAKING_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"cores\",\"outputs\":[{\"name\":\"coreIp\",\"type\":\"bytes32\"},{\"name\":\"corePublicKey\",\"type\":\"bytes32\"},{\"name\":\"isStaked\",\"type\":\"bool\"},{\"name\":\"isHealthy\",\"type\":\"bool\"},{\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"name\":\"stakeLockedUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NODE_STAKING_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coresArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CORE_STAKING_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NODE_STAKING_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"NodeStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_publicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"NodeStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_corePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_isHealthy\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"CoreStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_corePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_isHealthy\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"CoreStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"name\":\"_corePublicKey\",\"type\":\"bytes32\"}],\"name\":\"stakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"}],\"name\":\"updateStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"name\":\"_corePublicKey\",\"type\":\"bytes32\"}],\"name\":\"updateStakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unStakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"totalStakedFor\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"unlocks_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCoreCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isHealthyCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"whitelistCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChpRegistry is an auto generated Go binding around an Ethereum contract.
type ChpRegistry struct {
	ChpRegistryCaller     // Read-only binding to the contract
	ChpRegistryTransactor // Write-only binding to the contract
	ChpRegistryFilterer   // Log filterer for contract events
}

// ChpRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChpRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChpRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChpRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChpRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChpRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChpRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChpRegistrySession struct {
	Contract     *ChpRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChpRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChpRegistryCallerSession struct {
	Contract *ChpRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChpRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChpRegistryTransactorSession struct {
	Contract     *ChpRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChpRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChpRegistryRaw struct {
	Contract *ChpRegistry // Generic contract binding to access the raw methods on
}

// ChpRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChpRegistryCallerRaw struct {
	Contract *ChpRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ChpRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChpRegistryTransactorRaw struct {
	Contract *ChpRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChpRegistry creates a new instance of ChpRegistry, bound to a specific deployed contract.
func NewChpRegistry(address common.Address, backend bind.ContractBackend) (*ChpRegistry, error) {
	contract, err := bindChpRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChpRegistry{ChpRegistryCaller: ChpRegistryCaller{contract: contract}, ChpRegistryTransactor: ChpRegistryTransactor{contract: contract}, ChpRegistryFilterer: ChpRegistryFilterer{contract: contract}}, nil
}

// NewChpRegistryCaller creates a new read-only instance of ChpRegistry, bound to a specific deployed contract.
func NewChpRegistryCaller(address common.Address, caller bind.ContractCaller) (*ChpRegistryCaller, error) {
	contract, err := bindChpRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryCaller{contract: contract}, nil
}

// NewChpRegistryTransactor creates a new write-only instance of ChpRegistry, bound to a specific deployed contract.
func NewChpRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChpRegistryTransactor, error) {
	contract, err := bindChpRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryTransactor{contract: contract}, nil
}

// NewChpRegistryFilterer creates a new log filterer instance of ChpRegistry, bound to a specific deployed contract.
func NewChpRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChpRegistryFilterer, error) {
	contract, err := bindChpRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryFilterer{contract: contract}, nil
}

// bindChpRegistry binds a generic wrapper to an already deployed contract.
func bindChpRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChpRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChpRegistry *ChpRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChpRegistry.Contract.ChpRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChpRegistry *ChpRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.Contract.ChpRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChpRegistry *ChpRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChpRegistry.Contract.ChpRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChpRegistry *ChpRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChpRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChpRegistry *ChpRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChpRegistry *ChpRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChpRegistry.Contract.contract.Transact(opts, method, params...)
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCaller) CORESTAKINGAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "CORE_STAKING_AMOUNT")
	return *ret0, err
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_ChpRegistry *ChpRegistrySession) CORESTAKINGAMOUNT() (*big.Int, error) {
	return _ChpRegistry.Contract.CORESTAKINGAMOUNT(&_ChpRegistry.CallOpts)
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCallerSession) CORESTAKINGAMOUNT() (*big.Int, error) {
	return _ChpRegistry.Contract.CORESTAKINGAMOUNT(&_ChpRegistry.CallOpts)
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCaller) CORESTAKINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "CORE_STAKING_DURATION")
	return *ret0, err
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_ChpRegistry *ChpRegistrySession) CORESTAKINGDURATION() (*big.Int, error) {
	return _ChpRegistry.Contract.CORESTAKINGDURATION(&_ChpRegistry.CallOpts)
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCallerSession) CORESTAKINGDURATION() (*big.Int, error) {
	return _ChpRegistry.Contract.CORESTAKINGDURATION(&_ChpRegistry.CallOpts)
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCaller) NODESTAKINGAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "NODE_STAKING_AMOUNT")
	return *ret0, err
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_ChpRegistry *ChpRegistrySession) NODESTAKINGAMOUNT() (*big.Int, error) {
	return _ChpRegistry.Contract.NODESTAKINGAMOUNT(&_ChpRegistry.CallOpts)
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCallerSession) NODESTAKINGAMOUNT() (*big.Int, error) {
	return _ChpRegistry.Contract.NODESTAKINGAMOUNT(&_ChpRegistry.CallOpts)
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCaller) NODESTAKINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "NODE_STAKING_DURATION")
	return *ret0, err
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_ChpRegistry *ChpRegistrySession) NODESTAKINGDURATION() (*big.Int, error) {
	return _ChpRegistry.Contract.NODESTAKINGDURATION(&_ChpRegistry.CallOpts)
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCallerSession) NODESTAKINGDURATION() (*big.Int, error) {
	return _ChpRegistry.Contract.NODESTAKINGDURATION(&_ChpRegistry.CallOpts)
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_ChpRegistry *ChpRegistryCaller) Cores(opts *bind.CallOpts, arg0 common.Address) (struct {
	CoreIp           [32]byte
	CorePublicKey    [32]byte
	IsStaked         bool
	IsHealthy        bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	ret := new(struct {
		CoreIp           [32]byte
		CorePublicKey    [32]byte
		IsStaked         bool
		IsHealthy        bool
		AmountStaked     *big.Int
		StakeLockedUntil *big.Int
	})
	out := ret
	err := _ChpRegistry.contract.Call(opts, out, "cores", arg0)
	return *ret, err
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_ChpRegistry *ChpRegistrySession) Cores(arg0 common.Address) (struct {
	CoreIp           [32]byte
	CorePublicKey    [32]byte
	IsStaked         bool
	IsHealthy        bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _ChpRegistry.Contract.Cores(&_ChpRegistry.CallOpts, arg0)
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_ChpRegistry *ChpRegistryCallerSession) Cores(arg0 common.Address) (struct {
	CoreIp           [32]byte
	CorePublicKey    [32]byte
	IsStaked         bool
	IsHealthy        bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _ChpRegistry.Contract.Cores(&_ChpRegistry.CallOpts, arg0)
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistryCaller) CoresArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "coresArr", arg0)
	return *ret0, err
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistrySession) CoresArr(arg0 *big.Int) (common.Address, error) {
	return _ChpRegistry.Contract.CoresArr(&_ChpRegistry.CallOpts, arg0)
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistryCallerSession) CoresArr(arg0 *big.Int) (common.Address, error) {
	return _ChpRegistry.Contract.CoresArr(&_ChpRegistry.CallOpts, arg0)
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCaller) GetCoreCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "getCoreCount")
	return *ret0, err
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_ChpRegistry *ChpRegistrySession) GetCoreCount() (*big.Int, error) {
	return _ChpRegistry.Contract.GetCoreCount(&_ChpRegistry.CallOpts)
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_ChpRegistry *ChpRegistryCallerSession) GetCoreCount() (*big.Int, error) {
	return _ChpRegistry.Contract.GetCoreCount(&_ChpRegistry.CallOpts)
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_ChpRegistry *ChpRegistryCaller) IsHealthyCore(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "isHealthyCore", _address)
	return *ret0, err
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_ChpRegistry *ChpRegistrySession) IsHealthyCore(_address common.Address) (bool, error) {
	return _ChpRegistry.Contract.IsHealthyCore(&_ChpRegistry.CallOpts, _address)
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_ChpRegistry *ChpRegistryCallerSession) IsHealthyCore(_address common.Address) (bool, error) {
	return _ChpRegistry.Contract.IsHealthyCore(&_ChpRegistry.CallOpts, _address)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ChpRegistry *ChpRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ChpRegistry *ChpRegistrySession) IsOwner() (bool, error) {
	return _ChpRegistry.Contract.IsOwner(&_ChpRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ChpRegistry *ChpRegistryCallerSession) IsOwner() (bool, error) {
	return _ChpRegistry.Contract.IsOwner(&_ChpRegistry.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_ChpRegistry *ChpRegistryCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_ChpRegistry *ChpRegistrySession) IsPauser(account common.Address) (bool, error) {
	return _ChpRegistry.Contract.IsPauser(&_ChpRegistry.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_ChpRegistry *ChpRegistryCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ChpRegistry.Contract.IsPauser(&_ChpRegistry.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ChpRegistry *ChpRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ChpRegistry *ChpRegistrySession) Name() (string, error) {
	return _ChpRegistry.Contract.Name(&_ChpRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ChpRegistry *ChpRegistryCallerSession) Name() (string, error) {
	return _ChpRegistry.Contract.Name(&_ChpRegistry.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_ChpRegistry *ChpRegistryCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeIp           [32]byte
	NodePublicKey    [32]byte
	IsStaked         bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	ret := new(struct {
		NodeIp           [32]byte
		NodePublicKey    [32]byte
		IsStaked         bool
		AmountStaked     *big.Int
		StakeLockedUntil *big.Int
	})
	out := ret
	err := _ChpRegistry.contract.Call(opts, out, "nodes", arg0)
	return *ret, err
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_ChpRegistry *ChpRegistrySession) Nodes(arg0 common.Address) (struct {
	NodeIp           [32]byte
	NodePublicKey    [32]byte
	IsStaked         bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _ChpRegistry.Contract.Nodes(&_ChpRegistry.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_ChpRegistry *ChpRegistryCallerSession) Nodes(arg0 common.Address) (struct {
	NodeIp           [32]byte
	NodePublicKey    [32]byte
	IsStaked         bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _ChpRegistry.Contract.Nodes(&_ChpRegistry.CallOpts, arg0)
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistryCaller) NodesArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "nodesArr", arg0)
	return *ret0, err
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistrySession) NodesArr(arg0 *big.Int) (common.Address, error) {
	return _ChpRegistry.Contract.NodesArr(&_ChpRegistry.CallOpts, arg0)
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistryCallerSession) NodesArr(arg0 *big.Int) (common.Address, error) {
	return _ChpRegistry.Contract.NodesArr(&_ChpRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChpRegistry *ChpRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChpRegistry *ChpRegistrySession) Owner() (common.Address, error) {
	return _ChpRegistry.Contract.Owner(&_ChpRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChpRegistry *ChpRegistryCallerSession) Owner() (common.Address, error) {
	return _ChpRegistry.Contract.Owner(&_ChpRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ChpRegistry *ChpRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ChpRegistry *ChpRegistrySession) Paused() (bool, error) {
	return _ChpRegistry.Contract.Paused(&_ChpRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ChpRegistry *ChpRegistryCallerSession) Paused() (bool, error) {
	return _ChpRegistry.Contract.Paused(&_ChpRegistry.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_ChpRegistry *ChpRegistryCaller) RegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "registryAddress")
	return *ret0, err
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_ChpRegistry *ChpRegistrySession) RegistryAddress() (common.Address, error) {
	return _ChpRegistry.Contract.RegistryAddress(&_ChpRegistry.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_ChpRegistry *ChpRegistryCallerSession) RegistryAddress() (common.Address, error) {
	return _ChpRegistry.Contract.RegistryAddress(&_ChpRegistry.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_ChpRegistry *ChpRegistryCaller) Reward(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "reward")
	return *ret0, err
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_ChpRegistry *ChpRegistrySession) Reward() (bool, error) {
	return _ChpRegistry.Contract.Reward(&_ChpRegistry.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_ChpRegistry *ChpRegistryCallerSession) Reward() (bool, error) {
	return _ChpRegistry.Contract.Reward(&_ChpRegistry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChpRegistry *ChpRegistryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChpRegistry *ChpRegistrySession) Token() (common.Address, error) {
	return _ChpRegistry.Contract.Token(&_ChpRegistry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChpRegistry *ChpRegistryCallerSession) Token() (common.Address, error) {
	return _ChpRegistry.Contract.Token(&_ChpRegistry.CallOpts)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_ChpRegistry *ChpRegistryCaller) TotalStakedFor(opts *bind.CallOpts, addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	ret := new(struct {
		Amount    *big.Int
		UnlocksAt *big.Int
	})
	out := ret
	err := _ChpRegistry.contract.Call(opts, out, "totalStakedFor", addr)
	return *ret, err
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_ChpRegistry *ChpRegistrySession) TotalStakedFor(addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	return _ChpRegistry.Contract.TotalStakedFor(&_ChpRegistry.CallOpts, addr)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_ChpRegistry *ChpRegistryCallerSession) TotalStakedFor(addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	return _ChpRegistry.Contract.TotalStakedFor(&_ChpRegistry.CallOpts, addr)
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistryCaller) WhitelistedCoresArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChpRegistry.contract.Call(opts, out, "whitelistedCoresArr", arg0)
	return *ret0, err
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistrySession) WhitelistedCoresArr(arg0 *big.Int) (common.Address, error) {
	return _ChpRegistry.Contract.WhitelistedCoresArr(&_ChpRegistry.CallOpts, arg0)
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_ChpRegistry *ChpRegistryCallerSession) WhitelistedCoresArr(arg0 *big.Int) (common.Address, error) {
	return _ChpRegistry.Contract.WhitelistedCoresArr(&_ChpRegistry.CallOpts, arg0)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ChpRegistry *ChpRegistryTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ChpRegistry *ChpRegistrySession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ChpRegistry.Contract.AddPauser(&_ChpRegistry.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ChpRegistry *ChpRegistryTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ChpRegistry.Contract.AddPauser(&_ChpRegistry.TransactOpts, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ChpRegistry *ChpRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ChpRegistry *ChpRegistrySession) Pause() (*types.Transaction, error) {
	return _ChpRegistry.Contract.Pause(&_ChpRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ChpRegistry *ChpRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _ChpRegistry.Contract.Pause(&_ChpRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChpRegistry *ChpRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChpRegistry *ChpRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ChpRegistry.Contract.RenounceOwnership(&_ChpRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChpRegistry *ChpRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChpRegistry.Contract.RenounceOwnership(&_ChpRegistry.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ChpRegistry *ChpRegistryTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ChpRegistry *ChpRegistrySession) RenouncePauser() (*types.Transaction, error) {
	return _ChpRegistry.Contract.RenouncePauser(&_ChpRegistry.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ChpRegistry *ChpRegistryTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ChpRegistry.Contract.RenouncePauser(&_ChpRegistry.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactor) Stake(opts *bind.TransactOpts, _nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "stake", _nodeIp, _nodePublicKey)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistrySession) Stake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.Stake(&_ChpRegistry.TransactOpts, _nodeIp, _nodePublicKey)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactorSession) Stake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.Stake(&_ChpRegistry.TransactOpts, _nodeIp, _nodePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactor) StakeCore(opts *bind.TransactOpts, _coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "stakeCore", _coreIp, _corePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistrySession) StakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.StakeCore(&_ChpRegistry.TransactOpts, _coreIp, _corePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactorSession) StakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.StakeCore(&_ChpRegistry.TransactOpts, _coreIp, _corePublicKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChpRegistry *ChpRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChpRegistry *ChpRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChpRegistry.Contract.TransferOwnership(&_ChpRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChpRegistry *ChpRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChpRegistry.Contract.TransferOwnership(&_ChpRegistry.TransactOpts, newOwner)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_ChpRegistry *ChpRegistryTransactor) UnStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "unStake")
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_ChpRegistry *ChpRegistrySession) UnStake() (*types.Transaction, error) {
	return _ChpRegistry.Contract.UnStake(&_ChpRegistry.TransactOpts)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_ChpRegistry *ChpRegistryTransactorSession) UnStake() (*types.Transaction, error) {
	return _ChpRegistry.Contract.UnStake(&_ChpRegistry.TransactOpts)
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_ChpRegistry *ChpRegistryTransactor) UnStakeCore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "unStakeCore")
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_ChpRegistry *ChpRegistrySession) UnStakeCore() (*types.Transaction, error) {
	return _ChpRegistry.Contract.UnStakeCore(&_ChpRegistry.TransactOpts)
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_ChpRegistry *ChpRegistryTransactorSession) UnStakeCore() (*types.Transaction, error) {
	return _ChpRegistry.Contract.UnStakeCore(&_ChpRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ChpRegistry *ChpRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ChpRegistry *ChpRegistrySession) Unpause() (*types.Transaction, error) {
	return _ChpRegistry.Contract.Unpause(&_ChpRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ChpRegistry *ChpRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _ChpRegistry.Contract.Unpause(&_ChpRegistry.TransactOpts)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactor) UpdateStake(opts *bind.TransactOpts, _nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "updateStake", _nodeIp, _nodePublicKey)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistrySession) UpdateStake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.UpdateStake(&_ChpRegistry.TransactOpts, _nodeIp, _nodePublicKey)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactorSession) UpdateStake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.UpdateStake(&_ChpRegistry.TransactOpts, _nodeIp, _nodePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactor) UpdateStakeCore(opts *bind.TransactOpts, _coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "updateStakeCore", _coreIp, _corePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistrySession) UpdateStakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.UpdateStakeCore(&_ChpRegistry.TransactOpts, _coreIp, _corePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_ChpRegistry *ChpRegistryTransactorSession) UpdateStakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _ChpRegistry.Contract.UpdateStakeCore(&_ChpRegistry.TransactOpts, _coreIp, _corePublicKey)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_ChpRegistry *ChpRegistryTransactor) WhitelistCore(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ChpRegistry.contract.Transact(opts, "whitelistCore", _address)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_ChpRegistry *ChpRegistrySession) WhitelistCore(_address common.Address) (*types.Transaction, error) {
	return _ChpRegistry.Contract.WhitelistCore(&_ChpRegistry.TransactOpts, _address)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_ChpRegistry *ChpRegistryTransactorSession) WhitelistCore(_address common.Address) (*types.Transaction, error) {
	return _ChpRegistry.Contract.WhitelistCore(&_ChpRegistry.TransactOpts, _address)
}

// ChpRegistryCoreStakeUpdatedIterator is returned from FilterCoreStakeUpdated and is used to iterate over the raw logs and unpacked data for CoreStakeUpdated events raised by the ChpRegistry contract.
type ChpRegistryCoreStakeUpdatedIterator struct {
	Event *ChpRegistryCoreStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ChpRegistryCoreStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryCoreStakeUpdated)
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
		it.Event = new(ChpRegistryCoreStakeUpdated)
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
func (it *ChpRegistryCoreStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryCoreStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryCoreStakeUpdated represents a CoreStakeUpdated event raised by the ChpRegistry contract.
type ChpRegistryCoreStakeUpdated struct {
	Sender        common.Address
	CoreIp        [32]byte
	CorePublicKey [32]byte
	IsHealthy     bool
	AmountStaked  *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCoreStakeUpdated is a free log retrieval operation binding the contract event 0x4459ca768ace1f4b62fc2885f066e940d9ebd0ceaf21bb1cf430576bd72c5bad.
//
// Solidity: event CoreStakeUpdated(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) FilterCoreStakeUpdated(opts *bind.FilterOpts, _sender []common.Address) (*ChpRegistryCoreStakeUpdatedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "CoreStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryCoreStakeUpdatedIterator{contract: _ChpRegistry.contract, event: "CoreStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchCoreStakeUpdated is a free log subscription operation binding the contract event 0x4459ca768ace1f4b62fc2885f066e940d9ebd0ceaf21bb1cf430576bd72c5bad.
//
// Solidity: event CoreStakeUpdated(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) WatchCoreStakeUpdated(opts *bind.WatchOpts, sink chan<- *ChpRegistryCoreStakeUpdated, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "CoreStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryCoreStakeUpdated)
				if err := _ChpRegistry.contract.UnpackLog(event, "CoreStakeUpdated", log); err != nil {
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

// ChpRegistryCoreStakedIterator is returned from FilterCoreStaked and is used to iterate over the raw logs and unpacked data for CoreStaked events raised by the ChpRegistry contract.
type ChpRegistryCoreStakedIterator struct {
	Event *ChpRegistryCoreStaked // Event containing the contract specifics and raw log

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
func (it *ChpRegistryCoreStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryCoreStaked)
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
		it.Event = new(ChpRegistryCoreStaked)
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
func (it *ChpRegistryCoreStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryCoreStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryCoreStaked represents a CoreStaked event raised by the ChpRegistry contract.
type ChpRegistryCoreStaked struct {
	Sender        common.Address
	CoreIp        [32]byte
	CorePublicKey [32]byte
	IsHealthy     bool
	AmountStaked  *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCoreStaked is a free log retrieval operation binding the contract event 0xe9d817a2e042dbb5a1e846f677db950e6146e460e01eed27fe04631480ac782f.
//
// Solidity: event CoreStaked(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) FilterCoreStaked(opts *bind.FilterOpts, _sender []common.Address) (*ChpRegistryCoreStakedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "CoreStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryCoreStakedIterator{contract: _ChpRegistry.contract, event: "CoreStaked", logs: logs, sub: sub}, nil
}

// WatchCoreStaked is a free log subscription operation binding the contract event 0xe9d817a2e042dbb5a1e846f677db950e6146e460e01eed27fe04631480ac782f.
//
// Solidity: event CoreStaked(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) WatchCoreStaked(opts *bind.WatchOpts, sink chan<- *ChpRegistryCoreStaked, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "CoreStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryCoreStaked)
				if err := _ChpRegistry.contract.UnpackLog(event, "CoreStaked", log); err != nil {
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

// ChpRegistryNodeStakeUpdatedIterator is returned from FilterNodeStakeUpdated and is used to iterate over the raw logs and unpacked data for NodeStakeUpdated events raised by the ChpRegistry contract.
type ChpRegistryNodeStakeUpdatedIterator struct {
	Event *ChpRegistryNodeStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ChpRegistryNodeStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryNodeStakeUpdated)
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
		it.Event = new(ChpRegistryNodeStakeUpdated)
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
func (it *ChpRegistryNodeStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryNodeStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryNodeStakeUpdated represents a NodeStakeUpdated event raised by the ChpRegistry contract.
type ChpRegistryNodeStakeUpdated struct {
	Sender       common.Address
	NodeIp       [32]byte
	PublicKey    [32]byte
	AmountStaked *big.Int
	Duration     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNodeStakeUpdated is a free log retrieval operation binding the contract event 0xdee80d82aa175e7361ba064c7d854fd7d94a6b845747744eadd6231d89db8778.
//
// Solidity: event NodeStakeUpdated(address indexed _sender, bytes32 _nodeIp, bytes32 _publicKey, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) FilterNodeStakeUpdated(opts *bind.FilterOpts, _sender []common.Address) (*ChpRegistryNodeStakeUpdatedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "NodeStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryNodeStakeUpdatedIterator{contract: _ChpRegistry.contract, event: "NodeStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeStakeUpdated is a free log subscription operation binding the contract event 0xdee80d82aa175e7361ba064c7d854fd7d94a6b845747744eadd6231d89db8778.
//
// Solidity: event NodeStakeUpdated(address indexed _sender, bytes32 _nodeIp, bytes32 _publicKey, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) WatchNodeStakeUpdated(opts *bind.WatchOpts, sink chan<- *ChpRegistryNodeStakeUpdated, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "NodeStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryNodeStakeUpdated)
				if err := _ChpRegistry.contract.UnpackLog(event, "NodeStakeUpdated", log); err != nil {
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

// ChpRegistryNodeStakedIterator is returned from FilterNodeStaked and is used to iterate over the raw logs and unpacked data for NodeStaked events raised by the ChpRegistry contract.
type ChpRegistryNodeStakedIterator struct {
	Event *ChpRegistryNodeStaked // Event containing the contract specifics and raw log

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
func (it *ChpRegistryNodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryNodeStaked)
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
		it.Event = new(ChpRegistryNodeStaked)
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
func (it *ChpRegistryNodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryNodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryNodeStaked represents a NodeStaked event raised by the ChpRegistry contract.
type ChpRegistryNodeStaked struct {
	Sender        common.Address
	NodeIp        [32]byte
	NodePublicKey [32]byte
	AmountStaked  *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeStaked is a free log retrieval operation binding the contract event 0x8d0c0090f54cd42b9919624c7171a1a68ca1824fdf6ed7c7f1d618e89bf96713.
//
// Solidity: event NodeStaked(address indexed _sender, bytes32 _nodeIp, bytes32 _nodePublicKey, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) FilterNodeStaked(opts *bind.FilterOpts, _sender []common.Address) (*ChpRegistryNodeStakedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "NodeStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryNodeStakedIterator{contract: _ChpRegistry.contract, event: "NodeStaked", logs: logs, sub: sub}, nil
}

// WatchNodeStaked is a free log subscription operation binding the contract event 0x8d0c0090f54cd42b9919624c7171a1a68ca1824fdf6ed7c7f1d618e89bf96713.
//
// Solidity: event NodeStaked(address indexed _sender, bytes32 _nodeIp, bytes32 _nodePublicKey, uint256 _amountStaked, uint256 _duration)
func (_ChpRegistry *ChpRegistryFilterer) WatchNodeStaked(opts *bind.WatchOpts, sink chan<- *ChpRegistryNodeStaked, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "NodeStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryNodeStaked)
				if err := _ChpRegistry.contract.UnpackLog(event, "NodeStaked", log); err != nil {
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

// ChpRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChpRegistry contract.
type ChpRegistryOwnershipTransferredIterator struct {
	Event *ChpRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChpRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryOwnershipTransferred)
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
		it.Event = new(ChpRegistryOwnershipTransferred)
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
func (it *ChpRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ChpRegistry contract.
type ChpRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChpRegistry *ChpRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChpRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryOwnershipTransferredIterator{contract: _ChpRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChpRegistry *ChpRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChpRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryOwnershipTransferred)
				if err := _ChpRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ChpRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ChpRegistry contract.
type ChpRegistryPausedIterator struct {
	Event *ChpRegistryPaused // Event containing the contract specifics and raw log

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
func (it *ChpRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryPaused)
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
		it.Event = new(ChpRegistryPaused)
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
func (it *ChpRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryPaused represents a Paused event raised by the ChpRegistry contract.
type ChpRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ChpRegistry *ChpRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*ChpRegistryPausedIterator, error) {

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ChpRegistryPausedIterator{contract: _ChpRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ChpRegistry *ChpRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ChpRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryPaused)
				if err := _ChpRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ChpRegistryPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ChpRegistry contract.
type ChpRegistryPauserAddedIterator struct {
	Event *ChpRegistryPauserAdded // Event containing the contract specifics and raw log

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
func (it *ChpRegistryPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryPauserAdded)
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
		it.Event = new(ChpRegistryPauserAdded)
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
func (it *ChpRegistryPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryPauserAdded represents a PauserAdded event raised by the ChpRegistry contract.
type ChpRegistryPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ChpRegistry *ChpRegistryFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*ChpRegistryPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryPauserAddedIterator{contract: _ChpRegistry.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ChpRegistry *ChpRegistryFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ChpRegistryPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryPauserAdded)
				if err := _ChpRegistry.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ChpRegistryPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ChpRegistry contract.
type ChpRegistryPauserRemovedIterator struct {
	Event *ChpRegistryPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ChpRegistryPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryPauserRemoved)
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
		it.Event = new(ChpRegistryPauserRemoved)
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
func (it *ChpRegistryPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryPauserRemoved represents a PauserRemoved event raised by the ChpRegistry contract.
type ChpRegistryPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ChpRegistry *ChpRegistryFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*ChpRegistryPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ChpRegistryPauserRemovedIterator{contract: _ChpRegistry.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ChpRegistry *ChpRegistryFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ChpRegistryPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryPauserRemoved)
				if err := _ChpRegistry.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ChpRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ChpRegistry contract.
type ChpRegistryUnpausedIterator struct {
	Event *ChpRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *ChpRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChpRegistryUnpaused)
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
		it.Event = new(ChpRegistryUnpaused)
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
func (it *ChpRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChpRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChpRegistryUnpaused represents a Unpaused event raised by the ChpRegistry contract.
type ChpRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ChpRegistry *ChpRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ChpRegistryUnpausedIterator, error) {

	logs, sub, err := _ChpRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ChpRegistryUnpausedIterator{contract: _ChpRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ChpRegistry *ChpRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ChpRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _ChpRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChpRegistryUnpaused)
				if err := _ChpRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
