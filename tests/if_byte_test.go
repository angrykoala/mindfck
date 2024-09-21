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
	byte b
	a = 2
	b = 2
    if (a==b) {
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
	assert.Equal(t, []byte{2, 0, 0}, interpreter.Output)
}

func TestIfFalse(t *testing.T) {
	input := `
	byte a
	a = 2
    if (a==1b) {
        print a
    }
    print 0b
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{0}, interpreter.Output)
}

func TestNestedIf(t *testing.T) {
	input := `
	byte a
	byte b
	a = 2
	b = 3
    if (a==2b) {
        print a
		if (a+b==5b) {
			print 6b
		}
    }
    print 0b
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{2, 6, 0}, interpreter.Output)
}

func TestIfGTZero(t *testing.T) {
	input := `
	byte a
	byte b
	a = 2b
	b = 0b
    if (a>0b) {
        print a
    }
	if (b>0b) {
		print b
	}
    print 0b
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{2, 0}, interpreter.Output)
}

func TestIfElse(t *testing.T) {
	input := `
	byte a
	byte b
	a = 2b
	b = 0b
    if (a>0b) {
        print 1b
    } else {
		print 2b
	}

	if (b>0b) {
		print 3b
	} else {
		print 4b
	}
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{1, 4}, interpreter.Output)
}
