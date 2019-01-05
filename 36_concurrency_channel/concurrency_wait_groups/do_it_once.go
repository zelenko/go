// Run the function once, and never again.
// Even if asked multiple times, do NOT run more than once.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println("ask again", j)
			once.Do(onceBody)
			done <- true

		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
		fmt.Println("check", i)
	}
}
