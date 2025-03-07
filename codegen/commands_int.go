package codegen

import (
	"fmt"
	"mindfck/env"
)

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
	temp := c.env.DeclareAnonByte()
	defer c.Release(temp)

	secondByte := v.GetByte(1)
	c.IncByte(secondByte)
	c.NotByte(secondByte, temp)
	c.If(temp, func() {
		firstByte := v.GetByte(0)
		c.IncByte(firstByte)
	})
}

func (c *CommandHandler) DecInt(v env.Variable) {
	assertInt(v)
	temp := c.env.DeclareAnonByte()
	defer c.Release(temp)

	secondByte := v.GetByte(1)
	c.NotByte(secondByte, temp)
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

	c.GteInt(remainder, b, isRemainderBigger) // TODO: make this support ints

	c.While(isRemainderBigger, func() {
		c.IncInt(res)
		c.subToInt(b, remainder)

		c.GteInt(remainder, b, isRemainderBigger)
	})
}

func (c *CommandHandler) EqualsInt(x env.Variable, y env.Variable, res env.Variable) {
	assertInt(x)
	assertInt(y)
	assertBool(res)

	temp := c.env.DeclareAnonVariable(env.INT)
	defer c.env.ReleaseVariable(temp)
	isZero := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(isZero)

	c.Copy(x, temp)
	c.subToInt(y, temp)

	c.SetByte(res, 0)

	c.isZeroInt(temp, isZero)
	c.If(isZero, func() {
		c.SetByte(res, 1)
	})
}

// Compares x+1 > b, cheap Gte
func (c *CommandHandler) GteInt(x env.Variable, y env.Variable, res env.Variable) {
	x2 := c.env.DeclareAnonVariable(env.INT)
	defer c.env.ReleaseVariable(x2)

	// Because these are integers, we just compare GT with an x increased by 1
	c.Copy(x, x2)
	c.IncInt(x2)
	c.GtInt(x2, y, res)
}

func (c *CommandHandler) GtInt(x env.Variable, y env.Variable, res env.Variable) {
	assertInt(x)
	assertInt(y)
	assertBool(res)

	temp := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(temp)
	c.Reset(temp)

	c.GtByte(x.GetByte(0), y.GetByte(0), temp)

	c.If(temp, func() {
		c.SetByte(res, 1)
	})

	c.EqualsByte(x.GetByte(0), y.GetByte(0), temp)
	c.If(temp, func() {
		c.GtByte(x.GetByte(1), y.GetByte(1), res)
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
	// Step 1, sub first byte
	c.subToByte(a.GetByte(0), b.GetByte(0))

	// Step 2, sub second byte with int sub
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
	c.NotByte(isZero, isNotZero)

	c.While(isNotZero, func() {
		code()
		c.isZeroInt(condInt, isZero)
		c.NotByte(isZero, isNotZero)
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

	c.Copy(lastByte, to)
}

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

func (c *CommandHandler) NotInt(x env.Variable, res env.Variable) {
	assertInt(x)
	assertByte(res)
	c.SetByte(res, 0)

	temp := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(temp)

	c.isZeroInt(x, temp)

	c.If(temp, func() {
		c.IncByte(res)
	})
}

func (c *CommandHandler) PrintInt(v env.Variable) {
	assertInt(v)

	v_copy := c.env.DeclareAnonVariable(env.INT)
	defer c.env.ReleaseVariable(v_copy)
	c.Copy(v, v_copy)

	cond_result := c.env.DeclareAnonByte()
	cond_value := c.env.DeclareAnonByte()
	c.SetByte(cond_value, 10)

	// Smallest character
	v1 := c.env.DeclareAnonVariable(env.BYTE)
	v2 := c.env.DeclareAnonVariable(env.BYTE)
	v3 := c.env.DeclareAnonVariable(env.BYTE)
	v4 := c.env.DeclareAnonVariable(env.BYTE)
	v5 := c.env.DeclareAnonVariable(env.BYTE)
	c.Reset(v1)
	c.Reset(v2)
	c.Reset(v3)
	c.Reset(v4)
	c.Reset(v5)
	c.Reset(cond_result)
	defer c.env.ReleaseVariable(v1)
	defer c.env.ReleaseVariable(v2)
	defer c.env.ReleaseVariable(v3)
	defer c.env.ReleaseVariable(v4)
	defer c.env.ReleaseVariable(v5)
	defer c.env.ReleaseVariable(cond_result)
	defer c.env.ReleaseVariable(cond_value)

	c.whileInt(v_copy, func() {
		c.IncByte(v1)
		c.EqualsByte(v1, cond_value, cond_result)

		// If v1 == 10
		c.If(cond_result, func() {
			c.Reset(v1)
			c.IncByte(v2)
			c.EqualsByte(v2, cond_value, cond_result)
			c.If(cond_result, func() {
				c.Reset(v2)
				c.IncByte(v3)
				c.EqualsByte(v3, cond_value, cond_result)
				c.If(cond_result, func() {
					c.Reset(v3)
					c.IncByte(v4)
					c.EqualsByte(v4, cond_value, cond_result)
					c.If(cond_result, func() {
						c.Reset(v4)
						c.IncByte(v5)
						c.EqualsByte(v5, cond_value, cond_result)
					})
				})
			})
		})

		c.DecInt(v_copy)
	})

	char_diff := c.env.DeclareAnonByte()
	c.SetByte(char_diff, 48)
	defer c.env.ReleaseVariable(char_diff)

	// IF 5th chracter is not 0, print it
	c.If(v5, func() {
		c.addToByte(char_diff, v5)
		c.PrintByte(v5)
	})

	cond_value_int := c.env.DeclareAnonVariable(env.INT)
	defer c.env.ReleaseVariable(cond_value_int)

	c.SetInt(cond_value_int, 999)
	c.GtInt(v, cond_value_int, cond_result)
	// IF number > 999, print 4th character
	c.If(cond_result, func() {
		c.addToByte(char_diff, v4)
		c.PrintByte(v4)
	})

	c.SetInt(cond_value_int, 99)
	c.GtInt(v, cond_value_int, cond_result)
	// IF number > 99, print 3rd character
	c.If(cond_result, func() {
		c.addToByte(char_diff, v3)
		c.PrintByte(v3)
	})

	c.SetInt(cond_value_int, 9)
	c.GtInt(v, cond_value_int, cond_result)
	// IF number > 9, print 2nd character
	c.If(cond_result, func() {
		c.addToByte(char_diff, v2)
		c.PrintByte(v2)
	})

	c.addToByte(char_diff, v1)

	c.PrintByte(v1)
}

func assertInt(v env.Variable) {
	if v.Type() != env.INT {
		panic(fmt.Sprintf("invalid type %s, %s expected", v.Type(), env.INT))
	}
}
