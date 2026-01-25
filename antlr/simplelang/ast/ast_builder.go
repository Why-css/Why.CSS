package ast

import (
	"simplelang/parser"
	"strconv"
)

type ASTBuilder struct {
	parser.BaseSimpleLangVisitor
}

func (v *ASTBuilder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	var stats []Statement
	for _, statCtx := range ctx.AllStatement() {
		stat := statCtx.Accept(v).(Statement)
		stats = append(stats, stat)
	}
	return stats
}

func (v *ASTBuilder) VisitImport(ctx *parser.ImportContext) interface{} {
	name := ctx.IMPORTNAMES().GetText()
	return Import{Name: name}
}

func (v *ASTBuilder) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	name := ctx.ID().GetText()
	name = name[1:] // Remove the leading '$'
	expr := ctx.Expression().Accept(v).(Expression)
	return Assign{Name: name, Expr: expr}
}

func (v *ASTBuilder) VisitPrint(ctx *parser.PrintContext) interface{} {
	expr := ctx.Expression().Accept(v).(Expression)
	return Print{Expr: expr}
}

func (v *ASTBuilder) VisitInt(ctx *parser.IntContext) interface{} {
	val, _ := strconv.Atoi(ctx.INT().GetText())
	return IntLiteral{Value: val}
}

func (v *ASTBuilder) VisitId(ctx *parser.IdContext) interface{} {
	name := ctx.ID().GetText()
	name = name[1:] // Remove the leading '$'

	return VarRef{Name: name}
}

func (v *ASTBuilder) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	right := ctx.Expression(1).Accept(v).(Expression)
	op := ctx.GetOp().GetText()
	return BinOp{Left: left, Op: op, Right: right}
}

func (v *ASTBuilder) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	right := ctx.Expression(1).Accept(v).(Expression)
	op := ctx.GetOp().GetText()
	return BinOp{Left: left, Op: op, Right: right}
}

func (v *ASTBuilder) VisitParens(ctx *parser.ParensContext) interface{} {
	return ctx.Expression().Accept(v).(Expression)
}
