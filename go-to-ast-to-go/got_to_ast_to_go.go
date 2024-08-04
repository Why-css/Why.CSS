package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	// Define the source code file to parse
	filename := "./example/example.go"

	// Create a new token file set
	fset := token.NewFileSet()

	// Parse the Go source file to generate an AST
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Print the AST
	ast.Print(fset, node)

	// Print the AST as Go code
	err = printer.Fprint(os.Stdout, fset, node)
	if err != nil {
		panic(err)
	}
}
