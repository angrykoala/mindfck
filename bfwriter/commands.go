package bfwriter

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

	NIL  = 6 // Always 0
	MAIN = 7
)

// TODO: registry lock

type CommandHandler struct {
	writer *writer

	pointer int
}

func NewCommandHandler() *CommandHandler {
	cmd := CommandHandler{
		writer: NewWriter(),
	}

	return &cmd
}

// Move pointer to position (static)
func (c *CommandHandler) goTo(cell int) {
	// if pos < MAIN {
	// 	panic("Out of bounds")
	// }
	var diff = cell - c.pointer
	c.movePointer(diff)
}

func (c *CommandHandler) Set(cell int, value int) {
	c.Reset(cell)
	c.Add(cell, value)
}

func (c *CommandHandler) Add(cell int, count int) {
	c.goTo(cell)
	if count > 0 {
		for i := 0; i < count; i++ {
			c.writer.Command(BFInc)
		}
	}

	if count < 0 {
		for i := 0; i > count; i-- {
			c.writer.Command(BFDec)
		}
	}
}

func (c *CommandHandler) Sub(cell int, count int) {
	c.Add(cell, -1)
}

func (c *CommandHandler) Inc(cell int) {
	c.goTo(cell)
	c.writer.Command(BFInc)
}

func (c *CommandHandler) Dec(cell int) {
	c.goTo(cell)
	c.writer.Command(BFDec)
}

// Resets cell to 0
func (c *CommandHandler) Reset(cell int) {
	c.Loop(cell, func() {
		c.writer.Command(BFDec)
	})
}

// Move the value of current cell to target cell (counting from current cell)
func (c *CommandHandler) Shift(from int, to int) {
	c.Reset(to)
	c.Loop(from, func() {
		c.Inc(to)
		c.Dec(from)
	})
}

// Copy current cell into to, using temp cell, ends in origin and resets temp
func (c *CommandHandler) Copy(from int, to int) {
	if from == TEMP0 || to == TEMP0 {
		panic("Invalid COPY, using copy registers")
	}
	// Reset temp and to
	c.Reset(TEMP0)
	c.Reset(to)

	c.Loop(from, func() {
		c.Inc(to)
		c.Inc(TEMP0)
		c.Dec(from)
	})

	c.Shift(TEMP0, from)
}

// Compares a and b, result is a boolean in res
func (c *CommandHandler) Equals(x int, y int, res int) {
	c.Copy(x, TEMP1)
	c.SubCell(y, TEMP1)

	c.Set(res, 1)

	c.If(TEMP1, func() {
		c.Set(res, 0)
	})
}

func (c *CommandHandler) And(x int, y int, res int) {
	c.Reset(res)
	c.If(x, func() {
		c.If(y, func() {
			c.Set(res, 1)
		})
	})
}

// Boolean NOT
func (c *CommandHandler) Not(x int, res int) {
	c.Set(res, 1)
	c.If(x, func() {
		c.Dec(res)
	})

}

// Adds cell a to b, b is modified
func (c *CommandHandler) AddTo(a int, b int) {
	c.Reset(TEMP0)

	c.Loop(a, func() {
		c.Inc(TEMP0)
		c.Inc(b)
		c.Dec(a)
	})

	c.Shift(TEMP0, a)
}

// Substracts cell a to b, b is modified
func (c *CommandHandler) SubCell(a int, b int) {
	c.Reset(TEMP0)

	c.Loop(a, func() {
		c.Inc(TEMP0)
		c.Dec(b)
		c.Dec(a)
	})

	c.Shift(TEMP0, a)
}

// Multiply cell a and b
func (c *CommandHandler) MultCell(a int, b int, res int) {
	c.Copy(a, TEMP1)
	c.Reset(res)

	c.Loop(TEMP1, func() {
		c.AddTo(b, res)
		c.Dec(TEMP1)
	})
}

// Runs code inside if the current cell is truthy
// Accepts a code function using the command handler
// The resulting function must end in the same position as it begins!
func (c *CommandHandler) If(cond int, code func()) {
	c.goTo(cond)
	c.beginLoop()
	code()
	c.goTo(NIL)
	c.endLoop()
}

// Loops, ensuring that the loop begins and ends in condCell
func (c *CommandHandler) Loop(condCell int, code func()) {
	c.goTo(condCell)
	c.beginLoop()
	code()
	c.goTo(condCell)
	c.endLoop()
}

func (c *CommandHandler) Comment(comment string) {
	c.writer.Comment(comment)
}

func (c *CommandHandler) Print() {
	c.writer.Print()
}

// Core functionality

func (c *CommandHandler) beginLoop() {
	c.writer.Command(BFLoopBegin)
}

func (c *CommandHandler) endLoop() {
	c.writer.Command(BFLoopEnd)
}

// Move pointer n positions, left or right
func (c *CommandHandler) movePointer(pos int) {
	c.pointer += pos
	if pos > 0 {
		for i := 0; i < pos; i++ {
			c.writer.Command(BFIncPointer)
		}
	}

	if pos < 0 {
		for i := 0; i > pos; i-- {
			c.writer.Command(BFDecPointer)
		}
	}
}
