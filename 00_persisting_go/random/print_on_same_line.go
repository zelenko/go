package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 11; i++ {
		fmt.Print("\r", "line ", strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\r%s", strings.Repeat(" ", 35)) // clear all characters
	fmt.Print("\r", "line1")
	time.Sleep(5 * time.Second)
	fmt.Print("\r", "line 2")
}
