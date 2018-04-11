package main

import (
	"fmt"
	"unsafe"
)

// main is the entry point for the program.
func main() {
	copy((*(*[512]byte)(unsafe.Pointer(argp)))[:], data[:])
	fmt.Print(">")
}
