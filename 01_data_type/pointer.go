package main

import "fmt"

// Person is data type
type Person struct {
	firstName string
	lastName  string
}

func changeName(p *Person) {
	p.firstName = "Bob"
}

// main is the entry point for the program.
func main() {
	person := Person{
		firstName: "Alice",
		lastName:  "Dow",
	}

	changeName(&person)

	fmt.Println(person)
}
