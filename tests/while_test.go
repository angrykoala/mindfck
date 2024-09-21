package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhileSimple(t *testing.T) {
	input := `
	byte a
	a = 1
    while (a==1b) {
        print a
        a=a-1b
    }
    print 0b
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{1, 0}, interpreter.Output)
}

func TestWhileSimpleGT(t *testing.T) {
	input := `
	byte a
	a = 3
    while (a>1b) {
        print a
        a=a-1b
    }
    print 0b
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{3, 2, 0}, interpreter.Output)
}

func TestWhileGT(t *testing.T) {
	input := `
	byte a
	a = 10
    while (a>0b) {
        print a
        a = a-1b
    }
    print 0b
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, interpreter.Output)
}

func TestWhileLT(t *testing.T) {
	input := `
	byte a
	a = 0
    while (a<10b) {
        print a
        a = a+1b
    }
    print 0b
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, interpreter.Output)
}

func TestWhileNested(t *testing.T) {
	input := `
	byte acc
	acc = 0
	byte i
	i=0
	byte j
    while (i<10b) {
		j=0
        while(j<5b) {
			acc = acc+1b
			j=j+1b
		}
        i=i+1b
    }
    print acc
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{50}, interpreter.Output)
}
