// Quick Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 10; i++ {
		fmt.Print(95+rand.Intn(6), " ")
	}

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func sort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sort(a[:left])
	sort(a[left+1:])

	return a
}
