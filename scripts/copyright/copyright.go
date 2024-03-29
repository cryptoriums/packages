// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	// license compatible for Go and Proto files.
	license = []byte(`// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

`)
)

func applyLicenseToProtoAndGo() error {
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Filter out stuff that does not need copyright.
		if info.IsDir() {
			switch path {
			case "vendor":
				return filepath.SkipDir
			case filepath.Join("pkg", "contracts"):
				return filepath.SkipDir
			}

			return nil
		}
		if strings.HasSuffix(path, ".pb.go") {
			return nil
		}
		if (filepath.Ext(path) != ".proto" && filepath.Ext(path) != ".go") ||
			// We copied this file and we want maintain its license (MIT).
			path == "pkg/testutil/testutil.go" ||
			strings.Contains(path, "testing/contracts/bindings") || // Generated bindings.
			strings.Contains(path, "contracts/bindings") { // Generated bindings.
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if !strings.HasPrefix(string(b), string(license)) {
			log.Println("file", path, "is missing Copyright header. Adding.")

			var bb bytes.Buffer
			_, _ = bb.Write(license)
			_, _ = bb.Write(b)
			if err = os.WriteFile(path, bb.Bytes(), 0666); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	if err := applyLicenseToProtoAndGo(); err != nil {
		log.Fatal(err)
	}
}
