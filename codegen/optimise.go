package codegen

import "strings"

func optimize(code string) string {
	// These are naive (and, poorly implemented) optimisations on generated brainfuck
	// It rooughly reduces the codebase ~3%, so not really useful
	retry := true
	previousOptimized := code

	for retry {
		optimized := strings.ReplaceAll(previousOptimized, "<>", "")
		optimized = strings.ReplaceAll(optimized, "><", "")
		optimized = strings.ReplaceAll(optimized, "+-", "")
		optimized = strings.ReplaceAll(optimized, "-+", "")
		optimized = strings.ReplaceAll(optimized, "[]", "")
		optimized = strings.ReplaceAll(optimized, "[-]][-]", "[-]]")

		if len(optimized) != len(previousOptimized) {
			previousOptimized = optimized
			retry = true
		} else {
			previousOptimized = optimized
			retry = false
		}

	}
	return previousOptimized
}
