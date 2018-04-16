// Find pairs in an integer array whose sum is equal to 10

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 7, 8, -2, 12, 4}
	sum := 10
	fmt.Println("Pair that adds up to", sum, "from the following array are:")
	fmt.Println(arr)
	fmt.Println(findpairs(arr, sum))
}

func findpairs(arr []int, sum int) [][]int {
	vmap := make(map[int]int, len(arr)) // init map with len(arr) for speed
	dups := make(map[int]int)
	pairs := [][]int{}

	// Add every value to the map. O(n)
	for i, val := range arr {
		vmap[val] = i
	}

	// Check if 'sum - a value = something in the map'. O(1)
	for a, chk := range arr {
		if b, ok := vmap[sum-chk]; ok {

			// Check if the index has already been found
			_, a_dup := dups[a]
			_, b_dup := dups[b]

			if !a_dup && !b_dup {
				pairs = append(pairs, [][]int{{arr[a], arr[b]}}...)
				dups[a] = 1
				dups[b] = 1
			}
		}
	}

	return pairs
}
