package main

import (
	"fmt"
)

func main() {
	a := "asdfghgfdsa"
	fmt.Printf("Is '%s' a palindrome? %t\n", a, checkPalindrome(a))
}

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
