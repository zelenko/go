package main

import "fmt"

func main() {

	// Specify the size (3), otherwize will throw error
	queue := make(chan string, 3) // buffered channel
	queue <- "one"
	addMore(queue, "two")
	addMore(queue, "three")
	addMore(queue, "four")

	// No need to close loop for Println.
	// Printing values removes them from channel.
	// fmt.Println(">", <-queue)
	// fmt.Println(">", <-queue)

	// Channel must be closed before using for loop.
	close(queue)
	//fmt.Println("len:", len(queue))

	// Loop through rest of values in channel.
	// Looping through values removes them from channel.
	for elem := range queue {
		fmt.Println("==>", elem)
	}

	// No values left in the channel.
	fmt.Println(">", <-queue)
	fmt.Println(">", <-queue)
}

func addMore(c chan<- string, input string) {
	if len(c) < cap(c) {
		c <- input
	} else {
		fmt.Println("FULL CHANNEL ==> CANNOT ADD", input, "cap:", cap(c), "len:", len(c))
	}
}
