package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 110; i++ {
		print(randAB(1, 20), " ")
	}

	cryptRand()
}

func rand5() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(5)
	return num + 1
}

func rand7() int {
	//num := (rand5() + rand5() + rand5())%7
	num := 0
	for i := 0; i <= rand5(); i++ {
		num += rand5()
	}

	return num%7 + 1

}

func randAB(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	return a + rand.Intn(b-a+1) // a ≤ n ≤ b
}

func cryptRand() {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("b:", b, " | ", string(b))

}
