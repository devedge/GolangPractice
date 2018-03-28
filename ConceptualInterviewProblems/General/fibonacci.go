package main

import (
	"fmt"
)

// Memory map to speed up fibonacci calculation
var mem map[int]int = make(map[int]int)

func main() {
	rounds := 22
	fmt.Println("iterativeFib  :", iterativeFib(rounds))
	fmt.Println("recursiveFib  :", recursiveFib(rounds))
	fmt.Println("dynamicprogFib:", dynamicprogFib(rounds))
}

func iterativeFib(rounds int) int {
	i, a, b := 0, 0, 1

	for i < rounds {
		tmp := a + b
		a = b
		b = tmp
		i++
	}

	return a
}

func recursiveFib(rounds int) int {
	if rounds == 0 {
		return 0
	} else if rounds <= 2 {
		return 1
	} else {
		return recursiveFib(rounds-1) + recursiveFib(rounds-2)
	}
}

func dynamicprogFib(rounds int) int {
	// If the value exists in the map, return it immediately
	value, ok := mem[rounds]
	if ok {
		return value
	}

	// Otherwise, continue the regular fibonacci calculation
	if rounds == 0 {
		return 0
	} else if rounds <= 2 {
		return 1
	} else {
		return dynamicprogFib(rounds-1) + dynamicprogFib(rounds-2)
	}
}
