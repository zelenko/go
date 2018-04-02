package main

import (
	"fmt"
)

func main() {
	test([]int{1, 2, 3, 54, 6, 7, 78}, 2)
	test([]int{1, 2, 3, 54, 6, 7, 78}, 3)
	test([]int{1, 2, 3, 54, 6, 7, 78, 1}, 3)
	test([]int{1, 1, 1, 1, 1}, 5)
}

func test(n []int, t int) {

	out := 0
	for j := 0; j < len(n); j++ {

		out = out + n[j]
		if (j+1)%t == 0 {
			fmt.Printf("%d, ", out)
			out = 0
		}
	}
	if out != 0 {
		fmt.Printf("%d ", out)
	}
	fmt.Println()
}
