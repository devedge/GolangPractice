package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abcdefghijklmnopqrstuvwxyz"
	b := "abcdefghijklmnopqrstauvwxyz"
	c := "supercalif"
	fmt.Printf("Is '%s' only composed of unique characters?: %t\n", a, isUniquelyComposed(a))
	fmt.Printf("Is '%s' only composed of unique characters?: %t\n", b, isUniquelyComposed(b))
	fmt.Printf("Is '%s' only composed of unique characters?: %t\n", c, isUniquelyComposed(c))
}

// Is the string composed entirely of unique characters (eg., no
// repeat occurences of any characters)
func isUniquelyComposed(str string) bool {
	// make a map of unique chars
	umap := make(map[string]int)

	// for the entire range of the string
	for _, c := range str {
		val := strings.ToLower(string(c)) // lowercase for consistency

		// increment the character's appearance value
		umap[val] += 1

		// if that character has appeared more than once, it's
		// not unique
		if umap[val] > 1 {
			return false
		}
	}

	return true
}
