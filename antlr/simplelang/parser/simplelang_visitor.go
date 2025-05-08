// Code generated from SimpleLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SimpleLang

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SimpleLangParser.
type SimpleLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#Import.
	VisitImport(ctx *ImportContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#Assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#Print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#Id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#Int.
	VisitInt(ctx *IntContext) interface{}
}
