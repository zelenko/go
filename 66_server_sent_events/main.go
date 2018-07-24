package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/antage/eventsource.v1"
)

func main() {
	es := eventsource.New(nil, nil)
	defer es.Close()
	http.Handle("/events", es)

	// send the message
	go func() {
		id := 1
		for {
			es.SendEventMessage("tick: "+strconv.Itoa(id), "", strconv.Itoa(id))
			id++
			time.Sleep(1 * time.Second)
		}
	}()

	// static content
	http.Handle("/", http.FileServer(http.Dir(".")))

	// run the server
	fmt.Println(http.ListenAndServe(":8080", nil))
}
