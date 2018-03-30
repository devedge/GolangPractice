package main

import (
	"fmt"
	"strings"
)

func main() {
	val := 2373
	fmt.Println("Binary representation of", val, "is", binaryrep(val))
}

func binaryrep(a int) string {
	// Edge case
	if a == 0 {
		return "0"
	}

	// Create the result string slice
	s := []string{}

	// While 'a' is above '0'
	for a > 0 {

		// 'a & 1' is '1' if odd, and '0' if even
		if a&1 == 1 {
			s = append([]string{"1"}, s...) // prepend a string slice '1'
		} else {
			s = append([]string{"0"}, s...) // prepend a string slice '0'
		}

		a /= 2 // divide 'a' in half
	}

	// Join the slice back together as a string, and return
	return strings.Join(s, "")
}
