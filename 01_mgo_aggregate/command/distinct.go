package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("onlinestore").C("products3")

	var result []string
	err = c.Find(bson.M{"category": bson.M{"$exists": 1}}).Distinct("category", &result)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Results All: ", result)
	for _, v := range result {
		fmt.Println(v, "\t=")
	}

}
