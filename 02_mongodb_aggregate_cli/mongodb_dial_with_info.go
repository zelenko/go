package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func main() {
	dialInfo, err := mgo.ParseURL("mongodb://localhost:27017")
	dialInfo.Direct = true
	dialInfo.FailFast = true
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
}
