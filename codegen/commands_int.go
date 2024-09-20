package codegen

import (
	"fmt"
	"mindfck/env"
)

// TODO: Fix so it is proper print
func (c *CommandHandler) PrintInt(v env.Variable) {
	assertInt(v)
	c.goTo(v)
	c.out()
	c.shift(1)
	c.out()
}
func (c *CommandHandler) SetInt(v env.Variable, value int) {
	assertInt(v)

	leadingByte := value / 256
	remainder := value % 256

	b1 := v.GetByte(0)
	b2 := v.GetByte(1)

	c.SetByte(b1, leadingByte)
	c.SetByte(b2, remainder)
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

func (c *CommandHandler) AddInt(a env.Variable, b env.Variable, res env.Variable) {
	assertInt(a)
	assertInt(b)
	assertInt(res)

	c.Copy(a, res)
	c.addToInt(b, res)
}

func (c *CommandHandler) SubInt(a env.Variable, b env.Variable, res env.Variable) {
	assertInt(a)
	assertInt(b)
	assertInt(res)

	c.Copy(a, res)
	c.subToInt(b, res)
}

// Multiply cell a and b
func (c *CommandHandler) MultInt(a env.Variable, b env.Variable, res env.Variable) {
	println("MULTBYTE")
	temp := c.env.DeclareAnonVariable(env.INT)
	defer c.env.ReleaseVariable(temp)
	c.Copy(a, temp)
	c.Reset(res)

	c.whileInt(temp, func() {
		c.addToInt(b, res)
		c.DecInt(temp)
	})
}

// // Divide cell a and b
func (c *CommandHandler) DivInt(a env.Variable, b env.Variable, res env.Variable) {
	assertInt(a)
	assertInt(b)
	assertInt(res)

	remainder := c.env.DeclareAnonVariable(env.INT)
	defer c.env.ReleaseVariable(remainder)
	isRemainderBigger := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(isRemainderBigger)

	c.Reset(res)
	c.Copy(a, remainder)

	c.Gte(remainder, b, isRemainderBigger) // TODO: make this support ints

	c.While(isRemainderBigger, func() {
		c.IncInt(res)
		c.subToInt(b, remainder)

		c.Gte(remainder, b, isRemainderBigger)
	})
}

func (c *CommandHandler) addToInt(a env.Variable, b env.Variable) {
	// Step 1, add first byte
	c.addToByte(a.GetByte(0), b.GetByte(0))

	// Step 2, add second byte with int inc
	aCopy := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(aCopy)
	c.CopyByte(a.GetByte(1), aCopy)

	c.While(aCopy, func() {
		c.IncInt(b)
		c.DecByte(aCopy)
	})
}

func (c *CommandHandler) subToInt(a env.Variable, b env.Variable) {
	// Step 1, sub second byte
	c.subToByte(a.GetByte(1), b.GetByte(1))

	// Step 2, add second byte with int inc
	aCopy := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(aCopy)
	c.CopyByte(a.GetByte(1), aCopy)

	c.While(aCopy, func() {
		c.DecInt(b)
		c.DecByte(aCopy)
	})
}

func (c *CommandHandler) whileInt(condInt env.Variable, code func()) {
	assertInt(condInt)
	isZero := c.env.DeclareAnonByte()
	isNotZero := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(isZero)
	defer c.env.ReleaseVariable(isNotZero)
	c.isZeroInt(condInt, isZero)
	c.Not(isZero, isNotZero)

	c.While(isNotZero, func() {
		code()
		c.isZeroInt(condInt, isZero)
		c.Not(isZero, isNotZero)
	})

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
