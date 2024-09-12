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
		scope:      newScope(),
		scopeStack: []Scope{},
	}
	// TODO How to do the error handling here?
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.Ast(), nil
}

type AstGenerator struct {
	*mindfck.BasemindfckListener

	scope Scope

	scopeStack []Scope
}

func (l *AstGenerator) Ast() []mfast.Stmt {
	return l.scope.stmts
}

func (l *AstGenerator) EnterStatement(c *mindfck.StatementContext) {
	l.scope.clearExpr()
}

func (l *AstGenerator) ExitDeclaration(decl *mindfck.DeclarationContext) {
	current := &mfast.Declare{
		Label: decl.Identifier().IDENTIFIER().GetText(),
	}
	l.scope.appendStmt(current)
}

func (l *AstGenerator) ExitAssignment(assignment *mindfck.AssignmentContext) {
	current := &mfast.Assign{
		To:   assignment.Identifier().GetText(),
		From: l.scope.popLeftExpr(),
	}

	l.scope.appendStmt(current)
}

func (l *AstGenerator) ExitPrint(print *mindfck.PrintContext) {
	current := &mfast.Print{
		Value: l.scope.popLeftExpr(),
	}

	l.scope.appendStmt(current)
}

// func (l *AstGenerator) EnterIfConditional(cond *mindfck.IfConditionalContext) {
// 	// Begin block scope
// 	l.beginScope()
// }

// func (l *AstGenerator) ExitIfConditional(cond *mindfck.IfConditionalContext) {
// 	innerBlock := l.popScope()
// 	current := &mfast.If{
// 		Condition: l.expr[0],
// 		Block:     innerBlock,
// 	}
// 	l.ast = append(l.ast, current)
// }

func (l *AstGenerator) ExitExpression(c *mindfck.ExpressionContext) {
	if c.Literal() != nil {
		l.scope.pushExpr(&mfast.Literal{
			Value: utils.ToInt(c.Literal().GetText()),
		})

	} else if c.Identifier() != nil {
		l.scope.pushExpr(&mfast.VariableExpr{
			Label: c.Identifier().GetText(),
		})
	} else if c.Expression(0) != nil && c.Expression(1) != nil {
		l.scope.pushExpr(&mfast.BinaryExpr{
			Operator: mfast.Operand(c.Operand().GetText()),
			Left:     l.scope.popLeftExpr(),
			Right:    l.scope.popLeftExpr(),
		})
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
	l.scopeStack = append(l.scopeStack, l.scope)
	l.scope = newScope()
}

// Pops last scope into current, returns current scope
func (l *AstGenerator) popScope() Scope {
	current := l.scope
	l.scope = l.scopeStack[len(l.scopeStack)-1]
	l.scopeStack = l.scopeStack[:len(l.scopeStack)-1]
	return current
}
