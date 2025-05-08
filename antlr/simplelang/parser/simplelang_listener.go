// Code generated from SimpleLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SimpleLang

import "github.com/antlr4-go/antlr/v4"

// SimpleLangListener is a complete listener for a parse tree produced by SimpleLangParser.
type SimpleLangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterImport is called when entering the Import production.
	EnterImport(c *ImportContext)

	// EnterAssignment is called when entering the Assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterId is called when entering the Id production.
	EnterId(c *IdContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitImport is called when exiting the Import production.
	ExitImport(c *ImportContext)

	// ExitAssignment is called when exiting the Assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitId is called when exiting the Id production.
	ExitId(c *IdContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)
}
