package main

import (
	"fmt"
	"sync"
)

var completion sync.WaitGroup

func main() {
	fmt.Println("start")

	completion.Add(1) // wait for one thing
	go doSomething()

	fmt.Println("end")
	completion.Wait() // wait for all things to be done
}

func doSomething() {
	fmt.Println("do something")
	completion.Done() // this is done
}
