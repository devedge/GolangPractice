// Given 2 integer arrays, determine if the seconds array
// is a rotated version of the first array

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 5, 6, 7, 8}
	b := []int{5, 6, 7, 8, 1, 2, 3}
	fmt.Println("Determine if array 'b' is the rotated version of 'a'")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("Rotated?:", isRotated(a, b))
}

func isRotated(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// Concatenate 'a' with itself, and try to find 'b' in it
	concat := append(a, a...)

	for i := 0; i < len(concat); i++ {
		if concat[i] == b[0] {
			for j := 1; j < len(b); j++ {
				i++                    // continue iterating over 'a'
				if concat[i] != b[j] { // on any difference, break
					break
				} else if j == len(b)-1 { // if all of 'b' has been traversed
					return true
				}
			}
		}
	}
	return false
}
