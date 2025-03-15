package codegen

import (
	"fmt"
	"mindfck/env"
)

// Commands for multi byte variables

func (c *CommandHandler) SetArray(v env.Variable, value []int) {
	assertArrayOfSize(v, len(value))
	c.Reset(v)

	c.iterateArray(v, func(b env.Variable, i int, total_index int) {
		c.SetByte(b, value[i])
	})
}

func (c *CommandHandler) ReadIndex(v env.Variable, index env.Variable, to env.Variable) {
	assertArray(v)
	assertByte(to)
	assertByte(index)

	// "Zipper Alg"
	head := c.initializeArrayHead(v, index)
	c.goTo(head.Buffer)

	// Move head to index

	c.shift(1) // Go to index
	c.beginLoop()
	c.shift(-1) // Go to buffer
	c.moveHeadRight()
	c.shift(1) // Move to index
	c.endLoop()

	// Copy array[0] to data
	c.shift(3)        // Move to arr[0]
	c.rawCopy(-1, -4) // Copy to data, using Buffer as buffer

	// Move head back
	c.shift(-2) // Go to returnIndex
	c.beginLoop()
	c.shift(-2) // Go to buffer
	c.moveHeadLeft()
	c.shift(2) // Go to new returnIndex
	c.endLoop()

	c.shift(-2) // Return to original position

	c.CopyByte(head.Data, to)
	head.resetHead(c)

}

func (c *CommandHandler) WriteIndex(v env.Variable, index env.Variable, value env.Variable) {
	assertArray(v)
	assertByte(value)
	assertByte(index)

	// "Zipper Alg"
	head := c.initializeArrayHead(v, index)
	c.Copy(value, head.Data)
	c.goTo(head.Buffer)

	// Move head to index

	c.shift(1) // Go to index
	c.beginLoop()
	c.shift(-1) // Go to buffer
	c.moveHeadRight()
	c.shift(1) // Move to index
	c.endLoop()

	// Copy array[0] to data
	c.shift(2)                // Move to Data
	c.rawCopyWithReset(1, -3) // Copy to arr[0], using Buffer as buffer
	c.rawResetByte()          // Reset Data byte

	// Move head back
	c.shift(-1) // Go to returnIndex
	c.beginLoop()
	c.shift(-2) // Go to buffer
	c.moveHeadLeft()
	c.shift(2) // Go to new returnIndex
	c.endLoop()

	c.shift(-2) // Return to original position
	head.resetHead(c)
}

// Iterate array data, skipping the head
func (c *CommandHandler) iterateArray(v env.Variable, cb func(b env.Variable, i int, total_index int)) {
	c.iterateBytes(v, func(b env.Variable, i int) {
		if i >= env.ARRAY_HEAD_SIZE {
			cb(b, i-env.ARRAY_HEAD_SIZE, i)
		}
	})
}

func (c *CommandHandler) PrintArray(v env.Variable) {
	assertArray(v)

	temp := c.env.DeclareAnonByte()
	defer c.Release(temp)
	c.SetByte(temp, 91) // [
	c.PrintByte(temp)
	c.SetByte(temp, 44) // ,
	c.iterateArray(v, func(b env.Variable, i int, total_index int) {
		c.PrintByte(b)
		if total_index < v.Size()-1 {
			c.PrintByte(temp)
		}
	})
	c.SetByte(temp, 93) // ]
	c.PrintByte(temp)
}

func assertArrayOfSize(v env.Variable, size int) {
	assertArray(v)
	assertSize(v, size+env.ARRAY_HEAD_SIZE)
}

func assertSize(var1 env.Variable, size int) {
	if var1.Size() != size {
		panic("Variable has not correct size")
	}
}

func assertArray(v env.Variable) {
	if v.Type() != env.ARRAY {
		panic(fmt.Sprintf("invalid type %s, %s expected", v.Type(), env.ARRAY))
	}
}

type arrayHead struct {
	Buffer      env.Variable
	Index       env.Variable
	ReturnIndex env.Variable
	Data        env.Variable
}

func (head *arrayHead) resetHead(cmd *CommandHandler) {
	cmd.Reset(head.Data)
}

func (c *CommandHandler) initializeArrayHead(arr env.Variable, index env.Variable) *arrayHead {

	head := &arrayHead{
		Buffer:      arr.GetByte(0),
		Index:       arr.GetByte(1),
		ReturnIndex: arr.GetByte(2),
		Data:        arr.GetByte(3),
	}

	c.CopyByte(index, head.Index)
	return head
}

// Move head right 1 position, begins at buffer
// Ends in new buffer
func (c *CommandHandler) moveHeadRight() {
	c.shift(4)                // move to array[0]
	c.rawMoveByte(-4)         // Move first item of array to Buffer
	c.shift(-1)               // move to Data
	c.rawMoveByteWithReset(1) // Move dataP
	c.shift(-1)
	c.increment()             // Increment returnIndexP
	c.rawMoveByteWithReset(1) // Move returnIndexP
	c.shift(-1)
	c.decrement()             // Decrement indexP
	c.rawMoveByteWithReset(1) // Move indexP
	c.rawResetByte()          // Reset indexP
}

// Move head left 1 position, begins at buffer
// Ends in new buffer
func (c *CommandHandler) moveHeadLeft() {
	c.shift(2) // Move to returnIndex
	c.decrement()
	c.rawMoveByte(-1) // Move returnIndex
	c.shift(1)        // Move to data
	c.rawMoveByte(-1) // Move Data

	c.shift(-4)      // Go to arr[-1], this is the new buffer
	c.rawMoveByte(4) // Move arr[-1] to arr[0]
}

// Move current position to relative index
// Ends in same position
func (c *CommandHandler) rawMoveByte(relativeIndex int) {
	c.beginLoop()
	c.shift(relativeIndex)
	c.increment()
	c.shift(-relativeIndex)
	c.decrement()
	c.writer.command(BFLoopEnd)
}

// Move current position to relative index
// Ends in same position
// Resets target byte first
func (c *CommandHandler) rawMoveByteWithReset(relativeIndex int) {
	c.shift(relativeIndex)
	c.rawResetByte()
	c.shift(-relativeIndex)
	c.rawMoveByte(relativeIndex)
}

// Move current position to relative index
// Ends in same position
func (c *CommandHandler) rawCopy(toRelativeIndex int, bufferRelativeIndex int) {
	c.beginLoop()
	c.shift(toRelativeIndex)
	c.increment()
	c.shift(-toRelativeIndex + bufferRelativeIndex)
	c.increment()
	c.shift(-bufferRelativeIndex)
	c.decrement()
	c.writer.command(BFLoopEnd)

	c.shift(bufferRelativeIndex)
	c.rawMoveByte(-bufferRelativeIndex)
	c.shift(-bufferRelativeIndex)
}

// Move current position to relative index
// Ends in same position
// Resets target byte first
func (c *CommandHandler) rawCopyWithReset(toRelativeIndex int, bufferRelativeIndex int) {
	c.shift(toRelativeIndex)
	c.rawResetByte()
	c.shift(-toRelativeIndex)
	c.rawCopy(toRelativeIndex, bufferRelativeIndex)
}

func (c *CommandHandler) rawResetByte() {
	c.beginLoop()
	c.decrement()
	c.endLoop()
}
