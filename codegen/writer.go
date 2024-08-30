package codegen

import (
	"fmt"
	"strings"
)

type BFCommand string

const (
	BFInc        BFCommand = "+"
	BFDec        BFCommand = "-"
	BFIncPointer BFCommand = ">"
	BFDecPointer BFCommand = "<"
	BFOut        BFCommand = "."
	BFIn         BFCommand = ","
	BFLoopBegin  BFCommand = "["
	BFLoopEnd    BFCommand = "]"
)

type writer struct {
	sb *strings.Builder
}

func newWriter() *writer {
	var sb strings.Builder
	var wrt writer
	wrt.sb = &sb

	return &wrt
}

func (wrt *writer) command(command BFCommand) {
	wrt.sb.WriteString(string(command))
}

func (wrt *writer) write(code string) {
	wrt.sb.WriteString(code)
}

func (wrt *writer) comment(comment string) {
	wrt.sb.WriteString("  ")
	wrt.sb.WriteString(comment)
	wrt.sb.WriteString("\n")
}

func (wrt *writer) print() {
	fmt.Println(wrt.sb.String())
}
