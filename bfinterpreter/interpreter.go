package bfinterpreter

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

func (interpreter *Interpreter) Run(code string) {
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
		}
	}
}

func (interpreter *Interpreter) currentValue() byte {
	return interpreter.Memory[interpreter.memPtr]
}
