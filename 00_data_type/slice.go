package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	for _, num := range s {
		fmt.Printf("Output: %s\n", num)
	}
	fmt.Println("len:", len(s))

}
