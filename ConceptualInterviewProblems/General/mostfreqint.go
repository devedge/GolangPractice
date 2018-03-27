// Find the most frequent integer in an array

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 6, 1, 7, 8, 2, 3, 4}
	fmt.Println("Most frequent repeating number:", getMostFreq(arr))
}

func getMostFreq(arr []int) int {
	// map of number --> count
	countmap := make(map[int]int)
	countmax := 0  // the highest count
	countnum := -1 // the number with highest count

	// iterate over the entire array
	for _, val := range arr { // := 0; i < len(arr); i++ {
		countmap[val] += 1 // increment the number's count

		if countmap[val] > countmax {
			countmax = countmap[val]
			countnum = val
		}
	}

	return countnum
}
