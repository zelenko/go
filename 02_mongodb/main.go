package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"      // same as ==> labix.org/v2/mgo
	"gopkg.in/mgo.v2/bson" // sames as ==> labix.org/v2/mgo/bson
	"time"
	"strconv"
)

type person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

var (
	//IsDrop = true
	IsDrop = false
)

func main() {
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	c := session.DB("test").C("people")

	// Index
	index := mgo.Index{
		Key:        []string{"name", "phone"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		//panic(err)
		fmt.Println("LINE 55", err)
	}

	// Insert Datas
	err = c.Insert(&person{Name: "Alex1", Phone: "+55 88 1234 4321", Timestamp: time.Now()},
		&person{Name: "Clara1", Phone: "+66 88 1234 5678", Timestamp: time.Now()})

	if err != nil {
		//panic(err)
		fmt.Println("LINE 63", err)
	}

	// Query One
	result := person{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone", result)

	// Query All
	var results []person
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		//panic(err)
		fmt.Println("LINE 89", err)
	}

	// Query All
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	fmt.Printf("%T\n", results)

	for k, v := range results {
		fmt.Printf("%s -> %s\n", strconv.Itoa(k), v)
	}

}
