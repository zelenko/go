package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 6, 7}

	check(numbers, 2)
}

func check(number []int, total int) {
	j := 0
	for i := range number {
		j = number[i] + j
		//fmt.Print(number[i], total, "\n")
		if (i+1)%total == 0 {
			fmt.Print(j, ",")
			j = 0
		}
	}
}
