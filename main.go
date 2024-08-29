package main

import "mindfck/bfwriter"

func main() {
	cmd := bfwriter.NewCommandHandler()

	cmd.Add(10)
	cmd.Shift(10)
	cmd.MovePointer(10)
	cmd.Comment("Move Pointer")
	cmd.Copy(-2, -9)
	cmd.Comment("Copy")
	cmd.Print()
}

// var sb strings.Builder

// for i := 0; i < 1000; i++ {
//     sb.WriteString("a")
// }

// fmt.Println(sb.String())
