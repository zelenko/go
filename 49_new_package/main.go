package main

import (
	"fmt"

	"./toy"
)

func main() {
	car := toy.NewToy("Traxxas LaTrax Teton 1/18 4WD RTR RC Monster Truck")
	fmt.Println(car.GetName())
	fmt.Println(car.Length())

	truck := toy.NewToy([]byte(`one`))
	fmt.Println(truck.GetName())
	fmt.Println(truck.Length())

	truck1 := toy.NewToy(12345)
	fmt.Println(truck1.GetName())
	fmt.Println(truck1.Length())

	truck2 := toy.NewToy("12345kit", "boat23")
	fmt.Println(truck2.GetName())
	fmt.Println(truck2.Length())

	bikes, _ := toy.NewToys("bicycle", "motor", "boat", 345, []byte(`test-bike`))
	for i := range bikes {
		fmt.Print(bikes[i].GetName(), " ", bikes[i].Length(), "\n")
	}

	carts, err := toy.NewToys()
	if err != nil {
		//panic(err.Error())
		fmt.Println("Error:", err.Error())
		return
	}
	for i := range carts {
		fmt.Print(carts[i].GetName(), " ", carts[i].Length(), "\n")
	}

}
