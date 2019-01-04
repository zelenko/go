package main

import (
	"fmt"
	"strconv"
)

func pinger(c chan string) {
	for i := 0; i < 3; i++ {
		c <- "ping " + strconv.Itoa(i)
	}
	close(c)
}

func main() {
	var c = make(chan string)

	go pinger(c)

	// Channel must be closed before using for loop.
	for msg := range c {
		fmt.Println(msg)
	}
}
