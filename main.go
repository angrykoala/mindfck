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

	// cmd.Shift(10)
	// cmd.Copy(-2, -9)
	// cmd.Comment("Copy")

	cmd.Add(8, 4)
	cmd.Add(9, 3)
	cmd.And(8, 9, 10)
	cmd.Not(11, 12)
	cmd.Print()
}

// var sb strings.Builder

// for i := 0; i < 1000; i++ {
//     sb.WriteString("a")
// }

// fmt.Println(sb.String())
