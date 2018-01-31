// value-type or a reference-type can satisfy an interface
// input interface, export specific types
// source: http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/
// Want to write to your own type from any io.Reader? Simply implement the Write(p []byte) (n int, err error) function and you can use io.Copy.

package main

import (
	"fmt"
)

// Both an Application value and a User reference satisfy the interface:
type (

	// Application datatype: struct
	Application struct {
		id string
	}

	// User type
	User struct {
		name string
	}

	// CacheItem interface type
	CacheItem interface {
		GetID() string
	}
)

// GetID returns string
func (a Application) GetID() string {
	return a.id
	// return `test1`
}

// GetID returns string
func (u *User) GetID() string {
	return u.name
	// return ``

}

// And thus, both (Application and User types) can be treated as a CacheItem.

// main is the entry point for the program.
func main() {

	Store(&User{name: "Leto"})
	Store(Application{id: "9001"})

}

// Store accepts interface
func Store(item CacheItem) {
	fmt.Println("item:", item.GetID())
}
