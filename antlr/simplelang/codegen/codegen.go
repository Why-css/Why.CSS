package codegen

import (
	"fmt"
	"simplelang/ast"
	"strings"
)

// GenerateGoCode takes a list of Stat nodes and returns Go code as a string.
func GenerateGoCode(stats []ast.Statement) string {
	var b strings.Builder
	b.WriteString("package main\n\nimport \"fmt\"\n\nfunc main() {\n")
	for _, stat := range stats {
		b.WriteString("    ")
		b.WriteString(genStat(stat))
		b.WriteString("\n")
	}
	b.WriteString("}\n")
	return b.String()
}

func genStat(stat ast.Statement) string {
	switch s := stat.(type) {
	case ast.Assign:
		return fmt.Sprintf("%s := %s", s.Name, genExpr(s.Expr))
	case ast.Print:
		return fmt.Sprintf("fmt.Println(%s)", genExpr(s.Expr))
	default:
		return "// unknown statement"
	}
}

func genExpr(expr ast.Expression) string {
	switch e := expr.(type) {
	case ast.IntLiteral:
		return fmt.Sprintf("%d", e.Value)
	case ast.VarRef:
		return e.Name
	case ast.BinOp:
		return fmt.Sprintf("(%s %s %s)", genExpr(e.Left), e.Op, genExpr(e.Right))
	default:
		return "/* unknown expr */"
	}
}
