package bfwriter

// Based on https://esolangs.org/wiki/Brainfuck_algorithms
// and https://gist.github.com/roachhd/dce54bec8ba55fb17d3a

type CommandHandler struct {
	writer          *writer
	lastPointerMove int
}

func NewCommandHandler() *CommandHandler {
	cmd := CommandHandler{
		writer: NewWriter(),
	}

	return &cmd
}

func (c *CommandHandler) Inc() {
	c.writer.Command(BFInc)
}
func (c *CommandHandler) Dec() {
	c.writer.Command(BFDec)
}

func (c *CommandHandler) Add(count int) {
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

func (c *CommandHandler) BeginLoop() {
	c.writer.Command(BFLoopBegin)
}

func (c *CommandHandler) EndLoop() {
	c.writer.Command(BFLoopEnd)
}

// Resets cell to 0
// Based on https://esolangs.org/wiki/Brainfuck_algorithms#x_=_0
func (c *CommandHandler) Reset() {
	c.BeginLoop()
	c.Dec()
	c.EndLoop()

}

// Move pointer n positions, left or right
func (c *CommandHandler) MovePointer(pos int) {
	c.lastPointerMove = pos
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

// Move to previous pointer
func (c *CommandHandler) pop() {
	if c.lastPointerMove == 0 {
		panic("Cannot Pop, no last Pointer moved")
	}
	c.MovePointer(-c.lastPointerMove)
	c.lastPointerMove = 0
}

// Move the value of current cell to target cell (counting from current cell)
// Ends in origin
// Based on https://gist.github.com/roachhd/dce54bec8ba55fb17d3a
func (c *CommandHandler) Shift(to int) {
	c.BeginLoop()
	c.MovePointer(to)
	c.Inc()
	c.pop()
	c.Dec()
	c.EndLoop()
}

// Copy current cell into to, using temp cell, ends in origin and resets temp
func (c *CommandHandler) Copy(to int, temp int) {
	// Reset temp and to
	c.MovePointer(temp)
	c.Reset()
	c.pop()
	c.MovePointer(to)
	c.Reset()
	c.pop()

	// Copy origin to to and temp
	c.BeginLoop()
	c.MovePointer(to)
	c.Inc()
	c.pop()
	c.MovePointer(temp)
	c.Inc()
	c.pop()
	c.Dec()
	c.EndLoop()

	// Shift temp to origin
	c.MovePointer(temp)
	c.Shift(-temp)
}

// Adds cell to current cell, using temp
func (c *CommandHandler) AddCell(y int, temp int) {
	// Reset temp and to
	c.MovePointer(temp)
	c.Reset()
	c.pop()

	c.MovePointer(y)
	// Increment x and temp with y
	c.BeginLoop() // Loop begins at y
	c.pop()
	c.Inc()
	c.MovePointer(temp)
	c.Inc()
	c.pop()
	c.MovePointer(y)
	c.Dec()
	c.EndLoop()

	// Move temp to y
	c.pop()
	c.MovePointer(temp)
	c.Shift(y - temp)
}

// Rus code inside if the current cell is truthy
// Accepts a code function using the command handler
// The resulting function must end in the same position as it begins!
func (c *CommandHandler) If(temp int, code func()) {
	c.MovePointer(temp)
	c.Reset()
	c.pop()

	c.BeginLoop()
	code()
	c.MovePointer(temp)
	c.EndLoop()
	c.pop()
}

func (c *CommandHandler) Comment(comment string) {
	c.writer.Comment(comment)
}

func (c *CommandHandler) Print() {
	c.writer.Print()
}
