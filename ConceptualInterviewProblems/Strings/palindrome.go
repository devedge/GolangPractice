// Determine if a string is a palindrome

package main

import (
	"fmt"
)

func main() {
	a := "asdfghgfdsa"
	fmt.Printf("Is '%s' a palindrome? %t\n", a, checkPalindrome(a))
}

// Check if the string is a palindrome by iterating from the start and
// end at the same time, and checking that the values are equal. Increment
// the start index & decrement the end index at the same time, comparing
// both values until they are at the same index.
func checkPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
