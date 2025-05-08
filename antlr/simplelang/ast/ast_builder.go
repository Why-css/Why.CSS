package ast

import (
	"simplelang/parser"
	"strconv"
)

type ASTBuilder struct {
	parser.BaseSimpleLangVisitor
}

func (v *ASTBuilder) VisitProg(ctx *parser.ProgContext) interface{} {
	var stats []Stat
	for _, statCtx := range ctx.AllStat() {
		stat := statCtx.Accept(v).(Stat)
		stats = append(stats, stat)
	}
	return stats
}

func (v *ASTBuilder) VisitStat(ctx *parser.StatContext) interface{} {
	if ctx.ID() == nil { // print statement
		expr := ctx.Expr().Accept(v).(Expr)
		return Print{Expr: expr}
	}
	name := ctx.ID().GetText()
	expr := ctx.Expr().Accept(v).(Expr)
	return Assign{Name: name, Expr: expr}
}

func (v *ASTBuilder) VisitInt(ctx *parser.IntContext) interface{} {
	val, _ := strconv.Atoi(ctx.INT().GetText())
	return IntLiteral{Value: val}
}

func (v *ASTBuilder) VisitId(ctx *parser.IdContext) interface{} {
	return VarRef{Name: ctx.ID().GetText()}
}

func (v *ASTBuilder) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := ctx.Expr(0).Accept(v).(Expr)
	right := ctx.Expr(1).Accept(v).(Expr)
	op := ctx.GetOp().GetText()
	return BinOp{Left: left, Op: op, Right: right}
}

func (v *ASTBuilder) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := ctx.Expr(0).Accept(v).(Expr)
	right := ctx.Expr(1).Accept(v).(Expr)
	op := ctx.GetOp().GetText()
	return BinOp{Left: left, Op: op, Right: right}
}

func (v *ASTBuilder) VisitParens(ctx *parser.ParensContext) interface{} {
	return ctx.Expr().Accept(v).(Expr)
}
