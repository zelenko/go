package config

import (
	"crypto/tls"
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"net"
)

// DB variables
var (
	DB    *mgo.Database
	Books *mgo.Collection
	s     *mgo.Session
)

func init() {
	dialInfo, err := mgo.ParseURL(dblogin.BookstoreM0)
	if err != nil {
		fmt.Println("Cannot parse mongodb URL: " + err.Error())
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("web")
	Books = DB.C("books")

	fmt.Println("You connected to your mongo database.")
}

// CloseSession closes mongodb session
func CloseSession() {
	s.Close()
}
