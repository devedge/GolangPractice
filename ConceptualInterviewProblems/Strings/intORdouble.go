// Parse an input string, and determine whether it is an int or double

package main

import (
	"fmt"
)

// Valid ascii keypoints for ints or doubles.
// An int can contain a '-', any value for 0-9, and a '.' ONLY if it's followed by zeros
// A double can contain any of those above values
const int_start = 48
const int_end = 57
const asc_dash = 45
const asc_dot = 46

func main() {
	valDouble1 := "3499.020"
	valDouble2 := "-23336.78"
	valDoubleInvalid := "-.005f"
	valDoubleInvalid2 := "3.00.2"
	valInt1 := "24"
	valInt2 := "-24"
	valInt3 := "-24.00"
	valIntInvalid := "-24v00"
	fmt.Println("Value of", valDouble1, "=", parseValue(valDouble1))
	fmt.Println("Value of", valDouble2, "=", parseValue(valDouble2))
	fmt.Println("Value of", valDoubleInvalid, "=", parseValue(valDoubleInvalid))
	fmt.Println("Value of", valDoubleInvalid2, "=", parseValue(valDoubleInvalid2))
	fmt.Println("Value of", valInt1, "=", parseValue(valInt1))
	fmt.Println("Value of", valInt2, "=", parseValue(valInt2))
	fmt.Println("Value of", valInt3, "=", parseValue(valInt3))
	fmt.Println("Value of", valIntInvalid, "=", parseValue(valIntInvalid))
}

// Parses the input string character by character, and returns a string
// value of either 'int', 'double', or 'invalid'
func parseValue(str string) string {
	strVal := "int"  // default value of int
	decimal := false // no decimal point assumed

	for index, char := range str {
		// Check if the first character is a minus sign. If the dash (-)
		// is in any other position, it's invalid.
		if char == asc_dash {
			if index == 0 {
				continue // first character is '-', skip to next loop
			} else {
				return "invalid"
			}
		}

		// If this is the first decimal, make a note of it. There can
		// only be one decimal point, and if there are any numbers other
		// than zero after it, the number is a double (or float, in Go terms)
		if char == asc_dot {
			if !decimal {
				decimal = true
				continue // character was a '.', skip to next loop
			} else {
				return "invalid"
			}
		}

		// Finally, check if the value is a valid number. If not, return
		// "invalid". If there is a decimal, and the number is not equal to
		// zero (int_start), then this is a double
		if char >= int_start && char <= int_end {
			if decimal && char != int_start {
				strVal = "double"
			}
		} else {
			return "invalid"
		}
	}

	return strVal
}
