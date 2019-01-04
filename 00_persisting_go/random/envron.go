package main

import (
	"fmt"
	"syscall"
)

func main() {
	list := syscall.Environ()
	for _, v := range list {
		fmt.Println(v)
	}
}
