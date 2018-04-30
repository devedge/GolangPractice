// Given a string, print out all possible permutations of it to stdout

package main

import (
	"fmt"
	"strings"
)

func main() {
	permutation("abc")
}

// Wrapper function to turn a string into a slice
func permutation(str string) {
	// Convert the string into a slice to make index operations on strings
	slc := strings.Split(str, "")

	// Permute the whole string, from 0 to it's length
	permute(slc, 0, len(slc)-1)
}

// Recursive permutation function.
// Functionality explained here: https://stackoverflow.com/a/16989774
func permute(slc []string, start, end int) {
	if start == end {
		// Recursive base case
		// We've reached the end of a full permutation, so print it
		fmt.Println(strings.Join(slc, ""))
	} else {
		for index := start; index <= end; index++ {
			// swap the next charater
			swap(slc, start, index)

			// permute the rest of the string
			permute(slc, start+1, end)

			// undo the swap
			swap(slc, start, index)
		}
	}
}

// Function to swap the i & j elements in a slice
func swap(input []string, i, j int) {
	tmp := input[i]
	input[i] = input[j]
	input[j] = tmp
}
