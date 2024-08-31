package utils

func GetFirstRune(str string) rune {
	for _, c := range str {
		return c
	}

	panic("Empty string in GetFirstRune")
}
