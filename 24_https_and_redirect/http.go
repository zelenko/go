package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP port :80")

	n := httprouter.New()
	n.GET("/", notSecure)

	//  Start HTTP
	err := http.ListenAndServe(":80", n)
	if err != nil {
		log.Fatal("Web server (HTTP): ", err)
	}
}

// notSecure is for HTTP
func notSecure(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Transport layer is NOT secure.\n"))
}
