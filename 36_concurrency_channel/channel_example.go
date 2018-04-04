package main

import "fmt"

func main() {
	learnConcurrency()
}

// c is a channel, a concurrency-safe communication object.
func inc(i int, c chan int) {
	c <- i + 1 // <- is the "send" operator when a channel appears on the left.
}

// We'll use inc to increment some numbers concurrently.
func learnConcurrency() {
	// Same make function used earlier to make a slice. Make allocates and
	// initializes slices, maps, and channels.
	c := make(chan int)
	// Start three concurrent goroutines. Numbers will be incremented
	// concurrently, perhaps in parallel if the machine is capable and
	// properly configured. All three send to the same channel.
	go inc(0, c) // go is a statement that starts a new goroutine.
	go inc(10, c)
	go inc(-805, c)
	// Read three results from the channel and print them out.
	// There is no telling in what order the results will arrive!
	fmt.Println(<-c, <-c, <-c) // channel on right, <- is "receive" operator.

	cs := make(chan string)       // Another channel, this one handles strings.
	ccs := make(chan chan string) // A channel of string channels.
	go func() { c <- 84 }()       // Start a new goroutine just to send a value.
	go func() { cs <- "wordy" }() // Again, for cs this time.
	// Select has syntax like a switch statement but each case involves
	// a channel operation. It selects a case at random out of the cases
	// that are ready to communicate.
	select {
	case i := <-c: // The value received can be assigned to a variable,
		fmt.Printf("it's a %T", i)
	case <-cs: // or the value received can be discarded.
		fmt.Println("it's a string")
	case <-ccs: // Empty channel, not ready for communication.
		fmt.Println("didn't happen.")
	}
	// At this point a value was taken from either c or cs. One of the two
	// goroutines started above has completed, the other will remain blocked.
	go func() { cs <- "test 3" }()
	go func() { cs <- "test 4" }()
	fmt.Println(<-cs)
	fmt.Println(<-cs)
}
