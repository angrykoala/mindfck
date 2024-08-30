package env

import (
	"fmt"
	"mindfck/common"
)

// Same as codegen.Main
const RESERVED_MEMORY = 7

type MindfuckEnv struct {
	variables      map[string]int
	reservedMemory common.ItemSet
	freedMemory    []int
	memoryBegin    int
}

func New() *MindfuckEnv {
	return &MindfuckEnv{
		variables:      make(map[string]int),
		reservedMemory: common.ItemSet{},
		freedMemory:    []int{},
		memoryBegin:    RESERVED_MEMORY + 1,
	}
}

func (env *MindfuckEnv) ReserveMemory(label string) int {
	_, hasLabel := env.variables[label]

	if hasLabel {
		panic("Cannot reserve label, already reserved")
	}

	var varPos int

	if len(env.freedMemory) > 0 {
		// Reuse position
		varPos = env.freedMemory[0]
		env.freedMemory = env.freedMemory[1:]
	} else {
		varPos = env.reservedMemory.Size() + env.memoryBegin
	}

	env.variables[label] = varPos
	env.reservedMemory.Add(varPos)

	return varPos
}

func (env *MindfuckEnv) GetPosition(label string) int {
	position, hasLabel := env.variables[label]

	if !hasLabel {
		panic(fmt.Sprintf("Label %s not found", label))
	}

	return position
}
