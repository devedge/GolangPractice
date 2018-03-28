// Find the common elements between two arrays

package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 4, 5, 6}
	b := []int{4, 6, 7, 8, 9}
	fmt.Println("The common elements in arrays a & b are:")
	fmt.Println(findcommon(a, b))
}

func findcommon(a []int, b []int) []int {
	// Make a map and a blank array of common values
	cmap := make(map[int]int)
	common := []int{}

	// Iterate over the first array, and add all of
	// them to the map
	for i := 0; i < len(a); i++ {
		cmap[a[i]] = 1
	}

	// Iterate over the second array, and if the value exists
	// in the map, add to the 'common' array
	for j := 0; j < len(b); j++ {
		_, ok := cmap[b[j]]
		if ok {
			common = append(common, b[j])
		}
	}

	return common
}
