package convert

import (
	"github.com/cryptoriums/packages/env"
	ethereum_p "github.com/cryptoriums/packages/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func EnvAccountsToEthAccounts(accs []env.Account) ([]*ethereum_p.Account, error) {
	var ethAccs []*ethereum_p.Account
	for _, acc := range accs {
		ethAcc, err := ethereum_p.AccountFromPrvKey(acc.Priv)
		if err != nil {
			return nil, errors.Wrap(err, "getting private key to ECDSA")
		}
		ethAccs = append(ethAccs, ethAcc)
	}
	return ethAccs, nil
}

func ContractsToAddresses(contracts []env.Contract) []common.Address {
	var addrses []common.Address
	for _, contract := range contracts {
		addrses = append(addrses, contract.Address)
	}
	return addrses
}

func EnvAccountsToEthAccountsMap(accs map[common.Address]env.Account) (map[common.Address]*ethereum_p.Account, error) {
	ethAccs := make(map[common.Address]*ethereum_p.Account)
	for _, acc := range accs {
		ethAcc, err := ethereum_p.AccountFromPrvKey(acc.Priv)
		if err != nil {
			return nil, errors.Wrap(err, "getting private key to ECDSA")
		}
		ethAccs[acc.Pub] = ethAcc
	}
	return ethAccs, nil
}
