package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci10(t *testing.T) {
	input := `
	byte n
	n = 10
	byte i
	i = 0

	byte a
	byte b
	byte c
	a=0
	b=1

    while (i<n) {
        print a
		c = b
		b = a+b
		a=c
		byte one
		one = 1
		i=i+one
    }
	`

	ast, err := parser.Parse(input)
	assert.Nil(t, err)

	code, err := compiler.Compile(ast)
	assert.Nil(t, err)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, []byte{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, interpreter.Output)
}
