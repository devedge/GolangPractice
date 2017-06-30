package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Direct call
	f("direct")

	// Call as a goroutine
	go f("goroutine")

	// Start a goroutine as an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going") // run immediately

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
