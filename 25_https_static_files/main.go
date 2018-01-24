package main

import (
	"../25_public_html/lib"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTPS port :443")
	fmt.Println("HTTP port :80")

	r := httprouter.New()
	r.GET("/", zelenko.Secure)
	r.GET("/test/", zelenko.Test)
	r.GET("/test", zelenko.Test2)
	r.NotFound = http.FileServer(http.Dir("public"))
	// r.ServeFiles("/static/*filepath", http.Dir("/var/www/public/"))
	//http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/home/www/"))))

	//  Start HTTP
	go func() {
		err := http.ListenAndServe(":80", http.HandlerFunc(zelenko.Redirect))
		if err != nil {
			log.Fatalln("Web server (HTTP): ", err)
		}
	}()

	//  Start HTTPS
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", r)
	if err != nil {
		log.Fatal("Web server (HTTPS): ", err)
	}
}
