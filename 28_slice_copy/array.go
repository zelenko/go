/****************************************************************************************
Slice variable does not hold actual value, but it is a pointer to memory address.
Here is an example of how slice can get "corrupted".
This was mentioned in: https://www.upguard.com/blog/our-experience-with-golang
*****************************************************************************************/

package main

import "fmt"

func main() {

	array := []string{"a", "b", "c", "d", "e", "f"}
	slice1 := array[:3]
	slice2 := array[3:]
	fmt.Printf("so far so good: slice1 %v slice2 %v\n", slice1, slice2)
	slice1 = append(slice1, "BANG")

	fmt.Printf("append to slice1: %v\n", slice1)
	fmt.Printf("slice2 is now corrupt: %v\n", slice2)
	fmt.Printf("full array: %v\n", array)
	fmt.Printf("5th element of array: %v\n", array[4])

	for j, i := range array {
		fmt.Printf("%v ==> %v, \n", j, i)
	}

}
