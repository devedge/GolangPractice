package main

import (
	"fmt"
)

func main() {
	island := [][]int{{1, 0, 1, 0, 1},
		{1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1}}

	fmt.Println("For matrix:")
	fmt.Println(island[0])
	fmt.Println(island[1])
	fmt.Println(island[2])
	fmt.Println("3 clusters expected, found:", findIslands(island), "clusters")
	fmt.Println("")

	island2 := [][]int{{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 0, 1}}

	fmt.Println("For matrix:")
	fmt.Println(island2[0])
	fmt.Println(island2[1])
	fmt.Println(island2[2])
	fmt.Println(island2[3])
	fmt.Println(island2[4])
	fmt.Println("4 clusters expected, found:", findIslands(island2), "clusters")
}

// Returns the number of clusters found
func findIslands(arr [][]int) int {
	clusters := 0

	for i, y := range arr { // for every element in the array
		for j, _ := range y {
			if arr[i][j] == 1 { // if it's an island (1) and not yet visited (-1)
				clusters++                                   // increment cluster count
				markIsland(i, j, len(arr), len(arr[0]), arr) // recursively mark this island
			}
		}
	}

	return clusters // return the number of discovered island clusters
}

// Mark an island of '1's as traveled, or '-1'
func markIsland(i, j, max_i, max_j int, arr [][]int) {
	arr[i][j] = -1 // mark node as 'visited'

	// iterate all around the current value
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {

			// check bounds to ensure they're safe to access in the array
			if x >= 0 && x < max_i &&
				y >= 0 && y < max_j && !(x == i && y == j) { // ignore current index

				// i & j are valid, check if the value is 1
				if arr[x][y] == 1 {
					markIsland(x, y, max_i, max_j, arr) // recursively mark neighbors
				}
			}
		}
	}
}
