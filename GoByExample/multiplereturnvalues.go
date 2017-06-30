package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	// Multiple return values can be used to return both
	// result and error

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
