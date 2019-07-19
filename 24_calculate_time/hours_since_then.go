package main

import "fmt"
import "time"

func main() {
	// Create new date.
	a := time.Date(2018, 4, 11, 0, 0, 0, 0, time.UTC)

	delta := time.Now().Sub(a)
	fmt.Println(delta.Hours())

	duration := time.Since(a)
	fmt.Println(duration.Hours())
}
