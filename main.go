//go:generate antlr4 parser/mindfck.g4 -Dlanguage=Go -o parser/antlr/ -Xexact-output-dir -visitor

package main

import (
	"fmt"
	"mindfck/bfinterpreter"
	"mindfck/compiler"
	"mindfck/parser"
	"os"
	"slices"
)

func main() {
	filename := os.Args[1]

	run := slices.Contains(os.Args, "--run")

	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	ast, err := parser.Parse(string(input))
	if err != nil {
		panic(err)
	}
	code, err := compiler.Compile(ast)
	if err != nil {
		panic(err)
	}

	if run {
		interpreter := bfinterpreter.New()
		interpreter.Run(code)
		fmt.Println(string(interpreter.Output))
		// fmt.Println(interpreter.Memory)
	} else {
		fmt.Println(code)
	}
}
