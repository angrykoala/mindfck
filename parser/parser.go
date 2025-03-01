package parser

import (
	"fmt"
	"mindfck/mfast"
	mindfck "mindfck/parser/antlr"

	"github.com/antlr4-go/antlr/v4"
)

type CustomSyntaxError struct {
	line, column int
	msg          string
}

func (m *CustomSyntaxError) Error() string {
	return fmt.Sprintf("Error in line %d, column %d: %s", m.line, m.column, m.msg)
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func Parse(input string) ([]mfast.Stmt, error) {

	inputStream := antlr.NewInputStream(input)

	lexerErrors := &CustomErrorListener{}
	lexer := mindfck.NewmindfckLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserErrors := &CustomErrorListener{}
	parser := mindfck.NewmindfckParser(tokenStream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)

	tree := parser.Statements()
	visitor := &AstGeneratorVisitor{}
	ast := visitor.Visit(tree)

	if len(lexerErrors.Errors) > 0 {
		for _, e := range lexerErrors.Errors {
			fmt.Println("\t", e.Error())
		}
		return nil, fmt.Errorf("parsing error")
	}

	if len(parserErrors.Errors) > 0 {
		for _, e := range parserErrors.Errors {
			fmt.Println("\t", e.Error())
		}
		return nil, fmt.Errorf("parsing error")
	}

	return ast.([]mfast.Stmt), nil
}
