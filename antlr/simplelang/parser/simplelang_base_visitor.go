// Code generated from SimpleLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SimpleLang

import "github.com/antlr4-go/antlr/v4"

type BaseSimpleLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitImport(ctx *ImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}
