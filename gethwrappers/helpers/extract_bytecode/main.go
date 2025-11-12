package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	// Directories are relative to the EVM root (i.e. chains/evm)
	gobindingsSubdir    = "gobindings/generated"
	ccvGobindingsSubdir = "../../ccv/chains/evm/gobindings/generated"
	bytecodeSubdir      = "bytecode"
	abiSubdir           = "abi"
)

// Metadata holds the extracted bytecode and ABI from a contract binding
type Metadata struct {
	Bytecode string
	ABI      string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Determine the correct paths based on the current working directory
	gobindingsDir, ccvGobindingsDir, bytecodeDir, abiDir, err := findProjectDirs()
	if err != nil {
		return fmt.Errorf("failed to find project directories: %w", err)
	}

	fmt.Printf("Using gobindings dir: %s\n", gobindingsDir)
	fmt.Printf("Using ccv gobindings dir: %s\n", ccvGobindingsDir)
	fmt.Printf("Using bytecode dir: %s\n", bytecodeDir)
	fmt.Printf("Using ABI dir: %s\n", abiDir)

	// Get all versioned directories (exclude "latest" because it's unaudited)
	entries, err := os.ReadDir(gobindingsDir)
	if err != nil {
		return fmt.Errorf("failed to read gobindings directory: %w", err)
	}
	preCCVLen := len(entries)
	ccvEntries, err := os.ReadDir(ccvGobindingsDir)
	if err != nil {
		return fmt.Errorf("failed to read ccv gobindings directory: %w", err)
	}
	entries = append(entries, ccvEntries...)

	for i, entry := range entries {
		if !entry.IsDir() || entry.Name() == "latest" {
			continue
		}
		bindingsDir := gobindingsDir
		if i >= preCCVLen {
			bindingsDir = ccvGobindingsDir
		}

		versionDir := filepath.Join(bindingsDir, entry.Name())
		fmt.Printf("Processing version: %s\n", entry.Name())

		if err := processVersionDir(versionDir, bindingsDir, bytecodeDir, abiDir); err != nil {
			return fmt.Errorf("failed to process version %s: %w", entry.Name(), err)
		}
	}

	return nil
}

func findProjectDirs() (gobindingsDir, ccvGobindingsDir, bytecodeDir, abiDir string, err error) {
	// Find the chains/evm root directory
	evmRoot, err := findEvmRoot()
	if err != nil {
		return "", "", "", "", err
	}

	// Build all paths from the root using constants
	gobindingsDir = filepath.Join(evmRoot, gobindingsSubdir)
	ccvGobindingsDir = filepath.Join(evmRoot, ccvGobindingsSubdir)
	bytecodeDir = filepath.Join(evmRoot, bytecodeSubdir)
	abiDir = filepath.Join(evmRoot, abiSubdir)

	return gobindingsDir, ccvGobindingsDir, bytecodeDir, abiDir, nil
}

// findEvmRoot finds the chains/evm directory by walking up from the current directory
func findEvmRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Walk up the directory tree looking for chains/evm root
	dir := cwd
	for {
		// Check if gobindings/generated exists in this directory (indicates we found chains/evm)
		gobindingsPath := filepath.Join(dir, gobindingsSubdir)
		if info, err := os.Stat(gobindingsPath); err == nil && info.IsDir() {
			return dir, nil
		}

		// Check if we're in a subdirectory of chains/evm (like gobindings or gobindings/cmd)
		// by checking if the parent has gobindings/generated
		parent := filepath.Dir(dir)
		if parent != dir {
			parentGobindingsPath := filepath.Join(parent, gobindingsSubdir)
			if info, err := os.Stat(parentGobindingsPath); err == nil && info.IsDir() {
				return parent, nil
			}
		}

		// Check if chains/evm exists as a subdirectory (we might be in project root)
		chainsEvmPath := filepath.Join(dir, "chains", "evm", gobindingsSubdir)
		if info, err := os.Stat(chainsEvmPath); err == nil && info.IsDir() {
			return filepath.Join(dir, "chains", "evm"), nil
		}

		// Move up one directory
		if parent == dir {
			// We've reached the filesystem root
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("could not find chains/evm root directory. Please run from within the chainlink-ccip project")
}

func processVersionDir(versionDir, gobindingsDir, bytecodeDir, abiDir string) error {
	return filepath.Walk(versionDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process .go files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		return processGoFile(path, gobindingsDir, bytecodeDir, abiDir)
	})
}

func processGoFile(path, gobindingsDir, bytecodeDir, abiDir string) error {
	// Extract metadata from the file
	metadata, err := extractMetadata(path)
	if err != nil {
		return err
	}

	if metadata.Bytecode == "" && metadata.ABI == "" {
		// No metadata found in this file, skip
		return nil
	}

	// Determine the base output path - flatten to version/filename structure
	relPath, err := filepath.Rel(gobindingsDir, path)
	if err != nil {
		return fmt.Errorf("failed to get relative path for %s: %w", path, err)
	}

	// Extract version (first directory) and base filename
	parts := strings.Split(filepath.ToSlash(relPath), "/")
	if len(parts) < 2 {
		return fmt.Errorf("unexpected path structure: %s", relPath)
	}
	version := parts[0]
	baseFilename := strings.TrimSuffix(filepath.Base(path), ".go")

	// Extract bytecode if present
	if metadata.Bytecode != "" {
		if err := writeBytecode(bytecodeDir, version, baseFilename, metadata.Bytecode); err != nil {
			return err
		}
	}

	// Extract ABI if present
	if metadata.ABI != "" {
		if err := writeABI(abiDir, version, baseFilename, metadata.ABI); err != nil {
			return err
		}
	}

	return nil
}

func writeBytecode(bytecodeDir, version, baseFilename, bytecode string) error {
	filename := baseFilename + ".bin"
	filePath := filepath.Join(bytecodeDir, version, filename)
	relPath := filepath.Join(version, filename)

	if err := os.MkdirAll(filepath.Dir(filePath), 0750); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(filePath), err)
	}

	if err := os.WriteFile(filePath, []byte(bytecode), 0600); err != nil {
		return fmt.Errorf("failed to write bytecode to %s: %w", filePath, err)
	}

	fmt.Printf("  ✓ Extracted bytecode: %s\n", relPath)
	return nil
}

// fixABIInternalTypes fixes malformed internalType fields in ABI JSON strings.
// The Go binding generator sometimes omits spaces after type keywords like "contract", "struct", and "enum".
// For example: "contractIERC20" should be "contract IERC20"
func fixABIInternalTypes(abi string) string {
	// Pattern matches: internalType":"<keyword><CapitalLetter>
	// where keyword is "contract", "struct", or "enum"
	// This regex finds cases where there's no space between the keyword and the type name
	pattern := regexp.MustCompile(`("internalType":")(contract|struct|enum)([A-Z])`)

	// Replace with a space between the keyword and the type name
	// $1 = '"internalType":"', $2 = keyword (contract/struct/enum), $3 = capital letter
	return pattern.ReplaceAllString(abi, `${1}${2} ${3}`)
}

func writeABI(abiDir, version, baseFilename, abi string) error {
	// Fix malformed internalType fields before writing
	abi = fixABIInternalTypes(abi)

	filename := baseFilename + ".json"
	filePath := filepath.Join(abiDir, version, filename)
	relPath := filepath.Join(version, filename)

	if err := os.MkdirAll(filepath.Dir(filePath), 0750); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(filePath), err)
	}

	if err := os.WriteFile(filePath, []byte(abi), 0600); err != nil {
		return fmt.Errorf("failed to write ABI to %s: %w", filePath, err)
	}

	fmt.Printf("  ✓ Extracted ABI: %s\n", relPath)
	return nil
}

func extractMetadata(filePath string) (Metadata, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to parse file: %w", err)
	}

	var metadata Metadata

	// Walk through the AST to find the MetaData variable with Bin and ABI fields
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for variable declarations
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			return true
		}

		if md := extractMetadataFromVarDecl(genDecl); md.Bytecode != "" || md.ABI != "" {
			metadata = md
			return false // Stop walking, we found what we need
		}

		return true
	})

	return metadata, nil
}

// extractMetadataFromVarDecl extracts the Bin and ABI fields from a variable declaration if it's a MetaData variable
func extractMetadataFromVarDecl(genDecl *ast.GenDecl) Metadata {
	for _, spec := range genDecl.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok || len(valueSpec.Values) == 0 {
			continue
		}

		// Check if this is a MetaData variable
		if !isMetaDataVar(valueSpec) {
			continue
		}

		// Extract the Bin and ABI fields from the composite literal
		if md := extractMetadataFromValue(valueSpec.Values[0]); md.Bytecode != "" || md.ABI != "" {
			return md
		}
	}
	return Metadata{}
}

// isMetaDataVar checks if a value spec is a MetaData variable
func isMetaDataVar(valueSpec *ast.ValueSpec) bool {
	for _, name := range valueSpec.Names {
		if strings.HasSuffix(name.Name, "MetaData") {
			return true
		}
	}
	return false
}

// extractMetadataFromValue extracts the Bin and ABI fields from a value expression
func extractMetadataFromValue(value ast.Expr) Metadata {
	// Look for composite literal (&bind.MetaData{...})
	unaryExpr, ok := value.(*ast.UnaryExpr)
	if !ok {
		return Metadata{}
	}

	compositeLit, ok := unaryExpr.X.(*ast.CompositeLit)
	if !ok {
		return Metadata{}
	}

	return extractMetadataFromCompositeLit(compositeLit)
}

// extractMetadataFromCompositeLit extracts the Bin and ABI fields from a composite literal
func extractMetadataFromCompositeLit(compositeLit *ast.CompositeLit) Metadata {
	var metadata Metadata

	for _, elt := range compositeLit.Elts {
		kvExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}

		key, ok := kvExpr.Key.(*ast.Ident)
		if !ok {
			continue
		}

		// Extract the string literal value
		if basicLit, ok := kvExpr.Value.(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
			// Use strconv.Unquote to properly handle Go string literals with escape sequences
			value, err := strconv.Unquote(basicLit.Value)
			if err != nil {
				// Fallback to simple trim if unquote fails
				value = strings.Trim(basicLit.Value, `"`)
			}

			switch key.Name {
			case "Bin":
				metadata.Bytecode = value
			case "ABI":
				metadata.ABI = value
			}
		}
	}

	return metadata
}
