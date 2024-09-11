package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfTrue(t *testing.T) {
	input := `
	byte a
	a = 2
    if (a==2) {
        print a
    }
    print 0
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{2, 0}, interpreter.Output)
}

func TestIfFalse(t *testing.T) {
	input := `
	byte a
	a = 2
    if (a==1) {
        print a
    }
    print 0
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{0}, interpreter.Output)
}
