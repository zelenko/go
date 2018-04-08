// Copyright © 2018 Zelenko. All rights reserved.

package main

import (
	"./sess"
	"fmt"
)

// main is the entry point for the program.
func main() {
	mgr := sess.NewManager()
	defer mgr.Close()

	mgr.RegisterUser("user8", "userPass5")
	mgr.RegisterUser("user10", "userPass61")

	for _, i := range mgr.ShowUsers() {
		fmt.Println(i)
	}

	fmt.Println("> " + mgr.RandomSession())
	fmt.Println("> " + mgr.RandomSession())
	fmt.Println("> " + mgr.RandomSession())
	fmt.Println("> " + mgr.RandomSession())
	fmt.Println("> " + mgr.RandomSession())
	fmt.Println("> " + mgr.RandomSession())
	fmt.Println("> " + mgr.RandomSession())
}
