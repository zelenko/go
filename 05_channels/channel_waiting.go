package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	m := map[int]int{}
	m[4] = 7
	m[3] = 87
	m[72] = 19

	ch := make(chan int)

	// for closing ch
	ch2 := make(chan int)
	wg.Add(2)
	go func() {
		fmt.Println("start of func 1")
		var i int
		for n := range ch2 {
			i += n
			fmt.Println("ch2 n:", n, "i:", i)
			if i == len(m) {
				close(ch)
				fmt.Println("closing ch", len(m))
			}
		}
		fmt.Println("end of func 1")
		wg.Done()
	}()

	// populating
	go func() {
		fmt.Println("start of func 2")
		for _, v := range m {
			ch <- v
			ch2 <- 1
			fmt.Println("populating:", v)
			//fmt.Println("ch2 length", len(ch2))
		}
		fmt.Println("end of func 2")
		wg.Done()
	}()

	// printing ch
	for v := range ch {
		fmt.Println("ch:", v)
	}

	//for h := range ch2 {
	//	fmt.Println("h:",h)
	//}

	fmt.Println("closing ch2")
	// good housekeeping
	close(ch2)
	wg.Wait()
}
