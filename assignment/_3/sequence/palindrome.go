package sequence

import (
	"strings"
)

func Palindromic(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, ",", "")

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}
