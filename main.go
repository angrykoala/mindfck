package main

import (
	"fmt"
	"mindfck/parser"
)

func main() {
	// cmd := codegen.New()

	// var1 := cmd.Declare("var1")
	// // var2 := cmd.Declare("var2")
	// // var3 := cmd.Declare("var3")

	// cmd.Set(var1, 3)

	// code := cmd.Compile()

	// fmt.Println(code)

	// interpreter := bfinterpreter.New()
	// interpreter.Run(code)
	// fmt.Println(string(interpreter.Output))
	// fmt.Println(interpreter.Memory)

	tokens, err := parser.Tokenizer("byte  a")
	if err != nil {
		panic(err)
	}

	fmt.Println(tokens)
}
