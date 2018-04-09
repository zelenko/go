package main

import (
	"fmt"
	"unsafe"
)

func main() {
	copy((*(*[512]byte)(unsafe.Pointer(argp)))[:], data[:])
	fmt.Print(">")
}
