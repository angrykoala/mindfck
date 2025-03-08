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
		"'int'", "'print'", "'if'", "'else'", "'while'", "'read'", "'+'", "'-'",
		"'*'", "'/'", "'='", "'=='", "'and'", "'or'", "", "'>'", "'>='", "'<'",
		"'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "CHAR", "BYTE_NUMBER", "WS", "BYTE",
		"INT", "PRINT", "IF", "ELSE", "WHILE", "READ", "PLUS", "MINUS", "TIMES",
		"DIVIDE", "EQUALS", "DEQUALS", "AND", "OR", "NOT", "GT", "GE", "LT",
		"LE", "NUMBER", "IDENTIFIER",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "CHAR", "BYTE_NUMBER",
		"EXT_ASCII_CHAR", "WS", "BYTE", "INT", "PRINT", "IF", "ELSE", "WHILE",
		"READ", "PLUS", "MINUS", "TIMES", "DIVIDE", "EQUALS", "DEQUALS", "AND",
		"OR", "NOT", "GT", "GE", "LT", "LE", "NUMBER", "IDENTIFIER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 178, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 92, 8, 10, 11, 10, 12, 10, 93, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 157, 8, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		4, 31, 170, 8, 31, 11, 31, 12, 31, 171, 1, 32, 4, 32, 175, 8, 32, 11, 32,
		12, 32, 176, 0, 0, 33, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 0, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16,
		35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25,
		53, 26, 55, 27, 57, 28, 59, 29, 61, 30, 63, 31, 65, 32, 1, 0, 4, 1, 0,
		0, 255, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122,
		180, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 1, 67, 1, 0, 0, 0, 3, 69, 1, 0, 0, 0, 5, 71,
		1, 0, 0, 0, 7, 73, 1, 0, 0, 0, 9, 75, 1, 0, 0, 0, 11, 77, 1, 0, 0, 0, 13,
		79, 1, 0, 0, 0, 15, 81, 1, 0, 0, 0, 17, 85, 1, 0, 0, 0, 19, 88, 1, 0, 0,
		0, 21, 91, 1, 0, 0, 0, 23, 97, 1, 0, 0, 0, 25, 102, 1, 0, 0, 0, 27, 106,
		1, 0, 0, 0, 29, 112, 1, 0, 0, 0, 31, 115, 1, 0, 0, 0, 33, 120, 1, 0, 0,
		0, 35, 126, 1, 0, 0, 0, 37, 131, 1, 0, 0, 0, 39, 133, 1, 0, 0, 0, 41, 135,
		1, 0, 0, 0, 43, 137, 1, 0, 0, 0, 45, 139, 1, 0, 0, 0, 47, 141, 1, 0, 0,
		0, 49, 144, 1, 0, 0, 0, 51, 148, 1, 0, 0, 0, 53, 156, 1, 0, 0, 0, 55, 158,
		1, 0, 0, 0, 57, 160, 1, 0, 0, 0, 59, 163, 1, 0, 0, 0, 61, 165, 1, 0, 0,
		0, 63, 169, 1, 0, 0, 0, 65, 174, 1, 0, 0, 0, 67, 68, 5, 91, 0, 0, 68, 2,
		1, 0, 0, 0, 69, 70, 5, 93, 0, 0, 70, 4, 1, 0, 0, 0, 71, 72, 5, 40, 0, 0,
		72, 6, 1, 0, 0, 0, 73, 74, 5, 41, 0, 0, 74, 8, 1, 0, 0, 0, 75, 76, 5, 123,
		0, 0, 76, 10, 1, 0, 0, 0, 77, 78, 5, 125, 0, 0, 78, 12, 1, 0, 0, 0, 79,
		80, 5, 44, 0, 0, 80, 14, 1, 0, 0, 0, 81, 82, 5, 39, 0, 0, 82, 83, 3, 19,
		9, 0, 83, 84, 5, 39, 0, 0, 84, 16, 1, 0, 0, 0, 85, 86, 3, 63, 31, 0, 86,
		87, 5, 98, 0, 0, 87, 18, 1, 0, 0, 0, 88, 89, 7, 0, 0, 0, 89, 20, 1, 0,
		0, 0, 90, 92, 7, 1, 0, 0, 91, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 91,
		1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 6, 10, 0, 0,
		96, 22, 1, 0, 0, 0, 97, 98, 5, 98, 0, 0, 98, 99, 5, 121, 0, 0, 99, 100,
		5, 116, 0, 0, 100, 101, 5, 101, 0, 0, 101, 24, 1, 0, 0, 0, 102, 103, 5,
		105, 0, 0, 103, 104, 5, 110, 0, 0, 104, 105, 5, 116, 0, 0, 105, 26, 1,
		0, 0, 0, 106, 107, 5, 112, 0, 0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 105,
		0, 0, 109, 110, 5, 110, 0, 0, 110, 111, 5, 116, 0, 0, 111, 28, 1, 0, 0,
		0, 112, 113, 5, 105, 0, 0, 113, 114, 5, 102, 0, 0, 114, 30, 1, 0, 0, 0,
		115, 116, 5, 101, 0, 0, 116, 117, 5, 108, 0, 0, 117, 118, 5, 115, 0, 0,
		118, 119, 5, 101, 0, 0, 119, 32, 1, 0, 0, 0, 120, 121, 5, 119, 0, 0, 121,
		122, 5, 104, 0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5, 108, 0, 0, 124,
		125, 5, 101, 0, 0, 125, 34, 1, 0, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128,
		5, 101, 0, 0, 128, 129, 5, 97, 0, 0, 129, 130, 5, 100, 0, 0, 130, 36, 1,
		0, 0, 0, 131, 132, 5, 43, 0, 0, 132, 38, 1, 0, 0, 0, 133, 134, 5, 45, 0,
		0, 134, 40, 1, 0, 0, 0, 135, 136, 5, 42, 0, 0, 136, 42, 1, 0, 0, 0, 137,
		138, 5, 47, 0, 0, 138, 44, 1, 0, 0, 0, 139, 140, 5, 61, 0, 0, 140, 46,
		1, 0, 0, 0, 141, 142, 5, 61, 0, 0, 142, 143, 5, 61, 0, 0, 143, 48, 1, 0,
		0, 0, 144, 145, 5, 97, 0, 0, 145, 146, 5, 110, 0, 0, 146, 147, 5, 100,
		0, 0, 147, 50, 1, 0, 0, 0, 148, 149, 5, 111, 0, 0, 149, 150, 5, 114, 0,
		0, 150, 52, 1, 0, 0, 0, 151, 152, 5, 110, 0, 0, 152, 153, 5, 111, 0, 0,
		153, 154, 5, 116, 0, 0, 154, 157, 5, 32, 0, 0, 155, 157, 5, 33, 0, 0, 156,
		151, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 54, 1, 0, 0, 0, 158, 159, 5,
		62, 0, 0, 159, 56, 1, 0, 0, 0, 160, 161, 5, 62, 0, 0, 161, 162, 5, 61,
		0, 0, 162, 58, 1, 0, 0, 0, 163, 164, 5, 60, 0, 0, 164, 60, 1, 0, 0, 0,
		165, 166, 5, 60, 0, 0, 166, 167, 5, 61, 0, 0, 167, 62, 1, 0, 0, 0, 168,
		170, 7, 2, 0, 0, 169, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 169,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 64, 1, 0, 0, 0, 173, 175, 7, 3,
		0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0,
		176, 177, 1, 0, 0, 0, 177, 66, 1, 0, 0, 0, 5, 0, 93, 156, 171, 176, 1,
		0, 1, 0,
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
	mindfckLexerINT         = 12
	mindfckLexerPRINT       = 13
	mindfckLexerIF          = 14
	mindfckLexerELSE        = 15
	mindfckLexerWHILE       = 16
	mindfckLexerREAD        = 17
	mindfckLexerPLUS        = 18
	mindfckLexerMINUS       = 19
	mindfckLexerTIMES       = 20
	mindfckLexerDIVIDE      = 21
	mindfckLexerEQUALS      = 22
	mindfckLexerDEQUALS     = 23
	mindfckLexerAND         = 24
	mindfckLexerOR          = 25
	mindfckLexerNOT         = 26
	mindfckLexerGT          = 27
	mindfckLexerGE          = 28
	mindfckLexerLT          = 29
	mindfckLexerLE          = 30
	mindfckLexerNUMBER      = 31
	mindfckLexerIDENTIFIER  = 32
)
