package main

import "mindfck/bfwriter"

func main() {
	cmd := bfwriter.NewCommandHandler()

	// cmd.MovePointer(2)
	// cmd.Add(10)
	// cmd.MovePointer(1)
	// cmd.Add(5)
	// cmd.MovePointer(-1)
	// cmd.Comment("Setup")
	// cmd.AddCell(1, -1)
	// cmd.Comment("Add cell 2 using 3 as temp")

	cmd.Add(1)
	cmd.If(2, func() {
		cmd.Add(2)
	})

	// cmd.Shift(10)
	// cmd.Copy(-2, -9)
	// cmd.Comment("Copy")

	cmd.Print()
}

// var sb strings.Builder

// for i := 0; i < 1000; i++ {
//     sb.WriteString("a")
// }

// fmt.Println(sb.String())
