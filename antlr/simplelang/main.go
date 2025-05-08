package main

import (
	"fmt"
	"os"

	"simplelang/ast"
	"simplelang/codegen"
	"simplelang/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// input, _ := antlr.NewFileStream(os.Args[1])
	input, _ := antlr.NewFileStream("input")
	lexer := parser.NewSimpleLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSimpleLangParser(stream)
	tree := p.Prog()

	builtAst := tree.Accept(&ast.ASTBuilder{})
	fmt.Printf("%#v\n", builtAst)

	// Generate Go code
	goCode := codegen.GenerateGoCode(builtAst.([]ast.Stat))
	err := os.MkdirAll("./generated", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	_, err = os.Create("./generated/generated.go")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	err = os.WriteFile("generated/generated.go", []byte(goCode), 0644)
	if err != nil {
		fmt.Println("Error writing Go code:", err)
		return
	}
	fmt.Println("Go code written to generated.go")
}
