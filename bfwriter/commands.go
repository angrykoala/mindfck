package bfwriter

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
func (c *CommandHandler) Pop() {
	c.MovePointer(-c.lastPointerMove)
}

// Move the value of current cell to target cell (counting from current cell)
// Ends in origin
// Based on https://gist.github.com/roachhd/dce54bec8ba55fb17d3a
func (c *CommandHandler) Shift(to int) {
	c.BeginLoop()
	c.MovePointer(to)
	c.Inc()
	c.Pop()
	c.Dec()
	c.EndLoop()
}

// y[x+temp0+y-]
// temp0[y+temp0-]

// Copy current cell into to, using temp, ends in origin and resets temp
func (c *CommandHandler) Copy(to int, temp int) {
	// Reset temp and to
	c.MovePointer(temp)
	c.Reset()
	c.Pop()
	c.MovePointer(to)
	c.Reset()
	c.Pop()

	// Copy origin to to and temp
	c.BeginLoop()
	c.MovePointer(to)
	c.Inc()
	c.Pop()
	c.MovePointer(temp)
	c.Inc()
	c.Pop()
	c.Dec()
	c.EndLoop()

	// Shift temp to origin
	c.MovePointer(temp)
	c.Shift(-temp)

}

func (c *CommandHandler) Comment(comment string) {
	c.writer.Comment(comment)
}

func (c *CommandHandler) Print() {
	c.writer.Print()
}
