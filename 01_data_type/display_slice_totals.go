package main

import (
	"fmt"
)

func main() {
	quiz([]int{11, 9, 6, 11, 5}, 3)       // expected answer "26, 16"
	quiz([]int{11, 9, 6, 11, 5, 1}, 3)    // expected answer "26, 17"
	quiz([]int{11, 9, 6, 11, 5, 1, 1}, 3) // expected answer "26, 17, 1"
}

// quiz sums every n elements of arr and print.
// also print the sum of any remaining elements
func quiz(arr []int, n int) {

	out := 0

	for i := 0; i < len(arr); i++ {
		out += arr[i]

		if (i+1)%n == 0 { // because slice index starts with zero
			fmt.Printf("%d, ", out)
			out = 0
		} else if i+1 == len(arr) { // if the end of the slice
			fmt.Printf("%d ", out)
			out = 0
		}

	}
	fmt.Println("")
}
