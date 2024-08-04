// package main

// import (
// 	"fmt"
// 	"go/ast"
// 	"go/printer"
// 	"go/token"
// 	"os"
// )

// func main() {
// 	// Create a new token file set
// 	fset := token.NewFileSet()

// 	// Create the minimal AST for a "Hello, World!" program
// 	helloWorldAST := createHelloWorldAST()

// 	// Create an output file to write the generated Go code
// 	outputGoFile, err := os.Create("./output/output.go")
// 	if err != nil {
// 		fmt.Println("Error creating output file:", err)
// 		return
// 	}
// 	defer outputGoFile.Close()

// 	// Print the AST as Go code
// 	err = printer.Fprint(outputGoFile, fset, helloWorldAST)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // createHelloWorldAST creates the AST for a "Hello, World!" program
// func createHelloWorldAST() *ast.File {
// 	// Create package declaration
// 	pkg := &ast.Ident{Name: "main"}

// 	// Create import declaration for "fmt"
// 	importSpec := &ast.ImportSpec{
// 		Path: &ast.BasicLit{
// 			Kind:  token.STRING,
// 			Value: `"fmt"`,
// 		},
// 	}
// 	importDecl := &ast.GenDecl{
// 		Tok:   token.IMPORT,
// 		Specs: []ast.Spec{importSpec},
// 	}

// 	// Create the variable declaration for the "Hello, World!" string
// 	helloWorldVar := &ast.ValueSpec{
// 		Names: []*ast.Ident{ast.NewIdent("helloWorld")},
// 		Type:  ast.NewIdent("string"),
// 		Values: []ast.Expr{
// 			&ast.BasicLit{
// 				Kind:  token.STRING,
// 				Value: `"Hello, World!"`,
// 			},
// 		},
// 	}
// 	varDecl := &ast.GenDecl{
// 		Tok:   token.VAR,
// 		Specs: []ast.Spec{helloWorldVar},
// 	}

// 	// Create the main function declaration
// 	mainFunc := &ast.FuncDecl{
// 		Name: &ast.Ident{Name: "main"},
// 		Type: &ast.FuncType{
// 			Params: &ast.FieldList{},
// 		},
// 		Body: &ast.BlockStmt{
// 			List: []ast.Stmt{
// 				&ast.DeclStmt{Decl: varDecl},
// 				&ast.ExprStmt{
// 					X: &ast.CallExpr{
// 						Fun: &ast.SelectorExpr{
// 							X:   &ast.Ident{Name: "fmt"},
// 							Sel: &ast.Ident{Name: "Println"},
// 						},
// 						Args: []ast.Expr{
// 							&ast.Ident{Name: "helloWorld"},
// 							&ast.BasicLit{
// 								Kind:  token.STRING,
// 								Value: `"From CSS??"`,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	// Create the file node
// 	file := &ast.File{
// 		Name:  pkg,
// 		Decls: []ast.Decl{importDecl, mainFunc},
// 	}

// 	return file
// }

package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	// Create a new token file set
	fset := token.NewFileSet()

	// Create the minimal AST for a "Hello, World!" program
	helloWorldAST := createHelloWorldAST()

	// Create an output file to write the generated Go code
	outputGoFile, err := os.Create("./output/output.go")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputGoFile.Close()

	// Print the AST as Go code
	err = printer.Fprint(outputGoFile, fset, helloWorldAST)
	if err != nil {
		panic(err)
	}
}

// createHelloWorldAST creates the AST for a "Hello, World!" program
func createHelloWorldAST() *ast.File {
	// Create package declaration
	pkg := &ast.Ident{Name: "main"}

	// TODO create function for this
	importSpec := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"fmt"`,
		},
	}
	importDecl := &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{importSpec},
	}

	varDecl := getStringVariableDeclaration("greet", "Hello, World!")

	mainFunc := getFunctionDeclaration("main", getBlockStatement([]ast.Stmt{
		&ast.DeclStmt{Decl: varDecl},
		getExpressionStatement("fmt", "Println", []ast.Expr{
			getIndent("greet"),
			getStringLiteral("From CSS??"),
		}),
	}))

	// Create the file node
	file := &ast.File{
		Name:  pkg,
		Decls: []ast.Decl{importDecl, mainFunc},
	}

	return file
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

func getStringVariableDeclaration(name string, value string) *ast.GenDecl {
	variable := &ast.ValueSpec{
		Names: []*ast.Ident{ast.NewIdent(name)},
		Type:  ast.NewIdent("string"),
		Values: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"" + value + "\"",
			},
		},
	}

	return &ast.GenDecl{
		Tok:   token.VAR,
		Specs: []ast.Spec{variable},
	}
}
