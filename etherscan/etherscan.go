// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package etherscan

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/nanmu42/etherscan-api"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

func DownloadContracts(network etherscan.Network, address string, dstPath string) (map[string]string, error) {
	client := etherscan.New(network, "")
	rep, err := client.ContractSource(address)
	if err != nil {
		return nil, errors.Wrap(err, "get contract source")
	}

	if _, err := os.Stat(dstPath); !os.IsNotExist(err) {
		os.RemoveAll(dstPath)
	}
	if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
		return nil, errors.Wrapf(err, "create download folder:%v", dstPath)
	}

	var contractFiles = make(map[string]string)

	if srcCodes, ok := isMultiContract(rep[0].SourceCode); ok {
		for filePath := range srcCodes {
			content := srcCodes[filePath].Content
			filePath := filepath.Join(dstPath, filepath.Base(filePath))
			if err := write(filePath, content); err != nil {
				return nil, err
			}
			contractFiles[filePath] = strings.Split(rep[0].CompilerVersion, "+")[0]
		}
	} else {
		if strings.Contains(rep[0].CompilerVersion, "vyper") {
			contractName, err := stringInBetween(rep[0].SourceCode, "@title", "@author")
			if err != nil {
				return nil, errors.Wrap(err, "getting contract name")
			}
			filePath := filepath.Join(dstPath, contractName+".vy")
			if err := write(filePath, rep[0].SourceCode); err != nil {
				return nil, err
			}
			contractFiles[filePath] = "v" + strings.Split(rep[0].CompilerVersion, ":")[1]
		} else {
			contractName, err := stringInBetween(rep[0].SourceCode, "contract", "{")
			if err != nil {
				return nil, errors.Wrap(err, "getting contract name")
			}
			filePath := filepath.Join(dstPath, contractName+".sol")
			if err := write(filePath, rep[0].SourceCode); err != nil {
				return nil, err
			}
			contractFiles[filePath] = strings.Split(rep[0].CompilerVersion, "+")[0]
		}
	}

	return contractFiles, nil
}

func write(filePath, content string) (errFinal error) {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			errFinal = multierr.Append(errFinal, err)
		}
	}()
	w := bufio.NewWriter(f)
	defer func() {
		if err := w.Flush(); err != nil {
			errFinal = multierr.Append(errFinal, err)
		}
	}()

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		// Flatten the import subtree.
		// Rewrite the imports to remove all parent folder.
		if strings.HasPrefix(line, "import") {
			last := strings.LastIndex(line, "/")
			line = "import \"./" + line[last+1:len(line)-2] + "\";"
		}
		line += "\n"
		if _, err := w.Write([]byte(line)); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

type MultiContract struct {
	Language string
	Sources  map[string]Src
}

type Src struct {
	Content string
}

func stringInBetween(src, keywordStart, keywordEnd string) (string, error) {
	s := strings.Index(src, keywordStart)
	if s == -1 {
		return "", errors.New(keywordStart + " keyword not found")
	}
	s += len(keywordStart)
	e := strings.Index(src, keywordEnd)
	if e == -1 {
		return "", errors.New(keywordEnd + " keyword not found")
	}
	return strings.TrimSpace(src[s : e-1]), nil
}

func isMultiContract(s string) (map[string]Src, bool) {
	out := &MultiContract{}

	// Etherscan has inconsistent api responses so need to deal with these here.
	if err := json.Unmarshal([]byte(s), &out.Sources); err == nil {
		return out.Sources, true
	}

	s = strings.ReplaceAll(s, "{{", "{") // Deal with another wierdness of etherscan.
	s = strings.ReplaceAll(s, "}}", "}")

	if err := json.Unmarshal([]byte(s), out); err == nil {
		return out.Sources, true
	}
	return nil, false
}
