package env

type Variable interface {
	Position() int
	Size() int
	HasLabel() bool
	Label() string
}

type ByteVariable struct {
	position int
	label    string
}

func (v *ByteVariable) Position() int {
	return v.position
}

func (v *ByteVariable) HasLabel() bool {
	return v.label != ""
}

func (v *ByteVariable) Label() string {
	return v.label
}

func (v *ByteVariable) Size() int {
	return 1
}

type ArrayVariable struct {
	position int
	label    string
	size     int
}

func (v *ArrayVariable) Position() int {
	return v.position
}

func (v *ArrayVariable) HasLabel() bool {
	return v.label != ""
}

func (v *ArrayVariable) Label() string {
	return v.label
}

func (v *ArrayVariable) Size() int {
	return v.size
}

func (v *ArrayVariable) Get(index int) *ByteVariable {
	if index >= v.size {
		panic("Invalid index in array variable")
	}
	return &ByteVariable{
		position: v.position + index,
	}
}
