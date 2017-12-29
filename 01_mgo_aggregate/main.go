// This code selects aggregate data from mongodb and displays it in CLI

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

	/*
		// Same command for use in mongodb Shell
		db.products.aggregate([
			{$group : {_id:"$category", count:{$sum:1}}},
			{$sort:{"count":-1}},
			{$limit: 15}
		  ])
	*/

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":   "$category",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{
				"count": -1,
			},
		},
		{
			"$limit": 15,
		},
	}

	pipe := c.Pipe(pipeline)

	//pipe := c.Pipe([]bson.M{{"$match": bson.M{"user":"John"}}})
	//resp := []bson.M{}

	// Declaring type
	resp := []struct {
		ID    string `json:"id,omitempty" bson:"_id"`
		Count int    `json:"count,omitempty" bson:"count"`
	}{{}}

	err = pipe.All(&resp)
	checkErr(err)

	//fmt.Println(resp) // simple print proving it's working

	for _, v := range resp {
		fmt.Println(v.Count, "\t", v.ID)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
