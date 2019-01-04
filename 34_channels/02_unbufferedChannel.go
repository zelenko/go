package main

import (
	"fmt"
)

func sendMore(s int, c chan<- int) { // send-only type chan<-, send TO channel
	c <- s
}

func main() {
	c := make(chan int) // unbuffered channel of ints
	go sendMore(20, c)  // must be goroutine because unbuffered channel blocks
	go sendMore(30, c)
	show(c)
	show(c)
	//show(c) // Extra receive FROM unbuffered channel will return error.
}

func show(c <-chan int) { //receive-only type <-chan, receive FROM  channel
	fmt.Println(<-c)
}
