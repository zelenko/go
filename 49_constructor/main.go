package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"./play"
)

func main() {
	car := play.NewToy("Traxxas LaTrax Teton 1/18 4WD RTR RC Monster Truck")
	fmt.Println(car.String())
	fmt.Println(car.Length())

	truck := play.NewToy([]byte(`one`))
	fmt.Println(truck.String())
	fmt.Println(truck.Length())

	truck1 := play.NewToy(12345)
	fmt.Println(truck1.String())
	fmt.Println(truck1.Length())

	truck2 := play.NewToy("12345kit", "boat23")
	fmt.Println(truck2.String())
	fmt.Println(truck2.Length())

	carts, err := play.NewToys()
	if err != nil {
		//panic("Error: " + err.Error())
		//log.Fatalln("Error:", err.Error())
		fmt.Println("Error:", err.Error()) // Display error message and continue control flow.
	}
	for i := range carts { // If "carts" is nil, then this loop will not run
		fmt.Print(i, carts[i].String(), " ", carts[i].Length(), "\n")
	}

	bikes, _ := play.NewToys("bicycle", "motor", "boat", 345, []byte(`test-bike`))
	for i := range bikes {
		fmt.Print(i, "-", bikes[i].String(), " ", bikes[i].Length(), "\n")
	}

	// Experimenting with Readers and Writers

	//drone := &play.Toy{}		   // both allocaions work
	drone := play.NewToy()         // both allocaions work
	fmt.Fprint(drone, "Drone 456") // writer (Toy has writer interface, so you can write to it).
	fmt.Println("name:", drone)

	//tesla := &play.Toy{}
	tesla := play.NewToy()
	tesla.Read([]byte("Tesla 12345676")) // tesla reader
	fmt.Println("name:", tesla)

	n, err := io.Copy(tesla, strings.NewReader("New Tesla")) // tesla writer and new reader
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println(n, "Tesla name:", tesla)

	io.WriteString(tesla, "Hello Tesla") // tesla writer
	fmt.Println("Tesla H:", tesla)

	// writer
	io.Copy(drone, bytes.NewReader([]byte("new drone")))
	fmt.Println("drone n:", drone)

	// writer
	_, err = io.Copy(drone, tesla.Reader()) // drone writer and tesla reader
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println("drone t:", drone)

	iot := bytes.NewReader([]byte("Internet of Things")) // creates a reader
	iot.WriteTo(drone)                                   // iot reader writes to drone writer
	fmt.Print("drone: I", drone)

}
