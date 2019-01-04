package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func pinger(c chan string, id int) {
	i := 1
	for {
		c <- strconv.Itoa(id) + " ping\t" + strconv.Itoa(i) + "\t" + bson.NewObjectId().Hex() + "\t" + strconv.FormatInt(time.Now().UnixNano(), 10)
		i++
		// if i > 30 {
		// 	break
		// }
		time.Sleep(time.Millisecond * time.Duration(4))
	}
	//close(c)
}

func main() {
	var c = make(chan string)
	for i := 1; i < 21; i++ {
		go pinger(c, i)
	}

	i := 0
	for {
		i++
		msg, opened := <-c
		if !opened {
			break
		}
		fmt.Println(msg)
	}
}
