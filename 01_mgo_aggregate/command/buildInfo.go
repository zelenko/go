package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var result struct {
		Version        string
		VersionArray   []int  `bson:"versionArray"` // On MongoDB 2.0+; assembled from Version otherwise
		GitVersion     string `bson:"gitVersion"`
		OpenSSLVersion string `bson:"OpenSSLVersion"`
		SysInfo        string `bson:"sysInfo"` // Deprecated and empty on MongoDB 3.2+.
		Bits           int
		Debug          bool
		MaxObjectSize  int `bson:"maxBsonObjectSize"`
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	result, _ = session.BuildInfo()

	fmt.Println("Results All: ", result.Version)
	for _, v := range result.VersionArray {
		fmt.Println(v, "\t=")
	}

}
