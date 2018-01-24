//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
// Generator that counts to n //
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
package main

import (
	"fmt"
)

func main() {
	for i := range count(10) {
		fmt.Println("Counted", i)
	}
}

// count will create a channel
func count(n int) chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}
