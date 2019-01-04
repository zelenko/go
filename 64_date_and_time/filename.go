package main

import (
	"fmt"
	"time"
)

func main() {
	filename := time.Now().Format("2006-01-02_03-04-05pm") + ".csv"
	fmt.Printf(filename)
	filename = time.Now().Format("2006-01-02_3pm") + ".csv"
	fmt.Printf(filename)
}
