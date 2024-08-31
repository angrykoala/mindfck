package main

import (
	"fmt"
	"mindfck/bfinterpreter"
	"mindfck/codegen"
)

func main() {
	cmd := codegen.New()

	var1 := cmd.Declare("var1")
	// var2 := cmd.Declare("var2")
	// var3 := cmd.Declare("var3")

	cmd.Set(var1, 3)

	// cmd.Add(var1, var2, var3)
	code := cmd.Compile()

	fmt.Println(code)

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	fmt.Println(string(interpreter.Output))
	fmt.Println(interpreter.Memory)
}
