package parser

import (
	"mindfck/mfast"
	mindfck "mindfck/parser/antlr"

	"github.com/antlr4-go/antlr/v4"
)

func Parse(input string) ([]mfast.Stmt, error) {
	inputStream := antlr.NewInputStream(input)
	lexer := mindfck.NewmindfckLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := mindfck.NewmindfckParser(tokenStream)
	tree := parser.Statements()
	visitor := &AstGeneratorVisitor{}

	ast := visitor.Visit(tree)
	return ast.([]mfast.Stmt), nil
}
