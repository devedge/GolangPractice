// Implementation of FizzBuzz

package main

import "fmt"

// Just call fizzbuzz()
func main() {
	fizzbuzz(20)
}

// Print "FizzBuzz" if the current value is divisible by 3 & 5
// Print "Fizz" if the variable is divisible by 3
// Print "Buzz" if the variable is divisible by 5
func fizzbuzz(input int) {
	for i := 1; i < input+1; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
