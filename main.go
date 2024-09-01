package main

import (
	"fmt"
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"mindfck/parser/tokens"
)

func main() {
	// code()
	tokens, err := tokens.Tokenizer(`
	byte a
	byte b
	a = 3 + a
	a = 33 + a
	a = a + 2
	byte c
	a = a + 21
	a = a + 2
	b = 10 + a
	c = a + b
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

	fmt.Println(code)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	fmt.Println(string(interpreter.Output))
	fmt.Println(interpreter.Memory)
}
