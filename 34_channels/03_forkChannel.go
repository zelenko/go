package main

import (
	"fmt"
	"sync"
)

func main() {
	c := increment()
	c3 := make(chan int)   // send only channel
	c2 := duplicate(c, c3) // WORKING
	//duplicate3(c, c3) // works
	//c2 := duplicate4(c) // works
	//c2 := duplicateWorking(c, c3) // works

	cSum := pull(c2)
	for n := range cSum {
		fmt.Println("total:", n)
	}
	for n := range c3 {
		fmt.Println("copy:", n)
	}
}

// working
func duplicate(c <-chan int, output chan<- int) <-chan int {
	out1 := make(chan int)
	go func() {
		wg := new(sync.WaitGroup) // 1
		for n := range c {
			wg.Add(1) // 2
			go func(count int, c1 chan int, c2 chan<- int) {
				defer wg.Done() // 3
				c1 <- count
				c2 <- count
			}(n, out1, output)
			fmt.Println("-", n)
		}
		close(out1) // works, order is important
		wg.Wait()   // 4
		close(output)
	}()
	return out1
}

// working
func duplicate6(c <-chan int, output chan<- int) <-chan int {
	out1 := make(chan int)
	go func() {
		wg := new(sync.WaitGroup) // 1
		for n := range c {
			temp := n

			//			out1 <- temp
			//output <- temp
			wg.Add(1) // 2
			go func(count int, c chan int, c2 chan<- int) {
				defer wg.Done() // 3
				c <- count
				c2 <- count
			}(temp, out1, output)

			// wg.Add(1) // 2
			// go func(count int, c chan int) {
			// 	defer wg.Done() // 3
			// 	c <- count
			// }(temp, out1)

			// wg.Add(1) // 2
			// go func(count int, c chan<- int) {
			// 	defer wg.Done() // 3
			// 	output <- count
			// }(temp, output)

			fmt.Println("-", temp)
		}

		//close(output)
		close(out1) // works, order is important
		wg.Wait()   // 4
		//close(out1) // will not work in this order
		close(output)
		//close(out1) // will not work in this order
	}()
	return out1
}

// working
func duplicate4(c <-chan int) <-chan int {
	out1 := make(chan int)
	go func() {
		wg := new(sync.WaitGroup) // 1
		for n := range c {
			temp := n
			//output <- temp
			wg.Add(1) // 2
			go func(count int, c chan int) {
				defer wg.Done() // 3
				c <- count
			}(temp, out1)
			fmt.Println("-", temp)
		}
		//close(output)
		wg.Wait() // 4
		close(out1)
	}()
	return out1
}

// works
func duplicate3(c <-chan int, output chan<- int) {
	//out1 := make(chan int)
	go func() {
		//wg := new(sync.WaitGroup) // 1
		for n := range c {
			temp := n
			output <- temp
			// wg.Add(1) // 2
			// go func(count int) {
			// 	defer wg.Done() // 3
			// 	out1 <- count
			// }(temp)
			fmt.Println("-", temp)
		}
		close(output)
		//wg.Wait() // 4
		//close(out1)
	}()
	//return out1
}

// working
func duplicate5(c <-chan int, output chan<- int) <-chan int {
	out1 := make(chan int)
	go func() {
		wg := new(sync.WaitGroup) // 1
		for n := range c {
			temp := n
			//output <- temp
			wg.Add(1) // 2
			go func(count int, c chan<- int) {
				defer wg.Done() // 3
				output <- count
			}(temp, output)

			//out1 <- temp
			wg.Add(1) // 2
			go func(count int, c chan int) {
				defer wg.Done() // 3
				c <- count
			}(temp, out1)

			fmt.Println("-", temp)
		}

		close(out1) // works, order is important
		wg.Wait()   // 4
		//close(out1)  // will not work in this order
		close(output)
		//close(out1) // will not work in this order
	}()
	return out1
}

// working
func duplicateWorking(c <-chan int, output chan<- int) <-chan int {
	out1 := make(chan int)
	go func() {
		wg := new(sync.WaitGroup) // 1
		for n := range c {
			temp := n
			out1 <- temp
			wg.Add(1) // 2
			go func(count int) {
				defer wg.Done() // 3
				output <- count
			}(temp)
			fmt.Println("-", temp)
		}
		close(out1)
		wg.Wait() // 4
		close(output)
	}()
	return out1
}

func increment() <-chan int { //receive only channel
	out := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func pull(c <-chan int) <-chan int { // returns receive-only type <-chan, receive FROM  channel
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
