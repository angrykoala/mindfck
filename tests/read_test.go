package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	input := `
	byte a
	byte b
	read a
    read b
	print a + b
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
	interpreter.RunWithInput(code, []byte{10, 5})
	assert.Equal(t, []byte{15}, interpreter.Output)
}
