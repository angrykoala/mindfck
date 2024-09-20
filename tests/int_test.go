package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	input := `
	int a
    byte b
	a = 100
    b = 10
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
	assert.Equal(t, []byte{0, 100, 10}, interpreter.Output)
}

func TestIntAdd(t *testing.T) {
	input := `
	int a
    int b
	a = 250
    b = 300 + a
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
	assert.Equal(t, []byte{2, 38}, interpreter.Output)
}
func TestIntMult(t *testing.T) {
	input := `
	int a
    int b
	a = 10
    b = 300 * a
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
	assert.Equal(t, []byte{11, 184}, interpreter.Output)
}
