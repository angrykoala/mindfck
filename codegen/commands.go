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

func (c *CommandHandler) Compile() string {
	return c.writer.print()
}

func (c *CommandHandler) Declare(label string, varType env.VarType) env.Variable {
	return c.env.DeclareVariable(label, varType)
}

func (c *CommandHandler) Release(v env.Variable) {
	c.env.ReleaseVariable(v)
}

func (c *CommandHandler) Print(v env.Variable) {
	c.goTo(v)
	c.out()
}

func (c *CommandHandler) DebugBreak() {
	c.writer.command(BFDebug)
}

// Global Ops

// Resets cell to 0
func (c *CommandHandler) Reset(v env.Variable) {
	c.iterateBytes(v, func(b env.Variable, _ int) {
		c.ResetByte(b)
	})
}

// from -> to
func (c *CommandHandler) Move(from env.Variable, to env.Variable) {
	assertSameSize(from, to)
	c.Reset(to)
	c.iterateBytes(from, func(fromByte env.Variable, i int) {
		targetByte := to.GetByte(i)
		c.MoveByte(fromByte, targetByte)
	})
}

// Copy current cell into to, using temp cell, ends in origin and resets temp
func (c *CommandHandler) Copy(from env.Variable, to env.Variable) {
	assertSameSize(from, to)

	c.iterateBytes(from, func(fromByte env.Variable, i int) {
		targetByte := to.GetByte(i)
		c.CopyByte(fromByte, targetByte)

	})
}

func (c *CommandHandler) Inc(v env.Variable) {
	c.goTo(v)
	c.increment()
}

func (c *CommandHandler) Dec(v env.Variable) {
	c.goTo(v)
	c.decrement()
}

// Control Flow

func (c *CommandHandler) IfElse(cond env.Variable, ifCode func(), elseCode func()) {
	temp0 := c.env.DeclareAnonByte()
	temp1 := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(temp0)
	defer c.env.ReleaseVariable(temp1)
	c.CopyByte(cond, temp0)
	c.SetByte(temp1, 1)

	c.While(temp0, func() {
		ifCode()
		c.SetByte(temp0, 0)
		c.SetByte(temp1, 0)
	})

	c.While(temp1, func() {
		elseCode()
		c.SetByte(temp1, 0)
	})

}

// Runs code inside if the current cell is truthy
// Accepts a code function using the command handler
// The resulting function must end in the same position as it begins!
func (c *CommandHandler) If(cond env.Variable, code func()) {
	temp := c.env.DeclareAnonByte()
	defer c.env.ReleaseVariable(temp)
	c.CopyByte(cond, temp)
	c.While(temp, func() {
		code()
		c.SetByte(temp, 0)
	})

}

// Loops, ensuring that the loop begins and ends in condCell
func (c *CommandHandler) While(condCell env.Variable, code func()) {
	c.goTo(condCell)
	c.beginLoop()
	code()
	c.goTo(condCell)
	c.endLoop()
}

// Core functionality

func (c *CommandHandler) Comment(comment string) {
	c.writer.comment(comment)
}

func (c *CommandHandler) beginLoop() {
	c.writer.command(BFLoopBegin)
}

func (c *CommandHandler) endLoop() {
	c.writer.command(BFLoopEnd)
}

func (c *CommandHandler) out() {
	c.writer.command(BFOut)
}

// Move pointer to first byte of variable
func (c *CommandHandler) goTo(v env.Variable) {
	cell := v.Position()
	var diff = cell - c.pointer
	c.shift(diff)
}

// Move pointer n positions, left or right
func (c *CommandHandler) shift(pos int) {
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

func (c *CommandHandler) increment() {
	c.writer.command(BFInc)
}

func (c *CommandHandler) decrement() {
	c.writer.command(BFDec)
}

// Helpers

func (c *CommandHandler) iterateBytes(v env.Variable, cb func(b env.Variable, i int)) {
	c.goTo(v)
	for i := 0; i < v.Size(); i++ {
		cb(v.GetByte(i), i)
		c.shift(1)
	}
}

// Assertions

func assertSameSize(a env.Variable, b env.Variable) {
	if b.Size() != b.Size() {
		panic("Incompatible size of variables")
	}
}
