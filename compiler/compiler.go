// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package compiler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func CompilerVersion(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", errors.Wrap(err, "opening the source file")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		switch ext := filepath.Ext(fileName); ext {
		case ".sol":
			if strings.Contains(line, "pragma solidity") {
				idxStart := strings.Index(line, "0")
				idxEnd := strings.Index(line, ";")

				version := line[idxStart:idxEnd]
				if len(version) == 3 {
					version += ".0"
				}

				return "v" + version, nil
			}
		case ".vy":
			if strings.Contains(line, "@version") {
				idxStart := strings.Index(line, "0")
				return "v" + line[idxStart:], nil
			}
		default:
			return "", errors.Errorf("unsupported file extension:%v", ext)
		}
	}
	return "", errors.New("source file doesn't contain compiler version")
}

func GetContractObjects(contractFiles map[string]string, compilerArgs []string) (types []string, abis []string, bins []string, sigs []map[string]string, libs map[string]string, err error) {
	libs = make(map[string]string)
	for contractPath, compilerVersion := range contractFiles {
		var contracts map[string]*Contract
		if filepath.Ext(contractPath) == ".sol" {
			compilerPath, err := DownloadSolc(compilerVersion)
			if err != nil {
				return nil, nil, nil, nil, nil, errors.Wrap(err, "download solc")
			}
			contracts, err = CompileSolidity(compilerPath, []string{contractPath}, compilerArgs)
			if err != nil {
				return nil, nil, nil, nil, nil, errors.Wrap(err, "build Solidity contract")
			}
		} else {
			compilerPath, err := DownloadVyper(compilerVersion)
			if err != nil {
				return nil, nil, nil, nil, nil, errors.Wrap(err, "download solc")
			}
			output, err := CompileVyper(compilerPath, []string{contractPath}, compilerArgs)
			if err != nil {
				return nil, nil, nil, nil, nil, errors.Wrap(err, "build Vyper contract")
			}
			contracts = make(map[string]*Contract)
			for n, contract := range output {
				name := n
				// Sanitize the combined json names to match the
				// format expected by solidity.
				if !strings.Contains(n, ":") {
					// Remove extra path components
					name = abi.ToCamelCase(strings.TrimSuffix(filepath.Base(name), ".vy"))
				}
				contracts[name] = contract
			}
		}

		for name, contract := range contracts {
			abi, err := json.Marshal(contract.Info.AbiDefinition)
			if err != nil {
				return nil, nil, nil, nil, nil, errors.Wrap(err, "flatten the compiler parse")
			}
			abis = append(abis, string(abi))
			bins = append(bins, contract.Code)
			sigs = append(sigs, contract.Hashes)
			nameParts := strings.Split(name, ":")
			types = append(types, nameParts[len(nameParts)-1])

			libPattern := crypto.Keccak256Hash([]byte(name)).String()[2:36]
			libs[libPattern] = nameParts[len(nameParts)-1]
		}

	}

	return types, abis, bins, sigs, libs, nil
}

func ExportABI(folder string, abis []string) error {
	filename := filepath.Base(folder)
	var a []byte
	for _, abi := range abis {
		if len(abi) > 2 {
			a = append(a, abi[1:len(abi)-1]...)
			a = append(a, []byte(",")...)

		}
	}
	a = a[:len(a)-1] // Remove the last comma from the array.
	a = append([]byte(`[`), a...)
	a = append(a, []byte("]")...)

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return errors.Wrapf(err, "create destination folder:%v", folder)
	}

	fpath := filepath.Join(folder, filename+".json")
	if err := os.WriteFile(fpath, a, os.ModePerm); err != nil {
		return errors.Wrapf(err, "write file:%v", fpath)
	}

	return nil
}

func ExportBin(folder string, types, bins []string) error {
	for i, t := range types {
		fpath := filepath.Join(folder, t+".bin")
		if err := os.WriteFile(fpath, []byte(bins[i]), os.ModePerm); err != nil {
			return errors.Wrapf(err, "write file:%v", fpath)
		}
	}
	return nil
}

func ExportPackage(pkgFolder string, types []string, abis []string, bins []string, sigs []map[string]string, libs map[string]string, aliases map[string]string) error {
	pkgName := filepath.Base(pkgFolder)
	code, err := bind.Bind(types, abis, bins, sigs, pkgName, libs, aliases)
	if err != nil {
		return errors.Wrapf(err, "generate the Go wrapper:%v", pkgName)
	}

	pkgPath := filepath.Join(pkgFolder, pkgName+".go")

	if _, err := os.Stat(pkgFolder); !os.IsNotExist(err) {
		os.RemoveAll(pkgFolder)
	}
	if err := os.MkdirAll(pkgFolder, os.ModePerm); err != nil {
		return errors.Wrapf(err, "create destination folder:%v", pkgFolder)
	}

	if err := os.WriteFile(pkgPath, []byte(code), os.ModePerm); err != nil {
		return errors.Wrapf(err, "write package file:%v", pkgPath)
	}
	return nil
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "gettings the file")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("downloading solc returned unexpected status code:%v", resp.StatusCode)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return errors.Wrap(err, "creating destination file")
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return errors.Wrap(err, "writing the file")
}

// DownloadSolc will download @solcVersion of the Solc compiler to tmp/solc directory.
func DownloadSolc(version string) (string, error) {
	solcDir := filepath.Join("tmp", "solc")
	if err := os.MkdirAll(solcDir, os.ModePerm); err != nil {
		return "", err
	}
	solcPath := filepath.Join(solcDir, version)
	if _, err := os.Stat(solcPath); os.IsNotExist(err) {
		log.Println("downloading solc version", version)

		srcFile := ""
		switch runtime.GOOS {
		case "darwin":
			srcFile = "solc-macos"
		case "linux":
			srcFile = "solc-static-linux"
		default:
			return "", errors.Errorf("unsuported OS:%v", runtime.GOOS)
		}

		err = downloadFile(solcPath, fmt.Sprintf("https://github.com/ethereum/solidity/releases/download/%s/%s", version, srcFile))
		if err != nil {
			return "", err
		}
		if err := os.Chmod(solcPath, os.ModePerm); err != nil {
			return "", err
		}
	}
	return solcPath, nil
}

// DownloadVyper will download the vyper copmiler.
func DownloadVyper(version string) (string, error) {
	compilerDir := filepath.Join("tmp", "vyper")
	if err := os.MkdirAll(compilerDir, os.ModePerm); err != nil {
		return "", err
	}
	vyperPath := filepath.Join(compilerDir, version)
	if _, err := os.Stat(vyperPath); os.IsNotExist(err) {
		log.Println("downloading vyper version", version)

		srcFile := ""
		switch version {
		case "v0.2.16":
			srcFile = "0.2.16+commit.59e1bdd"
		case "v0.2.6":
			srcFile = "0.2.6+commit.35467d5"
		case "v0.2.5":
			srcFile = "0.2.5+commit.a0c561c"
		case "v0.2.4":
			srcFile = "0.2.4+commit.7949850"
		default:
			return "", errors.Errorf("unrecognized version:%v", version)
		}
		switch runtime.GOOS {
		case "darwin":
			srcFile = "vyper." + srcFile + ".darwin"
		case "linux":
			srcFile = "vyper." + srcFile + ".linux"
		default:
			return "", errors.Errorf("unsuported OS:%v", runtime.GOOS)
		}

		err = downloadFile(vyperPath, fmt.Sprintf("https://github.com/vyperlang/vyper/releases/download/%s/%s", version, srcFile))
		if err != nil {
			return "", err
		}
		if err := os.Chmod(vyperPath, os.ModePerm); err != nil {
			return "", err
		}
	}
	return vyperPath, nil
}
