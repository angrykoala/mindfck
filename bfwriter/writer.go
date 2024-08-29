package bfwriter

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

func NewWriter() *writer {
	var sb strings.Builder
	var wrt writer
	wrt.sb = &sb

	return &wrt
}

func (wrt *writer) Command(command BFCommand) {
	wrt.sb.WriteString(string(command))
}

func (wrt *writer) Comment(comment string) {
	wrt.sb.WriteString("  ")
	wrt.sb.WriteString(comment)
	wrt.sb.WriteString("\n")
}

func (wrt *writer) Print() {
	fmt.Println(wrt.sb.String())
}
