package main

import (
	"fmt"
)

func main() {
	ok := []int{1,2,3,4,5}
	//for i := 0; i < len(ok); i++ {
	
	for n := range ok {
		fmt.Print(" " , ok[n])
	}
	fmt.Println("")
	
	for i := 0; i < len(ok); i++ {
		fmt.Print(" " , ok[i])
	}
	fmt.Println("")
	
	for i, n := range ok {
		fmt.Print(" " , ok[i],n)
	}
}
