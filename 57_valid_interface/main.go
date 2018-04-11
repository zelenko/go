package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Valid interface has OK method
type Valid interface {
	OK() error
}

// Person has name
type Person struct {
	Name string
}

// OK method returns nil or error
func (p Person) OK() error {
	if p.Name == "" {
		return errors.New("name required")
	}
	return nil
}

// Decode interface takes reader and converts
func Decode(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return errors.New("decoding error: " + err.Error())
	}
	obj, ok := v.(Valid) // cast as type "Valid" interface
	// obj, ok := v.(Person)
	if !ok {
		//return nil // no OK method
		return errors.New("cannot convert 'v' to interface 'Valid'")
	}
	err = obj.OK()
	if err != nil {
		return errors.New("invalid interface: " + err.Error())
	}
	return nil
}

// main is the entry point for the program.
func main() {
	json := `{"Name": "John Doe"}`
	//user := &struct{ Name string }{}
	user := &Person{}
	err := Decode(strings.NewReader(json), user)
	if err != nil {
		fmt.Println("error:", err)
		goto end
	}
	fmt.Println("name:", user.Name)

end:
}
