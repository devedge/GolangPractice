package main

import (
	"fmt"
	"math"
)

func main() {
	val := 100
	fmt.Println("The first", val, "primes are:")
	fmt.Println(dynsieve(val))
}

// Find the first X primes using dynamic programming
func dynsieve(x int) []int {
	// Approximate the upper limit using an algorithm from
	// stack overflow:
	// https://stackoverflow.com/questions/7559236/dynamic-sieve-algorithms-for-prime-generation
	x_float := float64(x)      // put x into float64
	x_log := math.Log(x_float) // calculate one math.Log

	// run the approximation function. round down to an int.
	upperlimit := int(math.Floor(x_float * (x_log + math.Log(x_log))))
	fmt.Println("Using upper limit approx", upperlimit, "for", x)

	// initialize the 'sieve' & 'found primes' array
	sieve := make([]int, upperlimit)
	found := []int{}
	numprimes := 0

	// factorization code
	for i := 2; i < upperlimit; i++ {
		// if 'i' in the sieve has not been marked off as non-prime
		if sieve[i] == 0 {

			// increment the number of primes seen. 'i' is the latest one.
			numprimes++
			// add the primes to the 'found' array
			found = append(found, i)
			// if the number of 'numprimes' equals the requested amount,
			// return the array
			if numprimes == x {
				return found
			}

			// grab every other value after it, and determine its
			// multiples
			for j := i; j*i < upperlimit; j++ {
				sieve[j*i] = 1
			}
		}
	}

	return found
}

/**
  * Blindly determines the primes under a number 'x'
  for i := 2; i < x; i++ {
    // if 'i' in the sieve has not been marked off as non-prime
  	if sieve[i] == 0 {
  	  // grab every other value after it, and determine it's
  	  // multiples
  	  for j := i; j * i < x; j++ {
        sieve[j * i] = 1
      }
    }
  }
 **/

// determine primes
/*for a := 0; a < x; a++ {
	if sieve[a] == 0 {
		primes = append(primes, a)
	}
}*/
