package main

import (
	"fmt"
	"os"
)

func main() {

	// number of  arguments
	l := len(os.Args)
	
	// display arguments using switch
	switch l {
	case 1:
		fmt.Println("No arguments")

	case 2:
		fmt.Print(1, ": ", os.Args[1], "\n")

	default:
		fmt.Printf("Number of arguments: %v\n", l)
	}

	// display arguments using loop
	if len(os.Args) > 1 {

		// loop through all arguments
		for i := 1; i < len(os.Args); i++ {
			fmt.Print(i, ": ", os.Args[i], "\n")
		}

	} else {
		fmt.Println("No arguments")
	}
}
