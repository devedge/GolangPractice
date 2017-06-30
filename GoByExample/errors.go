package main

import (
	"fmt"	
	"errors"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't compute the meaning of life")
	}
	return arg + 3, nil
}


type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "still can't work with it"}
	}
	return arg + 3, nil
}


func main() {

	// Test the first error function
	for _, i := range []int{7, 42} {
		if r,e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	// Test the second one
	for _,i := range[]int{7, 42} {
		if r,e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Access different fields of the Error struct
	_,e := f2(42)
	if ae,ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
