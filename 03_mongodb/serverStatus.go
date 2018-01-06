package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {

	session, err := mgo.Dial(dblogin.Bookstore) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err)
	}
	defer session.Close()

	db := session.DB("test")

	//collectionName := "people"

	result := &bson.D{}
	err = db.Run(&bson.D{bson.DocElem{"serverStatus", 1}}, result)
	if err != nil {
		log.Fatalf("Failed to get collection stats: %s", err)
	}

	c := result.Map()["connections"]
	cu := result.Map()["current"]
	pid := result.Map()["pid"]

	fmt.Println("connections: ", c)
	fmt.Println("connections: ", cu)
	fmt.Println("connections pid: ", pid)

}

/*
		 "host" : "mongodb",
        "version" : "3.4.6",
        "process" : "mongod",
        "pid" : NumberLong(535),
        "uptime" : 4247807,
        "uptimeMillis" : NumberLong("4247807370"),
        "uptimeEstimate" : NumberLong(4247807),
        "localTime" : ISODate("2018-01-06T04:05:34.783Z"),
        "asserts" : {
                "regular" : 0,
                "warning" : 0,
                "msg" : 0,
                "user" : 113,
                "rollovers" : 0
        },
        "connections" : {
                "current" : 14,
                "available" : 805,
                "totalCreated" : 717
        },
        "extra_info" : {
                "note" : "fields vary by platform",
                "page_faults" : 515
*/
