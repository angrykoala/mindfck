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
	assert.Equal(t, []byte{'1', '0', '0', 10}, interpreter.Output)
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
	assert.Equal(t, []byte{'5', '5', '0'}, interpreter.Output)
}

func TestIntSub(t *testing.T) {
	input := `
	int a
    int b
	a = 250
    b = 300 - a
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
	assert.Equal(t, []byte{'5', '0'}, interpreter.Output)
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
	assert.Equal(t, []byte{'3', '0', '0', '0'}, interpreter.Output)
}

func TestIntDiv(t *testing.T) {
	input := `
    int b
    b = 100 / 2
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
	assert.Equal(t, []byte{'5', '0'}, interpreter.Output)
}
