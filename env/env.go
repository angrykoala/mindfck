package env

import (
	"fmt"
	"mindfck/utils"
)

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

func (env *MindfuckEnv) DeclareVariable(label string) *ByteVariable {
	_, hasLabel := env.labels[label]

	if hasLabel {
		panic("Cannot reserve label, already reserved")
	}

	var newVar = &ByteVariable{
		position: env.reserveMemory(),
		label:    label,
	}
	env.labels[label] = newVar

	return newVar
}

func (env *MindfuckEnv) DeclareArrayVariable(label string) *ArrayVariable {
	_, hasLabel := env.labels[label]

	if hasLabel {
		panic("Cannot reserve label, already reserved")
	}

	var newVar = &ArrayVariable{
		position: env.reserveMemory(),
		label:    label,
	}
	env.labels[label] = newVar

	return newVar
}

func (env *MindfuckEnv) DeclareAnonVariable() Variable {
	return &ByteVariable{
		position: env.reserveMemory(),
	}
}

func (env *MindfuckEnv) ReleaseVariable(v Variable) {
	if v.Position() < env.memoryBegin {
		panic("release: out of bounds")
	}

	env.reservedMemory.Delete(v.Position())
	env.freedMemory = append(env.freedMemory, v.Position())
	// slices.Sort(env.freedMemory)

	if v.HasLabel() {
		env.releaseLabel(v.Label())
	}
}

func (env *MindfuckEnv) ResolveLabel(label string) Variable {
	variable, hasLabel := env.labels[label]

	if !hasLabel {
		panic(fmt.Sprintf("env: label %s not found", label))
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
	// return env.reserveMemoryOfSize(1)
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

func (env *MindfuckEnv) reserveMemoryOfSize(size int) int {
	if size < 1 {
		panic("Invalid size in reserve Memory")
	}
	var varPos int

	freePos, ok := findFirstConsecutiveSet(env.freedMemory, size)
	if ok {
		// Reuse position if possible
		varPos = env.freedMemory[freePos]
		env.freedMemory = append(env.freedMemory[0:freePos-1], env.freedMemory[freePos+size+1:]...)

	} else {
		varPos = len(env.freedMemory) + env.reservedMemory.Size() + env.memoryBegin
	}

	for i := 0; i < size; i++ {
		env.reservedMemory.Add(varPos + i)
	}

	return varPos
}

func findFirstConsecutiveSet(nums []int, s int) (pos int, ok bool) {
	// Ensure there are enough elements to form a consecutive set of size s
	if len(nums) < s {
		return 0, false
	}

	// Iterate over the slice, looking for a consecutive sequence of size s
	for i := 0; i <= len(nums)-s; i++ {
		isConsecutive := true
		for j := 1; j < s; j++ {
			if nums[i+j] != nums[i+j-1]+1 {
				isConsecutive = false
				break
			}
		}
		if isConsecutive {
			return i, true
		}
	}
	return 0, false
}
