package main

import (
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 110; i++ {
		print(rand7(), " ")
	}
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
