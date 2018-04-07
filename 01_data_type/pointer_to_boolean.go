package main

import (
	"fmt"
)

var auto = new(bool) // auto is the memory address

func main() {
	*auto = true
	fmt.Println(*auto) // display whats at the address

	*auto = false
	fmt.Println(*auto)

	on(auto)
	test(auto)

	off(auto)
	test(auto)
}

func test(in *bool) {
	fmt.Println(*in)
}

func off(in *bool) {
	*in = false // what ever is at that address, set it to false
}

func on(in *bool) {
	*in = true
}
