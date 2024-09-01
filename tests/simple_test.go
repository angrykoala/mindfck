package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"mindfck/parser/tokens"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	tokens, err := tokens.Tokenizer(`
	byte a
	byte b
	a = 3
	a = 33 + a
	b = a + 2
	print a
    print b
    print 5
	`)
	if err != nil {
		panic(err)
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		panic(err)
	}

	code, err := compiler.Compile(ast)
	if err != nil {
		panic(err)
	}

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, interpreter.Output, []byte{36, 38, 5})
}

func TestLogical(t *testing.T) {
	tokens, err := tokens.Tokenizer(`
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
	`)
	if err != nil {
		panic(err)
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		panic(err)
	}

	code, err := compiler.Compile(ast)
	if err != nil {
		panic(err)
	}

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, interpreter.Output, []byte{1, 1, 0, 1, 0})
}
