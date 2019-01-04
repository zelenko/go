package main

import "fmt"

// Generator returns a channel that produces the numbers 1, 2, 3,…
// To stop the underlying goroutine, close the channel.
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number)
	// …
}

// https://programming.guide/go/stop-goroutine.html
