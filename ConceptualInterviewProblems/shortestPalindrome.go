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

	fmt.Println("Longest palindrome in 'what':", longestPalindrome("what"))
}

// Given a string, return the smallest palindrome that can be created
// with it
// func createShortest(str string) string {
//
// }


//
// func shortestPalindrome(str string) string {
//
// }
//
func longestPalindrome(str string) string {
  // Use Manchester's Algorithm to find all the palindrome centers. 'start'
  // and 'end' are the
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
  startrange := (centeridx - 1 - maxlen)/2
  endrange := maxlen

  // Return the palindrome subslice
  return str[startrange:endrange]
}


// Implementation of Manchester's Algorithm found on :
// https://articles.leetcode.com/longest-palindromic-substring-part-ii/
// And explained here: https://stackoverflow.com/a/10468753 and here:
// https://tarokuriyama.com/projects/palindrome2.php
func manchester(str string) ([]int, int, int) {
  // Closure to find minimum integer, since Go doesn't provide one
  min := func(a, b int) int { if a < b { return a }; return b }

  /*
  // The new, 'transformed' string contains:
  // - A '#' between each character (including front and back):
  //    length = length * 2 + 1
  // - The first character is "^":
  //    length = length + 1
  // - The last character is "$":
  //    length = length + 1
  newsize := len(str) * 2 + 3
  trsfm := make([]string, newsize)
*/
  size := len(str) * 2 + 3
  start, end := 1, size - 1

  // In Golang, string charaters cannot be referenced by index. However,
  // we can make a 'slice', which behaves the same way. This code takes
  // the 'str' string, and both converts it to a slice & transform it.
  transfm := make([]string, size)

  // Transform the string passed in, with '^' at the start, '$' at the
  // end, and each character interspersed with '#'.
  transfm[0] = "^" // Set the first character
  for i, char := range str {
    // The 'transfm' array is twice the size of 'str' (offset of i * 2),
    // and the index starts at 1. Since this adds 2 characters, the first
    // gets added at '1', and the second at '2'
    transfm[(i * 2) + 1] = "#"
    transfm[(i * 2) + 2] = string(char)
  }
  // Set the last two characters as '#$'
  transfm[size - 2] = "#"
  transfm[size - 1] = "$"

  // The magic: an int array that keeps track of the length of palindrome
  // centers. This is the same size as the transfm string.
  pcenters := make([]int, size)
  // Keep track of a current palindrome's center, and
  center := 0
  region := 0

  // Iterate over both 'transfm' and 'pcenters' at the same time
  for i := start; i < end; i++ {
    index_mirror := 2 * center - i

    if region > i {
      pcenters[i] = min(region - i, pcenters[index_mirror])
    } else {
      pcenters[i] = 0
    }

    // Attempt to expand a palindrome at this index, i
    for transfm[i + 1 + pcenters[i]] == transfm[i - 1 - pcenters[i]] {
      pcenters[i]++
    }

    // If the palindrome centered at i expands past region, adjust
    // center based on expanded palindrome.
    if i + pcenters[i] > region {
      center = i
      region = i + pcenters[i]
    }
  }

  return pcenters, start, end
}
