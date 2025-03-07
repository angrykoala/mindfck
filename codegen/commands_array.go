package codegen

import (
	"fmt"
	"mindfck/env"
)

// // Commands for multi byte variables

func (c *CommandHandler) SetArray(v env.Variable, value []int) {
	assertArrayOfSize(v, len(value))
	c.Reset(v)

	c.iterateBytes(v, func(b env.Variable, i int) {
		c.SetByte(b, value[i])
	})
}

func (c *CommandHandler) PrintArray(v env.Variable) {
	assertArray(v)

	c.iterateBytes(v, func(b env.Variable, _ int) {
		c.PrintByte(b)
	})
}

func assertArrayOfSize(v env.Variable, size int) {
	assertArray(v)
	assertSize(v, size)
}

func assertSize(var1 env.Variable, size int) {
	if var1.Size() != size {
		panic("Variable has not correct size")
	}
}

func assertArray(v env.Variable) {
	if v.Type() != env.ARRAY {
		panic(fmt.Sprintf("invalid type %s, %s expected", v.Type(), env.ARRAY))
	}
}
