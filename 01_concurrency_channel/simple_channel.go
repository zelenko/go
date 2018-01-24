package main

import "fmt"

func main() {

	// Specify the size (3), otherwize will throw error
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"

	// No need to close loop for Println.
	// Printing values removes them from channel.
	fmt.Println(">", <-queue)
	fmt.Println(">", <-queue)

	// Channel must be closed before using for loop.
	close(queue)

	// Loop through rest of values in channel.
	// Looping through values removes them from channel.
	for elem := range queue {
		fmt.Println("==>", elem)
	}

	// No values left in the channel.
	fmt.Println(">", <-queue)
}
