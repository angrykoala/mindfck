package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	input := `
    byte test[3]
    test = [97,98,99]
    print test
    
    print ' '
    
    test[2]=100
    
    print test
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
	assert.Equal(t, []byte{'[', 97, ',', 98, ',', 99, ']', ' ', '[', 97, ',', 98, ',', 100, ']'}, interpreter.Output)
}
