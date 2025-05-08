// Code generated from SimpleLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SimpleLang

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SimpleLangParser struct {
	*antlr.BaseParser
}

var SimpleLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplelangParserInit() {
	staticData := &SimpleLangParserStaticData
	staticData.LiteralNames = []string{
		"", "'@use'", "'\"'", "';'", "'='", "'print'", "'('", "')'", "'*'",
		"'/'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "ID", "INT", "IMPORTNAMES",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 50, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 4, 0, 8, 8, 0,
		11, 0, 12, 0, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 45, 8, 2, 10, 2, 12, 2, 48, 9, 2, 1, 2, 0, 1, 4, 3, 0, 2, 4,
		0, 2, 1, 0, 8, 9, 1, 0, 10, 11, 53, 0, 7, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0,
		4, 36, 1, 0, 0, 0, 6, 8, 3, 2, 1, 0, 7, 6, 1, 0, 0, 0, 8, 9, 1, 0, 0, 0,
		9, 7, 1, 0, 0, 0, 9, 10, 1, 0, 0, 0, 10, 1, 1, 0, 0, 0, 11, 12, 5, 1, 0,
		0, 12, 13, 5, 2, 0, 0, 13, 14, 5, 14, 0, 0, 14, 15, 5, 2, 0, 0, 15, 28,
		5, 3, 0, 0, 16, 17, 5, 12, 0, 0, 17, 18, 5, 4, 0, 0, 18, 19, 3, 4, 2, 0,
		19, 20, 5, 3, 0, 0, 20, 28, 1, 0, 0, 0, 21, 22, 5, 5, 0, 0, 22, 23, 5,
		6, 0, 0, 23, 24, 3, 4, 2, 0, 24, 25, 5, 7, 0, 0, 25, 26, 5, 3, 0, 0, 26,
		28, 1, 0, 0, 0, 27, 11, 1, 0, 0, 0, 27, 16, 1, 0, 0, 0, 27, 21, 1, 0, 0,
		0, 28, 3, 1, 0, 0, 0, 29, 30, 6, 2, -1, 0, 30, 37, 5, 13, 0, 0, 31, 37,
		5, 12, 0, 0, 32, 33, 5, 6, 0, 0, 33, 34, 3, 4, 2, 0, 34, 35, 5, 7, 0, 0,
		35, 37, 1, 0, 0, 0, 36, 29, 1, 0, 0, 0, 36, 31, 1, 0, 0, 0, 36, 32, 1,
		0, 0, 0, 37, 46, 1, 0, 0, 0, 38, 39, 10, 5, 0, 0, 39, 40, 7, 0, 0, 0, 40,
		45, 3, 4, 2, 6, 41, 42, 10, 4, 0, 0, 42, 43, 7, 1, 0, 0, 43, 45, 3, 4,
		2, 5, 44, 38, 1, 0, 0, 0, 44, 41, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44,
		1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 5, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0,
		5, 9, 27, 36, 44, 46,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SimpleLangParserInit initializes any static state used to implement SimpleLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSimpleLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleLangParserInit() {
	staticData := &SimpleLangParserStaticData
	staticData.once.Do(simplelangParserInit)
}

// NewSimpleLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSimpleLangParser(input antlr.TokenStream) *SimpleLangParser {
	SimpleLangParserInit()
	this := new(SimpleLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SimpleLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SimpleLang.g4"

	return this
}

// SimpleLangParser tokens.
const (
	SimpleLangParserEOF         = antlr.TokenEOF
	SimpleLangParserT__0        = 1
	SimpleLangParserT__1        = 2
	SimpleLangParserT__2        = 3
	SimpleLangParserT__3        = 4
	SimpleLangParserT__4        = 5
	SimpleLangParserT__5        = 6
	SimpleLangParserT__6        = 7
	SimpleLangParserT__7        = 8
	SimpleLangParserT__8        = 9
	SimpleLangParserT__9        = 10
	SimpleLangParserT__10       = 11
	SimpleLangParserID          = 12
	SimpleLangParserINT         = 13
	SimpleLangParserIMPORTNAMES = 14
	SimpleLangParserWS          = 15
)

// SimpleLangParser rules.
const (
	SimpleLangParserRULE_program    = 0
	SimpleLangParserRULE_statement  = 1
	SimpleLangParserRULE_expression = 2
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleLangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4130) != 0) {
		{
			p.SetState(6)
			p.Statement()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignmentContext struct {
	StatementContext
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintContext struct {
	StatementContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImportContext struct {
	StatementContext
}

func NewImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportContext {
	var p = new(ImportContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportContext) IMPORTNAMES() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserIMPORTNAMES, 0)
}

func (s *ImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterImport(s)
	}
}

func (s *ImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitImport(s)
	}
}

func (s *ImportContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitImport(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleLangParserRULE_statement)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SimpleLangParserT__0:
		localctx = NewImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(11)
			p.Match(SimpleLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(12)
			p.Match(SimpleLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(13)
			p.Match(SimpleLangParserIMPORTNAMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(14)
			p.Match(SimpleLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(15)
			p.Match(SimpleLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SimpleLangParserID:
		localctx = NewAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Match(SimpleLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.Match(SimpleLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)
			p.expression(0)
		}
		{
			p.SetState(19)
			p.Match(SimpleLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SimpleLangParserT__4:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(21)
			p.Match(SimpleLangParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.Match(SimpleLangParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.expression(0)
		}
		{
			p.SetState(24)
			p.Match(SimpleLangParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Match(SimpleLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExpressionContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExpressionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	ExpressionContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SimpleLangParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, SimpleLangParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SimpleLangParserINT:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(30)
			p.Match(SimpleLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SimpleLangParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.Match(SimpleLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SimpleLangParserT__5:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Match(SimpleLangParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.expression(0)
		}
		{
			p.SetState(34)
			p.Match(SimpleLangParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(39)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleLangParserT__7 || _la == SimpleLangParserT__8) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)
					p.expression(6)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleLangParserT__9 || _la == SimpleLangParserT__10) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(43)
					p.expression(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SimpleLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimpleLangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
