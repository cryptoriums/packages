// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package ethereum

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"regexp"
	"strings"
	"time"

	big_p "github.com/cryptoriums/packages/big"
	client "github.com/cryptoriums/packages/client"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
)

const (
	TxGasOverHead      = 21_000
	ComponentName      = "ethereum"
	BlockTime          = float64(15)
	BlocksPerSecond    = float64(1 / BlockTime)
	BlocksPerMinute    = float64(60 / BlockTime)
	ReorgEventWaitSafe = time.Minute
	ReorgEventWaitSlow = 3 * time.Minute
	ReorgEventWaitFast = 30 * time.Second

	MainnetName = "mainnet"
	RopstenName = "ropsten"
	GoerliName  = "goerli"
	RinkebyName = "rinkeby"
	HardhatName = "hardhat"

	MainnetID = 1
	RopstenID = 3
	GoerliID  = 4
	RinkebyID = 5
	HardhatID = 31337

	MaxBlockGasLimit = 30000000
	MaxGasPriceGwei  = 10_000 // To have some failsafe when creating TX and passing WEI instead of GWEI.
)

var NetworksByID = map[int64]string{
	MainnetID: MainnetName,
	RopstenID: RopstenName,
	RinkebyID: RinkebyName,
	GoerliID:  GoerliName,
	HardhatID: HardhatName,
}

var NetworksByName = map[string]int64{
	MainnetName: MainnetID,
	RopstenName: RopstenID,
	RinkebyName: RinkebyID,
	GoerliName:  GoerliID,
	HardhatName: HardhatID,
}

var ethAddressRE *regexp.Regexp = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// ValidateAddress checks if an ethereum URL is valid?
func ValidateAddress(address string) error {
	if match := ethAddressRE.MatchString(address); !match {
		return errors.New("invalid ethereum address")
	}
	return nil
}

type Account struct {
	Tags       []string
	PublicKey  common.Address
	PrivateKey *ecdsa.PrivateKey
}

func AccountFromPrvKey(pkey string) (Account, error) {
	pkey = strings.TrimPrefix(pkey, "0x")
	privateKey, err := crypto.HexToECDSA(strings.TrimSpace(pkey))
	if err != nil {
		return Account{}, errors.Wrap(err, "getting private key to ECDSA")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Account{}, errors.New("casting public key to ECDSA")
	}

	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return Account{PublicKey: publicAddress, PrivateKey: privateKey}, nil
}

func NewSignedTX(
	ctx context.Context,
	prvKey *ecdsa.PrivateKey,
	to common.Address,
	abis string,
	nonce uint64,
	netID int64,
	methodName string,
	args []interface{},
	gasLimit uint64,
	gasMaxFee float64,
	gasTip float64,
	value float64,
) (*types.Transaction, string, error) {

	// When just sending ether the data field is empty.
	data := []byte{}
	if abis != "" {
		abiP, err := abi.JSON(strings.NewReader(abis))
		if err != nil {
			return nil, "", errors.Wrap(err, "read contract ABI")
		}

		data, err = abiP.Pack(methodName, args...)
		if err != nil {
			return nil, "", errors.Wrap(err, "packing ABI")
		}
	}

	if gasMaxFee > MaxGasPriceGwei || gasTip > MaxGasPriceGwei {
		return nil, "", errors.Errorf("gas fee:%v or gas tip:%v higher than the maximum allowed:%v", gasMaxFee, gasTip, MaxGasPriceGwei)
	}

	if gasMaxFee == 0 {
		return nil, "", errors.New("for EIP1559 TXs the gasMaxFee should not be zero")
	}

	signer := types.LatestSignerForChainID(big.NewInt(netID))

	tx, err := types.SignNewTx(prvKey, signer, &types.DynamicFeeTx{
		ChainID:   big.NewInt(netID),
		Nonce:     nonce,
		GasFeeCap: big_p.FloatToBigIntMul(gasMaxFee, params.GWei),
		GasTipCap: big_p.FloatToBigIntMul(gasTip, params.GWei),
		Gas:       gasLimit,
		To:        &to,
		Data:      data,
		Value:     big_p.FloatToBigIntMul(value, params.Ether),
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

func NewTxOpts(
	ctx context.Context,
	client client.EthClient,
	nonce uint64,
	account Account,
	gasMaxFee float64,
	gasMaxTip float64,
	gasLimit uint64,
) (*bind.TransactOpts, error) {

	var gasMaxFeeWei *big.Int
	var gasMaxTipWei *big.Int
	var err error
	if gasMaxFee > 0 {
		gasMaxFeeWei = big_p.FloatToBigIntMul(gasMaxFee, params.GWei)
	}
	if gasMaxTip > 0 {
		gasMaxTipWei = big_p.FloatToBigIntMul(gasMaxTip, params.GWei)
	}

	if gasMaxFee > MaxGasPriceGwei || gasMaxTip > MaxGasPriceGwei {
		return nil, errors.Errorf("gas fee:%v or gas tip:%v higher than the maximum allowed:%v", gasMaxFee, gasMaxTip, MaxGasPriceGwei)
	}

	if nonce == 0 {
		nonce, err = client.PendingNonceAt(ctx, account.PublicKey)
		if err != nil {
			return nil, errors.Wrap(err, "getting pending nonce")
		}
	}

	if gasMaxTipWei == nil {
		gasMaxTipWei, err = client.SuggestGasTipCap(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "getting suggested gas tip")
		}
	}
	if gasMaxFeeWei == nil {
		if NetworksByID[client.NetworkID()] == HardhatName {
			return nil, errors.New("gasMaxFee is required for the hardhat network as it doesn't support the eth_maxPriorityFeePerGas method for getting the current max fee")
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
		gasMaxFeeWei = big.NewInt(0).Add(baseFee, gasMaxTipWei)
	}

	ethBalance, err := client.BalanceAt(ctx, account.PublicKey, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting balance")
	}

	cost := new(big.Int)
	cost.Mul(gasMaxFeeWei, big.NewInt(int64(gasLimit)))
	if ethBalance.Cmp(cost) < 0 {
		return nil, errors.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, big.NewInt(client.NetworkID()))
	if err != nil {
		return nil, errors.Wrap(err, "creating transactor")
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)

	opts.GasLimit = gasLimit
	opts.GasTipCap = gasMaxTipWei
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
	return "https://" + prefix + "etherscan.io"
}

type SendTransactionOpts struct {
	From     common.Address `json:"from"`               // The address the transaction is send from.
	To       common.Address `json:"to,omitempty"`       // (optional when creating new contract) The address the transaction is directed to.
	Gas      string         `json:"gas,omitempty"`      //  (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
	GasPrice string         `json:"gasPrice,omitempty"` // (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas.
	Value    string         `json:"value,omitempty"`    // (optional) Integer of the value sent with this transaction,
	Data     string         `json:"data"`               // The compiled code of a contract OR the hash of the invoked method signature and encoded parameters.
	Nonce    string         `json:"nonce,omitempty"`    // (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
}

func TestSignMessage(pubExp common.Address, priv *ecdsa.PrivateKey) error {
	msg := crypto.Keccak256([]byte("foo"))
	sig, err := crypto.Sign(msg, priv)
	if err != nil {
		return errors.Wrap(err, "crypto.Sign")
	}
	recoveredPub, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		return errors.Wrap(err, "crypto.Ecrecover")
	}
	_pubKeyAct, _ := crypto.UnmarshalPubkey(recoveredPub)
	pubKeyAct := crypto.PubkeyToAddress(*_pubKeyAct)
	if pubExp != pubKeyAct {
		return errors.Errorf("Address mismatch: want: %x have: %x", pubExp, pubKeyAct)
	}
	return nil
}
