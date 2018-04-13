package main

import (
	"fmt"
	"net/http"
)

// AnyType defines a type for the response
// type AnyTypeCopy struct{}
// type AnyType string
// type AnyType interface {}  // interface does not work
type AnyType int

// let that type implement the ServeHTTP method (defined in interface http.Handler)
func (h AnyType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Any type can say Hello!")
}

func main() {
	var h AnyType
	fmt.Println(http.ListenAndServe(":8080", h))
}
