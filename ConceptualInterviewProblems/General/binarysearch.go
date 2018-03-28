package main

import (
	"fmt"
)

// Functions were sort of found on rosetta code
func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8, 12, 14, 26, 45, 53, 76}
	val := 6
	fmt.Println("Array:", arr)
	fmt.Println("Iterative binary search for", val,
		", found at index:", binarySearchIter(val, arr))

	fmt.Println("Recursive binary search for", val,
		", found at index:", binarySearchRec(val, 0, len(arr)-1, arr))
}

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

func binarySearchRec(val, min, max int, arr []int) int {
	if max < min {
		return -1
	} // while the range is valid

	mid := (min + max) / 2 // calculate a new median

	if arr[mid] > val { // if the median value is less than 'val'
		return binarySearchRec(val, min, mid-1, arr) // median - 1 is new max
	} else if arr[mid] < val { // if the median is greater than 'val'
		return binarySearchRec(val, mid+1, max, arr) // median + 1 is new min
	}

	// recursive base case
	return mid
}
