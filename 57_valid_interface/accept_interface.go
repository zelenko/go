package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct{ name string }{name: "Go"}
	describe(strt)
	num := [...]byte{1, 2, 3}
	describe(num)
	slice := []byte{1, 2, 3}
	describe(slice)
	b := true
	describe(b)
	d := -1.2
	describe(d)
	p := &num
	describe(p)
}
