package env

import "fmt"

type VarType string

const (
	BYTE VarType = "byte"
	INT  VarType = "int"
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

// type ByteVariable struct {
// 	position int
// 	label    string
// }

// func (v *ByteVariable) Position() int {
// 	return v.position
// }

// func (v *ByteVariable) HasLabel() bool {
// 	return v.label != ""
// }

// func (v *ByteVariable) Label() string {
// 	return v.label
// }

// func (v *ByteVariable) Size() int {
// 	return 1
// }

// func (v *ByteVariable) Type() VarType {
// 	return BYTE
// }

// func (v *ByteVariable) GetByte(i int) *ByteVariable {
// 	if i > 0 {
// 		panic("invalid byte")
// 	}
// 	return v
// }

// type ArrayVariable struct {
// 	position int
// 	label    string
// 	size     int
// }

// func (v *ArrayVariable) Position() int {
// 	return v.position
// }

// func (v *ArrayVariable) HasLabel() bool {
// 	return v.label != ""
// }

// func (v *ArrayVariable) Label() string {
// 	return v.label
// }

// func (v *ArrayVariable) Size() int {
// 	return v.size
// }

// func (v *ArrayVariable) Get(index int) *ByteVariable {
// 	if index >= v.size {
// 		panic("Invalid index in array variable")
// 	}
// 	return &ByteVariable{
// 		position: v.position + index,
// 	}
// }
