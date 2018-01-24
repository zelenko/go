package config

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
)

// DB bookstore database
var DB *mgo.Database

// OS onlinstore database
var OS *mgo.Database

// Books collection
var Books *mgo.Collection

// Products3 collection
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
