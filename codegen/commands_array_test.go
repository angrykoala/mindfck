package codegen

import (
	"mindfck/bfinterpreter"
	"mindfck/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetArray(t *testing.T) {
	cmd := New()

	var1 := cmd.DeclareArray("var1", 5)
	cmd.SetArray(var1, []int{49, 50, 51, 52, 53})
	cmd.PrintArray(var1)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, 10, len(interpreter.Memory))
	assert.Equal(t, []byte("[1,2,3,4,5]"), interpreter.Output)
}

func TestReadIndex(t *testing.T) {

	cmd := New()

	var1 := cmd.DeclareArray("var1", 5)
	var2 := cmd.Declare("var2", env.BYTE)
	index := cmd.Declare("index", env.BYTE)
	cmd.SetArray(var1, []int{49, 50, 51, 52, 53})
	cmd.SetByte(index, 2)
	cmd.ReadIndex(var1, index, var2)
	cmd.PrintByte(var2)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{51}, interpreter.Output)
}
