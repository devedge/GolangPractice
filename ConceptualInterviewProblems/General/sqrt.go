package main

import (
	"fmt"
	"math"
)

func main() {
	input := 26.0
	fmt.Println("Newtonian  : ", newtonSqrt(input))
	fmt.Println("math.Sqrt(): ", math.Sqrt(input))
}

func newtonSqrt(val float64) float64 {
	F := func(x float64) float64 {
		return math.Pow(x, 2) - val
	}

	dF := func(x float64) float64 {
		return math.Pow(x, 2)
	}

	guess := val

	for i := 0; i <= 95; i++ {
		guess = guess - (F(guess) / dF(guess))
	}

	return guess
}
