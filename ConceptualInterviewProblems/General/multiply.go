// Multiply 2 integers without using *

package main

import (
	"fmt"
)

func main() {
	a := 223
	b := 13
	fmt.Println("Multiplying without using * :")
	fmt.Println(a, "*", b, "is", multiply(a, b))
}

func multiply(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	if b > 0 {
		return a + multiply(a, b-1)
	} else {
		return -multiply(a, -b)
	}
}
