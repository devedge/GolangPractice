package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "supercalifragilisticexpialidocious"
	fmt.Println("Reversing string", str)
	fmt.Println("Iteratively:", reverseIter(str))
	fmt.Println("Recursively:", reverseRecur(str))
}

func reverseIter(s string) string {
	res := make([]rune, len(s))
	max := len(s)

	for _, char := range s {
		if char != utf8.RuneError {
			max--
			res[max] = char
		}
	}

	return string(res[0:])
}

func reverseRecur(s string) string {
	if len(s) <= 0 { // base case
		return s
	}

	// return 'reverse' of rest of string + first character
	return reverseRecur(s[1:len(s)]) + string(s[0])
}
