//go:generate antlr4 parser/mindfck.g4 -Dlanguage=Go -o parser/antlr/ -Xexact-output-dir

package main

import (
	"fmt"
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
)

func main() {
	// code()
	input := `
	byte a
	a = 2
    if (a==2) {
        print a
    }
    print 0
	`

	ast, err := parser.Parse(input)
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
