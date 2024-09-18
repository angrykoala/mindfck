package codegen

import (
	"fmt"
	"mindfck/env"
)

func (c *CommandHandler) PrintInt(v env.Variable) {
	assertInt(v)
	c.goTo(v)
	c.out()
	c.shift(1)
	c.out()
}

func (c *CommandHandler) IncInt(v env.Variable) {
	assertInt(v)
	zero := c.env.DeclareAnonByte()
	temp := c.env.DeclareAnonByte()
	defer c.Release(zero)
	defer c.Release(temp)
	c.Reset(zero)

	secondByte := v.GetByte(1)
	c.IncByte(secondByte)
	c.Equals(secondByte, zero, temp)
	c.If(temp, func() {
		firstByte := v.GetByte(0)
		c.IncByte(firstByte)
	})
}
func (c *CommandHandler) DecInt(v env.Variable) {
	assertInt(v)
	zero := c.env.DeclareAnonByte()
	temp := c.env.DeclareAnonByte()
	defer c.Release(zero)
	defer c.Release(temp)
	c.Reset(zero)

	secondByte := v.GetByte(1)
	c.Equals(secondByte, zero, temp)
	c.If(temp, func() {
		firstByte := v.GetByte(0)
		c.DecByte(firstByte)
	})
	c.DecByte(secondByte)
}

func (c *CommandHandler) CastByteToInt(from env.Variable, to env.Variable) {
	assertInt(to)
	assertByte(from)
	c.Reset(to)

	lastByte := to.GetByte(1)

	c.Copy(from, lastByte)
}

// NOTE: this cast may cause issues
func (c *CommandHandler) CastIntToByte(from env.Variable, to env.Variable) {
	assertInt(from)
	assertByte(to)
	c.Reset(to)

	lastByte := from.GetByte(1)

	c.Copy(lastByte, from)
}

func assertInt(v env.Variable) {
	if v.Type() != env.INT {
		panic(fmt.Sprintf("invalid type %s, %s expected", v.Type(), env.INT))
	}
}

// func (c *CommandHandler) PrintNumber(label string) {
// 	c.goTo(label)

// 	c.writer.write("x>>++++++++++<<[->+>-[>+>>]>[+[-<+>]>+>>]<<<<<<]>>[-]>>>++++++++++<[->-[>+>>]>[+[-<+>]>+>>]<<<<<]>[-]>>[>++++++[-<++++++++>]<.<<+>+>[-]]<[<[->-<]++++++[->++++++++<]>.[-]]<<++++++[-<++++++++>]<.[-]<<[-<+>]<")
// }
