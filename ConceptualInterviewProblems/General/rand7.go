// Implement rand7() only using a rand5() function. The
// challenge is creating an even distribution from 1-7
// from a function that can only provide an even
// distribution from 1-5.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Random number from 1-7 using rand5():", rand7())

	fmt.Println("Testing randomness spread from 1-7 over 100,000 iterations...")
	countarr := [7]int{}
	for i := 0; i < 100000; i++ {
		countarr[rand7()-1] += 1
	}
	fmt.Println(countarr)

	fmt.Println("")
	fmt.Println("Using the second, complex rand7:", rand7mod())
	fmt.Println("Testing randomness spread from 0-6 over 100,000 iterations...")
	countarr2 := [7]int{}
	for i := 0; i < 100000; i++ {
		countarr2[rand7mod()] += 1
	}
	fmt.Println(countarr2)
}

func rand7() int {
	// A 5x5 array with numbers from 0-7 that get picked
	// by two rand5() calls. The 0s are included instead
	// of other variables to ensure an even spread.
	// The zeros are also interleaved instead of being
	// all appended to the end, to (possibly?) prevent
	// a bias from all of them clumped at the end.
	res := 0
	pick := [5][5]int{{0, 1, 2, 3, 4},
		{5, 6, 7, 0, 1},
		{2, 3, 4, 5, 6},
		{7, 0, 1, 2, 3},
		{4, 5, 6, 7, 0}}

	// Filter only values from 1-7
	for res == 0 {
		res = pick[rand5()-1][rand5()-1]
	}

	return res
}

// rand7 implementation found online (possibly stackoverflow).
// I'm not sure how it works, but its distribution is worse
// than the one above.
func rand7mod() int {
	// covers all multiples of 5 between 5 & 25, plus a value
	// from 1 to 5 to reach all values in between
	i := 5*rand5() + rand5()
	max := 25 // the maximum possible value

	for i < max%7 { // if 'i' is less than 4 (initially)
		i *= 5       // i = i * 5
		i += rand5() // add a value from 1-5 to 'i'
		max %= 7     // max = max % 7
		max *= 5     // max = max * 5
	}

	return i % 7
}

// Returns a value between 1-5
func rand5() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(5) + 1
}
