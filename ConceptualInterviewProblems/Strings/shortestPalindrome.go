// Note: while this question asks for the "shortest palindrome IN a string",
// it could mean "find the shortest palindrome you can construct from a string".
//
// Nevertheless, this code takes on three solutions:
// - Constructing the smallest palindrome from a string
// - Finding the shortest palindrome in a string ????
// - Find the largest palindrome in a string. This one is very hard to do
//    in O(N) time, and requires an algorithm called Manchester's Algorithm.

package main

import (
	"fmt"
)

func main() {
	p1 := "aaaaaaaaaaaaa"
	p2 := "xxxxxxxaxxx"
	p3 := "abxbaabba"
	p4 := "hannah"

	fmt.Printf("Longest palindrome in '%s': %s\n", p1, longestPalindrome(p1))
	fmt.Printf("Shortest palindrome in '%s': %s\n", p1, shortestPalindrome(p1))
	fmt.Printf("Longest palindrome in '%s': %s\n", p2, longestPalindrome(p2))
	fmt.Printf("Shortest palindrome in '%s': %s\n", p2, shortestPalindrome(p2))
	fmt.Printf("Longest palindrome in '%s': %s\n", p3, longestPalindrome(p3))
	fmt.Printf("Shortest palindrome in '%s': %s\n", p3, shortestPalindrome(p3))
	fmt.Printf("Longest palindrome in '%s': %s\n", p4, longestPalindrome(p4))
	fmt.Printf("Shortest palindrome in '%s': %s\n", p4, shortestPalindrome(p4))
}

// Given a string, return the smallest palindrome that can be created
// with it
// func createShortest(str string) string {
//
// }

// Given a string, find the shortest palindrome in it, using Manchester's
// Algorithm.
func shortestPalindrome(str string) string {
	// Use Manchester's Algorithm to find all the palindrome centers. 'start'
	// and 'end' are the valid start+end indices to 'pcenters'
	pcenters, start, end := manchester(str)
	minlen, centeridx := len(str), -1

	// Iterate over pcenters, and update the smallest value found. Technically,
	// the smallest palindrome is one letter
	for i := start; i < end; i++ {
		if pcenters[i] < minlen && pcenters[i] > 0 {
			minlen = pcenters[i]
			centeridx = i
		}
	}

	// If no minimum value was found, return "-1"
	if centeridx == -1 {
		return "-1"
	} else {
		// Define the new substring slice
		startrange := (centeridx - 1 - minlen) / 2
		endrange := minlen

		// Return the palindrome subslice
		return str[startrange:endrange]
	}
}

// Given a string, find the longest palindrome in it, using Manchester's
// Algorithm.
func longestPalindrome(str string) string {
	// Use Manchester's Algorithm to find all the palindrome centers. 'start'
	// and 'end' are the valid start+end indices to 'pcenters'
	pcenters, start, end := manchester(str)
	maxlen, centeridx := 0, 0

	// Iterate over pcenters, and update the largest values found
	for i := start; i < end; i++ {
		if pcenters[i] > maxlen {
			maxlen = pcenters[i]
			centeridx = i
		}
	}

	// Define the new substring slice
	startrange := (centeridx - 1 - maxlen) / 2
	endrange := maxlen

	// Return the palindrome subslice
	return str[startrange:endrange]
}

// Implementation of Manchester's Algorithm found on :
// https://articles.leetcode.com/longest-palindromic-substring-part-ii/
// And explained here: https://stackoverflow.com/a/10468753 and here:
// https://tarokuriyama.com/projects/palindrome2.php
func manchester(str string) ([]int, int, int) {
	// The first step is to transform the string, with '^' at the start &
	// '$' at the end (+2 characters), and each character interspersed
	// with '#', including start and end (2 * length + 1)
	size := len(str)*2 + 3
	start, end := 1, size-1

	// In Golang, string charaters cannot be referenced by index. However,
	// we can make a 'slice', which behaves the same way. This code takes
	// the 'str' string, and both converts it to a slice & transform it.
	transfm := make([]string, size)

	// Transform the string passed in, but inside the transfm slice.
	transfm[0] = "^" // Set the first character
	for i, char := range str {
		// The 'transfm' array is twice the size of 'str' (offset of i * 2),
		// and the index starts at 1. Since this adds 2 characters, the first
		// gets added at '1', and the second at '2'
		transfm[(i*2)+1] = "#"
		transfm[(i*2)+2] = string(char)
	}
	// Set the last two characters as '#$'
	transfm[size-2] = "#"
	transfm[size-1] = "$"

	// The magic: an int array that keeps track of the length of palindrome
	// centers. This is the same size as the transfm string.
	pcenters := make([]int, size)
	// Keep track of a current palindrome's center, and
	center := 0
	region := 0

	// Iterate over both 'transfm' and 'pcenters' at the same time
	for i := start; i < end; i++ {
		index_mirror := 2*center - i

		if region > i {
			pcenters[i] = min(region-i, pcenters[index_mirror])
		} else {
			pcenters[i] = 0
		}

		// Attempt to expand a palindrome at this index, i
		for transfm[i+1+pcenters[i]] == transfm[i-1-pcenters[i]] {
			pcenters[i]++
		}

		// If the palindrome centered at i expands past region, adjust
		// center based on expanded palindrome.
		if i+pcenters[i] > region {
			center = i
			region = i + pcenters[i]
		}
	}

	// Return the palindrome centers array, and the start+end indices in it
	return pcenters, start, end
}

// Closure to find minimum integer, since Go doesn't provide one
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
