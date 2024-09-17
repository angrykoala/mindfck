package env

import (
	"fmt"
	"mindfck/utils"
	"slices"
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

func (env *MindfuckEnv) DeclareVariable(label string, varType VarType) Variable {
	position := env.reserveMemory(getSize(varType))

	var newVar = NewVariable(position, varType, label)
	if newVar.HasLabel() {
		_, hasLabel := env.labels[label]

		if hasLabel {
			panic(fmt.Sprintf("Cannot reserve label [%s], already reserved", label))
		}

		env.labels[label] = newVar
	}
	return newVar
}

func (env *MindfuckEnv) DeclareAnonByte() Variable {
	return env.DeclareVariable("", BYTE)
}

func (env *MindfuckEnv) ReleaseVariable(v Variable) {
	if v.Position() < env.memoryBegin {
		panic("release: out of bounds")
	}

	env.releaseMemory(v.Position(), v.Size())

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

func (env *MindfuckEnv) reserveMemory(size int) int {
	if size < 1 {
		panic("Invalid size in reserve Memory")
	}
	var varPos int

	freePos, ok := findFirstConsecutiveSet(env.freedMemory, size)
	if ok {
		// Reuse position if possible
		varPos = env.freedMemory[freePos]
		env.freedMemory = append(env.freedMemory[0:freePos], env.freedMemory[freePos+size:]...)

	} else {
		varPos = len(env.freedMemory) + env.reservedMemory.Size() + env.memoryBegin
	}

	for i := 0; i < size; i++ {
		env.reservedMemory.Add(varPos + i)
	}

	return varPos
}

func (env *MindfuckEnv) releaseMemory(pos int, size int) {
	for i := pos; i < pos+size; i++ {
		if !env.reservedMemory.Has(i) {
			panic(fmt.Sprintf("release unallocated memory: %d. Memory: %d", i, env.reservedMemory.Items()))
		}

		env.reservedMemory.Delete(i)

		env.freedMemory = append(env.freedMemory, i)
	}
	slices.Sort(env.freedMemory)
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
