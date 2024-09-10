package parser

import (
	mindfck "mindfck/parser/antlr"
	"mindfck/parser/mfast"
	"mindfck/utils"

	"github.com/antlr4-go/antlr/v4"
)

type AstGenerator struct {
	*mindfck.BasemindfckListener

	ast  []mfast.Stmt
	expr []mfast.Expr
}

func (l *AstGenerator) EnterStatement(c *mindfck.StatementContext) {
	l.expr = []mfast.Expr{}
}

func (l *AstGenerator) ExitStatement(c *mindfck.StatementContext) {
	decl := c.Declaration()
	assignment := c.Assignment()
	print := c.Print_()

	if decl != nil {
		current := &mfast.Declare{
			Label: decl.Identifier().IDENTIFIER().GetText(),
		}
		l.ast = append(l.ast, current)
	} else if assignment != nil {
		current := &mfast.Assign{
			To:   assignment.Identifier().GetText(),
			From: l.expr[0],
		}

		l.ast = append(l.ast, current)
	} else if print != nil {
		current := &mfast.Print{
			Value: l.expr[0],
		}
		l.ast = append(l.ast, current)
	}
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

func Parse(input string) ([]mfast.Stmt, error) {
	inputStream := antlr.NewInputStream(input)
	lexer := mindfck.NewmindfckLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := mindfck.NewmindfckParser(tokenStream)
	tree := parser.Statements()

	listener := &AstGenerator{}
	// TODO How to do the error handling here?
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.ast, nil
}
