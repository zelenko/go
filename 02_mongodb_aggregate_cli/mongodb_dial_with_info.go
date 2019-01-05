package main

import (
	"gopkg.in/mgo.v2"
)

func main() {
	dialInfo, err := mgo.ParseURL("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	dialInfo.Direct = true
	dialInfo.FailFast = true
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
}
