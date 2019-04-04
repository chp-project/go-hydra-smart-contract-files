package ethcontracts

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"

	"github.com/chainpoint/chainpoint-core/go-abci-service/util"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/tendermint/tendermint/libs/log"
)

//EthClient : maintains connection information
type EthClient struct {
	EthURL               string
	TokenContractAddr    string
	RegistryContractAddr string
	Logger               log.Logger
	Client               ethclient.Client
}

//NewClient : creates a new ethclient with a connection to a particular url and particular token and registry contracts
func NewClient(ethURL string, tokenContract string, registryContract string, logger log.Logger) (*EthClient, error) {
	client, err := ethclient.Dial(ethURL)
	if util.LoggerError(logger, err) != nil {
		return nil, err
	}
	return &EthClient{
		EthURL:               ethURL,
		TokenContractAddr:    tokenContract,
		RegistryContractAddr: registryContract,
		Logger:               logger,
		Client:               *client,
	}, nil
}

//CheckEthBalance : Check Ethereum balance at a hex address
func (eth *EthClient) CheckEthBalance(addr string) (*big.Float, error) {
	address := common.HexToAddress(addr)
	balance, err := eth.Client.BalanceAt(context.Background(), address, nil)
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue, nil
}

//TokenBalanceOf : takes account hex string and gets a token balance
func (eth *EthClient) TokenBalanceOf(account string) (*big.Int, error) {
	tokenAddr := common.HexToAddress(eth.TokenContractAddr)
	tokenInstance, err := NewTNT(tokenAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	balance, err := tokenInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(account))
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	return balance, nil
}

//HighestBlock : Gets latest ethereum block
func (eth *EthClient) HighestBlock() (*big.Int, error) {
	header, err := eth.Client.HeaderByNumber(context.Background(), nil)
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	return header.Number, nil
}

//GetLastMintedAt : returns the ethereum block of the last minting event. Used for timing mint events
func (eth *EthClient) GetLastMintedAt() (*big.Int, error) {
	tokenAddr := common.HexToAddress(eth.TokenContractAddr)
	tokenInstance, err := NewTNT(tokenAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	balance, err := tokenInstance.LastMintedAtBlock(&bind.CallOpts{})
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	return balance, nil
}

//GetPastNodesStakedEvents : gets past nodes staking events
func (eth *EthClient) GetPastNodesStakedEvents() ([]ChpRegistryNodeStaked, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeStaked{}, err
	}
	opt := &bind.FilterOpts{}
	s := []common.Address{}
	staked, err := registryInstance.FilterNodeStaked(opt, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeStaked{}, err
	}
	nodeStructSlice := make([]ChpRegistryNodeStaked, 0)
	notEmpty := true
	for notEmpty {
		notEmpty = staked.Next()
		if notEmpty {
			nodeStructSlice = append(nodeStructSlice, *staked.Event)
		}
	}
	return nodeStructSlice, nil
}

//GetPastCoresStakedEvents : gets past cores staking events
func (eth *EthClient) GetPastCoresStakedEvents() ([]ChpRegistryCoreStaked, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreStaked{}, err
	}
	opt := &bind.FilterOpts{}
	s := []common.Address{}
	staked, err := registryInstance.FilterCoreStaked(opt, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreStaked{}, err
	}
	nodeStructSlice := make([]ChpRegistryCoreStaked, 0)
	notEmpty := true
	for notEmpty {
		notEmpty = staked.Next()
		if notEmpty {
			nodeStructSlice = append(nodeStructSlice, *staked.Event)
		}
	}
	return nodeStructSlice, nil
}

//GetPastNodesStakeUpdatedEvents : gets past nodes staking events
func (eth *EthClient) GetPastNodesStakeUpdatedEvents() ([]ChpRegistryNodeStakeUpdated, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeStakeUpdated{}, err
	}
	opt := &bind.FilterOpts{}
	s := []common.Address{}
	staked, err := registryInstance.FilterNodeStakeUpdated(opt, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeStakeUpdated{}, err
	}
	nodeStructSlice := make([]ChpRegistryNodeStakeUpdated, 0)
	notEmpty := true
	for notEmpty {
		notEmpty = staked.Next()
		if notEmpty {
			nodeStructSlice = append(nodeStructSlice, *staked.Event)
		}
	}
	return nodeStructSlice, nil
}

//GetPastCoresStakeUpdatedEvents : gets past core stake events
func (eth *EthClient) GetPastCoresStakeUpdatedEvents() ([]ChpRegistryCoreStakeUpdated, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreStakeUpdated{}, err
	}
	opt := &bind.FilterOpts{}
	s := []common.Address{}
	staked, err := registryInstance.FilterCoreStakeUpdated(opt, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreStakeUpdated{}, err
	}
	nodeStructSlice := make([]ChpRegistryCoreStakeUpdated, 0)
	notEmpty := true
	for notEmpty {
		notEmpty = staked.Next()
		if notEmpty {
			nodeStructSlice = append(nodeStructSlice, *staked.Event)
		}
	}
	return nodeStructSlice, nil
}

type nodeHandler func(ChpRegistryNodeStaked) error

//WatchNodeStakeEvents : watches for node stake events and passes them to the handler method
func (eth *EthClient) WatchNodeStakeEvents(handler nodeHandler, startBlock big.Int) error {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	opt := &bind.WatchOpts{}
	start := startBlock.Uint64()
	opt.Start = &start
	s := []common.Address{}
	ch := make(chan *ChpRegistryNodeStaked)
	sub, err := registryInstance.WatchNodeStaked(opt, ch, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	for {
		select {
		case event := <-ch:
			handler(*event)
			break
		case err := <-sub.Err():
			eth.Logger.Error("Subscription broken")
			return err
		}
	}
}

type coreHandler func(ChpRegistryCoreStaked) error

//WatchCoreStakeEvents : watches for core stake events and passes them to the handler method
func (eth *EthClient) WatchCoreStakeEvents(handler coreHandler, startBlock big.Int) error {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	opt := &bind.WatchOpts{}
	start := startBlock.Uint64()
	opt.Start = &start
	s := []common.Address{}
	ch := make(chan *ChpRegistryCoreStaked)
	sub, err := registryInstance.WatchCoreStaked(opt, ch, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	for {
		select {
		case event := <-ch:
			handler(*event)
			break
		case err := <-sub.Err():
			eth.Logger.Error("Subscription broken")
			return err
		}
	}
}

type nodeUpdatedHandler func(updated ChpRegistryNodeStakeUpdated) error

//WatchNodeStakeUpdatedEvents : watches for node stake updates and passes them to the handler method
func (eth *EthClient) WatchNodeStakeUpdatedEvents(handler nodeUpdatedHandler, startBlock big.Int) error {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	opt := &bind.WatchOpts{}
	start := startBlock.Uint64()
	opt.Start = &start
	s := []common.Address{}
	ch := make(chan *ChpRegistryNodeStakeUpdated)
	sub, err := registryInstance.WatchNodeStakeUpdated(opt, ch, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	for {
		select {
		case event := <-ch:
			handler(*event)
			break
		case err := <-sub.Err():
			eth.Logger.Error("Subscription broken")
			return err
		}
	}
}

type coreUpdatedHandler func(updated ChpRegistryCoreStakeUpdated) error

//WatchCoreStakeUpdatedEvents : watches for core stake updates and passes them to the handler method
func (eth *EthClient) WatchCoreStakeUpdatedEvents(handler coreUpdatedHandler, startBlock big.Int) error {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	opt := &bind.WatchOpts{}
	start := startBlock.Uint64()
	opt.Start = &start
	s := []common.Address{}
	ch := make(chan *ChpRegistryCoreStakeUpdated)
	sub, err := registryInstance.WatchCoreStakeUpdated(opt, ch, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	for {
		select {
		case event := <-ch:
			handler(*event)
			break
		case err := <-sub.Err():
			eth.Logger.Error("Subscription broken")
			return err
		}
	}
}

//PrivateKeyToAddress : converts hex private key to eth address
func PrivateKeyToAddress(privKey string) string {
	fromPrivateKey, _ := crypto.HexToECDSA(privKey)
	publicKey := fromPrivateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}
