package config

// import (
// 	"log"
// 	"time"

// 	"gopkg.in/mgo.v2"
// )

// // connect establishes connection
// func connect() {

// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:    []string{MongoDBHosts},
// 		Timeout:  100 * 365 * 24 * time.Hour,
// 		Database: AuthDatabase,
// 		Username: AuthUserName,
// 		Password: AuthPassword,
// 		FailFast: true,
// 	}

// 	// session maintains a pool of socket connections
// 	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		log.Fatalf("CreateSession: %s\n", err)
// 	}

// }
