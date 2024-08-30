package main

import (
	"fmt"
	"mindfck/bfinterpreter"
	"mindfck/codegen"
)

func main() {
	cmd := codegen.New()

	// cmd.MovePointer(2)
	// cmd.Add(10)
	// cmd.MovePointer(1)
	// cmd.Add(5)
	// cmd.MovePointer(-1)
	// cmd.Comment("Setup")
	// cmd.AddCell(1, -1)
	// cmd.Comment("Add cell 2 using 3 as temp")

	// cmd.Shift(10)
	// cmd.Copy(-2, -9)
	// cmd.Comment("Copy")

	// cmd.Add(8, 4)
	// cmd.Add(9, 3)
	cmd.Or(8, 9, 10)
	cmd.Print()

	interpreter := bfinterpreter.New()
	interpreter.Run("++>++.>>++++")
	fmt.Println(interpreter.Output)
	fmt.Println(interpreter.Memory)
}
