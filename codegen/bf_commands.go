package codegen

func (c *CommandHandler) DebugBreak() {
	c.writer.command(BFDebug)
}

func (c *CommandHandler) Comment(comment string) {
	c.writer.comment(comment)
}

func (c *CommandHandler) inc() {
	c.writer.command(BFInc)
}

func (c *CommandHandler) dec() {
	c.writer.command(BFDec)
}

func (c *CommandHandler) moveR() {
	c.writer.command(BFIncPointer)
}

func (c *CommandHandler) moveL() {
	c.writer.command(BFDecPointer)
}

func (c *CommandHandler) in() {
	c.writer.command(BFIn)
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
