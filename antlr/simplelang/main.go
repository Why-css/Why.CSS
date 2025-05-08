package main

import (
	"fmt"
	"os"

	"simplelang/ast"
	"simplelang/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewSimpleLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSimpleLangParser(stream)
	tree := p.Prog()

	ast := tree.Accept(&ast.ASTBuilder{})
	fmt.Printf("%#v\n", ast)
}
