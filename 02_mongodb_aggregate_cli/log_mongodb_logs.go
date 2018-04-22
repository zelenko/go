package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	mgo.SetLogger(logger)
	mgo.SetDebug(true)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
}
