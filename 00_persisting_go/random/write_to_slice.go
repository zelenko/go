// My example of how type can have Writer, or Stringer interface.
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

var list = &pList{}

func (l *pList) Write(p []byte) (n int, err error) {
	l.items = append(l.items, pipe{name: string(p)})
	return 0, nil
}

func (l *pList) String() string {
	out := ""
	for _, i := range l.items{
		out = out + i.name + ", "
	}
	return out
}

func main() {
	list = &pList{items: []pipe{pipe{name: "one"}, pipe{name: "two"}}}
	fmt.Println(list)

	fmt.Fprint(list, "three")
	fmt.Fprint(list, "four")
	fmt.Println(list)
}
