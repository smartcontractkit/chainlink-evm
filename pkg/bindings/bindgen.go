package bindings

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink-evm/pkg/bindings/abigen"
)

//go:embed sourcecre.go.tpl
var tpl string

func GenerateBindings(
	combinedJSONPath string, // path to combined-json, or ""
	abiPath string, // path to a single ABI JSON, or ""
	pkgName string, // generated Go package name
	typeName string, // Go struct name for single-ABI mode (defaults to pkgName)
	outPath string, // where to write the .go file
) error {
	var (
		types   []string
		abis    []string
		bins    []string
		libs    = make(map[string]string)
		aliases = make(map[string]string)
	)

	switch {
	case combinedJSONPath != "":
		// Combined-JSON mode
		data, err := os.ReadFile(combinedJSONPath)
		if err != nil {
			return fmt.Errorf("read combined-json %q: %w", combinedJSONPath, err)
		}
		contracts, err := compiler.ParseCombinedJSON(data, "", "", "", "")
		if err != nil {
			return fmt.Errorf("parse combined-json %q: %w", combinedJSONPath, err)
		}
		for name, c := range contracts {
			parts := strings.Split(name, ":")
			tn := parts[len(parts)-1]
			abiDef, err := json.Marshal(c.Info.AbiDefinition)
			if err != nil {
				return fmt.Errorf("marshal ABI for %s: %w", name, err)
			}
			types = append(types, tn)
			abis = append(abis, string(abiDef))
			bins = append(bins, c.Code)

			// library placeholders
			prefix := crypto.Keccak256Hash([]byte(name)).String()[2:36]
			libs[prefix] = tn
		}

	case abiPath != "":
		// Single-ABI mode
		abiBytes, err := os.ReadFile(abiPath)
		if err != nil {
			return fmt.Errorf("read ABI %q: %w", abiPath, err)
		}
		// validate JSON
		if err := json.Unmarshal(abiBytes, new(interface{})); err != nil {
			return fmt.Errorf("invalid ABI JSON %q: %w", abiPath, err)
		}
		if typeName == "" {
			typeName = pkgName
		}
		types = []string{typeName}
		abis = []string{string(abiBytes)}
		bins = []string{""} // no deploy bytecode
		// no libraries in single-ABI mode

	default:
		return errors.New("must provide either combinedJSONPath or abiPath")
	}

	// Generate w/ forked abigen
	outSrc, err := abigen.BindV2(types, abis, bins, pkgName, libs, aliases, tpl)
	if err != nil {
		return fmt.Errorf("BindV2: %w", err)
	}

	// Write file
	if err := os.WriteFile(outPath, []byte(outSrc), 0o600); err != nil {
		return fmt.Errorf("write %q: %w", outPath, err)
	}
	return nil
}
