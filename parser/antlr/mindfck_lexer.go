// Code generated from parser/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "'('", "')'", "'{'", "'}'", "", "'byte'", "'print'", "'if'", "'else'",
		"'while'", "'+'", "'-'", "'*'", "'/'", "'='", "'=='", "'and'", "'or'",
		"", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "WS", "BYTE", "PRINT", "IF", "ELSE", "WHILE", "PLUS",
		"MINUS", "TIMES", "DIVIDE", "EQUALS", "DEQUALS", "AND", "OR", "NOT",
		"GT", "GE", "LT", "LE", "IDENTIFIER", "NUMBER", "CHAR",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "WS", "BYTE", "PRINT", "IF", "ELSE",
		"WHILE", "PLUS", "MINUS", "TIMES", "DIVIDE", "EQUALS", "DEQUALS", "AND",
		"OR", "NOT", "GT", "GE", "LT", "LE", "IDENTIFIER", "NUMBER", "CHAR",
		"EXT_ASCII_CHAR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 148, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 4,
		4, 65, 8, 4, 11, 4, 12, 4, 66, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 121, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 23, 4, 23, 134, 8, 23, 11, 23, 12, 23, 135, 1,
		24, 4, 24, 139, 8, 24, 11, 24, 12, 24, 140, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 0, 1, 0, 4, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65, 90, 97, 122, 1,
		0, 48, 57, 1, 0, 0, 255, 150, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 1, 55, 1, 0, 0, 0, 3, 57, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61,
		1, 0, 0, 0, 9, 64, 1, 0, 0, 0, 11, 70, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0,
		15, 81, 1, 0, 0, 0, 17, 84, 1, 0, 0, 0, 19, 89, 1, 0, 0, 0, 21, 95, 1,
		0, 0, 0, 23, 97, 1, 0, 0, 0, 25, 99, 1, 0, 0, 0, 27, 101, 1, 0, 0, 0, 29,
		103, 1, 0, 0, 0, 31, 105, 1, 0, 0, 0, 33, 108, 1, 0, 0, 0, 35, 112, 1,
		0, 0, 0, 37, 120, 1, 0, 0, 0, 39, 122, 1, 0, 0, 0, 41, 124, 1, 0, 0, 0,
		43, 127, 1, 0, 0, 0, 45, 129, 1, 0, 0, 0, 47, 133, 1, 0, 0, 0, 49, 138,
		1, 0, 0, 0, 51, 142, 1, 0, 0, 0, 53, 146, 1, 0, 0, 0, 55, 56, 5, 40, 0,
		0, 56, 2, 1, 0, 0, 0, 57, 58, 5, 41, 0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5,
		123, 0, 0, 60, 6, 1, 0, 0, 0, 61, 62, 5, 125, 0, 0, 62, 8, 1, 0, 0, 0,
		63, 65, 7, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 6, 4, 0, 0, 69,
		10, 1, 0, 0, 0, 70, 71, 5, 98, 0, 0, 71, 72, 5, 121, 0, 0, 72, 73, 5, 116,
		0, 0, 73, 74, 5, 101, 0, 0, 74, 12, 1, 0, 0, 0, 75, 76, 5, 112, 0, 0, 76,
		77, 5, 114, 0, 0, 77, 78, 5, 105, 0, 0, 78, 79, 5, 110, 0, 0, 79, 80, 5,
		116, 0, 0, 80, 14, 1, 0, 0, 0, 81, 82, 5, 105, 0, 0, 82, 83, 5, 102, 0,
		0, 83, 16, 1, 0, 0, 0, 84, 85, 5, 101, 0, 0, 85, 86, 5, 108, 0, 0, 86,
		87, 5, 115, 0, 0, 87, 88, 5, 101, 0, 0, 88, 18, 1, 0, 0, 0, 89, 90, 5,
		119, 0, 0, 90, 91, 5, 104, 0, 0, 91, 92, 5, 105, 0, 0, 92, 93, 5, 108,
		0, 0, 93, 94, 5, 101, 0, 0, 94, 20, 1, 0, 0, 0, 95, 96, 5, 43, 0, 0, 96,
		22, 1, 0, 0, 0, 97, 98, 5, 45, 0, 0, 98, 24, 1, 0, 0, 0, 99, 100, 5, 42,
		0, 0, 100, 26, 1, 0, 0, 0, 101, 102, 5, 47, 0, 0, 102, 28, 1, 0, 0, 0,
		103, 104, 5, 61, 0, 0, 104, 30, 1, 0, 0, 0, 105, 106, 5, 61, 0, 0, 106,
		107, 5, 61, 0, 0, 107, 32, 1, 0, 0, 0, 108, 109, 5, 97, 0, 0, 109, 110,
		5, 110, 0, 0, 110, 111, 5, 100, 0, 0, 111, 34, 1, 0, 0, 0, 112, 113, 5,
		111, 0, 0, 113, 114, 5, 114, 0, 0, 114, 36, 1, 0, 0, 0, 115, 116, 5, 110,
		0, 0, 116, 117, 5, 111, 0, 0, 117, 118, 5, 116, 0, 0, 118, 121, 5, 32,
		0, 0, 119, 121, 5, 33, 0, 0, 120, 115, 1, 0, 0, 0, 120, 119, 1, 0, 0, 0,
		121, 38, 1, 0, 0, 0, 122, 123, 5, 62, 0, 0, 123, 40, 1, 0, 0, 0, 124, 125,
		5, 62, 0, 0, 125, 126, 5, 61, 0, 0, 126, 42, 1, 0, 0, 0, 127, 128, 5, 60,
		0, 0, 128, 44, 1, 0, 0, 0, 129, 130, 5, 60, 0, 0, 130, 131, 5, 61, 0, 0,
		131, 46, 1, 0, 0, 0, 132, 134, 7, 1, 0, 0, 133, 132, 1, 0, 0, 0, 134, 135,
		1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 48, 1, 0,
		0, 0, 137, 139, 7, 2, 0, 0, 138, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 50, 1, 0, 0, 0, 142, 143,
		5, 39, 0, 0, 143, 144, 3, 53, 26, 0, 144, 145, 5, 39, 0, 0, 145, 52, 1,
		0, 0, 0, 146, 147, 7, 3, 0, 0, 147, 54, 1, 0, 0, 0, 5, 0, 66, 120, 135,
		140, 1, 0, 1, 0,
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
	mindfckLexerELSE       = 9
	mindfckLexerWHILE      = 10
	mindfckLexerPLUS       = 11
	mindfckLexerMINUS      = 12
	mindfckLexerTIMES      = 13
	mindfckLexerDIVIDE     = 14
	mindfckLexerEQUALS     = 15
	mindfckLexerDEQUALS    = 16
	mindfckLexerAND        = 17
	mindfckLexerOR         = 18
	mindfckLexerNOT        = 19
	mindfckLexerGT         = 20
	mindfckLexerGE         = 21
	mindfckLexerLT         = 22
	mindfckLexerLE         = 23
	mindfckLexerIDENTIFIER = 24
	mindfckLexerNUMBER     = 25
	mindfckLexerCHAR       = 26
)
