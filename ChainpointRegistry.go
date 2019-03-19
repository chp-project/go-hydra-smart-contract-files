// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"name\":\"nodeIp\",\"type\":\"bytes32\"},{\"name\":\"nodePublicKey\",\"type\":\"bytes32\"},{\"name\":\"isStaked\",\"type\":\"bool\"},{\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"name\":\"stakeLockedUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedCoresArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CORE_STAKING_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"cores\",\"outputs\":[{\"name\":\"coreIp\",\"type\":\"bytes32\"},{\"name\":\"corePublicKey\",\"type\":\"bytes32\"},{\"name\":\"isStaked\",\"type\":\"bool\"},{\"name\":\"isHealthy\",\"type\":\"bool\"},{\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"name\":\"stakeLockedUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NODE_STAKING_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coresArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CORE_STAKING_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NODE_STAKING_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"NodeStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_publicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"NodeStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_corePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_isHealthy\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"CoreStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_corePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_isHealthy\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"CoreStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"name\":\"_corePublicKey\",\"type\":\"bytes32\"}],\"name\":\"stakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"}],\"name\":\"updateStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"name\":\"_corePublicKey\",\"type\":\"bytes32\"}],\"name\":\"updateStakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unStakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"totalStakedFor\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"unlocks_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCoreCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isHealthyCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"whitelistCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_Registry *RegistryCaller) CORESTAKINGAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "CORE_STAKING_AMOUNT")
	return *ret0, err
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_Registry *RegistrySession) CORESTAKINGAMOUNT() (*big.Int, error) {
	return _Registry.Contract.CORESTAKINGAMOUNT(&_Registry.CallOpts)
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_Registry *RegistryCallerSession) CORESTAKINGAMOUNT() (*big.Int, error) {
	return _Registry.Contract.CORESTAKINGAMOUNT(&_Registry.CallOpts)
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_Registry *RegistryCaller) CORESTAKINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "CORE_STAKING_DURATION")
	return *ret0, err
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_Registry *RegistrySession) CORESTAKINGDURATION() (*big.Int, error) {
	return _Registry.Contract.CORESTAKINGDURATION(&_Registry.CallOpts)
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_Registry *RegistryCallerSession) CORESTAKINGDURATION() (*big.Int, error) {
	return _Registry.Contract.CORESTAKINGDURATION(&_Registry.CallOpts)
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_Registry *RegistryCaller) NODESTAKINGAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "NODE_STAKING_AMOUNT")
	return *ret0, err
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_Registry *RegistrySession) NODESTAKINGAMOUNT() (*big.Int, error) {
	return _Registry.Contract.NODESTAKINGAMOUNT(&_Registry.CallOpts)
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_Registry *RegistryCallerSession) NODESTAKINGAMOUNT() (*big.Int, error) {
	return _Registry.Contract.NODESTAKINGAMOUNT(&_Registry.CallOpts)
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_Registry *RegistryCaller) NODESTAKINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "NODE_STAKING_DURATION")
	return *ret0, err
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_Registry *RegistrySession) NODESTAKINGDURATION() (*big.Int, error) {
	return _Registry.Contract.NODESTAKINGDURATION(&_Registry.CallOpts)
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_Registry *RegistryCallerSession) NODESTAKINGDURATION() (*big.Int, error) {
	return _Registry.Contract.NODESTAKINGDURATION(&_Registry.CallOpts)
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Registry *RegistryCaller) Cores(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Registry.contract.Call(opts, out, "cores", arg0)
	return *ret, err
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Registry *RegistrySession) Cores(arg0 common.Address) (struct {
	CoreIp           [32]byte
	CorePublicKey    [32]byte
	IsStaked         bool
	IsHealthy        bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Registry.Contract.Cores(&_Registry.CallOpts, arg0)
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Registry *RegistryCallerSession) Cores(arg0 common.Address) (struct {
	CoreIp           [32]byte
	CorePublicKey    [32]byte
	IsStaked         bool
	IsHealthy        bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Registry.Contract.Cores(&_Registry.CallOpts, arg0)
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_Registry *RegistryCaller) CoresArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "coresArr", arg0)
	return *ret0, err
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_Registry *RegistrySession) CoresArr(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.CoresArr(&_Registry.CallOpts, arg0)
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_Registry *RegistryCallerSession) CoresArr(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.CoresArr(&_Registry.CallOpts, arg0)
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_Registry *RegistryCaller) GetCoreCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getCoreCount")
	return *ret0, err
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_Registry *RegistrySession) GetCoreCount() (*big.Int, error) {
	return _Registry.Contract.GetCoreCount(&_Registry.CallOpts)
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_Registry *RegistryCallerSession) GetCoreCount() (*big.Int, error) {
	return _Registry.Contract.GetCoreCount(&_Registry.CallOpts)
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_Registry *RegistryCaller) IsHealthyCore(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isHealthyCore", _address)
	return *ret0, err
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_Registry *RegistrySession) IsHealthyCore(_address common.Address) (bool, error) {
	return _Registry.Contract.IsHealthyCore(&_Registry.CallOpts, _address)
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_Registry *RegistryCallerSession) IsHealthyCore(_address common.Address) (bool, error) {
	return _Registry.Contract.IsHealthyCore(&_Registry.CallOpts, _address)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistrySession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCallerSession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Registry *RegistryCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Registry *RegistrySession) IsPauser(account common.Address) (bool, error) {
	return _Registry.Contract.IsPauser(&_Registry.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Registry *RegistryCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Registry.Contract.IsPauser(&_Registry.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Registry *RegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Registry *RegistrySession) Name() (string, error) {
	return _Registry.Contract.Name(&_Registry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Registry *RegistryCallerSession) Name() (string, error) {
	return _Registry.Contract.Name(&_Registry.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Registry *RegistryCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Registry.contract.Call(opts, out, "nodes", arg0)
	return *ret, err
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Registry *RegistrySession) Nodes(arg0 common.Address) (struct {
	NodeIp           [32]byte
	NodePublicKey    [32]byte
	IsStaked         bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Registry.Contract.Nodes(&_Registry.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Registry *RegistryCallerSession) Nodes(arg0 common.Address) (struct {
	NodeIp           [32]byte
	NodePublicKey    [32]byte
	IsStaked         bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Registry.Contract.Nodes(&_Registry.CallOpts, arg0)
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_Registry *RegistryCaller) NodesArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "nodesArr", arg0)
	return *ret0, err
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_Registry *RegistrySession) NodesArr(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.NodesArr(&_Registry.CallOpts, arg0)
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_Registry *RegistryCallerSession) NodesArr(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.NodesArr(&_Registry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Registry *RegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Registry *RegistrySession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Registry *RegistryCallerSession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_Registry *RegistryCaller) RegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "registryAddress")
	return *ret0, err
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_Registry *RegistrySession) RegistryAddress() (common.Address, error) {
	return _Registry.Contract.RegistryAddress(&_Registry.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_Registry *RegistryCallerSession) RegistryAddress() (common.Address, error) {
	return _Registry.Contract.RegistryAddress(&_Registry.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_Registry *RegistryCaller) Reward(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "reward")
	return *ret0, err
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_Registry *RegistrySession) Reward() (bool, error) {
	return _Registry.Contract.Reward(&_Registry.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_Registry *RegistryCallerSession) Reward() (bool, error) {
	return _Registry.Contract.Reward(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistrySession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistryCallerSession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_Registry *RegistryCaller) TotalStakedFor(opts *bind.CallOpts, addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	ret := new(struct {
		Amount    *big.Int
		UnlocksAt *big.Int
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "totalStakedFor", addr)
	return *ret, err
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_Registry *RegistrySession) TotalStakedFor(addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	return _Registry.Contract.TotalStakedFor(&_Registry.CallOpts, addr)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_Registry *RegistryCallerSession) TotalStakedFor(addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	return _Registry.Contract.TotalStakedFor(&_Registry.CallOpts, addr)
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_Registry *RegistryCaller) WhitelistedCoresArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "whitelistedCoresArr", arg0)
	return *ret0, err
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_Registry *RegistrySession) WhitelistedCoresArr(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.WhitelistedCoresArr(&_Registry.CallOpts, arg0)
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_Registry *RegistryCallerSession) WhitelistedCoresArr(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.WhitelistedCoresArr(&_Registry.CallOpts, arg0)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Registry *RegistryTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Registry *RegistrySession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddPauser(&_Registry.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Registry *RegistryTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddPauser(&_Registry.TransactOpts, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistrySession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Registry *RegistryTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Registry *RegistrySession) RenouncePauser() (*types.Transaction, error) {
	return _Registry.Contract.RenouncePauser(&_Registry.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Registry *RegistryTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Registry.Contract.RenouncePauser(&_Registry.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Registry *RegistryTransactor) Stake(opts *bind.TransactOpts, _nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "stake", _nodeIp, _nodePublicKey)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Registry *RegistrySession) Stake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.Stake(&_Registry.TransactOpts, _nodeIp, _nodePublicKey)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Registry *RegistryTransactorSession) Stake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.Stake(&_Registry.TransactOpts, _nodeIp, _nodePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Registry *RegistryTransactor) StakeCore(opts *bind.TransactOpts, _coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "stakeCore", _coreIp, _corePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Registry *RegistrySession) StakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.StakeCore(&_Registry.TransactOpts, _coreIp, _corePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Registry *RegistryTransactorSession) StakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.StakeCore(&_Registry.TransactOpts, _coreIp, _corePublicKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_Registry *RegistryTransactor) UnStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unStake")
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_Registry *RegistrySession) UnStake() (*types.Transaction, error) {
	return _Registry.Contract.UnStake(&_Registry.TransactOpts)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_Registry *RegistryTransactorSession) UnStake() (*types.Transaction, error) {
	return _Registry.Contract.UnStake(&_Registry.TransactOpts)
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_Registry *RegistryTransactor) UnStakeCore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unStakeCore")
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_Registry *RegistrySession) UnStakeCore() (*types.Transaction, error) {
	return _Registry.Contract.UnStakeCore(&_Registry.TransactOpts)
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_Registry *RegistryTransactorSession) UnStakeCore() (*types.Transaction, error) {
	return _Registry.Contract.UnStakeCore(&_Registry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistrySession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Registry *RegistryTransactor) UpdateStake(opts *bind.TransactOpts, _nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateStake", _nodeIp, _nodePublicKey)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Registry *RegistrySession) UpdateStake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateStake(&_Registry.TransactOpts, _nodeIp, _nodePublicKey)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Registry *RegistryTransactorSession) UpdateStake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateStake(&_Registry.TransactOpts, _nodeIp, _nodePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Registry *RegistryTransactor) UpdateStakeCore(opts *bind.TransactOpts, _coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateStakeCore", _coreIp, _corePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Registry *RegistrySession) UpdateStakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateStakeCore(&_Registry.TransactOpts, _coreIp, _corePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Registry *RegistryTransactorSession) UpdateStakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateStakeCore(&_Registry.TransactOpts, _coreIp, _corePublicKey)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_Registry *RegistryTransactor) WhitelistCore(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "whitelistCore", _address)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_Registry *RegistrySession) WhitelistCore(_address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.WhitelistCore(&_Registry.TransactOpts, _address)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_Registry *RegistryTransactorSession) WhitelistCore(_address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.WhitelistCore(&_Registry.TransactOpts, _address)
}

// RegistryCoreStakeUpdatedIterator is returned from FilterCoreStakeUpdated and is used to iterate over the raw logs and unpacked data for CoreStakeUpdated events raised by the Registry contract.
type RegistryCoreStakeUpdatedIterator struct {
	Event *RegistryCoreStakeUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryCoreStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryCoreStakeUpdated)
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
		it.Event = new(RegistryCoreStakeUpdated)
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
func (it *RegistryCoreStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryCoreStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryCoreStakeUpdated represents a CoreStakeUpdated event raised by the Registry contract.
type RegistryCoreStakeUpdated struct {
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
func (_Registry *RegistryFilterer) FilterCoreStakeUpdated(opts *bind.FilterOpts, _sender []common.Address) (*RegistryCoreStakeUpdatedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "CoreStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryCoreStakeUpdatedIterator{contract: _Registry.contract, event: "CoreStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchCoreStakeUpdated is a free log subscription operation binding the contract event 0x4459ca768ace1f4b62fc2885f066e940d9ebd0ceaf21bb1cf430576bd72c5bad.
//
// Solidity: event CoreStakeUpdated(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_Registry *RegistryFilterer) WatchCoreStakeUpdated(opts *bind.WatchOpts, sink chan<- *RegistryCoreStakeUpdated, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "CoreStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryCoreStakeUpdated)
				if err := _Registry.contract.UnpackLog(event, "CoreStakeUpdated", log); err != nil {
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

// RegistryCoreStakedIterator is returned from FilterCoreStaked and is used to iterate over the raw logs and unpacked data for CoreStaked events raised by the Registry contract.
type RegistryCoreStakedIterator struct {
	Event *RegistryCoreStaked // Event containing the contract specifics and raw log

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
func (it *RegistryCoreStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryCoreStaked)
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
		it.Event = new(RegistryCoreStaked)
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
func (it *RegistryCoreStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryCoreStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryCoreStaked represents a CoreStaked event raised by the Registry contract.
type RegistryCoreStaked struct {
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
func (_Registry *RegistryFilterer) FilterCoreStaked(opts *bind.FilterOpts, _sender []common.Address) (*RegistryCoreStakedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "CoreStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryCoreStakedIterator{contract: _Registry.contract, event: "CoreStaked", logs: logs, sub: sub}, nil
}

// WatchCoreStaked is a free log subscription operation binding the contract event 0xe9d817a2e042dbb5a1e846f677db950e6146e460e01eed27fe04631480ac782f.
//
// Solidity: event CoreStaked(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_Registry *RegistryFilterer) WatchCoreStaked(opts *bind.WatchOpts, sink chan<- *RegistryCoreStaked, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "CoreStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryCoreStaked)
				if err := _Registry.contract.UnpackLog(event, "CoreStaked", log); err != nil {
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

// RegistryNodeStakeUpdatedIterator is returned from FilterNodeStakeUpdated and is used to iterate over the raw logs and unpacked data for NodeStakeUpdated events raised by the Registry contract.
type RegistryNodeStakeUpdatedIterator struct {
	Event *RegistryNodeStakeUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryNodeStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNodeStakeUpdated)
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
		it.Event = new(RegistryNodeStakeUpdated)
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
func (it *RegistryNodeStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNodeStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNodeStakeUpdated represents a NodeStakeUpdated event raised by the Registry contract.
type RegistryNodeStakeUpdated struct {
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
func (_Registry *RegistryFilterer) FilterNodeStakeUpdated(opts *bind.FilterOpts, _sender []common.Address) (*RegistryNodeStakeUpdatedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NodeStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNodeStakeUpdatedIterator{contract: _Registry.contract, event: "NodeStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeStakeUpdated is a free log subscription operation binding the contract event 0xdee80d82aa175e7361ba064c7d854fd7d94a6b845747744eadd6231d89db8778.
//
// Solidity: event NodeStakeUpdated(address indexed _sender, bytes32 _nodeIp, bytes32 _publicKey, uint256 _amountStaked, uint256 _duration)
func (_Registry *RegistryFilterer) WatchNodeStakeUpdated(opts *bind.WatchOpts, sink chan<- *RegistryNodeStakeUpdated, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NodeStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNodeStakeUpdated)
				if err := _Registry.contract.UnpackLog(event, "NodeStakeUpdated", log); err != nil {
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

// RegistryNodeStakedIterator is returned from FilterNodeStaked and is used to iterate over the raw logs and unpacked data for NodeStaked events raised by the Registry contract.
type RegistryNodeStakedIterator struct {
	Event *RegistryNodeStaked // Event containing the contract specifics and raw log

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
func (it *RegistryNodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNodeStaked)
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
		it.Event = new(RegistryNodeStaked)
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
func (it *RegistryNodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNodeStaked represents a NodeStaked event raised by the Registry contract.
type RegistryNodeStaked struct {
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
func (_Registry *RegistryFilterer) FilterNodeStaked(opts *bind.FilterOpts, _sender []common.Address) (*RegistryNodeStakedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NodeStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNodeStakedIterator{contract: _Registry.contract, event: "NodeStaked", logs: logs, sub: sub}, nil
}

// WatchNodeStaked is a free log subscription operation binding the contract event 0x8d0c0090f54cd42b9919624c7171a1a68ca1824fdf6ed7c7f1d618e89bf96713.
//
// Solidity: event NodeStaked(address indexed _sender, bytes32 _nodeIp, bytes32 _nodePublicKey, uint256 _amountStaked, uint256 _duration)
func (_Registry *RegistryFilterer) WatchNodeStaked(opts *bind.WatchOpts, sink chan<- *RegistryNodeStaked, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NodeStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNodeStaked)
				if err := _Registry.contract.UnpackLog(event, "NodeStaked", log); err != nil {
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

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Registry contract.
type RegistryPausedIterator struct {
	Event *RegistryPaused // Event containing the contract specifics and raw log

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
func (it *RegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPaused)
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
		it.Event = new(RegistryPaused)
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
func (it *RegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPaused represents a Paused event raised by the Registry contract.
type RegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistryPausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistryPausedIterator{contract: _Registry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistryPaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPaused)
				if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// RegistryPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Registry contract.
type RegistryPauserAddedIterator struct {
	Event *RegistryPauserAdded // Event containing the contract specifics and raw log

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
func (it *RegistryPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPauserAdded)
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
		it.Event = new(RegistryPauserAdded)
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
func (it *RegistryPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPauserAdded represents a PauserAdded event raised by the Registry contract.
type RegistryPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_Registry *RegistryFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*RegistryPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &RegistryPauserAddedIterator{contract: _Registry.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_Registry *RegistryFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *RegistryPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPauserAdded)
				if err := _Registry.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// RegistryPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Registry contract.
type RegistryPauserRemovedIterator struct {
	Event *RegistryPauserRemoved // Event containing the contract specifics and raw log

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
func (it *RegistryPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPauserRemoved)
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
		it.Event = new(RegistryPauserRemoved)
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
func (it *RegistryPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPauserRemoved represents a PauserRemoved event raised by the Registry contract.
type RegistryPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_Registry *RegistryFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*RegistryPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &RegistryPauserRemovedIterator{contract: _Registry.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_Registry *RegistryFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *RegistryPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPauserRemoved)
				if err := _Registry.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// RegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Registry contract.
type RegistryUnpausedIterator struct {
	Event *RegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUnpaused)
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
		it.Event = new(RegistryUnpaused)
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
func (it *RegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUnpaused represents a Unpaused event raised by the Registry contract.
type RegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistryUnpausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistryUnpausedIterator{contract: _Registry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUnpaused)
				if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
