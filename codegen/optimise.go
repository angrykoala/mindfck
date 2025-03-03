package codegen

import "strings"

func optimize(code string) string {
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
