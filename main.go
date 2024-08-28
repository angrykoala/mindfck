package main

import (
	"mindfck/bfwriter"
)

func main() {
	writer := bfwriter.NewWriter()

	writer.AddCommand(bfwriter.BFDec)
	writer.AddCommand(bfwriter.BFDec)
	writer.Print()
}

// var sb strings.Builder

// for i := 0; i < 1000; i++ {
//     sb.WriteString("a")
// }

// fmt.Println(sb.String())
