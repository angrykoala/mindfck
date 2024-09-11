package parser

import (
	mindfck "mindfck/parser/antlr"
	"mindfck/parser/mfast"
	"mindfck/utils"

	"github.com/antlr4-go/antlr/v4"
)

func Parse(input string) ([]mfast.Stmt, error) {
	inputStream := antlr.NewInputStream(input)
	lexer := mindfck.NewmindfckLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := mindfck.NewmindfckParser(tokenStream)
	tree := parser.Statements()

	listener := &AstGenerator{
		currentScope: newScope(),
		scopeStack:   []scope{},
	}
	// TODO How to do the error handling here?
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.ast, nil
}

type scope struct {
	stmt []mfast.Stmt
	expr []mfast.Expr
}

func newScope() scope {
	return scope{
		stmt: []mfast.Stmt{},
		expr: []mfast.Expr{},
	}
}

type AstGenerator struct {
	*mindfck.BasemindfckListener

	currentScope scope

	scopeStack []scope

	// ast  []mfast.Stmt
	// expr []mfast.Expr

	blockStack [][]mfast.Stmt // Scope should also keep expressions?
}

func (l *AstGenerator) Ast() []mfast.Stmt {
	return l.currentScope.stmt
}

func (l *AstGenerator) EnterStatement(c *mindfck.StatementContext) {
	l.expr = []mfast.Expr{}
}

func (l *AstGenerator) ExitDeclaration(decl *mindfck.DeclarationContext) {
	current := &mfast.Declare{
		Label: decl.Identifier().IDENTIFIER().GetText(),
	}
	l.ast = append(l.ast, current)
}

func (l *AstGenerator) ExitAssignment(assignment *mindfck.AssignmentContext) {
	current := &mfast.Assign{
		To:   assignment.Identifier().GetText(),
		From: l.expr[0],
	}

	l.ast = append(l.ast, current)
}

func (l *AstGenerator) ExitPrint(print *mindfck.PrintContext) {
	current := &mfast.Print{
		Value: l.expr[0],
	}
	l.ast = append(l.ast, current)
}

func (l *AstGenerator) EnterIfConditional(cond *mindfck.IfConditionalContext) {
	// Begin block scope
	l.beginScope()
}

func (l *AstGenerator) ExitIfConditional(cond *mindfck.IfConditionalContext) {
	innerBlock := l.popScope()
	current := &mfast.If{
		Condition: l.expr[0],
		Block:     innerBlock,
	}
	l.ast = append(l.ast, current)
}

func (l *AstGenerator) ExitExpression(c *mindfck.ExpressionContext) {
	if c.Literal() != nil {
		l.expr = append(l.expr, &mfast.Literal{
			Value: utils.ToInt(c.Literal().GetText()),
		})
	} else if c.Identifier() != nil {
		l.expr = append(l.expr, &mfast.VariableExpr{
			Label: c.Identifier().GetText(),
		})
	} else if c.Expression(0) != nil && c.Expression(1) != nil {
		l.expr = []mfast.Expr{&mfast.BinaryExpr{
			Operator: mfast.Operand(c.Operand().GetText()),
			Left:     l.expr[0],
			Right:    l.expr[1],
		}}
	}
}

// func (l *AstGenerator) beginScope() {
// 	l.blockStack = append(l.blockStack, l.ast)
// 	l.ast = []mfast.Stmt{}
// }

// // Pops last scope into current, returns current ast
// func (l *AstGenerator) popScope() []mfast.Stmt {
// 	current := l.ast
// 	l.ast = l.blockStack[len(l.blockStack)-1]
// 	l.blockStack = l.blockStack[:len(l.blockStack)-1]
// 	return current
// }

func (l *AstGenerator) beginScope() {
	l.scopeStack = append(l.scopeStack, l.currentScope)
	l.currentScope = newScope()
}

// Pops last scope into current, returns current scope
func (l *AstGenerator) popScope() scope {
	current := l.currentScope
	l.currentScope = l.scopeStack[len(l.scopeStack)-1]
	l.scopeStack = l.scopeStack[:len(l.scopeStack)-1]
	return current
}
