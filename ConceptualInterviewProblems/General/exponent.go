package main

import (
	"fmt"
)

func main() {
	base := 3
	power := -5
	fmt.Println("expSimple:", base, "^", -power, "is", expSimple(base, -power))
	fmt.Println("expOpt   :", base, "^", power, "is", expOpt(float64(base), power))
}

// O(n), no support for negative exponents
func expSimple(a int, b int) int {
	if b == 0 {
		return 1
	} else if b%2 == 0 {
		return expSimple(a, b/2) * expSimple(a, b/2)
	} else {
		return a * expSimple(a, b/2) * expSimple(a, b/2)
	}
}

// O(log(n))
func expOpt(a float64, b int) float64 {
	if b == 0 {
		return 1
	}
	if a == 0 {
		return 0
	}

	tmp := expOpt(a, b/2)

	if b%2 == 0 {
		return tmp * tmp
	} else {
		// Check for negative exponent
		if b > 0 {
			return a * tmp * tmp
		} else {
			return (tmp * tmp) / a
		}
	}
}
