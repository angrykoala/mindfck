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
	// If the variable has been used before. This is an optimization to avoid unused Resets
	IsDirty() bool
	SetDirty()
}

type variable struct {
	position int
	label    string
	size     int
	varType  VarType
	dirty    bool
}

func NewVariable(position int, varType VarType, label string, dirty bool) Variable {
	return &variable{
		position: position,
		label:    label,
		size:     getSize(varType),
		varType:  varType,
		dirty:    dirty,
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

func (v *variable) IsDirty() bool {
	return v.dirty
}

func (v *variable) SetDirty() {
	v.dirty = true
}

func (v *variable) GetByte(i int) Variable {
	if i > v.size {
		panic("invalid byte")
	}
	return NewVariable(v.position+i, BYTE, "", true)
}
