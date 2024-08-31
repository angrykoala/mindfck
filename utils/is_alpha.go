package utils

func IsAlphaNumeric(str string) bool {
	for _, c := range str {
		if !isAlphaNumericChar(c) {
			return false
		}
	}

	return true
}

func IsAlphabeticChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isAlphaNumericChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')

}
