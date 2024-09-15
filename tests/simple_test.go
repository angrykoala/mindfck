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
	a = 3
	a = 33 + a
	b = a + 2
	print a
    print b
    print 5
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
    a = 3
    b = 6
    c = a < 6
    print c
    c = a < b
    print c
    c = a > b
    print c
    c = a == 3
    print c
    c = a == 4
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
	a = 2 < 6
    b = 1
    c = 0
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
	a = 0
    b = 1
    c = 0
	d = 1
	e = a + b + c + 8 + d + 3
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
	a = a - 1
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
	print 3 > 0
	print 0 > 0
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
	print 3 < 10
	print 10 < 10
	print 12 < 10
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
	print 3 >= 10
	print 10 >= 10
	print 12 >= 10
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
