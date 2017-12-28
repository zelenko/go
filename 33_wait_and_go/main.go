package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	wg.Add(1) // wait for one thing
	go doSomething()
	fmt.Println("end")
	wg.Wait() // wait for all things to be done
	// end of program
}

func doSomething() {
	fmt.Println("do something")
	wg.Done() // this is done
}
