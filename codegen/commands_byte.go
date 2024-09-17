package codegen

import (
	"fmt"
	"mindfck/env"
)

func (c *CommandHandler) SetByte(v env.Variable, value int) {
	c.ResetByte(v)
	c.add(v, value)
}

// Resets cell to 0
func (c *CommandHandler) ResetByte(v env.Variable) {
	c.While(v, func() {
		c.writer.command(BFDec)
	})
}

// from -> to
func (c *CommandHandler) MoveByte(from env.Variable, to env.Variable) {
	c.ResetByte(to)
	c.While(from, func() {
		c.Inc(to)
		c.Dec(from)
	})
}

// Copy current cell into to, using temp cell, ends in origin and resets temp
func (c *CommandHandler) CopyByte(from env.Variable, to env.Variable) {
	temp0 := c.env.DeclareAnonVariable()
	temp1 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp0)
	defer c.env.ReleaseVariable(temp1)

	// Reset temp and to
	c.ResetByte(temp0)
	c.ResetByte(to)

	c.While(from, func() {
		c.Inc(to)
		c.Inc(temp0)
		c.Dec(from)
	})

	c.MoveByte(temp0, from)
}

// Compares x == y, result is a boolean in res
func (c *CommandHandler) Equals(x env.Variable, y env.Variable, res env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.CopyByte(x, temp)
	c.subTo(y, temp)

	c.SetByte(res, 1)

	c.If(temp, func() {
		c.SetByte(res, 0)
	})
}

// Compares x > b, based on https://esolangs.org/wiki/Brainfuck_algorithms#z_=_x_%3E_y
func (c *CommandHandler) Gt(x env.Variable, y env.Variable, z env.Variable) {
	temp0 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp0)
	temp1 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp1)

	x2 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(x2)
	y2 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(y2)
	c.CopyByte(x, x2)
	c.CopyByte(y, y2)

	c.SetByte(temp0, 0)
	c.SetByte(temp1, 0)
	c.SetByte(z, 0)

	c.While(x2, func() {
		c.Inc(temp0)

		c.While(y2, func() {
			c.Dec(y2)
			c.SetByte(temp0, 0)
			c.Inc(temp1)
		})

		c.While(temp0, func() {
			c.Dec(temp0)
			c.Inc(z)
		})
		c.While(temp1, func() {
			c.Dec(temp1)
			c.Inc(y2)
		})

		c.Dec(y2)
		c.Dec(x2)
	})
}

// Compares x+1 > b, cheap Gte
func (c *CommandHandler) Gte(x env.Variable, y env.Variable, res env.Variable) {
	x2 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(x2)

	// Because these are integers, we just compare GT with an x increased by 1
	c.CopyByte(x, x2)
	c.Inc(x2)
	c.Gt(x2, y, res)
}

func (c *CommandHandler) And(x env.Variable, y env.Variable, res env.Variable) {
	c.ResetByte(res)
	c.If(x, func() {
		c.If(y, func() {
			c.SetByte(res, 1)
		})
	})
}

func (c *CommandHandler) Or(x env.Variable, y env.Variable, res env.Variable) {
	c.ResetByte(res)
	c.If(x, func() {
		c.SetByte(res, 1)
	})

	c.If(y, func() {
		c.SetByte(res, 1)
	})
}

func (c *CommandHandler) Not(x env.Variable, res env.Variable) {
	c.SetByte(res, 1)
	c.If(x, func() {
		c.Dec(res)
	})
}

func (c *CommandHandler) Add(a env.Variable, b env.Variable, res env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.CopyByte(a, temp)
	c.addTo(b, temp)

	c.MoveByte(temp, res)
}

func (c *CommandHandler) Sub(a env.Variable, b env.Variable, res env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.CopyByte(a, temp)
	c.subTo(b, temp)

	c.MoveByte(temp, res)
}

// Multiply cell a and b
func (c *CommandHandler) Mult(a env.Variable, b env.Variable, res env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.CopyByte(a, temp)
	c.ResetByte(res)

	c.While(temp, func() {
		c.addTo(b, res)
		c.Dec(temp)
	})
}

// Divide cell a and b
func (c *CommandHandler) Div(a env.Variable, b env.Variable, res env.Variable) {
	remainder := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(remainder)
	isRemainderBigger := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(isRemainderBigger)

	c.ResetByte(res)
	c.CopyByte(a, remainder)

	c.Gte(remainder, b, isRemainderBigger)

	c.While(isRemainderBigger, func() {
		c.Inc(res)
		c.subTo(b, remainder)

		c.Gte(remainder, b, isRemainderBigger)
	})
}

func (c *CommandHandler) add(v env.Variable, count int) {
	assertByte(v)
	c.goTo(v)
	if count > 0 {
		for i := 0; i < count; i++ {
			c.writer.command(BFInc)
		}
	}

	if count < 0 {
		for i := 0; i > count; i-- {
			c.writer.command(BFDec)
		}
	}
}

func assertByte(v env.Variable) {
	if v.Type() != env.BYTE {
		panic(fmt.Sprintf("invalid type %s, %s expected", v.Type(), env.BYTE))
	}
}

// Substracts cell a to b, b is modified
func (c *CommandHandler) subTo(a env.Variable, b env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.ResetByte(temp)

	c.While(a, func() {
		c.Inc(temp)
		c.Dec(b)
		c.Dec(a)
	})

	c.MoveByte(temp, a)
}

// Adds cell a to b, b is modified
func (c *CommandHandler) addTo(a env.Variable, b env.Variable) {
	temp0 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp0)
	c.ResetByte(temp0)

	c.While(a, func() {
		c.Inc(temp0)
		c.Inc(b)
		c.Dec(a)
	})

	c.MoveByte(temp0, a)
}
