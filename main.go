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
	debug := slices.Contains(os.Args, "--debug")

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

		if debug {
			fmt.Println("Code Size:", len(code), "characters")
			fmt.Println("Memory:", interpreter.Memory)
			fmt.Println("Memory Size:", len(interpreter.Memory), "bytes")
			fmt.Println("Executed Instructions:", interpreter.ExecInstructions)
		}
	} else {
		fmt.Println(code)
	}
}
