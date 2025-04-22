package zksyncwrapper

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"strings"
)

func ReadBytecodeFromForgeJson(srcFile string) string {
	jsonData, err := os.ReadFile(srcFile)
	if err != nil {
		panic(err)
	}

	var bytecodeData struct {
		Bytecode struct {
			Object string `json:"object"`
		} `json:"bytecode"`
	}
	if err := json.Unmarshal(jsonData, &bytecodeData); err != nil {
		panic(err)
	}

	return bytecodeData.Bytecode.Object
}

//go:embed template.go
var zksyncDeployTemplate string

func WrapZksyncDeploy(bytecode, className, pkgName, outPath string) {
	fmt.Printf("Generating zk bytecode binding for %s\n", pkgName)

	fileNode := &ast.File{
		Name: ast.NewIdent(pkgName),
		Decls: []ast.Decl{
			declareImports(),
			declareDeployFunction(className),
			declareBytecodeVar(bytecode)}}

	writeFile(fileNode, outPath)
}

const comment = `// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
`

var importValues = []string{
	`"context"`,
	`"crypto/rand"`,
	`"fmt"`,
	`"github.com/ethereum/go-ethereum/accounts/abi/bind"`,
	`"github.com/ethereum/go-ethereum/common"`,
	`"github.com/zksync-sdk/zksync2-go/accounts"`,
	`"github.com/zksync-sdk/zksync2-go/clients"`,
	`"github.com/zksync-sdk/zksync2-go/types"`,
}

func declareImports() ast.Decl {

	specs := make([]ast.Spec, len(importValues))
	for i, value := range importValues {
		specs[i] = &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: value}}
	}

	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs}
}

func declareDeployFunction(contractName string) ast.Decl {
	template := zksyncDeployTemplate

	sep := "\n"
	lines := strings.Split(template, sep)
	from := 0
	to := 0
	// get the func body as string
	for !strings.Contains(lines[to], "return address, receipt, contract, nil") {
		if strings.Contains(lines[to], "DeployPlaceholderContractNameZk") {
			from = to
		}
		to++
	}
	template = strings.Join(lines[from+1:to+1], sep)
	template = template[1:] // remove the first space
	template = strings.Replace(template, "PlaceholderContractName", contractName, 2)

	return &ast.FuncDecl{
		Name: ast.NewIdent("Deploy" + contractName + "Zk"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{{
					Names: []*ast.Ident{ast.NewIdent("deployOpts")},
					Type:  &ast.Ident{Name: "*accounts.TransactOpts"}}, {
					Names: []*ast.Ident{ast.NewIdent("client")},
					Type:  &ast.Ident{Name: "*clients.Client"}}, {
					Names: []*ast.Ident{ast.NewIdent("wallet")},
					Type:  &ast.Ident{Name: "*accounts.Wallet"}}, {
					Names: []*ast.Ident{ast.NewIdent("backend")},
					Type:  &ast.Ident{Name: "bind.ContractBackend"}}, {
					Names: []*ast.Ident{ast.NewIdent("args")},
					Type:  &ast.Ellipsis{Elt: &ast.Ident{Name: "interface{}"}}}}},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{Type: &ast.Ident{Name: "common.Address"}},
					{Type: &ast.Ident{Name: "*types.Receipt"}},
					{Type: &ast.StarExpr{X: &ast.Ident{Name: contractName}}},
					{Type: &ast.Ident{Name: "error"}}}}},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.BasicLit{
						Kind:  token.STRING,
						Value: template}}}}}
}

func declareBytecodeVar(bytecode string) ast.Decl {
	return &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{ast.NewIdent("ZkBytecode")},
				Values: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("common"),
							Sel: ast.NewIdent("Hex2Bytes")},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf(`"%s"`, bytecode)}}}}}}}
}

func writeFile(fileNode *ast.File, dstFile string) {
	var buf bytes.Buffer
	fset := token.NewFileSet()
	if err := format.Node(&buf, fset, fileNode); err != nil {
		panic(err)
	}

	bs := buf.Bytes()
	bs = append([]byte(comment), bs...)

	if err := os.WriteFile(dstFile, bs, 0600); err != nil {
		panic(err)
	}
}
