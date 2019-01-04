package main

import (
	"fmt"
	"sync"
)

func main() {
	c := increment()
	fmt.Println("len c:", len(c))

	c3 := make(chan<- int) // send only channel

	c2 := duplicate(c, c3)

	fmt.Println("len c2:", len(c2))
	//fmt.Println("len c3:", len(c3))

	// for i := range c3 {
	// 	fmt.Println(">", i)
	// }

	cSum := pull(c2)
	for n := range cSum {
		fmt.Println("total:", n)
	}
	// for n := range c3 {
	// 	fmt.Println("copy:", n)
	// }

	// copy1 := <-c3
	// fmt.Println("copy:", copy1)
}
func increment() <-chan int { //receive only channel
	//wg := sync.WaitGroup
	out := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			out <- i
			//wg.Add()
		}
		close(out)
	}()

	return out
}
func pull(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func duplicate2(c <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		for n := range c {
			temp := n
			out1 <- temp
			out2 <- temp
		}
		close(out1)
		close(out2)
	}()
	return out1, out2
}

func duplicate(c <-chan int, output chan<- int) <-chan int { // chan<- is receiving; <-chan is sending
	out1 := make(chan int)
	//out2 := make(chan int)
	wg := sync.WaitGroup{}
	fmt.Println("len c:", len(c))

	send := func(in int) {
		go func(number int) {
			//out2 <- number
			fmt.Println("-", number)
			wg.Done()
		}(in)
	}
	go func() {
		//wg.Add(len(c))
		for n := range c {
			temp := n
			out1 <- temp
			//fmt.Println("=>", temp)
			wg.Add(1)
			send(temp)
			//output <- temp
			//wg.Done()
		}
		close(out1)
		//close(out2)
	}()
	wg.Wait()
	// go func() {
	// 	out2 <- 22
	// 	close(out2)

	// }()

	fmt.Println("len out1:", len(out1))
	//fmt.Println("len out2:", len(out2))
	return out1
}
