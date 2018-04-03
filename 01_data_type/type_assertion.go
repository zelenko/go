// A type assertion provides access to an interface value's underlying concrete value.
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)

	t := interface{}("test") // string as interface
	fmt.Println(t)

	s, ok = t.(string)
	fmt.Println(s, ok)
}
