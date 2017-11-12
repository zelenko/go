// This code selects one record from MongoDB and displays it in CLI

package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Pline string
	Bline string
}

func main() {
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("onlinestore").C("products3")
	//err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//           &Person{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//        log.Fatal(err)
	//}

	result := Person{}
	//out := bson.D()
	err = c.Find(bson.M{"_id": 69466}).One(&result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Phone:", result.Pline)
	fmt.Println("bline:", result.Bline)
	//fmt.Println("bline:", out.Bline)
	fmt.Println("bline:", bson.Now())

	// this examample is  from: https://labix.org/mgo
	// https://github.com/GoesToEleven/golang-web-dev/blob/master/046_mongodb/16_go-mongo/books/models.go  <== more examples here
}
