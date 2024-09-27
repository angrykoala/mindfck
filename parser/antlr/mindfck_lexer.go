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
		"", "'('", "')'", "'{'", "'}'", "", "", "", "'byte'", "'int'", "'print'",
		"'if'", "'else'", "'while'", "'read'", "'+'", "'-'", "'*'", "'/'", "'='",
		"'=='", "'and'", "'or'", "", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "CHAR", "BYTE_NUMBER", "WS", "BYTE", "INT", "PRINT",
		"IF", "ELSE", "WHILE", "READ", "PLUS", "MINUS", "TIMES", "DIVIDE", "EQUALS",
		"DEQUALS", "AND", "OR", "NOT", "GT", "GE", "LT", "LE", "NUMBER", "IDENTIFIER",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "CHAR", "BYTE_NUMBER", "EXT_ASCII_CHAR",
		"WS", "BYTE", "INT", "PRINT", "IF", "ELSE", "WHILE", "READ", "PLUS",
		"MINUS", "TIMES", "DIVIDE", "EQUALS", "DEQUALS", "AND", "OR", "NOT",
		"GT", "GE", "LT", "LE", "NUMBER", "IDENTIFIER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 166, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 4, 7, 80, 8, 7, 11, 7, 12, 7, 81, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 145, 8, 23, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 4, 28, 158, 8,
		28, 11, 28, 12, 28, 159, 1, 29, 4, 29, 163, 8, 29, 11, 29, 12, 29, 164,
		0, 0, 30, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 0, 15, 7, 17, 8, 19,
		9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 17, 37,
		18, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55,
		27, 57, 28, 59, 29, 1, 0, 4, 1, 0, 0, 255, 3, 0, 9, 10, 13, 13, 32, 32,
		1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 168, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0,
		0, 0, 1, 61, 1, 0, 0, 0, 3, 63, 1, 0, 0, 0, 5, 65, 1, 0, 0, 0, 7, 67, 1,
		0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 73, 1, 0, 0, 0, 13, 76, 1, 0, 0, 0, 15,
		79, 1, 0, 0, 0, 17, 85, 1, 0, 0, 0, 19, 90, 1, 0, 0, 0, 21, 94, 1, 0, 0,
		0, 23, 100, 1, 0, 0, 0, 25, 103, 1, 0, 0, 0, 27, 108, 1, 0, 0, 0, 29, 114,
		1, 0, 0, 0, 31, 119, 1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 123, 1, 0, 0,
		0, 37, 125, 1, 0, 0, 0, 39, 127, 1, 0, 0, 0, 41, 129, 1, 0, 0, 0, 43, 132,
		1, 0, 0, 0, 45, 136, 1, 0, 0, 0, 47, 144, 1, 0, 0, 0, 49, 146, 1, 0, 0,
		0, 51, 148, 1, 0, 0, 0, 53, 151, 1, 0, 0, 0, 55, 153, 1, 0, 0, 0, 57, 157,
		1, 0, 0, 0, 59, 162, 1, 0, 0, 0, 61, 62, 5, 40, 0, 0, 62, 2, 1, 0, 0, 0,
		63, 64, 5, 41, 0, 0, 64, 4, 1, 0, 0, 0, 65, 66, 5, 123, 0, 0, 66, 6, 1,
		0, 0, 0, 67, 68, 5, 125, 0, 0, 68, 8, 1, 0, 0, 0, 69, 70, 5, 39, 0, 0,
		70, 71, 3, 13, 6, 0, 71, 72, 5, 39, 0, 0, 72, 10, 1, 0, 0, 0, 73, 74, 3,
		57, 28, 0, 74, 75, 5, 98, 0, 0, 75, 12, 1, 0, 0, 0, 76, 77, 7, 0, 0, 0,
		77, 14, 1, 0, 0, 0, 78, 80, 7, 1, 0, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83,
		84, 6, 7, 0, 0, 84, 16, 1, 0, 0, 0, 85, 86, 5, 98, 0, 0, 86, 87, 5, 121,
		0, 0, 87, 88, 5, 116, 0, 0, 88, 89, 5, 101, 0, 0, 89, 18, 1, 0, 0, 0, 90,
		91, 5, 105, 0, 0, 91, 92, 5, 110, 0, 0, 92, 93, 5, 116, 0, 0, 93, 20, 1,
		0, 0, 0, 94, 95, 5, 112, 0, 0, 95, 96, 5, 114, 0, 0, 96, 97, 5, 105, 0,
		0, 97, 98, 5, 110, 0, 0, 98, 99, 5, 116, 0, 0, 99, 22, 1, 0, 0, 0, 100,
		101, 5, 105, 0, 0, 101, 102, 5, 102, 0, 0, 102, 24, 1, 0, 0, 0, 103, 104,
		5, 101, 0, 0, 104, 105, 5, 108, 0, 0, 105, 106, 5, 115, 0, 0, 106, 107,
		5, 101, 0, 0, 107, 26, 1, 0, 0, 0, 108, 109, 5, 119, 0, 0, 109, 110, 5,
		104, 0, 0, 110, 111, 5, 105, 0, 0, 111, 112, 5, 108, 0, 0, 112, 113, 5,
		101, 0, 0, 113, 28, 1, 0, 0, 0, 114, 115, 5, 114, 0, 0, 115, 116, 5, 101,
		0, 0, 116, 117, 5, 97, 0, 0, 117, 118, 5, 100, 0, 0, 118, 30, 1, 0, 0,
		0, 119, 120, 5, 43, 0, 0, 120, 32, 1, 0, 0, 0, 121, 122, 5, 45, 0, 0, 122,
		34, 1, 0, 0, 0, 123, 124, 5, 42, 0, 0, 124, 36, 1, 0, 0, 0, 125, 126, 5,
		47, 0, 0, 126, 38, 1, 0, 0, 0, 127, 128, 5, 61, 0, 0, 128, 40, 1, 0, 0,
		0, 129, 130, 5, 61, 0, 0, 130, 131, 5, 61, 0, 0, 131, 42, 1, 0, 0, 0, 132,
		133, 5, 97, 0, 0, 133, 134, 5, 110, 0, 0, 134, 135, 5, 100, 0, 0, 135,
		44, 1, 0, 0, 0, 136, 137, 5, 111, 0, 0, 137, 138, 5, 114, 0, 0, 138, 46,
		1, 0, 0, 0, 139, 140, 5, 110, 0, 0, 140, 141, 5, 111, 0, 0, 141, 142, 5,
		116, 0, 0, 142, 145, 5, 32, 0, 0, 143, 145, 5, 33, 0, 0, 144, 139, 1, 0,
		0, 0, 144, 143, 1, 0, 0, 0, 145, 48, 1, 0, 0, 0, 146, 147, 5, 62, 0, 0,
		147, 50, 1, 0, 0, 0, 148, 149, 5, 62, 0, 0, 149, 150, 5, 61, 0, 0, 150,
		52, 1, 0, 0, 0, 151, 152, 5, 60, 0, 0, 152, 54, 1, 0, 0, 0, 153, 154, 5,
		60, 0, 0, 154, 155, 5, 61, 0, 0, 155, 56, 1, 0, 0, 0, 156, 158, 7, 2, 0,
		0, 157, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 58, 1, 0, 0, 0, 161, 163, 7, 3, 0, 0, 162, 161, 1,
		0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0,
		0, 165, 60, 1, 0, 0, 0, 5, 0, 81, 144, 159, 164, 1, 0, 1, 0,
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
	mindfckLexerT__0        = 1
	mindfckLexerT__1        = 2
	mindfckLexerT__2        = 3
	mindfckLexerT__3        = 4
	mindfckLexerCHAR        = 5
	mindfckLexerBYTE_NUMBER = 6
	mindfckLexerWS          = 7
	mindfckLexerBYTE        = 8
	mindfckLexerINT         = 9
	mindfckLexerPRINT       = 10
	mindfckLexerIF          = 11
	mindfckLexerELSE        = 12
	mindfckLexerWHILE       = 13
	mindfckLexerREAD        = 14
	mindfckLexerPLUS        = 15
	mindfckLexerMINUS       = 16
	mindfckLexerTIMES       = 17
	mindfckLexerDIVIDE      = 18
	mindfckLexerEQUALS      = 19
	mindfckLexerDEQUALS     = 20
	mindfckLexerAND         = 21
	mindfckLexerOR          = 22
	mindfckLexerNOT         = 23
	mindfckLexerGT          = 24
	mindfckLexerGE          = 25
	mindfckLexerLT          = 26
	mindfckLexerLE          = 27
	mindfckLexerNUMBER      = 28
	mindfckLexerIDENTIFIER  = 29
)
