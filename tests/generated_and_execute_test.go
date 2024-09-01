package tests

import (
	"mindfck/bfinterpreter"
	"mindfck/codegen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedAndRun(t *testing.T) {
	cmd := codegen.New()

	var1 := cmd.Declare("var1")
	var2 := cmd.Declare("var2")

	cmd.Set(var1, 20)
	cmd.Set(var2, 50)

	cmd.Add(var1, var2, var1)
	var3 := cmd.Declare("var3")

	cmd.Copy(var1, var3)
	cmd.Print(var3)

	code := cmd.Compile()

	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, string(interpreter.Output), "F")
	assert.Equal(t, interpreter.Memory, []byte{70, 50, 0, 70, 0})
}

func TestIf(t *testing.T) {
	cmd := codegen.New()

	var1 := cmd.Declare("var1")
	var2 := cmd.Declare("var2")
	var3 := cmd.Declare("var3")

	cmd.Set(var1, 20)
	cmd.Set(var2, 0)
	cmd.Set(var3, 5)

	cmd.If(var1, func() {
		cmd.Print(var3)
	})
	cmd.If(var2, func() {
		cmd.Print(var1)
	})

	code := cmd.Compile()

	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, interpreter.Output, []byte{5})
	assert.Equal(t, interpreter.Memory, []byte{20, 0, 5, 0, 0, 0})
}

func TestBt(t *testing.T) {
	cmd := codegen.New()

	var1 := cmd.Declare("var1")
	var2 := cmd.Declare("var2")
	var3 := cmd.Declare("var3")
	res := cmd.Declare("res")

	cmd.Set(var1, 20)
	cmd.Set(var2, 5)
	cmd.Set(var3, 5)

	cmd.Gt(var1, var2, res)
	cmd.Print(res)
	cmd.Gt(var2, var1, res)
	cmd.Print(res)
	cmd.Gt(var2, var3, res)
	cmd.Print(res)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, interpreter.Output, []byte{1, 0, 0})
}

func TestDiv(t *testing.T) {
	cmd := codegen.New()

	var1 := cmd.Declare("var1")
	var2 := cmd.Declare("var2")
	var3 := cmd.Declare("var3")
	var4 := cmd.Declare("var4")

	cmd.Set(var1, 20)
	cmd.Set(var2, 5)

	cmd.Div(var1, var2, var3)
	cmd.Print(var3)
	cmd.Div(var2, var3, var4)
	cmd.Print(var4)
	cmd.Div(var3, var2, var4)
	cmd.Print(var4)

	code := cmd.Compile()

	interpreter := bfinterpreter.New()
	interpreter.Run(code)
	assert.Equal(t, interpreter.Output, []byte{4, 1, 0})
}
