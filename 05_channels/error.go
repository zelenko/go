package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Go errors:
	result, err := SomeFunction("b")
	checkErr(err)
	fmt.Println("result: ", result)
}

func SomeFunction(a string) (b int, err error) {

	b, err = strconv.Atoi(a)

	return
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("My Error: ", err.Error())
		//http.Error(w, "Error: " + err.Error(), 500)
		return
	}
}
