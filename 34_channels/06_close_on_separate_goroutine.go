package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	messages := make(chan string)
	for x := 1; x <= 5; x++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()
			sayhello(messages, count)
		}(x)
	}

	// will not work:
	// wg.Wait()
	// close(messages)

	// But this works:
	go func() {
		wg.Wait()
		close(messages)
	}()

	for msg := range messages {
		fmt.Println(msg)
	}
}

func sayhello(messages chan<- string, count int) {
	time.Sleep(time.Millisecond * time.Duration(1000))
	messages <- fmt.Sprintf("hello: %d", count)
}
