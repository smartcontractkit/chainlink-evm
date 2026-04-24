package main

import (
	"flag"
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
	// Parse command-line flags
	inputDir := flag.String("input", "", "Input directory containing Go wrapper files (required)")
	bytecodeDir := flag.String("bytecode", "", "Output directory for bytecode files (required)")
	abiDir := flag.String("abi", "", "Output directory for ABI files (required)")
	includeLatest := flag.Bool("include-latest", false, "Include 'latest' directories in processing (default: false, excludes 'latest' directories)")
	flag.Parse()

	// Validate required arguments
	if *bytecodeDir == "" {
		return fmt.Errorf("bytecode directory is required (use -bytecode flag)")
	}
	if *abiDir == "" {
		return fmt.Errorf("abi directory is required (use -abi flag)")
	}

	// Validate input directory exists
	if _, err := os.Stat(*inputDir); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("input directory does not exist: %s", *inputDir)
		}
		return fmt.Errorf("failed to stat input directory %s: %w", *inputDir, err)
	}

	fmt.Printf("Input dir: %s\n", *inputDir)
	fmt.Printf("Bytecode dir: %s\n", *bytecodeDir)
	fmt.Printf("ABI dir: %s\n", *abiDir)
	if *includeLatest {
		fmt.Printf("Including 'latest' directories\n")
	}

	// Process the input directory
	excludeLatest := !*includeLatest
	return processDirectory(*inputDir, *inputDir, *bytecodeDir, *abiDir, excludeLatest)
}

// processDirectory processes a directory and all its subdirectories for Go wrapper files
func processDirectory(dir, baseDir, bytecodeDir, abiDir string, excludeLatest bool) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())

		// Check if "latest" appears anywhere in the relative path
		if excludeLatest {
			relPath, err := filepath.Rel(baseDir, fullPath)
			if err == nil && containsLatestInPath(relPath) {
				continue
			}
		}

		if entry.IsDir() {
			// Recursively process subdirectories
			if err := processDirectory(fullPath, baseDir, bytecodeDir, abiDir, excludeLatest); err != nil {
				return err
			}
		} else if strings.HasSuffix(entry.Name(), ".go") {
			// Process Go files
			if err := processGoFile(fullPath, baseDir, bytecodeDir, abiDir); err != nil {
				return err
			}
		}
	}

	return nil
}

// containsLatestInPath checks if "latest" appears as a path component
func containsLatestInPath(path string) bool {
	parts := strings.Split(filepath.ToSlash(path), "/")
	for _, part := range parts {
		if part == "latest" {
			return true
		}
	}
	return false
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
