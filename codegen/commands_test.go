package codegen

import (
	"mindfck/bfinterpreter"
	"mindfck/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedAndRun(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.BYTE)
	var2 := cmd.Declare("var2", env.BYTE)

	cmd.SetByte(var1, 20)
	cmd.SetByte(var2, 50)

	cmd.Add(var1, var2, var1)
	var3 := cmd.Declare("var3", env.BYTE)

	cmd.CopyByte(var1, var3)
	cmd.PrintByte(var3)

	code := cmd.Compile()

	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, "F", string(interpreter.Output))
	assert.Equal(t, []byte{70, 50, 70, 0}, interpreter.Memory)
}

func TestIf(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.BYTE)
	var2 := cmd.Declare("var2", env.BYTE)
	var3 := cmd.Declare("var3", env.BYTE)

	cmd.SetByte(var1, 20)
	cmd.SetByte(var2, 0)
	cmd.SetByte(var3, 5)

	cmd.If(var1, func() {
		cmd.PrintByte(var3)
	})
	cmd.If(var2, func() {
		cmd.PrintByte(var1)
	})

	code := cmd.Compile()

	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{5}, interpreter.Output)
	assert.Equal(t, []byte{20, 0, 5, 0, 0}, interpreter.Memory)
}

func TestBt(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.BYTE)
	var2 := cmd.Declare("var2", env.BYTE)
	var3 := cmd.Declare("var3", env.BYTE)
	res := cmd.Declare("res", env.BYTE)

	cmd.SetByte(var1, 20)
	cmd.SetByte(var2, 5)
	cmd.SetByte(var3, 5)

	cmd.Gt(var1, var2, res)
	cmd.PrintByte(res)
	cmd.Gt(var2, var1, res)
	cmd.PrintByte(res)
	cmd.Gt(var2, var3, res)
	cmd.PrintByte(res)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, interpreter.Output, []byte{1, 0, 0})
}

func TestDiv(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.BYTE)
	var2 := cmd.Declare("var2", env.BYTE)
	var3 := cmd.Declare("var3", env.BYTE)
	var4 := cmd.Declare("var4", env.BYTE)

	cmd.SetByte(var1, 20)
	cmd.SetByte(var2, 5)

	cmd.Div(var1, var2, var3)
	cmd.PrintByte(var3)
	cmd.Div(var2, var3, var4)
	cmd.PrintByte(var4)
	cmd.Div(var3, var2, var4)
	cmd.PrintByte(var4)

	code := cmd.Compile()

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, interpreter.Output, []byte{4, 1, 0})
}
