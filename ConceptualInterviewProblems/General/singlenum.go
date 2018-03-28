// Find the number that occurs only once in an array

package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 4, 234, 234, 6, 6, 71, 71, 3, 3, 20, 22, 22, 9, 9}
	fmt.Println("Only non-repeating number:", findNonRepeatNum(arr))
}

func findNonRepeatNum(arr []int) int {
	l := len(arr)

	// overwrite the current value with the XOR of it &
	// the previous one
	for i := 1; i < l; i++ {
		arr[i] ^= arr[i-1]
	}

	// the unmatched value translates itself to the end of
	// the array
	return arr[l-1]
}
