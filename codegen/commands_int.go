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
	c.EqualsByte(secondByte, zero, temp)
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
	c.EqualsByte(secondByte, zero, temp)
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

// func (c *CommandHandler) EqualsInt(x env.Variable, y env.Variable, res env.Variable) {
// 	assertInt(x)
// 	assertInt(y)
// 	assertBool(res)

// 	temp := c.env.DeclareAnonByte()
// 	defer c.env.ReleaseVariable(temp)
// 	c.CopyByte(x, temp)
// 	c.subTo(y, temp)

// 	c.SetByte(res, 1)

// 	c.If(temp, func() {
// 		c.SetByte(res, 0)
// 	})
// }

// Checks if variable is zero, result is written in res
func (c *CommandHandler) isZeroInt(a env.Variable, res env.Variable) {
	assertInt(a)
	assertBool(res)
	c.SetByte(res, 1)

	c.If(a.GetByte(0), func() {
		c.SetByte(res, 0)
	})
	c.If(a.GetByte(1), func() {
		c.SetByte(res, 0)
	})
}

// // // Substracts int a to b, b is modified
// func (c *CommandHandler) subToInt(a env.Variable, b env.Variable) {
// 	assertInt(a)
// 	assertInt(b)
// 	temp := c.env.DeclareAnonVariable(env.INT)
// 	defer c.env.ReleaseVariable(temp)
// 	c.Reset(temp)

// 	c.While(a, func() {
// 		c.IncByte(temp)
// 		c.DecByte(b)
// 		c.DecByte(a)
// 	})

// 	c.MoveByte(temp, a)
// }

func assertInt(v env.Variable) {
	if v.Type() != env.INT {
		panic(fmt.Sprintf("invalid type %s, %s expected", v.Type(), env.INT))
	}
}

// func (c *CommandHandler) PrintNumber(label string) {
// 	c.goTo(label)

// 	c.writer.write("x>>++++++++++<<[->+>-[>+>>]>[+[-<+>]>+>>]<<<<<<]>>[-]>>>++++++++++<[->-[>+>>]>[+[-<+>]>+>>]<<<<<]>[-]>>[>++++++[-<++++++++>]<.<<+>+>[-]]<[<[->-<]++++++[->++++++++<]>.[-]]<<++++++[-<++++++++>]<.[-]<<[-<+>]<")
// }
