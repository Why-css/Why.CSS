// Code generated from SimpleLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SimpleLang

import "github.com/antlr4-go/antlr/v4"

// BaseSimpleLangListener is a complete listener for a parse tree produced by SimpleLangParser.
type BaseSimpleLangListener struct{}

var _ SimpleLangListener = &BaseSimpleLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseSimpleLangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseSimpleLangListener) ExitProg(ctx *ProgContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseSimpleLangListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseSimpleLangListener) ExitStat(ctx *StatContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseSimpleLangListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseSimpleLangListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseSimpleLangListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseSimpleLangListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseSimpleLangListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseSimpleLangListener) ExitParens(ctx *ParensContext) {}

// EnterId is called when production Id is entered.
func (s *BaseSimpleLangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production Id is exited.
func (s *BaseSimpleLangListener) ExitId(ctx *IdContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseSimpleLangListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseSimpleLangListener) ExitInt(ctx *IntContext) {}
