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
	assert.Equal(t, interpreter.Memory, []byte{0, 70, 50, 0, 70, 0})
}
