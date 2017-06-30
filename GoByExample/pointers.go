package main

import "fmt"

// Set the passed in value to zero
func zeroval(ival int) {
	ival = 0
}

// Zero out the actual memory address of the value
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
