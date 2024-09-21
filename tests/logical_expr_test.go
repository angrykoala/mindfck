package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogicalExpressions(t *testing.T) {
	input := `
	print 1b == 1b and 1b == 0b
    print not 1b
    print !0b
    print 1b or 0b
	`

	ast, err := parser.Parse(input)
	if err != nil {
		panic(err)
	}

	code, err := compiler.Compile(ast)
	if err != nil {
		panic(err)
	}

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{0, 0, 1, 1}, interpreter.Output)
}

func TestLogicalExpressionsGrouping(t *testing.T) {
	input := `
	print !(1b == 1b and 1b == 0b)
	`

	ast, err := parser.Parse(input)
	if err != nil {
		panic(err)
	}

	code, err := compiler.Compile(ast)
	if err != nil {
		panic(err)
	}

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{1}, interpreter.Output)
}
