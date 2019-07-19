package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("unix:", time.Now().Unix())
	fmt.Println("now:", time.Unix(time.Now().Unix(), 0).Format("2006-01-02_03-04-05pm MST"))      // 2001-09-31_09-05-03am
	fmt.Println("old date:", time.Unix(int64(1009807503), 0).Format("2006-01-02_03-04-05pm MST")) // 2001-09-31_09-05-03am
}
