package main

import (
	"fmt"
)

func main() {
	arr := []int{12, 14, 26, 45, 53, 76, 1, 2, 3, 5, 6, 7, 8}
	val := 7
	fmt.Println("Array:", arr)
	fmt.Println("Rotated binary search for", val, ", found at index:", binarySearchRotated(val, arr))
}

func binarySearchRotated(val int, arr []int) int {
	// First determine the start index
	foundidx := false
	startidx := 0
	for i := 1; i < len(arr); i++ { // start at 1
		if arr[i] < arr[i-1] { // if the current item is smaller than the prev
			startidx = i // start index is found
			foundidx = true
			break
		}
	}

	if !foundidx {
		return -1
	} // if the start idx wasn't found

	// split the array in two
	a := arr[startidx:len(arr)]
	b := arr[0:startidx]
	adjustedidx := 0

	// find out which split may contain the target value
	if val <= a[len(a)-1] {
		arr = a              // set 'arr' to the first slice
		adjustedidx = len(b) // offset the actual index
	} else if val >= b[0] {
		arr = b // set 'arr' to the second slice
	} else {
		return startidx // it's the target value
	}

	// binary search the target split
	return binarySearchIter(val, arr) + adjustedidx
}

// from binarysearch.go
func binarySearchIter(val int, arr []int) int {
	min := 0            // set the min search range
	max := len(arr) - 1 // set the max search range

	for min <= max { // while the minimum range
		mid := (min + max) / 2 // calculate the new median

		if arr[mid] > val { // if the median value is less than 'val'
			max = mid - 1 // set the median as the new max
		} else if arr[mid] < val { // if the median is greater than 'val'
			min = mid + 1 // set the median as the new minimum
		} else {
			return mid // otherwise, 'val' has been found
		}
	}

	return -1
}
