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
