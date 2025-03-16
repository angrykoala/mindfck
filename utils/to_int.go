package utils

import (
	"strconv"
)

// Converts a string to int ("15" -> 15)
func ToInt(str string) int {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("unexpected string in parseInt")
	}

	return int(number)

}
