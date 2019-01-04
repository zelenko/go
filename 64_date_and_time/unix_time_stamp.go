package main

import (
	"fmt"
	"time"
)

func main() {
	// today := time.Now()
	// sec := today.Unix()
	// fmt.Println("unix:", today.Unix())
	fmt.Println("unix:", time.Now().Unix())

	// oldDate := time.Unix(int64(1009807503), 0)
	// fmt.Println("old date:", oldDate.Format("2006-01-02_03-04-05pm MST")) // 2001-09-31_09-05-03am
	fmt.Println("now:", time.Unix(time.Now().Unix(), 0).Format("2006-01-02_03-04-05pm MST"))

	fmt.Println("old date:", time.Unix(int64(1009807503), 0).Format("2006-01-02_03-04-05pm MST"))
}
