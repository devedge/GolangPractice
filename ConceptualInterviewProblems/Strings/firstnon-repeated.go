// Find the first non-repeated character in a given string

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "supercalifragilisticexpialidocious"
	fmt.Println("Searching string '", str, "' for first non-repeated character:")
	idx, res := findSingleChar(str)

	if idx != -1 {
		fmt.Println("Found", res, "at index", idx)
	} else {
		fmt.Println("No non-repeated character found")
	}
}

func findSingleChar(str string) (int, string) {
	countmap := make(map[string]int)

	// use a key-value map to index every item
	for _, char := range str {
		countmap[strings.ToLower(string(char))] += 1
	}

	// iterate over the string again, and if any value in the
	// map == 1, return it
	for i, char2 := range str {
		if countmap[strings.ToLower(string(char2))] == 1 {
			return i, string(char2)
		}
	}

	return -1, ""
}
