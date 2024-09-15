package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExprPrecedence(t *testing.T) {
	input := `
	byte a
	byte b
	a = 3 + 2 * 2
	b = 3 * (2  + (2 / 2)) - 5/5
	print a
	print b
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
	assert.Equal(t, []byte{7, 8}, interpreter.Output)
}
