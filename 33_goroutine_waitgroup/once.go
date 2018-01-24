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
	asking := make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("ask again", i)
			asking <- "ask again "
			once.Do(onceBody)
			done <- true

		}()
	}

	for i := 0; i < 10; i++ {
		<-done
		fmt.Println("check", i)
	}

	close(asking)

	for elem := range asking {
		//	for range asking {
		//fmt.Println(<- asking)
		fmt.Println(elem)
	}

}
