// "select case" selects only one, whichever is faster.
package main

import "time"
import "fmt"

func main() {

    // executing an external call that returns its result on a channel `c1`
    // after 2s.
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()
	
	// which one is faster?
	// 1 sec is faster than 2 sec, so "timeout 1" is printed
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }



	// Here is another example: 2 sec vs 3 sec

    // If we allow a longer timeout of 3s, then the receive
    // from `c2` will succeed and we'll print the result.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}