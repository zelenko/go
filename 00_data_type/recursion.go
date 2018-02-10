// Examples of Recursion in Go

package main

import (
	"fmt"
	"strconv"
)

// factorial recursion
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

// countdown return decremental list of numbers
func countdown(i int, s []string) (int, []string) {
	if i < 1 {
		return i, s
	}

	return countdown(i-1, append(s, strconv.Itoa(i)))
}

// main is the entry point for the program.
func main() {
	var i = 15
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	_, s := countdown(i, []string{})
	fmt.Println("Countdown:", s)
}
