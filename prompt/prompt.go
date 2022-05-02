// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package prompt

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/cryptoriums/packages/env"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/peterh/liner"
	"github.com/pkg/errors"
)

func Proxy(contracts []env.Contract) (*common.Address, error) {
	for i, contract := range contracts {
		fmt.Println(strconv.Itoa(i) + ":" + contract.Address.Hex() + " " + strings.Join(contract.Tags, ","))
	}
	for {
		resp, err := prompt.Stdin.PromptInput("Enter proxy address or leave empty: ")
		if err != nil {
			return nil, err
		}

		if resp == "" {
			return nil, nil
		}

		if !common.IsHexAddress(resp) {
			fmt.Println("Input is not a valid index or a token address")
			continue
		}

		selectedProxy := common.HexToAddress(resp)

		for _, contract := range contracts {
			if contract.Address.Hex() == selectedProxy.Hex() {
				return &selectedProxy, nil
			}
		}
		fmt.Println("proxy address not found")
	}
}

func Nonce(ctx context.Context, client ethereum.PendingStateReader, addr common.Address) (uint64, error) {
	for {
		nonce, err := client.PendingNonceAt(ctx, addr)
		if err != nil {
			return 0, errors.Wrap(err, "running PendingNonceAt")
		}
		for {
			_nonce, err := prompt.Stdin.PromptWithSuggestion("Nonce: ", strconv.Itoa(int(nonce)), 0)
			if err != nil {
				return 0, errors.Wrap(err, "PromptWithSuggestion for nonce")
			}
			nonceI, err := strconv.Atoi(_nonce)
			if err != nil {
				fmt.Println("parsing nonce input:", err)
				continue
			}
			return uint64(nonceI), nil
		}
	}
}

func ReadFile() ([]byte, string, error) {
	for {
		filePath, err := prompt.Stdin.PromptWithSuggestion("Enter file path: ", "/home/krasi/src/github.com/cryptoriums/packages/env.json", 0)
		if err != nil {
			if err == liner.ErrPromptAborted {
				return nil, "", err
			}
			//lint:ignore faillint for prompts is better with prints.
			fmt.Println("getting file path from terminal:", err)
			continue
		}
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			//lint:ignore faillint for prompts is better with prints.
			fmt.Println(err)
			continue
		}
		return content, filePath, nil
	}
}

func Token(netID int64) (*env.Token, error) {
	for i, token := range env.TOKENS {
		fmt.Println(strconv.Itoa(i) + ":" + token.Name + " " + token.Address[netID].Hex())
	}
	for {
		resp, err := prompt.Stdin.PromptInput("Select token index or enter custom token contract address: ")
		if err != nil {
			return nil, err
		}

		if common.IsHexAddress(resp) {
			return &env.Token{
				Name: "custom",
				Address: map[int64]common.Address{
					netID: common.HexToAddress(resp),
				},
			}, nil
		} else {
			index, err := strconv.Atoi(resp)
			if err != nil {
				fmt.Println("Input is not a valid index or a token address")
				continue
			}
			if index < 0 || index > len(env.TOKENS)-1 {
				fmt.Println("Input is outside the TOKENS range")
				continue
			}
			if _, ok := env.TOKENS[index].Address[netID]; !ok {
				return nil, errors.Errorf("token address unknown for current network:%v", netID)
			}
			return &env.TOKENS[index], nil
		}
	}
}

func Float(msg string, min, max float64) (float64, error) {
	for {
		_input, err := prompt.Stdin.PromptInput(msg)
		if err != nil {
			return 0, errors.Wrap(err, "PromptInput")
		}
		input, err := strconv.ParseFloat(_input, 64)
		if err != nil {
			fmt.Println("casting input to float err:", err)
			continue
		}
		if input < min || input > max {
			fmt.Println("input outside the allowed range - " + fmt.Sprintf("%.2f", min) + ":" + fmt.Sprintf("%.2f", max))
			continue
		}
		return input, nil
	}
}

func Int(msg string, min, max int) (int64, error) {
	for {
		_input, err := prompt.Stdin.PromptInput(msg)
		if err != nil {
			return 0, errors.Wrap(err, "PromptInput")
		}
		input, err := strconv.Atoi(_input)
		if err != nil {
			fmt.Println("casting input to float err:", err)
			continue
		}
		if input < min || input > max {
			fmt.Println("input outside the allowed range - " + strconv.Itoa(min) + ":" + strconv.Itoa(max))
			continue
		}
		return int64(input), nil
	}
}
