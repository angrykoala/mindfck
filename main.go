package main

import (
	"mindfck/codegen"
)

func main() {
	cmd := codegen.New()

	var1 := cmd.Declare("var1")
	var2 := cmd.Declare("var2")
	var3 := cmd.Declare("var3")

	cmd.Set(var1, 65)
	cmd.Set(var2, 3)

	cmd.Add(var1, var2, var3)
	cmd.Out(var3)
	cmd.Print()
}
