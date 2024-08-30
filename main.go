package main

import (
	"mindfck/codegen"
	"mindfck/env"
)

func main() {
	cmd := codegen.New()
	varEnv := env.New()

	varEnv.ReserveMemory("var1")
	varEnv.ReserveMemory("var2")

	// cmd.Copy(-2, -9)
	// cmd.Comment("Copy")

	cmd.Set(varEnv.GetPosition("var1"), 4)
	cmd.Set(varEnv.GetPosition("var2"), 3)

	cmd.AddTo(varEnv.GetPosition("var1"), varEnv.GetPosition("var2"))
	cmd.Out(varEnv.GetPosition("var2"))
	cmd.Print()

	// interpreter := bfinterpreter.New()
	// interpreter.Run("++>++.>>++++")
	// fmt.Println(interpreter.Output)
	// fmt.Println(interpreter.Memory)
}
