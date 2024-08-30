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
	for _, v := range []byte(code) {
		switch v {
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
		// case '[':
		// interpreter.memPtr -= 1
		// case ']':
		// interpreter.memPtr -= 1
		case '.':
			var value = interpreter.Memory[interpreter.memPtr]
			interpreter.Output = append(interpreter.Output, value)
		}
	}
}
