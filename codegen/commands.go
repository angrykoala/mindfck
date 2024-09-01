package codegen

import "mindfck/env"

// Based on https://esolangs.org/wiki/Brainfuck_algorithms
// and https://gist.github.com/roachhd/dce54bec8ba55fb17d3a

type CommandHandler struct {
	writer  *writer
	env     *env.MindfuckEnv
	pointer int
}

func New() *CommandHandler {
	cmd := CommandHandler{
		writer: newWriter(),
		env:    env.New(0),
	}

	return &cmd
}

func (c *CommandHandler) Env() *env.MindfuckEnv {
	return c.env
}

func (c *CommandHandler) Declare(label string) env.Variable {
	return c.env.DeclareVariable(label)
}

func (c *CommandHandler) Release(v env.Variable) {
	c.env.ReleaseVariable(v)
}

func (c *CommandHandler) Print(v env.Variable) {
	c.goTo(v)
	c.writer.command(BFOut)
}

// func (c *CommandHandler) PrintNumber(label string) {
// 	c.goTo(label)

// 	c.writer.write("x>>++++++++++<<[->+>-[>+>>]>[+[-<+>]>+>>]<<<<<<]>>[-]>>>++++++++++<[->-[>+>>]>[+[-<+>]>+>>]<<<<<]>[-]>>[>++++++[-<++++++++>]<.<<+>+>[-]]<[<[->-<]++++++[->++++++++<]>.[-]]<<++++++[-<++++++++>]<.[-]<<[-<+>]<")
// }

func (c *CommandHandler) Set(v env.Variable, value int) {
	c.Reset(v)
	c.add(v, value)
}

func (c *CommandHandler) Sub(v env.Variable, count int) {
	c.add(v, -1)
}

func (c *CommandHandler) Inc(v env.Variable) {
	c.goTo(v)
	c.writer.command(BFInc)
}

func (c *CommandHandler) Dec(v env.Variable) {
	c.goTo(v)
	c.writer.command(BFDec)
}

// Resets cell to 0
func (c *CommandHandler) Reset(v env.Variable) {
	c.Loop(v, func() {
		c.writer.command(BFDec)
	})
}

// from -> to
func (c *CommandHandler) Move(from env.Variable, to env.Variable) {
	c.Reset(to)
	c.Loop(from, func() {
		c.Inc(to)
		c.Dec(from)
	})
}

// Copy current cell into to, using temp cell, ends in origin and resets temp
func (c *CommandHandler) Copy(from env.Variable, to env.Variable) {
	temp0 := c.env.DeclareAnonVariable()
	temp1 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp0)
	defer c.env.ReleaseVariable(temp1)

	// Reset temp and to
	c.Reset(temp0)
	c.Reset(to)

	c.Loop(from, func() {
		c.Inc(to)
		c.Inc(temp0)
		c.Dec(from)
	})

	c.Move(temp0, from)
}

// Compares a and b, result is a boolean in res
func (c *CommandHandler) Equals(x env.Variable, y env.Variable, res env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.Copy(x, temp)
	c.SubCell(y, temp)

	c.Set(res, 1)

	c.If(temp, func() {
		c.Set(res, 0)
	})
}

func (c *CommandHandler) And(x env.Variable, y env.Variable, res env.Variable) {
	c.Reset(res)
	c.If(x, func() {
		c.If(y, func() {
			c.Set(res, 1)
		})
	})
}

func (c *CommandHandler) Or(x env.Variable, y env.Variable, res env.Variable) {
	c.Reset(res)
	c.If(x, func() {
		c.Set(res, 1)
	})

	c.If(y, func() {
		c.Set(res, 1)
	})
}

// Boolean NOT
func (c *CommandHandler) Not(x env.Variable, res env.Variable) {
	c.Set(res, 1)
	c.If(x, func() {
		c.Dec(res)
	})
}

// Boolean NOT
func (c *CommandHandler) Add(a env.Variable, b env.Variable, res env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.Copy(a, temp)
	c.addTo(b, temp)

	c.Move(temp, res)
}

// Substracts cell a to b, b is modified
func (c *CommandHandler) SubCell(a env.Variable, b env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.Reset(temp)

	c.Loop(a, func() {
		c.Inc(temp)
		c.Dec(b)
		c.Dec(a)
	})

	c.Move(temp, a)
}

// Multiply cell a and b
func (c *CommandHandler) Mult(a env.Variable, b env.Variable, res env.Variable) {
	temp := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp)
	c.Copy(a, temp)
	c.Reset(res)

	c.Loop(temp, func() {
		c.addTo(b, res)
		c.Dec(temp)
	})
}

// Runs code inside if the current cell is truthy
// Accepts a code function using the command handler
// The resulting function must end in the same position as it begins!
func (c *CommandHandler) If(cond env.Variable, code func()) {
	c.goTo(cond)
	c.beginLoop()
	code()
	c.goTo(env.NIL)
	c.endLoop()
}

// Loops, ensuring that the loop begins and ends in condCell
func (c *CommandHandler) Loop(condCell env.Variable, code func()) {
	c.goTo(condCell)
	c.beginLoop()
	code()
	c.goTo(condCell)
	c.endLoop()
}

func (c *CommandHandler) Comment(comment string) {
	c.writer.comment(comment)
}

func (c *CommandHandler) Compile() string {
	return c.writer.print()
}

// Core functionality

func (c *CommandHandler) beginLoop() {
	c.writer.command(BFLoopBegin)
}

func (c *CommandHandler) endLoop() {
	c.writer.command(BFLoopEnd)
}

// Move pointer n positions, left or right
func (c *CommandHandler) movePointer(pos int) {
	c.pointer += pos
	if pos > 0 {
		for i := 0; i < pos; i++ {
			c.writer.command(BFIncPointer)
		}
	}

	if pos < 0 {
		for i := 0; i > pos; i-- {
			c.writer.command(BFDecPointer)
		}
	}
}

func (c *CommandHandler) getPos(v env.Variable) int {
	return v.Position()
}

// Move pointer to position (static)
func (c *CommandHandler) goTo(v env.Variable) {
	cell := c.getPos(v)
	var diff = cell - c.pointer
	c.movePointer(diff)
}

func (c *CommandHandler) add(v env.Variable, count int) {
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

// Adds cell a to b, b is modified
func (c *CommandHandler) addTo(a env.Variable, b env.Variable) {
	temp0 := c.env.DeclareAnonVariable()
	defer c.env.ReleaseVariable(temp0)
	c.Reset(temp0)

	c.Loop(a, func() {
		c.Inc(temp0)
		c.Inc(b)
		c.Dec(a)
	})

	c.Move(temp0, a)
}
