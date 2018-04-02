package main

import (
	"fmt"
)

type (
	rec string
)

const (
	testCost = "collection"
)

var (
	list   = []rec{"a", "b"}
	output = []rec{"a", "b"}
)

func main() {

	output = append(output, "c")

	// multiple ways to loop through slice
	for i := range output {
		fmt.Println("i:", output[i], i)
	}

	for k, j := range output {
		fmt.Println("k:", k, j)
	}

	m := 0
	for range output {
		fmt.Println("m:", output[m])
		m++
	}

	for n := 0; n < len(output); n++ {
		fmt.Println("n:", output[n])
	}

	fmt.Println(testCost)
}
