package codegen

import "mindfck/env"

// Based on https://esolangs.org/wiki/Brainfuck_algorithms
// and https://gist.github.com/roachhd/dce54bec8ba55fb17d3a

// Registry cells
const (
	TEMP0 = 0
	TEMP1 = 1
	TEMP2 = 2

	REG0 = 3
	REG1 = 4
	REG2 = 5

	NIL  = "NIL" // Always 0
	MAIN = 7
)

// TODO: registry lock

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

	cmd.env.ReserveMemory(NIL)

	return &cmd
}

func (c *CommandHandler) Declare(label string) string {
	return c.env.ReserveMemory(label)
}

func (c *CommandHandler) Print(label string) {
	c.goTo(label)
	c.writer.command(BFOut)
}

func (c *CommandHandler) PrintNumber(label string) {
	c.goTo(label)

	c.writer.write("x >>++++++++++<<[->+>-[>+>>]>[+[-<+>]>+>>]<<<<<<]>>[-]>>>++++++++++<[->-[>+>>]>[+[-<+>]>+>>]<<<<<]>[-]>>[>++++++[-<++++++++>]<.<<+>+>[-]]<[<[->-<]++++++[->++++++++<]>.[-]]<<++++++[-<++++++++>]<.[-]<<[-<+>]<")
}

func (c *CommandHandler) Set(label string, value int) {
	c.Reset(label)
	c.add(label, value)
}

func (c *CommandHandler) Sub(label string, count int) {
	c.add(label, -1)
}

func (c *CommandHandler) Inc(label string) {
	c.goTo(label)
	c.writer.command(BFInc)
}

func (c *CommandHandler) Dec(label string) {
	c.goTo(label)
	c.writer.command(BFDec)
}

// Resets cell to 0
func (c *CommandHandler) Reset(label string) {
	c.Loop(label, func() {
		c.writer.command(BFDec)
	})
}

// Move the value of current cell to target cell (counting from current cell)
func (c *CommandHandler) Move(from string, to string) {
	c.Reset(to)
	c.Loop(from, func() {
		c.Inc(to)
		c.Dec(from)
	})
}

// Copy current cell into to, using temp cell, ends in origin and resets temp
func (c *CommandHandler) Copy(from string, to string) {
	temp0 := c.env.ReserveMemory("_temp0")
	temp1 := c.env.ReserveMemory("_temp1")
	defer c.env.ReleaseMemory(temp0)
	defer c.env.ReleaseMemory(temp1)

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
func (c *CommandHandler) Equals(x string, y string, res string) {
	temp := c.env.ReserveMemory("_temp2")
	defer c.env.ReleaseMemory(temp)
	c.Copy(x, temp)
	c.SubCell(y, temp)

	c.Set(res, 1)

	c.If(temp, func() {
		c.Set(res, 0)
	})
}

func (c *CommandHandler) And(x string, y string, res string) {
	c.Reset(res)
	c.If(x, func() {
		c.If(y, func() {
			c.Set(res, 1)
		})
	})
}

func (c *CommandHandler) Or(x string, y string, res string) {
	c.Reset(res)
	c.If(x, func() {
		c.Set(res, 1)
	})

	c.If(y, func() {
		c.Set(res, 1)
	})
}

// Boolean NOT
func (c *CommandHandler) Not(x string, res string) {
	c.Set(res, 1)
	c.If(x, func() {
		c.Dec(res)
	})
}

// Boolean NOT
func (c *CommandHandler) Add(a string, b string, res string) {
	c.Copy(a, res)
	c.addTo(b, res)
}

// Substracts cell a to b, b is modified
func (c *CommandHandler) SubCell(a string, b string) {
	temp := c.env.ReserveMemory("_temp3")
	defer c.env.ReleaseMemory(temp)
	c.Reset(temp)

	c.Loop(a, func() {
		c.Inc(temp)
		c.Dec(b)
		c.Dec(a)
	})

	c.Move(temp, a)
}

// Multiply cell a and b
func (c *CommandHandler) Mult(a string, b string, res string) {
	temp := c.env.ReserveMemory("_temp4")
	defer c.env.ReleaseMemory(temp)
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
func (c *CommandHandler) If(cond string, code func()) {
	c.goTo(cond)
	c.beginLoop()
	code()
	c.goTo(NIL)
	c.endLoop()
}

// Loops, ensuring that the loop begins and ends in condCell
func (c *CommandHandler) Loop(condCell string, code func()) {
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

func (c *CommandHandler) getPos(label string) int {
	return c.env.GetPosition(label)
}

// Move pointer to position (static)
func (c *CommandHandler) goTo(label string) {
	cell := c.getPos(label)
	var diff = cell - c.pointer
	c.movePointer(diff)
}

func (c *CommandHandler) add(label string, count int) {
	c.goTo(label)
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
func (c *CommandHandler) addTo(a string, b string) {
	temp0 := c.env.ReserveMemory("_temp0")
	defer c.env.ReleaseMemory(temp0)
	c.Reset(temp0)

	c.Loop(a, func() {
		c.Inc(temp0)
		c.Inc(b)
		c.Dec(a)
	})

	c.Move(temp0, a)
}
