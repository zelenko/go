// code samples using the sort package
package main

import (
	"fmt"
	"sort"
)

func main() {
	sl := []string{"mumbai", "london", "tokyo", "seattle"}
	//	sort.Strings(sl)
	sort.Sort(sort.Reverse(sort.StringSlice(sl)))
	//	sort.Sort(sort.StringSlice(sl))
	fmt.Println(sl)

	intSlice := []int{3, 5, 6, 4, 2, 293, -34}
	sort.Ints(intSlice)
	//sort.Sort(sort.IntSlice(intSlice))
	fmt.Println(intSlice)

	a := []int{5, 3, 4, 7, 8, 9}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, v := range a {
		fmt.Print(v, " ")
	}

	// sort slice
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

}
