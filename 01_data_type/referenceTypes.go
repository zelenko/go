/*
The following are "reference" types:
	slices
	maps
	channels
	pointers
	functions


	otherwise it is Primitive
*/

package main

import (
	"fmt"
)

// main is the entry point for the program.
func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"

	// show results before making changes via function
	fmt.Println("emp:", s)

	test(s) // update slice; s is pointer

	fmt.Println("apd:", s) // show results
}

// slice is "reference" type, so no need to return value
func test(s []string) { // Passing by pointer
	s[2] = "c"
}
