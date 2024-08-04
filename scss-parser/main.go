package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"
	"strings"
)

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func readFileContents(file *os.File) []byte {
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return byteValue
}

func unmarsalData(byteValue *[]byte) (data []interface{}) {
	err := json.Unmarshal(*byteValue, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func parseFunctionCall(data *[]interface{}, length int, writer io.Writer, iterator *int) {
	x := ""
	sel := ""
	argumentValue := ""
	for (*iterator) < length {
		for key, value := range (*data)[*iterator].(map[string]interface{}) {
			value := value.(string)
			*iterator++

			if key == "name" {
				// split value from "."
				value := strings.Split(value, ".")
				x = value[0]
				sel = value[1]
				continue
			}
			if key == "string" {
				argumentValue = value
			}
			if value == ";" {
				break
			}
		}
	}

	fmt.Println(getExpressionStatement(x, sel, []ast.Expr{getStringLiteral(argumentValue)}))
}

func getStringLiteral(stringValue string) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: "\"" + stringValue + "\"",
	}
}

func getIndent(name string) *ast.Ident {
	return &ast.Ident{Name: name}
}

func getExpressionStatement(x string, sel string, arguments []ast.Expr) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: x},
				Sel: &ast.Ident{Name: sel},
			},
			Args: arguments,
		},
	}
}

func getFunctionDeclaration(name string, body *ast.BlockStmt) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: &ast.Ident{Name: name},
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
		},
		Body: body,
	}
}

func getBlockStatement(list []ast.Stmt) *ast.BlockStmt {
	return &ast.BlockStmt{
		List: list,
	}
}

func test() {
	getFunctionDeclaration("main", getBlockStatement([]ast.Stmt{getExpressionStatement("fmt", "Println", []ast.Expr{getStringLiteral("Hello World")})}))
}

func transverseAST(data *[]interface{}, length int, end string, writer io.Writer, iterator *int) {

	if length > *iterator {

		for key, value := range (*data)[*iterator].(map[string]interface{}) {
			value := value.(string)
			*iterator++

			if value == end {
				return
			} else if key == "at" && value == "import" {
				// this creates "&ast.ExprStmt{"
				// work until ";"
				// within that find "fmt" "println" and "hello world"

				// fmt.Println("starting function call")
				// transverseAST(data, length, ";", writer, iterator)
				// fmt.Print("ending function call")
				parseFunctionCall(data, length, writer, iterator)
			} else {
				fmt.Println(key, value)
			}

		}

		*iterator++
		transverseAST(data, length, end, writer, iterator)
	}
}

func main() {
	file := openFile("tokenisedScss.json")
	defer file.Close()

	byteValue := readFileContents(file)

	data := unmarsalData(&byteValue)

	// iterate over data
	// for _, item := range data.([]interface{}) {
	// if type is "FunctionCall"
	// -> get the name until the first "("
	// -> get the arguments until the last ")"
	// -> end on ";"
	iterator := 0
	transverseAST(&data, len(data), "", os.Stdout, &iterator)

}
