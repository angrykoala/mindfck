package main

import (
	"fmt"
	"mindfck/bfinterpreter"
	"mindfck/codegen"
	"mindfck/parser"
	"mindfck/parser/tokens"
)

func main() {
	// code()
	tokens, err := tokens.Tokenizer(`
	byte a
	a = 3 + 2
	`)
	if err != nil {
		panic(err)
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		panic(err)
	}

	fmt.Println(tokens)
	fmt.Println(ast)
}

func code() string {
	cmd := codegen.New()

	var1 := cmd.Declare("var1")
	var2 := cmd.Declare("var2")

	cmd.Set(var1, 20)
	cmd.Set(var2, 50)

	cmd.Add(var1, var2, var1)
	var3 := cmd.Declare("var3")

	cmd.Copy(var1, var3)
	cmd.Print(var3)

	code := cmd.Compile()

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	fmt.Println(string(interpreter.Output))
	fmt.Println(interpreter.Memory)
	return code
}
