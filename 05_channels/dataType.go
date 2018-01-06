// Figure out what type it is: maps, slices, or arrays!

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Declaring local variables
	map1 := map[string]string{"name": "John", "desc": "Golang"}
	map2 := map[string]int{"apple": 23, "tomato": 13}
	slice1 := []int{1, 2, 3}
	array1 := [3]int{1, 2, 3}
	// var m map[string]int
	// m = make(map[string]int)
	// More info Here: https://blog.golang.org/go-maps-in-action

	// Type, such as map[string]string, []int, [3]int
	fmt.Println("map1:", reflect.TypeOf(map1))
	fmt.Println("map2:", reflect.TypeOf(map2))
	fmt.Println("slice1:", reflect.TypeOf(slice1))
	fmt.Println("array1:", reflect.TypeOf(array1))

	// Value, such as map, slice, array.
	fmt.Println("map1:", reflect.ValueOf(map1).Kind())
	fmt.Println("map2:", reflect.ValueOf(map2).Kind())
	fmt.Println("slice1:", reflect.ValueOf(slice1).Kind())
	fmt.Println("array1:", reflect.ValueOf(array1).Kind())

	// True/False statement inside Printf
	fmt.Printf("%v is a map? %v\n", map1, reflect.ValueOf(map1).Kind() == reflect.Map)
	fmt.Printf("%v is a map? %v\n", map2, reflect.ValueOf(map2).Kind() == reflect.Map)
	fmt.Printf("%v is a map? %v\n", slice1, reflect.ValueOf(slice1).Kind() == reflect.Map)

	/*  More about reflect package: https://golang.org/pkg/reflect/
		 	Invalid Kind = iota
	        Bool
	        Int
	        Int8
	        Int16
	        Int32
	        Int64
	        Uint
	        Uint8
	        Uint16
	        Uint32
	        Uint64
	        Uintptr
	        Float32
	        Float64
	        Complex64
	        Complex128
	        Array
	        Chan
	        Func
	        Interface
	        Map
	        Ptr
	        Slice
	        String
	        Struct
	        UnsafePointer
	*/

}
