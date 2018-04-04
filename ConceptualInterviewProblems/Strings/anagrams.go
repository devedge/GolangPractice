package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "rail safety"
	b := "fairy tales"
	fmt.Println("Is '", a, "' an anagram of '", b, "'?: ", isAnagram(a, b))
}

func isAnagram(a, b string) bool {
	charmap := make(map[string]int)
	a = strings.ToLower(a) // make strings lowercase
	b = strings.ToLower(b)

	// increment every character in 'a'
	for _, a_char := range a {
		charmap[string(a_char)] += 1
	}

	// decrement every character in 'b'
	for _, b_char := range b {
		charmap[string(b_char)] -= 1
	}

	// for every value in the charmap,
	for _, val := range charmap {
		if val != 0 {
			return false
		}
	}

	return true
}
