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

	collectionName := "people"

	result := &bson.D{}
	err = db.Run(&bson.D{bson.DocElem{"collstats", collectionName}}, result)
	if err != nil {
		log.Fatalf("Failed to get collection stats: %s", err)
	}

	storageSize := result.Map()["storageSize"]

	size := result.Map()["size"]
	count := result.Map()["count"]
	avgObjSize := result.Map()["avgObjSize"]

	fmt.Println("Storage === Size: ", size)
	fmt.Println("Storage == count: ", count)
	fmt.Println("Storage StorageS: ", storageSize)
	fmt.Println("Stora avgObjSize: ", avgObjSize)

}

/*
		"ns" : "onlinestore.traffic",
        "size" : 166377,
        "count" : 638,
        "avgObjSize" : 260,
        "storageSize" : 77824,
*/
