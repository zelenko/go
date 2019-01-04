package main

import (
	"fmt"
	"time"
)

// main is the entry point for the program.
func main() {

	messages := make(chan string)

	go work(messages)

	msg := <-messages
	fmt.Println(msg)

	///
	go work2(messages)
	fmt.Println(<-messages)
}

func work(messages1 chan<- string) {
	messages1 <- "done"
}

func work2(messages2 chan<- string) {
	time.Sleep(5 * time.Second)
	messages2 <- "done2"
}
