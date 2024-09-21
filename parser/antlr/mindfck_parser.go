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
		"statements", "statement", "declaration", "assignment", "print", "ifConditional",
		"whileLoop", "read", "block", "expression", "identifier", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 115, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 61, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 5, 8, 75, 8, 8, 10, 8, 12,
		8, 78, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		89, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 106, 8, 9, 10, 9, 12, 9, 109, 9, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 0, 1, 18, 12, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 0, 7, 1, 0, 8, 9, 1, 0, 17, 18, 1, 0, 15, 16, 2, 0, 19,
		20, 24, 27, 1, 0, 21, 21, 1, 0, 22, 22, 2, 0, 5, 6, 28, 28, 118, 0, 27,
		1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8,
		45, 1, 0, 0, 0, 10, 48, 1, 0, 0, 0, 12, 62, 1, 0, 0, 0, 14, 70, 1, 0, 0,
		0, 16, 76, 1, 0, 0, 0, 18, 88, 1, 0, 0, 0, 20, 110, 1, 0, 0, 0, 22, 112,
		1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0,
		27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 1, 1, 0, 0, 0, 29, 27, 1, 0,
		0, 0, 30, 37, 3, 4, 2, 0, 31, 37, 3, 6, 3, 0, 32, 37, 3, 8, 4, 0, 33, 37,
		3, 10, 5, 0, 34, 37, 3, 12, 6, 0, 35, 37, 3, 14, 7, 0, 36, 30, 1, 0, 0,
		0, 36, 31, 1, 0, 0, 0, 36, 32, 1, 0, 0, 0, 36, 33, 1, 0, 0, 0, 36, 34,
		1, 0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 3, 1, 0, 0, 0, 38, 39, 7, 0, 0, 0,
		39, 40, 3, 20, 10, 0, 40, 5, 1, 0, 0, 0, 41, 42, 3, 20, 10, 0, 42, 43,
		5, 19, 0, 0, 43, 44, 3, 18, 9, 0, 44, 7, 1, 0, 0, 0, 45, 46, 5, 10, 0,
		0, 46, 47, 3, 18, 9, 0, 47, 9, 1, 0, 0, 0, 48, 49, 5, 11, 0, 0, 49, 50,
		5, 1, 0, 0, 50, 51, 3, 18, 9, 0, 51, 52, 5, 2, 0, 0, 52, 53, 5, 3, 0, 0,
		53, 54, 3, 16, 8, 0, 54, 60, 5, 4, 0, 0, 55, 56, 5, 12, 0, 0, 56, 57, 5,
		3, 0, 0, 57, 58, 3, 16, 8, 0, 58, 59, 5, 4, 0, 0, 59, 61, 1, 0, 0, 0, 60,
		55, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 11, 1, 0, 0, 0, 62, 63, 5, 13,
		0, 0, 63, 64, 5, 1, 0, 0, 64, 65, 3, 18, 9, 0, 65, 66, 5, 2, 0, 0, 66,
		67, 5, 3, 0, 0, 67, 68, 3, 16, 8, 0, 68, 69, 5, 4, 0, 0, 69, 13, 1, 0,
		0, 0, 70, 71, 5, 14, 0, 0, 71, 72, 3, 20, 10, 0, 72, 15, 1, 0, 0, 0, 73,
		75, 3, 2, 1, 0, 74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 76, 77, 1, 0, 0, 0, 77, 17, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80,
		6, 9, -1, 0, 80, 89, 3, 20, 10, 0, 81, 89, 3, 22, 11, 0, 82, 83, 5, 1,
		0, 0, 83, 84, 3, 18, 9, 0, 84, 85, 5, 2, 0, 0, 85, 89, 1, 0, 0, 0, 86,
		87, 5, 23, 0, 0, 87, 89, 3, 18, 9, 6, 88, 79, 1, 0, 0, 0, 88, 81, 1, 0,
		0, 0, 88, 82, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 107, 1, 0, 0, 0, 90,
		91, 10, 5, 0, 0, 91, 92, 7, 1, 0, 0, 92, 106, 3, 18, 9, 6, 93, 94, 10,
		4, 0, 0, 94, 95, 7, 2, 0, 0, 95, 106, 3, 18, 9, 5, 96, 97, 10, 3, 0, 0,
		97, 98, 7, 3, 0, 0, 98, 106, 3, 18, 9, 4, 99, 100, 10, 2, 0, 0, 100, 101,
		7, 4, 0, 0, 101, 106, 3, 18, 9, 3, 102, 103, 10, 1, 0, 0, 103, 104, 7,
		5, 0, 0, 104, 106, 3, 18, 9, 2, 105, 90, 1, 0, 0, 0, 105, 93, 1, 0, 0,
		0, 105, 96, 1, 0, 0, 0, 105, 99, 1, 0, 0, 0, 105, 102, 1, 0, 0, 0, 106,
		109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 19, 1,
		0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 111, 5, 29, 0, 0, 111, 21, 1, 0, 0,
		0, 112, 113, 7, 6, 0, 0, 113, 23, 1, 0, 0, 0, 7, 27, 36, 60, 76, 88, 105,
		107,
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
	mindfckParserCHAR        = 5
	mindfckParserBYTE_NUMBER = 6
	mindfckParserWS          = 7
	mindfckParserBYTE        = 8
	mindfckParserINT         = 9
	mindfckParserPRINT       = 10
	mindfckParserIF          = 11
	mindfckParserELSE        = 12
	mindfckParserWHILE       = 13
	mindfckParserREAD        = 14
	mindfckParserPLUS        = 15
	mindfckParserMINUS       = 16
	mindfckParserTIMES       = 17
	mindfckParserDIVIDE      = 18
	mindfckParserEQUALS      = 19
	mindfckParserDEQUALS     = 20
	mindfckParserAND         = 21
	mindfckParserOR          = 22
	mindfckParserNOT         = 23
	mindfckParserGT          = 24
	mindfckParserGE          = 25
	mindfckParserLT          = 26
	mindfckParserLE          = 27
	mindfckParserNUMBER      = 28
	mindfckParserIDENTIFIER  = 29
)

// mindfckParser rules.
const (
	mindfckParserRULE_statements    = 0
	mindfckParserRULE_statement     = 1
	mindfckParserRULE_declaration   = 2
	mindfckParserRULE_assignment    = 3
	mindfckParserRULE_print         = 4
	mindfckParserRULE_ifConditional = 5
	mindfckParserRULE_whileLoop     = 6
	mindfckParserRULE_read          = 7
	mindfckParserRULE_block         = 8
	mindfckParserRULE_expression    = 9
	mindfckParserRULE_identifier    = 10
	mindfckParserRULE_literal       = 11
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536899328) != 0 {
		{
			p.SetState(24)
			p.Statement()
		}

		p.SetState(29)
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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case mindfckParserBYTE, mindfckParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Declaration()
		}

	case mindfckParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Assignment()
		}

	case mindfckParserPRINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.Print_()
		}

	case mindfckParserIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
			p.IfConditional()
		}

	case mindfckParserWHILE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(34)
			p.WhileLoop()
		}

	case mindfckParserREAD:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(35)
			p.Read()
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
		p.SetState(38)

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
		p.SetState(39)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	EQUALS() antlr.TerminalNode
	Expression() IExpressionContext

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
	p.EnterRule(localctx, 6, mindfckParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Identifier()
	}
	{
		p.SetState(42)
		p.Match(mindfckParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
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
	p.EnterRule(localctx, 8, mindfckParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(mindfckParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
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
	p.EnterRule(localctx, 10, mindfckParserRULE_ifConditional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(mindfckParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(mindfckParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.expression(0)
	}
	{
		p.SetState(51)
		p.Match(mindfckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(mindfckParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Block()
	}
	{
		p.SetState(54)
		p.Match(mindfckParserT__3)
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

	if _la == mindfckParserELSE {
		{
			p.SetState(55)
			p.Match(mindfckParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(mindfckParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Block()
		}
		{
			p.SetState(58)
			p.Match(mindfckParserT__3)
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
	p.EnterRule(localctx, 12, mindfckParserRULE_whileLoop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(mindfckParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.expression(0)
	}
	{
		p.SetState(65)
		p.Match(mindfckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(mindfckParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Block()
	}
	{
		p.SetState(68)
		p.Match(mindfckParserT__3)
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
	p.EnterRule(localctx, 14, mindfckParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(mindfckParserREAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
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
	p.EnterRule(localctx, 16, mindfckParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536899328) != 0 {
		{
			p.SetState(73)
			p.Statement()
		}

		p.SetState(78)
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
	_startState := 18
	p.EnterRecursionRule(localctx, 18, mindfckParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case mindfckParserIDENTIFIER:
		{
			p.SetState(80)
			p.Identifier()
		}

	case mindfckParserCHAR, mindfckParserBYTE_NUMBER, mindfckParserNUMBER:
		{
			p.SetState(81)
			p.Literal()
		}

	case mindfckParserT__0:
		{
			p.SetState(82)
			p.Match(mindfckParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.expression(0)
		}
		{
			p.SetState(84)
			p.Match(mindfckParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case mindfckParserNOT:
		{
			p.SetState(86)
			p.Match(mindfckParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.expression(6)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(91)

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
					p.SetState(92)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(94)

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
					p.SetState(95)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(97)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&253231104) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(98)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(100)

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
					p.SetState(101)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, mindfckParserRULE_expression)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(103)

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
					p.SetState(104)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 20, mindfckParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
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
	p.EnterRule(localctx, 22, mindfckParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435552) != 0) {
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

func (p *mindfckParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
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

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
