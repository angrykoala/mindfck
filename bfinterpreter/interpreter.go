package bfinterpreter

import "fmt"

type Interpreter struct {
	Memory []byte
	memPtr int
	Output []byte
}

func New() *Interpreter {
	return &Interpreter{
		Memory: []byte{0},
	}
}

func (interpreter *Interpreter) RunWithInput(code string, input []byte) {
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			interpreter.Memory[interpreter.memPtr] += 1
		case '-':
			interpreter.Memory[interpreter.memPtr] -= 1
		case '>':
			interpreter.memPtr += 1
			if len(interpreter.Memory) <= interpreter.memPtr {
				interpreter.Memory = append(interpreter.Memory, 0)
			}
		case '<':
			interpreter.memPtr -= 1
			if interpreter.memPtr < 0 {
				panic("Kaboom")
			}
		case '[':
			if interpreter.currentValue() == 0 {
				var skips int = 1
				for skips != 0 {
					i = i + 1
					if code[i] == '[' {
						skips = skips + 1
					}
					if code[i] == ']' {
						skips = skips - 1
					}
				}
			}

		case ']':
			if interpreter.currentValue() != 0 {
				var skips int = 1
				for skips != 0 {
					i = i - 1
					if code[i] == ']' {
						skips = skips + 1
					}
					if code[i] == '[' {
						skips = skips - 1
					}
				}
			}
		case '.':
			var value = interpreter.Memory[interpreter.memPtr]
			interpreter.Output = append(interpreter.Output, value)
		case ',':
			inByte := input[0]
			input = input[1:]
			interpreter.Memory[interpreter.memPtr] = inByte
		case '#':
			interpreter.Debug()
			// return
		}
	}
}

func (interpreter *Interpreter) Run(code string) {
	interpreter.RunWithInput(code, []byte{})
}

func (interpreter *Interpreter) Debug() {
	fmt.Println("Memory:", interpreter.Memory)
	fmt.Println("Output:", interpreter.Output)
	fmt.Println("Pointer:", interpreter.memPtr)
}

func (interpreter *Interpreter) currentValue() byte {
	return interpreter.Memory[interpreter.memPtr]
}
