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

func (c *CommandHandler) CastByteToInt(from env.Variable, to env.Variable) {
	assertInt(to)
	assertByte(from)
	c.Reset(to)

	lastByte := to.GetByte(1)

	c.Copy(from, lastByte)
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
