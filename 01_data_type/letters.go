package main

import "fmt"

func main() {
	// Capital Letters
	for i := 65; i < 91; i++ {
		fmt.Print(string(i) + " ")
	}
	fmt.Println()

	// Small Letters
	for i := 97; i < 123; i++ {
		fmt.Print(string(i) + " ")
	}
}
