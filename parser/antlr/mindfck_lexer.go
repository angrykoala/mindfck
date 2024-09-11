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
		"", "'('", "')'", "'{'", "'}'", "", "'byte'", "'print'", "'if'", "'+'",
		"'-'", "'*'", "'/'", "'='", "'=='", "", "", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "WS", "BYTE", "PRINT", "IF", "PLUS", "MINUS", "TIMES",
		"DIVIDE", "EQUALS", "DEQUALS", "AND", "OR", "GT", "GE", "LT", "LE",
		"IDENTIFIER", "NUMBER",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "WS", "BYTE", "PRINT", "IF", "PLUS",
		"MINUS", "TIMES", "DIVIDE", "EQUALS", "DEQUALS", "AND", "OR", "GT",
		"GE", "LT", "LE", "IDENTIFIER", "NUMBER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 121, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		4, 4, 55, 8, 4, 11, 4, 12, 4, 56, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 94, 8, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 3, 15, 100, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 4, 20, 113, 8, 20, 11, 20, 12, 20,
		114, 1, 21, 4, 21, 118, 8, 21, 11, 21, 12, 21, 119, 0, 0, 22, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 1, 0, 3, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65, 90, 97, 122, 1,
		0, 48, 57, 125, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0,
		0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0,
		0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1,
		0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1, 45,
		1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 49, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0, 9,
		54, 1, 0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 65, 1, 0, 0, 0, 15, 71, 1, 0, 0,
		0, 17, 74, 1, 0, 0, 0, 19, 76, 1, 0, 0, 0, 21, 78, 1, 0, 0, 0, 23, 80,
		1, 0, 0, 0, 25, 82, 1, 0, 0, 0, 27, 84, 1, 0, 0, 0, 29, 93, 1, 0, 0, 0,
		31, 99, 1, 0, 0, 0, 33, 101, 1, 0, 0, 0, 35, 103, 1, 0, 0, 0, 37, 106,
		1, 0, 0, 0, 39, 108, 1, 0, 0, 0, 41, 112, 1, 0, 0, 0, 43, 117, 1, 0, 0,
		0, 45, 46, 5, 40, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48, 5, 41, 0, 0, 48, 4,
		1, 0, 0, 0, 49, 50, 5, 123, 0, 0, 50, 6, 1, 0, 0, 0, 51, 52, 5, 125, 0,
		0, 52, 8, 1, 0, 0, 0, 53, 55, 7, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1,
		0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58,
		59, 6, 4, 0, 0, 59, 10, 1, 0, 0, 0, 60, 61, 5, 98, 0, 0, 61, 62, 5, 121,
		0, 0, 62, 63, 5, 116, 0, 0, 63, 64, 5, 101, 0, 0, 64, 12, 1, 0, 0, 0, 65,
		66, 5, 112, 0, 0, 66, 67, 5, 114, 0, 0, 67, 68, 5, 105, 0, 0, 68, 69, 5,
		110, 0, 0, 69, 70, 5, 116, 0, 0, 70, 14, 1, 0, 0, 0, 71, 72, 5, 105, 0,
		0, 72, 73, 5, 102, 0, 0, 73, 16, 1, 0, 0, 0, 74, 75, 5, 43, 0, 0, 75, 18,
		1, 0, 0, 0, 76, 77, 5, 45, 0, 0, 77, 20, 1, 0, 0, 0, 78, 79, 5, 42, 0,
		0, 79, 22, 1, 0, 0, 0, 80, 81, 5, 47, 0, 0, 81, 24, 1, 0, 0, 0, 82, 83,
		5, 61, 0, 0, 83, 26, 1, 0, 0, 0, 84, 85, 5, 61, 0, 0, 85, 86, 5, 61, 0,
		0, 86, 28, 1, 0, 0, 0, 87, 88, 5, 65, 0, 0, 88, 89, 5, 78, 0, 0, 89, 94,
		5, 68, 0, 0, 90, 91, 5, 97, 0, 0, 91, 92, 5, 110, 0, 0, 92, 94, 5, 100,
		0, 0, 93, 87, 1, 0, 0, 0, 93, 90, 1, 0, 0, 0, 94, 30, 1, 0, 0, 0, 95, 96,
		5, 79, 0, 0, 96, 100, 5, 82, 0, 0, 97, 98, 5, 111, 0, 0, 98, 100, 5, 114,
		0, 0, 99, 95, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 32, 1, 0, 0, 0, 101,
		102, 5, 62, 0, 0, 102, 34, 1, 0, 0, 0, 103, 104, 5, 62, 0, 0, 104, 105,
		5, 61, 0, 0, 105, 36, 1, 0, 0, 0, 106, 107, 5, 60, 0, 0, 107, 38, 1, 0,
		0, 0, 108, 109, 5, 60, 0, 0, 109, 110, 5, 61, 0, 0, 110, 40, 1, 0, 0, 0,
		111, 113, 7, 1, 0, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 42, 1, 0, 0, 0, 116, 118, 7,
		2, 0, 0, 117, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 117, 1, 0, 0,
		0, 119, 120, 1, 0, 0, 0, 120, 44, 1, 0, 0, 0, 6, 0, 56, 93, 99, 114, 119,
		1, 0, 1, 0,
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
	mindfckLexerT__0       = 1
	mindfckLexerT__1       = 2
	mindfckLexerT__2       = 3
	mindfckLexerT__3       = 4
	mindfckLexerWS         = 5
	mindfckLexerBYTE       = 6
	mindfckLexerPRINT      = 7
	mindfckLexerIF         = 8
	mindfckLexerPLUS       = 9
	mindfckLexerMINUS      = 10
	mindfckLexerTIMES      = 11
	mindfckLexerDIVIDE     = 12
	mindfckLexerEQUALS     = 13
	mindfckLexerDEQUALS    = 14
	mindfckLexerAND        = 15
	mindfckLexerOR         = 16
	mindfckLexerGT         = 17
	mindfckLexerGE         = 18
	mindfckLexerLT         = 19
	mindfckLexerLE         = 20
	mindfckLexerIDENTIFIER = 21
	mindfckLexerNUMBER     = 22
)
