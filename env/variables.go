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
