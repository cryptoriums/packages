// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"regexp"
	"strings"
	"time"

	math_t "github.com/cryptoriums/packages/math"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type EthClient interface {
	bind.ContractBackend
	ethereum.ChainStateReader
	ethereum.ChainReader
	ethereum.TransactionReader
	NetworkID() int64
	BlockNumber(ctx context.Context) (uint64, error)
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

const (
	PrivateKeysEnvName = "ETH_PRIVATE_KEYS"
	NodeURLEnvName     = "NODE_URLS"
	ComponentName      = "ethereum"
	BlockTime          = float64(15)
	BlocksPerSecond    = float64(1 / BlockTime)
	BlocksPerMinute    = float64(60 / BlockTime)
	ReorgEventWaitSafe = time.Minute
	ReorgEventWaitFast = 30 * time.Second
)

var ethAddressRE *regexp.Regexp = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// ValidateAddress checks if an ethereum URL is valid?
func ValidateAddress(address string) error {
	if match := ethAddressRE.MatchString(address); !match {
		return errors.New("invalid ethereum address")
	}
	return nil
}

// GetAddressForNetwork returns an ethereum address based on ethereum node network id.
func GetAddressForNetwork(addresses string, networkID int64) (string, error) {
	// Parse addresses to a map.
	networkToAddress := make(map[string]string)
	_addresses := strings.Split(addresses, ",")
	for _, address := range _addresses {
		parts := strings.Split(strings.TrimSpace(address), ":")
		if len(parts) != 2 {
			return "", errors.New("malformed ethereum <network:address> string")
		}
		if err := ValidateAddress(parts[1]); err != nil {
			return "", err
		}
		networkToAddress[parts[0]] = parts[1]
	}

	switch networkID {
	case 1:
		if val, ok := networkToAddress["Mainnet"]; ok {
			return val, nil
		}
		return "", errors.New("address for the Mainnet network not found in the address list")
	case 4:
		if val, ok := networkToAddress["Rinkeby"]; ok {
			return val, nil
		}
		return "", errors.New("address for the Rinkeby network not found in the address list")
	default:
		return "", errors.New("unhandled network id")
	}
}

func DecodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func Keccak256(input string) [32]byte {
	hash := crypto.Keccak256([]byte(input))
	var hashed [32]byte
	copy(hashed[:], hash)

	return hashed
}

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func (a *Account) GetAddress() common.Address {
	return a.Address
}

func (a *Account) GetPrivateKey() *ecdsa.PrivateKey {
	return a.PrivateKey
}

func GetAccountByPubAddress(logger log.Logger, pubAddr string, envVars map[string]string) (*Account, error) {
	accounts, err := GetAccounts(logger, envVars)
	if err != nil {
		return nil, errors.Wrap(err, "getting accounts")
	}

	for i, addr := range accounts {
		if addr.Address.Hex() == pubAddr {
			return accounts[i], nil
		}
	}
	return nil, errors.Errorf("account not found:%v", pubAddr)
}

// GetAccounts returns a slice of Account from private keys in
// PrivateKeysEnvName environment variable.
func GetAccounts(logger log.Logger, envVars map[string]string) ([]*Account, error) {
	_privateKeys, ok := envVars[PrivateKeysEnvName]
	if !ok {
		return nil, errors.New("private key env var is missing")
	}
	privateKeys := strings.Split(_privateKeys, ",")

	// Create an Account instance per private keys.
	accounts := make([]*Account, len(privateKeys))
	for i, pkey := range privateKeys {
		privateKey, err := crypto.HexToECDSA(strings.TrimSpace(pkey))
		if err != nil {
			return nil, errors.Wrap(err, "getting private key to ECDSA")
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		accounts[i] = &Account{Address: publicAddress, PrivateKey: privateKey}
		level.Info(logger).Log("msg", "registered account", "addr", publicAddress.Hex())
	}
	return accounts, nil
}

func NewClient(ctx context.Context, logger log.Logger, envVars map[string]string) (EthClient, error) {
	nodeURL, ok := envVars[NodeURLEnvName]
	if !ok {
		return nil, errors.Errorf("missing NodeURLEnvNam:%v", NodeURLEnvName)
	}
	nodes := strings.Split(nodeURL, ",")
	if len(nodes) == 0 {
		return nil, errors.New("the env file doesn't contain any node urls")
	}

	ethClient, rpcClient, netID, err := NewClients(ctx, logger, nodes[0])
	if err != nil {
		return nil, err
	}

	return &ClientCachedNetID{
		Client:    ethClient,
		netID:     netID,
		rpcClient: rpcClient,
	}, nil
}

func NewClients(ctx context.Context, logger log.Logger, nodeURL string) (*ethclient.Client, *rpc.Client, int64, error) {
	rpcClient, err := rpc.DialContext(ctx, nodeURL)
	if err != nil {
		return nil, nil, 0, err
	}
	ethClient := ethclient.NewClient(rpcClient)

	if !strings.Contains(strings.ToLower(nodeURL), "arbitrum") { // Arbitrum nodes doesn't support sync checking.
		// Issue #55, halt if client is still syncing with Ethereum network
		s, err := ethClient.SyncProgress(ctx)
		if err != nil {
			return nil, nil, 0, errors.Wrap(err, "determining if Ethereum client is syncing")
		}
		if s != nil {
			return nil, nil, 0, errors.New("ethereum node is still syncing with the network")
		}
	}

	netID, err := ethClient.NetworkID(ctx)
	if err != nil {
		return nil, nil, 0, errors.Wrap(err, "get nerwork ID")
	}

	level.Info(logger).Log("msg", "created ethereum client", "netID", netID.Int64())

	return ethClient, rpcClient, netID.Int64(), nil
}

type ClientCachedNetID struct {
	*ethclient.Client
	netID     int64
	rpcClient *rpc.Client
}

func (self *ClientCachedNetID) NetworkID() int64 {
	return self.netID
}

func (self *ClientCachedNetID) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return self.rpcClient.CallContext(ctx, result, method, args...)
}

func NewSignedTX(
	to common.Address,
	data []byte,
	nonce uint64,
	prvKey *ecdsa.PrivateKey,
	netID int64,
	gasLimit uint64,
	gasMaxFee float64,
	value float64,
) (*types.Transaction, string, error) {

	if gasMaxFee == 0 {
		return nil, "", errors.New("for EIP1559 TXs the gasMaxFee should not be zero")
	}

	signer := types.LatestSignerForChainID(big.NewInt(netID))

	tx, err := types.SignNewTx(prvKey, signer, &types.DynamicFeeTx{
		ChainID:   big.NewInt(netID),
		Nonce:     nonce,
		GasFeeCap: math_t.FloatToBigIntMul(gasMaxFee, params.GWei),
		GasTipCap: math_t.FloatToBigIntMul(gasMaxFee, params.GWei),
		Gas:       gasLimit,
		To:        &to,
		Data:      data,
		Value:     math_t.FloatToBigIntMul(value, params.Ether),
	})
	if err != nil {
		return nil, "", errors.Wrap(err, "sign transaction")
	}
	dataM, err := tx.MarshalBinary()
	if err != nil {
		return nil, "", errors.Wrap(err, "marshal tx data")
	}

	return tx, hexutil.Encode(dataM), nil
}

func PrepareTxOpts(
	ctx context.Context,
	client EthClient,
	account *Account,
	gasMaxFee float64,
	gasLimit uint64,
) (*bind.TransactOpts, error) {

	var gasMaxFeeWei *big.Int
	if gasMaxFee > 0 {
		gasMaxFeeWei = math_t.FloatToBigIntMul(gasMaxFee, params.GWei)
	}

	nonce, err := client.PendingNonceAt(ctx, account.GetAddress())
	if err != nil {
		return nil, errors.Wrap(err, "getting pending nonce")
	}

	if gasMaxFeeWei == nil {
		gasMaxTip, err := client.SuggestGasTipCap(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "getting suggested gas tip")
		}
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "getting chain header")
		}
		// Add 25% more for the base fee as a safe margin in case of a network load surge.
		// At high network load the base fee increases 12.5% per block
		// so 25% will allow including the TX in the next 2 blocks if the network load surges.
		safeMargin := big.NewInt(0).Div(header.BaseFee, big.NewInt(4))
		baseFee := big.NewInt(0).Add(header.BaseFee, safeMargin)
		gasMaxFeeWei = big.NewInt(0).Add(baseFee, gasMaxTip)
	}

	ethBalance, err := client.BalanceAt(ctx, account.GetAddress(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting balance")
	}

	cost := new(big.Int)
	cost.Mul(gasMaxFeeWei, big.NewInt(int64(gasLimit)))
	if ethBalance.Cmp(cost) < 0 {
		return nil, errors.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(account.GetPrivateKey(), big.NewInt(client.NetworkID()))
	if err != nil {
		return nil, errors.Wrap(err, "creating transactor")
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)

	opts.GasLimit = gasLimit
	opts.GasTipCap = gasMaxFeeWei
	opts.GasFeeCap = gasMaxFeeWei
	opts.Context = ctx
	return opts, nil
}

func GetEtherscanURL(netID int64) string {
	var prefix string
	switch netID {
	case 4:
		prefix = "rinkeby."
	case 5:
		prefix = "goerli."
	}
	return "https://" + prefix + "etherscan.io/"
}