package main

import (
	"fmt"
)

var auto bool

func main() {
	auto = true
	fmt.Println(auto)

	off(&auto) // The `&auto` syntax gives the memory address of `auto`,
	fmt.Println(auto)

	on(&auto)
	test()

	off(&auto)
	test2()
}

func off(in *bool) {
	*in = false // *in dereferences the pointer from its
	// memory address to the current value at that address.
}

func on(in *bool) {
	*in = true
}

func test() {
	fmt.Println(auto)
}

func test2() {
	fmt.Println(auto)
}
