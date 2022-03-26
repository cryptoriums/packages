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
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/peterh/liner"
	"github.com/pkg/errors"
)

type Envs map[string]Env

type Env struct {
	Nodes    []string
	ApiKeys  map[string]string
	Accounts []Account
}

type Account struct {
	Pub  common.Address
	Priv string
}

func createHash(key string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(key)); err != nil {
		return "", errors.Wrap(err, "hasher.Write")
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

const EncryptIndicator = "@"

func Decrypt(_input string, pass string) (string, error) {
	if _input[0] != EncryptIndicator[0] {
		return "", errors.Errorf("input doesn't start with a encryption indicator:%v so probably not encrypted", EncryptIndicator)
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

func Encrypt(_input string, pass string) (string, error) {
	if _input[0] == EncryptIndicator[0] {
		return "", errors.Errorf("input already starts with the encryption indicator:%v so probably already encrypted", EncryptIndicator)
	}
	input := []byte(_input)

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
	ciphered := gcm.Seal(nonce, nonce, input, nil)
	return EncryptIndicator + hex.EncodeToString(ciphered), nil
}

func DecryptEnvWithWebPassword(ctx context.Context, logger log.Logger, header string, env Env, host string, port uint) *Env {
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

func DecryptEnvWithPasswordLoop(env Env) (Env, error) {
	for {
		pass, err := prompt.Stdin.PromptPassword("Decrypt Password: ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return Env{}, err
			}
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("getting password from terminal:", err)
			continue
		}
		env, err = DecryptEnv(env, pass)
		if err != nil {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("Decrypt error try again:", err)
			continue
		}
		return env, nil
	}
}

func DecryptEnv(env Env, pass string) (Env, error) {
	for i, account := range env.Accounts {
		decryped, err := Decrypt(account.Priv, pass)
		if err != nil {
			return Env{}, errors.Wrapf(err, "decrypting account:%v", account.Pub)
		}
		account.Priv = decryped
		env.Accounts[i] = account
	}

	for name, value := range env.ApiKeys {
		if value[0] == EncryptIndicator[0] {
			decryped, err := Decrypt(value, pass)
			if err != nil {
				return Env{}, errors.Wrapf(err, "decrypting key:%v", name)
			}
			env.ApiKeys[name] = decryped
		}
	}
	return env, nil
}

func EncryptWithPasswordLoop(input string) (string, error) {
	for {
		pass, err := prompt.Stdin.PromptPassword("Encryption Password:")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return "", err
			}
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("getting password from terminal:", err)
			continue
		}
		if pass == "" {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("password can't be empty")
			continue
		}

		output, err := Encrypt(input, pass)
		if err != nil {
			return "", errors.Wrap(err, "encrypt input")
		}

		return output, nil
	}
}

func DecryptWithPasswordLoop(input string) (string, error) {
	for {
		pass, err := prompt.Stdin.PromptPassword("Decryption Password: ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return "", err
			}
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("getting password from terminal:", err)
			continue
		}
		if pass == "" {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println("password can't be empty")
			continue
		}

		output, err := Decrypt(input, pass)
		if err != nil {
			return "", errors.Wrap(err, "decrypt input")
		}

		return output, nil
	}
}

func EnvForNetwork(envs Envs, netName string) (Env, bool) {
	if env, ok := envs[strings.ToLower(netName)]; ok {
		return env, ok
	}

	env, ok := envs[strings.Title(netName)]
	return env, ok
}

func LoadFromEnvVarOrFile(envName, envFilePath string) (Envs, error) {
	envs, errEnvName := LoadFromEnv(envName)
	if errEnvName == nil {
		return envs, nil
	}

	envs, errEnvFile := LoadFromFile(envFilePath)
	if errEnvFile == nil {
		return envs, nil
	}
	return nil, errors.Errorf("env not found in envName:%v and envFilePath:%v errEnvName:%v, errEnvFile:%v", envName, envFilePath, errEnvName, errEnvFile)
}

func LoadFromEnv(envName string) (Envs, error) {
	content := os.Getenv(envName)
	if content == "" {
		return nil, errors.New(envName + " is empty")
	}

	envs := Envs{}
	err := json.Unmarshal([]byte(content), &envs)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Unmarshal the env var:%v", envName)
	}

	return envs, nil
}

func LoadFromFile(envFilePath string) (Envs, error) {
	content, err := os.ReadFile(envFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "reading the env file:%v", envFilePath)
	}

	envs := Envs{}
	err = json.Unmarshal(content, &envs)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Unmarshal the env file:%v", envFilePath)
	}

	return envs, nil
}
