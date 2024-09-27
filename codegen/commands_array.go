package codegen

// // Commands for multi byte variables

// // Resets arr to 0
// func (c *CommandHandler) ResetArr(v *env.ArrayVariable) {
// 	c.iterateArr(v, func(i int) {
// 		c.ResetByte(v.Get(i))
// 	})
// }

// func (c *CommandHandler) SetArr(v *env.ArrayVariable, value []int) {
// 	assertSize(v, len(value))

// 	c.iterateArr(v, func(i int) {
// 		c.SetByte(v.Get(i), value[i])
// 	})
// }

// func assertCompatible(var1 env.Variable, var2 env.Variable) {
// 	if var1.Size() != var2.Size() {
// 		panic("Variables are not compatible")
// 	}
// }

// func assertSize(var1 env.Variable, size int) {
// 	if var1.Size() != size {
// 		panic("Variable has not correct size")
// 	}
// }
