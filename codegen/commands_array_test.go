package codegen

import (
	"mindfck/bfinterpreter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetArray(t *testing.T) {
	cmd := New()

	var1 := cmd.DeclareArray("var1", 5)
	cmd.SetArray(var1, []int{1, 2, 1, 4, 5})
	cmd.PrintArray(var1)

	code := cmd.Compile()
	interpreter := bfinterpreter.New()
	interpreter.Run(code)

	assert.Equal(t, []byte{1, 2, 1, 4, 5, 0}, interpreter.Memory)
	assert.Equal(t, []byte{1, 2, 1, 4, 5}, interpreter.Output)
}
