package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

func main() {

	// Create a mgo session
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	// Close the session when the function ends
	defer session.Close()

	fmt.Println("You connected to your mongo database.")

	// Run a command
	var serverStatus = &bson.M{}
	if err := session.Run("serverStatus", serverStatus); err != nil {
		panic(err)
	} else {
		//fmt.Println(*serverStatus)
		printMap(serverStatus)
	}

	// Run a command with an argument
	var startupWarnings = &bson.M{}
	if err := session.Run(bson.D{{"getLog", "startupWarnings"}}, startupWarnings); err != nil {
		panic(err)
	} else {
		//fmt.Println(*startupWarnings)
		//printMap(startupWarnings)
	}

}

func printMap(input *bson.M) {

	for k, v := range *input {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			// print map
			printMap(*v)
			fmt.Println("it is a map")
		} else {
			fmt.Println("k:", k, "v:", v)
		}

	}
}
