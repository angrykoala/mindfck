package codegen

import (
	"mindfck/bfinterpreter"
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
