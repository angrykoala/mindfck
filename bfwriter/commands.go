package bfwriter

// Based on https://esolangs.org/wiki/Brainfuck_algorithms
// and https://gist.github.com/roachhd/dce54bec8ba55fb17d3a

// Registry cells
const (
	TEMP1 = 0
	NIL   = 6 // Always 0
	MAIN  = 7
)

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
	c.inc()
}

func (c *CommandHandler) Dec(cell int) {
	c.goTo(cell)
	c.dec()
}

// Resets cell to 0
func (c *CommandHandler) Reset(cell int) {
	c.loop(cell, func() {
		c.dec()
	})
}

// Move the value of current cell to target cell (counting from current cell)
func (c *CommandHandler) Shift(from int, to int) {
	c.loop(from, func() {
		c.Inc(to)
		c.Dec(from)
	})
}

// Copy current cell into to, using temp cell, ends in origin and resets temp
func (c *CommandHandler) Copy(from int, to int) {
	// Reset temp and to
	c.Reset(TEMP1)
	c.Reset(to)

	c.loop(from, func() {
		c.Inc(to)
		c.Inc(TEMP1)
		c.Dec(from)
	})

	c.Shift(TEMP1, from)
}

// Adds cell a to b
func (c *CommandHandler) AddCell(a int, b int) {
	c.Reset(TEMP1)

	c.loop(a, func() {
		c.Inc(TEMP1)
		c.Inc(b)
		c.Dec(a)
	})

	c.Shift(TEMP1, a)
}

// temp0[-]
// y[x-temp0+y-]
// temp0[y+temp0-]

// Substracts cell a to b
func (c *CommandHandler) SubCell(a int, b int) {
	c.Reset(TEMP1)

	c.loop(a, func() {
		c.Inc(TEMP1)
		c.Dec(b)
		c.Dec(a)
	})

	c.Shift(TEMP1, a)
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

func (c *CommandHandler) Comment(comment string) {
	c.writer.Comment(comment)
}

func (c *CommandHandler) Print() {
	c.writer.Print()
}

// Core functionality

// Loops, ensuring that the loop begins and ends in condCell
func (c *CommandHandler) loop(condCell int, code func()) {
	c.goTo(condCell)
	c.beginLoop()
	code()
	c.goTo(condCell)
	c.endLoop()
}

func (c *CommandHandler) inc() {
	c.writer.Command(BFInc)
}
func (c *CommandHandler) dec() {
	c.writer.Command(BFDec)
}

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
