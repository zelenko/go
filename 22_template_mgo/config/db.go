package config

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
)

// database
var DB *mgo.Database
var OS *mgo.Database

// collections
var Books *mgo.Collection
var Products3 *mgo.Collection

func init() {
	// get a mongo sessions
	s, err := mgo.Dial(dblogin.Bookstore) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	OS = s.DB("onlinestore")
	Products3 = OS.C("products3")

	fmt.Println("You connected to your mongo database.")
}
