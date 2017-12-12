package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	b := a[:3]
	c := append(b, "lol")

	e := make([]string, len(c))
	copy(e, c)
	d := append(b, "woot")

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
