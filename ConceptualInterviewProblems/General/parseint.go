// Parse a string, and determine whether or not it's an int

package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Use ascii keypoints to determine the ints
const asc_start = 48
const asc_end = 57

func main() {
	str1 := "1424a55458"
	str2 := "7386545454"
	fmt.Println("Parse string", str1, "into an int:")
	res1, err1 := parseint(str1)
	res2, err2 := parseint(str2)

	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(res1, "- type:", reflect.TypeOf(res1))
	}

	fmt.Println("")
	fmt.Println("Parse string", str2, "into an int:")

	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(res2, "- type:", reflect.TypeOf(res2))
	}
}

func parseint(str string) (int, error) {
	val := 0

	for _, char := range str { // iterate over every character in the string
		if char >= asc_start && char <= asc_end {

			// the result is calculated from left to right
			// multiply the previous result by 10 to shift it one digit to the left
			// add the next digit from the char
			val = val*10 + (int(char) - asc_start)
		} else {
			return -1, errors.New("Value '" + string(char) + "' isn't a number")
		}
	}

	return val, nil
}
