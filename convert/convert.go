// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

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

func ApiKeysToMap(keys []env.ApiKey) map[string]string {
	keysMap := make(map[string]string)
	for _, key := range keys {
		keysMap[key.Name] = key.Value
	}
	return keysMap
}

func ContractsToAddresses(contracts []env.Contract) []common.Address {
	var addrses []common.Address
	for _, contract := range contracts {
		addrses = append(addrses, contract.Address)
	}
	return addrses
}
