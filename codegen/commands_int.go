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

func assertInt(v env.Variable) {
	if v.Type() != env.INT {
		panic(fmt.Sprintf("invalid type %s, %s expected", v.Type(), env.INT))
	}
}

// func (c *CommandHandler) PrintNumber(label string) {
// 	c.goTo(label)

// 	c.writer.write("x>>++++++++++<<[->+>-[>+>>]>[+[-<+>]>+>>]<<<<<<]>>[-]>>>++++++++++<[->-[>+>>]>[+[-<+>]>+>>]<<<<<]>[-]>>[>++++++[-<++++++++>]<.<<+>+>[-]]<[<[->-<]++++++[->++++++++<]>.[-]]<<++++++[-<++++++++>]<.[-]<<[-<+>]<")
// }
