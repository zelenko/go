package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1 * time.Millisecond)

	println(strconv.FormatInt(time.Since(start).Nanoseconds()/int64(time.Millisecond), 10))
	println(strconv.FormatInt(time.Since(start).Nanoseconds()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10))
	println(time.Since(start))
	println(time.Since(start).Seconds())
	fmt.Printf("Seconds [%v]", time.Since(start))

	diff := now.Sub(start)
	second := int(diff.Seconds())
	fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)
}
