// Code generated from parser/antlr/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type mindfckLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MindfckLexerLexerStaticData struct {
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

func mindfcklexerLexerInit() {
	staticData := &MindfckLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'byte'", "'print'", "'+'", "'-'", "'*'", "'/'", "'='", "'=='",
		"", "", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "BYTE", "PRINT", "PLUS", "MINUS", "TIMES", "DIVIDE", "EQUALS",
		"DEQUALS", "AND", "OR", "GT", "GE", "LT", "LE", "IDENTIFIER", "NUMBER",
	}
	staticData.RuleNames = []string{
		"WS", "BYTE", "PRINT", "PLUS", "MINUS", "TIMES", "DIVIDE", "EQUALS",
		"DEQUALS", "AND", "OR", "GT", "GE", "LT", "LE", "IDENTIFIER", "NUMBER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 100, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 4, 0, 37, 8, 0, 11, 0, 12, 0, 38, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 73, 8, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 3, 10, 79, 8, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 15, 4, 15, 92, 8, 15, 11, 15, 12, 15, 93, 1, 16,
		4, 16, 97, 8, 16, 11, 16, 12, 16, 98, 0, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 1, 0, 3, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65,
		90, 97, 122, 1, 0, 48, 57, 104, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 36, 1, 0,
		0, 0, 3, 42, 1, 0, 0, 0, 5, 47, 1, 0, 0, 0, 7, 53, 1, 0, 0, 0, 9, 55, 1,
		0, 0, 0, 11, 57, 1, 0, 0, 0, 13, 59, 1, 0, 0, 0, 15, 61, 1, 0, 0, 0, 17,
		63, 1, 0, 0, 0, 19, 72, 1, 0, 0, 0, 21, 78, 1, 0, 0, 0, 23, 80, 1, 0, 0,
		0, 25, 82, 1, 0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 87, 1, 0, 0, 0, 31, 91,
		1, 0, 0, 0, 33, 96, 1, 0, 0, 0, 35, 37, 7, 0, 0, 0, 36, 35, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 1,
		0, 0, 0, 40, 41, 6, 0, 0, 0, 41, 2, 1, 0, 0, 0, 42, 43, 5, 98, 0, 0, 43,
		44, 5, 121, 0, 0, 44, 45, 5, 116, 0, 0, 45, 46, 5, 101, 0, 0, 46, 4, 1,
		0, 0, 0, 47, 48, 5, 112, 0, 0, 48, 49, 5, 114, 0, 0, 49, 50, 5, 105, 0,
		0, 50, 51, 5, 110, 0, 0, 51, 52, 5, 116, 0, 0, 52, 6, 1, 0, 0, 0, 53, 54,
		5, 43, 0, 0, 54, 8, 1, 0, 0, 0, 55, 56, 5, 45, 0, 0, 56, 10, 1, 0, 0, 0,
		57, 58, 5, 42, 0, 0, 58, 12, 1, 0, 0, 0, 59, 60, 5, 47, 0, 0, 60, 14, 1,
		0, 0, 0, 61, 62, 5, 61, 0, 0, 62, 16, 1, 0, 0, 0, 63, 64, 5, 61, 0, 0,
		64, 65, 5, 61, 0, 0, 65, 18, 1, 0, 0, 0, 66, 67, 5, 65, 0, 0, 67, 68, 5,
		78, 0, 0, 68, 73, 5, 68, 0, 0, 69, 70, 5, 97, 0, 0, 70, 71, 5, 110, 0,
		0, 71, 73, 5, 100, 0, 0, 72, 66, 1, 0, 0, 0, 72, 69, 1, 0, 0, 0, 73, 20,
		1, 0, 0, 0, 74, 75, 5, 79, 0, 0, 75, 79, 5, 82, 0, 0, 76, 77, 5, 111, 0,
		0, 77, 79, 5, 114, 0, 0, 78, 74, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 22,
		1, 0, 0, 0, 80, 81, 5, 62, 0, 0, 81, 24, 1, 0, 0, 0, 82, 83, 5, 62, 0,
		0, 83, 84, 5, 61, 0, 0, 84, 26, 1, 0, 0, 0, 85, 86, 5, 60, 0, 0, 86, 28,
		1, 0, 0, 0, 87, 88, 5, 60, 0, 0, 88, 89, 5, 61, 0, 0, 89, 30, 1, 0, 0,
		0, 90, 92, 7, 1, 0, 0, 91, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 91,
		1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 32, 1, 0, 0, 0, 95, 97, 7, 2, 0, 0,
		96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 34, 1, 0, 0, 0, 6, 0, 38, 72, 78, 93, 98, 1, 0, 1, 0,
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

// mindfckLexerInit initializes any static state used to implement mindfckLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewmindfckLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MindfckLexerInit() {
	staticData := &MindfckLexerLexerStaticData
	staticData.once.Do(mindfcklexerLexerInit)
}

// NewmindfckLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewmindfckLexer(input antlr.CharStream) *mindfckLexer {
	MindfckLexerInit()
	l := new(mindfckLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MindfckLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "mindfck.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// mindfckLexer tokens.
const (
	mindfckLexerWS         = 1
	mindfckLexerBYTE       = 2
	mindfckLexerPRINT      = 3
	mindfckLexerPLUS       = 4
	mindfckLexerMINUS      = 5
	mindfckLexerTIMES      = 6
	mindfckLexerDIVIDE     = 7
	mindfckLexerEQUALS     = 8
	mindfckLexerDEQUALS    = 9
	mindfckLexerAND        = 10
	mindfckLexerOR         = 11
	mindfckLexerGT         = 12
	mindfckLexerGE         = 13
	mindfckLexerLT         = 14
	mindfckLexerLE         = 15
	mindfckLexerIDENTIFIER = 16
	mindfckLexerNUMBER     = 17
)
