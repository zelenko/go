package main

import "fmt"

// main is the entry point for the program.
func main() {

	fileName := func(i string) {
		switch i {
		case "one.txt":
			one()

		case "two.txt":
			fmt.Println("I'm an two")

		default:
			fmt.Printf("Don't know: %s\n", i)
		}
	}

	name := "one.txt"
	fileName(name)

	name = "four.txt"
	fileName(name)

	name = "two.txt"
	fileName(name)

	name = "two.txt"
	findOut(name)

}

func findOut(i string) {

	switch i {

	case "one.txt":
		//fmt.Println("I'm a one")
		one()

	case "two.txt":
		fmt.Println("I'm an two")

	default:
		fmt.Printf("Don't know: %s\n", i)

	}

}

func one() {
	fmt.Println("I'm a one")

}
