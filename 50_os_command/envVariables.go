package main

import (
	"fmt"
	"os"
)

func main() {
	// Not working on Windows 10
	// shows one variable
	fmt.Println(os.Getenv("GOPATH"))

	// works on Windows 10
	// Shows all the Environmental variables
	list := os.Environ()
	for i := range list {
		fmt.Println(list[i])
	}
}
