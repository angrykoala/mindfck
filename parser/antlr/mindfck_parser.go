// Code generated from parser/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindfck

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

type mindfckParser struct {
	*antlr.BaseParser
}

var MindfckParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mindfckParserInit() {
	staticData := &MindfckParserStaticData
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
		"statements", "statement", "declaration", "arrayDeclaration", "assignment",
		"print", "ifConditional", "whileLoop", "read", "block", "expression",
		"identifier", "literal", "arrayLiteral", "arrayItem", "arrayIndex",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 168, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 46, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 52, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 61, 8, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		88, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 5, 9, 102, 8, 9, 10, 9, 12, 9, 105, 9, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 116, 8, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 138, 8,
		10, 10, 10, 12, 10, 141, 9, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 149, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 155, 8, 13, 10, 13,
		12, 13, 158, 9, 13, 3, 13, 160, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 0, 1, 20, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 0, 8, 1, 0, 11, 12, 1, 0, 11, 11, 1, 0, 20, 21, 1, 0, 18,
		19, 2, 0, 22, 23, 27, 30, 1, 0, 24, 24, 1, 0, 25, 25, 2, 0, 9, 9, 31, 31,
		177, 0, 35, 1, 0, 0, 0, 2, 45, 1, 0, 0, 0, 4, 47, 1, 0, 0, 0, 6, 53, 1,
		0, 0, 0, 8, 62, 1, 0, 0, 0, 10, 72, 1, 0, 0, 0, 12, 75, 1, 0, 0, 0, 14,
		89, 1, 0, 0, 0, 16, 97, 1, 0, 0, 0, 18, 103, 1, 0, 0, 0, 20, 115, 1, 0,
		0, 0, 22, 142, 1, 0, 0, 0, 24, 148, 1, 0, 0, 0, 26, 150, 1, 0, 0, 0, 28,
		163, 1, 0, 0, 0, 30, 165, 1, 0, 0, 0, 32, 34, 3, 2, 1, 0, 33, 32, 1, 0,
		0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 1,
		1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 46, 3, 4, 2, 0, 39, 46, 3, 6, 3, 0,
		40, 46, 3, 8, 4, 0, 41, 46, 3, 10, 5, 0, 42, 46, 3, 12, 6, 0, 43, 46, 3,
		14, 7, 0, 44, 46, 3, 16, 8, 0, 45, 38, 1, 0, 0, 0, 45, 39, 1, 0, 0, 0,
		45, 40, 1, 0, 0, 0, 45, 41, 1, 0, 0, 0, 45, 42, 1, 0, 0, 0, 45, 43, 1,
		0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 48, 7, 0, 0, 0, 48,
		51, 3, 22, 11, 0, 49, 50, 5, 22, 0, 0, 50, 52, 3, 20, 10, 0, 51, 49, 1,
		0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 5, 1, 0, 0, 0, 53, 54, 7, 1, 0, 0, 54,
		55, 3, 22, 11, 0, 55, 56, 5, 1, 0, 0, 56, 57, 3, 30, 15, 0, 57, 60, 5,
		2, 0, 0, 58, 59, 5, 22, 0, 0, 59, 61, 3, 20, 10, 0, 60, 58, 1, 0, 0, 0,
		60, 61, 1, 0, 0, 0, 61, 7, 1, 0, 0, 0, 62, 67, 3, 22, 11, 0, 63, 64, 5,
		1, 0, 0, 64, 65, 3, 30, 15, 0, 65, 66, 5, 2, 0, 0, 66, 68, 1, 0, 0, 0,
		67, 63, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 5,
		22, 0, 0, 70, 71, 3, 20, 10, 0, 71, 9, 1, 0, 0, 0, 72, 73, 5, 13, 0, 0,
		73, 74, 3, 20, 10, 0, 74, 11, 1, 0, 0, 0, 75, 76, 5, 14, 0, 0, 76, 77,
		5, 3, 0, 0, 77, 78, 3, 20, 10, 0, 78, 79, 5, 4, 0, 0, 79, 80, 5, 5, 0,
		0, 80, 81, 3, 18, 9, 0, 81, 87, 5, 6, 0, 0, 82, 83, 5, 15, 0, 0, 83, 84,
		5, 5, 0, 0, 84, 85, 3, 18, 9, 0, 85, 86, 5, 6, 0, 0, 86, 88, 1, 0, 0, 0,
		87, 82, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 13, 1, 0, 0, 0, 89, 90, 5,
		16, 0, 0, 90, 91, 5, 3, 0, 0, 91, 92, 3, 20, 10, 0, 92, 93, 5, 4, 0, 0,
		93, 94, 5, 5, 0, 0, 94, 95, 3, 18, 9, 0, 95, 96, 5, 6, 0, 0, 96, 15, 1,
		0, 0, 0, 97, 98, 5, 17, 0, 0, 98, 99, 3, 22, 11, 0, 99, 17, 1, 0, 0, 0,
		100, 102, 3, 2, 1, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 19, 1, 0, 0, 0, 105, 103, 1,
		0, 0, 0, 106, 107, 6, 10, -1, 0, 107, 116, 3, 22, 11, 0, 108, 116, 3, 24,
		12, 0, 109, 110, 5, 3, 0, 0, 110, 111, 3, 20, 10, 0, 111, 112, 5, 4, 0,
		0, 112, 116, 1, 0, 0, 0, 113, 114, 5, 26, 0, 0, 114, 116, 3, 20, 10, 6,
		115, 106, 1, 0, 0, 0, 115, 108, 1, 0, 0, 0, 115, 109, 1, 0, 0, 0, 115,
		113, 1, 0, 0, 0, 116, 139, 1, 0, 0, 0, 117, 118, 10, 5, 0, 0, 118, 119,
		7, 2, 0, 0, 119, 138, 3, 20, 10, 6, 120, 121, 10, 4, 0, 0, 121, 122, 7,
		3, 0, 0, 122, 138, 3, 20, 10, 5, 123, 124, 10, 3, 0, 0, 124, 125, 7, 4,
		0, 0, 125, 138, 3, 20, 10, 4, 126, 127, 10, 2, 0, 0, 127, 128, 7, 5, 0,
		0, 128, 138, 3, 20, 10, 3, 129, 130, 10, 1, 0, 0, 130, 131, 7, 6, 0, 0,
		131, 138, 3, 20, 10, 2, 132, 133, 10, 8, 0, 0, 133, 134, 5, 1, 0, 0, 134,
		135, 3, 30, 15, 0, 135, 136, 5, 2, 0, 0, 136, 138, 1, 0, 0, 0, 137, 117,
		1, 0, 0, 0, 137, 120, 1, 0, 0, 0, 137, 123, 1, 0, 0, 0, 137, 126, 1, 0,
		0, 0, 137, 129, 1, 0, 0, 0, 137, 132, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0,
		139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 21, 1, 0, 0, 0, 141, 139,
		1, 0, 0, 0, 142, 143, 5, 32, 0, 0, 143, 23, 1, 0, 0, 0, 144, 149, 5, 31,
		0, 0, 145, 149, 5, 8, 0, 0, 146, 149, 5, 9, 0, 0, 147, 149, 3, 26, 13,
		0, 148, 144, 1, 0, 0, 0, 148, 145, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148,
		147, 1, 0, 0, 0, 149, 25, 1, 0, 0, 0, 150, 159, 5, 1, 0, 0, 151, 156, 3,
		28, 14, 0, 152, 153, 5, 7, 0, 0, 153, 155, 3, 28, 14, 0, 154, 152, 1, 0,
		0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0,
		157, 160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 151, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 2, 0, 0, 162, 27, 1,
		0, 0, 0, 163, 164, 7, 7, 0, 0, 164, 29, 1, 0, 0, 0, 165, 166, 5, 31, 0,
		0, 166, 31, 1, 0, 0, 0, 13, 35, 45, 51, 60, 67, 87, 103, 115, 137, 139,
		148, 156, 159,
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

// mindfckParserInit initializes any static state used to implement mindfckParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewmindfckParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MindfckParserInit() {
	staticData := &MindfckParserStaticData
	staticData.once.Do(mindfckParserInit)
}

// NewmindfckParser produces a new parser instance for the optional input antlr.TokenStream.
func NewmindfckParser(input antlr.TokenStream) *mindfckParser {
	MindfckParserInit()
	this := new(mindfckParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MindfckParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "mindfck.g4"

	return this
}

// mindfckParser tokens.
const (
	mindfckParserEOF         = antlr.TokenEOF
	mindfckParserT__0        = 1
	mindfckParserT__1        = 2
	mindfckParserT__2        = 3
	mindfckParserT__3        = 4
	mindfckParserT__4        = 5
	mindfckParserT__5        = 6
	mindfckParserT__6        = 7
	mindfckParserCHAR        = 8
	mindfckParserBYTE_NUMBER = 9
	mindfckParserWS          = 10
	mindfckParserBYTE        = 11
	mindfckParserINT         = 12
	mindfckParserPRINT       = 13
	mindfckParserIF          = 14
	mindfckParserELSE        = 15
	mindfckParserWHILE       = 16
	mindfckParserREAD        = 17
	mindfckParserPLUS        = 18
	mindfckParserMINUS       = 19
	mindfckParserTIMES       = 20
	mindfckParserDIVIDE      = 21
	mindfckParserEQUALS      = 22
	mindfckParserDEQUALS     = 23
	mindfckParserAND         = 24
	mindfckParserOR          = 25
	mindfckParserNOT         = 26
	mindfckParserGT          = 27
	mindfckParserGE          = 28
	mindfckParserLT          = 29
	mindfckParserLE          = 30
	mindfckParserNUMBER      = 31
	mindfckParserIDENTIFIER  = 32
)

// mindfckParser rules.
const (
	mindfckParserRULE_statements       = 0
	mindfckParserRULE_statement        = 1
	mindfckParserRULE_declaration      = 2
	mindfckParserRULE_arrayDeclaration = 3
	mindfckParserRULE_assignment       = 4
	mindfckParserRULE_print            = 5
	mindfckParserRULE_ifConditional    = 6
	mindfckParserRULE_whileLoop        = 7
	mindfckParserRULE_read             = 8
	mindfckParserRULE_block            = 9
	mindfckParserRULE_expression       = 10
	mindfckParserRULE_identifier       = 11
	mindfckParserRULE_literal          = 12
	mindfckParserRULE_arrayLiteral     = 13
	mindfckParserRULE_arrayItem        = 14
	mindfckParserRULE_arrayIndex       = 15
)

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
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

func (s *StatementsContext) Statement(i int) IStatementContext {
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

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mindfckParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4295194624) != 0 {
		{
			p.SetState(32)
			p.Statement()
		}

		p.SetState(37)
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

	// Getter signatures
	Declaration() IDeclarationContext
	ArrayDeclaration() IArrayDeclarationContext
	Assignment() IAssignmentContext
	Print_() IPrintContext
	IfConditional() IIfConditionalContext
	WhileLoop() IWhileLoopContext
	Read() IReadContext

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
	p.RuleIndex = mindfckParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) ArrayDeclaration() IArrayDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) IfConditional() IIfConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfConditionalContext)
}

func (s *StatementContext) WhileLoop() IWhileLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *StatementContext) Read() IReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, mindfckParserRULE_statement)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.ArrayDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Assignment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Print_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(42)
			p.IfConditional()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(43)
			p.WhileLoop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(44)
			p.Read()
		}

	case antlr.ATNInvalidAltNumber:
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType_ returns the type_ token.
	GetType_() antlr.Token

	// SetType_ sets the type_ token.
	SetType_(antlr.Token)

	// Getter signatures
	Identifier() IIdentifierContext
	BYTE() antlr.TerminalNode
	INT() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expression() IExpressionContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	type_  antlr.Token
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetType_() antlr.Token { return s.type_ }

func (s *DeclarationContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *DeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DeclarationContext) BYTE() antlr.TerminalNode {
	return s.GetToken(mindfckParserBYTE, 0)
}

func (s *DeclarationContext) INT() antlr.TerminalNode {
	return s.GetToken(mindfckParserINT, 0)
}

func (s *DeclarationContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(mindfckParserEQUALS, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
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

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mindfckParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*DeclarationContext).type_ = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == mindfckParserBYTE || _la == mindfckParserINT) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*DeclarationContext).type_ = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(48)
		p.Identifier()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindfckParserEQUALS {
		{
			p.SetState(49)
			p.Match(mindfckParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.expression(0)
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayDeclarationContext is an interface to support dynamic dispatch.
type IArrayDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType_ returns the type_ token.
	GetType_() antlr.Token

	// SetType_ sets the type_ token.
	SetType_(antlr.Token)

	// Getter signatures
	Identifier() IIdentifierContext
	ArrayIndex() IArrayIndexContext
	BYTE() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expression() IExpressionContext

	// IsArrayDeclarationContext differentiates from other interfaces.
	IsArrayDeclarationContext()
}

type ArrayDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	type_  antlr.Token
}

func NewEmptyArrayDeclarationContext() *ArrayDeclarationContext {
	var p = new(ArrayDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayDeclaration
	return p
}

func InitEmptyArrayDeclarationContext(p *ArrayDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayDeclaration
}

func (*ArrayDeclarationContext) IsArrayDeclarationContext() {}

func NewArrayDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclarationContext {
	var p = new(ArrayDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_arrayDeclaration

	return p
}

func (s *ArrayDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclarationContext) GetType_() antlr.Token { return s.type_ }

func (s *ArrayDeclarationContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *ArrayDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArrayDeclarationContext) ArrayIndex() IArrayIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
}

func (s *ArrayDeclarationContext) BYTE() antlr.TerminalNode {
	return s.GetToken(mindfckParserBYTE, 0)
}

func (s *ArrayDeclarationContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(mindfckParserEQUALS, 0)
}

func (s *ArrayDeclarationContext) Expression() IExpressionContext {
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

func (s *ArrayDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterArrayDeclaration(s)
	}
}

func (s *ArrayDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitArrayDeclaration(s)
	}
}

func (s *ArrayDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitArrayDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) ArrayDeclaration() (localctx IArrayDeclarationContext) {
	localctx = NewArrayDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mindfckParserRULE_arrayDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ArrayDeclarationContext).type_ = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == mindfckParserBYTE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ArrayDeclarationContext).type_ = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(54)
		p.Identifier()
	}
	{
		p.SetState(55)
		p.Match(mindfckParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.ArrayIndex()
	}
	{
		p.SetState(57)
		p.Match(mindfckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindfckParserEQUALS {
		{
			p.SetState(58)
			p.Match(mindfckParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.expression(0)
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	EQUALS() antlr.TerminalNode
	Expression() IExpressionContext
	ArrayIndex() IArrayIndexContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(mindfckParserEQUALS, 0)
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

func (s *AssignmentContext) ArrayIndex() IArrayIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mindfckParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Identifier()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindfckParserT__0 {
		{
			p.SetState(63)
			p.Match(mindfckParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.ArrayIndex()
		}
		{
			p.SetState(65)
			p.Match(mindfckParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(69)
		p.Match(mindfckParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.expression(0)
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

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(mindfckParserPRINT, 0)
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

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mindfckParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(mindfckParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.expression(0)
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

// IIfConditionalContext is an interface to support dynamic dispatch.
type IIfConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIfConditionalContext differentiates from other interfaces.
	IsIfConditionalContext()
}

type IfConditionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfConditionalContext() *IfConditionalContext {
	var p = new(IfConditionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_ifConditional
	return p
}

func InitEmptyIfConditionalContext(p *IfConditionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_ifConditional
}

func (*IfConditionalContext) IsIfConditionalContext() {}

func NewIfConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfConditionalContext {
	var p = new(IfConditionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_ifConditional

	return p
}

func (s *IfConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *IfConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(mindfckParserIF, 0)
}

func (s *IfConditionalContext) Expression() IExpressionContext {
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

func (s *IfConditionalContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfConditionalContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfConditionalContext) ELSE() antlr.TerminalNode {
	return s.GetToken(mindfckParserELSE, 0)
}

func (s *IfConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterIfConditional(s)
	}
}

func (s *IfConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitIfConditional(s)
	}
}

func (s *IfConditionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitIfConditional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) IfConditional() (localctx IIfConditionalContext) {
	localctx = NewIfConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, mindfckParserRULE_ifConditional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(mindfckParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(mindfckParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.expression(0)
	}
	{
		p.SetState(78)
		p.Match(mindfckParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(mindfckParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Block()
	}
	{
		p.SetState(81)
		p.Match(mindfckParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindfckParserELSE {
		{
			p.SetState(82)
			p.Match(mindfckParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(mindfckParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Block()
		}
		{
			p.SetState(85)
			p.Match(mindfckParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileLoopContext is an interface to support dynamic dispatch.
type IWhileLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsWhileLoopContext differentiates from other interfaces.
	IsWhileLoopContext()
}

type WhileLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileLoopContext() *WhileLoopContext {
	var p = new(WhileLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_whileLoop
	return p
}

func InitEmptyWhileLoopContext(p *WhileLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_whileLoop
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_whileLoop

	return p
}

func (s *WhileLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoopContext) WHILE() antlr.TerminalNode {
	return s.GetToken(mindfckParserWHILE, 0)
}

func (s *WhileLoopContext) Expression() IExpressionContext {
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

func (s *WhileLoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterWhileLoop(s)
	}
}

func (s *WhileLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitWhileLoop(s)
	}
}

func (s *WhileLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitWhileLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) WhileLoop() (localctx IWhileLoopContext) {
	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, mindfckParserRULE_whileLoop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(mindfckParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(mindfckParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.expression(0)
	}
	{
		p.SetState(92)
		p.Match(mindfckParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(mindfckParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Block()
	}
	{
		p.SetState(95)
		p.Match(mindfckParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_read
	return p
}

func InitEmptyReadContext(p *ReadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_read
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) READ() antlr.TerminalNode {
	return s.GetToken(mindfckParserREAD, 0)
}

func (s *ReadContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitRead(s)
	}
}

func (s *ReadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitRead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Read() (localctx IReadContext) {
	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, mindfckParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(mindfckParserREAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Identifier()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
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

func (s *BlockContext) Statement(i int) IStatementContext {
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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, mindfckParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4295194624) != 0 {
		{
			p.SetState(100)
			p.Statement()
		}

		p.SetState(105)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Identifier() IIdentifierContext
	Literal() ILiteralContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	NOT() antlr.TerminalNode
	TIMES() antlr.TerminalNode
	DIVIDE() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	DEQUALS() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	ArrayIndex() IArrayIndexContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(mindfckParserNOT, 0)
}

func (s *ExpressionContext) TIMES() antlr.TerminalNode {
	return s.GetToken(mindfckParserTIMES, 0)
}

func (s *ExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(mindfckParserDIVIDE, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(mindfckParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(mindfckParserMINUS, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(mindfckParserGT, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(mindfckParserGE, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(mindfckParserLT, 0)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(mindfckParserLE, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(mindfckParserEQUALS, 0)
}

func (s *ExpressionContext) DEQUALS() antlr.TerminalNode {
	return s.GetToken(mindfckParserDEQUALS, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(mindfckParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(mindfckParserOR, 0)
}

func (s *ExpressionContext) ArrayIndex() IArrayIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *mindfckParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, mindfckParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case mindfckParserIDENTIFIER:
		{
			p.SetState(107)
			p.Identifier()
		}

	case mindfckParserT__0, mindfckParserCHAR, mindfckParserBYTE_NUMBER, mindfckParserNUMBER:
		{
			p.SetState(108)
			p.Literal()
		}

	case mindfckParserT__2:
		{
			p.SetState(109)
			p.Match(mindfckParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.expression(0)
		}
		{
			p.SetState(111)
			p.Match(mindfckParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case mindfckParserNOT:
		{
			p.SetState(113)
			p.Match(mindfckParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.expression(6)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(118)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mindfckParserTIMES || _la == mindfckParserDIVIDE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(119)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(121)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mindfckParserPLUS || _la == mindfckParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(122)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(124)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2025848832) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(125)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(127)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mindfckParserAND) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(128)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(130)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == mindfckParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(131)
					p.expression(2)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(133)
					p.Match(mindfckParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(134)
					p.ArrayIndex()
				}
				{
					p.SetState(135)
					p.Match(mindfckParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(mindfckParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, mindfckParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(mindfckParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	BYTE_NUMBER() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(mindfckParserNUMBER, 0)
}

func (s *LiteralContext) CHAR() antlr.TerminalNode {
	return s.GetToken(mindfckParserCHAR, 0)
}

func (s *LiteralContext) BYTE_NUMBER() antlr.TerminalNode {
	return s.GetToken(mindfckParserBYTE_NUMBER, 0)
}

func (s *LiteralContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, mindfckParserRULE_literal)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case mindfckParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Match(mindfckParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case mindfckParserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(mindfckParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case mindfckParserBYTE_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Match(mindfckParserBYTE_NUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case mindfckParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
			p.ArrayLiteral()
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

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArrayItem() []IArrayItemContext
	ArrayItem(i int) IArrayItemContext

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllArrayItem() []IArrayItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayItemContext); ok {
			len++
		}
	}

	tst := make([]IArrayItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayItemContext); ok {
			tst[i] = t.(IArrayItemContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) ArrayItem(i int) IArrayItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayItemContext); ok {
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

	return t.(IArrayItemContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, mindfckParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(mindfckParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mindfckParserBYTE_NUMBER || _la == mindfckParserNUMBER {
		{
			p.SetState(151)
			p.ArrayItem()
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == mindfckParserT__6 {
			{
				p.SetState(152)
				p.Match(mindfckParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(153)
				p.ArrayItem()
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(161)
		p.Match(mindfckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayItemContext is an interface to support dynamic dispatch.
type IArrayItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	BYTE_NUMBER() antlr.TerminalNode

	// IsArrayItemContext differentiates from other interfaces.
	IsArrayItemContext()
}

type ArrayItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayItemContext() *ArrayItemContext {
	var p = new(ArrayItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayItem
	return p
}

func InitEmptyArrayItemContext(p *ArrayItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayItem
}

func (*ArrayItemContext) IsArrayItemContext() {}

func NewArrayItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayItemContext {
	var p = new(ArrayItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_arrayItem

	return p
}

func (s *ArrayItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayItemContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(mindfckParserNUMBER, 0)
}

func (s *ArrayItemContext) BYTE_NUMBER() antlr.TerminalNode {
	return s.GetToken(mindfckParserBYTE_NUMBER, 0)
}

func (s *ArrayItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterArrayItem(s)
	}
}

func (s *ArrayItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitArrayItem(s)
	}
}

func (s *ArrayItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitArrayItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) ArrayItem() (localctx IArrayItemContext) {
	localctx = NewArrayItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, mindfckParserRULE_arrayItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !(_la == mindfckParserBYTE_NUMBER || _la == mindfckParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayIndexContext is an interface to support dynamic dispatch.
type IArrayIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsArrayIndexContext differentiates from other interfaces.
	IsArrayIndexContext()
}

type ArrayIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayIndexContext() *ArrayIndexContext {
	var p = new(ArrayIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayIndex
	return p
}

func InitEmptyArrayIndexContext(p *ArrayIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mindfckParserRULE_arrayIndex
}

func (*ArrayIndexContext) IsArrayIndexContext() {}

func NewArrayIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayIndexContext {
	var p = new(ArrayIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mindfckParserRULE_arrayIndex

	return p
}

func (s *ArrayIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayIndexContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(mindfckParserNUMBER, 0)
}

func (s *ArrayIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.EnterArrayIndex(s)
	}
}

func (s *ArrayIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mindfckListener); ok {
		listenerT.ExitArrayIndex(s)
	}
}

func (s *ArrayIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mindfckVisitor:
		return t.VisitArrayIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mindfckParser) ArrayIndex() (localctx IArrayIndexContext) {
	localctx = NewArrayIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, mindfckParserRULE_arrayIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(mindfckParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *mindfckParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *mindfckParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
