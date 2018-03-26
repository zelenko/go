package main

import (
	"./webpage"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	log.Println("HTTPS port :443")
	log.Println("HTTP port :80")

	r := httprouter.New()
	r.GET("/", webpage.Secure)
	r.GET("/test/", webpage.Test)
	r.GET("/test", webpage.Test2)
	r.GET("/v", webpage.List)
	r.GET("/html", webpage.HTMLPage)
	r.NotFound = http.FileServer(http.Dir("public"))

	// r.ServeFiles("/static/*filepath", http.Dir("/var/www/public/"))
	// http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/home/www/"))))

	//  Start HTTP
	go func() {
		// redirect all HTTP to HTTPS
		err := http.ListenAndServe(":80", http.HandlerFunc(webpage.Redirect))
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
