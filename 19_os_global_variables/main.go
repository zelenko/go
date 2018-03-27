package main

import (
	"fmt"
	"os"
)

func main() {
	// Works on Windows 10
	fmt.Println(os.Getenv("GOPATH"))

	// Works on Windows 10; Shows all the Environmental variables
	list := os.Environ()
	for i := range list {
		fmt.Println(list[i])
	}
}
