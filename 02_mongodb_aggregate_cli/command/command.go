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

	// Run command, works
	result := bson.M{}
	if err := session.DB("admin").Run("serverStatus", &result); err != nil {
		check(err)
	} else {
		fmt.Println(result)
	}

	// Run a command, works
	var serverStatus = &bson.M{}
	if err := session.Run("serverStatus", serverStatus); err != nil {
		check(err)
	} else {
		fmt.Println(*serverStatus)
	}

	// Run a command with an argument, works
	var startupWarnings = &bson.M{}
	if err := session.Run(bson.D{{Name: "getLog", Value: "startupWarnings"}}, startupWarnings); err != nil {
		check(err)
	} else {
		fmt.Println(*startupWarnings)
	}

	if err := session.DB("test").Run("dbstats", &result); err != nil {
		check(err)
	} else {
		fmt.Println(result)
	}
	if err := session.DB("admin").Run("replSetGetStatus", &result); err != nil {
		check(err)
	} else {
		fmt.Println(result)
	}

}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		//http.Error(w, "Error: " + err.Error(), 500)
	}
}
