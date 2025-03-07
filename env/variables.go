package env

import "fmt"

type VarType string

const (
	BYTE  VarType = "byte"
	INT   VarType = "int"
	ARRAY VarType = "array"
)

func getSize(varType VarType) int {
	switch varType {
	case BYTE:
		return 1
	case INT:
		return 2
	}

	panic(fmt.Sprintf("invalid variable type %s", varType))
}

type Variable interface {
	Position() int
	Size() int
	HasLabel() bool
	Label() string
	Type() VarType
	GetByte(pos int) Variable
}

type variable struct {
	position int
	label    string
	size     int
	varType  VarType
}

func NewVariable(position int, varType VarType, label string) Variable {
	return &variable{
		position: position,
		label:    label,
		size:     getSize(varType),
		varType:  varType,
	}
}

func NewArrayVariable(position int, label string, size int) Variable {
	return &variable{
		position: position,
		label:    label,
		size:     size,
		varType:  ARRAY,
	}
}

func (v *variable) Position() int {
	return v.position
}

func (v *variable) HasLabel() bool {
	return v.label != ""
}

func (v *variable) Label() string {
	return v.label
}

func (v *variable) Size() int {
	return v.size
}

func (v *variable) Type() VarType {
	return v.varType
}

func (v *variable) GetByte(i int) Variable {
	if i > v.size {
		panic("invalid byte")
	}
	return NewVariable(v.position+i, BYTE, "")
}
