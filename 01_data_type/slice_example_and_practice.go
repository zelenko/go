// practicing with slices, struct, loop, switch, and randoms
// this is just a test
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {

	// creating a slice of users
	user := []struct {
		Name, City string
		level      int
	}{
		{"John", "Nashville", 45},
		{"Joe", "Orlando", 34},
	}

	for _, i := range user {
		fmt.Println(i.Name + " is in " + i.City)
	}

	for i := 0; i < len(user); i++ {
		fmt.Printf("%s is in %s\n", user[i].Name, user[i].City)
	}

	// u is a type
	type u struct {
		Name, City string
		level      int
	}

	// declaring and initializing
	user1 := &u{"Eric", "Austin", 14}

	user = append(user, *user1)

	for i := 0; i < len(user); i++ {
		fmt.Printf("%s is in %s\n", user[i].Name, user[i].City)
	}

	// generate random list.  c must be outside of loop to access it later.
	c := ""
	for i := 0; i < 20; i++ {
		a := strconv.Itoa(rand.Intn(3))
		b := []byte(a)

		for _, i := range b {
			c += string(i) + ", "
		}
	}

	fmt.Println(c)

	err := getInterface(12)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	fmt.Println(useInterface("test"))
	fmt.Println(useInterface([]byte("bytes")))
	fmt.Println(useInterface(123456))
	fmt.Println(useInterface(byte('a')))
	fmt.Println(useInterface([]string{"test", "two", "three", "four"}))
	fmt.Println(useInterface(1.2))

}

// getInterface prints a string or returns an error if something other than string given
func getInterface(b interface{}) error {
	a, ok := b.(string)
	if !ok {
		return fmt.Errorf("this is not a string")
	}

	fmt.Println("This is a string: " + a)
	return nil
}

// userInterface can accept these types are: string, int, byte, []byte, []string
func useInterface(a interface{}) (string, error) {
	s := ""
	switch a.(type) {
	case string:
		s = a.(string)
	case int:
		s = strconv.Itoa(a.(int))
	case byte:
		s = string(a.(byte))
	case []byte:
		s = string(a.([]byte))
	case []string:
		s = strings.Join(a.([]string), " ")
	default:
		return s, fmt.Errorf("invalid type;  valid types are: string, int, byte, []byte, []string")
	}
	return s, nil
}
