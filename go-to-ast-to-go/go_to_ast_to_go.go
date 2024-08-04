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

	// Create an output file to write the AST
	outputAstFile, err := os.Create("./output/ast_output.txt")
	if err != nil {
		fmt.Println("Error creating AST output file:", err)
		return
	}
	defer outputAstFile.Close()

	// Print the AST as text to the output file
	err = ast.Fprint(outputAstFile, fset, node, nil)
	if err != nil {
		fmt.Println("Error printing AST:", err)
		return
	}

	// Create an output file to write the generated Go code
	outputGoFile, err := os.Create("./output/output.go")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputGoFile.Close()

	// Print the AST as Go code
	err = printer.Fprint(outputGoFile, fset, node)
	if err != nil {
		panic(err)
	}
}
