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
const EthcontractsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"name\":\"nodeIp\",\"type\":\"bytes32\"},{\"name\":\"nodePublicKey\",\"type\":\"bytes32\"},{\"name\":\"isStaked\",\"type\":\"bool\"},{\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"name\":\"stakeLockedUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedCoresArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CORE_STAKING_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"cores\",\"outputs\":[{\"name\":\"coreIp\",\"type\":\"bytes32\"},{\"name\":\"corePublicKey\",\"type\":\"bytes32\"},{\"name\":\"isStaked\",\"type\":\"bool\"},{\"name\":\"isHealthy\",\"type\":\"bool\"},{\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"name\":\"stakeLockedUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NODE_STAKING_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coresArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CORE_STAKING_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NODE_STAKING_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"NodeStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_publicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"NodeStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_corePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_isHealthy\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"CoreStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_corePublicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_isHealthy\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"CoreStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"name\":\"_corePublicKey\",\"type\":\"bytes32\"}],\"name\":\"stakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeIp\",\"type\":\"bytes32\"},{\"name\":\"_nodePublicKey\",\"type\":\"bytes32\"}],\"name\":\"updateStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coreIp\",\"type\":\"bytes32\"},{\"name\":\"_corePublicKey\",\"type\":\"bytes32\"}],\"name\":\"updateStakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unStake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unStakeCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"totalStakedFor\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"unlocks_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCoreCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isHealthyCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"whitelistCore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) CORESTAKINGAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "CORE_STAKING_AMOUNT")
	return *ret0, err
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) CORESTAKINGAMOUNT() (*big.Int, error) {
	return _Ethcontracts.Contract.CORESTAKINGAMOUNT(&_Ethcontracts.CallOpts)
}

// CORESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xe5d0bf7d.
//
// Solidity: function CORE_STAKING_AMOUNT() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) CORESTAKINGAMOUNT() (*big.Int, error) {
	return _Ethcontracts.Contract.CORESTAKINGAMOUNT(&_Ethcontracts.CallOpts)
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) CORESTAKINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "CORE_STAKING_DURATION")
	return *ret0, err
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) CORESTAKINGDURATION() (*big.Int, error) {
	return _Ethcontracts.Contract.CORESTAKINGDURATION(&_Ethcontracts.CallOpts)
}

// CORESTAKINGDURATION is a free data retrieval call binding the contract method 0x7c8f3840.
//
// Solidity: function CORE_STAKING_DURATION() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) CORESTAKINGDURATION() (*big.Int, error) {
	return _Ethcontracts.Contract.CORESTAKINGDURATION(&_Ethcontracts.CallOpts)
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) NODESTAKINGAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "NODE_STAKING_AMOUNT")
	return *ret0, err
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) NODESTAKINGAMOUNT() (*big.Int, error) {
	return _Ethcontracts.Contract.NODESTAKINGAMOUNT(&_Ethcontracts.CallOpts)
}

// NODESTAKINGAMOUNT is a free data retrieval call binding the contract method 0xcfc635d4.
//
// Solidity: function NODE_STAKING_AMOUNT() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) NODESTAKINGAMOUNT() (*big.Int, error) {
	return _Ethcontracts.Contract.NODESTAKINGAMOUNT(&_Ethcontracts.CallOpts)
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) NODESTAKINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "NODE_STAKING_DURATION")
	return *ret0, err
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) NODESTAKINGDURATION() (*big.Int, error) {
	return _Ethcontracts.Contract.NODESTAKINGDURATION(&_Ethcontracts.CallOpts)
}

// NODESTAKINGDURATION is a free data retrieval call binding the contract method 0xf353e749.
//
// Solidity: function NODE_STAKING_DURATION() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) NODESTAKINGDURATION() (*big.Int, error) {
	return _Ethcontracts.Contract.NODESTAKINGDURATION(&_Ethcontracts.CallOpts)
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Ethcontracts *EthcontractsCaller) Cores(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Ethcontracts.contract.Call(opts, out, "cores", arg0)
	return *ret, err
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Ethcontracts *EthcontractsSession) Cores(arg0 common.Address) (struct {
	CoreIp           [32]byte
	CorePublicKey    [32]byte
	IsStaked         bool
	IsHealthy        bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Ethcontracts.Contract.Cores(&_Ethcontracts.CallOpts, arg0)
}

// Cores is a free data retrieval call binding the contract method 0x85ad32c4.
//
// Solidity: function cores(address ) constant returns(bytes32 coreIp, bytes32 corePublicKey, bool isStaked, bool isHealthy, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Ethcontracts *EthcontractsCallerSession) Cores(arg0 common.Address) (struct {
	CoreIp           [32]byte
	CorePublicKey    [32]byte
	IsStaked         bool
	IsHealthy        bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Ethcontracts.Contract.Cores(&_Ethcontracts.CallOpts, arg0)
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsCaller) CoresArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "coresArr", arg0)
	return *ret0, err
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsSession) CoresArr(arg0 *big.Int) (common.Address, error) {
	return _Ethcontracts.Contract.CoresArr(&_Ethcontracts.CallOpts, arg0)
}

// CoresArr is a free data retrieval call binding the contract method 0xd2c27afc.
//
// Solidity: function coresArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsCallerSession) CoresArr(arg0 *big.Int) (common.Address, error) {
	return _Ethcontracts.Contract.CoresArr(&_Ethcontracts.CallOpts, arg0)
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_Ethcontracts *EthcontractsCaller) GetCoreCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "getCoreCount")
	return *ret0, err
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_Ethcontracts *EthcontractsSession) GetCoreCount() (*big.Int, error) {
	return _Ethcontracts.Contract.GetCoreCount(&_Ethcontracts.CallOpts)
}

// GetCoreCount is a free data retrieval call binding the contract method 0xf1730874.
//
// Solidity: function getCoreCount() constant returns(uint256)
func (_Ethcontracts *EthcontractsCallerSession) GetCoreCount() (*big.Int, error) {
	return _Ethcontracts.Contract.GetCoreCount(&_Ethcontracts.CallOpts)
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_Ethcontracts *EthcontractsCaller) IsHealthyCore(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "isHealthyCore", _address)
	return *ret0, err
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_Ethcontracts *EthcontractsSession) IsHealthyCore(_address common.Address) (bool, error) {
	return _Ethcontracts.Contract.IsHealthyCore(&_Ethcontracts.CallOpts, _address)
}

// IsHealthyCore is a free data retrieval call binding the contract method 0x9273cce8.
//
// Solidity: function isHealthyCore(address _address) constant returns(bool)
func (_Ethcontracts *EthcontractsCallerSession) IsHealthyCore(_address common.Address) (bool, error) {
	return _Ethcontracts.Contract.IsHealthyCore(&_Ethcontracts.CallOpts, _address)
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

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Ethcontracts *EthcontractsCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
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
	err := _Ethcontracts.contract.Call(opts, out, "nodes", arg0)
	return *ret, err
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Ethcontracts *EthcontractsSession) Nodes(arg0 common.Address) (struct {
	NodeIp           [32]byte
	NodePublicKey    [32]byte
	IsStaked         bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Ethcontracts.Contract.Nodes(&_Ethcontracts.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address ) constant returns(bytes32 nodeIp, bytes32 nodePublicKey, bool isStaked, uint256 amountStaked, uint256 stakeLockedUntil)
func (_Ethcontracts *EthcontractsCallerSession) Nodes(arg0 common.Address) (struct {
	NodeIp           [32]byte
	NodePublicKey    [32]byte
	IsStaked         bool
	AmountStaked     *big.Int
	StakeLockedUntil *big.Int
}, error) {
	return _Ethcontracts.Contract.Nodes(&_Ethcontracts.CallOpts, arg0)
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsCaller) NodesArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "nodesArr", arg0)
	return *ret0, err
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsSession) NodesArr(arg0 *big.Int) (common.Address, error) {
	return _Ethcontracts.Contract.NodesArr(&_Ethcontracts.CallOpts, arg0)
}

// NodesArr is a free data retrieval call binding the contract method 0x2aa5f50e.
//
// Solidity: function nodesArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsCallerSession) NodesArr(arg0 *big.Int) (common.Address, error) {
	return _Ethcontracts.Contract.NodesArr(&_Ethcontracts.CallOpts, arg0)
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

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_Ethcontracts *EthcontractsCaller) RegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "registryAddress")
	return *ret0, err
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_Ethcontracts *EthcontractsSession) RegistryAddress() (common.Address, error) {
	return _Ethcontracts.Contract.RegistryAddress(&_Ethcontracts.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() constant returns(address)
func (_Ethcontracts *EthcontractsCallerSession) RegistryAddress() (common.Address, error) {
	return _Ethcontracts.Contract.RegistryAddress(&_Ethcontracts.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_Ethcontracts *EthcontractsCaller) Reward(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "reward")
	return *ret0, err
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_Ethcontracts *EthcontractsSession) Reward() (bool, error) {
	return _Ethcontracts.Contract.Reward(&_Ethcontracts.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() constant returns(bool)
func (_Ethcontracts *EthcontractsCallerSession) Reward() (bool, error) {
	return _Ethcontracts.Contract.Reward(&_Ethcontracts.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Ethcontracts *EthcontractsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Ethcontracts *EthcontractsSession) Token() (common.Address, error) {
	return _Ethcontracts.Contract.Token(&_Ethcontracts.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Ethcontracts *EthcontractsCallerSession) Token() (common.Address, error) {
	return _Ethcontracts.Contract.Token(&_Ethcontracts.CallOpts)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_Ethcontracts *EthcontractsCaller) TotalStakedFor(opts *bind.CallOpts, addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	ret := new(struct {
		Amount    *big.Int
		UnlocksAt *big.Int
	})
	out := ret
	err := _Ethcontracts.contract.Call(opts, out, "totalStakedFor", addr)
	return *ret, err
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_Ethcontracts *EthcontractsSession) TotalStakedFor(addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	return _Ethcontracts.Contract.TotalStakedFor(&_Ethcontracts.CallOpts, addr)
}

// TotalStakedFor is a free data retrieval call binding the contract method 0x4b341aed.
//
// Solidity: function totalStakedFor(address addr) constant returns(uint256 amount, uint256 unlocks_at)
func (_Ethcontracts *EthcontractsCallerSession) TotalStakedFor(addr common.Address) (struct {
	Amount    *big.Int
	UnlocksAt *big.Int
}, error) {
	return _Ethcontracts.Contract.TotalStakedFor(&_Ethcontracts.CallOpts, addr)
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsCaller) WhitelistedCoresArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ethcontracts.contract.Call(opts, out, "whitelistedCoresArr", arg0)
	return *ret0, err
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsSession) WhitelistedCoresArr(arg0 *big.Int) (common.Address, error) {
	return _Ethcontracts.Contract.WhitelistedCoresArr(&_Ethcontracts.CallOpts, arg0)
}

// WhitelistedCoresArr is a free data retrieval call binding the contract method 0x79a382fe.
//
// Solidity: function whitelistedCoresArr(uint256 ) constant returns(address)
func (_Ethcontracts *EthcontractsCallerSession) WhitelistedCoresArr(arg0 *big.Int) (common.Address, error) {
	return _Ethcontracts.Contract.WhitelistedCoresArr(&_Ethcontracts.CallOpts, arg0)
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

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) Stake(opts *bind.TransactOpts, _nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "stake", _nodeIp, _nodePublicKey)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsSession) Stake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Stake(&_Ethcontracts.TransactOpts, _nodeIp, _nodePublicKey)
}

// Stake is a paid mutator transaction binding the contract method 0x4f8b174e.
//
// Solidity: function stake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) Stake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.Stake(&_Ethcontracts.TransactOpts, _nodeIp, _nodePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) StakeCore(opts *bind.TransactOpts, _coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "stakeCore", _coreIp, _corePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsSession) StakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.StakeCore(&_Ethcontracts.TransactOpts, _coreIp, _corePublicKey)
}

// StakeCore is a paid mutator transaction binding the contract method 0xb0bc92e2.
//
// Solidity: function stakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) StakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.StakeCore(&_Ethcontracts.TransactOpts, _coreIp, _corePublicKey)
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

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_Ethcontracts *EthcontractsTransactor) UnStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "unStake")
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_Ethcontracts *EthcontractsSession) UnStake() (*types.Transaction, error) {
	return _Ethcontracts.Contract.UnStake(&_Ethcontracts.TransactOpts)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) UnStake() (*types.Transaction, error) {
	return _Ethcontracts.Contract.UnStake(&_Ethcontracts.TransactOpts)
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_Ethcontracts *EthcontractsTransactor) UnStakeCore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "unStakeCore")
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_Ethcontracts *EthcontractsSession) UnStakeCore() (*types.Transaction, error) {
	return _Ethcontracts.Contract.UnStakeCore(&_Ethcontracts.TransactOpts)
}

// UnStakeCore is a paid mutator transaction binding the contract method 0xdc964893.
//
// Solidity: function unStakeCore() returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) UnStakeCore() (*types.Transaction, error) {
	return _Ethcontracts.Contract.UnStakeCore(&_Ethcontracts.TransactOpts)
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

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) UpdateStake(opts *bind.TransactOpts, _nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "updateStake", _nodeIp, _nodePublicKey)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsSession) UpdateStake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.UpdateStake(&_Ethcontracts.TransactOpts, _nodeIp, _nodePublicKey)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x85b2fa42.
//
// Solidity: function updateStake(bytes32 _nodeIp, bytes32 _nodePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) UpdateStake(_nodeIp [32]byte, _nodePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.UpdateStake(&_Ethcontracts.TransactOpts, _nodeIp, _nodePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) UpdateStakeCore(opts *bind.TransactOpts, _coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "updateStakeCore", _coreIp, _corePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsSession) UpdateStakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.UpdateStakeCore(&_Ethcontracts.TransactOpts, _coreIp, _corePublicKey)
}

// UpdateStakeCore is a paid mutator transaction binding the contract method 0x9466291c.
//
// Solidity: function updateStakeCore(bytes32 _coreIp, bytes32 _corePublicKey) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) UpdateStakeCore(_coreIp [32]byte, _corePublicKey [32]byte) (*types.Transaction, error) {
	return _Ethcontracts.Contract.UpdateStakeCore(&_Ethcontracts.TransactOpts, _coreIp, _corePublicKey)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_Ethcontracts *EthcontractsTransactor) WhitelistCore(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Ethcontracts.contract.Transact(opts, "whitelistCore", _address)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_Ethcontracts *EthcontractsSession) WhitelistCore(_address common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.WhitelistCore(&_Ethcontracts.TransactOpts, _address)
}

// WhitelistCore is a paid mutator transaction binding the contract method 0xd8a8788d.
//
// Solidity: function whitelistCore(address _address) returns(bool)
func (_Ethcontracts *EthcontractsTransactorSession) WhitelistCore(_address common.Address) (*types.Transaction, error) {
	return _Ethcontracts.Contract.WhitelistCore(&_Ethcontracts.TransactOpts, _address)
}

// EthcontractsCoreStakeUpdatedIterator is returned from FilterCoreStakeUpdated and is used to iterate over the raw logs and unpacked data for CoreStakeUpdated events raised by the Ethcontracts contract.
type EthcontractsCoreStakeUpdatedIterator struct {
	Event *EthcontractsCoreStakeUpdated // Event containing the contract specifics and raw log

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
func (it *EthcontractsCoreStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsCoreStakeUpdated)
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
		it.Event = new(EthcontractsCoreStakeUpdated)
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
func (it *EthcontractsCoreStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsCoreStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsCoreStakeUpdated represents a CoreStakeUpdated event raised by the Ethcontracts contract.
type EthcontractsCoreStakeUpdated struct {
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
func (_Ethcontracts *EthcontractsFilterer) FilterCoreStakeUpdated(opts *bind.FilterOpts, _sender []common.Address) (*EthcontractsCoreStakeUpdatedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "CoreStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsCoreStakeUpdatedIterator{contract: _Ethcontracts.contract, event: "CoreStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchCoreStakeUpdated is a free log subscription operation binding the contract event 0x4459ca768ace1f4b62fc2885f066e940d9ebd0ceaf21bb1cf430576bd72c5bad.
//
// Solidity: event CoreStakeUpdated(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_Ethcontracts *EthcontractsFilterer) WatchCoreStakeUpdated(opts *bind.WatchOpts, sink chan<- *EthcontractsCoreStakeUpdated, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "CoreStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsCoreStakeUpdated)
				if err := _Ethcontracts.contract.UnpackLog(event, "CoreStakeUpdated", log); err != nil {
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

// EthcontractsCoreStakedIterator is returned from FilterCoreStaked and is used to iterate over the raw logs and unpacked data for CoreStaked events raised by the Ethcontracts contract.
type EthcontractsCoreStakedIterator struct {
	Event *EthcontractsCoreStaked // Event containing the contract specifics and raw log

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
func (it *EthcontractsCoreStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsCoreStaked)
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
		it.Event = new(EthcontractsCoreStaked)
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
func (it *EthcontractsCoreStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsCoreStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsCoreStaked represents a CoreStaked event raised by the Ethcontracts contract.
type EthcontractsCoreStaked struct {
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
func (_Ethcontracts *EthcontractsFilterer) FilterCoreStaked(opts *bind.FilterOpts, _sender []common.Address) (*EthcontractsCoreStakedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "CoreStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsCoreStakedIterator{contract: _Ethcontracts.contract, event: "CoreStaked", logs: logs, sub: sub}, nil
}

// WatchCoreStaked is a free log subscription operation binding the contract event 0xe9d817a2e042dbb5a1e846f677db950e6146e460e01eed27fe04631480ac782f.
//
// Solidity: event CoreStaked(address indexed _sender, bytes32 _coreIp, bytes32 _corePublicKey, bool _isHealthy, uint256 _amountStaked, uint256 _duration)
func (_Ethcontracts *EthcontractsFilterer) WatchCoreStaked(opts *bind.WatchOpts, sink chan<- *EthcontractsCoreStaked, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "CoreStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsCoreStaked)
				if err := _Ethcontracts.contract.UnpackLog(event, "CoreStaked", log); err != nil {
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

// EthcontractsNodeStakeUpdatedIterator is returned from FilterNodeStakeUpdated and is used to iterate over the raw logs and unpacked data for NodeStakeUpdated events raised by the Ethcontracts contract.
type EthcontractsNodeStakeUpdatedIterator struct {
	Event *EthcontractsNodeStakeUpdated // Event containing the contract specifics and raw log

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
func (it *EthcontractsNodeStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsNodeStakeUpdated)
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
		it.Event = new(EthcontractsNodeStakeUpdated)
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
func (it *EthcontractsNodeStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsNodeStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsNodeStakeUpdated represents a NodeStakeUpdated event raised by the Ethcontracts contract.
type EthcontractsNodeStakeUpdated struct {
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
func (_Ethcontracts *EthcontractsFilterer) FilterNodeStakeUpdated(opts *bind.FilterOpts, _sender []common.Address) (*EthcontractsNodeStakeUpdatedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "NodeStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsNodeStakeUpdatedIterator{contract: _Ethcontracts.contract, event: "NodeStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeStakeUpdated is a free log subscription operation binding the contract event 0xdee80d82aa175e7361ba064c7d854fd7d94a6b845747744eadd6231d89db8778.
//
// Solidity: event NodeStakeUpdated(address indexed _sender, bytes32 _nodeIp, bytes32 _publicKey, uint256 _amountStaked, uint256 _duration)
func (_Ethcontracts *EthcontractsFilterer) WatchNodeStakeUpdated(opts *bind.WatchOpts, sink chan<- *EthcontractsNodeStakeUpdated, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "NodeStakeUpdated", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsNodeStakeUpdated)
				if err := _Ethcontracts.contract.UnpackLog(event, "NodeStakeUpdated", log); err != nil {
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

// EthcontractsNodeStakedIterator is returned from FilterNodeStaked and is used to iterate over the raw logs and unpacked data for NodeStaked events raised by the Ethcontracts contract.
type EthcontractsNodeStakedIterator struct {
	Event *EthcontractsNodeStaked // Event containing the contract specifics and raw log

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
func (it *EthcontractsNodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthcontractsNodeStaked)
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
		it.Event = new(EthcontractsNodeStaked)
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
func (it *EthcontractsNodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthcontractsNodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthcontractsNodeStaked represents a NodeStaked event raised by the Ethcontracts contract.
type EthcontractsNodeStaked struct {
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
func (_Ethcontracts *EthcontractsFilterer) FilterNodeStaked(opts *bind.FilterOpts, _sender []common.Address) (*EthcontractsNodeStakedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.FilterLogs(opts, "NodeStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return &EthcontractsNodeStakedIterator{contract: _Ethcontracts.contract, event: "NodeStaked", logs: logs, sub: sub}, nil
}

// WatchNodeStaked is a free log subscription operation binding the contract event 0x8d0c0090f54cd42b9919624c7171a1a68ca1824fdf6ed7c7f1d618e89bf96713.
//
// Solidity: event NodeStaked(address indexed _sender, bytes32 _nodeIp, bytes32 _nodePublicKey, uint256 _amountStaked, uint256 _duration)
func (_Ethcontracts *EthcontractsFilterer) WatchNodeStaked(opts *bind.WatchOpts, sink chan<- *EthcontractsNodeStaked, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Ethcontracts.contract.WatchLogs(opts, "NodeStaked", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthcontractsNodeStaked)
				if err := _Ethcontracts.contract.UnpackLog(event, "NodeStaked", log); err != nil {
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
