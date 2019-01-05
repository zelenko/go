// Using io.Writer to write to underlying slice
// fmt.Fprint is printing to a writer
package main

import (
	"fmt"
)

type pipe struct {
	name string
}

type pList struct {
	items []pipe
}

// list is a pointer to pList type
var list = &pList{}

// Write method for pList type to implement io.Writer
func (l *pList) Write(p []byte) (n int, err error) {
	l.items = append(l.items, pipe{name: string(p)})
	return 0, nil
}

// String method for pList type to implement stringer interface
func (l *pList) String() string {
	out := ""
	for _, i := range l.items {
		out = out + i.name + ", "
	}
	return out
}

// main is the entry point for the program.
func main() {
	list = &pList{items: []pipe{{name: "one"}, {name: "two"}}}
	fmt.Println(list)

	fmt.Fprint(list, "three")
	fmt.Fprint(list, "four")
	fmt.Println(list)
}
