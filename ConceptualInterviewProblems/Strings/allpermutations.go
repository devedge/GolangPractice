// Given a string, print out all possible permutations of it to stdout

package main

import (
	"fmt"
  "strings"
)

func main() {
	fmt.Println(permutation("abc"))
}

func permute(slc []string, start, length int) []string {

}

func permutation(str string) string {
  // Convert the string into a slice to make index operations on strings
  slc := strings.Split(str, "")

  // Permute the whole string, from 0 to it's length
  res := permute(slc, 0, len(slc))

  // Join the slice back into a string, and return it
  return strings.Join(res, "")

  // split into a slice
  // grab the first element
  // swap every other one

  // fmt.Println(slc)
  // p(slc, 0, len(slc) - 1)


	// if len(str) <= 1 {
	// 	return str
	// }
  //
	// p := permute(str[1:])
	// //c := string(str[0])
  //
	// fmt.Println(p)
	//fmt.Println(c)

	//for i, char := range(p) {
	//	fmt.Println(p[:i] + string(char) + p[i:])
	//}
}

// Lambda to swap elements in a slice. The i & j indices are swapped.
func swap(slc []string, i, j int) {
  tmp := slc[i]
  slc[i] = slc[j]
  slc[j] = tmp
}

func p(slc []string, x, y int) {
  if x == y {
    for i := 0; i < len(slc); i++ {
      fmt.Printf(slc[i])
    }
    fmt.Println()
  } else {
    for i := x; i < y; i++ {
      swap(slc, y, i)
      p(slc, x + 1, y)
      swap(slc, y, i)
    }
  }
}
