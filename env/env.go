// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package env

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt" //lint:ignore faillint for prompts it is better than logs.
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jinzhu/copier"
	"github.com/peterh/liner"
	"github.com/pkg/errors"
)

type Contract struct {
	Address common.Address
	Tags    []string `json:",omitempty"`
}

type Node struct {
	URL  string
	Tags []string `json:",omitempty"`
}

type ApiKey struct {
	Name  string
	Value string
	Tags  []string `json:",omitempty"`
}

type Env struct {
	Contracts []Contract `json:",omitempty"`
	Nodes     []Node
	ApiKeys   []ApiKey
	Accounts  []Account
}

type Account struct {
	Pub  common.Address
	Priv string
	Tags []string `json:",omitempty"`
}

func createHash(key string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(key)); err != nil {
		return "", errors.Wrap(err, "hasher.Write")
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

const EncryptIndicator = "@"

func IsEncryptedEnv(envr Env) bool {
	for _, acc := range envr.Accounts {
		if IsEncrypted(acc.Priv) {
			return true
		}
	}
	for _, key := range envr.ApiKeys {
		if IsEncrypted(key.Value) {
			return true
		}
	}

	return false
}
func IsEncrypted(_input string) bool {
	if len(_input) == 0 {
		return false
	}
	return _input[0] == EncryptIndicator[0]
}

func Decrypt(_input string, pass string) (string, error) {
	if !IsEncrypted(_input) {
		return "", errors.Errorf("input is not encrypted")
	}
	input, err := hex.DecodeString(_input[1:])
	if err != nil {
		return "", errors.Wrap(err, "decode hex string")
	}
	h, err := createHash(pass)

	if err != nil {
		return "", errors.Wrap(err, "create pass hash")
	}
	block, err := aes.NewCipher([]byte(h))
	if err != nil {
		return "", errors.Wrap(err, "NewCipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "NewGCM")
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := input[:nonceSize], input[nonceSize:]
	decyphered, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.Wrap(err, "gcm.Open")
	}
	return string(decyphered), nil
}

func Encrypt(input string, pass string) (string, error) {
	if IsEncrypted(input) {
		return "", errors.Errorf("input is already encrypted")
	}

	h, err := createHash(pass)
	if err != nil {
		return "", errors.Wrap(err, "create pass hash")
	}
	block, _ := aes.NewCipher([]byte(h))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "NewGCM")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Wrap(err, "read full")
	}
	ciphered := gcm.Seal(nonce, nonce, []byte(input), nil)
	return EncryptIndicator + hex.EncodeToString(ciphered), nil
}

func DecryptEnvWithWebPassword(ctx context.Context, logger log.Logger, header string, env Env, host string, port uint) *Env {
	if !IsEncryptedEnv(env) {
		return &env
	}
	level.Error(logger).Log("msg", "env is encrypted so use the web server to input the password to decrypt")

	envDecrypted := make(chan *Env)
	srv := &http.Server{Addr: host + ":" + strconv.Itoa(int(port))}
	defer func() {
		if err := srv.Shutdown(ctx); err != nil {
			level.Error(logger).Log("msg", "shutting down the password prompt web server", "err", err)
		}
	}()

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			level.Error(logger).Log("msg", "starting the password web server", "err", err)
			envDecrypted <- nil
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		postResult := ""
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf("parsing form data:%v", err), http.StatusInternalServerError)
				return
			}
			pass := r.PostForm.Get("decryptPass")

			env, err := DecryptEnv(env, pass)
			if err == nil {
				fmt.Fprintf(w, `Env decrypted, execution will continue!`)
				envDecrypted <- &env
				return
			}
			postResult = err.Error()
		}
		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang='en'>
		<head>
			<meta name='viewport' content='width=device-width, initial-scale=1, maximum-scale=1'>
			<title>Decrypt password</title>
			<style>
				body {
					font-family: arial;
				}
				label {
					min-width: 9em;
					float: left;
				}
			</style>
		</head>
		<body>
		`+header+`
		`+postResult+`
		<form id="data" method="post">
			<label for="decryptPass">Decrypt pass:</label><input type="password" name="decryptPass" id="decryptPass"/>
			<input type="submit" value="GO">
		</form>
		</body>
		</html>
		`)

	})

	return <-envDecrypted
}

func ReEncryptEnvWithPasswordLoop(envOrig Env) (Env, string, error) {
	for {
		pass, err := prompt.Stdin.PromptPassword("Decrypt password: ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return Env{}, "", err
			}
			fmt.Println("getting password from terminal:", err)
			continue
		}
		passNew, err := prompt.Stdin.PromptInput("New encrypt password: ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return Env{}, "", err
			}
			fmt.Println("getting password from terminal:", err)
			continue
		}
		env, err := ReEnryptEnv(envOrig, pass, passNew)
		if err != nil {
			fmt.Println("Decrypt error try again:", err)
			continue
		}
		return env, passNew, nil
	}
}

func DecryptEnvWithPasswordLoop(envOrig Env) (Env, string, error) {
	for {
		pass, err := prompt.Stdin.PromptPassword("Decrypt password: ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return Env{}, "", err
			}
			fmt.Println("getting password from terminal:", err)
			continue
		}
		env, err := DecryptEnv(envOrig, pass)
		if err != nil {
			fmt.Println("Decrypt error try again:", err)
			continue
		}
		return env, pass, nil
	}
}

func ReEnryptEnv(envOrig Env, pass, passNew string) (Env, error) {
	if passNew == "" {
		return Env{}, errors.New("new pass shouldn't be empty")
	}
	// Make a copy to not modify the original env as slices and maps are passed by reference.
	env := Env{}
	err := copier.CopyWithOption(&env, envOrig, copier.Option{DeepCopy: true})
	if err != nil {
		return Env{}, errors.Wrapf(err, "copier.CopyWithOption")
	}

	for id, account := range env.Accounts {
		if account.Priv == "" {
			continue
		}
		var decrypted string

		if !IsEncrypted(account.Priv) {
			_, err := crypto.HexToECDSA(strings.TrimPrefix(account.Priv, "0x"))
			if err != nil {
				return Env{}, errors.Wrapf(err, "error parsing unencrypted priv key:%v", account.Priv)
			}
			continue
		}

		var err error
		decrypted, err = Decrypt(account.Priv, pass)
		if err != nil {
			return Env{}, errors.Wrapf(err, "decrypting account:%v", account.Pub)
		}

		encrypted, err := Encrypt(decrypted, passNew)
		if err != nil {
			return Env{}, errors.Wrapf(err, "encrypting account:%v", account.Pub)
		}

		env.Accounts[id] = Account{
			Priv: encrypted,
			Pub:  account.Pub,
			Tags: account.Tags,
		}
	}

	for id, key := range env.ApiKeys {
		if IsEncrypted(key.Value) {
			decrypted, err := Decrypt(key.Value, pass)
			if err != nil {
				return Env{}, errors.Wrapf(err, "decrypting key:%v", key.Name)
			}
			encrypted, err := Encrypt(decrypted, passNew)
			if err != nil {
				return Env{}, errors.Wrapf(err, "encrypting key:%v", key.Name)
			}
			env.ApiKeys[id] = ApiKey{
				Name:  key.Name,
				Value: encrypted,
				Tags:  key.Tags,
			}
		}
	}
	return env, nil
}

func DecryptEnv(envOrig Env, pass string) (Env, error) {
	// Make a copy to not modify the original env as slices and maps are passed by reference.
	env := Env{}
	err := copier.CopyWithOption(&env, envOrig, copier.Option{DeepCopy: true})
	if err != nil {
		return Env{}, errors.Wrapf(err, "copier.CopyWithOption")
	}

	for id, account := range envOrig.Accounts {
		if account.Priv == "" {
			continue
		}
		if !IsEncrypted(account.Priv) {
			_, err := crypto.HexToECDSA(strings.TrimPrefix(account.Priv, "0x"))
			if err != nil {
				return Env{}, errors.Wrapf(err, "error parsing unencrypted priv key:%v", account.Priv)
			}
			continue
		}
		decrypted, err := Decrypt(account.Priv, pass)
		if err != nil {
			return Env{}, errors.Wrapf(err, "decrypting account:%v", account.Pub)
		}

		acc := Account{
			Priv: decrypted,
			Pub:  account.Pub,
			Tags: account.Tags,
		}
		env.Accounts[id] = acc
	}

	for id, apiKey := range env.ApiKeys {
		if IsEncrypted(apiKey.Value) {
			var err error
			apiKey.Value, err = Decrypt(apiKey.Value, pass)
			if err != nil {
				return Env{}, errors.Wrapf(err, "decrypting key:%v", apiKey.Name)
			}
			env.ApiKeys[id] = apiKey
		}
	}
	return env, nil
}

func EncryptAccounts(accsOrig []Account, pass string) ([]Account, error) {
	var accs []Account
	err := copier.CopyWithOption(&accs, accsOrig, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.Wrapf(err, "copier.CopyWithOption")
	}
	for i, account := range accs {
		if account.Priv == "" {
			continue
		}
		encrypted, err := Encrypt(account.Priv, pass)
		if err != nil {
			return nil, errors.Wrapf(err, "encrypting account:%v", account.Pub)
		}
		accs[i].Priv = encrypted
	}

	return accs, nil
}

func DecryptAccounts(accsOrig []Account, pass string) ([]Account, error) {
	// Deep copy to not modify the original slice.
	accs := []Account{}
	err := copier.CopyWithOption(&accs, accsOrig, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.Wrapf(err, "copier.CopyWithOption")
	}
	for i, account := range accs {
		if account.Priv == "" {
			continue
		}
		decrypted, err := Decrypt(account.Priv, pass)
		if err != nil {
			return nil, errors.Wrapf(err, "decrypting account:%v", account.Pub)
		}
		accs[i].Priv = decrypted
	}

	return accs, nil
}

func EncryptWithPasswordLoop(input string) (string, string, error) {
	for {
		pass, err := prompt.Stdin.PromptPassword("Encryption password: ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return "", "", err
			}
			fmt.Println("getting password from terminal:", err)
			continue
		}
		if pass == "" {
			fmt.Println("password can't be empty")
			continue
		}

		output, err := Encrypt(input, pass)
		if err != nil {
			return "", "", errors.Wrap(err, "encrypt input")
		}

		return output, pass, nil
	}
}

func DecryptWithPasswordLoop(input string) (string, error) {
	for {
		pass, err := prompt.Stdin.PromptPassword("Decryption password: ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return "", err
			}
			fmt.Println("getting password from terminal:", err)
			continue
		}
		if pass == "" {
			fmt.Println("password can't be empty")
			continue
		}

		output, err := Decrypt(input, pass)
		if err != nil {
			fmt.Println("decrypt input error:", err)
			continue
		}

		return output, nil
	}
}

func LoadFromEnvVarOrFile(envName, envFilePath string, tags ...string) (Env, error) {
	env, err1 := LoadFromEnv(envName, tags...)
	if err1 == nil {
		return env, nil
	}

	env, err2 := LoadFromFile(envFilePath, tags...)
	if err2 == nil {
		return env, nil
	}
	return Env{}, errors.Errorf("env not found in envName:%v and envFilePath:%v errEnvName:%v, errEnvFile:%v", envName, envFilePath, err1, err2)
}

func LoadFromEnv(envName string, tags ...string) (Env, error) {
	content := os.Getenv(envName)
	if content == "" {
		return Env{}, errors.New(envName + " is empty")
	}

	env := Env{}
	err := json.Unmarshal([]byte(content), &env)
	if err != nil {
		return Env{}, errors.Wrapf(err, "json.Unmarshal the env var:%v", envName)
	}

	return ApplyFilter(env, tags...), nil
}

func LoadFromFile(envFilePath string, tags ...string) (Env, error) {
	content, err := os.ReadFile(envFilePath)
	if err != nil {
		return Env{}, errors.Wrapf(err, "reading the env file:%v", envFilePath)
	}

	env := Env{}
	err = json.Unmarshal(content, &env)
	if err != nil {
		return Env{}, errors.Wrapf(err, "json.Unmarshal the env file:%v", envFilePath)
	}

	return ApplyFilter(env, tags...), nil
}

func ApplyFilter(env Env, tags ...string) Env {
	if len(tags) == 0 || tags[0] == "" {
		return env
	}

	var (
		nodes     []Node
		accounts  []Account
		apiKeys   []ApiKey
		contracts []Contract
	)

	for _, acc := range env.Accounts {
		if Contains(tags, acc.Tags) {
			accounts = append(accounts, acc)
		}
	}
	env.Accounts = accounts

	for _, node := range env.Nodes {
		if Contains(tags, node.Tags) {
			nodes = append(nodes, node)
		}
	}
	env.Nodes = nodes

	for _, key := range env.ApiKeys {
		if Contains(tags, key.Tags) {
			apiKeys = append(apiKeys, key)
		}
	}
	env.ApiKeys = apiKeys

	for _, contract := range env.Contracts {
		if Contains(tags, contract.Tags) {
			contracts = append(contracts, contract)
		}
	}
	env.Contracts = contracts

	return env
}

func Contains(tagsA []string, tagsB []string) bool {
	if len(tagsA) == 0 || len(tagsB) == 0 {
		return true
	}
	for _, tagA := range tagsA {
		for _, tagB := range tagsB {
			if strings.EqualFold(tagA, tagB) {
				return true
			}
		}

	}
	return false
}

type Token struct {
	Name    string
	Address map[int64]common.Address // Token addresses for each supported network.
}

var ETH_TOKEN = TOKENS[0]

var TOKENS = []Token{
	{
		"ETH",
		map[int64]common.Address{
			1:     {},
			3:     {},
			4:     {},
			5:     {},
			56:    {},
			31337: {},
		},
	},
	{
		"TRB",
		map[int64]common.Address{
			1: common.HexToAddress("0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"),
			4: common.HexToAddress("0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"),
		},
	},
}

func SelectAccount(accounts []Account, print bool, msg string) (Account, error) {
	if print {
		for i, acc := range accounts {
			noPrivate := ""
			if acc.Priv == "" {
				noPrivate = "*no private key"
			}
			fmt.Println(strconv.Itoa(i) + " " + acc.Pub.Hex() + " " + strings.Join(acc.Tags, ",") + " " + noPrivate)
		}
	}
	for {
		accAddr, err := prompt.Stdin.PromptInput(msg + " ")
		if err != nil {
			return Account{}, errors.Wrap(err, "select account prompt")
		}

		if !common.IsHexAddress(accAddr) {
			fmt.Println("input not an address:", accAddr)
		}

		for _, acc := range accounts {
			if acc.Pub.Hex() == common.HexToAddress(accAddr).Hex() {
				return acc, nil
			}
		}

		fmt.Println("account not found, try again:", accAddr)
	}
}

func SelectAccountAndDecrypt(accounts []Account, print bool, msg string) (Account, error) {
	for {
		account, err := SelectAccount(accounts, print, msg)
		if err != nil {
			return Account{}, err
		}
		print = false
		if account.Priv == "" {
			fmt.Println("selected account doesn't have a private key:", account.Pub)
			continue
		}

		if IsEncrypted(account.Priv) {
			account.Priv, err = DecryptWithPasswordLoop(account.Priv)
			if err != nil {
				return Account{}, errors.Wrap(err, "DecryptWithPasswordLoop")
			}
		}
		return account, nil
	}
}
