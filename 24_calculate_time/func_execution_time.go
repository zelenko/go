package main

import (
	"log"
	"time"
)

func main() {
	// Add these two lines to show function execution time
	stop := StartTimer("exampleFunc")
	defer stop()

	time.Sleep(1 * time.Second)
}

// StartTimer measures time it took to run function
func StartTimer(name string) func() {
	startTime := time.Now()
	log.Println(name, "started")
	return func() {
		duration := time.Now().Sub(startTime)
		log.Println(name, "took", duration)

	}
}
