package env

import (
	"fmt"
	"mindfck/utils"
)

type Variable interface {
	Position() int
	hasLabel() bool
	label() string
}

type NamedVariable struct {
	position int
	_label   string
}

func (v *NamedVariable) Position() int {
	return v.position
}

func (v *NamedVariable) hasLabel() bool {
	return true
}

func (v *NamedVariable) label() string {
	return v._label
}

type AnonVariable struct {
	position int
}

func (v *AnonVariable) Position() int {
	return v.position
}

func (v *AnonVariable) hasLabel() bool {
	return false
}

func (v *AnonVariable) label() string {
	return ""
}

type MindfuckEnv struct {
	labels         map[string]Variable
	reservedMemory utils.ItemSet
	freedMemory    []int
	memoryBegin    int
}

func New(begin int) *MindfuckEnv {
	return &MindfuckEnv{
		labels:         make(map[string]Variable),
		reservedMemory: utils.ItemSet{},
		freedMemory:    []int{},
		memoryBegin:    begin + GlobalsCount,
	}
}

func (env *MindfuckEnv) DeclareVariable(label string) Variable {
	_, hasLabel := env.labels[label]

	if hasLabel {
		panic("Cannot reserve label, already reserved")
	}

	var newVar = &NamedVariable{
		position: env.reserveMemory(),
		_label:   label,
	}
	env.labels[label] = newVar

	return newVar
}

func (env *MindfuckEnv) DeclareAnonVariable() Variable {
	return &AnonVariable{
		position: env.reserveMemory(),
	}
}

func (env *MindfuckEnv) ReleaseVariable(v Variable) {
	if v.Position() < env.memoryBegin {
		panic("release: out of bounds")
	}

	env.reservedMemory.Delete(v.Position())
	env.freedMemory = append(env.freedMemory, v.Position())

	if v.hasLabel() {
		env.releaseLabel(v.label())
	}
}

func (env *MindfuckEnv) ResolveLabel(label string) Variable {
	variable, hasLabel := env.labels[label]

	if !hasLabel {
		panic(fmt.Sprintf("Label %s not found", label))
	}

	return variable
}

func (env *MindfuckEnv) releaseLabel(label string) {
	_, hasLabel := env.labels[label]

	if !hasLabel {
		panic(fmt.Sprintf("Label %s not found", label))
	}
	delete(env.labels, label)
}

func (env *MindfuckEnv) reserveMemory() int {
	var varPos int

	if len(env.freedMemory) > 0 {
		// Reuse position
		varPos = env.freedMemory[0]
		env.freedMemory = env.freedMemory[1:]
	} else {
		varPos = env.reservedMemory.Size() + env.memoryBegin
	}

	env.reservedMemory.Add(varPos)

	return varPos
}
