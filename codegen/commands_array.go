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

	temp := c.env.DeclareAnonByte()
	defer c.Release(temp)
	c.SetByte(temp, 91) // [
	c.PrintByte(temp)
	c.SetByte(temp, 44) // ,
	c.iterateBytes(v, func(b env.Variable, i int) {
		c.PrintByte(b)
		if i < v.Size()-1 {
			c.PrintByte(temp)
		}
	})
	c.SetByte(temp, 93) // ]
	c.PrintByte(temp)
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
