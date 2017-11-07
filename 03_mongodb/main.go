package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

func main() {
	p := fmt.Println

	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Collection People
	c := session.DB("test").C("people")
	var results []Person

	// Query All
	err = c.Find(bson.M{}).Sort("-timestamp").All(&results)
	if err != nil {
		panic(err)
	}

	for _, v := range results {
		//fmt.Printf("%s -> %s\n", k, v)
		//t1, e := time.Parse(time.RFC3339,			"2012-08-11T22:08:41+00:00")
		p(v.Phone, "\t", v.Timestamp.Format("2006-01-02 3:04PM"), "\t", v.Name, "\t")
	}
	fmt.Printf("Total results: %d\n", len(results))
}
