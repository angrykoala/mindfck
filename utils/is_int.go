package utils

import (
	"strconv"
)

func IsInt(str string) bool {
	if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		return true
	}
	return false
}
