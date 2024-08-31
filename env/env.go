package env

import (
	"fmt"
)

type MindfuckEnv struct {
	variables      map[string]int
	reservedMemory utils.ItemSet
	freedMemory    []int
	memoryBegin    int
}

func New(begin int) *MindfuckEnv {
	return &MindfuckEnv{
		variables:      make(map[string]int),
		reservedMemory: utils.ItemSet{},
		freedMemory:    []int{},
		memoryBegin:    begin,
	}
}

func (env *MindfuckEnv) ReserveMemory(label string) string {
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

	return label
}

func (env *MindfuckEnv) ReleaseMemory(label string) {
	pos := env.GetPosition(label)

	env.reservedMemory.Delete(pos)
	env.freedMemory = append(env.freedMemory, pos)
	delete(env.variables, label)
}

func (env *MindfuckEnv) GetPosition(label string) int {
	position, hasLabel := env.variables[label]

	if !hasLabel {
		panic(fmt.Sprintf("Label %s not found", label))
	}

	return position
}
