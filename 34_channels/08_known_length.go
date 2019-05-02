package main

import "fmt"

func main() {
	n := 10
	var queue = make(chan string, n) // n is recommended, but not required
	for i := 0; i < n; i++ {
		go addMore(queue, fmt.Sprintf("step - %d", i))
	}

	for i := 0; i < n; i++ {
		fmt.Println(<-queue)
	}
}

func addMore(c chan<- string, input string) {
	c <- input
}
