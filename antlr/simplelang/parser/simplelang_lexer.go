// Code generated from SimpleLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SimpleLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SimpleLangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplelanglexerLexerInit() {
	staticData := &SimpleLangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "';'", "'print'", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "ID", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"ID", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 67, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 51, 8, 9, 10, 9, 12, 9, 54, 9, 9, 1,
		10, 4, 10, 57, 8, 10, 11, 10, 12, 10, 58, 1, 11, 4, 11, 62, 8, 11, 11,
		11, 12, 11, 63, 1, 11, 1, 11, 0, 0, 12, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 1, 0, 5, 1, 0, 36, 36,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 69, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3,
		27, 1, 0, 0, 0, 5, 29, 1, 0, 0, 0, 7, 35, 1, 0, 0, 0, 9, 37, 1, 0, 0, 0,
		11, 39, 1, 0, 0, 0, 13, 41, 1, 0, 0, 0, 15, 43, 1, 0, 0, 0, 17, 45, 1,
		0, 0, 0, 19, 47, 1, 0, 0, 0, 21, 56, 1, 0, 0, 0, 23, 61, 1, 0, 0, 0, 25,
		26, 5, 61, 0, 0, 26, 2, 1, 0, 0, 0, 27, 28, 5, 59, 0, 0, 28, 4, 1, 0, 0,
		0, 29, 30, 5, 112, 0, 0, 30, 31, 5, 114, 0, 0, 31, 32, 5, 105, 0, 0, 32,
		33, 5, 110, 0, 0, 33, 34, 5, 116, 0, 0, 34, 6, 1, 0, 0, 0, 35, 36, 5, 40,
		0, 0, 36, 8, 1, 0, 0, 0, 37, 38, 5, 41, 0, 0, 38, 10, 1, 0, 0, 0, 39, 40,
		5, 42, 0, 0, 40, 12, 1, 0, 0, 0, 41, 42, 5, 47, 0, 0, 42, 14, 1, 0, 0,
		0, 43, 44, 5, 43, 0, 0, 44, 16, 1, 0, 0, 0, 45, 46, 5, 45, 0, 0, 46, 18,
		1, 0, 0, 0, 47, 48, 7, 0, 0, 0, 48, 52, 7, 1, 0, 0, 49, 51, 7, 2, 0, 0,
		50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1,
		0, 0, 0, 53, 20, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 57, 7, 3, 0, 0, 56,
		55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0,
		0, 59, 22, 1, 0, 0, 0, 60, 62, 7, 4, 0, 0, 61, 60, 1, 0, 0, 0, 62, 63,
		1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0,
		65, 66, 6, 11, 0, 0, 66, 24, 1, 0, 0, 0, 4, 0, 52, 58, 63, 1, 6, 0, 0,
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

// SimpleLangLexerInit initializes any static state used to implement SimpleLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleLangLexerInit() {
	staticData := &SimpleLangLexerLexerStaticData
	staticData.once.Do(simplelanglexerLexerInit)
}

// NewSimpleLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleLangLexer(input antlr.CharStream) *SimpleLangLexer {
	SimpleLangLexerInit()
	l := new(SimpleLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SimpleLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SimpleLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleLangLexer tokens.
const (
	SimpleLangLexerT__0 = 1
	SimpleLangLexerT__1 = 2
	SimpleLangLexerT__2 = 3
	SimpleLangLexerT__3 = 4
	SimpleLangLexerT__4 = 5
	SimpleLangLexerT__5 = 6
	SimpleLangLexerT__6 = 7
	SimpleLangLexerT__7 = 8
	SimpleLangLexerT__8 = 9
	SimpleLangLexerID   = 10
	SimpleLangLexerINT  = 11
	SimpleLangLexerWS   = 12
)
