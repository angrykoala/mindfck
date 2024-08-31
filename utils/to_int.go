package utils

import (
	"strconv"
)

func ToInt(str string) int {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("unexpected string in parseInt")
	}

	return int(number)

}
