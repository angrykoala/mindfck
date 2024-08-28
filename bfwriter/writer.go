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

type Writer struct {
	sb *strings.Builder
}

func NewWriter() *Writer {
	var sb strings.Builder
	var writer Writer
	writer.sb = &sb

	return &writer
}

func (writer Writer) AddCommand(command BFCommand) {
	writer.sb.WriteString(string(command))
}

func (writer Writer) Print() {
	fmt.Println(writer.sb.String())
}
