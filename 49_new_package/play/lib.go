package play

import (
	"fmt"
	"strconv"
	"strings"
)

// Toy is the datatype
type Toy struct {
	Name string
}

// NewToy returns a new Toy type. It supports the following optional params:
func NewToy(v ...interface{}) *Toy {
	if len(v) < 1 || v[0] == nil {
		return &Toy{}
	}

	r := &Toy{Name: ""}

	// if multiple parameters, only the last one assigned
	for i := range v {
		switch v[i].(type) {
		case string:
			//r.Name = v[i].(string) // convert interface to string
			r = &Toy{Name: v[i].(string)}
		case []byte:
			r.Name = string(v[i].([]byte)) // convert interface to []byte
		case int:
			r.Name = strconv.Itoa(v[i].(int)) // convert interface to int
		}
	}

	return r
}

// Length returns the length as a string
func (r *Toy) Length() string {
	return strconv.Itoa(len(r.Name))
}

// String returns a string
func (r *Toy) String() string {
	return r.Name
}

// Write is to implement the Writer interface
func (r *Toy) Write(p []byte) (n int, err error) {
	r.Name = string(p) + "(w)"
	return len(p), nil
}

// Read is to implement the Reader interface
func (r *Toy) Read(p []byte) (n int, err error) {
	r.Name = string(p) + "(r)"
	return len(p), nil
}

// Reader returns a reader
func (r *Toy) Reader() *strings.Reader {
	return strings.NewReader(r.Name)
}

// NewToys returns a slice of Toy
func NewToys(args ...interface{}) ([]Toy, error) {
	manyToys := []Toy{}
	if len(args) < 1 || args[0] == nil {
		return nil, fmt.Errorf("Nothing or blank")
	}

	// if multiple parameters, return a slice of Toy
	for i := range args {
		toy1 := Toy{Name: ""}
		switch args[i].(type) {
		case string:
			//toy1.Name = args[i].(string) // convert interface to string
			toy1 = Toy{Name: args[i].(string)} // this way works too
		case []byte:
			toy1.Name = string(args[i].([]byte)) // convert interface to []byte
		case int:
			toy1.Name = strconv.Itoa(args[i].(int)) // convert interface to int
		}

		manyToys = append(manyToys, toy1)
	}

	return manyToys, nil
}
