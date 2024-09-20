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

func TestSubInt(t *testing.T) {
	cmd := New()

	a := cmd.Declare("a", env.INT)
	b := cmd.Declare("b", env.INT)
	c := cmd.Declare("c", env.INT)
	cmd.SetInt(a, 300)
	cmd.SetInt(b, 250)
	cmd.SubInt(a, b, c)
	cmd.PrintInt(a)
	cmd.PrintInt(b)
	cmd.PrintInt(c)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{1, 44, 0, 250, 0, 50}, interpreter.Output)
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

func TestGTInt(t *testing.T) {
	cmd := New()

	var1 := cmd.Declare("var1", env.INT)
	var2 := cmd.Declare("var2", env.INT)
	res := cmd.Declare("res", env.BYTE)
	cmd.SetInt(var1, 10)
	cmd.SetInt(var2, 20)
	cmd.GtInt(var2, var1, res)
	cmd.PrintByte(res)
	cmd.GtInt(var1, var2, res)
	cmd.PrintByte(res)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{1, 0}, interpreter.Output)
}
