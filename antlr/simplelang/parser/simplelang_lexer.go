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
		"", "'@use'", "'\"'", "';'", "'='", "'print'", "'('", "')'", "'*'",
		"'/'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "ID", "INT", "IMPORTNAMES",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "ID", "INT", "IMPORTNAMES", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 87, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 64, 8, 11, 10, 11,
		12, 11, 67, 9, 11, 1, 12, 4, 12, 70, 8, 12, 11, 12, 12, 12, 71, 1, 13,
		1, 13, 5, 13, 76, 8, 13, 10, 13, 12, 13, 79, 9, 13, 1, 14, 4, 14, 82, 8,
		14, 11, 14, 12, 14, 83, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 1, 0, 5, 1, 0, 36, 36, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32,
		90, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0,
		0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0,
		0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1,
		0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 31,
		1, 0, 0, 0, 3, 36, 1, 0, 0, 0, 5, 38, 1, 0, 0, 0, 7, 40, 1, 0, 0, 0, 9,
		42, 1, 0, 0, 0, 11, 48, 1, 0, 0, 0, 13, 50, 1, 0, 0, 0, 15, 52, 1, 0, 0,
		0, 17, 54, 1, 0, 0, 0, 19, 56, 1, 0, 0, 0, 21, 58, 1, 0, 0, 0, 23, 60,
		1, 0, 0, 0, 25, 69, 1, 0, 0, 0, 27, 73, 1, 0, 0, 0, 29, 81, 1, 0, 0, 0,
		31, 32, 5, 64, 0, 0, 32, 33, 5, 117, 0, 0, 33, 34, 5, 115, 0, 0, 34, 35,
		5, 101, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 34, 0, 0, 37, 4, 1, 0, 0, 0,
		38, 39, 5, 59, 0, 0, 39, 6, 1, 0, 0, 0, 40, 41, 5, 61, 0, 0, 41, 8, 1,
		0, 0, 0, 42, 43, 5, 112, 0, 0, 43, 44, 5, 114, 0, 0, 44, 45, 5, 105, 0,
		0, 45, 46, 5, 110, 0, 0, 46, 47, 5, 116, 0, 0, 47, 10, 1, 0, 0, 0, 48,
		49, 5, 40, 0, 0, 49, 12, 1, 0, 0, 0, 50, 51, 5, 41, 0, 0, 51, 14, 1, 0,
		0, 0, 52, 53, 5, 42, 0, 0, 53, 16, 1, 0, 0, 0, 54, 55, 5, 47, 0, 0, 55,
		18, 1, 0, 0, 0, 56, 57, 5, 43, 0, 0, 57, 20, 1, 0, 0, 0, 58, 59, 5, 45,
		0, 0, 59, 22, 1, 0, 0, 0, 60, 61, 7, 0, 0, 0, 61, 65, 7, 1, 0, 0, 62, 64,
		7, 2, 0, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0,
		65, 66, 1, 0, 0, 0, 66, 24, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 70, 7,
		3, 0, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 26, 1, 0, 0, 0, 73, 77, 7, 1, 0, 0, 74, 76, 7, 2, 0,
		0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 28, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 82, 7, 4, 0, 0,
		81, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 6, 14, 0, 0, 86, 30, 1, 0, 0, 0, 5,
		0, 65, 71, 77, 83, 1, 6, 0, 0,
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
	SimpleLangLexerT__0        = 1
	SimpleLangLexerT__1        = 2
	SimpleLangLexerT__2        = 3
	SimpleLangLexerT__3        = 4
	SimpleLangLexerT__4        = 5
	SimpleLangLexerT__5        = 6
	SimpleLangLexerT__6        = 7
	SimpleLangLexerT__7        = 8
	SimpleLangLexerT__8        = 9
	SimpleLangLexerT__9        = 10
	SimpleLangLexerT__10       = 11
	SimpleLangLexerID          = 12
	SimpleLangLexerINT         = 13
	SimpleLangLexerIMPORTNAMES = 14
	SimpleLangLexerWS          = 15
)
