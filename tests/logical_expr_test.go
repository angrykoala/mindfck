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
	print 1 == 1 and 1 == 0
    print not 1
    print !0
    print 1 or 0
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
	print !(1 == 1 and 1 == 0)
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
