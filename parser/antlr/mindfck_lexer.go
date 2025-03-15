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
		"", "'['", "']'", "'('", "')'", "'{'", "'}'", "','", "", "", "", "'byte'",
		"'debug'", "'int'", "'print'", "'if'", "'else'", "'while'", "'read'",
		"'+'", "'-'", "'*'", "'/'", "'='", "'=='", "'and'", "'or'", "", "'>'",
		"'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "CHAR", "BYTE_NUMBER", "WS", "BYTE",
		"DEBUG", "INT", "PRINT", "IF", "ELSE", "WHILE", "READ", "PLUS", "MINUS",
		"TIMES", "DIVIDE", "EQUALS", "DEQUALS", "AND", "OR", "NOT", "GT", "GE",
		"LT", "LE", "NUMBER", "IDENTIFIER",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "CHAR", "BYTE_NUMBER",
		"EXT_ASCII_CHAR", "WS", "BYTE", "DEBUG", "INT", "PRINT", "IF", "ELSE",
		"WHILE", "READ", "PLUS", "MINUS", "TIMES", "DIVIDE", "EQUALS", "DEQUALS",
		"AND", "OR", "NOT", "GT", "GE", "LT", "LE", "NUMBER", "IDENTIFIER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 186, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 94, 8, 10, 11, 10, 12, 10,
		95, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 165, 8, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 4, 32, 178, 8,
		32, 11, 32, 12, 32, 179, 1, 33, 4, 33, 183, 8, 33, 11, 33, 12, 33, 184,
		0, 0, 34, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		0, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 17, 37,
		18, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55,
		27, 57, 28, 59, 29, 61, 30, 63, 31, 65, 32, 67, 33, 1, 0, 4, 1, 0, 0, 255,
		3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 188,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3, 71, 1, 0,
		0, 0, 5, 73, 1, 0, 0, 0, 7, 75, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 79,
		1, 0, 0, 0, 13, 81, 1, 0, 0, 0, 15, 83, 1, 0, 0, 0, 17, 87, 1, 0, 0, 0,
		19, 90, 1, 0, 0, 0, 21, 93, 1, 0, 0, 0, 23, 99, 1, 0, 0, 0, 25, 104, 1,
		0, 0, 0, 27, 110, 1, 0, 0, 0, 29, 114, 1, 0, 0, 0, 31, 120, 1, 0, 0, 0,
		33, 123, 1, 0, 0, 0, 35, 128, 1, 0, 0, 0, 37, 134, 1, 0, 0, 0, 39, 139,
		1, 0, 0, 0, 41, 141, 1, 0, 0, 0, 43, 143, 1, 0, 0, 0, 45, 145, 1, 0, 0,
		0, 47, 147, 1, 0, 0, 0, 49, 149, 1, 0, 0, 0, 51, 152, 1, 0, 0, 0, 53, 156,
		1, 0, 0, 0, 55, 164, 1, 0, 0, 0, 57, 166, 1, 0, 0, 0, 59, 168, 1, 0, 0,
		0, 61, 171, 1, 0, 0, 0, 63, 173, 1, 0, 0, 0, 65, 177, 1, 0, 0, 0, 67, 182,
		1, 0, 0, 0, 69, 70, 5, 91, 0, 0, 70, 2, 1, 0, 0, 0, 71, 72, 5, 93, 0, 0,
		72, 4, 1, 0, 0, 0, 73, 74, 5, 40, 0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 41,
		0, 0, 76, 8, 1, 0, 0, 0, 77, 78, 5, 123, 0, 0, 78, 10, 1, 0, 0, 0, 79,
		80, 5, 125, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5, 44, 0, 0, 82, 14, 1, 0,
		0, 0, 83, 84, 5, 39, 0, 0, 84, 85, 3, 19, 9, 0, 85, 86, 5, 39, 0, 0, 86,
		16, 1, 0, 0, 0, 87, 88, 3, 65, 32, 0, 88, 89, 5, 98, 0, 0, 89, 18, 1, 0,
		0, 0, 90, 91, 7, 0, 0, 0, 91, 20, 1, 0, 0, 0, 92, 94, 7, 1, 0, 0, 93, 92,
		1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 97, 1, 0, 0, 0, 97, 98, 6, 10, 0, 0, 98, 22, 1, 0, 0, 0, 99, 100, 5,
		98, 0, 0, 100, 101, 5, 121, 0, 0, 101, 102, 5, 116, 0, 0, 102, 103, 5,
		101, 0, 0, 103, 24, 1, 0, 0, 0, 104, 105, 5, 100, 0, 0, 105, 106, 5, 101,
		0, 0, 106, 107, 5, 98, 0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 103,
		0, 0, 109, 26, 1, 0, 0, 0, 110, 111, 5, 105, 0, 0, 111, 112, 5, 110, 0,
		0, 112, 113, 5, 116, 0, 0, 113, 28, 1, 0, 0, 0, 114, 115, 5, 112, 0, 0,
		115, 116, 5, 114, 0, 0, 116, 117, 5, 105, 0, 0, 117, 118, 5, 110, 0, 0,
		118, 119, 5, 116, 0, 0, 119, 30, 1, 0, 0, 0, 120, 121, 5, 105, 0, 0, 121,
		122, 5, 102, 0, 0, 122, 32, 1, 0, 0, 0, 123, 124, 5, 101, 0, 0, 124, 125,
		5, 108, 0, 0, 125, 126, 5, 115, 0, 0, 126, 127, 5, 101, 0, 0, 127, 34,
		1, 0, 0, 0, 128, 129, 5, 119, 0, 0, 129, 130, 5, 104, 0, 0, 130, 131, 5,
		105, 0, 0, 131, 132, 5, 108, 0, 0, 132, 133, 5, 101, 0, 0, 133, 36, 1,
		0, 0, 0, 134, 135, 5, 114, 0, 0, 135, 136, 5, 101, 0, 0, 136, 137, 5, 97,
		0, 0, 137, 138, 5, 100, 0, 0, 138, 38, 1, 0, 0, 0, 139, 140, 5, 43, 0,
		0, 140, 40, 1, 0, 0, 0, 141, 142, 5, 45, 0, 0, 142, 42, 1, 0, 0, 0, 143,
		144, 5, 42, 0, 0, 144, 44, 1, 0, 0, 0, 145, 146, 5, 47, 0, 0, 146, 46,
		1, 0, 0, 0, 147, 148, 5, 61, 0, 0, 148, 48, 1, 0, 0, 0, 149, 150, 5, 61,
		0, 0, 150, 151, 5, 61, 0, 0, 151, 50, 1, 0, 0, 0, 152, 153, 5, 97, 0, 0,
		153, 154, 5, 110, 0, 0, 154, 155, 5, 100, 0, 0, 155, 52, 1, 0, 0, 0, 156,
		157, 5, 111, 0, 0, 157, 158, 5, 114, 0, 0, 158, 54, 1, 0, 0, 0, 159, 160,
		5, 110, 0, 0, 160, 161, 5, 111, 0, 0, 161, 162, 5, 116, 0, 0, 162, 165,
		5, 32, 0, 0, 163, 165, 5, 33, 0, 0, 164, 159, 1, 0, 0, 0, 164, 163, 1,
		0, 0, 0, 165, 56, 1, 0, 0, 0, 166, 167, 5, 62, 0, 0, 167, 58, 1, 0, 0,
		0, 168, 169, 5, 62, 0, 0, 169, 170, 5, 61, 0, 0, 170, 60, 1, 0, 0, 0, 171,
		172, 5, 60, 0, 0, 172, 62, 1, 0, 0, 0, 173, 174, 5, 60, 0, 0, 174, 175,
		5, 61, 0, 0, 175, 64, 1, 0, 0, 0, 176, 178, 7, 2, 0, 0, 177, 176, 1, 0,
		0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0,
		180, 66, 1, 0, 0, 0, 181, 183, 7, 3, 0, 0, 182, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 68, 1, 0,
		0, 0, 5, 0, 95, 164, 179, 184, 1, 0, 1, 0,
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
	mindfckLexerT__4        = 5
	mindfckLexerT__5        = 6
	mindfckLexerT__6        = 7
	mindfckLexerCHAR        = 8
	mindfckLexerBYTE_NUMBER = 9
	mindfckLexerWS          = 10
	mindfckLexerBYTE        = 11
	mindfckLexerDEBUG       = 12
	mindfckLexerINT         = 13
	mindfckLexerPRINT       = 14
	mindfckLexerIF          = 15
	mindfckLexerELSE        = 16
	mindfckLexerWHILE       = 17
	mindfckLexerREAD        = 18
	mindfckLexerPLUS        = 19
	mindfckLexerMINUS       = 20
	mindfckLexerTIMES       = 21
	mindfckLexerDIVIDE      = 22
	mindfckLexerEQUALS      = 23
	mindfckLexerDEQUALS     = 24
	mindfckLexerAND         = 25
	mindfckLexerOR          = 26
	mindfckLexerNOT         = 27
	mindfckLexerGT          = 28
	mindfckLexerGE          = 29
	mindfckLexerLT          = 30
	mindfckLexerLE          = 31
	mindfckLexerNUMBER      = 32
	mindfckLexerIDENTIFIER  = 33
)
