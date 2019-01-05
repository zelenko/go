// wait for goroutine to finish.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 1

func main() {
	fmt.Println("start")

	wg.Add(1) // 2
	go doSomething()

	wg.Wait() // 4
	fmt.Println("end")
}

func doSomething() {
	fmt.Println("do something")
	wg.Done() // 3
}
