package codegen

import (
	"fmt"
	"mindfck/env"
)

// // Commands for multi byte variables

func (c *CommandHandler) SetArray(v env.Variable, value []int) {
	assertArrayOfSize(v, len(value))
	c.Reset(v)

	c.iterateArray(v, func(b env.Variable, i int, total_index int) {
		c.SetByte(b, value[i])
	})
}

// Iterate array data, skipping the head
func (c *CommandHandler) iterateArray(v env.Variable, cb func(b env.Variable, i int, total_index int)) {
	c.iterateBytes(v, func(b env.Variable, i int) {
		if i >= env.ARRAY_HEAD_SIZE {
			cb(b, i-env.ARRAY_HEAD_SIZE, i)
		}
	})
}

func (c *CommandHandler) PrintArray(v env.Variable) {
	assertArray(v)

	temp := c.env.DeclareAnonByte()
	defer c.Release(temp)
	c.SetByte(temp, 91) // [
	c.PrintByte(temp)
	c.SetByte(temp, 44) // ,
	c.iterateArray(v, func(b env.Variable, i int, total_index int) {
		c.PrintByte(b)
		if total_index < v.Size()-1 {
			c.PrintByte(temp)
		}
	})
	c.SetByte(temp, 93) // ]
	c.PrintByte(temp)
}

func assertArrayOfSize(v env.Variable, size int) {
	assertArray(v)
	assertSize(v, size+env.ARRAY_HEAD_SIZE)
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
