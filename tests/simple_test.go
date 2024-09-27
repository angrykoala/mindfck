package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	input := `
	byte a
	byte b
	a = 3b
	a = 33b + a
	b = a + 2b
	print a
    print b
    print 5b
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
	assert.Equal(t, []byte{36, 38, 5}, interpreter.Output)
}

func TestComparison(t *testing.T) {
	input := `
	byte a
	byte b
	byte c
    a = 3b
    b = 6b
    c = a < 6b
    print c
    c = a < b
    print c
    c = a > b
    print c
    c = a == 3b
    print c
    c = a == 4b
    print c
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
	assert.Equal(t, []byte{1, 1, 0, 1, 0}, interpreter.Output)
}

func TestLogical(t *testing.T) {
	input := `
	byte a
	byte b
	byte c
    byte d
	a = 2b < 6b
    b = 1b
    c = 0b
	d = a and b
	print d
	d = b and c
	print d
	d = b or c
	print d
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
	assert.Equal(t, []byte{1, 0, 1}, interpreter.Output)
}

func TestComplexExpressions(t *testing.T) {
	input := `
	byte a
	byte b
	byte c
    byte d
	byte e
	a = 0b
    b = 1b
    c = 0b
	d = 1b
	e = a + b + c + 8b + d + 3b
	print e
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
	assert.Equal(t, []byte{13}, interpreter.Output)
}

func TestDecrement(t *testing.T) {
	input := `
	byte a
	a = 2
	a = a - 1b
	print a
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

func TestGTZero(t *testing.T) {
	input := `
	print 3b > 0b
	print 0b > 0b
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
	assert.Equal(t, []byte{1, 0}, interpreter.Output)
}
func TestLT(t *testing.T) {
	input := `
	print 3b < 10b
	print 10b < 10b
	print 12b < 10b
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
	assert.Equal(t, []byte{1, 0, 0}, interpreter.Output)
}
func TestGTE(t *testing.T) {
	input := `
	print 3b >= 10b
	print 10b >= 10b
	print 12b >= 10b
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
	assert.Equal(t, []byte{0, 1, 1}, interpreter.Output)
}
