package codegen

import (
	"mindfck/bfinterpreter"
	"mindfck/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncIntSingleByte(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.INT)
	cmd.IncInt(var1)
	cmd.PrintInt(var1)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{0, 1, 0, 0, 0, 0, 0}, interpreter.Memory)
	assert.Equal(t, []byte{0, 1}, interpreter.Output)
}

func TestIncInt(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.INT)
	cmd.SetByte(var1.GetByte(0), 2)
	cmd.SetByte(var1.GetByte(1), 254)
	cmd.IncInt(var1)
	cmd.IncInt(var1)
	cmd.IncInt(var1)
	cmd.PrintInt(var1)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{3, 1}, interpreter.Output)
}

func TestDecInt(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.INT)
	cmd.SetByte(var1.GetByte(0), 2)
	cmd.SetByte(var1.GetByte(1), 1)
	cmd.DecInt(var1)
	cmd.DecInt(var1)
	cmd.DecInt(var1)
	cmd.PrintInt(var1)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{1, 254}, interpreter.Output)
}

func TestIsZeroInt(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.INT)
	res1 := cmd.Declare("res1", env.BYTE)
	res2 := cmd.Declare("res2", env.BYTE)
	cmd.isZeroInt(var1, res1)
	cmd.SetByte(var1.GetByte(0), 2)
	cmd.SetByte(var1.GetByte(1), 1)
	cmd.isZeroInt(var1, res2)

	cmd.PrintByte(res1)
	cmd.PrintByte(res2)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{1, 0}, interpreter.Output)
}
