package ethcontracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"

	"github.com/chainpoint/chainpoint-core/go-abci-service/util"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/chainpoint/tendermint/libs/log"
)

//EthClient : maintains connection information
type EthClient struct {
	EthURL               string
	EthPrivateKey        string
	TokenContractAddr    string
	RegistryContractAddr string
	Logger               log.Logger
	Client               ethclient.Client
}

//NewClient : creates a new ethclient with a connection to a particular url and particular token and registry contracts
func NewClient(ethURL string, ethPrivateKey string, tokenContract string, registryContract string, logger log.Logger) (*EthClient, error) {
	client, err := ethclient.Dial(ethURL)
	if util.LoggerError(logger, err) != nil {
		fmt.Println("Unable to connect to ethereum url")
		return nil, err
	}
	return &EthClient{
		EthURL:               ethURL,
		EthPrivateKey:        ethPrivateKey,
		TokenContractAddr:    tokenContract,
		RegistryContractAddr: registryContract,
		Logger:               logger,
		Client:               *client,
	}, nil
}

//MintNodes : mints tokens for node reward candidates
func (eth *EthClient) MintNodes(rewardCandidates []common.Address, rcHash []byte, sigs [][]byte) error {
	tokenInstance, err := NewTNT(common.HexToAddress(eth.TokenContractAddr), &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	privateKey, _ := crypto.HexToECDSA(eth.EthPrivateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get nonce for fromAddress
	nonce, err := eth.Client.PendingNonceAt(context.Background(), fromAddress)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}

	// Get suggested gas price
	gasPrice, err := eth.Client.SuggestGasPrice(context.Background())
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}

	/// New keyed transactor
	transactOps := bind.NewKeyedTransactor(privateKey)
	transactOps.Nonce = big.NewInt(int64(nonce))
	transactOps.Value = big.NewInt(0)
	transactOps.GasLimit = uint64(300000)
	transactOps.GasPrice = gasPrice
	var rcHashBytes [32]byte
	copy(rcHashBytes[:], rcHash[:32])
	tx, err := tokenInstance.Mint(transactOps, rewardCandidates, rcHashBytes, sigs[0], sigs[1], sigs[2], sigs[3], sigs[4], sigs[5])
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	eth.Logger.Info(fmt.Sprintf("NodeMint: Issued ETH Tx: %#v", *tx))
	return nil
}

//MintCores : mints tokens for core reward candidates
func (eth *EthClient) MintCores(rewardCandidates []common.Address, rcHash []byte, sigs [126][]byte) error {
	tokenInstance, err := NewTNT(common.HexToAddress(eth.TokenContractAddr), &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	privateKey, _ := crypto.HexToECDSA(eth.EthPrivateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get nonce for fromAddress
	nonce, err := eth.Client.PendingNonceAt(context.Background(), fromAddress)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}

	// Get suggested gas price
	gasPrice, err := eth.Client.SuggestGasPrice(context.Background())
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}

	/// New keyed transactor
	transactOps := bind.NewKeyedTransactor(privateKey)
	transactOps.Nonce = big.NewInt(int64(nonce))
	transactOps.Value = big.NewInt(0)
	transactOps.GasLimit = uint64(800000)
	transactOps.GasPrice = gasPrice
	var rcHashBytes [32]byte
	copy(rcHashBytes[:], rcHash[:32])
	tx, err := tokenInstance.MintCores(transactOps, rewardCandidates, rcHashBytes, sigs)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	eth.Logger.Info(fmt.Sprintf("CoreMint: Issued ETH Tx: %#v", *tx))
	return nil
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

//GetNodeLastMintedAt : returns the ethereum block of the last minting event. Used for timing mint events
func (eth *EthClient) GetNodeLastMintedAt() (*big.Int, error) {
	tokenAddr := common.HexToAddress(eth.TokenContractAddr)
	tokenInstance, err := NewTNT(tokenAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	balance, err := tokenInstance.NodeLastMintedAtBlock(&bind.CallOpts{})
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	return balance, nil
}

//GetCoreLastMintedAt : returns the ethereum block of the last minting event. Used for timing mint events
func (eth *EthClient) GetCoreLastMintedAt() (*big.Int, error) {
	tokenAddr := common.HexToAddress(eth.TokenContractAddr)
	tokenInstance, err := NewTNT(tokenAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	balance, err := tokenInstance.CoreLastMintedAtBlock(&bind.CallOpts{})
	if util.LoggerError(eth.Logger, err) != nil {
		return nil, err
	}
	return balance, nil
}

//GetPastNodesStakedEvents : gets past nodes staking events
func (eth *EthClient) GetPastNodesStakedEvents(startBlock big.Int) ([]ChpRegistryNodeStaked, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeStaked{}, err
	}
	opt := &bind.FilterOpts{}
	start := startBlock.Uint64()
	opt.Start = start
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

//GetPastNodesStakeUpdatedEvents : gets past nodes staking events
func (eth *EthClient) GetPastNodesStakeUpdatedEvents(startBlock big.Int) ([]ChpRegistryNodeStakeUpdated, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeStakeUpdated{}, err
	}
	opt := &bind.FilterOpts{}
	start := startBlock.Uint64()
	opt.Start = start
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

//GetPastNodesUnstakeEvents : get past unstake events so they can be removed
func (eth *EthClient) GetPastNodesUnstakeEvents(startBlock big.Int) ([]ChpRegistryNodeUnStaked, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeUnStaked{}, err
	}
	opt := &bind.FilterOpts{}
	start := startBlock.Uint64()
	opt.Start = start
	s := []common.Address{}
	unStaked, err := registryInstance.FilterNodeUnStaked(opt, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryNodeUnStaked{}, err
	}
	nodeStructSlice := make([]ChpRegistryNodeUnStaked, 0)
	notEmpty := true
	for notEmpty {
		notEmpty = unStaked.Next()
		if notEmpty {
			nodeStructSlice = append(nodeStructSlice, *unStaked.Event)
		}
	}
	return nodeStructSlice, nil
}

//GetPastCoresStakedEvents : gets past cores staking events
func (eth *EthClient) GetPastCoresStakedEvents(startBlock big.Int) ([]ChpRegistryCoreStaked, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreStaked{}, err
	}
	opt := &bind.FilterOpts{}
	start := startBlock.Uint64()
	opt.Start = start
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

//GetPastCoresStakeUpdatedEvents : gets past core stake events
func (eth *EthClient) GetPastCoresStakeUpdatedEvents(startBlock big.Int) ([]ChpRegistryCoreStakeUpdated, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreStakeUpdated{}, err
	}
	opt := &bind.FilterOpts{}
	start := startBlock.Uint64()
	opt.Start = start
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

//GetPastCoresUnstakeEvents : get past unstake events so they can be removed
func (eth *EthClient) GetPastCoresUnstakeEvents(startBlock big.Int) ([]ChpRegistryCoreUnStaked, error) {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreUnStaked{}, err
	}
	opt := &bind.FilterOpts{}
	start := startBlock.Uint64()
	opt.Start = start
	s := []common.Address{}
	unStaked, err := registryInstance.FilterCoreUnStaked(opt, s)
	if util.LoggerError(eth.Logger, err) != nil {
		return []ChpRegistryCoreUnStaked{}, err
	}
	nodeStructSlice := make([]ChpRegistryCoreUnStaked, 0)
	notEmpty := true
	for notEmpty {
		notEmpty = unStaked.Next()
		if notEmpty {
			nodeStructSlice = append(nodeStructSlice, *unStaked.Event)
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
		eth.Logger.Info("smart contract connection failed")
		return err
	}
	opt := &bind.WatchOpts{}
	start := startBlock.Uint64()
	opt.Start = &start
	s := []common.Address{}
	ch := make(chan *ChpRegistryNodeStaked)
	sub, err := registryInstance.WatchNodeStaked(opt, ch, s)
	if util.LoggerError(eth.Logger, err) != nil {
		eth.Logger.Info("watch event connection failed")
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

type nodeUnstakeHandler func(updated ChpRegistryNodeUnStaked) error

//WatchNodeUnstakeEvents : watches for node unstaking events and passes them to the handler
func (eth *EthClient) WatchNodeUnstakeEvents(handler nodeUnstakeHandler, startBlock big.Int) error {
	registryAddr := common.HexToAddress(eth.RegistryContractAddr)
	registryInstance, err := NewChpRegistry(registryAddr, &eth.Client)
	if util.LoggerError(eth.Logger, err) != nil {
		return err
	}
	opt := &bind.WatchOpts{}
	start := startBlock.Uint64()
	opt.Start = &start
	s := []common.Address{}
	ch := make(chan *ChpRegistryNodeUnStaked)
	sub, err := registryInstance.WatchNodeUnStaked(opt, ch, s)
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
	fromPrivateKey, err := crypto.HexToECDSA(privKey)
	if util.LogError(err) != nil {
		return ""
	}
	publicKey := fromPrivateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}

//AddressesToHash : converts hex addresses to hash for minting
func AddressesToHash(addrs []common.Address) []byte {
	addrStrs := make([]string, 0)
	for _, addr := range addrs {
		addrStrs = append(addrStrs, addr.Hex())
	}
	hashAddr := solsha3.SoliditySHA3(
		// types
		[]string{"address[]"},
		// values
		[]interface{}{addrStrs},
	)
	return hashAddr
}

//SignMsg : produces ethereum signature of msg using corresponding privKey
func SignMsg(msg []byte, privKey string) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return []byte{}, err
	}
	signature, err := crypto.Sign(msg, privateKey)
	return signature, err
}
