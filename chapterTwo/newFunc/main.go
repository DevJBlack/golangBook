package main

import (
	"fmt"
)

func main() {
	p := new(int)   // p, of type *int, points to an unmaned int variable
	fmt.Println(*p) // "0"
	*p = 2          // sets the unmaned int to 2
	fmt.Println(*p) // "2"
}

/*
	These two functions are the exact same

	func newInt() *int {
		return new(int)
	}

	func newInt() *int {
		var dummy int
		return &dummy
	}
*/
