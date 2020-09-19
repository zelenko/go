package main

import "fmt"

func main() {
	run([]string{"1", "2", "3"})
	run("ok")
	run()
}

// run function accepts any parameter type
func run(i ...interface{}) {
	if i == nil {
		fmt.Println("input is null")
		return
	}

	for _, input := range i {
		switch input.(type) {
		case int:
			fmt.Println("integer:", input)
		case string:
			fmt.Println("string:", input)
		case bool:
			fmt.Println("bool:", input)
		case []string:
			fmt.Println("slice of string:", input)
		case []int:
			fmt.Println("slice of int:", input)
		default:
			fmt.Println("some other type:", input)
		}
	}
}
