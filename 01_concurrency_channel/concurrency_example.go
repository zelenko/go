// Concurrency example
// Five seconds given to type something, else program exits.

package main

import (
	"fmt"
	"time"
)

// channel is "reference" type, so no need to return value
func readword(cha chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	cha <- word
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- true
}

// main is the entry point for the program.
func main() {
	t := make(chan bool)
	go timeout(t)

	ch := make(chan string)
	go readword(ch)

	// select whoever brings results first
	// select the fastest channel
	select {
	case word := <-ch:
		fmt.Println("Received", word)
	case <-t:
		fmt.Println("Timeout.")
	}
}
